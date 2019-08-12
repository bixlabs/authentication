package main

import (
	"encoding/json"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
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

	sender := emailProviders.NewEmailProvider()

	passwordManager := implementation.NewPasswordManager(userRepo, sender)
	auth := implementation.NewAuthenticator(userRepo, sender)

	_, _ = auth.Signup(structures.User{Email: "acabrera@bixlabs.com", Password: "secured_password"})
	user, _ := auth.Login("acabrera@bixlabs.com", "secured_password")
	jsonUser, _ := json.Marshal(user)
	tools.Log().Info(string(jsonUser))
	_ = passwordManager.ChangePassword(structures.User{Email: "acabrera@bixlabs.com",
		Password: "secured_password"}, "secured_password2")
	user, _ = auth.Login("acabrera@bixlabs.com", "secured_password2")
	jsonUser, _ = json.Marshal(user)
	tools.Log().Info(string(jsonUser))
	code, _ := passwordManager.ForgotPassword("acabrera@bixlabs.com")
	_ = passwordManager.ResetPassword("acabrera@bixlabs.com", code, "secured_password3")
	user, _ = auth.Login("acabrera@bixlabs.com", "secured_password3")
	jsonUser, _ = json.Marshal(user)
	tools.Log().Info(string(jsonUser))
}
