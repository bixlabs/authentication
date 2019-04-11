package main

import (
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/tools"
)

func main() {
	tools.InitializeLogger()
	authOperations := interactors.NewAuthenticator()

	_, _ = authOperations.Login("", "")
	_, _ = authOperations.Signup(structures.User{})
	_ = authOperations.ChangePassword("", "")
	_ = authOperations.ResetPassword("")
}
