import * as rafPolyfill from './raf-polyfill';
import * as Animation from './anime';

const namespace = {
  ...rafPolyfill,
  ...Animation,
}

export function mountTo(parent: object & { animations: object }) {
  parent.animations = namespace;
}

export default namespace;
