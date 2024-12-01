package handler

import (
	"context"
	"server/infra/applogger"
	"server/infra/dao"
	"server/pkg/ent"
	oas "server/pkg/oas"
	"server/usecase"
)

func NewHandler(db *ent.Client) oas.Handler {
	return &Handler{
		db: db,
	}
}

type Handler struct {
	db *ent.Client
}

func (h *Handler) TodoGet(ctx context.Context) (oas.TodoGetRes, error) {
	applogger.Info("TodoGet")

	repository := dao.NewTodoDao(h.db)
	usecase := usecase.NewTodoUsecase(repository)

	todos, err := usecase.GetTodos(ctx)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
