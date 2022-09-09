package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func(c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "John", ZipCode: "123", DateOfBirth: "2013-10-1", Status: "true"},
		{Id: "2", Name: "Amir", ZipCode: "123", DateOfBirth: "2013-10-1", Status: "true"},
	}
	return CustomerRepositoryStub{customers}
}