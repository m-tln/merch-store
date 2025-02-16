package usecase

import (
	"fmt"
	"merch-store/internal/repository"
)

type InfoUseCase struct {
	userRepo        repository.UsersRepository
	goodsRepo       repository.ProductsRepository
	transactionRepo repository.TransactionsRepository
	purchaseRepo    repository.PurchasesRepository
}

func NewInfoUseCase(userRepo repository.UsersRepository, goodsRepo repository.ProductsRepository,
	transactionRepo repository.TransactionsRepository, purchaseRepo repository.PurchasesRepository) *InfoUseCase {
	return &InfoUseCase{
		userRepo:        userRepo,
		goodsRepo:       goodsRepo,
		transactionRepo: transactionRepo,
		purchaseRepo:    purchaseRepo,
	}
}

func (uc *InfoUseCase) GetBalance(id int) (int32, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return 0, fmt.Errorf("user with username %v can't be found in db, error: %v", id, err)
	}
	return int32(user.Balance), nil
}

func (uc *InfoUseCase) GetInvetory(id int) (map[string]int32, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return map[string]int32{}, fmt.Errorf("user with username %v can't be found in db, error: %v", id, err)
	}

	purchases, err := uc.purchaseRepo.FindByUserID(user.ID)
	if err != nil {
		return map[string]int32{}, fmt.Errorf("purchases of %v can't be found, error: %v", id, err)
	}

	res := make(map[string]int32)
	for _, purchase := range purchases {
		item, err := uc.goodsRepo.FindByID(int(purchase.IDItem))
		if err != nil {
			return map[string]int32{}, fmt.Errorf("%v can't be found in db, error: %v", purchase.IDItem, err)
		}
		res[item.Name] += int32(purchase.Volume)
	}

	return res, nil
}


func (uc *InfoUseCase) GetRecieved(id int) (map[string][]uint, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return map[string][]uint{}, fmt.Errorf("user with username %v can't be found in db, error: %v", id, err)
	}

	transactions, err := uc.transactionRepo.GetTransactionsByIDTo(user.ID)
	if err != nil {
		return map[string][]uint{}, fmt.Errorf("transactions of %v can't be found, error: %v", id, err)
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

func (uc *InfoUseCase) GetSent(id int) (map[string][]uint, error) {
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return map[string][]uint{}, fmt.Errorf("user with username %v can't be found in db, error: %v", id, err)
	}

	transactions, err := uc.transactionRepo.GetTransactionsByIDFrom(user.ID)
	if err != nil {
		return map[string][]uint{}, fmt.Errorf("transactions of %v can't be found, error: %v", id, err)
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