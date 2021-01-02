package miru

import (
	"fmt"
	"go/format"
	"regexp"
	"strconv"
	"strings"

	"github.com/influx6/npkg/nerror"
	"github.com/influx6/npkg/nunsafe"

	"github.com/influx6/groundlayer/pkg/miru/parse"
)

const (
	DefaultLeftDelimiter  = "{{"
	DefaultRightDelimiter = "}}"

	baseParentCtx      = "ctx"
	baseNodeCtx        = "rootDoc"
	attachNodeFuncName = "attachToNode"

	helperAlias        = "helpers"
	helperImportPath   = "github.com/influx6/groundlayer/pkg/helpers"
	komponentGroupName = "group"
)

var (
	DefaultOption = Options{
		LeftDelimiter:  DefaultLeftDelimiter,
		RightDelimiter: DefaultRightDelimiter,
		Packages:       nil,
	}

	rootVar     = singleVar{Name: baseParentCtx}
	htmlRootVar = singleVar{Name: baseNodeCtx}

	allLineSpacesTabs = regexp.MustCompile(`[\n\f\r\t\[\]]+`)
)

var builtinFunctions = map[string]string{
	// basic functions
	"noop":     "noop",
	"and":      "helpers.And",
	"index":    "helpers.Index",
	"slice":    "helpers.Slice",
	"len":      "helpers.Len",
	"cap":      "helpers.Cap",
	"not":      "helpers.Not",
	"or":       "helpers.Or",
	"print":    "helpers.Print",
	"printf":   "helpers.Printf",
	"println":  "helpers.Println",
	"urlquery": "helpers.URLQuery",

	// Comparisons
	"eq": "helpers.Equal",              // ==
	"ge": "helpers.GreaterThanEqualTo", // >=
	"gt": "helpers.GreaterThan",        // >
	"le": "helpers.LessThanEqualTo",    // >=
	"lt": "helpers.LessThan",           // <
	"ne": "helpers.NotEqual",           // !=

	// internal functions we use
	"isInvalid":    "helpers.IsInvalid",
	"attachToNode": "helpers.AttachToNode",
}

// Options contains config fields for the TextWriter.
type Options struct {
	LeftDelimiter  string
	RightDelimiter string

	// Packages contain packages we wish added
	// into import path. It provides an alternative
	// to one declared directly in the template using
	// the {{import ... }} directive.
	//
	// The key is the alias and value the package path which
	// is to be imported.
	Packages map[string]string

	lastVariableCount  int
	lastKomponentCount int
}

type Treeset map[string]*parse.Tree

func New(text string, functions []map[string]string, sf SourceFileSystem) (*ParsedData, error) {
	return Parse("root", text, DefaultOption, functions, sf, nil)
}

// Parse parses provided string into underline collector.
//
// @param(require): name represent the special name to represent this parsed tree
// @param(required): text contains the content to be parsed
// @param(required): functions is the list of maps defining actions that should be considered functions
// @param(required): sf is the root file system to be used when retrieving template blocks for parsing
// @param(optional): tset is the previous Treeset from a previous run where this is considered a child of.
//
func Parse(name string, text string, ops Options, funcNames []map[string]string, sf SourceFileSystem, tset Treeset) (*ParsedData, error) {
	return ParseTree(name, text, ops, parse.MergeMaps(funcNames...), sf, tset)
}

func ParseTree(name string, text string, ops Options, fset map[string]string, sf SourceFileSystem, tset Treeset) (*ParsedData, error) {
	var mmap = parse.FuncMaps{
		Funcs:         parse.MergeMaps(builtinFunctions, fset),
		MethodSet:     map[string]*parse.MethodNode{},
		ModelTypeSet:  map[string]*parse.ModelTypeNode{},
		ModelSet:      map[string]*parse.ModelNode{},
		DefinedBlocks: map[string]string{},
	}

	return ParseTreeWithFuncMaps(name, text, ops, mmap, sf, tset)
}

func ParseTreeWithFuncMaps(name string, text string, ops Options, maps parse.FuncMaps, sf SourceFileSystem, set Treeset) (*ParsedData, error) {
	if ops.Packages == nil {
		ops.Packages = map[string]string{}
	}

	var treeSet, parseErr = parse.Parse(name, text, ops.LeftDelimiter, ops.RightDelimiter, maps, set)
	if parseErr != nil {
		return nil, nerror.WrapOnly(parseErr)
	}

	// parse out the root first.
	var root = treeSet[name]

	var builder = new(strings.Builder)
	var funcBuilder = new(strings.Builder)
	var typeBuilder = new(strings.Builder)

	var tw = newTextWriter(builder, funcBuilder, typeBuilder, root, ops, maps, treeSet, sf, rootVar, htmlRootVar)

	var parseTreeErr = tw.ParseTree()
	if parseTreeErr != nil {
		return nil, parseTreeErr
	}

	var parsed = newParsedData(root.TypeName, typeBuilder, funcBuilder, builder, tw.rootNode, ops)
	return parsed, nil
}

// NewTextWriter returns a new instance of a TextWriter using provided
// builder.
func newTextWriter(
	builder *strings.Builder,
	funcBuilder *strings.Builder,
	typeBuilder *strings.Builder,
	tree *parse.Tree,
	ops Options,
	funcMaps parse.FuncMaps,
	ts Treeset,
	sf SourceFileSystem,
	varRoot vard,
	htmlRoot vard,
) *TextWriter {
	return &TextWriter{
		base:        tree,
		options:     ops,
		TreeSet:     ts,
		funcMaps:    funcMaps,
		typeBuilder: typeBuilder,
		builder:     builder,
		funcBuilder: funcBuilder,
		filesystem:  sf,
		rootQueue: &varQueue{
			queue: []vard{varRoot},
		},
		htmlQueue: &varQueue{
			queue: []vard{htmlRoot},
		},
		variables:  ops.lastVariableCount,
		komponents: ops.lastKomponentCount,
	}
}

// TextWriter implements a template to go parser.
type TextWriter struct {
	options            Options
	variables          int
	komponents         int
	rootQueue          *varQueue
	htmlQueue          *varQueue
	TreeSet            Treeset
	base               *parse.Tree
	funcMaps           parse.FuncMaps
	filesystem         SourceFileSystem
	rootNode           *parse.RootNode
	builder            *strings.Builder
	KomponentBuilder   *strings.Builder
	funcBuilder        *strings.Builder
	typeBuilder        *strings.Builder
	inMethodDefinition bool
	inHTMLAttr         bool
}

func (t *TextWriter) addImport(alias string, importPath string) {
	importPath = strings.TrimSpace(strings.Trim(importPath, `"`))

	var existingPath, ok = t.options.Packages[alias]
	if !ok {
		t.options.Packages[alias] = importPath
		return
	}

	// if already added then ignore.
	if existingPath == importPath {
		return
	}

	// ensure the path is not already existing with a different alias
	for key, value := range t.options.Packages {
		if value == importPath {
			panic(fmt.Sprintf("import path already exists with alias %q pointed at %q", key, value))
		}
	}

	t.options.Packages[alias] = importPath
}

// swapBuilder swaps current string builder with provided argument, returning
// the previous builder.
func (t *TextWriter) swapBuilder(sd *strings.Builder) *strings.Builder {
	var old *strings.Builder
	old = t.builder
	t.builder = sd
	return old
}

// ParseTree parses the tree with provided Treeset and parent.
func (t *TextWriter) ParseTree() error {
	t.addImport("peji", "github.com/influx6/groundlayer/pkg/peji")
	t.addImport("styled", "github.com/influx6/groundlayer/pkg/styled")

	var err = t.walk(t.base, t.base.Root, nil)
	if err != nil {
		return nerror.WrapOnly(err)
	}
	return nil
}

// func (t *TextWriter) walk(tree *parse.Tree, node parse.Node, parent parse.Node) error {
//	return t.walk(tree, node, parent, t.funcMaps.Komponents)
// }

