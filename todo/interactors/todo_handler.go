package interactors

import (
	"github.com/bixlabs/go-layout/todo/structures"
	"github.com/bixlabs/go-layout/tools"
	"github.com/sirupsen/logrus"
)

type TodoOperations interface {
	Create(todo structures.Todo) *structures.Todo
	Read(id string) *structures.Todo
	Update(todo structures.Todo) *structures.Todo
	Delete(id string) bool
}

// TODO: open to discussion, I'm not sure where the implementation should be.

type TodoOperationsHandler struct{}

func NewTodoOperationsHandler() TodoOperationsHandler {
	return TodoOperationsHandler{}
}

func (handler TodoOperationsHandler) Create(todo structures.Todo) *structures.Todo {
	tools.Log().WithFields(logrus.Fields{"ID": todo.ID, "Name": todo.Name}).Info("A todo was created")
	return &todo
}

func (handler TodoOperationsHandler) Read(id string) *structures.Todo {
	tools.Log().WithFields(logrus.Fields{"ID": id}).Info("A todo was retrieved")
	return &structures.Todo{ID: id}
}

func (handler TodoOperationsHandler) Update(todo structures.Todo) *structures.Todo {
	tools.Log().WithFields(logrus.Fields{"ID": todo.ID, "Name": todo.Name}).Info("A todo was updated")
	return &todo
}

func (handler TodoOperationsHandler) Delete(id string) bool {
	tools.Log().WithFields(logrus.Fields{"ID": id}).Info("A todo was deleted")
	return true
}
