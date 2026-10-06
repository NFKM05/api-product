package domain

import "time"

type Product struct {
	ID          int32
	ProductName string
	Quantity    int32
	Price       string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
