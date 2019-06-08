package util

import (
	"github.com/bixlabs/authentication/tools"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const PasswordManager = 8
const EmailValidationRegex = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])"
const SignupInvalidEmailMessage = "Email is not valid"
const SignupPasswordLengthMessage = "Password should have at least 8 characters"

func IsValidEmail(email string) error {
	if isValidEmail, _ := regexp.MatchString(EmailValidationRegex, email); !isValidEmail {
		tools.Log().Debug("An invalid email was provided: " + email)
		return InvalidEmailError{}
	}
	return nil
}

type InvalidEmailError struct{}

func (e InvalidEmailError) Error() string {
	return SignupInvalidEmailMessage
}

func CheckPasswordLength(password string) error {
	if len(password) < PasswordManager {
		tools.Log().Debug("A password with incorrect length was provided")
		return PasswordLengthError{}
	}
	return nil
}

type PasswordLengthError struct{}

func (e PasswordLengthError) Error() string {
	return SignupPasswordLengthMessage
}

func HashPassword(password string) (string, error) {
	pass := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		tools.Log().WithField("error", err).Error("Password hash failed")
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, plainPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		tools.Log().Debug("A wrong password was provided")
		return WrongCredentialsError{}
	}
	return nil
}

type WrongCredentialsError struct{}

func (e WrongCredentialsError) Error() string {
	return "wrong credentials"
}
