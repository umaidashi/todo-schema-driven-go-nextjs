package handler

import (
	"context"
	"server/infra/applogger"
	"server/pkg/ent"
	oas "server/pkg/oas"
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
	return nil, nil
}
