package mapper

import (
	"io.github.com/api-product/internal/db"
	"io.github.com/api-product/internal/domain"
	"io.github.com/api-product/internal/utils"
)

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

func ToDomainUser(u db.User) domain.User {
	return domain.User{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		IsActive:   u.IsActive.Valid,
		InactiveDt: u.InactiveDt.Time,
	}
}
