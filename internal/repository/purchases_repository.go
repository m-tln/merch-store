package repository

import "merch-store/internal/domain"

type PurchasesRepository interface {
	Create(purchase *domain.Purchase) error
	FindByUserID(id int) ([]domain.Purchase, error)
}
