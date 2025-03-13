package router

import (
	"myapi/infrastructure/middleware"
	"myapi/internal/handler"
	"net/http"
)

func SetupStocksRoutes(mux *http.ServeMux, handler *handler.StockHandler) {
	mux.Handle("/stocks", middleware.JSONResponseMiddleware(
		middleware.MethodHandler(http.MethodGet, http.HandlerFunc(handler.GetStocks)),
	)) // GET /stocks
}
