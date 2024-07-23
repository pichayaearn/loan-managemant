package model

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	digit = 2
)

type Payment struct {
	id             int64
	loanId         int32
	monthlyAmount  decimal.Decimal
	loanAmount     decimal.Decimal
	interestAmount decimal.Decimal
	loanBalance    decimal.Decimal
	payDate        time.Time
	nextPayDate    time.Time
	createdAt      time.Time
	createdBy      uuid.UUID
	updatedAt      time.Time
	updatedBy      uuid.UUID
	deletedAt      time.Time
	deletedBy      uuid.UUID
}

func (p Payment) ID() int64                       { return p.id }
func (p Payment) LoanID() int32                   { return p.loanId }
func (p Payment) MonthlyAmount() decimal.Decimal  { return p.monthlyAmount }
func (p Payment) LoanAmount() decimal.Decimal     { return p.loanAmount }
func (p Payment) InterestAmount() decimal.Decimal { return p.interestAmount }
func (p Payment) LoanBalance() decimal.Decimal    { return p.loanBalance }
func (p Payment) PayDate() time.Time              { return p.payDate }
func (p Payment) NextPayDate() time.Time          { return p.nextPayDate }
func (p Payment) CreatedAt() time.Time            { return p.createdAt }
func (p Payment) CreatedBy() uuid.UUID            { return p.createdBy }
func (p Payment) UpdatedAt() time.Time            { return p.updatedAt }
func (p Payment) UpdatedBy() uuid.UUID            { return p.updatedBy }
func (p Payment) DeletedAt() time.Time            { return p.deletedAt }
func (p Payment) DeletedBy() uuid.UUID            { return p.deletedBy }

func (p *Payment) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&p.monthlyAmount, validator.Required),
		validator.Field(&p.interestAmount, validator.Required),
		validator.Field(&p.loanAmount, validator.Required),
		validator.Field(&p.loanBalance, validator.Required),
		validator.Field(&p.payDate, validator.Required),
		validator.Field(&p.nextPayDate, validator.Required),
		validator.Field(&p.createdAt, validator.Required),
		validator.Field(&p.createdBy, validator.Required),
	}

	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}

	if err := validator.ValidateStruct(p, rules...); err != nil {
		return err
	}

	return nil
}

type NewPaymentOpts struct {
	LoanId         int32
	MonthlyAmount  decimal.Decimal
	LoanAmount     decimal.Decimal
	InterestAmount decimal.Decimal
	LoanBalance    decimal.Decimal
	DebtPayDate    int
	CreatedBy      uuid.UUID
}

func NewPayment(opts NewPaymentOpts) (*Payment, error) {
	now := time.Now()
	nextPayDate := now.AddDate(0, 0, opts.DebtPayDate)

	monthlyAmount, err := fixFormatDecimal(opts.MonthlyAmount)
	if err != nil {
		return nil, err
	}
	loanAmount, err := fixFormatDecimal(opts.LoanAmount)
	if err != nil {
		return nil, err
	}
	loanBalance, err := fixFormatDecimal(opts.LoanBalance)
	if err != nil {
		return nil, err
	}
	interest, err := fixFormatDecimal(opts.InterestAmount)
	if err != nil {
		return nil, err
	}
	payment := Payment{
		loanId:         opts.LoanId,
		monthlyAmount:  monthlyAmount,
		loanAmount:     loanAmount,
		interestAmount: interest,
		loanBalance:    loanBalance,
		payDate:        now,
		nextPayDate:    nextPayDate,
		createdAt:      now,
		createdBy:      opts.CreatedBy,
	}

	if err := payment.Validate(); err != nil {
		return nil, err
	}

	return &payment, nil
}

func fixFormatDecimal(value decimal.Decimal) (decimal.Decimal, error) {
	formattedValue := value.StringFixed(digit)
	finalValue, err := decimal.NewFromString(formattedValue)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return finalValue, nil
}
