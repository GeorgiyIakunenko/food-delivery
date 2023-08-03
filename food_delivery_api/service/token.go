package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	ID  int    `json:"id"`
	UID string `json:"uid"`
	jwt.RegisteredClaims
}

type TokenMapping struct {
	AccessUid  string `json:"access_uid"`
	RefreshUid string `json:"refresh_uid"`
}

func GenerateToken(userID, lifeTimeMinutes int, secret string) (string, string, error) {

	// creating a uuid using google uuid package

	uid := uuid.New().String()

	claims := &JwtCustomClaims{
		userID,
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(lifeTimeMinutes))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	return tokenString, uid, nil
}

func ValidateToken(tokenString, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
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
