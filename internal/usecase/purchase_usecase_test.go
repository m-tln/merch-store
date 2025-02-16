package usecase_test

import (
	"merch-store/internal/domain"
	"merch-store/internal/infrastructure/mocks"
	"merch-store/internal/usecase"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestPurchaseUseCase_MakePurchase(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUserRepo := mocks.NewMockUsersRepository(ctrl)
	mockProductsRepo := mocks.NewMockProductsRepository(ctrl)
	mockPurchaseRepo := mocks.NewMockPurchasesRepository(ctrl)

	purchaseUseCase := usecase.NewPurchaseUseCase(mockPurchaseRepo, mockProductsRepo, mockUserRepo)

	expectedProduct := &domain.Product{
		ID:    0,
		Name:  "table",
		Price: 500,
	}
	mockProductsRepo.EXPECT().FindByName(gomock.Eq(expectedProduct.Name)).Return(expectedProduct, nil)

	expectedUser := &domain.User{
		ID:      0,
		Name:    "Joe",
		Balance: 1000,
	}
	mockUserRepo.EXPECT().FindByID(gomock.Eq(expectedUser.ID)).Return(expectedUser, nil)
	mockUserRepo.EXPECT().UpdateBalance(gomock.Eq(expectedUser.ID),
		gomock.Eq(int(expectedUser.Balance-expectedProduct.Price))).Return(nil)

	mockPurchaseRepo.EXPECT().Create(gomock.Any()).Return(nil)

	err := purchaseUseCase.MakePurchase(expectedUser.ID, expectedProduct.Name)

	if err != nil {
		t.Errorf("Make purchase return error: %v", err)
	}
}
