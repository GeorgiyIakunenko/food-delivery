package middleware

import (
	"context"
	"food_delivery/config"
	"food_delivery/server/response"
	"food_delivery/service"
	"net/http"
)

func ValidateAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := service.GetTokenFromBearerString(authHeader)

		claims, err := service.ValidateToken(token, config.NewConfig().AccessSecret)
		if err != nil {
			response.SendInvalidCredentials(w)
			return
		}

		// Pass the claims to the next handler
		ctx := context.WithValue(r.Context(), config.NewConfig().AccessSecret, claims)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func ValidateRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := service.GetTokenFromBearerString(authHeader)

		claims, err := service.ValidateToken(token, config.NewConfig().RefreshSecret)
		if err != nil {
			response.SendInvalidCredentials(w)
			return
		}

		// Pass the claims to the next handler
		ctx := context.WithValue(r.Context(), config.NewConfig().RefreshSecret, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
