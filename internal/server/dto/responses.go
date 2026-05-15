package dto

import (
	"github.com/google/uuid"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ProductResponse struct {
	ID   uuid.UUID
	Name string
}
