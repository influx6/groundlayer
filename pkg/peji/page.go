package peji

import (
	"fmt"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/influx6/npkg/nstorage/plain"

	"github.com/influx6/groundlayer/pkg/domu"
)

var (
	multipleSlashes = regexp.MustCompile(`/{2,}`)
	allSlashes      = regexp.MustCompile(`^/$`)
)

// LiveDOMHandler returns a Handler which provides the params as the data to be giving to
// a Live DOM render method.
func LiveDOMHandler(target string, page *Page) Handler {
	return HandlerFunc(func(d Data) *domu.Node {
		return page.Live(target).Render(d)
	})
}

// StaticDOMHandler returns a Handler which provides the params as the data to be giving to
// a Live DOM render method.
func StaticDOMHandler(target string, page *Page) Handler {
	return HandlerFunc(func(d Data) *domu.Node {
		return page.Static(target).Render(d)
	})
}

// Page implements the Document interface and represents
// a single unique html page.
type Page struct {
	UseChanges       bool
	Name             string
	RoutePath        string
	AnyRouteInPage   string
	Router           *Mux
	Layout           Layout
	Last             *domu.Node
	Registry         *DOMRegistry
	LiveRegistry     *DOMRegistry
	LiveRouteMapping *plain.StringMap
}

func WithPage(name string, layout Layout, notFound Handler) *Page {
	return NewPageUsingRouter(handlePath(name), layout, NewMux(MuxConfig{
		RootPath: name,
		NotFound: notFound,
	}))
}

func NewPageUsingRouter(routePath string, layout Layout, router *Mux) *Page {
	var p Page
	p.Name = routePath
	p.Layout = layout
	p.Router = router
	p.RoutePath = handlePath(routePath)
	p.AnyRouteInPage = handlePath(routePath, "*path")

	fmt.Printf("NewPage: %q -> %q -- %q \n", routePath, p.RoutePath, p.AnyRouteInPage)

	var pPointer = &p
	var handler = HandlerFunc(func(d Data) *domu.Node {
		return pPointer.Render(d)
	})

	p.Router.Serve(p.RoutePath, handler)
	p.Router.Serve(p.AnyRouteInPage, handler)
	return &p
}

func (pg *Page) Close() {
	pg.LiveRegistry.Reset()
	pg.Registry.Reset()
}

func (pg *Page) check() {
	if pg.Registry == nil {
		pg.Registry = NewDOMRegistry()
	}
	if pg.LiveRegistry == nil {
		pg.LiveRegistry = NewDOMRegistry()
	}
	if pg.LiveRouteMapping == nil {
		pg.LiveRouteMapping = plain.NewStringMap(10)
	}
}

func (pg *Page) SetReconciliation(b bool) {
	pg.UseChanges = b
}

// Cleanup goes through the live instance releasing
// any which may have surpassed the maxLastUsed period.
func (pg *Page) Cleanup(maxDuration time.Duration) {
	pg.LiveRegistry.Clean(maxDuration)
}

// HasLive returns true if giving route exists either as a live dom or dom list.
func (pg *Page) HasLive(route string) bool {
	pg.check()
	return pg.LiveRegistry.Has(route) || pg.LiveRegistry.HasList(route)
}

// HasStatic returns true if giving route exists either as a static dom or dom list.
func (pg *Page) HasStatic(name string) bool {
	pg.check()
	return pg.Registry.Has(name) || pg.Registry.HasList(name)
}

func (pg *Page) Serve(d Data) *domu.Node {
	return pg.Router.ServeRoute(d)
}

func (pg *Page) ServeStatic(name string, handler Handler) {
	name = cleanName(name)
	if !pg.HasStatic(name) {
		panic(fmt.Sprintf("static dom %q does not exists", name))
	}

	var targetRoute = handlePath(pg.Name, name)
	var targetRouteWithWildcard = handlePath(pg.Name, name, "*path")

	pg.Router.Serve(targetRoute, handler)
	pg.Router.Serve(targetRouteWithWildcard, handler)
}

func (pg *Page) ServeLive(route string, handler Handler) {
	if !pg.HasLive(route) {
		panic(fmt.Sprintf("route %q does not exists for any registered dom", route))
	}

	var targetRoute = handlePath(path.Join(pg.Name, route))
	var targetRouteWithWildcard = handlePath(pg.Name, route, "*path")

	pg.Router.Serve(targetRoute, handler)
	pg.Router.Serve(targetRouteWithWildcard, handler)
}

// GetRouting returns the route for giving name.
func (pg *Page) GetRouting(name string) string {
	pg.check()
	return pg.LiveRouteMapping.Get(name)
}

// GetStatic returns the DOM for giving name, if not found, nil is returned.
func (pg *Page) Static(name string) DOM {
	pg.check()

	if pg.Registry.HasList(name) {
		return DOMSet(pg.Registry.GetList(name))
	}
	return pg.Registry.Get(name)
}

// GetLive returns the DOM for giving name, if not found, nil is returned.
func (pg *Page) Live(route string) DOM {
	pg.check()

	if !pg.HasLive(route) {
		return nil
	}

	if pg.LiveRegistry.Has(route) {
		return DOMSet(pg.LiveRegistry.GetList(route))
	}
	return pg.LiveRegistry.Get(route)
}

