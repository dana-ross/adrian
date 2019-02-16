// @ts-check

const { Console } = require('console')
const logger = new Console(process.stdout, process.stderr)
require('console-stamp')(logger)
module.exports = logger
