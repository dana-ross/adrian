package fonts

import "fmt"

// FontFaceCSS generates the CSS for a font
func FontFaceCSS(font FontData) string {

	return fmt.Sprintf(`@font-face {
  font-family: '%s';
  font-style: normal;
  font-weight: %d;
  src: local('%s'), url(/font/%s.%s) format('%s');
}`, font.Name, font.CSSWeight, font.Name, font.UniqueID, font.Type, font.CSSFormat)
}
