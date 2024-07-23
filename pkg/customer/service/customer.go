package service

import (
	"github.com/pichayaearn/loan-management/pkg/customer/model"
)

type CustomerService struct {
	customerRepo model.CustomerRepo
}

type NewCustomerServiceCfg struct {
	CustomerRepo model.CustomerRepo
}

func NewCustomerService(cfg NewCustomerServiceCfg) model.CustomerService {
	return &CustomerService{
		customerRepo: cfg.CustomerRepo,
	}
}

func (cs CustomerService) Create(opts model.CreateCustomerOpts) error {
	customer := model.NewCustomer(opts)
	if err := cs.customerRepo.Create(*customer); err != nil {
		return err
	}
	return nil
}