// Retire retires the dom (live or static) associated with the name
// and every route associated with it.
func (pg *Page) Retire(name string) {
	pg.check()

	if !pg.HasStatic(name) {
		return
	}

	pg.Registry.Delete(name)
	pg.Registry.DeleteList(name)

	if pg.LiveRouteMapping.Has(name) {
		var mapping = pg.LiveRouteMapping.Get(name)
		pg.LiveRegistry.Delete(mapping)
		pg.LiveRegistry.DeleteList(mapping)
		pg.LiveRouteMapping.SetMany(func(items map[string]string) {
			delete(items, mapping)
		})
	}
}

// AddStatic adds a new name mapping for a provided list, ensuring these
// are referenced with giving component, if there is more than 1 item then they
// are treated as a whole and referenced by provided name.
//
// Note: It clears all live routes for this giving name, if we are replacing
// an existing live key.
func (pg *Page) AddStatic(name string, items ...DOM) {
	pg.check()

	var itemCount = len(items)
	if itemCount == 0 {
		panic("DOM set must not be a count of zero")
	}

	if pg.HasStatic(name) {
		panic("names must be unique across list and single component types")
	}

	if itemCount == 1 {
		pg.Registry.Add(items[0], name)
	} else {
		pg.Registry.AddList(items, name)
	}

	var set = DOMSet(items)
	pg.ServeStatic(name, HandlerFunc(func(data Data) *domu.Node {
		return set.Render(data)
	}))

	if pg.LiveRouteMapping.Has(name) {
		var mapping = pg.LiveRouteMapping.Get(name)
		pg.LiveRegistry.Delete(mapping)
		pg.LiveRegistry.DeleteList(mapping)
		pg.LiveRouteMapping.SetMany(func(items map[string]string) {
			delete(items, mapping)
		})
	}
}

// AddLive returns a new live DOM by creating a route mapping between
// a target component or list of components in the page to be serviceable
// by the route ensuring it's always reconciled with it's last state.
func (pg *Page) AddLive(route string, targetName string) DOM {
	pg.check()

	if strings.Trim(route, "/") == targetName {
		panic("route can not be the same as targetName")
	}

	if pg.LiveRouteMapping.Has(targetName) {
		panic(fmt.Sprintf("target %q already has a route %q", targetName, pg.LiveRouteMapping.Get(targetName)))
	}

	// if we do not have such a registry, then ignore
	if !pg.HasStatic(targetName) {
		panic(fmt.Sprintf("no component with target name %s", targetName))
	}

	if pg.Registry.Has(targetName) {
		var target = pg.Registry.Get(targetName)
		var liveTarget = NewLiveDOM(target)
		pg.LiveRouteMapping.Set(targetName, route)
		pg.LiveRegistry.Add(liveTarget, route)

		pg.ServeLive(route, HandlerFunc(func(data Data) *domu.Node {
			return liveTarget.Render(data)
		}))
		return liveTarget
	}

	var instances = pg.Registry.GetList(targetName)
	var liveInstances = DOMToLiveDOM(instances)
	pg.LiveRouteMapping.Set(targetName, route)
	pg.LiveRegistry.AddList(liveInstances, route)

	var set = DOMSet(liveInstances)
	pg.ServeLive(route, HandlerFunc(func(data Data) *domu.Node {
		return set.Render(data)
	}))
	return set
}

// Render renders new domu.Node from existing page which is reconciled
// with it's previous render.
//
// The markup node is pure, which means it's the complete picture
// generated from a Page object.
func (pg *Page) Render(d Data) *domu.Node {
	if !pg.UseChanges {
		return pg.render(d)
	}

	var lastRender = pg.Last
	var newRender = pg.render(d)
	pg.Last = newRender.Clone(true)
	_ = newRender.Reconcile(lastRender)
	return newRender
}

// Render returns a new domu.Node with Document node type.
//
// If the IsHtml flag is set on then we expect the page
// root element to be a maximum of 2, where the first is
// a DocType node and the other a HTML node.
// If during IsHtml on flag, and only one child is present
// then it must be a HTML tag node else an error is returned.
func (pg *Page) render(data Data) *domu.Node {
	var doc = domu.HTMLDoc()
	pg.Layout.Render(pg, data, doc)
	return doc
}

func handlePath(targetPath string, more ...string) string {
	var base = targetPath
	if !strings.HasPrefix(targetPath, "/") {
		base = "/" + targetPath
	}
	if len(more) == 0 {
		return multipleSlashes.ReplaceAllString(base, "/")
	}
	var joined = path.Join(append([]string{base}, more...)...)
	return multipleSlashes.ReplaceAllString(joined, "/")
}

func cleanName(targetPath string) string {
	targetPath = strings.TrimSpace(targetPath)
	targetPath = strings.TrimLeft(targetPath, "/")
	targetPath = strings.TrimRight(targetPath, "/")
	return targetPath
}

func cleanPath(t string) string {
	return multipleSlashes.ReplaceAllString(t, "/")
}

func cleanAllSlashes(t string) string {
	return allSlashes.ReplaceAllString(t, "")
}

func reduceMultipleSlashToOne(t string) string {
	return multipleSlashes.ReplaceAllString(t, "/")
}
