package mappers

import (
	"github.com/bixlabs/authentication/api/authentication/structures/changepass"
	api "github.com/bixlabs/authentication/api/authentication/structures/login"
	"github.com/bixlabs/authentication/api/authentication/structures/signup"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
)

func ChangePasswordRequestToUser(request changepass.Request) structures.User {
	return structures.User{Email: request.Email, Password: request.OldPassword}
}

func LoginResponseToResult(r login.Response) api.Result {
	return api.Result{Response: r}
}

func SignupRequestToUser(request signup.Request) structures.User {
	return structures.User{Email: request.Email, Password: request.Password, GivenName: request.GivenName,
		SecondName: request.SecondName, FamilyName: request.FamilyName, SecondFamilyName: request.SecondFamilyName}
}
