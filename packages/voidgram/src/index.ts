import * as realm from '../../realm/src';
import * as markup from '../../markup/src';
import * as animations from '../../animation/src';
import * as http from '../../http/src';
import * as promises from '../../promises/src';

const voidgram = {
    realm: {},
    animations: {},
    http: {},
    promises: {},
    markup: {},
}

// mount all to object
realm.mountTo(voidgram)
markup.mountTo(voidgram)
animations.mountTo(voidgram)
http.mountTo(voidgram)
promises.mountTo(voidgram)

if (window) {
    //@ts-ignore
    window.voidgram = voidgram;
}

export default voidgram
