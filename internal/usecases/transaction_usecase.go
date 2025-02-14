package usecase

import (
	"merch-store/internal/domain"
	"merch-store/internal/repository"
)


type TransactionUseCase struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionUseCase(transactionRepo repository.TransactionRepository) (*TransactionUseCase) {
	return &TransactionUseCase{transactionRepo: transactionRepo}
}

func (uc *TransactionUseCase) Create(transacrion *domain.Transaction) error {
	return uc.transactionRepo.Create(transacrion)
}

func (uc *TransactionUseCase) GetTransactionsByID(id int) ([]domain.Transaction, error) {
	return uc.transactionRepo.GetTransactionsByID(id)
}