import moxios from 'moxios';
import sinon from 'sinon';
import api from '.';

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
});
