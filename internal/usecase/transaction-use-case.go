package usecase

import (
	"fmt"
	"myapi/internal/domain"
)

type TransactionUseCase struct {
	transactionRepo TransactionRepository
	stockRepo       StockRepository
}

func NewTransactionUseCase(transactionRepo TransactionRepository, stockRepo StockRepository) *TransactionUseCase {
	return &TransactionUseCase{
		transactionRepo: transactionRepo,
		stockRepo:       stockRepo,
	}
}
func (uc *TransactionUseCase) CreateTransactionBuy(transaction domain.Transaction) (domain.Transaction, error) {
	var stock []domain.Stock = uc.stockRepo.GetStocks([]string{transaction.Ticker})

	if len(stock) == 0 {
		return domain.Transaction{}, fmt.Errorf("error creating transaction: stock not found")
	}

	newTransaction := uc.transactionRepo.CreateTransactionBuy(transaction)

	return newTransaction, nil
}
