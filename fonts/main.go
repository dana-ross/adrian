package fonts

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/ConradIrwin/font/sfnt"
	adrianConfig "github.com/daveross/adrian/config"
)

// FontFileData describes a font (format) belonging to a font family
type FontFileData struct {
	Path      string
	FileName  string
	Extension string
	CSSFormat string
}

// FontData describes a font file and the various metadata associated with it.
type FontData struct {
	Name      string
	Family    string
	SubFamily string
	CSSWeight int
	CSS       string
	UniqueID  string
	Metadata  map[sfnt.NameID]string
	Files     map[string]FontFileData
}

var supportedFormats = map[string]bool{
	"ttf":   true,
	"otf":   true,
	"woff":  true,
	"woff2": true,
}

var fonts = make(map[string]FontData)

// LoadFont loads a font into memory
func LoadFont(filePath string, config adrianConfig.Config) {

	fontFormat := GetCanonicalExtension(filePath)
	if _, ok := supportedFormats[fontFormat]; !ok {
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	font, parseError := sfnt.Parse(file)
	if parseError != nil {
		log.Fatal(parseError)
	}

	nameTable, err := font.NameTable()
	if err != nil {
		log.Fatal(err)
	}

	metadata := make(map[sfnt.NameID]string)
	for _, nameEntry := range nameTable.List() {
		metadata[nameEntry.NameID] = nameEntry.String()
	}

	fontName := metadata[sfnt.NameFull]
	if fontName == "" {
		log.Fatalf("Font has no name! Using file name instead.")
	}

	fontFamily := metadata[sfnt.NamePreferredFamily]
	if fontFamily == "" {
		if v, ok := metadata[sfnt.NameFontFamily]; ok {
			fontFamily = v
		} else {
			log.Fatalf("Font %v has no font family!", fontName)
		}
	}

	fontData, ok := fonts[fontFamily]
	if !ok {
		fontData = FontData{
			Name:     fontName,
			Metadata: metadata,
			Family:   fontFamily,
			Files:    make(map[string]FontFileData),
		}
	}

	fontData.SubFamily = fontData.Metadata[sfnt.NameFontSubfamily]

	fontData.UniqueID = calcUniqueID(fontData, config)

	fontData.CSSWeight = guessFontCSSWeight(fontData)

	fontFileData := FontFileData{
		CSSFormat: fontCSSFormat(fontFormat),
		Extension: fontFormat,
	}

	fontData.Files[fontFormat] = fontFileData
	fontData.CSS = fontFaceCSS(fontData)

	log.Printf("Loaded font: %s (%s)", fontData.Name, fontFormat)
	fonts[fontFamily] = fontData
}

// GetFont returns the data for a single font by its name
func GetFont(name string) (fontData FontData, err error) {
	name = strings.ToLower(name)
	for _, fontData := range fonts {
		if strings.ToLower(fontData.Name) == name {
			return fontData, nil
		}
	}

	return FontData{}, errors.New("Font not found")
}

// GetFontByUniqueID returns the data for a single font by its unique ID
func GetFontByUniqueID(uniqueID string) (fontData FontData, err error) {
	for _, fontData := range fonts {
		if fontData.UniqueID == uniqueID {
			return fontData, nil
		}
	}

	return FontData{}, errors.New("Font not found")
}

// calcUniqueID generates a unique ID for a font, optionally obfuscating it
func calcUniqueID(fontData FontData, config adrianConfig.Config) string {
	if config.Global.ObfuscateFilenames {
		hash := sha256.New()
		hash.Write([]byte(fontData.Family + fontData.SubFamily))
		return hex.EncodeToString(hash.Sum(nil))
	}

	return fontData.Family
}

// fontCSSFormat determines the appropriate CSS font format given a FontData struct with Type set
func fontCSSFormat(fileType string) string {
	switch fileType {
	case "ttf":
		return "truetype"
	default:
		return fileType
	}
}

// GetCanonicalExtension is
func GetCanonicalExtension(filePath string) string {
	return strings.ToLower(regexp.MustCompile("^\\.").ReplaceAllLiteralString(path.Ext(filePath), ""))
}
