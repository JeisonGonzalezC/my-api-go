package repository

import (
	"fmt"
	"myapi/internal/domain"

	"gorm.io/gorm"
)

// StockRepo the implementation of StockRepository
type StockRepo struct {
	db *gorm.DB
}

// NewStockRepository return a new instance of StockRepo
func NewStockRepository(db *gorm.DB) *StockRepo {
	if db == nil {
		panic("NewStockRepository: database instance is nil")
	}
	return &StockRepo{db: db}
}

func (r *StockRepo) GetStocks(tickers []string) []domain.Stock {
	var stocks []domain.Stock

	if err := r.db.Preload("Transactions").Where("ticker IN ?", tickers).Find(&stocks).Error; err != nil {
		return []domain.Stock{}
	}

	return stocks
}

func (r *StockRepo) CreateStocks(stocks []domain.Stock) error {
	if len(stocks) == 0 {
		return nil
	}

	if err := r.db.Create(&stocks).Error; err != nil {
		return fmt.Errorf("error creating stocks: %w", err)
	}

	return nil
}
