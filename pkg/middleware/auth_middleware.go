package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

// AuthMiddlewareConfig holds configuration for the authentication middleware
type AuthMiddlewareConfig interface {
	ValidateToken(tokenString string) (*jwt.Token, error) // Function to validate JWT tokens
}

type contextKeyUserID string
const KeyUserID contextKeyUserID = "userID"

// AuthMiddleware creates a new authentication middleware
func AuthMiddleware(config AuthMiddlewareConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract the token from the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing authorization header", http.StatusUnauthorized)
				return
			}

			// Extract the token from the "Bearer <token>" format
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				http.Error(w, "Invalid token format", http.StatusUnauthorized)
				return
			}

			// Validate the token using the provided function
			token, err := config.ValidateToken(tokenString)
			if err != nil {
				http.Error(w, "Failed to validate token", http.StatusInternalServerError)
				return
			}
			if token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Missing claims", http.StatusInternalServerError)
			}
	
			userID, ok := claims["user_id"].(string)
			if !ok {
				http.Error(w, "Missing user_id in token", http.StatusInternalServerError)
			}
			ctx := context.WithValue(r.Context(), KeyUserID, userID)

			// Token is valid, proceed to the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