func (t *TextWriter) walk(tree *parse.Tree, node parse.Node, parent parse.Node) error {
	var tmp = getBuffer()
	var oldBuilder = t.swapBuilder(tmp)
	defer func() {
		oldBuilder.WriteString(tmp.String())
		_ = t.swapBuilder(oldBuilder)

		releaseBuffer(tmp)
	}()

	switch node := node.(type) {
	case *parse.MountLiveNode:
		return t.walkMountLiveKomponent(tree, node, parent)
	case *parse.MountNode:
		return t.walkMountKomponent(tree, node, parent)
	case *parse.RootNode:
		if t.rootNode != nil {
			return nerror.New("root type already has being defined: see %#v\n", t.rootNode)
		}
		t.rootNode = node
	case *parse.ActionNode:
		defer t.Write("\n")
		return t.walkActionNode(tree, node, parent)
	case *parse.ImportDeclr:
		return t.walkImportDeclaration(tree, node, parent)
	case *parse.ImportedFieldFunc:
		return t.walkImportFuncField(tree, node, parent)
	case *parse.IfNode:
		return t.walkIfNode(tree, node, parent)
	case *parse.RangeNode:
		return t.walkRangeNode(tree, node, parent)
	case *parse.WithNode:
		return t.walkWithNode(tree, node, parent)
	case *parse.TextNode:
		return t.walkTextNode(tree, node, parent)
	case *parse.DOMIdentifier:
		return t.walkDOMIdentifier(tree, node, parent)
	case *parse.TextIdentifierNode:
		return t.walkTextIdentifier(tree, node, parent)
	case *parse.ListNode:
		return t.walkListNode(tree, node, parent)
	case *parse.HTMLNode:
		return t.walkHTMLNode(tree, node, parent)
	case *parse.TemplateNode:
		return t.walkTemplateNode(tree, node, parent)
	case *parse.FieldNode:
		return t.walkField(tree, node, parent)
	case *parse.CommandNode:
		return t.walkCommand(tree, node, parent)
	case *parse.ChainNode:
		return t.walkChainNode(tree, node, parent)
	case *parse.IdentifierNode:
		return t.walkIdentifier(tree, node, parent)
	case *parse.PipeNode:
		return t.walkPipeNode(tree, node, parent)
	case *parse.VariableNode:
		return t.walkVariableIdentifier(tree, node, parent)
	case *parse.BoolNode:
		t.BooleanValue(node.True)
	case *parse.DotNode:
		t.DotValue()
	case *parse.NumberNode:
		return t.NumberValue(node)
	case *parse.StringNode:
		return t.walkStringNode(tree, node, parent)
	case *parse.ModelTypeNode:
		return t.walkModelTypeNode(tree, node, parent)
	case *parse.ModelNode:
		return t.walkModelNode(tree, node, parent)
	case *parse.MethodNode:
		return t.walkMethodNode(tree, node, parent)
	case *parse.NilNode:
		t.NilValue()
	default:
		return nerror.New("unknown node type")
	}

	return nil
}

func (t *TextWriter) walkImportFuncField(tree *parse.Tree, node *parse.ImportedFieldFunc, parent parse.Node) error {
	if node.IsField() {
		return t.walkImportField(tree, node, parent)
	}
	return t.walkImportFunc(tree, node, parent)
}

func (t *TextWriter) walkImportField(tree *parse.Tree, node *parse.ImportedFieldFunc, parent parse.Node) error {
	if parent != nil {
		t.Write(node.ToName())
		return nil
	}

	if err := t.attachContentToNode(tree, node.ToName()); err != nil {
		return nerror.WrapOnly(err)
	}
	return nil
}

func (t *TextWriter) walkImportFunc(tree *parse.Tree, node *parse.ImportedFieldFunc, parent parse.Node) error {
	var fieldBuffer = getBuffer()
	defer releaseBuffer(fieldBuffer)

	var oldBuilder = t.swapBuilder(fieldBuffer)

	for _, field := range node.Fields {
		if err := t.walk(tree, field, node); err != nil {
			return nerror.WrapOnly(err)
		}
		t.Write(", ")
	}

	var funcBuffer = getBuffer()
	defer releaseBuffer(funcBuffer)
	funcBuffer.WriteString(node.ToName())
	funcBuffer.WriteString("(")
	funcBuffer.WriteString(fieldBuffer.String())
	funcBuffer.WriteString(")")

	var commands = getBuffer()
	defer releaseBuffer(commands)
	_ = t.swapBuilder(commands)

	var cmder CommandPipe
	cmder.insertText = funcBuffer.String()
	if err := cmder.Prepare(node.Commands); err != nil {
		return nerror.WrapOnly(err)
	}

	if cmdErr := cmder.Parse(t, tree); cmdErr != nil {
		return nerror.WrapOnly(cmdErr)
	}

	t.swapBuilder(oldBuilder)

	if !cmder.NoAttaching {
		if err := t.attachContentToNode(tree, commands.String()); err != nil {
			return nerror.WrapOnly(err)
		}
		return nil
	}

	oldBuilder.WriteString(commands.String())
	oldBuilder.WriteString("\n")
	return nil
}

func (t *TextWriter) walkImportDeclaration(tree *parse.Tree, node *parse.ImportDeclr, parent parse.Node) error {
	var importPath = strings.TrimSpace(strings.Trim(node.Path, `"`))
	t.addImport(node.Alias, importPath)
	return nil
}

func (t *TextWriter) walkTextIdentifier(tree *parse.Tree, node *parse.TextIdentifierNode, parent parse.Node) error {
	t.Write(node.Text)
	return nil
}

func (t *TextWriter) walkStringNode(tree *parse.Tree, node *parse.StringNode, parent parse.Node) error {
	t.Write(node.Quoted)
	return nil
}

func (t *TextWriter) walkMountLiveKomponent(tree *parse.Tree, node *parse.MountLiveNode, parent parse.Node) error {
	var mountBuffer = getBuffer()
	defer releaseBuffer(mountBuffer)

	var oldBuilder = t.swapBuilder(mountBuffer)
	defer t.swapBuilder(oldBuilder)

	t.Write("\n")
	t.Write("if page.HasLive(")
	t.Write("\"")
	t.Write(node.Name)
	t.Write("\")")
	t.Write(" { ")
	t.Write("\n\t")

	t.Write("page.Live(")
	t.Write("\"")
	t.Write(node.Route)
	t.Write("\"")
	t.Write(", ")
	t.Write("\"")
	t.Write(node.Name)
	t.Write("\"")
	t.Write(").")
	t.Write("Render(")

	var mountDataBuffer = getBuffer()
	defer releaseBuffer(mountDataBuffer)

	if node.Target != nil {
		_ = t.swapBuilder(mountDataBuffer)

		if err := t.walk(tree, node.Target, node); err != nil {
			return nerror.WrapOnly(err)
		}

		t.swapBuilder(mountBuffer)

		t.Write(mountDataBuffer.String())
	}

	if node.Target == nil && t.inMethodDefinition {
		return nerror.New("when mounting component within a method you be explicit as to argument used as context")
	}

	if node.Target == nil {
		_ = t.swapBuilder(mountDataBuffer)

		if currentRoot, ok := t.rootQueue.peek().(singleVar); ok {
			mountDataBuffer.WriteString(currentRoot.Name)
		}

		t.swapBuilder(mountBuffer)

		t.Write(mountDataBuffer.String())
	}

	var currentNode, ok = t.htmlQueue.peek().(singleVar)
	if !ok {
		return nerror.New("expected a html node in html queue")
	}
	t.Write(").Mount(")
	t.Write(currentNode.Name)
	t.Write(")")
	t.Write("\n")
	t.Write("}")
	t.Write("\n")

	oldBuilder.WriteString(mountBuffer.String())
	return nil
}

func (t *TextWriter) walkMountKomponent(tree *parse.Tree, node *parse.MountNode, parent parse.Node) error {
	var mountBuffer = getBuffer()
	defer releaseBuffer(mountBuffer)

	var oldBuilder = t.swapBuilder(mountBuffer)
	defer t.swapBuilder(oldBuilder)

	t.Write("\n")
	t.Write("if page.HasStatic(")
	t.Write("\"")
	t.Write(node.Name)
	t.Write("\")")
	t.Write(" { ")
	t.Write("\n\t")

	t.Write("page.Static(")
	t.Write("\"")
	t.Write(node.Name)
	t.Write("\").")
	t.Write("Render(")

	var mountDataBuffer = getBuffer()
	defer releaseBuffer(mountDataBuffer)

	if node.Target != nil {
		_ = t.swapBuilder(mountDataBuffer)

		if err := t.walk(tree, node.Target, node); err != nil {
			return nerror.WrapOnly(err)
		}

		t.swapBuilder(mountBuffer)

		t.Write(mountDataBuffer.String())
	}

	if node.Target == nil && t.inMethodDefinition {
		return nerror.New("when mounting component within a method you be explicit as to argument used as context")
	}

	if node.Target == nil {
		_ = t.swapBuilder(mountDataBuffer)

		if currentRoot, ok := t.rootQueue.peek().(singleVar); ok {
			mountDataBuffer.WriteString(currentRoot.Name)
		}

		t.swapBuilder(mountBuffer)

		t.Write(mountDataBuffer.String())
	}

	var currentNode, ok = t.htmlQueue.peek().(singleVar)
	if !ok {
		return nerror.New("expected a html node in html queue")
	}
	t.Write(").Mount(")
	t.Write(currentNode.Name)
	t.Write(")")
	t.Write("\n")
	t.Write("}")
	t.Write("\n")

	oldBuilder.WriteString(mountBuffer.String())
	return nil
}

