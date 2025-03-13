package bootstrap

import (
	"myapi/infrastructure/api"
	"myapi/internal/handler"
	"myapi/internal/repository"
	"myapi/internal/usecase"
	"myapi/router"
)

func InitApp() router.Handlers {
	// Init APIs
	apiClient := api.NewStockAPIClient()

	// Init repositories
	stockRepo := repository.NewStockRepository()
	transactionRepo := repository.NewTransactionRepository()

	// Init use cases
	stockUseCase := usecase.NewStockUseCase(stockRepo, apiClient)
	transactionUseCase := usecase.NewTransactionUseCase(transactionRepo, stockRepo)

	// Init handlers
	stockHandler := handler.NewStockHandler(stockUseCase)
	transactionHandler := handler.NewTransactionHandler(transactionUseCase)

	// Retornar los handlers
	return router.Handlers{
		StockHandler:       stockHandler,
		TransactionHandler: transactionHandler,
	}
}
