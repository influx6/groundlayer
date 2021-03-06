/**
 * A cached reference to the hasOwnProperty function.
 */
const hasOwnProperty = Object.prototype.hasOwnProperty;

// KeyMap defines a map type with string keys and boolean pairs.
export type KeyMap = {
  [x: string]: boolean;
};

/**
 * A constructor function that will create blank objects.
 */
export function Blank() {}
Blank.prototype = Object.create(null);

/**
 * Used to prevent property collisions between our "map" and its prototype.
 * @param map The map to check.
 * @param property The property to check.
 * @return Whether map has property.
 */
export function has(map: object, property: string): boolean {
  return hasOwnProperty.call(map, property);
}

/**
 * Creates an map object without a prototype.
 */
// tslint:disable-next-line:no-any
export function createMap(): any {
  // tslint:disable-next-line:no-any
  return new (Blank as any)();
}

/**
 * Truncates an array, removing items up until length.
 * @param arr The array to truncate.
 * @param length The new length of the array.
 */
export function truncateArray(arr: Array<{} | null | undefined>, length: number) {
  while (arr.length > length) {
    arr.pop();
  }
}

/**
 * RandomID generates unique IDs for use as pseudo-private/protected names.
 * Similar in concept to
 * <http://wiki.ecmascript.org/doku.php?id=strawman:names>.
 *
 * The goals of this function are twofold:
 *
 * - Provide a way to generate a string guaranteed to be unique when compared
 *  to other strings generated by this function.
 * - Make the string complex enough that it is highly unlikely to be
 *  accidentally duplicated by hand (this is key if you're using `ID`
 *  as a private/protected name on an object).
 *
 *	Use:
 *
 *			var privateName = ID();
 *			var o = { 'public': 'foo' };
 *			o[privateName] = 'bar';
 *
 * Taken from https://gist.github.com/gordonbrander/2230317.
 */
export function RandomID() {
  // Math.random should be unique because of its seeding algorithm.
  // Convert it to base 36 (numbers + letters), and grab the first 9 characters
  // after the decimal.
  return Math.random().toString(36).substr(2, 9);
}

/**
 * IsEqual returns true/false if giving properties are equal.
 * @param a
 * @param b
 */
export function isEqual(a, b) {
  // Create arrays of property names
  const aProps = Object.getOwnPropertyNames(a);
  const bProps = Object.getOwnPropertyNames(b);

  // If number of properties is different,
  // objects are not equivalent
  if (aProps.length != bProps.length) {
    return false;
  }

  for (let i = 0; i < aProps.length; i++) {
    const propName = aProps[i];

    // If values of same property are not equal,
    // objects are not equivalent
    if (a[propName] !== b[propName]) {
      return false;
    }
  }

  // If we made it this far, objects
  // are considered equivalent
  return true;
}

export function decodeFormBody(body) {
  let form = new FormData();
  body
    .trim()
    .split('&')
    .forEach(function (bytes) {
      if (bytes) {
        let split = bytes.split('=');
        let name = split.shift().replace(/\+/g, ' ');
        let value = split.join('=').replace(/\+/g, ' ');
        form.append(decodeURIComponent(name), decodeURIComponent(value));
      }
    });
  return form;
}

export function parseHTTPHeaders(rawHeaders) {
  let headers = new Headers();
  // Replace instances of \r\n and \n followed by at least one space or horizontal tab with a space
  // https://tools.ietf.org/html/rfc7230#section-3.2;
  let preProcessedHeaders = rawHeaders.replace(/\r?\n[\t ]+/g, ' ');
  preProcessedHeaders.split(/\r?\n/).forEach(function (line) {
    let parts = line.split(':');
    let key = parts.shift().trim();
    if (key) {
      let value = parts.join(':').trim();
      headers.append(key, value);
    }
  });
  return headers;
}

export function fileReaderReady(reader) {
  return new Promise(function (resolve, reject) {
    reader.onload = function () {
      resolve(reader.result);
    };
    reader.onerror = function () {
      reject(reader.error);
    };
  });
}

export function readBlobAsArrayBuffer(blob) {
  let reader = new FileReader();
  let promise = fileReaderReady(reader);
  reader.readAsArrayBuffer(blob);
  return promise;
}

export function readBlobAsText(blob) {
  let reader = new FileReader();
  let promise = fileReaderReady(reader);
  reader.readAsText(blob);
  return promise;
}

export function readArrayBufferAsText(buf) {
  let view = new Uint8Array(buf);
  let chars = new Array(view.length);

  for (let i = 0; i < view.length; i++) {
    chars[i] = String.fromCharCode(view[i]);
  }
  return chars.join('');
}

export function bufferClone(buf) {
  if (buf.slice) {
    return buf.slice(0);
  } else {
    let view = new Uint8Array(buf.byteLength);
    view.set(new Uint8Array(buf));
    return view.buffer;
  }
}
