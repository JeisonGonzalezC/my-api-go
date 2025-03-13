package usecase_test

import (
	"myapi/internal/domain"
	"myapi/internal/usecase"
	"myapi/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionBuyCase(t *testing.T) {
	mockRepo := new(mocks.TransactionRepository)
	ucStockRepo := new(mocks.StockRepository)
	uc := usecase.NewTransactionUseCase(mockRepo, ucStockRepo)

	ucStockRepo.On("GetStocks", mock.Anything).Return([]domain.Stock{{Ticker: "AAPL"}})
	mockRepo.On("CreateTransactionBuy", mock.Anything).Return(domain.Transaction{}, nil)

	result, err := uc.CreateTransactionBuy(domain.Transaction{})

	assert.Nil(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
	ucStockRepo.AssertExpectations(t)
}
