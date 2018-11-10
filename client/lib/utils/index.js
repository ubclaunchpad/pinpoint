import axios from 'axios';

export default function newRequester({ baseURL = 'localhost', token = '' }) {
  return axios.create({
    baseURL,
    timeout: 1000,
    headers: {
      authentication: token ? `bearer ${token}` : '',
    },
  });
}
