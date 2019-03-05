package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	adrianConfig "Adrian2.0/config"
	adrianFonts "Adrian2.0/fonts"
	adrianServer "Adrian2.0/server"

	"github.com/labstack/echo"
)

func main() {

	log.Println("Starting Adrian 2.0")
	log.Println("Loading adrian.yaml")
	config := adrianConfig.LoadConfig("./adrian.yaml")
	log.Println("Initializing web server")
	e := adrianServer.Instantiate(config)
	log.Println("Loading fonts and starting watchers")
	for _, folder := range config.Global.Directories {
		adrianFonts.FindFonts(folder, config)
		adrianFonts.InstantiateWatcher(folder, config)
	}
	log.Println("Defining paths")

	e.GET("/font/css/:filename", func(c echo.Context) error {
		if filepath.Ext(c.Param("filename")) == ".css" {
			fontData, err := adrianFonts.GetFont(basename(c.Param("filename")))
			if err != nil {
				return adrianServer.Return404(c)
			}
			c.Response().Header().Set(echo.HeaderContentType, "text/css")
			return c.String(http.StatusOK, fontData.CSS)
		}

		return adrianServer.Return404(c)
	})

	log.Printf("Listening on port %d", config.Global.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Global.Port)))
}

// Basename gets the base filename (minus the last extension)
func basename(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n >= 0 {
		return s[:n]
	}
	return s
}
