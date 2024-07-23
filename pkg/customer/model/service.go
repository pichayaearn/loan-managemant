package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateCustomerOpts struct {
	Firstname string
	Lastname  string
	Mobile    string
	Email     string
	CreatedBy uuid.UUID
}

type CustomerService interface {
	Create(opts CreateCustomerOpts) error
}

type CreateLoanOpts struct {
	CustomerID  uuid.UUID
	Amount      decimal.Decimal
	Interest    decimal.Decimal
	StartDate   time.Time
	Period      int
	Unit        string
	DebtPayDate int
	CreatedBy   uuid.UUID
}

type LoanService interface {
	Create(opts CreateLoanOpts) error
}

type CreatePaymentOpts struct {
	MonthlyAmount decimal.Decimal
	LoanID        int32
	CreatedBy     uuid.UUID
}
type PaymentService interface {
	Create(opts CreatePaymentOpts) error
}
