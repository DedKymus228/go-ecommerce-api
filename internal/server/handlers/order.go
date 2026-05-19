package handlers

import "C"
import (
	"e-commerce-api/internal/constants"
	"e-commerce-api/internal/server/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (h *Handler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.logger.Error("failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	userIDAny, exists := c.Get(constants.UserIDKey)
	userID := userIDAny.(uuid.UUID)
	if !exists {
		h.logger.Error("userID not found in context")
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "unauthorized"})
		return
	}
	orderID, err := h.orderService.CreateOrder(c, userID, req.ShippingAddress)
	if err != nil {
		h.logger.Error("failed to create order", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}

	h.logger.Info("order created", zap.String("order_id", orderID.String()))
	c.JSON(http.StatusCreated, dto.OrderResponse{ID: orderID.String()})
	return
}

func (h *Handler) ListHistoryOrder(c *gin.Context) {
	userIDAny, exists := c.Get(constants.UserIDKey)
	if !exists {
		h.logger.Error("userID not found in context")
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "unauthorized"})
		return
	}
	userID := userIDAny.(uuid.UUID)
	orders, err := h.orderService.GetHistoryOrders(c, userID)
	if err != nil {
		h.logger.Error("failed to get history orders", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}

	var resp = make([]dto.OrderResponse, len(orders))
	for i, order := range orders {
		resp[i] = resp[i].ToOrderDTO(order)
	}

}

func (h *Handler) GetOrderByID(c *gin.Context) {
	var req dto.GetOrderRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.logger.Error("failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	userIDAny, exists := c.Get(constants.UserIDKey)
	userID := userIDAny.(uuid.UUID)
	if !exists {
		h.logger.Error("userID not found in context")
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "unauthorized"})
		return
	}

	order, err := h.orderService.GetOrder(c, req.OrderId, userID)
	if err != nil {
		h.logger.Error("failed to get order", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}

	var resp dto.OrderResponse
	resp = resp.ToOrderDTO(order)

	c.JSON(http.StatusOK, resp)
}
