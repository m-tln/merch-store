package domain

import "time"

type User struct {
	ID           int `gorm:"primaryKey"`
	Name         string `gorm:"uniqueIndex"`
	PasswordHash string
	Balance      uint64
}

func (User) TableName() string {
	return "users"
}

type Token struct {
	AccessToken string
}

type Goods struct {
	ID    uint64 `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Price uint64 `gorm:"not null"`
}

func (Goods) TableName() string {
	return "goods"
}

type Purchase struct {
	IDCostumer uint64 `gorm:"not null"`
	IDItem     uint64 `gorm:"not null"`
	Volume     uint64 
	CreatedAt  time.Time
}

func (Purchase) TableName() string {
	return "purchases"
}

type Transaction struct {
	IDFrom    uint64 `gorm:"not null"`
	IDTo      uint64 `gorm:"not null"`
	Volume    uint64
	CreatedAt time.Time
}

func (Transaction) TableName() string {
	return "transactions"
}
