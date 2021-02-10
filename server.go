package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Instantiate returns a pre-configured Echo instance
func Instantiate(config Config) *echo.Echo {

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.Global.AllowedOrigins,
		AllowHeaders: []string{echo.HeaderOrigin},
		AllowMethods: []string{http.MethodGet},
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(SetServerHeader)
	e.Use(SetCacheControlHeaders(config))

	e.Use(readFromCache)

	return e

}

// return404 sends an appropriate message back to the browser on a 404
func return404(c echo.Context) error {
	status := make(map[string]string)
	status["message"] = "Not Found"
	return c.JSON(http.StatusNotFound, status)
}
