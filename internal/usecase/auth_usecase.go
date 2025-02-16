package usecase

import (
	"fmt"
	"merch-store/internal/domain"
	"merch-store/internal/repository"
	"merch-store/internal/service"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo   repository.UserRepository
	jwtService *service.JWTService
}

func NewAuthUseCase(userRepo repository.UserRepository, jwtService *service.JWTService) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo, jwtService: jwtService}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (uc *AuthUseCase) GetToken(username, password string) (*string, error) {
	user, err := uc.userRepo.FindByUsername(username)
	fmt.Println("GetToken    username: ", username, " password: ", password)
	if err != nil {
		// Automatically create user if not found
		fmt.Println("Error: ", err)
		hash, err := HashPassword(password)
		if err != nil {
			return nil, fmt.Errorf("error in getting hash of password, %v", err)
		}
		user = &domain.User{
			Name:         username,
			PasswordHash: hash,
			Balance: 1000,
		}
		if err := uc.userRepo.Create(user); err != nil {
			return nil, fmt.Errorf("error in creating user, %v", err)
		}
	}
	fmt.Println("GetToken:   ", user.Name, "   ", user.PasswordHash)

	// Validate password (in a real app, compare hashed passwords)
	if !CheckPasswordHash(password, user.PasswordHash) {
		return nil, fmt.Errorf("invalid credentials, %v", err)
	}

	// Generate JWT token
	tokenString, err := uc.jwtService.GenerateToken(user.ID)
	fmt.Print(tokenString)
	if err != nil {
		return nil, fmt.Errorf("error in generating token, %v", err)
	}

	return &tokenString, nil
}
