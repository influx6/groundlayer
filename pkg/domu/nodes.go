package domu

import (
	"errors"
	"io"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/influx6/npkg/natomic"
	"github.com/influx6/npkg/nerror"
	"github.com/influx6/npkg/nunsafe"
	"github.com/influx6/npkg/nxid"

	"github.com/influx6/groundlayer/pkg/styled"
)

const (
	emptyContent  = ""
	commentPrefix = "comment"
	textPrefix    = "text"
)

var (
	ErrInvalidNodeType = nerror.New("invalid node type, unsupported")
	ErrNotFound        = errors.New("not found")
)

// comment and tag constants
const (
	commentBegin  = "<!-- "
	commentEnd    = " -->"
	blockBegin    = "<"
	blockEnd      = ">"
	spacer        = " "
	equalSign     = "="
	quotation     = "\""
	hashString    = "#"
	forwardSlash  = "/"
	eventHeader   = "events"
	blockEndBegin = "</"
	newline       = "\n"
	tabulation    = "\t"
	htmlTag       = "html"
)

var spaceRegExp = regexp.MustCompile(`\s+`)

// NodeType defines a giving type of node.
type NodeType int

// const of node types.
const (
	DocumentNode         NodeType = 9
	DocumentFragmentNode NodeType = 11
	ElementNode          NodeType = 1
	TextNode             NodeType = 3
	CommentNode          NodeType = 8
	CarrierNode          NodeType = 100
	ReplicateNode        NodeType = 101
)

// String returns a text representation of a NodeType.
func (n NodeType) String() string {
	switch n {
	case DocumentNode:
		return "#DOCUMENT"
	case DocumentFragmentNode:
		return "#DOCUMENTFragment"
	case ElementNode:
		return "Element"
	case TextNode:
		return "PageName"
	case CommentNode:
		return "Comment"
	default:
		return "Unknown"
	}
}

// NodeFn defines a type which applies an operation to a node.
type NodeFn func(*Node) *Node

// Nodes defines an interface which expose a method for encoding
// a giving implementer sing provider NodeEncoder.
type Nodes interface {
	EncodeNode(NodeEncoder)
}

// NodeEncoder exposes method which provide means of encoding a giving
// node which is a combination of a name, attributes and child nodes.
type NodeEncoder interface {
	AttrEncodable

	EncodeChild(Nodes) error
	EncodeName(string) error
}

// Stringer defines an interface which exposes a method to
// get it's underline text using it's String() method.
type Stringer interface {
	String() string
}

// TextContent implements the Stringer interface for type string.
type TextContent string

// String returns the underline string type.
func (t TextContent) String() string {
	return string(t)
}

// Matchable requires implementers to provide methods for comparing
// against a giving node instance.
type Matchable interface {
	Match(*Node) bool
}

// Mounter defines an interface which exposes a method to mount a provided
// implementer into a provided Node instance. Basically the provided node
// is the parent to be mounted into.
type Mounter interface {
	Mount(parent *Node)
}

// FunctionApplier defines a function type that implements the Mounter interface.
type FunctionApplier func(*Node) error

// Mounter implements the Mounter interface.
func (fn FunctionApplier) Mounter(n *Node) error {
	return fn(n)
}

// HTMLDoc returns the document node type which has no parent
// and should be the start position of all nodes.
func HTMLDoc(renders ...Mounter) *Node {
	doc := NewNode(DocumentNode, "html")
	for _, mounter := range renders {
		mounter.Mount(doc)
		if err := doc.Err(); err != nil {
			panic(err.Error())
		}
	}
	return doc
}

// Replicate returns the Replicate node type which has no parent
// and should be the start position of a set nodes which will be
// applied to the parent to be mounted to.
func Replicate(renders ...Mounter) *Node {
	doc := NewNode(ReplicateNode, "doc")
	for _, mounter := range renders {
		mounter.Mount(doc)
		if err := doc.Err(); err != nil {
			panic(err.Error())
		}
	}
	return doc
}

// Carrier returns the Carrier node type which has no parent
// and should be the start position of a set nodes which will be
// applied to the parent to be mounted to.
func Carrier(renders ...Mounter) *Node {
	doc := NewNode(CarrierNode, "doc")
	for _, mounter := range renders {
		mounter.Mount(doc)
		if err := doc.Err(); err != nil {
			panic(err.Error())
		}
	}
	return doc
}

