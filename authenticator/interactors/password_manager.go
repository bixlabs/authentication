package interactors

import "github.com/bixlabs/authentication/authenticator/structures"

type PasswordManager interface {
	ChangePassword(user structures.User, newPassword string) error
	ResetPassword(email string, code int, newPassword string) error
	ForgotPassword(email string) (string, error)
}
