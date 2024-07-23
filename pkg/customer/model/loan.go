package model

import (
	"errors"
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	Month = "month"
	Year  = "year"
	Day   = "day"
)

type Loan struct {
	id          int32
	customerID  uuid.UUID
	amount      decimal.Decimal
	interest    decimal.Decimal
	startDate   time.Time
	endDate     time.Time
	debtPayDate int
	createdAt   time.Time
	createdBy   uuid.UUID
	updatedAt   time.Time
	updatedBy   uuid.UUID
	deletedAt   time.Time
	deletedBy   uuid.UUID
}

func (l Loan) ID() int32                 { return l.id }
func (l Loan) CustomerID() uuid.UUID     { return l.customerID }
func (l Loan) Amount() decimal.Decimal   { return l.amount }
func (l Loan) Interest() decimal.Decimal { return l.interest }
func (l Loan) StartDate() time.Time      { return l.startDate }
func (l Loan) EndDate() time.Time        { return l.endDate }
func (l Loan) DebtPayDate() int          { return l.debtPayDate }
func (l Loan) CreatedAt() time.Time      { return l.createdAt }
func (l Loan) CreatedBy() uuid.UUID      { return l.createdBy }
func (l Loan) UpdatedAt() time.Time      { return l.updatedAt }
func (l Loan) UpdatedBy() uuid.UUID      { return l.updatedBy }
func (l Loan) DeletedAt() time.Time      { return l.deletedAt }
func (l Loan) DeletedBy() uuid.UUID      { return l.deletedBy }

func (l *Loan) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&l.id, validator.Required),
		validator.Field(&l.customerID, validator.Required),
		validator.Field(&l.amount, validator.Required),
		validator.Field(&l.startDate, validator.Required),
		validator.Field(&l.endDate, validator.Required),
		validator.Field(&l.createdAt, validator.Required),
		validator.Field(&l.createdBy, validator.Required),
	}

	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}

	if err := validator.ValidateStruct(l, rules...); err != nil {
		return err
	}

	return nil
}
func NewLoan(opts CreateLoanOpts) (*Loan, error) {
	var endDate time.Time
	switch opts.Unit {
	case Day:
		endDate = opts.StartDate.AddDate(0, 0, opts.Period)
	case Month:
		endDate = opts.StartDate.AddDate(0, opts.Period, 0)
	case Year:
		endDate = opts.StartDate.AddDate(opts.Period, 0, 0)
	default:
	}

	if endDate.IsZero() {
		return nil, errors.New("create new loan failed because end date is empty")
	}
	return &Loan{
		customerID:  opts.CustomerID,
		amount:      opts.Amount,
		interest:    opts.Interest,
		startDate:   opts.StartDate,
		endDate:     endDate,
		debtPayDate: opts.DebtPayDate,
		createdAt:   time.Now(),
		createdBy:   opts.CreatedBy,
	}, nil
}
