import { createMap, has, truncateArray } from './utils';
import * as exts from './extensions';

// ELEMENT_NODE represents element node type value.
export const ELEMENT_NODE = 1;

// DOCUMENT_FRAGMENT_NODE represents documentFragment node type value.
export const DOCUMENT_FRAGMENT_NODE = 11;

// DOCUMENT_NODE represents document node type value.
export const DOCUMENT_NODE = 9;

// TEXT_NODE represents text node type value.
export const TEXT_NODE = 3;

// COMMENT_NODE represents comment node type value.
export const COMMENT_NODE = 8;

// ElementConstructor defines a interface which exposes a method to
// create a new Element.
export interface ElementConstructor {
  new (): HTMLElement;
}

// AttrMutatorMap defines a map of key and mutator pairs.
export type AttrMutatorMap = {
  [x: string]: Mutator;
};

// Attributes defines a map of key - value string pairs.
export type Attributes = {
  [x: string]: string;
};

// Key defines a unique string type to represent a giving key.
export type Key = string;

// NameOrCtorDef defines a type which can be either a string or a
// ElementConstructor.
export type NameOrCtorDef = string | ElementConstructor;

// Mutator is a function which may mutate the underline html element
// passed to it.
export type Mutator = (n: HTMLElement, nm: string, v: any) => void;

const attributes: AttrMutatorMap = createMap() as AttrMutatorMap;
attributes['style'] = applyStyle;

/**
 * Checks if the node is the root of a document. This is either a Document
 * or ShadowRoot. DocumentFragments are included for simplicity of the
 * implementation, though we only want to consider Documents or ShadowRoots.
 * @param node The node to check.
 * @return True if the node the root of a document, false otherwise.
 */
export function isDocumentRoot(node: Node): node is Document | ShadowRoot {
  return node.nodeType === 11 || node.nodeType === 9;
}

/**
 * Checks if the node is an Element. This is faster than an instanceof check.
 * @param node The node to check.
 * @return Whether or not the node is an Element.
 */
export function isElement(node: Node): node is Element {
  return node.nodeType === 1;
}

/**
 * Checks if the node is a text node. This is faster than an instanceof check.
 * @param node The node to check.
 * @return Whether or not the node is a Text.
 */
export function isText(node: Node): node is Text {
  return node.nodeType === 3;
}

/**
 * @param  node The node to start at, inclusive.
 * @param  root The root ancestor to get until, exclusive.
 * @return The ancestry of DOM nodes.
 */
export function getAncestry(node: Node, root: Node | null) {
  const ancestry: Node[] = [];
  let cur: Node | null = node;

  while (cur !== root) {
    const n: Node = cur!;
    ancestry.push(n);
    cur = n.parentNode;
  }

  return ancestry;
}

/**
 * return The root node of the DOM tree that contains this node.
 */
const getRootNode =
  // tslint:disable-next-line:no-any b/79476176
  (Node as any).prototype.getRootNode ||
  function (this: Node) {
    // tslint:disable-next-line:no-unnecessary-type-assertion b/77361044
    let cur: Node | null = this as Node;
    let prev = cur;

    while (cur) {
      prev = cur;
      cur = cur.parentNode;
    }

    return prev;
  };

// NodeMatcher defines an interface which defines a function to be
// used for matching a Node against some criteria.
export interface NodeMatcher {
  (n: Node): boolean;
}

/**
 * reverseCollectNodeWithBreadth runs across giving parent node in a breadth first manner, never
 * going below the parent's children level, checking if giving children node match
 * provided matcher.
 *
 * Walks the tree in reverse from last child of parent to first child.
 * It collects all matching nodes into provide matches array.
 *
 * @param parent the parent node to match through.
 * @param matcher the matcher function to use for match checks.
 * @param matches the array used to reverseCollected matched nodes.
 */
