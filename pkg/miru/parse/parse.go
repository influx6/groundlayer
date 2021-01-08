// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// WARNING:
// The code has dramatically changed compared to the current go
// implementation as of August 2020. Hence certain things
// will not match previous behaviour, like texts.
// Due to the need to optimize how they are collected,
// their positions are thrown out and they are merged into
// a single whole across both text and html parsing.

// Package parse builds parse trees for templates as defined by text/template
// and html/template. Clients should use those packages to construct templates
// rather than this one, which provides shared internal data structures not
// intended for general use.
package parse

import (
	"bytes"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/influx6/npkg/nerror"
)

var (
	allSpace          = regexp.MustCompile(`\s+`)
	allLineSpacesTabs = regexp.MustCompile(`[\n\f\r\t\[\]]+`)
	modelTypeBase     = regexp.MustCompile(`([0-9a-zA-Z]+)\(([0-9a-zA-Z]+),?([0-9a-zA-Z]+)?\)`)
)

const (
	themeAttrName = "theme"
)

var domFunctions = map[string]bool{
	"dom":       true,
	"rootdom":   true,
	"parentdom": true,
}

// IsDomFuncs returns true/false if the name is a dom function.
func IsDomFuncs(name string) bool {
	return domFunctions[name]
}

// TextList implements a wrapper for a list of TextNode.
type TextList []*TextNode

func (t *TextList) add(tm *TextNode) {
	*t = append(*t, tm)
}

func (t *TextList) print(l *ListNode, tt *Tree) {
	if len(*t) == 0 {
		return
	}
	var newTextNode = tt.newText(0, t.String(), tt.withinTheme)
	l.append(newTextNode)
}

func (t TextList) String() string {
	var bs strings.Builder
	for _, txn := range t {
		bs.Write(txn.Text)
	}
	return bs.String()
}

// FuncMaps provides a wrapper of fields.
type FuncMaps struct {
	Funcs         map[string]string
	ModelSet      map[string]*ModelNode
	ModelTypeSet  map[string]*ModelTypeNode
	MethodSet     map[string]*MethodNode
	DefinedBlocks map[string]string
}

// Tree is the representation of a single parsed template.
type Tree struct {
	FuncMaps
	badErr         error
	seenEOF        bool
	FromDefinition bool
	Name           string // name of the template represented by the tree.
	ParseName      string // name of the top-level template during parsing, for error messages.
	TypeName       string
	Root           *ListNode // top-level root of the tree.
	textCache      []*TextNode
	text           string // text parsed to create the template (or its parent)
	// Parsing only; cleared after parse.
	lex           *lexer
	token         [3]item // three-token lookahead for parser.
	peekCount     int
	vars          []string // variables defined at the moment.
	treeSet       map[string]*Tree
	currentMethod *MethodNode
	rootNode      *RootNode
	seenRighDelim bool
	withinHTML    bool
	withinTheme   bool
}

// Parsing.

// New allocates a new parse tree with the given name.
func New(name string) *Tree {
	return &Tree{
		Name:     name,
		TypeName: "interface{}",
		FuncMaps: FuncMaps{
			Funcs:         map[string]string{},
			MethodSet:     map[string]*MethodNode{},
			ModelTypeSet:  map[string]*ModelTypeNode{},
			ModelSet:      map[string]*ModelNode{},
			DefinedBlocks: map[string]string{},
		},
	}
}

// NewWithParseMap allocates a new parse tree with the given name.
func NewWithParseMap(name string, maps FuncMaps) *Tree {
	return &Tree{
		FuncMaps: maps,
		Name:     name,
		TypeName: "interface{}",
	}
}

// Copy returns a copy of the Tree. Any parsing state is discarded.
func (t *Tree) Copy() *Tree {
	if t == nil {
		return nil
	}
	return &Tree{
		Name:      t.Name,
		ParseName: t.ParseName,
		Root:      t.Root.CopyList(),
		text:      t.text,
		TypeName:  t.TypeName,
	}
}

// Parse returns a map from template name to parse.Tree, created by parsing the
// templates described in the argument string. The top-level template will be
// given the specified name. If an error is encountered, parsing stops and an
// empty map is returned with the error.
func Parse(name, text, leftDelim, rightDelim string, maps FuncMaps, treeSet map[string]*Tree) (map[string]*Tree, error) {
	if treeSet == nil {
		treeSet = make(map[string]*Tree)
	}

	t := NewWithParseMap(name, maps)
	t.text = text
	_, err := t.Parse(text, leftDelim, rightDelim, treeSet, maps)
	return treeSet, err
}

// next returns the next token.
func (t *Tree) next() item {
	if t.peekCount > 0 {
		t.peekCount--
	} else {
		t.token[0] = t.lex.nextItem()
	}
	return t.token[t.peekCount]
}

// backup backs the input stream up one token.
func (t *Tree) backup() {
	t.peekCount++
}

// backup2 backs the input stream up two tokens.
// The zeroth token is already there.
func (t *Tree) backup2(t1 item) {
	t.token[1] = t1
	t.peekCount = 2
}

// backup3 backs the input stream up three tokens
// The zeroth token is already there.
func (t *Tree) backup3(t2, t1 item) { // Reverse order: we're pushing back.
	t.token[1] = t1
	t.token[2] = t2
	t.peekCount = 3
}

// peek returns but does not consume the next token.
func (t *Tree) peek() item {
	if t.peekCount > 0 {
		return t.token[t.peekCount-1]
	}
	t.peekCount = 1
	t.token[0] = t.lex.nextItem()
	return t.token[0]
}

