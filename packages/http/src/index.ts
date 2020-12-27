import * as fetch from './fetch';
import * as http from './http';
import * as websocket from './websocket';
import * as rafPolyfill from "../../animation/src/raf-polyfill";
import * as Animation from "../../animation/src/anime";

const namespace = {
  ...fetch,
  ...http,
  ...websocket,
}

export function mountTo(parent: object & { http: object }) {
  parent.http = namespace;
}

export default namespace;
