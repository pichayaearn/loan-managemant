package route

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/pichayaearn/loan-management/pkg/customer/model/mocks"
	"github.com/pichayaearn/loan-management/pkg/customer/serializer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetCustomer(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		name        string
		parameter   serializer.GetCustomerReq
		mockService func(m *mocks.CustomerService)
		isError     bool
	}{
		{
			name:      "failed, invalid parameter",
			parameter: serializer.GetCustomerReq{},
			isError:   true,
		},

		{
			name: "failed, error from customer service",
			parameter: serializer.GetCustomerReq{
				ID: uuid.New(),
			},
			mockService: func(m *mocks.CustomerService) {
				m.On("GetByID", mock.IsType(uuid.UUID{}), context.Background()).Return(nil, errors.New("error from service"))
			},
			isError: true,
		},
		{
			name: "success",
			parameter: serializer.GetCustomerReq{
				ID: uuid.New(),
			},
			mockService: func(m *mocks.CustomerService) {
				m.On("GetByID", mock.IsType(uuid.UUID{}), context.Background()).Return(&model.Customer{}, nil)
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/customer", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			q := req.URL.Query()
			q.Set("id", tt.parameter.ID.String())
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)
			c.Set("ActionBy", uuid.New())

			mockService := mocks.NewCustomerService(t)
			if tt.mockService != nil {
				tt.mockService(mockService)
			}

			err := GetCustomer(GetCustomerCfg{
				CustomerService: mockService,
			})(c)

			if tt.isError {
				r.Error(err)
			} else {
				r.NoError(err)
				r.NotEmpty(rec.Body)
				r.Equal(http.StatusOK, rec.Result().StatusCode)

			}
		})
	}
}
