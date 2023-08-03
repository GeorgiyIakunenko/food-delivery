package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type JwtCustomClaims2 struct {
	ID int `json:"id"`
}

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, lifeTimeMinutes int, secret string) (string, error) {
	claims := &JwtCustomClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(lifeTimeMinutes))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {

		fmt.Println("token, err", err)

		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {

		fmt.Println("claim, ok", err)

		return nil, err
	}

	return claims, nil
}

func GetTokenFromBearerString(bearerString string) string {
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
