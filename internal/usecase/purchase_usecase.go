package usecase

import (
	"errors"
	"fmt"
	"merch-store/internal/domain"
	"merch-store/internal/repository"
)

type PurchaseUseCase struct {
	purchaseRepo repository.PurchasesRepository
	productsRepo repository.ProductsRepository
	userRepo     repository.UsersRepository
}

func NewPurchaseUseCase(purchaseRepo repository.PurchasesRepository, productsRepo repository.ProductsRepository,
	userRepo repository.UsersRepository) *PurchaseUseCase {
	return &PurchaseUseCase{purchaseRepo: purchaseRepo, productsRepo: productsRepo, userRepo: userRepo}
}

const SmallBalanceToBuy string = "not enough coins to buy"

func (uc *PurchaseUseCase) MakePurchase(id int, item string) error {
	product, err := uc.productsRepo.FindByName(item)
	if err != nil {
		return fmt.Errorf("%s can't be found in db, error: %v", item, err)
	}
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("user with username %v can't be found in db, error: %v", id, err)
	}
	if user.Balance < product.Price {
		return errors.New(SmallBalanceToBuy)
	}

	err = uc.userRepo.UpdateBalance(user.ID, int(user.Balance-product.Price))

	if err != nil {
		return fmt.Errorf("balance of %v can't be updated, error: %v", id, err)
	}

	err = uc.purchaseRepo.Create(&domain.Purchase{
		IDCostumer: uint64(user.ID),
		IDItem:     product.ID,
		Volume:     1,
	})

	if err != nil {
		return fmt.Errorf("new purchase can't be created username: %v, item: %s, error: %v", id, item, err)
	}

	return nil
}
