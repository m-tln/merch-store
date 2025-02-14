package middleware

import (
    "net/http"
    "strings"
    "merch-store/internal/service"
)

func AuthMiddleware(jwtService *service.JWTService, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing authorization header", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwtService.ValidateToken(tokenString)
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Proceed to the next handler
        next.ServeHTTP(w, r)
    })
}