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

    expect.assertions(1);
    a.getStatus().then(onFulfilled)
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

    expect.assertions(1);
    a.createAccount({
      email: 'bob@gmail.com',
      password: 'blah',
      name: 'bob',
    }).then(onFulfilled)
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

    expect.assertions(1);
    a.login({
      email: 'bob@gmail.com',
      password: 'blah',
    }).then(onFulfilled)
      .catch(err => {
        expect(err).toEqual(Error('error 500'));
        done();
      });
  });
});

describe('verify', () => {
  test('ok', (done) => {
    moxios.stubRequest('/user/verify', {
      status: 200,
      response: true,
    });

    a.verify({
      hash: '1337h4x0r',
    }).then(onFulfilled);
    moxios.wait(() => {
      const response = onFulfilled.getCall(0).args[0];
      expect(response).toEqual(true);
      done();
    });
  });

  test('fail', (done) => {
    moxios.stubRequest('/user/verify', {
      status: 500,
      response: false,
    });

    expect.assertions(1);
    a.verify({
      hash: 'invalid',
    }).then(onFulfilled)
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
        clubID: 'UBC Launchpad',
      },
    });

    a.createClub({
      clubID: 'UBC Launchpad',
      description: 'The best software engineering club',
    }).then(onFulfilled);
    moxios.wait(() => {
      const response = onFulfilled.getCall(0).args[0];
      expect(response).toEqual('UBC Launchpad');
      done();
    });
  });

  test('fail', (done) => {
    moxios.stubRequest('/club/create', {
      status: 404,
      response: {
        clubID: 'UBC Launchpad',
      },
    });

    expect.assertions(1);
    a.createClub({
      clubID: 'UBC Launchpad',
      description: 'The best software engineering club',
    }).then(onFulfilled)
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
        period: '1234',
      },
    });

    a.createPeriod({
      period: '1234',
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
        period: '1234',
      },
    });

    expect.assertions(1);
    a.createPeriod({

    }).then(onFulfilled)
      .catch(err => {
        expect(err).toEqual(Error('error 400'));
        done();
      });
  });
});