// Document returns the document node type which has no parent
// and should be the start position of all nodes.
func Document(renders ...Mounter) *Node {
	doc := NewNode(DocumentNode, "doc")
	for _, mounter := range renders {
		mounter.Mount(doc)
		if err := doc.Err(); err != nil {
			panic(err.Error())
		}
	}
	return doc
}

// Element returns a element node type which can be added into a parent
// or use as a base for other nodes.
func Element(name string, renders ...Mounter) *Node {
	doc := NewNode(ElementNode, name)
	for _, mounter := range renders {
		mounter.Mount(doc)
		if err := doc.Err(); err != nil {
			panic(err.Error())
		}
	}
	return doc
}

// TextType returns a element node type which can be added into a parent
// or use as a base for other nodes.
func TextType(textType NodeType, name string, content Stringer) *Node {
	doc := NewNode(textType, name)
	doc.content = content
	return doc
}

// Text returns a new Node of Text Type which has no children
// or attributes.
func Text(text string) *Node {
	return TextType(TextNode, TextNode.String(), TextContent(text))
}

// Comment returns a new Node of Comment Type which has no children
// or attributes.
func Comment(comment string) *Node {
	return TextType(CommentNode, CommentNode.String(), TextContent(comment))
}

// NodeList defines a type for slice of nodes, implementing the Mounter interface.
type NodeList []*Node

// Mount applies giving nodes in list to provided parent node.
func (n NodeList) Mount(parent *Node) error {
	for _, elem := range n {
		parent.AppendChild(elem)
		if err := parent.Err(); err != nil {
			return err
		}
	}
	return nil
}

// Node defines a concrete type implementing a combined linked-list with
// root, next and previous siblings using a underline growing array as
// the basis.
type Node struct {
	Themes      styled.ThemeDirective
	Attrs       AttrList
	Events      *EventHashList
	parent      *Node
	tid         string
	atid        string
	idAttr      string
	tagName     string
	routePrefix string
	content     Stringer
	nodeType    NodeType
	removed     bool
	changed     bool
	attrChanged bool
	index       *natomic.IntSwitch
	next        *natomic.IntSwitch
	prev        *natomic.IntSwitch
	kids        []*Node

	err error
}

// NewNode returns a new Node instance with the giving Node as
// underline parent pointer. It uses the provided `tagName` as
// name of node (i.e div or section) and the provided `idAttr`
// as id of giving node (i.e <div id={idAttr}>). This must be
// unique across a node child list.
//
// To important things to note in creating a node:
//
// 1. The tagName must be provided always as it tells the rendering
// system what the node represent.
//
// 2. The idAttr must both be provided an unique across all nodes, as
// it is the unique identifier to be used for referencing, replacement
// and swaps by the rendering system.
//
func NewNode(nt NodeType, tagName string) *Node {
	if tagName == "" {
		panic("tagName can not be empty")
	}

	var child Node
	child.nodeType = nt
	child.atid = child.tid
	child.tagName = tagName
	child.tid = nxid.New().String()
	child.Themes = []string{}

	child.Events = NewEventHashList()
	child.next = &natomic.IntSwitch{}
	child.prev = &natomic.IntSwitch{}
	child.index = &natomic.IntSwitch{}

	child.next.Flip(-1)
	child.prev.Flip(-1)
	child.index.Flip(-1)
	return &child
}

// Clone returns a clone of a giving node depending on the deepClone flag
// where if false then only the node and it's properties are cloned without the
// children, else all children are also cloned.
func (n *Node) Clone(deepClone bool) *Node {
	var newNode = NewNode(n.nodeType, n.tagName)
	newNode.tid = n.tid
	newNode.atid = n.atid
	newNode.idAttr = n.idAttr
	newNode.content = n.content

	for key, content := range n.Events.nodes {
		var events = make([]string, len(content))
		var copied = copy(events, content)
		newNode.Events.nodes[key] = events[:copied]
	}

	n.Attrs.Each(func(attr Attr) bool {
		newNode.Attrs.Add(attr.Clone())
		return true
	})

	if deepClone {
		newNode.kids = make([]*Node, 0, n.ChildCount())
		for _, node := range n.kids {
			var newChildNode = node.Clone(deepClone)
			newNode.AppendChild(newChildNode)
			if err := newNode.Err(); err != nil {
				break
			}
		}
	}

	return newNode
}

func (n *Node) Find(fn func(*Node, int) bool) (*Node, error) {
	for index, node := range n.kids {
		if fn(node, index) {
			return node, nil
		}
	}
	return nil, nerror.New("invalid op")
}

