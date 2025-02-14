package repository

import (
	"merch-store/internal/domain"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepositoryImpl(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) Create(transacrion *domain.Transaction) error {
	return r.db.Create(transacrion).Error
}

func (r *TransactionRepositoryImpl) GetTransactionsByID(id int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Where("id_from = ?", id).Find(&transactions).Error
	return transactions, err
}