// nextNonSpace returns the next non-space token.
func (t *Tree) nextNonSpace() (token item) {
	for {
		token = t.next()
		if token.typ != itemSpace {
			break
		}
	}
	return token
}

// peekNonSpace returns but does not consume the next non-space token.
func (t *Tree) peekNonSpace() item {
	token := t.nextNonSpace()
	t.backup()
	return token
}

// ErrorContext returns a textual representation of the location of the node in the input text.
// The receiver is only used when the node does not have a pointer to the tree inside,
// which can occur in old code.
func (t *Tree) ErrorContext(n Node) (location, context string) {
	pos := int(n.Position())
	tree := n.tree()
	if tree == nil {
		tree = t
	}
	text := tree.text[:pos]
	byteNum := strings.LastIndex(text, "\n")
	if byteNum == -1 {
		byteNum = pos // On first line.
	} else {
		byteNum++ // After the newline.
		byteNum = pos - byteNum
	}
	lineNum := 1 + strings.Count(text, "\n")
	context = n.String()
	return fmt.Sprintf("%s:%d:%d", tree.ParseName, lineNum, byteNum), context
}

// errorf formats the error and terminates processing.
func (t *Tree) errorf(format string, args ...interface{}) {
	t.Root = nil
	format = fmt.Sprintf("template: %s:%d: %s", t.ParseName, t.token[0].line, format)
	panic(nerror.New(format, args...))
}

// error terminates processing.
func (t *Tree) error(err error) {
	t.errorf("%s", err)
}

// expectTypeOrRightDelimiter consumes the next token and guarantees it has the required type.
func (t *Tree) expectTypeOrRightDelimiter(context string) item {
	token := t.nextNonSpace()
	switch token.typ {
	case itemText, itemRawString, itemString, itemIdentifier:
		var err error
		var typeName = token.val
		if strings.Contains(typeName, "\"") || strings.Contains(typeName, "'") {
			typeName, err = strconv.Unquote(token.val)
			if err != nil {
				t.errorf("Failed to unquote string %q -> %q", token.val, err.Error())
			}
		}
		t.TypeName = typeName
	case itemRightDelim:
		// do nothing
		return token
	default:
		t.unexpected(token, context)
	}

	// return token
	return t.expect(itemRightDelim, context)
}

// expect consumes the next token and guarantees it has the required type.
func (t *Tree) expect(expected itemType, context string) item {
	token := t.nextNonSpace()
	if token.typ != expected {
		t.unexpected(token, context)
	}
	return token
}

// expectOneOf consumes the next token and guarantees it has one of the required types.
func (t *Tree) expectOneOf(expected1, expected2 itemType, context string) item {
	token := t.nextNonSpace()
	if token.typ != expected1 && token.typ != expected2 {
		t.unexpected(token, context)
	}
	return token
}

// unexpected complains about the token and terminates processing.
func (t *Tree) unexpected(token item, context string) {
	t.errorf("unexpected %s in %s", token, context)
}

// recover is the handler that turns panics into returns from the top level of Parse.
func (t *Tree) recover(errp *error) {
	e := recover()
	if e != nil {
		if _, ok := e.(runtime.Error); ok {
			panic(e)
		}
		if t != nil {
			t.lex.drain()
			t.stopParse()
		}

		*errp = e.(error)
	}
}

// startParse initializes the parser, using the lexer.
func (t *Tree) startParse(maps FuncMaps, lex *lexer, treeSet map[string]*Tree) {
	t.Root = nil
	t.lex = lex
	t.vars = []string{"$"}
	t.FuncMaps = maps
	t.treeSet = treeSet
}

// stopParse terminates parsing.
func (t *Tree) stopParse() {
	t.lex = nil
	t.vars = nil
	t.Funcs = nil
	t.treeSet = nil
	t.ModelSet = nil
	t.MethodSet = nil
	t.ModelTypeSet = nil
}

// Parse parses the template definition string to construct a representation of
// the template for execution. If either action delimiter string is empty, the
// default ("{{" or "}}") is used. Embedded template definitions are added to
// the treeSet map.
func (t *Tree) Parse(text, leftDelim, rightDelim string, treeSet map[string]*Tree, maps FuncMaps) (tree *Tree, err error) {
	defer t.recover(&err)

	t.ParseName = t.Name
	t.startParse(maps, lex(t.Name, text, leftDelim, rightDelim), treeSet)
	t.text = text
	t.parse()
	t.add()
	t.stopParse()

	tree = t
	return
}

// add adds tree to t.treeSet.
func (t *Tree) add() {
	tree := t.treeSet[t.Name]
	if tree == nil || IsEmptyTree(tree.Root) {
		t.treeSet[t.Name] = t
		return
	}

	if !IsEmptyTree(t.Root) {
		t.errorf("template: multiple definition of template %q, maybe you named a file and a definition block the same name or your files containing define blocks have contents outside of define blocks, only the root file can have those ?", t.Name)
	}
}

// IsEmptyTree reports whether this tree (node) is empty of everything but space.
func IsEmptyTree(n Node) bool {
	switch n := n.(type) {
	case nil:
		return true
	case *ActionNode:
	case *IfNode:
	case *ListNode:
		for _, node := range n.Nodes {
			if !IsEmptyTree(node) {
				return false
			}
		}
		return true
	case *RangeNode:
	case *TemplateNode:
	case *TextNode:
		return len(bytes.TrimSpace(n.Text)) == 0
	case *WithNode:
	case *HTMLNode:
		for _, node := range n.Children.Nodes {
			if !IsEmptyTree(node) {
				return false
			}
		}
		return true
	default:
		panic(fmt.Sprintf("unknown node: %#v", n))
	}
	return false
}

