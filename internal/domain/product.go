package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          int32
	ProductName string
	Quantity    int32
	Price       string
	Description *string
	UserID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Product) Validate() error {
	if p.ProductName == "" {
		return errors.New("product name is required")
	}
	if p.Quantity <= 0 {
		return errors.New("quantity must be over 0")
	}
	if p.Price == "" {
		return errors.New("product price is required")
	}
	if p.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	return nil
}

func (p *Product) Normalize() {
	p.ProductName = strings.TrimSpace(p.ProductName)
}
