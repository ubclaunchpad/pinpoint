
class API {
  constructor(req = require('axios')) {
    this.req = req;
  }

  status() {
    return this.req.get('/status');
  }
}

module.exports = {
  API,
};
