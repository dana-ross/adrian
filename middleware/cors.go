package middleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	coreMiddleware "github.com/labstack/echo/middleware"
)

// LockedDownCORSWithConfig is good
func LockedDownCORSWithConfig(config coreMiddleware.CORSConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		CORSWithConfig := coreMiddleware.CORSWithConfig(config)
		CORSWithConfigHandler := CORSWithConfig(next)

		return func(c echo.Context) error {

			req := c.Request()
			origin := req.Header.Get(echo.HeaderOrigin)
			allowOrigin := ""

			// Check allowed origins
			for _, o := range config.AllowOrigins {
				if o == "*" && config.AllowCredentials {
					allowOrigin = origin
					break
				}
				if o == "*" || o == origin {
					allowOrigin = o
					break
				}
			}

			log.Printf("allowOrigin is %s", allowOrigin)

			if allowOrigin == "" {
				return c.NoContent(http.StatusForbidden)
			}

			return CORSWithConfigHandler(c)

		}
	}
}
