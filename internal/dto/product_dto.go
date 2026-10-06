package dto

import (
	"time"

	"io.github.com/api-product/internal/domain"
)

type CreateProduct struct {
	ProductName string `json:"product_name"`
	Quantity    int32  `json:"quantity"`
	Price       string `json:"price"`
	Description string `json:"description,omitempty"`
}

type CreateProductResponse struct {
	ID          int32     `json:"id"`
	ProductName string    `json:"product_name"`
	Quantity    int32     `json:"quantity"`
	Price       string    `json:"price"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

func ProductResponseToDomain(p domain.Product) CreateProductResponse {
	return CreateProductResponse{
		ID:          p.ID,
		ProductName: p.ProductName,
		Quantity:    p.Quantity,
		Price:       p.Price,
		Description: *p.Description,
		CreatedAt:   p.CreatedAt,
	}
}
