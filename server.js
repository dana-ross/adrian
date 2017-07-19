// @ts-check

const yaml = require('js-yaml')
const fs = require('fs')
const dir = require('node-dir')
const fontkit = require('fontkit')
const crypto = require('crypto')
const express = require('express')
const url = require('url')
const memoize = require('fast-memoize')

/**
 * Read & parse a YAML configuration file
 * @param {string} filename
 * @return {object} configuration 
 */
const readConfig = (filename) => yaml.safeLoad(fs.readFileSync(filename, 'utf8'))

/**
 * 
 * @param {object} obj 
 * @param {string} key 
 */
const has = (obj, key) => obj && hasOwnProperty.call(obj, key)

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

function fontCSS(font) {
    return `@font-face { font-family: '${font.familyName}'; font-style: normal; font-weight: 400; src: local('Roboto'), local('Roboto-Regular'), url(https://fonts.gstatic.com/s/roboto/v16/ek4gzZ-GeXAPcSbHtCeQI_esZW2xOQ-xsNqO47m55DA.woff2) format('woff2'); unicode-range: U+0460-052F, U+20B4, U+2DE0-2DFF, U+A640-A69F;`
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

const fontWeight = () => {

}

const app = express()

/**
 * Middleware to set CORS headers
 */
app.use((req, res, next) => {
    let refererDomain
    const domains = (has(config, 'global') && has(config.global, 'domains')) ? config.global.domains : []

    if(req.get('referer') && (refererDomain = url.parse(req.get('referer'))) && domains.includes(refererDomain.hostname)) {
        res.setHeader('Access-Control-Allow-Origin',  refererDomain.protocol + '//' + refererDomain.host)
        res.setHeader('Access-Control-Allow-Methods', 'GET')
        next()
    }
    else {
        res.sendStatus(403)
    }
})

/**
 * Middleware to set caching headers
 */
app.use((req, res, next) =>{
    // 2628000 seconds = 30 days
    if(res.statusCode === 200) {
        res.setHeader('Cache-Control', 'max-age=2628000, public')
        res.setHeader('Vary', 'Origin')
    }
    next()
})

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
