package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/pichayaearn/loan-management/pkg/customer/serializer"
	"github.com/shopspring/decimal"
)

type CreatePaymentCfg struct {
	PaymentService model.PaymentService
}

func CreatePayment(cfg CreatePaymentCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := serializer.BindUserIDFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Bind user id "+err.Error())
		}

		req := serializer.NewCreatePaymentReq(userID)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		amount, err := decimal.NewFromString(req.PaymentAmount)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		if err := cfg.PaymentService.Create(model.CreatePaymentOpts{
			LoanID:        req.LoanId,
			MonthlyAmount: amount,
			CreatedBy:     req.CreatedBy,
		}); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Create  "+err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}
