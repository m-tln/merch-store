package usecase_test

import (
	"merch-store/internal/domain"
	"merch-store/internal/infrastructure/mocks"
	"merch-store/internal/usecase"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestInfoUseCase_GetBalance(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUserRepo := mocks.NewMockUsersRepository(ctrl)
	mockProductsRepo := mocks.NewMockProductsRepository(ctrl)
	mockTransactionsRepo := mocks.NewMockTransactionsRepository(ctrl)
	mockPurchaseRepo := mocks.NewMockPurchasesRepository(ctrl)
	expectedUser := &domain.User{
		ID:      0,
		Name:    "Joe",
		Balance: 1000,
	}
	mockUserRepo.EXPECT().FindByID(expectedUser.ID).Return(expectedUser, nil)

	infoUseCase := usecase.NewInfoUseCase(mockUserRepo, mockProductsRepo, mockTransactionsRepo, mockPurchaseRepo)

	balance, err := infoUseCase.GetBalance(expectedUser.ID)
	if err != nil {
		t.Errorf("GetBalance return error: %v", err)
	}

	if expectedUser.Balance != uint64(balance) {
		t.Errorf("Wrong balance")
	}

}

func TestInfoUseCase_GetInvetory(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUserRepo := mocks.NewMockUsersRepository(ctrl)
	mockProductsRepo := mocks.NewMockProductsRepository(ctrl)
	mockTransactionsRepo := mocks.NewMockTransactionsRepository(ctrl)
	mockPurchaseRepo := mocks.NewMockPurchasesRepository(ctrl)

	expectedUser := &domain.User{
		ID:      0,
		Name:    "Joe",
		Balance: 1000,
	}
	mockUserRepo.EXPECT().FindByID(expectedUser.ID).Return(expectedUser, nil)

	expectedPurchase := []domain.Purchase{
		{
			IDCostumer: 0,
			IDItem:     2,
			Volume:     2,
		},
		{
			IDCostumer: 0,
			IDItem:     2,
			Volume:     3,
		},
		{
			IDCostumer: 0,
			IDItem:     0,
			Volume:     1,
		},
	}
	mockPurchaseRepo.EXPECT().FindByUserID(gomock.Any()).Return(expectedPurchase, nil)
	mockProductsRepo.EXPECT().FindByID(gomock.Eq(2)).Return(&domain.Product{Name: "table"}, nil).Times(2)
	mockProductsRepo.EXPECT().FindByID(gomock.Eq(0)).Return(&domain.Product{Name: "chair"}, nil).Times(1)

	infoUseCase := usecase.NewInfoUseCase(mockUserRepo, mockProductsRepo, mockTransactionsRepo, mockPurchaseRepo)

	res, err := infoUseCase.GetInvetory(expectedUser.ID)

	if err != nil {
		t.Errorf("GetInventory return error: %v", err)
	}

	expected := map[string]int32{"chair": 1, "table": 5}
	for item, quantity := range res {
		if q, ok := expected[item]; q != quantity || !ok {
			t.Errorf("GetInventary return isn't similar with expectence")
		}
	}

}

func TestInfoUseCase_GetSent(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUserRepo := mocks.NewMockUsersRepository(ctrl)
	mockProductsRepo := mocks.NewMockProductsRepository(ctrl)
	mockTransactionsRepo := mocks.NewMockTransactionsRepository(ctrl)
	mockPurchaseRepo := mocks.NewMockPurchasesRepository(ctrl)

	expectedUser := &domain.User{
		ID:      0,
		Name:    "Joe",
		Balance: 1000,
	}
	mockUserRepo.EXPECT().FindByID(expectedUser.ID).Return(expectedUser, nil)

	expectedTransactions := []domain.Transaction{
		{
			IDFrom: 0,
			IDTo:   2,
			Volume: 100,
		},
		{
			IDFrom: 0,
			IDTo:   3,
			Volume: 200,
		},
		{
			IDFrom: 0,
			IDTo:   2,
			Volume: 100,
		},
	}
	mockTransactionsRepo.EXPECT().GetTransactionsByIDFrom(gomock.Eq(expectedUser.ID)).Return(expectedTransactions, nil)
	mockUserRepo.EXPECT().FindByID(gomock.Eq(2)).Return(&domain.User{Name: "Bob"}, nil).Times(2)
	mockUserRepo.EXPECT().FindByID(gomock.Eq(3)).Return(&domain.User{Name: "Ivan"}, nil)

	infoUseCase := usecase.NewInfoUseCase(mockUserRepo, mockProductsRepo, mockTransactionsRepo, mockPurchaseRepo)

	res, err := infoUseCase.GetSent(expectedUser.ID)
	if err != nil {
		t.Errorf("GetSent return error: %v", err)
	}

	expected := map[string][]uint{"Bob": {100, 100}, "Ivan": {200}}
	for name, transactions := range res {
		if !reflect.DeepEqual(transactions, expected[name]) {
			t.Errorf("GetSent return isn't similar with expectence")
		}
	}
}
