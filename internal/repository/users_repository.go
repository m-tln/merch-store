package repository

import "merch-store/internal/domain"

type UsersRepository interface {
	Create(*domain.User) error
	FindByID(int) (*domain.User, error)
	UpdateBalance(int, int) error
	FindByUsername(string) (*domain.User, error)
}
