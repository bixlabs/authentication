package main

import (
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/in_memory"
	"github.com/bixlabs/authentication/tools"
)

func main() {
	tools.InitializeLogger()
	authOperations := implementation.NewAuthenticator(in_memory.NewUserRepo())

	_, _ = authOperations.Login("", "")
	_, _ = authOperations.Signup(structures.User{})
	_ = authOperations.ChangePassword("", "")
	_ = authOperations.ResetPassword("")
}
