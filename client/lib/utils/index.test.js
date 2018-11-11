const utils = require('.');

describe('newRequester', () => {
  test('no params', () => {
    const req = utils.newRequester();
    expect(req).toBeTruthy();
  });
});
