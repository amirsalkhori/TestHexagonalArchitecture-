package service

import (
	"goHexagonal/domain"
	errs "goHexagonal/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetById(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func(d DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return d.repo.FindAll()
}

func (d DefaultCustomerService) GetById(id string) (*domain.Customer, *errs.AppError) {
	return d.repo.FindById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}