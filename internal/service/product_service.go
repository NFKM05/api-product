package service

import (
	"context"

	"io.github.com/api-product/internal/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product string, quantity int32, price, description string) (domain.Product, error)
}

type ProductService struct {
	repository ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repository: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product string, quantity int32, price, description string) (domain.Product, error) {
	return s.repository.Create(ctx, product, quantity, price, description)
}
