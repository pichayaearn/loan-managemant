package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/pichayaearn/loan-management/pkg/customer/serializer"
)

type UpdateCustomerCfg struct {
	CustomerService model.CustomerService
}

func UpdateCustomer(cfg UpdateCustomerCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := serializer.BindUserIDFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Bind user id "+err.Error())
		}

		req := serializer.NewUpdateCustomerReq(userID)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		//create meeting
		if err := cfg.CustomerService.Update(model.UpdateCustomerOpts{
			CustomerID: req.CustomerID,
			Firstname:  req.Firstname,
			Lastname:   req.Lastname,
			Email:      req.Email,
			Mobile:     req.Mobile,
			UpdatedBy:  req.UpdatedBy,
		}); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Updated  "+err.Error())
		}

		return c.NoContent(http.StatusAccepted)
	}
}
