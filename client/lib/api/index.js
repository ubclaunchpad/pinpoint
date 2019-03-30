
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

  async verify({ hash }) {
    try {
      await this.req.post('/user/verify', { hash });
      return true;
    } catch (error) {
      switch (error.response.status) {
        default: throw new Error(`error ${error.response.status}`);
      }
    }
  }

  async createClub({ clubID, description }) {
    try {
      const response = await this.req.post('/club/create', { clubID, description });
      return response.data.clubID;
    } catch (error) {
      switch (error.response.status) {
        default: throw new Error(`error ${error.response.status}`);
      }
    }
  }

  async createPeriod({ period }) {
    try {
      const response = await this.req.post('/club/period/create', { period });
      return response.data.period;
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
