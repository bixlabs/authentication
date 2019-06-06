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
	userRepo, sender := in_memory.NewUserRepo(), in_memory.DummySender{}
	passwordManager := implementation.NewPasswordManager(userRepo, sender)
	auth := implementation.NewAuthenticator(userRepo, sender)

	_, _ = auth.Signup(structures.User{Email: "email@bixlabs.com", Password: "secured_password"})
	user, _ := auth.Login("email@bixlabs.com", "secured_password")
	jsonUser, _ := json.Marshal(user)
	println(string(jsonUser))
	_ = passwordManager.ChangePassword(structures.User{Email: "email@bixlabs.com", Password: "secured_password"}, "secured_password2")
	_, _ = auth.Login("email@bixlabs.com", "secured_password2")
	_ = passwordManager.SendResetPasswordRequest("email@bixlabs.com")
	_ = passwordManager.ResetPassword("", "", "")
}
