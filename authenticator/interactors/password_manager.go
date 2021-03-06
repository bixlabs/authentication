package interactors

import "github.com/bixlabs/authentication/authenticator/structures"

type PasswordManager interface {
	ChangePassword(user structures.User, newPassword string) error
	StartResetPassword(email string) (string, error)
	FinishResetPassword(email string, code string, newPassword string) error
}
