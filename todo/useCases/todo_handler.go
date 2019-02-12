package useCases

import (
	. "github.com/bixlabs/go-layout/todo/structures"
	"github.com/bixlabs/go-layout/tools"
	"github.com/sirupsen/logrus"
)

type TodoOperations interface {
	Create(todo Todo) *Todo
	Read(id string) *Todo
	Update(todo Todo) *Todo
	Delete(id string) bool
}

/*

Using the import like this "github.com/bixlabs/go-layout/business/structures" (without the dot) this is how you will have to reference the To_do struct:

type TodoOperations interface {
	create(todo structures.Todo) structures.Todo
}

*/

// TODO: open to discussion, I'm not sure where the implementation should be.

type TodoOperationsHandler struct{}

func NewTodoOperationsHandler() TodoOperationsHandler {
	return TodoOperationsHandler{}
}

func (handler TodoOperationsHandler) Create(todo Todo) *Todo {
	tools.Log().WithFields(logrus.Fields{"ID": todo.ID, "Name": todo.Name}).Info("A todo was created")
	return &todo
}

func (handler TodoOperationsHandler) Read(id string) *Todo {
	tools.Log().WithFields(logrus.Fields{"ID": id}).Info("A todo was retrieved")
	return &Todo{ID:id}
}

func (handler TodoOperationsHandler) Update(todo Todo) *Todo {
	tools.Log().WithFields(logrus.Fields{"ID": todo.ID, "Name": todo.Name}).Info("A todo was updated")
	return &todo
}

func (handler TodoOperationsHandler) Delete(id string) bool {
	tools.Log().WithFields(logrus.Fields{"ID": id}).Info("A todo was deleted")
	return true
}
