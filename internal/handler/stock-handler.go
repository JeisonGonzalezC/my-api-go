package handler

import (
	"encoding/json"
	"myapi/internal/usecase"
	"net/http"
)

type StockHandler struct {
	useCase *usecase.StockUseCase
}

func NewStockHandler(useCase *usecase.StockUseCase) *StockHandler {
	return &StockHandler{useCase: useCase}
}

func (h *StockHandler) GetStocks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	allowedParams := map[string]bool{
		"nextPage": true,
	}

	for key := range query {
		if !allowedParams[key] {
			http.Error(w, "Param not allowed: "+key, http.StatusBadRequest)
			return
		}
	}

	nextPage := query.Get("nextPage")

	stocksFromDb := h.useCase.GetStocksUseCase(nextPage)

	json.NewEncoder(w).Encode(stocksFromDb)
}
