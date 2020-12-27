import * as dom from './dom';
import * as utils from './utils';
import * as exts from './extensions';

import { COMMENT_NODE, DOCUMENT_FRAGMENT_NODE, ELEMENT_NODE, TEXT_NODE } from './dom';

// NodeDictator defines an interface type which exposes methods
// that operate on DOM Nodes to patch into a patchTree process allowing
// the checks of a giving node if they are the same and changed.
export interface NodeDictator {
  Same(n: Node, m: Node): boolean;
  Changed(n: Node, m: Node): boolean;
}

// DefaultNodeDictator implements the NodeDictator interface.
export const DefaultNodeDictator: NodeDictator = {
  Same: (n: Node, m: Node): boolean => {
    const sameNode = n.nodeType === m.nodeType && n.nodeName === m.nodeName;
    if (!sameNode) {
      return false;
    }

    if (n.nodeType !== dom.ELEMENT_NODE) {
      return true;
    }

    const nElem = n as Element;
    const mElem = m as Element;
    return nElem.id === mElem.id;
  },
  Changed: (n: Node, m: Node): boolean => {
    if (n.nodeType === dom.TEXT_NODE && m.nodeType === dom.TEXT_NODE) {
      return n.textContent === m.textContent;
    }
    if (n.nodeType === dom.COMMENT_NODE && m.nodeType === dom.COMMENT_NODE) {
      return n.textContent === m.textContent;
    }

    const nElem = n as Element;
    const mElem = m as Element;
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

/*
 * JSONDictator defines an interface which exposes method to verify
 * the same and change state of a DOM node and a JSON node.
 */
export interface JSONDictator {
  Same(n: Node, m: JSONNode): boolean;
  Changed(n: Node, m: JSONNode): boolean;
}

// DefaultJSONDictator implements the JSONDictator interface.
export const DefaultJSONDictator: JSONDictator = {
  Same: (n: Node, m: JSONNode): boolean => {
    const sameNode = n.nodeType === m.type && n.nodeName === m.name;
    if (!sameNode) {
      return false;
    }

    if (n.nodeType !== dom.ELEMENT_NODE) {
      return true;
    }

    const nElem = n as Element;
    return nElem.id === m.id;
  },
  Changed: (n: Node, m: JSONNode): boolean => {
    if (n.nodeType === dom.TEXT_NODE && m.type === dom.TEXT_NODE) {
      return n.textContent !== m.content;
    }
    if (n.nodeType === dom.COMMENT_NODE && m.type === dom.COMMENT_NODE) {
      return n.textContent !== m.content;
    }

    const tnode = n as Element;
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

// JSONAttr represent giving json format for a attribute for
// a giving json node.
export interface JSONAttr {
  Key: string;
  Value: any;
}

// JSONEvent represent giving json format for a event for
// a giving json node.
export interface JSONEvent {
  Name: string;
  Targets: Array<string>;
}

// JSONNode defines an interface exposing the expected
// type representing a giving JSON encoded version of
// a DOM update node or request.
export interface JSONNode {
  id: string; // id defines the dom id attribute for giving node.
  tid: string; // tid gives the unique id dom node.
  atid: string; // atid tells the ancestor-id which was the reconciliation process. It tells us what the id of old node.
  ref: string; // ref defines a reference value provided to match nodes from rendering and rendered.
  type: number; // type tells us which type number representing of node e.g 3, 8, 11.
  typeName: string; // typeName tells us which type of node e.g document-fragment, text-node
  name: string; // name tells the name of giving node e.g div, span.
  content: string; // content represents text content if this is a content.
  namespace: string; // namespace tells the DOM namespace for giving node.
  removed: boolean; // removed tells us if this represents a removal JSONNode instruction.
  attrs: Array<JSONAttr>; // attrs defines the array of node attributes.
  events: Array<JSONEvent>; // events defines the event map for giving attribute.
  children: Array<JSONNode>; // children contains child node descriptions for giving root.
}

// JSONNodeFunction defines a function interface for applying
// a function to a JSONNode.
export interface JSONNodeFunction {
  (n: JSONNode): void;
}

/**
 * ToJSONNode transforms a DOM node into a JSONNode description object.
 * It recursively walks down the dom tree and generates giving node unless if
 * shallow is set to true.
 *
 * If a node has no ID then a random one is generated for it.
 *
 * @param node is a DOM node to transform info a JSONNode.
 * @param shallow is a flag if only the node itself without the children should be transformed.
 * @param parentNode is the reference to the parent JSONNode if any.
 * @constructor
 */
export function ToJSONNode(
  node: Node | string,
  shallow: boolean | false,
  parentNode: JSONNode | null,
): Array<JSONNode> {
  const list = new Array<JSONNode>();

  if (typeof node === 'string') {
    const pub = document.createElement('div');
    pub.innerHTML = (node as string).trim();
    dom.applyChildNode(pub, function (child: Node) {
      list.push(NodeToJSONNode(child, shallow, parentNode));
    });
    return list;
  }

  list.push(NodeToJSONNode(node as Node, shallow, parentNode));
  return list;
}

/**
 * ToJSONNode transforms a DOM node into a JSONNode description object.
 * It recursively walks down the dom tree and generates giving node unless if
 * shallow is set to true.
 *
 * If a node has no ID then a random one is generated for it.
 *
 * @param node is a DOM node to transform info a JSONNode.
 * @param shallow is a flag if only the node itself without the children should be transformed.
 * @param parentNode is the reference to the parent JSONNode if any.
 * @constructor
 */
export function NodeToJSONNode(node: Node, shallow: boolean | false, parentNode: JSONNode | null): JSONNode {
  const jnode = {} as JSONNode;
  jnode.children = [];
  jnode.events = [];
  jnode.attrs = [];
  jnode.namespace = '';
  jnode.type = node.nodeType;
  jnode.name = node.nodeName.toLowerCase();
  jnode.id = exts.Objects.GetAttrWith(node, '_id') as string;
  jnode.tid = exts.Objects.GetAttrWith(node, '_tid') as string;
  jnode.ref = exts.Objects.GetAttrWith(node, '_ref') as string;
  jnode.atid = exts.Objects.GetAttrWith(node, '_atid') as string;

  const elem = node as Element;
  if (elem === null) return jnode;

  // @ts-ignore
  if (node._tid) {
    // @ts-ignore
    jnode.tid = node._tid;
  }

  switch (node.nodeType) {
    case TEXT_NODE:
      jnode.typeName = 'Text';
      jnode.content = node.textContent!;
      return jnode;
    case COMMENT_NODE:
      jnode.typeName = 'Comment';
      jnode.content = node.textContent!;
      return jnode;
    case ELEMENT_NODE:
      jnode.typeName = 'Element';
      jnode.children = new Array<JSONNode>();
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
    jnode.id = elem.getAttribute('id')!;
  } else {
    jnode.id = utils.RandomID();
  }

  if (jnode.ref === '' && !exts.Objects.isNullOrUndefined(parentNode)) {
    jnode.ref = jnode.id;
  }

  if (elem.hasAttribute('_ref')) {
    jnode.ref = elem.getAttribute('_ref')!;
  } else {
    if (!exts.Objects.isNullOrUndefined(parentNode)) {
      if (parentNode!.ref !== '') {
        jnode.ref = parentNode!.ref + '/' + jnode.id;
      } else {
        jnode.ref = parentNode!.id + '/' + jnode.id;
      }
    }
  }
  if (elem.hasAttribute('_tid')) {
    jnode.tid = elem.getAttribute('_tid')!;
  }

  if (elem.hasAttribute('_atid')) {
    jnode.atid = elem.getAttribute('_atid')!;
  }

  // map out events:
  for (var i = 0; i < elem.attributes.length; i++) {
    let attr = elem.attributes.item(i);
    if (attr == null) continue;
    if (!attr.name.startsWith('event-')) {
      continue;
    }
    let eventName = attr.name.replace('event-', '');
    jnode.events.push(JSONEvent(eventName, attr.value.split('|')));
  }

  if (!shallow) {
    dom.applyChildNode(node, function (child: Node) {
      if (child instanceof Text || child instanceof Comment) {
        jnode.children.push(NodeToJSONNode(child, false, jnode));
        return;
      }

      const childElem = child as Element;
      if (!childElem.hasAttribute('id')) {
        childElem.id = utils.RandomID();
      }
      jnode.children.push(NodeToJSONNode(childElem, false, jnode));
    });
  }

  return jnode;
}

/**
 * JSONEvent returns an JSONEvent object for giving string describing
 * an event subscription.
 *
 * @param eventDesc is the event string in JSONEvent format. <EventName>-<PreventDefaultBit><StopPropagationBit>.
 * @constructor
 */
export function JSONEvent(name: string, targets: Array<string>): JSONEvent {
  const event = {} as JSONEvent;
  event.Name = name;
  event.Targets = targets;
  return event;
}

/**
 * applyJSONNodeFunction applies provided function to the json node and it's
 * children list.
 *
 * @param node is the target JSONNode.
 * @param fn is function to apply to JSONNode.
 */
export function applyJSONNodeFunction(node: JSONNode, fn: JSONNodeFunction): void {
  fn(node);
  node.children.forEach(function (child) {
    applyJSONNodeFunction(child, fn);
  });
}

/**
 * applyJSONNodeKidsFunction applies provided function to the json node children first
 * then parent.
 *
 * @param node is the target JSONNode.
 * @param fn is function to apply to JSONNode.
 */
export function applyJSONNodeKidsFunction(node: JSONNode, fn: JSONNodeFunction): void {
  node.children.forEach(function (child) {
    applyJSONNodeFunction(child, fn);
  });
  fn(node);
}

/**
 * isJSONNode provides a type checker for verifying the type of n to be
 * a JSONNode;
 * @param n
 */
export function isJSONNode(n: any): n is JSONNode {
  const hasID = typeof (<JSONNode>n).id !== 'undefined';
  const hasRef = typeof (<JSONNode>n).ref !== 'undefined';
  const hasTid = typeof (<JSONNode>n).tid !== 'undefined';
  const hasTypeName = typeof (<JSONNode>n).typeName !== 'undefined';
  return hasID && hasRef && hasTypeName && hasTid;
}

/**
 * findElement seeks giving DOMNode if available targeted by giving JSONNode in provided
 * parent
 *
 * @param desc JSONNode describing target element.
 * @param parent contain nodes we could be targeting.
 */
export function findElement(desc: JSONNode, parent: Element): Element | null {
  const selector: string = desc.name + '#' + desc.id;
  const targets = parent.querySelectorAll(selector)!;

  // if non is found then attempt the use of attribute selectors.
  // do a specific attribute selector with a tid.
  if (targets.length === 0) {
    let attrSelector = desc.name + `[_tid='${desc.tid}']`;
    let target = parent.querySelector(attrSelector)!;
    if (target) {
      return target;
    }

    // could it be the tid represents a change, so the element is still represented by it's
    // ancestral tid?
    attrSelector = desc.name + `[_atid='${desc.atid}']`;
    target = parent.querySelector(attrSelector)!;
    if (target) {
      return target;
    }

    // last will be to search by reference value
    attrSelector = desc.name + `[_ref='${desc.ref}']`;
    return parent.querySelector(attrSelector)!;
  }

  if (targets.length === 1) {
    return targets[0];
  }

  // search the list for any match item.
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

/**
 * findElementByRef seeks giving DOMNode if available by running through parents children
 * using the ref string which represents a tree of parent to child ID attribute list.
 *
 * @param ref JSONNode describing target element.
 * @param parent contain nodes we could be targeting.
 */
export function findElementbyRef(ref: string, parent: Element): Element | null {
  // split id map and attach "#" to indicate we are dealing with a id attribute.
  const ids: Array<string> = ref.split('/').map(function (elem) {
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

  // parent checking can be tricky, because one example could be we are provided
  // a static parent which has no relation to the id tree, so we would need to check
  // from the first ID if any such element exists, but another instance is the parent
  // could actually be the first ID in the list, so we would have to plan accordingly
  // in that form.

  // if we are dealing with a parent who is the actual first id in the list then
  // shift our list to account for this.
  const first = ids[0];
  if (parent.id == first.substr(1)) {
    ids.shift();
  }

  // Go through the id list and pull out the next child in the tree.
  let cur: Element = parent.querySelector(ids.shift()!)!;
  while (cur) {
    if (ids.length === 0) {
      return cur;
    }
    cur = cur.querySelector(ids.shift()!)!;
  }
  return cur;
}

/**
 * findElementParentByRef seeks giving DOMNode parent if available by running through parents children
 * using the ref string until it reaches a node it can't find which will represent where the node should
 * exists. It returns the parent of that node..
 *
 * @param ref JSONNode describing target element.
 * @param parent contain nodes we could be targeting.
 */
export function findElementParentbyRef(ref: string, parent: Element): Element | null {
  // split id map and attach "#" to indicate we are dealing with a id attribute.
  const ids: Array<string> = ref.split('/').map(function (elem) {
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

  // pop the last target as it's the parent we want.
  ids.pop();

  // if we are dealing with a parent who is the actual first id in the list then
  // shift our list to account for this.
  const first = ids[0];
  if (parent.id == first.substr(1)) {
    ids.shift();
  }

  // Go through the id list and pull out the next child in the tree.
  let cur: Element = parent.querySelector(ids.shift()!)!;
  while (cur) {
    if (ids.length === 0) {
      return cur;
    }
    cur = cur.querySelector(ids.shift()!)!;
  }
  return cur;
}

// JSONMaker defines an interface which exposes a method to create
// a DOM node from a JSON node.
export interface JSONMaker {
  /**
   * Make will take provided JSONNode and create
   * appropriate node representation for giving json.
   *
   * @param doc is the Document to use for creating elements.
   * @param n is the JSONNode describing the DOM node to generate.
   * @param shallow indicates if we want to generate a deep node or without it's children.
   * @param skipRemoved indicates if we want to skip removed fragments in the children list of JSON Node.
   */
  Make(doc: Document, n: JSONNode, shallow: boolean, skipRemoved: boolean): Node;
}

// DefaultJSONMaker implements the JSONMaker interface.
export const DefaultJSONMaker: JSONMaker = {
  Make: jsonMaker,
};

// jsonMaker will generate a complete DOM node either complete with children or shallow
// which represented the giving JSONNode description.
export function jsonMaker(
  doc: Document,
  descNode: JSONNode,
  shallow: boolean | false,
  skipRemoved: boolean | true,
): Node {
  if (descNode.type === COMMENT_NODE) {
    const node = doc.createComment(descNode.content);
    exts.Objects.PatchWith(node, '_id', descNode.id);
    exts.Objects.PatchWith(node, '_ref', descNode.ref);
    exts.Objects.PatchWith(node, '_tid', descNode.tid);
    exts.Objects.PatchWith(node, '_atid', descNode.atid);
    return node;
  }

  if (descNode.type === TEXT_NODE) {
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

  let node: Element;
  if (descNode.namespace.length !== 0) {
    node = doc.createElement(descNode.name);
  } else {
    node = doc.createElementNS(descNode.namespace, descNode.name);
  }

  // attach information as from format into node JS Object itself.
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

  // if it has a removed marker, i am yet unsure why you would want
  // to render a removed JSON Node but rather than add the other bits
  // you just get a shallow render
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

// NodeMap defines a map of Node keys - boolean values pairs.
export type NodeMap = {
  // @ts-ignore
  [x: Node]: boolean;
};

// NodeMap defines a map of string keys - boolean values pairs.
export type NodeTidMap = {
  [x: string]: boolean;
};

/**
 * PatchJSONNodeTree handles a general case algorithm with a simplified steps for handling
 * tree patching and updating for a giving live node with that of a JSON description of
 * giving node changes.
 *
 * PatchJSONNodeTree works with a whole DOM node or a DOM fragment, in that we mean a complete
 * set of DOM nodes for both the old and new set, which matches the children to be updated.
 * The idea is when using this, you provide the static node where old nodes are mounted
 * and you use the JSON fragment to power the update.
 *
 * Consider using StreamJSONNodes to use a more stream or partial friendly update mechanism
 * which allows updating sparse parts of the DOM trees.
 *
 * Note: This function assumes the oldNodeOrMount for the initial call is either empty or
 * contains previously rendered nodes only. That is there exists no static nodes within, as
 * anything not matching the update will be removed after reconciliation.
 *
 * IMPORTANT: The ID value of a DOM node is very important in this algorithm, as JSONNode.id
 * value will be used to query the parent for the target node. if a giving node's id is not found
 * then an attempt is made with the Ancestor ID (JSONNode.aid), if it fails, then the node is created
 * and swapped with whatever node is found in that position.
 *
 * @param fragment: A JSONNode fragment which will be used for DOM update.
 * @param mount: A static node or last parent node which holds the target node for the giving JSON fragment.
 * @param dictator: Dictator helps us decide if the node is the same or changed.
 * @param maker: Maker helps create a real DOM node from a JSONNode.
 * @return
 **/
export function PatchJSONNodeTree(fragment: JSONNode, mount: Element, dictator: JSONDictator, maker: JSONMaker): void {
  // Attempt to find the giving node within the mount, if we can't find it then add it as a new node.
  let targetNode: Element = findElement(fragment, mount)!;
  if (exts.Objects.isNullOrUndefined(targetNode)) {
    const tNode = maker.Make(document, fragment, false, true);
    mount.appendChild(tNode);
    return;
  }

  PatchJSONNode(fragment, targetNode, dictator, maker!);
}

/**
 * PatchJSONNode provides a DOM node patching approach for a giving JSON fragment and it's target DOM node.
 * It uses a recursive function which runs through a fragment and it's target node and the children changes
 * of that target node as well.
 *
 * Generally you should use PatchJSONNode if you wish to apply a JSONode change directly to the target node as
 * PatchJSONNodeTree will first seek out the target node from the parent or mount node.
 *
 * @param fragment is the JSONNode containing change for giving target.
 * @param targetNode is the target DOM node to apply changes to.
 * @param dictator is used to decide the same-ness and change-ness state of target node.
 * @param maker is used to create a new DOM node from a JSONNode.
 * @constructor
 */
export function PatchJSONNode(fragment: JSONNode, targetNode: Element, dictator: JSONDictator, maker: JSONMaker): void {
  // if we are not dealing with exactly the same node as fragment as far as the.
  // then dictator is concerned, then we must do a swap after creating node.
  if (!dictator.Same(targetNode, fragment)) {
    const tNode = maker.Make(document, fragment, false, true);
    dom.replaceNode(targetNode.parentNode!, targetNode, tNode);
    return;
  }

  if (!dictator.Changed(targetNode, fragment)) {
    return;
  }

  // Get current attributes from of target in comparison with fragment.
  PatchJSONAttributes(fragment, targetNode);

  // Loop over children of giving node
  const kids = dom.nodeListToArray(targetNode.childNodes);
  const totalKids = kids.length;

  const fragmentKids = fragment.children.length;

  let i = 0;
  for (; i < totalKids; i++) {
    const childNode = kids[i];
    if (i >= fragmentKids) {
      const chnode = childNode as ChildNode;
      if (chnode) {
        chnode.remove();
      }
      continue;
    }

    const childFragment = fragment.children[i];
    PatchJSONNode(childFragment, childNode as Element, dictator, maker);
  }

  for (; i < fragmentKids; i++) {
    const tNode = maker.Make(document, fragment, false, true);
    targetNode.appendChild(tNode);
  }
  return;
}

/**
 * StreamJSONNodes provides a change stream like patching mechanism which uses a array of incoming JSONNode changes
 * which arrangements have no connected correlation to the actual dom but will be treated as individual change stream which
 * will be used to locate the specific elements referenced by using their JSONNode.ref for tracking and update.
 *
 * @param fragment
 * @param mount
 * @param dictator
 * @param maker
 * @constructor
 */
export function StreamJSONNodes(
  fragment: Array<JSONNode>,
  mount: Element,
  dictator: JSONDictator,
  maker: JSONMaker,
): void {
  const changes = fragment.filter(function (elem) {
    return !elem.removed;
  });

  // filter out all removals and ensure no removals match some
  // change or update before we actually make changes to the dom.
  fragment
    .filter(function (elem) {
      if (!elem.removed) {
        return false;
      }

      // if this removals has a directly related change, then filter it out also
      // from removals.
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

    // if giving targetNode is not found then this is possibly
    // a new node, create and append into parent.
    if (exts.Objects.isNullOrUndefined(targetNode)) {
      // Get parent of giving change at-least.
      const targetNodeParent = findElementParentbyRef(change.ref, mount)!;
      if (exts.Objects.isNullOrUndefined(targetNodeParent)) {
        console.log('Unable to apply new change stream: ', change);
        return;
      }

      // create new node and append into parent.
      const tNode = maker.Make(document, change, false, true);
      targetNodeParent.appendChild(tNode);
      return;
    }

    // // create new node and replace node in parent.
    // const tNode = maker.Make(document, change, false, true);
    // dom.replaceNode(targetNode!.parentNode!, targetNode!, tNode);
    ApplyStreamNode(change, targetNode!, dictator, maker);
  });
  return;
}

/**
 * ApplyStreamNode provides a DOM node patching approach for a giving JSON fragment and it's target DOM node.
 * It recursively applies changes from the JSONNode to the target, it does not optimize by skipping nodes
 * which have not changed and their children but runs through the changes in the JSONNode till end.
 *
 * Unlike PatchJSONNode it will not remove nodes left after it has gone through all JSONNode as it considers this
 * a stream of changes, the node must by nature be removed by the provision of a JSONNode that marks it as removed
 *
 * Generally this provides a very special case when you are dealing with change streams and not a your general patch.
 *
 * @param fragment is the JSONNode containing change for giving target.
 * @param targetNode is the target DOM node to apply changes to.
 * @param dictator is used to decide the same-ness and change-ness state of target node.
 * @param maker is used to create a new DOM node from a JSONNode.
 * @constructor
 */
export function ApplyStreamNode(
  fragment: JSONNode,
  targetNode: Element,
  dictator: JSONDictator,
  maker: JSONMaker,
): void {
  // if we are not dealing with exactly the same node as fragment as far as the.
  // then dictator is concerned, then we must do a swap after creating node.
  if (!dictator.Same(targetNode, fragment)) {
    const tNode = maker.Make(document, fragment, false, true);
    dom.replaceNode(targetNode.parentNode!, targetNode, tNode);
    return;
  }

  // if there there is a change, then apply attribute update.
  if (dictator.Changed(targetNode, fragment)) {
    PatchJSONAttributes(fragment, targetNode);
  }

  // Loop over children of giving node
  const totalKids = targetNode.childNodes.length;
  const fragmentKids = fragment.children.length;

  let i = 0;
  for (; i < totalKids; i++) {
    const childNode = targetNode.childNodes[i];
    if (i >= fragmentKids) {
      return;
    }

    const childFragment = fragment.children[i];
    PatchJSONNode(childFragment, childNode as Element, dictator, maker);
  }

  for (; i < fragmentKids; i++) {
    const tNode = maker.Make(document, fragment, false, true);
    targetNode.appendChild(tNode);
  }
  return;
}

// PatchTextCommentWithJSON runs the process of patching a corresponding text or comment node against
export function PatchTextCommentWithJSON(fragment: JSONNode, target: Node): void {
  if (fragment.type !== COMMENT_NODE && fragment.type !== TEXT_NODE) {
    return;
  }

  if (fragment.type !== COMMENT_NODE && fragment.type !== TEXT_NODE) {
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

// PatchJSONAttributes runs the attribute patching of a target dom element
// in comparison with a JSONNode.
export function PatchJSONAttributes(node: JSONNode, target: Element): void {
  const oldNodeAttrs: dom.Attributes = dom.recordAttributes(target!);

  node.attrs.forEach(function (attr) {
    const oldValue = oldNodeAttrs[attr.Key];

    // delete attribute from oldNodeAttr map.
    // Any attribute left in this map are ones
    // the new node has no use of anymore, so
    // we gotta remove it.
    delete oldNodeAttrs[attr.Key];

    // if the value is the same, skip it.
    if (attr.Value === oldValue) {
      return null;
    }

    // get attribute and apply to giving node.
    target.setAttribute(attr.Key, attr.Value);
  });

  // Remove all attributes left in this map.
  // New node did not update them, so we can remove them.
  for (let index in oldNodeAttrs) {
    target.removeAttribute(index);
  }

  // Patch attributes of id, tid, ref , atid and events.
  // target.setAttribute("id", node.id);
  target.setAttribute('_tid', node.tid);
  target.setAttribute('_ref', node.ref);
  target.setAttribute('_atid', node.atid);

  // Set the events.
  node.events.forEach(function events(event) {
    target.setAttribute('event-' + event.Name, event.Targets.join('|'));
  });

  // Patch object itself to have id, ref, tid and atid values.
  exts.Objects.PatchWith(target, '_id', node.id);
  exts.Objects.PatchWith(target, '_ref', node.ref);
  exts.Objects.PatchWith(target, '_tid', node.tid);
  exts.Objects.PatchWith(target, '_atid', node.atid);
}

/*
 * PatchDOMTree handles a general case algorithm with a simplified steps for handling
 * tree patching and updating for a giving live node with that of new changes loaded
 * into a document fragment or node.
 *
 * PatchDOMTree works with a whole DOM node or a DOM fragment, in that we mean a complete
 * set of DOM nodes for both the old and new set, which matches the children to be updated.
 * The idea is when using this, you get the giving static node where old nodes are mounted
 * and you use a document fragment containing nodes changes for the new-fragment on the
 * initial call.
 *
 * Consider using PartialPatchTree to use a more stream or partial friendly update mechanism
 * which allows updating sparse parts of the DOM trees.
 *
 * Note: This function assumes the oldNodeOrMount for the initial call is either empty or
 * contains previously rendered nodes only. That is there exists no static nodes within, as
 * anything not matching the update will be removed after reconciliation.
 *
 * @param newFragment: A DOMFragment or DOM node which will be used for replacement.
 * @param oldNodeOrMount: A static node or child node which holds the previous DOM tree.
 * @param dictator: Dictator helps us decide if the node is the same or changed.
 * @param isChildRecursion: A bool which tells the patch algorithm that it is comparing inner nodes of roots.
 * @return
 **/
export function PatchDOMTree(
  newFragment: Node,
  oldNodeOrMount: Node,
  dictator: NodeDictator,
  isChildRecursion: boolean | false,
) {
  if (isChildRecursion) {
    const rootNode = oldNodeOrMount.parentNode;

    // if we are dealing with de-similar nodes, then we swap giving nodes
    // instead of dealing with unmatched behaviours.
    if (!dictator.Same(oldNodeOrMount, newFragment)) {
      dom.replaceNode(rootNode!, oldNodeOrMount, newFragment);
      return null;
    }

    // if its not the same, then we are dealing with difference and must
    // reconcile.
    if (!oldNodeOrMount.hasChildNodes()) {
      dom.replaceNode(rootNode!, oldNodeOrMount, newFragment);
      return null;
    }
  }

  const newChildren = dom.nodeListToArray(newFragment.childNodes);
  const oldChildren = dom.nodeListToArray(oldNodeOrMount.childNodes);

  const oldChildrenLength = oldChildren.length;
  const newChildrenLength = newChildren.length;
  const removeOldLeft = newChildrenLength < oldChildrenLength;

  let lastIndex = 0;
  let lastNode: Node;
  let newChildNode: Node;
  let lastNodeNextSibling: Node | null = null;

  for (; lastIndex < oldChildrenLength; lastIndex++) {
    // if we've reached the index point where the new nodes
    // parent has stopped, then break.
    if (lastIndex >= newChildrenLength) {
      break;
    }

    lastNode = oldChildren[lastIndex];
    newChildNode = newChildren[lastIndex];
    lastNodeNextSibling = lastNode.nextSibling!;

    // Since they are the same type or node, we check content if text and verify
    // contents do not match, update and continue.
    if (
      (lastNode.nodeType === dom.TEXT_NODE || lastNode.nodeType === dom.COMMENT_NODE) &&
      newChildNode.nodeType === lastNode.nodeType
    ) {
      if (lastNode.textContent !== newChildNode.textContent) {
        lastNode.textContent = newChildNode.textContent;
      }
      continue;
    }

    // if giving nodes are not the same, then replace them,
    // possible case is we are dealing with different node
    // types entirely.
    if (!dictator.Same(lastNode, newChildNode)) {
      dom.replaceNode(oldNodeOrMount, lastNode, newChildNode);
      continue;
    }

    // if there was no actual change between nodes, then
    // skip this node and it's children as we see no reason
    // to go further down the rabbit hole.
    if (!dictator.Changed(lastNode, newChildNode)) {
      continue;
    }

    if (lastNode.nodeType !== newChildNode.nodeType) {
      dom.replaceNode(oldNodeOrMount, lastNode, newChildNode);
      continue;
    }

    // if this giving node has no children but the new one has, no need
    // bothering to do attribute checks, just replace node.
    if (!lastNode.hasChildNodes() && newChildNode.hasChildNodes()) {
      dom.replaceNode(oldNodeOrMount, lastNode, newChildNode);
      continue;
    }

    // if old node has kids but we do not have children, then replace
    // no-need for further checks.
    if (lastNode.hasChildNodes() && !newChildNode.hasChildNodes()) {
      dom.replaceNode(oldNodeOrMount, lastNode, newChildNode);
      continue;
    }

    const lastElement = lastNode as Element;
    const newElement = newChildNode as Element;

    // PatchDOMAttributes of giving element.
    PatchDOMAttributes(newElement, lastElement);

    // Add new attribute that says we've patched this already.
    lastElement.setAttribute('_patched', 'true');

    // Run Patching down the node set for giving element kids.
    PatchDOMTree(newElement, lastElement, dictator, true);

    // Remove patched attribute that says we've patched this already.
    lastElement.removeAttribute('_patched');
  }

  // if we are to remove all children nodes from the old parent here.
  if (removeOldLeft && lastNodeNextSibling !== null) {
    dom.removeFromNode(lastNodeNextSibling, null);
    return null;
  }

  // Since we got here, it means the the new fragment has more children
  // than the parent.
  for (; lastIndex < newChildrenLength; lastIndex++) {
    let newNode = newChildren[lastIndex];
    if (!exts.Objects.isNullOrUndefined(newNode)) {
      oldNodeOrMount.appendChild(newNode);
    }
  }
}

// PatchDOMAttributes runs the underline process to patch giving attributes
// of two elements.
export function PatchDOMAttributes(newElement: Element, oldElement: Element) {
  const oldNodeAttrs: dom.Attributes = dom.recordAttributes(oldElement!);

  for (let index in newElement.attributes) {
    const attr = newElement.attributes[index];
    const oldValue = oldNodeAttrs[attr.name];

    // delete attribute from oldNodeAttr map.
    // Any attribute left in this map are ones
    // the new node has no use of anymore, so
    // we gotta remove it.
    delete oldNodeAttrs[attr.name];

    // if the value is the same, skip it.
    if (attr.value === oldValue) {
      continue;
    }

    // get attribute and apply to giving node.
    oldElement.setAttribute(attr.name, attr.value);
  }

  // Remove all attributes left in this map.
  // New node did not update them, so we can remove them.
  for (let index in oldNodeAttrs) {
    oldElement.removeAttribute(index);
  }
}
