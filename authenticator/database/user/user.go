package user

import "github.com/bixlabs/authentication/authenticator/structures"

type Repository interface {
	Create(user structures.User) (error, structures.User)
	IsEmailAvailable(email string) (error, bool)
	VerifyPassword(password string) (error, bool)
	ChangePassword(password string) error
	SaveResetPasswordToken(token string) error
	VerifyResetPasswordToken(token string) (error, bool)
}
