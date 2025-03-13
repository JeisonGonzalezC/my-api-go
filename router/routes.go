package router

import (
	"myapi/internal/handler"
	"net/http"
)

type Handlers struct {
	StockHandler       *handler.StockHandler
	TransactionHandler *handler.TransactionHandler
}

func SetupRoutes(handlers Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	SetupStocksRoutes(mux, handlers.StockHandler)
	SetupTransactionRoutes(mux, handlers.TransactionHandler)

	return mux
}
