package usecase

import (
	"log"
	"myapi/infrastructure/api"
	"myapi/internal/domain"
)

type StockUseCase struct {
	repo StockRepository
	api  StockAPI
}

type StockResult struct {
	Items    []domain.Stock
	NextPage string
}

func NewStockUseCase(repo StockRepository, api StockAPI) *StockUseCase {
	return &StockUseCase{repo: repo, api: api}
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

func (uc *StockUseCase) GetStocksUseCase(nextPage string) StockResult {
	stocksFromApi, err := uc.api.GetStocksFromAPI(nextPage)
	if err != nil {
		return StockResult{}
	}

	tickers := []string{}
	for _, item := range stocksFromApi.Items {
		tickers = append(tickers, item.Ticker)
	}

	existingStocks := uc.repo.GetStocks(tickers)
	newStocks := checkNewStocksToSave(existingStocks, stocksFromApi)

	err = uc.repo.CreateStocks(newStocks)
	if err != nil {
		log.Printf("Error creatin stocks: %v", err)
	}

	newAndExistingStocks := append(existingStocks, newStocks...)
	for i := range newAndExistingStocks {
		newAndExistingStocks[i].EvaluateRecommendation()
	}

	return StockResult{Items: newAndExistingStocks, NextPage: stocksFromApi.NextPage}
}
