package router

import (
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	SetupStocksRoutes(mux)
	SetupTransactionRoutes(mux)

	return mux
}
