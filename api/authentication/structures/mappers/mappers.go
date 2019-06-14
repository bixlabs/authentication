package mappers

import (
	"github.com/bixlabs/authentication/api/authentication/structures/change_password"
	api "github.com/bixlabs/authentication/api/authentication/structures/login"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
)

func ChangePasswordRequestToUser(request change_password.Request) structures.User {
	return structures.User{Email: request.Email, Password: request.OldPassword}
}

func LoginResponseToResult(r login.Response) *api.Result {
	return &api.Result{Response: r}
}
