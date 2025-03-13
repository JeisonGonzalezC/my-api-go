package repository

import (
	"myapi/infrastructure/database"
	"myapi/internal/domain"
)

func CreateTransactionBuy(transaction domain.Transaction) domain.Transaction {
	err := database.DB.Create(&transaction).Error
	if err != nil {
		return domain.Transaction{}
	}

	return transaction
}
