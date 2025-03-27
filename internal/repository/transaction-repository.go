package repository

import (
	"myapi/internal/domain"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepo {
	if db == nil {
		panic("NewTransactionRepository: database instance is nil")
	}
	return &TransactionRepo{db: db}
}

func (t *TransactionRepo) CreateTransactionBuy(transaction domain.Transaction) domain.Transaction {
	err := t.db.Create(&transaction).Error
	if err != nil {
		return domain.Transaction{}
	}

	return transaction
}
