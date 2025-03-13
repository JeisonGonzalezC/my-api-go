package repository

import (
	"myapi/infrastructure/database"
	"myapi/internal/domain"
)

type TransactionRepo struct{}

func NewTransactionRepository() *TransactionRepo {
	return &TransactionRepo{}
}

func (t *TransactionRepo) CreateTransactionBuy(transaction domain.Transaction) domain.Transaction {
	err := database.DB.Create(&transaction).Error
	if err != nil {
		return domain.Transaction{}
	}

	return transaction
}
