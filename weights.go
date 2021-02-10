package main

import (
	"regexp"
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

func guessFontCSSWeight(fontVariant FontVariant) int {
	variantName := strings.ToLower(fontVariant.Name)
	var weightName string

	if match, _ := regexp.MatchString("black$", variantName); match {
		weightName = "black"
	} else if match, _ := regexp.MatchString("italic$", variantName); match {
		weightName = "normal"
	} else {
		weightName = variantName
	}

	if _, ok := fontWeights[weightName]; ok {
		return fontWeights[weightName]
	}

	return fontWeights["normal"]

}