func (n *Node) EachChild(fn func(*Node, int) bool) {
	for index, node := range n.kids {
		if !fn(node, index) {
			return
		}
	}
}

// Apply applies giving set of functions to the node.
func (n *Node) Apply(fns ...NodeFn) *Node {
	if n.nodeType == ReplicateNode || n.nodeType == CarrierNode {
		n.EachChild(func(node *Node, i int) bool {
			node.Apply(fns...)
			return true
		})
		return n
	}

	var current = n
	for _, fn := range fns {
		current = fn(current)
	}
	return current
}

// Err returns the current error found when adding into this node.
func (n *Node) Err() error {
	return n.err
}

func (n *Node) SetErr(err error) {
	if n.parent != nil {
		n.parent.SetErr(err)
	}
	n.err = err
}

// Render implements the component interface.
func (n *Node) Render(data interface{}) (*Node, error) {
	return n, nil
}

// RenderHTML renders giving Nodes using a html markup syntax format.
//
// Underneath it calls Node.RenderHTMLTo (See comments for method).
func (n *Node) RenderHTML(w io.Writer, indented bool) error {
	if err := n.renderNode(w, indented, 0, true); err != nil {
		return nerror.WrapOnly(err)
	}
	return nil
}

// RenderHTMLDiff renders giving Nodes using a html markup syntax format.
//
// Underneath it calls Node.RenderHTMLTo (See comments for method).
func (n *Node) RenderHTMLDiff(w io.Writer, indented bool) error {
	if err := n.renderNode(w, indented, 0, false); err != nil {
		return nerror.WrapOnly(err)
	}
	return nil
}

// RenderHTMLTo renders giving Nodes using a html markup syntax format
// to provided strings.Builder. This allows memory efficient rendering.
//
// It implements an efficient means of using HTML as the defactor means of
// visualizing the produced output of a giving node and it's children.
//
// It runs depth-first collected all internal representation of a node, it's
// attributes and children.
func (n *Node) RenderHTMLTo(content io.Writer, indented bool, renderRemoved bool) error {
	if err := n.renderNode(content, indented, 0, renderRemoved); err != nil {
		return nerror.WrapOnly(err)
	}
	return nil
}

func (n *Node) renderNode(build io.Writer, indented bool, indentCount int, renderRemoved bool) error {
	// create the current tag for giving type of node.
	// Rules are:
	switch n.nodeType {
	case ReplicateNode, CarrierNode:
		return n.renderChildren(build, indented, indentCount, renderRemoved)
	case DocumentNode:
		// 1. if Document node then skip and render children, except for
		// html node.
		if n.Name() == "html" && n.parent == nil {
			return n.renderRoot(build, indented, indentCount, renderRemoved)
		}
		return n.renderChildren(build, indented, indentCount, renderRemoved)
	case TextNode:
		// 2. If text then render as a text node with no intricate tag as html allows
		// wrapping text.
		return n.renderText(build, indented, indentCount, renderRemoved)
	case CommentNode:
		// 2. If comment then render as a html comment node with appropriate
		// prefix and suffix.
		return n.renderComment(build, indented, indentCount, renderRemoved)
	case ElementNode:
		// 3. If a element node, then render the name, attributes, events then
		// children with enclosing tag.
		return n.renderElement(build, indented, indentCount, renderRemoved)
	default:
		// 4. If it's not a known type then return error.
		return ErrInvalidNodeType
	}
}

// ReconcileNotifier defines a function type which can be passed
// into the Node.Reconcile method which will be notified for changes
// be they a new node, removal of old node or swapping of changed nodes.
//
// The function is expected to return a boolean value which indicates
// if due to some internal error wants to force an immediate stop
// reconciliation operations to allow return.
type ReconcileNotifier func(changed *Node)

