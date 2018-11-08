
class API {
  constructor(req = require('axios')) {
    this.req = req;
  }

  getStatus() {
    return this.req.get('/status');
  }
}

module.exports = {
  API,
};
