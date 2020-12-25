package middleware

import (
	"github.com/labstack/echo"
	adrianConfig "github.com/dana-ross/adrian/config"
	"strconv"
)

// SetServerHeader sets a Server header
func SetServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Adrian 2.1.1")
		return next(c)
	}
}

// SetCacheControlHeaders sets headers for control over caching
func SetCacheControlHeaders(config adrianConfig.Config) func(echo.HandlerFunc) echo.HandlerFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "max-age=" + strconv.FormatUint(uint64(config.Global.CacheControlLifetime), 10) + ", public")
			return next(c)
		}
	}
}