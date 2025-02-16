package repository

import "merch-store/internal/domain"

func ProductFromDomainToRepo(p *domain.Product) *Product {
	return &Product{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}
}

func ProductFromRepoToDomain(p *Product) *domain.Product {
	return &domain.Product{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}
}

func UserFromDomainToRepo(u *domain.User) *User {
	return &User{
		ID:           u.ID,
		Name:         u.Name,
		PasswordHash: u.PasswordHash,
		Balance:      u.Balance,
	}
}

func UserFromRepoToDomain(u *User) *domain.User {
	return &domain.User{
		ID:           u.ID,
		Name:         u.Name,
		PasswordHash: u.PasswordHash,
		Balance:      u.Balance,
	}
}

func PurchaseFromDomainToRepo(p *domain.Purchase) *Purchase {
	return &Purchase{
		IDCostumer: p.IDCostumer,
		IDItem:     p.IDItem,
		Volume:     p.Volume,
		CreatedAt:  p.CreatedAt,
	}
}

func PurchaseFromRepoToDomain(p *Purchase) *domain.Purchase {
	return &domain.Purchase{
		IDCostumer: p.IDCostumer,
		IDItem:     p.IDItem,
		Volume:     p.Volume,
		CreatedAt:  p.CreatedAt,
	}
}

func TransactionFromDomainToRepo(t *domain.Transaction) *Transaction {
	return &Transaction{
		IDFrom:    t.IDFrom,
		IDTo:      t.IDTo,
		Volume:    t.Volume,
		CreatedAt: t.CreatedAt,
	}
}

func TransactionFromRepoToDomain(t *Transaction) *domain.Transaction {
	return &domain.Transaction{
		IDFrom:    t.IDFrom,
		IDTo:      t.IDTo,
		Volume:    t.Volume,
		CreatedAt: t.CreatedAt,
	}
}
