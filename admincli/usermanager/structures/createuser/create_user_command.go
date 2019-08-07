package createuser

import "github.com/bixlabs/authentication/admincli/usermanager/structures/finduser"

type Command struct {
	Email            string
	Password         string
	GivenName        string
	SecondName       string
	FamilyName       string
	SecondFamilyName string
}

type Result finduser.Result
