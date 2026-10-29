package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"io.github.com/api-product/internal/dto"
	"io.github.com/api-product/internal/utils"
)

type AuthContextKey string

const TokenDataKey AuthContextKey = "tokenData"

func Auth(jwtService *utils.JwtService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var headerData dto.HeaderData

			headerData.Authorization = r.Header.Get("Authorization")

			if err := headerData.Validate(); err != nil {
				utils.NewApiError(w, http.StatusBadRequest, errors.New(""))
				return
			}

			token := strings.TrimPrefix(headerData.Authorization, "Bearer")
			tokenData, err := jwtService.VeririfyJwtToken(token)
			if err != nil {
				utils.NewApiError(w, http.StatusUnauthorized, errors.New(""))
				return
			}

			ctx := context.WithValue(r.Context(), TokenDataKey, tokenData)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetTokenData(r *http.Request) *dto.TokenData {
	tokenData, _ := r.Context().Value(TokenDataKey).(*dto.TokenData)
	return tokenData
}
