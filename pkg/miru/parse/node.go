// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse nodes.

package parse

import (
	"fmt"
	"strconv"
	"strings"
)

var baseTypes = map[string]bool{
	"Uint":      true,
	"Any":       true,
	"any":       true,
	"Uint32":    true,
	"Uint64":    true,
	"Int":       true,
	"Int32":     true,
	"Int64":     true,
	"Time":      true,
	"Float":     true,
	"Float32":   true,
	"Float64":   true,
	"Complex64": true,
	"Rune":      true,
	"Byte":      true,
	"String":    true,
}

var baseTypesRepresentation = map[string]string{
	"Any":       "interface{}",
	"any":       "interface{}",
	"Time":      "time.Time",
	"time":      "time.Time",
	"String":    "string",
	"Uint":      "uint",
	"Uint32":    "uint32",
	"Uint64":    "uint64",
	"Int":       "int",
	"Int32":     "int32",
	"Int64":     "int64",
	"Float":     "float64",
	"Float32":   "float32",
	"Float64":   "float64",
	"Complex64": "complex64",
	"Rune":      "rune",
	"Byte":      "byte",
	"ByteSlice": "[]byte",
}

var textFormat = "%s" // Changed to "%q" in tests for better error messages.

type NodeTyper interface {
	Type() NodeType
}

// A Node is an element in the parse tree. The interface is trivial.
// The interface contains an unexported method so that only
// types local to this package can satisfy it.
type Node interface {
	NodeTyper

	String() string
	// Copy does a deep copy of the Node and all its components.
	// To avoid type assertions, some XxxNodes also have specialized
	// CopyXxx methods that return *XxxNode.
	Copy() Node
	Position() Pos // byte position of start of node in full original input string
	// tree returns the containing *Tree.
	// It is unexported so all implementations of Node are in this package.
	tree() *Tree
	// writeTo writes the String output to the builder.
	writeTo(*strings.Builder)
}

// NodeType identifies the type of a parse tree node.
type NodeType int

// Pos represents a byte position in the original input text from which
// this template was parsed.
type Pos int

func (p Pos) Position() Pos {
	return p
}

// Type returns itself and provides an easy default implementation
// for embedding in a Node. Embedded in all non-trivial Nodes.
func (t NodeType) Type() NodeType {
	return t
}

func (t NodeType) String() string {
	switch t {
	case NodeText:
		return "NodeText"
	case NodeAction:
		return "NodeAction"
	default:
		return "Unknown"
	}
}

const (
	NodeText             NodeType = iota // Plain text.
	NodeDOM                              // DOM identifer.
	NodeAction                           // A non-control action such as a field evaluation.
	NodeBool                             // A boolean constant.
	NodeChain                            // A sequence of field accesses.
	NodeCommand                          // An element of a pipeline.
	NodeDot                              // The cursor, dot.
	nodeElse                             // An else action. Not added to tree.
	nodeEnd                              // An end action. Not added to tree.
	nodeMethodEnd                        // An end action. Not added to tree.
	NodeField                            // A field or method name.
	NodeIdentifier                       // An identifier; always a function name.
	NodeTextIdentifier                   // An identifier; always a function name.
	NodeMethodIdentifier                 // An identifier; always a function name.
	NodeIf                               // An if action.
	NodeTag                              // A html tag .
	NodeTagEnd                           // A html ending tag .
	NodeImport                           // A import declaration
	NodeImportedFunc                     // A imported func declaration
	NodeImportedField                    // A imported field declaration
	NodeModel                            // A model declaration
	NodeModelType                        // A model type declaration
	NodeModelField                       // A model field declaration
	NodeModelTypeField                   // A model type field declaration
	NodeMethod                           // A method declaration
	NodeMethodArgument                   // A model agument declaration
	NodeAttr                             // A html attr node .
	NodeKids                             // A html attr node .
	NodeList                             // A list of Nodes.
	NodeNil                              // An untyped nil constant.
	NodeNumber                           // A numerical constant.
	NodePipe                             // A pipeline of commands.
	NodeRange                            // A range action.
	NodeString                           // A string constant.
	NodeRoot                             // A root node constant.
	NodeKomponent                        // A komponent node constant.
	NodeKomponentMount                   // A Mountkomponent node constant.
	NodeTemplate                         // A template invocation action.
	NodeVariable                         // A $ variable.
	NodeWith                             // A with action.
	nodeFinish                           // A finish signal
)

// Nodes.

// ListNode holds a sequence of nodes.
type ListNode struct {
	NodeType
	Pos
	tr          *Tree
	Nodes       []Node // The element nodes in lexical order.
	InHTML      bool
	HasDOMMount bool
}

func (t *Tree) newList(pos Pos) *ListNode {
	return &ListNode{tr: t, InHTML: t.withinHTML, NodeType: NodeList, Pos: pos}
}

func (l *ListNode) append(n Node) {
	l.Nodes = append(l.Nodes, n)
}

func (l *ListNode) tree() *Tree {
	return l.tr
}

func (l *ListNode) String() string {
	var sb strings.Builder
	l.writeTo(&sb)
	return sb.String()
}

func (l *ListNode) writeTo(sb *strings.Builder) {
	for _, n := range l.Nodes {
		n.writeTo(sb)
	}
}

func (l *ListNode) CopyList() *ListNode {
	if l == nil {
		return l
	}
	n := l.tr.newList(l.Pos)
	for _, elem := range l.Nodes {
		n.append(elem.Copy())
	}
	return n
}

func (l *ListNode) Copy() Node {
	return l.CopyList()
}

// TextNode holds plain text.
type TextNode struct {
	NodeType
	Pos
	tr   *Tree
	Text []byte // The text; may span newlines.
}

func (t *Tree) newText(pos Pos, text string) *TextNode {
	return &TextNode{tr: t, NodeType: NodeText, Pos: pos, Text: []byte(text)}
}

