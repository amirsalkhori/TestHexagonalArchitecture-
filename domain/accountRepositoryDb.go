package domain

import (
	"goHexagonal/errs"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func(account AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, status) VALUES (?, ?, ?, ?)"
	result, err := account.client.Exec(sqlInsert, a.CustomerId, a.OpenigDate, a.AccountType, a.Status)
	if err != nil {
		log.Fatal("Error while insert data into db", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Error while getting data into account table", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb{
	return AccountRepositoryDb{dbClient}
}