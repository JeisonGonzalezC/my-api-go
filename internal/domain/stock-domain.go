package domain

import (
	"myapi/pkg"
	"strings"
)

type Stock struct {
	Ticker       string
	TargetFrom   string
	TargetTo     string
	Company      string
	Action       string
	Brokerage    string
	RatingFrom   string
	RatingTo     string
	Time         string
	Recommended  bool          `gorm:"-"` // not stored in database
	Transactions []Transaction `gorm:"foreignKey:Ticker;references:Ticker"`
}

func (s *Stock) EvaluateRecommendation() {
	priceIncrease := pkg.FromStringToPositiveNumber(s.TargetTo) > pkg.FromStringToPositiveNumber(s.TargetFrom)

	ratingImproved := s.isRatingUpgraded()

	actionPositive := strings.Contains(s.Action, "raised") || strings.Contains(s.Action, "upgraded") || strings.Contains(s.Action, "initiated")

	s.Recommended = (priceIncrease && ratingImproved) || (priceIncrease && actionPositive) || (ratingImproved && actionPositive)
}

func (s *Stock) isRatingUpgraded() bool {
	ratingOrder := map[string]int{
		"Sell":         1,
		"Underweight":  2,
		"Hold":         3,
		"Neutral":      3,
		"Equal Weight": 3,
		"Buy":          4,
		"Overweight":   4,
	}

	return ratingOrder[s.RatingTo] > ratingOrder[s.RatingFrom]
}
