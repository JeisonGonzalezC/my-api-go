package handler

import (
	"encoding/json"
	"myapi/internal/usecase"
	"net/http"
)

func GetStocks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	allowedParams := map[string]bool{
		"nextPage": true,
	}

	for key := range query {
		if !allowedParams[key] {
			http.Error(w, "Param not allowred: "+key, http.StatusBadRequest)
			return
		}
	}

	var nextPage string = query.Get("nextPage")

	stocksFromDb := usecase.GetStocksUseCase(nextPage)
	json.NewEncoder(w).Encode(stocksFromDb)
}
