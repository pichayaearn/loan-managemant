package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/loan-management/pkg/customer/model"
	"github.com/pichayaearn/loan-management/pkg/customer/serializer"
)

type GetCustomerCfg struct {
	CustomerService model.CustomerService
}

func GetCustomer(cfg GetCustomerCfg) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := new(serializer.GetCustomerReq)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		customer, err := cfg.CustomerService.GetByID(req.ID, c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Get  "+err.Error())
		}

		resp := serializer.ToGetCustomerResponse(*customer)

		return c.JSON(http.StatusOK, resp)
	}
}
