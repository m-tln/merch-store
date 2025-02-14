package domain

import "time"

type User struct {
	ID           int `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	PasswordHash string
	Balance      uint64
}

type Token struct {
	AccessToken string
}

type Goods struct {
	ID    uint64 `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Price uint64 `gorm:"not null"`
}

type Purchase struct {
	IDCostumer uint64 `gorm:"not null"`
	IDItem     uint64 `gorm:"not null"`
	Volume     uint64 
	CreatedAt  time.Time
}

type Transaction struct {
	IDFrom    uint64 `gorm:"not null"`
	IDTo      uint64 `gorm:"not null"`
	Volume    uint64
	CreatedAt time.Time
}
