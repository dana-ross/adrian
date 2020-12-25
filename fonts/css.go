package fonts

import (
	"fmt"
	"net/url"
)

// FontFaceCSS generates the CSS for a font
func FontFaceCSS(font FontData, fontWeights []int, display string) string {
	var css = ""
	for variantName, variant := range font.Variants {
		if len(fontWeights) > 0 {
			for _, requestedWeight := range fontWeights {
				if requestedWeight == variant.CSSWeight {
					css = css + fmt.Sprintf(`@font-face{font-family: '%s';font-style: %s;font-weight: %d;%s%s}`, variantName, variant.CSSFontStyle, variant.CSSWeight, fontDisplayCSS(display), fontFaceSrc(variant.UniqueID, variantName, variant.Files))
				}
			}
		} else {
			css = css + fmt.Sprintf(`@font-face{font-family: '%s';font-style: %s;font-weight: %d;%s%s}`, variantName, variant.CSSFontStyle, variant.CSSWeight, fontDisplayCSS(display), fontFaceSrc(variant.UniqueID, variantName, variant.Files))
		}
	}
	return css
}

func fontFaceSrc(uniqueID string, fontFamily string, fontFiles map[string]FontFileData) string {
	css := fmt.Sprintf("src: ")
	for fontFileIndex, fontFileData := range fontFiles {
		if(fontFileIndex > 0) {
			css = css + ', '
		}
		css = css + fmt.Sprintf(`url('/font/%s.%s') format('%s')`,
			url.QueryEscape(uniqueID), url.QueryEscape(fontFileData.Extension), fontFileData.CSSFormat)
	}

	return css
}

func fontDisplayCSS(display string) string {
	switch display {
	case "auto", "block", "swap", "fallback", "optional":
		return fmt.Sprintf("font-display: %s;", display)
	default:
		return ""
	}
}
