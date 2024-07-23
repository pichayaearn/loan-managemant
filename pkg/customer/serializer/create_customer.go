package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type CreateCustomerReq struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`

	CreatedBy uuid.UUID `json:"-"`
}

func NewCreateCustomerReq(createdBy uuid.UUID) *CreateCustomerReq {
	return &CreateCustomerReq{
		CreatedBy: createdBy,
	}
}

func (req CreateCustomerReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.CreatedBy, validation.Required, is.UUIDv4),
		validation.Field(&req.Firstname, validation.Required),
		validation.Field(&req.Lastname, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Mobile, validation.Required),
	)
}
