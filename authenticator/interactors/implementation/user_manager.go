package implementation

import (
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/mappers"
	"github.com/bixlabs/authentication/tools"
)

type userManager struct {
	authenticator interactors.Authenticator
	repository    user.Repository
}

// NewUserManager is a constructor for the userManager struct
func NewUserManager(auth interactors.Authenticator, repo user.Repository) interactors.UserManager {
	return &userManager{authenticator: auth, repository: repo}
}

func (um userManager) Create(user structures.User) (structures.User, error) {
	if user, err := um.generatePasswordIfEmpty(&user); err != nil {
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
	if err := util.IsValidEmail(email); err != nil {
		return err
	}

	userToRemove, err := um.repository.Find(email)
	if err != nil {
		return util.UserNotFoundError{}
	}

	return um.repository.Delete(userToRemove)
}

func (um userManager) Find(email string) (structures.User, error) {
	if err := util.IsValidEmail(email); err != nil {
		return structures.User{}, err
	}

	user, err := um.repository.Find(email)
	if err != nil {
		return structures.User{}, util.UserNotFoundError{}
	}

	return user, nil
}

func (um userManager) Update(email string, updateAttrs structures.UpdateUser) (structures.User, error) {
	if err := util.IsValidEmail(email); err != nil {
		return structures.User{}, err
	}

	if updateAttrs.Email != "" {
		if err := util.IsValidEmail(updateAttrs.Email); err != nil {
			return structures.User{}, err
		}
	}

	if updateAttrs.Password != "" {
		if err := util.CheckPasswordLength(updateAttrs.Password); err != nil {
			return structures.User{}, err
		}
	}

	user, err := um.repository.Find(email)
	if err != nil {
		return user, util.UserNotFoundError{}
	}

	if updateAttrs.Email != user.Email {
		if isAvailable, err := um.repository.IsEmailAvailable(updateAttrs.Email); err != nil || !isAvailable {
			tools.Log().WithField("error", err).Debug("A duplicated email was provided")
			return user, util.DuplicatedEmailError{}
		}
	}

	if updateAttrs.Password != "" {
		hashedPassword, err := util.HashPassword(updateAttrs.Password)
		if err != nil {
			return user, err
		}
		user.Password = hashedPassword
	}

	return um.repository.Update(user.Email, mappers.AssignUserUpdateToUser(user, updateAttrs))
}