export function reverseCollectNodeWithBreadth(parent: Node, matcher: NodeMatcher, matches: Array<Node>): void {
  let cur: Node = parent.lastChild!;
  while (cur) {
    if (matcher(cur)) {
      matches.push(cur);
    }
    cur = cur.previousSibling!;
  }
}

/**
 * reverseFindNodeWithBreadth runs across giving parent node in a breadth first manner, never
 * going below the parent's children level, checking if giving children node match
 * provided matcher.
 *
 * Walks the tree in reverse from last child of parent to first child.
 *
 * @param parent the parent node to match through.
 * @param matcher the matcher function to use for match checks.
 */
export function reverseFindNodeWithBreadth(parent: Node, matcher: NodeMatcher): Node | null {
  let cur: Node = parent.lastChild!;

  while (cur) {
    if (matcher(cur)) {
      return cur;
    }
    cur = cur.previousSibling!;
  }
  return null;
}

/**
 * collectNodeWithBreadth runs across giving parent node in a breadth first manner, never
 * going below the parent's children level, checking if giving children node match
 * provided matcher.
 *
 * It collects all matching nodes into provide matches array.
 *
 * @param parent the parent node to match through.
 * @param matcher the matcher function to use for match checks.
 * @param matches the array used to collected matched nodes.
 */
export function collectNodeWithBreadth(parent: Node, matcher: NodeMatcher, matches: Array<Node>): void {
  let cur = parent.firstChild!;
  if (matcher(cur)) {
    matches.push(cur);
  }

  while (cur) {
    if (matcher(cur.nextSibling!)) {
      matches.push(cur);
    }

    cur = cur.nextSibling!;
  }
}

/**
 * collectNodeWithDepth runs across giving parent node in a depth first manner, never
 * going through other siblings, but going from parent to first child down the tree.
 *
 * It collects all matching nodes into provide matches array.
 *
 * @param parent the parent node to match through.
 * @param matcher the matcher function to use for match checks.
 * @param matches the array used to collected matched nodes.
 */
export function collectNodeWithDepth(parent: Node, matcher: NodeMatcher, matches: Array<Node>): void {
  let cur: Node = parent.firstChild!;
  while (cur) {
    if (matcher(cur)) {
      matches.push(cur);
    }
    cur = cur.firstChild!;
  }
}

/**
 * findNodeWithBreadth runs across giving parent node in a breadth first manner, never
 * going below the parent's children level, checking if giving children node match
 * provided matcher.
 *
 * @param parent the parent node to match through.
 * @param matcher the matcher function to use for match checks.
 */

export function findNodeWithBreadth(parent: Node, matcher: NodeMatcher): Node | null {
  let cur: Node = parent.firstChild!;
  while (cur) {
    if (matcher(cur)) {
      return cur;
    }
    cur = cur.nextSibling!;
  }
  return null;
}

/**
 * findNodeWithDepth runs across giving parent node in a depth first manner, never
 * going through other siblings, but going from parent to first child down the tree.
 *
 * @param parent the parent node to match through.
 * @param matcher the matcher function to use for match checks.
 */
export function findNodeWithDepth(parent: Node, matcher: NodeMatcher): Node | null {
  let cur: Node = parent.firstChild!;
  while (cur) {
    if (matcher(cur)) {
      return cur;
    }
    cur = cur.firstChild!;
  }
  return null;
}

/**
 * findDepthFirst runs across giving parent node in a depth first manner where each
 * child node is search depth first for a matching node.
 *
 * @param parent the parent node to match through.
 * @param matcher the matcher function to use for match checks.
 */
export function findDepthFirst(parent: Node, matcher: NodeMatcher): Node | null {
  let cur: Node = parent.firstChild!;
  while (cur) {
    const found = findNodeWithDepth(cur, matcher);
    if (found) {
      return found;
    }
    cur = cur.nextSibling!;
  }
  return null;
}

/**
 * collectDepthFirst runs across giving parent node in a depth first manner where each
 * child node is search depth first for a matching node.
 *
 * @param parent the parent node to match through.
 * @param matcher the matcher function to use for match checks.
 * @param matches the array used to collected matched nodes.
 */
