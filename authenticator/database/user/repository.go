package user

import "github.com/bixlabs/authentication/authenticator/structures"

type Repository interface {
	Create(user structures.User) (structures.User, error)
	IsEmailAvailable(email string) (bool, error)
	VerifyPassword(password string) (bool, error)
	ChangePassword(password string) error
	SaveResetPasswordToken(token string) error
	VerifyResetPasswordToken(token string) (bool, error)
}
