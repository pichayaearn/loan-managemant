package model

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

const (
	CustomerStatusActive   string = "active"
	CustomerStatusInActive string = "inactive"
)

type Customer struct {
	id        uuid.UUID
	firtname  string
	lastname  string
	mobile    string
	email     string
	status    string
	createdAt time.Time
	createdBy uuid.UUID
	updatedAt time.Time
	updatedBy uuid.UUID
	deletedAt time.Time
	deletedBy uuid.UUID
}

func (m Customer) ID() uuid.UUID        { return m.id }
func (m Customer) Firstname() string    { return m.firtname }
func (m Customer) Lastname() string     { return m.lastname }
func (m Customer) Mobile() string       { return m.mobile }
func (m Customer) Email() string        { return m.email }
func (m Customer) Status() string       { return m.status }
func (m Customer) CreatedAt() time.Time { return m.createdAt }
func (m Customer) CreatedBy() uuid.UUID { return m.createdBy }
func (m Customer) UpdatedAt() time.Time { return m.updatedAt }
func (m Customer) UpdatedBy() uuid.UUID { return m.updatedBy }
func (m Customer) DeletedAt() time.Time { return m.deletedAt }
func (m Customer) DeletedBy() uuid.UUID { return m.deletedBy }

func (m *Customer) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&m.id, validator.Required),
		validator.Field(&m.firtname, validator.Required),
		validator.Field(&m.lastname, validator.Required),
		validator.Field(&m.mobile, validator.Required),
		validator.Field(&m.email, validator.Required, is.Email),
		validator.Field(&m.status, validator.Required, validator.In(CustomerStatusActive, CustomerStatusInActive)),
		validator.Field(&m.createdAt, validator.Required),
		validator.Field(&m.createdBy, validator.Required),
	}

	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}

	if err := validator.ValidateStruct(m, rules...); err != nil {
		return err
	}

	return nil
}

func (c *Customer) SetFirstname(firstname string, updatedBy uuid.UUID) error {
	c.firtname = firstname
	c.updatedBy = updatedBy
	c.updatedAt = time.Now()
	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}

func (c *Customer) SetLastname(lastname string, updatedBy uuid.UUID) error {
	c.lastname = lastname
	c.updatedBy = updatedBy
	c.updatedAt = time.Now()
	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}

func (c *Customer) SetMobile(mobile string, updatedBy uuid.UUID) error {
	c.mobile = mobile
	c.updatedBy = updatedBy
	c.updatedAt = time.Now()
	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}

func (c *Customer) SetEmail(email string, updatedBy uuid.UUID) error {
	c.email = email
	c.updatedBy = updatedBy
	c.updatedAt = time.Now()
	if err := c.Validate(); err != nil {
		return err
	}
	return nil
}
func NewCustomer(opts CreateCustomerOpts) *Customer {
	return &Customer{
		firtname:  opts.Firstname,
		lastname:  opts.Lastname,
		mobile:    opts.Mobile,
		email:     opts.Email,
		status:    CustomerStatusActive,
		createdAt: time.Now(),
		createdBy: opts.CreatedBy,
	}
}
