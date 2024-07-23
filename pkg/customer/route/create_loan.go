package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/pichayaearn/loan-management/pkg/customer/serializer"
	"github.com/shopspring/decimal"
)

type CreateLoanCfg struct {
	LoanService model.LoanService
}

func CreateLoan(cfg CreateLoanCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := serializer.BindUserIDFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Bind user id "+err.Error())
		}

		req := serializer.NewCreateLoanReq(userID)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		amount, err := decimal.NewFromString(req.Amount)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		interest, err := decimal.NewFromString(req.Interest)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		//create meeting
		if err := cfg.LoanService.Create(model.CreateLoanOpts{
			CustomerID: req.CustomerID,
			Amount:     amount,
			Interest:   interest,
			StartDate:  req.StartDate,
			Period:     req.Period,
			Unit:       req.Unit,
			CreatedBy:  req.CreatedBy,
		}); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Create  "+err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}
