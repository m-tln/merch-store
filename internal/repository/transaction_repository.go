package repository

import "merch-store/internal/domain"

type TransactionRepository interface {
	Create(*domain.Transaction) error
	GetTransactionsByIDTo(id int) ([]domain.Transaction, error)
	GetTransactionsByIDFrom(id int) ([]domain.Transaction, error)
}