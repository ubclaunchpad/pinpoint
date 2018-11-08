const Lib = require('.');

describe('lib', () => {
  test('constructor', () => {
    const a = new Lib({});
    expect(a.req).toBeTruthy();
  });
});
