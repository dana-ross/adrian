package main

import (
	"strconv"

	"github.com/labstack/echo"
)

// SetServerHeader sets a Server header
func SetServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Adrian 3.0.0")
		return next(c)
	}
}

// SetCacheControlHeaders sets headers for control over caching
func SetCacheControlHeaders(config Config) func(echo.HandlerFunc) echo.HandlerFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "max-age=" + strconv.FormatUint(uint64(config.Global.CacheControlLifetime), 10) + ", public")
			return next(c)
		}
	}
}