func (t *TextNode) String() string {
	return fmt.Sprintf(textFormat, t.Text)
}

func (t *TextNode) writeTo(sb *strings.Builder) {
	sb.WriteString(t.String())
}

func (t *TextNode) tree() *Tree {
	return t.tr
}

func (t *TextNode) Copy() Node {
	return &TextNode{tr: t.tr, NodeType: NodeText, Pos: t.Pos, Text: append([]byte{}, t.Text...)}
}

// DOMIdentifier holds plain text.
type DOMIdentifier struct {
	NodeType
	Pos
	tr   *Tree
	Text string
}

func (t *Tree) newDOMIdentifer(pos Pos, text string) *DOMIdentifier {
	return &DOMIdentifier{tr: t, NodeType: NodeDOM, Pos: pos, Text: text}
}

func (t *DOMIdentifier) String() string {
	return fmt.Sprintf(textFormat, t.Text)
}

func (t *DOMIdentifier) writeTo(sb *strings.Builder) {
	sb.WriteString(t.String())
}

func (t *DOMIdentifier) tree() *Tree {
	return t.tr
}

func (t *DOMIdentifier) Copy() Node {
	return &DOMIdentifier{tr: t.tr, NodeType: NodeTextIdentifier, Pos: t.Pos, Text: t.Text}
}

// TextIdentifierNode holds plain text.
type TextIdentifierNode struct {
	NodeType
	Pos
	tr   *Tree
	Text string
}

func (t *Tree) newTextIdentifer(pos Pos, text string) *TextIdentifierNode {
	return &TextIdentifierNode{tr: t, NodeType: NodeTextIdentifier, Pos: pos, Text: text}
}

func (t *TextIdentifierNode) String() string {
	return fmt.Sprintf(textFormat, t.Text)
}

func (t *TextIdentifierNode) writeTo(sb *strings.Builder) {
	sb.WriteString(t.String())
}

func (t *TextIdentifierNode) tree() *Tree {
	return t.tr
}

func (t *TextIdentifierNode) Copy() Node {
	return &TextIdentifierNode{tr: t.tr, NodeType: NodeTextIdentifier, Pos: t.Pos, Text: t.Text}
}

// PipeNode holds a pipeline with optional declaration
type PipeNode struct {
	isImported bool
	NodeType
	Pos
	tr          *Tree
	Line        int             // The line number in the input. Deprecated: Kept for compatibility.
	IsAssign    bool            // The variables are being assigned, not declared.
	Decl        []*VariableNode // Variables in lexical order.
	Cmds        []*CommandNode  // The commands in lexical order.
	HasDOMMount bool
}

func (t *Tree) newPipeline(pos Pos, line int, vars []*VariableNode) *PipeNode {
	return &PipeNode{tr: t, NodeType: NodePipe, Pos: pos, Line: line, Decl: vars}
}

func (p *PipeNode) append(command *CommandNode) {
	p.Cmds = append(p.Cmds, command)
}

func (p *PipeNode) String() string {
	var sb strings.Builder
	p.writeTo(&sb)
	return sb.String()
}

func (p *PipeNode) writeTo(sb *strings.Builder) {
	if len(p.Decl) > 0 {
		for i, v := range p.Decl {
			if i > 0 {
				sb.WriteString(", ")
			}
			v.writeTo(sb)
		}
		sb.WriteString(" := ")
	}
	for i, c := range p.Cmds {
		if i > 0 {
			sb.WriteString(" | ")
		}
		c.writeTo(sb)
	}
}

func (p *PipeNode) tree() *Tree {
	return p.tr
}

func (p *PipeNode) CopyPipe() *PipeNode {
	if p == nil {
		return p
	}
	vars := make([]*VariableNode, len(p.Decl))
	for i, d := range p.Decl {
		vars[i] = d.Copy().(*VariableNode)
	}
	n := p.tr.newPipeline(p.Pos, p.Line, vars)
	n.isImported = p.isImported
	n.IsAssign = p.IsAssign
	for _, c := range p.Cmds {
		n.append(c.Copy().(*CommandNode))
	}
	return n
}

func (p *PipeNode) Copy() Node {
	return p.CopyPipe()
}

// ActionNode holds an action (something bounded by delimiters).
// Control actions have their own nodes; ActionNode represents simple
// ones such as field evaluations and parenthesized pipelines.
type ActionNode struct {
	isImported bool
	NodeType
	Pos
	tr          *Tree
	Line        int       // The line number in the input. Deprecated: Kept for compatibility.
	Pipe        *PipeNode // The pipeline in the action.
	InHTML      bool
	HasDOMMount bool
}

func (t *Tree) newAction(pos Pos, line int, pipe *PipeNode) *ActionNode {
	return &ActionNode{tr: t, NodeType: NodeAction, Pos: pos, Line: line, Pipe: pipe, InHTML: t.withinHTML}
}

func (a *ActionNode) String() string {
	var sb strings.Builder
	a.writeTo(&sb)
	return sb.String()
}

func (a *ActionNode) writeTo(sb *strings.Builder) {
	sb.WriteString("{{")
	a.Pipe.writeTo(sb)
	sb.WriteString("}}")
}

func (a *ActionNode) tree() *Tree {
	return a.tr
}

func (a *ActionNode) Copy() Node {
	var ab = a.tr.newAction(a.Pos, a.Line, a.Pipe.CopyPipe())
	ab.isImported = a.isImported
	return ab
}

// CommandNode holds a command (a pipeline inside an evaluating action).
type CommandNode struct {
	isImported bool
	NodeType
	Pos
	tr          *Tree
	Line        int
	Args        []Node // Arguments in lexical order: Identifier, field, or constant.
	InHTML      bool
	HasDOMMount bool
}

func (t *Tree) newCommand(pos Pos) *CommandNode {
	return &CommandNode{tr: t, NodeType: NodeCommand, Pos: pos, InHTML: t.withinHTML}
}

