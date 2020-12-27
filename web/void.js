(function(){function r(e,n,t){function o(i,f){if(!n[i]){if(!e[i]){var c="function"==typeof require&&require;if(!f&&c)return c(i,!0);if(u)return u(i,!0);var a=new Error("Cannot find module '"+i+"'");throw a.code="MODULE_NOT_FOUND",a}var p=n[i]={exports:{}};e[i][0].call(p.exports,function(r){var n=e[i][1][r];return o(n||r)},p,p.exports,r,e,n,t)}return n[i].exports}for(var u="function"==typeof require&&require,i=0;i<t.length;i++)o(t[i]);return o}return r})()({1:[function(require,module,exports){
// shim for using process in browser
var process = module.exports = {};

// cached from whatever global is present so that test runners that stub it
// don't break things.  But we need to wrap it in a try catch in case it is
// wrapped in strict mode code which doesn't define any globals.  It's inside a
// function because try/catches deoptimize in certain engines.

var cachedSetTimeout;
var cachedClearTimeout;

function defaultSetTimout() {
    throw new Error('setTimeout has not been defined');
}
function defaultClearTimeout () {
    throw new Error('clearTimeout has not been defined');
}
(function () {
    try {
        if (typeof setTimeout === 'function') {
            cachedSetTimeout = setTimeout;
        } else {
            cachedSetTimeout = defaultSetTimout;
        }
    } catch (e) {
        cachedSetTimeout = defaultSetTimout;
    }
    try {
        if (typeof clearTimeout === 'function') {
            cachedClearTimeout = clearTimeout;
        } else {
            cachedClearTimeout = defaultClearTimeout;
        }
    } catch (e) {
        cachedClearTimeout = defaultClearTimeout;
    }
} ())
function runTimeout(fun) {
    if (cachedSetTimeout === setTimeout) {
        //normal enviroments in sane situations
        return setTimeout(fun, 0);
    }
    // if setTimeout wasn't available but was latter defined
    if ((cachedSetTimeout === defaultSetTimout || !cachedSetTimeout) && setTimeout) {
        cachedSetTimeout = setTimeout;
        return setTimeout(fun, 0);
    }
    try {
        // when when somebody has screwed with setTimeout but no I.E. maddness
        return cachedSetTimeout(fun, 0);
    } catch(e){
        try {
            // When we are in I.E. but the script has been evaled so I.E. doesn't trust the global object when called normally
            return cachedSetTimeout.call(null, fun, 0);
        } catch(e){
            // same as above but when it's a version of I.E. that must have the global object for 'this', hopfully our context correct otherwise it will throw a global error
            return cachedSetTimeout.call(this, fun, 0);
        }
    }


}
function runClearTimeout(marker) {
    if (cachedClearTimeout === clearTimeout) {
        //normal enviroments in sane situations
        return clearTimeout(marker);
    }
    // if clearTimeout wasn't available but was latter defined
    if ((cachedClearTimeout === defaultClearTimeout || !cachedClearTimeout) && clearTimeout) {
        cachedClearTimeout = clearTimeout;
        return clearTimeout(marker);
    }
    try {
        // when when somebody has screwed with setTimeout but no I.E. maddness
        return cachedClearTimeout(marker);
    } catch (e){
        try {
            // When we are in I.E. but the script has been evaled so I.E. doesn't  trust the global object when called normally
            return cachedClearTimeout.call(null, marker);
        } catch (e){
            // same as above but when it's a version of I.E. that must have the global object for 'this', hopfully our context correct otherwise it will throw a global error.
            // Some versions of I.E. have different rules for clearTimeout vs setTimeout
            return cachedClearTimeout.call(this, marker);
        }
    }



}
var queue = [];
var draining = false;
var currentQueue;
var queueIndex = -1;

function cleanUpNextTick() {
    if (!draining || !currentQueue) {
        return;
    }
    draining = false;
    if (currentQueue.length) {
        queue = currentQueue.concat(queue);
    } else {
        queueIndex = -1;
    }
    if (queue.length) {
        drainQueue();
    }
}

function drainQueue() {
    if (draining) {
        return;
    }
    var timeout = runTimeout(cleanUpNextTick);
    draining = true;

    var len = queue.length;
    while(len) {
        currentQueue = queue;
        queue = [];
        while (++queueIndex < len) {
            if (currentQueue) {
                currentQueue[queueIndex].run();
            }
        }
        queueIndex = -1;
        len = queue.length;
    }
    currentQueue = null;
    draining = false;
    runClearTimeout(timeout);
}

process.nextTick = function (fun) {
    var args = new Array(arguments.length - 1);
    if (arguments.length > 1) {
        for (var i = 1; i < arguments.length; i++) {
            args[i - 1] = arguments[i];
        }
    }
    queue.push(new Item(fun, args));
    if (queue.length === 1 && !draining) {
        runTimeout(drainQueue);
    }
};

// v8 likes predictible objects
function Item(fun, array) {
    this.fun = fun;
    this.array = array;
}
Item.prototype.run = function () {
    this.fun.apply(null, this.array);
};
process.title = 'browser';
process.browser = true;
process.env = {};
process.argv = [];
process.version = ''; // empty string to agroundlayer regexp issues
process.versions = {};

function noop() {}

process.on = noop;
process.addListener = noop;
process.once = noop;
process.off = noop;
process.removeListener = noop;
process.removeAllListeners = noop;
process.emit = noop;
process.prependListener = noop;
process.prependOnceListener = noop;

process.listeners = function (name) { return [] }

process.binding = function (name) {
    throw new Error('process.binding is not supported');
};

process.cwd = function () { return '/' };
process.chdir = function (dir) {
    throw new Error('process.chdir is not supported');
};
process.umask = function() { return 0; };

},{}],2:[function(require,module,exports){
(function (setImmediate){
'use strict';

/**
 * @this {Promise}
 */
function finallyConstructor(callback) {
  var constructor = this.constructor;
  return this.then(
    function(value) {
      return constructor.resolve(callback()).then(function() {
        return value;
      });
    },
    function(reason) {
      return constructor.resolve(callback()).then(function() {
        return constructor.reject(reason);
      });
    }
  );
}

// Store setTimeout reference so promise-polyfill will be unaffected by
// other code modifying setTimeout (like sinon.useFakeTimers())
var setTimeoutFunc = setTimeout;

function noop() {}

// Polyfill for Function.prototype.bind
function bind(fn, thisArg) {
  return function() {
    fn.apply(thisArg, arguments);
  };
}

/**
 * @constructor
 * @param {Function} fn
 */
function Promise(fn) {
  if (!(this instanceof Promise))
    throw new TypeError('Promises must be constructed via new');
  if (typeof fn !== 'function') throw new TypeError('not a function');
  /** @type {!number} */
  this._state = 0;
  /** @type {!boolean} */
  this._handled = false;
  /** @type {Promise|undefined} */
  this._value = undefined;
  /** @type {!Array<!Function>} */
  this._deferreds = [];

  doResolve(fn, this);
}

function handle(self, deferred) {
  while (self._state === 3) {
    self = self._value;
  }
  if (self._state === 0) {
    self._deferreds.push(deferred);
    return;
  }
  self._handled = true;
  Promise._immediateFn(function() {
    var cb = self._state === 1 ? deferred.onFulfilled : deferred.onRejected;
    if (cb === null) {
      (self._state === 1 ? resolve : reject)(deferred.promise, self._value);
      return;
    }
    var ret;
    try {
      ret = cb(self._value);
    } catch (e) {
      reject(deferred.promise, e);
      return;
    }
    resolve(deferred.promise, ret);
  });
}

function resolve(self, newValue) {
  try {
    // Promise Resolution Procedure: https://github.com/promises-aplus/promises-spec#the-promise-resolution-procedure
    if (newValue === self)
      throw new TypeError('A promise cannot be resolved with itself.');
    if (
      newValue &&
      (typeof newValue === 'object' || typeof newValue === 'function')
    ) {
      var then = newValue.then;
      if (newValue instanceof Promise) {
        self._state = 3;
        self._value = newValue;
        finale(self);
        return;
      } else if (typeof then === 'function') {
        doResolve(bind(then, newValue), self);
        return;
      }
    }
    self._state = 1;
    self._value = newValue;
    finale(self);
  } catch (e) {
    reject(self, e);
  }
}

function reject(self, newValue) {
  self._state = 2;
  self._value = newValue;
  finale(self);
}

function finale(self) {
  if (self._state === 2 && self._deferreds.length === 0) {
    Promise._immediateFn(function() {
      if (!self._handled) {
        Promise._unhandledRejectionFn(self._value);
      }
    });
  }

  for (var i = 0, len = self._deferreds.length; i < len; i++) {
    handle(self, self._deferreds[i]);
  }
  self._deferreds = null;
}

/**
 * @constructor
 */
function Handler(onFulfilled, onRejected, promise) {
  this.onFulfilled = typeof onFulfilled === 'function' ? onFulfilled : null;
  this.onRejected = typeof onRejected === 'function' ? onRejected : null;
  this.promise = promise;
}

/**
 * Take a potentially misbehaving resolver function and make sure
 * onFulfilled and onRejected are only called once.
 *
 * Makes no guarantees about asynchrony.
 */
function doResolve(fn, self) {
  var done = false;
  try {
    fn(
      function(value) {
        if (done) return;
        done = true;
        resolve(self, value);
      },
      function(reason) {
        if (done) return;
        done = true;
        reject(self, reason);
      }
    );
  } catch (ex) {
    if (done) return;
    done = true;
    reject(self, ex);
  }
}

Promise.prototype['catch'] = function(onRejected) {
  return this.then(null, onRejected);
};

Promise.prototype.then = function(onFulfilled, onRejected) {
  // @ts-ignore
  var prom = new this.constructor(noop);

  handle(this, new Handler(onFulfilled, onRejected, prom));
  return prom;
};

Promise.prototype['finally'] = finallyConstructor;

Promise.all = function(arr) {
  return new Promise(function(resolve, reject) {
    if (!arr || typeof arr.length === 'undefined')
      throw new TypeError('Promise.all accepts an array');
    var args = Array.prototype.slice.call(arr);
    if (args.length === 0) return resolve([]);
    var remaining = args.length;

    function res(i, val) {
      try {
        if (val && (typeof val === 'object' || typeof val === 'function')) {
          var then = val.then;
          if (typeof then === 'function') {
            then.call(
              val,
              function(val) {
                res(i, val);
              },
              reject
            );
            return;
          }
        }
        args[i] = val;
        if (--remaining === 0) {
          resolve(args);
        }
      } catch (ex) {
        reject(ex);
      }
    }

    for (var i = 0; i < args.length; i++) {
      res(i, args[i]);
    }
  });
};

Promise.resolve = function(value) {
  if (value && typeof value === 'object' && value.constructor === Promise) {
    return value;
  }

  return new Promise(function(resolve) {
    resolve(value);
  });
};

Promise.reject = function(value) {
  return new Promise(function(resolve, reject) {
    reject(value);
  });
};

Promise.race = function(values) {
  return new Promise(function(resolve, reject) {
    for (var i = 0, len = values.length; i < len; i++) {
      values[i].then(resolve, reject);
    }
  });
};

// Use polyfill for setImmediate for performance gains
Promise._immediateFn =
  (typeof setImmediate === 'function' &&
    function(fn) {
      setImmediate(fn);
    }) ||
  function(fn) {
    setTimeoutFunc(fn, 0);
  };

Promise._unhandledRejectionFn = function _unhandledRejectionFn(err) {
  if (typeof console !== 'undefined' && console) {
    console.warn('Possible Unhandled Promise Rejection:', err); // eslint-disable-line no-console
  }
};

module.exports = Promise;

}).call(this,require("timers").setImmediate)

},{"timers":3}],3:[function(require,module,exports){
(function (setImmediate,clearImmediate){
var nextTick = require('process/browser.js').nextTick;
var apply = Function.prototype.apply;
var slice = Array.prototype.slice;
var immediateIds = {};
var nextImmediateId = 0;

// DOM APIs, for completeness

exports.setTimeout = function() {
  return new Timeout(apply.call(setTimeout, window, arguments), clearTimeout);
};
exports.setInterval = function() {
  return new Timeout(apply.call(setInterval, window, arguments), clearInterval);
};
exports.clearTimeout =
exports.clearInterval = function(timeout) { timeout.close(); };

function Timeout(id, clearFn) {
  this._id = id;
  this._clearFn = clearFn;
}
Timeout.prototype.unref = Timeout.prototype.ref = function() {};
Timeout.prototype.close = function() {
  this._clearFn.call(window, this._id);
};

// Does not start the time, just sets up the members needed.
exports.enroll = function(item, msecs) {
  clearTimeout(item._idleTimeoutId);
  item._idleTimeout = msecs;
};

exports.unenroll = function(item) {
  clearTimeout(item._idleTimeoutId);
  item._idleTimeout = -1;
};

exports._unrefActive = exports.active = function(item) {
  clearTimeout(item._idleTimeoutId);

  var msecs = item._idleTimeout;
  if (msecs >= 0) {
    item._idleTimeoutId = setTimeout(function onTimeout() {
      if (item._onTimeout)
        item._onTimeout();
    }, msecs);
  }
};

// That's not how node.js implements it but the exposed api is the same.
exports.setImmediate = typeof setImmediate === "function" ? setImmediate : function(fn) {
  var id = nextImmediateId++;
  var args = arguments.length < 2 ? false : slice.call(arguments, 1);

  immediateIds[id] = true;

  nextTick(function onNextTick() {
    if (immediateIds[id]) {
      // fn.call() is faster so we optimize for the common use-case
      // @see http://jsperf.com/call-apply-segu
      if (args) {
        fn.apply(null, args);
      } else {
        fn.call(null);
      }
      // Prevent ids from leaking
      exports.clearImmediate(id);
    }
  });

  return id;
};

exports.clearImmediate = typeof clearImmediate === "function" ? clearImmediate : function(id) {
  delete immediateIds[id];
};
}).call(this,require("timers").setImmediate,require("timers").clearImmediate)

},{"process/browser.js":1,"timers":3}],4:[function(require,module,exports){
(function (global, factory) {
  typeof exports === 'object' && typeof module !== 'undefined' ? factory(exports) :
  typeof define === 'function' && define.amd ? define(['exports'], factory) :
  (factory((global.WHATWGFetch = {})));
}(this, (function (exports) { 'use strict';

  var support = {
    searchParams: 'URLSearchParams' in self,
    iterable: 'Symbol' in self && 'iterator' in Symbol,
    blob:
      'FileReader' in self &&
      'Blob' in self &&
      (function() {
        try {
          new Blob();
          return true
        } catch (e) {
          return false
        }
      })(),
    formData: 'FormData' in self,
    arrayBuffer: 'ArrayBuffer' in self
  };

  function isDataView(obj) {
    return obj && DataView.prototype.isPrototypeOf(obj)
  }

  if (support.arrayBuffer) {
    var viewClasses = [
      '[object Int8Array]',
      '[object Uint8Array]',
      '[object Uint8ClampedArray]',
      '[object Int16Array]',
      '[object Uint16Array]',
      '[object Int32Array]',
      '[object Uint32Array]',
      '[object Float32Array]',
      '[object Float64Array]'
    ];

    var isArrayBufferView =
      ArrayBuffer.isView ||
      function(obj) {
        return obj && viewClasses.indexOf(Object.prototype.toString.call(obj)) > -1
      };
  }

  function normalizeName(name) {
    if (typeof name !== 'string') {
      name = String(name);
    }
    if (/[^a-z0-9\-#$%&'*+.^_`|~]/i.test(name)) {
      throw new TypeError('Invalid character in header field name')
    }
    return name.toLowerCase()
  }

  function normalizeValue(value) {
    if (typeof value !== 'string') {
      value = String(value);
    }
    return value
  }

  // Build a destructive iterator for the value list
  function iteratorFor(items) {
    var iterator = {
      next: function() {
        var value = items.shift();
        return {done: value === undefined, value: value}
      }
    };

    if (support.iterable) {
      iterator[Symbol.iterator] = function() {
        return iterator
      };
    }

    return iterator
  }

  function Headers(headers) {
    this.map = {};

    if (headers instanceof Headers) {
      headers.forEach(function(value, name) {
        this.append(name, value);
      }, this);
    } else if (Array.isArray(headers)) {
      headers.forEach(function(header) {
        this.append(header[0], header[1]);
      }, this);
    } else if (headers) {
      Object.getOwnPropertyNames(headers).forEach(function(name) {
        this.append(name, headers[name]);
      }, this);
    }
  }

  Headers.prototype.append = function(name, value) {
    name = normalizeName(name);
    value = normalizeValue(value);
    var oldValue = this.map[name];
    this.map[name] = oldValue ? oldValue + ', ' + value : value;
  };

  Headers.prototype['delete'] = function(name) {
    delete this.map[normalizeName(name)];
  };

  Headers.prototype.get = function(name) {
    name = normalizeName(name);
    return this.has(name) ? this.map[name] : null
  };

  Headers.prototype.has = function(name) {
    return this.map.hasOwnProperty(normalizeName(name))
  };

  Headers.prototype.set = function(name, value) {
    this.map[normalizeName(name)] = normalizeValue(value);
  };

  Headers.prototype.forEach = function(callback, thisArg) {
    for (var name in this.map) {
      if (this.map.hasOwnProperty(name)) {
        callback.call(thisArg, this.map[name], name, this);
      }
    }
  };

  Headers.prototype.keys = function() {
    var items = [];
    this.forEach(function(value, name) {
      items.push(name);
    });
    return iteratorFor(items)
  };

  Headers.prototype.values = function() {
    var items = [];
    this.forEach(function(value) {
      items.push(value);
    });
    return iteratorFor(items)
  };

  Headers.prototype.entries = function() {
    var items = [];
    this.forEach(function(value, name) {
      items.push([name, value]);
    });
    return iteratorFor(items)
  };

  if (support.iterable) {
    Headers.prototype[Symbol.iterator] = Headers.prototype.entries;
  }

  function consumed(body) {
    if (body.bodyUsed) {
      return Promise.reject(new TypeError('Already read'))
    }
    body.bodyUsed = true;
  }

  function fileReaderReady(reader) {
    return new Promise(function(resolve, reject) {
      reader.onload = function() {
        resolve(reader.result);
      };
      reader.onerror = function() {
        reject(reader.error);
      };
    })
  }

  function readBlobAsArrayBuffer(blob) {
    var reader = new FileReader();
    var promise = fileReaderReady(reader);
    reader.readAsArrayBuffer(blob);
    return promise
  }

  function readBlobAsText(blob) {
    var reader = new FileReader();
    var promise = fileReaderReady(reader);
    reader.readAsText(blob);
    return promise
  }

  function readArrayBufferAsText(buf) {
    var view = new Uint8Array(buf);
    var chars = new Array(view.length);

    for (var i = 0; i < view.length; i++) {
      chars[i] = String.fromCharCode(view[i]);
    }
    return chars.join('')
  }

  function bufferClone(buf) {
    if (buf.slice) {
      return buf.slice(0)
    } else {
      var view = new Uint8Array(buf.byteLength);
      view.set(new Uint8Array(buf));
      return view.buffer
    }
  }

  function Body() {
    this.bodyUsed = false;

    this._initBody = function(body) {
      this._bodyInit = body;
      if (!body) {
        this._bodyText = '';
      } else if (typeof body === 'string') {
        this._bodyText = body;
      } else if (support.blob && Blob.prototype.isPrototypeOf(body)) {
        this._bodyBlob = body;
      } else if (support.formData && FormData.prototype.isPrototypeOf(body)) {
        this._bodyFormData = body;
      } else if (support.searchParams && URLSearchParams.prototype.isPrototypeOf(body)) {
        this._bodyText = body.toString();
      } else if (support.arrayBuffer && support.blob && isDataView(body)) {
        this._bodyArrayBuffer = bufferClone(body.buffer);
        // IE 10-11 can't handle a DataView body.
        this._bodyInit = new Blob([this._bodyArrayBuffer]);
      } else if (support.arrayBuffer && (ArrayBuffer.prototype.isPrototypeOf(body) || isArrayBufferView(body))) {
        this._bodyArrayBuffer = bufferClone(body);
      } else {
        this._bodyText = body = Object.prototype.toString.call(body);
      }

      if (!this.headers.get('content-type')) {
        if (typeof body === 'string') {
          this.headers.set('content-type', 'text/plain;charset=UTF-8');
        } else if (this._bodyBlob && this._bodyBlob.type) {
          this.headers.set('content-type', this._bodyBlob.type);
        } else if (support.searchParams && URLSearchParams.prototype.isPrototypeOf(body)) {
          this.headers.set('content-type', 'application/x-www-form-urlencoded;charset=UTF-8');
        }
      }
    };

    if (support.blob) {
      this.blob = function() {
        var rejected = consumed(this);
        if (rejected) {
          return rejected
        }

        if (this._bodyBlob) {
          return Promise.resolve(this._bodyBlob)
        } else if (this._bodyArrayBuffer) {
          return Promise.resolve(new Blob([this._bodyArrayBuffer]))
        } else if (this._bodyFormData) {
          throw new Error('could not read FormData body as blob')
        } else {
          return Promise.resolve(new Blob([this._bodyText]))
        }
      };

      this.arrayBuffer = function() {
        if (this._bodyArrayBuffer) {
          return consumed(this) || Promise.resolve(this._bodyArrayBuffer)
        } else {
          return this.blob().then(readBlobAsArrayBuffer)
        }
      };
    }

    this.text = function() {
      var rejected = consumed(this);
      if (rejected) {
        return rejected
      }

      if (this._bodyBlob) {
        return readBlobAsText(this._bodyBlob)
      } else if (this._bodyArrayBuffer) {
        return Promise.resolve(readArrayBufferAsText(this._bodyArrayBuffer))
      } else if (this._bodyFormData) {
        throw new Error('could not read FormData body as text')
      } else {
        return Promise.resolve(this._bodyText)
      }
    };

    if (support.formData) {
      this.formData = function() {
        return this.text().then(decode)
      };
    }

    this.json = function() {
      return this.text().then(JSON.parse)
    };

    return this
  }

  // HTTP methods whose capitalization should be normalized
  var methods = ['DELETE', 'GET', 'HEAD', 'OPTIONS', 'POST', 'PUT'];

  function normalizeMethod(method) {
    var upcased = method.toUpperCase();
    return methods.indexOf(upcased) > -1 ? upcased : method
  }

  function Request(input, options) {
    options = options || {};
    var body = options.body;

    if (input instanceof Request) {
      if (input.bodyUsed) {
        throw new TypeError('Already read')
      }
      this.url = input.url;
      this.credentials = input.credentials;
      if (!options.headers) {
        this.headers = new Headers(input.headers);
      }
      this.method = input.method;
      this.mode = input.mode;
      this.signal = input.signal;
      if (!body && input._bodyInit != null) {
        body = input._bodyInit;
        input.bodyUsed = true;
      }
    } else {
      this.url = String(input);
    }

    this.credentials = options.credentials || this.credentials || 'same-origin';
    if (options.headers || !this.headers) {
      this.headers = new Headers(options.headers);
    }
    this.method = normalizeMethod(options.method || this.method || 'GET');
    this.mode = options.mode || this.mode || null;
    this.signal = options.signal || this.signal;
    this.referrer = null;

    if ((this.method === 'GET' || this.method === 'HEAD') && body) {
      throw new TypeError('Body not allowed for GET or HEAD requests')
    }
    this._initBody(body);
  }

  Request.prototype.clone = function() {
    return new Request(this, {body: this._bodyInit})
  };

  function decode(body) {
    var form = new FormData();
    body
      .trim()
      .split('&')
      .forEach(function(bytes) {
        if (bytes) {
          var split = bytes.split('=');
          var name = split.shift().replace(/\+/g, ' ');
          var value = split.join('=').replace(/\+/g, ' ');
          form.append(decodeURIComponent(name), decodeURIComponent(value));
        }
      });
    return form
  }

  function parseHeaders(rawHeaders) {
    var headers = new Headers();
    // Replace instances of \r\n and \n followed by at least one space or horizontal tab with a space
    // https://tools.ietf.org/html/rfc7230#section-3.2
    var preProcessedHeaders = rawHeaders.replace(/\r?\n[\t ]+/g, ' ');
    preProcessedHeaders.split(/\r?\n/).forEach(function(line) {
      var parts = line.split(':');
      var key = parts.shift().trim();
      if (key) {
        var value = parts.join(':').trim();
        headers.append(key, value);
      }
    });
    return headers
  }

  Body.call(Request.prototype);

  function Response(bodyInit, options) {
    if (!options) {
      options = {};
    }

    this.type = 'default';
    this.status = options.status === undefined ? 200 : options.status;
    this.ok = this.status >= 200 && this.status < 300;
    this.statusText = 'statusText' in options ? options.statusText : 'OK';
    this.headers = new Headers(options.headers);
    this.url = options.url || '';
    this._initBody(bodyInit);
  }

  Body.call(Response.prototype);

  Response.prototype.clone = function() {
    return new Response(this._bodyInit, {
      status: this.status,
      statusText: this.statusText,
      headers: new Headers(this.headers),
      url: this.url
    })
  };

  Response.error = function() {
    var response = new Response(null, {status: 0, statusText: ''});
    response.type = 'error';
    return response
  };

  var redirectStatuses = [301, 302, 303, 307, 308];

  Response.redirect = function(url, status) {
    if (redirectStatuses.indexOf(status) === -1) {
      throw new RangeError('Invalid status code')
    }

    return new Response(null, {status: status, headers: {location: url}})
  };

  exports.DOMException = self.DOMException;
  try {
    new exports.DOMException();
  } catch (err) {
    exports.DOMException = function(message, name) {
      this.message = message;
      this.name = name;
      var error = Error(message);
      this.stack = error.stack;
    };
    exports.DOMException.prototype = Object.create(Error.prototype);
    exports.DOMException.prototype.constructor = exports.DOMException;
  }

  function fetch(input, init) {
    return new Promise(function(resolve, reject) {
      var request = new Request(input, init);

      if (request.signal && request.signal.aborted) {
        return reject(new exports.DOMException('Aborted', 'AbortError'))
      }

      var xhr = new XMLHttpRequest();

      function abortXhr() {
        xhr.abort();
      }

      xhr.onload = function() {
        var options = {
          status: xhr.status,
          statusText: xhr.statusText,
          headers: parseHeaders(xhr.getAllResponseHeaders() || '')
        };
        options.url = 'responseURL' in xhr ? xhr.responseURL : options.headers.get('X-Request-URL');
        var body = 'response' in xhr ? xhr.response : xhr.responseText;
        resolve(new Response(body, options));
      };

      xhr.onerror = function() {
        reject(new TypeError('Network request failed'));
      };

      xhr.ontimeout = function() {
        reject(new TypeError('Network request failed'));
      };

      xhr.onabort = function() {
        reject(new exports.DOMException('Aborted', 'AbortError'));
      };

      xhr.open(request.method, request.url, true);

      if (request.credentials === 'include') {
        xhr.withCredentials = true;
      } else if (request.credentials === 'omit') {
        xhr.withCredentials = false;
      }

      if ('responseType' in xhr && support.blob) {
        xhr.responseType = 'blob';
      }

      request.headers.forEach(function(value, name) {
        xhr.setRequestHeader(name, value);
      });

      if (request.signal) {
        request.signal.addEventListener('abort', abortXhr);

        xhr.onreadystatechange = function() {
          // DONE (success or failure)
          if (xhr.readyState === 4) {
            request.signal.removeEventListener('abort', abortXhr);
          }
        };
      }

      xhr.send(typeof request._bodyInit === 'undefined' ? null : request._bodyInit);
    })
  }

  fetch.polyfill = true;

  if (!self.fetch) {
    self.fetch = fetch;
    self.Headers = Headers;
    self.Request = Request;
    self.Response = Response;
  }

  exports.Headers = Headers;
  exports.Request = Request;
  exports.Response = Response;
  exports.fetch = fetch;

  Object.defineProperty(exports, '__esModule', { value: true });

})));

},{}],5:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AFrame = exports.AnimationQueue = groundlayer 0;
const rafPolyfill = require("./raf-polyfill");
class AnimationQueue {
    constructor() {
        this.skip = false;
        this.binded = false;
        this.requestAnimationID = -1;
        this.frames = new Array();
        this.bindCycle = this.cycle.bind(this);
        this.rafProvider = rafPolyfill.GetRAF();
    }
    new() {
        const newFrame = new AFrame(this.frames.length, this);
        this.frames.push(newFrame);
        this.bind();
        return newFrame;
    }
    add(f) {
        f.queueIndex = this.frames.length;
        f.queue = this;
        this.frames.push(f);
        this.bind();
    }
    resume() {
        this.skip = false;
        this.bind();
    }
    pause() {
        this.skip = true;
    }
    unbind() {
        if (!this.binded) {
            return null;
        }
        this.rafProvider.cancelAnimationFrame(this.requestAnimationID);
    }
    bind() {
        if (this.binded)
            return null;
        this.requestAnimationID = this.rafProvider.requestAnimationFrame(this.bindCycle, null);
        this.binded = true;
    }
    cycle(ms) {
        if (this.frames.length === 0) {
            this.binded = false;
            return;
        }
        this.frames.forEach(function (f) {
            if (!f.paused()) {
                f.animate(ms);
            }
        });
        this.bind();
    }
}
exports.AnimationQueue = AnimationQueue;
class AFrame {
    constructor(index, queue) {
        this.skip = false;
        this.queue = queue;
        this.queueIndex = index;
        this.callbacks = new Array();
    }
    add(callback) {
        this.callbacks.push(callback);
    }
    clear() {
        this.callbacks.length = 0;
    }
    paused() {
        return this.skip;
    }
    pause() {
        this.skip = true;
    }
    stop() {
        this.pause();
        if (this.queueIndex === -1) {
            return null;
        }
        if (this.queue.frames.length == 0) {
            this.queue = undefined;
            this.queueIndex = -1;
            return null;
        }
        const total = this.queue.frames.length;
        if (total == 1) {
            this.queue.frames.pop();
            this.queue = undefined;
            this.queueIndex = -1;
            return null;
        }
        this.queue.frames[this.queueIndex] = this.queue.frames[total - 1];
        this.queue.frames.length = total - 1;
        this.queue = undefined;
        this.queueIndex = -1;
    }
    animate(ts) {
        for (let index in this.callbacks) {
            const callback = this.callbacks[index];
            callback(ts);
        }
    }
}
exports.AFrame = AFrame;
class ChangeManager {
    constructor(queue) {
        this.reads = new Array();
        this.writes = new Array();
        this.readState = false;
        this.inReadCall = false;
        this.inWriteCall = false;
        this.scheduled = false;
        this.frame = queue.new();
    }
    static drainTasks(q, wrapper) {
        let task = q.shift();
        while (task) {
            if (wrapper !== null) {
                wrapper(task);
                task = q.shift();
                continue;
            }
            task();
            task = q.shift();
        }
    }
    mutate(fn) {
        this.writes.push(fn);
        this._schedule();
    }
    read(fn) {
        this.reads.push(fn);
        this._schedule();
    }
    _schedule() {
        if (this.scheduled) {
            return;
        }
        this.scheduled = true;
        this.frame.add(this._runTasks.bind(this));
    }
    _runTasks() {
        const readError = this._runReads();
        if (readError !== null && readError !== undefined) {
            this.scheduled = false;
            this._schedule();
            throw readError;
        }
        const writeError = this._runWrites();
        if (writeError !== null && writeError !== undefined) {
            this.scheduled = false;
            this._schedule();
            throw writeError;
        }
        if (this.reads.length > 0 || this.writes.length > 0) {
            this.scheduled = false;
            this._schedule();
            return;
        }
        this.scheduled = false;
    }
    _runReads() {
        try {
            ChangeManager.drainTasks(this.reads, this._execReads.bind(this));
        }
        catch (e) {
            return e;
        }
        return null;
    }
    _execReads(task) {
        this.inReadCall = true;
        task();
        this.inReadCall = false;
    }
    _runWrites() {
        try {
            ChangeManager.drainTasks(this.writes, this._execWrite.bind(this));
        }
        catch (e) {
            return e;
        }
        return null;
    }
    _execWrite(task) {
        this.inWriteCall = true;
        task();
        this.inWriteCall = false;
    }
}

},{"./raf-polyfill":7}],6:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.mountTo = groundlayer 0;
const rafPolyfill = require("./raf-polyfill");
const Animation = require("./anime");
const namespace = Object.assign(Object.assign({}, rafPolyfill), Animation);
function mountTo(parent) {
    parent.animations = namespace;
}
exports.mountTo = mountTo;
exports.default = namespace;

},{"./anime":5,"./raf-polyfill":7}],7:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.GetRAF = groundlayer 0;
const now = (function () {
    if (self.hasOwnProperty('performance')) {
        return ((performance.now ? performance.now.bind(performance) : null) ||
            (performance.mozNow ? performance.mozNow.bind(performance) : null) ||
            (performance.msNow ? performance.msNow.bind(performance) : null) ||
            (performance.oNow ? performance.oNow.bind(performance) : null) ||
            (performance.webkitNow ? performance.webkitNow.bind(performance) : null) ||
            Date.now.bind(Date));
    }
    return Date.now.bind(Date);
})();
const frameRate = 1000 / 60;
const vendors = ['ms', 'moz', 'webkit', 'o'];
function GetRAF() {
    let lastTime = 0;
    const mod = {};
    for (var x = 0; x < vendors.length && !window.requestAnimationFrame; ++x) {
        mod.requestAnimationFrame = window[vendors[x] + 'RequestAnimationFrame'];
        mod.cancelAnimationFrame =
            window[vendors[x] + 'CancelAnimationFrame'] || window[vendors[x] + 'RequestCancelAnimationFrame'];
    }
    if (!mod.requestAnimationFrame || !mod.cancelAnimationFrame)
        mod.requestAnimationFrame = function (callback, element) {
            const currTime = now();
            const timeToCall = Math.max(0, frameRate - (currTime - lastTime));
            const id = self.setTimeout(function () {
                try {
                    callback(currTime + timeToCall);
                }
                catch (e) {
                    console.log('Error: ', e);
                    setTimeout(function () {
                        throw e;
                    }, 0);
                }
            }, timeToCall);
            lastTime = currTime + timeToCall;
            return id;
        };
    if (!mod.cancelAnimationFrame) {
        mod.cancelAnimationFrame = function (id) {
            clearTimeout(id);
        };
    }
    return mod;
}
exports.GetRAF = GetRAF;

},{}],8:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DOMException = exports.Response = exports.Request = exports.Headers = exports.fetch = groundlayer 0;
const fetch = require("whatwg-fetch");
if (!self.fetch) {
    self.fetch = exports.fetch.fetch;
    self.Headers = exports.fetch.Headers;
    self.Request = exports.fetch.Request;
    self.Response = exports.fetch.Response;
    self.DOMException = exports.fetch.DOMException;
}
exports.fetch = self.fetch;
exports.Headers = self.Headers;
exports.Request = self.Request;
exports.Response = self.Response;
exports.DOMException = self.DOMException;

},{"whatwg-fetch":4}],9:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.default = {};

},{}],10:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.mountTo = void 0;
const fetch = require("./fetch");
const http = require("./http");
const websocket = require("./websocket");
const namespace = Object.assign(Object.assign(Object.assign({}, fetch), http), websocket);
function mountTo(parent) {
    parent.http = namespace;
}
exports.mountTo = mountTo;
exports.default = namespace;

},{"./fetch":8,"./http":9,"./websocket":11}],11:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Socket = exports.OneMinute = exports.OneSecond = void 0;
exports.OneSecond = 1000;
exports.OneMinute = exports.OneSecond * 60;
class Socket {
    constructor(addr, reader, exponent, maxReconnects, maxWait) {
        this.addr = addr;
        this.socket = null;
        this.reader = reader;
        this.maxWait = maxWait;
        this.userClosed = false;
        this.exponent = exponent;
        this.disconnected = false;
        this.attemptedConnects = 0;
        this.lastWait = exports.OneSecond;
        this.maxReconnect = maxReconnects;
        this.writeBuffer = new Array();
    }
    connect() {
        if (this.socket) {
            return;
        }
        if (this.attemptedConnects >= this.maxReconnect) {
            this.reader.Exhausted(this);
            return;
        }
        const socket = new WebSocket(this.addr);
        socket.addEventListener('open', this._opened.bind(this));
        socket.addEventListener('error', this._errored.bind(this));
        socket.addEventListener('message', this._messaged.bind(this));
        socket.addEventListener('close', this._disconnected.bind(this));
        this.socket = socket;
        this.disconnected = false;
    }
    send(message) {
        if (this.disconnected) {
            this.writeBuffer.push(message);
            return;
        }
        this.socket.send(message);
    }
    reset() {
        this.attemptedConnects = 0;
        this.lastWait = exports.OneSecond;
    }
    end() {
        this.userClosed = true;
        this.reader.Closed(this);
        this.socket.close();
        this.socket = null;
    }
    _disconnected(event) {
        this.reader.Disconnected(event, this);
        this.disconnected = true;
        this.socket = null;
        if (this.userClosed) {
            return;
        }
        let nextWait = this.lastWait;
        if (this.exponent) {
            nextWait = this.exponent(nextWait);
            if (nextWait > this.maxWait) {
                nextWait = this.maxWait;
            }
        }
        setTimeout(this.connect.bind(this), nextWait);
        this.attemptedConnects++;
    }
    _opened(event) {
        this.reader.Connected(event, this);
        while (this.writeBuffer.length > 0) {
            const message = this.writeBuffer.shift();
            this.socket.send(message);
        }
    }
    _errored(event) {
        this.reader.Errored(event, this);
    }
    _messaged(event) {
        this.reader.Message(event, this);
    }
}
exports.Socket = Socket;

},{}],12:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getMatchingNode = exports.toDataTransfer = exports.toGamepad = exports.toTouches = exports.toMediaStream = exports.toRotationData = exports.toMotionData = exports.toInputSourceCapability = exports.fromFile = exports.fromBlob = exports.removeFromNode = exports.createText = exports.createElement = exports.recordAttributes = exports.getNamespaceForTag = exports.applyAttributeTyped = exports.applySVGStyles = exports.applyStyles = exports.applyStyle = exports.applySVGStyle = exports.setStyleValue = exports.applyProp = exports.applyAttrs = exports.applyAttr = exports.getNamespace = exports.replaceNodeIf = exports.replaceNode = exports.insertBefore = exports.moveBefore = exports.getFocusedPath = exports.getActiveElement = exports.collectBreadthFirst = exports.toHTML = exports.eachNodeAndChild = exports.nodeListToArray = exports.eachChildAndNode = exports.reverseApplyEachNode = exports.applyEachNode = exports.applyEachChildNode = exports.applyChildNode = exports.findBreadthFirst = exports.collectDepthFirst = exports.findDepthFirst = exports.findNodeWithDepth = exports.findNodeWithBreadth = exports.collectNodeWithDepth = exports.collectNodeWithBreadth = exports.reverseFindNodeWithBreadth = exports.reverseCollectNodeWithBreadth = exports.getAncestry = exports.isText = exports.isElement = exports.isDocumentRoot = exports.COMMENT_NODE = exports.TEXT_NODE = exports.DOCUMENT_NODE = exports.DOCUMENT_FRAGMENT_NODE = exports.ELEMENT_NODE = void 0;
const utils_1 = require("./utils");
const exts = require("./extensions");
exports.ELEMENT_NODE = 1;
exports.DOCUMENT_FRAGMENT_NODE = 11;
exports.DOCUMENT_NODE = 9;
exports.TEXT_NODE = 3;
exports.COMMENT_NODE = 8;
const attributes = utils_1.createMap();
attributes['style'] = applyStyle;
function isDocumentRoot(node) {
    return node.nodeType === 11 || node.nodeType === 9;
}
exports.isDocumentRoot = isDocumentRoot;
function isElement(node) {
    return node.nodeType === 1;
}
exports.isElement = isElement;
function isText(node) {
    return node.nodeType === 3;
}
exports.isText = isText;
function getAncestry(node, root) {
    const ancestry = [];
    let cur = node;
    while (cur !== root) {
        const n = cur;
        ancestry.push(n);
        cur = n.parentNode;
    }
    return ancestry;
}
exports.getAncestry = getAncestry;
const getRootNode = Node.prototype.getRootNode ||
    function () {
        let cur = this;
        let prev = cur;
        while (cur) {
            prev = cur;
            cur = cur.parentNode;
        }
        return prev;
    };
function reverseCollectNodeWithBreadth(parent, matcher, matches) {
    let cur = parent.lastChild;
    while (cur) {
        if (matcher(cur)) {
            matches.push(cur);
        }
        cur = cur.previousSibling;
    }
}
exports.reverseCollectNodeWithBreadth = reverseCollectNodeWithBreadth;
function reverseFindNodeWithBreadth(parent, matcher) {
    let cur = parent.lastChild;
    while (cur) {
        if (matcher(cur)) {
            return cur;
        }
        cur = cur.previousSibling;
    }
    return null;
}
exports.reverseFindNodeWithBreadth = reverseFindNodeWithBreadth;
function collectNodeWithBreadth(parent, matcher, matches) {
    let cur = parent.firstChild;
    if (matcher(cur)) {
        matches.push(cur);
    }
    while (cur) {
        if (matcher(cur.nextSibling)) {
            matches.push(cur);
        }
        cur = cur.nextSibling;
    }
}
exports.collectNodeWithBreadth = collectNodeWithBreadth;
function collectNodeWithDepth(parent, matcher, matches) {
    let cur = parent.firstChild;
    while (cur) {
        if (matcher(cur)) {
            matches.push(cur);
        }
        cur = cur.firstChild;
    }
}
exports.collectNodeWithDepth = collectNodeWithDepth;
function findNodeWithBreadth(parent, matcher) {
    let cur = parent.firstChild;
    while (cur) {
        if (matcher(cur)) {
            return cur;
        }
        cur = cur.nextSibling;
    }
    return null;
}
exports.findNodeWithBreadth = findNodeWithBreadth;
function findNodeWithDepth(parent, matcher) {
    let cur = parent.firstChild;
    while (cur) {
        if (matcher(cur)) {
            return cur;
        }
        cur = cur.firstChild;
    }
    return null;
}
exports.findNodeWithDepth = findNodeWithDepth;
function findDepthFirst(parent, matcher) {
    let cur = parent.firstChild;
    while (cur) {
        const found = findNodeWithDepth(cur, matcher);
        if (found) {
            return found;
        }
        cur = cur.nextSibling;
    }
    return null;
}
exports.findDepthFirst = findDepthFirst;
function collectDepthFirst(parent, matcher, matches) {
    let cur = parent.firstChild;
    while (cur) {
        collectNodeWithDepth(cur, matcher, matches);
        cur = cur.nextSibling;
    }
    return;
}
exports.collectDepthFirst = collectDepthFirst;
function findBreadthFirst(parent, matcher) {
    let cur = parent.firstChild;
    while (cur) {
        const found = findNodeWithBreadth(cur, matcher);
        if (found) {
            return found;
        }
        cur = cur.firstChild;
    }
    return null;
}
exports.findBreadthFirst = findBreadthFirst;
function applyChildNode(parent, fn) {
    let cur = parent.firstChild;
    while (cur) {
        fn(cur);
        cur = cur.nextSibling;
    }
}
exports.applyChildNode = applyChildNode;
function applyEachChildNode(parent, fn) {
    let cur = parent.firstChild;
    while (cur) {
        fn(cur);
        applyEachChildNode(cur, fn);
        cur = cur.nextSibling;
    }
}
exports.applyEachChildNode = applyEachChildNode;
function applyEachNode(parent, fn) {
    fn(parent);
    let cur = parent.firstChild;
    while (cur) {
        applyEachNode(cur, fn);
        cur = cur.nextSibling;
    }
}
exports.applyEachNode = applyEachNode;
function reverseApplyEachNode(parent, fn) {
    let cur = parent.lastChild;
    while (cur) {
        reverseApplyEachNode(cur, fn);
        fn(cur);
        cur = cur.previousSibling;
    }
    fn(parent);
}
exports.reverseApplyEachNode = reverseApplyEachNode;
function eachChildAndNode(node, fn) {
    const list = node.childNodes;
    for (let i = 0; i < list.length; i++) {
        fn(list[i]);
    }
    fn(node);
}
exports.eachChildAndNode = eachChildAndNode;
function nodeListToArray(items) {
    const list = [];
    for (let i = 0; i < items.length; i++) {
        list.push(items[i]);
    }
    return list;
}
exports.nodeListToArray = nodeListToArray;
function eachNodeAndChild(node, fn) {
    fn(node);
    const list = node.childNodes;
    for (let i = 0; i < list.length; i++) {
        fn(list[i]);
    }
}
exports.eachNodeAndChild = eachNodeAndChild;
function toHTML(node, shallow) {
    const div = document.createElement('div');
    div.appendChild(node.cloneNode(!shallow));
    return div.innerHTML;
}
exports.toHTML = toHTML;
function collectBreadthFirst(parent, matcher, matches) {
    let cur = parent.firstChild;
    while (cur) {
        collectNodeWithBreadth(cur, matcher, matches);
        cur = cur.firstChild;
    }
    return;
}
exports.collectBreadthFirst = collectBreadthFirst;
function getActiveElement(node) {
    const root = getRootNode.call(node);
    return isDocumentRoot(root) ? root.activeElement : null;
}
exports.getActiveElement = getActiveElement;
function getFocusedPath(node, root) {
    const activeElement = getActiveElement(node);
    if (!activeElement || !node.contains(activeElement)) {
        return [];
    }
    return getAncestry(activeElement, root);
}
exports.getFocusedPath = getFocusedPath;
function moveBefore(parentNode, node, referenceNode) {
    const insertReferenceNode = node.nextSibling;
    let cur = referenceNode;
    while (cur !== null && cur !== node) {
        const next = cur.nextSibling;
        parentNode.insertBefore(cur, insertReferenceNode);
        cur = next;
    }
}
exports.moveBefore = moveBefore;
function insertBefore(parentNode, node, referenceNode) {
    if (referenceNode === null) {
        parentNode.appendChild(node);
        return null;
    }
    parentNode.insertBefore(node, referenceNode);
    return null;
}
exports.insertBefore = insertBefore;
function replaceNode(parentNode, node, replacement) {
    if (replacement === null) {
        return null;
    }
    parentNode.replaceChild(replacement, node);
    return null;
}
exports.replaceNode = replaceNode;
function replaceNodeIf(targetNode, replacement) {
    if (replacement === null) {
        return false;
    }
    const parent = targetNode.parentNode;
    if (!parent) {
        return false;
    }
    parent.replaceChild(replacement, targetNode);
    return true;
}
exports.replaceNodeIf = replaceNodeIf;
function getNamespace(name) {
    if (name.lastIndexOf('xml:', 0) === 0) {
        return 'http://www.w3.org/XML/1998/namespace';
    }
    if (name.lastIndexOf('xlink:', 0) === 0) {
        return 'http://www.w3.org/1999/xlink';
    }
    return undefined;
}
exports.getNamespace = getNamespace;
function applyAttr(el, name, value) {
    if (value == null) {
        el.removeAttribute(name);
    }
    else {
        const attrNS = getNamespace(name);
        if (attrNS) {
            el.setAttributeNS(attrNS, name, String(value));
        }
        else {
            el.setAttribute(name, String(value));
        }
    }
}
exports.applyAttr = applyAttr;
function applyAttrs(el, values) {
    for (let key in values) {
        if (values[key] == null) {
            el.removeAttribute(key);
            continue;
        }
        el.setAttribute(key, values[key]);
    }
}
exports.applyAttrs = applyAttrs;
function applyProp(el, name, value) {
    el[name] = value;
}
exports.applyProp = applyProp;
function setStyleValue(style, prop, value) {
    if (prop.indexOf('-') >= 0) {
        style.setProperty(prop, value);
    }
    else {
        style[prop] = value;
    }
}
exports.setStyleValue = setStyleValue;
function applySVGStyle(el, name, style) {
    if (typeof style === 'string') {
        el.style.cssText = style;
    }
    else {
        el.style.cssText = '';
        const elStyle = el.style;
        for (const prop in style) {
            if (utils_1.has(style, prop)) {
                setStyleValue(elStyle, prop, style[prop]);
            }
        }
    }
}
exports.applySVGStyle = applySVGStyle;
function applyStyle(el, name, style) {
    if (typeof style === 'string') {
        el.style.cssText = style;
    }
    else {
        el.style.cssText = '';
        const elStyle = el.style;
        for (const prop in style) {
            if (utils_1.has(style, prop)) {
                setStyleValue(elStyle, prop, style[prop]);
            }
        }
    }
}
exports.applyStyle = applyStyle;
function applyStyles(el, style) {
    if (typeof style === 'string') {
        el.style.cssText = style;
    }
    else {
        el.style.cssText = '';
        const elStyle = el.style;
        for (const prop in style) {
            if (utils_1.has(style, prop)) {
                setStyleValue(elStyle, prop, style[prop]);
            }
        }
    }
}
exports.applyStyles = applyStyles;
function applySVGStyles(el, style) {
    if (typeof style === 'string') {
        el.style.cssText = style;
    }
    else {
        el.style.cssText = '';
        const elStyle = el.style;
        for (const prop in style) {
            if (utils_1.has(style, prop)) {
                setStyleValue(elStyle, prop, style[prop]);
            }
        }
    }
}
exports.applySVGStyles = applySVGStyles;
function applyAttributeTyped(el, name, value) {
    const type = typeof value;
    if (type === 'object' || type === 'function') {
        applyProp(el, name, value);
    }
    else {
        applyAttr(el, name, value);
    }
}
exports.applyAttributeTyped = applyAttributeTyped;
function getNamespaceForTag(tag, parent) {
    if (tag === 'svg') {
        return 'http://www.w3.org/2000/svg';
    }
    if (tag === 'math') {
        return 'http://www.w3.org/1998/Math/MathML';
    }
    if (parent == null) {
        return null;
    }
    return parent.namespaceURI;
}
exports.getNamespaceForTag = getNamespaceForTag;
function recordAttributes(node) {
    const attrs = {};
    const attributes = node.attributes;
    const length = attributes.length;
    if (!length) {
        return attrs;
    }
    for (let i = 0, j = 0; i < length; i += 1, j += 2) {
        const attr = attributes[i];
        attrs[attr.name] = attr.value;
    }
    return attrs;
}
exports.recordAttributes = recordAttributes;
function createElement(doc, nameOrCtor, key, content, attributes, namespace) {
    let el;
    if (typeof nameOrCtor === 'function') {
        el = new nameOrCtor();
        return el;
    }
    namespace = namespace.trim();
    if (namespace.length > 0) {
        switch (nameOrCtor) {
            case 'svg':
                el = doc.createElementNS('http://www.w3.org/2000/svg', nameOrCtor);
                break;
            case 'math':
                el = doc.createElementNS('http://www.w3.org/1998/Math/MathML', nameOrCtor);
                break;
            default:
                el = doc.createElementNS(namespace, nameOrCtor);
        }
    }
    else {
        el = doc.createElement(nameOrCtor);
    }
    el.setAttribute('_key', key);
    if (attributes) {
        applyAttrs(el, attributes);
    }
    if (content.length > 0) {
        el.innerHTML = content;
    }
    return el;
}
exports.createElement = createElement;
function createText(doc, text, key) {
    const node = doc.createTextNode(text);
    exts.Objects.PatchWith(node, 'key', key);
    return node;
}
exports.createText = createText;
function removeFromNode(fromNode, endNode) {
    const parentNode = fromNode.parentNode;
    let child = fromNode;
    while (child !== endNode) {
        const next = child.nextSibling;
        parentNode.removeChild(child);
        child = next;
    }
}
exports.removeFromNode = removeFromNode;
function fromBlob(o) {
    if (o === null || o === undefined) {
        return;
    }
    let data = null;
    const fileReader = new FileReader();
    fileReader.onloadend = function () {
        data = new Uint8Array(fileReader.result);
    };
    fileReader.readAsArrayBuffer(o);
    return data;
}
exports.fromBlob = fromBlob;
function fromFile(o) {
    if (o === null || o === undefined) {
        return;
    }
    let data = null;
    const fileReader = new FileReader();
    fileReader.onloadend = function () {
        data = new Uint8Array(fileReader.result);
    };
    fileReader.readAsArrayBuffer(o);
    return data;
}
exports.fromFile = fromFile;
function toInputSourceCapability(o) {
    if (o === null || o === undefined) {
        return;
    }
    return {
        FiresTouchEvent: o.firesTouchEvent,
    };
}
exports.toInputSourceCapability = toInputSourceCapability;
function toMotionData(o) {
    let md = { X: 0.0, Y: 0.0, Z: 0.0 };
    if (o === null || o === undefined) {
        return md;
    }
    md.X = o.x;
    md.Y = o.y;
    md.Z = o.z;
    return md;
}
exports.toMotionData = toMotionData;
function toRotationData(o) {
    if (o === null || o === undefined) {
        return;
    }
    const md = {};
    md.Alpha = o.alpha;
    md.Beta = o.beta;
    md.Gamma = o.gamma;
    return md;
}
exports.toRotationData = toRotationData;
function toMediaStream(o) {
    if (o === null || o === undefined) {
        return;
    }
    const stream = { Audios: [], Videos: [] };
    stream.Active = o.active;
    stream.Ended = o.ended;
    stream.ID = o.id;
    stream.Audios = [];
    stream.Videos = [];
    let audioTracks = o.getAudioTracks();
    if (audioTracks !== null && audioTracks !== undefined) {
        for (let i = 0; i < audioTracks.length; i++) {
            let track = audioTracks[i];
            let settings = track.getSettings();
            stream.Audios.push({
                Enabled: track.enabled,
                ID: track.id,
                Kind: track.kind,
                Label: track.label,
                Muted: track.muted,
                ReadyState: track.readyState,
                Remote: track.remote,
                AudioSettings: {
                    ChannelCount: settings.channelCount.Int(),
                    EchoCancellation: settings.echoCancellation,
                    Latency: settings.latency,
                    SampleRate: settings.sampleRate.Int64(),
                    SampleSize: settings.sampleSize.Int64(),
                    Volume: settings.volume,
                    MediaTrackSettings: {
                        DeviceID: settings.deviceId,
                        GroupID: settings.groupId,
                    },
                },
            });
        }
    }
    let videosTracks = o.getVideoTracks();
    if (videosTracks !== null && videosTracks !== undefined) {
        for (let i = 0; i < videosTracks.length; i++) {
            let track = videosTracks[i];
            let settings = track.getSettings();
            stream.Videos.push({
                Enabled: track.enabled,
                ID: track.id,
                Kind: track.kind,
                Label: track.label,
                Muted: track.muted,
                ReadyState: track.readyState,
                Remote: track.remote,
                VideoSettings: {
                    AspectRatio: settings.aspectRation,
                    FrameRate: settings.frameRate,
                    Height: settings.height.Int64(),
                    Width: settings.width.Int64(),
                    FacingMode: settings.facingMode,
                    MediaTrackSettings: {
                        DeviceID: settings.deviceId,
                        GroupID: settings.groupId,
                    },
                },
            });
        }
    }
    return stream;
}
exports.toMediaStream = toMediaStream;
function toTouches(o) {
    if (o === null || o === undefined) {
        return;
    }
    const touches = [];
    for (let i = 0; i < o.length; i++) {
        let ev = o.item(i);
        touches.push({
            ClientX: ev.clientX,
            ClientY: ev.clientY,
            OffsetX: ev.offsetX,
            OffsetY: ev.offsetY,
            PageX: ev.pageX,
            PageY: ev.pageY,
            ScreenX: ev.screenX,
            ScreenY: ev.screenY,
            Identifier: ev.identifier,
        });
    }
    return touches;
}
exports.toTouches = toTouches;
function toGamepad(o) {
    let pad = {};
    if (o === null || o === undefined) {
        return pad;
    }
    pad.DisplayID = o.displayId;
    pad.ID = o.id;
    pad.Index = o.index.Int();
    pad.Mapping = o.mapping;
    pad.Connected = o.connected;
    pad.Timestamp = o.timestamp;
    pad.Axes = [];
    pad.Buttons = [];
    let axes = o.axes;
    if (axes !== null && axes !== undefined) {
        for (let i = 0; i < axes.length; i++) {
            pad.Axes.push(axes[i]);
        }
    }
    let buttons = o.buttons;
    if (buttons !== null && buttons !== undefined) {
        for (let i = 0; i < buttons.length; i++) {
            const button = buttons[i];
            pad.Buttons.push({
                Value: button.value,
                Pressed: button.pressed,
            });
        }
    }
    return pad;
}
exports.toGamepad = toGamepad;
function toDataTransfer(o) {
    if (o === null || o === undefined) {
        return;
    }
    let dt = {};
    dt.DropEffect = o.dropEffect;
    dt.EffectAllowed = o.effectAllowed;
    dt.Types = o.types;
    dt.Items = [];
    const items = o.items;
    if (items !== null && items !== undefined) {
        for (let i = 0; i < items.length; i++) {
            const item = items.DataTransferItem(i);
            dt.Items.push({
                Name: item.name,
                Size: item.size.Int(),
                Data: fromFile(item),
            });
        }
    }
    dt.Files = [];
    const files = o.files;
    if (files !== null && files !== undefined) {
        for (let i = 0; i < files.length; i++) {
            const item = files[i];
            dFiles.push({
                Name: item.name,
                Size: item.size.Int(),
                Data: fromFile(item),
            });
        }
    }
    return dt;
}
exports.toDataTransfer = toDataTransfer;
function getMatchingNode(matchNode, matcher) {
    if (!matchNode) {
        return null;
    }
    if (matcher(matchNode)) {
        return matchNode;
    }
    while ((matchNode = matchNode.nextSibling)) {
        if (matcher(matchNode)) {
            return matchNode;
        }
    }
    return null;
}
exports.getMatchingNode = getMatchingNode;

},{"./extensions":13,"./utils":17}],13:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Objects = void 0;
var Objects;
(function (Objects) {
    function PatchWith(elem, attrName, attrs) {
        elem[attrName] = attrs;
    }
    Objects.PatchWith = PatchWith;
    function GetAttrWith(elem, attrName) {
        const value = elem[attrName];
        return value ? value : '';
    }
    Objects.GetAttrWith = GetAttrWith;
    function isNullOrUndefined(elem) {
        return elem === null || elem === undefined;
    }
    Objects.isNullOrUndefined = isNullOrUndefined;
    function isAny(elem, ...values) {
        for (let index of values) {
            if (elem === index) {
                return true;
            }
        }
        return false;
    }
    Objects.isAny = isAny;
})(Objects = exports.Objects || (exports.Objects = {}));

},{}],14:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.mountTo = void 0;
const utils = require("./utils");
const exts = require("./extensions");
const patch = require("./patch");
const mount = require("./mount");
const dom = require("./dom");
const namespace = Object.assign(Object.assign(Object.assign(Object.assign(Object.assign({}, utils), exts), patch), dom), mount);
function mountTo(parent) {
    parent.markup = namespace;
}
exports.mountTo = mountTo;
exports.default = namespace;

},{"./dom":12,"./extensions":13,"./mount":15,"./patch":16,"./utils":17}],15:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DOMMount = void 0;
const dom = require("./dom");
const patch = require("./patch");
const utils = require("./utils");
class DOMMount {
    constructor(document, target, notifier) {
        if (typeof document !== 'object') {
            throw new Error('document should be an object');
        }
        this.doc = document;
        this.notifier = notifier;
        this.events = {};
        this.handler = this.handleEvent.bind(this);
        if (typeof target === 'string') {
            const targetSelector = target;
            const node = this.doc.querySelector(targetSelector);
            if (node === null || node === undefined) {
                throw new Error(`unable to locate node for ${{ targetSelector }}`);
            }
            this.mountNode = node;
        }
        else {
            this.mountNode = target;
        }
    }
    handleEvent(event) {
        if (!this.events[event.type]) {
            return;
        }
        event.stopPropagation();
        const target = event.target;
        if (target.nodeType !== dom.ELEMENT_NODE) {
            return;
        }
        const targetElement = target;
        const kebabEventName = 'event-' + utils.ToKebabCase(event.type);
        if (!targetElement.hasAttribute(kebabEventName)) {
            return;
        }
        const triggers = targetElement.getAttribute(kebabEventName);
        if (triggers == null) {
            return;
        }
        if (this.notifier && typeof this.notifier === 'function') {
            this.notifier(event, triggers.split('|'), targetElement);
        }
    }
    patch(change) {
        this.patchWith(change, patch.DefaultNodeDictator, patch.DefaultJSONDictator, patch.DefaultJSONMaker);
    }
    patchWith(change, nodeDictator, jsonDictator, jsonMaker) {
        if (change instanceof DocumentFragment) {
            const fragment = change;
            this.registerNodeEvents(fragment);
            patch.PatchDOMTree(fragment, this.mountNode, nodeDictator, false);
            return;
        }
        if (typeof change === 'string') {
            const node = document.createElement('div');
            node.innerHTML = change.trim();
            this.registerNodeEvents(node);
            patch.PatchDOMTree(node, this.mountNode, nodeDictator, false);
            return;
        }
        if (!patch.isJSONNode(change)) {
            return;
        }
        const node = change;
        this.registerJSONNodeEvents(node);
        patch.PatchJSONNodeTree(node, this.mountNode, jsonDictator, jsonMaker);
    }
    patchList(changes) {
        changes.forEach(this.patch.bind(this));
    }
    stream(changes) {
        const nodes = JSON.parse(changes);
        return this.streamList(nodes);
    }
    streamList(changes) {
        this.streamListWith(changes, patch.DefaultJSONDictator, patch.DefaultJSONMaker);
    }
    streamListWith(changes, dictator, maker) {
        changes.forEach(this.registerJSONNodeEvents.bind(this));
        patch.StreamJSONNodes(changes, this.mountNode, patch.DefaultJSONDictator, patch.DefaultJSONMaker);
    }
    registerNodeEvents(n) {
        const binder = this;
        dom.applyEachNode(n, function (node) {
            if (node.nodeType !== dom.ELEMENT_NODE) {
                return;
            }
            const elem = node;
            const events = elem.getAttribute('events');
            if (events) {
                events.split(' ').forEach(function (desc) {
                    const eventName = desc.substr(0, desc.length - 3);
                    binder.registerEvent(eventName);
                    switch (desc.substr(desc.length - 2, desc.length)) {
                        case '01':
                            break;
                        case '10':
                            n.addEventListener(eventName, DOMMount.preventDefault, false);
                            break;
                        case '11':
                            n.addEventListener(eventName, DOMMount.preventDefault, false);
                            break;
                    }
                });
            }
        });
    }
    registerJSONNodeEvents(node) {
        const binder = this;
        patch.applyJSONNodeFunction(node, function (n) {
            if (n.removed) {
                n.events.forEach(function (desc) {
                    binder.unregisterEvent(desc.Name);
                });
                return;
            }
            n.events.forEach(function (desc) {
                binder.registerEvent(desc.Name);
            });
        });
    }
    textContent() {
        return this.mountNode.textContent;
    }
    innerHTML() {
        return this.mountNode.innerHTML.trim();
    }
    Html() {
        return dom.toHTML(this.mountNode, false);
    }
    registerEvent(eventName) {
        if (this.events[eventName]) {
            return;
        }
        this.mountNode.addEventListener(eventName, this.handler, true);
        this.events[eventName] = true;
    }
    unregisterEvent(eventName) {
        if (!this.events[eventName]) {
            return;
        }
        this.mountNode.removeEventListener(eventName, this.handler, true);
        this.events[eventName] = false;
    }
    static preventDefault(event) {
        event.preventDefault();
    }
    static stopPropagation(event) {
        event.stopPropagation();
    }
}
exports.DOMMount = DOMMount;

},{"./dom":12,"./patch":16,"./utils":17}],16:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PatchDOMAttributes = exports.PatchDOMTree = exports.PatchJSONAttributes = exports.PatchTextCommentWithJSON = exports.ApplyStreamNode = exports.StreamJSONNodes = exports.PatchJSONNode = exports.PatchJSONNodeTree = exports.jsonMaker = exports.DefaultJSONMaker = exports.findElementParentbyRef = exports.findElementbyRef = exports.findElement = exports.isJSONNode = exports.applyJSONNodeKidsFunction = exports.applyJSONNodeFunction = exports.JSONEvent = exports.NodeToJSONNode = exports.ToJSONNode = exports.DefaultJSONDictator = exports.DefaultNodeDictator = void 0;
const dom = require("./dom");
const utils = require("./utils");
const exts = require("./extensions");
const dom_1 = require("./dom");
exports.DefaultNodeDictator = {
    Same: (n, m) => {
        const sameNode = n.nodeType === m.nodeType && n.nodeName === m.nodeName;
        if (!sameNode) {
            return false;
        }
        if (n.nodeType !== dom.ELEMENT_NODE) {
            return true;
        }
        const nElem = n;
        const mElem = m;
        return nElem.id === mElem.id;
    },
    Changed: (n, m) => {
        if (n.nodeType === dom.TEXT_NODE && m.nodeType === dom.TEXT_NODE) {
            return n.textContent === m.textContent;
        }
        if (n.nodeType === dom.COMMENT_NODE && m.nodeType === dom.COMMENT_NODE) {
            return n.textContent === m.textContent;
        }
        const nElem = n;
        const mElem = m;
        const nAttr = dom.recordAttributes(nElem);
        const mAttr = dom.recordAttributes(mElem);
        if (!nAttr.hasOwnProperty('_tid') && mAttr.hasOwnProperty('_tid')) {
            return true;
        }
        if (nAttr.hasOwnProperty('_atid')) {
            return !(mAttr._atid === nAttr._tid && mAttr._tid === mAttr._atid);
        }
        delete mAttr['_atid'];
        delete nAttr['_atid'];
        delete mAttr['_tid'];
        delete nAttr['_tid'];
        for (let key in mAttr) {
            if (!nAttr.hasOwnProperty(key)) {
                return true;
            }
            if (mAttr[key] !== nAttr[key]) {
                return true;
            }
        }
        return nElem.innerHTML !== mElem.innerHTML;
    },
};
exports.DefaultJSONDictator = {
    Same: (n, m) => {
        const sameNode = n.nodeType === m.type && n.nodeName === m.name;
        if (!sameNode) {
            return false;
        }
        if (n.nodeType !== dom.ELEMENT_NODE) {
            return true;
        }
        const nElem = n;
        return nElem.id === m.id;
    },
    Changed: (n, m) => {
        if (n.nodeType === dom.TEXT_NODE && m.type === dom.TEXT_NODE) {
            return n.textContent !== m.content;
        }
        if (n.nodeType === dom.COMMENT_NODE && m.type === dom.COMMENT_NODE) {
            return n.textContent !== m.content;
        }
        const tnode = n;
        if (tnode.hasAttribute('id')) {
            const id = tnode.getAttribute('id');
            if (id !== m.id) {
                return true;
            }
        }
        if (tnode.hasAttribute('_ref')) {
            const ref = tnode.getAttribute('_ref');
            if (ref !== m.ref) {
                return true;
            }
        }
        const tid = tnode.getAttribute('_tid');
        const atid = tnode.getAttribute('_atid');
        if (tnode.hasAttribute('_tid')) {
            if (tid !== m.tid) {
                return true;
            }
            if (tnode.hasAttribute('_atid')) {
                if (tid !== m.tid && atid !== m.atid) {
                    return true;
                }
                if (tid !== m.tid && atid === m.atid) {
                    return true;
                }
            }
        }
        if (!tnode.hasAttribute('events') && m.events.length !== 0) {
            return true;
        }
        for (var index = 0; index < m.events.length; index++) {
            let event = m.events[index];
            let attrName = 'event-' + event.Name;
            let attrValue = event.Targets.join('|');
            let nodeAttr = tnode.attributes.getNamedItem(attrName);
            if (nodeAttr == null) {
                return true;
            }
            if (nodeAttr.value != attrValue) {
                return true;
            }
        }
        return true;
    },
};
function ToJSONNode(node, shallow, parentNode) {
    const list = new Array();
    if (typeof node === 'string') {
        const pub = document.createElement('div');
        pub.innerHTML = node.trim();
        dom.applyChildNode(pub, function (child) {
            list.push(NodeToJSONNode(child, shallow, parentNode));
        });
        return list;
    }
    list.push(NodeToJSONNode(node, shallow, parentNode));
    return list;
}
exports.ToJSONNode = ToJSONNode;
function NodeToJSONNode(node, shallow, parentNode) {
    const jnode = {};
    jnode.children = [];
    jnode.events = [];
    jnode.attrs = [];
    jnode.namespace = '';
    jnode.type = node.nodeType;
    jnode.name = node.nodeName.toLowerCase();
    jnode.id = exts.Objects.GetAttrWith(node, '_id');
    jnode.tid = exts.Objects.GetAttrWith(node, '_tid');
    jnode.ref = exts.Objects.GetAttrWith(node, '_ref');
    jnode.atid = exts.Objects.GetAttrWith(node, '_atid');
    const elem = node;
    if (elem === null)
        return jnode;
    if (node._tid) {
        jnode.tid = node._tid;
    }
    switch (node.nodeType) {
        case dom_1.TEXT_NODE:
            jnode.typeName = 'Text';
            jnode.content = node.textContent;
            return jnode;
        case dom_1.COMMENT_NODE:
            jnode.typeName = 'Comment';
            jnode.content = node.textContent;
            return jnode;
        case dom_1.ELEMENT_NODE:
            jnode.typeName = 'Element';
            jnode.children = new Array();
            break;
        default:
            throw new Error(`unable to handle node type ${node.nodeType}`);
    }
    if (exts.Objects.isNullOrUndefined(elem)) {
        if (jnode.id === '') {
            jnode.id = utils.RandomID();
        }
        return jnode;
    }
    if (elem.hasAttribute('id')) {
        jnode.id = elem.getAttribute('id');
    }
    else {
        jnode.id = utils.RandomID();
    }
    if (jnode.ref === '' && !exts.Objects.isNullOrUndefined(parentNode)) {
        jnode.ref = jnode.id;
    }
    if (elem.hasAttribute('_ref')) {
        jnode.ref = elem.getAttribute('_ref');
    }
    else {
        if (!exts.Objects.isNullOrUndefined(parentNode)) {
            if (parentNode.ref !== '') {
                jnode.ref = parentNode.ref + '/' + jnode.id;
            }
            else {
                jnode.ref = parentNode.id + '/' + jnode.id;
            }
        }
    }
    if (elem.hasAttribute('_tid')) {
        jnode.tid = elem.getAttribute('_tid');
    }
    if (elem.hasAttribute('_atid')) {
        jnode.atid = elem.getAttribute('_atid');
    }
    for (var i = 0; i < elem.attributes.length; i++) {
        let attr = elem.attributes.item(i);
        if (attr == null)
            continue;
        if (!attr.name.startsWith('event-')) {
            continue;
        }
        let eventName = attr.name.replace('event-', '');
        jnode.events.push(JSONEvent(eventName, attr.value.split('|')));
    }
    if (!shallow) {
        dom.applyChildNode(node, function (child) {
            if (child instanceof Text || child instanceof Comment) {
                jnode.children.push(NodeToJSONNode(child, false, jnode));
                return;
            }
            const childElem = child;
            if (!childElem.hasAttribute('id')) {
                childElem.id = utils.RandomID();
            }
            jnode.children.push(NodeToJSONNode(childElem, false, jnode));
        });
    }
    return jnode;
}
exports.NodeToJSONNode = NodeToJSONNode;
function JSONEvent(name, targets) {
    const event = {};
    event.Name = name;
    event.Targets = targets;
    return event;
}
exports.JSONEvent = JSONEvent;
function applyJSONNodeFunction(node, fn) {
    fn(node);
    node.children.forEach(function (child) {
        applyJSONNodeFunction(child, fn);
    });
}
exports.applyJSONNodeFunction = applyJSONNodeFunction;
function applyJSONNodeKidsFunction(node, fn) {
    node.children.forEach(function (child) {
        applyJSONNodeFunction(child, fn);
    });
    fn(node);
}
exports.applyJSONNodeKidsFunction = applyJSONNodeKidsFunction;
function isJSONNode(n) {
    const hasID = typeof n.id !== 'undefined';
    const hasRef = typeof n.ref !== 'undefined';
    const hasTid = typeof n.tid !== 'undefined';
    const hasTypeName = typeof n.typeName !== 'undefined';
    return hasID && hasRef && hasTypeName && hasTid;
}
exports.isJSONNode = isJSONNode;
function findElement(desc, parent) {
    const selector = desc.name + '#' + desc.id;
    const targets = parent.querySelectorAll(selector);
    if (targets.length === 0) {
        let attrSelector = desc.name + `[_tid='${desc.tid}']`;
        let target = parent.querySelector(attrSelector);
        if (target) {
            return target;
        }
        attrSelector = desc.name + `[_atid='${desc.atid}']`;
        target = parent.querySelector(attrSelector);
        if (target) {
            return target;
        }
        attrSelector = desc.name + `[_ref='${desc.ref}']`;
        return parent.querySelector(attrSelector);
    }
    if (targets.length === 1) {
        return targets[0];
    }
    const total = targets.length;
    for (let i = 0; i < total; i++) {
        const elem = targets.item(i);
        if (elem.getAttribute('_tid') === desc.tid) {
            return elem;
        }
        if (elem.getAttribute('_atid') === desc.atid) {
            return elem;
        }
        if (elem.getAttribute('_ref') === desc.ref) {
            return elem;
        }
    }
    return null;
}
exports.findElement = findElement;
function findElementbyRef(ref, parent) {
    const ids = ref.split('/').map(function (elem) {
        if (elem.trim() === '') {
            return '';
        }
        return '#' + elem;
    });
    if (ids.length === 0) {
        return null;
    }
    if (ids[0] === '' || ids[0].trim() === '') {
        ids.shift();
    }
    const first = ids[0];
    if (parent.id == first.substr(1)) {
        ids.shift();
    }
    let cur = parent.querySelector(ids.shift());
    while (cur) {
        if (ids.length === 0) {
            return cur;
        }
        cur = cur.querySelector(ids.shift());
    }
    return cur;
}
exports.findElementbyRef = findElementbyRef;
function findElementParentbyRef(ref, parent) {
    const ids = ref.split('/').map(function (elem) {
        if (elem.trim() === '') {
            return '';
        }
        return '#' + elem;
    });
    if (ids.length === 0) {
        return null;
    }
    if (ids[0] === '' || ids[0].trim() === '') {
        ids.shift();
    }
    ids.pop();
    const first = ids[0];
    if (parent.id == first.substr(1)) {
        ids.shift();
    }
    let cur = parent.querySelector(ids.shift());
    while (cur) {
        if (ids.length === 0) {
            return cur;
        }
        cur = cur.querySelector(ids.shift());
    }
    return cur;
}
exports.findElementParentbyRef = findElementParentbyRef;
exports.DefaultJSONMaker = {
    Make: jsonMaker,
};
function jsonMaker(doc, descNode, shallow, skipRemoved) {
    if (descNode.type === dom_1.COMMENT_NODE) {
        const node = doc.createComment(descNode.content);
        exts.Objects.PatchWith(node, '_id', descNode.id);
        exts.Objects.PatchWith(node, '_ref', descNode.ref);
        exts.Objects.PatchWith(node, '_tid', descNode.tid);
        exts.Objects.PatchWith(node, '_atid', descNode.atid);
        return node;
    }
    if (descNode.type === dom_1.TEXT_NODE) {
        const node = doc.createTextNode(descNode.content);
        exts.Objects.PatchWith(node, '_id', descNode.id);
        exts.Objects.PatchWith(node, '_ref', descNode.ref);
        exts.Objects.PatchWith(node, '_tid', descNode.tid);
        exts.Objects.PatchWith(node, '_atid', descNode.atid);
        return node;
    }
    if (descNode.id === '') {
        descNode.id = utils.RandomID();
    }
    let node;
    if (descNode.namespace.length !== 0) {
        node = doc.createElement(descNode.name);
    }
    else {
        node = doc.createElementNS(descNode.namespace, descNode.name);
    }
    exts.Objects.PatchWith(node, '_id', descNode.id);
    exts.Objects.PatchWith(node, '_ref', descNode.ref);
    exts.Objects.PatchWith(node, '_tid', descNode.tid);
    exts.Objects.PatchWith(node, '_atid', descNode.atid);
    node.setAttribute('id', descNode.id);
    node.setAttribute('_tid', descNode.tid);
    node.setAttribute('_ref', descNode.ref);
    node.setAttribute('_atid', descNode.atid);
    descNode.events.forEach(function events(event) {
        node.setAttribute('event-' + event.Name, event.Targets.join('|'));
    });
    descNode.attrs.forEach(function attrs(attr) {
        node.setAttribute(attr.Key, attr.Value);
    });
    if (descNode.removed) {
        node.setAttribute('_removed', 'true');
        return node;
    }
    if (!shallow) {
        descNode.children.forEach(function (kidJSON) {
            if (skipRemoved && kidJSON.removed) {
                return;
            }
            node.appendChild(jsonMaker(doc, kidJSON, shallow, skipRemoved));
        });
    }
    return node;
}
exports.jsonMaker = jsonMaker;
function PatchJSONNodeTree(fragment, mount, dictator, maker) {
    let targetNode = findElement(fragment, mount);
    if (exts.Objects.isNullOrUndefined(targetNode)) {
        const tNode = maker.Make(document, fragment, false, true);
        mount.appendChild(tNode);
        return;
    }
    PatchJSONNode(fragment, targetNode, dictator, maker);
}
exports.PatchJSONNodeTree = PatchJSONNodeTree;
function PatchJSONNode(fragment, targetNode, dictator, maker) {
    if (!dictator.Same(targetNode, fragment)) {
        const tNode = maker.Make(document, fragment, false, true);
        dom.replaceNode(targetNode.parentNode, targetNode, tNode);
        return;
    }
    if (!dictator.Changed(targetNode, fragment)) {
        return;
    }
    PatchJSONAttributes(fragment, targetNode);
    const kids = dom.nodeListToArray(targetNode.childNodes);
    const totalKids = kids.length;
    const fragmentKids = fragment.children.length;
    let i = 0;
    for (; i < totalKids; i++) {
        const childNode = kids[i];
        if (i >= fragmentKids) {
            const chnode = childNode;
            if (chnode) {
                chnode.remove();
            }
            continue;
        }
        const childFragment = fragment.children[i];
        PatchJSONNode(childFragment, childNode, dictator, maker);
    }
    for (; i < fragmentKids; i++) {
        const tNode = maker.Make(document, fragment, false, true);
        targetNode.appendChild(tNode);
    }
    return;
}
exports.PatchJSONNode = PatchJSONNode;
function StreamJSONNodes(fragment, mount, dictator, maker) {
    const changes = fragment.filter(function (elem) {
        return !elem.removed;
    });
    fragment
        .filter(function (elem) {
        if (!elem.removed) {
            return false;
        }
        let filtered = true;
        changes.forEach(function (el) {
            if (elem.tid === el.tid || elem.tid == el.atid || elem.ref === el.ref) {
                filtered = false;
            }
        });
        return filtered;
    })
        .forEach(function (removal) {
        const target = findElement(removal, mount);
        if (target) {
            target.remove();
        }
    });
    changes.forEach(function (change) {
        const targetNode = findElement(change, mount);
        if (exts.Objects.isNullOrUndefined(targetNode)) {
            const targetNodeParent = findElementParentbyRef(change.ref, mount);
            if (exts.Objects.isNullOrUndefined(targetNodeParent)) {
                console.log('Unable to apply new change stream: ', change);
                return;
            }
            const tNode = maker.Make(document, change, false, true);
            targetNodeParent.appendChild(tNode);
            return;
        }
        ApplyStreamNode(change, targetNode, dictator, maker);
    });
    return;
}
exports.StreamJSONNodes = StreamJSONNodes;
function ApplyStreamNode(fragment, targetNode, dictator, maker) {
    if (!dictator.Same(targetNode, fragment)) {
        const tNode = maker.Make(document, fragment, false, true);
        dom.replaceNode(targetNode.parentNode, targetNode, tNode);
        return;
    }
    if (dictator.Changed(targetNode, fragment)) {
        PatchJSONAttributes(fragment, targetNode);
    }
    const totalKids = targetNode.childNodes.length;
    const fragmentKids = fragment.children.length;
    let i = 0;
    for (; i < totalKids; i++) {
        const childNode = targetNode.childNodes[i];
        if (i >= fragmentKids) {
            return;
        }
        const childFragment = fragment.children[i];
        PatchJSONNode(childFragment, childNode, dictator, maker);
    }
    for (; i < fragmentKids; i++) {
        const tNode = maker.Make(document, fragment, false, true);
        targetNode.appendChild(tNode);
    }
    return;
}
exports.ApplyStreamNode = ApplyStreamNode;
function PatchTextCommentWithJSON(fragment, target) {
    if (fragment.type !== dom_1.COMMENT_NODE && fragment.type !== dom_1.TEXT_NODE) {
        return;
    }
    if (fragment.type !== dom_1.COMMENT_NODE && fragment.type !== dom_1.TEXT_NODE) {
        return;
    }
    if (target.textContent === fragment.content) {
        return;
    }
    target.textContent = fragment.content;
    exts.Objects.PatchWith(target, '_ref', fragment.ref);
    exts.Objects.PatchWith(target, '_tid', fragment.tid);
    exts.Objects.PatchWith(target, '_atid', fragment.atid);
}
exports.PatchTextCommentWithJSON = PatchTextCommentWithJSON;
function PatchJSONAttributes(node, target) {
    const oldNodeAttrs = dom.recordAttributes(target);
    node.attrs.forEach(function (attr) {
        const oldValue = oldNodeAttrs[attr.Key];
        delete oldNodeAttrs[attr.Key];
        if (attr.Value === oldValue) {
            return null;
        }
        target.setAttribute(attr.Key, attr.Value);
    });
    for (let index in oldNodeAttrs) {
        target.removeAttribute(index);
    }
    target.setAttribute('_tid', node.tid);
    target.setAttribute('_ref', node.ref);
    target.setAttribute('_atid', node.atid);
    node.events.forEach(function events(event) {
        target.setAttribute('event-' + event.Name, event.Targets.join('|'));
    });
    exts.Objects.PatchWith(target, '_id', node.id);
    exts.Objects.PatchWith(target, '_ref', node.ref);
    exts.Objects.PatchWith(target, '_tid', node.tid);
    exts.Objects.PatchWith(target, '_atid', node.atid);
}
exports.PatchJSONAttributes = PatchJSONAttributes;
function PatchDOMTree(newFragment, oldNodeOrMount, dictator, isChildRecursion) {
    if (isChildRecursion) {
        const rootNode = oldNodeOrMount.parentNode;
        if (!dictator.Same(oldNodeOrMount, newFragment)) {
            dom.replaceNode(rootNode, oldNodeOrMount, newFragment);
            return null;
        }
        if (!oldNodeOrMount.hasChildNodes()) {
            dom.replaceNode(rootNode, oldNodeOrMount, newFragment);
            return null;
        }
    }
    const newChildren = dom.nodeListToArray(newFragment.childNodes);
    const oldChildren = dom.nodeListToArray(oldNodeOrMount.childNodes);
    const oldChildrenLength = oldChildren.length;
    const newChildrenLength = newChildren.length;
    const removeOldLeft = newChildrenLength < oldChildrenLength;
    let lastIndex = 0;
    let lastNode;
    let newChildNode;
    let lastNodeNextSibling = null;
    for (; lastIndex < oldChildrenLength; lastIndex++) {
        if (lastIndex >= newChildrenLength) {
            break;
        }
        lastNode = oldChildren[lastIndex];
        newChildNode = newChildren[lastIndex];
        lastNodeNextSibling = lastNode.nextSibling;
        if ((lastNode.nodeType === dom.TEXT_NODE || lastNode.nodeType === dom.COMMENT_NODE) &&
            newChildNode.nodeType === lastNode.nodeType) {
            if (lastNode.textContent !== newChildNode.textContent) {
                lastNode.textContent = newChildNode.textContent;
            }
            continue;
        }
        if (!dictator.Same(lastNode, newChildNode)) {
            dom.replaceNode(oldNodeOrMount, lastNode, newChildNode);
            continue;
        }
        if (!dictator.Changed(lastNode, newChildNode)) {
            continue;
        }
        if (lastNode.nodeType !== newChildNode.nodeType) {
            dom.replaceNode(oldNodeOrMount, lastNode, newChildNode);
            continue;
        }
        if (!lastNode.hasChildNodes() && newChildNode.hasChildNodes()) {
            dom.replaceNode(oldNodeOrMount, lastNode, newChildNode);
            continue;
        }
        if (lastNode.hasChildNodes() && !newChildNode.hasChildNodes()) {
            dom.replaceNode(oldNodeOrMount, lastNode, newChildNode);
            continue;
        }
        const lastElement = lastNode;
        const newElement = newChildNode;
        PatchDOMAttributes(newElement, lastElement);
        lastElement.setAttribute('_patched', 'true');
        PatchDOMTree(newElement, lastElement, dictator, true);
        lastElement.removeAttribute('_patched');
    }
    if (removeOldLeft && lastNodeNextSibling !== null) {
        dom.removeFromNode(lastNodeNextSibling, null);
        return null;
    }
    for (; lastIndex < newChildrenLength; lastIndex++) {
        let newNode = newChildren[lastIndex];
        if (!exts.Objects.isNullOrUndefined(newNode)) {
            oldNodeOrMount.appendChild(newNode);
        }
    }
}
exports.PatchDOMTree = PatchDOMTree;
function PatchDOMAttributes(newElement, oldElement) {
    const oldNodeAttrs = dom.recordAttributes(oldElement);
    for (let index in newElement.attributes) {
        const attr = newElement.attributes[index];
        const oldValue = oldNodeAttrs[attr.name];
        delete oldNodeAttrs[attr.name];
        if (attr.value === oldValue) {
            continue;
        }
        oldElement.setAttribute(attr.name, attr.value);
    }
    for (let index in oldNodeAttrs) {
        oldElement.removeAttribute(index);
    }
}
exports.PatchDOMAttributes = PatchDOMAttributes;

},{"./dom":12,"./extensions":13,"./utils":17}],17:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isEqual = exports.RandomID = exports.truncateArray = exports.createMap = exports.has = exports.Blank = exports.ToKebabCase = void 0;
const hasOwnProperty = Object.prototype.hasOwnProperty;
function ToKebabCase(str) {
    const result = str.replace(/[A-Z\u00C0-\u00D6\u00D8-\u00DE]/g, (match) => '-' + match.toLowerCase());
    return str[0] === str[0].toUpperCase() ? result.substring(1) : result;
}
exports.ToKebabCase = ToKebabCase;
function Blank() { }
exports.Blank = Blank;
Blank.prototype = Object.create(null);
function has(map, property) {
    return hasOwnProperty.call(map, property);
}
exports.has = has;
function createMap() {
    return new Blank();
}
exports.createMap = createMap;
function truncateArray(arr, length) {
    while (arr.length > length) {
        arr.pop();
    }
}
exports.truncateArray = truncateArray;
function RandomID() {
    return Math.random().toString(36).substr(2, 9);
}
exports.RandomID = RandomID;
function isEqual(a, b) {
    const aProps = Object.getOwnPropertyNames(a);
    const bProps = Object.getOwnPropertyNames(b);
    if (aProps.length != bProps.length) {
        return false;
    }
    for (let i = 0; i < aProps.length; i++) {
        const propName = aProps[i];
        if (a[propName] !== b[propName]) {
            return false;
        }
    }
    return true;
}
exports.isEqual = isEqual;

},{}],18:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.mountTo = void 0;
const promise = require("promise-polyfill");
if (!self.Promise) {
    self.Promise = promise;
}
const namespace = self.Promise;
function mountTo(parent) {
    parent.promises = namespace;
}
exports.mountTo = mountTo;
exports.default = namespace;

},{"promise-polyfill":2}],19:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.mountTo = void 0;
const namespace = {};
function mountTo(parent) {
    parent.realm = namespace;
}
exports.mountTo = mountTo;
exports.default = namespace;

},{}],20:[function(require,module,exports){
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const realm = require("../../realm/src");
const markup = require("../../markup/src");
const animations = require("../../animation/src");
const http = require("../../http/src");
const promises = require("../../promises/src");
const voidgram = {
    realm: {},
    animations: {},
    http: {},
    promises: {},
    markup: {},
};
realm.mountTo(voidgram);
markup.mountTo(voidgram);
animations.mountTo(voidgram);
http.mountTo(voidgram);
promises.mountTo(voidgram);
if (window) {
    window.voidgram = voidgram;
}
exports.default = voidgram;

},{"../../animation/src":6,"../../http/src":10,"../../markup/src":14,"../../promises/src":18,"../../realm/src":19}]},{},[20])
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIm5vZGVfbW9kdWxlcy9icm93c2VyLXBhY2svX3ByZWx1ZGUuanMiLCJub2RlX21vZHVsZXMvcHJvY2Vzcy9icm93c2VyLmpzIiwibm9kZV9tb2R1bGVzL3Byb21pc2UtcG9seWZpbGwvbGliL2luZGV4LmpzIiwibm9kZV9tb2R1bGVzL3RpbWVycy1icm93c2VyaWZ5L21haW4uanMiLCJub2RlX21vZHVsZXMvd2hhdHdnLWZldGNoL2Rpc3QvZmV0Y2gudW1kLmpzIiwicGFja2FnZXMvdm9pZGdyYW0vbGliL2FuaW1hdGlvbi9zcmMvYW5pbWUuanMiLCJwYWNrYWdlcy92b2lkZ3JhbS9saWIvYW5pbWF0aW9uL3NyYy9pbmRleC5qcyIsInBhY2thZ2VzL3ZvaWRncmFtL2xpYi9hbmltYXRpb24vc3JjL3JhZi1wb2x5ZmlsbC5qcyIsInBhY2thZ2VzL3ZvaWRncmFtL2xpYi9odHRwL3NyYy9mZXRjaC5qcyIsInBhY2thZ2VzL3ZvaWRncmFtL2xpYi9odHRwL3NyYy9odHRwLmpzIiwicGFja2FnZXMvdm9pZGdyYW0vbGliL2h0dHAvc3JjL2luZGV4LmpzIiwicGFja2FnZXMvdm9pZGdyYW0vbGliL2h0dHAvc3JjL3dlYnNvY2tldC5qcyIsInBhY2thZ2VzL3ZvaWRncmFtL2xpYi9tYXJrdXAvc3JjL2RvbS5qcyIsInBhY2thZ2VzL3ZvaWRncmFtL2xpYi9tYXJrdXAvc3JjL2V4dGVuc2lvbnMuanMiLCJwYWNrYWdlcy92b2lkZ3JhbS9saWIvbWFya3VwL3NyYy9pbmRleC5qcyIsInBhY2thZ2VzL3ZvaWRncmFtL2xpYi9tYXJrdXAvc3JjL21vdW50LmpzIiwicGFja2FnZXMvdm9pZGdyYW0vbGliL21hcmt1cC9zcmMvcGF0Y2guanMiLCJwYWNrYWdlcy92b2lkZ3JhbS9saWIvbWFya3VwL3NyYy91dGlscy5qcyIsInBhY2thZ2VzL3ZvaWRncmFtL2xpYi9wcm9taXNlcy9zcmMvaW5kZXguanMiLCJwYWNrYWdlcy92b2lkZ3JhbS9saWIvcmVhbG0vc3JjL2luZGV4LmpzIiwicGFja2FnZXMvdm9pZGdyYW0vbGliL3ZvaWRncmFtL3NyYy9pbmRleC5qcyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiQUFBQTtBQ0FBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7OztBQ3hMQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOzs7OztBQ25RQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTs7OztBQzNFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTs7QUNuaEJBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FDak1BO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTs7QUNYQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FDbERBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FDaEJBO0FBQ0E7QUFDQTtBQUNBOztBQ0hBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQ1pBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQ3BGQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTs7QUNyckJBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FDNUJBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTs7QUNkQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FDOUpBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTs7QUNwb0JBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQzdDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQ2JBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQ1RBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQSIsImZpbGUiOiJnZW5lcmF0ZWQuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlc0NvbnRlbnQiOlsiKGZ1bmN0aW9uKCl7ZnVuY3Rpb24gcihlLG4sdCl7ZnVuY3Rpb24gbyhpLGYpe2lmKCFuW2ldKXtpZighZVtpXSl7dmFyIGM9XCJmdW5jdGlvblwiPT10eXBlb2YgcmVxdWlyZSYmcmVxdWlyZTtpZighZiYmYylyZXR1cm4gYyhpLCEwKTtpZih1KXJldHVybiB1KGksITApO3ZhciBhPW5ldyBFcnJvcihcIkNhbm5vdCBmaW5kIG1vZHVsZSAnXCIraStcIidcIik7dGhyb3cgYS5jb2RlPVwiTU9EVUxFX05PVF9GT1VORFwiLGF9dmFyIHA9bltpXT17ZXhwb3J0czp7fX07ZVtpXVswXS5jYWxsKHAuZXhwb3J0cyxmdW5jdGlvbihyKXt2YXIgbj1lW2ldWzFdW3JdO3JldHVybiBvKG58fHIpfSxwLHAuZXhwb3J0cyxyLGUsbix0KX1yZXR1cm4gbltpXS5leHBvcnRzfWZvcih2YXIgdT1cImZ1bmN0aW9uXCI9PXR5cGVvZiByZXF1aXJlJiZyZXF1aXJlLGk9MDtpPHQubGVuZ3RoO2krKylvKHRbaV0pO3JldHVybiBvfXJldHVybiByfSkoKSIsIi8vIHNoaW0gZm9yIHVzaW5nIHByb2Nlc3MgaW4gYnJvd3NlclxudmFyIHByb2Nlc3MgPSBtb2R1bGUuZXhwb3J0cyA9IHt9O1xuXG4vLyBjYWNoZWQgZnJvbSB3aGF0ZXZlciBnbG9iYWwgaXMgcHJlc2VudCBzbyB0aGF0IHRlc3QgcnVubmVycyB0aGF0IHN0dWIgaXRcbi8vIGRvbid0IGJyZWFrIHRoaW5ncy4gIEJ1dCB3ZSBuZWVkIHRvIHdyYXAgaXQgaW4gYSB0cnkgY2F0Y2ggaW4gY2FzZSBpdCBpc1xuLy8gd3JhcHBlZCBpbiBzdHJpY3QgbW9kZSBjb2RlIHdoaWNoIGRvZXNuJ3QgZGVmaW5lIGFueSBnbG9iYWxzLiAgSXQncyBpbnNpZGUgYVxuLy8gZnVuY3Rpb24gYmVjYXVzZSB0cnkvY2F0Y2hlcyBkZW9wdGltaXplIGluIGNlcnRhaW4gZW5naW5lcy5cblxudmFyIGNhY2hlZFNldFRpbWVvdXQ7XG52YXIgY2FjaGVkQ2xlYXJUaW1lb3V0O1xuXG5mdW5jdGlvbiBkZWZhdWx0U2V0VGltb3V0KCkge1xuICAgIHRocm93IG5ldyBFcnJvcignc2V0VGltZW91dCBoYXMgbm90IGJlZW4gZGVmaW5lZCcpO1xufVxuZnVuY3Rpb24gZGVmYXVsdENsZWFyVGltZW91dCAoKSB7XG4gICAgdGhyb3cgbmV3IEVycm9yKCdjbGVhclRpbWVvdXQgaGFzIG5vdCBiZWVuIGRlZmluZWQnKTtcbn1cbihmdW5jdGlvbiAoKSB7XG4gICAgdHJ5IHtcbiAgICAgICAgaWYgKHR5cGVvZiBzZXRUaW1lb3V0ID09PSAnZnVuY3Rpb24nKSB7XG4gICAgICAgICAgICBjYWNoZWRTZXRUaW1lb3V0ID0gc2V0VGltZW91dDtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAgIGNhY2hlZFNldFRpbWVvdXQgPSBkZWZhdWx0U2V0VGltb3V0O1xuICAgICAgICB9XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgICBjYWNoZWRTZXRUaW1lb3V0ID0gZGVmYXVsdFNldFRpbW91dDtcbiAgICB9XG4gICAgdHJ5IHtcbiAgICAgICAgaWYgKHR5cGVvZiBjbGVhclRpbWVvdXQgPT09ICdmdW5jdGlvbicpIHtcbiAgICAgICAgICAgIGNhY2hlZENsZWFyVGltZW91dCA9IGNsZWFyVGltZW91dDtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAgIGNhY2hlZENsZWFyVGltZW91dCA9IGRlZmF1bHRDbGVhclRpbWVvdXQ7XG4gICAgICAgIH1cbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICAgIGNhY2hlZENsZWFyVGltZW91dCA9IGRlZmF1bHRDbGVhclRpbWVvdXQ7XG4gICAgfVxufSAoKSlcbmZ1bmN0aW9uIHJ1blRpbWVvdXQoZnVuKSB7XG4gICAgaWYgKGNhY2hlZFNldFRpbWVvdXQgPT09IHNldFRpbWVvdXQpIHtcbiAgICAgICAgLy9ub3JtYWwgZW52aXJvbWVudHMgaW4gc2FuZSBzaXR1YXRpb25zXG4gICAgICAgIHJldHVybiBzZXRUaW1lb3V0KGZ1biwgMCk7XG4gICAgfVxuICAgIC8vIGlmIHNldFRpbWVvdXQgd2Fzbid0IGF2YWlsYWJsZSBidXQgd2FzIGxhdHRlciBkZWZpbmVkXG4gICAgaWYgKChjYWNoZWRTZXRUaW1lb3V0ID09PSBkZWZhdWx0U2V0VGltb3V0IHx8ICFjYWNoZWRTZXRUaW1lb3V0KSAmJiBzZXRUaW1lb3V0KSB7XG4gICAgICAgIGNhY2hlZFNldFRpbWVvdXQgPSBzZXRUaW1lb3V0O1xuICAgICAgICByZXR1cm4gc2V0VGltZW91dChmdW4sIDApO1xuICAgIH1cbiAgICB0cnkge1xuICAgICAgICAvLyB3aGVuIHdoZW4gc29tZWJvZHkgaGFzIHNjcmV3ZWQgd2l0aCBzZXRUaW1lb3V0IGJ1dCBubyBJLkUuIG1hZGRuZXNzXG4gICAgICAgIHJldHVybiBjYWNoZWRTZXRUaW1lb3V0KGZ1biwgMCk7XG4gICAgfSBjYXRjaChlKXtcbiAgICAgICAgdHJ5IHtcbiAgICAgICAgICAgIC8vIFdoZW4gd2UgYXJlIGluIEkuRS4gYnV0IHRoZSBzY3JpcHQgaGFzIGJlZW4gZXZhbGVkIHNvIEkuRS4gZG9lc24ndCB0cnVzdCB0aGUgZ2xvYmFsIG9iamVjdCB3aGVuIGNhbGxlZCBub3JtYWxseVxuICAgICAgICAgICAgcmV0dXJuIGNhY2hlZFNldFRpbWVvdXQuY2FsbChudWxsLCBmdW4sIDApO1xuICAgICAgICB9IGNhdGNoKGUpe1xuICAgICAgICAgICAgLy8gc2FtZSBhcyBhYm92ZSBidXQgd2hlbiBpdCdzIGEgdmVyc2lvbiBvZiBJLkUuIHRoYXQgbXVzdCBoYXZlIHRoZSBnbG9iYWwgb2JqZWN0IGZvciAndGhpcycsIGhvcGZ1bGx5IG91ciBjb250ZXh0IGNvcnJlY3Qgb3RoZXJ3aXNlIGl0IHdpbGwgdGhyb3cgYSBnbG9iYWwgZXJyb3JcbiAgICAgICAgICAgIHJldHVybiBjYWNoZWRTZXRUaW1lb3V0LmNhbGwodGhpcywgZnVuLCAwKTtcbiAgICAgICAgfVxuICAgIH1cblxuXG59XG5mdW5jdGlvbiBydW5DbGVhclRpbWVvdXQobWFya2VyKSB7XG4gICAgaWYgKGNhY2hlZENsZWFyVGltZW91dCA9PT0gY2xlYXJUaW1lb3V0KSB7XG4gICAgICAgIC8vbm9ybWFsIGVudmlyb21lbnRzIGluIHNhbmUgc2l0dWF0aW9uc1xuICAgICAgICByZXR1cm4gY2xlYXJUaW1lb3V0KG1hcmtlcik7XG4gICAgfVxuICAgIC8vIGlmIGNsZWFyVGltZW91dCB3YXNuJ3QgYXZhaWxhYmxlIGJ1dCB3YXMgbGF0dGVyIGRlZmluZWRcbiAgICBpZiAoKGNhY2hlZENsZWFyVGltZW91dCA9PT0gZGVmYXVsdENsZWFyVGltZW91dCB8fCAhY2FjaGVkQ2xlYXJUaW1lb3V0KSAmJiBjbGVhclRpbWVvdXQpIHtcbiAgICAgICAgY2FjaGVkQ2xlYXJUaW1lb3V0ID0gY2xlYXJUaW1lb3V0O1xuICAgICAgICByZXR1cm4gY2xlYXJUaW1lb3V0KG1hcmtlcik7XG4gICAgfVxuICAgIHRyeSB7XG4gICAgICAgIC8vIHdoZW4gd2hlbiBzb21lYm9keSBoYXMgc2NyZXdlZCB3aXRoIHNldFRpbWVvdXQgYnV0IG5vIEkuRS4gbWFkZG5lc3NcbiAgICAgICAgcmV0dXJuIGNhY2hlZENsZWFyVGltZW91dChtYXJrZXIpO1xuICAgIH0gY2F0Y2ggKGUpe1xuICAgICAgICB0cnkge1xuICAgICAgICAgICAgLy8gV2hlbiB3ZSBhcmUgaW4gSS5FLiBidXQgdGhlIHNjcmlwdCBoYXMgYmVlbiBldmFsZWQgc28gSS5FLiBkb2Vzbid0ICB0cnVzdCB0aGUgZ2xvYmFsIG9iamVjdCB3aGVuIGNhbGxlZCBub3JtYWxseVxuICAgICAgICAgICAgcmV0dXJuIGNhY2hlZENsZWFyVGltZW91dC5jYWxsKG51bGwsIG1hcmtlcik7XG4gICAgICAgIH0gY2F0Y2ggKGUpe1xuICAgICAgICAgICAgLy8gc2FtZSBhcyBhYm92ZSBidXQgd2hlbiBpdCdzIGEgdmVyc2lvbiBvZiBJLkUuIHRoYXQgbXVzdCBoYXZlIHRoZSBnbG9iYWwgb2JqZWN0IGZvciAndGhpcycsIGhvcGZ1bGx5IG91ciBjb250ZXh0IGNvcnJlY3Qgb3RoZXJ3aXNlIGl0IHdpbGwgdGhyb3cgYSBnbG9iYWwgZXJyb3IuXG4gICAgICAgICAgICAvLyBTb21lIHZlcnNpb25zIG9mIEkuRS4gaGF2ZSBkaWZmZXJlbnQgcnVsZXMgZm9yIGNsZWFyVGltZW91dCB2cyBzZXRUaW1lb3V0XG4gICAgICAgICAgICByZXR1cm4gY2FjaGVkQ2xlYXJUaW1lb3V0LmNhbGwodGhpcywgbWFya2VyKTtcbiAgICAgICAgfVxuICAgIH1cblxuXG5cbn1cbnZhciBxdWV1ZSA9IFtdO1xudmFyIGRyYWluaW5nID0gZmFsc2U7XG52YXIgY3VycmVudFF1ZXVlO1xudmFyIHF1ZXVlSW5kZXggPSAtMTtcblxuZnVuY3Rpb24gY2xlYW5VcE5leHRUaWNrKCkge1xuICAgIGlmICghZHJhaW5pbmcgfHwgIWN1cnJlbnRRdWV1ZSkge1xuICAgICAgICByZXR1cm47XG4gICAgfVxuICAgIGRyYWluaW5nID0gZmFsc2U7XG4gICAgaWYgKGN1cnJlbnRRdWV1ZS5sZW5ndGgpIHtcbiAgICAgICAgcXVldWUgPSBjdXJyZW50UXVldWUuY29uY2F0KHF1ZXVlKTtcbiAgICB9IGVsc2Uge1xuICAgICAgICBxdWV1ZUluZGV4ID0gLTE7XG4gICAgfVxuICAgIGlmIChxdWV1ZS5sZW5ndGgpIHtcbiAgICAgICAgZHJhaW5RdWV1ZSgpO1xuICAgIH1cbn1cblxuZnVuY3Rpb24gZHJhaW5RdWV1ZSgpIHtcbiAgICBpZiAoZHJhaW5pbmcpIHtcbiAgICAgICAgcmV0dXJuO1xuICAgIH1cbiAgICB2YXIgdGltZW91dCA9IHJ1blRpbWVvdXQoY2xlYW5VcE5leHRUaWNrKTtcbiAgICBkcmFpbmluZyA9IHRydWU7XG5cbiAgICB2YXIgbGVuID0gcXVldWUubGVuZ3RoO1xuICAgIHdoaWxlKGxlbikge1xuICAgICAgICBjdXJyZW50UXVldWUgPSBxdWV1ZTtcbiAgICAgICAgcXVldWUgPSBbXTtcbiAgICAgICAgd2hpbGUgKCsrcXVldWVJbmRleCA8IGxlbikge1xuICAgICAgICAgICAgaWYgKGN1cnJlbnRRdWV1ZSkge1xuICAgICAgICAgICAgICAgIGN1cnJlbnRRdWV1ZVtxdWV1ZUluZGV4XS5ydW4oKTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICBxdWV1ZUluZGV4ID0gLTE7XG4gICAgICAgIGxlbiA9IHF1ZXVlLmxlbmd0aDtcbiAgICB9XG4gICAgY3VycmVudFF1ZXVlID0gbnVsbDtcbiAgICBkcmFpbmluZyA9IGZhbHNlO1xuICAgIHJ1bkNsZWFyVGltZW91dCh0aW1lb3V0KTtcbn1cblxucHJvY2Vzcy5uZXh0VGljayA9IGZ1bmN0aW9uIChmdW4pIHtcbiAgICB2YXIgYXJncyA9IG5ldyBBcnJheShhcmd1bWVudHMubGVuZ3RoIC0gMSk7XG4gICAgaWYgKGFyZ3VtZW50cy5sZW5ndGggPiAxKSB7XG4gICAgICAgIGZvciAodmFyIGkgPSAxOyBpIDwgYXJndW1lbnRzLmxlbmd0aDsgaSsrKSB7XG4gICAgICAgICAgICBhcmdzW2kgLSAxXSA9IGFyZ3VtZW50c1tpXTtcbiAgICAgICAgfVxuICAgIH1cbiAgICBxdWV1ZS5wdXNoKG5ldyBJdGVtKGZ1biwgYXJncykpO1xuICAgIGlmIChxdWV1ZS5sZW5ndGggPT09IDEgJiYgIWRyYWluaW5nKSB7XG4gICAgICAgIHJ1blRpbWVvdXQoZHJhaW5RdWV1ZSk7XG4gICAgfVxufTtcblxuLy8gdjggbGlrZXMgcHJlZGljdGlibGUgb2JqZWN0c1xuZnVuY3Rpb24gSXRlbShmdW4sIGFycmF5KSB7XG4gICAgdGhpcy5mdW4gPSBmdW47XG4gICAgdGhpcy5hcnJheSA9IGFycmF5O1xufVxuSXRlbS5wcm90b3R5cGUucnVuID0gZnVuY3Rpb24gKCkge1xuICAgIHRoaXMuZnVuLmFwcGx5KG51bGwsIHRoaXMuYXJyYXkpO1xufTtcbnByb2Nlc3MudGl0bGUgPSAnYnJvd3Nlcic7XG5wcm9jZXNzLmJyb3dzZXIgPSB0cnVlO1xucHJvY2Vzcy5lbnYgPSB7fTtcbnByb2Nlc3MuYXJndiA9IFtdO1xucHJvY2Vzcy52ZXJzaW9uID0gJyc7IC8vIGVtcHR5IHN0cmluZyB0byBhdm9pZCByZWdleHAgaXNzdWVzXG5wcm9jZXNzLnZlcnNpb25zID0ge307XG5cbmZ1bmN0aW9uIG5vb3AoKSB7fVxuXG5wcm9jZXNzLm9uID0gbm9vcDtcbnByb2Nlc3MuYWRkTGlzdGVuZXIgPSBub29wO1xucHJvY2Vzcy5vbmNlID0gbm9vcDtcbnByb2Nlc3Mub2ZmID0gbm9vcDtcbnByb2Nlc3MucmVtb3ZlTGlzdGVuZXIgPSBub29wO1xucHJvY2Vzcy5yZW1vdmVBbGxMaXN0ZW5lcnMgPSBub29wO1xucHJvY2Vzcy5lbWl0ID0gbm9vcDtcbnByb2Nlc3MucHJlcGVuZExpc3RlbmVyID0gbm9vcDtcbnByb2Nlc3MucHJlcGVuZE9uY2VMaXN0ZW5lciA9IG5vb3A7XG5cbnByb2Nlc3MubGlzdGVuZXJzID0gZnVuY3Rpb24gKG5hbWUpIHsgcmV0dXJuIFtdIH1cblxucHJvY2Vzcy5iaW5kaW5nID0gZnVuY3Rpb24gKG5hbWUpIHtcbiAgICB0aHJvdyBuZXcgRXJyb3IoJ3Byb2Nlc3MuYmluZGluZyBpcyBub3Qgc3VwcG9ydGVkJyk7XG59O1xuXG5wcm9jZXNzLmN3ZCA9IGZ1bmN0aW9uICgpIHsgcmV0dXJuICcvJyB9O1xucHJvY2Vzcy5jaGRpciA9IGZ1bmN0aW9uIChkaXIpIHtcbiAgICB0aHJvdyBuZXcgRXJyb3IoJ3Byb2Nlc3MuY2hkaXIgaXMgbm90IHN1cHBvcnRlZCcpO1xufTtcbnByb2Nlc3MudW1hc2sgPSBmdW5jdGlvbigpIHsgcmV0dXJuIDA7IH07XG4iLCIndXNlIHN0cmljdCc7XG5cbi8qKlxuICogQHRoaXMge1Byb21pc2V9XG4gKi9cbmZ1bmN0aW9uIGZpbmFsbHlDb25zdHJ1Y3RvcihjYWxsYmFjaykge1xuICB2YXIgY29uc3RydWN0b3IgPSB0aGlzLmNvbnN0cnVjdG9yO1xuICByZXR1cm4gdGhpcy50aGVuKFxuICAgIGZ1bmN0aW9uKHZhbHVlKSB7XG4gICAgICByZXR1cm4gY29uc3RydWN0b3IucmVzb2x2ZShjYWxsYmFjaygpKS50aGVuKGZ1bmN0aW9uKCkge1xuICAgICAgICByZXR1cm4gdmFsdWU7XG4gICAgICB9KTtcbiAgICB9LFxuICAgIGZ1bmN0aW9uKHJlYXNvbikge1xuICAgICAgcmV0dXJuIGNvbnN0cnVjdG9yLnJlc29sdmUoY2FsbGJhY2soKSkudGhlbihmdW5jdGlvbigpIHtcbiAgICAgICAgcmV0dXJuIGNvbnN0cnVjdG9yLnJlamVjdChyZWFzb24pO1xuICAgICAgfSk7XG4gICAgfVxuICApO1xufVxuXG4vLyBTdG9yZSBzZXRUaW1lb3V0IHJlZmVyZW5jZSBzbyBwcm9taXNlLXBvbHlmaWxsIHdpbGwgYmUgdW5hZmZlY3RlZCBieVxuLy8gb3RoZXIgY29kZSBtb2RpZnlpbmcgc2V0VGltZW91dCAobGlrZSBzaW5vbi51c2VGYWtlVGltZXJzKCkpXG52YXIgc2V0VGltZW91dEZ1bmMgPSBzZXRUaW1lb3V0O1xuXG5mdW5jdGlvbiBub29wKCkge31cblxuLy8gUG9seWZpbGwgZm9yIEZ1bmN0aW9uLnByb3RvdHlwZS5iaW5kXG5mdW5jdGlvbiBiaW5kKGZuLCB0aGlzQXJnKSB7XG4gIHJldHVybiBmdW5jdGlvbigpIHtcbiAgICBmbi5hcHBseSh0aGlzQXJnLCBhcmd1bWVudHMpO1xuICB9O1xufVxuXG4vKipcbiAqIEBjb25zdHJ1Y3RvclxuICogQHBhcmFtIHtGdW5jdGlvbn0gZm5cbiAqL1xuZnVuY3Rpb24gUHJvbWlzZShmbikge1xuICBpZiAoISh0aGlzIGluc3RhbmNlb2YgUHJvbWlzZSkpXG4gICAgdGhyb3cgbmV3IFR5cGVFcnJvcignUHJvbWlzZXMgbXVzdCBiZSBjb25zdHJ1Y3RlZCB2aWEgbmV3Jyk7XG4gIGlmICh0eXBlb2YgZm4gIT09ICdmdW5jdGlvbicpIHRocm93IG5ldyBUeXBlRXJyb3IoJ25vdCBhIGZ1bmN0aW9uJyk7XG4gIC8qKiBAdHlwZSB7IW51bWJlcn0gKi9cbiAgdGhpcy5fc3RhdGUgPSAwO1xuICAvKiogQHR5cGUgeyFib29sZWFufSAqL1xuICB0aGlzLl9oYW5kbGVkID0gZmFsc2U7XG4gIC8qKiBAdHlwZSB7UHJvbWlzZXx1bmRlZmluZWR9ICovXG4gIHRoaXMuX3ZhbHVlID0gdW5kZWZpbmVkO1xuICAvKiogQHR5cGUgeyFBcnJheTwhRnVuY3Rpb24+fSAqL1xuICB0aGlzLl9kZWZlcnJlZHMgPSBbXTtcblxuICBkb1Jlc29sdmUoZm4sIHRoaXMpO1xufVxuXG5mdW5jdGlvbiBoYW5kbGUoc2VsZiwgZGVmZXJyZWQpIHtcbiAgd2hpbGUgKHNlbGYuX3N0YXRlID09PSAzKSB7XG4gICAgc2VsZiA9IHNlbGYuX3ZhbHVlO1xuICB9XG4gIGlmIChzZWxmLl9zdGF0ZSA9PT0gMCkge1xuICAgIHNlbGYuX2RlZmVycmVkcy5wdXNoKGRlZmVycmVkKTtcbiAgICByZXR1cm47XG4gIH1cbiAgc2VsZi5faGFuZGxlZCA9IHRydWU7XG4gIFByb21pc2UuX2ltbWVkaWF0ZUZuKGZ1bmN0aW9uKCkge1xuICAgIHZhciBjYiA9IHNlbGYuX3N0YXRlID09PSAxID8gZGVmZXJyZWQub25GdWxmaWxsZWQgOiBkZWZlcnJlZC5vblJlamVjdGVkO1xuICAgIGlmIChjYiA9PT0gbnVsbCkge1xuICAgICAgKHNlbGYuX3N0YXRlID09PSAxID8gcmVzb2x2ZSA6IHJlamVjdCkoZGVmZXJyZWQucHJvbWlzZSwgc2VsZi5fdmFsdWUpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cbiAgICB2YXIgcmV0O1xuICAgIHRyeSB7XG4gICAgICByZXQgPSBjYihzZWxmLl92YWx1ZSk7XG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgcmVqZWN0KGRlZmVycmVkLnByb21pc2UsIGUpO1xuICAgICAgcmV0dXJuO1xuICAgIH1cbiAgICByZXNvbHZlKGRlZmVycmVkLnByb21pc2UsIHJldCk7XG4gIH0pO1xufVxuXG5mdW5jdGlvbiByZXNvbHZlKHNlbGYsIG5ld1ZhbHVlKSB7XG4gIHRyeSB7XG4gICAgLy8gUHJvbWlzZSBSZXNvbHV0aW9uIFByb2NlZHVyZTogaHR0cHM6Ly9naXRodWIuY29tL3Byb21pc2VzLWFwbHVzL3Byb21pc2VzLXNwZWMjdGhlLXByb21pc2UtcmVzb2x1dGlvbi1wcm9jZWR1cmVcbiAgICBpZiAobmV3VmFsdWUgPT09IHNlbGYpXG4gICAgICB0aHJvdyBuZXcgVHlwZUVycm9yKCdBIHByb21pc2UgY2Fubm90IGJlIHJlc29sdmVkIHdpdGggaXRzZWxmLicpO1xuICAgIGlmIChcbiAgICAgIG5ld1ZhbHVlICYmXG4gICAgICAodHlwZW9mIG5ld1ZhbHVlID09PSAnb2JqZWN0JyB8fCB0eXBlb2YgbmV3VmFsdWUgPT09ICdmdW5jdGlvbicpXG4gICAgKSB7XG4gICAgICB2YXIgdGhlbiA9IG5ld1ZhbHVlLnRoZW47XG4gICAgICBpZiAobmV3VmFsdWUgaW5zdGFuY2VvZiBQcm9taXNlKSB7XG4gICAgICAgIHNlbGYuX3N0YXRlID0gMztcbiAgICAgICAgc2VsZi5fdmFsdWUgPSBuZXdWYWx1ZTtcbiAgICAgICAgZmluYWxlKHNlbGYpO1xuICAgICAgICByZXR1cm47XG4gICAgICB9IGVsc2UgaWYgKHR5cGVvZiB0aGVuID09PSAnZnVuY3Rpb24nKSB7XG4gICAgICAgIGRvUmVzb2x2ZShiaW5kKHRoZW4sIG5ld1ZhbHVlKSwgc2VsZik7XG4gICAgICAgIHJldHVybjtcbiAgICAgIH1cbiAgICB9XG4gICAgc2VsZi5fc3RhdGUgPSAxO1xuICAgIHNlbGYuX3ZhbHVlID0gbmV3VmFsdWU7XG4gICAgZmluYWxlKHNlbGYpO1xuICB9IGNhdGNoIChlKSB7XG4gICAgcmVqZWN0KHNlbGYsIGUpO1xuICB9XG59XG5cbmZ1bmN0aW9uIHJlamVjdChzZWxmLCBuZXdWYWx1ZSkge1xuICBzZWxmLl9zdGF0ZSA9IDI7XG4gIHNlbGYuX3ZhbHVlID0gbmV3VmFsdWU7XG4gIGZpbmFsZShzZWxmKTtcbn1cblxuZnVuY3Rpb24gZmluYWxlKHNlbGYpIHtcbiAgaWYgKHNlbGYuX3N0YXRlID09PSAyICYmIHNlbGYuX2RlZmVycmVkcy5sZW5ndGggPT09IDApIHtcbiAgICBQcm9taXNlLl9pbW1lZGlhdGVGbihmdW5jdGlvbigpIHtcbiAgICAgIGlmICghc2VsZi5faGFuZGxlZCkge1xuICAgICAgICBQcm9taXNlLl91bmhhbmRsZWRSZWplY3Rpb25GbihzZWxmLl92YWx1ZSk7XG4gICAgICB9XG4gICAgfSk7XG4gIH1cblxuICBmb3IgKHZhciBpID0gMCwgbGVuID0gc2VsZi5fZGVmZXJyZWRzLmxlbmd0aDsgaSA8IGxlbjsgaSsrKSB7XG4gICAgaGFuZGxlKHNlbGYsIHNlbGYuX2RlZmVycmVkc1tpXSk7XG4gIH1cbiAgc2VsZi5fZGVmZXJyZWRzID0gbnVsbDtcbn1cblxuLyoqXG4gKiBAY29uc3RydWN0b3JcbiAqL1xuZnVuY3Rpb24gSGFuZGxlcihvbkZ1bGZpbGxlZCwgb25SZWplY3RlZCwgcHJvbWlzZSkge1xuICB0aGlzLm9uRnVsZmlsbGVkID0gdHlwZW9mIG9uRnVsZmlsbGVkID09PSAnZnVuY3Rpb24nID8gb25GdWxmaWxsZWQgOiBudWxsO1xuICB0aGlzLm9uUmVqZWN0ZWQgPSB0eXBlb2Ygb25SZWplY3RlZCA9PT0gJ2Z1bmN0aW9uJyA/IG9uUmVqZWN0ZWQgOiBudWxsO1xuICB0aGlzLnByb21pc2UgPSBwcm9taXNlO1xufVxuXG4vKipcbiAqIFRha2UgYSBwb3RlbnRpYWxseSBtaXNiZWhhdmluZyByZXNvbHZlciBmdW5jdGlvbiBhbmQgbWFrZSBzdXJlXG4gKiBvbkZ1bGZpbGxlZCBhbmQgb25SZWplY3RlZCBhcmUgb25seSBjYWxsZWQgb25jZS5cbiAqXG4gKiBNYWtlcyBubyBndWFyYW50ZWVzIGFib3V0IGFzeW5jaHJvbnkuXG4gKi9cbmZ1bmN0aW9uIGRvUmVzb2x2ZShmbiwgc2VsZikge1xuICB2YXIgZG9uZSA9IGZhbHNlO1xuICB0cnkge1xuICAgIGZuKFxuICAgICAgZnVuY3Rpb24odmFsdWUpIHtcbiAgICAgICAgaWYgKGRvbmUpIHJldHVybjtcbiAgICAgICAgZG9uZSA9IHRydWU7XG4gICAgICAgIHJlc29sdmUoc2VsZiwgdmFsdWUpO1xuICAgICAgfSxcbiAgICAgIGZ1bmN0aW9uKHJlYXNvbikge1xuICAgICAgICBpZiAoZG9uZSkgcmV0dXJuO1xuICAgICAgICBkb25lID0gdHJ1ZTtcbiAgICAgICAgcmVqZWN0KHNlbGYsIHJlYXNvbik7XG4gICAgICB9XG4gICAgKTtcbiAgfSBjYXRjaCAoZXgpIHtcbiAgICBpZiAoZG9uZSkgcmV0dXJuO1xuICAgIGRvbmUgPSB0cnVlO1xuICAgIHJlamVjdChzZWxmLCBleCk7XG4gIH1cbn1cblxuUHJvbWlzZS5wcm90b3R5cGVbJ2NhdGNoJ10gPSBmdW5jdGlvbihvblJlamVjdGVkKSB7XG4gIHJldHVybiB0aGlzLnRoZW4obnVsbCwgb25SZWplY3RlZCk7XG59O1xuXG5Qcm9taXNlLnByb3RvdHlwZS50aGVuID0gZnVuY3Rpb24ob25GdWxmaWxsZWQsIG9uUmVqZWN0ZWQpIHtcbiAgLy8gQHRzLWlnbm9yZVxuICB2YXIgcHJvbSA9IG5ldyB0aGlzLmNvbnN0cnVjdG9yKG5vb3ApO1xuXG4gIGhhbmRsZSh0aGlzLCBuZXcgSGFuZGxlcihvbkZ1bGZpbGxlZCwgb25SZWplY3RlZCwgcHJvbSkpO1xuICByZXR1cm4gcHJvbTtcbn07XG5cblByb21pc2UucHJvdG90eXBlWydmaW5hbGx5J10gPSBmaW5hbGx5Q29uc3RydWN0b3I7XG5cblByb21pc2UuYWxsID0gZnVuY3Rpb24oYXJyKSB7XG4gIHJldHVybiBuZXcgUHJvbWlzZShmdW5jdGlvbihyZXNvbHZlLCByZWplY3QpIHtcbiAgICBpZiAoIWFyciB8fCB0eXBlb2YgYXJyLmxlbmd0aCA9PT0gJ3VuZGVmaW5lZCcpXG4gICAgICB0aHJvdyBuZXcgVHlwZUVycm9yKCdQcm9taXNlLmFsbCBhY2NlcHRzIGFuIGFycmF5Jyk7XG4gICAgdmFyIGFyZ3MgPSBBcnJheS5wcm90b3R5cGUuc2xpY2UuY2FsbChhcnIpO1xuICAgIGlmIChhcmdzLmxlbmd0aCA9PT0gMCkgcmV0dXJuIHJlc29sdmUoW10pO1xuICAgIHZhciByZW1haW5pbmcgPSBhcmdzLmxlbmd0aDtcblxuICAgIGZ1bmN0aW9uIHJlcyhpLCB2YWwpIHtcbiAgICAgIHRyeSB7XG4gICAgICAgIGlmICh2YWwgJiYgKHR5cGVvZiB2YWwgPT09ICdvYmplY3QnIHx8IHR5cGVvZiB2YWwgPT09ICdmdW5jdGlvbicpKSB7XG4gICAgICAgICAgdmFyIHRoZW4gPSB2YWwudGhlbjtcbiAgICAgICAgICBpZiAodHlwZW9mIHRoZW4gPT09ICdmdW5jdGlvbicpIHtcbiAgICAgICAgICAgIHRoZW4uY2FsbChcbiAgICAgICAgICAgICAgdmFsLFxuICAgICAgICAgICAgICBmdW5jdGlvbih2YWwpIHtcbiAgICAgICAgICAgICAgICByZXMoaSwgdmFsKTtcbiAgICAgICAgICAgICAgfSxcbiAgICAgICAgICAgICAgcmVqZWN0XG4gICAgICAgICAgICApO1xuICAgICAgICAgICAgcmV0dXJuO1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICBhcmdzW2ldID0gdmFsO1xuICAgICAgICBpZiAoLS1yZW1haW5pbmcgPT09IDApIHtcbiAgICAgICAgICByZXNvbHZlKGFyZ3MpO1xuICAgICAgICB9XG4gICAgICB9IGNhdGNoIChleCkge1xuICAgICAgICByZWplY3QoZXgpO1xuICAgICAgfVxuICAgIH1cblxuICAgIGZvciAodmFyIGkgPSAwOyBpIDwgYXJncy5sZW5ndGg7IGkrKykge1xuICAgICAgcmVzKGksIGFyZ3NbaV0pO1xuICAgIH1cbiAgfSk7XG59O1xuXG5Qcm9taXNlLnJlc29sdmUgPSBmdW5jdGlvbih2YWx1ZSkge1xuICBpZiAodmFsdWUgJiYgdHlwZW9mIHZhbHVlID09PSAnb2JqZWN0JyAmJiB2YWx1ZS5jb25zdHJ1Y3RvciA9PT0gUHJvbWlzZSkge1xuICAgIHJldHVybiB2YWx1ZTtcbiAgfVxuXG4gIHJldHVybiBuZXcgUHJvbWlzZShmdW5jdGlvbihyZXNvbHZlKSB7XG4gICAgcmVzb2x2ZSh2YWx1ZSk7XG4gIH0pO1xufTtcblxuUHJvbWlzZS5yZWplY3QgPSBmdW5jdGlvbih2YWx1ZSkge1xuICByZXR1cm4gbmV3IFByb21pc2UoZnVuY3Rpb24ocmVzb2x2ZSwgcmVqZWN0KSB7XG4gICAgcmVqZWN0KHZhbHVlKTtcbiAgfSk7XG59O1xuXG5Qcm9taXNlLnJhY2UgPSBmdW5jdGlvbih2YWx1ZXMpIHtcbiAgcmV0dXJuIG5ldyBQcm9taXNlKGZ1bmN0aW9uKHJlc29sdmUsIHJlamVjdCkge1xuICAgIGZvciAodmFyIGkgPSAwLCBsZW4gPSB2YWx1ZXMubGVuZ3RoOyBpIDwgbGVuOyBpKyspIHtcbiAgICAgIHZhbHVlc1tpXS50aGVuKHJlc29sdmUsIHJlamVjdCk7XG4gICAgfVxuICB9KTtcbn07XG5cbi8vIFVzZSBwb2x5ZmlsbCBmb3Igc2V0SW1tZWRpYXRlIGZvciBwZXJmb3JtYW5jZSBnYWluc1xuUHJvbWlzZS5faW1tZWRpYXRlRm4gPVxuICAodHlwZW9mIHNldEltbWVkaWF0ZSA9PT0gJ2Z1bmN0aW9uJyAmJlxuICAgIGZ1bmN0aW9uKGZuKSB7XG4gICAgICBzZXRJbW1lZGlhdGUoZm4pO1xuICAgIH0pIHx8XG4gIGZ1bmN0aW9uKGZuKSB7XG4gICAgc2V0VGltZW91dEZ1bmMoZm4sIDApO1xuICB9O1xuXG5Qcm9taXNlLl91bmhhbmRsZWRSZWplY3Rpb25GbiA9IGZ1bmN0aW9uIF91bmhhbmRsZWRSZWplY3Rpb25GbihlcnIpIHtcbiAgaWYgKHR5cGVvZiBjb25zb2xlICE9PSAndW5kZWZpbmVkJyAmJiBjb25zb2xlKSB7XG4gICAgY29uc29sZS53YXJuKCdQb3NzaWJsZSBVbmhhbmRsZWQgUHJvbWlzZSBSZWplY3Rpb246JywgZXJyKTsgLy8gZXNsaW50LWRpc2FibGUtbGluZSBuby1jb25zb2xlXG4gIH1cbn07XG5cbm1vZHVsZS5leHBvcnRzID0gUHJvbWlzZTtcbiIsInZhciBuZXh0VGljayA9IHJlcXVpcmUoJ3Byb2Nlc3MvYnJvd3Nlci5qcycpLm5leHRUaWNrO1xudmFyIGFwcGx5ID0gRnVuY3Rpb24ucHJvdG90eXBlLmFwcGx5O1xudmFyIHNsaWNlID0gQXJyYXkucHJvdG90eXBlLnNsaWNlO1xudmFyIGltbWVkaWF0ZUlkcyA9IHt9O1xudmFyIG5leHRJbW1lZGlhdGVJZCA9IDA7XG5cbi8vIERPTSBBUElzLCBmb3IgY29tcGxldGVuZXNzXG5cbmV4cG9ydHMuc2V0VGltZW91dCA9IGZ1bmN0aW9uKCkge1xuICByZXR1cm4gbmV3IFRpbWVvdXQoYXBwbHkuY2FsbChzZXRUaW1lb3V0LCB3aW5kb3csIGFyZ3VtZW50cyksIGNsZWFyVGltZW91dCk7XG59O1xuZXhwb3J0cy5zZXRJbnRlcnZhbCA9IGZ1bmN0aW9uKCkge1xuICByZXR1cm4gbmV3IFRpbWVvdXQoYXBwbHkuY2FsbChzZXRJbnRlcnZhbCwgd2luZG93LCBhcmd1bWVudHMpLCBjbGVhckludGVydmFsKTtcbn07XG5leHBvcnRzLmNsZWFyVGltZW91dCA9XG5leHBvcnRzLmNsZWFySW50ZXJ2YWwgPSBmdW5jdGlvbih0aW1lb3V0KSB7IHRpbWVvdXQuY2xvc2UoKTsgfTtcblxuZnVuY3Rpb24gVGltZW91dChpZCwgY2xlYXJGbikge1xuICB0aGlzLl9pZCA9IGlkO1xuICB0aGlzLl9jbGVhckZuID0gY2xlYXJGbjtcbn1cblRpbWVvdXQucHJvdG90eXBlLnVucmVmID0gVGltZW91dC5wcm90b3R5cGUucmVmID0gZnVuY3Rpb24oKSB7fTtcblRpbWVvdXQucHJvdG90eXBlLmNsb3NlID0gZnVuY3Rpb24oKSB7XG4gIHRoaXMuX2NsZWFyRm4uY2FsbCh3aW5kb3csIHRoaXMuX2lkKTtcbn07XG5cbi8vIERvZXMgbm90IHN0YXJ0IHRoZSB0aW1lLCBqdXN0IHNldHMgdXAgdGhlIG1lbWJlcnMgbmVlZGVkLlxuZXhwb3J0cy5lbnJvbGwgPSBmdW5jdGlvbihpdGVtLCBtc2Vjcykge1xuICBjbGVhclRpbWVvdXQoaXRlbS5faWRsZVRpbWVvdXRJZCk7XG4gIGl0ZW0uX2lkbGVUaW1lb3V0ID0gbXNlY3M7XG59O1xuXG5leHBvcnRzLnVuZW5yb2xsID0gZnVuY3Rpb24oaXRlbSkge1xuICBjbGVhclRpbWVvdXQoaXRlbS5faWRsZVRpbWVvdXRJZCk7XG4gIGl0ZW0uX2lkbGVUaW1lb3V0ID0gLTE7XG59O1xuXG5leHBvcnRzLl91bnJlZkFjdGl2ZSA9IGV4cG9ydHMuYWN0aXZlID0gZnVuY3Rpb24oaXRlbSkge1xuICBjbGVhclRpbWVvdXQoaXRlbS5faWRsZVRpbWVvdXRJZCk7XG5cbiAgdmFyIG1zZWNzID0gaXRlbS5faWRsZVRpbWVvdXQ7XG4gIGlmIChtc2VjcyA+PSAwKSB7XG4gICAgaXRlbS5faWRsZVRpbWVvdXRJZCA9IHNldFRpbWVvdXQoZnVuY3Rpb24gb25UaW1lb3V0KCkge1xuICAgICAgaWYgKGl0ZW0uX29uVGltZW91dClcbiAgICAgICAgaXRlbS5fb25UaW1lb3V0KCk7XG4gICAgfSwgbXNlY3MpO1xuICB9XG59O1xuXG4vLyBUaGF0J3Mgbm90IGhvdyBub2RlLmpzIGltcGxlbWVudHMgaXQgYnV0IHRoZSBleHBvc2VkIGFwaSBpcyB0aGUgc2FtZS5cbmV4cG9ydHMuc2V0SW1tZWRpYXRlID0gdHlwZW9mIHNldEltbWVkaWF0ZSA9PT0gXCJmdW5jdGlvblwiID8gc2V0SW1tZWRpYXRlIDogZnVuY3Rpb24oZm4pIHtcbiAgdmFyIGlkID0gbmV4dEltbWVkaWF0ZUlkKys7XG4gIHZhciBhcmdzID0gYXJndW1lbnRzLmxlbmd0aCA8IDIgPyBmYWxzZSA6IHNsaWNlLmNhbGwoYXJndW1lbnRzLCAxKTtcblxuICBpbW1lZGlhdGVJZHNbaWRdID0gdHJ1ZTtcblxuICBuZXh0VGljayhmdW5jdGlvbiBvbk5leHRUaWNrKCkge1xuICAgIGlmIChpbW1lZGlhdGVJZHNbaWRdKSB7XG4gICAgICAvLyBmbi5jYWxsKCkgaXMgZmFzdGVyIHNvIHdlIG9wdGltaXplIGZvciB0aGUgY29tbW9uIHVzZS1jYXNlXG4gICAgICAvLyBAc2VlIGh0dHA6Ly9qc3BlcmYuY29tL2NhbGwtYXBwbHktc2VndVxuICAgICAgaWYgKGFyZ3MpIHtcbiAgICAgICAgZm4uYXBwbHkobnVsbCwgYXJncyk7XG4gICAgICB9IGVsc2Uge1xuICAgICAgICBmbi5jYWxsKG51bGwpO1xuICAgICAgfVxuICAgICAgLy8gUHJldmVudCBpZHMgZnJvbSBsZWFraW5nXG4gICAgICBleHBvcnRzLmNsZWFySW1tZWRpYXRlKGlkKTtcbiAgICB9XG4gIH0pO1xuXG4gIHJldHVybiBpZDtcbn07XG5cbmV4cG9ydHMuY2xlYXJJbW1lZGlhdGUgPSB0eXBlb2YgY2xlYXJJbW1lZGlhdGUgPT09IFwiZnVuY3Rpb25cIiA/IGNsZWFySW1tZWRpYXRlIDogZnVuY3Rpb24oaWQpIHtcbiAgZGVsZXRlIGltbWVkaWF0ZUlkc1tpZF07XG59OyIsIihmdW5jdGlvbiAoZ2xvYmFsLCBmYWN0b3J5KSB7XG4gIHR5cGVvZiBleHBvcnRzID09PSAnb2JqZWN0JyAmJiB0eXBlb2YgbW9kdWxlICE9PSAndW5kZWZpbmVkJyA/IGZhY3RvcnkoZXhwb3J0cykgOlxuICB0eXBlb2YgZGVmaW5lID09PSAnZnVuY3Rpb24nICYmIGRlZmluZS5hbWQgPyBkZWZpbmUoWydleHBvcnRzJ10sIGZhY3RvcnkpIDpcbiAgKGZhY3RvcnkoKGdsb2JhbC5XSEFUV0dGZXRjaCA9IHt9KSkpO1xufSh0aGlzLCAoZnVuY3Rpb24gKGV4cG9ydHMpIHsgJ3VzZSBzdHJpY3QnO1xuXG4gIHZhciBzdXBwb3J0ID0ge1xuICAgIHNlYXJjaFBhcmFtczogJ1VSTFNlYXJjaFBhcmFtcycgaW4gc2VsZixcbiAgICBpdGVyYWJsZTogJ1N5bWJvbCcgaW4gc2VsZiAmJiAnaXRlcmF0b3InIGluIFN5bWJvbCxcbiAgICBibG9iOlxuICAgICAgJ0ZpbGVSZWFkZXInIGluIHNlbGYgJiZcbiAgICAgICdCbG9iJyBpbiBzZWxmICYmXG4gICAgICAoZnVuY3Rpb24oKSB7XG4gICAgICAgIHRyeSB7XG4gICAgICAgICAgbmV3IEJsb2IoKTtcbiAgICAgICAgICByZXR1cm4gdHJ1ZVxuICAgICAgICB9IGNhdGNoIChlKSB7XG4gICAgICAgICAgcmV0dXJuIGZhbHNlXG4gICAgICAgIH1cbiAgICAgIH0pKCksXG4gICAgZm9ybURhdGE6ICdGb3JtRGF0YScgaW4gc2VsZixcbiAgICBhcnJheUJ1ZmZlcjogJ0FycmF5QnVmZmVyJyBpbiBzZWxmXG4gIH07XG5cbiAgZnVuY3Rpb24gaXNEYXRhVmlldyhvYmopIHtcbiAgICByZXR1cm4gb2JqICYmIERhdGFWaWV3LnByb3RvdHlwZS5pc1Byb3RvdHlwZU9mKG9iailcbiAgfVxuXG4gIGlmIChzdXBwb3J0LmFycmF5QnVmZmVyKSB7XG4gICAgdmFyIHZpZXdDbGFzc2VzID0gW1xuICAgICAgJ1tvYmplY3QgSW50OEFycmF5XScsXG4gICAgICAnW29iamVjdCBVaW50OEFycmF5XScsXG4gICAgICAnW29iamVjdCBVaW50OENsYW1wZWRBcnJheV0nLFxuICAgICAgJ1tvYmplY3QgSW50MTZBcnJheV0nLFxuICAgICAgJ1tvYmplY3QgVWludDE2QXJyYXldJyxcbiAgICAgICdbb2JqZWN0IEludDMyQXJyYXldJyxcbiAgICAgICdbb2JqZWN0IFVpbnQzMkFycmF5XScsXG4gICAgICAnW29iamVjdCBGbG9hdDMyQXJyYXldJyxcbiAgICAgICdbb2JqZWN0IEZsb2F0NjRBcnJheV0nXG4gICAgXTtcblxuICAgIHZhciBpc0FycmF5QnVmZmVyVmlldyA9XG4gICAgICBBcnJheUJ1ZmZlci5pc1ZpZXcgfHxcbiAgICAgIGZ1bmN0aW9uKG9iaikge1xuICAgICAgICByZXR1cm4gb2JqICYmIHZpZXdDbGFzc2VzLmluZGV4T2YoT2JqZWN0LnByb3RvdHlwZS50b1N0cmluZy5jYWxsKG9iaikpID4gLTFcbiAgICAgIH07XG4gIH1cblxuICBmdW5jdGlvbiBub3JtYWxpemVOYW1lKG5hbWUpIHtcbiAgICBpZiAodHlwZW9mIG5hbWUgIT09ICdzdHJpbmcnKSB7XG4gICAgICBuYW1lID0gU3RyaW5nKG5hbWUpO1xuICAgIH1cbiAgICBpZiAoL1teYS16MC05XFwtIyQlJicqKy5eX2B8fl0vaS50ZXN0KG5hbWUpKSB7XG4gICAgICB0aHJvdyBuZXcgVHlwZUVycm9yKCdJbnZhbGlkIGNoYXJhY3RlciBpbiBoZWFkZXIgZmllbGQgbmFtZScpXG4gICAgfVxuICAgIHJldHVybiBuYW1lLnRvTG93ZXJDYXNlKClcbiAgfVxuXG4gIGZ1bmN0aW9uIG5vcm1hbGl6ZVZhbHVlKHZhbHVlKSB7XG4gICAgaWYgKHR5cGVvZiB2YWx1ZSAhPT0gJ3N0cmluZycpIHtcbiAgICAgIHZhbHVlID0gU3RyaW5nKHZhbHVlKTtcbiAgICB9XG4gICAgcmV0dXJuIHZhbHVlXG4gIH1cblxuICAvLyBCdWlsZCBhIGRlc3RydWN0aXZlIGl0ZXJhdG9yIGZvciB0aGUgdmFsdWUgbGlzdFxuICBmdW5jdGlvbiBpdGVyYXRvckZvcihpdGVtcykge1xuICAgIHZhciBpdGVyYXRvciA9IHtcbiAgICAgIG5leHQ6IGZ1bmN0aW9uKCkge1xuICAgICAgICB2YXIgdmFsdWUgPSBpdGVtcy5zaGlmdCgpO1xuICAgICAgICByZXR1cm4ge2RvbmU6IHZhbHVlID09PSB1bmRlZmluZWQsIHZhbHVlOiB2YWx1ZX1cbiAgICAgIH1cbiAgICB9O1xuXG4gICAgaWYgKHN1cHBvcnQuaXRlcmFibGUpIHtcbiAgICAgIGl0ZXJhdG9yW1N5bWJvbC5pdGVyYXRvcl0gPSBmdW5jdGlvbigpIHtcbiAgICAgICAgcmV0dXJuIGl0ZXJhdG9yXG4gICAgICB9O1xuICAgIH1cblxuICAgIHJldHVybiBpdGVyYXRvclxuICB9XG5cbiAgZnVuY3Rpb24gSGVhZGVycyhoZWFkZXJzKSB7XG4gICAgdGhpcy5tYXAgPSB7fTtcblxuICAgIGlmIChoZWFkZXJzIGluc3RhbmNlb2YgSGVhZGVycykge1xuICAgICAgaGVhZGVycy5mb3JFYWNoKGZ1bmN0aW9uKHZhbHVlLCBuYW1lKSB7XG4gICAgICAgIHRoaXMuYXBwZW5kKG5hbWUsIHZhbHVlKTtcbiAgICAgIH0sIHRoaXMpO1xuICAgIH0gZWxzZSBpZiAoQXJyYXkuaXNBcnJheShoZWFkZXJzKSkge1xuICAgICAgaGVhZGVycy5mb3JFYWNoKGZ1bmN0aW9uKGhlYWRlcikge1xuICAgICAgICB0aGlzLmFwcGVuZChoZWFkZXJbMF0sIGhlYWRlclsxXSk7XG4gICAgICB9LCB0aGlzKTtcbiAgICB9IGVsc2UgaWYgKGhlYWRlcnMpIHtcbiAgICAgIE9iamVjdC5nZXRPd25Qcm9wZXJ0eU5hbWVzKGhlYWRlcnMpLmZvckVhY2goZnVuY3Rpb24obmFtZSkge1xuICAgICAgICB0aGlzLmFwcGVuZChuYW1lLCBoZWFkZXJzW25hbWVdKTtcbiAgICAgIH0sIHRoaXMpO1xuICAgIH1cbiAgfVxuXG4gIEhlYWRlcnMucHJvdG90eXBlLmFwcGVuZCA9IGZ1bmN0aW9uKG5hbWUsIHZhbHVlKSB7XG4gICAgbmFtZSA9IG5vcm1hbGl6ZU5hbWUobmFtZSk7XG4gICAgdmFsdWUgPSBub3JtYWxpemVWYWx1ZSh2YWx1ZSk7XG4gICAgdmFyIG9sZFZhbHVlID0gdGhpcy5tYXBbbmFtZV07XG4gICAgdGhpcy5tYXBbbmFtZV0gPSBvbGRWYWx1ZSA/IG9sZFZhbHVlICsgJywgJyArIHZhbHVlIDogdmFsdWU7XG4gIH07XG5cbiAgSGVhZGVycy5wcm90b3R5cGVbJ2RlbGV0ZSddID0gZnVuY3Rpb24obmFtZSkge1xuICAgIGRlbGV0ZSB0aGlzLm1hcFtub3JtYWxpemVOYW1lKG5hbWUpXTtcbiAgfTtcblxuICBIZWFkZXJzLnByb3RvdHlwZS5nZXQgPSBmdW5jdGlvbihuYW1lKSB7XG4gICAgbmFtZSA9IG5vcm1hbGl6ZU5hbWUobmFtZSk7XG4gICAgcmV0dXJuIHRoaXMuaGFzKG5hbWUpID8gdGhpcy5tYXBbbmFtZV0gOiBudWxsXG4gIH07XG5cbiAgSGVhZGVycy5wcm90b3R5cGUuaGFzID0gZnVuY3Rpb24obmFtZSkge1xuICAgIHJldHVybiB0aGlzLm1hcC5oYXNPd25Qcm9wZXJ0eShub3JtYWxpemVOYW1lKG5hbWUpKVxuICB9O1xuXG4gIEhlYWRlcnMucHJvdG90eXBlLnNldCA9IGZ1bmN0aW9uKG5hbWUsIHZhbHVlKSB7XG4gICAgdGhpcy5tYXBbbm9ybWFsaXplTmFtZShuYW1lKV0gPSBub3JtYWxpemVWYWx1ZSh2YWx1ZSk7XG4gIH07XG5cbiAgSGVhZGVycy5wcm90b3R5cGUuZm9yRWFjaCA9IGZ1bmN0aW9uKGNhbGxiYWNrLCB0aGlzQXJnKSB7XG4gICAgZm9yICh2YXIgbmFtZSBpbiB0aGlzLm1hcCkge1xuICAgICAgaWYgKHRoaXMubWFwLmhhc093blByb3BlcnR5KG5hbWUpKSB7XG4gICAgICAgIGNhbGxiYWNrLmNhbGwodGhpc0FyZywgdGhpcy5tYXBbbmFtZV0sIG5hbWUsIHRoaXMpO1xuICAgICAgfVxuICAgIH1cbiAgfTtcblxuICBIZWFkZXJzLnByb3RvdHlwZS5rZXlzID0gZnVuY3Rpb24oKSB7XG4gICAgdmFyIGl0ZW1zID0gW107XG4gICAgdGhpcy5mb3JFYWNoKGZ1bmN0aW9uKHZhbHVlLCBuYW1lKSB7XG4gICAgICBpdGVtcy5wdXNoKG5hbWUpO1xuICAgIH0pO1xuICAgIHJldHVybiBpdGVyYXRvckZvcihpdGVtcylcbiAgfTtcblxuICBIZWFkZXJzLnByb3RvdHlwZS52YWx1ZXMgPSBmdW5jdGlvbigpIHtcbiAgICB2YXIgaXRlbXMgPSBbXTtcbiAgICB0aGlzLmZvckVhY2goZnVuY3Rpb24odmFsdWUpIHtcbiAgICAgIGl0ZW1zLnB1c2godmFsdWUpO1xuICAgIH0pO1xuICAgIHJldHVybiBpdGVyYXRvckZvcihpdGVtcylcbiAgfTtcblxuICBIZWFkZXJzLnByb3RvdHlwZS5lbnRyaWVzID0gZnVuY3Rpb24oKSB7XG4gICAgdmFyIGl0ZW1zID0gW107XG4gICAgdGhpcy5mb3JFYWNoKGZ1bmN0aW9uKHZhbHVlLCBuYW1lKSB7XG4gICAgICBpdGVtcy5wdXNoKFtuYW1lLCB2YWx1ZV0pO1xuICAgIH0pO1xuICAgIHJldHVybiBpdGVyYXRvckZvcihpdGVtcylcbiAgfTtcblxuICBpZiAoc3VwcG9ydC5pdGVyYWJsZSkge1xuICAgIEhlYWRlcnMucHJvdG90eXBlW1N5bWJvbC5pdGVyYXRvcl0gPSBIZWFkZXJzLnByb3RvdHlwZS5lbnRyaWVzO1xuICB9XG5cbiAgZnVuY3Rpb24gY29uc3VtZWQoYm9keSkge1xuICAgIGlmIChib2R5LmJvZHlVc2VkKSB7XG4gICAgICByZXR1cm4gUHJvbWlzZS5yZWplY3QobmV3IFR5cGVFcnJvcignQWxyZWFkeSByZWFkJykpXG4gICAgfVxuICAgIGJvZHkuYm9keVVzZWQgPSB0cnVlO1xuICB9XG5cbiAgZnVuY3Rpb24gZmlsZVJlYWRlclJlYWR5KHJlYWRlcikge1xuICAgIHJldHVybiBuZXcgUHJvbWlzZShmdW5jdGlvbihyZXNvbHZlLCByZWplY3QpIHtcbiAgICAgIHJlYWRlci5vbmxvYWQgPSBmdW5jdGlvbigpIHtcbiAgICAgICAgcmVzb2x2ZShyZWFkZXIucmVzdWx0KTtcbiAgICAgIH07XG4gICAgICByZWFkZXIub25lcnJvciA9IGZ1bmN0aW9uKCkge1xuICAgICAgICByZWplY3QocmVhZGVyLmVycm9yKTtcbiAgICAgIH07XG4gICAgfSlcbiAgfVxuXG4gIGZ1bmN0aW9uIHJlYWRCbG9iQXNBcnJheUJ1ZmZlcihibG9iKSB7XG4gICAgdmFyIHJlYWRlciA9IG5ldyBGaWxlUmVhZGVyKCk7XG4gICAgdmFyIHByb21pc2UgPSBmaWxlUmVhZGVyUmVhZHkocmVhZGVyKTtcbiAgICByZWFkZXIucmVhZEFzQXJyYXlCdWZmZXIoYmxvYik7XG4gICAgcmV0dXJuIHByb21pc2VcbiAgfVxuXG4gIGZ1bmN0aW9uIHJlYWRCbG9iQXNUZXh0KGJsb2IpIHtcbiAgICB2YXIgcmVhZGVyID0gbmV3IEZpbGVSZWFkZXIoKTtcbiAgICB2YXIgcHJvbWlzZSA9IGZpbGVSZWFkZXJSZWFkeShyZWFkZXIpO1xuICAgIHJlYWRlci5yZWFkQXNUZXh0KGJsb2IpO1xuICAgIHJldHVybiBwcm9taXNlXG4gIH1cblxuICBmdW5jdGlvbiByZWFkQXJyYXlCdWZmZXJBc1RleHQoYnVmKSB7XG4gICAgdmFyIHZpZXcgPSBuZXcgVWludDhBcnJheShidWYpO1xuICAgIHZhciBjaGFycyA9IG5ldyBBcnJheSh2aWV3Lmxlbmd0aCk7XG5cbiAgICBmb3IgKHZhciBpID0gMDsgaSA8IHZpZXcubGVuZ3RoOyBpKyspIHtcbiAgICAgIGNoYXJzW2ldID0gU3RyaW5nLmZyb21DaGFyQ29kZSh2aWV3W2ldKTtcbiAgICB9XG4gICAgcmV0dXJuIGNoYXJzLmpvaW4oJycpXG4gIH1cblxuICBmdW5jdGlvbiBidWZmZXJDbG9uZShidWYpIHtcbiAgICBpZiAoYnVmLnNsaWNlKSB7XG4gICAgICByZXR1cm4gYnVmLnNsaWNlKDApXG4gICAgfSBlbHNlIHtcbiAgICAgIHZhciB2aWV3ID0gbmV3IFVpbnQ4QXJyYXkoYnVmLmJ5dGVMZW5ndGgpO1xuICAgICAgdmlldy5zZXQobmV3IFVpbnQ4QXJyYXkoYnVmKSk7XG4gICAgICByZXR1cm4gdmlldy5idWZmZXJcbiAgICB9XG4gIH1cblxuICBmdW5jdGlvbiBCb2R5KCkge1xuICAgIHRoaXMuYm9keVVzZWQgPSBmYWxzZTtcblxuICAgIHRoaXMuX2luaXRCb2R5ID0gZnVuY3Rpb24oYm9keSkge1xuICAgICAgdGhpcy5fYm9keUluaXQgPSBib2R5O1xuICAgICAgaWYgKCFib2R5KSB7XG4gICAgICAgIHRoaXMuX2JvZHlUZXh0ID0gJyc7XG4gICAgICB9IGVsc2UgaWYgKHR5cGVvZiBib2R5ID09PSAnc3RyaW5nJykge1xuICAgICAgICB0aGlzLl9ib2R5VGV4dCA9IGJvZHk7XG4gICAgICB9IGVsc2UgaWYgKHN1cHBvcnQuYmxvYiAmJiBCbG9iLnByb3RvdHlwZS5pc1Byb3RvdHlwZU9mKGJvZHkpKSB7XG4gICAgICAgIHRoaXMuX2JvZHlCbG9iID0gYm9keTtcbiAgICAgIH0gZWxzZSBpZiAoc3VwcG9ydC5mb3JtRGF0YSAmJiBGb3JtRGF0YS5wcm90b3R5cGUuaXNQcm90b3R5cGVPZihib2R5KSkge1xuICAgICAgICB0aGlzLl9ib2R5Rm9ybURhdGEgPSBib2R5O1xuICAgICAgfSBlbHNlIGlmIChzdXBwb3J0LnNlYXJjaFBhcmFtcyAmJiBVUkxTZWFyY2hQYXJhbXMucHJvdG90eXBlLmlzUHJvdG90eXBlT2YoYm9keSkpIHtcbiAgICAgICAgdGhpcy5fYm9keVRleHQgPSBib2R5LnRvU3RyaW5nKCk7XG4gICAgICB9IGVsc2UgaWYgKHN1cHBvcnQuYXJyYXlCdWZmZXIgJiYgc3VwcG9ydC5ibG9iICYmIGlzRGF0YVZpZXcoYm9keSkpIHtcbiAgICAgICAgdGhpcy5fYm9keUFycmF5QnVmZmVyID0gYnVmZmVyQ2xvbmUoYm9keS5idWZmZXIpO1xuICAgICAgICAvLyBJRSAxMC0xMSBjYW4ndCBoYW5kbGUgYSBEYXRhVmlldyBib2R5LlxuICAgICAgICB0aGlzLl9ib2R5SW5pdCA9IG5ldyBCbG9iKFt0aGlzLl9ib2R5QXJyYXlCdWZmZXJdKTtcbiAgICAgIH0gZWxzZSBpZiAoc3VwcG9ydC5hcnJheUJ1ZmZlciAmJiAoQXJyYXlCdWZmZXIucHJvdG90eXBlLmlzUHJvdG90eXBlT2YoYm9keSkgfHwgaXNBcnJheUJ1ZmZlclZpZXcoYm9keSkpKSB7XG4gICAgICAgIHRoaXMuX2JvZHlBcnJheUJ1ZmZlciA9IGJ1ZmZlckNsb25lKGJvZHkpO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgdGhpcy5fYm9keVRleHQgPSBib2R5ID0gT2JqZWN0LnByb3RvdHlwZS50b1N0cmluZy5jYWxsKGJvZHkpO1xuICAgICAgfVxuXG4gICAgICBpZiAoIXRoaXMuaGVhZGVycy5nZXQoJ2NvbnRlbnQtdHlwZScpKSB7XG4gICAgICAgIGlmICh0eXBlb2YgYm9keSA9PT0gJ3N0cmluZycpIHtcbiAgICAgICAgICB0aGlzLmhlYWRlcnMuc2V0KCdjb250ZW50LXR5cGUnLCAndGV4dC9wbGFpbjtjaGFyc2V0PVVURi04Jyk7XG4gICAgICAgIH0gZWxzZSBpZiAodGhpcy5fYm9keUJsb2IgJiYgdGhpcy5fYm9keUJsb2IudHlwZSkge1xuICAgICAgICAgIHRoaXMuaGVhZGVycy5zZXQoJ2NvbnRlbnQtdHlwZScsIHRoaXMuX2JvZHlCbG9iLnR5cGUpO1xuICAgICAgICB9IGVsc2UgaWYgKHN1cHBvcnQuc2VhcmNoUGFyYW1zICYmIFVSTFNlYXJjaFBhcmFtcy5wcm90b3R5cGUuaXNQcm90b3R5cGVPZihib2R5KSkge1xuICAgICAgICAgIHRoaXMuaGVhZGVycy5zZXQoJ2NvbnRlbnQtdHlwZScsICdhcHBsaWNhdGlvbi94LXd3dy1mb3JtLXVybGVuY29kZWQ7Y2hhcnNldD1VVEYtOCcpO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfTtcblxuICAgIGlmIChzdXBwb3J0LmJsb2IpIHtcbiAgICAgIHRoaXMuYmxvYiA9IGZ1bmN0aW9uKCkge1xuICAgICAgICB2YXIgcmVqZWN0ZWQgPSBjb25zdW1lZCh0aGlzKTtcbiAgICAgICAgaWYgKHJlamVjdGVkKSB7XG4gICAgICAgICAgcmV0dXJuIHJlamVjdGVkXG4gICAgICAgIH1cblxuICAgICAgICBpZiAodGhpcy5fYm9keUJsb2IpIHtcbiAgICAgICAgICByZXR1cm4gUHJvbWlzZS5yZXNvbHZlKHRoaXMuX2JvZHlCbG9iKVxuICAgICAgICB9IGVsc2UgaWYgKHRoaXMuX2JvZHlBcnJheUJ1ZmZlcikge1xuICAgICAgICAgIHJldHVybiBQcm9taXNlLnJlc29sdmUobmV3IEJsb2IoW3RoaXMuX2JvZHlBcnJheUJ1ZmZlcl0pKVxuICAgICAgICB9IGVsc2UgaWYgKHRoaXMuX2JvZHlGb3JtRGF0YSkge1xuICAgICAgICAgIHRocm93IG5ldyBFcnJvcignY291bGQgbm90IHJlYWQgRm9ybURhdGEgYm9keSBhcyBibG9iJylcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICByZXR1cm4gUHJvbWlzZS5yZXNvbHZlKG5ldyBCbG9iKFt0aGlzLl9ib2R5VGV4dF0pKVxuICAgICAgICB9XG4gICAgICB9O1xuXG4gICAgICB0aGlzLmFycmF5QnVmZmVyID0gZnVuY3Rpb24oKSB7XG4gICAgICAgIGlmICh0aGlzLl9ib2R5QXJyYXlCdWZmZXIpIHtcbiAgICAgICAgICByZXR1cm4gY29uc3VtZWQodGhpcykgfHwgUHJvbWlzZS5yZXNvbHZlKHRoaXMuX2JvZHlBcnJheUJ1ZmZlcilcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICByZXR1cm4gdGhpcy5ibG9iKCkudGhlbihyZWFkQmxvYkFzQXJyYXlCdWZmZXIpXG4gICAgICAgIH1cbiAgICAgIH07XG4gICAgfVxuXG4gICAgdGhpcy50ZXh0ID0gZnVuY3Rpb24oKSB7XG4gICAgICB2YXIgcmVqZWN0ZWQgPSBjb25zdW1lZCh0aGlzKTtcbiAgICAgIGlmIChyZWplY3RlZCkge1xuICAgICAgICByZXR1cm4gcmVqZWN0ZWRcbiAgICAgIH1cblxuICAgICAgaWYgKHRoaXMuX2JvZHlCbG9iKSB7XG4gICAgICAgIHJldHVybiByZWFkQmxvYkFzVGV4dCh0aGlzLl9ib2R5QmxvYilcbiAgICAgIH0gZWxzZSBpZiAodGhpcy5fYm9keUFycmF5QnVmZmVyKSB7XG4gICAgICAgIHJldHVybiBQcm9taXNlLnJlc29sdmUocmVhZEFycmF5QnVmZmVyQXNUZXh0KHRoaXMuX2JvZHlBcnJheUJ1ZmZlcikpXG4gICAgICB9IGVsc2UgaWYgKHRoaXMuX2JvZHlGb3JtRGF0YSkge1xuICAgICAgICB0aHJvdyBuZXcgRXJyb3IoJ2NvdWxkIG5vdCByZWFkIEZvcm1EYXRhIGJvZHkgYXMgdGV4dCcpXG4gICAgICB9IGVsc2Uge1xuICAgICAgICByZXR1cm4gUHJvbWlzZS5yZXNvbHZlKHRoaXMuX2JvZHlUZXh0KVxuICAgICAgfVxuICAgIH07XG5cbiAgICBpZiAoc3VwcG9ydC5mb3JtRGF0YSkge1xuICAgICAgdGhpcy5mb3JtRGF0YSA9IGZ1bmN0aW9uKCkge1xuICAgICAgICByZXR1cm4gdGhpcy50ZXh0KCkudGhlbihkZWNvZGUpXG4gICAgICB9O1xuICAgIH1cblxuICAgIHRoaXMuanNvbiA9IGZ1bmN0aW9uKCkge1xuICAgICAgcmV0dXJuIHRoaXMudGV4dCgpLnRoZW4oSlNPTi5wYXJzZSlcbiAgICB9O1xuXG4gICAgcmV0dXJuIHRoaXNcbiAgfVxuXG4gIC8vIEhUVFAgbWV0aG9kcyB3aG9zZSBjYXBpdGFsaXphdGlvbiBzaG91bGQgYmUgbm9ybWFsaXplZFxuICB2YXIgbWV0aG9kcyA9IFsnREVMRVRFJywgJ0dFVCcsICdIRUFEJywgJ09QVElPTlMnLCAnUE9TVCcsICdQVVQnXTtcblxuICBmdW5jdGlvbiBub3JtYWxpemVNZXRob2QobWV0aG9kKSB7XG4gICAgdmFyIHVwY2FzZWQgPSBtZXRob2QudG9VcHBlckNhc2UoKTtcbiAgICByZXR1cm4gbWV0aG9kcy5pbmRleE9mKHVwY2FzZWQpID4gLTEgPyB1cGNhc2VkIDogbWV0aG9kXG4gIH1cblxuICBmdW5jdGlvbiBSZXF1ZXN0KGlucHV0LCBvcHRpb25zKSB7XG4gICAgb3B0aW9ucyA9IG9wdGlvbnMgfHwge307XG4gICAgdmFyIGJvZHkgPSBvcHRpb25zLmJvZHk7XG5cbiAgICBpZiAoaW5wdXQgaW5zdGFuY2VvZiBSZXF1ZXN0KSB7XG4gICAgICBpZiAoaW5wdXQuYm9keVVzZWQpIHtcbiAgICAgICAgdGhyb3cgbmV3IFR5cGVFcnJvcignQWxyZWFkeSByZWFkJylcbiAgICAgIH1cbiAgICAgIHRoaXMudXJsID0gaW5wdXQudXJsO1xuICAgICAgdGhpcy5jcmVkZW50aWFscyA9IGlucHV0LmNyZWRlbnRpYWxzO1xuICAgICAgaWYgKCFvcHRpb25zLmhlYWRlcnMpIHtcbiAgICAgICAgdGhpcy5oZWFkZXJzID0gbmV3IEhlYWRlcnMoaW5wdXQuaGVhZGVycyk7XG4gICAgICB9XG4gICAgICB0aGlzLm1ldGhvZCA9IGlucHV0Lm1ldGhvZDtcbiAgICAgIHRoaXMubW9kZSA9IGlucHV0Lm1vZGU7XG4gICAgICB0aGlzLnNpZ25hbCA9IGlucHV0LnNpZ25hbDtcbiAgICAgIGlmICghYm9keSAmJiBpbnB1dC5fYm9keUluaXQgIT0gbnVsbCkge1xuICAgICAgICBib2R5ID0gaW5wdXQuX2JvZHlJbml0O1xuICAgICAgICBpbnB1dC5ib2R5VXNlZCA9IHRydWU7XG4gICAgICB9XG4gICAgfSBlbHNlIHtcbiAgICAgIHRoaXMudXJsID0gU3RyaW5nKGlucHV0KTtcbiAgICB9XG5cbiAgICB0aGlzLmNyZWRlbnRpYWxzID0gb3B0aW9ucy5jcmVkZW50aWFscyB8fCB0aGlzLmNyZWRlbnRpYWxzIHx8ICdzYW1lLW9yaWdpbic7XG4gICAgaWYgKG9wdGlvbnMuaGVhZGVycyB8fCAhdGhpcy5oZWFkZXJzKSB7XG4gICAgICB0aGlzLmhlYWRlcnMgPSBuZXcgSGVhZGVycyhvcHRpb25zLmhlYWRlcnMpO1xuICAgIH1cbiAgICB0aGlzLm1ldGhvZCA9IG5vcm1hbGl6ZU1ldGhvZChvcHRpb25zLm1ldGhvZCB8fCB0aGlzLm1ldGhvZCB8fCAnR0VUJyk7XG4gICAgdGhpcy5tb2RlID0gb3B0aW9ucy5tb2RlIHx8IHRoaXMubW9kZSB8fCBudWxsO1xuICAgIHRoaXMuc2lnbmFsID0gb3B0aW9ucy5zaWduYWwgfHwgdGhpcy5zaWduYWw7XG4gICAgdGhpcy5yZWZlcnJlciA9IG51bGw7XG5cbiAgICBpZiAoKHRoaXMubWV0aG9kID09PSAnR0VUJyB8fCB0aGlzLm1ldGhvZCA9PT0gJ0hFQUQnKSAmJiBib2R5KSB7XG4gICAgICB0aHJvdyBuZXcgVHlwZUVycm9yKCdCb2R5IG5vdCBhbGxvd2VkIGZvciBHRVQgb3IgSEVBRCByZXF1ZXN0cycpXG4gICAgfVxuICAgIHRoaXMuX2luaXRCb2R5KGJvZHkpO1xuICB9XG5cbiAgUmVxdWVzdC5wcm90b3R5cGUuY2xvbmUgPSBmdW5jdGlvbigpIHtcbiAgICByZXR1cm4gbmV3IFJlcXVlc3QodGhpcywge2JvZHk6IHRoaXMuX2JvZHlJbml0fSlcbiAgfTtcblxuICBmdW5jdGlvbiBkZWNvZGUoYm9keSkge1xuICAgIHZhciBmb3JtID0gbmV3IEZvcm1EYXRhKCk7XG4gICAgYm9keVxuICAgICAgLnRyaW0oKVxuICAgICAgLnNwbGl0KCcmJylcbiAgICAgIC5mb3JFYWNoKGZ1bmN0aW9uKGJ5dGVzKSB7XG4gICAgICAgIGlmIChieXRlcykge1xuICAgICAgICAgIHZhciBzcGxpdCA9IGJ5dGVzLnNwbGl0KCc9Jyk7XG4gICAgICAgICAgdmFyIG5hbWUgPSBzcGxpdC5zaGlmdCgpLnJlcGxhY2UoL1xcKy9nLCAnICcpO1xuICAgICAgICAgIHZhciB2YWx1ZSA9IHNwbGl0LmpvaW4oJz0nKS5yZXBsYWNlKC9cXCsvZywgJyAnKTtcbiAgICAgICAgICBmb3JtLmFwcGVuZChkZWNvZGVVUklDb21wb25lbnQobmFtZSksIGRlY29kZVVSSUNvbXBvbmVudCh2YWx1ZSkpO1xuICAgICAgICB9XG4gICAgICB9KTtcbiAgICByZXR1cm4gZm9ybVxuICB9XG5cbiAgZnVuY3Rpb24gcGFyc2VIZWFkZXJzKHJhd0hlYWRlcnMpIHtcbiAgICB2YXIgaGVhZGVycyA9IG5ldyBIZWFkZXJzKCk7XG4gICAgLy8gUmVwbGFjZSBpbnN0YW5jZXMgb2YgXFxyXFxuIGFuZCBcXG4gZm9sbG93ZWQgYnkgYXQgbGVhc3Qgb25lIHNwYWNlIG9yIGhvcml6b250YWwgdGFiIHdpdGggYSBzcGFjZVxuICAgIC8vIGh0dHBzOi8vdG9vbHMuaWV0Zi5vcmcvaHRtbC9yZmM3MjMwI3NlY3Rpb24tMy4yXG4gICAgdmFyIHByZVByb2Nlc3NlZEhlYWRlcnMgPSByYXdIZWFkZXJzLnJlcGxhY2UoL1xccj9cXG5bXFx0IF0rL2csICcgJyk7XG4gICAgcHJlUHJvY2Vzc2VkSGVhZGVycy5zcGxpdCgvXFxyP1xcbi8pLmZvckVhY2goZnVuY3Rpb24obGluZSkge1xuICAgICAgdmFyIHBhcnRzID0gbGluZS5zcGxpdCgnOicpO1xuICAgICAgdmFyIGtleSA9IHBhcnRzLnNoaWZ0KCkudHJpbSgpO1xuICAgICAgaWYgKGtleSkge1xuICAgICAgICB2YXIgdmFsdWUgPSBwYXJ0cy5qb2luKCc6JykudHJpbSgpO1xuICAgICAgICBoZWFkZXJzLmFwcGVuZChrZXksIHZhbHVlKTtcbiAgICAgIH1cbiAgICB9KTtcbiAgICByZXR1cm4gaGVhZGVyc1xuICB9XG5cbiAgQm9keS5jYWxsKFJlcXVlc3QucHJvdG90eXBlKTtcblxuICBmdW5jdGlvbiBSZXNwb25zZShib2R5SW5pdCwgb3B0aW9ucykge1xuICAgIGlmICghb3B0aW9ucykge1xuICAgICAgb3B0aW9ucyA9IHt9O1xuICAgIH1cblxuICAgIHRoaXMudHlwZSA9ICdkZWZhdWx0JztcbiAgICB0aGlzLnN0YXR1cyA9IG9wdGlvbnMuc3RhdHVzID09PSB1bmRlZmluZWQgPyAyMDAgOiBvcHRpb25zLnN0YXR1cztcbiAgICB0aGlzLm9rID0gdGhpcy5zdGF0dXMgPj0gMjAwICYmIHRoaXMuc3RhdHVzIDwgMzAwO1xuICAgIHRoaXMuc3RhdHVzVGV4dCA9ICdzdGF0dXNUZXh0JyBpbiBvcHRpb25zID8gb3B0aW9ucy5zdGF0dXNUZXh0IDogJ09LJztcbiAgICB0aGlzLmhlYWRlcnMgPSBuZXcgSGVhZGVycyhvcHRpb25zLmhlYWRlcnMpO1xuICAgIHRoaXMudXJsID0gb3B0aW9ucy51cmwgfHwgJyc7XG4gICAgdGhpcy5faW5pdEJvZHkoYm9keUluaXQpO1xuICB9XG5cbiAgQm9keS5jYWxsKFJlc3BvbnNlLnByb3RvdHlwZSk7XG5cbiAgUmVzcG9uc2UucHJvdG90eXBlLmNsb25lID0gZnVuY3Rpb24oKSB7XG4gICAgcmV0dXJuIG5ldyBSZXNwb25zZSh0aGlzLl9ib2R5SW5pdCwge1xuICAgICAgc3RhdHVzOiB0aGlzLnN0YXR1cyxcbiAgICAgIHN0YXR1c1RleHQ6IHRoaXMuc3RhdHVzVGV4dCxcbiAgICAgIGhlYWRlcnM6IG5ldyBIZWFkZXJzKHRoaXMuaGVhZGVycyksXG4gICAgICB1cmw6IHRoaXMudXJsXG4gICAgfSlcbiAgfTtcblxuICBSZXNwb25zZS5lcnJvciA9IGZ1bmN0aW9uKCkge1xuICAgIHZhciByZXNwb25zZSA9IG5ldyBSZXNwb25zZShudWxsLCB7c3RhdHVzOiAwLCBzdGF0dXNUZXh0OiAnJ30pO1xuICAgIHJlc3BvbnNlLnR5cGUgPSAnZXJyb3InO1xuICAgIHJldHVybiByZXNwb25zZVxuICB9O1xuXG4gIHZhciByZWRpcmVjdFN0YXR1c2VzID0gWzMwMSwgMzAyLCAzMDMsIDMwNywgMzA4XTtcblxuICBSZXNwb25zZS5yZWRpcmVjdCA9IGZ1bmN0aW9uKHVybCwgc3RhdHVzKSB7XG4gICAgaWYgKHJlZGlyZWN0U3RhdHVzZXMuaW5kZXhPZihzdGF0dXMpID09PSAtMSkge1xuICAgICAgdGhyb3cgbmV3IFJhbmdlRXJyb3IoJ0ludmFsaWQgc3RhdHVzIGNvZGUnKVxuICAgIH1cblxuICAgIHJldHVybiBuZXcgUmVzcG9uc2UobnVsbCwge3N0YXR1czogc3RhdHVzLCBoZWFkZXJzOiB7bG9jYXRpb246IHVybH19KVxuICB9O1xuXG4gIGV4cG9ydHMuRE9NRXhjZXB0aW9uID0gc2VsZi5ET01FeGNlcHRpb247XG4gIHRyeSB7XG4gICAgbmV3IGV4cG9ydHMuRE9NRXhjZXB0aW9uKCk7XG4gIH0gY2F0Y2ggKGVycikge1xuICAgIGV4cG9ydHMuRE9NRXhjZXB0aW9uID0gZnVuY3Rpb24obWVzc2FnZSwgbmFtZSkge1xuICAgICAgdGhpcy5tZXNzYWdlID0gbWVzc2FnZTtcbiAgICAgIHRoaXMubmFtZSA9IG5hbWU7XG4gICAgICB2YXIgZXJyb3IgPSBFcnJvcihtZXNzYWdlKTtcbiAgICAgIHRoaXMuc3RhY2sgPSBlcnJvci5zdGFjaztcbiAgICB9O1xuICAgIGV4cG9ydHMuRE9NRXhjZXB0aW9uLnByb3RvdHlwZSA9IE9iamVjdC5jcmVhdGUoRXJyb3IucHJvdG90eXBlKTtcbiAgICBleHBvcnRzLkRPTUV4Y2VwdGlvbi5wcm90b3R5cGUuY29uc3RydWN0b3IgPSBleHBvcnRzLkRPTUV4Y2VwdGlvbjtcbiAgfVxuXG4gIGZ1bmN0aW9uIGZldGNoKGlucHV0LCBpbml0KSB7XG4gICAgcmV0dXJuIG5ldyBQcm9taXNlKGZ1bmN0aW9uKHJlc29sdmUsIHJlamVjdCkge1xuICAgICAgdmFyIHJlcXVlc3QgPSBuZXcgUmVxdWVzdChpbnB1dCwgaW5pdCk7XG5cbiAgICAgIGlmIChyZXF1ZXN0LnNpZ25hbCAmJiByZXF1ZXN0LnNpZ25hbC5hYm9ydGVkKSB7XG4gICAgICAgIHJldHVybiByZWplY3QobmV3IGV4cG9ydHMuRE9NRXhjZXB0aW9uKCdBYm9ydGVkJywgJ0Fib3J0RXJyb3InKSlcbiAgICAgIH1cblxuICAgICAgdmFyIHhociA9IG5ldyBYTUxIdHRwUmVxdWVzdCgpO1xuXG4gICAgICBmdW5jdGlvbiBhYm9ydFhocigpIHtcbiAgICAgICAgeGhyLmFib3J0KCk7XG4gICAgICB9XG5cbiAgICAgIHhoci5vbmxvYWQgPSBmdW5jdGlvbigpIHtcbiAgICAgICAgdmFyIG9wdGlvbnMgPSB7XG4gICAgICAgICAgc3RhdHVzOiB4aHIuc3RhdHVzLFxuICAgICAgICAgIHN0YXR1c1RleHQ6IHhoci5zdGF0dXNUZXh0LFxuICAgICAgICAgIGhlYWRlcnM6IHBhcnNlSGVhZGVycyh4aHIuZ2V0QWxsUmVzcG9uc2VIZWFkZXJzKCkgfHwgJycpXG4gICAgICAgIH07XG4gICAgICAgIG9wdGlvbnMudXJsID0gJ3Jlc3BvbnNlVVJMJyBpbiB4aHIgPyB4aHIucmVzcG9uc2VVUkwgOiBvcHRpb25zLmhlYWRlcnMuZ2V0KCdYLVJlcXVlc3QtVVJMJyk7XG4gICAgICAgIHZhciBib2R5ID0gJ3Jlc3BvbnNlJyBpbiB4aHIgPyB4aHIucmVzcG9uc2UgOiB4aHIucmVzcG9uc2VUZXh0O1xuICAgICAgICByZXNvbHZlKG5ldyBSZXNwb25zZShib2R5LCBvcHRpb25zKSk7XG4gICAgICB9O1xuXG4gICAgICB4aHIub25lcnJvciA9IGZ1bmN0aW9uKCkge1xuICAgICAgICByZWplY3QobmV3IFR5cGVFcnJvcignTmV0d29yayByZXF1ZXN0IGZhaWxlZCcpKTtcbiAgICAgIH07XG5cbiAgICAgIHhoci5vbnRpbWVvdXQgPSBmdW5jdGlvbigpIHtcbiAgICAgICAgcmVqZWN0KG5ldyBUeXBlRXJyb3IoJ05ldHdvcmsgcmVxdWVzdCBmYWlsZWQnKSk7XG4gICAgICB9O1xuXG4gICAgICB4aHIub25hYm9ydCA9IGZ1bmN0aW9uKCkge1xuICAgICAgICByZWplY3QobmV3IGV4cG9ydHMuRE9NRXhjZXB0aW9uKCdBYm9ydGVkJywgJ0Fib3J0RXJyb3InKSk7XG4gICAgICB9O1xuXG4gICAgICB4aHIub3BlbihyZXF1ZXN0Lm1ldGhvZCwgcmVxdWVzdC51cmwsIHRydWUpO1xuXG4gICAgICBpZiAocmVxdWVzdC5jcmVkZW50aWFscyA9PT0gJ2luY2x1ZGUnKSB7XG4gICAgICAgIHhoci53aXRoQ3JlZGVudGlhbHMgPSB0cnVlO1xuICAgICAgfSBlbHNlIGlmIChyZXF1ZXN0LmNyZWRlbnRpYWxzID09PSAnb21pdCcpIHtcbiAgICAgICAgeGhyLndpdGhDcmVkZW50aWFscyA9IGZhbHNlO1xuICAgICAgfVxuXG4gICAgICBpZiAoJ3Jlc3BvbnNlVHlwZScgaW4geGhyICYmIHN1cHBvcnQuYmxvYikge1xuICAgICAgICB4aHIucmVzcG9uc2VUeXBlID0gJ2Jsb2InO1xuICAgICAgfVxuXG4gICAgICByZXF1ZXN0LmhlYWRlcnMuZm9yRWFjaChmdW5jdGlvbih2YWx1ZSwgbmFtZSkge1xuICAgICAgICB4aHIuc2V0UmVxdWVzdEhlYWRlcihuYW1lLCB2YWx1ZSk7XG4gICAgICB9KTtcblxuICAgICAgaWYgKHJlcXVlc3Quc2lnbmFsKSB7XG4gICAgICAgIHJlcXVlc3Quc2lnbmFsLmFkZEV2ZW50TGlzdGVuZXIoJ2Fib3J0JywgYWJvcnRYaHIpO1xuXG4gICAgICAgIHhoci5vbnJlYWR5c3RhdGVjaGFuZ2UgPSBmdW5jdGlvbigpIHtcbiAgICAgICAgICAvLyBET05FIChzdWNjZXNzIG9yIGZhaWx1cmUpXG4gICAgICAgICAgaWYgKHhoci5yZWFkeVN0YXRlID09PSA0KSB7XG4gICAgICAgICAgICByZXF1ZXN0LnNpZ25hbC5yZW1vdmVFdmVudExpc3RlbmVyKCdhYm9ydCcsIGFib3J0WGhyKTtcbiAgICAgICAgICB9XG4gICAgICAgIH07XG4gICAgICB9XG5cbiAgICAgIHhoci5zZW5kKHR5cGVvZiByZXF1ZXN0Ll9ib2R5SW5pdCA9PT0gJ3VuZGVmaW5lZCcgPyBudWxsIDogcmVxdWVzdC5fYm9keUluaXQpO1xuICAgIH0pXG4gIH1cblxuICBmZXRjaC5wb2x5ZmlsbCA9IHRydWU7XG5cbiAgaWYgKCFzZWxmLmZldGNoKSB7XG4gICAgc2VsZi5mZXRjaCA9IGZldGNoO1xuICAgIHNlbGYuSGVhZGVycyA9IEhlYWRlcnM7XG4gICAgc2VsZi5SZXF1ZXN0ID0gUmVxdWVzdDtcbiAgICBzZWxmLlJlc3BvbnNlID0gUmVzcG9uc2U7XG4gIH1cblxuICBleHBvcnRzLkhlYWRlcnMgPSBIZWFkZXJzO1xuICBleHBvcnRzLlJlcXVlc3QgPSBSZXF1ZXN0O1xuICBleHBvcnRzLlJlc3BvbnNlID0gUmVzcG9uc2U7XG4gIGV4cG9ydHMuZmV0Y2ggPSBmZXRjaDtcblxuICBPYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgJ19fZXNNb2R1bGUnLCB7IHZhbHVlOiB0cnVlIH0pO1xuXG59KSkpO1xuIiwiXCJ1c2Ugc3RyaWN0XCI7XG5PYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgXCJfX2VzTW9kdWxlXCIsIHsgdmFsdWU6IHRydWUgfSk7XG5leHBvcnRzLkFGcmFtZSA9IGV4cG9ydHMuQW5pbWF0aW9uUXVldWUgPSB2b2lkIDA7XG5jb25zdCByYWZQb2x5ZmlsbCA9IHJlcXVpcmUoXCIuL3JhZi1wb2x5ZmlsbFwiKTtcbmNsYXNzIEFuaW1hdGlvblF1ZXVlIHtcbiAgICBjb25zdHJ1Y3RvcigpIHtcbiAgICAgICAgdGhpcy5za2lwID0gZmFsc2U7XG4gICAgICAgIHRoaXMuYmluZGVkID0gZmFsc2U7XG4gICAgICAgIHRoaXMucmVxdWVzdEFuaW1hdGlvbklEID0gLTE7XG4gICAgICAgIHRoaXMuZnJhbWVzID0gbmV3IEFycmF5KCk7XG4gICAgICAgIHRoaXMuYmluZEN5Y2xlID0gdGhpcy5jeWNsZS5iaW5kKHRoaXMpO1xuICAgICAgICB0aGlzLnJhZlByb3ZpZGVyID0gcmFmUG9seWZpbGwuR2V0UkFGKCk7XG4gICAgfVxuICAgIG5ldygpIHtcbiAgICAgICAgY29uc3QgbmV3RnJhbWUgPSBuZXcgQUZyYW1lKHRoaXMuZnJhbWVzLmxlbmd0aCwgdGhpcyk7XG4gICAgICAgIHRoaXMuZnJhbWVzLnB1c2gobmV3RnJhbWUpO1xuICAgICAgICB0aGlzLmJpbmQoKTtcbiAgICAgICAgcmV0dXJuIG5ld0ZyYW1lO1xuICAgIH1cbiAgICBhZGQoZikge1xuICAgICAgICBmLnF1ZXVlSW5kZXggPSB0aGlzLmZyYW1lcy5sZW5ndGg7XG4gICAgICAgIGYucXVldWUgPSB0aGlzO1xuICAgICAgICB0aGlzLmZyYW1lcy5wdXNoKGYpO1xuICAgICAgICB0aGlzLmJpbmQoKTtcbiAgICB9XG4gICAgcmVzdW1lKCkge1xuICAgICAgICB0aGlzLnNraXAgPSBmYWxzZTtcbiAgICAgICAgdGhpcy5iaW5kKCk7XG4gICAgfVxuICAgIHBhdXNlKCkge1xuICAgICAgICB0aGlzLnNraXAgPSB0cnVlO1xuICAgIH1cbiAgICB1bmJpbmQoKSB7XG4gICAgICAgIGlmICghdGhpcy5iaW5kZWQpIHtcbiAgICAgICAgICAgIHJldHVybiBudWxsO1xuICAgICAgICB9XG4gICAgICAgIHRoaXMucmFmUHJvdmlkZXIuY2FuY2VsQW5pbWF0aW9uRnJhbWUodGhpcy5yZXF1ZXN0QW5pbWF0aW9uSUQpO1xuICAgIH1cbiAgICBiaW5kKCkge1xuICAgICAgICBpZiAodGhpcy5iaW5kZWQpXG4gICAgICAgICAgICByZXR1cm4gbnVsbDtcbiAgICAgICAgdGhpcy5yZXF1ZXN0QW5pbWF0aW9uSUQgPSB0aGlzLnJhZlByb3ZpZGVyLnJlcXVlc3RBbmltYXRpb25GcmFtZSh0aGlzLmJpbmRDeWNsZSwgbnVsbCk7XG4gICAgICAgIHRoaXMuYmluZGVkID0gdHJ1ZTtcbiAgICB9XG4gICAgY3ljbGUobXMpIHtcbiAgICAgICAgaWYgKHRoaXMuZnJhbWVzLmxlbmd0aCA9PT0gMCkge1xuICAgICAgICAgICAgdGhpcy5iaW5kZWQgPSBmYWxzZTtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICB0aGlzLmZyYW1lcy5mb3JFYWNoKGZ1bmN0aW9uIChmKSB7XG4gICAgICAgICAgICBpZiAoIWYucGF1c2VkKCkpIHtcbiAgICAgICAgICAgICAgICBmLmFuaW1hdGUobXMpO1xuICAgICAgICAgICAgfVxuICAgICAgICB9KTtcbiAgICAgICAgdGhpcy5iaW5kKCk7XG4gICAgfVxufVxuZXhwb3J0cy5BbmltYXRpb25RdWV1ZSA9IEFuaW1hdGlvblF1ZXVlO1xuY2xhc3MgQUZyYW1lIHtcbiAgICBjb25zdHJ1Y3RvcihpbmRleCwgcXVldWUpIHtcbiAgICAgICAgdGhpcy5za2lwID0gZmFsc2U7XG4gICAgICAgIHRoaXMucXVldWUgPSBxdWV1ZTtcbiAgICAgICAgdGhpcy5xdWV1ZUluZGV4ID0gaW5kZXg7XG4gICAgICAgIHRoaXMuY2FsbGJhY2tzID0gbmV3IEFycmF5KCk7XG4gICAgfVxuICAgIGFkZChjYWxsYmFjaykge1xuICAgICAgICB0aGlzLmNhbGxiYWNrcy5wdXNoKGNhbGxiYWNrKTtcbiAgICB9XG4gICAgY2xlYXIoKSB7XG4gICAgICAgIHRoaXMuY2FsbGJhY2tzLmxlbmd0aCA9IDA7XG4gICAgfVxuICAgIHBhdXNlZCgpIHtcbiAgICAgICAgcmV0dXJuIHRoaXMuc2tpcDtcbiAgICB9XG4gICAgcGF1c2UoKSB7XG4gICAgICAgIHRoaXMuc2tpcCA9IHRydWU7XG4gICAgfVxuICAgIHN0b3AoKSB7XG4gICAgICAgIHRoaXMucGF1c2UoKTtcbiAgICAgICAgaWYgKHRoaXMucXVldWVJbmRleCA9PT0gLTEpIHtcbiAgICAgICAgICAgIHJldHVybiBudWxsO1xuICAgICAgICB9XG4gICAgICAgIGlmICh0aGlzLnF1ZXVlLmZyYW1lcy5sZW5ndGggPT0gMCkge1xuICAgICAgICAgICAgdGhpcy5xdWV1ZSA9IHVuZGVmaW5lZDtcbiAgICAgICAgICAgIHRoaXMucXVldWVJbmRleCA9IC0xO1xuICAgICAgICAgICAgcmV0dXJuIG51bGw7XG4gICAgICAgIH1cbiAgICAgICAgY29uc3QgdG90YWwgPSB0aGlzLnF1ZXVlLmZyYW1lcy5sZW5ndGg7XG4gICAgICAgIGlmICh0b3RhbCA9PSAxKSB7XG4gICAgICAgICAgICB0aGlzLnF1ZXVlLmZyYW1lcy5wb3AoKTtcbiAgICAgICAgICAgIHRoaXMucXVldWUgPSB1bmRlZmluZWQ7XG4gICAgICAgICAgICB0aGlzLnF1ZXVlSW5kZXggPSAtMTtcbiAgICAgICAgICAgIHJldHVybiBudWxsO1xuICAgICAgICB9XG4gICAgICAgIHRoaXMucXVldWUuZnJhbWVzW3RoaXMucXVldWVJbmRleF0gPSB0aGlzLnF1ZXVlLmZyYW1lc1t0b3RhbCAtIDFdO1xuICAgICAgICB0aGlzLnF1ZXVlLmZyYW1lcy5sZW5ndGggPSB0b3RhbCAtIDE7XG4gICAgICAgIHRoaXMucXVldWUgPSB1bmRlZmluZWQ7XG4gICAgICAgIHRoaXMucXVldWVJbmRleCA9IC0xO1xuICAgIH1cbiAgICBhbmltYXRlKHRzKSB7XG4gICAgICAgIGZvciAobGV0IGluZGV4IGluIHRoaXMuY2FsbGJhY2tzKSB7XG4gICAgICAgICAgICBjb25zdCBjYWxsYmFjayA9IHRoaXMuY2FsbGJhY2tzW2luZGV4XTtcbiAgICAgICAgICAgIGNhbGxiYWNrKHRzKTtcbiAgICAgICAgfVxuICAgIH1cbn1cbmV4cG9ydHMuQUZyYW1lID0gQUZyYW1lO1xuY2xhc3MgQ2hhbmdlTWFuYWdlciB7XG4gICAgY29uc3RydWN0b3IocXVldWUpIHtcbiAgICAgICAgdGhpcy5yZWFkcyA9IG5ldyBBcnJheSgpO1xuICAgICAgICB0aGlzLndyaXRlcyA9IG5ldyBBcnJheSgpO1xuICAgICAgICB0aGlzLnJlYWRTdGF0ZSA9IGZhbHNlO1xuICAgICAgICB0aGlzLmluUmVhZENhbGwgPSBmYWxzZTtcbiAgICAgICAgdGhpcy5pbldyaXRlQ2FsbCA9IGZhbHNlO1xuICAgICAgICB0aGlzLnNjaGVkdWxlZCA9IGZhbHNlO1xuICAgICAgICB0aGlzLmZyYW1lID0gcXVldWUubmV3KCk7XG4gICAgfVxuICAgIHN0YXRpYyBkcmFpblRhc2tzKHEsIHdyYXBwZXIpIHtcbiAgICAgICAgbGV0IHRhc2sgPSBxLnNoaWZ0KCk7XG4gICAgICAgIHdoaWxlICh0YXNrKSB7XG4gICAgICAgICAgICBpZiAod3JhcHBlciAhPT0gbnVsbCkge1xuICAgICAgICAgICAgICAgIHdyYXBwZXIodGFzayk7XG4gICAgICAgICAgICAgICAgdGFzayA9IHEuc2hpZnQoKTtcbiAgICAgICAgICAgICAgICBjb250aW51ZTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIHRhc2soKTtcbiAgICAgICAgICAgIHRhc2sgPSBxLnNoaWZ0KCk7XG4gICAgICAgIH1cbiAgICB9XG4gICAgbXV0YXRlKGZuKSB7XG4gICAgICAgIHRoaXMud3JpdGVzLnB1c2goZm4pO1xuICAgICAgICB0aGlzLl9zY2hlZHVsZSgpO1xuICAgIH1cbiAgICByZWFkKGZuKSB7XG4gICAgICAgIHRoaXMucmVhZHMucHVzaChmbik7XG4gICAgICAgIHRoaXMuX3NjaGVkdWxlKCk7XG4gICAgfVxuICAgIF9zY2hlZHVsZSgpIHtcbiAgICAgICAgaWYgKHRoaXMuc2NoZWR1bGVkKSB7XG4gICAgICAgICAgICByZXR1cm47XG4gICAgICAgIH1cbiAgICAgICAgdGhpcy5zY2hlZHVsZWQgPSB0cnVlO1xuICAgICAgICB0aGlzLmZyYW1lLmFkZCh0aGlzLl9ydW5UYXNrcy5iaW5kKHRoaXMpKTtcbiAgICB9XG4gICAgX3J1blRhc2tzKCkge1xuICAgICAgICBjb25zdCByZWFkRXJyb3IgPSB0aGlzLl9ydW5SZWFkcygpO1xuICAgICAgICBpZiAocmVhZEVycm9yICE9PSBudWxsICYmIHJlYWRFcnJvciAhPT0gdW5kZWZpbmVkKSB7XG4gICAgICAgICAgICB0aGlzLnNjaGVkdWxlZCA9IGZhbHNlO1xuICAgICAgICAgICAgdGhpcy5fc2NoZWR1bGUoKTtcbiAgICAgICAgICAgIHRocm93IHJlYWRFcnJvcjtcbiAgICAgICAgfVxuICAgICAgICBjb25zdCB3cml0ZUVycm9yID0gdGhpcy5fcnVuV3JpdGVzKCk7XG4gICAgICAgIGlmICh3cml0ZUVycm9yICE9PSBudWxsICYmIHdyaXRlRXJyb3IgIT09IHVuZGVmaW5lZCkge1xuICAgICAgICAgICAgdGhpcy5zY2hlZHVsZWQgPSBmYWxzZTtcbiAgICAgICAgICAgIHRoaXMuX3NjaGVkdWxlKCk7XG4gICAgICAgICAgICB0aHJvdyB3cml0ZUVycm9yO1xuICAgICAgICB9XG4gICAgICAgIGlmICh0aGlzLnJlYWRzLmxlbmd0aCA+IDAgfHwgdGhpcy53cml0ZXMubGVuZ3RoID4gMCkge1xuICAgICAgICAgICAgdGhpcy5zY2hlZHVsZWQgPSBmYWxzZTtcbiAgICAgICAgICAgIHRoaXMuX3NjaGVkdWxlKCk7XG4gICAgICAgICAgICByZXR1cm47XG4gICAgICAgIH1cbiAgICAgICAgdGhpcy5zY2hlZHVsZWQgPSBmYWxzZTtcbiAgICB9XG4gICAgX3J1blJlYWRzKCkge1xuICAgICAgICB0cnkge1xuICAgICAgICAgICAgQ2hhbmdlTWFuYWdlci5kcmFpblRhc2tzKHRoaXMucmVhZHMsIHRoaXMuX2V4ZWNSZWFkcy5iaW5kKHRoaXMpKTtcbiAgICAgICAgfVxuICAgICAgICBjYXRjaCAoZSkge1xuICAgICAgICAgICAgcmV0dXJuIGU7XG4gICAgICAgIH1cbiAgICAgICAgcmV0dXJuIG51bGw7XG4gICAgfVxuICAgIF9leGVjUmVhZHModGFzaykge1xuICAgICAgICB0aGlzLmluUmVhZENhbGwgPSB0cnVlO1xuICAgICAgICB0YXNrKCk7XG4gICAgICAgIHRoaXMuaW5SZWFkQ2FsbCA9IGZhbHNlO1xuICAgIH1cbiAgICBfcnVuV3JpdGVzKCkge1xuICAgICAgICB0cnkge1xuICAgICAgICAgICAgQ2hhbmdlTWFuYWdlci5kcmFpblRhc2tzKHRoaXMud3JpdGVzLCB0aGlzLl9leGVjV3JpdGUuYmluZCh0aGlzKSk7XG4gICAgICAgIH1cbiAgICAgICAgY2F0Y2ggKGUpIHtcbiAgICAgICAgICAgIHJldHVybiBlO1xuICAgICAgICB9XG4gICAgICAgIHJldHVybiBudWxsO1xuICAgIH1cbiAgICBfZXhlY1dyaXRlKHRhc2spIHtcbiAgICAgICAgdGhpcy5pbldyaXRlQ2FsbCA9IHRydWU7XG4gICAgICAgIHRhc2soKTtcbiAgICAgICAgdGhpcy5pbldyaXRlQ2FsbCA9IGZhbHNlO1xuICAgIH1cbn1cbi8vIyBzb3VyY2VNYXBwaW5nVVJMPWFuaW1lLmpzLm1hcCIsIlwidXNlIHN0cmljdFwiO1xuT2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsIFwiX19lc01vZHVsZVwiLCB7IHZhbHVlOiB0cnVlIH0pO1xuZXhwb3J0cy5tb3VudFRvID0gdm9pZCAwO1xuY29uc3QgcmFmUG9seWZpbGwgPSByZXF1aXJlKFwiLi9yYWYtcG9seWZpbGxcIik7XG5jb25zdCBBbmltYXRpb24gPSByZXF1aXJlKFwiLi9hbmltZVwiKTtcbmNvbnN0IG5hbWVzcGFjZSA9IE9iamVjdC5hc3NpZ24oT2JqZWN0LmFzc2lnbih7fSwgcmFmUG9seWZpbGwpLCBBbmltYXRpb24pO1xuZnVuY3Rpb24gbW91bnRUbyhwYXJlbnQpIHtcbiAgICBwYXJlbnQuYW5pbWF0aW9ucyA9IG5hbWVzcGFjZTtcbn1cbmV4cG9ydHMubW91bnRUbyA9IG1vdW50VG87XG5leHBvcnRzLmRlZmF1bHQgPSBuYW1lc3BhY2U7XG4vLyMgc291cmNlTWFwcGluZ1VSTD1pbmRleC5qcy5tYXAiLCJcInVzZSBzdHJpY3RcIjtcbk9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBcIl9fZXNNb2R1bGVcIiwgeyB2YWx1ZTogdHJ1ZSB9KTtcbmV4cG9ydHMuR2V0UkFGID0gdm9pZCAwO1xuY29uc3Qgbm93ID0gKGZ1bmN0aW9uICgpIHtcbiAgICBpZiAoc2VsZi5oYXNPd25Qcm9wZXJ0eSgncGVyZm9ybWFuY2UnKSkge1xuICAgICAgICByZXR1cm4gKChwZXJmb3JtYW5jZS5ub3cgPyBwZXJmb3JtYW5jZS5ub3cuYmluZChwZXJmb3JtYW5jZSkgOiBudWxsKSB8fFxuICAgICAgICAgICAgKHBlcmZvcm1hbmNlLm1vek5vdyA/IHBlcmZvcm1hbmNlLm1vek5vdy5iaW5kKHBlcmZvcm1hbmNlKSA6IG51bGwpIHx8XG4gICAgICAgICAgICAocGVyZm9ybWFuY2UubXNOb3cgPyBwZXJmb3JtYW5jZS5tc05vdy5iaW5kKHBlcmZvcm1hbmNlKSA6IG51bGwpIHx8XG4gICAgICAgICAgICAocGVyZm9ybWFuY2Uub05vdyA/IHBlcmZvcm1hbmNlLm9Ob3cuYmluZChwZXJmb3JtYW5jZSkgOiBudWxsKSB8fFxuICAgICAgICAgICAgKHBlcmZvcm1hbmNlLndlYmtpdE5vdyA/IHBlcmZvcm1hbmNlLndlYmtpdE5vdy5iaW5kKHBlcmZvcm1hbmNlKSA6IG51bGwpIHx8XG4gICAgICAgICAgICBEYXRlLm5vdy5iaW5kKERhdGUpKTtcbiAgICB9XG4gICAgcmV0dXJuIERhdGUubm93LmJpbmQoRGF0ZSk7XG59KSgpO1xuY29uc3QgZnJhbWVSYXRlID0gMTAwMCAvIDYwO1xuY29uc3QgdmVuZG9ycyA9IFsnbXMnLCAnbW96JywgJ3dlYmtpdCcsICdvJ107XG5mdW5jdGlvbiBHZXRSQUYoKSB7XG4gICAgbGV0IGxhc3RUaW1lID0gMDtcbiAgICBjb25zdCBtb2QgPSB7fTtcbiAgICBmb3IgKHZhciB4ID0gMDsgeCA8IHZlbmRvcnMubGVuZ3RoICYmICF3aW5kb3cucmVxdWVzdEFuaW1hdGlvbkZyYW1lOyArK3gpIHtcbiAgICAgICAgbW9kLnJlcXVlc3RBbmltYXRpb25GcmFtZSA9IHdpbmRvd1t2ZW5kb3JzW3hdICsgJ1JlcXVlc3RBbmltYXRpb25GcmFtZSddO1xuICAgICAgICBtb2QuY2FuY2VsQW5pbWF0aW9uRnJhbWUgPVxuICAgICAgICAgICAgd2luZG93W3ZlbmRvcnNbeF0gKyAnQ2FuY2VsQW5pbWF0aW9uRnJhbWUnXSB8fCB3aW5kb3dbdmVuZG9yc1t4XSArICdSZXF1ZXN0Q2FuY2VsQW5pbWF0aW9uRnJhbWUnXTtcbiAgICB9XG4gICAgaWYgKCFtb2QucmVxdWVzdEFuaW1hdGlvbkZyYW1lIHx8ICFtb2QuY2FuY2VsQW5pbWF0aW9uRnJhbWUpXG4gICAgICAgIG1vZC5yZXF1ZXN0QW5pbWF0aW9uRnJhbWUgPSBmdW5jdGlvbiAoY2FsbGJhY2ssIGVsZW1lbnQpIHtcbiAgICAgICAgICAgIGNvbnN0IGN1cnJUaW1lID0gbm93KCk7XG4gICAgICAgICAgICBjb25zdCB0aW1lVG9DYWxsID0gTWF0aC5tYXgoMCwgZnJhbWVSYXRlIC0gKGN1cnJUaW1lIC0gbGFzdFRpbWUpKTtcbiAgICAgICAgICAgIGNvbnN0IGlkID0gc2VsZi5zZXRUaW1lb3V0KGZ1bmN0aW9uICgpIHtcbiAgICAgICAgICAgICAgICB0cnkge1xuICAgICAgICAgICAgICAgICAgICBjYWxsYmFjayhjdXJyVGltZSArIHRpbWVUb0NhbGwpO1xuICAgICAgICAgICAgICAgIH1cbiAgICAgICAgICAgICAgICBjYXRjaCAoZSkge1xuICAgICAgICAgICAgICAgICAgICBjb25zb2xlLmxvZygnRXJyb3I6ICcsIGUpO1xuICAgICAgICAgICAgICAgICAgICBzZXRUaW1lb3V0KGZ1bmN0aW9uICgpIHtcbiAgICAgICAgICAgICAgICAgICAgICAgIHRocm93IGU7XG4gICAgICAgICAgICAgICAgICAgIH0sIDApO1xuICAgICAgICAgICAgICAgIH1cbiAgICAgICAgICAgIH0sIHRpbWVUb0NhbGwpO1xuICAgICAgICAgICAgbGFzdFRpbWUgPSBjdXJyVGltZSArIHRpbWVUb0NhbGw7XG4gICAgICAgICAgICByZXR1cm4gaWQ7XG4gICAgICAgIH07XG4gICAgaWYgKCFtb2QuY2FuY2VsQW5pbWF0aW9uRnJhbWUpIHtcbiAgICAgICAgbW9kLmNhbmNlbEFuaW1hdGlvbkZyYW1lID0gZnVuY3Rpb24gKGlkKSB7XG4gICAgICAgICAgICBjbGVhclRpbWVvdXQoaWQpO1xuICAgICAgICB9O1xuICAgIH1cbiAgICByZXR1cm4gbW9kO1xufVxuZXhwb3J0cy5HZXRSQUYgPSBHZXRSQUY7XG4vLyMgc291cmNlTWFwcGluZ1VSTD1yYWYtcG9seWZpbGwuanMubWFwIiwiXCJ1c2Ugc3RyaWN0XCI7XG5PYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgXCJfX2VzTW9kdWxlXCIsIHsgdmFsdWU6IHRydWUgfSk7XG5leHBvcnRzLkRPTUV4Y2VwdGlvbiA9IGV4cG9ydHMuUmVzcG9uc2UgPSBleHBvcnRzLlJlcXVlc3QgPSBleHBvcnRzLkhlYWRlcnMgPSBleHBvcnRzLmZldGNoID0gdm9pZCAwO1xuY29uc3QgZmV0Y2ggPSByZXF1aXJlKFwid2hhdHdnLWZldGNoXCIpO1xuaWYgKCFzZWxmLmZldGNoKSB7XG4gICAgc2VsZi5mZXRjaCA9IGV4cG9ydHMuZmV0Y2guZmV0Y2g7XG4gICAgc2VsZi5IZWFkZXJzID0gZXhwb3J0cy5mZXRjaC5IZWFkZXJzO1xuICAgIHNlbGYuUmVxdWVzdCA9IGV4cG9ydHMuZmV0Y2guUmVxdWVzdDtcbiAgICBzZWxmLlJlc3BvbnNlID0gZXhwb3J0cy5mZXRjaC5SZXNwb25zZTtcbiAgICBzZWxmLkRPTUV4Y2VwdGlvbiA9IGV4cG9ydHMuZmV0Y2guRE9NRXhjZXB0aW9uO1xufVxuZXhwb3J0cy5mZXRjaCA9IHNlbGYuZmV0Y2g7XG5leHBvcnRzLkhlYWRlcnMgPSBzZWxmLkhlYWRlcnM7XG5leHBvcnRzLlJlcXVlc3QgPSBzZWxmLlJlcXVlc3Q7XG5leHBvcnRzLlJlc3BvbnNlID0gc2VsZi5SZXNwb25zZTtcbmV4cG9ydHMuRE9NRXhjZXB0aW9uID0gc2VsZi5ET01FeGNlcHRpb247XG4vLyMgc291cmNlTWFwcGluZ1VSTD1mZXRjaC5qcy5tYXAiLCJcInVzZSBzdHJpY3RcIjtcbk9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBcIl9fZXNNb2R1bGVcIiwgeyB2YWx1ZTogdHJ1ZSB9KTtcbmV4cG9ydHMuZGVmYXVsdCA9IHt9O1xuLy8jIHNvdXJjZU1hcHBpbmdVUkw9aHR0cC5qcy5tYXAiLCJcInVzZSBzdHJpY3RcIjtcbk9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBcIl9fZXNNb2R1bGVcIiwgeyB2YWx1ZTogdHJ1ZSB9KTtcbmV4cG9ydHMubW91bnRUbyA9IHZvaWQgMDtcbmNvbnN0IGZldGNoID0gcmVxdWlyZShcIi4vZmV0Y2hcIik7XG5jb25zdCBodHRwID0gcmVxdWlyZShcIi4vaHR0cFwiKTtcbmNvbnN0IHdlYnNvY2tldCA9IHJlcXVpcmUoXCIuL3dlYnNvY2tldFwiKTtcbmNvbnN0IG5hbWVzcGFjZSA9IE9iamVjdC5hc3NpZ24oT2JqZWN0LmFzc2lnbihPYmplY3QuYXNzaWduKHt9LCBmZXRjaCksIGh0dHApLCB3ZWJzb2NrZXQpO1xuZnVuY3Rpb24gbW91bnRUbyhwYXJlbnQpIHtcbiAgICBwYXJlbnQuaHR0cCA9IG5hbWVzcGFjZTtcbn1cbmV4cG9ydHMubW91bnRUbyA9IG1vdW50VG87XG5leHBvcnRzLmRlZmF1bHQgPSBuYW1lc3BhY2U7XG4vLyMgc291cmNlTWFwcGluZ1VSTD1pbmRleC5qcy5tYXAiLCJcInVzZSBzdHJpY3RcIjtcbk9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBcIl9fZXNNb2R1bGVcIiwgeyB2YWx1ZTogdHJ1ZSB9KTtcbmV4cG9ydHMuU29ja2V0ID0gZXhwb3J0cy5PbmVNaW51dGUgPSBleHBvcnRzLk9uZVNlY29uZCA9IHZvaWQgMDtcbmV4cG9ydHMuT25lU2Vjb25kID0gMTAwMDtcbmV4cG9ydHMuT25lTWludXRlID0gZXhwb3J0cy5PbmVTZWNvbmQgKiA2MDtcbmNsYXNzIFNvY2tldCB7XG4gICAgY29uc3RydWN0b3IoYWRkciwgcmVhZGVyLCBleHBvbmVudCwgbWF4UmVjb25uZWN0cywgbWF4V2FpdCkge1xuICAgICAgICB0aGlzLmFkZHIgPSBhZGRyO1xuICAgICAgICB0aGlzLnNvY2tldCA9IG51bGw7XG4gICAgICAgIHRoaXMucmVhZGVyID0gcmVhZGVyO1xuICAgICAgICB0aGlzLm1heFdhaXQgPSBtYXhXYWl0O1xuICAgICAgICB0aGlzLnVzZXJDbG9zZWQgPSBmYWxzZTtcbiAgICAgICAgdGhpcy5leHBvbmVudCA9IGV4cG9uZW50O1xuICAgICAgICB0aGlzLmRpc2Nvbm5lY3RlZCA9IGZhbHNlO1xuICAgICAgICB0aGlzLmF0dGVtcHRlZENvbm5lY3RzID0gMDtcbiAgICAgICAgdGhpcy5sYXN0V2FpdCA9IGV4cG9ydHMuT25lU2Vjb25kO1xuICAgICAgICB0aGlzLm1heFJlY29ubmVjdCA9IG1heFJlY29ubmVjdHM7XG4gICAgICAgIHRoaXMud3JpdGVCdWZmZXIgPSBuZXcgQXJyYXkoKTtcbiAgICB9XG4gICAgY29ubmVjdCgpIHtcbiAgICAgICAgaWYgKHRoaXMuc29ja2V0KSB7XG4gICAgICAgICAgICByZXR1cm47XG4gICAgICAgIH1cbiAgICAgICAgaWYgKHRoaXMuYXR0ZW1wdGVkQ29ubmVjdHMgPj0gdGhpcy5tYXhSZWNvbm5lY3QpIHtcbiAgICAgICAgICAgIHRoaXMucmVhZGVyLkV4aGF1c3RlZCh0aGlzKTtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICBjb25zdCBzb2NrZXQgPSBuZXcgV2ViU29ja2V0KHRoaXMuYWRkcik7XG4gICAgICAgIHNvY2tldC5hZGRFdmVudExpc3RlbmVyKCdvcGVuJywgdGhpcy5fb3BlbmVkLmJpbmQodGhpcykpO1xuICAgICAgICBzb2NrZXQuYWRkRXZlbnRMaXN0ZW5lcignZXJyb3InLCB0aGlzLl9lcnJvcmVkLmJpbmQodGhpcykpO1xuICAgICAgICBzb2NrZXQuYWRkRXZlbnRMaXN0ZW5lcignbWVzc2FnZScsIHRoaXMuX21lc3NhZ2VkLmJpbmQodGhpcykpO1xuICAgICAgICBzb2NrZXQuYWRkRXZlbnRMaXN0ZW5lcignY2xvc2UnLCB0aGlzLl9kaXNjb25uZWN0ZWQuYmluZCh0aGlzKSk7XG4gICAgICAgIHRoaXMuc29ja2V0ID0gc29ja2V0O1xuICAgICAgICB0aGlzLmRpc2Nvbm5lY3RlZCA9IGZhbHNlO1xuICAgIH1cbiAgICBzZW5kKG1lc3NhZ2UpIHtcbiAgICAgICAgaWYgKHRoaXMuZGlzY29ubmVjdGVkKSB7XG4gICAgICAgICAgICB0aGlzLndyaXRlQnVmZmVyLnB1c2gobWVzc2FnZSk7XG4gICAgICAgICAgICByZXR1cm47XG4gICAgICAgIH1cbiAgICAgICAgdGhpcy5zb2NrZXQuc2VuZChtZXNzYWdlKTtcbiAgICB9XG4gICAgcmVzZXQoKSB7XG4gICAgICAgIHRoaXMuYXR0ZW1wdGVkQ29ubmVjdHMgPSAwO1xuICAgICAgICB0aGlzLmxhc3RXYWl0ID0gZXhwb3J0cy5PbmVTZWNvbmQ7XG4gICAgfVxuICAgIGVuZCgpIHtcbiAgICAgICAgdGhpcy51c2VyQ2xvc2VkID0gdHJ1ZTtcbiAgICAgICAgdGhpcy5yZWFkZXIuQ2xvc2VkKHRoaXMpO1xuICAgICAgICB0aGlzLnNvY2tldC5jbG9zZSgpO1xuICAgICAgICB0aGlzLnNvY2tldCA9IG51bGw7XG4gICAgfVxuICAgIF9kaXNjb25uZWN0ZWQoZXZlbnQpIHtcbiAgICAgICAgdGhpcy5yZWFkZXIuRGlzY29ubmVjdGVkKGV2ZW50LCB0aGlzKTtcbiAgICAgICAgdGhpcy5kaXNjb25uZWN0ZWQgPSB0cnVlO1xuICAgICAgICB0aGlzLnNvY2tldCA9IG51bGw7XG4gICAgICAgIGlmICh0aGlzLnVzZXJDbG9zZWQpIHtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICBsZXQgbmV4dFdhaXQgPSB0aGlzLmxhc3RXYWl0O1xuICAgICAgICBpZiAodGhpcy5leHBvbmVudCkge1xuICAgICAgICAgICAgbmV4dFdhaXQgPSB0aGlzLmV4cG9uZW50KG5leHRXYWl0KTtcbiAgICAgICAgICAgIGlmIChuZXh0V2FpdCA+IHRoaXMubWF4V2FpdCkge1xuICAgICAgICAgICAgICAgIG5leHRXYWl0ID0gdGhpcy5tYXhXYWl0O1xuICAgICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICAgIHNldFRpbWVvdXQodGhpcy5jb25uZWN0LmJpbmQodGhpcyksIG5leHRXYWl0KTtcbiAgICAgICAgdGhpcy5hdHRlbXB0ZWRDb25uZWN0cysrO1xuICAgIH1cbiAgICBfb3BlbmVkKGV2ZW50KSB7XG4gICAgICAgIHRoaXMucmVhZGVyLkNvbm5lY3RlZChldmVudCwgdGhpcyk7XG4gICAgICAgIHdoaWxlICh0aGlzLndyaXRlQnVmZmVyLmxlbmd0aCA+IDApIHtcbiAgICAgICAgICAgIGNvbnN0IG1lc3NhZ2UgPSB0aGlzLndyaXRlQnVmZmVyLnNoaWZ0KCk7XG4gICAgICAgICAgICB0aGlzLnNvY2tldC5zZW5kKG1lc3NhZ2UpO1xuICAgICAgICB9XG4gICAgfVxuICAgIF9lcnJvcmVkKGV2ZW50KSB7XG4gICAgICAgIHRoaXMucmVhZGVyLkVycm9yZWQoZXZlbnQsIHRoaXMpO1xuICAgIH1cbiAgICBfbWVzc2FnZWQoZXZlbnQpIHtcbiAgICAgICAgdGhpcy5yZWFkZXIuTWVzc2FnZShldmVudCwgdGhpcyk7XG4gICAgfVxufVxuZXhwb3J0cy5Tb2NrZXQgPSBTb2NrZXQ7XG4vLyMgc291cmNlTWFwcGluZ1VSTD13ZWJzb2NrZXQuanMubWFwIiwiXCJ1c2Ugc3RyaWN0XCI7XG5PYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgXCJfX2VzTW9kdWxlXCIsIHsgdmFsdWU6IHRydWUgfSk7XG5leHBvcnRzLmdldE1hdGNoaW5nTm9kZSA9IGV4cG9ydHMudG9EYXRhVHJhbnNmZXIgPSBleHBvcnRzLnRvR2FtZXBhZCA9IGV4cG9ydHMudG9Ub3VjaGVzID0gZXhwb3J0cy50b01lZGlhU3RyZWFtID0gZXhwb3J0cy50b1JvdGF0aW9uRGF0YSA9IGV4cG9ydHMudG9Nb3Rpb25EYXRhID0gZXhwb3J0cy50b0lucHV0U291cmNlQ2FwYWJpbGl0eSA9IGV4cG9ydHMuZnJvbUZpbGUgPSBleHBvcnRzLmZyb21CbG9iID0gZXhwb3J0cy5yZW1vdmVGcm9tTm9kZSA9IGV4cG9ydHMuY3JlYXRlVGV4dCA9IGV4cG9ydHMuY3JlYXRlRWxlbWVudCA9IGV4cG9ydHMucmVjb3JkQXR0cmlidXRlcyA9IGV4cG9ydHMuZ2V0TmFtZXNwYWNlRm9yVGFnID0gZXhwb3J0cy5hcHBseUF0dHJpYnV0ZVR5cGVkID0gZXhwb3J0cy5hcHBseVNWR1N0eWxlcyA9IGV4cG9ydHMuYXBwbHlTdHlsZXMgPSBleHBvcnRzLmFwcGx5U3R5bGUgPSBleHBvcnRzLmFwcGx5U1ZHU3R5bGUgPSBleHBvcnRzLnNldFN0eWxlVmFsdWUgPSBleHBvcnRzLmFwcGx5UHJvcCA9IGV4cG9ydHMuYXBwbHlBdHRycyA9IGV4cG9ydHMuYXBwbHlBdHRyID0gZXhwb3J0cy5nZXROYW1lc3BhY2UgPSBleHBvcnRzLnJlcGxhY2VOb2RlSWYgPSBleHBvcnRzLnJlcGxhY2VOb2RlID0gZXhwb3J0cy5pbnNlcnRCZWZvcmUgPSBleHBvcnRzLm1vdmVCZWZvcmUgPSBleHBvcnRzLmdldEZvY3VzZWRQYXRoID0gZXhwb3J0cy5nZXRBY3RpdmVFbGVtZW50ID0gZXhwb3J0cy5jb2xsZWN0QnJlYWR0aEZpcnN0ID0gZXhwb3J0cy50b0hUTUwgPSBleHBvcnRzLmVhY2hOb2RlQW5kQ2hpbGQgPSBleHBvcnRzLm5vZGVMaXN0VG9BcnJheSA9IGV4cG9ydHMuZWFjaENoaWxkQW5kTm9kZSA9IGV4cG9ydHMucmV2ZXJzZUFwcGx5RWFjaE5vZGUgPSBleHBvcnRzLmFwcGx5RWFjaE5vZGUgPSBleHBvcnRzLmFwcGx5RWFjaENoaWxkTm9kZSA9IGV4cG9ydHMuYXBwbHlDaGlsZE5vZGUgPSBleHBvcnRzLmZpbmRCcmVhZHRoRmlyc3QgPSBleHBvcnRzLmNvbGxlY3REZXB0aEZpcnN0ID0gZXhwb3J0cy5maW5kRGVwdGhGaXJzdCA9IGV4cG9ydHMuZmluZE5vZGVXaXRoRGVwdGggPSBleHBvcnRzLmZpbmROb2RlV2l0aEJyZWFkdGggPSBleHBvcnRzLmNvbGxlY3ROb2RlV2l0aERlcHRoID0gZXhwb3J0cy5jb2xsZWN0Tm9kZVdpdGhCcmVhZHRoID0gZXhwb3J0cy5yZXZlcnNlRmluZE5vZGVXaXRoQnJlYWR0aCA9IGV4cG9ydHMucmV2ZXJzZUNvbGxlY3ROb2RlV2l0aEJyZWFkdGggPSBleHBvcnRzLmdldEFuY2VzdHJ5ID0gZXhwb3J0cy5pc1RleHQgPSBleHBvcnRzLmlzRWxlbWVudCA9IGV4cG9ydHMuaXNEb2N1bWVudFJvb3QgPSBleHBvcnRzLkNPTU1FTlRfTk9ERSA9IGV4cG9ydHMuVEVYVF9OT0RFID0gZXhwb3J0cy5ET0NVTUVOVF9OT0RFID0gZXhwb3J0cy5ET0NVTUVOVF9GUkFHTUVOVF9OT0RFID0gZXhwb3J0cy5FTEVNRU5UX05PREUgPSB2b2lkIDA7XG5jb25zdCB1dGlsc18xID0gcmVxdWlyZShcIi4vdXRpbHNcIik7XG5jb25zdCBleHRzID0gcmVxdWlyZShcIi4vZXh0ZW5zaW9uc1wiKTtcbmV4cG9ydHMuRUxFTUVOVF9OT0RFID0gMTtcbmV4cG9ydHMuRE9DVU1FTlRfRlJBR01FTlRfTk9ERSA9IDExO1xuZXhwb3J0cy5ET0NVTUVOVF9OT0RFID0gOTtcbmV4cG9ydHMuVEVYVF9OT0RFID0gMztcbmV4cG9ydHMuQ09NTUVOVF9OT0RFID0gODtcbmNvbnN0IGF0dHJpYnV0ZXMgPSB1dGlsc18xLmNyZWF0ZU1hcCgpO1xuYXR0cmlidXRlc1snc3R5bGUnXSA9IGFwcGx5U3R5bGU7XG5mdW5jdGlvbiBpc0RvY3VtZW50Um9vdChub2RlKSB7XG4gICAgcmV0dXJuIG5vZGUubm9kZVR5cGUgPT09IDExIHx8IG5vZGUubm9kZVR5cGUgPT09IDk7XG59XG5leHBvcnRzLmlzRG9jdW1lbnRSb290ID0gaXNEb2N1bWVudFJvb3Q7XG5mdW5jdGlvbiBpc0VsZW1lbnQobm9kZSkge1xuICAgIHJldHVybiBub2RlLm5vZGVUeXBlID09PSAxO1xufVxuZXhwb3J0cy5pc0VsZW1lbnQgPSBpc0VsZW1lbnQ7XG5mdW5jdGlvbiBpc1RleHQobm9kZSkge1xuICAgIHJldHVybiBub2RlLm5vZGVUeXBlID09PSAzO1xufVxuZXhwb3J0cy5pc1RleHQgPSBpc1RleHQ7XG5mdW5jdGlvbiBnZXRBbmNlc3RyeShub2RlLCByb290KSB7XG4gICAgY29uc3QgYW5jZXN0cnkgPSBbXTtcbiAgICBsZXQgY3VyID0gbm9kZTtcbiAgICB3aGlsZSAoY3VyICE9PSByb290KSB7XG4gICAgICAgIGNvbnN0IG4gPSBjdXI7XG4gICAgICAgIGFuY2VzdHJ5LnB1c2gobik7XG4gICAgICAgIGN1ciA9IG4ucGFyZW50Tm9kZTtcbiAgICB9XG4gICAgcmV0dXJuIGFuY2VzdHJ5O1xufVxuZXhwb3J0cy5nZXRBbmNlc3RyeSA9IGdldEFuY2VzdHJ5O1xuY29uc3QgZ2V0Um9vdE5vZGUgPSBOb2RlLnByb3RvdHlwZS5nZXRSb290Tm9kZSB8fFxuICAgIGZ1bmN0aW9uICgpIHtcbiAgICAgICAgbGV0IGN1ciA9IHRoaXM7XG4gICAgICAgIGxldCBwcmV2ID0gY3VyO1xuICAgICAgICB3aGlsZSAoY3VyKSB7XG4gICAgICAgICAgICBwcmV2ID0gY3VyO1xuICAgICAgICAgICAgY3VyID0gY3VyLnBhcmVudE5vZGU7XG4gICAgICAgIH1cbiAgICAgICAgcmV0dXJuIHByZXY7XG4gICAgfTtcbmZ1bmN0aW9uIHJldmVyc2VDb2xsZWN0Tm9kZVdpdGhCcmVhZHRoKHBhcmVudCwgbWF0Y2hlciwgbWF0Y2hlcykge1xuICAgIGxldCBjdXIgPSBwYXJlbnQubGFzdENoaWxkO1xuICAgIHdoaWxlIChjdXIpIHtcbiAgICAgICAgaWYgKG1hdGNoZXIoY3VyKSkge1xuICAgICAgICAgICAgbWF0Y2hlcy5wdXNoKGN1cik7XG4gICAgICAgIH1cbiAgICAgICAgY3VyID0gY3VyLnByZXZpb3VzU2libGluZztcbiAgICB9XG59XG5leHBvcnRzLnJldmVyc2VDb2xsZWN0Tm9kZVdpdGhCcmVhZHRoID0gcmV2ZXJzZUNvbGxlY3ROb2RlV2l0aEJyZWFkdGg7XG5mdW5jdGlvbiByZXZlcnNlRmluZE5vZGVXaXRoQnJlYWR0aChwYXJlbnQsIG1hdGNoZXIpIHtcbiAgICBsZXQgY3VyID0gcGFyZW50Lmxhc3RDaGlsZDtcbiAgICB3aGlsZSAoY3VyKSB7XG4gICAgICAgIGlmIChtYXRjaGVyKGN1cikpIHtcbiAgICAgICAgICAgIHJldHVybiBjdXI7XG4gICAgICAgIH1cbiAgICAgICAgY3VyID0gY3VyLnByZXZpb3VzU2libGluZztcbiAgICB9XG4gICAgcmV0dXJuIG51bGw7XG59XG5leHBvcnRzLnJldmVyc2VGaW5kTm9kZVdpdGhCcmVhZHRoID0gcmV2ZXJzZUZpbmROb2RlV2l0aEJyZWFkdGg7XG5mdW5jdGlvbiBjb2xsZWN0Tm9kZVdpdGhCcmVhZHRoKHBhcmVudCwgbWF0Y2hlciwgbWF0Y2hlcykge1xuICAgIGxldCBjdXIgPSBwYXJlbnQuZmlyc3RDaGlsZDtcbiAgICBpZiAobWF0Y2hlcihjdXIpKSB7XG4gICAgICAgIG1hdGNoZXMucHVzaChjdXIpO1xuICAgIH1cbiAgICB3aGlsZSAoY3VyKSB7XG4gICAgICAgIGlmIChtYXRjaGVyKGN1ci5uZXh0U2libGluZykpIHtcbiAgICAgICAgICAgIG1hdGNoZXMucHVzaChjdXIpO1xuICAgICAgICB9XG4gICAgICAgIGN1ciA9IGN1ci5uZXh0U2libGluZztcbiAgICB9XG59XG5leHBvcnRzLmNvbGxlY3ROb2RlV2l0aEJyZWFkdGggPSBjb2xsZWN0Tm9kZVdpdGhCcmVhZHRoO1xuZnVuY3Rpb24gY29sbGVjdE5vZGVXaXRoRGVwdGgocGFyZW50LCBtYXRjaGVyLCBtYXRjaGVzKSB7XG4gICAgbGV0IGN1ciA9IHBhcmVudC5maXJzdENoaWxkO1xuICAgIHdoaWxlIChjdXIpIHtcbiAgICAgICAgaWYgKG1hdGNoZXIoY3VyKSkge1xuICAgICAgICAgICAgbWF0Y2hlcy5wdXNoKGN1cik7XG4gICAgICAgIH1cbiAgICAgICAgY3VyID0gY3VyLmZpcnN0Q2hpbGQ7XG4gICAgfVxufVxuZXhwb3J0cy5jb2xsZWN0Tm9kZVdpdGhEZXB0aCA9IGNvbGxlY3ROb2RlV2l0aERlcHRoO1xuZnVuY3Rpb24gZmluZE5vZGVXaXRoQnJlYWR0aChwYXJlbnQsIG1hdGNoZXIpIHtcbiAgICBsZXQgY3VyID0gcGFyZW50LmZpcnN0Q2hpbGQ7XG4gICAgd2hpbGUgKGN1cikge1xuICAgICAgICBpZiAobWF0Y2hlcihjdXIpKSB7XG4gICAgICAgICAgICByZXR1cm4gY3VyO1xuICAgICAgICB9XG4gICAgICAgIGN1ciA9IGN1ci5uZXh0U2libGluZztcbiAgICB9XG4gICAgcmV0dXJuIG51bGw7XG59XG5leHBvcnRzLmZpbmROb2RlV2l0aEJyZWFkdGggPSBmaW5kTm9kZVdpdGhCcmVhZHRoO1xuZnVuY3Rpb24gZmluZE5vZGVXaXRoRGVwdGgocGFyZW50LCBtYXRjaGVyKSB7XG4gICAgbGV0IGN1ciA9IHBhcmVudC5maXJzdENoaWxkO1xuICAgIHdoaWxlIChjdXIpIHtcbiAgICAgICAgaWYgKG1hdGNoZXIoY3VyKSkge1xuICAgICAgICAgICAgcmV0dXJuIGN1cjtcbiAgICAgICAgfVxuICAgICAgICBjdXIgPSBjdXIuZmlyc3RDaGlsZDtcbiAgICB9XG4gICAgcmV0dXJuIG51bGw7XG59XG5leHBvcnRzLmZpbmROb2RlV2l0aERlcHRoID0gZmluZE5vZGVXaXRoRGVwdGg7XG5mdW5jdGlvbiBmaW5kRGVwdGhGaXJzdChwYXJlbnQsIG1hdGNoZXIpIHtcbiAgICBsZXQgY3VyID0gcGFyZW50LmZpcnN0Q2hpbGQ7XG4gICAgd2hpbGUgKGN1cikge1xuICAgICAgICBjb25zdCBmb3VuZCA9IGZpbmROb2RlV2l0aERlcHRoKGN1ciwgbWF0Y2hlcik7XG4gICAgICAgIGlmIChmb3VuZCkge1xuICAgICAgICAgICAgcmV0dXJuIGZvdW5kO1xuICAgICAgICB9XG4gICAgICAgIGN1ciA9IGN1ci5uZXh0U2libGluZztcbiAgICB9XG4gICAgcmV0dXJuIG51bGw7XG59XG5leHBvcnRzLmZpbmREZXB0aEZpcnN0ID0gZmluZERlcHRoRmlyc3Q7XG5mdW5jdGlvbiBjb2xsZWN0RGVwdGhGaXJzdChwYXJlbnQsIG1hdGNoZXIsIG1hdGNoZXMpIHtcbiAgICBsZXQgY3VyID0gcGFyZW50LmZpcnN0Q2hpbGQ7XG4gICAgd2hpbGUgKGN1cikge1xuICAgICAgICBjb2xsZWN0Tm9kZVdpdGhEZXB0aChjdXIsIG1hdGNoZXIsIG1hdGNoZXMpO1xuICAgICAgICBjdXIgPSBjdXIubmV4dFNpYmxpbmc7XG4gICAgfVxuICAgIHJldHVybjtcbn1cbmV4cG9ydHMuY29sbGVjdERlcHRoRmlyc3QgPSBjb2xsZWN0RGVwdGhGaXJzdDtcbmZ1bmN0aW9uIGZpbmRCcmVhZHRoRmlyc3QocGFyZW50LCBtYXRjaGVyKSB7XG4gICAgbGV0IGN1ciA9IHBhcmVudC5maXJzdENoaWxkO1xuICAgIHdoaWxlIChjdXIpIHtcbiAgICAgICAgY29uc3QgZm91bmQgPSBmaW5kTm9kZVdpdGhCcmVhZHRoKGN1ciwgbWF0Y2hlcik7XG4gICAgICAgIGlmIChmb3VuZCkge1xuICAgICAgICAgICAgcmV0dXJuIGZvdW5kO1xuICAgICAgICB9XG4gICAgICAgIGN1ciA9IGN1ci5maXJzdENoaWxkO1xuICAgIH1cbiAgICByZXR1cm4gbnVsbDtcbn1cbmV4cG9ydHMuZmluZEJyZWFkdGhGaXJzdCA9IGZpbmRCcmVhZHRoRmlyc3Q7XG5mdW5jdGlvbiBhcHBseUNoaWxkTm9kZShwYXJlbnQsIGZuKSB7XG4gICAgbGV0IGN1ciA9IHBhcmVudC5maXJzdENoaWxkO1xuICAgIHdoaWxlIChjdXIpIHtcbiAgICAgICAgZm4oY3VyKTtcbiAgICAgICAgY3VyID0gY3VyLm5leHRTaWJsaW5nO1xuICAgIH1cbn1cbmV4cG9ydHMuYXBwbHlDaGlsZE5vZGUgPSBhcHBseUNoaWxkTm9kZTtcbmZ1bmN0aW9uIGFwcGx5RWFjaENoaWxkTm9kZShwYXJlbnQsIGZuKSB7XG4gICAgbGV0IGN1ciA9IHBhcmVudC5maXJzdENoaWxkO1xuICAgIHdoaWxlIChjdXIpIHtcbiAgICAgICAgZm4oY3VyKTtcbiAgICAgICAgYXBwbHlFYWNoQ2hpbGROb2RlKGN1ciwgZm4pO1xuICAgICAgICBjdXIgPSBjdXIubmV4dFNpYmxpbmc7XG4gICAgfVxufVxuZXhwb3J0cy5hcHBseUVhY2hDaGlsZE5vZGUgPSBhcHBseUVhY2hDaGlsZE5vZGU7XG5mdW5jdGlvbiBhcHBseUVhY2hOb2RlKHBhcmVudCwgZm4pIHtcbiAgICBmbihwYXJlbnQpO1xuICAgIGxldCBjdXIgPSBwYXJlbnQuZmlyc3RDaGlsZDtcbiAgICB3aGlsZSAoY3VyKSB7XG4gICAgICAgIGFwcGx5RWFjaE5vZGUoY3VyLCBmbik7XG4gICAgICAgIGN1ciA9IGN1ci5uZXh0U2libGluZztcbiAgICB9XG59XG5leHBvcnRzLmFwcGx5RWFjaE5vZGUgPSBhcHBseUVhY2hOb2RlO1xuZnVuY3Rpb24gcmV2ZXJzZUFwcGx5RWFjaE5vZGUocGFyZW50LCBmbikge1xuICAgIGxldCBjdXIgPSBwYXJlbnQubGFzdENoaWxkO1xuICAgIHdoaWxlIChjdXIpIHtcbiAgICAgICAgcmV2ZXJzZUFwcGx5RWFjaE5vZGUoY3VyLCBmbik7XG4gICAgICAgIGZuKGN1cik7XG4gICAgICAgIGN1ciA9IGN1ci5wcmV2aW91c1NpYmxpbmc7XG4gICAgfVxuICAgIGZuKHBhcmVudCk7XG59XG5leHBvcnRzLnJldmVyc2VBcHBseUVhY2hOb2RlID0gcmV2ZXJzZUFwcGx5RWFjaE5vZGU7XG5mdW5jdGlvbiBlYWNoQ2hpbGRBbmROb2RlKG5vZGUsIGZuKSB7XG4gICAgY29uc3QgbGlzdCA9IG5vZGUuY2hpbGROb2RlcztcbiAgICBmb3IgKGxldCBpID0gMDsgaSA8IGxpc3QubGVuZ3RoOyBpKyspIHtcbiAgICAgICAgZm4obGlzdFtpXSk7XG4gICAgfVxuICAgIGZuKG5vZGUpO1xufVxuZXhwb3J0cy5lYWNoQ2hpbGRBbmROb2RlID0gZWFjaENoaWxkQW5kTm9kZTtcbmZ1bmN0aW9uIG5vZGVMaXN0VG9BcnJheShpdGVtcykge1xuICAgIGNvbnN0IGxpc3QgPSBbXTtcbiAgICBmb3IgKGxldCBpID0gMDsgaSA8IGl0ZW1zLmxlbmd0aDsgaSsrKSB7XG4gICAgICAgIGxpc3QucHVzaChpdGVtc1tpXSk7XG4gICAgfVxuICAgIHJldHVybiBsaXN0O1xufVxuZXhwb3J0cy5ub2RlTGlzdFRvQXJyYXkgPSBub2RlTGlzdFRvQXJyYXk7XG5mdW5jdGlvbiBlYWNoTm9kZUFuZENoaWxkKG5vZGUsIGZuKSB7XG4gICAgZm4obm9kZSk7XG4gICAgY29uc3QgbGlzdCA9IG5vZGUuY2hpbGROb2RlcztcbiAgICBmb3IgKGxldCBpID0gMDsgaSA8IGxpc3QubGVuZ3RoOyBpKyspIHtcbiAgICAgICAgZm4obGlzdFtpXSk7XG4gICAgfVxufVxuZXhwb3J0cy5lYWNoTm9kZUFuZENoaWxkID0gZWFjaE5vZGVBbmRDaGlsZDtcbmZ1bmN0aW9uIHRvSFRNTChub2RlLCBzaGFsbG93KSB7XG4gICAgY29uc3QgZGl2ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnZGl2Jyk7XG4gICAgZGl2LmFwcGVuZENoaWxkKG5vZGUuY2xvbmVOb2RlKCFzaGFsbG93KSk7XG4gICAgcmV0dXJuIGRpdi5pbm5lckhUTUw7XG59XG5leHBvcnRzLnRvSFRNTCA9IHRvSFRNTDtcbmZ1bmN0aW9uIGNvbGxlY3RCcmVhZHRoRmlyc3QocGFyZW50LCBtYXRjaGVyLCBtYXRjaGVzKSB7XG4gICAgbGV0IGN1ciA9IHBhcmVudC5maXJzdENoaWxkO1xuICAgIHdoaWxlIChjdXIpIHtcbiAgICAgICAgY29sbGVjdE5vZGVXaXRoQnJlYWR0aChjdXIsIG1hdGNoZXIsIG1hdGNoZXMpO1xuICAgICAgICBjdXIgPSBjdXIuZmlyc3RDaGlsZDtcbiAgICB9XG4gICAgcmV0dXJuO1xufVxuZXhwb3J0cy5jb2xsZWN0QnJlYWR0aEZpcnN0ID0gY29sbGVjdEJyZWFkdGhGaXJzdDtcbmZ1bmN0aW9uIGdldEFjdGl2ZUVsZW1lbnQobm9kZSkge1xuICAgIGNvbnN0IHJvb3QgPSBnZXRSb290Tm9kZS5jYWxsKG5vZGUpO1xuICAgIHJldHVybiBpc0RvY3VtZW50Um9vdChyb290KSA/IHJvb3QuYWN0aXZlRWxlbWVudCA6IG51bGw7XG59XG5leHBvcnRzLmdldEFjdGl2ZUVsZW1lbnQgPSBnZXRBY3RpdmVFbGVtZW50O1xuZnVuY3Rpb24gZ2V0Rm9jdXNlZFBhdGgobm9kZSwgcm9vdCkge1xuICAgIGNvbnN0IGFjdGl2ZUVsZW1lbnQgPSBnZXRBY3RpdmVFbGVtZW50KG5vZGUpO1xuICAgIGlmICghYWN0aXZlRWxlbWVudCB8fCAhbm9kZS5jb250YWlucyhhY3RpdmVFbGVtZW50KSkge1xuICAgICAgICByZXR1cm4gW107XG4gICAgfVxuICAgIHJldHVybiBnZXRBbmNlc3RyeShhY3RpdmVFbGVtZW50LCByb290KTtcbn1cbmV4cG9ydHMuZ2V0Rm9jdXNlZFBhdGggPSBnZXRGb2N1c2VkUGF0aDtcbmZ1bmN0aW9uIG1vdmVCZWZvcmUocGFyZW50Tm9kZSwgbm9kZSwgcmVmZXJlbmNlTm9kZSkge1xuICAgIGNvbnN0IGluc2VydFJlZmVyZW5jZU5vZGUgPSBub2RlLm5leHRTaWJsaW5nO1xuICAgIGxldCBjdXIgPSByZWZlcmVuY2VOb2RlO1xuICAgIHdoaWxlIChjdXIgIT09IG51bGwgJiYgY3VyICE9PSBub2RlKSB7XG4gICAgICAgIGNvbnN0IG5leHQgPSBjdXIubmV4dFNpYmxpbmc7XG4gICAgICAgIHBhcmVudE5vZGUuaW5zZXJ0QmVmb3JlKGN1ciwgaW5zZXJ0UmVmZXJlbmNlTm9kZSk7XG4gICAgICAgIGN1ciA9IG5leHQ7XG4gICAgfVxufVxuZXhwb3J0cy5tb3ZlQmVmb3JlID0gbW92ZUJlZm9yZTtcbmZ1bmN0aW9uIGluc2VydEJlZm9yZShwYXJlbnROb2RlLCBub2RlLCByZWZlcmVuY2VOb2RlKSB7XG4gICAgaWYgKHJlZmVyZW5jZU5vZGUgPT09IG51bGwpIHtcbiAgICAgICAgcGFyZW50Tm9kZS5hcHBlbmRDaGlsZChub2RlKTtcbiAgICAgICAgcmV0dXJuIG51bGw7XG4gICAgfVxuICAgIHBhcmVudE5vZGUuaW5zZXJ0QmVmb3JlKG5vZGUsIHJlZmVyZW5jZU5vZGUpO1xuICAgIHJldHVybiBudWxsO1xufVxuZXhwb3J0cy5pbnNlcnRCZWZvcmUgPSBpbnNlcnRCZWZvcmU7XG5mdW5jdGlvbiByZXBsYWNlTm9kZShwYXJlbnROb2RlLCBub2RlLCByZXBsYWNlbWVudCkge1xuICAgIGlmIChyZXBsYWNlbWVudCA9PT0gbnVsbCkge1xuICAgICAgICByZXR1cm4gbnVsbDtcbiAgICB9XG4gICAgcGFyZW50Tm9kZS5yZXBsYWNlQ2hpbGQocmVwbGFjZW1lbnQsIG5vZGUpO1xuICAgIHJldHVybiBudWxsO1xufVxuZXhwb3J0cy5yZXBsYWNlTm9kZSA9IHJlcGxhY2VOb2RlO1xuZnVuY3Rpb24gcmVwbGFjZU5vZGVJZih0YXJnZXROb2RlLCByZXBsYWNlbWVudCkge1xuICAgIGlmIChyZXBsYWNlbWVudCA9PT0gbnVsbCkge1xuICAgICAgICByZXR1cm4gZmFsc2U7XG4gICAgfVxuICAgIGNvbnN0IHBhcmVudCA9IHRhcmdldE5vZGUucGFyZW50Tm9kZTtcbiAgICBpZiAoIXBhcmVudCkge1xuICAgICAgICByZXR1cm4gZmFsc2U7XG4gICAgfVxuICAgIHBhcmVudC5yZXBsYWNlQ2hpbGQocmVwbGFjZW1lbnQsIHRhcmdldE5vZGUpO1xuICAgIHJldHVybiB0cnVlO1xufVxuZXhwb3J0cy5yZXBsYWNlTm9kZUlmID0gcmVwbGFjZU5vZGVJZjtcbmZ1bmN0aW9uIGdldE5hbWVzcGFjZShuYW1lKSB7XG4gICAgaWYgKG5hbWUubGFzdEluZGV4T2YoJ3htbDonLCAwKSA9PT0gMCkge1xuICAgICAgICByZXR1cm4gJ2h0dHA6Ly93d3cudzMub3JnL1hNTC8xOTk4L25hbWVzcGFjZSc7XG4gICAgfVxuICAgIGlmIChuYW1lLmxhc3RJbmRleE9mKCd4bGluazonLCAwKSA9PT0gMCkge1xuICAgICAgICByZXR1cm4gJ2h0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsnO1xuICAgIH1cbiAgICByZXR1cm4gdW5kZWZpbmVkO1xufVxuZXhwb3J0cy5nZXROYW1lc3BhY2UgPSBnZXROYW1lc3BhY2U7XG5mdW5jdGlvbiBhcHBseUF0dHIoZWwsIG5hbWUsIHZhbHVlKSB7XG4gICAgaWYgKHZhbHVlID09IG51bGwpIHtcbiAgICAgICAgZWwucmVtb3ZlQXR0cmlidXRlKG5hbWUpO1xuICAgIH1cbiAgICBlbHNlIHtcbiAgICAgICAgY29uc3QgYXR0ck5TID0gZ2V0TmFtZXNwYWNlKG5hbWUpO1xuICAgICAgICBpZiAoYXR0ck5TKSB7XG4gICAgICAgICAgICBlbC5zZXRBdHRyaWJ1dGVOUyhhdHRyTlMsIG5hbWUsIFN0cmluZyh2YWx1ZSkpO1xuICAgICAgICB9XG4gICAgICAgIGVsc2Uge1xuICAgICAgICAgICAgZWwuc2V0QXR0cmlidXRlKG5hbWUsIFN0cmluZyh2YWx1ZSkpO1xuICAgICAgICB9XG4gICAgfVxufVxuZXhwb3J0cy5hcHBseUF0dHIgPSBhcHBseUF0dHI7XG5mdW5jdGlvbiBhcHBseUF0dHJzKGVsLCB2YWx1ZXMpIHtcbiAgICBmb3IgKGxldCBrZXkgaW4gdmFsdWVzKSB7XG4gICAgICAgIGlmICh2YWx1ZXNba2V5XSA9PSBudWxsKSB7XG4gICAgICAgICAgICBlbC5yZW1vdmVBdHRyaWJ1dGUoa2V5KTtcbiAgICAgICAgICAgIGNvbnRpbnVlO1xuICAgICAgICB9XG4gICAgICAgIGVsLnNldEF0dHJpYnV0ZShrZXksIHZhbHVlc1trZXldKTtcbiAgICB9XG59XG5leHBvcnRzLmFwcGx5QXR0cnMgPSBhcHBseUF0dHJzO1xuZnVuY3Rpb24gYXBwbHlQcm9wKGVsLCBuYW1lLCB2YWx1ZSkge1xuICAgIGVsW25hbWVdID0gdmFsdWU7XG59XG5leHBvcnRzLmFwcGx5UHJvcCA9IGFwcGx5UHJvcDtcbmZ1bmN0aW9uIHNldFN0eWxlVmFsdWUoc3R5bGUsIHByb3AsIHZhbHVlKSB7XG4gICAgaWYgKHByb3AuaW5kZXhPZignLScpID49IDApIHtcbiAgICAgICAgc3R5bGUuc2V0UHJvcGVydHkocHJvcCwgdmFsdWUpO1xuICAgIH1cbiAgICBlbHNlIHtcbiAgICAgICAgc3R5bGVbcHJvcF0gPSB2YWx1ZTtcbiAgICB9XG59XG5leHBvcnRzLnNldFN0eWxlVmFsdWUgPSBzZXRTdHlsZVZhbHVlO1xuZnVuY3Rpb24gYXBwbHlTVkdTdHlsZShlbCwgbmFtZSwgc3R5bGUpIHtcbiAgICBpZiAodHlwZW9mIHN0eWxlID09PSAnc3RyaW5nJykge1xuICAgICAgICBlbC5zdHlsZS5jc3NUZXh0ID0gc3R5bGU7XG4gICAgfVxuICAgIGVsc2Uge1xuICAgICAgICBlbC5zdHlsZS5jc3NUZXh0ID0gJyc7XG4gICAgICAgIGNvbnN0IGVsU3R5bGUgPSBlbC5zdHlsZTtcbiAgICAgICAgZm9yIChjb25zdCBwcm9wIGluIHN0eWxlKSB7XG4gICAgICAgICAgICBpZiAodXRpbHNfMS5oYXMoc3R5bGUsIHByb3ApKSB7XG4gICAgICAgICAgICAgICAgc2V0U3R5bGVWYWx1ZShlbFN0eWxlLCBwcm9wLCBzdHlsZVtwcm9wXSk7XG4gICAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICB9XG59XG5leHBvcnRzLmFwcGx5U1ZHU3R5bGUgPSBhcHBseVNWR1N0eWxlO1xuZnVuY3Rpb24gYXBwbHlTdHlsZShlbCwgbmFtZSwgc3R5bGUpIHtcbiAgICBpZiAodHlwZW9mIHN0eWxlID09PSAnc3RyaW5nJykge1xuICAgICAgICBlbC5zdHlsZS5jc3NUZXh0ID0gc3R5bGU7XG4gICAgfVxuICAgIGVsc2Uge1xuICAgICAgICBlbC5zdHlsZS5jc3NUZXh0ID0gJyc7XG4gICAgICAgIGNvbnN0IGVsU3R5bGUgPSBlbC5zdHlsZTtcbiAgICAgICAgZm9yIChjb25zdCBwcm9wIGluIHN0eWxlKSB7XG4gICAgICAgICAgICBpZiAodXRpbHNfMS5oYXMoc3R5bGUsIHByb3ApKSB7XG4gICAgICAgICAgICAgICAgc2V0U3R5bGVWYWx1ZShlbFN0eWxlLCBwcm9wLCBzdHlsZVtwcm9wXSk7XG4gICAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICB9XG59XG5leHBvcnRzLmFwcGx5U3R5bGUgPSBhcHBseVN0eWxlO1xuZnVuY3Rpb24gYXBwbHlTdHlsZXMoZWwsIHN0eWxlKSB7XG4gICAgaWYgKHR5cGVvZiBzdHlsZSA9PT0gJ3N0cmluZycpIHtcbiAgICAgICAgZWwuc3R5bGUuY3NzVGV4dCA9IHN0eWxlO1xuICAgIH1cbiAgICBlbHNlIHtcbiAgICAgICAgZWwuc3R5bGUuY3NzVGV4dCA9ICcnO1xuICAgICAgICBjb25zdCBlbFN0eWxlID0gZWwuc3R5bGU7XG4gICAgICAgIGZvciAoY29uc3QgcHJvcCBpbiBzdHlsZSkge1xuICAgICAgICAgICAgaWYgKHV0aWxzXzEuaGFzKHN0eWxlLCBwcm9wKSkge1xuICAgICAgICAgICAgICAgIHNldFN0eWxlVmFsdWUoZWxTdHlsZSwgcHJvcCwgc3R5bGVbcHJvcF0pO1xuICAgICAgICAgICAgfVxuICAgICAgICB9XG4gICAgfVxufVxuZXhwb3J0cy5hcHBseVN0eWxlcyA9IGFwcGx5U3R5bGVzO1xuZnVuY3Rpb24gYXBwbHlTVkdTdHlsZXMoZWwsIHN0eWxlKSB7XG4gICAgaWYgKHR5cGVvZiBzdHlsZSA9PT0gJ3N0cmluZycpIHtcbiAgICAgICAgZWwuc3R5bGUuY3NzVGV4dCA9IHN0eWxlO1xuICAgIH1cbiAgICBlbHNlIHtcbiAgICAgICAgZWwuc3R5bGUuY3NzVGV4dCA9ICcnO1xuICAgICAgICBjb25zdCBlbFN0eWxlID0gZWwuc3R5bGU7XG4gICAgICAgIGZvciAoY29uc3QgcHJvcCBpbiBzdHlsZSkge1xuICAgICAgICAgICAgaWYgKHV0aWxzXzEuaGFzKHN0eWxlLCBwcm9wKSkge1xuICAgICAgICAgICAgICAgIHNldFN0eWxlVmFsdWUoZWxTdHlsZSwgcHJvcCwgc3R5bGVbcHJvcF0pO1xuICAgICAgICAgICAgfVxuICAgICAgICB9XG4gICAgfVxufVxuZXhwb3J0cy5hcHBseVNWR1N0eWxlcyA9IGFwcGx5U1ZHU3R5bGVzO1xuZnVuY3Rpb24gYXBwbHlBdHRyaWJ1dGVUeXBlZChlbCwgbmFtZSwgdmFsdWUpIHtcbiAgICBjb25zdCB0eXBlID0gdHlwZW9mIHZhbHVlO1xuICAgIGlmICh0eXBlID09PSAnb2JqZWN0JyB8fCB0eXBlID09PSAnZnVuY3Rpb24nKSB7XG4gICAgICAgIGFwcGx5UHJvcChlbCwgbmFtZSwgdmFsdWUpO1xuICAgIH1cbiAgICBlbHNlIHtcbiAgICAgICAgYXBwbHlBdHRyKGVsLCBuYW1lLCB2YWx1ZSk7XG4gICAgfVxufVxuZXhwb3J0cy5hcHBseUF0dHJpYnV0ZVR5cGVkID0gYXBwbHlBdHRyaWJ1dGVUeXBlZDtcbmZ1bmN0aW9uIGdldE5hbWVzcGFjZUZvclRhZyh0YWcsIHBhcmVudCkge1xuICAgIGlmICh0YWcgPT09ICdzdmcnKSB7XG4gICAgICAgIHJldHVybiAnaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmcnO1xuICAgIH1cbiAgICBpZiAodGFnID09PSAnbWF0aCcpIHtcbiAgICAgICAgcmV0dXJuICdodHRwOi8vd3d3LnczLm9yZy8xOTk4L01hdGgvTWF0aE1MJztcbiAgICB9XG4gICAgaWYgKHBhcmVudCA9PSBudWxsKSB7XG4gICAgICAgIHJldHVybiBudWxsO1xuICAgIH1cbiAgICByZXR1cm4gcGFyZW50Lm5hbWVzcGFjZVVSSTtcbn1cbmV4cG9ydHMuZ2V0TmFtZXNwYWNlRm9yVGFnID0gZ2V0TmFtZXNwYWNlRm9yVGFnO1xuZnVuY3Rpb24gcmVjb3JkQXR0cmlidXRlcyhub2RlKSB7XG4gICAgY29uc3QgYXR0cnMgPSB7fTtcbiAgICBjb25zdCBhdHRyaWJ1dGVzID0gbm9kZS5hdHRyaWJ1dGVzO1xuICAgIGNvbnN0IGxlbmd0aCA9IGF0dHJpYnV0ZXMubGVuZ3RoO1xuICAgIGlmICghbGVuZ3RoKSB7XG4gICAgICAgIHJldHVybiBhdHRycztcbiAgICB9XG4gICAgZm9yIChsZXQgaSA9IDAsIGogPSAwOyBpIDwgbGVuZ3RoOyBpICs9IDEsIGogKz0gMikge1xuICAgICAgICBjb25zdCBhdHRyID0gYXR0cmlidXRlc1tpXTtcbiAgICAgICAgYXR0cnNbYXR0ci5uYW1lXSA9IGF0dHIudmFsdWU7XG4gICAgfVxuICAgIHJldHVybiBhdHRycztcbn1cbmV4cG9ydHMucmVjb3JkQXR0cmlidXRlcyA9IHJlY29yZEF0dHJpYnV0ZXM7XG5mdW5jdGlvbiBjcmVhdGVFbGVtZW50KGRvYywgbmFtZU9yQ3Rvciwga2V5LCBjb250ZW50LCBhdHRyaWJ1dGVzLCBuYW1lc3BhY2UpIHtcbiAgICBsZXQgZWw7XG4gICAgaWYgKHR5cGVvZiBuYW1lT3JDdG9yID09PSAnZnVuY3Rpb24nKSB7XG4gICAgICAgIGVsID0gbmV3IG5hbWVPckN0b3IoKTtcbiAgICAgICAgcmV0dXJuIGVsO1xuICAgIH1cbiAgICBuYW1lc3BhY2UgPSBuYW1lc3BhY2UudHJpbSgpO1xuICAgIGlmIChuYW1lc3BhY2UubGVuZ3RoID4gMCkge1xuICAgICAgICBzd2l0Y2ggKG5hbWVPckN0b3IpIHtcbiAgICAgICAgICAgIGNhc2UgJ3N2Zyc6XG4gICAgICAgICAgICAgICAgZWwgPSBkb2MuY3JlYXRlRWxlbWVudE5TKCdodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZycsIG5hbWVPckN0b3IpO1xuICAgICAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICAgICAgY2FzZSAnbWF0aCc6XG4gICAgICAgICAgICAgICAgZWwgPSBkb2MuY3JlYXRlRWxlbWVudE5TKCdodHRwOi8vd3d3LnczLm9yZy8xOTk4L01hdGgvTWF0aE1MJywgbmFtZU9yQ3Rvcik7XG4gICAgICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgICAgICBkZWZhdWx0OlxuICAgICAgICAgICAgICAgIGVsID0gZG9jLmNyZWF0ZUVsZW1lbnROUyhuYW1lc3BhY2UsIG5hbWVPckN0b3IpO1xuICAgICAgICB9XG4gICAgfVxuICAgIGVsc2Uge1xuICAgICAgICBlbCA9IGRvYy5jcmVhdGVFbGVtZW50KG5hbWVPckN0b3IpO1xuICAgIH1cbiAgICBlbC5zZXRBdHRyaWJ1dGUoJ19rZXknLCBrZXkpO1xuICAgIGlmIChhdHRyaWJ1dGVzKSB7XG4gICAgICAgIGFwcGx5QXR0cnMoZWwsIGF0dHJpYnV0ZXMpO1xuICAgIH1cbiAgICBpZiAoY29udGVudC5sZW5ndGggPiAwKSB7XG4gICAgICAgIGVsLmlubmVySFRNTCA9IGNvbnRlbnQ7XG4gICAgfVxuICAgIHJldHVybiBlbDtcbn1cbmV4cG9ydHMuY3JlYXRlRWxlbWVudCA9IGNyZWF0ZUVsZW1lbnQ7XG5mdW5jdGlvbiBjcmVhdGVUZXh0KGRvYywgdGV4dCwga2V5KSB7XG4gICAgY29uc3Qgbm9kZSA9IGRvYy5jcmVhdGVUZXh0Tm9kZSh0ZXh0KTtcbiAgICBleHRzLk9iamVjdHMuUGF0Y2hXaXRoKG5vZGUsICdrZXknLCBrZXkpO1xuICAgIHJldHVybiBub2RlO1xufVxuZXhwb3J0cy5jcmVhdGVUZXh0ID0gY3JlYXRlVGV4dDtcbmZ1bmN0aW9uIHJlbW92ZUZyb21Ob2RlKGZyb21Ob2RlLCBlbmROb2RlKSB7XG4gICAgY29uc3QgcGFyZW50Tm9kZSA9IGZyb21Ob2RlLnBhcmVudE5vZGU7XG4gICAgbGV0IGNoaWxkID0gZnJvbU5vZGU7XG4gICAgd2hpbGUgKGNoaWxkICE9PSBlbmROb2RlKSB7XG4gICAgICAgIGNvbnN0IG5leHQgPSBjaGlsZC5uZXh0U2libGluZztcbiAgICAgICAgcGFyZW50Tm9kZS5yZW1vdmVDaGlsZChjaGlsZCk7XG4gICAgICAgIGNoaWxkID0gbmV4dDtcbiAgICB9XG59XG5leHBvcnRzLnJlbW92ZUZyb21Ob2RlID0gcmVtb3ZlRnJvbU5vZGU7XG5mdW5jdGlvbiBmcm9tQmxvYihvKSB7XG4gICAgaWYgKG8gPT09IG51bGwgfHwgbyA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIHJldHVybjtcbiAgICB9XG4gICAgbGV0IGRhdGEgPSBudWxsO1xuICAgIGNvbnN0IGZpbGVSZWFkZXIgPSBuZXcgRmlsZVJlYWRlcigpO1xuICAgIGZpbGVSZWFkZXIub25sb2FkZW5kID0gZnVuY3Rpb24gKCkge1xuICAgICAgICBkYXRhID0gbmV3IFVpbnQ4QXJyYXkoZmlsZVJlYWRlci5yZXN1bHQpO1xuICAgIH07XG4gICAgZmlsZVJlYWRlci5yZWFkQXNBcnJheUJ1ZmZlcihvKTtcbiAgICByZXR1cm4gZGF0YTtcbn1cbmV4cG9ydHMuZnJvbUJsb2IgPSBmcm9tQmxvYjtcbmZ1bmN0aW9uIGZyb21GaWxlKG8pIHtcbiAgICBpZiAobyA9PT0gbnVsbCB8fCBvID09PSB1bmRlZmluZWQpIHtcbiAgICAgICAgcmV0dXJuO1xuICAgIH1cbiAgICBsZXQgZGF0YSA9IG51bGw7XG4gICAgY29uc3QgZmlsZVJlYWRlciA9IG5ldyBGaWxlUmVhZGVyKCk7XG4gICAgZmlsZVJlYWRlci5vbmxvYWRlbmQgPSBmdW5jdGlvbiAoKSB7XG4gICAgICAgIGRhdGEgPSBuZXcgVWludDhBcnJheShmaWxlUmVhZGVyLnJlc3VsdCk7XG4gICAgfTtcbiAgICBmaWxlUmVhZGVyLnJlYWRBc0FycmF5QnVmZmVyKG8pO1xuICAgIHJldHVybiBkYXRhO1xufVxuZXhwb3J0cy5mcm9tRmlsZSA9IGZyb21GaWxlO1xuZnVuY3Rpb24gdG9JbnB1dFNvdXJjZUNhcGFiaWxpdHkobykge1xuICAgIGlmIChvID09PSBudWxsIHx8IG8gPT09IHVuZGVmaW5lZCkge1xuICAgICAgICByZXR1cm47XG4gICAgfVxuICAgIHJldHVybiB7XG4gICAgICAgIEZpcmVzVG91Y2hFdmVudDogby5maXJlc1RvdWNoRXZlbnQsXG4gICAgfTtcbn1cbmV4cG9ydHMudG9JbnB1dFNvdXJjZUNhcGFiaWxpdHkgPSB0b0lucHV0U291cmNlQ2FwYWJpbGl0eTtcbmZ1bmN0aW9uIHRvTW90aW9uRGF0YShvKSB7XG4gICAgbGV0IG1kID0geyBYOiAwLjAsIFk6IDAuMCwgWjogMC4wIH07XG4gICAgaWYgKG8gPT09IG51bGwgfHwgbyA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIHJldHVybiBtZDtcbiAgICB9XG4gICAgbWQuWCA9IG8ueDtcbiAgICBtZC5ZID0gby55O1xuICAgIG1kLlogPSBvLno7XG4gICAgcmV0dXJuIG1kO1xufVxuZXhwb3J0cy50b01vdGlvbkRhdGEgPSB0b01vdGlvbkRhdGE7XG5mdW5jdGlvbiB0b1JvdGF0aW9uRGF0YShvKSB7XG4gICAgaWYgKG8gPT09IG51bGwgfHwgbyA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIHJldHVybjtcbiAgICB9XG4gICAgY29uc3QgbWQgPSB7fTtcbiAgICBtZC5BbHBoYSA9IG8uYWxwaGE7XG4gICAgbWQuQmV0YSA9IG8uYmV0YTtcbiAgICBtZC5HYW1tYSA9IG8uZ2FtbWE7XG4gICAgcmV0dXJuIG1kO1xufVxuZXhwb3J0cy50b1JvdGF0aW9uRGF0YSA9IHRvUm90YXRpb25EYXRhO1xuZnVuY3Rpb24gdG9NZWRpYVN0cmVhbShvKSB7XG4gICAgaWYgKG8gPT09IG51bGwgfHwgbyA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIHJldHVybjtcbiAgICB9XG4gICAgY29uc3Qgc3RyZWFtID0geyBBdWRpb3M6IFtdLCBWaWRlb3M6IFtdIH07XG4gICAgc3RyZWFtLkFjdGl2ZSA9IG8uYWN0aXZlO1xuICAgIHN0cmVhbS5FbmRlZCA9IG8uZW5kZWQ7XG4gICAgc3RyZWFtLklEID0gby5pZDtcbiAgICBzdHJlYW0uQXVkaW9zID0gW107XG4gICAgc3RyZWFtLlZpZGVvcyA9IFtdO1xuICAgIGxldCBhdWRpb1RyYWNrcyA9IG8uZ2V0QXVkaW9UcmFja3MoKTtcbiAgICBpZiAoYXVkaW9UcmFja3MgIT09IG51bGwgJiYgYXVkaW9UcmFja3MgIT09IHVuZGVmaW5lZCkge1xuICAgICAgICBmb3IgKGxldCBpID0gMDsgaSA8IGF1ZGlvVHJhY2tzLmxlbmd0aDsgaSsrKSB7XG4gICAgICAgICAgICBsZXQgdHJhY2sgPSBhdWRpb1RyYWNrc1tpXTtcbiAgICAgICAgICAgIGxldCBzZXR0aW5ncyA9IHRyYWNrLmdldFNldHRpbmdzKCk7XG4gICAgICAgICAgICBzdHJlYW0uQXVkaW9zLnB1c2goe1xuICAgICAgICAgICAgICAgIEVuYWJsZWQ6IHRyYWNrLmVuYWJsZWQsXG4gICAgICAgICAgICAgICAgSUQ6IHRyYWNrLmlkLFxuICAgICAgICAgICAgICAgIEtpbmQ6IHRyYWNrLmtpbmQsXG4gICAgICAgICAgICAgICAgTGFiZWw6IHRyYWNrLmxhYmVsLFxuICAgICAgICAgICAgICAgIE11dGVkOiB0cmFjay5tdXRlZCxcbiAgICAgICAgICAgICAgICBSZWFkeVN0YXRlOiB0cmFjay5yZWFkeVN0YXRlLFxuICAgICAgICAgICAgICAgIFJlbW90ZTogdHJhY2sucmVtb3RlLFxuICAgICAgICAgICAgICAgIEF1ZGlvU2V0dGluZ3M6IHtcbiAgICAgICAgICAgICAgICAgICAgQ2hhbm5lbENvdW50OiBzZXR0aW5ncy5jaGFubmVsQ291bnQuSW50KCksXG4gICAgICAgICAgICAgICAgICAgIEVjaG9DYW5jZWxsYXRpb246IHNldHRpbmdzLmVjaG9DYW5jZWxsYXRpb24sXG4gICAgICAgICAgICAgICAgICAgIExhdGVuY3k6IHNldHRpbmdzLmxhdGVuY3ksXG4gICAgICAgICAgICAgICAgICAgIFNhbXBsZVJhdGU6IHNldHRpbmdzLnNhbXBsZVJhdGUuSW50NjQoKSxcbiAgICAgICAgICAgICAgICAgICAgU2FtcGxlU2l6ZTogc2V0dGluZ3Muc2FtcGxlU2l6ZS5JbnQ2NCgpLFxuICAgICAgICAgICAgICAgICAgICBWb2x1bWU6IHNldHRpbmdzLnZvbHVtZSxcbiAgICAgICAgICAgICAgICAgICAgTWVkaWFUcmFja1NldHRpbmdzOiB7XG4gICAgICAgICAgICAgICAgICAgICAgICBEZXZpY2VJRDogc2V0dGluZ3MuZGV2aWNlSWQsXG4gICAgICAgICAgICAgICAgICAgICAgICBHcm91cElEOiBzZXR0aW5ncy5ncm91cElkLFxuICAgICAgICAgICAgICAgICAgICB9LFxuICAgICAgICAgICAgICAgIH0sXG4gICAgICAgICAgICB9KTtcbiAgICAgICAgfVxuICAgIH1cbiAgICBsZXQgdmlkZW9zVHJhY2tzID0gby5nZXRWaWRlb1RyYWNrcygpO1xuICAgIGlmICh2aWRlb3NUcmFja3MgIT09IG51bGwgJiYgdmlkZW9zVHJhY2tzICE9PSB1bmRlZmluZWQpIHtcbiAgICAgICAgZm9yIChsZXQgaSA9IDA7IGkgPCB2aWRlb3NUcmFja3MubGVuZ3RoOyBpKyspIHtcbiAgICAgICAgICAgIGxldCB0cmFjayA9IHZpZGVvc1RyYWNrc1tpXTtcbiAgICAgICAgICAgIGxldCBzZXR0aW5ncyA9IHRyYWNrLmdldFNldHRpbmdzKCk7XG4gICAgICAgICAgICBzdHJlYW0uVmlkZW9zLnB1c2goe1xuICAgICAgICAgICAgICAgIEVuYWJsZWQ6IHRyYWNrLmVuYWJsZWQsXG4gICAgICAgICAgICAgICAgSUQ6IHRyYWNrLmlkLFxuICAgICAgICAgICAgICAgIEtpbmQ6IHRyYWNrLmtpbmQsXG4gICAgICAgICAgICAgICAgTGFiZWw6IHRyYWNrLmxhYmVsLFxuICAgICAgICAgICAgICAgIE11dGVkOiB0cmFjay5tdXRlZCxcbiAgICAgICAgICAgICAgICBSZWFkeVN0YXRlOiB0cmFjay5yZWFkeVN0YXRlLFxuICAgICAgICAgICAgICAgIFJlbW90ZTogdHJhY2sucmVtb3RlLFxuICAgICAgICAgICAgICAgIFZpZGVvU2V0dGluZ3M6IHtcbiAgICAgICAgICAgICAgICAgICAgQXNwZWN0UmF0aW86IHNldHRpbmdzLmFzcGVjdFJhdGlvbixcbiAgICAgICAgICAgICAgICAgICAgRnJhbWVSYXRlOiBzZXR0aW5ncy5mcmFtZVJhdGUsXG4gICAgICAgICAgICAgICAgICAgIEhlaWdodDogc2V0dGluZ3MuaGVpZ2h0LkludDY0KCksXG4gICAgICAgICAgICAgICAgICAgIFdpZHRoOiBzZXR0aW5ncy53aWR0aC5JbnQ2NCgpLFxuICAgICAgICAgICAgICAgICAgICBGYWNpbmdNb2RlOiBzZXR0aW5ncy5mYWNpbmdNb2RlLFxuICAgICAgICAgICAgICAgICAgICBNZWRpYVRyYWNrU2V0dGluZ3M6IHtcbiAgICAgICAgICAgICAgICAgICAgICAgIERldmljZUlEOiBzZXR0aW5ncy5kZXZpY2VJZCxcbiAgICAgICAgICAgICAgICAgICAgICAgIEdyb3VwSUQ6IHNldHRpbmdzLmdyb3VwSWQsXG4gICAgICAgICAgICAgICAgICAgIH0sXG4gICAgICAgICAgICAgICAgfSxcbiAgICAgICAgICAgIH0pO1xuICAgICAgICB9XG4gICAgfVxuICAgIHJldHVybiBzdHJlYW07XG59XG5leHBvcnRzLnRvTWVkaWFTdHJlYW0gPSB0b01lZGlhU3RyZWFtO1xuZnVuY3Rpb24gdG9Ub3VjaGVzKG8pIHtcbiAgICBpZiAobyA9PT0gbnVsbCB8fCBvID09PSB1bmRlZmluZWQpIHtcbiAgICAgICAgcmV0dXJuO1xuICAgIH1cbiAgICBjb25zdCB0b3VjaGVzID0gW107XG4gICAgZm9yIChsZXQgaSA9IDA7IGkgPCBvLmxlbmd0aDsgaSsrKSB7XG4gICAgICAgIGxldCBldiA9IG8uaXRlbShpKTtcbiAgICAgICAgdG91Y2hlcy5wdXNoKHtcbiAgICAgICAgICAgIENsaWVudFg6IGV2LmNsaWVudFgsXG4gICAgICAgICAgICBDbGllbnRZOiBldi5jbGllbnRZLFxuICAgICAgICAgICAgT2Zmc2V0WDogZXYub2Zmc2V0WCxcbiAgICAgICAgICAgIE9mZnNldFk6IGV2Lm9mZnNldFksXG4gICAgICAgICAgICBQYWdlWDogZXYucGFnZVgsXG4gICAgICAgICAgICBQYWdlWTogZXYucGFnZVksXG4gICAgICAgICAgICBTY3JlZW5YOiBldi5zY3JlZW5YLFxuICAgICAgICAgICAgU2NyZWVuWTogZXYuc2NyZWVuWSxcbiAgICAgICAgICAgIElkZW50aWZpZXI6IGV2LmlkZW50aWZpZXIsXG4gICAgICAgIH0pO1xuICAgIH1cbiAgICByZXR1cm4gdG91Y2hlcztcbn1cbmV4cG9ydHMudG9Ub3VjaGVzID0gdG9Ub3VjaGVzO1xuZnVuY3Rpb24gdG9HYW1lcGFkKG8pIHtcbiAgICBsZXQgcGFkID0ge307XG4gICAgaWYgKG8gPT09IG51bGwgfHwgbyA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIHJldHVybiBwYWQ7XG4gICAgfVxuICAgIHBhZC5EaXNwbGF5SUQgPSBvLmRpc3BsYXlJZDtcbiAgICBwYWQuSUQgPSBvLmlkO1xuICAgIHBhZC5JbmRleCA9IG8uaW5kZXguSW50KCk7XG4gICAgcGFkLk1hcHBpbmcgPSBvLm1hcHBpbmc7XG4gICAgcGFkLkNvbm5lY3RlZCA9IG8uY29ubmVjdGVkO1xuICAgIHBhZC5UaW1lc3RhbXAgPSBvLnRpbWVzdGFtcDtcbiAgICBwYWQuQXhlcyA9IFtdO1xuICAgIHBhZC5CdXR0b25zID0gW107XG4gICAgbGV0IGF4ZXMgPSBvLmF4ZXM7XG4gICAgaWYgKGF4ZXMgIT09IG51bGwgJiYgYXhlcyAhPT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIGZvciAobGV0IGkgPSAwOyBpIDwgYXhlcy5sZW5ndGg7IGkrKykge1xuICAgICAgICAgICAgcGFkLkF4ZXMucHVzaChheGVzW2ldKTtcbiAgICAgICAgfVxuICAgIH1cbiAgICBsZXQgYnV0dG9ucyA9IG8uYnV0dG9ucztcbiAgICBpZiAoYnV0dG9ucyAhPT0gbnVsbCAmJiBidXR0b25zICE9PSB1bmRlZmluZWQpIHtcbiAgICAgICAgZm9yIChsZXQgaSA9IDA7IGkgPCBidXR0b25zLmxlbmd0aDsgaSsrKSB7XG4gICAgICAgICAgICBjb25zdCBidXR0b24gPSBidXR0b25zW2ldO1xuICAgICAgICAgICAgcGFkLkJ1dHRvbnMucHVzaCh7XG4gICAgICAgICAgICAgICAgVmFsdWU6IGJ1dHRvbi52YWx1ZSxcbiAgICAgICAgICAgICAgICBQcmVzc2VkOiBidXR0b24ucHJlc3NlZCxcbiAgICAgICAgICAgIH0pO1xuICAgICAgICB9XG4gICAgfVxuICAgIHJldHVybiBwYWQ7XG59XG5leHBvcnRzLnRvR2FtZXBhZCA9IHRvR2FtZXBhZDtcbmZ1bmN0aW9uIHRvRGF0YVRyYW5zZmVyKG8pIHtcbiAgICBpZiAobyA9PT0gbnVsbCB8fCBvID09PSB1bmRlZmluZWQpIHtcbiAgICAgICAgcmV0dXJuO1xuICAgIH1cbiAgICBsZXQgZHQgPSB7fTtcbiAgICBkdC5Ecm9wRWZmZWN0ID0gby5kcm9wRWZmZWN0O1xuICAgIGR0LkVmZmVjdEFsbG93ZWQgPSBvLmVmZmVjdEFsbG93ZWQ7XG4gICAgZHQuVHlwZXMgPSBvLnR5cGVzO1xuICAgIGR0Lkl0ZW1zID0gW107XG4gICAgY29uc3QgaXRlbXMgPSBvLml0ZW1zO1xuICAgIGlmIChpdGVtcyAhPT0gbnVsbCAmJiBpdGVtcyAhPT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIGZvciAobGV0IGkgPSAwOyBpIDwgaXRlbXMubGVuZ3RoOyBpKyspIHtcbiAgICAgICAgICAgIGNvbnN0IGl0ZW0gPSBpdGVtcy5EYXRhVHJhbnNmZXJJdGVtKGkpO1xuICAgICAgICAgICAgZHQuSXRlbXMucHVzaCh7XG4gICAgICAgICAgICAgICAgTmFtZTogaXRlbS5uYW1lLFxuICAgICAgICAgICAgICAgIFNpemU6IGl0ZW0uc2l6ZS5JbnQoKSxcbiAgICAgICAgICAgICAgICBEYXRhOiBmcm9tRmlsZShpdGVtKSxcbiAgICAgICAgICAgIH0pO1xuICAgICAgICB9XG4gICAgfVxuICAgIGR0LkZpbGVzID0gW107XG4gICAgY29uc3QgZmlsZXMgPSBvLmZpbGVzO1xuICAgIGlmIChmaWxlcyAhPT0gbnVsbCAmJiBmaWxlcyAhPT0gdW5kZWZpbmVkKSB7XG4gICAgICAgIGZvciAobGV0IGkgPSAwOyBpIDwgZmlsZXMubGVuZ3RoOyBpKyspIHtcbiAgICAgICAgICAgIGNvbnN0IGl0ZW0gPSBmaWxlc1tpXTtcbiAgICAgICAgICAgIGRGaWxlcy5wdXNoKHtcbiAgICAgICAgICAgICAgICBOYW1lOiBpdGVtLm5hbWUsXG4gICAgICAgICAgICAgICAgU2l6ZTogaXRlbS5zaXplLkludCgpLFxuICAgICAgICAgICAgICAgIERhdGE6IGZyb21GaWxlKGl0ZW0pLFxuICAgICAgICAgICAgfSk7XG4gICAgICAgIH1cbiAgICB9XG4gICAgcmV0dXJuIGR0O1xufVxuZXhwb3J0cy50b0RhdGFUcmFuc2ZlciA9IHRvRGF0YVRyYW5zZmVyO1xuZnVuY3Rpb24gZ2V0TWF0Y2hpbmdOb2RlKG1hdGNoTm9kZSwgbWF0Y2hlcikge1xuICAgIGlmICghbWF0Y2hOb2RlKSB7XG4gICAgICAgIHJldHVybiBudWxsO1xuICAgIH1cbiAgICBpZiAobWF0Y2hlcihtYXRjaE5vZGUpKSB7XG4gICAgICAgIHJldHVybiBtYXRjaE5vZGU7XG4gICAgfVxuICAgIHdoaWxlICgobWF0Y2hOb2RlID0gbWF0Y2hOb2RlLm5leHRTaWJsaW5nKSkge1xuICAgICAgICBpZiAobWF0Y2hlcihtYXRjaE5vZGUpKSB7XG4gICAgICAgICAgICByZXR1cm4gbWF0Y2hOb2RlO1xuICAgICAgICB9XG4gICAgfVxuICAgIHJldHVybiBudWxsO1xufVxuZXhwb3J0cy5nZXRNYXRjaGluZ05vZGUgPSBnZXRNYXRjaGluZ05vZGU7XG4vLyMgc291cmNlTWFwcGluZ1VSTD1kb20uanMubWFwIiwiXCJ1c2Ugc3RyaWN0XCI7XG5PYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgXCJfX2VzTW9kdWxlXCIsIHsgdmFsdWU6IHRydWUgfSk7XG5leHBvcnRzLk9iamVjdHMgPSB2b2lkIDA7XG52YXIgT2JqZWN0cztcbihmdW5jdGlvbiAoT2JqZWN0cykge1xuICAgIGZ1bmN0aW9uIFBhdGNoV2l0aChlbGVtLCBhdHRyTmFtZSwgYXR0cnMpIHtcbiAgICAgICAgZWxlbVthdHRyTmFtZV0gPSBhdHRycztcbiAgICB9XG4gICAgT2JqZWN0cy5QYXRjaFdpdGggPSBQYXRjaFdpdGg7XG4gICAgZnVuY3Rpb24gR2V0QXR0cldpdGgoZWxlbSwgYXR0ck5hbWUpIHtcbiAgICAgICAgY29uc3QgdmFsdWUgPSBlbGVtW2F0dHJOYW1lXTtcbiAgICAgICAgcmV0dXJuIHZhbHVlID8gdmFsdWUgOiAnJztcbiAgICB9XG4gICAgT2JqZWN0cy5HZXRBdHRyV2l0aCA9IEdldEF0dHJXaXRoO1xuICAgIGZ1bmN0aW9uIGlzTnVsbE9yVW5kZWZpbmVkKGVsZW0pIHtcbiAgICAgICAgcmV0dXJuIGVsZW0gPT09IG51bGwgfHwgZWxlbSA9PT0gdW5kZWZpbmVkO1xuICAgIH1cbiAgICBPYmplY3RzLmlzTnVsbE9yVW5kZWZpbmVkID0gaXNOdWxsT3JVbmRlZmluZWQ7XG4gICAgZnVuY3Rpb24gaXNBbnkoZWxlbSwgLi4udmFsdWVzKSB7XG4gICAgICAgIGZvciAobGV0IGluZGV4IG9mIHZhbHVlcykge1xuICAgICAgICAgICAgaWYgKGVsZW0gPT09IGluZGV4KSB7XG4gICAgICAgICAgICAgICAgcmV0dXJuIHRydWU7XG4gICAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cbiAgICBPYmplY3RzLmlzQW55ID0gaXNBbnk7XG59KShPYmplY3RzID0gZXhwb3J0cy5PYmplY3RzIHx8IChleHBvcnRzLk9iamVjdHMgPSB7fSkpO1xuLy8jIHNvdXJjZU1hcHBpbmdVUkw9ZXh0ZW5zaW9ucy5qcy5tYXAiLCJcInVzZSBzdHJpY3RcIjtcbk9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBcIl9fZXNNb2R1bGVcIiwgeyB2YWx1ZTogdHJ1ZSB9KTtcbmV4cG9ydHMubW91bnRUbyA9IHZvaWQgMDtcbmNvbnN0IHV0aWxzID0gcmVxdWlyZShcIi4vdXRpbHNcIik7XG5jb25zdCBleHRzID0gcmVxdWlyZShcIi4vZXh0ZW5zaW9uc1wiKTtcbmNvbnN0IHBhdGNoID0gcmVxdWlyZShcIi4vcGF0Y2hcIik7XG5jb25zdCBtb3VudCA9IHJlcXVpcmUoXCIuL21vdW50XCIpO1xuY29uc3QgZG9tID0gcmVxdWlyZShcIi4vZG9tXCIpO1xuY29uc3QgbmFtZXNwYWNlID0gT2JqZWN0LmFzc2lnbihPYmplY3QuYXNzaWduKE9iamVjdC5hc3NpZ24oT2JqZWN0LmFzc2lnbihPYmplY3QuYXNzaWduKHt9LCB1dGlscyksIGV4dHMpLCBwYXRjaCksIGRvbSksIG1vdW50KTtcbmZ1bmN0aW9uIG1vdW50VG8ocGFyZW50KSB7XG4gICAgcGFyZW50Lm1hcmt1cCA9IG5hbWVzcGFjZTtcbn1cbmV4cG9ydHMubW91bnRUbyA9IG1vdW50VG87XG5leHBvcnRzLmRlZmF1bHQgPSBuYW1lc3BhY2U7XG4vLyMgc291cmNlTWFwcGluZ1VSTD1pbmRleC5qcy5tYXAiLCJcInVzZSBzdHJpY3RcIjtcbk9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBcIl9fZXNNb2R1bGVcIiwgeyB2YWx1ZTogdHJ1ZSB9KTtcbmV4cG9ydHMuRE9NTW91bnQgPSB2b2lkIDA7XG5jb25zdCBkb20gPSByZXF1aXJlKFwiLi9kb21cIik7XG5jb25zdCBwYXRjaCA9IHJlcXVpcmUoXCIuL3BhdGNoXCIpO1xuY29uc3QgdXRpbHMgPSByZXF1aXJlKFwiLi91dGlsc1wiKTtcbmNsYXNzIERPTU1vdW50IHtcbiAgICBjb25zdHJ1Y3Rvcihkb2N1bWVudCwgdGFyZ2V0LCBub3RpZmllcikge1xuICAgICAgICBpZiAodHlwZW9mIGRvY3VtZW50ICE9PSAnb2JqZWN0Jykge1xuICAgICAgICAgICAgdGhyb3cgbmV3IEVycm9yKCdkb2N1bWVudCBzaG91bGQgYmUgYW4gb2JqZWN0Jyk7XG4gICAgICAgIH1cbiAgICAgICAgdGhpcy5kb2MgPSBkb2N1bWVudDtcbiAgICAgICAgdGhpcy5ub3RpZmllciA9IG5vdGlmaWVyO1xuICAgICAgICB0aGlzLmV2ZW50cyA9IHt9O1xuICAgICAgICB0aGlzLmhhbmRsZXIgPSB0aGlzLmhhbmRsZUV2ZW50LmJpbmQodGhpcyk7XG4gICAgICAgIGlmICh0eXBlb2YgdGFyZ2V0ID09PSAnc3RyaW5nJykge1xuICAgICAgICAgICAgY29uc3QgdGFyZ2V0U2VsZWN0b3IgPSB0YXJnZXQ7XG4gICAgICAgICAgICBjb25zdCBub2RlID0gdGhpcy5kb2MucXVlcnlTZWxlY3Rvcih0YXJnZXRTZWxlY3Rvcik7XG4gICAgICAgICAgICBpZiAobm9kZSA9PT0gbnVsbCB8fCBub2RlID09PSB1bmRlZmluZWQpIHtcbiAgICAgICAgICAgICAgICB0aHJvdyBuZXcgRXJyb3IoYHVuYWJsZSB0byBsb2NhdGUgbm9kZSBmb3IgJHt7IHRhcmdldFNlbGVjdG9yIH19YCk7XG4gICAgICAgICAgICB9XG4gICAgICAgICAgICB0aGlzLm1vdW50Tm9kZSA9IG5vZGU7XG4gICAgICAgIH1cbiAgICAgICAgZWxzZSB7XG4gICAgICAgICAgICB0aGlzLm1vdW50Tm9kZSA9IHRhcmdldDtcbiAgICAgICAgfVxuICAgIH1cbiAgICBoYW5kbGVFdmVudChldmVudCkge1xuICAgICAgICBpZiAoIXRoaXMuZXZlbnRzW2V2ZW50LnR5cGVdKSB7XG4gICAgICAgICAgICByZXR1cm47XG4gICAgICAgIH1cbiAgICAgICAgZXZlbnQuc3RvcFByb3BhZ2F0aW9uKCk7XG4gICAgICAgIGNvbnN0IHRhcmdldCA9IGV2ZW50LnRhcmdldDtcbiAgICAgICAgaWYgKHRhcmdldC5ub2RlVHlwZSAhPT0gZG9tLkVMRU1FTlRfTk9ERSkge1xuICAgICAgICAgICAgcmV0dXJuO1xuICAgICAgICB9XG4gICAgICAgIGNvbnN0IHRhcmdldEVsZW1lbnQgPSB0YXJnZXQ7XG4gICAgICAgIGNvbnN0IGtlYmFiRXZlbnROYW1lID0gJ2V2ZW50LScgKyB1dGlscy5Ub0tlYmFiQ2FzZShldmVudC50eXBlKTtcbiAgICAgICAgaWYgKCF0YXJnZXRFbGVtZW50Lmhhc0F0dHJpYnV0ZShrZWJhYkV2ZW50TmFtZSkpIHtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICBjb25zdCB0cmlnZ2VycyA9IHRhcmdldEVsZW1lbnQuZ2V0QXR0cmlidXRlKGtlYmFiRXZlbnROYW1lKTtcbiAgICAgICAgaWYgKHRyaWdnZXJzID09IG51bGwpIHtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICBpZiAodGhpcy5ub3RpZmllciAmJiB0eXBlb2YgdGhpcy5ub3RpZmllciA9PT0gJ2Z1bmN0aW9uJykge1xuICAgICAgICAgICAgdGhpcy5ub3RpZmllcihldmVudCwgdHJpZ2dlcnMuc3BsaXQoJ3wnKSwgdGFyZ2V0RWxlbWVudCk7XG4gICAgICAgIH1cbiAgICB9XG4gICAgcGF0Y2goY2hhbmdlKSB7XG4gICAgICAgIHRoaXMucGF0Y2hXaXRoKGNoYW5nZSwgcGF0Y2guRGVmYXVsdE5vZGVEaWN0YXRvciwgcGF0Y2guRGVmYXVsdEpTT05EaWN0YXRvciwgcGF0Y2guRGVmYXVsdEpTT05NYWtlcik7XG4gICAgfVxuICAgIHBhdGNoV2l0aChjaGFuZ2UsIG5vZGVEaWN0YXRvciwganNvbkRpY3RhdG9yLCBqc29uTWFrZXIpIHtcbiAgICAgICAgaWYgKGNoYW5nZSBpbnN0YW5jZW9mIERvY3VtZW50RnJhZ21lbnQpIHtcbiAgICAgICAgICAgIGNvbnN0IGZyYWdtZW50ID0gY2hhbmdlO1xuICAgICAgICAgICAgdGhpcy5yZWdpc3Rlck5vZGVFdmVudHMoZnJhZ21lbnQpO1xuICAgICAgICAgICAgcGF0Y2guUGF0Y2hET01UcmVlKGZyYWdtZW50LCB0aGlzLm1vdW50Tm9kZSwgbm9kZURpY3RhdG9yLCBmYWxzZSk7XG4gICAgICAgICAgICByZXR1cm47XG4gICAgICAgIH1cbiAgICAgICAgaWYgKHR5cGVvZiBjaGFuZ2UgPT09ICdzdHJpbmcnKSB7XG4gICAgICAgICAgICBjb25zdCBub2RlID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnZGl2Jyk7XG4gICAgICAgICAgICBub2RlLmlubmVySFRNTCA9IGNoYW5nZS50cmltKCk7XG4gICAgICAgICAgICB0aGlzLnJlZ2lzdGVyTm9kZUV2ZW50cyhub2RlKTtcbiAgICAgICAgICAgIHBhdGNoLlBhdGNoRE9NVHJlZShub2RlLCB0aGlzLm1vdW50Tm9kZSwgbm9kZURpY3RhdG9yLCBmYWxzZSk7XG4gICAgICAgICAgICByZXR1cm47XG4gICAgICAgIH1cbiAgICAgICAgaWYgKCFwYXRjaC5pc0pTT05Ob2RlKGNoYW5nZSkpIHtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICBjb25zdCBub2RlID0gY2hhbmdlO1xuICAgICAgICB0aGlzLnJlZ2lzdGVySlNPTk5vZGVFdmVudHMobm9kZSk7XG4gICAgICAgIHBhdGNoLlBhdGNoSlNPTk5vZGVUcmVlKG5vZGUsIHRoaXMubW91bnROb2RlLCBqc29uRGljdGF0b3IsIGpzb25NYWtlcik7XG4gICAgfVxuICAgIHBhdGNoTGlzdChjaGFuZ2VzKSB7XG4gICAgICAgIGNoYW5nZXMuZm9yRWFjaCh0aGlzLnBhdGNoLmJpbmQodGhpcykpO1xuICAgIH1cbiAgICBzdHJlYW0oY2hhbmdlcykge1xuICAgICAgICBjb25zdCBub2RlcyA9IEpTT04ucGFyc2UoY2hhbmdlcyk7XG4gICAgICAgIHJldHVybiB0aGlzLnN0cmVhbUxpc3Qobm9kZXMpO1xuICAgIH1cbiAgICBzdHJlYW1MaXN0KGNoYW5nZXMpIHtcbiAgICAgICAgdGhpcy5zdHJlYW1MaXN0V2l0aChjaGFuZ2VzLCBwYXRjaC5EZWZhdWx0SlNPTkRpY3RhdG9yLCBwYXRjaC5EZWZhdWx0SlNPTk1ha2VyKTtcbiAgICB9XG4gICAgc3RyZWFtTGlzdFdpdGgoY2hhbmdlcywgZGljdGF0b3IsIG1ha2VyKSB7XG4gICAgICAgIGNoYW5nZXMuZm9yRWFjaCh0aGlzLnJlZ2lzdGVySlNPTk5vZGVFdmVudHMuYmluZCh0aGlzKSk7XG4gICAgICAgIHBhdGNoLlN0cmVhbUpTT05Ob2RlcyhjaGFuZ2VzLCB0aGlzLm1vdW50Tm9kZSwgcGF0Y2guRGVmYXVsdEpTT05EaWN0YXRvciwgcGF0Y2guRGVmYXVsdEpTT05NYWtlcik7XG4gICAgfVxuICAgIHJlZ2lzdGVyTm9kZUV2ZW50cyhuKSB7XG4gICAgICAgIGNvbnN0IGJpbmRlciA9IHRoaXM7XG4gICAgICAgIGRvbS5hcHBseUVhY2hOb2RlKG4sIGZ1bmN0aW9uIChub2RlKSB7XG4gICAgICAgICAgICBpZiAobm9kZS5ub2RlVHlwZSAhPT0gZG9tLkVMRU1FTlRfTk9ERSkge1xuICAgICAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGNvbnN0IGVsZW0gPSBub2RlO1xuICAgICAgICAgICAgY29uc3QgZXZlbnRzID0gZWxlbS5nZXRBdHRyaWJ1dGUoJ2V2ZW50cycpO1xuICAgICAgICAgICAgaWYgKGV2ZW50cykge1xuICAgICAgICAgICAgICAgIGV2ZW50cy5zcGxpdCgnICcpLmZvckVhY2goZnVuY3Rpb24gKGRlc2MpIHtcbiAgICAgICAgICAgICAgICAgICAgY29uc3QgZXZlbnROYW1lID0gZGVzYy5zdWJzdHIoMCwgZGVzYy5sZW5ndGggLSAzKTtcbiAgICAgICAgICAgICAgICAgICAgYmluZGVyLnJlZ2lzdGVyRXZlbnQoZXZlbnROYW1lKTtcbiAgICAgICAgICAgICAgICAgICAgc3dpdGNoIChkZXNjLnN1YnN0cihkZXNjLmxlbmd0aCAtIDIsIGRlc2MubGVuZ3RoKSkge1xuICAgICAgICAgICAgICAgICAgICAgICAgY2FzZSAnMDEnOlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICAgICAgICAgICAgICAgICAgY2FzZSAnMTAnOlxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIG4uYWRkRXZlbnRMaXN0ZW5lcihldmVudE5hbWUsIERPTU1vdW50LnByZXZlbnREZWZhdWx0LCBmYWxzZSk7XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgICAgICAgICAgICAgICAgICBjYXNlICcxMSc6XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgbi5hZGRFdmVudExpc3RlbmVyKGV2ZW50TmFtZSwgRE9NTW91bnQucHJldmVudERlZmF1bHQsIGZhbHNlKTtcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBicmVhaztcbiAgICAgICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfVxuICAgICAgICB9KTtcbiAgICB9XG4gICAgcmVnaXN0ZXJKU09OTm9kZUV2ZW50cyhub2RlKSB7XG4gICAgICAgIGNvbnN0IGJpbmRlciA9IHRoaXM7XG4gICAgICAgIHBhdGNoLmFwcGx5SlNPTk5vZGVGdW5jdGlvbihub2RlLCBmdW5jdGlvbiAobikge1xuICAgICAgICAgICAgaWYgKG4ucmVtb3ZlZCkge1xuICAgICAgICAgICAgICAgIG4uZXZlbnRzLmZvckVhY2goZnVuY3Rpb24gKGRlc2MpIHtcbiAgICAgICAgICAgICAgICAgICAgYmluZGVyLnVucmVnaXN0ZXJFdmVudChkZXNjLk5hbWUpO1xuICAgICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIG4uZXZlbnRzLmZvckVhY2goZnVuY3Rpb24gKGRlc2MpIHtcbiAgICAgICAgICAgICAgICBiaW5kZXIucmVnaXN0ZXJFdmVudChkZXNjLk5hbWUpO1xuICAgICAgICAgICAgfSk7XG4gICAgICAgIH0pO1xuICAgIH1cbiAgICB0ZXh0Q29udGVudCgpIHtcbiAgICAgICAgcmV0dXJuIHRoaXMubW91bnROb2RlLnRleHRDb250ZW50O1xuICAgIH1cbiAgICBpbm5lckhUTUwoKSB7XG4gICAgICAgIHJldHVybiB0aGlzLm1vdW50Tm9kZS5pbm5lckhUTUwudHJpbSgpO1xuICAgIH1cbiAgICBIdG1sKCkge1xuICAgICAgICByZXR1cm4gZG9tLnRvSFRNTCh0aGlzLm1vdW50Tm9kZSwgZmFsc2UpO1xuICAgIH1cbiAgICByZWdpc3RlckV2ZW50KGV2ZW50TmFtZSkge1xuICAgICAgICBpZiAodGhpcy5ldmVudHNbZXZlbnROYW1lXSkge1xuICAgICAgICAgICAgcmV0dXJuO1xuICAgICAgICB9XG4gICAgICAgIHRoaXMubW91bnROb2RlLmFkZEV2ZW50TGlzdGVuZXIoZXZlbnROYW1lLCB0aGlzLmhhbmRsZXIsIHRydWUpO1xuICAgICAgICB0aGlzLmV2ZW50c1tldmVudE5hbWVdID0gdHJ1ZTtcbiAgICB9XG4gICAgdW5yZWdpc3RlckV2ZW50KGV2ZW50TmFtZSkge1xuICAgICAgICBpZiAoIXRoaXMuZXZlbnRzW2V2ZW50TmFtZV0pIHtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICB0aGlzLm1vdW50Tm9kZS5yZW1vdmVFdmVudExpc3RlbmVyKGV2ZW50TmFtZSwgdGhpcy5oYW5kbGVyLCB0cnVlKTtcbiAgICAgICAgdGhpcy5ldmVudHNbZXZlbnROYW1lXSA9IGZhbHNlO1xuICAgIH1cbiAgICBzdGF0aWMgcHJldmVudERlZmF1bHQoZXZlbnQpIHtcbiAgICAgICAgZXZlbnQucHJldmVudERlZmF1bHQoKTtcbiAgICB9XG4gICAgc3RhdGljIHN0b3BQcm9wYWdhdGlvbihldmVudCkge1xuICAgICAgICBldmVudC5zdG9wUHJvcGFnYXRpb24oKTtcbiAgICB9XG59XG5leHBvcnRzLkRPTU1vdW50ID0gRE9NTW91bnQ7XG4vLyMgc291cmNlTWFwcGluZ1VSTD1tb3VudC5qcy5tYXAiLCJcInVzZSBzdHJpY3RcIjtcbk9iamVjdC5kZWZpbmVQcm9wZXJ0eShleHBvcnRzLCBcIl9fZXNNb2R1bGVcIiwgeyB2YWx1ZTogdHJ1ZSB9KTtcbmV4cG9ydHMuUGF0Y2hET01BdHRyaWJ1dGVzID0gZXhwb3J0cy5QYXRjaERPTVRyZWUgPSBleHBvcnRzLlBhdGNoSlNPTkF0dHJpYnV0ZXMgPSBleHBvcnRzLlBhdGNoVGV4dENvbW1lbnRXaXRoSlNPTiA9IGV4cG9ydHMuQXBwbHlTdHJlYW1Ob2RlID0gZXhwb3J0cy5TdHJlYW1KU09OTm9kZXMgPSBleHBvcnRzLlBhdGNoSlNPTk5vZGUgPSBleHBvcnRzLlBhdGNoSlNPTk5vZGVUcmVlID0gZXhwb3J0cy5qc29uTWFrZXIgPSBleHBvcnRzLkRlZmF1bHRKU09OTWFrZXIgPSBleHBvcnRzLmZpbmRFbGVtZW50UGFyZW50YnlSZWYgPSBleHBvcnRzLmZpbmRFbGVtZW50YnlSZWYgPSBleHBvcnRzLmZpbmRFbGVtZW50ID0gZXhwb3J0cy5pc0pTT05Ob2RlID0gZXhwb3J0cy5hcHBseUpTT05Ob2RlS2lkc0Z1bmN0aW9uID0gZXhwb3J0cy5hcHBseUpTT05Ob2RlRnVuY3Rpb24gPSBleHBvcnRzLkpTT05FdmVudCA9IGV4cG9ydHMuTm9kZVRvSlNPTk5vZGUgPSBleHBvcnRzLlRvSlNPTk5vZGUgPSBleHBvcnRzLkRlZmF1bHRKU09ORGljdGF0b3IgPSBleHBvcnRzLkRlZmF1bHROb2RlRGljdGF0b3IgPSB2b2lkIDA7XG5jb25zdCBkb20gPSByZXF1aXJlKFwiLi9kb21cIik7XG5jb25zdCB1dGlscyA9IHJlcXVpcmUoXCIuL3V0aWxzXCIpO1xuY29uc3QgZXh0cyA9IHJlcXVpcmUoXCIuL2V4dGVuc2lvbnNcIik7XG5jb25zdCBkb21fMSA9IHJlcXVpcmUoXCIuL2RvbVwiKTtcbmV4cG9ydHMuRGVmYXVsdE5vZGVEaWN0YXRvciA9IHtcbiAgICBTYW1lOiAobiwgbSkgPT4ge1xuICAgICAgICBjb25zdCBzYW1lTm9kZSA9IG4ubm9kZVR5cGUgPT09IG0ubm9kZVR5cGUgJiYgbi5ub2RlTmFtZSA9PT0gbS5ub2RlTmFtZTtcbiAgICAgICAgaWYgKCFzYW1lTm9kZSkge1xuICAgICAgICAgICAgcmV0dXJuIGZhbHNlO1xuICAgICAgICB9XG4gICAgICAgIGlmIChuLm5vZGVUeXBlICE9PSBkb20uRUxFTUVOVF9OT0RFKSB7XG4gICAgICAgICAgICByZXR1cm4gdHJ1ZTtcbiAgICAgICAgfVxuICAgICAgICBjb25zdCBuRWxlbSA9IG47XG4gICAgICAgIGNvbnN0IG1FbGVtID0gbTtcbiAgICAgICAgcmV0dXJuIG5FbGVtLmlkID09PSBtRWxlbS5pZDtcbiAgICB9LFxuICAgIENoYW5nZWQ6IChuLCBtKSA9PiB7XG4gICAgICAgIGlmIChuLm5vZGVUeXBlID09PSBkb20uVEVYVF9OT0RFICYmIG0ubm9kZVR5cGUgPT09IGRvbS5URVhUX05PREUpIHtcbiAgICAgICAgICAgIHJldHVybiBuLnRleHRDb250ZW50ID09PSBtLnRleHRDb250ZW50O1xuICAgICAgICB9XG4gICAgICAgIGlmIChuLm5vZGVUeXBlID09PSBkb20uQ09NTUVOVF9OT0RFICYmIG0ubm9kZVR5cGUgPT09IGRvbS5DT01NRU5UX05PREUpIHtcbiAgICAgICAgICAgIHJldHVybiBuLnRleHRDb250ZW50ID09PSBtLnRleHRDb250ZW50O1xuICAgICAgICB9XG4gICAgICAgIGNvbnN0IG5FbGVtID0gbjtcbiAgICAgICAgY29uc3QgbUVsZW0gPSBtO1xuICAgICAgICBjb25zdCBuQXR0ciA9IGRvbS5yZWNvcmRBdHRyaWJ1dGVzKG5FbGVtKTtcbiAgICAgICAgY29uc3QgbUF0dHIgPSBkb20ucmVjb3JkQXR0cmlidXRlcyhtRWxlbSk7XG4gICAgICAgIGlmICghbkF0dHIuaGFzT3duUHJvcGVydHkoJ190aWQnKSAmJiBtQXR0ci5oYXNPd25Qcm9wZXJ0eSgnX3RpZCcpKSB7XG4gICAgICAgICAgICByZXR1cm4gdHJ1ZTtcbiAgICAgICAgfVxuICAgICAgICBpZiAobkF0dHIuaGFzT3duUHJvcGVydHkoJ19hdGlkJykpIHtcbiAgICAgICAgICAgIHJldHVybiAhKG1BdHRyLl9hdGlkID09PSBuQXR0ci5fdGlkICYmIG1BdHRyLl90aWQgPT09IG1BdHRyLl9hdGlkKTtcbiAgICAgICAgfVxuICAgICAgICBkZWxldGUgbUF0dHJbJ19hdGlkJ107XG4gICAgICAgIGRlbGV0ZSBuQXR0clsnX2F0aWQnXTtcbiAgICAgICAgZGVsZXRlIG1BdHRyWydfdGlkJ107XG4gICAgICAgIGRlbGV0ZSBuQXR0clsnX3RpZCddO1xuICAgICAgICBmb3IgKGxldCBrZXkgaW4gbUF0dHIpIHtcbiAgICAgICAgICAgIGlmICghbkF0dHIuaGFzT3duUHJvcGVydHkoa2V5KSkge1xuICAgICAgICAgICAgICAgIHJldHVybiB0cnVlO1xuICAgICAgICAgICAgfVxuICAgICAgICAgICAgaWYgKG1BdHRyW2tleV0gIT09IG5BdHRyW2tleV0pIHtcbiAgICAgICAgICAgICAgICByZXR1cm4gdHJ1ZTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICByZXR1cm4gbkVsZW0uaW5uZXJIVE1MICE9PSBtRWxlbS5pbm5lckhUTUw7XG4gICAgfSxcbn07XG5leHBvcnRzLkRlZmF1bHRKU09ORGljdGF0b3IgPSB7XG4gICAgU2FtZTogKG4sIG0pID0+IHtcbiAgICAgICAgY29uc3Qgc2FtZU5vZGUgPSBuLm5vZGVUeXBlID09PSBtLnR5cGUgJiYgbi5ub2RlTmFtZSA9PT0gbS5uYW1lO1xuICAgICAgICBpZiAoIXNhbWVOb2RlKSB7XG4gICAgICAgICAgICByZXR1cm4gZmFsc2U7XG4gICAgICAgIH1cbiAgICAgICAgaWYgKG4ubm9kZVR5cGUgIT09IGRvbS5FTEVNRU5UX05PREUpIHtcbiAgICAgICAgICAgIHJldHVybiB0cnVlO1xuICAgICAgICB9XG4gICAgICAgIGNvbnN0IG5FbGVtID0gbjtcbiAgICAgICAgcmV0dXJuIG5FbGVtLmlkID09PSBtLmlkO1xuICAgIH0sXG4gICAgQ2hhbmdlZDogKG4sIG0pID0+IHtcbiAgICAgICAgaWYgKG4ubm9kZVR5cGUgPT09IGRvbS5URVhUX05PREUgJiYgbS50eXBlID09PSBkb20uVEVYVF9OT0RFKSB7XG4gICAgICAgICAgICByZXR1cm4gbi50ZXh0Q29udGVudCAhPT0gbS5jb250ZW50O1xuICAgICAgICB9XG4gICAgICAgIGlmIChuLm5vZGVUeXBlID09PSBkb20uQ09NTUVOVF9OT0RFICYmIG0udHlwZSA9PT0gZG9tLkNPTU1FTlRfTk9ERSkge1xuICAgICAgICAgICAgcmV0dXJuIG4udGV4dENvbnRlbnQgIT09IG0uY29udGVudDtcbiAgICAgICAgfVxuICAgICAgICBjb25zdCB0bm9kZSA9IG47XG4gICAgICAgIGlmICh0bm9kZS5oYXNBdHRyaWJ1dGUoJ2lkJykpIHtcbiAgICAgICAgICAgIGNvbnN0IGlkID0gdG5vZGUuZ2V0QXR0cmlidXRlKCdpZCcpO1xuICAgICAgICAgICAgaWYgKGlkICE9PSBtLmlkKSB7XG4gICAgICAgICAgICAgICAgcmV0dXJuIHRydWU7XG4gICAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgICAgaWYgKHRub2RlLmhhc0F0dHJpYnV0ZSgnX3JlZicpKSB7XG4gICAgICAgICAgICBjb25zdCByZWYgPSB0bm9kZS5nZXRBdHRyaWJ1dGUoJ19yZWYnKTtcbiAgICAgICAgICAgIGlmIChyZWYgIT09IG0ucmVmKSB7XG4gICAgICAgICAgICAgICAgcmV0dXJuIHRydWU7XG4gICAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgICAgY29uc3QgdGlkID0gdG5vZGUuZ2V0QXR0cmlidXRlKCdfdGlkJyk7XG4gICAgICAgIGNvbnN0IGF0aWQgPSB0bm9kZS5nZXRBdHRyaWJ1dGUoJ19hdGlkJyk7XG4gICAgICAgIGlmICh0bm9kZS5oYXNBdHRyaWJ1dGUoJ190aWQnKSkge1xuICAgICAgICAgICAgaWYgKHRpZCAhPT0gbS50aWQpIHtcbiAgICAgICAgICAgICAgICByZXR1cm4gdHJ1ZTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGlmICh0bm9kZS5oYXNBdHRyaWJ1dGUoJ19hdGlkJykpIHtcbiAgICAgICAgICAgICAgICBpZiAodGlkICE9PSBtLnRpZCAmJiBhdGlkICE9PSBtLmF0aWQpIHtcbiAgICAgICAgICAgICAgICAgICAgcmV0dXJuIHRydWU7XG4gICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgICAgIGlmICh0aWQgIT09IG0udGlkICYmIGF0aWQgPT09IG0uYXRpZCkge1xuICAgICAgICAgICAgICAgICAgICByZXR1cm4gdHJ1ZTtcbiAgICAgICAgICAgICAgICB9XG4gICAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgICAgaWYgKCF0bm9kZS5oYXNBdHRyaWJ1dGUoJ2V2ZW50cycpICYmIG0uZXZlbnRzLmxlbmd0aCAhPT0gMCkge1xuICAgICAgICAgICAgcmV0dXJuIHRydWU7XG4gICAgICAgIH1cbiAgICAgICAgZm9yICh2YXIgaW5kZXggPSAwOyBpbmRleCA8IG0uZXZlbnRzLmxlbmd0aDsgaW5kZXgrKykge1xuICAgICAgICAgICAgbGV0IGV2ZW50ID0gbS5ldmVudHNbaW5kZXhdO1xuICAgICAgICAgICAgbGV0IGF0dHJOYW1lID0gJ2V2ZW50LScgKyBldmVudC5OYW1lO1xuICAgICAgICAgICAgbGV0IGF0dHJWYWx1ZSA9IGV2ZW50LlRhcmdldHMuam9pbignfCcpO1xuICAgICAgICAgICAgbGV0IG5vZGVBdHRyID0gdG5vZGUuYXR0cmlidXRlcy5nZXROYW1lZEl0ZW0oYXR0ck5hbWUpO1xuICAgICAgICAgICAgaWYgKG5vZGVBdHRyID09IG51bGwpIHtcbiAgICAgICAgICAgICAgICByZXR1cm4gdHJ1ZTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGlmIChub2RlQXR0ci52YWx1ZSAhPSBhdHRyVmFsdWUpIHtcbiAgICAgICAgICAgICAgICByZXR1cm4gdHJ1ZTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICByZXR1cm4gdHJ1ZTtcbiAgICB9LFxufTtcbmZ1bmN0aW9uIFRvSlNPTk5vZGUobm9kZSwgc2hhbGxvdywgcGFyZW50Tm9kZSkge1xuICAgIGNvbnN0IGxpc3QgPSBuZXcgQXJyYXkoKTtcbiAgICBpZiAodHlwZW9mIG5vZGUgPT09ICdzdHJpbmcnKSB7XG4gICAgICAgIGNvbnN0IHB1YiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuICAgICAgICBwdWIuaW5uZXJIVE1MID0gbm9kZS50cmltKCk7XG4gICAgICAgIGRvbS5hcHBseUNoaWxkTm9kZShwdWIsIGZ1bmN0aW9uIChjaGlsZCkge1xuICAgICAgICAgICAgbGlzdC5wdXNoKE5vZGVUb0pTT05Ob2RlKGNoaWxkLCBzaGFsbG93LCBwYXJlbnROb2RlKSk7XG4gICAgICAgIH0pO1xuICAgICAgICByZXR1cm4gbGlzdDtcbiAgICB9XG4gICAgbGlzdC5wdXNoKE5vZGVUb0pTT05Ob2RlKG5vZGUsIHNoYWxsb3csIHBhcmVudE5vZGUpKTtcbiAgICByZXR1cm4gbGlzdDtcbn1cbmV4cG9ydHMuVG9KU09OTm9kZSA9IFRvSlNPTk5vZGU7XG5mdW5jdGlvbiBOb2RlVG9KU09OTm9kZShub2RlLCBzaGFsbG93LCBwYXJlbnROb2RlKSB7XG4gICAgY29uc3Qgam5vZGUgPSB7fTtcbiAgICBqbm9kZS5jaGlsZHJlbiA9IFtdO1xuICAgIGpub2RlLmV2ZW50cyA9IFtdO1xuICAgIGpub2RlLmF0dHJzID0gW107XG4gICAgam5vZGUubmFtZXNwYWNlID0gJyc7XG4gICAgam5vZGUudHlwZSA9IG5vZGUubm9kZVR5cGU7XG4gICAgam5vZGUubmFtZSA9IG5vZGUubm9kZU5hbWUudG9Mb3dlckNhc2UoKTtcbiAgICBqbm9kZS5pZCA9IGV4dHMuT2JqZWN0cy5HZXRBdHRyV2l0aChub2RlLCAnX2lkJyk7XG4gICAgam5vZGUudGlkID0gZXh0cy5PYmplY3RzLkdldEF0dHJXaXRoKG5vZGUsICdfdGlkJyk7XG4gICAgam5vZGUucmVmID0gZXh0cy5PYmplY3RzLkdldEF0dHJXaXRoKG5vZGUsICdfcmVmJyk7XG4gICAgam5vZGUuYXRpZCA9IGV4dHMuT2JqZWN0cy5HZXRBdHRyV2l0aChub2RlLCAnX2F0aWQnKTtcbiAgICBjb25zdCBlbGVtID0gbm9kZTtcbiAgICBpZiAoZWxlbSA9PT0gbnVsbClcbiAgICAgICAgcmV0dXJuIGpub2RlO1xuICAgIGlmIChub2RlLl90aWQpIHtcbiAgICAgICAgam5vZGUudGlkID0gbm9kZS5fdGlkO1xuICAgIH1cbiAgICBzd2l0Y2ggKG5vZGUubm9kZVR5cGUpIHtcbiAgICAgICAgY2FzZSBkb21fMS5URVhUX05PREU6XG4gICAgICAgICAgICBqbm9kZS50eXBlTmFtZSA9ICdUZXh0JztcbiAgICAgICAgICAgIGpub2RlLmNvbnRlbnQgPSBub2RlLnRleHRDb250ZW50O1xuICAgICAgICAgICAgcmV0dXJuIGpub2RlO1xuICAgICAgICBjYXNlIGRvbV8xLkNPTU1FTlRfTk9ERTpcbiAgICAgICAgICAgIGpub2RlLnR5cGVOYW1lID0gJ0NvbW1lbnQnO1xuICAgICAgICAgICAgam5vZGUuY29udGVudCA9IG5vZGUudGV4dENvbnRlbnQ7XG4gICAgICAgICAgICByZXR1cm4gam5vZGU7XG4gICAgICAgIGNhc2UgZG9tXzEuRUxFTUVOVF9OT0RFOlxuICAgICAgICAgICAgam5vZGUudHlwZU5hbWUgPSAnRWxlbWVudCc7XG4gICAgICAgICAgICBqbm9kZS5jaGlsZHJlbiA9IG5ldyBBcnJheSgpO1xuICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgIGRlZmF1bHQ6XG4gICAgICAgICAgICB0aHJvdyBuZXcgRXJyb3IoYHVuYWJsZSB0byBoYW5kbGUgbm9kZSB0eXBlICR7bm9kZS5ub2RlVHlwZX1gKTtcbiAgICB9XG4gICAgaWYgKGV4dHMuT2JqZWN0cy5pc051bGxPclVuZGVmaW5lZChlbGVtKSkge1xuICAgICAgICBpZiAoam5vZGUuaWQgPT09ICcnKSB7XG4gICAgICAgICAgICBqbm9kZS5pZCA9IHV0aWxzLlJhbmRvbUlEKCk7XG4gICAgICAgIH1cbiAgICAgICAgcmV0dXJuIGpub2RlO1xuICAgIH1cbiAgICBpZiAoZWxlbS5oYXNBdHRyaWJ1dGUoJ2lkJykpIHtcbiAgICAgICAgam5vZGUuaWQgPSBlbGVtLmdldEF0dHJpYnV0ZSgnaWQnKTtcbiAgICB9XG4gICAgZWxzZSB7XG4gICAgICAgIGpub2RlLmlkID0gdXRpbHMuUmFuZG9tSUQoKTtcbiAgICB9XG4gICAgaWYgKGpub2RlLnJlZiA9PT0gJycgJiYgIWV4dHMuT2JqZWN0cy5pc051bGxPclVuZGVmaW5lZChwYXJlbnROb2RlKSkge1xuICAgICAgICBqbm9kZS5yZWYgPSBqbm9kZS5pZDtcbiAgICB9XG4gICAgaWYgKGVsZW0uaGFzQXR0cmlidXRlKCdfcmVmJykpIHtcbiAgICAgICAgam5vZGUucmVmID0gZWxlbS5nZXRBdHRyaWJ1dGUoJ19yZWYnKTtcbiAgICB9XG4gICAgZWxzZSB7XG4gICAgICAgIGlmICghZXh0cy5PYmplY3RzLmlzTnVsbE9yVW5kZWZpbmVkKHBhcmVudE5vZGUpKSB7XG4gICAgICAgICAgICBpZiAocGFyZW50Tm9kZS5yZWYgIT09ICcnKSB7XG4gICAgICAgICAgICAgICAgam5vZGUucmVmID0gcGFyZW50Tm9kZS5yZWYgKyAnLycgKyBqbm9kZS5pZDtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGVsc2Uge1xuICAgICAgICAgICAgICAgIGpub2RlLnJlZiA9IHBhcmVudE5vZGUuaWQgKyAnLycgKyBqbm9kZS5pZDtcbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgIH1cbiAgICBpZiAoZWxlbS5oYXNBdHRyaWJ1dGUoJ190aWQnKSkge1xuICAgICAgICBqbm9kZS50aWQgPSBlbGVtLmdldEF0dHJpYnV0ZSgnX3RpZCcpO1xuICAgIH1cbiAgICBpZiAoZWxlbS5oYXNBdHRyaWJ1dGUoJ19hdGlkJykpIHtcbiAgICAgICAgam5vZGUuYXRpZCA9IGVsZW0uZ2V0QXR0cmlidXRlKCdfYXRpZCcpO1xuICAgIH1cbiAgICBmb3IgKHZhciBpID0gMDsgaSA8IGVsZW0uYXR0cmlidXRlcy5sZW5ndGg7IGkrKykge1xuICAgICAgICBsZXQgYXR0ciA9IGVsZW0uYXR0cmlidXRlcy5pdGVtKGkpO1xuICAgICAgICBpZiAoYXR0ciA9PSBudWxsKVxuICAgICAgICAgICAgY29udGludWU7XG4gICAgICAgIGlmICghYXR0ci5uYW1lLnN0YXJ0c1dpdGgoJ2V2ZW50LScpKSB7XG4gICAgICAgICAgICBjb250aW51ZTtcbiAgICAgICAgfVxuICAgICAgICBsZXQgZXZlbnROYW1lID0gYXR0ci5uYW1lLnJlcGxhY2UoJ2V2ZW50LScsICcnKTtcbiAgICAgICAgam5vZGUuZXZlbnRzLnB1c2goSlNPTkV2ZW50KGV2ZW50TmFtZSwgYXR0ci52YWx1ZS5zcGxpdCgnfCcpKSk7XG4gICAgfVxuICAgIGlmICghc2hhbGxvdykge1xuICAgICAgICBkb20uYXBwbHlDaGlsZE5vZGUobm9kZSwgZnVuY3Rpb24gKGNoaWxkKSB7XG4gICAgICAgICAgICBpZiAoY2hpbGQgaW5zdGFuY2VvZiBUZXh0IHx8IGNoaWxkIGluc3RhbmNlb2YgQ29tbWVudCkge1xuICAgICAgICAgICAgICAgIGpub2RlLmNoaWxkcmVuLnB1c2goTm9kZVRvSlNPTk5vZGUoY2hpbGQsIGZhbHNlLCBqbm9kZSkpO1xuICAgICAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGNvbnN0IGNoaWxkRWxlbSA9IGNoaWxkO1xuICAgICAgICAgICAgaWYgKCFjaGlsZEVsZW0uaGFzQXR0cmlidXRlKCdpZCcpKSB7XG4gICAgICAgICAgICAgICAgY2hpbGRFbGVtLmlkID0gdXRpbHMuUmFuZG9tSUQoKTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGpub2RlLmNoaWxkcmVuLnB1c2goTm9kZVRvSlNPTk5vZGUoY2hpbGRFbGVtLCBmYWxzZSwgam5vZGUpKTtcbiAgICAgICAgfSk7XG4gICAgfVxuICAgIHJldHVybiBqbm9kZTtcbn1cbmV4cG9ydHMuTm9kZVRvSlNPTk5vZGUgPSBOb2RlVG9KU09OTm9kZTtcbmZ1bmN0aW9uIEpTT05FdmVudChuYW1lLCB0YXJnZXRzKSB7XG4gICAgY29uc3QgZXZlbnQgPSB7fTtcbiAgICBldmVudC5OYW1lID0gbmFtZTtcbiAgICBldmVudC5UYXJnZXRzID0gdGFyZ2V0cztcbiAgICByZXR1cm4gZXZlbnQ7XG59XG5leHBvcnRzLkpTT05FdmVudCA9IEpTT05FdmVudDtcbmZ1bmN0aW9uIGFwcGx5SlNPTk5vZGVGdW5jdGlvbihub2RlLCBmbikge1xuICAgIGZuKG5vZGUpO1xuICAgIG5vZGUuY2hpbGRyZW4uZm9yRWFjaChmdW5jdGlvbiAoY2hpbGQpIHtcbiAgICAgICAgYXBwbHlKU09OTm9kZUZ1bmN0aW9uKGNoaWxkLCBmbik7XG4gICAgfSk7XG59XG5leHBvcnRzLmFwcGx5SlNPTk5vZGVGdW5jdGlvbiA9IGFwcGx5SlNPTk5vZGVGdW5jdGlvbjtcbmZ1bmN0aW9uIGFwcGx5SlNPTk5vZGVLaWRzRnVuY3Rpb24obm9kZSwgZm4pIHtcbiAgICBub2RlLmNoaWxkcmVuLmZvckVhY2goZnVuY3Rpb24gKGNoaWxkKSB7XG4gICAgICAgIGFwcGx5SlNPTk5vZGVGdW5jdGlvbihjaGlsZCwgZm4pO1xuICAgIH0pO1xuICAgIGZuKG5vZGUpO1xufVxuZXhwb3J0cy5hcHBseUpTT05Ob2RlS2lkc0Z1bmN0aW9uID0gYXBwbHlKU09OTm9kZUtpZHNGdW5jdGlvbjtcbmZ1bmN0aW9uIGlzSlNPTk5vZGUobikge1xuICAgIGNvbnN0IGhhc0lEID0gdHlwZW9mIG4uaWQgIT09ICd1bmRlZmluZWQnO1xuICAgIGNvbnN0IGhhc1JlZiA9IHR5cGVvZiBuLnJlZiAhPT0gJ3VuZGVmaW5lZCc7XG4gICAgY29uc3QgaGFzVGlkID0gdHlwZW9mIG4udGlkICE9PSAndW5kZWZpbmVkJztcbiAgICBjb25zdCBoYXNUeXBlTmFtZSA9IHR5cGVvZiBuLnR5cGVOYW1lICE9PSAndW5kZWZpbmVkJztcbiAgICByZXR1cm4gaGFzSUQgJiYgaGFzUmVmICYmIGhhc1R5cGVOYW1lICYmIGhhc1RpZDtcbn1cbmV4cG9ydHMuaXNKU09OTm9kZSA9IGlzSlNPTk5vZGU7XG5mdW5jdGlvbiBmaW5kRWxlbWVudChkZXNjLCBwYXJlbnQpIHtcbiAgICBjb25zdCBzZWxlY3RvciA9IGRlc2MubmFtZSArICcjJyArIGRlc2MuaWQ7XG4gICAgY29uc3QgdGFyZ2V0cyA9IHBhcmVudC5xdWVyeVNlbGVjdG9yQWxsKHNlbGVjdG9yKTtcbiAgICBpZiAodGFyZ2V0cy5sZW5ndGggPT09IDApIHtcbiAgICAgICAgbGV0IGF0dHJTZWxlY3RvciA9IGRlc2MubmFtZSArIGBbX3RpZD0nJHtkZXNjLnRpZH0nXWA7XG4gICAgICAgIGxldCB0YXJnZXQgPSBwYXJlbnQucXVlcnlTZWxlY3RvcihhdHRyU2VsZWN0b3IpO1xuICAgICAgICBpZiAodGFyZ2V0KSB7XG4gICAgICAgICAgICByZXR1cm4gdGFyZ2V0O1xuICAgICAgICB9XG4gICAgICAgIGF0dHJTZWxlY3RvciA9IGRlc2MubmFtZSArIGBbX2F0aWQ9JyR7ZGVzYy5hdGlkfSddYDtcbiAgICAgICAgdGFyZ2V0ID0gcGFyZW50LnF1ZXJ5U2VsZWN0b3IoYXR0clNlbGVjdG9yKTtcbiAgICAgICAgaWYgKHRhcmdldCkge1xuICAgICAgICAgICAgcmV0dXJuIHRhcmdldDtcbiAgICAgICAgfVxuICAgICAgICBhdHRyU2VsZWN0b3IgPSBkZXNjLm5hbWUgKyBgW19yZWY9JyR7ZGVzYy5yZWZ9J11gO1xuICAgICAgICByZXR1cm4gcGFyZW50LnF1ZXJ5U2VsZWN0b3IoYXR0clNlbGVjdG9yKTtcbiAgICB9XG4gICAgaWYgKHRhcmdldHMubGVuZ3RoID09PSAxKSB7XG4gICAgICAgIHJldHVybiB0YXJnZXRzWzBdO1xuICAgIH1cbiAgICBjb25zdCB0b3RhbCA9IHRhcmdldHMubGVuZ3RoO1xuICAgIGZvciAobGV0IGkgPSAwOyBpIDwgdG90YWw7IGkrKykge1xuICAgICAgICBjb25zdCBlbGVtID0gdGFyZ2V0cy5pdGVtKGkpO1xuICAgICAgICBpZiAoZWxlbS5nZXRBdHRyaWJ1dGUoJ190aWQnKSA9PT0gZGVzYy50aWQpIHtcbiAgICAgICAgICAgIHJldHVybiBlbGVtO1xuICAgICAgICB9XG4gICAgICAgIGlmIChlbGVtLmdldEF0dHJpYnV0ZSgnX2F0aWQnKSA9PT0gZGVzYy5hdGlkKSB7XG4gICAgICAgICAgICByZXR1cm4gZWxlbTtcbiAgICAgICAgfVxuICAgICAgICBpZiAoZWxlbS5nZXRBdHRyaWJ1dGUoJ19yZWYnKSA9PT0gZGVzYy5yZWYpIHtcbiAgICAgICAgICAgIHJldHVybiBlbGVtO1xuICAgICAgICB9XG4gICAgfVxuICAgIHJldHVybiBudWxsO1xufVxuZXhwb3J0cy5maW5kRWxlbWVudCA9IGZpbmRFbGVtZW50O1xuZnVuY3Rpb24gZmluZEVsZW1lbnRieVJlZihyZWYsIHBhcmVudCkge1xuICAgIGNvbnN0IGlkcyA9IHJlZi5zcGxpdCgnLycpLm1hcChmdW5jdGlvbiAoZWxlbSkge1xuICAgICAgICBpZiAoZWxlbS50cmltKCkgPT09ICcnKSB7XG4gICAgICAgICAgICByZXR1cm4gJyc7XG4gICAgICAgIH1cbiAgICAgICAgcmV0dXJuICcjJyArIGVsZW07XG4gICAgfSk7XG4gICAgaWYgKGlkcy5sZW5ndGggPT09IDApIHtcbiAgICAgICAgcmV0dXJuIG51bGw7XG4gICAgfVxuICAgIGlmIChpZHNbMF0gPT09ICcnIHx8IGlkc1swXS50cmltKCkgPT09ICcnKSB7XG4gICAgICAgIGlkcy5zaGlmdCgpO1xuICAgIH1cbiAgICBjb25zdCBmaXJzdCA9IGlkc1swXTtcbiAgICBpZiAocGFyZW50LmlkID09IGZpcnN0LnN1YnN0cigxKSkge1xuICAgICAgICBpZHMuc2hpZnQoKTtcbiAgICB9XG4gICAgbGV0IGN1ciA9IHBhcmVudC5xdWVyeVNlbGVjdG9yKGlkcy5zaGlmdCgpKTtcbiAgICB3aGlsZSAoY3VyKSB7XG4gICAgICAgIGlmIChpZHMubGVuZ3RoID09PSAwKSB7XG4gICAgICAgICAgICByZXR1cm4gY3VyO1xuICAgICAgICB9XG4gICAgICAgIGN1ciA9IGN1ci5xdWVyeVNlbGVjdG9yKGlkcy5zaGlmdCgpKTtcbiAgICB9XG4gICAgcmV0dXJuIGN1cjtcbn1cbmV4cG9ydHMuZmluZEVsZW1lbnRieVJlZiA9IGZpbmRFbGVtZW50YnlSZWY7XG5mdW5jdGlvbiBmaW5kRWxlbWVudFBhcmVudGJ5UmVmKHJlZiwgcGFyZW50KSB7XG4gICAgY29uc3QgaWRzID0gcmVmLnNwbGl0KCcvJykubWFwKGZ1bmN0aW9uIChlbGVtKSB7XG4gICAgICAgIGlmIChlbGVtLnRyaW0oKSA9PT0gJycpIHtcbiAgICAgICAgICAgIHJldHVybiAnJztcbiAgICAgICAgfVxuICAgICAgICByZXR1cm4gJyMnICsgZWxlbTtcbiAgICB9KTtcbiAgICBpZiAoaWRzLmxlbmd0aCA9PT0gMCkge1xuICAgICAgICByZXR1cm4gbnVsbDtcbiAgICB9XG4gICAgaWYgKGlkc1swXSA9PT0gJycgfHwgaWRzWzBdLnRyaW0oKSA9PT0gJycpIHtcbiAgICAgICAgaWRzLnNoaWZ0KCk7XG4gICAgfVxuICAgIGlkcy5wb3AoKTtcbiAgICBjb25zdCBmaXJzdCA9IGlkc1swXTtcbiAgICBpZiAocGFyZW50LmlkID09IGZpcnN0LnN1YnN0cigxKSkge1xuICAgICAgICBpZHMuc2hpZnQoKTtcbiAgICB9XG4gICAgbGV0IGN1ciA9IHBhcmVudC5xdWVyeVNlbGVjdG9yKGlkcy5zaGlmdCgpKTtcbiAgICB3aGlsZSAoY3VyKSB7XG4gICAgICAgIGlmIChpZHMubGVuZ3RoID09PSAwKSB7XG4gICAgICAgICAgICByZXR1cm4gY3VyO1xuICAgICAgICB9XG4gICAgICAgIGN1ciA9IGN1ci5xdWVyeVNlbGVjdG9yKGlkcy5zaGlmdCgpKTtcbiAgICB9XG4gICAgcmV0dXJuIGN1cjtcbn1cbmV4cG9ydHMuZmluZEVsZW1lbnRQYXJlbnRieVJlZiA9IGZpbmRFbGVtZW50UGFyZW50YnlSZWY7XG5leHBvcnRzLkRlZmF1bHRKU09OTWFrZXIgPSB7XG4gICAgTWFrZToganNvbk1ha2VyLFxufTtcbmZ1bmN0aW9uIGpzb25NYWtlcihkb2MsIGRlc2NOb2RlLCBzaGFsbG93LCBza2lwUmVtb3ZlZCkge1xuICAgIGlmIChkZXNjTm9kZS50eXBlID09PSBkb21fMS5DT01NRU5UX05PREUpIHtcbiAgICAgICAgY29uc3Qgbm9kZSA9IGRvYy5jcmVhdGVDb21tZW50KGRlc2NOb2RlLmNvbnRlbnQpO1xuICAgICAgICBleHRzLk9iamVjdHMuUGF0Y2hXaXRoKG5vZGUsICdfaWQnLCBkZXNjTm9kZS5pZCk7XG4gICAgICAgIGV4dHMuT2JqZWN0cy5QYXRjaFdpdGgobm9kZSwgJ19yZWYnLCBkZXNjTm9kZS5yZWYpO1xuICAgICAgICBleHRzLk9iamVjdHMuUGF0Y2hXaXRoKG5vZGUsICdfdGlkJywgZGVzY05vZGUudGlkKTtcbiAgICAgICAgZXh0cy5PYmplY3RzLlBhdGNoV2l0aChub2RlLCAnX2F0aWQnLCBkZXNjTm9kZS5hdGlkKTtcbiAgICAgICAgcmV0dXJuIG5vZGU7XG4gICAgfVxuICAgIGlmIChkZXNjTm9kZS50eXBlID09PSBkb21fMS5URVhUX05PREUpIHtcbiAgICAgICAgY29uc3Qgbm9kZSA9IGRvYy5jcmVhdGVUZXh0Tm9kZShkZXNjTm9kZS5jb250ZW50KTtcbiAgICAgICAgZXh0cy5PYmplY3RzLlBhdGNoV2l0aChub2RlLCAnX2lkJywgZGVzY05vZGUuaWQpO1xuICAgICAgICBleHRzLk9iamVjdHMuUGF0Y2hXaXRoKG5vZGUsICdfcmVmJywgZGVzY05vZGUucmVmKTtcbiAgICAgICAgZXh0cy5PYmplY3RzLlBhdGNoV2l0aChub2RlLCAnX3RpZCcsIGRlc2NOb2RlLnRpZCk7XG4gICAgICAgIGV4dHMuT2JqZWN0cy5QYXRjaFdpdGgobm9kZSwgJ19hdGlkJywgZGVzY05vZGUuYXRpZCk7XG4gICAgICAgIHJldHVybiBub2RlO1xuICAgIH1cbiAgICBpZiAoZGVzY05vZGUuaWQgPT09ICcnKSB7XG4gICAgICAgIGRlc2NOb2RlLmlkID0gdXRpbHMuUmFuZG9tSUQoKTtcbiAgICB9XG4gICAgbGV0IG5vZGU7XG4gICAgaWYgKGRlc2NOb2RlLm5hbWVzcGFjZS5sZW5ndGggIT09IDApIHtcbiAgICAgICAgbm9kZSA9IGRvYy5jcmVhdGVFbGVtZW50KGRlc2NOb2RlLm5hbWUpO1xuICAgIH1cbiAgICBlbHNlIHtcbiAgICAgICAgbm9kZSA9IGRvYy5jcmVhdGVFbGVtZW50TlMoZGVzY05vZGUubmFtZXNwYWNlLCBkZXNjTm9kZS5uYW1lKTtcbiAgICB9XG4gICAgZXh0cy5PYmplY3RzLlBhdGNoV2l0aChub2RlLCAnX2lkJywgZGVzY05vZGUuaWQpO1xuICAgIGV4dHMuT2JqZWN0cy5QYXRjaFdpdGgobm9kZSwgJ19yZWYnLCBkZXNjTm9kZS5yZWYpO1xuICAgIGV4dHMuT2JqZWN0cy5QYXRjaFdpdGgobm9kZSwgJ190aWQnLCBkZXNjTm9kZS50aWQpO1xuICAgIGV4dHMuT2JqZWN0cy5QYXRjaFdpdGgobm9kZSwgJ19hdGlkJywgZGVzY05vZGUuYXRpZCk7XG4gICAgbm9kZS5zZXRBdHRyaWJ1dGUoJ2lkJywgZGVzY05vZGUuaWQpO1xuICAgIG5vZGUuc2V0QXR0cmlidXRlKCdfdGlkJywgZGVzY05vZGUudGlkKTtcbiAgICBub2RlLnNldEF0dHJpYnV0ZSgnX3JlZicsIGRlc2NOb2RlLnJlZik7XG4gICAgbm9kZS5zZXRBdHRyaWJ1dGUoJ19hdGlkJywgZGVzY05vZGUuYXRpZCk7XG4gICAgZGVzY05vZGUuZXZlbnRzLmZvckVhY2goZnVuY3Rpb24gZXZlbnRzKGV2ZW50KSB7XG4gICAgICAgIG5vZGUuc2V0QXR0cmlidXRlKCdldmVudC0nICsgZXZlbnQuTmFtZSwgZXZlbnQuVGFyZ2V0cy5qb2luKCd8JykpO1xuICAgIH0pO1xuICAgIGRlc2NOb2RlLmF0dHJzLmZvckVhY2goZnVuY3Rpb24gYXR0cnMoYXR0cikge1xuICAgICAgICBub2RlLnNldEF0dHJpYnV0ZShhdHRyLktleSwgYXR0ci5WYWx1ZSk7XG4gICAgfSk7XG4gICAgaWYgKGRlc2NOb2RlLnJlbW92ZWQpIHtcbiAgICAgICAgbm9kZS5zZXRBdHRyaWJ1dGUoJ19yZW1vdmVkJywgJ3RydWUnKTtcbiAgICAgICAgcmV0dXJuIG5vZGU7XG4gICAgfVxuICAgIGlmICghc2hhbGxvdykge1xuICAgICAgICBkZXNjTm9kZS5jaGlsZHJlbi5mb3JFYWNoKGZ1bmN0aW9uIChraWRKU09OKSB7XG4gICAgICAgICAgICBpZiAoc2tpcFJlbW92ZWQgJiYga2lkSlNPTi5yZW1vdmVkKSB7XG4gICAgICAgICAgICAgICAgcmV0dXJuO1xuICAgICAgICAgICAgfVxuICAgICAgICAgICAgbm9kZS5hcHBlbmRDaGlsZChqc29uTWFrZXIoZG9jLCBraWRKU09OLCBzaGFsbG93LCBza2lwUmVtb3ZlZCkpO1xuICAgICAgICB9KTtcbiAgICB9XG4gICAgcmV0dXJuIG5vZGU7XG59XG5leHBvcnRzLmpzb25NYWtlciA9IGpzb25NYWtlcjtcbmZ1bmN0aW9uIFBhdGNoSlNPTk5vZGVUcmVlKGZyYWdtZW50LCBtb3VudCwgZGljdGF0b3IsIG1ha2VyKSB7XG4gICAgbGV0IHRhcmdldE5vZGUgPSBmaW5kRWxlbWVudChmcmFnbWVudCwgbW91bnQpO1xuICAgIGlmIChleHRzLk9iamVjdHMuaXNOdWxsT3JVbmRlZmluZWQodGFyZ2V0Tm9kZSkpIHtcbiAgICAgICAgY29uc3QgdE5vZGUgPSBtYWtlci5NYWtlKGRvY3VtZW50LCBmcmFnbWVudCwgZmFsc2UsIHRydWUpO1xuICAgICAgICBtb3VudC5hcHBlbmRDaGlsZCh0Tm9kZSk7XG4gICAgICAgIHJldHVybjtcbiAgICB9XG4gICAgUGF0Y2hKU09OTm9kZShmcmFnbWVudCwgdGFyZ2V0Tm9kZSwgZGljdGF0b3IsIG1ha2VyKTtcbn1cbmV4cG9ydHMuUGF0Y2hKU09OTm9kZVRyZWUgPSBQYXRjaEpTT05Ob2RlVHJlZTtcbmZ1bmN0aW9uIFBhdGNoSlNPTk5vZGUoZnJhZ21lbnQsIHRhcmdldE5vZGUsIGRpY3RhdG9yLCBtYWtlcikge1xuICAgIGlmICghZGljdGF0b3IuU2FtZSh0YXJnZXROb2RlLCBmcmFnbWVudCkpIHtcbiAgICAgICAgY29uc3QgdE5vZGUgPSBtYWtlci5NYWtlKGRvY3VtZW50LCBmcmFnbWVudCwgZmFsc2UsIHRydWUpO1xuICAgICAgICBkb20ucmVwbGFjZU5vZGUodGFyZ2V0Tm9kZS5wYXJlbnROb2RlLCB0YXJnZXROb2RlLCB0Tm9kZSk7XG4gICAgICAgIHJldHVybjtcbiAgICB9XG4gICAgaWYgKCFkaWN0YXRvci5DaGFuZ2VkKHRhcmdldE5vZGUsIGZyYWdtZW50KSkge1xuICAgICAgICByZXR1cm47XG4gICAgfVxuICAgIFBhdGNoSlNPTkF0dHJpYnV0ZXMoZnJhZ21lbnQsIHRhcmdldE5vZGUpO1xuICAgIGNvbnN0IGtpZHMgPSBkb20ubm9kZUxpc3RUb0FycmF5KHRhcmdldE5vZGUuY2hpbGROb2Rlcyk7XG4gICAgY29uc3QgdG90YWxLaWRzID0ga2lkcy5sZW5ndGg7XG4gICAgY29uc3QgZnJhZ21lbnRLaWRzID0gZnJhZ21lbnQuY2hpbGRyZW4ubGVuZ3RoO1xuICAgIGxldCBpID0gMDtcbiAgICBmb3IgKDsgaSA8IHRvdGFsS2lkczsgaSsrKSB7XG4gICAgICAgIGNvbnN0IGNoaWxkTm9kZSA9IGtpZHNbaV07XG4gICAgICAgIGlmIChpID49IGZyYWdtZW50S2lkcykge1xuICAgICAgICAgICAgY29uc3QgY2hub2RlID0gY2hpbGROb2RlO1xuICAgICAgICAgICAgaWYgKGNobm9kZSkge1xuICAgICAgICAgICAgICAgIGNobm9kZS5yZW1vdmUoKTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGNvbnRpbnVlO1xuICAgICAgICB9XG4gICAgICAgIGNvbnN0IGNoaWxkRnJhZ21lbnQgPSBmcmFnbWVudC5jaGlsZHJlbltpXTtcbiAgICAgICAgUGF0Y2hKU09OTm9kZShjaGlsZEZyYWdtZW50LCBjaGlsZE5vZGUsIGRpY3RhdG9yLCBtYWtlcik7XG4gICAgfVxuICAgIGZvciAoOyBpIDwgZnJhZ21lbnRLaWRzOyBpKyspIHtcbiAgICAgICAgY29uc3QgdE5vZGUgPSBtYWtlci5NYWtlKGRvY3VtZW50LCBmcmFnbWVudCwgZmFsc2UsIHRydWUpO1xuICAgICAgICB0YXJnZXROb2RlLmFwcGVuZENoaWxkKHROb2RlKTtcbiAgICB9XG4gICAgcmV0dXJuO1xufVxuZXhwb3J0cy5QYXRjaEpTT05Ob2RlID0gUGF0Y2hKU09OTm9kZTtcbmZ1bmN0aW9uIFN0cmVhbUpTT05Ob2RlcyhmcmFnbWVudCwgbW91bnQsIGRpY3RhdG9yLCBtYWtlcikge1xuICAgIGNvbnN0IGNoYW5nZXMgPSBmcmFnbWVudC5maWx0ZXIoZnVuY3Rpb24gKGVsZW0pIHtcbiAgICAgICAgcmV0dXJuICFlbGVtLnJlbW92ZWQ7XG4gICAgfSk7XG4gICAgZnJhZ21lbnRcbiAgICAgICAgLmZpbHRlcihmdW5jdGlvbiAoZWxlbSkge1xuICAgICAgICBpZiAoIWVsZW0ucmVtb3ZlZCkge1xuICAgICAgICAgICAgcmV0dXJuIGZhbHNlO1xuICAgICAgICB9XG4gICAgICAgIGxldCBmaWx0ZXJlZCA9IHRydWU7XG4gICAgICAgIGNoYW5nZXMuZm9yRWFjaChmdW5jdGlvbiAoZWwpIHtcbiAgICAgICAgICAgIGlmIChlbGVtLnRpZCA9PT0gZWwudGlkIHx8IGVsZW0udGlkID09IGVsLmF0aWQgfHwgZWxlbS5yZWYgPT09IGVsLnJlZikge1xuICAgICAgICAgICAgICAgIGZpbHRlcmVkID0gZmFsc2U7XG4gICAgICAgICAgICB9XG4gICAgICAgIH0pO1xuICAgICAgICByZXR1cm4gZmlsdGVyZWQ7XG4gICAgfSlcbiAgICAgICAgLmZvckVhY2goZnVuY3Rpb24gKHJlbW92YWwpIHtcbiAgICAgICAgY29uc3QgdGFyZ2V0ID0gZmluZEVsZW1lbnQocmVtb3ZhbCwgbW91bnQpO1xuICAgICAgICBpZiAodGFyZ2V0KSB7XG4gICAgICAgICAgICB0YXJnZXQucmVtb3ZlKCk7XG4gICAgICAgIH1cbiAgICB9KTtcbiAgICBjaGFuZ2VzLmZvckVhY2goZnVuY3Rpb24gKGNoYW5nZSkge1xuICAgICAgICBjb25zdCB0YXJnZXROb2RlID0gZmluZEVsZW1lbnQoY2hhbmdlLCBtb3VudCk7XG4gICAgICAgIGlmIChleHRzLk9iamVjdHMuaXNOdWxsT3JVbmRlZmluZWQodGFyZ2V0Tm9kZSkpIHtcbiAgICAgICAgICAgIGNvbnN0IHRhcmdldE5vZGVQYXJlbnQgPSBmaW5kRWxlbWVudFBhcmVudGJ5UmVmKGNoYW5nZS5yZWYsIG1vdW50KTtcbiAgICAgICAgICAgIGlmIChleHRzLk9iamVjdHMuaXNOdWxsT3JVbmRlZmluZWQodGFyZ2V0Tm9kZVBhcmVudCkpIHtcbiAgICAgICAgICAgICAgICBjb25zb2xlLmxvZygnVW5hYmxlIHRvIGFwcGx5IG5ldyBjaGFuZ2Ugc3RyZWFtOiAnLCBjaGFuZ2UpO1xuICAgICAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGNvbnN0IHROb2RlID0gbWFrZXIuTWFrZShkb2N1bWVudCwgY2hhbmdlLCBmYWxzZSwgdHJ1ZSk7XG4gICAgICAgICAgICB0YXJnZXROb2RlUGFyZW50LmFwcGVuZENoaWxkKHROb2RlKTtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICBBcHBseVN0cmVhbU5vZGUoY2hhbmdlLCB0YXJnZXROb2RlLCBkaWN0YXRvciwgbWFrZXIpO1xuICAgIH0pO1xuICAgIHJldHVybjtcbn1cbmV4cG9ydHMuU3RyZWFtSlNPTk5vZGVzID0gU3RyZWFtSlNPTk5vZGVzO1xuZnVuY3Rpb24gQXBwbHlTdHJlYW1Ob2RlKGZyYWdtZW50LCB0YXJnZXROb2RlLCBkaWN0YXRvciwgbWFrZXIpIHtcbiAgICBpZiAoIWRpY3RhdG9yLlNhbWUodGFyZ2V0Tm9kZSwgZnJhZ21lbnQpKSB7XG4gICAgICAgIGNvbnN0IHROb2RlID0gbWFrZXIuTWFrZShkb2N1bWVudCwgZnJhZ21lbnQsIGZhbHNlLCB0cnVlKTtcbiAgICAgICAgZG9tLnJlcGxhY2VOb2RlKHRhcmdldE5vZGUucGFyZW50Tm9kZSwgdGFyZ2V0Tm9kZSwgdE5vZGUpO1xuICAgICAgICByZXR1cm47XG4gICAgfVxuICAgIGlmIChkaWN0YXRvci5DaGFuZ2VkKHRhcmdldE5vZGUsIGZyYWdtZW50KSkge1xuICAgICAgICBQYXRjaEpTT05BdHRyaWJ1dGVzKGZyYWdtZW50LCB0YXJnZXROb2RlKTtcbiAgICB9XG4gICAgY29uc3QgdG90YWxLaWRzID0gdGFyZ2V0Tm9kZS5jaGlsZE5vZGVzLmxlbmd0aDtcbiAgICBjb25zdCBmcmFnbWVudEtpZHMgPSBmcmFnbWVudC5jaGlsZHJlbi5sZW5ndGg7XG4gICAgbGV0IGkgPSAwO1xuICAgIGZvciAoOyBpIDwgdG90YWxLaWRzOyBpKyspIHtcbiAgICAgICAgY29uc3QgY2hpbGROb2RlID0gdGFyZ2V0Tm9kZS5jaGlsZE5vZGVzW2ldO1xuICAgICAgICBpZiAoaSA+PSBmcmFnbWVudEtpZHMpIHtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICBjb25zdCBjaGlsZEZyYWdtZW50ID0gZnJhZ21lbnQuY2hpbGRyZW5baV07XG4gICAgICAgIFBhdGNoSlNPTk5vZGUoY2hpbGRGcmFnbWVudCwgY2hpbGROb2RlLCBkaWN0YXRvciwgbWFrZXIpO1xuICAgIH1cbiAgICBmb3IgKDsgaSA8IGZyYWdtZW50S2lkczsgaSsrKSB7XG4gICAgICAgIGNvbnN0IHROb2RlID0gbWFrZXIuTWFrZShkb2N1bWVudCwgZnJhZ21lbnQsIGZhbHNlLCB0cnVlKTtcbiAgICAgICAgdGFyZ2V0Tm9kZS5hcHBlbmRDaGlsZCh0Tm9kZSk7XG4gICAgfVxuICAgIHJldHVybjtcbn1cbmV4cG9ydHMuQXBwbHlTdHJlYW1Ob2RlID0gQXBwbHlTdHJlYW1Ob2RlO1xuZnVuY3Rpb24gUGF0Y2hUZXh0Q29tbWVudFdpdGhKU09OKGZyYWdtZW50LCB0YXJnZXQpIHtcbiAgICBpZiAoZnJhZ21lbnQudHlwZSAhPT0gZG9tXzEuQ09NTUVOVF9OT0RFICYmIGZyYWdtZW50LnR5cGUgIT09IGRvbV8xLlRFWFRfTk9ERSkge1xuICAgICAgICByZXR1cm47XG4gICAgfVxuICAgIGlmIChmcmFnbWVudC50eXBlICE9PSBkb21fMS5DT01NRU5UX05PREUgJiYgZnJhZ21lbnQudHlwZSAhPT0gZG9tXzEuVEVYVF9OT0RFKSB7XG4gICAgICAgIHJldHVybjtcbiAgICB9XG4gICAgaWYgKHRhcmdldC50ZXh0Q29udGVudCA9PT0gZnJhZ21lbnQuY29udGVudCkge1xuICAgICAgICByZXR1cm47XG4gICAgfVxuICAgIHRhcmdldC50ZXh0Q29udGVudCA9IGZyYWdtZW50LmNvbnRlbnQ7XG4gICAgZXh0cy5PYmplY3RzLlBhdGNoV2l0aCh0YXJnZXQsICdfcmVmJywgZnJhZ21lbnQucmVmKTtcbiAgICBleHRzLk9iamVjdHMuUGF0Y2hXaXRoKHRhcmdldCwgJ190aWQnLCBmcmFnbWVudC50aWQpO1xuICAgIGV4dHMuT2JqZWN0cy5QYXRjaFdpdGgodGFyZ2V0LCAnX2F0aWQnLCBmcmFnbWVudC5hdGlkKTtcbn1cbmV4cG9ydHMuUGF0Y2hUZXh0Q29tbWVudFdpdGhKU09OID0gUGF0Y2hUZXh0Q29tbWVudFdpdGhKU09OO1xuZnVuY3Rpb24gUGF0Y2hKU09OQXR0cmlidXRlcyhub2RlLCB0YXJnZXQpIHtcbiAgICBjb25zdCBvbGROb2RlQXR0cnMgPSBkb20ucmVjb3JkQXR0cmlidXRlcyh0YXJnZXQpO1xuICAgIG5vZGUuYXR0cnMuZm9yRWFjaChmdW5jdGlvbiAoYXR0cikge1xuICAgICAgICBjb25zdCBvbGRWYWx1ZSA9IG9sZE5vZGVBdHRyc1thdHRyLktleV07XG4gICAgICAgIGRlbGV0ZSBvbGROb2RlQXR0cnNbYXR0ci5LZXldO1xuICAgICAgICBpZiAoYXR0ci5WYWx1ZSA9PT0gb2xkVmFsdWUpIHtcbiAgICAgICAgICAgIHJldHVybiBudWxsO1xuICAgICAgICB9XG4gICAgICAgIHRhcmdldC5zZXRBdHRyaWJ1dGUoYXR0ci5LZXksIGF0dHIuVmFsdWUpO1xuICAgIH0pO1xuICAgIGZvciAobGV0IGluZGV4IGluIG9sZE5vZGVBdHRycykge1xuICAgICAgICB0YXJnZXQucmVtb3ZlQXR0cmlidXRlKGluZGV4KTtcbiAgICB9XG4gICAgdGFyZ2V0LnNldEF0dHJpYnV0ZSgnX3RpZCcsIG5vZGUudGlkKTtcbiAgICB0YXJnZXQuc2V0QXR0cmlidXRlKCdfcmVmJywgbm9kZS5yZWYpO1xuICAgIHRhcmdldC5zZXRBdHRyaWJ1dGUoJ19hdGlkJywgbm9kZS5hdGlkKTtcbiAgICBub2RlLmV2ZW50cy5mb3JFYWNoKGZ1bmN0aW9uIGV2ZW50cyhldmVudCkge1xuICAgICAgICB0YXJnZXQuc2V0QXR0cmlidXRlKCdldmVudC0nICsgZXZlbnQuTmFtZSwgZXZlbnQuVGFyZ2V0cy5qb2luKCd8JykpO1xuICAgIH0pO1xuICAgIGV4dHMuT2JqZWN0cy5QYXRjaFdpdGgodGFyZ2V0LCAnX2lkJywgbm9kZS5pZCk7XG4gICAgZXh0cy5PYmplY3RzLlBhdGNoV2l0aCh0YXJnZXQsICdfcmVmJywgbm9kZS5yZWYpO1xuICAgIGV4dHMuT2JqZWN0cy5QYXRjaFdpdGgodGFyZ2V0LCAnX3RpZCcsIG5vZGUudGlkKTtcbiAgICBleHRzLk9iamVjdHMuUGF0Y2hXaXRoKHRhcmdldCwgJ19hdGlkJywgbm9kZS5hdGlkKTtcbn1cbmV4cG9ydHMuUGF0Y2hKU09OQXR0cmlidXRlcyA9IFBhdGNoSlNPTkF0dHJpYnV0ZXM7XG5mdW5jdGlvbiBQYXRjaERPTVRyZWUobmV3RnJhZ21lbnQsIG9sZE5vZGVPck1vdW50LCBkaWN0YXRvciwgaXNDaGlsZFJlY3Vyc2lvbikge1xuICAgIGlmIChpc0NoaWxkUmVjdXJzaW9uKSB7XG4gICAgICAgIGNvbnN0IHJvb3ROb2RlID0gb2xkTm9kZU9yTW91bnQucGFyZW50Tm9kZTtcbiAgICAgICAgaWYgKCFkaWN0YXRvci5TYW1lKG9sZE5vZGVPck1vdW50LCBuZXdGcmFnbWVudCkpIHtcbiAgICAgICAgICAgIGRvbS5yZXBsYWNlTm9kZShyb290Tm9kZSwgb2xkTm9kZU9yTW91bnQsIG5ld0ZyYWdtZW50KTtcbiAgICAgICAgICAgIHJldHVybiBudWxsO1xuICAgICAgICB9XG4gICAgICAgIGlmICghb2xkTm9kZU9yTW91bnQuaGFzQ2hpbGROb2RlcygpKSB7XG4gICAgICAgICAgICBkb20ucmVwbGFjZU5vZGUocm9vdE5vZGUsIG9sZE5vZGVPck1vdW50LCBuZXdGcmFnbWVudCk7XG4gICAgICAgICAgICByZXR1cm4gbnVsbDtcbiAgICAgICAgfVxuICAgIH1cbiAgICBjb25zdCBuZXdDaGlsZHJlbiA9IGRvbS5ub2RlTGlzdFRvQXJyYXkobmV3RnJhZ21lbnQuY2hpbGROb2Rlcyk7XG4gICAgY29uc3Qgb2xkQ2hpbGRyZW4gPSBkb20ubm9kZUxpc3RUb0FycmF5KG9sZE5vZGVPck1vdW50LmNoaWxkTm9kZXMpO1xuICAgIGNvbnN0IG9sZENoaWxkcmVuTGVuZ3RoID0gb2xkQ2hpbGRyZW4ubGVuZ3RoO1xuICAgIGNvbnN0IG5ld0NoaWxkcmVuTGVuZ3RoID0gbmV3Q2hpbGRyZW4ubGVuZ3RoO1xuICAgIGNvbnN0IHJlbW92ZU9sZExlZnQgPSBuZXdDaGlsZHJlbkxlbmd0aCA8IG9sZENoaWxkcmVuTGVuZ3RoO1xuICAgIGxldCBsYXN0SW5kZXggPSAwO1xuICAgIGxldCBsYXN0Tm9kZTtcbiAgICBsZXQgbmV3Q2hpbGROb2RlO1xuICAgIGxldCBsYXN0Tm9kZU5leHRTaWJsaW5nID0gbnVsbDtcbiAgICBmb3IgKDsgbGFzdEluZGV4IDwgb2xkQ2hpbGRyZW5MZW5ndGg7IGxhc3RJbmRleCsrKSB7XG4gICAgICAgIGlmIChsYXN0SW5kZXggPj0gbmV3Q2hpbGRyZW5MZW5ndGgpIHtcbiAgICAgICAgICAgIGJyZWFrO1xuICAgICAgICB9XG4gICAgICAgIGxhc3ROb2RlID0gb2xkQ2hpbGRyZW5bbGFzdEluZGV4XTtcbiAgICAgICAgbmV3Q2hpbGROb2RlID0gbmV3Q2hpbGRyZW5bbGFzdEluZGV4XTtcbiAgICAgICAgbGFzdE5vZGVOZXh0U2libGluZyA9IGxhc3ROb2RlLm5leHRTaWJsaW5nO1xuICAgICAgICBpZiAoKGxhc3ROb2RlLm5vZGVUeXBlID09PSBkb20uVEVYVF9OT0RFIHx8IGxhc3ROb2RlLm5vZGVUeXBlID09PSBkb20uQ09NTUVOVF9OT0RFKSAmJlxuICAgICAgICAgICAgbmV3Q2hpbGROb2RlLm5vZGVUeXBlID09PSBsYXN0Tm9kZS5ub2RlVHlwZSkge1xuICAgICAgICAgICAgaWYgKGxhc3ROb2RlLnRleHRDb250ZW50ICE9PSBuZXdDaGlsZE5vZGUudGV4dENvbnRlbnQpIHtcbiAgICAgICAgICAgICAgICBsYXN0Tm9kZS50ZXh0Q29udGVudCA9IG5ld0NoaWxkTm9kZS50ZXh0Q29udGVudDtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGNvbnRpbnVlO1xuICAgICAgICB9XG4gICAgICAgIGlmICghZGljdGF0b3IuU2FtZShsYXN0Tm9kZSwgbmV3Q2hpbGROb2RlKSkge1xuICAgICAgICAgICAgZG9tLnJlcGxhY2VOb2RlKG9sZE5vZGVPck1vdW50LCBsYXN0Tm9kZSwgbmV3Q2hpbGROb2RlKTtcbiAgICAgICAgICAgIGNvbnRpbnVlO1xuICAgICAgICB9XG4gICAgICAgIGlmICghZGljdGF0b3IuQ2hhbmdlZChsYXN0Tm9kZSwgbmV3Q2hpbGROb2RlKSkge1xuICAgICAgICAgICAgY29udGludWU7XG4gICAgICAgIH1cbiAgICAgICAgaWYgKGxhc3ROb2RlLm5vZGVUeXBlICE9PSBuZXdDaGlsZE5vZGUubm9kZVR5cGUpIHtcbiAgICAgICAgICAgIGRvbS5yZXBsYWNlTm9kZShvbGROb2RlT3JNb3VudCwgbGFzdE5vZGUsIG5ld0NoaWxkTm9kZSk7XG4gICAgICAgICAgICBjb250aW51ZTtcbiAgICAgICAgfVxuICAgICAgICBpZiAoIWxhc3ROb2RlLmhhc0NoaWxkTm9kZXMoKSAmJiBuZXdDaGlsZE5vZGUuaGFzQ2hpbGROb2RlcygpKSB7XG4gICAgICAgICAgICBkb20ucmVwbGFjZU5vZGUob2xkTm9kZU9yTW91bnQsIGxhc3ROb2RlLCBuZXdDaGlsZE5vZGUpO1xuICAgICAgICAgICAgY29udGludWU7XG4gICAgICAgIH1cbiAgICAgICAgaWYgKGxhc3ROb2RlLmhhc0NoaWxkTm9kZXMoKSAmJiAhbmV3Q2hpbGROb2RlLmhhc0NoaWxkTm9kZXMoKSkge1xuICAgICAgICAgICAgZG9tLnJlcGxhY2VOb2RlKG9sZE5vZGVPck1vdW50LCBsYXN0Tm9kZSwgbmV3Q2hpbGROb2RlKTtcbiAgICAgICAgICAgIGNvbnRpbnVlO1xuICAgICAgICB9XG4gICAgICAgIGNvbnN0IGxhc3RFbGVtZW50ID0gbGFzdE5vZGU7XG4gICAgICAgIGNvbnN0IG5ld0VsZW1lbnQgPSBuZXdDaGlsZE5vZGU7XG4gICAgICAgIFBhdGNoRE9NQXR0cmlidXRlcyhuZXdFbGVtZW50LCBsYXN0RWxlbWVudCk7XG4gICAgICAgIGxhc3RFbGVtZW50LnNldEF0dHJpYnV0ZSgnX3BhdGNoZWQnLCAndHJ1ZScpO1xuICAgICAgICBQYXRjaERPTVRyZWUobmV3RWxlbWVudCwgbGFzdEVsZW1lbnQsIGRpY3RhdG9yLCB0cnVlKTtcbiAgICAgICAgbGFzdEVsZW1lbnQucmVtb3ZlQXR0cmlidXRlKCdfcGF0Y2hlZCcpO1xuICAgIH1cbiAgICBpZiAocmVtb3ZlT2xkTGVmdCAmJiBsYXN0Tm9kZU5leHRTaWJsaW5nICE9PSBudWxsKSB7XG4gICAgICAgIGRvbS5yZW1vdmVGcm9tTm9kZShsYXN0Tm9kZU5leHRTaWJsaW5nLCBudWxsKTtcbiAgICAgICAgcmV0dXJuIG51bGw7XG4gICAgfVxuICAgIGZvciAoOyBsYXN0SW5kZXggPCBuZXdDaGlsZHJlbkxlbmd0aDsgbGFzdEluZGV4KyspIHtcbiAgICAgICAgbGV0IG5ld05vZGUgPSBuZXdDaGlsZHJlbltsYXN0SW5kZXhdO1xuICAgICAgICBpZiAoIWV4dHMuT2JqZWN0cy5pc051bGxPclVuZGVmaW5lZChuZXdOb2RlKSkge1xuICAgICAgICAgICAgb2xkTm9kZU9yTW91bnQuYXBwZW5kQ2hpbGQobmV3Tm9kZSk7XG4gICAgICAgIH1cbiAgICB9XG59XG5leHBvcnRzLlBhdGNoRE9NVHJlZSA9IFBhdGNoRE9NVHJlZTtcbmZ1bmN0aW9uIFBhdGNoRE9NQXR0cmlidXRlcyhuZXdFbGVtZW50LCBvbGRFbGVtZW50KSB7XG4gICAgY29uc3Qgb2xkTm9kZUF0dHJzID0gZG9tLnJlY29yZEF0dHJpYnV0ZXMob2xkRWxlbWVudCk7XG4gICAgZm9yIChsZXQgaW5kZXggaW4gbmV3RWxlbWVudC5hdHRyaWJ1dGVzKSB7XG4gICAgICAgIGNvbnN0IGF0dHIgPSBuZXdFbGVtZW50LmF0dHJpYnV0ZXNbaW5kZXhdO1xuICAgICAgICBjb25zdCBvbGRWYWx1ZSA9IG9sZE5vZGVBdHRyc1thdHRyLm5hbWVdO1xuICAgICAgICBkZWxldGUgb2xkTm9kZUF0dHJzW2F0dHIubmFtZV07XG4gICAgICAgIGlmIChhdHRyLnZhbHVlID09PSBvbGRWYWx1ZSkge1xuICAgICAgICAgICAgY29udGludWU7XG4gICAgICAgIH1cbiAgICAgICAgb2xkRWxlbWVudC5zZXRBdHRyaWJ1dGUoYXR0ci5uYW1lLCBhdHRyLnZhbHVlKTtcbiAgICB9XG4gICAgZm9yIChsZXQgaW5kZXggaW4gb2xkTm9kZUF0dHJzKSB7XG4gICAgICAgIG9sZEVsZW1lbnQucmVtb3ZlQXR0cmlidXRlKGluZGV4KTtcbiAgICB9XG59XG5leHBvcnRzLlBhdGNoRE9NQXR0cmlidXRlcyA9IFBhdGNoRE9NQXR0cmlidXRlcztcbi8vIyBzb3VyY2VNYXBwaW5nVVJMPXBhdGNoLmpzLm1hcCIsIlwidXNlIHN0cmljdFwiO1xuT2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsIFwiX19lc01vZHVsZVwiLCB7IHZhbHVlOiB0cnVlIH0pO1xuZXhwb3J0cy5pc0VxdWFsID0gZXhwb3J0cy5SYW5kb21JRCA9IGV4cG9ydHMudHJ1bmNhdGVBcnJheSA9IGV4cG9ydHMuY3JlYXRlTWFwID0gZXhwb3J0cy5oYXMgPSBleHBvcnRzLkJsYW5rID0gZXhwb3J0cy5Ub0tlYmFiQ2FzZSA9IHZvaWQgMDtcbmNvbnN0IGhhc093blByb3BlcnR5ID0gT2JqZWN0LnByb3RvdHlwZS5oYXNPd25Qcm9wZXJ0eTtcbmZ1bmN0aW9uIFRvS2ViYWJDYXNlKHN0cikge1xuICAgIGNvbnN0IHJlc3VsdCA9IHN0ci5yZXBsYWNlKC9bQS1aXFx1MDBDMC1cXHUwMEQ2XFx1MDBEOC1cXHUwMERFXS9nLCAobWF0Y2gpID0+ICctJyArIG1hdGNoLnRvTG93ZXJDYXNlKCkpO1xuICAgIHJldHVybiBzdHJbMF0gPT09IHN0clswXS50b1VwcGVyQ2FzZSgpID8gcmVzdWx0LnN1YnN0cmluZygxKSA6IHJlc3VsdDtcbn1cbmV4cG9ydHMuVG9LZWJhYkNhc2UgPSBUb0tlYmFiQ2FzZTtcbmZ1bmN0aW9uIEJsYW5rKCkgeyB9XG5leHBvcnRzLkJsYW5rID0gQmxhbms7XG5CbGFuay5wcm90b3R5cGUgPSBPYmplY3QuY3JlYXRlKG51bGwpO1xuZnVuY3Rpb24gaGFzKG1hcCwgcHJvcGVydHkpIHtcbiAgICByZXR1cm4gaGFzT3duUHJvcGVydHkuY2FsbChtYXAsIHByb3BlcnR5KTtcbn1cbmV4cG9ydHMuaGFzID0gaGFzO1xuZnVuY3Rpb24gY3JlYXRlTWFwKCkge1xuICAgIHJldHVybiBuZXcgQmxhbmsoKTtcbn1cbmV4cG9ydHMuY3JlYXRlTWFwID0gY3JlYXRlTWFwO1xuZnVuY3Rpb24gdHJ1bmNhdGVBcnJheShhcnIsIGxlbmd0aCkge1xuICAgIHdoaWxlIChhcnIubGVuZ3RoID4gbGVuZ3RoKSB7XG4gICAgICAgIGFyci5wb3AoKTtcbiAgICB9XG59XG5leHBvcnRzLnRydW5jYXRlQXJyYXkgPSB0cnVuY2F0ZUFycmF5O1xuZnVuY3Rpb24gUmFuZG9tSUQoKSB7XG4gICAgcmV0dXJuIE1hdGgucmFuZG9tKCkudG9TdHJpbmcoMzYpLnN1YnN0cigyLCA5KTtcbn1cbmV4cG9ydHMuUmFuZG9tSUQgPSBSYW5kb21JRDtcbmZ1bmN0aW9uIGlzRXF1YWwoYSwgYikge1xuICAgIGNvbnN0IGFQcm9wcyA9IE9iamVjdC5nZXRPd25Qcm9wZXJ0eU5hbWVzKGEpO1xuICAgIGNvbnN0IGJQcm9wcyA9IE9iamVjdC5nZXRPd25Qcm9wZXJ0eU5hbWVzKGIpO1xuICAgIGlmIChhUHJvcHMubGVuZ3RoICE9IGJQcm9wcy5sZW5ndGgpIHtcbiAgICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cbiAgICBmb3IgKGxldCBpID0gMDsgaSA8IGFQcm9wcy5sZW5ndGg7IGkrKykge1xuICAgICAgICBjb25zdCBwcm9wTmFtZSA9IGFQcm9wc1tpXTtcbiAgICAgICAgaWYgKGFbcHJvcE5hbWVdICE9PSBiW3Byb3BOYW1lXSkge1xuICAgICAgICAgICAgcmV0dXJuIGZhbHNlO1xuICAgICAgICB9XG4gICAgfVxuICAgIHJldHVybiB0cnVlO1xufVxuZXhwb3J0cy5pc0VxdWFsID0gaXNFcXVhbDtcbi8vIyBzb3VyY2VNYXBwaW5nVVJMPXV0aWxzLmpzLm1hcCIsIlwidXNlIHN0cmljdFwiO1xuT2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsIFwiX19lc01vZHVsZVwiLCB7IHZhbHVlOiB0cnVlIH0pO1xuZXhwb3J0cy5tb3VudFRvID0gdm9pZCAwO1xuY29uc3QgcHJvbWlzZSA9IHJlcXVpcmUoXCJwcm9taXNlLXBvbHlmaWxsXCIpO1xuaWYgKCFzZWxmLlByb21pc2UpIHtcbiAgICBzZWxmLlByb21pc2UgPSBwcm9taXNlO1xufVxuY29uc3QgbmFtZXNwYWNlID0gc2VsZi5Qcm9taXNlO1xuZnVuY3Rpb24gbW91bnRUbyhwYXJlbnQpIHtcbiAgICBwYXJlbnQucHJvbWlzZXMgPSBuYW1lc3BhY2U7XG59XG5leHBvcnRzLm1vdW50VG8gPSBtb3VudFRvO1xuZXhwb3J0cy5kZWZhdWx0ID0gbmFtZXNwYWNlO1xuLy8jIHNvdXJjZU1hcHBpbmdVUkw9aW5kZXguanMubWFwIiwiXCJ1c2Ugc3RyaWN0XCI7XG5PYmplY3QuZGVmaW5lUHJvcGVydHkoZXhwb3J0cywgXCJfX2VzTW9kdWxlXCIsIHsgdmFsdWU6IHRydWUgfSk7XG5leHBvcnRzLm1vdW50VG8gPSB2b2lkIDA7XG5jb25zdCBuYW1lc3BhY2UgPSB7fTtcbmZ1bmN0aW9uIG1vdW50VG8ocGFyZW50KSB7XG4gICAgcGFyZW50LnJlYWxtID0gbmFtZXNwYWNlO1xufVxuZXhwb3J0cy5tb3VudFRvID0gbW91bnRUbztcbmV4cG9ydHMuZGVmYXVsdCA9IG5hbWVzcGFjZTtcbi8vIyBzb3VyY2VNYXBwaW5nVVJMPWluZGV4LmpzLm1hcCIsIlwidXNlIHN0cmljdFwiO1xuT2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsIFwiX19lc01vZHVsZVwiLCB7IHZhbHVlOiB0cnVlIH0pO1xuY29uc3QgcmVhbG0gPSByZXF1aXJlKFwiLi4vLi4vcmVhbG0vc3JjXCIpO1xuY29uc3QgbWFya3VwID0gcmVxdWlyZShcIi4uLy4uL21hcmt1cC9zcmNcIik7XG5jb25zdCBhbmltYXRpb25zID0gcmVxdWlyZShcIi4uLy4uL2FuaW1hdGlvbi9zcmNcIik7XG5jb25zdCBodHRwID0gcmVxdWlyZShcIi4uLy4uL2h0dHAvc3JjXCIpO1xuY29uc3QgcHJvbWlzZXMgPSByZXF1aXJlKFwiLi4vLi4vcHJvbWlzZXMvc3JjXCIpO1xuY29uc3Qgdm9pZGdyYW0gPSB7XG4gICAgcmVhbG06IHt9LFxuICAgIGFuaW1hdGlvbnM6IHt9LFxuICAgIGh0dHA6IHt9LFxuICAgIHByb21pc2VzOiB7fSxcbiAgICBtYXJrdXA6IHt9LFxufTtcbnJlYWxtLm1vdW50VG8odm9pZGdyYW0pO1xubWFya3VwLm1vdW50VG8odm9pZGdyYW0pO1xuYW5pbWF0aW9ucy5tb3VudFRvKHZvaWRncmFtKTtcbmh0dHAubW91bnRUbyh2b2lkZ3JhbSk7XG5wcm9taXNlcy5tb3VudFRvKHZvaWRncmFtKTtcbmlmICh3aW5kb3cpIHtcbiAgICB3aW5kb3cudm9pZGdyYW0gPSB2b2lkZ3JhbTtcbn1cbmV4cG9ydHMuZGVmYXVsdCA9IHZvaWRncmFtO1xuLy8jIHNvdXJjZU1hcHBpbmdVUkw9aW5kZXguanMubWFwIl19
