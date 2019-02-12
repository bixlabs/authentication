package main

import (
	. "github.com/bixlabs/go-layout/todo/structures"
	. "github.com/bixlabs/go-layout/todo/useCases"
	"github.com/bixlabs/go-layout/tools"
)

func main() {
	tools.InitializeLogger()
	todoOperations := NewTodoOperationsHandler()

	todoOperations.Create(Todo{})
	todoOperations.Read("1")
	todoOperations.Update(Todo{})
	todoOperations.Delete("1")
}
