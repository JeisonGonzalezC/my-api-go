package usecase

import (
	"fmt"
	"myapi/internal/domain"
	"myapi/internal/repository"
)

func CreateTransactionBuy(transaction domain.Transaction) (domain.Transaction, error) {
	var stock []domain.Stock = repository.GetStocksRepository([]string{transaction.Ticker})

	if len(stock) == 0 {
		return domain.Transaction{}, fmt.Errorf("error creating transaction: stock not found")
	}

	newTransaction := repository.CreateTransactionBuy(transaction)

	return newTransaction, nil
}