func (c *CommandNode) append(arg Node) {
	c.Args = append(c.Args, arg)
}

func (c *CommandNode) String() string {
	var sb strings.Builder
	c.writeTo(&sb)
	return sb.String()
}

func (c *CommandNode) writeTo(sb *strings.Builder) {
	for i, arg := range c.Args {
		if i > 0 {
			sb.WriteByte(' ')
		}
		if arg, ok := arg.(*PipeNode); ok {
			sb.WriteByte('(')
			arg.writeTo(sb)
			sb.WriteByte(')')
			continue
		}
		arg.writeTo(sb)
	}
}

func (c *CommandNode) tree() *Tree {
	return c.tr
}

func (c *CommandNode) Copy() Node {
	if c == nil {
		return c
	}
	n := c.tr.newCommand(c.Pos)
	for _, c := range c.Args {
		n.append(c.Copy())
	}
	return n
}

// IdentifierNode holds an identifier.
type IdentifierNode struct {
	NodeType
	Pos
	tr            *Tree
	Ident         string // The identifier's name.
	IsInbuiltFunc bool
	InHTML        bool
}

// NewIdentifier returns a new IdentifierNode with the given identifier name.
func NewIdentifier(ident string) *IdentifierNode {
	return &IdentifierNode{NodeType: NodeIdentifier, Ident: ident}
}

// NewMethodIdentifier returns a new IdentifierNode with the given identifier name.
func NewMethodIdentifier(ident string) *IdentifierNode {
	return &IdentifierNode{NodeType: NodeMethodIdentifier, Ident: ident}
}

// SetPos sets the position. NewIdentifier is a public method so we can't modify its signature.
// Chained for convenience.
// TODO: fix one day?
func (i *IdentifierNode) SetPos(pos Pos) *IdentifierNode {
	i.Pos = pos
	return i
}

// SetTree sets the parent tree for the node. NewIdentifier is a public method so we can't modify its signature.
// Chained for convenience.
// TODO: fix one day?
func (i *IdentifierNode) SetTree(t *Tree) *IdentifierNode {
	i.tr = t
	return i
}

func (i *IdentifierNode) setIsInbuiltFunction(v bool) *IdentifierNode {
	i.IsInbuiltFunc = v
	return i
}

func (i *IdentifierNode) setInHTML(InHTML bool) {
	i.InHTML = InHTML
}

func (i *IdentifierNode) String() string {
	return i.Ident
}

func (i *IdentifierNode) writeTo(sb *strings.Builder) {
	sb.WriteString(i.String())
}

func (i *IdentifierNode) tree() *Tree {
	return i.tr
}

func (i *IdentifierNode) Copy() Node {
	return NewIdentifier(i.Ident).SetTree(i.tr).SetPos(i.Pos)
}

// AssignNode holds a list of variable names, possibly with chained field
// accesses. The dollar sign is part of the (first) name.
type VariableNode struct {
	NodeType
	Pos
	tr     *Tree
	Ident  []string // Variable name and fields in lexical order.
	InHTML bool
}

func (t *Tree) newVariable(pos Pos, ident string) *VariableNode {
	return &VariableNode{tr: t, NodeType: NodeVariable, Pos: pos, InHTML: t.withinHTML, Ident: strings.Split(ident, ".")}
}

func (v *VariableNode) String() string {
	var sb strings.Builder
	v.writeTo(&sb)
	return sb.String()
}

func (v *VariableNode) writeTo(sb *strings.Builder) {
	for i, id := range v.Ident {
		if i > 0 {
			sb.WriteByte('.')
		}
		sb.WriteString(id)
	}
}

func (v *VariableNode) tree() *Tree {
	return v.tr
}

func (v *VariableNode) Copy() Node {
	return &VariableNode{tr: v.tr, NodeType: NodeVariable, Pos: v.Pos, Ident: append([]string{}, v.Ident...)}
}

// DotNode holds the special identifier '.'.
type DotNode struct {
	NodeType
	Pos
	tr *Tree
}

func (t *Tree) newDot(pos Pos) *DotNode {
	return &DotNode{tr: t, NodeType: NodeDot, Pos: pos}
}

func (d *DotNode) Type() NodeType {
	// Override method on embedded NodeType for API compatibility.
	// TODO: Not really a problem; could change API without effect but
	// api tool complains.
	return NodeDot
}

func (d *DotNode) String() string {
	return "."
}

func (d *DotNode) writeTo(sb *strings.Builder) {
	sb.WriteString(d.String())
}

func (d *DotNode) tree() *Tree {
	return d.tr
}

func (d *DotNode) Copy() Node {
	return d.tr.newDot(d.Pos)
}

// FinishNode holds the special identifier 'nil' representing an untyped nil constant.
type FinishNode struct {
	NodeType
	Pos
	tr *Tree
}

func (t *Tree) newFinish(pos Pos) *FinishNode {
	return &FinishNode{tr: t, NodeType: nodeFinish, Pos: pos}
}

func (n *FinishNode) Type() NodeType {
	return nodeFinish
}

func (n *FinishNode) String() string {
	return "nil"
}

func (n *FinishNode) writeTo(sb *strings.Builder) {
	sb.WriteString(n.String())
}

func (n *FinishNode) tree() *Tree {
	return n.tr
}

func (n *FinishNode) Copy() Node {
	return n.tr.newNil(n.Pos)
}

// NilNode holds the special identifier 'nil' representing an untyped nil constant.
type NilNode struct {
	NodeType
	Pos
	tr *Tree
}

func (t *Tree) newNil(pos Pos) *NilNode {
	return &NilNode{tr: t, NodeType: NodeNil, Pos: pos}
}

func (n *NilNode) Type() NodeType {
	// Override method on embedded NodeType for API compatibility.
	// TODO: Not really a problem; could change API without effect but
	// api tool complains.
	return NodeNil
}

