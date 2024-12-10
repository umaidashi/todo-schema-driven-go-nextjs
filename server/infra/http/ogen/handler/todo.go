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

	// TODO: ctx から id を取得する
	id := 1

	todo, err := usecase.GetTodo(ctx, id)
	if err != nil {
		return nil, err
	}

	return &oas.TodoBase{
		ID:          int(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		Status: oas.TodoStatus{
			Name:    todo.Status.Name,
			Label:   todo.Status.Label,
			Color:   todo.Status.Color,
			BgColor: todo.Status.BgColor,
		},
		Priority: oas.Priority{
			Name:    todo.Priority.Name,
			Label:   todo.Priority.Label,
			Color:   todo.Priority.Color,
			BgColor: todo.Priority.BgColor,
		},
		StartAt: todo.StartAt,
		EndAt:   oas.NewOptDateTime(*todo.EndAt),
	}, nil
}
