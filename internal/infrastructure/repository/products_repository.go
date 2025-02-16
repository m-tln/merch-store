package repository

import (
	"merch-store/internal/domain"

	"gorm.io/gorm"
)

type ProductsRepositoryImpl struct {
	db *gorm.DB
}

func NewProductsRepositoryImpl(db *gorm.DB) *ProductsRepositoryImpl {
	return &ProductsRepositoryImpl{db: db}
}

func (r *ProductsRepositoryImpl) Create(product *domain.Product) error {
	productImpl := ProductFromDomainToRepo(product)
	return r.db.Create(productImpl).Error
}

func (r *ProductsRepositoryImpl) FindByID(id int) (*domain.Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	return ProductFromRepoToDomain(&product), err
}

func (r *ProductsRepositoryImpl) FindByName(name string) (*domain.Product, error) {
	var product Product
	err := r.db.Where("name = ?", name).First(&product).Error
	return ProductFromRepoToDomain(&product), err
}
