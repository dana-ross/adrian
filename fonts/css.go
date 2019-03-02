package fonts

import "fmt"

// FontFactCSS generates the CSS for a font
func FontFaceCSS(font FontData, protocol string) string {

	return fmt.Sprintf(`@font-face {
  font-family: '%s';
  font-style: normal;
  font-weight: %d;
  src: local('%s'), url(${font.uniqueID}.${font.type}) format('${font.type}');
}`, font.Name, font.CSSWeight, font.Name)
}
