package usecase

import "myapi/internal/domain"

type TransactionRepository interface {
	CreateTransactionBuy(transaction domain.Transaction) domain.Transaction
}