// parse is the top-level parser for a template, essentially the same
// as itemList except it also parses {{define}} actions.
// It runs to EOF.
func (t *Tree) parse() {
	var textList TextList

	t.Root = t.newList(t.peek().pos)
	for t.peek().typ != itemEOF {
		if t.peek().typ == itemLeftDelim {
			delim := t.next()
			nextToken := t.nextNonSpace()
			if nextToken.typ == itemDefine {
				newT := New("definition") // name will be updated once we know it.
				newT.text = t.text
				newT.ParseName = t.ParseName
				newT.FromDefinition = true
				newT.startParse(t.FuncMaps, t.lex, t.treeSet)
				newT.parseDefinition()
				continue
			}
			t.backup2(delim)
		}

		n := t.textOrAction(false)
		switch n.Type() {
		case NodeKomponent:
		case nodeEnd, nodeElse:
			t.errorf("unexpected %s", n)
		case nodeFinish:
			return
		case NodeText:
			textList.add(n.(*TextNode))
		default:
			textList.print(t.Root, t)
			textList = textList[:0]

			t.Root.append(n)
		}
	}
}

// textOrAction:
//	text | action
func (t *Tree) textOrAction(allowSpace bool) Node {
	var token item
	if allowSpace {
		token = t.next()
	} else {
		token = t.nextNonSpace()
	}

	switch token.typ {
	case itemTag:
		nn := t.html(token)
		return nn
	case itemSpace:
		if allowSpace {
			return t.newText(token.pos, token.val, t.withinTheme)
		}
	case itemText:
		return t.newText(token.pos, token.val, t.withinTheme)
	case itemLeftDelim:
		return t.action(false)
	case itemEOF:
		return t.newFinish(token.pos)
	case itemIgnore:
		t.errorf("Should not receive ignore instruction for token %#v at line %d, did you forget to properly end a statement", token, token.pos)
	default:
		if t.withinTheme {
			return t.newText(token.pos, token.val, t.withinTheme)
		}
		t.unexpected(token, "input")
	}
	return nil
}

// HTML:
func (t *Tree) html(token item) (n Node) {
	var kids = new(ListNode)
	kids.NodeType = NodeKids

	var attr AttrList
	var tagNode = t.newHTMLTag(token.pos, token.line, token.val, attr, kids)

	var textList TextList

	t.withinHTML = true
	defer func() {
		t.withinHTML = false
	}()

	var unfinishedHtml = false
	var badToken item

htmlLoop:
	for {
		var nextToken = t.next()

		switch nextToken.typ {
		case itemIdentifier:
		case itemTagAttrStart:
			var newAttr = t.attr(nextToken)
			tagNode.Attr.append(newAttr)
		case itemText, itemSpace:
			var tx = t.newText(nextToken.pos, nextToken.val, false)
			textList.add(tx)
		case itemTag:
			textList.print(tagNode.Children, t)
			textList = textList[:0]

			tagNode.Children.append(t.html(nextToken))
			if t.seenEOF {
				unfinishedHtml = true
				badToken = nextToken
				break htmlLoop
			}
		case itemTagSelfEnd:
			textList.print(tagNode.Children, t)
			textList = textList[:0]

			tagNode.Closing = false
			break htmlLoop
		case itemTagEnd:
			textList.print(tagNode.Children, t)
			textList = textList[:0]

			tagNode.Closing = true
			break htmlLoop
		case itemMethodEnd:
			t.errorf("Tag was not properly closed as we see a endMethod")
		case itemLeftDelim:
			textList.print(tagNode.Children, t)
			textList = textList[:0]

			var node = t.action(false)
			if node.Type() == NodeKomponent {
				continue htmlLoop
			}

			if node.Type() != nodeEnd && node.Type() != nodeElse && node.Type() != nodeMethodEnd {
				tagNode.Children.append(node)
				continue htmlLoop
			}
			if node.Type() == nodeMethodEnd {
				t.errorf("Found {{endMethod}} occurring during html parsing, did you forget to close a tag ?")
			}
		case itemEOF:
			t.seenEOF = true
			break htmlLoop
		}
	}

	if unfinishedHtml {
		t.unexpected(badToken, fmt.Sprintf("tag %q was not properly closed at line %d and pos %d", tagNode.Tag, badToken.line, badToken.pos))
	}

	return tagNode
}

func (t *Tree) attr(token item) *AttrNode {
	newAttr := t.newAttr(token.pos, token.val, token.val, new(ListNode))
	newAttr.IsTheme = strings.TrimSpace(token.val) == themeAttrName

	t.withinTheme = newAttr.IsTheme

	var acceptSpaces = false
attrLoop:
	for {
		var nextToken = t.next()

		switch nextToken.typ {
		case itemTagAttrValueStart:
			acceptSpaces = true
		case itemTagAttrEnd:
			break attrLoop
		case itemList:
			if nextToken.val == "" {
				continue attrLoop
			}

			var text = strings.TrimSpace(allLineSpacesTabs.ReplaceAllString(nextToken.val, " "))
			for _, part := range strings.Split(text, " ") {
				var newText = t.newText(nextToken.pos, part, newAttr.IsTheme)
				newAttr.Values.append(newText)
			}
		case itemText:
			if nextToken.val == "" {
				continue attrLoop
			}
			if strings.TrimSpace(nextToken.val) == "" && newAttr.IsTheme {
				continue attrLoop
			}
			var newText = t.newText(nextToken.pos, nextToken.val, newAttr.IsTheme)
			newAttr.Values.append(newText)
		case itemSpace:
			if nextToken.val == "" {
				continue attrLoop
			}
			if acceptSpaces && !newAttr.IsTheme {
				newAttr.Values.append(t.newText(nextToken.pos, nextToken.val, newAttr.IsTheme))
			}
		case itemEOF:
			break attrLoop
		case itemLeftDelim:
			var node = t.action(false)
			if node.Type() == NodeKomponent {
				continue attrLoop
			}
			if node.Type() != nodeEnd && node.Type() != nodeElse && node.Type() != nodeMethodEnd {
				newAttr.Values.append(node)
				continue attrLoop
			}
			if node.Type() == nodeMethodEnd {
				t.errorf("Found {{endMethod}} occurring during html parsing, did you forget to close a tag ?")
				break attrLoop
			}
		default:
			// newAttr.Values.append(t.newText(nextToken.pos, nextToken.val))
		}
	}

	t.withinTheme = false
	return newAttr
}

