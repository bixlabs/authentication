package main

import (
	"github.com/bixlabs/go-layout/api"
	"github.com/bixlabs/go-layout/todo/useCases"
	"github.com/bixlabs/go-layout/tools"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	tools.InitializeLogger()
	todoOperations := useCases.NewTodoOperationsHandler()
	api.NewTodoRestConfigurator(todoOperations)
}
