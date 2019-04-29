package interactors

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
)

type Authenticator interface {
	Login(email, password string) (login.Response, error)
	Signup(user structures.User) (structures.User, error)
	ChangePassword(oldPassword, newPassword string) error
	ResetPassword(email string) error
}
