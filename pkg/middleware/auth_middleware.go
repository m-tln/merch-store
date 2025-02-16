package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

type AuthMiddlewareConfig interface {
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type contextKeyUserID string
const KeyUserID contextKeyUserID = "userID"

func AuthMiddleware(config AuthMiddlewareConfig) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing authorization header", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				http.Error(w, "Invalid token format", http.StatusUnauthorized)
				return
			}

			token, err := config.ValidateToken(tokenString)
			if err != nil {
				http.Error(w, "Failed to validate token", http.StatusInternalServerError)
				return
			}
			if !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Missing claims", http.StatusInternalServerError)
			}
	
			userID, ok := claims["user_id"]
			if !ok {
				http.Error(w, "Missing user_id in token", http.StatusInternalServerError)
			}
			ctx := context.WithValue(r.Context(), KeyUserID, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
