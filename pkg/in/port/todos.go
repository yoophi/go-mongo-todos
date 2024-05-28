package port

import (
	"go-mongo-todos/pkg/domain/entity/todo"
)

type CreateTodoRequest struct {
}

type CreateTodoResponse struct{}

type DeleteTodoRequest struct{}

type DeleteTodoResponse struct{}

type GetTodoRequest struct {
	ID string
}

type GetTodoResponse struct {
	Todo todo.Todo
}

type GetTodoListResponse struct {
	TodoList []todo.Todo
}

type GetTodoUseCase interface {
	GetTodo(request GetTodoRequest) (GetTodoResponse, error)
}

type GetTodoListUseCase interface {
	GetTodoList() (GetTodoListResponse, error)
}

type CreateTodoUseCase interface {
	CreateTodo(request CreateTodoRequest) (CreateTodoResponse, error)
}

type DeleteTodoUseCase interface {
	DeleteTodo(request DeleteTodoRequest) (DeleteTodoResponse, error)
}
