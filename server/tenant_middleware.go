package server

import (
	"context"
	"net/http"
	"regexp"

	"github.com/jSierraB3991/golang-multitenant/libs"
	"github.com/labstack/echo/v4"
)

func TenantMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tenantID := c.Request().Header.Get("X-Tenant-ID")
		if tenantID == "" {
			tenantID = "public"
		}

		matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, tenantID)
		if !matched {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid tenant ID format")
		}

		// Inyectar tenant en el contexto de la request
		ctx := context.WithValue(c.Request().Context(), libs.ContextTenantKey, tenantID)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
