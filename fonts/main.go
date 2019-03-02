package fonts

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/ConradIrwin/font/sfnt"
)

// FontData describes a font file and the various metadata associated with it.
type FontData struct {
	Path      string
	Name      string
	Family    string
	SubFamily string
	CSSWeight int
	FileName  string
	Metadata  map[sfnt.NameID]string
	// Data     []byte
}

var fonts []FontData

// LoadFont loads a font into memory
func LoadFont(path string) FontData {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	font, parseError := sfnt.Parse(file)
	if parseError != nil {
		log.Fatal(parseError)
	}

	fontData := FontData{
		Path:     "",
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

	fontData.CSSWeight = guessFontCSSWeight(fontData)

	fonts = append(fonts, fontData)
	return fontData
}

// GetFont returns the data for a single font
func GetFont(name string) (fontData FontData, err error) {
	name = strings.ToLower(name)
	for _, fontData := range fonts {
		if strings.ToLower(fontData.Name) == name {
			return fontData, nil
		}
	}

	return FontData{}, errors.New("Font not found")
}
