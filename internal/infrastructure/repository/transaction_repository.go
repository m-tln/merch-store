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
	return r.db.Create(TransactionFromDomainToRepo(transacrion)).Error
}

func (r *TransactionRepositoryImpl) GetTransactionsByIDFrom(id int) ([]domain.Transaction, error) {
	var transactionsRepo []Transaction
	err := r.db.Where("id_from = ?", id).Find(&transactionsRepo).Error
	transactions := make([]domain.Transaction, len(transactionsRepo))
	for _, transaction := range transactionsRepo {
		transactions = append(transactions, *TransactionFromRepoToDomain(&transaction))
	}
	return transactions, err
}

func (r *TransactionRepositoryImpl) GetTransactionsByIDTo(id int) ([]domain.Transaction, error) {
	var transactionsRepo []Transaction
	err := r.db.Where("id_to = ?", id).Find(&transactionsRepo).Error
	transactions := make([]domain.Transaction, len(transactionsRepo))
	for _, transaction := range transactionsRepo {
		transactions = append(transactions, *TransactionFromRepoToDomain(&transaction))
	}
	return transactions, err
}
