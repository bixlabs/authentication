package implementation

import (
	"errors"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"github.com/bixlabs/authentication/tools"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const signupDuplicateEmailMessage = "Email is already taken"
const emailValidationRegex = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])"
const signupInvalidEmailMessage = "Email is not valid"
const signupPasswordLengthMessage = "Password should be at least 8 characters"
const passwordMinLength = 8

type authenticator struct {
	repository user.Repository
}

func NewAuthenticator(repository user.Repository) *authenticator {
	return &authenticator{repository}
}

func (auth authenticator) Login(email, password string) (login.Response, error) {
	tools.Log().Warn("Login: Not Implemented yet")
	return login.Response{}, nil
}

func (auth authenticator) Signup(user structures.User) (structures.User, error) {
	if err := auth.hasValidationIssue(user); err != nil {
		return user, err
	}

	hashedPassword, err := auth.hashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword

	user, err = auth.repository.Create(user)
	if err != nil {
		return user, err
	}

	tools.Log().Info("A user was created")
	return user, nil
}

func (auth authenticator) hasValidationIssue(user structures.User) error {
	if isValidEmail, _ := regexp.MatchString(emailValidationRegex, user.Email); !isValidEmail {
		tools.Log().Debug("An invalid email was provided: " + user.Email)
		return errors.New(signupInvalidEmailMessage)
	}
	if isAvailable, err := auth.repository.IsEmailAvailable(user.Email); err != nil || !isAvailable {
		tools.Log().WithField("error", err).Debug("A duplicated email was provided")
		return errors.New(signupDuplicateEmailMessage)
	}
	if len(user.Password) < passwordMinLength {
		tools.Log().Debug("A password with incorrect length was provided")
		return errors.New(signupPasswordLengthMessage)
	}
	return nil
}

func (auth authenticator) hashPassword(password string) (string, error) {
	pass := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		tools.Log().WithField("error", err).Error("Password hash failed")
		return "", err
	}
	return string(hashedPassword), nil
}

func (auth authenticator) ChangePassword(oldPassword, newPassword string) error {
	tools.Log().Warn("ChangePassword: Not Implemented yet")
	return nil
}

func (auth authenticator) ResetPassword(email string) error {
	tools.Log().Warn("ResetPassword: Not Implemented yet")
	return nil
}
