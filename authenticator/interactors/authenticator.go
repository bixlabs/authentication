package interactors

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
)

type Authenticator interface {
	Login(email, password string) (*login.Response, error)
	Signup(user structures.User) (structures.User, error)
	VerifyJWT(jwt string) (structures.User, error)

	// TODO: Maybe we should move this to another interface for Users
	Create(user structures.User) (structures.User, error)
	Delete(email string) error
	Find(email string) (structures.User, error)
	Update(email string, updateAttrs structures.UserUpdate) (structures.User, error)
}