func (t *Tree) methods(token item) *MethodNode {
	node := t.newMethod(token.pos, token.val, nil, new(ListNode))
	if _, ok := t.MethodSet[token.val]; ok {
		t.errorf("it seems a methodSet with given name %q already exist, could you rename if not a duplicate ?", token.val)
	}
	t.MethodSet[token.val] = node

	t.currentMethod = node
	defer func() {
		t.currentMethod = nil
	}()

	var textList TextList

handleLoop:
	for {
		var nextToken = t.nextNonSpace()

		switch nextToken.typ {
		case itemModelField:
			textList.print(node.Body, t)
			textList = textList[:0]

			node.appendArgument(t.modelField(nextToken))
		case itemMethod, itemModelType, itemModel, itemDefine, itemTemplate:
			t.errorf("Nested methods, model or modelTypes, or  define, template blocks are not supported")
		case itemMethodEnd:
			textList.print(node.Body, t)
			textList = textList[:0]

			t.expect(itemRightDelim, "endMethod")
			break handleLoop
		case itemEOF:
			textList.print(node.Body, t)
			textList = textList[:0]

			t.backup()
			break handleLoop
		case itemTag:
			textList.print(node.Body, t)
			textList = textList[:0]

			node.Body.append(t.html(nextToken))
		case itemText:
			textList.add(t.newText(nextToken.pos, nextToken.val, t.withinTheme))
		case itemLeftDelim:
			textList.print(node.Body, t)
			textList = textList[:0]

			var actionResult = t.action(true)
			if actionResult.Type() != nodeEnd && actionResult.Type() != nodeElse && actionResult.Type() != nodeMethodEnd {
				node.Body.append(actionResult)
				continue handleLoop
			}
			if actionResult.Type() == nodeMethodEnd {
				break handleLoop
			}
		default:
		}
	}
	return node
}

func (t *Tree) rootType(token item) *RootNode {
	if t.rootNode != nil {
		t.errorf("only one RootType is allowed in all parsing of a file", token.val)
		return nil
	}

	var err error
	var content = token.val
	if strings.Contains(content, "\"") || strings.Contains(content, "'") {
		content, err = strconv.Unquote(token.val)
		if err != nil {
			t.errorf("Failed to unquote string %q -> %q", token.val, err.Error())
		}
	}

	var typeValue = t.getTypeName(content)
	var rootNodeType = t.newRootNode(token.pos, token.val, typeValue)
	t.rootNode = rootNodeType
	t.TypeName = typeValue

	return rootNodeType
}

func (t *Tree) modelType(token item) *ModelTypeNode {
	node := t.newModelType(token.pos, token.val, nil)
	if _, ok := t.ModelTypeSet[token.val]; ok {
		t.errorf("it seems a modelType with given name %q already exist, could you rename if not a duplicate ?", token.val)
	}
	t.ModelTypeSet[token.val] = node

handleLoop:
	for {
		var nextToken = t.nextNonSpace()
		switch nextToken.typ {
		case itemModelTypeBase:
			node.Detail = t.modelTypeBase(nextToken)
		case itemMethod, itemModelType, itemModel:
			t.errorf("Nested methods, model or modelTypes are not supported")
		case itemEnd, itemEOF, itemModelTypeEnd:
			break handleLoop
		}
	}

	return node
}

func (t *Tree) modelTypeBase(token item) Node {
	var tokenValue = allSpace.ReplaceAllString(token.val, "")
	if !modelTypeBase.MatchString(tokenValue) {
		return t.getRefType(token, tokenValue)
	}

	var collectionTypeList = modelTypeBase.FindStringSubmatch(tokenValue)
	var collectionName = collectionTypeList[1]
	var collectionKey = collectionTypeList[2]
	var collectionValue = collectionTypeList[3]

	var collectionNameType, isCollection = collectionType[collectionName]

	if !isCollection {
		t.errorf("not yet supporting collection %q type", collectionName)
	}

	var nodes = new(ListNode)
	switch collectionNameType {
	case ListType:
		nodes.append(t.getRefType(token, collectionKey))
	case MapType:
		if collectionValue == "" {
			t.errorf("Map require a value type as an empty is not support")
		}

		nodes.append(t.getRefType(token, collectionKey))
		nodes.append(t.getRefType(token, collectionValue))
	}

	return t.newCollection(token.pos, collectionNameType, nodes)
}

func (t *Tree) importDeclr(token item) *ImportDeclr {
	node := t.newImportDeclr(token.pos, token.val, "")

handleLoop:
	for {
		var nextToken = t.nextNonSpace()
		switch nextToken.typ {
		case itemImportDeclarationEnd:
			node.Path = nextToken.val
			break handleLoop
		case itemEnd, itemEOF:
			break handleLoop
		default:
			t.errorf("Nested methods, model or modelTypes are not supported")
		}
	}
	return node
}

