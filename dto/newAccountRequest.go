package dto

import (
	"goHexagonal/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId string `json:"customer_id"`
	AccountType string `json:"account_type"`
	Amount string `json:"amount"`
}

func(req NewAccountRequest) Validate() *errs.AppError{
	if strings.ToLower(req.AccountType) != "saveing" && strings.ToLower(req.AccountType) != "checking"{
		return errs.NewValidationError("Amount type should be saving or checking type")
	}

	return nil
}