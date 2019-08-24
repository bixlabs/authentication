package interactors

import (
	"github.com/bixlabs/authentication/authenticator/structures"
)

// UserManager manage the CRUD of users
type UserManager interface {
	Create(user structures.User) (structures.User, error)
	Delete(email string) error
	Find(email string) (structures.User, error)
	Update(email string, updateAttrs structures.UpdateUser) (structures.User, error)
}
