// @ts-check

const express = require('express')
const has = require('./has')
const url = require('url')

/**
 * @param {object} app
 * @param {object} config
 */
module.exports = (app, config) => {

    /**
     * Middleware to set X-Powered-By header
     */
    app.use((req, res, next) => {
        res.removeHeader("X-Powered-By")
        res.setHeader('X-Powered-By', 'Adrian 1.0')
        next()
    })

    /**
     * Middleware to set CORS headers
     */
    app.use((req, res, next) => {
        let refererDomain
        const domains = (has(config, 'global') && has(config.global, 'domains')) ? config.global.domains : []
        const debug = (has(config, 'global') && has(config.global, 'debug')) ? config.global.debug : false

        if (debug) {
            next()
        }
        else if(req.get('referer') && (refererDomain = url.parse(req.get('referer'))) && domains.includes(refererDomain.hostname)) {
            res.setHeader('Access-Control-Allow-Origin', refererDomain.protocol + '//' + refererDomain.host)
            res.setHeader('Access-Control-Allow-Methods', 'GET')
            next()
        } else {
            res.sendStatus(403)
        }
    })

    /**
     * Middleware to set caching headers
     */
    app.use((req, res, next) => {
        // 2628000 seconds = 30 days
        if (res.statusCode === 200) {
            res.setHeader('Cache-Control', 'max-age=2628000, public')
            res.setHeader('Vary', 'Origin')
        }
        next()
    })

}
