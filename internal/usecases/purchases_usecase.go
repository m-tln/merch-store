package usecase

import (
	"merch-store/internal/domain"
	"merch-store/internal/repository"
)


type PurchaseUseCase struct {
	purchaseRepo repository.PurchaseRepository
}

func NewPurchaseUseCase(purchaseRepo *repository.PurchaseRepository) (*PurchaseUseCase) {
	return	&PurchaseUseCase{purchaseRepo: *purchaseRepo}
}

func (uc *PurchaseUseCase) Create(purchase *domain.Purchase) error {
	return uc.purchaseRepo.Create(purchase)
}

func (uc *PurchaseUseCase) FindByUserID(id int) ([]domain.Purchase, error) {
	return uc.purchaseRepo.FindByUserID(id)
}
