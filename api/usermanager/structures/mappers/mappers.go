package mappers

import (
	"github.com/bixlabs/authentication/api/usermanager/structures/create"
	"github.com/bixlabs/authentication/api/usermanager/structures/update"
	"github.com/bixlabs/authentication/authenticator/structures"
)

func CreateRequestToUser(request create.Request) structures.User {
	return structures.User{Email: request.Email, Password: request.Password, GivenName: request.GivenName,
		SecondName: request.SecondName, FamilyName: request.FamilyName, SecondFamilyName: request.SecondFamilyName}
}

func UpdateRequestToUpdateUser(request update.Request) structures.UpdateUser {
	return structures.UpdateUser{Email: request.Email, Password: request.Password, GivenName: request.GivenName,
		SecondName: request.SecondName, FamilyName: request.FamilyName, SecondFamilyName: request.SecondFamilyName}
}
