package handlers

import (
	"e-commerce-api/internal/service"

	"go.uber.org/zap"
)

type Handler struct {
	productService service.Product
	authService    service.Auth
	logger         *zap.Logger
}

func NewHandler(auth service.Auth, product service.Product, logger *zap.Logger) *Handler {
	return &Handler{
		productService: product,
		authService:    auth,
		logger:         logger,
	}
}
