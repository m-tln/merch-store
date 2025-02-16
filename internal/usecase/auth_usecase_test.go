package usecase_test

import (
	"errors"
	"merch-store/internal/domain"
	"merch-store/internal/infrastructure/mocks"
	"merch-store/internal/service"
	"merch-store/internal/usecase"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAuthUseCase_GetToken(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUserRepo := mocks.NewMockUsersRepository(ctrl)

	hash, _ := usecase.HashPassword("password")
	expectedNewUser := &domain.User{
		ID:           0,
		Name:         "Joe",
		PasswordHash: hash,
		Balance:      1000,
	}

	mockUserRepo.EXPECT().Create(gomock.Any()).Return(nil).Times(1)
	mockUserRepo.EXPECT().FindByUsername(gomock.Any()).Return(nil, errors.New("No user with username"))

	authUseCase := usecase.NewAuthUseCase(mockUserRepo, service.NewJWTService("secret-code"))
	_, err := authUseCase.GetToken("Joe", "password")

	if err != nil {
		t.Errorf("GetToken return error for new user: %v", err)
	}

	mockUserRepo.EXPECT().FindByUsername(gomock.Any()).Return(expectedNewUser, nil)
	_, err = authUseCase.GetToken("Joe", "password")

	if err != nil {
		t.Errorf("GetToken return error for user who already have been in db: %v", err)
	}

	wrongPasswordUser := &domain.User{
		ID:           0,
		Name:         "Joe",
		PasswordHash: "11111111",
		Balance:      1000,
	}
	mockUserRepo.EXPECT().FindByUsername(gomock.Any()).Return(wrongPasswordUser, nil)
	_, err = authUseCase.GetToken("Joe", "password")
	if err == nil {
		t.Errorf("GetToken didn't catch wrong password")
	}
}
