package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
)

type GetCustomerReq struct {
	ID uuid.UUID `query:"id"`
}

func (req GetCustomerReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ID, is.UUIDv4),
	)
}

type GetCustomerResponse struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
}

func ToGetCustomerResponse(customer model.Customer) GetCustomerResponse {
	return GetCustomerResponse{
		Firstname: customer.Firstname(),
		Lastname:  customer.Lastname(),
		Email:     customer.Email(),
		Mobile:    customer.Mobile(),
	}
}
