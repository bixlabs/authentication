package interactors

import "github.com/bixlabs/authentication/authenticator/structures"

type PasswordManager interface {
	ChangePassword(user structures.User, newPassword string) error
	ResetPassword(email string, code string, newPassword string) error
	SendResetPasswordRequest(email string) error
}
