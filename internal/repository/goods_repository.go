package repository

import "merch-store/internal/domain"

type GoodsRepository interface {
	FindByID(id int) (*domain.Goods, error)
}