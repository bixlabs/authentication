package database

import "github.com/bixlabs/authentication/authenticator/structures"

type UserRepository interface {
	create(user structures.User) (error, structures.User)
	verifyPassword(password string) (error, bool)
	changePassword(password string) error
	saveResetPasswordToken(token string) error
	verifyResetPasswordToken(token string) (error, bool)
}
