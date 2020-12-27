const namespace = {};

export function mountTo(parent: object & { realm: object }) {
    parent.realm = namespace;
}

export default namespace;
