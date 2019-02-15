package useCases

import "github.com/bixlabs/authentication/authentication/structures"

type Authenticator interface {
	Login(email, password string) (error, structures.LoginResponse)
	Signup(user structures.User) (error, structures.LoginResponse)
	PasswordChange(oldPassword, newPassword string) error
	ResetPassword(email string) error
}
