package usecase

import (
	"merch-store/internal/domain"
	"merch-store/internal/repository"
)

type GoodsUseCase struct {
	goodsRepo repository.GoodsRepository
}

func NewGoodsUseCase(goodsRepo repository.GoodsRepository) (*GoodsUseCase) {
	return &GoodsUseCase{goodsRepo: goodsRepo}
}

func (uc *GoodsUseCase) FindByID(id int) (*domain.Goods, error) {
	return uc.goodsRepo.FindByID(id)
}