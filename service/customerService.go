package service

import (
	"goHexagonal/domain"
	"goHexagonal/dto"
	errs "goHexagonal/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetById(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func(d DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	return d.repo.FindAll(status)
}

func (d DefaultCustomerService) GetById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := d.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}