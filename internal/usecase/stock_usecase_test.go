package usecase_test

import (
	"myapi/infrastructure/api"
	"myapi/internal/domain"
	"myapi/internal/usecase"
	"myapi/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetStocksUseCase(t *testing.T) {
	mockRepo := new(mocks.StockRepository)
	mockAPI := new(mocks.StockAPI)
	uc := usecase.NewStockUseCase(mockRepo, mockAPI)

	stocksFromAPI := api.StocksResponse{
		Items: []api.StockItemResponse{
			{Ticker: "AAPL"},
			{Ticker: "MSFT"},
		},
		NextPage: "page_2",
	}

	stocksFromDB := []domain.Stock{
		{Ticker: "AAPL"},
	}

	mockAPI.On("GetStocksFromAPI", "").Return(stocksFromAPI, nil)

	mockRepo.On("GetStocks", mock.Anything).Return(stocksFromDB)

	mockRepo.On("CreateStocks", mock.Anything).Return(nil)

	result := uc.GetStocksUseCase("")

	assert.Len(t, result.Items, 2)
	assert.Equal(t, "page_2", result.NextPage)

	mockAPI.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
