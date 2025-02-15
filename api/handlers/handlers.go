package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"merch-store/adapter/logger"
	openapi "merch-store/api/generated/go"
	usecase "merch-store/internal/usecase"
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
func (s *CustomAPIService) ApiInfoGet(ctx context.Context) (openapi.ImplResponse, error) {
	s.log.Info("Get info", map[string]interface{}{})

	userIDraw := ctx.Value(middleware.KeyUserID)
	fmt.Println("raw: ", userIDraw)
	userIDstr, ok := userIDraw.(float64)
	fmt.Println("str: ", userIDstr)
	if !ok {
		s.log.Error("Missing userID in context", map[string]interface{}{})
		return openapi.Response(http.StatusUnauthorized, openapi.ErrorResponse{Errors: "Unauthorized: Missing userID"}), nil
	}

	userID := int(userIDstr)

	var err error
	// if err != nil {
	// 	s.log.Error("strcov failed", map[string]interface{}{"Error": err})
	// 	return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), nil
	// }

	s.log.Info("Request from user", map[string]interface{}{"userID": userID})
	responce := openapi.InfoResponse{}

	responce.Coins, err = s.infoUseCase.GetBalance(userID)
	if err != nil {
		s.log.Error("get balance failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), nil
	}

	inventory, err := s.infoUseCase.GetInvetory(userID)
	if err != nil {
		s.log.Error("get inventory failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), nil
	}

	for item, quantity := range inventory {
		responce.Inventory = append(responce.Inventory, openapi.InfoResponseInventoryInner{Type: item, Quantity: quantity})
	}

	historySent, err := s.infoUseCase.GetSent(userID)
	if err != nil {
		s.log.Error("get icoin history failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), nil
	}

	for userTo, transactions := range historySent {
		for _, transaction := range transactions {
			responce.CoinHistory.Sent = append(responce.CoinHistory.Sent,
				openapi.InfoResponseCoinHistorySentInner{ToUser: userTo, Amount: int32(transaction)})
		}
	}

	historyReceived, err := s.infoUseCase.GetRecieved(userID)
	if err != nil {
		s.log.Error("get icoin history failed", map[string]interface{}{"Error": err})
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), nil
	}

	for userFrom, transactions := range historyReceived {
		for _, transaction := range transactions {
			responce.CoinHistory.Received = append(responce.CoinHistory.Received,
				openapi.InfoResponseCoinHistoryReceivedInner{FromUser: userFrom, Amount: int32(transaction)})
		}
	}

	return openapi.Response(http.StatusOK, responce), nil
}

// ApiSendCoinPost - Отправить монеты другому пользователю.
func (s *CustomAPIService) ApiSendCoinPost(ctx context.Context, body openapi.SendCoinRequest) (openapi.ImplResponse, error) {
	s.log.Info("Send coin post", map[string]interface{}{})
	// TODO - update ApiSendCoinPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(401, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, ErrorResponse{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ApiSendCoinPost method not implemented")
}

// ApiBuyItemGet - Купить предмет за монеты.
func (s *CustomAPIService) ApiBuyItemGet(ctx context.Context, item string) (openapi.ImplResponse, error) {
	s.log.Info("Buy item get", map[string]interface{}{})
	// TODO - update ApiBuyItemGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(401, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, ErrorResponse{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ApiBuyItemGet method not implemented")
}

// ApiAuthPost - Аутентификация и получение JWT-токена.
func (s *CustomAPIService) ApiAuthPost(ctx context.Context, body openapi.AuthRequest) (openapi.ImplResponse, error) {
	s.log.Info("Auth post", map[string]interface{}{})

	str, err := s.authUseCase.GetToken(body.Username, body.Password)
	if err != nil {
		s.log.Error("get coin failed", map[string]interface{}{"Error": err})
	}

	return openapi.Response(http.StatusOK, openapi.AuthResponse{Token: *str}), nil
}
