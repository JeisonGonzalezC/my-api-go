package repository

import (
	"fmt"
	"myapi/infrastructure/database"
	"myapi/internal/domain"
)

// StockRepo the implementation of StockRepository
type StockRepo struct{}

// NewStockRepository return a new instance of StockRepo
func NewStockRepository() *StockRepo {
	return &StockRepo{}
}

func (r *StockRepo) GetStocks(tickers []string) []domain.Stock {
	var stocks []domain.Stock

	if err := database.DB.Preload("Transactions").Where("ticker IN ?", tickers).Find(&stocks).Error; err != nil {
		return []domain.Stock{}
	}

	return stocks
}

func (r *StockRepo) CreateStocks(stocks []domain.Stock) error {
	if len(stocks) == 0 {
		return nil
	}

	if err := database.DB.Create(&stocks).Error; err != nil {
		return fmt.Errorf("error creating stocks: %w", err)
	}

	return nil
}
