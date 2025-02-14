package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
    secretKey string
}

func NewJWTService(secretKey string) *JWTService {
    return &JWTService{secretKey: secretKey}
}

func (s *JWTService) GenerateToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.secretKey))
}

func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(s.secretKey), nil
    })
}

func (s *JWTService) GetUsername(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.secretKey), nil
	})

    if err != nil {
		return "", fmt.Errorf("jwt failed with error: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if username, ok := claims["username"].(string); ok {
            return username, nil
        }
		return "", fmt.Errorf("no username in jwt")
	} else {
		return "", fmt.Errorf("invalid token claims")
	}
}