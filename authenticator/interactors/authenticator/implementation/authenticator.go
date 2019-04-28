package implementation

import (
	"errors"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"github.com/bixlabs/authentication/tools"
	"regexp"
)

const signupDuplicateEmailMessage = "Email is already taken"
const emailValidationRegex = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])"
const signupInvalidEmailMessage = "Email is not valid"
const signupPasswordLengthMessage = "Password should be at least 8 characters"

type authenticator struct {
	repository user.Repository
}

func NewAuthenticator(repository user.Repository) *authenticator {
	return &authenticator{repository}
}

func (auth authenticator) Login(email, password string) (error, login.Response) {
	tools.Log().Warn("Login: Not Implemented yet")
	return nil, login.Response{}
}

func (auth authenticator) Signup(user structures.User) error {

	if err, isAvailable := auth.repository.IsEmailAvailable(user.Email); err != nil || !isAvailable {
		return errors.New(signupDuplicateEmailMessage)
	}

	if isValidEmail, _ := regexp.MatchString(emailValidationRegex, user.Email); !isValidEmail {
		return errors.New(signupInvalidEmailMessage)
	}

	if len(user.Password) < 8 {
		return errors.New(signupPasswordLengthMessage)
	}

	_, _ =auth.repository.Create(user)
	tools.Log().Info("A user was created")
	return nil

}

func (auth authenticator) ChangePassword(oldPassword, newPassword string) error {
	tools.Log().Warn("ChangePassword: Not Implemented yet")
	return nil
}

func (auth authenticator) ResetPassword(email string) error {
	tools.Log().Warn("ResetPassword: Not Implemented yet")
	return nil
}
