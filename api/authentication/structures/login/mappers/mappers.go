package mappers

import (
	api "github.com/bixlabs/authentication/api/authentication/structures/login"
	"github.com/bixlabs/authentication/authenticator/structures/login"
)

func LoginResponseToResult(r login.Response) *api.Result {
	return &api.Result{Response: r}
}
