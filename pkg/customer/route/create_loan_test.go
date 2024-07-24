package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/pichayaearn/loan-management/pkg/customer/model/mocks"
	"github.com/pichayaearn/loan-management/pkg/customer/serializer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateLoan(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		name        string
		parameter   serializer.CreateLoanReq
		mockService func(m *mocks.LoanService)
		isError     bool
	}{
		{
			name: "failed, invalid parameter",
			parameter: serializer.CreateLoanReq{
				CustomerID: uuid.New(),
				Amount:     "",
				Interest:   "3",
				StartDate:  time.Now(),
				Period:     2,
				Unit:       "month",
				CreatedBy:  uuid.New(),
			},
			isError: true,
		},
		{
			name: "failed, can not pass amount to decimal",
			parameter: serializer.CreateLoanReq{
				CustomerID: uuid.New(),
				Amount:     "xxx",
				Interest:   "3",
				StartDate:  time.Now(),
				Period:     2,
				Unit:       "month",
				CreatedBy:  uuid.New(),
			},
			isError: true,
		},
		{
			name: "failed, can not pass interest to decimal",
			parameter: serializer.CreateLoanReq{
				CustomerID: uuid.New(),
				Amount:     "10000",
				Interest:   "x",
				StartDate:  time.Now(),
				Period:     2,
				Unit:       "month",
				CreatedBy:  uuid.New(),
			},
			isError: true,
		},
		{
			name: "failed, error from loan service",
			parameter: serializer.CreateLoanReq{
				CustomerID: uuid.New(),
				Amount:     "10000",
				Interest:   "3",
				StartDate:  time.Now(),
				Period:     2,
				Unit:       "month",
				CreatedBy:  uuid.New(),
			},
			mockService: func(m *mocks.LoanService) {
				m.On("Create", mock.IsType(model.CreateLoanOpts{})).Return(errors.New("error from create loan service"))
			},
			isError: true,
		},
		{
			name: "success",
			parameter: serializer.CreateLoanReq{
				CustomerID: uuid.New(),
				Amount:     "10000",
				Interest:   "3",
				StartDate:  time.Now(),
				Period:     2,
				Unit:       "month",
				CreatedBy:  uuid.New(),
			},
			mockService: func(m *mocks.LoanService) {
				m.On("Create", mock.IsType(model.CreateLoanOpts{})).Return(nil)
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.parameter)
			r.NoError(err)
			req := httptest.NewRequest(http.MethodPost, "/customer/loan", strings.NewReader(string(reqBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)
			c.Set("ActionBy", tt.parameter.CreatedBy.String())

			mockService := mocks.NewLoanService(t)
			if tt.mockService != nil {
				tt.mockService(mockService)
			}

			err = CreateLoan(CreateLoanCfg{
				LoanService: mockService,
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
