package dto

import (
	db "e-commerce-api/internal/repository/db/sqlc"
	"fmt"
	"time"

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

type OrderResponse struct {
	ID              string `json:"id"`
	UserID          string `json:"user_id"`
	StatusID        *int32 `json:"status_id"`
	TotalAmount     string `json:"total_amount"`
	ShippingAddress string `json:"shipping_address"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (o *OrderResponse) ToOrderDTO(dbOrder db.Order) OrderResponse {
	dto := OrderResponse{
		ID:              dbOrder.ID.String(),
		UserID:          dbOrder.UserID.String(),
		ShippingAddress: dbOrder.ShippingAddress,
	}

	if dbOrder.StatusID.Valid {
		val := dbOrder.StatusID.Int32
		dto.StatusID = &val
	}

	if dbOrder.TotalAmount.Valid {

		if val, err := dbOrder.TotalAmount.Value(); err == nil && val != nil {
			dto.TotalAmount = fmt.Sprintf("%v", val)
		} else {
			dto.TotalAmount = "0.00"
		}
	} else {
		dto.TotalAmount = "0.00"
	}

	if dbOrder.CreatedAt.Valid {
		dto.CreatedAt = dbOrder.CreatedAt.Time.Format(time.RFC3339) // "2026-05-19T22:12:24Z"
	}
	if dbOrder.UpdatedAt.Valid {
		dto.UpdatedAt = dbOrder.UpdatedAt.Time.Format(time.RFC3339)
	}

	return dto
}