func (t *TextWriter) walkModelTypeNode(tree *parse.Tree, node *parse.ModelTypeNode, parent parse.Node) error {
	var modelTypeCondition = getBuffer()
	defer releaseBuffer(modelTypeCondition)

	var oldBuilder = t.swapBuilder(modelTypeCondition)

	t.Write("type")
	t.Space()
	t.Write(node.Typ)
	t.Space()

	var err error
	switch detailNode := node.Detail.(type) {
	case *parse.BaseTypeNode:
		err = t.walkBaseTypeNode(tree, detailNode, node)
	case *parse.CollectionTypeNode:
		err = t.walkCollectionTypeNode(tree, detailNode, node)
	}

	if err != nil {
		return nerror.WrapOnly(err)
	}

	t.NewLine()
	t.NewLine()
	t.swapBuilder(oldBuilder)

	t.typeBuilder.WriteString(modelTypeCondition.String())

	return nil
}

func (t *TextWriter) walkBaseTypeNode(tree *parse.Tree, node *parse.BaseTypeNode, parent parse.Node) error {
	t.Write(node.Typ)
	return nil
}

func (t *TextWriter) walkModelNodeName(tree *parse.Tree, node *parse.ModelNode, parent parse.Node) error {
	t.Write(node.Model)
	return nil
}

func (t *TextWriter) walkModelTypeNodeName(tree *parse.Tree, node *parse.ModelTypeNode, parent parse.Node) error {
	t.Write(node.Typ)
	return nil
}

func (t *TextWriter) walkTypeNodeName(tree *parse.Tree, node parse.Node, parent parse.Node) error {
	switch tnode := node.(type) {
	case *parse.BaseTypeNode:
		return t.walkBaseTypeNode(tree, tnode, parent)
	case *parse.ModelNode:
		return t.walkModelNodeName(tree, tnode, parent)
	case *parse.ModelTypeNode:
		return t.walkModelTypeNodeName(tree, tnode, parent)
	}
	return nerror.New("unknown node type received")
}

func (t *TextWriter) walkCollectionTypeNode(tree *parse.Tree, node *parse.CollectionTypeNode, parent parse.Node) error {
	switch node.Typ {
	case parse.MapType:
		key := node.Detail.Nodes[0]
		value := node.Detail.Nodes[1]
		t.Write("map[")
		_ = t.walkTypeNodeName(tree, key, parent)
		t.Write("]")
		_ = t.walkTypeNodeName(tree, value, parent)
	case parse.ListType:
		key := node.Detail.Nodes[0]
		t.Write("[]")
		_ = t.walkTypeNodeName(tree, key, parent)
	default:
		return nerror.New("unknown collection type %#v", node)
	}
	return nil
}

func (t *TextWriter) walkModelNode(tree *parse.Tree, node *parse.ModelNode, parent parse.Node) error {
	var modelTypeCondition = getBuffer()
	defer releaseBuffer(modelTypeCondition)

	var oldBuilder = t.swapBuilder(modelTypeCondition)

	t.Write("type")
	t.Space()
	t.Write(node.Model)
	t.Space()
	t.Write("struct {")
	t.NewLine()

	var fieldLen = len(node.Fields.Nodes)
	for index, field := range node.Fields.Nodes {
		var refField, ok = field.(*parse.RefNode)
		if !ok {
			return nerror.New("unknown filed type %#v", field)
		}

		t.Indent()
		t.Write(refField.Name)
		t.Space()
		t.Write(refField.Typ)

		if index < fieldLen-1 {
			t.NewLine()
		}
	}

	t.NewLine()
	t.Write("}")
	t.NewLine()
	t.NewLine()

	t.swapBuilder(oldBuilder)

	t.typeBuilder.WriteString(modelTypeCondition.String())
	return nil
}

func (t *TextWriter) walkMethodNode(tree *parse.Tree, node *parse.MethodNode, parent parse.Node) error {
	var methodContent = getBuffer()
	defer releaseBuffer(methodContent)

	var oldBuilder = t.swapBuilder(methodContent)
	defer t.swapBuilder(oldBuilder)

	t.Write("func ")
	t.Write(node.Model)
	t.Write("(page *peji.Page, ")

	for _, argNode := range node.Arguments {
		t.Write(argNode.Name)
		t.Space()
		t.Write(argNode.Typ)
		t.Write(",")
		t.Space()
	}

	var parentNodeName = "parentNode"
	t.Write(parentNodeName)
	t.Space()
	t.Write("*domu.Node")
	t.Write(") {")
	t.NewLine()
	t.NewLine()
	t.htmlQueue.pushSingle(parentNodeName)

	// flag that we are in a method.
	t.inMethodDefinition = true
	for _, methodChild := range node.Body.Nodes {
		if err := t.walk(tree, methodChild, node); err != nil {
			return err
		}
	}
	t.inMethodDefinition = false

	t.htmlQueue.pop()
	t.NewLine()
	t.Write("}")
	t.NewLine()

	t.funcBuilder.WriteString(methodContent.String())
	return nil
}

func (t *TextWriter) walkHTMLNode(tree *parse.Tree, node *parse.HTMLNode, parent parse.Node) error {
	var htmlContent = getBuffer()
	defer releaseBuffer(htmlContent)

	var oldBuilder = t.swapBuilder(htmlContent)
	defer t.swapBuilder(oldBuilder)

	var nodeName = t.newVar("node")
	t.NewElement(nodeName, node.Tag, nodeName)
	t.NewLine()

	for _, attr := range node.Attr {
		var attrName string

		if attr.IsTheme {
			attrName = t.newVar("theme")
		} else {
			attrName = t.newVar("attr")
		}

		if attr.IsTheme {
			t.NewHtmlTheme(attrName)
		} else {
			t.NewHtmlListAttr(attrName, attr.Key)
		}
		t.NewLine()

		t.htmlQueue.pushSingle(attrName)

		var htmlAttrContent = getBuffer()
		var currentBuffer = t.swapBuilder(htmlAttrContent)

		for _, value := range attr.Values.Nodes {
			if err := t.walkAttrValue(tree, value, node, attrName); err != nil {
				return nerror.WrapOnly(err)
			}

			currentBuffer.WriteString(htmlAttrContent.String())
			htmlAttrContent.Reset()
		}

		t.htmlQueue.pop()
		t.swapBuilder(currentBuffer)
		releaseBuffer(htmlAttrContent)

		if attr.IsTheme {
			t.AssignNodeTheme(attrName, nodeName)
		} else {
			t.AppendNode(nodeName, attrName)
		}

		t.NewLine()
		t.NewLine()
	}

	if totalChildren := len(node.Children.Nodes); totalChildren > 0 {
		// push node name into the stack
		t.htmlQueue.pushSingle(nodeName)

		for _, childNode := range node.Children.Nodes {
			if err := t.walk(tree, childNode, node); err != nil {
				return nerror.WrapOnly(err)
			}
		}

		// pop node name from the stack
		t.htmlQueue.pop()
	}

	var currentNode, ok = t.htmlQueue.peek().(singleVar)
	if !ok {
		return nerror.New("no dom node found in html variable queue")
	}

	t.NewLine()
	t.NewLine()
	t.AppendNode(currentNode.Name, nodeName)
	t.NewLine()
	t.NewLine()

	oldBuilder.WriteString(htmlContent.String())
	return nil
}

