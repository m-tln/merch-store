package app

// import (
// 	"fmt"
// 	"net/http"

// 	"merch-store/adapter/logger"
// 	"merch-store/adapter/repository"
// 	openapi "merch-store/api/generated/go"
// 	"merch-store/api/handlers"
// 	"merch-store/internal/config"
// 	"merch-store/internal/service"
// 	"merch-store/internal/usecases"
// 	"merch-store/pkg/middleware"
// )

// type Service struct {
// 	router http.Handler

// 	log *logger.CustomLogger
// 	cfg *config.Config
// }

// func NewService() (*Service, error) {

// 	log, err := logger.NewCustomLogger()
	
// 	if err != nil {
// 		panic(err)
// 	}

// 	cfg, err := config.NewConfig()

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Service{log: log, cfg: cfg}, nil
// }

// func Init() error {
// 	svc, err := NewService()

// 	if err != nil {
// 		return fmt.Errorf("can't make new service: %v", err)
// 	}

// 	dsn, err := svc.cfg.GetUserDSN()
// 	if err != nil {
// 		return fmt.Errorf("can't get user dsn: %v", err)
// 	}
// 	db, err := repository.InitBD(dsn)
// 	if err != nil {
// 		return fmt.Errorf("userDB can't be inited: %v", err)
// 	}

// 	userRepo := repository.NewUserRepositoryImpl(db)
// 	purchaseRepo := repository.NewPurchaseRepositoryImpl(db)
// 	goodsRepo := repository.NewGoodsRepository(db)
// 	transactionRepo := repository.NewTransactionRepositoryImpl(db)

// 	userRepoUseCase := usecase.NewUserUseCase(userRepo)
// 	purchaseRepoUseCase := usecase.NewPurchaseUseCase(purchaseRepo)
// 	goodsRepoUseCase := usecase.NewGoodsUseCase(goodsRepo)
// 	transactionRepoUseCase := usecase.NewTransactionUseCase(transactionRepo)


// 	APIService := handlers.NewCustomAPIService(userRepoUseCase, purchaseRepoUseCase, 
// 													  goodsRepoUseCase, transactionRepoUseCase)
// 	APIController := openapi.NewDefaultAPIController(APIService)

// 	router := openapi.NewRouter(APIController)

// 	authMiddleware := middleware.AuthMiddleware(service.NewJWTService("pronin"))
// 	protectRouter := authMiddleware(router)

// 	svc.router = protectRouter


// 	return http.ListenAndServe(":8080", protectRouter)
// }

// func Start() {

// }

// func Stop() {

// }