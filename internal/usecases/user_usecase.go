package usecase

import (
    "merch-store/internal/domain"
	"merch-store/internal/repository"
)

type UserUseCase struct {
    userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
    return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) CreateUser(name, passwordHash string, balance uint64) error {
    user := &domain.User{
        Name:         name,
        PasswordHash: passwordHash,
        Balance:      balance,
    }
    return uc.userRepo.Create(user)
}

func (uc *UserUseCase) GetUserByID(id int) (*domain.User, error) {
    return uc.userRepo.FindByID(id)
}

func (uc *UserUseCase) UpdateUserBalance(id int, balance int) error {
    return uc.userRepo.UpdateBalance(id, balance)
}