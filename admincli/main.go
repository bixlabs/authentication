package main

import (
	"github.com/bixlabs/authentication/admincli/cmd"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	email "github.com/bixlabs/authentication/authenticator/provider/email/implementation"
	"github.com/bixlabs/authentication/database/user/sqlite"
	"github.com/bixlabs/authentication/email/providers"
	"github.com/bixlabs/authentication/tools"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	tools.InitializeLogger()
	userRepo, closeDB := sqlite.NewSqliteStorage()
	defer closeDB()
	sender := email.NewSender(providers.NewEmailProvider())
	auth := implementation.NewAuthenticator(userRepo, sender)
	userManager := implementation.NewUserManager(auth, userRepo)
	cmd.SetUserManager(userManager)
	cmd.Execute()
}
