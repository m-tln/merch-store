package repository

import (
    "merch-store/internal/domain"
    "gorm.io/gorm"
)

type GoodsRepositoryImpl struct {
    db *gorm.DB
}

func NewGoodsRepository(db *gorm.DB) *GoodsRepositoryImpl {
    return &GoodsRepositoryImpl{db: db}
}

func (r *GoodsRepositoryImpl) Create(goods *domain.Goods) error {
    return r.db.Create(goods).Error
}

func (r *GoodsRepositoryImpl) FindByID(id int) (*domain.Goods, error) {
    var goods domain.Goods
    err := r.db.First(&goods, id).Error
    return &goods, err
}