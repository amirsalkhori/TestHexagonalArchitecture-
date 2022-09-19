package dto

type CustomerResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	ZipCode string `json:"zipcode"`
	DateOfBirth string `json:"dateOfBirth"`
	Status string `json:"status"`
}
