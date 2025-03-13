package domain

type Transaction struct {
	ID     string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Ticker string
	Amount float64
}
