package main

import (
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/in_memory"
	"github.com/bixlabs/authentication/database/user/sqlite"
	"github.com/bixlabs/authentication/tools"
	"github.com/gin-gonic/gin/json"
)

func main() {
	tools.InitializeLogger()
	userRepo, closeDB := sqlite.NewSqliteStorage()
	defer closeDB()
	sender := in_memory.DummySender{}
	passwordManager := implementation.NewPasswordManager(userRepo, sender)
	auth := implementation.NewAuthenticator(userRepo, sender)

	_, _ = auth.Signup(structures.User{Email: "email@bixlabs.com", Password: "secured_password"})
	user, _ := auth.Login("email@bixlabs.com", "secured_password")
	jsonUser, _ := json.Marshal(user)
	println(string(jsonUser))
	_ = passwordManager.ChangePassword(structures.User{Email: "email@bixlabs.com", Password: "secured_password"}, "secured_password2")
	user, _ = auth.Login("email@bixlabs.com", "secured_password2")
	jsonUser, _ = json.Marshal(user)
	println(string(jsonUser))
	code, _ := passwordManager.ForgotPassword("email@bixlabs.com")
	_ = passwordManager.ResetPassword("email@bixlabs.com", code, "secured_password3")
	user, _ = auth.Login("email@bixlabs.com", "secured_password3")
	jsonUser, _ = json.Marshal(user)
	println(string(jsonUser))

}
