package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ConradIrwin/font/sfnt"
)

// FontFileData describes a font (format) belonging to a font family
type FontFileData struct {
	Name      string
	Path      string
	FileName  string
	Extension string
	CSSFormat string
	MD5       string
}

// FontVariant describes a variant (ex: "Black", "Italic") of a font
type FontVariant struct {
	Name         string
	SubFamily    string
	UniqueID     string
	CSSFontStyle string
	CSSWeight    int
	Files        map[string]FontFileData
}

// FontData describes a font file and the various metadata associated with it.
type FontData struct {
	Family   string
	Metadata map[sfnt.NameID]string
	Variants map[string]FontVariant
}

// supportedFormats contains the font formats Adrian currently supports
var supportedFormats = map[string]bool{
	"ttf":   true,
	"otf":   true,
	"woff":  true,
	"woff2": true,
}

var fonts = make(map[string]FontData)
var uniqueIDXref = make(map[string]*FontVariant)

// LoadFont loads a font into memory
func LoadFont(filePath string, config Config) {

	fontFormat := GetCanonicalExtension(filePath)
	if _, ok := supportedFormats[fontFormat]; !ok {
		return
	}

	file, err := os.Open(filepath.Clean(filePath))
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
			Metadata: metadata,
			Family:   fontFamily,
			Variants: make(map[string]FontVariant),
		}
	}
	
	fontFileData := FontFileData{
		Name:      fontName,
		CSSFormat: fontCSSFormat(fontFormat),
		Extension: fontFormat,
		Path:      filePath,
		MD5:       calcFileMD5(filePath),
	}

	fontVariant, ok := fontData.Variants[fontName]
	if !ok {
		fontVariant = FontVariant{
			Name:      fontName,
			SubFamily: fontData.Metadata[sfnt.NameFontSubfamily],
			Files:     make(map[string]FontFileData),
		}
		fontVariant.CSSWeight = guessFontCSSWeight(fontVariant)
		fontVariant.CSSFontStyle = guessFontCSSStyle(fontVariant)

		fontVariant.UniqueID = calcUniqueID(fontVariant, config)
	}

	fontVariant.Files[fontFormat] = fontFileData
	fontData.Variants[fontName] = fontVariant
	log.Printf("Loaded font: %s (%s)", fontName, fontFormat)
	fonts[fontFamily] = fontData
	uniqueIDXref[fontVariant.UniqueID] = &fontVariant
}

// GetFont returns the data for a single font by its name
func GetFont(name string) (fontData FontData, err error) {
	name = strings.ToLower(name)
	for _, fontData := range fonts {
		if strings.ToLower(fontData.Family) == name {
			return fontData, nil
		}
	}

	return FontData{}, errors.New("Font not found")
}

// GetFontVariantByUniqueID returns the data for a single font by its unique ID
func GetFontVariantByUniqueID(uniqueID string) (fontVariant *FontVariant, err error) {
	fontVariant, ok := uniqueIDXref[uniqueID]
	if ok {
		return fontVariant, nil
	}

	return &FontVariant{}, errors.New("Font not found")
}

// calcUniqueID generates a unique ID for a font, optionally obfuscating it
func calcUniqueID(fontVariant FontVariant, config Config) string {
	if config.Global.ObfuscateFilenames {
		hash := sha256.New()
		hash.Write([]byte(fontVariant.Name))
		return hex.EncodeToString(hash.Sum(nil))
	}

	return fontVariant.Name
}

// calcFileMD5 calculates the MD5 hash of a file
func calcFileMD5(filePath string) string {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hash.Sum(nil))
}

// fontCSSFormat determines the appropriate CSS font format given a FontData struct with Type set
func fontCSSFormat(fileType string) string {
	switch fileType {
	case "ttf":
		return "truetype"
	case "otf":
		return "opentype"
	default:
		return fileType
	}
}

// GetCanonicalExtension is
func GetCanonicalExtension(filePath string) string {
	return strings.ToLower(regexp.MustCompile("^\\.").ReplaceAllLiteralString(path.Ext(filePath), ""))
}

// guessFontCSSStyle is
func guessFontCSSStyle(fontVariant FontVariant) string {
	variantName := strings.ToLower(fontVariant.Name)
	if match, _ := regexp.MatchString("italic$", variantName); match {
		return "italic"
	}

	return "normal"

}
