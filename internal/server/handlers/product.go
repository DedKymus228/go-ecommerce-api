package handlers

import (
	"e-commerce-api/internal/server/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (h *Handler) GetProductByID(c *gin.Context) {
	// 1. Получаем ID из URL
	idStr := c.Param("id")

	// 2. Парсим строку в UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid product ID format"})
		return
	}

	// 3. Вызываем сервис
	product, err := h.productService.GetProductByID(c.Request.Context(), id)
	if err != nil {
		// В будущем здесь нужно будет проверять, что ошибка - это "не найдено" (pgx.ErrNoRows)
		// и возвращать 404. Пока для простоты возвращаем 500.
		h.logger.Error("failed to get product by id", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}

	// 4. Отдаем успешный ответ
	c.JSON(http.StatusOK, product)
}

func (h *Handler) ListProducts(c *gin.Context) {
	products, err := h.productService.ListProducts(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to list products", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, products)
}
