package repository

import "merch-store/internal/domain"

type UserRepository interface {
	Create(*domain.User) error
	FindByID(int) (*domain.User, error)
	UpdateBalance(int, int) (error)
}