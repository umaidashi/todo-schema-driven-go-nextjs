package usecase

import (
	"context"
	"server/domain/model"
	"server/domain/repository"
)

type TodoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	return TodoUsecase{
		todoRepo: todoRepo,
	}
}

func (u TodoUsecase) GetTodos(ctx context.Context) ([]model.Todo, error) {
	return u.todoRepo.List(ctx)
}
