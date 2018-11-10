import axios from 'axios';

function newRequester({ baseURL = 'localhost', token = '' }) {
  return axios.create({
    baseURL,
    timeout: 1000,
    headers: {
      authentication: token ? `bearer ${token}` : '',
    },
  });
}

module.exports = {
  newRequester,
};