func (t *Tree) mountLiveType(token item) *MountLiveNode {
	var node = t.newMountLiveNode(token.pos, strings.TrimSpace(token.val), "", nil, token.typ == itemMountLiveList)

handleLoop:
	for {
		var nextToken = t.nextNonSpace()
		switch nextToken.typ {
		case itemKomponentName:
			node.Name = strings.TrimSpace(nextToken.val)
		case itemBool, itemImportedFieldFunc, itemDOMIdentifier, itemCharConstant, itemComplex, itemDot, itemField, itemIdentifier,
			itemNumber, itemNil, itemRawString, itemString, itemVariable:
			t.backup()
			node.Target = t.command()
			continue handleLoop
		case itemEnd, itemEOF, itemRightDelim:
			break handleLoop
		case itemLeftDelim:
			t.backup()
			break handleLoop
		default:
			t.unexpected(nextToken, "usage of mount, mount only supports dot, variables, strings, numbers and imported fields")
			break handleLoop
		}
	}
	return node
}

func (t *Tree) mountType(token item) *MountNode {
	var node = t.newMountNode(token.pos, strings.TrimSpace(token.val), nil, token.typ == itemMountList)

handleLoop:
	for {
		var nextToken = t.nextNonSpace()
		switch nextToken.typ {
		case itemBool, itemImportedFieldFunc, itemDOMIdentifier, itemCharConstant, itemComplex, itemDot, itemField, itemIdentifier,
			itemNumber, itemNil, itemRawString, itemString, itemVariable:
			t.backup()
			node.Target = t.command()
			continue handleLoop
		case itemEnd, itemEOF, itemRightDelim:
			break handleLoop
		case itemLeftDelim:
			t.backup()
			break handleLoop
		default:
			t.unexpected(nextToken, "usage of mount, mount only supports dot, variables, strings, numbers and imported fields")
			break handleLoop
		}
	}
	return node
}

func (t *Tree) importedFieldFunc(token item) *ImportedFieldFunc {
	var node *ImportedFieldFunc

handleLoop:
	for {
		var nextToken = t.nextNonSpace()
		switch nextToken.typ {
		case itemImportedFunc:
			node = t.newImportedFunc(token.pos, token.val)
			break handleLoop
		case itemImportedField:
			node = t.newImportedField(token.pos, token.val)
			break handleLoop
		case itemRightDelim:
			t.backup()
			break handleLoop
		default:
			t.errorf("unexpected token type %#v", nextToken)
		}
	}

	if fields := t.pipeFieldDOMAndVars(token); len(fields) > 0 {
		node.Fields = fields
	}

	node.Commands = t.pipekids(token)
	return node
}

// Pipeline:
//	declarations? command (command Field Field)*
func (t *Tree) pipeFieldDOMAndVars(token item) []Node {
	var nodes []Node

handleLoop:
	for {
		token := t.nextNonSpace()
		switch token.typ {
		case itemDOMIdentifier:
			nodes = append(nodes, t.newDOMIdentifer(token.pos, token.val))
		case itemField:
			nodes = append(nodes, t.newField(token.pos, token.val))
		case itemString:
			s, err := strconv.Unquote(token.val)
			if err != nil {
				t.error(err)
			}
			nodes = append(nodes, t.newString(token.pos, token.val, s))
		case itemVariable:
			nodes = append(nodes, t.newVariable(token.pos, token.val))
		default:
			t.backup()
			break handleLoop
		}
	}
	return nodes
}

// Pipeline:
//	declarations? command ('|' command)*
func (t *Tree) pipekids(token item) []*CommandNode {
	var commands []*CommandNode

handleLoop:
	for {
		token := t.nextNonSpace()
		switch token.typ {
		case itemEOF:
			t.backup()
			break handleLoop
		case itemRightDelim, itemRightParen:
			// At this point, the pipeline is complete
			if token.typ == itemRightParen {
				t.backup()
			}

			break handleLoop
		case itemPipe:
		case itemBool, itemImportedFieldFunc, itemCharConstant, itemComplex, itemDot, itemField, itemDOMIdentifier, itemIdentifier,
			itemNumber, itemNil, itemRawString, itemString, itemVariable, itemLeftParen:
			t.backup()
			commands = append(commands, t.command())
		default:
			t.unexpected(token, "pipekids")
		}
	}

	return commands
}

func (t *Tree) models(token item) *ModelNode {
	node := t.newModel(token.pos, token.val, new(ListNode))
	if _, ok := t.ModelSet[token.val]; ok {
		t.errorf("it seems a model with given name %q already exist, could you rename if not a duplicate ?", token.val)
	}
	t.ModelSet[token.val] = node

handleLoop:
	for {
		var nextToken = t.nextNonSpace()
		switch nextToken.typ {
		case itemModelField:
			node.Fields.append(t.modelField(nextToken))
		case itemMethod, itemModelType, itemModel:
			t.errorf("Nested methods, model or modelTypes are not supported")
		case itemEnd, itemEOF, itemModelEnd:
			break handleLoop
		}
	}
	return node
}

func (t *Tree) modelField(token item) *RefNode {
	var tokenValue = allSpace.ReplaceAllString(token.val, "")
	var fieldNameAndTypeList = strings.Split(tokenValue, ":")
	var fieldName = fieldNameAndTypeList[0]
	var fieldType = fieldNameAndTypeList[1]

	var detail = t.getRefType(token, fieldType)
	return t.newRefNode(token.pos, fieldName, t.getTypeName(fieldType), detail)
}

