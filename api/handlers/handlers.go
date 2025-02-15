package handlers

import (
	"context"
	"errors"
	"net/http"

	"merch-store/adapter/logger"
	openapi "merch-store/api/generated/go"
)

type CustomAPIService struct {
	// userRepo        repository.UserRepository
	// purchaseRepo    repository.PurchaseRepository
	// goodsRepo       repository.GoodsRepository
	// transactionRepo repository.TransactionRepository
	log             logger.CustomLogger
}

// NewDefaultAPIService creates a default api service
// func NewCustomAPIService(userRepo repository.UserRepository,
// 	purchaseRepo repository.PurchaseRepository,
// 	goodsRepo repository.GoodsRepository,
// 	transactionRepo repository.TransactionRepository) *CustomAPIService {
// 	return &CustomAPIService{userRepo: userRepo, purchaseRepo: purchaseRepo,
// 		goodsRepo: goodsRepo, transactionRepo: transactionRepo}
// }

func NewCustomAPIService() *CustomAPIService {
	return &CustomAPIService{log: logger.CustomLogger{}}
}

// ApiInfoGet - Получить информацию о монетах, инвентаре и истории транзакций.
func (s *CustomAPIService) ApiInfoGet(ctx context.Context) (openapi.ImplResponse, error) {
	s.log.Info("Get info", map[string]interface{}{})

	// userIDstr, ok := ctx.Value("userID").(string)
	// if !ok {
	// 	s.log.Error("Missing userID in context", map[string]interface{}{})
	// 	return openapi.Response(http.StatusUnauthorized, openapi.ErrorResponse{Errors: "Unauthorized: Missing userID"}), nil
	// }

	// userID, err := strconv.Atoi(userIDstr)

	// if err != nil {
	// 	s.log.Error("strcov failed", map[string]interface{}{"Error": err})
	// 	return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), nil
	// }

	// s.log.Info("Request from user", map[string]interface{}{"userID": userID})

	// user, err := s.userRepo.FindByID(userID)

	// if err != nil {
	// 	s.log.Error("user getting failed", map[string]interface{}{"Error": err})
	// }

	balance := 10 // user.Balance

	// _, err = s.transactionRepo.GetTransactionsByID(userID)

	// if err != nil {
	// 	s.log.Error("transactions getting failed", map[string]interface{}{"Error": err})
	// 	return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), nil
	// }

	// _, err = s.purchaseRepo.FindByUserID(userID)

	// if err != nil {
	// 	s.log.Error("purchases getting failed", map[string]interface{}{"Error": err})
	// 	return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{Errors: "Internal server error"}), nil
	// }


	return openapi.Response(http.StatusOK, openapi.InfoResponse{
		Coins: int32(balance),
	}), nil

	// TODO: Uncomment the next line to return response Response(200, InfoResponse{}) or use other options such as http.Ok ...
	// return Response(200, InfoResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(401, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, ErrorResponse{}), nil
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
	
	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ApiAuthPost method not implemented")	
}