export function collectDepthFirst(parent: Node, matcher: NodeMatcher, matches: Array<Node>): void {
  let cur: Node = parent.firstChild!;
  while (cur) {
    collectNodeWithDepth(cur, matcher, matches);
    cur = cur.nextSibling!;
  }
  return;
}

/**
 * findBreadthFirst runs across giving parent node in a depth first manner where each
 * child node is search depth first for a matching node.
 *
 * @param parent the parent node to match through.
 * @param matcher the matcher function to use for match checks.
 */
export function findBreadthFirst(parent: Node, matcher: NodeMatcher): Node | null {
  let cur: Node = parent.firstChild!;
  while (cur) {
    const found = findNodeWithBreadth(cur, matcher);
    if (found) {
      return found;
    }
    cur = cur.firstChild!;
  }
  return null;
}

// NodeFunction defines a function type which applies
// a change to a node.
export interface NodeFunction {
  (n: Node): void;
}

/**
 * applyChildNode applies a giving function to it's children recursively
 * down the tree in a depth first manner.
 *
 * @param parent parent node and tree to begin function apply.
 * @param fn function to apply against each node.
 */
export function applyChildNode(parent: Node, fn: NodeFunction) {
  let cur: Node = parent.firstChild!;
  while (cur) {
    fn(cur);
    cur = cur.nextSibling!;
  }
}

/**
 * applyEachChildNode applies a giving function to it's children recursively
 * down the tree in a depth first manner.
 *
 * @param parent parent node and tree to begin function apply.
 * @param fn function to apply against each node.
 */
export function applyEachChildNode(parent: Node, fn: NodeFunction) {
  let cur: Node = parent.firstChild!;
  while (cur) {
    fn(cur);

    applyEachChildNode(cur, fn);
    cur = cur.nextSibling!;
  }
}

/**
 * applyEachNode applies a giving function to a parent and it's children recursively
 * down the tree in a depth first manner.
 *
 * @param parent parent node and tree to begin function apply.
 * @param fn function to apply against each node.
 */
export function applyEachNode(parent: Node, fn: NodeFunction) {
  fn(parent);

  let cur: Node = parent.firstChild!;
  while (cur) {
    applyEachNode(cur, fn);
    cur = cur.nextSibling!;
  }
}

/**
 * reverseApplyEachNode applies a giving function to a child to parent.
 * It'recursively applies function from down to up the tree.
 *
 * @param parent parent node and tree to begin function apply.
 * @param fn function to apply against each node.
 */
export function reverseApplyEachNode(parent: Node, fn: NodeFunction) {
  let cur: Node = parent.lastChild!;
  while (cur) {
    reverseApplyEachNode(cur, fn);
    fn(cur);
    cur = cur.previousSibling!;
  }

  fn(parent);
}

/**
 * eachChildNode applies giving function to giving child nodes of
 * giving parent, applying the function to the node itself last.
 *
 * @param node
 * @param fn
 */
export function eachChildAndNode(node: Node, fn: NodeFunction): void {
  const list = node.childNodes;
  for (let i = 0; i < list.length; i++) {
    fn(list[i]);
  }
  fn(node);
}

/**
 * nodeListToArray returns a list of Nodes from a list NodeList.
 * @param items
 */
export function nodeListToArray(items: NodeListOf<ChildNode>): Array<Node> {
  const list = [] as Array<Node>;
  for (let i = 0; i < items.length; i++) {
    list.push(items[i]);
  }
  return list;
}

/**
 * eachNodeChild applies giving function to giving child nodes of
 * giving parent, applying the function to the node itself fist.
 *
 * @param node
 * @param fn
 */
export function eachNodeAndChild(node: Node, fn: NodeFunction): void {
  fn(node);
  const list = node.childNodes;
  for (let i = 0; i < list.length; i++) {
    fn(list[i]);
  }
}

/**
 * toHTML returns the html representation of giving node.
 * @param node
 * @param shallow
 */
