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

}

func (h *Handler) AddToCart(c *gin.Context) {
	userIDStr := c.GetHeader(constants.UserIDKey)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		h.logger.Error("failed to parse user id", zap.Error(err))
		return
	}
	var req dto.AddItemToCartRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		h.logger.Error("failed to bind request", zap.Error(err))
		return
	}
	err = h.cartService.AddToCart(c, userID, req.ProductID, req.Quantity)
	if errors.Is(err, apperrors.ErrNotEnoughStock) {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) RemoveFromCart(c *gin.Context) {

}
