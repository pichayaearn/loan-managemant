package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/pichayaearn/loan-management/pkg/customer/model/mocks"
	"github.com/pichayaearn/loan-management/pkg/customer/serializer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateCustomer(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		name        string
		parameter   serializer.CreateCustomerReq
		mockService func(m *mocks.CustomerService)
		isError     bool
	}{
		{
			name: "failed, invalid parameter",
			parameter: serializer.CreateCustomerReq{
				Firstname: "",
				Lastname:  "",
				Email:     "",
				Mobile:    "",
				CreatedBy: uuid.UUID{},
			},
			isError: true,
		},
		{
			name: "failed, invalid parameter email is not email type",
			parameter: serializer.CreateCustomerReq{
				Firstname: "Test",
				Lastname:  "Test",
				Email:     "Test",
				Mobile:    "0808988888",
				CreatedBy: uuid.New(),
			},
			isError: true,
		},
		{
			name: "failed, error from create customer service",
			parameter: serializer.CreateCustomerReq{
				Firstname: "Test",
				Lastname:  "Test",
				Email:     "test@gmail.com",
				Mobile:    "0869522373",
				CreatedBy: uuid.New(),
			},
			mockService: func(m *mocks.CustomerService) {
				m.On("Create", mock.IsType(model.CreateCustomerOpts{})).Return(errors.New("error from service"))
			},
			isError: true,
		},
		{
			name: "success",
			parameter: serializer.CreateCustomerReq{
				Firstname: "Test",
				Lastname:  "Test",
				Email:     "test@gmail.com",
				Mobile:    "0869522373",
				CreatedBy: uuid.New(),
			},
			mockService: func(m *mocks.CustomerService) {
				m.On("Create", mock.IsType(model.CreateCustomerOpts{})).Return(nil)
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.parameter)
			r.NoError(err)
			req := httptest.NewRequest(http.MethodPost, "/customer", strings.NewReader(string(reqBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)
			c.Set("ActionBy", tt.parameter.CreatedBy.String())

			mockService := mocks.NewCustomerService(t)
			if tt.mockService != nil {
				tt.mockService(mockService)
			}

			err = CreateCustomer(CreateCustomerCfg{
				CustomerService: mockService,
			})(c)

			if tt.isError {
				r.Error(err)
			} else {
				r.NoError(err)
				r.Equal(http.StatusCreated, rec.Result().StatusCode)
			}
		})
	}
}
