const moxios = require('moxios');
const sinon = require('sinon');
const api = require('.');

beforeEach(() => {
  moxios.install();
});

afterEach(() => {
  moxios.uninstall();
});

describe('API', () => {
  test('constructor', () => {
    const a = new api.API();
    expect(a.req).toBeTruthy();
  });

  describe('getStatus', () => {
    test('ok', (done) => {
      const a = new api.API();
      moxios.stubRequest('/status', {
        status: 200,
        response: {},
      });

      const onFulfilled = sinon.spy();
      a.getStatus().then(onFulfilled);
      moxios.wait(() => {
        const response = onFulfilled.getCall(0).args[0];
        expect(response.status).toEqual(200);
        done();
      });
    });
  });

  describe('createAccount', () => {
    test('ok', (done) => {
      const a = new api.API();
      moxios.stubRequest('/user/create', {
        status: 200,
        response: {},
      });

      const onFulfilled = sinon.spy();
      a.createAccount({
        email: 'bob@gmail.com',
        password: 'blah',
        name: 'bob',
      }).then(onFulfilled);
      moxios.wait(() => {
        const response = onFulfilled.getCall(0).args[0];
        expect(response.status).toEqual(200);
        done();
      });
    });
  });

  describe('login', () => {
    test('ok', (done) => {
      const a = new api.API();
      moxios.stubRequest('/user/login', {
        status: 200,
        response: {},
      });

      const onFulfilled = sinon.spy();
      a.login({
        email: 'bob@gmail.com',
        password: 'blah',
      }).then(onFulfilled);
      moxios.wait(() => {
        const response = onFulfilled.getCall(0).args[0];
        expect(response.status).toEqual(200);
        done();
      });
    });
  });

  describe('createClub', () => {
    test('ok', (done) => {
      const a = new api.API();
      moxios.stubRequest('/club/create', {
        status: 200,
        response: {},
      });

      const onFulfilled = sinon.spy();
      a.createClub({
        name: 'UBC Launchpad',
        desc: 'The best software engineering club',
      }).then(onFulfilled);
      moxios.wait(() => {
        const response = onFulfilled.getCall(0).args[0];
        expect(response.status).toEqual(200);
        done();
      });
    });
  });

  describe('createPeriod', () => {
    test('ok', (done) => {
      const a = new api.API();
      moxios.stubRequest('/club/period/create', {
        status: 200,
        response: {},
      });

      const onFulfilled = sinon.spy();
      a.createPeriod({
        name: 'Winter Semester',
        start: '2018-08-09',
        end: '2018-08-12',
      }).then(onFulfilled);
      moxios.wait(() => {
        const response = onFulfilled.getCall(0).args[0];
        expect(response.status).toEqual(200);
        done();
      });
    });
  });
});
