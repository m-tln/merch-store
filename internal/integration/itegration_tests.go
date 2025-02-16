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
			DATABASE_HOST:     "localhost",
			DATABASE_PORT:     "5432",
			DATABASE_USER:     "postgres",
			DATABASE_PASSWORD: "password",
			DATABASE_NAME:     "shop",
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
