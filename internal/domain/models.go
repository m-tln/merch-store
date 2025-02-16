package domain

import "time"

type User struct {
	ID           int
	Name         string
	PasswordHash string
	Balance      uint64
}

type Token struct {
	AccessToken string
}

type Product struct {
	ID    uint64
	Name  string
	Price uint64
}

type Purchase struct {
	IDCostumer uint64
	IDItem     uint64
	Volume     uint64
	CreatedAt  time.Time
}

type Transaction struct {
	IDFrom    uint64
	IDTo      uint64
	Volume    uint64
	CreatedAt time.Time
}
