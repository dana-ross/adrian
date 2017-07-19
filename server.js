// @ts-check

const yaml = require('js-yaml')
const fs = require('fs')
const dir = require('node-dir')
const fontkit = require('fontkit')
const crypto = require('crypto')
const express = require('express')
const memoize = require('fast-memoize')
const middleware = require('./middleware')
const has = require('./has')

/**
 * Read & parse a YAML configuration file
 * @param {string} filename
 * @return {object} configuration 
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
const findFont = (fonts, id) => fonts.filter((x) => x.uniqueID === id).pop()

/**
 * 
 * @param {object} font
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

const config = readConfig('font-server.yaml')

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

function fontFaceCSS(font, protocol) {
    return `@font-face {
  font-family: '${font.familyName}';
  font-style: normal;
  font-weight: 500;
  src: local('${font.fullName}'), local('${font.familyName}-${font.subfamilyName}'), url(/${font.uniqueID}.${font.type}) format('${font.type}');
}`
}

/**
 * 
 * @param {string} variant
 * @return {number} 
 */
const fontVariantToCSSWeight = (variant) => {

    const variants = {
        'thin': 100,
        'extra light': 200,
        'light': 300,
        'normal': 400,
        'medium': 500,
        'semi bold': 600,
        'bold': 700,
        'extra bold': 800,
        'black': 900
    }

    return variants.includes(variant.toLowerCase()) ? 
            variants[variant.toLowerCase()] : variants['normal']

}

const app = express()

middleware(app, config)

/**
 * Route to serve fonts
 */
app.get('/font/:id\.(otf|ttf|woff|woff2)', (req, res) => {
    fs.createReadStream(findFont(fonts, req.params.id).filename).pipe(res)
})

/**
 * Route to serve CSS
 */
app.get('/font/:name\.css', (req, res)=> {
    res.send(fontFaceCSS(fonts[0], req.protocol))
})

app.listen(3000)
