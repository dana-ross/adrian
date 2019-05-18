package fonts

// Effects contains the CSS for supported font effects
var Effects = map[string]string{
	"anaglyph": ".font-effect-anaglyph {text-shadow: -0.06em 0 red, 0.06em 0 cyan;}",
	// brick-sign requires a texture image from Google
	// canvas-print requires a texture image from Google
	// crackle requires a texture image from Google
	// decaying requires a texture image from Google
	// destruction requires a texture image from Google
	// distressed requires a texture image from Google
	// distressed-wood requires a texture image from Google
	"emboss":         ".font-effect-emboss {text-shadow: 0px 1px 1px #fff, 0 -1px 1px #000;color: #ddd;}",
	"fire":           ".font-effect-fire {text-shadow: 0 -0.05em 0.2em #FFF, 0.01em -0.02em 0.15em #FE0, 0.01em -0.05em 0.15em #FC0, 0.02em -0.15em 0.2em #F90, 0.04em -0.20em 0.3em #F70, 0.05em -0.25em 0.4em #F70, 0.06em -0.2em 0.9em #F50, 0.1em -0.1em 1.0em #F40;color: #ffe;}",
	"fire-animation": "@-webkit-keyframes font-effect-fire-animation-keyframes {0% {text-shadow: 0 -0.05em 0.2em #FFF, 0.01em -0.02em 0.15em #FE0, 0.01em -0.05em 0.15em #FC0, 0.02em -0.15em 0.2em #F90, 0.04em -0.20em 0.3em #F70,0.05em -0.25em 0.4em #F70, 0.06em -0.2em 0.9em #F50, 0.1em -0.1em 1.0em #F40;}25% {text-shadow: 0 -0.05em 0.2em #FFF, 0 -0.05em 0.17em #FE0, 0.04em -0.12em 0.22em #FC0, 0.04em -0.13em 0.27em #F90, 0.05em -0.23em 0.33em #F70, 0.07em -0.28em 0.47em #F70, 0.1em -0.3em 0.8em #F50, 0.1em -0.3em 0.9em #F40;}50% {text-shadow: 0 -0.05em 0.2em #FFF, 0.01em -0.02em 0.15em #FE0, 0.01em -0.05em 0.15em #FC0, 0.02em -0.15em 0.2em #F90, 0.04em -0.20em 0.3em #F70,0.05em -0.25em 0.4em #F70, 0.06em -0.2em 0.9em #F50, 0.1em -0.1em 1.0em #F40;}75% {text-shadow: 0 -0.05em 0.2em #FFF, 0 -0.06em 0.18em #FE0, 0.05em -0.15em 0.23em #FC0, 0.05em -0.15em 0.3em #F90, 0.07em -0.25em 0.4em #F70, 0.09em -0.3em 0.5em #F70, 0.1em -0.3em 0.9em #F50, 0.1em -0.3em 1.0em #F40;}100% {text-shadow: 0 -0.05em 0.2em #FFF, 0.01em -0.02em 0.15em #FE0, 0.01em -0.05em 0.15em #FC0, 0.02em -0.15em 0.2em #F90, 0.04em -0.20em 0.3em #F70,0.05em -0.25em 0.4em #F70, 0.06em -0.2em 0.9em #F50, 0.1em -0.1em 1.0em #F40;}}.font-effect-fire-animation {-webkit-animation-duration:0.8s;-webkit-animation-name:font-effect-fire-animation-keyframes;-webkit-animation-iteration-count:infinite;-webkit-animation-direction:alternate;color: #ffe;}",
	// fragile requires a texture image from Google
	// grass requires a texture image from Google
	// ice requires a texture image from Google
	// mitosis requires a texture image from Google
	"neon":    ".font-effect-neon {text-shadow: 0 0 0.1em #fff, 0 0 0.2em #fff, 0 0 0.3em #fff, 0 0 0.4em #f7f,0 0 0.6em #f0f, 0 0 0.8em #f0f, 0 0 1.0em #f0f, 0 0 1.2em #f0f;color: #fff;}",
	"outline": ".font-effect-outline {text-shadow:0 1px 1px #000, 0 -1px 1px #000, 1px 0 1px #000, -1px 0 1px #000;color: #fff;}",
	// putting-green requires a texture image from Google
	// scuffed-steel requires a texture image from Google
	"shadow-multiple": ".font-effect-shadow-multiple {text-shadow: .04em .04em 0 #fff,.08em .08em 0 #aaa;-webkit-text-shadow: .04em .04em 0 #fff, .08em .08em 0 #aaa;}",
	// splintered requires a texture image from Google
	// static requires a texture image from Google
	// stonewash requires a texture image from Google
	"3d":       ".font-effect-3d {text-shadow: 0px 1px 0px #c7c8ca, 0px 2px 0px #b1b3b6, 0px 3px 0px #9d9fa2, 0px 4px 0px #8a8c8e, 0px 5px 0px #77787b, 0px 6px 0px #636466, 0px 7px 0px #4d4d4f, 0px 8px 7px #001135;color: #fff;}",
	"3d-float": ".font-effect-3d-float {text-shadow: 0 0.032em 0 #b0b0b0, 0px 0.15em 0.11em rgba(0,0,0,0.15), 0px 0.25em 0.021em rgba(0,0,0,0.1), 0px 0.32em 0.32em rgba(0,0,0,0.1);color: #fff;}",
	// vintage requires a texture image from Google
	// wallpaper requires a texture image from Google
}
