package dto

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type HeaderData struct {
	Authorization string
}

func (h *HeaderData) Validate() error {
	if h.Authorization == "" {
		return errors.New("authorization header is required")
	}

	parts := strings.SplitN(h.Authorization, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return errors.New("authorization header must be in the format: Bearer <token>")
	}
	return nil
}

type TokenData struct {
	UserID uuid.UUID
	Email  string
}
