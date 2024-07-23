package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LoanFactoryOpts struct {
	Id          int32
	CustomerID  uuid.UUID
	Amount      decimal.Decimal
	Interest    decimal.Decimal
	StartDate   time.Time
	EndDate     time.Time
	DebtPayDate int
	CreatedAt   time.Time
	CreatedBy   uuid.UUID
	UpdatedAt   time.Time
	UpdatedBy   uuid.UUID
	DeletedAt   time.Time
	DeletedBy   uuid.UUID
}

func LoanFactory(opts LoanFactoryOpts) (*Loan, error) {
	loan := Loan{
		id:          opts.Id,
		customerID:  opts.CustomerID,
		amount:      opts.Amount,
		interest:    opts.Interest,
		startDate:   opts.StartDate,
		endDate:     opts.EndDate,
		debtPayDate: opts.DebtPayDate,
		createdAt:   opts.CreatedAt,
		createdBy:   opts.CreatedBy,
		updatedAt:   opts.UpdatedAt,
		updatedBy:   opts.UpdatedBy,
		deletedAt:   opts.DeletedAt,
		deletedBy:   opts.DeletedBy,
	}

	if err := loan.Validate(); err != nil {
		return nil, err
	}

	return &loan, nil
}

type PaymentFactoryOpts struct {
	Id             int64
	LoanId         int32
	MonthlyAmount  decimal.Decimal
	LoanAmount     decimal.Decimal
	InterestAmount decimal.Decimal
	LoanBalance    decimal.Decimal
	PayDate        time.Time
	NextPayDate    time.Time
	CreatedA       time.Time
	CreatedBy      uuid.UUID
	UpdatedAt      time.Time
	UpdatedBy      uuid.UUID
	DeletedAt      time.Time
	DeletedBy      uuid.UUID
}

func PaymentFactory(opts PaymentFactoryOpts) (*Payment, error) {
	payment := Payment{
		id:             opts.Id,
		loanId:         opts.LoanId,
		monthlyAmount:  opts.MonthlyAmount,
		loanAmount:     opts.LoanAmount,
		interestAmount: opts.InterestAmount,
		loanBalance:    opts.LoanBalance,
		payDate:        opts.PayDate,
		nextPayDate:    opts.NextPayDate,
		createdAt:      opts.CreatedA,
		createdBy:      opts.CreatedBy,
		updatedAt:      opts.UpdatedAt,
		updatedBy:      opts.UpdatedBy,
		deletedAt:      opts.DeletedAt,
		deletedBy:      opts.DeletedBy,
	}

	if err := payment.Validate(validation.Field(&payment.id, validation.Required)); err != nil {
		return nil, err
	}
	return &payment, nil
}
