import * as dom from './dom';
import * as patch from './patch';
import * as utils from './utils';

/**
 * DOMChange defines a type which represent a string of html, a DocumentFragment containing
 *
 */
export type DOMChange = string | DocumentFragment | patch.JSONNode;

/**
 * EventHandler defines a function interface type for handling events.
 */
export interface EventHandler {
  (e: Event): void;
}

/**
 * EventNotifier defines a function interface type for sending notification
 * for the occurrence of an event for a giving node.
 *
 * Note the element and Node are the same object, but if giving node is not
 * an Element then element would be null.
 */
export interface EventNotifier {
  (e: Event, triggers: Array<string>, elem: Element): void;
}

/* DOMMount exists to provide a focus DOM operation on a giving underline static node,
 *  which will be used for mounting an ever updating series of changes, nodes and html elements.
 * It acts as the bridge for event management, propagation and update, just like in react, the mount
 * node will be where your components are rendered.
 *
 * Static node means a DOM node never to be removed, changed by other scripts, it was added by
 * html text within original html file)
 *
 */
export class DOMMount {
  public readonly doc: Document;
  public readonly mountNode: Element;
  private readonly events: utils.KeyMap;
  private readonly handler: EventHandler;
  private readonly notifier: EventNotifier;

  constructor(document: Document, target: string | Element, notifier: EventNotifier) {
    if (typeof document !== 'object') {
      throw new Error('document should be an object');
    }

    this.doc = document;
    this.notifier = notifier;
    this.events = {} as utils.KeyMap;
    this.handler = this.handleEvent.bind(this);

    // if it's a string, then attempt using of document.querySelector
    if (typeof target === 'string') {
      const targetSelector = target as string;
      const node: Element = this.doc.querySelector(targetSelector)!;
      if (node === null || node === undefined) {
        throw new Error(`unable to locate node for ${{ targetSelector }}`);
      }

      this.mountNode = node;
    } else {
      this.mountNode = target as Element;
    }
  }

