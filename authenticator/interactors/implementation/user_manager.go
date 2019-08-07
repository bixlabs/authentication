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
	repository user.Repository
}

func NewUserManager(repository user.Repository) interactors.UserManager {
	return &userManager{repository: repository}
}

func (um userManager) Create(user structures.User) (structures.User, error) {
	if user, err := um.generatePasswordIfEmpty(&user); err != nil {
		return user, err
	}

	return um.createUser(user)
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

func (um userManager) createUser(user structures.User) (structures.User, error) {
	if err := um.hasValidationIssue(user); err != nil {
		return user, err
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword

	user, err = um.repository.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (um userManager) hasValidationIssue(user structures.User) error {
	if err := util.IsValidEmail(user.Email); err != nil {
		return err
	}

	if isAvailable, err := um.repository.IsEmailAvailable(user.Email); err != nil || !isAvailable {
		tools.Log().WithField("error", err).Debug("A duplicated email was provided")
		return util.DuplicatedEmailError{}
	}

	if err := util.CheckPasswordLength(user.Password); err != nil {
		return err
	}

	return nil
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
