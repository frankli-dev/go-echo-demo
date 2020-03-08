package http

import (
	"net/http"

	"github.com/ivan-marquez/golang-demo/pkg/listing"
	"github.com/labstack/echo/v4"
)

// getRenewableResources handler retrieves renewable resources data
// from the listing service interface.
func getRenewableResources(s listing.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/html; application/json")
		res, err := s.GetRenewableResources()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.Render(http.StatusOK, "renewables", res)
	}
}

// Handler sets up all app routes
func Handler(e *echo.Echo, s listing.Service) {
	e.GET("/renewables", getRenewableResources(s))
}
