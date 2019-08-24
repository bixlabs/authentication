package updateuser

import "github.com/bixlabs/authentication/admincli/usermanager/structures/finduser"

// Command received by the update user command
type Command struct {
	Email            string
	Password         string
	GivenName        string
	SecondName       string
	FamilyName       string
	SecondFamilyName string
}

// Result to the update user command
type Result finduser.Result
