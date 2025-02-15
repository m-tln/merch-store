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

func (uc *UserUseCase) Create(user *domain.User) error {
    return uc.userRepo.Create(user)
}

func (uc *UserUseCase) FindByID(id int) (*domain.User, error) {
    return uc.userRepo.FindByID(id)
}

func (uc *UserUseCase) UpdateBalance(id int, balance int) error {
    return uc.userRepo.UpdateBalance(id, balance)
}

func (uc *UserUseCase) FindByUsername(username string) (*domain.User, error) {
    return uc.userRepo.FindByUsername(username)
}