package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"io.github.com/api-product/internal/domain"
)

// Define o contrato de acesso a dados que do ProductService
type ProductRepository interface {
	Create(ctx context.Context, product string, quantity int32, price string, description *string, userID uuid.UUID) (domain.Product, error)
	UpdateProduct(ctx context.Context, id int32, name string, quantity int32, price string, description *string) (domain.Product, error)
	ProductById(ctx context.Context, id int32) (domain.Product, error)
	AllProducts(ctx context.Context) ([]domain.Product, error)
	DeleteProduct(ctx context.Context, id int32) error
}

// Concentra as regras de negócio da aplicação
type ProductService struct {
	repository ProductRepository
}

// Injeta a dependência do repository no service
func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repository: repo}
}

// Função auxiliar que une as duas funções para validação e normalização dos campos
func (s *ProductService) prepare(p *domain.Product) error {
	p.Normalize()
	return p.Validate()
}

// CreateProduct valida e cria um novo produto no banco vinculado a um usuário de criação
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

// UpdateProduct atualiza um produto existente na base de dados
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

// ProductById busca um produto pelo seu identificador
func (s *ProductService) ProductById(ctx context.Context, id int32) (domain.Product, error) {
	return s.repository.ProductById(ctx, id)
}

// AllProducts lista todos os produtos cadastrados na base
func (s *ProductService) AllProducts(ctx context.Context) ([]domain.Product, error) {
	return s.repository.AllProducts(ctx)
}

// DeleteProduct deleta um produto existente na base
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
