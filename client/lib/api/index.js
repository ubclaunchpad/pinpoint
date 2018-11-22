
class API {
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

  async createClub({ name, desc }) {
    return this.req.post('/club/create', { name, desc });
  }

  async createPeriod({ name, start, end }) {
    return this.req.post('/club/period/create', { name, start, end });
  }
}

module.exports = {
  API,
};
