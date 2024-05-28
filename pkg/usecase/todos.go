package usecase

import (
	"go-mongo-todos/pkg/domain/entity/todo"
)

type FetchTodoUseCase struct{}

func (f FetchTodoUseCase) GetTodo(id todo.ID) (todo.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (f FetchTodoUseCase) GetTodoList() ([]todo.Todo, error) {
	//TODO implement me
	panic("implement me")
}
