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
} // @name UserRequest

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	IsActive bool      `json:"is_active"`
} // @name UserResponse

type UpdateNameRequest struct {
	Name string `json:"name"`
} // @name UpdateNameRequest

type UpdateEmailRequest struct {
	Email string `json:"email"`
} // @name UpdateEmailRequest

type UpdatePasswordRequest struct {
	Password    string `json:"current_password"`
	NewPassword string `json:"new_password"`
} // @name UpdatePasswordRequest

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
} // @name LoginRequest

type LoginResponse struct {
	Token string `json:"access_token"`
} // @name LoginResponse

type DesactivateUser struct {
	Username   string    `json:"name"`
	InactiveDt time.Time `json:"inactive_dt"`
} // @name DesactivateUser

func UserResponseFromDomain(u domain.User) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		IsActive: u.IsActive,
	}
}
