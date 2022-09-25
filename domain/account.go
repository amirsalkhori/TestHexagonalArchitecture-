package domain

import (
	"goHexagonal/dto"
	"goHexagonal/errs"
)

type Account struct {
	AccountID string `db:"account_id"`
	CustomerId string `db:"customer_id"`
	OpenigDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount string `db:"amount"`
	Status string `db:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func(a Account) ToNewAccountResponse() dto.NewAccountResponse{
	return dto.NewAccountResponse{a.AccountID}
}