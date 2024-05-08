package jwt

import (
	"time"

	"github.com/agusheryanto182/go-inventory-management/config"
	"github.com/golang-jwt/jwt"
)

type JWTInterface interface {
	GenerateJWT(ID, phone_number string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type JWTService struct{}

func NewJWTService() JWTInterface {
	return &JWTService{}
}

var SECRET = config.GetString("JWT_SECRET")

func (s *JWTService) GenerateJWT(ID, phone_number string) (string, error) {
	claims := jwt.MapClaims{
		"id":           ID,
		"phone_number": phone_number,
		"iat":          time.Now().Unix(),
		"exp":          time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
