
package main

import (
	. "github.com/bixlabs/go-layout/business/use_cases"
	. "github.com/bixlabs/go-layout/business/structures"
)

func main() {
	todoOperations := NewTodoOperationsHandler()

	todoOperations.Create(Todo{})
	todoOperations.Read("1")
	todoOperations.Update(Todo{})
	todoOperations.Delete("1")
}

