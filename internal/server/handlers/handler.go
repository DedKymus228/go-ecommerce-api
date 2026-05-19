package handlers

import (
	"e-commerce-api/internal/service"

	"go.uber.org/zap"
)

type Handler struct {
	orderService   service.OrderService
	cartService    service.Cart
	productService service.Product
	authService    service.Auth
	logger         *zap.Logger
}

func NewHandler(order service.OrderService, cart service.Cart, auth service.Auth, product service.Product, logger *zap.Logger) *Handler {
	return &Handler{
		orderService:   order,
		cartService:    cart,
		productService: product,
		authService:    auth,
		logger:         logger,
	}
}
