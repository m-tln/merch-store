package app

import (
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
	router http.Handler

	log *logger.CustomLogger
	cfg *config.Config
}

func NewService() (*Service, error) {

	log, err := logger.NewCustomLogger()

	if err != nil {
		panic(err)
	}

	cfg, err := config.NewConfig()

	if err != nil {
		return nil, err
	}

	return &Service{log: log, cfg: cfg}, nil
}

func Init() error {
	svc, err := NewService()

	if err != nil {
		return fmt.Errorf("can't make new service: %v", err)
	}

	dsn, err := svc.cfg.GetDSN()
	if err != nil {
		return fmt.Errorf("can't get dsn: %v", err)
	}
	db, err := repository.InitBD(dsn)
	if err != nil {
		return fmt.Errorf("userDB can't be inited: %v", err)
	}

	userRepo := repository.NewUserRepositoryImpl(db)
	purchaseRepo := repository.NewPurchaseRepositoryImpl(db)
	goodsRepo := repository.NewGoodsRepository(db)
	transactionRepo := repository.NewTransactionRepositoryImpl(db)

	jwtSecret, err := svc.cfg.GetSecretJWT()
	if err != nil {
		return fmt.Errorf("can't get jwt secret: %v", err)
	}

	infoUseCase := usecase.NewInfoUseCase(userRepo, goodsRepo, transactionRepo, purchaseRepo)
	sendCoinsUseCase := usecase.NewSendCoinUseCase(userRepo, transactionRepo)
	purchaseUseCase := usecase.NewPurchaseUseCase(purchaseRepo, goodsRepo, userRepo)
	authUseCase := usecase.NewAuthUseCase(userRepo, service.NewJWTService(jwtSecret))

	APIService := handlers.NewCustomAPIService(*infoUseCase, *sendCoinsUseCase, *purchaseUseCase, *authUseCase)
	APIController := controller.NewCustomAPIController(*APIService, service.NewJWTService(jwtSecret))

	router := openapi.NewRouter(APIController)

	svc.router = router

	return http.ListenAndServe(":8080", router)
}

func Start() {

}

func Stop() {

}
