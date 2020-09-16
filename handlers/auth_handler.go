package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//HealthHandler will send a true message if app is reachable.
func HealthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "True")
}
