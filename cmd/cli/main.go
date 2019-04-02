package main

import (
	"github.com/bixlabs/authenticationt/todo/interactors"
	"github.com/bixlabs/authentication/todo/structures"
	"github.com/bixlabs/authentication/tools"
)

func main() {
	tools.InitializeLogger()
	todoOperations := interactors.NewTodoOperationsHandler()

	todoOperations.Create(structures.Todo{})
	todoOperations.Read("1")
	todoOperations.Update(structures.Todo{})
	todoOperations.Delete("1")
}
