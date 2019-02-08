package main

import (
	"github.com/bixlabs/go-layout/api"
	"github.com/bixlabs/go-layout/todo/useCases"
)

func main() {
	todoOperations := useCases.NewTodoOperationsHandler()
	api.NewTodoRestConfigurator(todoOperations)
}