  /**
   * handleEvent handles the case of live events for the giving mount node.
   * Events which reach this method would not continue propagation but are
   * meant to be handled by this node if registered.
   *
   * If a event is not being handled or registered then it is allowed to propagate.
   *
   * @param event
   */
  handleEvent(event: Event): void {
    if (!this.events[event.type]) {
      return;
    }

    event.stopPropagation();

    const target = event.target as Node;
    if (target.nodeType !== dom.ELEMENT_NODE) {
      return;
    }

    const targetElement = (target as Element)!;
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

  /**
   * applyPatch applies a DOM Change which either is a string of html,
   * DocumentFragment containing changes or a JSON Node which represent
   * a complete DOM which will be applied to the static mount directly using
   * the dom.PatchTree. This means we expect the DOM described here to be
   * a complete change reflection for the nodes in the mount.
   *
   * It entirely relies on dom.PatchDOMTree for a DOM node and dom.PatchJSONNodeTree
   * for a JSONNode. Strings are considered html text.
   *
   * @param change the node change to be applied.
   */
  patch(change: DOMChange): void {
    this.patchWith(change, patch.DefaultNodeDictator, patch.DefaultJSONDictator, patch.DefaultJSONMaker);
  }

  patchWith(
    change: DOMChange,
    nodeDictator: patch.NodeDictator,
    jsonDictator: patch.JSONDictator,
    jsonMaker: patch.JSONMaker,
  ): void {
    if (change instanceof DocumentFragment) {
      const fragment = change as DocumentFragment;
      this.registerNodeEvents(fragment);
      patch.PatchDOMTree(fragment, this.mountNode, nodeDictator, false);
      return;
    }

    if (typeof change === 'string') {
      const node = document.createElement('div');
      node.innerHTML = (change as string).trim();

      this.registerNodeEvents(node);
      patch.PatchDOMTree(node, this.mountNode, nodeDictator, false);
      return;
    }

    if (!patch.isJSONNode(change)) {
      return;
    }

    const node = change as patch.JSONNode;
    this.registerJSONNodeEvents(node);
    patch.PatchJSONNodeTree(node, this.mountNode, jsonDictator, jsonMaker);
  }

  /**
   * patchList applies a array of DOM changes to be applied to the mount DOM.
   * It individually calls DOMMount.patch for each change.
   *
   * @param changes is a array of DOM change.
   */
  patchList(changes: Array<DOMChange>): void {
    changes.forEach(this.patch.bind(this));
  }

  /**
   * stream applies a string which should contain a array of JSONNode which will be
   * applied to the DOMMount node using DOMMount.streamList.
   *
   * @param changes a JSON string containing a list of JSONNode items.
   */
  stream(changes: string): void {
    const nodes = JSON.parse(changes) as Array<patch.JSONNode>;
    return this.streamList(nodes);
  }

  /***
   * streamList applies a list of JSONNode changes considered as partial changes
   * with no relation to their order, each is applied independently to the DOM,
   * where each change only applies to it's target. It is unlike the DOMMount.patch
   * method which expects a complete DOM representation as previously rendered with
   * the new changes.
   *
   * This can be used to effectively remove part of the nodes which have expired or
   * add or swap nodes within specific areas of the DOM.
   *
   * It relies on the dom.StreamJSONNodes.
   *
   * @param changes is a array of JSONNodes to apply to the dom in static mount node.
   */
  streamList(changes: Array<patch.JSONNode>): void {
    this.streamListWith(changes, patch.DefaultJSONDictator, patch.DefaultJSONMaker);
  }

  streamListWith(changes: Array<patch.JSONNode>, dictator: patch.JSONDictator, maker: patch.JSONMaker): void {
    changes.forEach(this.registerJSONNodeEvents.bind(this));
    patch.StreamJSONNodes(changes, this.mountNode, patch.DefaultJSONDictator, patch.DefaultJSONMaker);
  }

  registerNodeEvents(n: Node): void {
    const binder = this;
    dom.applyEachNode(n, function (node: Node): void {
      if (node.nodeType !== dom.ELEMENT_NODE) {
        return;
      }

      const elem = node as Element;
      const events = elem.getAttribute('events')!;
      if (events) {
        events.split(' ').forEach(function (desc) {
          const eventName = desc.substr(0, desc.length - 3);
          binder.registerEvent(eventName);

          // apply giving modifiers to ensure consistent behaviour for
          // prevent default.
          switch (desc.substr(desc.length - 2, desc.length)) {
            case '01':
              // we need propagation for live events to work.
              // node.addEventListener(eventName, MountNode.stopPropagation, false);
              break;
            case '10':
              n.addEventListener(eventName, DOMMount.preventDefault, false);
              break;
            case '11':
              n.addEventListener(eventName, DOMMount.preventDefault, false);

              // we need propagation for live events to work.
              // node.addEventListener(eventName, MountNode.stopPropagation, false);
              break;
          }
        });
      }
    });
  }

  registerJSONNodeEvents(node: patch.JSONNode): void {
    const binder = this;
    patch.applyJSONNodeFunction(node, function (n: patch.JSONNode): void {
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

  textContent(): string {
    return this.mountNode.textContent!;
  }

  innerHTML(): string {
    return this.mountNode.innerHTML.trim();
  }

  Html(): string {
    return dom.toHTML(this.mountNode, false);
  }

  registerEvent(eventName: string): void {
    if (this.events[eventName]) {
      return;
    }

    this.mountNode!.addEventListener(eventName, this.handler, true);
    this.events[eventName] = true;
  }

  unregisterEvent(eventName: string): void {
    if (!this.events[eventName]) {
      return;
    }
    this.mountNode!.removeEventListener(eventName, this.handler, true);
    this.events[eventName] = false;
  }

  static preventDefault(event): void {
    event.preventDefault();
  }

  static stopPropagation(event): void {
    event.stopPropagation();
  }
}