func (n *NilNode) String() string {
	return "nil"
}

func (n *NilNode) writeTo(sb *strings.Builder) {
	sb.WriteString(n.String())
}

func (n *NilNode) tree() *Tree {
	return n.tr
}

func (n *NilNode) Copy() Node {
	return n.tr.newNil(n.Pos)
}

// FieldNode holds a field (identifier starting with '.').
// The names may be chained ('.x.y').
// The period is dropped from each ident.
type FieldNode struct {
	NodeType
	Pos
	tr     *Tree
	Ident  []string // The identifiers in lexical order.
	InHTML bool
}

func (t *Tree) newField(pos Pos, ident string) *FieldNode {
	return &FieldNode{tr: t, InHTML: t.withinHTML, NodeType: NodeField, Pos: pos, Ident: strings.Split(ident[1:], ".")} // [1:] to drop leading period
}

func (f *FieldNode) String() string {
	var sb strings.Builder
	f.writeTo(&sb)
	return sb.String()
}

func (f *FieldNode) writeTo(sb *strings.Builder) {
	for _, id := range f.Ident {
		sb.WriteByte('.')
		sb.WriteString(id)
	}
}

func (f *FieldNode) tree() *Tree {
	return f.tr
}

func (f *FieldNode) Copy() Node {
	return &FieldNode{tr: f.tr, NodeType: NodeField, Pos: f.Pos, Ident: append([]string{}, f.Ident...)}
}

// ChainNode holds a term followed by a chain of field accesses (identifier starting with '.').
// The names may be chained ('.x.y').
// The periods are dropped from each ident.
type ChainNode struct {
	NodeType
	Pos
	tr     *Tree
	Node   Node
	Field  []string // The identifiers in lexical order.
	InHTML bool
}

func (t *Tree) newChain(pos Pos, node Node) *ChainNode {
	return &ChainNode{tr: t, InHTML: t.withinHTML, NodeType: NodeChain, Pos: pos, Node: node}
}

// Add adds the named field (which should start with a period) to the end of the chain.
func (c *ChainNode) Add(field string) {
	if len(field) == 0 || field[0] != '.' {
		panic("no dot in field")
	}
	field = field[1:] // Remove leading dot.
	if field == "" {
		panic("empty field")
	}
	c.Field = append(c.Field, field)
}

func (c *ChainNode) String() string {
	var sb strings.Builder
	c.writeTo(&sb)
	return sb.String()
}

func (c *ChainNode) writeTo(sb *strings.Builder) {
	if _, ok := c.Node.(*PipeNode); ok {
		sb.WriteByte('(')
		c.Node.writeTo(sb)
		sb.WriteByte(')')
	} else {
		c.Node.writeTo(sb)
	}
	for _, field := range c.Field {
		sb.WriteByte('.')
		sb.WriteString(field)
	}
}

func (c *ChainNode) tree() *Tree {
	return c.tr
}

func (c *ChainNode) Copy() Node {
	return &ChainNode{tr: c.tr, NodeType: NodeChain, Pos: c.Pos, Node: c.Node, Field: append([]string{}, c.Field...)}
}

// RefNode holds a reference to a type constant.
type RefNode struct {
	NodeType
	Pos
	Name   string
	Typ    string
	Detail Node
	tr     *Tree
}

func (t *Tree) newRefNode(pos Pos, name string, typ string, detail Node) *RefNode {
	return &RefNode{tr: t, NodeType: NodeModel, Name: name, Typ: typ, Detail: detail, Pos: pos}
}

func (b *RefNode) String() string {
	return b.Name
}

func (b *RefNode) writeTo(sb *strings.Builder) {
	sb.WriteString(b.String())
}

func (b *RefNode) tree() *Tree {
	return b.tr
}

func (b *RefNode) Copy() Node {
	return b.tr.newRefNode(b.Pos, b.Name, b.Typ, b.Detail)
}

// MethodNode holds a boolean constant.
type MethodNode struct {
	NodeType
	Pos
	Model     string
	tr        *Tree
	Arguments []*RefNode
	Body      *ListNode
}

func (t *Tree) newMethod(pos Pos, model string, arguments []*RefNode, body *ListNode) *MethodNode {
	return &MethodNode{tr: t, NodeType: NodeMethod, Model: model, Pos: pos, Arguments: arguments, Body: body}
}

func (b *MethodNode) appendArgument(ref *RefNode) {
	b.Arguments = append(b.Arguments, ref)
}

func (b *MethodNode) HasArgument(argName string) bool {
	for _, ref := range b.Arguments {
		if ref.Name == argName {
			return true
		}
	}
	return false
}

func (b *MethodNode) String() string {
	return b.Model
}

func (b *MethodNode) writeTo(sb *strings.Builder) {
	sb.WriteString(b.String())
}

func (b *MethodNode) tree() *Tree {
	return b.tr
}

func (b *MethodNode) Copy() Node {
	return b.tr.newMethod(b.Pos, b.Model, b.Arguments, b.Body)
}

// BaseTypeNode holds a boolean constant.
type BaseTypeNode struct {
	NodeType
	Typ string
	Pos
	tr     *Tree
	InHTML bool
}

func (t *Tree) newBase(pos Pos, typ string) *BaseTypeNode {
	return &BaseTypeNode{tr: t, InHTML: t.withinHTML, NodeType: NodeModel, Typ: typ, Pos: pos}
}

func (b *BaseTypeNode) String() string {
	return b.Typ
}

func (b *BaseTypeNode) writeTo(sb *strings.Builder) {
	sb.WriteString(b.String())
}

func (b *BaseTypeNode) tree() *Tree {
	return b.tr
}

func (b *BaseTypeNode) Copy() Node {
	return b.tr.newBase(b.Pos, b.Typ)
}

type CollectionType int