func (t *TextWriter) walkAttrValue(tree *parse.Tree, node parse.Node, parent parse.Node, varName string) error {
	t.inHTMLAttr = true
	switch value := node.(type) {
	case *parse.TextNode:
		var deQuoted = nunsafe.Bytes2String(value.Text)
		deQuoted = strings.Trim(deQuoted, "\"")
		if deQuoted != "" {
			t.AddValueToHtmlListAttr(varName, deQuoted)
			t.NewLine()
		}
	case *parse.ActionNode:
		if err := t.walkHTMLPipeNode(tree, value.Pipe, parent); err != nil {
			return nerror.WrapOnly(err)
		}
	case *parse.IfNode:
		if err := t.walkIfNode(tree, value, parent); err != nil {
			return nerror.WrapOnly(err)
		}
	default:
		return nerror.New("unable to handle attr value for node %#v", node)
	}
	t.inHTMLAttr = false
	return nil
}

func (t *TextWriter) walkTemplateNode(tree *parse.Tree, node *parse.TemplateNode, parent parse.Node) error {
	if node.Pipe == nil {
		node.Pipe = &parse.PipeNode{
			Cmds: []*parse.CommandNode{
				{
					Args: []parse.Node{
						&parse.DotNode{},
					},
				},
			},
		}
	}

	var pipeCondition = getBuffer()
	defer releaseBuffer(pipeCondition)

	var oldBuilder = t.swapBuilder(pipeCondition)

	if err := t.walkNonDisplayPipeNode(tree, node.Pipe, node); err != nil {
		return nerror.WrapOnly(err)
	}

	t.swapBuilder(oldBuilder)

	var fileName, blockName = parseTemplateName(node.Name)

	var funcName, hasFuncName = t.funcMaps.DefinedBlocks[blockName]
	if !hasFuncName {
		funcName = t.newVar("definedTemplate")
	}

	// do we already have it defined, else skip if not parse and
	// generate definition.
	if !hasFuncName {

		// do we have an existing tree parsed for giving requested template ?
		var targetTree, hasTree = t.TreeSet[blockName]
		if hasTree {
			var targetParseBuilder = getBuffer()
			var funcParseBuilder = getBuffer()
			var typeParseBuilder = getBuffer()
			var KomponentParseBuilder = getBuffer()

			var targetTextWriter = newTextWriter(
				targetParseBuilder,
				funcParseBuilder,
				typeParseBuilder,
				targetTree,
				t.options,
				t.funcMaps,
				t.TreeSet,
				t.filesystem,
				rootVar,
				htmlRootVar,
			)

			t.addVariableCount()
			targetTextWriter.variables = t.variables

			t.incKomponentCounts()
			targetTextWriter.komponents = t.komponents

			targetTextWriter.rootNode = t.rootNode
			if targetWriterErr := targetTextWriter.ParseTree(); targetWriterErr != nil {
				return nerror.WrapOnly(targetWriterErr)
			}

			if KomponentParseBuilder.Len() > 0 {
				t.KomponentBuilder.WriteString("\n")
				t.KomponentBuilder.WriteString(KomponentParseBuilder.String())
				t.KomponentBuilder.WriteString("\n")
			}

			t.typeBuilder.WriteString("\n")
			t.typeBuilder.WriteString(typeParseBuilder.String())
			t.typeBuilder.WriteString("\n")

			t.funcBuilder.WriteString(funcParseBuilder.String())
			t.funcBuilder.WriteString("\n")
			t.funcBuilder.WriteString("\n")
			t.funcBuilder.WriteString(`// `)
			t.funcBuilder.WriteString(funcName)
			t.funcBuilder.WriteString(" is a template defined in ")
			t.funcBuilder.WriteString(`"`)
			t.funcBuilder.WriteString(targetTree.Name)
			t.funcBuilder.WriteString(`"`)

			var wrappedContent = wrapWithBaseFunction(funcName, targetTree.TypeName, targetParseBuilder)
			t.funcBuilder.WriteString("\n")
			t.funcBuilder.WriteString(wrappedContent)
			t.funcBuilder.WriteString("\n")

			releaseBuffer(targetParseBuilder)
			releaseBuffer(funcParseBuilder)
			releaseBuffer(funcParseBuilder)
		} else {
			var targetFile, targetFileErr = t.filesystem.GetFile(fileName)
			if targetFileErr != nil {
				return nerror.WrapOnly(targetFileErr)
			}

			var newOption = t.options
			t.addVariableCount()
			newOption.lastVariableCount = t.variables
			t.incKomponentCounts()
			newOption.lastKomponentCount = t.getKomponentCount()

			var parsedFileTree, parsedFileErr = targetFile.ParseBlockFor(blockName, newOption, t.funcMaps, t.TreeSet)
			if parsedFileErr != nil {
				return nerror.WrapOnly(parsedFileErr)
			}

			// Write out the functions here into the list
			t.typeBuilder.WriteString("\n")
			t.typeBuilder.WriteString(parsedFileTree.types.String())
			t.typeBuilder.WriteString("\n")

			t.funcBuilder.WriteString("\n")
			t.funcBuilder.WriteString(parsedFileTree.functions.String())
			t.funcBuilder.WriteString("\n")

			t.funcBuilder.WriteString("\n")
			t.funcBuilder.WriteString("\n")
			t.funcBuilder.WriteString(`// `)
			t.funcBuilder.WriteString(funcName)
			t.funcBuilder.WriteString(" is a template defined in ")
			t.funcBuilder.WriteString(`"`)
			t.funcBuilder.WriteString(node.Name)
			t.funcBuilder.WriteString(`"`)

			var wrappedContent = wrapWithBaseFunction(funcName, parsedFileTree.typeName, parsedFileTree.body)
			t.funcBuilder.WriteString("\n")
			t.funcBuilder.WriteString(wrappedContent)
			t.funcBuilder.WriteString("\n")
		}
	}

	t.funcMaps.DefinedBlocks[blockName] = funcName

	t.NewLine()
	t.Write(funcName)
	t.Write("(")
	t.Write("page, ")
	t.Write(pipeCondition.String())
	t.Write(", ")

	var currentNode, ok = t.htmlQueue.peek().(singleVar)
	if !ok {
		return nerror.New("no dom node found in html variable queue")
	}

	t.Write(currentNode.Name)
	t.Write(")")
	t.NewLine()

	return nil
}

func (t *TextWriter) walkListNode(tree *parse.Tree, node *parse.ListNode, parent parse.Node) error {
	for _, listNode := range node.Nodes {

		// Are we dealing with a imported function ?
		if importedNode, isImportedNode := listNode.(*parse.ImportedFieldFunc); isImportedNode {
			var importedAction = &parse.ActionNode{
				Pipe: &parse.PipeNode{
					Cmds: append([]*parse.CommandNode{
						{
							Args: []parse.Node{
								importedNode,
							},
						},
					}, importedNode.Commands...),
				},
			}

			importedNode.Commands = nil
			if parseErr := t.walk(tree, importedAction, parent); parseErr != nil {
				return nerror.WrapOnly(parseErr)
			}
			continue
		}
		if parseErr := t.walk(tree, listNode, parent); parseErr != nil {
			return nerror.WrapOnly(parseErr)
		}
	}
	return nil
}

func (t *TextWriter) walkTextNode(tree *parse.Tree, node *parse.TextNode, parent parse.Node) error {
	var currentNode, ok = t.htmlQueue.peek().(singleVar)
	if !ok {
		return nerror.New("expected a singleVar in variable queue")
	}

	if t.inHTMLAttr {
		var text = strings.TrimSpace(allLineSpacesTabs.ReplaceAllString(string(node.Text), ""))
		for _, part := range strings.Split(text, " ") {
			t.AddToAttr(currentNode.Name, part)
			t.NewLine()
		}
		return nil
	}

	var content = string(node.Text)
	var varName = fmt.Sprintf("text%d", t.nextVariableCount())
	t.NewText(varName, content)
	t.NewLine()
	t.AppendNode(currentNode.Name, varName)
	t.NewLine()
	return nil
}

func (t *TextWriter) attachContentToNode(tree *parse.Tree, content string) error {
	var currentNode, ok = t.htmlQueue.peek().(singleVar)
	if !ok {
		return nerror.New("expected a singleVar in variable queue")
	}

	t.attachResultToNode(content, currentNode.Name)
	t.NewLine()
	return nil
}

