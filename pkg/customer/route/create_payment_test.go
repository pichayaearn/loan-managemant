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

func TestCreatePayment(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		name        string
		parameter   serializer.CreatePaymentReq
		mockService func(m *mocks.PaymentService)
		isError     bool
	}{
		{
			name: "failed, invalid parameter",
			parameter: serializer.CreatePaymentReq{
				LoanId:        1,
				PaymentAmount: "",
				CreatedBy:     uuid.New(),
			},
			isError: true,
		},
		{
			name: "failed, pass amount to decimal",
			parameter: serializer.CreatePaymentReq{
				LoanId:        1,
				PaymentAmount: "x",
				CreatedBy:     uuid.New(),
			},
			isError: true,
		},
		{
			name: "failed, error from payment service",
			parameter: serializer.CreatePaymentReq{
				LoanId:        1,
				PaymentAmount: "3000",
				CreatedBy:     uuid.New(),
			},
			mockService: func(m *mocks.PaymentService) {
				m.On("Create", mock.IsType(model.CreatePaymentOpts{})).Return(errors.New("error from payment service"))
			},
			isError: true,
		},
		{
			name: "failed, error from payment service",
			parameter: serializer.CreatePaymentReq{
				LoanId:        1,
				PaymentAmount: "3000",
				CreatedBy:     uuid.New(),
			},
			mockService: func(m *mocks.PaymentService) {
				m.On("Create", mock.IsType(model.CreatePaymentOpts{})).Return(errors.New("error from payment service"))
			},
			isError: true,
		},
		{
			name: "success",
			parameter: serializer.CreatePaymentReq{
				LoanId:        1,
				PaymentAmount: "3000",
				CreatedBy:     uuid.New(),
			},
			mockService: func(m *mocks.PaymentService) {
				m.On("Create", mock.IsType(model.CreatePaymentOpts{})).Return(nil)
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.parameter)
			r.NoError(err)
			req := httptest.NewRequest(http.MethodPost, "/customer/payment", strings.NewReader(string(reqBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)
			c.Set("ActionBy", tt.parameter.CreatedBy.String())

			mockService := mocks.NewPaymentService(t)
			if tt.mockService != nil {
				tt.mockService(mockService)
			}

			err = CreatePayment(CreatePaymentCfg{
				PaymentService: mockService,
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
