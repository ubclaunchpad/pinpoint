const moxios = require('moxios');
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
});
