package main

import (
	. "github.com/bixlabs/go-layout/todo/structures"
	. "github.com/bixlabs/go-layout/todo/useCases"
)

func main() {
	todoOperations := NewTodoOperationsHandler()

	todoOperations.Create(Todo{})
	todoOperations.Read("1")
	todoOperations.Update(Todo{})
	todoOperations.Delete("1")
}
