package main

import (
	"encoding/json"
	"fmt"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	email "github.com/bixlabs/authentication/authenticator/provider/email/implementation"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/sqlite"
	emailProviders "github.com/bixlabs/authentication/email/providers"
	"github.com/bixlabs/authentication/tools"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	tools.InitializeLogger()
	userRepo, closeDB := sqlite.NewSqliteStorage()
	defer closeDB()

	sender := email.NewSender(emailProviders.NewEmailProvider())

	passwordManager := implementation.NewPasswordManager(userRepo, sender)
	auth := implementation.NewAuthenticator(userRepo, sender)

	_, _ = auth.Signup(structures.User{Email: "example@bixlabs.com", Password: "secured_password"})
	user, _ := auth.Login("example@bixlabs.com", "secured_password")
	jsonUser, _ := json.Marshal(user)
	tools.Log().Info(string(jsonUser))
	_ = passwordManager.ChangePassword(structures.User{Email: "example@bixlabs.com",
		Password: "secured_password"}, "secured_password2")
	user, _ = auth.Login("example@bixlabs.com", "secured_password2")
	jsonUser, _ = json.Marshal(user)
	tools.Log().Info(string(jsonUser))
	code, ok := passwordManager.ForgotPassword("example@bixlabs.com")
	fmt.Println(ok)
	_ = passwordManager.ResetPassword("example@bixlabs.com", code, "secured_password3")
	user, _ = auth.Login("example@bixlabs.com", "secured_password3")
	jsonUser, _ = json.Marshal(user)
	tools.Log().Info(string(jsonUser))
}