const (
	ListType CollectionType = iota
	MapType
)

var collectionType = map[string]CollectionType{
	"List": ListType,
	"Map":  MapType,
}

var collectionTypeName = map[CollectionType]string{
	ListType: "List",
	MapType:  "Map",
}

// CollectionTypeNode holds a boolean constant.
type CollectionTypeNode struct {
	NodeType
	Typ    CollectionType
	Detail *ListNode
	Pos
	tr *Tree
}

func (t *Tree) newCollection(pos Pos, typ CollectionType, detail *ListNode) *CollectionTypeNode {
	return &CollectionTypeNode{tr: t, NodeType: NodeModel, Detail: detail, Typ: typ, Pos: pos}
}

func (b *CollectionTypeNode) String() string {
	return collectionTypeName[b.Typ]
}

func (b *CollectionTypeNode) writeTo(sb *strings.Builder) {
	sb.WriteString(b.String())
}

func (b *CollectionTypeNode) tree() *Tree {
	return b.tr
}

func (b *CollectionTypeNode) Copy() Node {
	return b.tr.newCollection(b.Pos, b.Typ, b.Detail)
}

// ModelTypeNode holds a boolean constant.
type ModelTypeNode struct {
	NodeType
	Typ    string
	Detail Node
	Pos
	tr *Tree
}

func (t *Tree) newModelType(pos Pos, model string, detail Node) *ModelTypeNode {
	return &ModelTypeNode{tr: t, NodeType: NodeModel, Typ: model, Pos: pos, Detail: detail}
}

func (b *ModelTypeNode) String() string {
	return b.Typ
}

func (b *ModelTypeNode) writeTo(sb *strings.Builder) {
	sb.WriteString(b.String())
}

func (b *ModelTypeNode) tree() *Tree {
	return b.tr
}

func (b *ModelTypeNode) Copy() Node {
	return b.tr.newModelType(b.Pos, b.Typ, b.Detail)
}

// ModelNode holds a boolean constant.
type ModelNode struct {
	NodeType
	Model string
	Pos
	tr     *Tree
	Fields *ListNode
}

func (t *Tree) newModel(pos Pos, model string, fields *ListNode) *ModelNode {
	return &ModelNode{tr: t, NodeType: NodeModel, Model: model, Pos: pos, Fields: fields}
}

func (b *ModelNode) String() string {
	return b.Model
}

func (b *ModelNode) writeTo(sb *strings.Builder) {
	sb.WriteString(b.String())
}

func (b *ModelNode) tree() *Tree {
	return b.tr
}

func (b *ModelNode) Copy() Node {
	return b.tr.newModel(b.Pos, b.Model, b.Fields)
}

// BoolNode holds a boolean constant.
type BoolNode struct {
	NodeType
	Pos
	tr     *Tree
	True   bool // The value of the boolean constant.
	InHTML bool
}

func (t *Tree) newBool(pos Pos, true bool) *BoolNode {
	return &BoolNode{tr: t, InHTML: t.withinHTML, NodeType: NodeBool, Pos: pos, True: true}
}

func (b *BoolNode) String() string {
	if b.True {
		return "true"
	}
	return "false"
}

func (b *BoolNode) writeTo(sb *strings.Builder) {
	sb.WriteString(b.String())
}

func (b *BoolNode) tree() *Tree {
	return b.tr
}

func (b *BoolNode) Copy() Node {
	return b.tr.newBool(b.Pos, b.True)
}

// NumberNode holds a number: signed or unsigned integer, float, or complex.
// The value is parsed and stored under all the types that can represent the value.
// This simulates in a small amount of code the behavior of Go's ideal constants.
type NumberNode struct {
	NodeType
	Pos
	tr         *Tree
	IsInt      bool       // Number has an integral value.
	IsUint     bool       // Number has an unsigned integral value.
	IsFloat    bool       // Number has a floating-point value.
	IsComplex  bool       // Number is complex.
	Int64      int64      // The signed integer value.
	Uint64     uint64     // The unsigned integer value.
	Float64    float64    // The floating-point value.
	Complex128 complex128 // The complex value.
	Text       string     // The original textual representation from the input.
	InHTML     bool
}

func (t *Tree) newNumber(pos Pos, text string, typ itemType) (*NumberNode, error) {
	n := &NumberNode{tr: t, InHTML: t.withinHTML, NodeType: NodeNumber, Pos: pos, Text: text}
	switch typ {
	case itemCharConstant:
		rune, _, tail, err := strconv.UnquoteChar(text[1:], text[0])
		if err != nil {
			return nil, err
		}
		if tail != "'" {
			return nil, fmt.Errorf("malformed character constant: %s", text)
		}
		n.Int64 = int64(rune)
		n.IsInt = true
		n.Uint64 = uint64(rune)
		n.IsUint = true
		n.Float64 = float64(rune) // odd but those are the rules.
		n.IsFloat = true
		return n, nil
	case itemComplex:
		// fmt.Sscan can parse the pair, so let it do the work.
		if _, err := fmt.Sscan(text, &n.Complex128); err != nil {
			return nil, err
		}
		n.IsComplex = true
		n.simplifyComplex()
		return n, nil
	}
	// Imaginary constants can only be complex unless they are zero.
	if len(text) > 0 && text[len(text)-1] == 'i' {
		f, err := strconv.ParseFloat(text[:len(text)-1], 64)
		if err == nil {
			n.IsComplex = true
			n.Complex128 = complex(0, f)
			n.simplifyComplex()
			return n, nil
		}
	}
	// Do integer test first so we get 0x123 etc.
	u, err := strconv.ParseUint(text, 0, 64) // will fail for -0; fixed below.
	if err == nil {
		n.IsUint = true
		n.Uint64 = u
	}
	i, err := strconv.ParseInt(text, 0, 64)
	if err == nil {
		n.IsInt = true
		n.Int64 = i
		if i == 0 {
			n.IsUint = true // in case of -0.
			n.Uint64 = u
		}
	}
	// If an integer extraction succeeded, promote the float.
	if n.IsInt {
		n.IsFloat = true
		n.Float64 = float64(n.Int64)
	} else if n.IsUint {
		n.IsFloat = true
		n.Float64 = float64(n.Uint64)
	} else {
		f, err := strconv.ParseFloat(text, 64)
		if err == nil {
			// If we parsed it as a float but it looks like an integer,
			// it's a huge number too large to fit in an int. Reject it.
			if !strings.ContainsAny(text, ".eEpP") {
				return nil, fmt.Errorf("integer overflow: %q", text)
			}
			n.IsFloat = true
			n.Float64 = f
			// If a floating-point extraction succeeded, extract the int if needed.
			if !n.IsInt && float64(int64(f)) == f {
				n.IsInt = true
				n.Int64 = int64(f)
			}
			if !n.IsUint && float64(uint64(f)) == f {
				n.IsUint = true
				n.Uint64 = uint64(f)
			}
		}
	}
	if !n.IsInt && !n.IsUint && !n.IsFloat {
		return nil, fmt.Errorf("illegal number syntax: %q", text)
	}
	return n, nil
}

