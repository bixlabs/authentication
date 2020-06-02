package mappers

import (
	"github.com/bixlabs/authentication/api/user_management/structures/create"
	"github.com/bixlabs/authentication/authenticator/structures"
)

func CreateResponseToResult(user structures.User) create.Result {
	return create.Result{User: user}
}
