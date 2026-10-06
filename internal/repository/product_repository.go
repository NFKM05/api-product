package repository

import (
	"context"
	"database/sql"

	"io.github.com/api-product/internal/db"
	"io.github.com/api-product/internal/domain"
	"io.github.com/api-product/internal/mapper"
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
		Description: sql.NullString{},
	})

	if err != nil {
		return domain.Product{}, err
	}

	return mapper.ToDomain(row), nil
}
