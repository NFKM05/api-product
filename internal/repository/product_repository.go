package repository

import (
	"context"

	"io.github.com/api-product/internal/db"
	"io.github.com/api-product/internal/domain"
	"io.github.com/api-product/internal/mapper"
	"io.github.com/api-product/internal/utils"
)

type ProductRepository struct {
	queries *db.Queries
}

func NewProductRepository(queries *db.Queries) *ProductRepository {
	return &ProductRepository{queries: queries}
}

func (r *ProductRepository) Create(ctx context.Context, product string, quantity int32,
	price, description string) (domain.Product, error) {
	row, err := r.queries.CreateProduct(ctx, db.CreateProductParams{
		ProductName: product,
		Quantity:    quantity,
		Price:       price,
		Description: utils.ToNullString(&description),
	})

	if err != nil {
		return domain.Product{}, err
	}

	return mapper.ToDomain(row), nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id int32, name string,
	quantity int32, price string, description *string) (domain.Product, error) {
	row, err := r.queries.UpdateProduct(ctx, db.UpdateProductParams{
		ID:          id,
		ProductName: name,
		Quantity:    quantity,
		Price:       price,
		Description: utils.ToNullString(description),
	})

	if err != nil {
		return domain.Product{}, err
	}

	return mapper.ToDomain(row), nil
}

func (r *ProductRepository) ProductById(ctx context.Context, id int32) (domain.Product, error) {
	row, err := r.queries.ProductById(ctx, id)

	if err != nil {
		return domain.Product{}, err
	}

	return mapper.ToDomain(row), nil
}

func (r *ProductRepository) AllProducts(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.queries.ListAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := make([]domain.Product, 0, len(rows))
	for _, row := range rows {
		products = append(products, mapper.ToDomain(row))
	}
	return products, nil
}
