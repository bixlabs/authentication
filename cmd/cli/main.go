package main

import (
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/in_memory"
	"github.com/bixlabs/authentication/tools"
	"github.com/gin-gonic/gin/json"
)

func main() {
	tools.InitializeLogger()
	authOperations := implementation.NewAuthenticator(in_memory.NewUserRepo())

	_, _ = authOperations.Login("", "")
	_, _ = authOperations.Signup(structures.User{Email: "email@bixlabs.com", Password: "secured_password"})
	user, _ := authOperations.Login("email@bixlabs.com", "secured_password")
	jsonUser, _ := json.Marshal(user)
	println(string(jsonUser))
	_ = authOperations.ChangePassword(structures.User{Email: "email@bixlabs.com", Password: "secured_password"}, "secured_password2")
	_ = authOperations.ResetPassword("")
}