func (t *TextWriter) walkWithNode(tree *parse.Tree, node *parse.WithNode, parent parse.Node) error {
	var rangeBranch = node.BranchNode

	var rangeCondition = getBuffer()
	defer releaseBuffer(rangeCondition)

	var oldBuilder = t.swapBuilder(rangeCondition)
	defer t.swapBuilder(oldBuilder)

	if err := t.walkNonDisplayPipeNode(tree, rangeBranch.Pipe, node); err != nil {
		return nerror.WrapOnly(err)
	}

	var rangeVariable = getBuffer()
	defer releaseBuffer(rangeVariable)

	_ = t.swapBuilder(rangeVariable)

	var varName = t.newVar("item")
	if len(rangeBranch.Pipe.Decl) == 0 {
		t.rootQueue.pushSingle(varName)
		defer t.rootQueue.pop()

		rangeBranch.Pipe.Decl = append(rangeBranch.Pipe.Decl, &parse.VariableNode{
			Ident: []string{varName},
		})
	}

	if err := t.walkPipeNodeVariable(tree, rangeBranch.Pipe, node); err != nil {
		return nerror.WrapOnly(err)
	}

	var rangeResult = getBuffer()
	defer releaseBuffer(rangeResult)
	_ = t.swapBuilder(rangeResult)

	// handle items within the range condition.
	for _, item := range rangeBranch.List.Nodes {
		if err := t.walk(tree, item, node); err != nil {
			return nerror.WrapOnly(err)
		}
	}

	var hasElseSection = false

	var rangeElseResult = getBuffer()
	defer releaseBuffer(rangeElseResult)
	_ = t.swapBuilder(rangeElseResult)

	// handle items within the range condition.
	if rangeBranch.ElseList != nil {
		hasElseSection = len(rangeBranch.ElseList.Nodes) != 0

		for _, item := range rangeBranch.ElseList.Nodes {
			if err := t.walk(tree, item, node); err != nil {
				return nerror.WrapOnly(err)
			}
		}
	}

	oldBuilder.WriteString("\n")
	oldBuilder.WriteString("\n")
	oldBuilder.WriteString(rangeVariable.String())
	oldBuilder.WriteString(rangeCondition.String())

	if hasElseSection {
		var isInvalidFuncName = t.funcMaps.Funcs["isInvalid"]

		oldBuilder.WriteString("\n")
		oldBuilder.WriteString("if ")
		oldBuilder.WriteString(isInvalidFuncName)
		oldBuilder.WriteString("(")
		oldBuilder.WriteString(varName)
		oldBuilder.WriteString(") { \n")
		oldBuilder.WriteString(rangeElseResult.String())
		oldBuilder.WriteString(" } else { ")
		oldBuilder.WriteString("\n")
	}

	oldBuilder.WriteString("\n")
	oldBuilder.WriteString(rangeResult.String())

	if hasElseSection {
		oldBuilder.WriteString("} \n")
		return nil
	}

	oldBuilder.WriteString("\n")
	return nil
}

func (t *TextWriter) walkRangeNode(tree *parse.Tree, node *parse.RangeNode, parent parse.Node) error {
	var rangeBranch = node.BranchNode

	var rangeCondition = getBuffer()
	defer releaseBuffer(rangeCondition)

	var oldBuilder = t.swapBuilder(rangeCondition)
	defer t.swapBuilder(oldBuilder)

	if err := t.walkNonDisplayPipeNode(tree, rangeBranch.Pipe, node); err != nil {
		return nerror.WrapOnly(err)
	}

	var rangeVariable = getBuffer()
	defer releaseBuffer(rangeVariable)
	_ = t.swapBuilder(rangeVariable)

	var varName = t.newVar("item")
	if len(rangeBranch.Pipe.Decl) == 0 {
		t.rootQueue.pushSingle(varName)
		defer t.rootQueue.pop()

		rangeBranch.Pipe.Decl = append(rangeBranch.Pipe.Decl, &parse.VariableNode{
			Ident: []string{"_"},
		}, &parse.VariableNode{
			Ident: []string{varName},
		})
	}

	if err := t.walkRangePipeNodeVariable(tree, rangeBranch.Pipe, node); err != nil {
		return nerror.WrapOnly(err)
	}

	var rangeResult = getBuffer()
	defer releaseBuffer(rangeResult)
	_ = t.swapBuilder(rangeResult)

	// handle items within the range condition.
	for _, item := range rangeBranch.List.Nodes {
		if err := t.walk(tree, item, node); err != nil {
			return nerror.WrapOnly(err)
		}
	}

	var hasElseSection = false

	var rangeElseResult = getBuffer()
	defer releaseBuffer(rangeElseResult)
	_ = t.swapBuilder(rangeElseResult)

	// handle items within the range condition.
	if rangeBranch.ElseList != nil {
		hasElseSection = len(rangeBranch.ElseList.Nodes) != 0

		for _, item := range rangeBranch.ElseList.Nodes {
			if err := t.walk(tree, item, node); err != nil {
				return nerror.WrapOnly(err)
			}
		}
	}

	if hasElseSection {
		oldBuilder.WriteString("\n")
		oldBuilder.WriteString("if helpers.Len(")
		oldBuilder.WriteString(rangeCondition.String())
		oldBuilder.WriteString(") == 0 { \n")
		oldBuilder.WriteString(rangeElseResult.String())
		oldBuilder.WriteString(" } else { ")
		oldBuilder.WriteString("\n")
	}

	oldBuilder.WriteString("\n")
	oldBuilder.WriteString("for ")
	oldBuilder.WriteString(rangeVariable.String())
	oldBuilder.WriteString(rangeCondition.String())
	oldBuilder.WriteString(" { \n")
	oldBuilder.WriteString(rangeResult.String())
	oldBuilder.WriteString(" }\n")

	if hasElseSection {
		oldBuilder.WriteString("} \n")
		return nil
	}

	oldBuilder.WriteString("\n")
	return nil
}

func (t *TextWriter) walkIfNode(tree *parse.Tree, node *parse.IfNode, parent parse.Node) error {
	var ifCondition = getBuffer()
	defer releaseBuffer(ifCondition)

	var oldBuilder = t.swapBuilder(ifCondition)
	defer t.swapBuilder(oldBuilder)

	if err := t.walkNonDisplayPipeNode(tree, node.Pipe, node); err != nil {
		return nerror.WrapOnly(err)
	}

	var ifResult = getBuffer()
	defer releaseBuffer(ifResult)
	_ = t.swapBuilder(ifResult)

	// handle items within the if condition.
	for _, item := range node.List.Nodes {
		if err := t.walk(tree, item, node); err != nil {
			return nerror.WrapOnly(err)
		}
	}

	var elseContent = getBuffer()
	defer releaseBuffer(elseContent)
	t.swapBuilder(elseContent)

	var hasElseContent = false
	var isAnyIfElse = false
	// handle items within the else condition.
	if node.ElseList != nil {
		var totalElseNodes = len(node.ElseList.Nodes)
		hasElseContent = totalElseNodes > 0

		if totalElseNodes == 1 {
			_, isAnyIfElse = node.ElseList.Nodes[0].(*parse.IfNode)
		}

		for _, item := range node.ElseList.Nodes {
			if err := t.walk(tree, item, node); err != nil {
				return nerror.WrapOnly(err)
			}
		}
	}

	oldBuilder.WriteString("\n")
	oldBuilder.WriteString("if ")
	oldBuilder.WriteString(ifCondition.String())
	oldBuilder.WriteString(" {\n")
	oldBuilder.WriteString(ifResult.String())
	oldBuilder.WriteString("\n\n}")

	if hasElseContent {
		if isAnyIfElse {
			oldBuilder.WriteString(" else ")
			oldBuilder.WriteString(elseContent.String())
		} else {
			oldBuilder.WriteString(" else {\n")
			oldBuilder.WriteString(elseContent.String())
			oldBuilder.WriteString("}")
			oldBuilder.WriteString("\n")
		}
	}

	oldBuilder.WriteString("\n")
	return nil
}

func (t *TextWriter) walkActionNode(tree *parse.Tree, node *parse.ActionNode, parent parse.Node) error {
	var copiedErr error
	var copiedPipe = *node.Pipe
	copiedPipe.Cmds, copiedErr = t.sortCommands(tree, node.Pipe.Cmds)
	if copiedErr != nil {
		return nerror.WrapOnly(copiedErr)
	}
	return t.walkPipeNode(tree, &copiedPipe, node)
}

