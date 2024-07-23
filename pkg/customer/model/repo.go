package model

import (
	"context"

	"github.com/google/uuid"
)

type CustomerRepo interface {
	Create(customer Customer) error
	FindByCustomerId(customerID uuid.UUID, ctx context.Context) (*Customer, error)
	Update(custmer Customer) error
}

type LoanRepo interface {
	Create(loan Loan) error
	GetByID(id int32, ctx context.Context) (*Loan, error)
}

type PaymentRepo interface {
	FindOnePaymentByLoanId(loainId int32, ctx context.Context) (*Payment, error)
	Create(payment Payment) error
}
