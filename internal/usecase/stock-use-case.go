package usecase

import (
	"log"
	"myapi/infrastructure/api"
	"myapi/internal/domain"
	"myapi/internal/repository"
)

type StockUseCase struct {
	Items    []domain.Stock
	NextPage string
}

func createStockUseCase(stocks []domain.Stock) {
	if len(stocks) == 0 {
		return
	}
	err := repository.CreateStocksRepository(stocks)
	if err != nil {
		log.Printf("Error creating stocks: %v", err)
	}
}

func checkNewStocksToSave(existingStocks []domain.Stock, stocksFromApi api.StocksResponse) []domain.Stock {
	existingMap := make(map[string]bool, len(existingStocks))

	for _, stock := range existingStocks {
		existingMap[stock.Ticker] = true
	}

	var newStocks []domain.Stock
	for _, item := range stocksFromApi.Items {
		if !existingMap[item.Ticker] {
			currentStock := domain.Stock{
				Ticker:     item.Ticker,
				TargetFrom: item.TargetFrom,
				TargetTo:   item.TargetTo,
				Company:    item.Company,
				Action:     item.Action,
				Brokerage:  item.Brokerage,
				RatingFrom: item.RatingFrom,
				RatingTo:   item.RatingTo,
				Time:       item.Time,
			}
			newStocks = append(newStocks, currentStock)
		}
	}

	return newStocks
}

func GetStocksUseCase(nextPage string) StockUseCase {
	stocksFromApi, err := api.GetStocksFromAPI(nextPage)
	if err != nil {
		return StockUseCase{}
	}

	tickers := []string{}
	for _, item := range stocksFromApi.Items {
		tickers = append(tickers, item.Ticker)
	}

	var existingStocks []domain.Stock = repository.GetStocksRepository(tickers)
	var newStocks []domain.Stock = checkNewStocksToSave(existingStocks, stocksFromApi)

	createStockUseCase(newStocks)

	var newAndExistingStocks []domain.Stock = append(existingStocks, newStocks...)
	for i := range newAndExistingStocks {
		newAndExistingStocks[i].EvaluateRecommendation()
	}

	return StockUseCase{Items: newAndExistingStocks, NextPage: stocksFromApi.NextPage}
}
