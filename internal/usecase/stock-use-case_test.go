package usecase

import (
	"errors"
	"myapi/infrastructure/api"
	"myapi/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock API
type MockAPI struct {
	mock.Mock
}

func (m *MockAPI) GetStocksFromAPI(nextPage string) (api.StocksResponse, error) {
	args := m.Called(nextPage)
	return args.Get(0).(api.StocksResponse), args.Error(1)
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetStocksRepository(tickers []string) []domain.Stock {
	args := m.Called(tickers)
	return args.Get(0).([]domain.Stock)
}

func (m *MockRepository) CreateStocksRepository(stocks []domain.Stock) error {
	args := m.Called(stocks)
	return args.Error(0)
}

func TestGetStocksUseCase_NewStocksAdded(t *testing.T) {
	mockAPI := new(MockAPI)
	mockRepo := new(MockRepository)

	apiResponse := api.StocksResponse{
		Items: []api.StockItemResponse{
			{Ticker: "AAPL", TargetFrom: "150", TargetTo: "170", Company: "Apple", Action: "upgraded", Brokerage: "Goldman", RatingFrom: "Hold", RatingTo: "Buy", Time: "2025-01-01"},
			{Ticker: "GOOGL", TargetFrom: "2500", TargetTo: "2700", Company: "Google", Action: "raised", Brokerage: "Morgan", RatingFrom: "Neutral", RatingTo: "Overweight", Time: "2025-01-01"},
		},
		NextPage: "page2",
	}

	mockAPI.On("GetStocksFromAPI", "").Return(apiResponse, nil)
	mockRepo.On("GetStocksRepository", mock.Anything).Return([]domain.Stock{})
	mockRepo.On("CreateStocksRepository", mock.Anything).Return(nil)

	result := GetStocksUseCase("")

	assert.Len(t, result.Items, 2)
	assert.Equal(t, "AAPL", result.Items[0].Ticker)
	assert.Equal(t, "GOOGL", result.Items[1].Ticker)
	mockRepo.AssertCalled(t, "CreateStocksRepository", mock.Anything)
}

func TestGetStocksUseCase_ExistingStocks(t *testing.T) {
	mockAPI := new(MockAPI)
	mockRepo := new(MockRepository)

	apiResponse := api.StocksResponse{
		Items: []api.StockItemResponse{
			{Ticker: "AAPL", TargetFrom: "150", TargetTo: "170", Company: "Apple", Action: "upgraded", Brokerage: "Goldman", RatingFrom: "Hold", RatingTo: "Buy", Time: "2025-01-01"},
		},
		NextPage: "page2",
	}

	existingStocks := []domain.Stock{
		{Ticker: "AAPL", TargetFrom: "150", TargetTo: "170", Company: "Apple", Action: "upgraded", Brokerage: "Goldman", RatingFrom: "Hold", RatingTo: "Buy", Time: "2025-01-01"},
	}

	mockAPI.On("GetStocksFromAPI", "").Return(apiResponse, nil)
	mockRepo.On("GetStocksRepository", mock.Anything).Return(existingStocks)
	mockRepo.On("CreateStocksRepository", mock.Anything).Return(nil)

	result := GetStocksUseCase("")

	assert.Len(t, result.Items, 1)
	mockRepo.AssertNotCalled(t, "CreateStocksRepository")
}

func TestGetStocksUseCase_ApiFails(t *testing.T) {
	mockAPI := new(MockAPI)

	mockAPI.On("GetStocksFromAPI", "").Return(api.StocksResponse{}, errors.New("API error"))

	result := GetStocksUseCase("")

	assert.Empty(t, result.Items)
}
