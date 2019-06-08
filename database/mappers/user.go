package mappers

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/model"
	"strconv"
)

func UserToDatabaseModel(user structures.User) model.User {
	id, err := strconv.Atoi(user.ID)
	if err != nil && user.ID != "" {
		panic(err)
	}
	return model.User{ID: uint(id),
		Email:            user.Email,
		Password:         user.Password,
		GivenName:        user.GivenName,
		SecondName:       user.SecondName,
		FamilyName:       user.FamilyName,
		SecondFamilyName: user.SecondFamilyName,
		ResetToken:       user.ResetToken}
}

func DatabaseModelToUser(user model.User) structures.User {
	id := strconv.Itoa(int(user.ID))
	return structures.User{ID: id,
		Email:            user.Email,
		Password:         user.Password,
		GivenName:        user.GivenName,
		SecondName:       user.SecondName,
		FamilyName:       user.FamilyName,
		SecondFamilyName: user.SecondFamilyName,
		ResetToken:       user.ResetToken}
}
