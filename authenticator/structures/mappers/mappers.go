package mappers

import (
	"github.com/bixlabs/authentication/authenticator/structures"
)

func AssignUserUpdateToUser(user structures.User, updateAttrs structures.UpdateUser) structures.User {
	email := updateAttrs.Email
	if updateAttrs.Email == "" {
		email = user.Email
	}

	givenName := updateAttrs.GivenName
	if givenName == "" {
		givenName = user.GivenName
	}

	secondName := updateAttrs.SecondName
	if secondName == "" {
		secondName = user.SecondName
	}

	familyName := updateAttrs.FamilyName
	if familyName == "" {
		familyName = user.FamilyName
	}

	secondFamilyName := updateAttrs.SecondFamilyName
	if secondFamilyName == "" {
		secondFamilyName = user.SecondFamilyName
	}

	return structures.User{
		ID:               user.ID,
		Email:            email,
		Password:         user.Password,
		GivenName:        givenName,
		SecondName:       secondName,
		FamilyName:       familyName,
		SecondFamilyName: secondFamilyName,
	}
}
