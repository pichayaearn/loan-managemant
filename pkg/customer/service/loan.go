package service

import (
	"github.com/pichayaearn/loan-management/pkg/customer/model"
)

type LoanService struct {
	repo model.LoanRepo
}

type NewLoanServiceCfg struct {
	Repo model.LoanRepo
}

func NewLoanService(cfg NewLoanServiceCfg) *LoanService {
	return &LoanService{
		repo: cfg.Repo,
	}
}

func (ls LoanService) Create(req model.CreateLoanOpts) error {
	newLoan, err := model.NewLoan(req)
	if err != nil {
		return err
	}
	if err := ls.repo.Create(*newLoan); err != nil {
		return err
	}
	return nil
}
