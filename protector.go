package protector

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var blockedMethods = []string{"POST", "PUT", "PATCH", "DELETE"}

func ProtectorMiddleware(isReadOnly func() bool) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if isReadOnly() {
				for _, s := range blockedMethods {
					if s == c.Request().Method {
						return echo.NewHTTPError(http.StatusServiceUnavailable, "Service is currently in read-only mode")
					}
				}
			}
			return next(c)
		}
	}
}