func (t *TextWriter) sortCommands(tree *parse.Tree, commands []*parse.CommandNode) ([]*parse.CommandNode, error) {
	var sorted []*parse.CommandNode

	for _, command := range commands {

		// Are we dealing with a imported function ?
		// If so, it will be the first element.
		if importedField, hasImportedField := command.Args[0].(*parse.ImportedFieldFunc); hasImportedField {

			sorted = append(sorted, &parse.CommandNode{
				Args: []parse.Node{
					importedField,
				},
			})

			// add these commands flat into the list.
			sorted = append(sorted, importedField.Commands...)
			importedField.Commands = nil
			continue
		}

		// Does this command have possible
		sorted = append(sorted, command)
	}

	return sorted, nil
}

func (t *TextWriter) walkIdentifier(tree *parse.Tree, node *parse.IdentifierNode, parent parse.Node) error {
	if node.IsInbuiltFunc {
		t.addImport(helperAlias, helperImportPath)
	}

	t.Write(node.Ident)
	return nil
}

func (t *TextWriter) walkField(tree *parse.Tree, node *parse.FieldNode, parent parse.Node) error {
	if current, ok := t.rootQueue.peek().(singleVar); ok {
		t.Write(current.Name)
		t.Write(".")
		t.Write(strings.Join(node.Ident, "."))
		return nil
	}

	return nerror.New("require root queue to have parent identifier")
}

func (t *TextWriter) walkDotValue(tree *parse.Tree, node *parse.DotNode, parent parse.Node) error {
	if current, ok := t.rootQueue.peek().(singleVar); ok {
		t.Write(current.Name)
		return nil
	}

	return nerror.New("require root queue to have parent identifier")
}

func (t *TextWriter) walkPipeNodeTarget(tree *parse.Tree, node *parse.PipeNode, parent parse.Node) error {
	var totalCmds = len(node.Cmds)
	if totalCmds == 0 || totalCmds > 1 {
		return nil
	}

	var targetCommand = node.Cmds[0]
	if itemErr := t.walkCommand(tree, targetCommand, node); itemErr != nil {
		return nerror.WrapOnly(itemErr)
	}
	return nil
}

func (t *TextWriter) walkRangePipeNodeVariable(tree *parse.Tree, node *parse.PipeNode, parent parse.Node) error {
	var hasVariables = len(node.Decl) > 0
	if !hasVariables {
		return nil
	}

	for index, variable := range node.Decl {
		if index > 0 {
			t.Write(", ")
		}
		var set = strings.Replace(strings.Join(variable.Ident, ","), "$", "", -1)
		t.Write(set)
	}
	t.Write(" := range ")

	return nil
}

func (t *TextWriter) walkPipeNodeVariable(tree *parse.Tree, node *parse.PipeNode, parent parse.Node) error {
	var hasVariables = len(node.Decl) > 0
	if !hasVariables {
		return nil
	}

	for index, variable := range node.Decl {
		if index > 0 {
			t.Write(", ")
		}
		var set = strings.Replace(strings.Join(variable.Ident, ","), "$", "", -1)
		t.Write(set)
	}

	if node.IsAssign {
		t.Write(" = ")
		return nil
	}

	t.Write(" := ")
	return nil
}

func (t *TextWriter) walkNonDisplayPipeNode(tree *parse.Tree, node *parse.PipeNode, parent parse.Node) error {
	var totalCmds = len(node.Cmds)
	if totalCmds == 0 {
		return nil
	}

	var tmp = getBuffer()
	defer releaseBuffer(tmp)

	var oldBuilder = t.swapBuilder(tmp)

	var cmder CommandPipe
	if err := cmder.Prepare(node.Cmds); err != nil {
		return nerror.WrapOnly(err)
	}

	if cmdErr := cmder.Parse(t, tree); cmdErr != nil {
		return nerror.WrapOnly(cmdErr)
	}

	oldBuilder.WriteString(tmp.String())
	t.swapBuilder(oldBuilder)
	return nil
}

func (t *TextWriter) walkHTMLPipeNode(tree *parse.Tree, node *parse.PipeNode, parent parse.Node) error {
	var tmp = getBuffer()
	defer releaseBuffer(tmp)

	var oldBuilder = t.swapBuilder(tmp)

	if varErr := t.walkPipeNodeVariable(tree, node, parent); varErr != nil {
		return nerror.WrapOnly(varErr)
	}

	var totalCmds = len(node.Cmds)
	if totalCmds == 0 {
		return nil
	}

	var cmder CommandPipe
	if err := cmder.Prepare(node.Cmds); err != nil {
		return nerror.WrapOnly(err)
	}

	if cmdErr := cmder.Parse(t, tree); cmdErr != nil {
		return nerror.WrapOnly(cmdErr)
	}

	t.swapBuilder(oldBuilder)

	if cmder.NoAttaching {
		oldBuilder.WriteString(tmp.String())
		return nil
	}

	if err := t.attachContentToNode(tree, tmp.String()); err != nil {
		return nerror.WrapOnly(err)
	}

	return nil
}

func (t *TextWriter) walkPipeNode(tree *parse.Tree, node *parse.PipeNode, parent parse.Node) error {
	var tmp = getBuffer()
	defer releaseBuffer(tmp)

	var oldBuilder = t.swapBuilder(tmp)

	if varErr := t.walkPipeNodeVariable(tree, node, parent); varErr != nil {
		return nerror.WrapOnly(varErr)
	}

	var totalCmds = len(node.Cmds)
	if totalCmds == 0 {
		return nil
	}

	var cmder CommandPipe
	if err := cmder.Prepare(node.Cmds); err != nil {
		return nerror.WrapOnly(err)
	}

	if cmdErr := cmder.Parse(t, tree); cmdErr != nil {
		return nerror.WrapOnly(cmdErr)
	}

	t.swapBuilder(oldBuilder)

	// if no variables then we are possibly dealing with a plain statement
	// so pipe it into current root node.
	if !cmder.NoAttaching {
		var _, isActionNode = parent.(*parse.ActionNode)
		if len(node.Decl) == 0 && isActionNode {
			if err := t.attachContentToNode(tree, tmp.String()); err != nil {
				return nerror.WrapOnly(err)
			}

			return nil
		}
	}

	oldBuilder.WriteString(tmp.String())
	return nil
}

func (t *TextWriter) walkCommand(tree *parse.Tree, node *parse.CommandNode, parent parse.Node) error {
	var tmp = getBuffer()
	defer releaseBuffer(tmp)

	var oldBuilder = t.swapBuilder(tmp)

	var cmder CommandPipe

	var targetList = []*parse.CommandNode{node}
	if err := cmder.Prepare(targetList); err != nil {
		return nerror.WrapOnly(err)
	}

	if cmdErr := cmder.Parse(t, tree); cmdErr != nil {
		return nerror.WrapOnly(cmdErr)
	}

	t.swapBuilder(oldBuilder)

	oldBuilder.WriteString(tmp.String())
	return nil
}

func (t *TextWriter) walkVariableIdentifier(tree *parse.Tree, node *parse.VariableNode, parent parse.Node) error {
	var tmp = getBuffer()
	defer releaseBuffer(tmp)

	var oldBuilder = t.swapBuilder(tmp)
	defer t.swapBuilder(oldBuilder)

	for index, field := range node.Ident {
		if index > 0 {
			t.Write(".")
		}
		t.Write(strings.Replace(field, "$", "", -1))
	}

	oldBuilder.WriteString(tmp.String())
	return nil
}

func (t *TextWriter) walkDOMIdentifier(tree *parse.Tree, node *parse.DOMIdentifier, parent parse.Node) error {
	switch node.Text {
	case "dom":
		if current, ok := t.htmlQueue.peek().(singleVar); ok {
			t.Write(current.Name)
			return nil
		}
	case "rootdom":
		if current, ok := t.htmlQueue.first().(singleVar); ok {
			t.Write(current.Name)
			return nil
		}
	case "parentdom":
		if current, ok := t.htmlQueue.beforeLast().(singleVar); ok {
			t.Write(current.Name)
			return nil
		}
	}

	return nerror.New("unknown dom identifier: %q", node.Text)
}

