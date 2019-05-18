package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	adrianConfig "github.com/daveross/adrian/config"
	adrianFonts "github.com/daveross/adrian/fonts"
	adrianServer "github.com/daveross/adrian/server"

	"github.com/labstack/echo"
)

func main() {

	versionParam := flag.Bool("version", false, "display the version number and exit")
	flag.Parse()

	// Handle the --version parameter
	if *versionParam {
		fmt.Printf("%s\n", "2.0.0")
		os.Exit(0)
	}

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
		display := c.QueryParam("display")
		effects := strings.Split(c.QueryParam("effect"), "|")
		var fontsCSS string
		for _, fontFilename := range fontFilenames {
			fontData, err := adrianFonts.GetFont(fontFilename)
			if err != nil {
				return adrianServer.Return404(c)
			}
			fontsCSS = fontsCSS + adrianFonts.FontCSS(fontData, display)
		}

		for _, effect := range effects {
			if adrianFonts.Effects[effect] != "" {
				fontsCSS = fontsCSS + adrianFonts.Effects[effect]
			}
		}

		return c.String(http.StatusOK, fontsCSS)
	})

	e.GET("/font/:filename/", func(c echo.Context) error {
		filename, error := url.QueryUnescape(c.Param("filename"))
		if error != nil {
			return adrianServer.Return404(c)
		}

		switch filepath.Ext(filename) {
		case ".ttf":
			return outputFont(c, "font/truetype")
		case ".woff":
			return outputFont(c, "font/woff")
		case ".woff2":
			return outputFont(c, "font/woff2")
		case ".otf":
			return outputFont(c, "font/opentype")
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
	filename, error := url.QueryUnescape(c.Param("filename"))
	if error != nil {
		return adrianServer.Return404(c)
	}
	fontVariant, err := adrianFonts.GetFontVariantByUniqueID(basename(filename))
	if err != nil {
		return adrianServer.Return404(c)
	}

	fontFileData, ok := fontVariant.Files[adrianFonts.GetCanonicalExtension(filename)]
	if !ok {
		log.Fatal("Invalid font format" + adrianFonts.GetCanonicalExtension(filename))
	}

	fontBinary, err := ioutil.ReadFile(fontFileData.Path) // just pass the file name
	if err != nil {
		log.Fatal("Can't read font file " + fontFileData.FileName)
	}

	c.Response().Header().Set("Content-Transfer-Encoding", "binary")
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	return c.Blob(http.StatusOK, mimeType, fontBinary)

}
