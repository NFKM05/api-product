package dto

import (
	"time"

	"github.com/google/uuid"
	"io.github.com/api-product/internal/domain"
)

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	IsActive bool      `json:"is_active"`
}

type UpdateNameRequest struct {
	Name string `json:"name"`
}

type UpdateEmailRequest struct {
	Email string `json:"email"`
}

type UpdatePasswordRequest struct {
	Password    string `json:"current_password"`
	NewPassword string `json:"new_password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"access_token"`
}

type DesactivateUser struct {
	Username   string    `json:"name"`
	InactiveDt time.Time `json:"inactive_dt"`
}

func UserResponseFromDomain(u domain.User) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		IsActive: u.IsActive,
	}
}
