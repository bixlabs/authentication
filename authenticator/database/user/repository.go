package user

import "github.com/bixlabs/authentication/authenticator/structures"

type Repository interface {
	Create(user structures.User) (structures.User, error)
	IsEmailAvailable(email string) (bool, error)
	GetHashedPassword(email string) (string, error)
	ChangePassword(email, newPassword string) error
	Find(email string) (structures.User, error)
	UpdateResetToken(email, resetToken string) error
}
