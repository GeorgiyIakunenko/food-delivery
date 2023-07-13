package service

import (
	"errors"
	"food_delivery/config"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type TokenService struct {
	cfg *config.Config
}

type TokenServiceI interface {
	GenerateToken(userID, lifeTimeMinutes int, secret string) (string, error)
	ValidateToken(tokenString, secret string) (*JwtCustomClaims, error)
	GetTokenFromBearerString(bearerString string) string
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func (t *TokenService) GenerateToken(userID, lifeTimeMinutes int, secret string) (string, error) {
	claims := &JwtCustomClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(lifeTimeMinutes))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (t *TokenService) ValidateToken(tokenString, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}

func (t *TokenService) GetTokenFromBearerString(bearerString string) string {
	if bearerString == "" {
		return ""
	}

	parts := strings.Split(bearerString, " ")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}
