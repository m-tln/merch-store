package app

import (
	"context"
	"fmt"
	"net/http"

	"merch-store/adapter/logger"
	"merch-store/adapter/repository"
	"merch-store/api/controller"
	openapi "merch-store/api/generated/go"
	"merch-store/api/handlers"
	"merch-store/internal/config"
	"merch-store/internal/service"
	"merch-store/internal/usecase"
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
	goodsRepo := repository.NewGoodsRepository(db)
	transactionRepo := repository.NewTransactionRepositoryImpl(db)

	jwtSecret, err := cfg.GetSecretJWT()
	if err != nil {
		return nil, fmt.Errorf("can't get jwt secret: %v", err)
	}

	infoUseCase := usecase.NewInfoUseCase(userRepo, goodsRepo, transactionRepo, purchaseRepo)
	sendCoinsUseCase := usecase.NewSendCoinUseCase(userRepo, transactionRepo)
	purchaseUseCase := usecase.NewPurchaseUseCase(purchaseRepo, goodsRepo, userRepo)
	authUseCase := usecase.NewAuthUseCase(userRepo, service.NewJWTService(jwtSecret))

	APIService := handlers.NewCustomAPIService(*infoUseCase, *sendCoinsUseCase, *purchaseUseCase, *authUseCase)
	APIController := controller.NewCustomAPIController(*APIService, service.NewJWTService(jwtSecret))

	router := openapi.NewRouter(APIController)

	serverAddress := fmt.Sprintf("%s:%s", cfg.GetHost(), cfg.GetPort())

	server :=&http.Server{
		Addr: serverAddress,
		Handler: router,
	}

	return &Service{server: server, Log: log, cfg: cfg}, nil
}

func (svc *Service) Start(ctx context.Context) error {
	svc.Log.Info("Starting server...", map[string]interface{}{})
	return svc.server.ListenAndServe()
}

func (svc *Service) Stop(ctx context.Context) error {
	svc.Log.Info("Stopping server...", map[string]interface{}{})
	return svc.server.Shutdown(ctx)
}
