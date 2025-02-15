package usecase

import (
	"errors"
	"fmt"
	"merch-store/internal/domain"
	"merch-store/internal/repository"
)

type PurchaseUseCase struct {
	purchaseRepo repository.PurchaseRepository
	goodsRepo repository.GoodsRepository
	userRepo repository.UserRepository
}

func NewPurchaseUseCase(purchaseRepo repository.PurchaseRepository, goodsRepo repository.GoodsRepository, 
						userRepo repository.UserRepository) *PurchaseUseCase {
	return &PurchaseUseCase{purchaseRepo: purchaseRepo, goodsRepo: goodsRepo, userRepo: userRepo}
}

const SmallBalanceToBuy string = "not enough coins to buy"

func (uc *PurchaseUseCase) MakePurchase(username string, item string) error {
	good, err := uc.goodsRepo.FindByName(item)
	if err != nil {
		return fmt.Errorf("%s can't be found in db, error: %v", item, err)
	}
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return fmt.Errorf("user with username %s can't be found in db, error: %v", username, err)
	}
	if user.Balance < good.Price {
		return errors.New(SmallBalanceToBuy)
	}

	err = uc.userRepo.UpdateBalance(user.ID, int(user.Balance - good.Price))

	if err != nil {
		return fmt.Errorf("balance of %s can't be updated, error: %v", username, err)
	}

	err = uc.purchaseRepo.Create(&domain.Purchase{
		IDCostumer: uint64(user.ID),
		IDItem: good.Price,
		Volume: 1,
	})

	if err != nil {
		return fmt.Errorf("new purchase can't be created username: %s, item: %s, error: %v", username, item, err)
	}

	return nil
}