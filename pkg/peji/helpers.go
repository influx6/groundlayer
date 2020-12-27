package peji

import "github.com/influx6/groundlayer/pkg/domu"

// Section composes a set of components wrapped into a Section component.
func Section(components ...DOM) DOMFunc {
	return func(data Data) *domu.Node {
		var parent = domu.HTMLSection()
		for _, component := range components {
			component.Render(data).Mount(parent)
		}
		return parent
	}
}