// simplifyComplex pulls out any other types that are represented by the complex number.
// These all require that the imaginary part be zero.
func (n *NumberNode) simplifyComplex() {
	n.IsFloat = imag(n.Complex128) == 0
	if n.IsFloat {
		n.Float64 = real(n.Complex128)
		n.IsInt = float64(int64(n.Float64)) == n.Float64
		if n.IsInt {
			n.Int64 = int64(n.Float64)
		}
		n.IsUint = float64(uint64(n.Float64)) == n.Float64
		if n.IsUint {
			n.Uint64 = uint64(n.Float64)
		}
	}
}

func (n *NumberNode) String() string {
	return n.Text
}

func (n *NumberNode) writeTo(sb *strings.Builder) {
	sb.WriteString(n.String())
}

func (n *NumberNode) tree() *Tree {
	return n.tr
}

func (n *NumberNode) Copy() Node {
	nn := new(NumberNode)
	*nn = *n // Easy, fast, correct.
	return nn
}

// AttrNode holds a string constant. The value has been "unquoted".
type AttrNode struct {
	NodeType
	Pos
	tr     *Tree
	Quoted string // The original text of the string, with quotes.
	Key    string // The string, after quote processing.
	Values *ListNode
}

func (t *Tree) newAttr(pos Pos, orig, key string, v *ListNode) *AttrNode {
	return &AttrNode{tr: t, NodeType: NodeString, Pos: pos, Quoted: orig, Key: key, Values: v}
}

func (s *AttrNode) String() string {
	return s.Quoted
}

func (s *AttrNode) writeTo(sb *strings.Builder) {
	sb.WriteString(s.String())
}

func (s *AttrNode) tree() *Tree {
	return s.tr
}

func (s *AttrNode) Copy() Node {
	return s.tr.newAttr(s.Pos, s.Quoted, s.Key, s.Values)
}

// MountLiveNode holds a string constant. The value has been "unquoted".
type MountLiveNode struct {
	NodeType
	Pos
	Target Node
	IsList bool
	tr     *Tree
	Name   string // The string, after quote processing.
	Route  string
}

func (t *Tree) newMountLiveNode(pos Pos, route string, name string, target Node, IsList bool) *MountLiveNode {
	return &MountLiveNode{tr: t, NodeType: NodeKomponentMount, Pos: pos, IsList: IsList, Target: target, Route: route, Name: name}
}

func (s *MountLiveNode) String() string {
	return s.Name
}

func (s *MountLiveNode) writeTo(sb *strings.Builder) {
	sb.WriteString(s.String())
}

func (s *MountLiveNode) tree() *Tree {
	return s.tr
}

func (s *MountLiveNode) Copy() Node {
	return s.tr.newMountLiveNode(s.Pos, s.Route, s.Name, s.Target, s.IsList)
}

// MountNode holds a string constant. The value has been "unquoted".
type MountNode struct {
	NodeType
	Pos
	Target Node
	IsList bool
	tr     *Tree
	Name   string // The string, after quote processing.
}

func (t *Tree) newMountNode(pos Pos, text string, target Node, IsList bool) *MountNode {
	return &MountNode{tr: t, NodeType: NodeKomponentMount, Pos: pos, IsList: IsList, Target: target, Name: text}
}

func (s *MountNode) String() string {
	return s.Name
}

func (s *MountNode) writeTo(sb *strings.Builder) {
	sb.WriteString(s.String())
}

func (s *MountNode) tree() *Tree {
	return s.tr
}

func (s *MountNode) Copy() Node {
	return s.tr.newMountNode(s.Pos, s.Name, s.Target, s.IsList)
}

// RootNode holds a string constant. The value has been "unquoted".
type RootNode struct {
	NodeType
	Pos
	tr     *Tree
	Quoted string // The original text of the string, with quotes.
	Text   string // The string, after quote processing.
}

func (t *Tree) newRootNode(pos Pos, orig, text string) *RootNode {
	return &RootNode{tr: t, NodeType: NodeRoot, Pos: pos, Quoted: orig, Text: text}
}

func (s *RootNode) String() string {
	return s.Quoted
}

func (s *RootNode) writeTo(sb *strings.Builder) {
	sb.WriteString(s.String())
}

func (s *RootNode) tree() *Tree {
	return s.tr
}

func (s *RootNode) Copy() Node {
	return s.tr.newRootNode(s.Pos, s.Quoted, s.Text)
}

// ImportDeclr holds a string constant. The value has been "unquoted".
type ImportDeclr struct {
	NodeType
	Pos
	tr    *Tree
	Alias string
	Path  string // The string, after quote processing.
}

