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
	return r.db.Create(purchase).Error
}

func (r *PurchaseRepositoryImpl) FindByUserID(userID int) ([]domain.Purchase, error) {
    var purchases []domain.Purchase
    err := r.db.Where("id_costumer = ?", userID).Find(&purchases).Error
    return purchases, err
}