// Reconcile attempts to reconcile old Node set with this node set
// returning an true/false when such occurs, else updates this node with
// information regarding changes such as removals of node children.
//
// Reconcile will return true if this node should be swapped with
// old node in it's tree, as both the root it self has changed.
//
//  Reconciliation is done breath first, where the node is checked first
// against it's counter part and if there is matching state then all it's
// child nodes are checked and will be accordingly set for swapping or
// updating if not matching.
//
func (n *Node) Reconcile(old *Node) bool {
	if n.nodeType == ReplicateNode || n.nodeType == CarrierNode {
		panic("a ReplicatorNode or CarrierNode should never exists within a nodelist, they serve diff purposes, check logic")
	}

	if old.nodeType == ReplicateNode || old.nodeType == CarrierNode {
		panic("a ReplicatorNode or CarrierNode should never exists within a nodelist, they serve diff purposes, check logic")
	}

	if n.nodeType != old.nodeType {
		n.changed = true
		var cloned = old.Clone(false)
		cloned.removed = true
		n.AppendChild(cloned)
		n.changed = true
		return true
	}

	if n.Name() != old.Name() {
		n.changed = true
		var cloned = old.Clone(false)
		cloned.removed = true
		n.AppendChild(cloned)
		n.changed = true
		return true
	}

	// Set id of old node to our own ancestorID.
	n.atid = old.tid

	if !n.Attrs.MatchAttrs(old.Attrs) {
		n.changed = true
		n.attrChanged = true
		return true
	}

	if !n.Events.MatchEvents(old.Events) {
		n.changed = true
		n.attrChanged = true
		return true
	}

	if n.Type() == CommentNode || n.Type() == TextNode {
		if n.Text() == old.Text() {
			return false
		}
		n.changed = true
		return true
	}

	// if total kids are different, more or less (does not matter)
	// we know new node has changed either way.
	if len(old.kids) != len(n.kids) {
		n.changed = true
		return true
	}

	var changed bool

	// if we reached here, definitely the parent has not changed
	// in anyway, so check the kids and do a check if there is
	// a change, if we've maxed out kids in new then set others
	// as removed.
	old.EachChild(func(node *Node, i int) bool {
		newChild, err := n.Get(i)

		// if we fail to find a node matching index, then mark
		// as removed.
		if err != nil {
			changed = true

			var cloned = old.Clone(false)
			cloned.removed = true
			n.AppendChild(cloned)
			return true
		}

		if newChild != nil && newChild.Reconcile(node) {
			changed = true
		}
		return true
	})

	if !changed {
		n.tid = old.tid
		n.changed = false
		return false
	}

	n.changed = true
	return true
}

// RenderShallowHTML renders as HTML giving node tag alone without
// rendering children.
func (n *Node) RenderShallowHTML(build io.Writer, indented bool) error {
	if _, err := build.Write(nunsafe.String2Bytes(blockBegin)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(n.tagName)); err != nil {
		return err
	}
	if err := n.renderAttributes(build, indented); err != nil {
		return err
	}
	if err := n.renderEvents(build, indented); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(blockEnd)); err != nil {
		return err
	}
	if indented {
		if _, err := build.Write(nunsafe.String2Bytes(newline)); err != nil {
			return err
		}
	}
	if indented {
		if _, err := build.Write(nunsafe.String2Bytes(newline)); err != nil {
			return err
		}
	}
	if _, err := build.Write(nunsafe.String2Bytes(blockEndBegin)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(n.tagName)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(blockEnd)); err != nil {
		return err
	}
	return nil
}

