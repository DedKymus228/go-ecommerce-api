package handlers

import (
	"e-commerce-api/internal/service"

	"go.uber.org/zap"
)

type Handler struct {
	authService service.Auth
	logger      *zap.Logger
}


func NewHandler(auth service.Auth, logger *zap.Logger) *Handler {
	return &Handler{
		authService: auth,
		logger:      logger,
	}
}
