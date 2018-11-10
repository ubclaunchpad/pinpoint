import Lib from '.';

describe('lib', () => {
  test('constructor', () => {
    const a = new Lib({});
    expect(a.req).toBeTruthy();
  });
});
