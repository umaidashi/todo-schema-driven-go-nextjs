package handler

import (
	"context"
	"server/infra/applogger"
	oas "server/presentation/http/ogen/oas"
)

func NewHandler() oas.Handler {
	return &Handler{}
}

type Handler struct {
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
	return nil, nil
}

func (h *Handler) UpdateTodo(ctx context.Context, req oas.OptTodoInformation, params oas.UpdateTodoParams) (oas.UpdateTodoRes, error) {
	applogger.Info("UpdateTodo")
	return nil, nil
}
