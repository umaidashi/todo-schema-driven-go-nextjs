package repository

import (
	"context"
	"server/domain/model"
)

type TodoRepository interface {
	One(ctx context.Context, id int) (model.Todo, error)
	List(ctx context.Context) ([]model.Todo, error)
}
