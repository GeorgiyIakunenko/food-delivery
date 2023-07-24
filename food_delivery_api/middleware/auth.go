package middleware

import (
	"context"
	"fmt"
	"food_delivery/config"
	"food_delivery/server/response"
	"food_delivery/service"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type AuthMiddleware struct {
	RedisClient *redis.Client
	cfg         *config.Config
}

func NewAuthMiddleware(redisClient *redis.Client, cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		RedisClient: redisClient,
		cfg:         cfg,
	}
}

func (am *AuthMiddleware) ValidateAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := service.GetTokenFromBearerString(authHeader)

		claims, err := service.ValidateToken(token, am.cfg.AccessSecret)
		if err != nil {
			response.SendTokenExpired(w)
			return
		}

		userID := claims.ID

		accessTokenKey := fmt.Sprintf("access_token:%d", userID)
		ctx := context.Background()
		storedToken, err := am.RedisClient.Get(ctx, accessTokenKey).Result()
		if err != nil {
			response.SendInvalidCredentials(w)
			return
		}

		if token != storedToken {
			response.SendInvalidCredentials(w)
			return
		}

		ctx = context.WithValue(r.Context(), am.cfg.AccessSecret, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (am *AuthMiddleware) ValidateRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := service.GetTokenFromBearerString(authHeader)

		claims, err := service.ValidateToken(token, am.cfg.RefreshSecret)
		if err != nil {
			response.SendTokenExpired(w)
			return
		}

		refreshTokenKey := fmt.Sprintf("refresh_token:%d", claims.ID)
		ctx := context.Background()
		storedToken, err := am.RedisClient.Get(ctx, refreshTokenKey).Result()
		if err != nil {
			response.SendInvalidCredentials(w)
			return
		}

		if token != storedToken {
			response.SendInvalidCredentials(w)
			return
		}

		ctx = context.WithValue(r.Context(), am.cfg.RefreshSecret, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
