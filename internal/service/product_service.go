package service

import (
	"context"

	"io.github.com/api-product/internal/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product string, quantity int32, price string, description *string) (domain.Product, error)
	UpdateProduct(ctx context.Context, id int32, name string, quantity int32, price string, description *string) (domain.Product, error)
	ProductById(ctx context.Context, id int32) (domain.Product, error)
	AllProducts(ctx context.Context) ([]domain.Product, error)
	DeleteProduct(ctx context.Context, id int32) error
}

type ProductService struct {
	repository ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repository: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product string, quantity int32, price string, description *string) (domain.Product, error) {
	return s.repository.Create(ctx, product, quantity, price, description)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int32, name string, quantity int32, price string, description *string) (domain.Product, error) {
	return s.repository.UpdateProduct(ctx, id, name, quantity, price, description)
}

func (s *ProductService) ProductById(ctx context.Context, id int32) (domain.Product, error) {
	return s.repository.ProductById(ctx, id)
}

func (s *ProductService) AllProducts(ctx context.Context) ([]domain.Product, error) {
	return s.repository.AllProducts(ctx)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int32) error {
	return s.repository.DeleteProduct(ctx, id)
}
