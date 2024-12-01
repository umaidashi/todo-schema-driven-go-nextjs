package handler

import (
	"context"
	"server/domain/model"
	"server/infra/applogger"
	"server/infra/dao"
	"server/pkg/ent"
	oas "server/pkg/oas"

	"github.com/samber/lo"
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

	todoDao := dao.NewTodoDao(h.db)
	todos, err := todoDao.List(ctx)
	if err != nil {
		return &oas.GetTodoInternalServerError{
			ErrorCode:    oas.NewOptFloat64(500),
			ErrorMessage: oas.NewOptString("Internal Server Error"),
		}, err
	}

	todoInfomations := lo.Map(todos, func(todo model.Todo, _ int) oas.TodoInformation {
		return oas.TodoInformation{
			Title:  todo.Title,
			Detail: oas.NewOptString(todo.Description),
		}
	})

	return (*oas.GetTodoOKApplicationJSON)(&todoInfomations), nil
}

func (h *Handler) UpdateTodo(ctx context.Context, req oas.OptTodoInformation, params oas.UpdateTodoParams) (oas.UpdateTodoRes, error) {
	applogger.Info("UpdateTodo")
	return nil, nil
}
