package handler

import (
	"encoding/json"
	"myapi/internal/domain"
	"myapi/internal/usecase"
	"net/http"
)

func CreateTransactionBuy(w http.ResponseWriter, r *http.Request) {
	var newTransaction CreateTransactionRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&newTransaction); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := newTransaction.Validate(); err != nil {
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	transaction := domain.Transaction{
		Ticker: newTransaction.Ticker,
		Amount: newTransaction.Amount,
	}
	createdTransaction, err := usecase.CreateTransactionBuy(transaction)

	if err != nil {
		http.Error(w, "Error creating transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdTransaction)
}
