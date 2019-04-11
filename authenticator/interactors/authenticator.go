package interactors

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
)

type Authenticator interface {
	Login(email, password string) (error, login.Response)
	Signup(user structures.User) (error, login.Response)
	PasswordChange(oldPassword, newPassword string) error
	ResetPassword(email string) error
}
