const utils = require('./utils');
const api = require('./api');

// todo: replace with real deployment
const defaultAPIPath = 'localhost';

function $main({ base = defaultAPIPath, auth } = {}) {
  return new api.API(utils.newRequester({ baseURL: base, token: auth }));
}

function addReadOnlyProperties(target, source) {
  Object.keys(source).forEach(key => Object.defineProperty(target, key, {
    value: source[key],
    configurable: false,
    writable: false,
  }));
}

addReadOnlyProperties($main, {
  // include library properties here
});

module.exports = $main;
