package v1

import (
	"github.com/labstack/echo"
	"net/http"
)

type healthCheck struct {
}

func NewHealthCheck(e *echo.Echo) *healthCheck {
	h := &healthCheck{}
	e.GET("/health", h.healthCheck)

	return h
}

func (h *healthCheck) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"status": "UP"})
}
