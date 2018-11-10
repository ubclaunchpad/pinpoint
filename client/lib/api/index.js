
export default class API {
  constructor(req = require('axios')) {
    this.req = req;
  }

  async getStatus() {
    return this.req.get('/status');
  }

  async createAccount({ email, name, password }) {
    return this.req.post('/user/create', { email, name, password });
  }

  async login({ email, password }) {
    return this.req.post('/user/login', { email, password });
  }
}
