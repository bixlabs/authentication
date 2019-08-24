package createuser

import "github.com/bixlabs/authentication/admincli/usermanager/structures/finduser"

// Command received by the create user command
type Command struct {
	Email            string
	Password         string
	GivenName        string
	SecondName       string
	FamilyName       string
	SecondFamilyName string
}

// Result to the create user command
type Result finduser.Result
