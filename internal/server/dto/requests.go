package dto

import "github.com/google/uuid"

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ProductByIDRequest struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

type AddItemToCartRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int32     `json:"quantity" binding:"required,gt=0"`
}

type DeleteItemFromCartRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
}

type CreateOrderRequest struct {
	ShippingAddress string `json:"shipping_address" binding:"required"`
}

type GetOrderRequest struct {
	OrderId uuid.UUID `json:"order_id" binding:"required"`
}
