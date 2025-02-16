package repository

import (
	"merch-store/internal/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	return r.db.Create(UserFromDomainToRepo(user)).Error
}

func (r *UserRepositoryImpl) FindByID(id int) (*domain.User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return UserFromRepoToDomain(&user), err
}

func (r *UserRepositoryImpl) UpdateBalance(id int, balance int) error {
	return r.db.Model(&User{}).Where("id = ?", id).Update("balance", balance).Error
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var user User
	err := r.db.Where("name = ?", username).First(&user).Error
	return UserFromRepoToDomain(&user), err
}
