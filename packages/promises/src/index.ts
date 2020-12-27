import * as promise from 'promise-polyfill';
import * as utils from "../../markup/src/utils";
import * as exts from "../../markup/src/extensions";
import * as patch from "../../markup/src/patch";
import * as dom from "../../markup/src/dom";
import * as mount from "../../markup/src/mount";

//@ts-ignore
if (!self.Promise!) {
  //@ts-ignore
  self.Promise = promise;
}

const namespace = self.Promise;

export function mountTo(parent: object & { promises: object }) {
  parent.promises = namespace;
}

export default namespace;
