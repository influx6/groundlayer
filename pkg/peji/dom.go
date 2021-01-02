package peji

import (
	"fmt"

	"github.com/influx6/sabuhp"

	"github.com/influx6/groundlayer/pkg/domu"
)

const (
	PlainHTML      = "text/html"
	VoidHTML       = "application/groundlayer+html"
	VoidHTMLDiff   = "application/groundlayer+html+diff"
	VoidJSON       = "application/groundlayer+json"
	VoidJSONStream = "application/groundlayer+json_stream"
)

type Data struct {
	Path      string
	SessionId string
	Data      interface{}
	Params    Params
	Message   *sabuhp.Message
}

type DefaultNotFound struct{}

func (d DefaultNotFound) Handle(dt Data) *domu.Node {
	return domu.Text(fmt.Sprintf("%s not found!", dt.Path))
}

type DOM interface {
	Render(ctx Data) *domu.Node
}

var _ DOM = (*DOMFunc)(nil)

type DOMFunc func(ctx Data) *domu.Node

func (pf DOMFunc) Render(ctx Data) *domu.Node {
	return pf(ctx)
}

func Portal(route string, name string, renders ...domu.Mounter) *domu.Node {
	var elem = domu.Element("portal", renders...)
	domu.NewStringAttr("route", route).Mount(elem)
	domu.NewStringAttr("name", name).Mount(elem)
	return elem
}

// LiveDOMFrom returns a DOM wrapped to return Nodes reconciled with their previous render.
func LiveDOMFrom(dom DOM, route string, name string) *LiveDOM {
	return &LiveDOM{dom: dom, route: route, name: name}
}

// LiveFromDOMList returns a DOM wrapped to return Nodes reconciled with their previous render.
func LiveFromDOMList(dom []DOM, route string, name string) *LiveDOM {
	return &LiveDOM{dom: DOMSet(dom), route: route, name: name}
}

type LiveDOMList []LiveDOM

// LiveDOM wraps a dom which keeps track of Last render and reconciling new renderings with the old.
type LiveDOM struct {
	dom   DOM
	name  string
	route string
}

// Render implements the DOM interface which returns
// a new instance of the DOM's rendered reconciled with
// the Last.
func (c *LiveDOM) Render(ctx Data) *domu.Node {
	return Portal(c.route, c.name, c.dom.Render(ctx))
}

// DOMSet defines a list of DOM items.
type DOMSet []DOM

// Render implements the DOM interface for a DOM set (a list of DOM objects)
func (dms DOMSet) Render(data Data) *domu.Node {
	var carrier = domu.Carrier()
	for _, item := range dms {
		item.Render(data).Mount(carrier)
	}
	return carrier
}

// Render implements the DOM interface for a DOM set (a list of DOM objects)
func (dms DOMSet) RenderTo(data Data, parent *domu.Node) {
	for _, item := range dms {
		item.Render(data).Mount(parent)
	}
}

type Layout interface {
	Render(p *Page, ctx Data, parent *domu.Node)
}

var _ Layout = (*LayoutFunc)(nil)

type LayoutFunc func(p *Page, ctx Data, parent *domu.Node)

func (pf LayoutFunc) Render(p *Page, ctx Data, parent *domu.Node) {
	pf(p, ctx, parent)
}
