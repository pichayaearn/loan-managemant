package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type CreatePaymentReq struct {
	LoanId        int32  `json:"loan_id"`
	PaymentAmount string `json:"payment_amount"`

	CreatedBy uuid.UUID `json:"-"`
}

func NewCreatePaymentReq(createdBy uuid.UUID) *CreatePaymentReq {
	return &CreatePaymentReq{
		CreatedBy: createdBy,
	}
}

func (req CreatePaymentReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.CreatedBy, validation.Required, is.UUIDv4),
		validation.Field(&req.LoanId, validation.Required),
		validation.Field(&req.PaymentAmount, validation.Required),
	)
}
