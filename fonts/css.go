package fonts

import "fmt"

// fontFaceCSS generates the CSS for a font
func fontFaceCSS(font FontData) string {
	var css string

	css = css + fmt.Sprintf(`@font-face {
    font-family: '%s';
    font-style: normal;
    font-weight: %d;
    %s
}`, font.Name, font.CSSWeight, fontFaceSrc(font))

	return css
}

func fontFaceSrc(font FontData) string {
	css := fmt.Sprintf("src: local('%s')", font.Family)
	for _, fontFileData := range font.Files {
		css = css + fmt.Sprintf(`, url(/font/%s.%s) format('%s')`,
			font.UniqueID, fontFileData.Extension, fontFileData.CSSFormat)
	}
	return css
}
