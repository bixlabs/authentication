package mappers

import (
	"github.com/bixlabs/authentication/admincli/authentication/structures/createuser"
	"github.com/bixlabs/authentication/admincli/authentication/structures/updateuser"
	"github.com/bixlabs/authentication/authenticator/structures"
)

func CreateUserCommandToUser(command createuser.Command) structures.User {
	return structures.User{Email: command.Email, Password: command.Password, GivenName: command.GivenName,
		SecondName: command.SecondName, FamilyName: command.FamilyName, SecondFamilyName: command.SecondFamilyName}
}

func UpdateUserCommandToUpdateUser(command updateuser.Command) structures.UpdateUser {
	return structures.UpdateUser{Email: command.Email, Password: command.Password, GivenName: command.GivenName,
		SecondName: command.SecondName, FamilyName: command.FamilyName, SecondFamilyName: command.SecondFamilyName}
}
