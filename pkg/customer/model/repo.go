package model

import "context"

type CustomerRepo interface {
	Create(customer Customer) error
}

type LoanRepo interface {
	Create(loan Loan) error
	GetByID(id int32, ctx context.Context) (*Loan, error)
}

type PaymentRepo interface {
	FindOnePaymentByLoanId(loainId int32, ctx context.Context) (*Payment, error)
	Create(payment Payment) error
}