func (t *TextWriter) walkChainNode(
	tree *parse.Tree,
	chain *parse.ChainNode,
	parent parse.Node,
) error {
	if chain.Node.Type() == parse.NodeNil {
		return nerror.New("ChainNode: indirect through explicit nil")
	}
	if len(chain.Field) == 0 {
		return nerror.New("no fields in ChainNode")
	}

	var tmp = getBuffer()
	defer releaseBuffer(tmp)
	tmp.Reset()

	var oldBuilder = t.swapBuilder(tmp)

	if parseErr := t.walk(tree, chain.Node, chain); parseErr != nil {
		return nerror.WrapOnly(parseErr)
	}

	var wrappedContent = "(" + tmp.String() + ")"
	tmp.Reset()
	tmp.WriteString(wrappedContent)

	tmp.WriteString(".")
	tmp.WriteString(strings.Join(chain.Field, "."))

	_, writeErr := oldBuilder.WriteString(tmp.String())
	if writeErr != nil {
		return nerror.WrapOnly(writeErr)
	}

	return nil
}

func (t *TextWriter) Reset() {
	t.builder.Reset()
}

func (t *TextWriter) Write(s string) {
	t.builder.WriteString(s)
}

func (t *TextWriter) currentVariableCount() int {
	return t.variables
}

func (t *TextWriter) nextVariableCount() int {
	t.variables++
	return t.variables
}

func (t *TextWriter) addVariableCount() {
	t.variables++
}

func (t *TextWriter) getVarCount() string {
	return strconv.Itoa(t.variables)
}

func (t *TextWriter) incKomponentCounts() {
	t.komponents++
}

func (t *TextWriter) getKomponentCount() int {
	return t.komponents
}

func (t *TextWriter) getKomponentCountString() string {
	return strconv.Itoa(t.komponents)
}

func (t *TextWriter) newVar(base string) string {
	return fmt.Sprintf("%s%d", base, t.nextVariableCount())
}

func (t *TextWriter) Var() {
	t.Write(`var `)
}

func (t *TextWriter) Identifier(id string) {
	t.Write(id)
}

func (t *TextWriter) NewVarStart(varName string) {
	t.Write(fmt.Sprintf(`var %s = `, varName))
}

func (t *TextWriter) NewVar(varName string, valName string) {
	t.Write(fmt.Sprintf(`var %s = %s`, varName, valName))
}

func (t *TextWriter) NewText(varName string, text string) {
	t.Write(fmt.Sprintf(`var %s = domu.Text(%q)`, varName, text))
}

func (t *TextWriter) AddToAttr(varName string, text string) {
	t.Write(fmt.Sprintf(`%s.Add(%q)`, varName, text))
}

func (t *TextWriter) attachResultToNode(content string, node string) {
	var attachToName, hasAttachFuncName = t.funcMaps.Funcs[attachNodeFuncName]
	if !hasAttachFuncName {
		panic("func name not found: " + attachNodeFuncName)
	}
	t.addImport(helperAlias, helperImportPath)
	t.Write(fmt.Sprintf(`%s(%s, %s)`, attachToName, content, node))
}

func (t *TextWriter) NewDocument(varName string) {
	t.Write(fmt.Sprintf(`var %s = domu.Document()`, varName))
}

func (t *TextWriter) NewHtmlAttr(name string, content string, rootNode string) {
	t.Write(fmt.Sprintf(`domu.NewStringAttr(%q, %q).Mount(%q)`, name, content, rootNode))
}

func (t *TextWriter) NewHtmlListAttr(varName string, name string) {
	t.Write(fmt.Sprintf(`var %s = domu.NewStringListAttr(%q, "")`, varName, name))
}

func (t *TextWriter) NewHtmlTheme(varName string) {
	t.Write(fmt.Sprintf(`var %s = styled.ThemeDirective{}`, varName))
}

func (t *TextWriter) AddValueToHtmlListAttr(varName string, content string) {
	t.Write(fmt.Sprintf(`%s.Add(%q)`, varName, content))
}

func (t *TextWriter) NewElement(varName string, tagName string, id string) {
	t.Write(fmt.Sprintf(`var %s = domu.Element(%q).UseID(%q)`, varName, tagName, id))
}

func (t *TextWriter) NewComment(varName string, content string) {
	t.Write(fmt.Sprintf(`var %s = domu.Comment(%q)`, varName, content))
}

func (t *TextWriter) NewNode(varName string, nodeType int, tagName string, id string) {
	t.Write(fmt.Sprintf(`var %s = domu.NewNode(%d, %q).UseID(%q)`, varName, nodeType, tagName, id))
}

func (t *TextWriter) AssignNodeTheme(rootName string, varName string) {
	t.Write(fmt.Sprintf(`%s.Themes = %s`, varName, rootName))
}

func (t *TextWriter) AppendNode(rootName string, varName string) {
	t.Write(fmt.Sprintf(`%s.Mount(%s)`, varName, rootName))
}

func (t *TextWriter) NewFuncStart(funcName string) {
	t.Write(fmt.Sprintf(`func %s(`, funcName))
}

func (t *TextWriter) NewLine() {
	t.Write("\n")
}
func (t *TextWriter) Indent() {
	t.Write("\t")
}

func (t *TextWriter) Space() {
	t.Write(" ")
}

func (t *TextWriter) SpaceN(length int) {
	for i := 0; i < length; i++ {
		t.Write(" ")
	}
}

func (t *TextWriter) DotValue() {
	t.Write(".")
}

func (t *TextWriter) NilValue() {
	t.Write("nil")
}

func (t *TextWriter) NumberValue(constant *parse.NumberNode) error {
	switch {
	case constant.IsComplex:
		t.Write(fmt.Sprintf(`%s`, constant.Text))
		return nil
	case constant.IsFloat &&
		!isHexInt(constant.Text) && !isRuneInt(constant.Text) &&
		strings.ContainsAny(constant.Text, ".eEpP"):
		t.Write(fmt.Sprintf(`%s`, constant.Text))
		return nil
	case constant.IsInt:
		t.Write(fmt.Sprintf(`%s`, constant.Text))
		return nil
	case constant.IsUint:
		t.Write(fmt.Sprintf(`%s`, constant.Text))
		return nil
	}
	return nerror.New("number %q not handled", constant.Text)
}

func (t *TextWriter) StringValue(res string) {
	t.Write(fmt.Sprintf(`%q`, res))
}

func (t *TextWriter) BooleanValue(res bool) {
	t.Write(fmt.Sprintf(`%t`, res))
}

func (t *TextWriter) BooleanVar(varName string, val bool) {
	t.Write(fmt.Sprintf(`%s bool = %t`, varName, val))
}

func (t *TextWriter) FuncArgStart(varName string, varType string) {
	t.Write(fmt.Sprintf(`%s %s`, varName, varType))
}

func (t *TextWriter) AddFuncArg(varName string, varType string) {
	t.Write(fmt.Sprintf(`, %s %s`, varName, varType))
}

func (t *TextWriter) FuncArgEnd() {
	t.Write(`)`)
}

func (t *TextWriter) FuncReturns(returnTypes ...string) {
	t.Write("(")
	for index, rt := range returnTypes {
		if index > 0 {
			t.Write(", ")
		}
		t.Write(rt)
	}
	t.Write(") ")
}

func (t *TextWriter) FuncNamedReturnStart(varName string, varType string) {
	t.Write(fmt.Sprintf(`(%s %s`, varName, varType))
}

func (t *TextWriter) AddFuncNamedReturn(varName string, varType string) {
	t.Write(fmt.Sprintf(`, %s %s`, varName, varType))
}

func (t *TextWriter) FuncNamedReturnEnd(varName string, varType string) {
	t.Write(`)`)
}

func isRuneInt(s string) bool {
	return len(s) > 0 && s[0] == '\''
}

func isHexInt(s string) bool {
	return len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') && !strings.ContainsAny(s, "pP")
}

func getBuffer() *strings.Builder {
	var b strings.Builder
	return &b
}

func releaseBuffer(br *strings.Builder) {
	br.Reset()
}

var spaces = regexp.MustCompile(`\s+`)

func cleanAllSpaces(content string) string {
	return spaces.ReplaceAllString(content, "")
}

func wrapWithBaseFunction(funcName string, typeName string, b *strings.Builder) string {
	var t = getBuffer()
	defer releaseBuffer(t)

	t.WriteString(`func `)
	t.WriteString(funcName)
	t.WriteString(`(page *peji.Page, `)
	t.WriteString(`ctx `)
	t.WriteString(typeName)
	t.WriteString(`, rootDoc *domu.Node) {`)
	t.WriteString("\n")
	t.WriteString(b.String())
	t.WriteString("\n")
	t.WriteString(`}`)

	return t.String()
}