func (n *Node) renderComment(build io.Writer, indented bool, indentCount int, renderRemoved bool) error {
	if n.removed && !renderRemoved {
		return nil
	}
	if n.content != nil {
		if indented {
			if indentCount > 0 {
				for i := indentCount; i > 0; i-- {
					if _, err := build.Write(nunsafe.String2Bytes(tabulation)); err != nil {
						return err
					}
				}
			}
		}
		if _, err := build.Write(nunsafe.String2Bytes(commentBegin)); err != nil {
			return err
		}
		if indented {
			if _, err := build.Write(nunsafe.String2Bytes(newline)); err != nil {
				return err
			}
			if indentCount > 0 {
				for i := indentCount; i > 0; i-- {
					if _, err := build.Write(nunsafe.String2Bytes(tabulation)); err != nil {
						return err
					}
				}
			}
		}
		if _, err := build.Write(nunsafe.String2Bytes(n.content.String())); err != nil {
			return err
		}
		if indented {
			if _, err := build.Write(nunsafe.String2Bytes(newline)); err != nil {
				return err
			}
			if indentCount > 0 {
				for i := indentCount; i > 0; i-- {
					if _, err := build.Write(nunsafe.String2Bytes(tabulation)); err != nil {
						return err
					}
				}
			}
		}
		if _, err := build.Write(nunsafe.String2Bytes(commentEnd)); err != nil {
			return err
		}
		if indented {
			if _, err := build.Write(nunsafe.String2Bytes(newline)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (n *Node) renderText(build io.Writer, indented bool, indentCount int, renderRemoved bool) error {
	if n.removed && !renderRemoved {
		return nil
	}
	if n.content != nil {
		if indented {
			if indentCount > 0 {
				for i := indentCount; i > 0; i-- {
					if _, err := build.Write(nunsafe.String2Bytes(tabulation)); err != nil {
						return err
					}
				}
			}
		}
		if _, err := build.Write(nunsafe.String2Bytes(n.content.String())); err != nil {
			return err
		}
		if indented {
			if _, err := build.Write(nunsafe.String2Bytes(newline)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (n *Node) renderAttributes(build io.Writer, indented bool) error {
	if _, err := build.Write(nunsafe.String2Bytes(spacer)); err != nil {
		return err
	}

	var err error
	var attrBuilder strings.Builder
	var encoder = DOMAttrEncoderWith("", &attrBuilder)

	if n.atid != n.tid && n.atid != "" {
		if err = encoder.QuotedString("atid", n.atid); err != nil {
			return err
		}
	}

	if err = encoder.QuotedString("_tid", n.tid); err != nil {
		return err
	}

	if n.removed {
		if err = encoder.QuotedString("_removed", "true"); err != nil {
			return err
		}
	}

	if err = encoder.QuotedString("_ref", n.RefTree()); err != nil {
		return err
	}

	if n.Attrs.Len() == 0 {
		_, _ = build.Write(nunsafe.String2Bytes(attrBuilder.String()))
		return nil
	}

	n.Attrs.Each(func(attr Attr) bool {
		if err = attr.EncodeAttr(encoder); err != nil {
			return false
		}
		return true
	})

	if err != nil {
		return err
	}

	_, _ = build.Write(nunsafe.String2Bytes(attrBuilder.String()))
	return nil
}

func (n *Node) renderEvents(build io.Writer, indented bool) error {
	if n.Events.Len() == 0 {
		return nil
	}

	if _, err := build.Write(nunsafe.String2Bytes(spacer)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(eventHeader)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(equalSign)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(quotation)); err != nil {
		return err
	}

	var eventBuild strings.Builder
	if err := n.Events.EncodeEvents(&eventBuild); err != nil {
		return err
	}

	_, _ = build.Write(nunsafe.String2Bytes(eventBuild.String()))
	if _, err := build.Write(nunsafe.String2Bytes(quotation)); err != nil {
		return err
	}
	return nil
}

func (n *Node) renderElement(build io.Writer, indented bool, indentCount int, renderRemoved bool) error {
	if indented {
		if indentCount > 0 {
			for i := indentCount; i > 0; i-- {
				if _, err := build.Write(nunsafe.String2Bytes(tabulation)); err != nil {
					return err
				}
			}
		}
	}

	if n.removed && !renderRemoved {
		return nil
	}

	if _, err := build.Write(nunsafe.String2Bytes(blockBegin)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(n.tagName)); err != nil {
		return err
	}
	if err := n.renderAttributes(build, indented); err != nil {
		return err
	}
	if err := n.renderEvents(build, indented); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(blockEnd)); err != nil {
		return err
	}
	if indented {
		if _, err := build.Write(nunsafe.String2Bytes(newline)); err != nil {
			return err
		}
	}

	var newIndent = indentCount + 1
	if err := n.renderChildren(build, indented, newIndent, renderRemoved); err != nil {
		return err
	}
	if indented {
		if indentCount > 0 {
			for i := indentCount; i > 0; i-- {
				if _, err := build.Write(nunsafe.String2Bytes(tabulation)); err != nil {
					return err
				}
			}
		}
	}
	if _, err := build.Write(nunsafe.String2Bytes(blockEndBegin)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(n.tagName)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(blockEnd)); err != nil {
		return err
	}
	if indented {
		if _, err := build.Write(nunsafe.String2Bytes(newline)); err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) renderRoot(build io.Writer, indented bool, indentCount int, renderRemoved bool) error {
	if _, err := build.Write(nunsafe.String2Bytes(blockBegin)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(htmlTag)); err != nil {
		return err
	}
	if err := n.renderAttributes(build, indented); err != nil {
		return err
	}
	if err := n.renderEvents(build, indented); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(blockEnd)); err != nil {
		return err
	}
	if indented {
		if _, err := build.Write(nunsafe.String2Bytes(newline)); err != nil {
			return err
		}
	}
	if err := n.renderChildren(build, indented, indentCount, renderRemoved); err != nil {
		return err
	}
	if indented {
		if _, err := build.Write(nunsafe.String2Bytes(newline)); err != nil {
			return err
		}
	}
	if _, err := build.Write(nunsafe.String2Bytes(blockEndBegin)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(htmlTag)); err != nil {
		return err
	}
	if _, err := build.Write(nunsafe.String2Bytes(blockEnd)); err != nil {
		return err
	}
	return nil
}

func (n *Node) renderChildren(build io.Writer, indented bool, indentCount int, renderRemoved bool) (err error) {
	n.EachChild(func(node *Node, _ int) bool {
		if err = node.renderNode(build, indented, indentCount, renderRemoved); err != nil {
			return false
		}
		return true
	})
	return
}

// Get returns a giving Node at giving index, if no removals
// have being done before this call, then insertion order of
// children nodes are maintained, else before using this method
// ensure to call Node.Balance() to restore insertion order.
func (n *Node) Get(index int) (*Node, error) {
	if index >= len(n.kids) {
		return nil, nerror.New("bad op")
	}
	return n.kids[index], nil
}

// RefID returns the reference id of giving node.
func (n *Node) RefID() string {
	return n.tid
}

func (n *Node) SetTid(id string) {
	n.tid = id
}

func (n *Node) SetAtid(id string) {
	n.atid = id
}

// SetRoutePart sets the routePrefix portion of this node.
func (n *Node) SetPrefix(route string) {
	n.routePrefix = strings.TrimSpace(route)
}

// RefTree returns the tree path for giving node by collecting
// the parents ID till this node.
func (n *Node) RefTree() string {
	var content = stringPool.Get().(*strings.Builder)
	defer stringPool.Put(content)
	content.Reset()

	// write out tree ref from root or ancestor parent.
	n.writeRefTree(content)
	return content.String()
}

// writeRefTree will walk up the node tree writing out
// all parent paths into provided string builder.
func (n *Node) writeRefTree(b *strings.Builder) {
	if n.parent != nil {
		n.parent.writeRefTree(b)
	}
	if n.routePrefix != "" {
		b.Write(nunsafe.String2Bytes(forwardSlash))
		b.Write(nunsafe.String2Bytes(n.routePrefix))
	}
	b.Write(nunsafe.String2Bytes(forwardSlash))
	b.Write(nunsafe.String2Bytes(n.idAttr))
}

// FindRefNode walks through provided ref node paths till the last node
// is found.
func (n *Node) FindRefNode(ref string) (*Node, error) {
	var refs = strings.Split(strings.TrimSpace(ref), "/")
	return n.findRefNodeFromList(refs)
}

func (n *Node) findRefNodeFromList(ref []string) (*Node, error) {
	var firstId = strings.TrimSpace(ref[0])
	firstId = strings.Trim(firstId, "/")
	if n.idAttr != firstId {
		return nil, nerror.New("node id %q does not match %q", n.idAttr, firstId)
	}

	var rest = ref[1:]
	if len(rest) == 0 {
		return n, nil
	}

	var secondId = strings.TrimSpace(rest[0])
	var targetNode, err = n.Find(func(node *Node, _ int) bool {
		return node.idAttr == secondId
	})
	if err != nil {
		return n, nerror.WrapOnly(err)
	}

	return targetNode.findRefNodeFromList(rest)
}

// ID returns user-provided id of giving node.
func (n *Node) ID() string {
	return n.idAttr
}

// Name returns the name of giving node (i.e the node name).
func (n *Node) Name() string {
	return n.tagName
}

// Text returns the underline text content of a node if it's a
// TextNode.
func (n *Node) Text() string {
	if n.content == nil {
		return emptyContent
	}
	return n.content.String()
}

// Type returns the underline type of giving node.
func (n *Node) Type() NodeType {
	return n.nodeType
}

// Parent returns the underline parent of giving Node.
func (n *Node) Parent() *Node {
	return n.parent
}

// Match returns true/false if provided node matches this node
// exactly in type, name and attributes.
func (n *Node) Match(other *Node) bool {
	if n.Type() != other.Type() {
		return false
	}

	if n.Type() != CommentNode && n.Type() != TextNode {
		if n.Name() != other.Name() {
			return false
		}

		if !n.Attrs.MatchAttrs(other.Attrs) {
			return false
		}

		return true
	}

	if n.Text() != other.Text() {
		return false
	}

	return true
}

// Mount implements the Mounter interface where it mounts the provided
// node as a child node of it self.
//
// Operation will fail if the node.Err() has an error.
func (n *Node) Mount(parent *Node) {
	parent.AppendChild(n)
}

func (n *Node) addChild(kid *Node) {
	if len(n.kids) == cap(n.kids) {
		var newKids = make([]*Node, cap(n.kids)*2)
		var written = copy(newKids, n.kids)
		n.kids = newKids[:written]
	}

	var last = len(n.kids)
	n.kids = append(n.kids, kid)

	kid.prev.Flip(last)
	kid.index.Flip(len(n.kids))
	kid.next.Flip(-1)

	var lastChild = n.kids[last]
	lastChild.next.Flip(len(n.kids))
}

// AppendChild adds new child into node tree creating a relationship of child
// and parent.
//
// Comment and Text nodes can't have children.
//
// Operation will fail if the node.Err() has an error.
func (n *Node) AppendChild(kid *Node) {
	if n.err != nil {
		return
	}

	if kid.err != nil {
		n.err = kid.err
		return
	}

	if n == kid {
		return
	}

	if n.Type() == CommentNode || n.Type() == TextNode {
		n.SetErr(nerror.New("noop"))
		return
	}

	if kid.nodeType == ReplicateNode {
		kid.EachChild(func(node *Node, _ int) bool {
			var cloned = node.Clone(true)
			cloned.parent = n
			n.addChild(cloned)
			return true
		})
		return
	}

	if kid.nodeType == CarrierNode {
		kid.EachChild(func(node *Node, _ int) bool {
			node.parent = n
			n.addChild(node)
			return true
		})
		kid.kids = kid.kids[:0]
		return
	}

	n.addChild(kid)
	kid.parent = n
}

// FirstChild returns first child of giving node if any,
// else returns an error.
func (n *Node) FirstChild() (*Node, error) {
	if len(n.kids) == 0 {
		return nil, nerror.New("noop")
	}
	return n.kids[0], nil
}

// LastChild returns last child of giving node if any,
// else returns an error.
func (n *Node) LastChild() (*Node, error) {
	if len(n.kids) == 0 {
		return nil, nerror.New("noop")
	}
	return n.kids[len(n.kids)-1], nil
}

// PreviousSibling returns a node that is previous to this giving
// node in it's parent's list.
func (n *Node) PreviousSibling() (*Node, error) {
	if n.parent == nil {
		return nil, nerror.New("noop")
	}
	return n.parent.Get(n.prev.Read())
}

// NextSibling returns a node that is next to this giving
// node in it's parent's list.
func (n *Node) NextSibling() (*Node, error) {
	if n.parent == nil {
		return nil, nerror.New("noop")
	}
	return n.parent.Get(n.next.Read())
}

// NodeAttr returns a Attr for giving node.
func (n *Node) NodeAttr() NodeAttr {
	return NodeAttr{
		Type: n.nodeType,
		Ref:  n.tid,
		ID:   n.idAttr,
		Name: n.tagName,
	}
}

// ChildCount returns the current total count of kids.
func (n *Node) ChildCount() int {
	return len(n.kids)
}

// ResetNode resets giving node alone without affecting it's underline sub-tree.
//
// It keeps it's children as it was for re-use. See ResetTree for a more
// expansive reset call.
func (n *Node) ResetNode() {
	n.reset()
	n.Events.Reset()
	n.Attrs = n.Attrs[:0]
}

// ResetTree resets both the node and it's children nodes in a depth-first manner.
//
// It accepts a function which will be called against this node and all children nodes
// to allow user provide a garbage action like adding nodes back into an object pool.
//
// The list of nodes is set back to empty once done, allowing this node to be re-used.
func (n *Node) ResetTree(doNode func(*Node)) {
	n.ResetNode()

	n.EachChild(func(child *Node, _ int) bool {
		child.ResetTree(doNode)
		return true
	})

	n.kids = n.kids[:0]

	if doNode != nil {
		doNode(n)
	}
}

func (n *Node) reset() {
	n.parent = nil
	n.next.Flip(-1)
	n.prev.Flip(-1)
}

// ****************************************************************************
// Class List and ID List
// ****************************************************************************

// IDList defines a map type containing a giving class and
// associated nodes that match said classes.
type IDList map[string]NodeHashList

// Add adds giving node if it has a class key into
// giving IDList.
//
// It panics if it sees a id of type already existing.
func (c IDList) Add(n *Node) {
	set, ok := c[n.ID()]
	if ok && set.Count() != 0 {
		panic("Node with id '" + n.ID() + "' already exists")
	}

	set.Add(n)
	c[n.ID()] = set
}

// Remove removes giving node from possible class list found
// in it.
func (c IDList) Remove(n *Node) {
	if set, ok := c[n.ID()]; ok {
		set.Remove(n)
		return
	}
}

// NodeHashList implements the a set list for Nodes using
// their Node.RefID() value as unique keys.
type NodeHashList struct {
	nodes map[string]*Node
}

// Reset resets the internal hashmap used for storing
// nodes. There by removing all registered nodes.
func (na *NodeHashList) Reset() {
	na.nodes = map[string]*Node{}
}

// Count returns the total content count of map
func (na *NodeHashList) Count() int {
	if na.nodes == nil {
		return 0
	}
	return len(na.nodes)
}

// Add adds giving node into giving list if it has
// giving attribute value.
func (na *NodeHashList) Add(n *Node) {
	if na.nodes == nil {
		na.nodes = map[string]*Node{}
	}
	na.nodes[n.RefID()] = n
}

// Remove removes giving node into giving list if it has
// giving attribute value.
func (na *NodeHashList) Remove(n *Node) {
	if na.nodes == nil {
		na.nodes = map[string]*Node{}
	}
	delete(na.nodes, n.RefID())
}

// NodeAttr contains basic information about a giving node.
type NodeAttr struct {
	Name string
	ID   string
	Ref  string
	Type NodeType
}

// NodeAttrList implements the a set list for Nodes using
// their Node.RefID() value as unique keys.
type NodeAttrList struct {
	nodes map[string]NodeAttr
}

// Reset resets the internal hashmap used for storing
// nodes. There by removing all registered nodes.
func (na *NodeAttrList) Reset() {
	na.nodes = map[string]NodeAttr{}
}

// Count returns the total content count of map
func (na *NodeAttrList) Count() int {
	if na.nodes == nil {
		return 0
	}
	return len(na.nodes)
}

// Add adds giving node into giving list if it has
// giving attribute value.
func (na *NodeAttrList) Add(n *Node) {
	if na.nodes == nil {
		na.nodes = map[string]NodeAttr{}
	}
	na.nodes[n.RefID()] = n.NodeAttr()
}

// Remove removes giving node into giving list if it has
// giving attribute value.
func (na *NodeAttrList) Remove(n *Node) {
	if na.nodes == nil {
		na.nodes = map[string]NodeAttr{}
	}
	delete(na.nodes, n.RefID())
}

// ****************************************************************************
// slidingList
// ****************************************************************************

// increasing factor provides a base increase size for
// a node's internal slice/array. It is used in the calculation
// of said size increment calculation for growth rate.
const increasingFactor = 5

// slidingList implements a efficient random access compact list
// where elements can be moved from end of list to fill up opened
// positions within the list.
//
// It uses a node index pointers which allows moving nodes around
// easily, this unfortunately breaks the consistency guarantees of
// node's within their index as their position, hence ensure to sort
// the list after any all possible remove operations to regain
// positional guarantees.
//
// slidingList is not safe for concurrent use.
type slidingList []*Node

// **********************************************
// utils
// **********************************************

var alphanums = []rune("bcdfghjklmnpqrstvwxz0123456789")

func randomString(length int) string {
	var total = len(alphanums)
	var nowTime = time.Now()
	var parts = strconv.Itoa(nowTime.Year())[:2] + strconv.Itoa(nowTime.Minute())
	var b = make([]rune, length-len(parts))

	var next int
	for i := range b {
		next = rand.Intn(total)
		if next == total {
			next--
		}
		b[i] = alphanums[next]
	}

	return parts + string(b)
}

func prefixRandomString(prefix string, length int) string {
	var nowTime = time.Now()
	var total = len(alphanums)
	var parts = strconv.Itoa(nowTime.Year())[:2] + strconv.Itoa(nowTime.Minute())
	var alloc = length - len(parts)
	var space = len(prefix) + len(parts) + alloc + 1

	var contents = make([]byte, 0, space)
	contents = append(contents, prefix...)
	contents = append(contents, '-')
	contents = append(contents, parts...)

	var next int
	var b = make([]byte, alloc)
	for i := range b {
		next = rand.Intn(total)
		if next == total {
			next--
		}
		b[i] = byte(alphanums[next])
	}

	contents = append(contents, b...)
	return string(contents)
}
