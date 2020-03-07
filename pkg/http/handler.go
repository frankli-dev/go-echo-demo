package http

import (
	"net/http"

	"github.com/ivan-marquez/golang-demo/pkg/listing"
	"github.com/labstack/echo/v4"
)

// TODO: add comment
func getRenewableResources(s listing.Service) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/html")
		// TODO: add error handling
		res := s.GetRenewableResources()
		return c.Render(http.StatusOK, "renewables", res)
	}
}

// TODO: add comment
func Handler(e *echo.Echo, s listing.Service) {
	e.GET("/renewables", getRenewableResources(s))
}