export function toHTML(node: Node, shallow: boolean | false): string {
  const div = document.createElement('div');
  div.appendChild(node.cloneNode(!shallow));
  return div.innerHTML;
}

/**
 * collectBreadthFirst runs across giving parent node in a depth first manner where each
 * child node is search depth first for a matching node.
 *
 * @param parent the parent node to match through.
 * @param matcher the matcher function to use for match checks.
 * @param matches the array used to collected matched nodes.
 */
export function collectBreadthFirst(parent: Node, matcher: NodeMatcher, matches: Array<Node>): void {
  let cur: Node = parent.firstChild!;
  while (cur) {
    collectNodeWithBreadth(cur, matcher, matches);
    cur = cur.firstChild!;
  }
  return;
}

/**
 * @param node The node to get the activeElement for.
 * @return The activeElement in the Document or ShadowRoot
 *     corresponding to node, if present.
 */
export function getActiveElement(node: Node): Element | null {
  const root = getRootNode.call(node);
  return isDocumentRoot(root) ? root.activeElement : null;
}

/**
 * Gets the path of nodes that contain the focused node in the same document as
 * a reference node, up until the root.
 * @param node The reference node to get the activeElement for.
 * @param root The root to get the focused path until.
 */
export function getFocusedPath(node: Node, root: Node | null): Node[] {
  const activeElement = getActiveElement(node);

  if (!activeElement || !node.contains(activeElement)) {
    return [];
  }

  return getAncestry(activeElement, root);
}

/**
 * Like insertBefore, but instead instead of moving the desired node, instead
 * moves all the other nodes after.
 * @param parentNode
 * @param node
 * @param referenceNode
 */
export function moveBefore(parentNode: Node, node: Node, referenceNode: Node | null) {
  const insertReferenceNode = node.nextSibling;
  let cur = referenceNode;

  while (cur !== null && cur !== node) {
    const next = cur.nextSibling;
    parentNode.insertBefore(cur, insertReferenceNode);
    cur = next;
  }
}

/**
 * insertBefore inserts giving node before reference node.
 * If reference node is null, then a normal appendChild is used with
 * parent to insert giving node.
 * @param parentNode
 * @param node
 * @param referenceNode
 */
export function insertBefore(parentNode: Node, node: Node, referenceNode: Node | null) {
  if (referenceNode === null) {
    parentNode.appendChild(node);
    return null;
  }

  parentNode.insertBefore(node, referenceNode);
  return null;
}

/**
 * replaceNode replaces giving node with provided replacement node.
 * @param parentNode
 * @param node
 * @param replacement
 */
export function replaceNode(parentNode: Node, node: Node, replacement: Node) {
  if (replacement === null) {
    return null;
  }

  parentNode.replaceChild(replacement, node);
  return null;
}

/**
 * replaceNodeIf replaces giving node with provided replacement node by using
 * parent of giving node if and only if it has one.
 *
 * It returns true/false if a replacement was actually done.
 * @param targetNode
 * @param replacement
 */
