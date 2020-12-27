package peji

import (
	"strings"

	"github.com/influx6/npkg"

	"github.com/influx6/groundlayer/pkg/domu"
)

type Params map[string]string

func (h Params) EncodeObject(encoder npkg.ObjectEncoder) {
	for key, value := range h {
		encoder.String(key, value)
	}
}

func (h Params) Get(k string) string {
	return h[k]
}

func (h Params) Set(k string, v string) {
	h[k] = v
}

func (h Params) Delete(k string) {
	delete(h, k)
}

type Header map[string][]string

func (h Header) Get(k string) string {
	if values, ok := h[k]; ok {
		return values[0]
	}
	return ""
}

func (h Header) Values(k string) []string {
	return h[k]
}

func (h Header) Add(k string, v string) {
	h[k] = append(h[k], v)
}

func (h Header) Set(k string, v string) {
	h[k] = append([]string{v}, v)
}

func (h Header) Delete(k string) {
	delete(h, k)
}

type Handler interface {
	Handle(Data) *domu.Node
}

type HandlerFunc func(Data) *domu.Node

func (h HandlerFunc) Handle(d Data) *domu.Node {
	return h(d)
}

type MuxConfig struct {
	RootPath string
	NotFound Handler
}

// Mux is Request multiplexer.
// It matches an event routePath or http url pattern to
// a specific TransportHandler which will be registered
// to the provided transport for handling specific events.
type Mux struct {
	config   MuxConfig
	rootPath string
	trie     *Trie
	NotFound Handler
}

func NewMux(config MuxConfig) *Mux {
	return &Mux{
		config:   config,
		rootPath: config.RootPath,
		trie:     NewTrie(),
		NotFound: config.NotFound,
	}
}

func (m *Mux) Serve(route string, handler Handler) {
	var searchRoute = reduceMultipleSlashToOne(m.rootPath + route)
	m.trie.Insert(searchRoute, WithHandler(handler))
}

// Match implements the Matcher interface.
//
// Allow a mux to be used as a matcher and handler elsewhere.
func (m *Mux) Match(target string) bool {
	var handler = m.trie.Search(target, Params{})
	return handler != nil
}

func (m *Mux) ServeRoute(d Data) *domu.Node {
	var params = d.Params
	if params == nil {
		d.Params = Params{}
		params = d.Params
	}

	var reqPath = d.Path
	if len(reqPath) > 1 && strings.HasSuffix(reqPath, "/") {
		// Remove trailing slash and client-permanent rule for redirection,
		// if configuration allows that and reqPath has an extra slash.

		// update the new reqPath and redirect.
		// use Trim to ensure there is no open redirect due to two leading slashes
		reqPath = pathSep + strings.Trim(reqPath, pathSep)
	}

	// r.URL.Query() is slow and will allocate a lot, although
	// the first idea was to not introduce a new type to the end-developers
	// so they are using this library as the std one, but we will have to do it
	// for the params, we keep that rule so a new ResponseWriter, which is an interface,
	// and it will be compatible with net/http will be introduced to store the params at least,
	// we don't want to add a third parameter or a global state to this library.
	var targetNode = m.trie.Search(reqPath, params)
	if targetNode == nil {
		return m.NotFound.Handle(d)
	}
	return targetNode.Handler.Handle(d)
}

func indexOfList(vs []string, m string) int {
	for index, v := range vs {
		if v == m {
			return index
		}
	}
	return -1
}
func toLower(vs []string) []string {
	for index, v := range vs {
		vs[index] = strings.ToLower(v)
	}
	return vs
}