func (t *Tree) getTypeName(fieldType string) string {
	if baseRep, ok := baseTypesRepresentation[fieldType]; ok {
		return baseRep
	}
	return fieldType
}

func (t *Tree) getRefType(token item, tokenFieldString string) Node {
	var fieldType = allSpace.ReplaceAllString(tokenFieldString, "")

	var isBaseType = baseTypes[fieldType]
	if !isBaseType {
		isBaseType = baseTypes[capitalize(fieldType)]
	}
	var customModel, isCustomModel = t.ModelSet[fieldType]
	var customModelType, isCustomModelType = t.ModelTypeSet[fieldType]

	if !isBaseType && !isCustomModel && !isCustomModelType {
		t.errorf("Model field type %q must already be defined, see line %d in %d", fieldType, token.line, token.pos)
	}

	var detail Node
	switch true {
	case isBaseType:
		var fieldTypeName = t.getTypeName(fieldType)
		detail = t.newBase(token.pos, fieldTypeName)
	case isCustomModel:
		detail = customModel
	case isCustomModelType:
		detail = customModelType
	}

	return detail
}

// parseDefinition parses a {{define}} ...  {{end}} template definition and
// installs the definition in t.treeSet. The "define" keyword has already
// been scanned.
func (t *Tree) parseDefinition() {
	const context = "define clause"

	name := t.expectOneOf(itemString, itemRawString, context)
	var err error
	t.Name, err = strconv.Unquote(name.val)
	if err != nil {
		t.error(err)
	}

	t.expectTypeOrRightDelimiter(context)
	// t.expect(itemRightDelim, context)

	var end Node
	t.Root, end = t.itemList()
	if end.Type() != nodeEnd {
		t.errorf("unexpected %s in %s", end, context)
	}
	t.add()
	t.stopParse()
}

// itemList:
//	textOrAction*
// Terminates at {{end}} or {{else}}, returned separately.
func (t *Tree) itemList() (list *ListNode, next Node) {
	var textList TextList

	list = t.newList(t.peek().pos)

doloop:
	for t.peek().typ != itemEOF {
		n := t.textOrAction(true)

		switch n.Type() {
		case NodeKomponent:
			continue doloop
		case NodeText:
			textList.add(n.(*TextNode))
			continue doloop
		case nodeEnd, nodeElse, nodeFinish:
			textList.print(list, t)
			textList = textList[:0]
			return list, n
		}

		textList.print(list, t)
		textList = textList[:0]

		list.append(n)
	}
	t.errorf("unexpected EOF")
	return
}

// Action:
//	control
//	command ("|" command)*
// Left delim is past. Now get actions.
// First word could be a keyword such as range.
func (t *Tree) action(inMethod bool) Node {
	token := t.nextNonSpace()
	switch token.typ {
	case itemImportDeclaration:
		return t.importDeclr(token)
	case itemImportedFieldFunc:
		return t.importedFieldFunc(token)
	case itemMethodEnd:
		return t.endMethodControl()
	case itemBlock:
		return t.blockControl()
	case itemElse:
		return t.elseControl()
	case itemEnd:
		return t.endControl()
	case itemModel:
		return t.models(token)
	case itemRoot:
		return t.rootType(token)
	case itemMountLive, itemMountLiveList:
		return t.mountLiveType(token)
	case itemMount, itemMountList:
		return t.mountType(token)
	case itemModelType:
		return t.modelType(token)
	case itemMethod:
		return t.methods(token)
	case itemIf:
		return t.ifControl()
	case itemRange:
		return t.rangeControl()
	case itemTemplate:
		return t.templateControl()
	case itemWith:
		return t.withControl()
	}

	t.backup()
	var nextToken = t.peek()
	// Do not pop variables; they persist until "end".
	var actionNode = t.newAction(nextToken.pos, nextToken.line, t.pipeline("command"))
	return actionNode
}

// Pipeline:
//	declarations? command ('|' command)*
func (t *Tree) pipeline(context string) (pipe *PipeNode) {
	token := t.peekNonSpace()
	pipe = t.newPipeline(token.pos, token.line, nil)
	// Are there declarations or assignments?
decls:
	if v := t.peekNonSpace(); v.typ == itemVariable {
		t.next()
		// Since space is a token, we need 3-token look-ahead here in the worst case:
		// in "$x foo" we need to read "foo" (as opposed to ":=") to know that $x is an
		// argument variable rather than a declaration. So remember the token
		// adjacent to the variable so we can push it back if necessary.
		tokenAfterVariable := t.peek()
		next := t.peekNonSpace()
		switch {
		case next.typ == itemAssign, next.typ == itemDeclare:
			pipe.IsAssign = next.typ == itemAssign
			t.nextNonSpace()
			pipe.Decl = append(pipe.Decl, t.newVariable(v.pos, v.val))
			t.vars = append(t.vars, v.val)
		case next.typ == itemChar && next.val == ",":
			t.nextNonSpace()
			pipe.Decl = append(pipe.Decl, t.newVariable(v.pos, v.val))
			t.vars = append(t.vars, v.val)
			if context == "range" && len(pipe.Decl) < 2 {
				switch t.peekNonSpace().typ {
				case itemVariable, itemRightDelim, itemRightParen:
					// second initialized variable in a range pipeline
					goto decls
				default:
					t.errorf("range can only initialize variables")
				}
			}
			t.errorf("too many declarations in %s", context)
		case tokenAfterVariable.typ == itemSpace:
			t.backup3(v, tokenAfterVariable)
		default:
			t.backup2(v)
		}
	}
	for {
		switch token := t.nextNonSpace(); token.typ {
		case itemRightDelim, itemRightParen:
			// At this point, the pipeline is complete
			t.checkPipeline(pipe, context)
			if token.typ == itemRightParen {
				t.backup()
			}
			return
		case itemBool, itemImportedFieldFunc, itemDOMIdentifier, itemCharConstant, itemComplex, itemDot, itemField, itemIdentifier,
			itemNumber, itemNil, itemRawString, itemString, itemVariable, itemLeftParen:
			t.backup()
			pipe.append(t.command())
		case itemEOF:
			t.backup()
			return
		default:
			t.unexpected(token, context)
		}
	}
}

