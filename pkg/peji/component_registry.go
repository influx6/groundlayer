package peji

// DOMRegistry implements the a component single for DOM objects.
type DOMRegistry struct {
	single map[string]DOM
	list   map[string][]DOM
}

func NewDOMRegistry() *DOMRegistry {
	return &DOMRegistry{
		single: map[string]DOM{},
		list:   map[string][]DOM{},
	}
}

func (k *DOMRegistry) Reset() {
	k.single = map[string]DOM{}
	k.list = map[string][]DOM{}
}

// HasType returns true/false if giving typeName exists.
func (k *DOMRegistry) HasList(refName string) bool {
	var _, ok = k.list[refName]
	return ok
}

// HasType returns true/false if giving typeName exists.
func (k *DOMRegistry) Has(refName string) bool {
	var _, ok = k.single[refName]
	return ok
}

// DeleteList deletes any key in the list map.
func (k *DOMRegistry) DeleteList(refName string) {
	delete(k.single, refName)
}

// Delete deletes any key in the singles map.
func (k *DOMRegistry) Delete(refName string) {
	delete(k.list, refName)
}

// GetList returns the DOM matching giving type routePath.
func (k *DOMRegistry) GetList(refName string) []DOM {
	return k.list[refName]
}

// Get returns the DOM matching giving type routePath.
func (k *DOMRegistry) Get(refName string) DOM {
	return k.single[refName]
}

// AddList adds giving component into underline register using the
// type routePath for giving component.
func (k *DOMRegistry) AddList(km []DOM, refName string) {
	k.list[refName] = km
}

// Add adds giving component into underline register using the
// type routePath for giving component.
func (k *DOMRegistry) Add(km DOM, refName string) {
	k.single[refName] = km
}
