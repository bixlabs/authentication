package main

import (
	"github.com/bixlabs/authentication/api/todo"
	_ "github.com/bixlabs/authentication/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/bixlabs/authentication/todo/useCases"
	"github.com/bixlabs/authentication/tools"
	_ "github.com/joho/godotenv/autoload"
)

// @title Go-Layout
// @version 1.0
// @description Simple template to use for golang projects

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email jarrieta@bixlabs.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /v1

func main() {
	tools.InitializeLogger()
	todoOperations := useCases.NewTodoOperationsHandler()
	todo.NewTodoRestConfigurator(todoOperations)
}
