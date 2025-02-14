package repository

import (
    "merch-store/internal/domain"
    "gorm.io/gorm"
)

type UserRepositoryImpl struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
    return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) FindByID(id int) (*domain.User, error) {
    var user domain.User
    err := r.db.First(&user, id).Error
    return &user, err
}

func (r *UserRepositoryImpl) UpdateBalance(id int, balance int) error {
    return r.db.Model(&domain.User{}).Where("id = ?", id).Update("balance", balance).Error
}