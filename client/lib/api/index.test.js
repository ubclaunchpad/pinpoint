const moxios = require('moxios');
const sinon = require('sinon');
const api = require('.');
let a;
let onFulfilled;

beforeEach(() => {
  moxios.install();
  a = new api.API();
  onFulfilled = sinon.spy();
});

afterEach(() => {
  moxios.uninstall();
});

describe('API', () => {
  test('constructor', () => {
    expect(a.req).toBeTruthy();
  });
});

describe('getStatus', () => {
  test('ok', (done) => {
    moxios.stubRequest('/status', {
      status: 200,
      response: {
        resp: 'active',
      },
    });

    a.getStatus().then(onFulfilled);
    moxios.wait(() => {
      const response = onFulfilled.getCall(0).args[0];
      expect(response).toEqual('active');
      done();
    });
  });

  test('fail', (done) => {
    moxios.stubRequest('/status', {
      status: 400,
      response: {},
    });

    a.getStatus().then(onFulfilled)
      .then(() => {
        expect(true).toBe(false);
      })
      .catch(err => {
        expect(err).toEqual(Error('error 400'));
        done();
      });
  });
});

describe('createAccount', () => {
  test('ok', (done) => {
    moxios.stubRequest('/user/create', {
      status: 200,
      response: {
        email: 'bob@gmail.com',
      },
    });

    a.createAccount({
      email: 'bob@gmail.com',
      password: 'blah',
      name: 'bob',
    }).then(onFulfilled);
    moxios.wait(() => {
      const response = onFulfilled.getCall(0).args[0];
      expect(response).toEqual('bob@gmail.com');
      done();
    });
  });

  test('fail', (done) => {
    moxios.stubRequest('/user/create', {
      status: 300,
      response: {},
    });

    a.createAccount({
      email: 'bob@gmail.com',
      password: 'blah',
      name: 'bob',
    }).then(onFulfilled)
      .then(() => {
        expect(true).toBe(false);
      })
      .catch(err => {
        expect(err).toEqual(Error('error 300'));
        done();
      });
  });
});

describe('login', () => {
  test('ok', (done) => {
    moxios.stubRequest('/user/login', {
      status: 200,
      response: {
        token: '1234',
      },
    });

    a.login({
      email: 'bob@gmail.com',
      password: 'blah',
    }).then(onFulfilled);
    moxios.wait(() => {
      const response = onFulfilled.getCall(0).args[0];
      expect(response).toEqual('1234');
      done();
    });
  });
  test('fail', (done) => {
    moxios.stubRequest('/user/login', {
      status: 500,
      response: {
        token: '1234',
      },
    });

    a.login({
      email: 'bob@gmail.com',
      password: 'blah',
    }).then(onFulfilled)
      .then(() => {
        expect(true).toBe(false);
      })
      .catch(err => {
        expect(err).toEqual(Error('error 500'));
        done();
      });
  });
});

describe('createClub', () => {
  test('ok', (done) => {
    moxios.stubRequest('/club/create', {
      status: 200,
      response: {
        ClubID: '1234',
      },
    });

    a.createClub({
      name: 'UBC Launchpad',
      desc: 'The best software engineering club',
    }).then(onFulfilled);
    moxios.wait(() => {
      const response = onFulfilled.getCall(0).args[0];
      expect(response).toEqual('1234');
      done();
    });
  });

  test('fail', (done) => {
    moxios.stubRequest('/club/create', {
      status: 404,
      response: {
        ClubID: '1234',
      },
    });

    a.createClub({
      name: 'UBC Launchpad',
      desc: 'The best software engineering club',
    }).then(onFulfilled)
      .then(() => {
        expect(true).toBe(false);
      })
      .catch(err => {
        expect(err).toEqual(Error('error 404'));
        done();
      });
  });
});

describe('createPeriod', () => {
  test('ok', (done) => {
    moxios.stubRequest('/club/period/create', {
      status: 200,
      response: {
        PeriodID: '1234',
      },
    });

    a.createPeriod({
      name: 'Winter Semester',
      start: '2018-08-09',
      end: '2018-08-12',
    }).then(onFulfilled);
    moxios.wait(() => {
      const response = onFulfilled.getCall(0).args[0];
      expect(response).toEqual('1234');
      done();
    });
  });

  test('fail', (done) => {
    moxios.stubRequest('/club/period/create', {
      status: 400,
      response: {
        PeriodID: '1234',
      },
    });

    a.createPeriod({
      name: 'Winter Semester',
      start: '2018-08-09',
      end: '2018-08-12',
    }).then(onFulfilled)
      .then(() => {
        expect(true).toBe(false);
      })
      .catch(err => {
        expect(err).toEqual(Error('error 400'));
        done();
      });
  });
});

describe('createPeriodFail', () => {
  test('ok', (done) => {
    moxios.stubRequest('/club/period/create', {
      status: 400,
      response: {
        PeriodID: '1234',
      },
    });

    a.createPeriod({
      name: 'Winter Semester',
      start: '2018-08-09',
      end: '2018-08-12',
    }).then(onFulfilled)
      .then(() => {
        expect(true).toBe(false);
      })
      .catch(err => {
        expect(err).toEqual(Error('error 400'));
        done();
      });
  });
});
