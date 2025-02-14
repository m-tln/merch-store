package repository

import (
	"fmt"

	"merch-store/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitBD(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error in initing db %v", err)
	}

	err = db.AutoMigrate(&domain.Goods{}, &domain.User{}, &domain.Transaction{}, &domain.Purchase{})

	if err != nil {
		return nil, fmt.Errorf("error auto migrate failed with error: %v", err)
	}

	return db, err
}