package mappers

import (
	"github.com/bixlabs/authentication/admincli/usermanager/structures/createuser"
	"github.com/bixlabs/authentication/admincli/usermanager/structures/finduser"
	"github.com/bixlabs/authentication/admincli/usermanager/structures/resetpassword"
	"github.com/bixlabs/authentication/admincli/usermanager/structures/updateuser"
	"github.com/bixlabs/authentication/authenticator/structures"
)

// CreateUserCommandToUser maps from a create user command to a business user
func CreateUserCommandToUser(command createuser.Command) structures.User {
	return structures.User{Email: command.Email, Password: command.Password, GivenName: command.GivenName,
		SecondName: command.SecondName, FamilyName: command.FamilyName, SecondFamilyName: command.SecondFamilyName}
}

// UserToCreateUserResult maps from a business user a create user command result
func UserToCreateUserResult(user structures.User) createuser.Result {
	return createuser.Result{ID: user.ID, Email: user.Email, Password: user.Password, GivenName: user.GivenName,
		SecondName: user.SecondName, FamilyName: user.FamilyName, SecondFamilyName: user.SecondFamilyName}
}

// UpdateUserCommandToUpdateUser maps from a update user command to a business update user
func UpdateUserCommandToUpdateUser(command updateuser.Command) structures.UpdateUser {
	return structures.UpdateUser{Email: command.Email, Password: command.Password, GivenName: command.GivenName,
		SecondName: command.SecondName, FamilyName: command.FamilyName, SecondFamilyName: command.SecondFamilyName}
}

// UserToUpdateUserResult maps from a business user to an update user command result
func UserToUpdateUserResult(user structures.User) updateuser.Result {
	return updateuser.Result{ID: user.ID, Email: user.Email, Password: user.Password, GivenName: user.GivenName,
		SecondName: user.SecondName, FamilyName: user.FamilyName, SecondFamilyName: user.SecondFamilyName}
}

// ResetUserCommandToUpdateUser maps from a reset password command to a business update user
func ResetUserCommandToUpdateUser(command resetpassword.Command) structures.UpdateUser {
	return structures.UpdateUser{Password: command.Password}
}

// UserToFindUserResult maps from a business user to a find user command result
func UserToFindUserResult(user structures.User) finduser.Result {
	return finduser.Result{ID: user.ID, Email: user.Email, Password: user.Password, GivenName: user.GivenName,
		SecondName: user.SecondName, FamilyName: user.FamilyName, SecondFamilyName: user.SecondFamilyName}
}
