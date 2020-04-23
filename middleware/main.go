package middleware

import (
	"github.com/labstack/echo"
)

// SetServerHeader sets a Server header
func SetServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Adrian 2.1.1")
		return next(c)
	}
}

// SetCacheControlHeaders sets headers for control over caching
func SetCacheControlHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "max-age=2628000, public")
		return next(c)
	}
}
