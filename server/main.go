package server

import (
	"net/http"

	adrianConfig "Adrian2.0/config"
	adrianMiddleware "Adrian2.0/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Instantiate returns a pre-configured Echo instance
func Instantiate(config adrianConfig.Config) *echo.Echo {

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(adrianMiddleware.LockedDownCORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.Global.AllowedOrigins,
		AllowHeaders: []string{echo.HeaderOrigin},
		AllowMethods: []string{http.MethodGet},
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	e.Use(adrianMiddleware.SetServerHeader)

	return e

}
