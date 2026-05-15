package service

import (
	"context"
	db "e-commerce-api/internal/repository/db/sqlc"

	"github.com/google/uuid"
)

// Product описывает бизнес-логику для работы с товарами
type Product interface {
	GetProductByID(ctx context.Context, id uuid.UUID) (db.Product, error)
	ListProducts(ctx context.Context) ([]db.Product, error)
}

// ProductService - реализация интерфейса Product
type ProductService struct {
	repo db.Querier
}

func NewProductService(repo db.Querier) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetProductByID(ctx context.Context, id uuid.UUID) (db.Product, error) {
	// На этом этапе просто вызываем репозиторий.
	// В будущем здесь может быть логика кеширования или проверки прав.
	return s.repo.GetProductByID(ctx, id)
}

func (s *ProductService) ListProducts(ctx context.Context) ([]db.Product, error) {
	return s.repo.ListProducts(ctx)
}
