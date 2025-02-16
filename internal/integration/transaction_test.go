package integration_test

import (
	"merch-store/internal/domain"
	"merch-store/internal/infrastructure/repository"
	"merch-store/internal/usecase"
	"testing"
)

func TestTransferCoinsScenario(t *testing.T) {
	// Setup
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	userRepo := repository.NewUserRepositoryImpl(db)
	transactionRepo := repository.NewTransactionRepositoryImpl(db)

	transactionUseCase := usecase.NewSendCoinUseCase(userRepo, transactionRepo)

	sendler := &domain.User{
		ID:           200,
		Name:         "sendler",
		PasswordHash: "password",
		Balance:      1000,
	}
	err = userRepo.Create(sendler)
	if err != nil {
		t.Fatalf("Failed to create sender: %v", err)
	}

	receiver := &domain.User{
		ID:           201,
		Name:         "receiver",
		PasswordHash: "password",
		Balance:      1000,
	}
	err = userRepo.Create(receiver)
	if err != nil {
		t.Fatalf("Failed to create receiver: %v", err)
	}

	// Transfer coins
	err = transactionUseCase.MakeTransaction(sendler.ID, receiver.Name, int32(sendler.Balance / 2))
	if err != nil {
		t.Fatalf("Failed to transfer coins: %v", err)
	}

	// Verify the transaction
	if err := db.Where("IDFrom = ? AND IDTo = ?", sendler.ID, receiver.ID).Error; err != nil {
		t.Fatalf("Failed to find transaction: %v", err)
	}

	// Verify the sender's balance
	senderAfter, err := userRepo.FindByUsername(sendler.Name)
	if err != nil {
		t.Fatalf("Failed to get sender: %v", err)
	}
	if senderAfter.Balance != sendler.Balance - (sendler.Balance / 2) {
		t.Error("Wrong sendler balance")
	}

	// Verify the receiver's balance
	receiverAfter, err := userRepo.FindByUsername(receiver.Name)
	if err != nil {
		t.Fatalf("Failed to get receiver: %v", err)
	}
	if receiverAfter.Balance != receiver.Balance + (sendler.Balance / 2) {
		t.Error("Wrong receiver balance")
	}

	// Cleanup
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM transactions")
}
