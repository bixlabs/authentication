package interactors

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"github.com/bixlabs/authentication/tools"
)

type Authenticator interface {
	Login(email, password string) (error, login.Response)
	Signup(user structures.User) (error, login.Response)
	ChangePassword(oldPassword, newPassword string) error
	ResetPassword(email string) error
}

// TODO: I'm not sure if this code should be here, should we separate the abstraction from the implementation in files?
type authenticator struct {
}

func NewAuthenticator() *authenticator {
	return &authenticator{}
}

func (authenticator) Login(email, password string) (error, login.Response) {
	tools.Log().Warn("Login: Not Implemented yet")
	return nil, login.Response{}
}

func (authenticator) Signup(user structures.User) (error, login.Response) {
	tools.Log().Warn("Signup: Not Implemented yet")
	return nil, login.Response{}

}

func (authenticator) ChangePassword(oldPassword, newPassword string) error {
	tools.Log().Warn("ChangePassword: Not Implemented yet")
	return nil
}

func (authenticator) ResetPassword(email string) error {
	tools.Log().Warn("ResetPassword: Not Implemented yet")
	return nil
}
