package service

import (
	"context"
	db "e-commerce-api/internal/repository/db/sqlc"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Order interface {
	CreateOrder(ctx context.Context, userID uuid.UUID, shippingAddress string) (uuid.UUID, error)
	GetHistoryOrders(ctx context.Context, userID uuid.UUID) ([]db.Order, error)
	GetOrder(ctx context.Context, orderID uuid.UUID, userID uuid.UUID) (db.Order, error)
	// TODO: PayOrder
}

type OrderService struct {
	repo db.Querier
}

func NewOrderService(repo db.Querier) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, userID uuid.UUID, shippingAddress string) (uuid.UUID, error) {
	createStatus, err := o.repo.GetOrderStatusByName(ctx, "created")
	if err != nil {
		err = fmt.Errorf("failed to get order status: %w", err)
		return uuid.Nil, err
	}
	arg := db.CreateOrderParams{
		UserID:          userID,
		StatusID:        pgtype.Int4{Int32: createStatus.ID, Valid: true},
		ShippingAddress: shippingAddress,
	}

	order, err := o.repo.CreateOrder(ctx, arg)
	if err != nil {
		err = fmt.Errorf("failed to create order: %w", err)
		return uuid.Nil, err
	}
	return order.ID, nil
}

func (o *OrderService) GetHistoryOrders(ctx context.Context, userID uuid.UUID) ([]db.Order, error) {
	orders, err := o.repo.ListUserOrders(ctx, userID)
	if err != nil {
		err = fmt.Errorf("failed to get orders: %w", err)
		return nil, err
	}
	return orders, nil
}

func (o *OrderService) GetOrder(ctx context.Context, orderID uuid.UUID, userID uuid.UUID) (db.Order, error) {
	var emptyOrder db.Order

	order, err := o.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return emptyOrder, err
	}
	if order.UserID != userID {
		err = fmt.Errorf("access to order not allowed")
		return emptyOrder, err
	}

	return order, nil
}