func (t *Tree) newImportDeclr(pos Pos, alias string, path string) *ImportDeclr {
	return &ImportDeclr{tr: t, NodeType: NodeString, Pos: pos, Alias: alias, Path: path}
}

func (s *ImportDeclr) String() string {
	return s.Alias + ` "` + s.Path + `"`
}

func (s *ImportDeclr) writeTo(sb *strings.Builder) {
	sb.WriteString(s.String())
}

func (s *ImportDeclr) tree() *Tree {
	return s.tr
}

func (s *ImportDeclr) Copy() Node {
	return s.tr.newImportDeclr(s.Pos, s.Alias, s.Path)
}

// ImportedFieldFunc holds a string constant. The value has been "unquoted".
type ImportedFieldFunc struct {
	NodeType
	Pos
	tr          *Tree
	Text        string // The string, after quote processing.
	AsArgument  bool
	Fields      []Node
	Commands    []*CommandNode
	InHTML      bool
	HasDOMMount bool
	IsNoOp      bool
}

func (t *Tree) newImportedField(pos Pos, text string) *ImportedFieldFunc {
	return &ImportedFieldFunc{tr: t, InHTML: t.withinHTML, NodeType: NodeImportedField, Pos: pos, Text: text}
}

func (t *Tree) newImportedFunc(pos Pos, text string) *ImportedFieldFunc {
	return &ImportedFieldFunc{tr: t, InHTML: t.withinHTML, NodeType: NodeImportedFunc, Pos: pos, Text: text}
}

func (s *ImportedFieldFunc) ToTextIdentifier() *TextIdentifierNode {
	return &TextIdentifierNode{
		NodeType: NodeIdentifier,
		Text:     s.ToName(),
		Pos:      s.Pos,
		tr:       s.tr,
	}
}

func (s *ImportedFieldFunc) ToIdentifier() *IdentifierNode {
	return &IdentifierNode{
		NodeType: NodeIdentifier,
		Ident:    s.ToName(),
		Pos:      s.Pos,
		tr:       s.tr,
	}
}

func (s *ImportedFieldFunc) ToName() string {
	if s.IsField() {
		return strings.TrimSpace(strings.Trim(s.Text, "@"))
	}
	return strings.TrimSpace(strings.Trim(s.Text, "#"))
}

func (s *ImportedFieldFunc) IsFunc() bool {
	return s.NodeType == NodeImportedFunc
}

func (s *ImportedFieldFunc) IsField() bool {
	return s.NodeType == NodeImportedField
}

func (s *ImportedFieldFunc) String() string {
	return s.Text
}

func (s *ImportedFieldFunc) writeTo(sb *strings.Builder) {
	sb.WriteString(s.String())
}

func (s *ImportedFieldFunc) tree() *Tree {
	return s.tr
}

func (s *ImportedFieldFunc) Copy() Node {
	if s.NodeType == NodeImportedField {
		return s.tr.newImportedField(s.Pos, s.Text)
	}
	return s.tr.newImportedFunc(s.Pos, s.Text)
}

// StringNode holds a string constant. The value has been "unquoted".
type StringNode struct {
	NodeType
	Pos
	tr     *Tree
	Quoted string // The original text of the string, with quotes.
	Text   string // The string, after quote processing.
	InHTML bool
}

func (t *Tree) newString(pos Pos, orig, text string) *StringNode {
	return &StringNode{tr: t, InHTML: t.withinHTML, NodeType: NodeString, Pos: pos, Quoted: orig, Text: text}
}

func (s *StringNode) String() string {
	return s.Quoted
}

func (s *StringNode) writeTo(sb *strings.Builder) {
	sb.WriteString(s.String())
}

func (s *StringNode) tree() *Tree {
	return s.tr
}

func (s *StringNode) Copy() Node {
	return s.tr.newString(s.Pos, s.Quoted, s.Text)
}

// endNode represents an {{end}} action.
// It does not appear in the final parse tree.
type endNode struct {
	NodeType
	Pos
	tr *Tree
}

func (t *Tree) newEnd(pos Pos) *endNode {
	return &endNode{tr: t, NodeType: nodeEnd, Pos: pos}
}

func (t *Tree) newMethodEnd(pos Pos) *endNode {
	return &endNode{tr: t, NodeType: nodeMethodEnd, Pos: pos}
}

func (e *endNode) String() string {
	return "{{end}}"
}

func (e *endNode) writeTo(sb *strings.Builder) {
	sb.WriteString(e.String())
}

func (e *endNode) tree() *Tree {
	return e.tr
}

func (e *endNode) Copy() Node {
	return e.tr.newEnd(e.Pos)
}

// elseNode represents an {{else}} action. Does not appear in the final tree.
type elseNode struct {
	NodeType
	Pos
	tr   *Tree
	Line int // The line number in the input. Deprecated: Kept for compatibility.
}

func (t *Tree) newElse(pos Pos, line int) *elseNode {
	return &elseNode{tr: t, NodeType: nodeElse, Pos: pos, Line: line}
}

func (e *elseNode) Type() NodeType {
	return nodeElse
}

func (e *elseNode) String() string {
	return "{{else}}"
}

func (e *elseNode) writeTo(sb *strings.Builder) {
	sb.WriteString(e.String())
}

func (e *elseNode) tree() *Tree {
	return e.tr
}

func (e *elseNode) Copy() Node {
	return e.tr.newElse(e.Pos, e.Line)
}

// HTMLNode is the common representation of if, range, and with.
type HTMLNode struct {
	NodeType
	Pos
	tr       *Tree
	Tag      string
	Closing  bool
	Line     int       // The line number in the input. Deprecated: Kept for compatibility.
	Attr     AttrList  // The pipeline to be evaluated.
	Children *ListNode // What are the children nodes for this.
}

type AttrList []*AttrNode

func (at *AttrList) append(n *AttrNode) {
	*at = append(*at, n)
}

