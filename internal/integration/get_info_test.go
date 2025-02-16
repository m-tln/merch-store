package integration_test

import (
	"encoding/json"
	"merch-store/internal/domain"
	"merch-store/internal/infrastructure/repository"
	"net/http"
	"testing"
	"time"
)

func TestGetInfoScenario(t *testing.T) {
	// Setup
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepositoryImpl(db)
	productsRepo := repository.NewProductsRepositoryImpl(db)
	transactionRepo := repository.NewTransactionRepositoryImpl(db)
	purchaseRepo := repository.NewPurchaseRepositoryImpl(db)

	// Create a user
	user := &domain.User{
		Name:     "testuser2",
		PasswordHash: "hashedpassword",
		Balance:      1000,
	}
	if err := userRepo.Create(user); err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Create items
	items := []*domain.Product{
		{Name: "t-shirt", Price: 80},
		{Name: "cup", Price: 20},
	}
	for _, item := range items {
		if err := productsRepo.Create(item); err != nil {
			t.Fatalf("Failed to create item: %v", err)
		}
	}

	// Create transactions
	transactions := []*domain.Transaction{
		{IDFrom: 1, IDTo: 2, Volume: 200, CreatedAt: time.Now()},
		{IDFrom: 2, IDTo: 1, Volume: 100, CreatedAt: time.Now()},
	}
	for _, transaction := range transactions {
		if err := transactionRepo.Create(transaction); err != nil {
			t.Fatalf("Failed to create transaction: %v", err)
		}
	}

	// Create purchases
	purchases := []*domain.Purchase{
		{IDCostumer: 1, IDItem: 1, Volume: 1, CreatedAt: time.Now()},
		{IDCostumer: 1, IDItem: 2, Volume: 2, CreatedAt: time.Now()},
	}
	for _, purchase := range purchases {
		if err := purchaseRepo.Create(purchase); err != nil {
			t.Fatalf("Failed to create purchase: %v", err)
		}
	}

	// Send a GET request to /api/info
	resp, err := http.Get("http://localhost:8080/api/info")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Parse the response body
	var infoResponse struct {
		Coins     int `json:"coins"`
		Inventory []struct {
			Type     string `json:"type"`
			Quantity int    `json:"quantity"`
		} `json:"inventory"`
		CoinHistory struct {
			Received []struct {
				FromUser string `json:"fromUser"`
				Amount   int    `json:"amount"`
			} `json:"received"`
			Sent []struct {
				ToUser string `json:"toUser"`
				Amount int    `json:"amount"`
			} `json:"sent"`
		} `json:"coinHistory"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&infoResponse); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Verify the user's balance
	if infoResponse.Coins != 1000 {
		t.Errorf("Expected user balance to be 1000, got %d", infoResponse.Coins)
	}

	// Verify the inventory
	expectedInventory := map[string]int{
		"t-shirt": 1,
		"cup":     2,
	}
	for _, item := range infoResponse.Inventory {
		if expectedInventory[item.Type] != item.Quantity {
			t.Errorf("Expected %s quantity to be %d, got %d", item.Type, expectedInventory[item.Type], item.Quantity)
		}
	}

	// Verify the coin history
	if len(infoResponse.CoinHistory.Received) != 1 || infoResponse.CoinHistory.Received[0].Amount != 100 {
		t.Errorf("Expected received coin history to contain 1 transaction with amount 100")
	}
	if len(infoResponse.CoinHistory.Sent) != 1 || infoResponse.CoinHistory.Sent[0].Amount != 200 {
		t.Errorf("Expected sent coin history to contain 1 transaction with amount 200")
	}

	// Cleanup
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM goods")
	db.Exec("DELETE FROM transactions")
	db.Exec("DELETE FROM purchases")
}
