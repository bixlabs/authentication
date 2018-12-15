package main

import (
	"github.com/bixlabs/go-layout/api"
	"github.com/bixlabs/go-layout/todo/use_cases"
)

func main() {
	todoOperations := use_cases.NewTodoOperationsHandler()
	api.NewTodoRestConfigurator(todoOperations)
}


