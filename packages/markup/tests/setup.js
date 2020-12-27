// import jest from 'jest';

jest.useFakeTimers();

var timeoutDummy = function setTimeoutDummy(fn, ms) {
  fn();
};

self.setTimeoutDummy = window.setTimeoutDummy = global.setTimeoutDummy = timeoutDummy;
