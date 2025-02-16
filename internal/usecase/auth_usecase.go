package usecase

import (
	"fmt"
	"merch-store/internal/domain"
	"merch-store/internal/repository"
	"merch-store/internal/service"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo   repository.UsersRepository
	jwtService *service.JWTService
}

func NewAuthUseCase(userRepo repository.UsersRepository, jwtService *service.JWTService) *AuthUseCase {
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
	if err != nil {
		hash, err := HashPassword(password)
		if err != nil {
			return nil, fmt.Errorf("error in getting hash of password, %v", err)
		}
		user = &domain.User{
			Name:         username,
			PasswordHash: hash,
			Balance:      1000,
		}
		if err := uc.userRepo.Create(user); err != nil {
			return nil, fmt.Errorf("error in creating user, %v", err)
		}
	}

	if !CheckPasswordHash(password, user.PasswordHash) {
		return nil, fmt.Errorf("invalid credentials, %v", err)
	}

	tokenString, err := uc.jwtService.GenerateToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("error in generating token, %v", err)
	}

	return &tokenString, nil
}
