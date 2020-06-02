package implementation

import (
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/mappers"
	"github.com/bixlabs/authentication/tools"
	"github.com/sirupsen/logrus"
)

type userManager struct {
	authenticator interactors.Authenticator
	repository    user.Repository
}

func NewUserManager(auth interactors.Authenticator, repo user.Repository) interactors.UserManager {
	return &userManager{authenticator: auth, repository: repo}
}

func (um userManager) Create(user structures.User) (structures.User, error) {
	loggerFields := logrus.Fields{"email": user.Email, "meth": "userManager:Create"}
	contextLogger := tools.Log().WithFields(loggerFields)

	if user, err := um.generatePasswordIfEmpty(&user); err != nil {
		contextLogger.WithError(err).Debug("error creating user with empty password")

		return user, err
	}

	return um.authenticator.Signup(user)
}

func (um userManager) generatePasswordIfEmpty(user *structures.User) (structures.User, error) {
	if user.Password != "" {
		return *user, nil
	}

	var err error
	user.GeneratedPassword, err = util.GenerateRandomPassword()
	if err != nil {
		return *user, err
	}

	user.Password = user.GeneratedPassword

	return *user, nil
}

func (um userManager) Delete(email string) error {
	loggerFields := logrus.Fields{"email": email, "meth": "userManager:Delete"}
	contextLogger := tools.Log().WithFields(loggerFields)

	if err := util.IsValidEmail(email); err != nil {
		contextLogger.WithError(err).Debug("email not valid")

		return err
	}

	userToRemove, err := um.repository.Find(email)
	if err != nil {
		contextLogger.Debug("email not found in repository")

		return util.UserNotFoundError{}
	}

	return um.repository.Delete(userToRemove)
}

func (um userManager) Find(email string) (structures.User, error) {
	loggerFields := logrus.Fields{"email": email, "method": "userManager:Find"}
	contextLogger := tools.Log().WithFields(loggerFields)

	if err := util.IsValidEmail(email); err != nil {
		contextLogger.WithError(err).Debug("email not valid")

		return structures.User{}, err
	}

	user, err := um.repository.Find(email)
	if err != nil {
		contextLogger.WithError(err).Debug("email not found in repository")

		return structures.User{}, util.UserNotFoundError{}
	}

	return user, nil
}

func (um userManager) Update(email string, updateAttrs structures.UpdateUser) (structures.User, error) {
	loggerFields := logrus.Fields{"email": email, "method": "userManager:Update"}
	contextLogger := tools.Log().WithFields(loggerFields)

	if err := util.IsValidEmail(email); err != nil {
		contextLogger.WithError(err).Debug("email is not valid")

		return structures.User{}, err
	}

	if updateAttrs.Email != "" {
		if err := util.IsValidEmail(updateAttrs.Email); err != nil {
			contextLogger.WithError(err).Debug("updated email is not valid")

			return structures.User{}, err
		}
	}

	if updateAttrs.Password != "" {
		if err := util.CheckPasswordLength(updateAttrs.Password); err != nil {
			contextLogger.WithError(err).Debug("updated password doesnt follow policy")

			return structures.User{}, err
		}
	}

	user, err := um.repository.Find(email)
	if err != nil {
		return user, util.UserNotFoundError{}
	}

	if updateAttrs.Email != user.Email {
		if isAvailable, err := um.repository.IsEmailAvailable(updateAttrs.Email); err != nil || !isAvailable {
			contextLogger.WithError(err).Debug("A duplicated email was provided")

			return user, util.DuplicatedEmailError{}
		}
	}

	if updateAttrs.Password != "" {
		hashedPassword, err := util.HashPassword(updateAttrs.Password)
		if err != nil {
			contextLogger.WithError(err).Debug("error hashing the updated password")

			return user, err
		}
		user.Password = hashedPassword
	}

	return um.repository.Update(user.Email, mappers.AssignUserUpdateToUser(user, updateAttrs))
}
