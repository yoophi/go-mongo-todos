package mem

import (
	"go-mongo-todos/pkg/domain/entity/todo"
	"go-mongo-todos/pkg/out/port/repository"
)

type TodoRepository struct{}

func (t TodoRepository) GetTodo(id todo.ID) (todo.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoRepository) GetTodoList() ([]todo.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoRepository) CreateTodo(input repository.CreateTodoModel) (todo.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoRepository) DeleteTodo(id todo.ID) error {
	//TODO implement me
	panic("implement me")
}
