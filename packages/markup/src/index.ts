import * as utils from './utils';
import * as exts from './extensions';
import * as patch from './patch';
import * as mount from './mount';
import * as dom from './dom';

const namespace = {
  ...utils,
  ...exts,
  ...patch,
  ...dom,
  ...mount,
};

export function mountTo(parent: object & { markup: object }) {
  parent.markup = namespace;
}

export default namespace;
