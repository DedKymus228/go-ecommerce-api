package handlers

import (
	apperrors "e-commerce-api/internal/errors"
	"e-commerce-api/internal/server/dto"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (u *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}

	userID, err := u.authService.Register(c.Request.Context(), req)

	switch {
	case errors.Is(err, apperrors.ErrUserAlreadyExists):
		c.JSON(409, dto.ErrorResponse{Error: err.Error()})
		u.logger.Error("user already exists", zap.Error(err))
		return
	case err == nil:
		c.JSON(201, dto.RegisterResponse{UserID: userID.String()})
		u.logger.Info("user registered", zap.String("user_id", userID.String()))
		return
	default:
		c.JSON(500, dto.ErrorResponse{Error: err.Error()})
		u.logger.Error("failed to register user", zap.Error(err))
		return
	}
}

func (u *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, dto.ErrorResponse{Error: err.Error()})
		return
	}

	token, err := u.authService.Login(c.Request.Context(), req)

	switch {
	case errors.Is(err, apperrors.ErrInvalidCredentials):
		c.JSON(401, dto.ErrorResponse{Error: err.Error()})
		u.logger.Error("invalid password or email", zap.Error(err))
		return
	case err == nil:
		c.JSON(200, dto.LoginResponse{Token: token})
		u.logger.Info("user logged in")
		return
	default:
		c.JSON(500, dto.ErrorResponse{Error: "internal server error"})
		u.logger.Error("failed to login user", zap.Error(err))
		return
	}
}
