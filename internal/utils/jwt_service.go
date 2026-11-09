package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"io.github.com/api-product/internal/dto"
)

// Chave utilizada para assinar o token JWT
const secretJwtKey = "supersecretkey"

// Struct que concentra geração e verificação de tokens JWT usados para autentição
type JwtService struct {
	secretKey string
}

// Função que cria um JwtService com a chave de assinatura configurada
func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: secretJwtKey,
	}
}

// Gera um token assinado contendo informações do email e id do usuário válido por 2 horas
func (j *JwtService) GenerateJwtToken(email string, userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(j.secretKey))
}

// Valida a assinatura e a expiração de um token JWT e extrai os dados do usuário autenticado
func (j *JwtService) VeririfyJwtToken(token string) (*dto.TokenData, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		// Verifica se o token foi assinado com HMAC e rejeita caso for diferente deste algoritmo
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, errors.New("could not parsed token")
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// Todo claim JWT é serializado e lido como string por isso o casting para string não
	// para UUID diretamente
	userIDstr, ok := claims["userID"].(string)
	if !ok {
		return nil, errors.New("invalid userID claims")
	}

	// Parse na sequência para reconstruir o tipo original
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		return nil, errors.New("invalid userID format")
	}

	email, _ := claims["email"].(string)

	return &dto.TokenData{
		UserID: userID,
		Email:  email,
	}, nil
}
