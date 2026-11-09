package dto

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// Struct que representa o Authorization recebido na requisição usado pelo Middleware
type HeaderData struct {
	Authorization string
}

// Função auxiliar o Middleware que valida o header da requisição
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

// Struct que contém os dados extraídos de um JWT já validado
type TokenData struct {
	UserID uuid.UUID
	Email  string
}
