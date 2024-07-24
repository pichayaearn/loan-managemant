package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type UpdatedCustomerReq struct {
	CustomerID uuid.UUID `json:"customer_id"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Email      string    `json:"email"`
	Mobile     string    `json:"mobile"`

	UpdatedBy uuid.UUID `json:"-"`
}

func NewUpdateCustomerReq(updatedBy uuid.UUID) *UpdatedCustomerReq {
	return &UpdatedCustomerReq{
		UpdatedBy: updatedBy,
	}
}

func (req UpdatedCustomerReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.CustomerID, validation.Required, is.UUIDv4),
		validation.Field(&req.Email, is.Email),
		validation.Field(&req.UpdatedBy, validation.Required, is.UUIDv4),
	)
}
