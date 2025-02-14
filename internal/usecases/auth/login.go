package auth

import (
	"fmt"
	"merch-store/internal/domain"
	"merch-store/internal/repository"
	"merch-store/internal/service"

	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	userRepo   repository.UserRepository
	jwtService *service.JWTService
}

func NewLoginUseCase(userRepo repository.UserRepository, jwtService *service.JWTService) *LoginUseCase {
	return &LoginUseCase{userRepo: userRepo, jwtService: jwtService}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (uc *LoginUseCase) Execute(username, password string) (*domain.Token, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		// Automatically create user if not found
		hash, err := HashPassword(password)
		if err != nil {
			return nil, fmt.Errorf("error in getting hash of password, %v", err)
		}
		user = &domain.User{
			Name:         username,
			PasswordHash: hash,
		}
		if err := uc.userRepo.Create(user); err != nil {
			return nil, fmt.Errorf("error in creating user, %v", err)
		}
	}

	// Validate password (in a real app, compare hashed passwords)
	if !CheckPasswordHash(password, user.PasswordHash) {
		return nil, fmt.Errorf("invalid credentials, %v", err)
	}

	// Generate JWT token
	tokenString, err := uc.jwtService.GenerateToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("error in generating token, %v", err)
	}

	return &domain.Token{AccessToken: tokenString}, nil
}
