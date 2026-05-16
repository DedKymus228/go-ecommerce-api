package handlers

import (
	"e-commerce-api/internal/constants"
	apperrors "e-commerce-api/internal/errors"
	"e-commerce-api/internal/server/dto"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (h *Handler) GetCart(c *gin.Context) {
	userIDAny, exists := c.Get(constants.UserIDKey)
	if !exists {
		h.logger.Error("userID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDAny.(uuid.UUID)
	if !ok {
		h.logger.Error("invalid userID type in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	item, err := h.cartService.GetCart(c, userID)
	if err != nil {
		h.logger.Error("failed to get cart", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) AddToCart(c *gin.Context) {
	userIDAny, exists := c.Get(constants.UserIDKey)
	if !exists {
		h.logger.Error("userID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDAny.(uuid.UUID)
	if !ok {
		h.logger.Error("invalid userID type in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	var req dto.AddItemToCartRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.logger.Error("failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	err = h.cartService.AddToCart(c, userID, req.ProductID, req.Quantity)

	switch {
	case errors.Is(err, apperrors.ErrNotEnoughStock):
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	case errors.Is(err, apperrors.ErrProductNotFound):
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: err.Error()})
		return
	case err != nil:
		h.logger.Error("failed to add to cart", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	default:
		c.JSON(http.StatusCreated, gin.H{})
	}
}

func (h *Handler) RemoveFromCart(c *gin.Context) {
	userIDAny, exists := c.Get(constants.UserIDKey)
	if !exists {
		h.logger.Error("userID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := userIDAny.(uuid.UUID)
	if !ok {
		h.logger.Error("invalid userID type in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	productIDStr := c.Param("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		h.logger.Error("failed to parse product ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid product id format"})
		return
	}

	err = h.cartService.RemoveFromCart(c, userID, productID)
	switch {
	case err == nil:
		c.JSON(http.StatusNoContent, gin.H{})
		return
	case errors.Is(err, apperrors.ErrCartItemNotFound):
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: err.Error()})
		return
	default:
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		h.logger.Error("failed to remove item from cart", zap.Error(err))
		return
	}
}
