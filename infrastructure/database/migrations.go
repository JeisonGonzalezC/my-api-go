package database

import (
	"log"

	"gorm.io/gorm"
)

type Transaction struct {
	ID     string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Ticker string  `gorm:"type:varchar(10);not null"`
	Amount float64 `gorm:"not null"`
}

type Stock struct {
	Ticker       string        `gorm:"primaryKey;type:varchar(10);not null"`
	TargetFrom   string        `gorm:"type:varchar(10);not null"`
	TargetTo     string        `gorm:"type:varchar(10);not null"`
	Company      string        `gorm:"type:varchar(255);not null"`
	Action       string        `gorm:"type:varchar(50);not null"`
	Brokerage    string        `gorm:"type:varchar(100);not null"`
	RatingFrom   string        `gorm:"type:varchar(50);not null"`
	RatingTo     string        `gorm:"type:varchar(50);not null"`
	Time         string        `gorm:"type:timestamptz;not null"`
	Transactions []Transaction `gorm:"foreignKey:Ticker;references:Ticker;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&Stock{})
	if err != nil {
		log.Fatalf("Error migrating Stock: %v", err)
	}

	err = db.AutoMigrate(&Transaction{})
	if err != nil {
		log.Fatalf("Error migrating Transaction: %v", err)
	}

	if err != nil {
		log.Fatalf("Error migration: %v", err)
	}
	log.Println("Migration Stock and Transaction tables created")
}
