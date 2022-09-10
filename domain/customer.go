package domain

import errs "goHexagonal/errs"

type Customer struct {
	Id string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	ZipCode string `json:"zip"`
	DateOfBirth string `json:"dateOfBirth"`
	Status string `json:"status"`
}

type CustomerRepository interface{
	FindAll() ([]Customer, error)
	FindById(id string) (*Customer, *errs.AppError)
}