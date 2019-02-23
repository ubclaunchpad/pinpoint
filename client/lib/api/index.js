
class API {
  constructor(req = require('axios')) {
    this.req = req;
  }

  async getStatus() {
    try {
      const response = await this.req.get('/status');
      return response.data.resp;
    } catch (error) {
      switch (error.response.status) {
        default: throw new Error(`error ${error.response.status}`);
      }
    }
  }

  async createAccount({ email, name, password }) {
    try {
      const response = await this.req.post('/user/create', { email, name, password });
      return response.data.email;
    } catch (error) {
      switch (error.response.status) {
        case 400: throw new Error(`${error.response.data.message}`);
        default: throw new Error(`error ${error.response.status}`);
      }
    }
  }

  async login({ email, password }) {
    try {
      const response = await this.req.post('/user/login', { email, password });
      return response.data.token;
    } catch (error) {
      switch (error.response.status) {
        case 401: throw new Error('Incorrect Credentials');
        default: throw new Error(`error ${error.response.status}`);
      }
    }
  }

  async createClub({ name, desc }) {
    try {
      const response = await this.req.post('/club/create', { name, desc });
      return response.data.ClubID;
    } catch (error) {
      switch (error.response.status) {
        default: throw new Error(`error ${error.response.status}`);
      }
    }
  }

  async createPeriod({ name, start, end }) {
    try {
      const response = await this.req.post('/club/period/create', { name, start, end });
      return response.data.PeriodID;
    } catch (error) {
      switch (error.response.status) {
        default: throw new Error(`error ${error.response.status}`);
      }
    }
  }
}

module.exports = {
  API,
};
