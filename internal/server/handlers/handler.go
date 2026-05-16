package handlers

import (
	"e-commerce-api/internal/service"

	"go.uber.org/zap"
)

type Handler struct {
	cartService    service.Cart
	productService service.Product
	authService    service.Auth
	logger         *zap.Logger
}

func NewHandler(cart service.Cart, auth service.Auth, product service.Product, logger *zap.Logger) *Handler {
	return &Handler{
		cartService:    cart,
		productService: product,
		authService:    auth,
		logger:         logger,
	}
}
