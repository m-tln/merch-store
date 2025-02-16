package app

import (
	"context"
	"fmt"
	"net/http"

	"merch-store/internal/infrastructure/repository"

	openapi "merch-store/api/generated/go"
	"merch-store/internal/config"
	httpapi "merch-store/internal/infrastructure/http_api"
	"merch-store/internal/service"
	"merch-store/internal/usecase"
	"merch-store/pkg/logger"
)

type Service struct {
	server *http.Server

	Log *logger.CustomLogger
	cfg *config.Config
}

func NewService() (*Service, error) {

	log, err := logger.NewCustomLogger()

	if err != nil {
		panic(err)
	}

	cfg, err := config.NewConfig()

	if err != nil {
		return nil, fmt.Errorf("can't get config, error: %v", err)
	}

	dsn, err := cfg.GetDSN()
	if err != nil {
		return nil, fmt.Errorf("can't get dsn: %v", err)
	}
	db, err := repository.InitBD(dsn)
	if err != nil {
		return nil, fmt.Errorf("userDB can't be inited: %v", err)
	}

	userRepo := repository.NewUserRepositoryImpl(db)
	purchaseRepo := repository.NewPurchaseRepositoryImpl(db)
	productRepo := repository.NewProductsRepository(db)
	transactionRepo := repository.NewTransactionRepositoryImpl(db)

	jwtSecret, err := cfg.GetSecretJWT()
	if err != nil {
		return nil, fmt.Errorf("can't get jwt secret: %v", err)
	}

	infoUseCase := usecase.NewInfoUseCase(userRepo, productRepo, transactionRepo, purchaseRepo)
	sendCoinsUseCase := usecase.NewSendCoinUseCase(userRepo, transactionRepo)
	purchaseUseCase := usecase.NewPurchaseUseCase(purchaseRepo, productRepo, userRepo)
	authUseCase := usecase.NewAuthUseCase(userRepo, service.NewJWTService(jwtSecret))

	APIService := httpapi.NewCustomAPIService(*infoUseCase, *sendCoinsUseCase, *purchaseUseCase, *authUseCase)
	APIController := httpapi.NewCustomAPIController(*APIService, service.NewJWTService(jwtSecret))

	router := openapi.NewRouter(APIController)

	serverAddress := cfg.GetServerAddress()

	server := &http.Server{
		Addr:    serverAddress,
		Handler: router,
	}

	return &Service{server: server, Log: log, cfg: cfg}, nil
}

func (svc *Service) Start(_ context.Context) error {
	svc.Log.Info("Starting server...", map[string]interface{}{})
	return svc.server.ListenAndServe()
}

func (svc *Service) Stop(ctx context.Context) error {
	svc.Log.Info("Stopping server...", map[string]interface{}{})
	return svc.server.Shutdown(ctx)
}
