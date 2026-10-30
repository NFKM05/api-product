package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"io.github.com/api-product/internal/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product string, quantity int32, price string, description *string, userID uuid.UUID) (domain.Product, error)
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

func (s *ProductService) prepare(p *domain.Product) error {
	p.Normalize()
	return p.Validate()
}

func (s *ProductService) CreateProduct(ctx context.Context, product string, quantity int32, price string, description *string, userID uuid.UUID) (domain.Product, error) {
	prod := domain.Product{
		ProductName: product,
		Quantity:    quantity,
		Price:       price,
		Description: description,
		UserID:      userID,
	}
	if err := s.prepare(&prod); err != nil {
		return domain.Product{}, err
	}
	return s.repository.Create(ctx, product, quantity, price, description, userID)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int32, name string, quantity int32, price string, description *string, userID uuid.UUID) (domain.Product, error) {

	product, err := s.repository.ProductById(ctx, id)

	if err != nil {
		return domain.Product{}, errors.New("product not found")
	}

	if product.UserID != userID {
		return domain.Product{}, errors.New("the user provided is not the same as the registered user.")
	}

	prod := domain.Product{
		ProductName: name,
		Quantity:    quantity,
		Price:       price,
		Description: description,
		UserID:      userID,
	}

	if err := s.prepare(&prod); err != nil {
		return domain.Product{}, err
	}
	return s.repository.UpdateProduct(ctx, id, name, quantity, price, description)
}

func (s *ProductService) ProductById(ctx context.Context, id int32) (domain.Product, error) {
	return s.repository.ProductById(ctx, id)
}

func (s *ProductService) AllProducts(ctx context.Context) ([]domain.Product, error) {
	return s.repository.AllProducts(ctx)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int32, userID uuid.UUID) error {
	product, err := s.repository.ProductById(ctx, id)

	if err != nil {
		return errors.New("product not found")
	}

	if product.UserID != userID {
		return errors.New("the user provided is not the same as the registered user.")
	}

	return s.repository.DeleteProduct(ctx, id)
}
