package fonts

import (
	"fmt"
	"net/url"
	"strings"
)

// fontFaceCSS generates the CSS for a font
func fontFaceCSS(font FontData) string {
	var css string
	for variantName, variant := range font.Variants {
		css = css + fmt.Sprintf(`@font-face{font-family: '%s';font-style: %s;font-weight: %d;%s}`, variantName, variant.CSSFontStyle, variant.CSSWeight, fontFaceSrc(variant.UniqueID, variantName, variant.Files))
	}
	return strings.Replace(css, "\n", "", -1)
}

func fontFaceSrc(uniqueID string, fontFamily string, fontFiles map[string]FontFileData) string {
	css := fmt.Sprintf("src: local('%s')", fontFamily)
	for _, fontFileData := range fontFiles {
		css = css + fmt.Sprintf(`,url(/font/%s.%s) format('%s')`,
			url.QueryEscape(uniqueID), url.QueryEscape(fontFileData.Extension), fontFileData.CSSFormat)
	}

	return css
}
