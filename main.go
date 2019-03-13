package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	adrianConfig "github.com/daveross/adrian/config"
	adrianFonts "github.com/daveross/adrian/fonts"
	adrianServer "github.com/daveross/adrian/server"

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

	e.GET("/css/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, "text/css")
		fontFilenames := strings.Split(c.QueryParam("family"), "|")
		var fontsCSS string
		for _, fontFilename := range fontFilenames {
			fontData, err := adrianFonts.GetFont(fontFilename)
			if err != nil {
				return adrianServer.Return404(c)
			}
			fontsCSS = fontsCSS + "\n" + fontData.CSS
			}
		return c.String(http.StatusOK, fontsCSS)
	})

	e.GET("/font/:filename", func(c echo.Context) error {
		switch filepath.Ext(c.Param("filename")) {
		case ".ttf":
			return outputFont(c, "font/ttf")
		case ".woff":
			return outputFont(c, "font/woff")
		case ".woff2":
			return outputFont(c, "font/woff2")
		case ".otf":
			return outputFont(c, "font/otf")
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

func outputFont(c echo.Context, mimeType string) error {
	fontData, err := adrianFonts.GetFontByUniqueID(basename(c.Param("filename")))
	if err != nil {
		return adrianServer.Return404(c)
	}
	fontBinary, err := ioutil.ReadFile(fontData.Path) // just pass the file name
	if err != nil {
		log.Fatal("Can't read font file " + fontData.FileName)
	}
	return c.Blob(http.StatusOK, mimeType, fontBinary)

}
