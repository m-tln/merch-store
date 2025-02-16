package repository

import (
	"merch-store/internal/domain"

	"gorm.io/gorm"
)

type PurchaseRepositoryImpl struct {
	db *gorm.DB
}

func NewPurchaseRepositoryImpl(db *gorm.DB) *PurchaseRepositoryImpl {
	return &PurchaseRepositoryImpl{db: db}
}

func (r *PurchaseRepositoryImpl) Create(purchase *domain.Purchase) error {
	return r.db.Create(PurchaseFromDomainToRepo(purchase)).Error
}

func (r *PurchaseRepositoryImpl) FindByUserID(userID int) ([]domain.Purchase, error) {
	var purchasesRepo []Purchase
	err := r.db.Where("id_costumer = ?", userID).Find(&purchasesRepo).Error
	purchases := make([]domain.Purchase, len(purchasesRepo))
	for _, purchase := range purchasesRepo {
		purchases = append(purchases, *PurchaseFromRepoToDomain(&purchase))
	}
	return purchases, err
}
