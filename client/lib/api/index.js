
class API {
  constructor(req = require('axios')) {
    this.req = req;
  }

  async getStatus() {
  try{
    let response = await this.req.get('/status');
    return response.data.resp;
  }
  catch (error) {
    throw "error " + error.response.status;
  }
  }

  async createAccount({ email, name, password }) {
    try{
      let response = await this.req.post('/user/create', { email, name, password });  
      return response.data.email;
    } 
    catch (error) {
      throw "error " + error.response.status;
    }
  }

  async login({ email, password }) {
    try{
      let response = await this.req.post('/user/login', { email, password })
      return response.data.token;
    }
    catch (error) {
      throw "error " + error.response.status;
    }
  }

  async createClub({ name, desc }) {
    try{
      let response = await this.req.post('/club/create', {  name, desc });
      return response.data.ClubID;
    } 
    catch (error) {
      throw "error " + error.response.status;
    }
  }

  async createPeriod({ name, start, end }) {
    try{
      let response = await this.req.post('/club/period/create', { name, start,end });
      return response.data.PeriodID;
    } 
    catch (error) {
      throw "error " + error.response.status;
    }
  }
}

module.exports = {
  API,
};
