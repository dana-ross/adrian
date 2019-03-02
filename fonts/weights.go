package fonts

import (
	"strings"
)

var fontWeights = map[string]int{
	"thin":       100,
	"extralight": 200,
	"ultralight": 200,
	"light":      300,
	"book":       400,
	"normal":     400,
	"regular":    400,
	"roman":      400,
	"medium":     500,
	"semibold":   600,
	"demibold":   600,
	"bold":       700,
	"extrabold":  800,
	"ultrabold":  900,
	"black":      900,
	"heavy":      900,
}

func guessFontCSSWeight(FontData FontData) int {

	fontVariant := strings.ToLower(FontData.SubFamily)

	if _, ok := fontWeights[fontVariant]; ok {
		return fontWeights[fontVariant]
	}

	if _, ok := fontWeights[strings.Replace(fontVariant, " italic", "", -1)]; ok {
		return fontWeights[strings.Replace(fontVariant, " italic", "", -1)]
	}

	return fontWeights["normal"]

}
