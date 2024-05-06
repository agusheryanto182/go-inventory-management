package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTInterface interface {
	GenerateJWT(ID, email string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) JWTInterface {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) GenerateJWT(ID, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    ID,
		"email": email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (j *JWT) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
