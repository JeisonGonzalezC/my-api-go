package usecase

import "myapi/internal/domain"

type StockRepository interface {
	GetStocks(tickers []string) []domain.Stock
	CreateStocks(stocks []domain.Stock) error
}
