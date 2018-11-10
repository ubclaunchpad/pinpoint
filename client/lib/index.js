import utils from './utils';
import api from './api';

// todo: replace with real deployment
const defaultAPIPath = 'http://localhost:8081/';

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
  API: api.API,
});

export default $main;
