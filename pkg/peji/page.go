package peji

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/influx6/groundlayer/pkg/domu"
	"github.com/influx6/groundlayer/pkg/styled"
)

var (
	multipleSlashes = regexp.MustCompile(`/{2,}`)
	allSlashes      = regexp.MustCompile(`^/$`)
)

// DOMHandler returns a Handler which provides the params as the data to be giving to
// a Live DOM render method.
func DOMHandler(target string, page *Page) Handler {
	return HandlerFunc(func(d Data) *domu.Node {
		return page.Live(target).Render(d)
	})
}

type OnPage func(route string, p *Page)

// Page implements the Document interface and represents
// a single unique html page.
type Page struct {
	UseChanges     bool
	Name           string
	RoutePath      string
	AnyRouteInPage string
	Router         *Mux
	Layout         Layout
	onPageAdded    []OnPage
	Theme          *styled.Theme
	Last           *domu.Node
	Registry       *DOMRegistry
}

func WithPage(name string, theme *styled.Theme, layout Layout, notFound Handler) *Page {
	return NewPageUsingRouter(handlePath(name), theme, layout, NewMux(MuxConfig{
		RootPath: name,
		NotFound: notFound,
	}))
}

func NewPageUsingRouter(routePath string, theme *styled.Theme, layout Layout, router *Mux) *Page {
	var p Page
	p.Name = routePath
	p.Layout = layout
	p.Router = router
	p.Theme = theme
	p.Registry = NewDOMRegistry()
	p.RoutePath = handlePath(routePath)
	p.AnyRouteInPage = handlePath(routePath, "*path")

	var pPointer = &p
	var handler = HandlerFunc(func(d Data) *domu.Node {
		return pPointer.Render(d)
	})

	p.Router.Serve(p.RoutePath, handler)
	p.Router.Serve(p.AnyRouteInPage, handler)
	return &p
}

func (pg *Page) Close() {
	pg.Registry.Reset()
}

func (pg *Page) check() {
	if pg.Registry == nil {
		pg.Registry = NewDOMRegistry()
	}
}

func (pg *Page) OnPageAdd(fn OnPage) {
	pg.onPageAdded = append(pg.onPageAdded, fn)
}

func (pg *Page) SetReconciliation(b bool) {
	pg.UseChanges = b
}

// Has returns true if giving route exists either as a static dom or dom list.
func (pg *Page) Has(route string) bool {
	pg.check()
	if strings.HasPrefix(route, pg.RoutePath) {
		route = handlePath(route)
		return pg.Registry.Has(route) || pg.Registry.HasList(route)
	}
	route = handlePath(pg.RoutePath, route)
	return pg.Registry.Has(route) || pg.Registry.HasList(route)
}

func (pg *Page) Serve(d Data) *domu.Node {
	return pg.Router.ServeRoute(d)
}

func (pg *Page) notifyOnPage(name string, p *Page) {
	for _, fn := range pg.onPageAdded {
		fn(name, p)
	}
}

func (pg *Page) ServeRoute(route string, handler Handler) {
	if !pg.Has(route) {
		panic(fmt.Sprintf("route %q does not exists for any registered dom", route))
	}

	if !strings.HasPrefix(route, pg.RoutePath) {
		route = handlePath(pg.RoutePath, route)
	} else {
		route = handlePath(route)
	}

	pg.serveRoute(route, handler)
}

func (pg *Page) serveRoute(route string, handler Handler) {
	var targetRoute = handlePath(pg.Name, route)
	var targetRouteWithWildcard = handlePath(pg.Name, route, "*path")

	pg.notifyOnPage(targetRoute, pg)

	pg.Router.Serve(targetRoute, handler)
	pg.Router.Serve(targetRouteWithWildcard, handler)
}

// GetLive returns the DOM for giving routePath, if not found, nil is returned.
func (pg *Page) Live(route string) DOM {
	pg.check()

	if !pg.Has(route) {
		return nil
	}

	if !strings.HasPrefix(route, pg.RoutePath) {
		route = handlePath(pg.RoutePath, route)
	} else {
		route = handlePath(route)
	}

	return pg.Registry.Get(route)
}

// AddLive adds a new routePath mapping for a provided list, ensuring these
// are referenced with giving component, if there is more than 1 item then they
// are treated as a whole and referenced by provided routePath.
//
// Note: It clears all live routes for this giving routePath, if we are replacing
// an existing live key.
func (pg *Page) AddLive(route string, items ...DOM) {
	pg.check()

	var itemCount = len(items)
	if itemCount == 0 {
		panic("DOM set must not be a count of zero")
	}

	if pg.Has(route) {
		panic("route must be unique across list and single component types")
	}

	var targetRoute string
	if !strings.HasPrefix(route, pg.RoutePath) {
		targetRoute = handlePath(pg.RoutePath, route)
	} else {
		targetRoute = handlePath(route)
	}

	var set *LiveDOM
	set = LiveFromDOMList(items, targetRoute)
	pg.Registry.Add(set, targetRoute)

	pg.serveRoute(targetRoute, HandlerFunc(func(data Data) *domu.Node {
		var rendered = set.Render(data)
		rendered.WalkTreeDeptFirst(func(child *domu.Node) bool {
			return handleThemeDirective(child, pg.Theme, rendered)
		})
		return rendered
	}))
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

	// Find head node
	var headNode, headNodeErr = doc.Find(findHead)
	if headNodeErr != nil {
		panic(fmt.Sprintf("Page layout requires the `head` element, we found no `head` node"))
	}

	// Find body node
	var _, bodyNodeErr = doc.Find(findBody)
	if bodyNodeErr != nil {
		panic(fmt.Sprintf("Page layout requires the `body` element, we found no `body` node"))
	}

	doc.WalkTreeDeptFirst(func(child *domu.Node) bool {
		return handleThemeDirective(child, pg.Theme, headNode)
	})

	return doc
}

func findBody(node *domu.Node, _ int) bool {
	if node.Name() == "body" {
		return true
	}
	return false
}

func findHead(node *domu.Node, _ int) bool {
	if node.Name() == "head" {
		return true
	}
	return false
}

func handleThemeDirective(child *domu.Node, theme *styled.Theme, root *domu.Node) bool {
	// ignore nodes without theme directives.
	if len(child.Themes.Directives) == 0 && child.Name() != groundLayerTagName {
		return true
	}

	// co-locate all the styles for a groundlayer-frame to it's groundlayer parent.
	if len(child.Themes.Directives) == 0 && child.Name() == groundLayerTagName {
		child.WalkTreeDeptFirst(func(node *domu.Node) bool {
			return handleThemeDirective(node, theme, child)
		})
		return false
	}

	// resolve themes for children into root.
	theme.Resolve(child.Themes, root)
	return true
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
