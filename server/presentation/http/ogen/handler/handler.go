package handler

import (
	"context"
	"server/infra/applogger"
	"server/pkg/ent"
	oas "server/pkg/oas"
	"time"
)

func NewHandler(db *ent.Client) oas.Handler {
	return &Handler{
		db: db,
	}
}

type Handler struct {
	db *ent.Client
}

func (h *Handler) CreateTodo(ctx context.Context, req oas.OptTodoInformation) (oas.CreateTodoRes, error) {
	applogger.Info("CreateTodo")
	return nil, nil
}

func (h *Handler) DeleteTodo(ctx context.Context, params oas.DeleteTodoParams) (oas.DeleteTodoRes, error) {
	applogger.Info("DeleteTodo")
	return nil, nil
}

func (h *Handler) GetTodo(ctx context.Context, params oas.GetTodoParams) (oas.GetTodoRes, error) {
	applogger.Info("GetTodo")

	todoInfomation := oas.TodoInformation{
		Title:     "title",
		Detail:    oas.NewOptString("description"),
		Progress:  20,
		StartDate: oas.NewOptDate(time.Now()),
		EndDate:   oas.NewOptDate(time.Now()),
	}
	return &oas.GetTodoOKApplicationJSON{
		todoInfomation,
	}, nil
}

func (h *Handler) UpdateTodo(ctx context.Context, req oas.OptTodoInformation, params oas.UpdateTodoParams) (oas.UpdateTodoRes, error) {
	applogger.Info("UpdateTodo")
	return nil, nil
}
