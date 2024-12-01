package dao

import (
	"context"
	"server/domain/model"
	"server/domain/repository"
	"server/pkg/ent"

	"github.com/samber/lo"
)

type TodoDao struct {
	db *ent.Client
}

func NewTodoDao(db *ent.Client) repository.TodoRepository {
	return &TodoDao{db}
}

func (d TodoDao) One(ctx context.Context, id int) (model.Todo, error) {
	todo, err := d.db.Todo.Get(ctx, id)
	if err != nil {
		return model.Todo{}, err
	}
	return model.Todo{
		ID:          model.TodoId(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      model.TodoStatusOf(todo.Status.String()),
		Priority:    model.PriorityOf(todo.Priority.String()),
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}, nil
}

func (d TodoDao) List(ctx context.Context) ([]model.Todo, error) {
	todos, err := d.db.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	result := lo.Map(todos, func(todo *ent.Todo, _ int) model.Todo {
		return model.Todo{
			ID:          model.TodoId(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			Status:      model.TodoStatusOf(todo.Status.String()),
			Priority:    model.PriorityOf(todo.Priority.String()),
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		}
	})
	return result, nil
}
