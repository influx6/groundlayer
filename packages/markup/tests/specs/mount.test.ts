import * as mount from '../../lib/mount';
import * as patch from '../../lib/patch';

describe('HTML DOMMount', function () {
  it('Should be able to create a DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = "<div id='app'></div>";

    const app = document.querySelector('div#app');

    const mounted = new mount.DOMMount(document, '#app', (event: Event, targets: Array<string>, target: Element) => {});

    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);

    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);

    const newContent = `
			<div id="rack-table">
				<div id="rack-item-1">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-2">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-3">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-4">
						<span>Wrecker</span>
				</div>
			</div>
		`.trim();

    mounted.patch(newContent);
    expect(mounted.innerHTML()).toBe(newContent);
  });

  it('Should be able to patch an existing DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = `
		<div id='app'>
					<div id="rack-table">
						<div id="rack-item-1">
								<span>Wrecker</span>
						</div>
						<div id="rack-item-2">
								<span>Wrecker</span>
						</div>
						<div id="rack-item-3">
								<span>Wrecker</span>
						</div>
						<div id="rack-item-4">
								<span>Wrecker</span>
						</div>
					</div>
		</div>`.trim();

    const app = document.querySelector('div#app');

    const mounted = new mount.DOMMount(document, '#app', (event: Event, targets: Array<string>, target: Element) => {});

    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);

    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);

    const newContent2 = `
			<div id="rack-table">
				<div id="rack-item-2" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-3" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-1" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-4" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
			</div>
		`.trim();

    mounted.patch(newContent2);
    expect(mounted.innerHTML()).toBe(newContent2);
  });

  it('Should be able to receive DOM events with a DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = "<div id='app'></div>";

    const app = document.querySelector('div#app');

    let eventReceived = 0;
    const mounted = new mount.DOMMount(document, '#app', (event: Event, targets: Array<string>, target: Element) => {
      eventReceived++;
    });

    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);

    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);

    const newContent = `
			<div id="rack-table">
				<div id="rack-item-2" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-3" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-1" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-4" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
			</div>
		`.trim();

    mounted.patch(newContent);
    expect(mounted.innerHTML()).toBe(newContent);

    const rackItem2 = mounted.mountNode.querySelector('div#rack-item-2') as Element;
    const clickEvent = new Event('rackup', { bubbles: true });
    rackItem2.dispatchEvent(clickEvent);
    expect(eventReceived).toBe(1);
  });
});

describe('JSON DOMMount Fragment', function () {
  it('Should be able to patch an existing DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = `
		<div id='app'>
					<div id="rack-table">
						<div id="rack-item-1">
								<span>Wrecker</span>
						</div>
						<div id="rack-item-2">
								<span>Wrecker</span>
						</div>
						<div id="rack-item-3">
								<span>Wrecker</span>
						</div>
						<div id="rack-item-4">
								<span>Wrecker</span>
						</div>
					</div>
		</div>`;

    const app = document.querySelector('div#app');

    const mounted = new mount.DOMMount(document, '#app', (event: Event, targets: Array<string>, target: Element) => {});

    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);

    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);

    const renderedContent = `
				<div id="rack-item-2" events-rackup="do_something" _ref="/rack-table/rack-item2">
						<span>Wrecker Walk away</span>
			</div>
		`;

    const patchContent = patch.ToJSONNode(renderedContent, false, null);
    mounted.streamList(patchContent);

    expect(mounted.innerHTML()).not.toBe('');
    expect(mounted.mountNode.querySelector('div#rack-item-2')).toBeDefined();
    expect(mounted.mountNode.querySelector('div#rack-item-2')!.getAttribute('events')).toBe('rackup-00');

    const spanWrack = mounted.mountNode.querySelector('div#rack-item-2')!.querySelector('span')!;
    expect(spanWrack).toBeDefined();
    expect(spanWrack.innerHTML).toBe('Wrecker Walk away');
    expect(spanWrack.textContent).toBe('Wrecker Walk away');
  });
});

describe('JSON DOMMount', function () {
  it('Should be able to create a DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = "<div id='app'></div>";

    const app = document.querySelector('div#app');

    const mounted = new mount.DOMMount(document, '#app', (event: Event, targets: Array<string>, target: Element) => {});

    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);

    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);

    const renderedContent = `
			<div id="rack-table">
				<div id="rack-item-1">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-2">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-3">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-4">
						<span>Wrecker</span>
				</div>
			</div>
		`;

    const patchContent = patch.ToJSONNode(renderedContent, false, null);
    mounted.patchList(patchContent);
    expect(mounted.innerHTML()).not.toBe('');
    expect(mounted.mountNode.querySelector('div#rack-item-1')).toBeDefined();
  });

  it('Should be able to patch an existing DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = `
		<div id='app'>
					<div id="rack-table">
						<div id="rack-item-1">
								<span>Wrecker</span>
						</div>
						<div id="rack-item-2">
								<span>Wrecker</span>
						</div>
						<div id="rack-item-3">
								<span>Wrecker</span>
						</div>
						<div id="rack-item-4">
								<span>Wrecker</span>
						</div>
					</div>
		</div>`;

    const app = document.querySelector('div#app');

    const mounted = new mount.DOMMount(document, '#app', (event: Event, targets: Array<string>, target: Element) => {});

    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);

    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);

    const renderedContent = `
			<div id="rack-table">
				<div id="rack-item-2" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-3" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-1" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-4" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
			</div>
		`;

    const patchContent = patch.ToJSONNode(renderedContent, false, null);
    mounted.patchList(patchContent);
    expect(mounted.innerHTML()).not.toBe('');
    expect(mounted.mountNode.querySelector('div#rack-item-1')).toBeDefined();
    expect(mounted.mountNode.querySelector('div#rack-item-1')!.getAttribute('events')).toBe('rackup-00');
  });

  it('Should be able to receive DOM events with a DOMMount', function () {
    // Set up our document body
    document.body.innerHTML = "<div id='app'></div>";

    const app = document.querySelector('div#app');

    let eventReceived = 0;
    const mounted = new mount.DOMMount(document, '#app', (event: Event, targets: Array<string>, target: Element) => {
      eventReceived++;
    });

    expect(mounted).toBeDefined();
    expect(mounted.doc).toBeDefined();
    expect(mounted.doc).toBe(document);

    expect(mounted.mountNode).toBeDefined();
    expect(mounted.mountNode).toBe(app);

    const renderedContent = `
			<div id="rack-table">
				<div id="rack-item-2" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-3" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-1" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
				<div id="rack-item-4" events-rackup="do_something">
						<span>Wrecker</span>
				</div>
			</div>
		`;

    const patchContent = patch.ToJSONNode(renderedContent, false, null);
    mounted.patchList(patchContent);
    expect(mounted.innerHTML()).not.toBe('');
    expect(mounted.mountNode.querySelector('div#rack-item-1')).toBeDefined();
    expect(mounted.mountNode.querySelector('div#rack-item-1')!.getAttribute('events')).toBe('rackup-00');

    const rackItem2 = mounted.mountNode.querySelector('div#rack-item-2') as Element;
    const clickEvent = new Event('rackup', { bubbles: true });
    rackItem2.dispatchEvent(clickEvent);
    expect(eventReceived).toBe(1);
  });
});
