package mapper

import (
	"io.github.com/api-product/internal/db"
	"io.github.com/api-product/internal/domain"
	"io.github.com/api-product/internal/utils"
)

func ToDomain(p db.Product) domain.Product {

	return domain.Product{
		ID:          p.ID,
		ProductName: p.ProductName,
		Quantity:    p.Quantity,
		Price:       p.Price,
		Description: utils.FromNullString(p.Description),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt.Time,
	}
}
