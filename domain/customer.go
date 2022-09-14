package domain

import errs "goHexagonal/errs"

type Customer struct {
	Id string `json:"id" db:"customer_id"`
	Name string `json:"name" db:"name"`
	City string `json:"city" db:"city"`
	ZipCode string `json:"zip" db:"zip_code"`
	DateOfBirth string `json:"dateOfBirth" db:"date_of_birth"`
	Status string `json:"status" db:"status"`
}

type CustomerRepository interface{
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(id string) (*Customer, *errs.AppError)
}