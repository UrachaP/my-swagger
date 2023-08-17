package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheckHandler godoc
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @produce json
// @response 200 {string} string "OK"
// @router /healthcheck [get]
func HealthCheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
