
function $main({ option1, option2 } = {}) {
  console.log(option1, option2);
  return {};
}

function addReadOnlyProperties(target, source) {
  Object.keys(source).forEach(key => Object.defineProperty(target, key, {
    value: source[key],
    configurable: false,
    writable: false,
  }));
}

addReadOnlyProperties($main, {
  // include library properties here
});

module.exports = $main;
