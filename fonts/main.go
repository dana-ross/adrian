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

	adrianConfig "Adrian2.0/config"
	"github.com/ConradIrwin/font/sfnt"
)

// FontData describes a font file and the various metadata associated with it.
type FontData struct {
	Path      string
	Name      string
	Family    string
	SubFamily string
	Type      string
	CSSFormat string
	CSSWeight int
	CSS       string
	FileName  string
	UniqueID  string
	Metadata  map[sfnt.NameID]string
}

var fonts []FontData

// LoadFont loads a font into memory
func LoadFont(filePath string, config adrianConfig.Config) FontData {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	font, parseError := sfnt.Parse(file)
	if parseError != nil {
		log.Fatal(parseError)
	}

	fontData := FontData{
		Path:     filePath,
		FileName: path.Base(filePath),
		Type:     strings.ToLower(regexp.MustCompile("^\\.").ReplaceAllLiteralString(path.Ext(filePath), "")),
		Metadata: make(map[sfnt.NameID]string),
	}

	nameTable, err := font.NameTable()
	if err != nil {
		log.Fatal(err)
	}

	for _, nameEntry := range nameTable.List() {
		fontData.Metadata[nameEntry.NameID] = nameEntry.String()
	}
	fontData.Name = fontData.Metadata[sfnt.NameFull]
	if fontData.Name == "" {
		log.Fatalf("Font has no name! Using file name instead.")
	}

	fontData.Family = fontData.Metadata[sfnt.NamePreferredFamily]
	if fontData.Family == "" {
		if v, ok := fontData.Metadata[sfnt.NameFontFamily]; ok {
			fontData.Family = v
		} else {
			log.Fatalf("Font %v has no font family!", fontData.Name)
		}
	}

	fontData.SubFamily = fontData.Metadata[sfnt.NameFontSubfamily]

	fontData.UniqueID = calcUniqueID(fontData, config)

	fontData.CSSWeight = guessFontCSSWeight(fontData)
	fontData.CSSFormat = fontCSSFormat(fontData)
	fontData.CSS = fontFaceCSS(fontData)

	log.Printf("Loaded font: %s", fontData.Name)
	fonts = append(fonts, fontData)
	return fontData
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
func fontCSSFormat(fontData FontData) string {
	switch fontData.Type {
	case "ttf":
		return "truetype"
	default:
		return fontData.Type
	}
}
