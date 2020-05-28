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
	"strconv"
	"strings"

	adrianConfig "github.com/daveross/adrian/config"
	adrianFonts "github.com/daveross/adrian/fonts"

	"github.com/labstack/echo"
)

func main() {

	versionParam := flag.Bool("version", false, "display the version number and exit")
	configParam := flag.String("config", "", "specify a config file")
	flag.Parse()

	// Handle the --version parameter
	if *versionParam {
		fmt.Printf("%s\n", "2.2.1")
		os.Exit(0)
	}

	log.Println("Starting Adrian 2.1.1")
	log.Println("Loading adrian.yaml")
	var config adrianConfig.Config
	if *configParam != "" {
		config = adrianConfig.LoadConfig(*configParam)
	} else {
		config = adrianConfig.LoadConfig("./adrian.yaml")
	}
	log.Println("Initializing web server")
	e := Instantiate(config)
	log.Println("Loading fonts and starting watchers")
	for _, folder := range config.Global.Directories {
		adrianFonts.FindFonts(folder, config)
		adrianFonts.InstantiateWatcher(folder, config)
	}

	accessLog := openAccessLog(config.Global.Logs.Access)
	errorLog := openErrorLog(config.Global.Logs.Error)

	log.Println("Defining paths")

	registerCSSPath(e, accessLog, errorLog)
	registerFontPath(e, accessLog, errorLog)

	log.Printf("Listening on port %d", config.Global.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Global.Port)))
}

func registerCSSPath(e *echo.Echo, accessLog *os.File, errorLog *os.File) {
	e.GET("/css/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, "text/css")
		fontRequests := strings.Split(c.QueryParam("family"), "|")
		display := c.QueryParam("display")
		var fontsCSS string
		for _, fontRequest := range fontRequests {
			fontRequestData := strings.SplitN(fontRequest, ":", 2)
			fontFamilyName := fontRequestData[0]
			var fontWeights []int
			if len(fontRequestData) > 1 {
				for _, weight := range strings.Split(fontRequestData[1], ",") {
					numericWeight, err := strconv.Atoi(weight)
					if err == nil {
						fontWeights = append(fontWeights, numericWeight)
					}
				}
				fontWeights = uniqueInts(fontWeights)
			}
			fontData, err := adrianFonts.GetFont(fontFamilyName)
			if err != nil {
				writeAccessLog(accessLog, errorLog, c, 404, 0)
				return return404(c)
			}
			fontsCSS = fontsCSS + adrianFonts.FontFaceCSS(fontData, fontWeights, display)
		}
		writeToCache(c, fontsCSS)
		writeAccessLog(accessLog, errorLog, c, 200, len(fontsCSS))
		return c.String(http.StatusOK, fontsCSS)
	})

	return
}

func registerFontPath(e *echo.Echo, accessLog *os.File, errorLog *os.File) {
	e.GET("/font/:filename/", func(c echo.Context) error {
		filename, error := url.QueryUnescape(c.Param("filename"))
		if error != nil {
			writeAccessLog(accessLog, errorLog, c, 404, 0)
			return return404(c)
		}

		switch filepath.Ext(filename) {
		case ".ttf":
			return outputFont(c, "font/truetype", accessLog, errorLog)
		case ".woff":
			return outputFont(c, "font/woff", accessLog, errorLog)
		case ".woff2":
			return outputFont(c, "font/woff2", accessLog, errorLog)
		case ".otf":
			return outputFont(c, "font/opentype", accessLog, errorLog)
		}

		writeAccessLog(accessLog, errorLog, c, 404, 0)
		return return404(c)
	})
	return
}
// Basename gets the base filename (minus the last extension)
func basename(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n >= 0 {
		return s[:n]
	}
	return s
}

func outputFont(c echo.Context, mimeType string, accessLog *os.File, errorLog *os.File) error {
	filename, error := url.QueryUnescape(c.Param("filename"))
	if error != nil {
		writeAccessLog(accessLog, errorLog, c, 404, 0)
		return return404(c)
	}
	fontVariant, err := adrianFonts.GetFontVariantByUniqueID(basename(filename))
	if err != nil {
		writeAccessLog(accessLog, errorLog, c, 404, 0)
		return return404(c)
	}

	fontFileData, ok := fontVariant.Files[adrianFonts.GetCanonicalExtension(filename)]
	if !ok {
		log.Fatal("Invalid font format" + adrianFonts.GetCanonicalExtension(filename))
	}

	for i := range c.Request().Header["If-None-Match"] {
		individualHashes := strings.Split(c.Request().Header["If-None-Match"][i], (", "))
		for j := range individualHashes {
			if individualHashes[j] == fontFileData.MD5 {
				status := make(map[string]string)
				status["message"] = "Not Modified"
				
				writeAccessLog(accessLog, errorLog, c, 304, 0)
				return c.JSON(http.StatusNotModified, status)	
			}
		}
	}

	fontBinary, err := ioutil.ReadFile(fontFileData.Path) // just pass the file name
	if err != nil {
		log.Fatal("Can't read font file " + fontFileData.FileName)
	}

	c.Response().Header().Set("Content-Transfer-Encoding", "binary")
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	c.Response().Header().Set("ETag", fontFileData.MD5)

	_, err = accessLog.WriteString(formatAccessLogMessage(c, 200, len(fontBinary)) + "\n")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error writing to access log: %s", err))
	}
	
	return c.Blob(http.StatusOK, mimeType, fontBinary)

}

// uniqueInts returns a unique subset of the int slice provided.
func uniqueInts(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