func (t *Tree) checkPipeline(pipe *PipeNode, context string) {
	// Reject empty pipelines
	if len(pipe.Cmds) == 0 {
		t.errorf("missing value for %s", context)
	}
	// Only the first command of a pipeline can start with a non executable operand
	for i, c := range pipe.Cmds[1:] {
		switch c.Args[0].Type() {
		case NodeBool, NodeDot, NodeNil, NodeNumber, NodeString:
			// With A|B|C, pipeline stage 2 is B
			t.errorf("non executable command in pipeline stage %d", i+2)
		}
	}
}

func (t *Tree) parseControl(allowElseIf bool, context string) (pos Pos, line int, pipe *PipeNode, list, elseList *ListNode) {
	defer t.popVars(len(t.vars))
	pipe = t.pipeline(context)
	var next Node
	list, next = t.itemList()
	switch next.Type() {
	case nodeEnd: // done
	case nodeElse:
		if allowElseIf {
			// Special case for "else if". If the "else" is followed immediately by an "if",
			// the elseControl will have left the "if" token pending. Treat
			//	{{if a}}_{{else if b}}_{{end}}
			// as
			//	{{if a}}_{{else}}{{if b}}_{{end}}{{end}}.
			// To do this, parse the if as usual and stop at it {{end}}; the subsequent{{end}}
			// is assumed. This technique works even for long if-else-if chains.
			// TODO: Should we allow else-if in with and range?
			if t.peek().typ == itemIf {
				t.next() // Consume the "if" token.
				elseList = t.newList(next.Position())
				elseList.append(t.ifControl())
				// Do not consume the next item - only one {{end}} required.
				break
			}
		}

		elseList, next = t.itemList()
		if next.Type() != nodeEnd {
			t.errorf("expected end; found %s", next)
		}
	}
	return pipe.Position(), pipe.Line, pipe, list, elseList
}

// If:
//	{{if pipeline}} itemList {{end}}
//	{{if pipeline}} itemList {{else}} itemList {{end}}
// If keyword is past.
func (t *Tree) ifControl() Node {
	return t.newIf(t.parseControl(true, "if"))
}

// Range:
//	{{range pipeline}} itemList {{end}}
//	{{range pipeline}} itemList {{else}} itemList {{end}}
// Range keyword is past.
func (t *Tree) rangeControl() Node {
	return t.newRange(t.parseControl(false, "range"))
}

// With:
//	{{with pipeline}} itemList {{end}}
//	{{with pipeline}} itemList {{else}} itemList {{end}}
// If keyword is past.
func (t *Tree) withControl() Node {
	return t.newWith(t.parseControl(false, "with"))
}

// EndMethod:
//	{{endMethod}}
// End keyword is past.
func (t *Tree) endMethodControl() Node {
	return t.newMethodEnd(t.expect(itemRightDelim, "end").pos)
}

// End:
//	{{end}}
// End keyword is past.
func (t *Tree) endControl() Node {
	return t.newEnd(t.expect(itemRightDelim, "end").pos)
}

// Else:
//	{{else}}
// Else keyword is past.
func (t *Tree) elseControl() Node {
	// Special case for "else if".
	peek := t.peekNonSpace()
	if peek.typ == itemIf {
		// We see "{{else if ... " but in effect rewrite it to {{else}}{{if ... ".
		return t.newElse(peek.pos, peek.line)
	}
	token := t.expect(itemRightDelim, "else")
	return t.newElse(token.pos, token.line)
}

// Block:
//	{{block stringValue pipeline}}
// Block keyword is past.
// The name must be something that can evaluate to a string.
// The pipeline is mandatory.
func (t *Tree) blockControl() Node {
	const context = "block clause"

	token := t.nextNonSpace()
	name := t.parseTemplateName(token, context)
	pipe := t.pipeline(context)

	block := New(name) // name will be updated once we know it.
	block.text = t.text
	block.ParseName = t.ParseName
	block.startParse(t.FuncMaps, t.lex, t.treeSet)
	var end Node
	block.Root, end = block.itemList()
	if end.Type() != nodeEnd {
		t.errorf("unexpected %s in %s", end, context)
	}
	block.add()
	block.stopParse()

	return t.newTemplate(token.pos, token.line, name, pipe)
}

// Template:
//	{{template stringValue pipeline}}
// Template keyword is past. The name must be something that can evaluate
// to a string.
func (t *Tree) templateControl() Node {
	const context = "template clause"
	token := t.nextNonSpace()
	name := t.parseTemplateName(token, context)
	var pipe *PipeNode
	if t.nextNonSpace().typ != itemRightDelim {
		t.backup()
		// Do not pop variables; they persist until "end".
		pipe = t.pipeline(context)
	}
	return t.newTemplate(token.pos, token.line, name, pipe)
}

func (t *Tree) parseTemplateName(token item, context string) (name string) {
	switch token.typ {
	case itemString, itemRawString:
		s, err := strconv.Unquote(token.val)
		if err != nil {
			t.error(err)
		}
		name = s
	default:
		t.unexpected(token, context)
	}
	return
}

