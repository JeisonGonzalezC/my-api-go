package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type StockItemResponse struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type StocksResponse struct {
	Items    []StockItemResponse
	NextPage string `json:"next_page"`
}

// StockAPIClient is a client for stock API
type StockAPIClient struct {
	BaseURL string
	Token   string
}

// NewStockAPIClient is a constructor for StockAPIClient
func NewStockAPIClient() *StockAPIClient {
	return &StockAPIClient{
		BaseURL: os.Getenv("API_URL"),
		Token:   os.Getenv("API_TOKEN"),
	}
}

func (c *StockAPIClient) GetStocksFromAPI(nextPage string) (StocksResponse, error) {
	url := c.BaseURL
	if nextPage != "" {
		url = fmt.Sprintf("%s?next_page=%s", c.BaseURL, nextPage)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return StocksResponse{}, fmt.Errorf("error request API: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return StocksResponse{}, fmt.Errorf("error request HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return StocksResponse{}, fmt.Errorf("error API request data, code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return StocksResponse{}, fmt.Errorf("error reading response: %w", err)
	}

	var response StocksResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return StocksResponse{}, fmt.Errorf("error JSON parse: %w", err)
	}

	return response, nil
}
