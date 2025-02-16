package usecase_test

import (
	"merch-store/internal/domain"
	"merch-store/internal/infrastructure/mocks"
	"merch-store/internal/usecase"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSendCoinUseCase_MakeTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUserRepo := mocks.NewMockUsersRepository(ctrl)
	mockTransactionsRepo := mocks.NewMockTransactionsRepository(ctrl)
	expectedUser1 := &domain.User{
		ID:      0,
		Name:    "Joe",
		Balance: 1000,
	}
	mockUserRepo.EXPECT().FindByID(expectedUser1.ID).Return(expectedUser1, nil)
	expectedUser2 := &domain.User{
		ID:      0,
		Name:    "Michael",
		Balance: 1000,
	}
	mockUserRepo.EXPECT().FindByUsername(expectedUser2.Name).Return(expectedUser2, nil)

	amount := 100
	mockUserRepo.EXPECT().UpdateBalance(expectedUser1.ID, int(expectedUser1.Balance)-amount)
	mockUserRepo.EXPECT().UpdateBalance(expectedUser2.ID, int(expectedUser2.Balance)+amount)

	expectedTransaction := &domain.Transaction{
		IDFrom: uint64(expectedUser1.ID),
		IDTo:   uint64(expectedUser2.ID),
		Volume: uint64(amount),
	}
	mockTransactionsRepo.EXPECT().Create(expectedTransaction).Return(nil)

	sendCoinUseCase := usecase.NewSendCoinUseCase(mockUserRepo, mockTransactionsRepo)
	if err := sendCoinUseCase.MakeTransaction(expectedUser1.ID, expectedUser2.Name, int32(amount)); err != nil {
		t.Errorf("MakeTransaction return error: %v", err)
	}

}
