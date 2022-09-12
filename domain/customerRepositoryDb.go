package domain

import (
	"database/sql"
	errs "goHexagonal/errs"
	"goHexagonal/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct{
	client *sql.DB
}

func(c CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError){	
	var findAllSql string
	var rows *sql.Rows
	var err error
	if status == ""{
		findAllSql = "SELECT * FROM customers"
		rows, err = c.client.Query(findAllSql)
	}else{
		findAllSql = "SELECT * FROM customers WHERE status = ?"
		rows, err = c.client.Query(findAllSql, status)
	}
	if err != nil{
		logger.Error("Error while query customer " + err.Error())
		return nil, errs.NewNotFoundError("Unexpected database error")
	}
	customers := make([]Customer, 0)
	for rows.Next(){
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scanning customers table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		} 
		customers = append(customers, c)
	}

	return customers, nil
}

func (c CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError){
	findById := "SELECT * FROM customers WHERE customer_id = ?"
	row := c.client.QueryRow(findById, id)
	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.DateOfBirth, &customer.ZipCode, &customer.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while query customer table" + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		}else{
			logger.Error("Error while query customer table" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb{
	client, err := sql.Open("mysql", "rooti:changeMe@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)		
	return CustomerRepositoryDb{client}
}