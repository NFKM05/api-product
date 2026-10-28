package dto

import (
	"time"

	"github.com/google/uuid"
	"io.github.com/api-product/internal/domain"
)

type CreateProduct struct {
	ProductName string    `json:"product_name"`
	Quantity    int32     `json:"quantity"`
	Price       string    `json:"price"`
	Description string    `json:"description,omitempty"`
	UserID      uuid.UUID `json:"user_id"`
}

type CreateProductResponse struct {
	ID          int32     `json:"id"`
	ProductName string    `json:"product_name"`
	Quantity    int32     `json:"quantity"`
	Price       string    `json:"price"`
	Description *string   `json:"description,omitempty"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateProduct struct {
	ProductName string `json:"product_name"`
	Quantity    int32  `json:"quantity"`
	Price       string `json:"price"`
	Description string `json:"description,omitempty"`
}

func ProductResponseFromDomain(p domain.Product) CreateProductResponse {
	return CreateProductResponse{
		ID:          p.ID,
		ProductName: p.ProductName,
		Quantity:    p.Quantity,
		Price:       p.Price,
		Description: p.Description,
		UserID:      p.UserID,
		CreatedAt:   p.CreatedAt,
	}
}
