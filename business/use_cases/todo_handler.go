package use_cases

import (
	. "github.com/bixlabs/go-layout/business/structures"
)


type TodoOperations interface {
	create(todo Todo) *Todo
	read(id string) *Todo
	update(todo Todo) *Todo
	delete(id string) bool
}

/*

Using the import like this "github.com/bixlabs/go-layout/business/structures" (without the dot) this is how you will have to reference the To_do struct:

type TodoOperations interface {
	create(todo structures.Todo) structures.Todo
}

*/

// TODO: open to discussion, I'm not sure where the implementation should be.

type TodoOperationsHandler struct {}

func (handler TodoOperationsHandler) create(todo Todo) *Todo {
	println("A Todo was created")
	return nil
}

func (handler TodoOperationsHandler) read(id string) *Todo {
	println("A Todo was retrieved")
	return nil
}

func (handler TodoOperationsHandler) update(todo Todo) *Todo {
	println("A Todo was updated")
	return nil
}

func (handler TodoOperationsHandler) delete(id string) bool {
	println("A Todo was deleted")
	return true
}


