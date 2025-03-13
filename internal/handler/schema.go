package handler

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func positiveNumber(fl validator.FieldLevel) bool {
	return fl.Field().Float() > 0
}

func init() {
	validate.RegisterValidation("positive", positiveNumber)
}

type CreateTransactionRequest struct {
	Amount float64 `json:"amount" validate:"required,positive"`
	Ticker string  `json:"ticker" validate:"required,max=100"`
}

func (r *CreateTransactionRequest) Validate() error {
	return validate.Struct(r)
}
