package httpapi

import (
	"context"
	"errors"
	"net/http"

	openapi "merch-store/api/generated/go"
	"merch-store/internal/usecase"
	"merch-store/pkg/logger"
	"merch-store/pkg/middleware"
)

type CustomAPIService struct {
	infoUseCase      usecase.InfoUseCase
	sendCoinsUseCase usecase.SendCoinUseCase
	purchaseUseCase  usecase.PurchaseUseCase
	authUseCase      usecase.AuthUseCase
	log              logger.CustomLogger
}

// NewDefaultAPIService creates a default api service
func NewCustomAPIService(infoUseCase usecase.InfoUseCase,
	sendCoinsUseCase usecase.SendCoinUseCase,
	purchaseUseCase usecase.PurchaseUseCase,
	authUseCase usecase.AuthUseCase) *CustomAPIService {
	return &CustomAPIService{infoUseCase: infoUseCase, purchaseUseCase: purchaseUseCase,
		sendCoinsUseCase: sendCoinsUseCase, authUseCase: authUseCase}
}

// ApiInfoGet - Получить информацию о монетах, инвентаре и истории транзакций.
func (s *CustomAPIService) APIInfoGet(ctx context.Context) (openapi.ImplResponse, error) {
	s.log.Info("Get info", map[string]interface{}{})

	userIDraw := ctx.Value(middleware.KeyUserID)
	userIDstr, ok := userIDraw.(float64)
	if !ok {
		s.log.Error("Missing userID in context", map[string]interface{}{})
		return openapi.Response(http.StatusUnauthorized, openapi.ErrorResponse{Errors: "Unauthorized: Missing userID"}), nil
	}

	userID := int(userIDstr)

	var err error

	s.log.Info("Request from user", map[string]interface{}{"userID": userID})
	response := openapi.InfoResponse{}

	response.Coins, err = s.infoUseCase.GetBalance(userID)
	if err != nil {
		s.log.Error("get balance failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), err
	}

	inventory, err := s.infoUseCase.GetInvetory(userID)
	if err != nil {
		s.log.Error("get inventory failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), err
	}

	for item, quantity := range inventory {
		response.Inventory = append(response.Inventory, openapi.InfoResponseInventoryInner{Type: item, Quantity: quantity})
	}

	historySent, err := s.infoUseCase.GetSent(userID)
	if err != nil {
		s.log.Error("get icoin history failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), err
	}

	for userTo, transactions := range historySent {
		for _, transaction := range transactions {
			response.CoinHistory.Sent = append(response.CoinHistory.Sent,
				openapi.InfoResponseCoinHistorySentInner{ToUser: userTo, Amount: int32(transaction)})
		}
	}

	historyReceived, err := s.infoUseCase.GetRecieved(userID)
	if err != nil {
		s.log.Error("get coin history failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), err
	}

	for userFrom, transactions := range historyReceived {
		for _, transaction := range transactions {
			response.CoinHistory.Received = append(response.CoinHistory.Received,
				openapi.InfoResponseCoinHistoryReceivedInner{FromUser: userFrom, Amount: int32(transaction)})
		}
	}

	return openapi.Response(http.StatusOK, response), nil
}

// ApiSendCoinPost - Отправить монеты другому пользователю.
func (s *CustomAPIService) APISendCoinPost(ctx context.Context, body openapi.SendCoinRequest) (openapi.ImplResponse, error) {
	s.log.Info("Send coin post", map[string]interface{}{})

	userIDraw := ctx.Value(middleware.KeyUserID)
	userIDstr, ok := userIDraw.(float64)
	if !ok {
		s.log.Error("Missing userID in context", map[string]interface{}{})
		return openapi.Response(http.StatusUnauthorized, openapi.ErrorResponse{Errors: "Unauthorized: Missing userID"}), nil
	}
	userID := int(userIDstr)

	err := s.sendCoinsUseCase.MakeTransaction(userID, body.ToUser, body.Amount)
	if err != nil {
		if errors.Is(err, errors.New(usecase.SmallBalanceToBuy)) {
			return openapi.Response(http.StatusBadRequest, openapi.ErrorResponse{Errors: usecase.SmallBalanceToBuy}), nil
		}
		s.log.Error("make transaction failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}

// ApiBuyItemGet - Купить предмет за монеты.
func (s *CustomAPIService) APIBuyItemGet(ctx context.Context, item string) (openapi.ImplResponse, error) {
	s.log.Info("Buy item get", map[string]interface{}{})

	userIDraw := ctx.Value(middleware.KeyUserID)
	userIDstr, ok := userIDraw.(float64)
	if !ok {
		s.log.Error("Missing userID in context", map[string]interface{}{})
		return openapi.Response(http.StatusUnauthorized, openapi.ErrorResponse{Errors: "Unauthorized: Missing userID"}), nil
	}
	userID := int(userIDstr)

	err := s.purchaseUseCase.MakePurchase(userID, item)
	if err != nil {
		if errors.Is(err, errors.New(usecase.SmallBalanceToSend)) {
			return openapi.Response(http.StatusBadRequest, openapi.ErrorResponse{Errors: usecase.SmallBalanceToSend}), nil
		}
		s.log.Error("make purchase failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), err
	}

	return openapi.Response(http.StatusOK, nil), nil
}

// ApiAuthPost - Аутентификация и получение JWT-токена.
func (s *CustomAPIService) APIAuthPost(ctx context.Context, body openapi.AuthRequest) (openapi.ImplResponse, error) {
	s.log.Info("Auth post", map[string]interface{}{})

	str, err := s.authUseCase.GetToken(body.Username, body.Password)
	if err != nil {
		s.log.Error("get coin failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), err
	}

	return openapi.Response(http.StatusOK, openapi.AuthResponse{Token: *str}), nil
}
