package service

import (
	"goHexagonal/domain"
	"goHexagonal/dto"
	"goHexagonal/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func(d DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountID: "",
		CustomerId: request.CustomerId,
		OpenigDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount: request.Amount,
		Status: "1",
	}
	
	newAccount, err := d.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponse()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}