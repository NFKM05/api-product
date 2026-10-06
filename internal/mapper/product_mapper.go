package mapper

import (
	"io.github.com/api-product/internal/db"
	"io.github.com/api-product/internal/domain"
)

func ToDomain(p db.Product) domain.Product {

	return domain.Product{
		ID:          int64(p.ID),
		ProductName: p.ProductName,
		Quantity:    p.Quantity,
		Price:       p.Price,
		Description: &p.Description.String,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt.Time,
	}
}