// command:
//	operand (space operand)*
// space-separated arguments up to a pipeline character or right delimiter.
// we consume the pipe character but leave the right delim to terminate the action.
func (t *Tree) command() *CommandNode {
	cmd := t.newCommand(t.peekNonSpace().pos)
	cmd.Line = t.peekNonSpace().line

handleLoop:
	for {
		t.peekNonSpace() // skip leading spaces.
		operand := t.operand()
		if operand != nil {
			cmd.append(operand)
		}
		token := t.next()
		switch token.typ {
		case itemSpace:
			continue handleLoop
		case itemError:
			t.errorf("%s", token.val)
		case itemRightDelim, itemRightParen:
			t.backup()
		case itemPipe:
		case itemEOF:
			t.backup()
		case itemImportedFieldFunc:
		default:
			t.errorf("unexpected %s in operand", token)
		}
		break
	}

	if len(cmd.Args) == 0 {
		t.errorf("empty command")
	}
	return cmd
}

// operand:
//	term .Field*
// An operand is a space-separated component of a command,
// a term possibly followed by field accesses.
// A nil return means the next item is not an operand.
func (t *Tree) operand() Node {
	node := t.term()
	if node == nil {
		return nil
	}
	if t.peek().typ == itemField {
		chain := t.newChain(t.peek().pos, node)
		for t.peek().typ == itemField {
			chain.Add(t.next().val)
		}
		// Compatibility with original API: If the term is of type NodeField
		// or NodeVariable, just put more fields on the original.
		// Otherwise, keep the Chain node.
		// Obvious parsing errors involving literal values are detected here.
		// More complex error cases will have to be handled at execution time.
		switch node.Type() {
		case NodeField:
			node = t.newField(chain.Position(), chain.String())
		case NodeVariable:
			node = t.newVariable(chain.Position(), chain.String())
		case NodeBool, NodeString, NodeNumber, NodeNil, NodeDot:
			t.errorf("unexpected . after term %q", node.String())
		default:
			node = chain
		}
	}
	return node
}

// term:
//	literal (number, string, nil, boolean)
//	function (identifier)
//	.
//	.Field
//	$
//	'(' pipeline ')'
// A term is a simple "expression".
// A nil return means the next item is not a term.
func (t *Tree) term() Node {
	token := t.nextNonSpace()
	switch token.typ {
	case itemError:
		t.errorf("%s", token.val)
	case itemDOMIdentifier:
		if !IsDomFuncs(token.val) {
			t.errorf("unable to handle dom identifier %q", token.val)
		}
		return t.newDOMIdentifer(token.pos, token.val)
	case itemIdentifier:
		// are we handling a method parsing ?
		if t.currentMethod != nil {
			if t.currentMethod.HasArgument(token.val) {
				return NewMethodIdentifier(token.val).SetTree(t).SetPos(token.pos)
			}
		}

		if !t.hasFunction(token.val) {
			t.errorf("function %q not defined", token.val)
		}

		var funcName = t.getFunctionName(token.val)
		return NewIdentifier(funcName).SetTree(t).SetPos(token.pos).setIsInbuiltFunction(true)
	case itemImportedFieldFunc:
		return t.importedFieldFunc(token)
	case itemDot:
		return t.newDot(token.pos)
	case itemNil:
		return t.newNil(token.pos)
	case itemVariable:
		return t.useVar(token.pos, token.val)
	case itemField:
		return t.newField(token.pos, token.val)
	case itemBool:
		return t.newBool(token.pos, token.val == "true")
	case itemCharConstant, itemComplex, itemNumber:
		number, err := t.newNumber(token.pos, token.val, token.typ)
		if err != nil {
			t.error(err)
		}
		return number
	case itemLeftParen:
		pipe := t.pipeline("parenthesized pipeline")
		if token := t.next(); token.typ != itemRightParen {
			t.errorf("unclosed right paren: unexpected %s", token)
		}
		return pipe
	case itemString, itemRawString:
		s, err := strconv.Unquote(token.val)
		if err != nil {
			t.error(err)
		}
		return t.newString(token.pos, token.val, s)
	}
	t.backup()
	return nil
}

// hasFunction reports if a function name exists in the Tree's maps.
func (t *Tree) hasFunction(name string) bool {
	var _, hasFunc = t.FuncMaps.Funcs[name]
	if !hasFunc {
		return domFunctions[name]
	}
	return hasFunc
}

// getFunctionName returns the intended name of function.
func (t *Tree) getFunctionName(name string) string {
	if fName, hasFName := t.FuncMaps.Funcs[name]; hasFName {
		return fName
	}
	return name
}

// popVars trims the variable list to the specified length
func (t *Tree) popVars(n int) {
	t.vars = t.vars[:n]
}

// useVar returns a node for a variable reference. It errors if the
// variable is not defined.
func (t *Tree) useVar(pos Pos, name string) Node {
	v := t.newVariable(pos, name)
	for _, varName := range t.vars {
		if varName == v.Ident[0] {
			return v
		}
	}
	t.errorf("undefined variable %q", v.Ident[0])
	return nil
}

// MergeMaps merges the list of maps into one.
func MergeMaps(items ...map[string]string) map[string]string {
	var elem = map[string]string{}
	for _, item := range items {
		for k, v := range item {
			elem[k] = v
		}
	}
	return elem
}

func reverseCommands(cmds []*CommandNode) []*CommandNode {
	var reversed = make([]*CommandNode, len(cmds))
	var cmdSize = len(cmds)
	var next = 0
	for i := cmdSize; i > 0; i-- {
		var item = cmds[i-1]
		reversed[next] = item
		next++
	}
	return reversed
}
