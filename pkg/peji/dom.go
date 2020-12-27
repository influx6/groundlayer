package peji

import (
	"fmt"
	"time"

	"github.com/influx6/sabuhp"

	"github.com/influx6/groundlayer/pkg/domu"
)

const (
	PlainHTML      = "text/html"
	VoidHTML       = "application/void+html"
	VoidHTMLDiff   = "application/void+html+diff"
	VoidJSON       = "application/void+json"
	VoidJSONStream = "application/void+json_stream"
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

// NewLiveDOM returns a DOM wrapped to return Nodes reconciled with their previous render.
func NewLiveDOM(dom DOM) *LiveDOM {
	if cdm, ok := dom.(*LiveDOM); ok {
		return cdm
	}
	return &LiveDOM{dom: dom}
}

type LiveDOMList []LiveDOM

// LiveDOM wraps a dom which keeps track of Last render and reconciling new renderings with the old.
type LiveDOM struct {
	dom  DOM
	last *domu.Node
	tick time.Time
}

// LastCalled returns the time when the Last called was made.
func (c *LiveDOM) LastCalled() time.Time {
	return c.tick
}

// Render implements the DOM interface which returns
// a new instance of the DOM's rendered reconciled with
// the Last.
func (c *LiveDOM) Render(ctx Data) *domu.Node {
	c.tick = time.Now()

	var prev = c.last
	var res = c.dom.Render(ctx)
	c.last = res.Clone(true)

	if prev != nil {
		res.Reconcile(prev)
	}
	return res
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

func DOMToLiveDOM(list []DOM) []DOM {
	var set = make([]DOM, len(list))
	for _, dm := range list {
		if cdm, ok := dm.(*LiveDOM); ok {
			set = append(set, cdm)
			continue
		}
		set = append(set, &LiveDOM{dom: dm})
	}
	return set
}

func LiveDOMToDOM(list []*LiveDOM) []DOM {
	var set = make([]DOM, len(list))
	for _, dm := range list {
		set = append(set, dm)
	}
	return set
}
