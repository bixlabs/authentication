package authenticator

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
)

type Authenticator interface {
	Login(email, password string) (error, login.Response)
	Signup(user structures.User) error
	ChangePassword(oldPassword, newPassword string) error
	ResetPassword(email string) error
}
