package mappers

import (
	"github.com/bixlabs/authentication/admincli/authentication/structures/createuser"
	"github.com/bixlabs/authentication/admincli/authentication/structures/finduser"
	"github.com/bixlabs/authentication/admincli/authentication/structures/resetpassword"
	"github.com/bixlabs/authentication/admincli/authentication/structures/updateuser"
	"github.com/bixlabs/authentication/authenticator/structures"
)

func CreateUserCommandToUser(command createuser.Command) structures.User {
	return structures.User{Email: command.Email, Password: command.Password, GivenName: command.GivenName,
		SecondName: command.SecondName, FamilyName: command.FamilyName, SecondFamilyName: command.SecondFamilyName}
}

func UserToCreateUserResult(user structures.User) createuser.Result {
	return createuser.Result{ID: user.ID, Email: user.Email, Password: user.Password, GivenName: user.GivenName,
		SecondName: user.SecondName, FamilyName: user.FamilyName, SecondFamilyName: user.SecondFamilyName}
}

func UpdateUserCommandToUpdateUser(command updateuser.Command) structures.UpdateUser {
	return structures.UpdateUser{Email: command.Email, Password: command.Password, GivenName: command.GivenName,
		SecondName: command.SecondName, FamilyName: command.FamilyName, SecondFamilyName: command.SecondFamilyName}
}

func UserToUpdateUserResult(user structures.User) updateuser.Result {
	return updateuser.Result{ID: user.ID, Email: user.Email, Password: user.Password, GivenName: user.GivenName,
		SecondName: user.SecondName, FamilyName: user.FamilyName, SecondFamilyName: user.SecondFamilyName}
}

func ResetUserCommandToUpdateUser(command resetpassword.Command) structures.UpdateUser {
	return structures.UpdateUser{Password: command.Password}
}

func UserToFindUserResult(user structures.User) finduser.Result {
	return finduser.Result{ID: user.ID, Email: user.Email, Password: user.Password, GivenName: user.GivenName,
		SecondName: user.SecondName, FamilyName: user.FamilyName, SecondFamilyName: user.SecondFamilyName}
}
