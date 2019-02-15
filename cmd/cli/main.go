package main

import (
	. "github.com/bixlabs/authentication/todo/structures"
	. "github.com/bixlabs/authentication/todo/useCases"
	"github.com/bixlabs/authentication/tools"
)

func main() {
	tools.InitializeLogger()
	todoOperations := NewTodoOperationsHandler()

	todoOperations.Create(Todo{})
	todoOperations.Read("1")
	todoOperations.Update(Todo{})
	todoOperations.Delete("1")
}
