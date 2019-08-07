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

	authOperations := implementation.NewAuthenticator(userRepo, memory.DummySender{})

	cmd.SetAuthenticator(authOperations)
	cmd.Execute()
}
