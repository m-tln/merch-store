package repository

import "merch-store/internal/domain"

type TransactionRepository interface {
	Create(*domain.Transaction) error
	GetTransactionsByID(id int) ([]domain.Transaction, error)
}