package domu

import (
	"errors"
	"io"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/influx6/npkg/natomic"
	"github.com/influx6/npkg/nerror"
	"github.com/influx6/npkg/nunsafe"
	"github.com/influx6/npkg/nxid"
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
	kids        *slidingList

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
	child.tagName = tagName
	child.kids = &slidingList{}
	child.tid = nxid.New().String()
	child.atid = child.tid

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
		n.kids.Each(func(node *Node, _ int) bool {
			var newChildNode = node.Clone(deepClone)
			newNode.AppendChild(newChildNode)
			if err := newNode.Err(); err != nil {
				return false
			}
			return true
		})
	}

	return newNode
}

func (n *Node) EachChild(fn func(*Node, int) bool) {
	n.kids.Each(fn)
}

// SwapAll swaps provided node with myself within
// parent list, the swapped node replaces me and
// my children nodes list.
func (n *Node) SwapAll(m *Node) error {
	if n.parent == nil {
		return ErrInvalidOp
	}

	if err := n.parent.kids.SwapNode(n.index.Read(), m, false); err != nil {
		return err
	}

	n.reset()
	return nil
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
	if old.kids.Length() != n.kids.Length() {
		n.changed = true
		return true
	}

	// if we matched, then swap our ids to ensure we can locate old node
	// in rendering.
	// ensure our list is also sorted.
	n.kids.SortList()

	// ensure child list of old node is sorted.
	old.kids.SortList()

	var changed bool

	// if we reached here, definitely the parent has not changed
	// in anyway, so check the kids and do a check if there is
	// a change, if we've maxed out kids in new then set others
	// as removed.
	old.kids.Each(func(node *Node, i int) bool {
		newChild, err := n.kids.Get(i)

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

	if n.atid != n.tid {
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
	n.kids.Each(func(node *Node, _ int) bool {
		if err = node.renderNode(build, indented, indentCount, renderRemoved); err != nil {
			return false
		}
		return true
	})
	return
}

// SwapNode swaps provided node with myself within parent's list. The swapped node
// takes over my children node list.
func (n *Node) SwapNode(m *Node) error {
	if n.parent == nil {
		return ErrInvalidOp
	}

	if err := n.parent.kids.SwapNode(n.index.Read(), m, true); err != nil {
		return err
	}

	// reset my properties.
	n.reset()
	return nil
}

// Get returns a giving Node at giving index, if no removals
// have being done before this call, then insertion order of
// children nodes are maintained, else before using this method
// ensure to call Node.Balance() to restore insertion order.
func (n *Node) Get(index int) (*Node, error) {
	return n.kids.Get(index)
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
	var targetNode, err = n.kids.Find(func(node *Node, _ int) bool {
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

// Remove removes this giving node from it's parent node list.
func (n *Node) Remove() error {
	if n.parent == nil {
		return ErrInvalidOp
	}

	if _, err := n.parent.kids.RemoveAndSwap(n.index.Read()); err != nil {
		return err
	}

	n.parent = nil
	return nil
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

// AppendChild adds new child into node tree creating a relationship of child
// and parent.
//
// Comment and Text nodes can have children but they must be of the same
// type as their parent.
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
		n.SetErr(ErrInvalidOp)
		return
	}

	if kid.nodeType == ReplicateNode {
		kid.EachChild(func(node *Node, _ int) bool {
			var cloned = node.Clone(true)
			cloned.parent = n
			if _, addErr := n.kids.Add(cloned); addErr != nil {
				n.SetErr(addErr)
			}
			return true
		})
		return
	}

	if kid.nodeType == CarrierNode {
		kid.EachChild(func(node *Node, _ int) bool {
			node.parent = n
			if _, addErr := n.kids.Add(node); addErr != nil {
				n.SetErr(addErr)
			}
			return true
		})
		kid.kids.Clear()
		return
	}

	if _, err := n.kids.Add(kid); err != nil {
		n.SetErr(ErrInvalidOp)
		return
	}

	kid.parent = n
}

// Balance runs optimization operations to the list of child nodes for
// this node. Node uses a sliding list underneath which trades positional
// guarantees (i.e consistently having a node at a giving index within the used list)
// for efficient memory management and random access speed during write operations like remove
// which can leave blank spots within underline list leaving your memory fragmented.
//
// This allows us the benefit of aggregating all write operations like remove an into a
// single set of calls, which can then be balance at the end using this function with
// little cost.
//
// If a removal is never done, or if only swaps are done, then Balance does nothing as those
// still maintain random access guarantees.
func (n *Node) Balance() {
	n.kids.SortList()
}

// FirstChild returns first child of giving node if any,
// else returns an error.
func (n *Node) FirstChild() (*Node, error) {
	return n.kids.FirstNode()
}

// LastChild returns last child of giving node if any,
// else returns an error.
func (n *Node) LastChild() (*Node, error) {
	return n.kids.LastNode()
}

// PreviousSibling returns a node that is previous to this giving
// node in it's parent's list.
func (n *Node) PreviousSibling() (*Node, error) {
	if n.parent == nil {
		return nil, ErrInvalidOp
	}
	return n.parent.Get(n.prev.Read())
}

// NextSibling returns a node that is next to this giving
// node in it's parent's list.
func (n *Node) NextSibling() (*Node, error) {
	if n.parent == nil {
		return nil, ErrInvalidOp
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
	return n.kids.Length()
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

	n.kids.Each(func(child *Node, _ int) bool {
		child.ResetTree(doNode)
		return true
	})

	n.kids.items = n.kids.items[:0]

	if doNode != nil {
		doNode(n)
	}
}

func (n *Node) reset() {
	n.parent = nil
	n.next.Flip(-1)
	n.prev.Flip(-1)
}

//****************************************************************************
// Class List and ID List
//****************************************************************************

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

//****************************************************************************
// slidingList
//****************************************************************************

// increasing factor provides a base increase size for
// a node's internal slice/array. It is used in the calculation
// of said size increment calculation for growth rate.
const increasingFactor = 5

// errors
var (
	// ErrInvalidIndex is returned when giving index is out of range or below 0.
	ErrInvalidIndex = errors.New("index is out of range")

	// ErrInvalidOp is returned when an operation is impossible.
	ErrInvalidOp = errors.New("operation can not be performed")

	// ErrEmptyList is returned when given list is empty.
	ErrEmptyList = errors.New("list is empty")

	// ErrIndexNotEmpty is returned when index has a element and not empty.
	ErrIndexNotEmpty = errors.New("index has element")

	// ErrEmptyIndex is returned when index has no element.
	ErrEmptyIndex = errors.New("index has no element")
)

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
type slidingList struct {
	items []*Node

	// lf is the lastAddFactor containing the last total items added before
	// an expansion.
	lf uint32

	// dirty signifies if list requires resorting to regain
	// positional guarantees.
	dirty uint32

	// lastNode represents the next index of last node.
	lastNode *natomic.IntSwitch

	// firstNode points to the giving node which stands in first
	// place as in
	firstNode *natomic.IntSwitch
}

// newNodeArrayList returns a new instance of a slidingList.
func newNodeArrayList() *slidingList {
	return &slidingList{}
}

// Capacity returns the underline size of the array of the used slice.
func (al *slidingList) Capacity() int {
	return cap(al.items)
}

// Clear resets the underline slice, emptying
// all elements within.
func (al *slidingList) Clear() {
	al.items = al.items[:0]
}

// Length returns the length of the underline slice.
func (al *slidingList) Length() int {
	return len(al.items)
}

// LastIndex returns the index value of the last item in
// the list.
func (al *slidingList) LastIndex() int {
	if al.lastNode == nil {
		return 0
	}
	return al.lastNode.Read()
}

// Available returns total space available to be filled
// left in array.
func (al *slidingList) Available() int {
	return cap(al.items) - len(al.items)
}

// List returns the underline slice of giving list.
func (al *slidingList) List() []*Node {
	return al.items
}

// Get returns giving element at provided index else returns an
// error if index is out of range. If slot is empty then returns
// error ErrEmptyIndex.
func (al *slidingList) Get(index int) (*Node, error) {
	if index >= len(al.items) || index < 0 || len(al.items) == 0 {
		return nil, ErrInvalidIndex
	}
	if al.items[index] == nil {
		return nil, ErrEmptyIndex
	}
	return al.items[index], nil
}

// LastNode returns possible last node within sliding list.
func (al *slidingList) LastNode() (*Node, error) {
	if al.lastNode != nil {
		return al.items[al.lastNode.Read()], nil
	}
	return nil, ErrInvalidOp
}

// FirstNode returns possible first node within sliding list.
func (al *slidingList) FirstNode() (*Node, error) {
	if al.firstNode != nil {
		return al.items[al.firstNode.Read()], nil
	}
	return nil, ErrInvalidOp
}

// GetNext returns the next Node from the giving index. It returns the
// next sibling if any from the provided index.
func (al *slidingList) GetNext(index int) (*Node, error) {
	if index >= len(al.items) || index < 0 || len(al.items) == 0 {
		return nil, ErrInvalidIndex
	}

	next := index + 1
	if al.items[next] == nil {
		return nil, ErrEmptyIndex
	}
	return al.items[next], nil
}

// GetPrevious returns the previous Node from the giving index. It returns the
// previous sibling if any from the provided index.
func (al *slidingList) GetPrevious(index int) (*Node, error) {
	if index >= len(al.items) || index < 0 || len(al.items) == 0 {
		return nil, ErrInvalidIndex
	}

	if index == 0 {
		return nil, ErrInvalidOp
	}

	prev := index - 1
	if al.items[prev] == nil {
		return nil, ErrEmptyIndex
	}
	return al.items[prev], nil
}

// Add adds giving Node into the array list, expanding giving node
// as needed. It returns the index where the item was added.
func (al *slidingList) Add(n *Node) (int, error) {
	if al.items == nil {
		al.items = make([]*Node, 0, al.getLastExpansion())
	}
	if al.Available() <= 2 {
		al.expandList()
	}

	var index = len(al.items)
	n.index.Flip(index)
	al.items = append(al.items, n)

	if al.firstNode == nil {
		al.firstNode = al.lastNode
	}

	if al.lastNode != nil {
		var lastNode = al.items[al.lastNode.Read()]
		lastNode.next.Flip(index)
		n.prev.Flip(al.lastNode.Read())
	}

	al.lastNode = n.index
	if al.firstNode == nil {
		al.firstNode = n.index
	}

	al.incrementLastExpansion(1)
	return index, nil
}

// Last returns the last node in the list.
func (al *slidingList) Last() *Node {
	if al.lastNode == nil {
		return nil
	}
	return al.items[al.lastNode.Read()]
}

// First returns the first node in the list.
func (al *slidingList) First() *Node {
	if al.firstNode == nil {
		return nil
	}
	return al.items[al.firstNode.Read()]
}

// Empty returns true/false if giving list is empty.
func (al *slidingList) Empty() bool {
	return len(al.items) == 0
}

// ToList returns a ordered list of Node items according
// to their link arrangement in the underline list.
func (al *slidingList) ToList() []*Node {
	// run through the list of items
	var count = len(al.items)
	if al.firstNode == nil && count > 1 {
		panic("Invalid pointer to first item, check logic")
	}

	al.SortList()
	return al.items
}

// SortList resets the underline list if dirty according to
// it's insertion order, returning the underline list back into
// positional guarantees.
func (al *slidingList) SortList() {
	if !al.isDirty() {
		return
	}

	if al.firstNode == nil && al.lastNode != nil {
		panic("Invalid pointer state for slidingList.firstNode")
	}

	var count = len(al.items)
	if count == 0 {
		return
	}

	var items = make([]*Node, 0, count)

	var index int
	var next = al.firstNode.Read()

	for next != -1 {
		item := al.items[next]
		next = item.next.Read()

		item.index.Flip(index)
		item.prev.Flip(index - 1)

		index++
		item.next.Flip(index)

		items = append(items, item)

		if next == -1 {
			item.next.Flip(-1)
			break
		}
	}

	//al.firstNode = items[0].index
	//al.lastNode = items[len(items)-1].index

	al.items = items
	al.resetDirty()
}

// Find returns giving node matching provided function.
func (al *slidingList) Find(fn func(*Node, int) bool) (*Node, error) {
	var count = len(al.items)
	if count == 0 {
		return nil, ErrNotFound
	}

	// run through the list of items
	if al.firstNode == nil && count > 1 {
		panic("Invalid pointer to first item, check logic")
	}

	var index int
	var next = al.firstNode
	if next == nil && al.lastNode != nil {
		panic("Invalid pointer state for slidingList.firstNode")
	}

	if next == nil {
		return nil, ErrNotFound
	}

	for next.Read() != -1 {
		item := al.items[next.Read()]
		if fn(item, index) {
			return item, nil
		}

		if item.next.Read() == -1 {
			break
		}

		index++
		next = item.next
	}
	return nil, ErrNotFound
}

// Each runs down the list of elements in array from first to last
// following not the arrangement in the slice but the links of next
// and previous for each node.
//
// It keeps walking down the sliding list node till either it reaches the
// node without a next index pointer or the function returns false.
func (al *slidingList) Each(fn func(*Node, int) bool) {
	var count = len(al.items)
	if count == 0 {
		return
	}

	// run through the list of items
	if al.firstNode == nil && count > 1 {
		panic("Invalid pointer to first item, check logic")
	}

	var index int
	var next = al.firstNode
	if next == nil && al.lastNode != nil {
		panic("Invalid pointer state for slidingList.firstNode")
	}

	if next == nil {
		return
	}

	for next.Read() != -1 {
		item := al.items[next.Read()]
		if !fn(item, index) {
			return
		}

		if item.next.Read() == -1 {
			return
		}

		index++
		next = item.next
	}
}

// EmptyIndex returns true/false if giving index is empty.
func (al *slidingList) EmptyIndex(index int) (bool, error) {
	if index < 0 || index >= len(al.items) {
		return false, ErrInvalidIndex
	}
	if len(al.items) == 0 {
		return false, ErrEmptyList
	}
	if al.items[index] == nil {
		return true, nil
	}
	return false, nil
}

// GetFirst returns the first node if any within list.
func (al *slidingList) GetFirst() (*Node, error) {
	return al.Get(0)
}

// GetLast returns the last node if any within list.
func (al *slidingList) GetLast() (*Node, error) {
	return al.Get(al.LastIndex())
}

// RemoveAndSwap removes a giving index and attempts to swap that
// index with the last element in the list to remove any blank spot.
func (al *slidingList) RemoveAndSwap(index int) (*Node, error) {
	node, err := al.RemoveIndex(index)
	if err != nil {
		return nil, err
	}

	if err = al.SwapIndex(index); err != nil && err != ErrEmptyList {
		return node, err
	}
	return node, nil
}

// RemoveIndex removes giving element at index from the array list,
// returning removed item, else an error if index is out of range
// or if index is empty.
//
// RemoveIndex creates a empty pocket when used and does not swap
// or maintain compactness, hence you must keep a manual history
// of removed indexes for swapping after your removals. This let's you
// optimize writes and maintain indexes for operations, which then are
// later swapped then sorted to maintain compactness.
func (al *slidingList) RemoveIndex(index int) (*Node, error) {
	var count = len(al.items)
	if index >= len(al.items) || index < 0 || count == 0 {
		return nil, ErrInvalidIndex
	}

	if item := al.items[index]; item != nil {
		al.items[index] = nil

		// if this is really the last item within slice
		// we need to slice down slice to avoid reference
		// nil slots at the end of slice.
		// We also need to use this to decide if swapping is
		// needed.
		if index == (count - 1) {
			al.items = al.items[:index]
		}

		if item.prev.Read() != -1 {
			pNode := al.items[item.prev.Read()]

			if item.next.Read() != -1 {
				nNode := al.items[item.next.Read()]
				pNode.next.Flip(nNode.index.Read())
				nNode.prev.Flip(pNode.index.Read())
			} else {
				pNode.next.Flip(-1)
			}

			if item.index == al.lastNode {
				al.lastNode = pNode.index
			}

			item.reset()
			return item, nil
		}

		if item.next.Read() != -1 {
			nNode := al.items[item.next.Read()]

			if item.prev.Read() != -1 {
				pNode := al.items[item.prev.Read()]
				pNode.next.Flip(nNode.index.Read())
				nNode.prev.Flip(pNode.index.Read())

				if item.index == al.lastNode {
					al.lastNode = pNode.index
				}
			} else {
				nNode.prev.Flip(-1)
			}

			if item.index == al.firstNode {
				al.firstNode = nNode.index
			}

			item.reset()
			return item, nil
		}

		item.reset()
		al.lastNode = nil
		al.firstNode = nil
		al.items = al.items[:0]
		return item, nil
	}
	return nil, ErrEmptyIndex
}

// SwapNode swaps giving node in index with provide node `m`. It will
// also swap kids list if `swapKids` is true.
// It returns an error if giving index is wrong
func (al *slidingList) SwapNode(index int, m *Node, swapKids bool) error {
	var count = len(al.items)
	if count == 0 {
		return ErrEmptyList
	}
	if index < 0 || index > count {
		return ErrInvalidOp
	}
	if index == count {
		return nil
	}

	var oldNode = al.items[index]

	m.next.Flip(oldNode.next.Read())
	m.prev.Flip(oldNode.prev.Read())
	m.index.Flip(oldNode.index.Read())

	if oldNode.index == al.firstNode {
		al.firstNode = m.index
	}

	if oldNode.index == al.lastNode {
		al.lastNode = m.index
	}

	if swapKids {
		var oldKids = oldNode.kids
		oldNode.kids = m.kids
		m.kids = oldKids
	}

	al.items[index] = m
	return nil
}

// SwapIndex swaps last item in the slice into the provided index shrinking
// the giving array by 1. It fails to perform the shrinking if the total element
// in the list is 1 or if list is empty.
func (al *slidingList) SwapIndex(index int) error {
	var count = len(al.items)
	if count == 0 {
		return ErrEmptyList
	}
	if index < 0 || index > count {
		return ErrInvalidOp
	}
	if index == count {
		return nil
	}

	if al.items[index] != nil {
		return ErrIndexNotEmpty
	}

	if count == 1 && al.items[0] == nil {
		al.lastNode = nil
		al.firstNode = nil
		al.items = al.items[:0]
		return nil
	}

	lastIndex := count - 1
	lastItem := al.items[lastIndex]

	// if we are nil, then we have open
	// pocket, close it and retry till you
	// find usable pocket.
	if lastItem == nil {
		al.items = al.items[:lastIndex]
		return al.SwapIndex(index)
	}

	al.items[index] = lastItem
	lastItem.index.Flip(index)
	if lastItem.prev.Read() != -1 {
		pNode := al.items[lastItem.prev.Read()]
		pNode.next.Flip(index)
	}
	if lastItem.next.Read() != -1 {
		nNode := al.items[lastItem.next.Read()]
		nNode.prev.Flip(index)
	}

	al.items = al.items[:lastIndex]
	al.setDirty()
	return nil
}

func (al *slidingList) expandList() {
	var nextCapacity = al.getNextCapacity()
	var newList = make([]*Node, nextCapacity)
	var copied = copy(newList, al.items)
	al.items = newList[:copied]
	al.resetLastExpansion()
}

// getNextCapacity returns a possible increasing expansion value
// based on a expansive but gradual increasing size
func (al *slidingList) getNextCapacity() int {
	var current = cap(al.items)
	if current == 0 {
		current = increasingFactor
	}
	var clen = len(al.items)
	if clen == 0 {
		clen = 1
	}

	var ld = al.getLastExpansion()
	if ld == 0 {
		ld = 1
	}

	var pb = (current * increasingFactor) * ld
	var inc = (pb / current) + increasingFactor
	return inc + (current / clen)
}

// resetLastExpansion resets the giving last addition value.
func (al *slidingList) resetLastExpansion() {
	atomic.StoreUint32(&al.lf, 0)
}

func (al *slidingList) resetDirty() {
	atomic.StoreUint32(&al.dirty, 0)
}

func (al *slidingList) setDirty() {
	atomic.StoreUint32(&al.dirty, 1)
}

func (al *slidingList) isDirty() bool {
	return atomic.LoadUint32(&al.dirty) == 1
}

// incrementLastExpansion increments the last addition value by count.
func (al *slidingList) incrementLastExpansion(n int) {
	atomic.AddUint32(&al.lf, uint32(n))
}

// getLastExpansion returns the current value of last addition value.
func (al *slidingList) getLastExpansion() int {
	return int(atomic.LoadUint32(&al.lf))
}

//**********************************************
// utils
//**********************************************

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
