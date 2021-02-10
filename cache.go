package main

import (
	"fmt"
	"log"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/labstack/echo"
)

var cache = fastcache.New(9048)

// ReadFromCache tries to load a cached response
func readFromCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if cache.Has([]byte(c.QueryParams().Encode())) {
			c.Response().Header().Set(echo.HeaderContentType, "text/css")
			_, err := c.Response().Write(cache.Get(nil, []byte(c.QueryParams().Encode())))
			if err != nil {
				log.Fatal(fmt.Sprintf("Error writing cached value to Response: %s", err))
			}
			return nil
		}
		return next(c)
	}
}

func writeToCache(c echo.Context, content string) {
	cache.Set([]byte(c.QueryParams().Encode()), []byte(content))
}
