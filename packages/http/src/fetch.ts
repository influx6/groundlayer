import * as fetch from 'whatwg-fetch';

// polyfill global context also.
if (!self.fetch) {
  // @ts-ignore
  self.fetch = fetch.fetch;
  // @ts-ignore
  self.Headers = fetch.Headers;
  // @ts-ignore
  self.Request = fetch.Request;
  // @ts-ignore
  self.Response = fetch.Response;
  // @ts-ignore
  self.DOMException = fetch.DOMException;
}

export const fetch = self.fetch;

// @ts-ignore
export const Headers = self.Headers;

// @ts-ignore
export const Request = self.Request;

// @ts-ignore
export const Response = self.Response;

// @ts-ignore
export const DOMException = self.DOMException;