func wrapWithArgumentFunction(funcName string, arguments []string, b *strings.Builder) (string, error) {
	var t = getBuffer()
	defer releaseBuffer(t)

	t.WriteString(`func `)
	t.WriteString(funcName)
	t.WriteString(`( `)

	var lastIndex = len(arguments) - 1
	for index, arg := range arguments {
		t.WriteString(arg)
		if lastIndex != index {
			t.WriteString(", ")
		}
	}

	t.WriteString(`) *domu.Node {`)
	t.WriteString("\n")
	t.WriteString(b.String())
	t.WriteString("\n")
	t.WriteString(`}`)

	var buf = nunsafe.String2Bytes(b.String())
	var res, err = format.Source(buf)
	return nunsafe.Bytes2String(res), err
}

type varType int

const (
	singleType varType = iota
	doubleType
)

type vard interface {
	Type() varType
}

type doubleVar struct {
	Left  string
	Right string
}

func (s doubleVar) Type() varType {
	return doubleType
}

type singleVar struct {
	Name string
}

func (s singleVar) Type() varType {
	return singleType
}

type varQueue struct {
	queue []vard
}

func (sq *varQueue) peek() vard {
	var index = len(sq.queue) - 1
	if index < 0 {
		index = 0
	}
	return sq.queue[index]
}

func (sq *varQueue) beforeLast() vard {
	var index = len(sq.queue) - 2
	if index < 0 {
		index = 0
	}
	return sq.queue[index]
}

func (sq *varQueue) first() vard {
	return sq.queue[0]
}

func (sq *varQueue) pop() vard {
	var cm = sq.peek()
	sq.queue = sq.queue[0 : len(sq.queue)-1]
	return cm
}

func (sq *varQueue) push(sm vard) {
	sq.queue = append(sq.queue, sm)
}

func (sq *varQueue) pushSingle(name string) {
	sq.queue = append(sq.queue, singleVar{Name: name})
}

func (sq *varQueue) pushDouble(left, right string) {
	sq.queue = append(sq.queue, doubleVar{Left: left, Right: right})
}

type CommandPipe struct {
	RootCommander *CommandPipe
	Root          *parse.CommandNode
	FuncName      *parse.IdentifierNode
	MoreFields    []parse.Node
	Arguments     []parse.NodeTyper
	NoAttaching   bool
	insertText    string
}

func (c *CommandPipe) Type() parse.NodeType {
	return -1
}

func (c *CommandPipe) Prepare(set []*parse.CommandNode) error {
	var totalItems = len(set)
	if err := c.PrepareCommand(set[totalItems-1], set[0:totalItems-1]); err != nil {
		return nerror.WrapOnly(err)
	}
	return nil
}

func (c *CommandPipe) PrepareCommand(cmd *parse.CommandNode, others []*parse.CommandNode) error {
	c.Root = cmd

	var argumentList = cmd.Args
	var argumentSize = len(cmd.Args)
	if argumentSize > 0 {
		switch firstNode := argumentList[0].(type) {
		case *parse.IdentifierNode:
			c.FuncName = firstNode
			c.NoAttaching = firstNode.Ident == "noop"
			argumentList = argumentList[1:]
		case *parse.ImportedFieldFunc:
			if firstNode.IsFunc() {
				c.FuncName = firstNode.ToIdentifier()
				c.MoreFields = firstNode.Fields
				argumentList = argumentList[1:]
			}
		}
	}

	var otherSize = len(others)
	if c.RootCommander != nil && c.FuncName != nil && c.FuncName.Ident == "noop" && otherSize > 0 {
		return nerror.New("A noop function call must be the last in the list of pipe calls, mixing isn't allowed")
	}

	if otherSize > 0 {
		var nextPipe = new(CommandPipe)
		nextPipe.RootCommander = c

		// pass along text for content
		if c.NoAttaching {
			nextPipe.insertText = c.insertText
		}

		if err := nextPipe.Prepare(others); err != nil {
			return nerror.WrapOnly(err)
		}
		c.Arguments = append(c.Arguments, nextPipe)
	}

	for _, item := range argumentList {
		if importedField, ok := item.(*parse.ImportedFieldFunc); ok {
			c.Arguments = append(c.Arguments, item)

			for _, field := range importedField.Fields {
				c.Arguments = append(c.Arguments, field)
			}
			continue
		}
		c.Arguments = append(c.Arguments, item)
	}

	return nil
}

func (c *CommandPipe) isSingle() bool {
	return len(c.insertText) == 0 && len(c.Arguments) == 0 && len(c.MoreFields) == 0
}

func (c *CommandPipe) Parse(t *TextWriter, tree *parse.Tree) error {
	var tmp2 = getBuffer()
	defer releaseBuffer(tmp2)

	var oldBuilder = t.swapBuilder(tmp2)
	defer t.swapBuilder(oldBuilder)

	if err := c.ParseArgs(t, tree); err != nil {
		return nerror.WrapOnly(err)
	}

	oldBuilder.WriteString(tmp2.String())
	return nil
}

func (c *CommandPipe) ParseArgs(t *TextWriter, tree *parse.Tree) error {
	var fieldBuffer = getBuffer()
	defer releaseBuffer(fieldBuffer)

	var oldBuilder = t.swapBuilder(fieldBuffer)
	defer t.swapBuilder(oldBuilder)

	var fieldSize = len(c.MoreFields)
	for fieldIndex, field := range c.MoreFields {
		if err := t.walk(tree, field, c.Root); err != nil {
			return nerror.WrapOnly(err)
		}

		if fieldIndex < fieldSize-1 {
			t.Write(", ")
		}
	}

	var tmp = getBuffer()
	defer releaseBuffer(tmp)
	t.swapBuilder(tmp)

	var err error
	var totalArgs = len(c.Arguments)
	for index, argument := range c.Arguments {
		switch realArgument := argument.(type) {
		case *CommandPipe:
			err = realArgument.Parse(t, tree)
		case *parse.ImportedFieldFunc:
			if realArgument.IsField() {
				err = t.walkTextIdentifier(tree, realArgument.ToTextIdentifier(), c.Root)
			}
			if realArgument.IsFunc() {
				err = nerror.New("Encountering imported func here which should have already being handled: %#q\n", realArgument)
			}
		case *parse.StringNode:
			err = t.walkStringNode(tree, realArgument, c.Root)
		case *parse.VariableNode:
			err = t.walkVariableIdentifier(tree, realArgument, c.Root)
		case *parse.DOMIdentifier:
			err = t.walkDOMIdentifier(tree, realArgument, c.Root)
		case *parse.IdentifierNode:
			err = t.walkIdentifier(tree, realArgument, c.Root)
		case *parse.FieldNode:
			err = t.walkField(tree, realArgument, c.Root)
		case *parse.DotNode:
			err = t.walkDotValue(tree, realArgument, c.Root)
		case parse.Node:
			err = t.walk(tree, realArgument, c.Root)
		}

		if err != nil {
			break
		}

		if index != (totalArgs - 1) {
			t.Write(", ")
		}
	}

	if err != nil {
		return nerror.WrapOnly(err)
	}

	if c.FuncName != nil && c.FuncName.Ident != "noop" {
		if c.FuncName.IsInbuiltFunc {
			t.addImport(helperAlias, helperImportPath)
		}

		oldBuilder.WriteString(c.FuncName.Ident)

		if c.isSingle() {
			return nil
		}

		oldBuilder.WriteString("(")

		if len(c.insertText) > 0 {
			oldBuilder.WriteString(c.insertText)
			oldBuilder.WriteString(",")
		}

		oldBuilder.WriteString(tmp.String())

		if fieldBuffer.Len() > 0 {
			if tmp.Len() > 0 {
				oldBuilder.WriteString(", ")
			}
			oldBuilder.WriteString(fieldBuffer.String())
		}

		oldBuilder.WriteString(")")
		return nil
	}

	oldBuilder.WriteString(tmp.String())
	return nil
}

func parseTemplateName(name string) (string, string) {
	var targetBlock = name
	if hashedIndex := strings.Index(name, "#"); hashedIndex != -1 {
		targetBlock = name[hashedIndex+1:]
		name = name[:hashedIndex]
	}
	return name, targetBlock
}
