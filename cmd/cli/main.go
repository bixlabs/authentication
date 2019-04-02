package main

import (
	"github.com/bixlabs/go-layout/todo/interactors"
	"github.com/bixlabs/go-layout/todo/structures"
	"github.com/bixlabs/go-layout/tools"
)

func main() {
	tools.InitializeLogger()
	todoOperations := interactors.NewTodoOperationsHandler()

	todoOperations.Create(structures.Todo{})
	todoOperations.Read("1")
	todoOperations.Update(structures.Todo{})
	todoOperations.Delete("1")
}
