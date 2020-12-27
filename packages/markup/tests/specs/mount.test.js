'use strict';
Object.defineProperty(exports, '__esModule', { value: true });
var mount = require('../../lib/mount');
var patch = require('../../lib/patch');
describe('HTML DOMMount', function () {
  it('Should be able to create a DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = "<div id='app'></div>";
    var app = document.querySelector('div#app');
    var mounted = new mount.DOMMount(document, '#app', function (event, target) {});
    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);
    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);
    var newContent = '\n\t\t\t<div id="rack-table">\n\t\t\t\t<div id="rack-item-1">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-2">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-3">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-4">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t</div>\n\t\t'.trim();
    mounted.patch(newContent);
    expect(mounted.innerHTML()).toBe(newContent);
  });
  it('Should be able to patch an existing DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = '\n\t\t<div id=\'app\'>\n\t\t\t\t\t<div id="rack-table">\n\t\t\t\t\t\t<div id="rack-item-1">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<div id="rack-item-2">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<div id="rack-item-3">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<div id="rack-item-4">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t</div>\n\t\t</div>'.trim();
    var app = document.querySelector('div#app');
    var mounted = new mount.DOMMount(document, '#app', function (event, target) {});
    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);
    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);
    var newContent2 = '\n\t\t\t<div id="rack-table">\n\t\t\t\t<div id="rack-item-2" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-3" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-1" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-4" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t</div>\n\t\t'.trim();
    mounted.patch(newContent2);
    expect(mounted.innerHTML()).toBe(newContent2);
  });
  it('Should be able to receive DOM events with a DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = "<div id='app'></div>";
    var app = document.querySelector('div#app');
    var eventReceived = 0;
    var mounted = new mount.DOMMount(document, '#app', function (event, target) {
      eventReceived++;
    });
    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);
    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);
    var newContent = '\n\t\t\t<div id="rack-table">\n\t\t\t\t<div id="rack-item-2" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-3" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-1" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-4" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t</div>\n\t\t'.trim();
    mounted.patch(newContent);
    expect(mounted.innerHTML()).toBe(newContent);
    var rackItem2 = mounted.mountNode.querySelector('div#rack-item-2');
    var clickEvent = new Event('rackup', { bubbles: true });
    rackItem2.dispatchEvent(clickEvent);
    expect(eventReceived).toBe(1);
  });
});
describe('JSON DOMMount Fragment', function () {
  it('Should be able to patch an existing DOMMount', function () {
    // Set up our document body
    document.body.innerHTML =
      '\n\t\t<div id=\'app\'>\n\t\t\t\t\t<div id="rack-table">\n\t\t\t\t\t\t<div id="rack-item-1">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<div id="rack-item-2">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<div id="rack-item-3">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<div id="rack-item-4">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t</div>\n\t\t</div>';
    var app = document.querySelector('div#app');
    var mounted = new mount.DOMMount(document, '#app', function (event, target) {});
    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);
    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);
    var renderedContent =
      '\n\t\t\t\t<div id="rack-item-2" events="rackup-00" _ref="/rack-table/rack-item2">\n\t\t\t\t\t\t<span>Wrecker Walk away</span>\n\t\t\t</div>\n\t\t';
    var patchContent = patch.ToJSONNode(renderedContent, false, null);
    mounted.streamList(patchContent);
    expect(mounted.innerHTML()).not.toBe('');
    expect(mounted.mountNode.querySelector('div#rack-item-2')).toBeDefined();
    expect(mounted.mountNode.querySelector('div#rack-item-2').getAttribute('events')).toBe('rackup-00');
    var spanWrack = mounted.mountNode.querySelector('div#rack-item-2').querySelector('span');
    expect(spanWrack).toBeDefined();
    expect(spanWrack.innerHTML).toBe('Wrecker Walk away');
    expect(spanWrack.textContent).toBe('Wrecker Walk away');
  });
});
describe('JSON DOMMount', function () {
  it('Should be able to create a DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = "<div id='app'></div>";
    var app = document.querySelector('div#app');
    var mounted = new mount.DOMMount(document, '#app', function (event, target) {});
    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);
    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);
    var renderedContent =
      '\n\t\t\t<div id="rack-table">\n\t\t\t\t<div id="rack-item-1">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-2">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-3">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-4">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t</div>\n\t\t';
    var patchContent = patch.ToJSONNode(renderedContent, false, null);
    mounted.patchList(patchContent);
    expect(mounted.innerHTML()).not.toBe('');
    expect(mounted.mountNode.querySelector('div#rack-item-1')).toBeDefined();
  });
  it('Should be able to patch an existing DOMMount', function () {
    // Set up our document body
    document.body.innerHTML =
      '\n\t\t<div id=\'app\'>\n\t\t\t\t\t<div id="rack-table">\n\t\t\t\t\t\t<div id="rack-item-1">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<div id="rack-item-2">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<div id="rack-item-3">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<div id="rack-item-4">\n\t\t\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t</div>\n\t\t</div>';
    var app = document.querySelector('div#app');
    var mounted = new mount.DOMMount(document, '#app', function (event, target) {});
    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);
    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);
    var renderedContent =
      '\n\t\t\t<div id="rack-table">\n\t\t\t\t<div id="rack-item-2" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-3" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-1" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-4" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t</div>\n\t\t';
    var patchContent = patch.ToJSONNode(renderedContent, false, null);
    mounted.patchList(patchContent);
    expect(mounted.innerHTML()).not.toBe('');
    expect(mounted.mountNode.querySelector('div#rack-item-1')).toBeDefined();
    expect(mounted.mountNode.querySelector('div#rack-item-1').getAttribute('events')).toBe('rackup-00');
  });
  it('Should be able to receive DOM events with a DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = "<div id='app'></div>";
    var app = document.querySelector('div#app');
    var eventReceived = 0;
    var mounted = new mount.DOMMount(document, '#app', function (event, target) {
      eventReceived++;
    });
    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);
    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);
    var renderedContent =
      '\n\t\t\t<div id="rack-table">\n\t\t\t\t<div id="rack-item-2" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-3" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-1" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t\t<div id="rack-item-4" events="rackup-00">\n\t\t\t\t\t\t<span>Wrecker</span>\n\t\t\t\t</div>\n\t\t\t</div>\n\t\t';
    var patchContent = patch.ToJSONNode(renderedContent, false, null);
    mounted.patchList(patchContent);
    expect(mounted.innerHTML()).not.toBe('');
    expect(mounted.mountNode.querySelector('div#rack-item-1')).toBeDefined();
    expect(mounted.mountNode.querySelector('div#rack-item-1').getAttribute('events')).toBe('rackup-00');
    var rackItem2 = mounted.mountNode.querySelector('div#rack-item-2');
    var clickEvent = new Event('rackup', { bubbles: true });
    rackItem2.dispatchEvent(clickEvent);
    expect(eventReceived).toBe(1);
  });
});
