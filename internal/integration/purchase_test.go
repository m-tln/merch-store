package integration_test

import (
	"merch-store/internal/domain"
	"merch-store/internal/infrastructure/repository"
	"merch-store/internal/usecase"
	"testing"
)

func TestPurchaseScenario(t *testing.T) {
	// Setup
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	usersRepo := repository.NewUserRepositoryImpl(db)
	productsRepo := repository.NewProductsRepositoryImpl(db)
	purchasesRepo := repository.NewPurchaseRepositoryImpl(db)

	purchaseUseCase := usecase.NewPurchaseUseCase(purchasesRepo, productsRepo, usersRepo)

	user := &domain.User{
		ID:           200,
		Name:         "Buyer",
		PasswordHash: "password",
		Balance:      1000,
	}
	err = usersRepo.Create(user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	item := &domain.Product{
		Name:  "t-shirt",
		Price: 80,
	}
	if err := db.Create(item).Error; err != nil {
		t.Fatalf("Failed to create item: %v", err)
	}

	// Purchase the item
	err = purchaseUseCase.MakePurchase(user.ID, item.Name)
	if err != nil {
		t.Fatalf("Failed to purchase item: %v", err)
	}

	// Verify the purchase
	if err := db.Where("IDCostumer = ? AND IDItem = ?", user.ID, item.ID).Error; err != nil {
		t.Fatalf("Failed to find purchase: %v", err)
	}

	// Verify the user's balance
	userToFind, err := usersRepo.FindByID(user.ID)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}
	if userToFind.Balance != user.Balance-item.Price {
		t.Error("Wrong balance")
	}

	// Cleanup
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM goods")
	db.Exec("DELETE FROM purchases")
}
