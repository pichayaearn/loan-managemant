package service

import (
	"context"
	"errors"
	"log"

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

func (cs CustomerService) Update(opts model.UpdateCustomerOpts) error {
	log.Printf("opts %+v", opts)
	customer, err := cs.customerRepo.FindByCustomerId(opts.CustomerID, context.Background())
	if err != nil {
		return err
	}

	if customer == nil {
		return errors.New("customer not found")
	}

	if opts.Firstname != "" {
		if err := customer.SetFirstname(opts.Firstname, opts.UpdatedBy); err != nil {
			return err
		}
	}

	if opts.Lastname != "" {
		if err := customer.SetLastname(opts.Lastname, opts.UpdatedBy); err != nil {
			return err
		}
	}

	if opts.Mobile != "" {
		if err := customer.SetMobile(opts.Mobile, opts.UpdatedBy); err != nil {
			return err
		}
	}

	if opts.Email != "" {
		if err := customer.SetEmail(opts.Email, opts.UpdatedBy); err != nil {
			return err
		}
	}

	if err := cs.customerRepo.Update(*customer); err != nil {
		return err
	}
	return nil

}
