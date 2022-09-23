package domain

import (
	"goHexagonal/dto"
	errs "goHexagonal/errs"
)

type Customer struct {
	Id string `db:"customer_id"`
	Name string `db:"name"`
	City string `db:"city"`
	ZipCode string `db:"zip_code"`
	DateOfBirth string `db:"date_of_birth"`
	Status string `db:"status"`
}

func (c Customer) statusAsText() string{
	statusAsText := "active"
	if c.Status == "inactive" {
		statusAsText = "inactive"
	}

	return statusAsText 
}
func(c Customer) ToDto() dto.CustomerResponse{
	
	response := dto.CustomerResponse{
		Id: 			c.Id,
		Name: 			c.Name,
		City: 			c.City,
		ZipCode: 		c.ZipCode,
		DateOfBirth: 	c.DateOfBirth,
		Status: 		c.statusAsText(),
	}

	return response
}
 
type CustomerRepository interface{
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(id string) (*Customer, *errs.AppError)
}