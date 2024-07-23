package serializer

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
)

type CreateLoanReq struct {
	CustomerID uuid.UUID `json:"customer_id"`
	Amount     string    `json:"amount"`
	Interest   string    `json:"interest"`
	StartDate  time.Time `json:"start_date"`
	Period     int       `json:"period"`
	Unit       string    `json:"unit"`

	CreatedBy uuid.UUID `json:"-"`
}

func NewCreateLoanReq(createdBy uuid.UUID) *CreateLoanReq {
	return &CreateLoanReq{
		CreatedBy: createdBy,
	}
}

func (req CreateLoanReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.CreatedBy, validation.Required, is.UUIDv4),
		validation.Field(&req.CustomerID, validation.Required, is.UUIDv4),
		validation.Field(&req.Amount, validation.Required),
		validation.Field(&req.Interest, validation.Required),
		validation.Field(&req.StartDate, validation.Required),
		validation.Field(&req.Period, validation.Required),
		validation.Field(&req.Unit, validation.Required, validation.In(model.Day, model.Month, model.Year)),
	)
}
