package service

import (
	"context"
	"log"

	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/shopspring/decimal"
)

type PaymentService struct {
	loanRepo    model.LoanRepo
	paymentRepo model.PaymentRepo
}

type NewPaymentServiceCfgs struct {
	LoanRepo    model.LoanRepo
	PaymentRepo model.PaymentRepo
}

func NewPaymentService(cfg NewPaymentServiceCfgs) *PaymentService {
	return &PaymentService{
		loanRepo:    cfg.LoanRepo,
		paymentRepo: cfg.PaymentRepo,
	}
}

func (ps PaymentService) Create(opts model.CreatePaymentOpts) error {
	//find loan by id
	loan, err := ps.loanRepo.GetByID(opts.LoanID, context.Background())
	if err != nil {
		return err
	}

	//find paymentExists by loanId is exists
	paymentExists, err := ps.paymentRepo.FindOnePaymentByLoanId(loan.ID(), context.Background())
	if err != nil {
		return err
	}

	var oldBalance decimal.Decimal
	if paymentExists != nil {
		oldBalance = paymentExists.LoanBalance()
	} else {
		oldBalance = loan.Amount()
	}

	interest := calculateInterest(loan.Interest(), oldBalance)
	loanAmountMonthly := calculateLoanAmount(interest, opts.MonthlyAmount)
	loanBalance := calculateLoanBalance(oldBalance, loanAmountMonthly)

	log.Printf("เงินต้น = %+v, ดอกเบี้ย = %+v, เงินต้น = %+v, เงินต้นคงเหลือ = %+v", oldBalance, interest, loanAmountMonthly, loanBalance)

	newPayment, err := model.NewPayment(model.NewPaymentOpts{
		LoanId:         opts.LoanID,
		MonthlyAmount:  opts.MonthlyAmount,
		LoanAmount:     loanAmountMonthly,
		InterestAmount: interest,
		LoanBalance:    loanBalance,
		CreatedBy:      opts.CreatedBy,
	})
	if err != nil {
		return err
	}

	if err := ps.paymentRepo.Create(*newPayment); err != nil {
		return err
	}
	return nil
}

// CalculateInterest calculates the interest portion of the monthly payment.
func calculateInterest(loan, interest decimal.Decimal) decimal.Decimal {
	monthlyRate := interest.Div(decimal.NewFromInt(100)).Div(decimal.NewFromInt(12))
	return loan.Mul(monthlyRate)
}

func calculateLoanAmount(interest decimal.Decimal, monthlyAmount decimal.Decimal) decimal.Decimal {
	return monthlyAmount.Sub(interest)
}

func calculateLoanBalance(loan decimal.Decimal, loanMonthly decimal.Decimal) decimal.Decimal {
	return loan.Sub(loanMonthly)
}
