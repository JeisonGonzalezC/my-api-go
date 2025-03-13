package router

import (
	"myapi/infrastructure/middleware"
	"myapi/internal/handler"
	"net/http"
)

func SetupTransactionRoutes(mux *http.ServeMux, handler *handler.TransactionHandler) {
	mux.Handle("/transaction", middleware.JSONResponseMiddleware(
		middleware.MethodHandler(http.MethodPost, http.HandlerFunc(handler.CreateTransactionBuy)),
	)) // POST /transaction
}