func (at *AttrList) WriteTo(sb *strings.Builder) {
	for _, attr := range *at {
		attr.writeTo(sb)
	}
}

func (b *HTMLNode) String() string {
	var sb strings.Builder
	b.writeTo(&sb)
	return sb.String()
}

func (b *HTMLNode) writeTo(sb *strings.Builder) {
	sb.WriteString("<")
	sb.WriteString(b.Tag)
	sb.WriteByte(' ')
	b.Attr.WriteTo(sb)
	sb.WriteString(">")

	if !b.Closing {
		sb.WriteString("/>")
		return
	}

	if b.Children != nil {
		b.Children.writeTo(sb)
	}
	sb.WriteString("</")
	sb.WriteString(b.Tag)
	sb.WriteString(">")
}

func (b *HTMLNode) tree() *Tree {
	return b.tr
}

func (b *HTMLNode) Copy() Node {
	return b.tr.newHTMLTag(b.Pos, b.Line, b.Tag, b.Attr, b.Children)
}

// BranchNode is the common representation of if, range, and with.
type BranchNode struct {
	NodeType
	Pos
	tr       *Tree
	Line     int       // The line number in the input. Deprecated: Kept for compatibility.
	Pipe     *PipeNode // The pipeline to be evaluated.
	List     *ListNode // What to execute if the value is non-empty.
	ElseList *ListNode // What to execute if the value is empty (nil if absent).
}

func (b *BranchNode) String() string {
	var sb strings.Builder
	b.writeTo(&sb)
	return sb.String()
}

func (b *BranchNode) writeTo(sb *strings.Builder) {
	name := ""
	switch b.NodeType {
	case NodeIf:
		name = "if"
	case NodeRange:
		name = "range"
	case NodeWith:
		name = "with"
	default:
		panic("unknown branch type")
	}
	sb.WriteString("{{")
	sb.WriteString(name)
	sb.WriteByte(' ')
	b.Pipe.writeTo(sb)
	sb.WriteString("}}")
	b.List.writeTo(sb)
	if b.ElseList != nil {
		sb.WriteString("{{else}}")
		b.ElseList.writeTo(sb)
	}
	sb.WriteString("{{end}}")
}

func (b *BranchNode) tree() *Tree {
	return b.tr
}

func (b *BranchNode) Copy() Node {
	switch b.NodeType {
	case NodeIf:
		return b.tr.newIf(b.Pos, b.Line, b.Pipe, b.List, b.ElseList)
	case NodeRange:
		return b.tr.newRange(b.Pos, b.Line, b.Pipe, b.List, b.ElseList)
	case NodeWith:
		return b.tr.newWith(b.Pos, b.Line, b.Pipe, b.List, b.ElseList)
	default:
		panic("unknown branch type")
	}
}

// IfNode represents an {{if}} action and its commands.
type IfNode struct {
	BranchNode
}

func (t *Tree) newIf(pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) *IfNode {
	return &IfNode{BranchNode{tr: t, NodeType: NodeIf, Pos: pos, Line: line, Pipe: pipe, List: list, ElseList: elseList}}
}

func (t *Tree) newHTMLTag(pos Pos, line int, tag string, attr AttrList, kids *ListNode) *HTMLNode {
	return &HTMLNode{tr: t, NodeType: NodeIf, Pos: pos, Tag: tag, Line: line, Attr: attr, Children: kids}
}

func (i *IfNode) Copy() Node {
	return i.tr.newIf(i.Pos, i.Line, i.Pipe.CopyPipe(), i.List.CopyList(), i.ElseList.CopyList())
}

// RangeNode represents a {{range}} action and its commands.
type RangeNode struct {
	BranchNode
}

func (t *Tree) newRange(pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) *RangeNode {
	return &RangeNode{BranchNode{tr: t, NodeType: NodeRange, Pos: pos, Line: line, Pipe: pipe, List: list, ElseList: elseList}}
}

func (r *RangeNode) Copy() Node {
	return r.tr.newRange(r.Pos, r.Line, r.Pipe.CopyPipe(), r.List.CopyList(), r.ElseList.CopyList())
}

// WithNode represents a {{with}} action and its commands.
type WithNode struct {
	BranchNode
}

func (t *Tree) newWith(pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) *WithNode {
	return &WithNode{BranchNode{tr: t, NodeType: NodeWith, Pos: pos, Line: line, Pipe: pipe, List: list, ElseList: elseList}}
}

func (w *WithNode) Copy() Node {
	return w.tr.newWith(w.Pos, w.Line, w.Pipe.CopyPipe(), w.List.CopyList(), w.ElseList.CopyList())
}

// TemplateNode represents a {{template}} action.
type TemplateNode struct {
	NodeType
	Pos
	tr   *Tree
	Line int       // The line number in the input. Deprecated: Kept for compatibility.
	Name string    // The name of the template (unquoted).
	Pipe *PipeNode // The command to evaluate as dot for the template.
}

func (t *Tree) newTemplate(pos Pos, line int, name string, pipe *PipeNode) *TemplateNode {
	return &TemplateNode{tr: t, NodeType: NodeTemplate, Pos: pos, Line: line, Name: name, Pipe: pipe}
}

func (t *TemplateNode) String() string {
	var sb strings.Builder
	t.writeTo(&sb)
	return sb.String()
}

func (t *TemplateNode) writeTo(sb *strings.Builder) {
	sb.WriteString("{{template ")
	sb.WriteString(strconv.Quote(t.Name))
	if t.Pipe != nil {
		sb.WriteByte(' ')
		t.Pipe.writeTo(sb)
	}
	sb.WriteString("}}")
}

func (t *TemplateNode) tree() *Tree {
	return t.tr
}

func (t *TemplateNode) Copy() Node {
	return t.tr.newTemplate(t.Pos, t.Line, t.Name, t.Pipe.CopyPipe())
}
