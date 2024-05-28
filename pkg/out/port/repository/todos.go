package repository

import "go-mongo-todos/pkg/domain/entity/todo"

type CreateTodoModel struct {
	Title string
}

type TodoRepository interface {
	GetTodo(id todo.ID) (todo.Todo, error)
	GetTodoList() ([]todo.Todo, error)
	CreateTodo(input CreateTodoModel) (todo.Todo, error)
	DeleteTodo(id todo.ID) error
}
