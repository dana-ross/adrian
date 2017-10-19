// @ts-check

const yaml = require('js-yaml')
const fs = require('fs')
const dir = require('node-dir')
const fontkit = require('fontkit')
const crypto = require('crypto')
const express = require('express')
const middleware = require('./middleware')
const apicache = require('apicache')
const has = require('./has')
const fontWeights = {
	"thin": 100,
	"extralight": 200,
	"ultralight": 200,
	"light": 300,
	"book": 400,
	"normal": 400,
	"regular": 400,
	"roman": 400,
	"medium": 500,
	"semibold": 600,
	"demibold": 600,
	"bold": 700,
	"extrabold": 800,
	"ultrabold": 800,
	"black": 900,
	"heavy": 900
}

/**
 * Read & parse a YAML configuration file
 * @param {string} filename
 * @return {Object} configuration 
 */
const readConfig = (filename) => yaml.safeLoad(fs.readFileSync(filename, 'utf8'))


/**
 * 
 * @param {array[string]} directories 
 */
function findFonts(directories) {
    return directories.map((directory) => {
        return dir.files(directory, { sync: true }).filter(
            (filename) => filename.match(/\/([^.])[^\/]*\.(otf|ttf|woff|woff2)$/i)
        )
    }).reduce((a,b) => a.concat(b))
}

/**
 * 
 * @param {array[object]} fonts 
 * @param {string} id
 * @return object
 */
const findFontByID = (fonts, id) => fonts.filter((x) => x.uniqueID === id).pop()
const findFontByName = (fonts, name) => fonts.filter((x) => x.fullName === name).pop()
const findFontsByFamilyName = (fonts, name) => fonts.filter((x) => x.fullName.toLowerCase().startsWith(name.toLowerCase()))

/**
 * 
 * @param {Object} font
 * @return string 
 */
function fontType(font) {
    switch(font.constructor.name) {
        case 'TTFFont': return 'ttf'
        case 'WOFF2Font': return 'woff2'
        case 'WOFFFont': return 'woff'
        default: return ''
    }
}

/**
 * Determine or guess a font's CSS weight
 * @param {Object} font 
 */
function guessFontCSSWeight(fontWeights, font) {
    
    let fontVariant = font.subfamilyName.toLowerCase() 
    if('regular' !== fontVariant) {
        if(fontWeights.hasOwnProperty(fontVariant)) {
            return fontWeights[fontVariant]
        }
    }

    for(let fontWeightIndex in fontWeights) {
        if(font.fullName.toLowerCase().replace(/( italic)$/, '').endsWith(' ' + fontWeightIndex)) {
            return fontWeights[fontWeightIndex]
        }
    }

    return fontWeights.regular

}

const config = readConfig('adrian.yaml')

const fontDirectories = (has(config, 'global') && has(config.global, 'directories')) ? config.global.directories : []
const fonts = findFonts(fontDirectories).map((filename) => {
    const font = fontkit.openSync(filename)

    return {
        filename: filename,
        type: fontType(font),
        fullName: font.fullName,
        familyName: font.familyName,
        subfamilyName: font.subfamilyName,
        copyright: font.copyright,
        uniqueID: crypto.createHash('sha256').update(font.familyName + ' ' + font.subfamilyName).digest('hex')
    }
})

const fontFaceCSS = (font, protocol) => {
    const fontWeight = guessFontCSSWeight(fontWeights, font)
    return `@font-face {
  font-family: '${font.familyName}';
  font-style: normal;
  font-weight: ${fontWeight};
  src: local('${font.fullName}'), url(${font.uniqueID}.${font.type}) format('${font.type}');
}`.split("\n").map((x) => x.replace(/^\s+/, '')).join(' ')
}

const app = express()

middleware(app, config)

/**
 * Route to serve fonts
 */
app.get('/font/:id\.(otf|ttf|woff|woff2)', (req, res) => {
    fs.createReadStream(findFontByID(fonts, req.params.id).filename).pipe(res)
})

/**
 * Route to serve CSS
 */
app.get('/font/:name\.css', apicache.middleware('5 minutes'), (req, res)=> {
    if(findFontByName(fonts, req.params.name)) {
        res.send(fontFaceCSS(findFontByName(fonts, req.params.name), req.protocol))
    }
    else {
        res.sendStatus(404)
    }
})

app.get('/font/family/:name.css', apicache.middleware('5 minutes'), (req, res) => {
    const familyMembers = findFontsByFamilyName(fonts, req.params.name)
    if(familyMembers.length) {
        familyMembers.forEach((font) => res.write(fontFaceCSS(font)) + "\n")
        res.end()
    }
    else {
        res.sendStatus(404)
    }
})

app.listen(3000)
