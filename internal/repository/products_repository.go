package repository

import "merch-store/internal/domain"

type ProductsRepository interface {
	FindByID(id int) (*domain.Product, error)
	FindByName(name string) (*domain.Product, error)
}
