package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"io.github.com/api-product/internal/dto"
)

const secretJwtKey = "supersecretkey"

type JwtService struct {
	secretKey string
}

func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: secretJwtKey,
	}
}

func (j *JwtService) GenerateJwtToken(email string, userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(j.secretKey))
}

func (j *JwtService) VeririfyJwtToken(token string) (*dto.TokenData, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
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

	userIDstr, ok := claims["userID"].(string)
	if !ok {
		return nil, errors.New("invalid userID claims")
	}

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
