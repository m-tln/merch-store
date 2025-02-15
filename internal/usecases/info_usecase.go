package usecase

import (
	"fmt"
	"merch-store/internal/repository"
)

type InfoUseCase struct {
	userRepo        repository.UserRepository
	goodsRepo       repository.GoodsRepository
	transactionRepo repository.TransactionRepository
	purchaseRepo    repository.PurchaseRepository
}

func NewInfoUseCase(userRepo repository.UserRepository, goodsRepo repository.GoodsRepository,
	transactionRepo repository.TransactionRepository, purchaseRepo repository.PurchaseRepository) *InfoUseCase {
	return &InfoUseCase{
		userRepo:        userRepo,
		goodsRepo:       goodsRepo,
		transactionRepo: transactionRepo,
		purchaseRepo:    purchaseRepo,
	}
}

func (uc *InfoUseCase) GetInvetory(username string) (map[string]uint, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return map[string]uint{}, fmt.Errorf("user with username %s can't be found in db, error: %v", username, err)
	}

	purchases, err := uc.purchaseRepo.FindByUserID(user.ID)
	if err != nil {
		return map[string]uint{}, fmt.Errorf("purchases of %s can't be found, error: %v", username, err)
	}

	res := make(map[string]uint)
	for _, purchase := range purchases {
		item, err := uc.goodsRepo.FindByID(int(purchase.IDItem))
		if err != nil {
			return map[string]uint{}, fmt.Errorf("%v can't be found in db, error: %v", purchase.IDItem, err)
		}
		res[item.Name] += uint(purchase.Volume)
	}

	return res, nil
}


func (uc *InfoUseCase) GetRecieved(username string) (map[string][]uint, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return map[string][]uint{}, fmt.Errorf("user with username %s can't be found in db, error: %v", username, err)
	}

	transactions, err := uc.transactionRepo.GetTransactionsByIDTo(user.ID)
	if err != nil {
		return map[string][]uint{}, fmt.Errorf("transactions of %s can't be found, error: %v", username, err)
	}

	res := make(map[string][]uint)
	for _, transaction := range transactions {
		userFrom, err := uc.userRepo.FindByID(int(transaction.IDFrom))
		if err != nil {
			return map[string][]uint{}, fmt.Errorf("user with username %v can't be found in db, error: %v", transaction.IDFrom, err)
		}
		res[userFrom.Name] = append(res[userFrom.Name], uint(transaction.Volume))
	}
	return res, nil
}

func (uc *InfoUseCase) GetSent(username string) (map[string][]uint, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return map[string][]uint{}, fmt.Errorf("user with username %s can't be found in db, error: %v", username, err)
	}

	transactions, err := uc.transactionRepo.GetTransactionsByIDFrom(user.ID)
	if err != nil {
		return map[string][]uint{}, fmt.Errorf("transactions of %s can't be found, error: %v", username, err)
	}

	res := make(map[string][]uint)
	for _, transaction := range transactions {
		userTo, err := uc.userRepo.FindByID(int(transaction.IDTo))
		if err != nil {
			return map[string][]uint{}, fmt.Errorf("user with username %v can't be found in db, error: %v", transaction.IDTo, err)
		}
		res[userTo.Name] = append(res[userTo.Name], uint(transaction.Volume))
	}
	return res, nil
}