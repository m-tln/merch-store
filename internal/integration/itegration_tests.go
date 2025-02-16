package integration_test

import (
	"fmt"
	"merch-store/internal/config"
	"merch-store/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	cfg := &config.Config{
		DBConfig: config.DBConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "password",
			Name:     "shop",
		},
	}
	dsn, err := cfg.GetDSN()
	if err != nil {
		return nil, fmt.Errorf("can't get dsn, error: %v", err)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&domain.User{}, &domain.Product{}, &domain.Purchase{}); err != nil {
		return nil, err
	}

	return db, nil
}