export function replaceNodeIf(targetNode: Node, replacement: Node): boolean {
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

/**
 * Returns the namespace to use for the attribute.
 *
 * @param name
 */
export function getNamespace(name: string): string | undefined {
  if (name.lastIndexOf('xml:', 0) === 0) {
    return 'http://www.w3.org/XML/1998/namespace';
  }

  if (name.lastIndexOf('xlink:', 0) === 0) {
    return 'http://www.w3.org/1999/xlink';
  }

  return undefined;
}

/**
 * Applies an attribute or property to a given Element. If the value is null
 * or undefined, it is removed from the Element. Otherwise, the value is set
 * as an attribute.
 */
// tslint:disable-next-line:no-any
export function applyAttr(el: Element, name: string, value: any) {
  if (value == null) {
    el.removeAttribute(name);
  } else {
    const attrNS = getNamespace(name);
    if (attrNS) {
      el.setAttributeNS(attrNS, name, String(value));
    } else {
      el.setAttribute(name, String(value));
    }
  }
}

// applyAttrs applies a giving map of key-value pairs into a attribute
// list. If the giving value of a key in a map is null, then that key is
// removed from the Element.
export function applyAttrs(el: Element, values: Attributes): void {
  for (let key in values) {
    if (values[key] == null) {
      el.removeAttribute(key);
      continue;
    }
    el.setAttribute(key, values[key]);
  }
}

/**
 * Applies a property to a given Element.
 */
// tslint:disable-next-line:no-any
export function applyProp(el: Element, name: string, value: any) {
  // tslint:disable-next-line:no-any
  (el as any)[name] = value;
}

/**
 * Applies a value to a style declaration. Supports CSS custom properties by
 * setting properties containing a dash using CSSStyleDeclaration.setProperty.
 */
export function setStyleValue(style: CSSStyleDeclaration, prop: string, value: string) {
  if (prop.indexOf('-') >= 0) {
    style.setProperty(prop, value);
  } else {
    // TODO(tomnguyen) Figure out why this is necessary.
    // tslint:disable-next-line:no-any
    (style as any)[prop] = value;
  }
}

/**
 * Applies a style to an Element. No vendor prefix expansion is done for
 * property names/values.
 * @param el
 * @param name The attribute's name.
 * @param  style The style to set. Either a string of css or an object
 *     containing property-value pairs.
 */
export function applySVGStyle(el: SVGElement, name: string, style: string | { [k: string]: string }) {
  if (typeof style === 'string') {
    el.style.cssText = style;
  } else {
    el.style.cssText = '';
    const elStyle = el.style;

    for (const prop in style) {
      if (has(style, prop)) {
        setStyleValue(elStyle, prop, style[prop]);
      }
    }
  }
}

/**
 * Applies a style to an Element. No vendor prefix expansion is done for
 * property names/values.
 * @param el
 * @param name The attribute's name.
 * @param  style The style to set. Either a string of css or an object
 *     containing property-value pairs.
 */
export function applyStyle(el: HTMLElement, name: string, style: string | { [k: string]: string }) {
  if (typeof style === 'string') {
    el.style.cssText = style;
  } else {
    el.style.cssText = '';
    const elStyle = el.style;

    for (const prop in style) {
      if (has(style, prop)) {
        setStyleValue(elStyle, prop, style[prop]);
      }
    }
  }
}

/**
 * applyStyles applies a map of style attributes to an Element.
 * No vendor prefix expansion is done for
 * property names/values.
 * @param el is the main HTMLElement
 * @param  style The style to set. Either a string of css or an object
 *     containing property-value pairs.
 */
export function applyStyles(el: HTMLElement, style: string | Attributes) {
  if (typeof style === 'string') {
    el.style.cssText = style;
  } else {
    el.style.cssText = '';
    const elStyle = el.style;

    for (const prop in style) {
      if (has(style, prop)) {
        setStyleValue(elStyle, prop, style[prop]);
      }
    }
  }
}

/**
 * applyStyles applies a map of style attributes to an Element.
 * No vendor prefix expansion is done for
 * property names/values.
 * @param el is the main SVGElement
 * @param  style The style to set. Either a string of css or an object
 *     containing property-value pairs.
 */
export function applySVGStyles(el: SVGElement, style: string | Attributes) {
  if (typeof style === 'string') {
    el.style.cssText = style;
  } else {
    el.style.cssText = '';
    const elStyle = el.style;

    for (const prop in style) {
      if (has(style, prop)) {
        setStyleValue(elStyle, prop, style[prop]);
      }
    }
  }
}

/**
 * Updates a single attribute on an Element.
 * @param el
 * @param name The attribute's name.
 * @param value The attribute's value. If the value is an object or
 *     function it is set on the Element, otherwise, it is set as an HTML
 *     attribute.
 */
export function applyAttributeTyped(el: HTMLElement, name: string, value: {}) {
  const type = typeof value;

  if (type === 'object' || type === 'function') {
    applyProp(el, name, value);
  } else {
    applyAttr(el, name, value);
  }
}

/**
 * Gets the namespace to create an element (of a given tag) in.
 */
export function getNamespaceForTag(tag: string, parent: Node | null) {
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

/**
 * Records the element's attributes.
 * @param node The Element that may have attributes
 */
export function recordAttributes(node: Element): Attributes {
  const attrs: Attributes = {};
  const attributes = node.attributes;
  const length = attributes.length;
  if (!length) {
    return attrs;
  }

  // Use a cached length. The attributes array is really a live NamedNodeMap,
  // which exists as a DOM "Host Object" (probably as C++ code). This makes the
  // usual constant length iteration very difficult to optimize in JITs.
  for (let i = 0, j = 0; i < length; i += 1, j += 2) {
    const attr = attributes[i];
    attrs[attr.name] = attr.value;
  }
  return attrs;
}

/**
 * Creates an Element.
 * @param doc The document with which to create the Element.
 * @param nameOrCtor The tag or constructor for the Element.
 * @param key A key to identify the Element.
 * @param  content The underline html content for the Element.
 * @param  attributes The attributes map to apply to element.
 * @param  namespace The node namespace if specific to be used in creating element.
 */
export function createElement(
  doc: Document,
  nameOrCtor: NameOrCtorDef,
  key: Key,
  content: string,
  attributes: Attributes | null,
  namespace: string | '',
): Element {
  let el: Element;

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
  } else {
    el = doc.createElement(nameOrCtor);
  }

  // add key as attribute.
  el.setAttribute('_key', key);
  if (attributes) {
    applyAttrs(el, attributes);
  }

  // if content is provided then pass the content to the inner html.
  if (content.length > 0) {
    el.innerHTML = content;
  }

  return el;
}

/**
 * Creates a Text Node.
 * @param doc The document with which to create the Element.
 * @param text The content of giving text node.
 * @param key The unique key of giving text node.
 * @return
 */
export function createText(doc: Document, text: string | '', key: Key | ''): Node {
  const node = doc.createTextNode(text);
  exts.Objects.PatchWith(node, 'key', key);
  return node;
}

/**
 * Clears out any unvisited Nodes in a given range.
 * @param fromNode
 * @param endNode The node to clear until, exclusive.
 */
export function removeFromNode(fromNode: Node, endNode: Node | null) {
  const parentNode = fromNode.parentNode;
  let child: Node = fromNode;

  while (child !== endNode) {
    const next = child!.nextSibling;
    parentNode!.removeChild(child!);
    child = next!;
  }
}

/**
 *
 * export function fromBlob transform the providded Object blob into a byte slice.
 * @param o
 */
export function fromBlob(o) {
  if (o === null || o === undefined) {
    return;
  }

  let data = null;

  const fileReader = new FileReader();
  fileReader.onloadend = function () {
    // @ts-ignore
    data = new Uint8Array(fileReader.result);
  };

  fileReader.readAsArrayBuffer(o);

  return data;
}

/**
 *
 * export function fromFile transform the provided Object blob into a byte slice.
 * @param o
 */
export function fromFile(o) {
  if (o === null || o === undefined) {
    return;
  }

  let data = null;

  const fileReader = new FileReader();
  fileReader.onloadend = function () {
    // @ts-ignore
    data = new Uint8Array(fileReader.result);
  };

  fileReader.readAsArrayBuffer(o);

  return data;
}

/**
 *
 * toInputSourceCapability returns the events.InputDeviceCapabilities from the object.
 * @param: o
 */
export function toInputSourceCapability(o) {
  if (o === null || o === undefined) {
    return;
  }

  return {
    FiresTouchEvent: o.firesTouchEvent,
  };
}

/**
 *
 * export function toMotionData returns a motionData object from the Object.
 * @param o
 */
export function toMotionData(o) {
  let md = { X: 0.0, Y: 0.0, Z: 0.0 };

  if (o === null || o === undefined) {
    return md;
  }

  md.X = o.x;
  md.Y = o.y;
  md.Z = o.z;
  return md;
}

/**
 * export function toRotationData returns a RotationData object from the Object.
 * @param o
 */
export function toRotationData(o) {
  if (o === null || o === undefined) {
    return;
  }

  const md = {};
  // @ts-ignore
  md.Alpha = o.alpha;
  // @ts-ignore
  md.Beta = o.beta;
  // @ts-ignore
  md.Gamma = o.gamma;
  return md;
}

/**
 *
 * export function toMediaStream returns a events.MediaStream object.
 * @param o
 */
export function toMediaStream(o) {
  if (o === null || o === undefined) {
    return;
  }

  const stream = { Audios: [], Videos: [] };
  // @ts-ignore
  stream.Active = o.active;
  // @ts-ignore
  stream.Ended = o.ended;
  // @ts-ignore
  stream.ID = o.id;
  // @ts-ignore
  stream.Audios = [];
  // @ts-ignore
  stream.Videos = [];

  let audioTracks = o.getAudioTracks();
  if (audioTracks !== null && audioTracks !== undefined) {
    for (let i = 0; i < audioTracks.length; i++) {
      let track = audioTracks[i];
      let settings = track.getSettings();

      stream.Audios.push({
        // @ts-ignore
        Enabled: track.enabled,
        // @ts-ignore
        ID: track.id,
        // @ts-ignore
        Kind: track.kind,
        // @ts-ignore
        Label: track.label,
        // @ts-ignore
        Muted: track.muted,
        // @ts-ignore
        ReadyState: track.readyState,
        // @ts-ignore
        Remote: track.remote,
        AudioSettings: {
          // @ts-ignore
          ChannelCount: settings.channelCount.Int(),
          // @ts-ignore
          EchoCancellation: settings.echoCancellation,
          // @ts-ignore
          Latency: settings.latency,
          // @ts-ignore
          SampleRate: settings.sampleRate.Int64(),
          // @ts-ignore
          SampleSize: settings.sampleSize.Int64(),
          // @ts-ignore
          Volume: settings.volume,
          // @ts-ignore
          MediaTrackSettings: {
            // @ts-ignore
            DeviceID: settings.deviceId,
            // @ts-ignore
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
        // @ts-ignore
        Enabled: track.enabled,
        // @ts-ignore
        ID: track.id,
        // @ts-ignore
        Kind: track.kind,
        // @ts-ignore
        Label: track.label,
        // @ts-ignore
        Muted: track.muted,
        // @ts-ignore
        ReadyState: track.readyState,
        // @ts-ignore
        Remote: track.remote,
        // @ts-ignore
        VideoSettings: {
          // @ts-ignore
          AspectRatio: settings.aspectRation,
          // @ts-ignore
          FrameRate: settings.frameRate,
          // @ts-ignore
          Height: settings.height.Int64(),
          // @ts-ignore
          Width: settings.width.Int64(),
          // @ts-ignore
          FacingMode: settings.facingMode,
          // @ts-ignore
          MediaTrackSettings: {
            // @ts-ignore
            DeviceID: settings.deviceId,
            // @ts-ignore
            GroupID: settings.groupId,
          },
        },
      });
    }
  }

  return stream;
}

/**
 * toTouches transform giving object into a touch list.
 * @param o
 */
export function toTouches(o) {
  if (o === null || o === undefined) {
    return;
  }

  const touches = [];
  for (let i = 0; i < o.length; i++) {
    let ev = o.item(i);
    touches.push({
      // @ts-ignore
      ClientX: ev.clientX,
      // @ts-ignore
      ClientY: ev.clientY,
      // @ts-ignore
      OffsetX: ev.offsetX,
      // @ts-ignore
      OffsetY: ev.offsetY,
      // @ts-ignore
      PageX: ev.pageX,
      // @ts-ignore
      PageY: ev.pageY,
      // @ts-ignore
      ScreenX: ev.screenX,
      // @ts-ignore
      ScreenY: ev.screenY,
      // @ts-ignore
      Identifier: ev.identifier,
    });
  }

  return touches;
}

/**
 * toGamepad returns a Gamepad struct from the js object.
 * @param o
 */
export function toGamepad(o) {
  let pad = {};

  if (o === null || o === undefined) {
    return pad;
  }

  // @ts-ignore
  pad.DisplayID = o.displayId;
  // @ts-ignore
  pad.ID = o.id;
  // @ts-ignore
  pad.Index = o.index.Int();
  // @ts-ignore
  pad.Mapping = o.mapping;
  // @ts-ignore
  pad.Connected = o.connected;
  // @ts-ignore
  pad.Timestamp = o.timestamp;
  // @ts-ignore
  pad.Axes = [];
  // @ts-ignore
  pad.Buttons = [];
  // @ts-ignore

  let axes = o.axes;
  if (axes !== null && axes !== undefined) {
    for (let i = 0; i < axes.length; i++) {
      // @ts-ignore
      pad.Axes.push(axes[i]);
    }
  }

  let buttons = o.buttons;
  if (buttons !== null && buttons !== undefined) {
    for (let i = 0; i < buttons.length; i++) {
      const button = buttons[i];
      // @ts-ignore
      pad.Buttons.push({
        Value: button.value,
        Pressed: button.pressed,
      });
    }
  }

  return pad;
}

// toDataTransfer returns a transfer object from the Object.
export function toDataTransfer(o) {
  if (o === null || o === undefined) {
    return;
  }

  let dt = {};
  // @ts-ignore
  dt.DropEffect = o.dropEffect;
  // @ts-ignore
  dt.EffectAllowed = o.effectAllowed;
  // @ts-ignore
  dt.Types = o.types;
  // @ts-ignore
  dt.Items = [];

  const items = o.items;
  if (items !== null && items !== undefined) {
    for (let i = 0; i < items.length; i++) {
      const item = items.DataTransferItem(i);

      // @ts-ignore
      dt.Items.push({
        Name: item.name,
        Size: item.size.Int(),
        Data: fromFile(item),
      });
    }
  }

  // @ts-ignore
  dt.Files = [];

  const files = o.files;
  if (files !== null && files !== undefined) {
    for (let i = 0; i < files.length; i++) {
      const item = files[i];

      // @ts-ignore
      dFiles.push({
        // @ts-ignore
        Name: item.name,
        // @ts-ignore
        Size: item.size.Int(),
        // @ts-ignore
        Data: fromFile(item),
      });
    }
  }

  return dt;
}

// /**
//  * Aligns the virtual Node definition with the actual DOM, moving the
//  * corresponding DOM node to the correct location or creating it if necessary.
//  * @param nameOrCtor The name or constructor for the Node.
//  * @param key The key used to identify the Node.
//  */
// function alignWithDOM(currentNode:Node, matcher:) {
//   const existingNode = getMatchingNode(currentNode, nameOrCtor, key);
//   const node = existingNode || createNode(nameOrCtor, key);
//
//   // If we are at the matching node, then we are done.
//   if (node === currentNode) {
//     return;
//   }
//
//   // Re-order the node into the right position, preserving focus if either
//   // node or currentNode are focused by making sure that they are not detached
//   // from the DOM.
//   if (focusPath.indexOf(node) >= 0) {
//     // Move everything else before the node.
//     moveBefore(currentParent!, node, currentNode);
//   } else {
//     currentParent!.insertBefore(node, currentNode);
//   }
//
//   currentNode = node;
// }

/**
 * Finds the matching node, starting at `node` and looking at the subsequent
 * siblings if a key is used.
 * @param matchNode The node to start looking at.
 * @param matcher The key used to identify the Node.
 */
export function getMatchingNode(matchNode: Node | null, matcher: NodeMatcher): Node | null {
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
