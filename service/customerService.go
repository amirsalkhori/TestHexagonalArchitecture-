package service

import "goHexagonal/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetById(id string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func(d DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return d.repo.FindAll()
}

func (d DefaultCustomerService) GetById(id string) (*domain.Customer, error) {
	return d.repo.FindById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}