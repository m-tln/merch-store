package usecase

import (
	"errors"
	"fmt"
	"merch-store/internal/domain"
	"merch-store/internal/repository"
)

type SendCoinUseCase struct {
	userRepo        repository.UserRepository
	transactionRepo repository.TransactionRepository
}

func NewSendCoinUseCase(userRepo repository.UserRepository,
	transactionRepo repository.TransactionRepository) *SendCoinUseCase {
	return &SendCoinUseCase{
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}

const SmallBalanceToSend string = "not enough coins to send"

func (uc *SendCoinUseCase) MakeTransaction(from string, to string, amount int32) error {
	userFrom, err := uc.userRepo.FindByUsername(from)
	if err != nil {
		return fmt.Errorf("user with username %s can't be found in db, error: %v", from, err)
	}

	if userFrom.Balance < uint64(amount) {
		return errors.New(SmallBalanceToSend)
	}

	userTo, err := uc.userRepo.FindByUsername(to)
	if err != nil {
		return fmt.Errorf("user with username %s can't be found in db, error: %v", to, err)
	}

	err = uc.userRepo.UpdateBalance(userFrom.ID, int(userFrom.Balance) - int(amount))
	if err != nil {
		return fmt.Errorf("balance of %s can't be updated, error: %v", from, err)
	}

	err = uc.userRepo.UpdateBalance(userTo.ID, int(userTo.Balance) + int(amount))
	if err != nil {
		return fmt.Errorf("balance of %s can't be updated, error: %v", to, err)
	}

	err = uc.transactionRepo.Create(&domain.Transaction{
		IDFrom: uint64(userFrom.ID),
		IDTo: uint64(userTo.ID),
		Volume: uint64(amount),
	})

	if err != nil {
		return fmt.Errorf("can't create new transaction")
	}

	return nil
}
