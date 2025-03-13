package repository

import (
	"fmt"
	"myapi/infrastructure/database"
	"myapi/internal/domain"
)

func GetStocksRepository(tickers []string) []domain.Stock {
	var stocks []domain.Stock

	if err := database.DB.Preload("Transactions").Where("ticker IN ?", tickers).Find(&stocks).Error; err != nil {
		return []domain.Stock{}
	}

	return stocks
}

func CreateStocksRepository(stocks []domain.Stock) error {
	if err := database.DB.Create(&stocks).Error; err != nil {
		return fmt.Errorf("error creating stocks: %w", err)
	}

	return nil
}
