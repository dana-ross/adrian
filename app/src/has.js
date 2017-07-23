// @ts-check

/**
 * 
 * @param {object} obj 
 * @param {string} key 
 * @return {boolean}
 */
module.exports = (obj, key) => obj && hasOwnProperty.call(obj, key)
