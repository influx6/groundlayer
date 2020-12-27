import * as realm from '../../realm/src';
import * as markup from '../../markup/src';
import * as animations from '../../animation/src';
import * as http from '../../http/src';
import * as promises from '../../promises/src';

const groundlayergram = {
    realm: {},
    animations: {},
    http: {},
    promises: {},
    markup: {},
}

// mount all to object
realm.mountTo(groundlayergram)
markup.mountTo(groundlayergram)
animations.mountTo(groundlayergram)
http.mountTo(groundlayergram)
promises.mountTo(groundlayergram)

if (window) {
    //@ts-ignore
    window.groundlayergram = groundlayergram;
}

export default groundlayergram
