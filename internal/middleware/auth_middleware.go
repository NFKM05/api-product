package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"io.github.com/api-product/internal/dto"
	"io.github.com/api-product/internal/utils"
)

// AuthContextKey é o tipo utilizado como chave context.Context para guardar dados de autenticação
type AuthContextKey string

// TokenDataKey é a chave usada para guardar e recuperar o token do context de um requisição autenticada
const TokenDataKey AuthContextKey = "tokenData"

// Devolve um Middleware que exige um token JWT válido pra acessar a rota
func Auth(jwtService *utils.JwtService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var headerData dto.HeaderData

			// Confere o formato do Header
			headerData.Authorization = r.Header.Get("Authorization")

			if err := headerData.Validate(); err != nil {
				utils.NewApiError(w, http.StatusUnauthorized, errors.New("authorization token is required for access resource"))
				return
			}

			// Verifica a assinatura e validade do token e se estiver correto guarda os dados do usuário
			token := strings.TrimPrefix(headerData.Authorization, "Bearer ")
			tokenData, err := jwtService.VeririfyJwtToken(token)
			if err != nil {
				utils.NewApiError(w, http.StatusUnauthorized, errors.New("unauthorized token"))
				return
			}

			ctx := context.WithValue(r.Context(), TokenDataKey, tokenData)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Recupera os dados do usuário autenticado a partir do context da requisição
func GetTokenData(r *http.Request) *dto.TokenData {
	tokenData, _ := r.Context().Value(TokenDataKey).(*dto.TokenData)
	return tokenData
}
