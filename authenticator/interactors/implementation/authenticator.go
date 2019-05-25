package implementation

import (
	"errors"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"github.com/bixlabs/authentication/tools"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const duplicatedEmailMessage = "Email is already taken"
const emailValidationRegex = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])"
const invalidEmailMessage = "Email is not valid"
const invalidPasswordLengthMessage = "Password should be at least 8 characters"
const passwordMinLength = 8

type authenticator struct {
	repository user.Repository
}

func NewAuthenticator(repository user.Repository) interactors.Authenticator {
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

	hashedPassword, err := hashPassword(user.Password)
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
	if err := isValidEmail(user.Email); err != nil {
		return err
	}
	if isAvailable, err := auth.repository.IsEmailAvailable(user.Email); err != nil || !isAvailable {
		tools.Log().WithField("error", err).Debug("A duplicated email was provided")
		return errors.New(duplicatedEmailMessage)
	}

	if err := checkPasswordLength(user.Password); err != nil {
		return err
	}

	return nil
}

func isValidEmail(email string) error {
	if isValidEmail, _ := regexp.MatchString(emailValidationRegex, email); !isValidEmail {
		tools.Log().Debug("An invalid email was provided: " + email)
		return errors.New(invalidEmailMessage)
	}
	return nil
}



func checkPasswordLength(password string) error {
	if len(password) < passwordMinLength {
		tools.Log().Debug("A password with incorrect length was provided")
		return errors.New(invalidPasswordLengthMessage)
	}
	return nil
}

func hashPassword(password string) (string, error) {
	pass := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		tools.Log().WithField("error", err).Error("Password hash failed")
		return "", err
	}
	return string(hashedPassword), nil
}

func (auth authenticator) ChangePassword(user structures.User, newPassword string) error {
	if err := isValidEmail(user.Email); err != nil {
		return err
	}

	if err := checkPasswordLength(newPassword); err != nil {
		return err
	}

	oldHashedPassword, err := auth.repository.GetHashPassword(user.Email)
	if err != nil {
		return err
	}

	if err := verifyPassword(oldHashedPassword, user.Password); err != nil {
		return err
	}

	hashedPassword, err := hashPassword(newPassword)

	if err != nil {
		return err
	}

	return auth.repository.ChangePassword(user.Email, hashedPassword)
}

func verifyPassword(oldHashedPassword, plainPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(oldHashedPassword), []byte(plainPassword)); err != nil {
		return errors.New("wrong old password")
	}
	return nil
}

func (auth authenticator) ResetPassword(email string) error {
	tools.Log().Warn("ResetPassword: Not Implemented yet")
	return nil
}
