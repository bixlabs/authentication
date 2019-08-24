package main

import (
	"github.com/bixlabs/authentication/admincli/cmd"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/database/user/memory"
	"github.com/bixlabs/authentication/database/user/sqlite"
	"github.com/bixlabs/authentication/tools"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	tools.InitializeLogger()
	userRepo, closeDB := sqlite.NewSqliteStorage()
	defer closeDB()
	auth := implementation.NewAuthenticator(userRepo, memory.DummySender{})
	userManager := implementation.NewUserManager(auth, userRepo)
	cmd.SetUserManager(userManager)
	cmd.Execute()
}
