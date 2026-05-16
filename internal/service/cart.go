package service

import (
	"context"
	"e-commerce-api/internal/errors"
	db "e-commerce-api/internal/repository/db/sqlc"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Cart interface {
	GetCart(ctx context.Context, userID uuid.UUID) ([]db.GetCartItemsRow, error)
	AddToCart(ctx context.Context, userID uuid.UUID, productID uuid.UUID, quantity int32) error
	RemoveFromCart(ctx context.Context, userID uuid.UUID, productID uuid.UUID) error
}

type CartService struct {
	repo db.Querier
}

func NewCartService(repo db.Querier) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) AddToCart(ctx context.Context, userID uuid.UUID, productID uuid.UUID, quantity int32) error {
	cart, err := s.getOrCreateCart(ctx, userID)
	if err != nil {
		return err
	}

	product, err := s.repo.GetProductByID(ctx, productID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return apperrors.ErrProductNotFound
		}
		return fmt.Errorf("failed to get product: %w", err)
	}

	if product.StockQuantity < quantity {
		return apperrors.ErrNotEnoughStock
	}

	existingItem, err := s.repo.GetCartItemByProduct(ctx, db.GetCartItemByProductParams{
		CartID:    cart.ID,
		ProductID: productID,
	})

	if err == nil {
		newQuantity := existingItem.Quantity + quantity
		if product.StockQuantity < newQuantity {
			return apperrors.ErrNotEnoughStock
		}
		return s.repo.UpdateCartItemQuantity(ctx, db.UpdateCartItemQuantityParams{
			CartID:    cart.ID,
			ProductID: productID,
			Quantity:  newQuantity,
		})
	}

	if errors.Is(err, pgx.ErrNoRows) {
		params := db.AddCartItemParams{
			CartID:    cart.ID,
			ProductID: productID,
			Quantity:  quantity,
		}
		_, err = s.repo.AddCartItem(ctx, params)
		if err != nil {
			return fmt.Errorf("failed to add item to cart: %w", err)
		}
		return nil
	}

	return fmt.Errorf("failed to check cart item: %w", err)
}

func (s *CartService) getOrCreateCart(ctx context.Context, userID uuid.UUID) (db.Cart, error) {
	cart, err := s.repo.GetCartByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			cart, err = s.repo.CreateCart(ctx, userID)
			if err != nil {
				return db.Cart{}, fmt.Errorf("failed to create cart: %w", err)
			}
			return cart, nil
		}
		return db.Cart{}, fmt.Errorf("failed to get cart: %w", err)
	}
	return cart, nil
}

func (s *CartService) GetCart(ctx context.Context, userID uuid.UUID) ([]db.GetCartItemsRow, error) {
	cart, err := s.repo.GetCartByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart: %w", err)
	}
	items ,err := s.repo.GetCartItems(ctx, cart.ID)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *CartService) RemoveFromCart(ctx context.Context, userID uuid.UUID, productID uuid.UUID) error {
	cart, err := s.repo.GetCartByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get cart: %w", err)
	}
	existingItem, err := s.repo.GetCartItemByProduct(ctx, db.GetCartItemByProductParams{
		CartID:    cart.ID,
		ProductID: productID,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return apperrors.ErrCartItemNotFound
	}

	err = s.repo.DeleteCartItem(ctx, db.DeleteCartItemParams{
		CartID:    existingItem.CartID,
		ProductID: productID,
	})
	if err != nil {
		return fmt.Errorf("failed to remove item from cart: %w", err)
	}
	return nil
}
