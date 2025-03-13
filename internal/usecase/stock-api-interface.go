package usecase

import "myapi/infrastructure/api"

type StockAPI interface {
	GetStocksFromAPI(nextPage string) (api.StocksResponse, error)
}
