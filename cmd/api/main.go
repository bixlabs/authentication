package main

import (
	"github.com/bixlabs/go-layout/api"
	"github.com/bixlabs/go-layout/business/use_cases"
)

func main() {
	todoOperations := use_cases.NewTodoOperationsHandler()
	api.NewTodoRestConfigurator(todoOperations)
}


