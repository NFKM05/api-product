package mapper

import (
	"io.github.com/api-product/internal/db"
	"io.github.com/api-product/internal/domain"
	"io.github.com/api-product/internal/utils"
)

// Converte um db.Product gerado pelo sqlc para domain.Product
// Para não expor o tipo gerado pelo sqlc
func ToDomainProduct(p db.Product) domain.Product {

	return domain.Product{
		ID:          p.ID,
		ProductName: p.ProductName,
		Quantity:    p.Quantity,
		Price:       p.Price,
		Description: utils.FromNullString(p.Description),
		UserID:      p.UserID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt.Time,
	}
}

// Converte um db.User gerado pelo sqlc para domain.Product
// Para não expor o tipo gerado pelo sqlc
func ToDomainUser(u db.User) domain.User {
	return domain.User{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		IsActive:   u.IsActive.Bool,
		InactiveDt: u.InactiveDt.Time,
	}
}
