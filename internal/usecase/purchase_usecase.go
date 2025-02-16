package usecase

import (
	"errors"
	"fmt"
	"merch-store/internal/domain"
	"merch-store/internal/repository"
)

type PurchaseUseCase struct {
	purchaseRepo repository.PurchasesRepository
	goodsRepo repository.ProductsRepository
	userRepo repository.UsersRepository
}

func NewPurchaseUseCase(purchaseRepo repository.PurchasesRepository, goodsRepo repository.ProductsRepository, 
						userRepo repository.UsersRepository) *PurchaseUseCase {
	return &PurchaseUseCase{purchaseRepo: purchaseRepo, goodsRepo: goodsRepo, userRepo: userRepo}
}

const SmallBalanceToBuy string = "not enough coins to buy"

func (uc *PurchaseUseCase) MakePurchase(id int, item string) error {
	good, err := uc.goodsRepo.FindByName(item)
	if err != nil {
		return fmt.Errorf("%s can't be found in db, error: %v", item, err)
	}
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("user with username %v can't be found in db, error: %v", id, err)
	}
	if user.Balance < good.Price {
		return errors.New(SmallBalanceToBuy)
	}

	err = uc.userRepo.UpdateBalance(user.ID, int(user.Balance - good.Price))

	if err != nil {
		return fmt.Errorf("balance of %v can't be updated, error: %v", id, err)
	}

	err = uc.purchaseRepo.Create(&domain.Purchase{
		IDCostumer: uint64(user.ID),
		IDItem: good.ID,
		Volume: 1,
	})

	if err != nil {
		return fmt.Errorf("new purchase can't be created username: %v, item: %s, error: %v", id, item, err)
	}

	return nil
}