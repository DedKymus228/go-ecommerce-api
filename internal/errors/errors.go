package apperrors

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

var (
	ErrProductNotFound    = errors.New("product not found")
	ErrNotEnoughStock     = errors.New("not enough stock for the product")
	ErrCartItemNotFound   = errors.New("item not found in cart")
	ErrUpdateItemQuantity = errors.New("failed to update item quantity")
)
