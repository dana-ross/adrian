package fonts

import (
	"fmt"
	"log"
	"os"

	"github.com/ConradIrwin/font/sfnt"
)

var fonts []*sfnt.Font

// LoadFont loads a font into memory
func LoadFont(path string) *sfnt.Font {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	font, parseError := sfnt.Parse(file)
	if parseError != nil {
		log.Fatal(parseError)
	}

	fonts = append(fonts, font)
	return font
}

func ListFonts() {
	fmt.Print(fonts)
}
