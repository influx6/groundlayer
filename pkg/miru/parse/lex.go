// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// item represents a token or text string returned from the scanner.
type item struct {
	typ  itemType // The type of this item.
	pos  Pos      // The starting position, in bytes, of this item in the input string.
	val  string   // The value of this item.
	line int      // The line number at the start of this item.
}

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	case i.typ > itemKeyword:
		return fmt.Sprintf("<%s>", i.val)
	case len(i.val) > 10:
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

// itemType identifies the type of lex items.
type itemType int

const (
	itemError                itemType = iota // error occurred; value is text of error
	itemIgnore                               // ignore keyword
	itemBool                                 // boolean constant
	itemChar                                 // printable ASCII character; grab bag for comma etc.
	itemCharConstant                         // character constant
	itemComplex                              // complex constant (1+2i); imaginary is just a number
	itemAssign                               // equals ('=') introducing an assignment
	itemDeclare                              // colon-equals (':=') introducing a declaration
	itemEOF                                  // eof
	itemField                                // alphanumeric identifier starting with '.'
	itemIdentifier                           // alphanumeric identifier not starting with '.'
	itemLeftDelim                            // left action delimiter
	itemLeftParen                            // '(' inside action
	itemNumber                               // simple number, including imaginary
	itemPipe                                 // pipe symbol
	itemRawString                            // raw quoted string (includes quotes)
	itemRightDelim                           // right action delimiter
	itemRightParen                           // ')' inside action
	itemSpace                                // run of spaces separating arguments
	itemString                               // quoted string (includes quotes)
	itemText                                 // plain text
	itemVariable                             // variable starting with '$', such as '$' or  '$1' or '$hello'
	itemKeyword                              // used only to delimit the keywords
	itemBlock                                // block keyword
	itemTagAttrValueStart                    // block keyword
	itemTagThemeAttrStart                    // block keyword
	itemTagAttrStart                         // block keyword
	itemTagAttrEnd                           // block keyword
	itemTag                                  // block keyword
	itemTagEnd                               // block keyword
	itemTagSelfEnd                           // block keyword
	itemDot                                  // the cursor, spelled '.'
	itemDefine                               // define keyword
	itemElse                                 // else keyword
	itemEnd                                  // end keyword
	itemIf                                   // if keyword
	itemNil                                  // the untyped nil constant, easiest to treat as a keyword
	itemRange                                // range keyword
	itemTemplate                             // template keyword
	itemWith                                 // with keyword
	itemMethod                               // method keyword
	itemMethodEnd                            // end of method section
	itemModel                                // model keyword
	itemModelField                           // model keyword
	itemModelEnd                             // end of model section
	itemModelType                            // modelType keyword
	itemModelTypeBase                        // model keyword
	itemModelTypeEnd                         // end of modelType keyword
	itemRoot                                 // rootType keyword
	itemList                                 // list keyword using [
	itemMount                                // mount keyword
	itemMountList                            // mount keyword
	itemMountLive                            // mountLive keyword
	itemMountLiveList                        // mountLiveList keyword
	itemImportedFieldFunc                    // package identifier package.identifier
	itemImportedFunc                         // package identifier package.identifier
	itemImportedField                        // package identifier package.identifier
	itemImportDeclaration                    // package import declaration
	itemImportDeclarationEnd                 // package import declaration
	itemDOMIdentifier                        // dom keyword
	itemKomponentName                        // komponent name keyword
)

var key = map[string]itemType{
	".":             itemDot,
	"block":         itemBlock,
	"define":        itemDefine,
	"else":          itemElse,
	"end":           itemEnd,
	"if":            itemIf,
	"range":         itemRange,
	"nil":           itemNil,
	"template":      itemTemplate,
	"with":          itemWith,
	"dom":           itemDOMIdentifier,
	"parentdom":     itemDOMIdentifier,
	"rootdom":       itemDOMIdentifier,
	"model":         itemModel,
	"method":        itemMethod,
	"modelType":     itemModelType,
	"endMethod":     itemMethodEnd,
	"rootType":      itemRoot,
	"mount":         itemMount,
	"mountLive":     itemMountLive,
	"mountList":     itemMountList,
	"mountLiveList": itemMountLiveList,
	"import":        itemImportDeclaration,
}

const eof = -1

// Trimming spaces.
// If the action begins "{{- " rather than "{{", then all space/tab/newlines
// preceding the action are trimmed; conversely if it ends " -}}" the
// leading spaces are trimmed. This is done entirely in the lexer; the
// parser never sees it happen. We require an ASCII space to be
// present to agroundlayer ambiguity with things like "{{-3}}". It reads
// better with the space present anyway. For simplicity, only ASCII
// space does the job.
const (
	spaceChars      = " \t\r\n" // These are the space characters defined by Go itself.
	leftTrimMarker  = "- "      // Attached to left delimiter, trims trailing spaces from preceding text.
	rightTrimMarker = " -"      // Attached to right delimiter, trims leading spaces from following text.
	trimMarkerLen   = Pos(len(leftTrimMarker))
)

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	name           string    // the name of the input; used only for error reports
	input          string    // the string being scanned
	leftDelim      string    // start of action
	rightDelim     string    // end of action
	trimRightDelim string    // end of action with trim marker
	pos            Pos       // current position in the input
	start          Pos       // start position of this item
	width          Pos       // width of last rune read from input
	items          chan item // channel of scanned items
	parenDepth     int       // nesting depth of ( ) exprs
	line           int       // 1+number of newlines seen
	startLine      int       // start line of this item
	htmlMode       bool      // htmlMode flag <
	htmlTagMode    bool      // htmlMode flag <
	htmlListMode   bool
	htmlStacks     []string
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	if r == '\n' {
		l.line++
	}
	return r
}

// pop of dom queue
func (l *lexer) popDOM() string {
	total := len(l.htmlStacks)
	last := l.htmlStacks[total-1]
	l.htmlStacks = l.htmlStacks[0 : total-1]
	return last
}

// push into dom queue
func (l *lexer) isLastDOM(dom string) bool {
	total := len(l.htmlStacks)
	last := l.htmlStacks[total-1]
	return last == dom
}

// push into dom queue
func (l *lexer) pushDOM(dom string) {
	l.htmlStacks = append(l.htmlStacks, dom)
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
	// Correct newline count.
	if l.width == 1 && l.input[l.pos] == '\n' {
		l.line--
	}
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.input[l.start:l.pos], l.startLine}
	l.start = l.pos
	l.startLine = l.line
}

func (l *lexer) skipStart() {
	l.start += l.width
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.line += strings.Count(l.input[l.start:l.pos], "\n")
	l.start = l.pos
	l.startLine = l.line
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...), l.startLine}
	return nil
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) nextItem() item {
	return <-l.items
}

// drain drains the output so the lexing goroutine will exit.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) drain() {
	if len(l.items) == 0 {
		return
	}
	for range l.items {
	}
}

// lex creates a new scanner for the input string.
func lex(name, input, left, right string) *lexer {
	if left == "" {
		left = leftDelim
	}
	if right == "" {
		right = rightDelim
	}
	l := &lexer{
		name:           name,
		input:          input,
		leftDelim:      left,
		rightDelim:     right,
		trimRightDelim: rightTrimMarker + right,
		items:          make(chan item, 1),
		line:           1,
		startLine:      1,
		htmlMode:       false,
	}
	go l.run()
	return l
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

// state functions

const (
	leftDelim              = "{{"
	rightDelim             = "}}"
	leftHTMLDelim          = "<"
	rightHTMLDelim         = ">"
	leftCloseHTMLDelim     = `</`
	rightCloseHTMLDelim    = `/>`
	leftComment            = "/*"
	rightComment           = "*/"
	singleQuote            = '\''
	doubleQuote            = '"'
	leftListDelim          = '['
	rightListDelim         = ']'
	leftAngleBracketDelim  = '<'
	rightAngleBracketDelim = '>'
)

// lexText scans until an opening action delimiter, "{{".
func lexText(l *lexer) stateFn {
	var totalItems = len(l.input)
	var firstDelim, _ = utf8.DecodeRuneInString(l.leftDelim)
	for {
		if l.pos == Pos(totalItems) {
			l.emit(itemEOF)
			return nil
		}

		var currentLexed = len(l.input[l.start:l.pos])
		if r := l.peek(); isSpace(r) && currentLexed > 1 {
			l.emit(itemText)
		}

		// lex out all the space
		if lexSpaceUntil(l) {
			l.emit(itemSpace)
		}

		if nr := l.next(); nr == firstDelim || nr == leftAngleBracketDelim {
			l.backup()

			l.emit(itemText)
			if nr == leftAngleBracketDelim {
				l.htmlMode = true
				return lexTag
			}

			var foundDelim, foundCount = lexIfDelimiter(l, l.leftDelim)
			lexBackupCount(l, foundCount)
			if foundDelim {
				return lexLeftDelim
			}

			l.next()
		}
	}
}

func lexMethod(l *lexer) stateFn {
	// Collect away possible space
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	// CollectName for model
	lexTextUntil(l)
	l.emit(itemMethod)

	// Collect away possible space
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	if r := l.next(); r != '|' {
		return l.errorf("A model must have a | rune after it's name to signify field declarations")
	}

	l.ignore()
	lexModelFields(l)
	l.ignore()
	return lexText
}

func lexModel(l *lexer) stateFn {
	// Collect away possible space
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	// CollectName for model
	lexTextUntil(l)
	l.emit(itemModel)

	// Collect away possible space
	if lexSpaceUntil(l) {
		l.ignore()
	}

	if r := l.next(); r != '|' {
		return l.errorf("A model must have a | rune after it's name to signify field declarations")
	}

	l.ignore()
	lexModelFields(l)
	l.emit(itemModelEnd)

	return lexText
}

func lexImportDeclaration(l *lexer) stateFn {
	// Collect away possible space
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	// CollectName for import alias
	if count := lexAllUntil(l, func(r rune) bool {
		return r == '_' || isAlphaNumeric(r)
	}); count == 0 {
		return l.errorf("A import requires an alias as next token")
	}

	l.emit(itemImportDeclaration)

	// Collect away possible space
	if lexSpaceUntil(l) {
		l.ignore()
	}

	if r := l.next(); r != '|' {
		return l.errorf("A model must have a | rune after it's name to signify field declarations")
	}

	l.ignore()

	// Collect away possible space
	if lexSpaceUntil(l) {
		l.ignore()
	}

	if count := lexAllUntil(l, func(r rune) bool {
		return r == '_' || r == '-' || r == '/' || r == '.' || isAlphaNumeric(r)
	}); count == 0 {
		return l.errorf("A import requires an the path to import")
	}
	l.emit(itemImportDeclarationEnd)

	// Collect away possible space
	if lexSpaceUntil(l) {
		l.ignore()
	}

	// if we found the ending right delimiter, then this
	// should be considered a field and not a function call
	var foundDelim, _ = lexIfDelimiter(l, l.rightDelim)
	if !foundDelim {
		return l.errorf("Unexpected token found at line %d: %q", l.line, l.input[l.start:])
	}

	l.ignore()
	return lexText
}

// lexImportFieldStart scans a import function: #Alphanumeric.Function
// The # has been scanned.
func lexImportFieldStart(l *lexer) stateFn {
	if l.atTerminator() { // Nothing interesting follows -> "$".
		l.emit(itemIgnore)
		return lexInsideAction
	}
	return lexImportField(l, itemImportedField)
}

// lexImportFuncStart scans a import function: #Alphanumeric.Function
// The # has been scanned.
func lexImportFuncStart(l *lexer) stateFn {
	if l.atTerminator() { // Nothing interesting follows -> "$".
		l.emit(itemIgnore)
		return lexInsideAction
	}
	return lexImportField(l, itemImportedFunc)
}

func lexImportField(l *lexer, t itemType) stateFn {
	if count := lexAllUntil(l, func(r rune) bool {
		return isAlphaNumeric(r) || r == '.'
	}); count == 0 {
		return l.errorf("expecting alphaneumeric set of token after #")
	}
	l.emit(itemImportedFieldFunc)

	if !l.atTerminator() {
		return l.errorf("bad character %#U", l.peek())
	}

	// Collect away possible space
	if lexSpaceUntil(l) {
		l.ignore()
	}

	l.emit(t)

	return lexInsideAction
}

func lexModelFields(l *lexer) {
doLoop:
	for {
		var foundDelim, _ = lexIfDelimiter(l, l.rightDelim)
		if foundDelim {
			break doLoop
		}

		// Collect away possible space
		if lexSpaceUntil(l) {
			l.emit(itemSpace)
		}

		lexAllUntil(l, func(r rune) bool {
			// Allow space and allow colon(:), so that we grab the fieldName and Type
			// in format: fieldName:fieldType
			return isSpace(r) || r == ':' || isAlphaNumeric(r)
		})

		l.emit(itemModelField)

		// Collect away possible space
		if lexSpaceUntil(l) {
			l.emit(itemSpace)
		}

		if commaRune := l.next(); commaRune == ',' {
			l.ignore()
			continue
		}

		l.backup()
		continue
	}
}

func lexMount(l *lexer) stateFn {
	return lexMountBase(l, false)
}

func lexMountList(l *lexer) stateFn {
	return lexMountBase(l, true)
}

func lexMountBase(l *lexer, isList bool) stateFn {
	// Collect away possible space
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	// collect route for mount
	if count := lexAllUntil(l, func(r rune) bool {
		return r == '/' || r == '_' || r == '-' || isAlphaNumeric(r)
	}); count == 0 {
		return l.errorf("mount path is required for this ")
	}

	if isList {
		l.emit(itemMountList)
	} else {
		l.emit(itemMount)
	}

	// Collect away possible space
	if lexSpaceUntil(l) {
		l.ignore()
	}

	var foundDelim, _ = lexIfDelimiter(l, l.rightDelim)
	if foundDelim {
		l.emit(itemRightDelim)
		return lexText
	}

	if lexSpaceUntil(l) {
		l.ignore()
	}

	return lexInsideAction
}

func lexRootType(l *lexer) stateFn {
	// Collect away possible space
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	// CollectName for root model
	lexTextAndDotUntil(l)
	l.emit(itemRoot)

	// Collect away possible space
	if lexSpaceUntil(l) {
		l.ignore()
	}

	var foundDelim, _ = lexIfDelimiter(l, l.rightDelim)
	if !foundDelim {
		return l.errorf("failed to close rootType declaration in line %d at %d", l.line, l.pos)
	}

	l.ignore()
	return lexText
}

func lexModelType(l *lexer) stateFn {
	// Collect away possible space
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	// CollectName for model
	lexTextUntil(l)
	l.emit(itemModelType)

	// Collect away possible space
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	if r := l.next(); r != '|' {
		return l.errorf("A model must have a | rune after it's name to signify field declarations")
	}

	l.ignore()
	lexModelTypeBase(l)
	l.emit(itemModelTypeEnd)

	return lexText
}

func lexModelTypeBase(l *lexer) {
	// Collect away possible space
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	lexAllUntil(l, func(r rune) bool {
		// Allow names like zip_code, allow space and allow colon(:)
		return r == '(' || r == ')' || r == ',' || isSpace(r) || isAlphaNumeric(r)
	})

	l.emit(itemModelTypeBase)

	// Collect away possible space
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	var foundDelim, _ = lexIfDelimiter(l, l.rightDelim)
	if foundDelim {
		return
	}

	l.errorf("Invalid ending of ModelType field expected %q ending", l.rightDelim)
}

func lexTag(l *lexer) stateFn {
	var pr = l.peek()
	if pr != leftAngleBracketDelim {
		return lexText
	}

	var nr2 = l.input[l.start : l.pos+2]
	if nr2 == "<!" {
		return lexDocTypeTag
	}

	if nr2 == "/>" {
		l.next()
		l.next()
		l.emit(itemTagSelfEnd)
		l.popDOM()
		return lexText
	}

	if nr2 == "</" {
		l.next()
		l.next()
		l.emit(itemIgnore)

		if !lexTextUntil(l) {
			return l.errorf("invalid html tag ending")
		}

		var domNode = l.input[l.start:l.pos]
		if l.isLastDOM(domNode) {
			l.popDOM()
			l.emit(itemTagEnd)

			if mr := l.next(); mr != rightAngleBracketDelim {
				return l.errorf("invalid html tag ending, expected '>', see line %d at %d", l.line, l.pos)
			}
			l.ignore()
			return lexText
		}
	}

	l.next()
	l.skipStart()
	lexTagName(l)
	l.pushDOM(l.input[l.start:l.pos])
	l.emit(itemTag)

	return lexSpaceWith(lexTagAttrs)
}

func lexDocTypeTag(l *lexer) stateFn {
	lexNextCount(l, 2)
	l.ignore()

	// gather all the space available.
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	// gather all the space available.
	lexTagName(l)
	l.emit(itemTag)

	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	lexAllUntil(l, func(r rune) bool {
		return !isOtherTypeOfSpaceOrGreaterThan(r)
	})
	l.emit(itemText)

	if l.next() == rightAngleBracketDelim {
		l.emit(itemTagEnd)
		return lexText
	}

	l.backup()
	return l.errorf("expected token(>) at the end of a doctype tag")
}

func lexTagAttrs(l *lexer) stateFn {
	// gather all the space available.
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	lexNextCount(l, 2)
	var nr2 = l.input[l.start:l.pos]
	if nr2 == "/>" {
		l.emit(itemTagAttrEnd)
		l.htmlListMode = false
		l.emit(itemTagSelfEnd)
		l.popDOM()
		return lexText
	}
	lexBackupCount(l, 2)

	// lexTextUntil(l)
	if found := lexAllUntil(l, func(r rune) bool {
		return isAlphaNumeric(r) || (r == '_') || (r == '.') || (r == '-') || r == '(' || r == ')'
	}); found > 0 {
		if l.htmlListMode {
			l.emit(itemText)
		} else {
			l.emit(itemTagAttrStart)
		}
	}

	// gather all the space available.
	if lexSpaceUntil(l) {
		l.emit(itemSpace)
	}

	var nr = l.next()
	if nr == '\\' {
		l.next()
		var ending = l.input[l.start:l.pos]
		if ending == "\\>" {
			l.emit(itemTagAttrEnd)
			l.htmlListMode = false
			return lexText
		}
		l.backup()
	}

	if nr == '{' {
		l.backup()

		var foundDelim, foundCount = lexIfDelimiter(l, l.leftDelim)
		lexBackupCount(l, foundCount)
		if foundDelim {
			l.emit(itemText)
			var nextFunc = lexLeftDelim
			for {
				nextFunc = nextFunc(l)
				if nextFunc == nil {
					break
				}
			}
		}
	}

	if nr == rightListDelim && l.htmlListMode {
		l.ignore()
		l.emit(itemTagAttrEnd)
		l.htmlListMode = false
		return lexTagAttrs
	}

	if nr != '=' && nr != rightAngleBracketDelim {
		// might be another tag in-front
		return lexTagAttrs
	}

	if nr == rightAngleBracketDelim {
		l.emit(itemTagAttrEnd)
		l.htmlListMode = false
		return lexText
	}

	if nr == rightListDelim && l.htmlListMode {
		l.ignore()
		l.emit(itemTagAttrEnd)
		l.htmlListMode = false
		return lexTagAttrs
	}

	if nr == '=' {
		if lexSpaceUntil(l) {
			l.emit(itemSpace)
		}

		l.emit(itemTagAttrValueStart)

		var nextNr = l.peek()

		if nextNr == leftListDelim {
			return lexTagAttrValueWithList
		}

		if nextNr == doubleQuote {
			return lexTagAttrValueWithDoubleQuote
		}

		if nextNr == singleQuote {
			return lexTagAttrValueWithSingleQuote
		}

		return lexTagAttrValueWithoutQuote
	}

	l.emit(itemTagAttrEnd)
	l.htmlListMode = false
	return lexText
}

func lexTagAttrValueWithList(l *lexer) stateFn {
	l.htmlListMode = true
	var foundEndingDelimiter = lexTextEnclosedInStarterAndEndsWith(l, leftListDelim, rightListDelim, l.htmlTagMode)
	l.emit(itemList)

	if foundEndingDelimiter {
		l.htmlTagMode = false
		l.emit(itemTagAttrEnd)
		return lexTagAttrs
	}

	r := l.peek()
	if r == rightAngleBracketDelim {
		if !foundEndingDelimiter {
			return l.errorf("expected ending quote(\") but found rightAngleBracketDelim")
		}

		l.next()
		l.emit(itemTagAttrEnd)
		return lexText
	}

	lexNextCount(l, 2)
	if l.input[l.start:l.pos] == "/>" {
		if !foundEndingDelimiter {
			return l.errorf("expected ending quote(\") but found rightAngleBracketDelim")
		}

		lexBackupCount(l, 2)
		l.emit(itemTagAttrEnd)
		lexNextCount(l, 2)

		l.emit(itemTagSelfEnd)

		l.popDOM()
		return lexText
	}

	lexBackupCount(l, 2)

	var foundDelim, foundCount = lexIfDelimiter(l, l.leftDelim)
	lexBackupCount(l, foundCount)
	if foundDelim {
		l.emit(itemText)
		l.htmlTagMode = true
		var nextFunc = lexLeftDelim
		for {
			nextFunc = nextFunc(l)
			if nextFunc == nil {
				break
			}
		}
	}

	return lexTagAttrs
}

func lexTagAttrValueWithDoubleQuote(l *lexer) stateFn {
	return lexTagAttrValueWithQuote(l, '"')
}

func lexTagAttrValueWithSingleQuote(l *lexer) stateFn {
	return lexTagAttrValueWithQuote(l, '\'')
}

func lexTagAttrValueWithQuote(l *lexer, quoteType rune) stateFn {
	var foundEndingDelimiter = lexTextEnclosedInStarterAndEndsWithStarter(l, quoteType, l.htmlTagMode)
	l.emit(itemText)

	if foundEndingDelimiter {
		l.htmlTagMode = false
		l.emit(itemTagAttrEnd)
		return lexTagAttrs
	}

	r := l.peek()
	if r == rightAngleBracketDelim {
		if !foundEndingDelimiter {
			return l.errorf("expected ending quote(\") but found rightAngleBracketDelim")
		}

		l.next()
		l.emit(itemTagAttrEnd)
		return lexText
	}

	lexNextCount(l, 2)
	if l.input[l.start:l.pos] == "/>" {
		if !foundEndingDelimiter {
			return l.errorf("expected ending quote(\") but found rightAngleBracketDelim")
		}

		lexBackupCount(l, 2)
		l.emit(itemTagAttrEnd)
		lexNextCount(l, 2)

		l.emit(itemTagSelfEnd)

		l.popDOM()
		return lexText
	}

	lexBackupCount(l, 2)

	var foundDelim, foundCount = lexIfDelimiter(l, l.leftDelim)
	lexBackupCount(l, foundCount)
	if foundDelim {
		l.emit(itemText)
		l.htmlTagMode = true
		var nextFunc = lexLeftDelim
		for {
			nextFunc = nextFunc(l)
			if nextFunc == nil {
				break
			}
		}

		if quoteType == '\'' {
			return lexTagAttrValueWithSingleQuote
		}
		return lexTagAttrValueWithDoubleQuote
	}

	return lexTagAttrs
}

func lexTagAttrValueWithoutQuote(l *lexer) stateFn {
	l.next()
	lexTextUntil(l)

	r := l.peek()
	if r == rightAngleBracketDelim {
		l.emit(itemTagAttrEnd)
		l.next()
		return lexText
	}

	l.emit(itemText)
	l.emit(itemTagAttrEnd)
	return lexTagAttrs
}

// rightTrimLength returns the length of the spaces at the end of the string.
func rightTrimLength(s string) Pos {
	return Pos(len(s) - len(strings.TrimRight(s, spaceChars)))
}

// atRightDelim reports whether the lexer is at a right delimiter, possibly preceded by a trim marker.
func (l *lexer) atRightDelim() (delim, trimSpaces bool) {
	if strings.HasPrefix(l.input[l.pos:], l.trimRightDelim) { // With trim marker.
		return true, true
	}
	if strings.HasPrefix(l.input[l.pos:], l.rightDelim) { // Without trim marker.
		return true, false
	}
	return false, false
}

// leftTrimLength returns the length of the spaces at the beginning of the string.
func leftTrimLength(s string) Pos {
	return Pos(len(s) - len(strings.TrimLeft(s, spaceChars)))
}

// lexLeftDelim scans the left delimiter, which is known to be present, possibly with a trim marker.
func lexLeftDelim(l *lexer) stateFn {
	l.pos += Pos(len(l.leftDelim))
	trimSpace := strings.HasPrefix(l.input[l.pos:], leftTrimMarker)
	afterMarker := Pos(0)
	if trimSpace {
		afterMarker = trimMarkerLen
	}
	if strings.HasPrefix(l.input[l.pos+afterMarker:], leftComment) {
		l.pos += afterMarker
		l.ignore()
		return lexComment
	}
	l.emit(itemLeftDelim)
	l.pos += afterMarker
	l.ignore()
	l.parenDepth = 0
	return lexInsideAction
}

// lexComment scans a comment. The left comment marker is known to be present.
func lexComment(l *lexer) stateFn {
	l.pos += Pos(len(leftComment))
	i := strings.Index(l.input[l.pos:], rightComment)
	if i < 0 {
		return l.errorf("unclosed comment")
	}
	l.pos += Pos(i + len(rightComment))
	delim, trimSpace := l.atRightDelim()
	if !delim {
		return l.errorf("comment ends before closing delimiter")
	}
	if trimSpace {
		l.pos += trimMarkerLen
	}
	l.pos += Pos(len(l.rightDelim))
	if trimSpace {
		l.pos += leftTrimLength(l.input[l.pos:])
	}
	l.ignore()
	return lexText
}

// lexRightDelim scans the right delimiter, which is known to be present, possibly with a trim marker.
func lexRightDelim(l *lexer) stateFn {
	trimSpace := strings.HasPrefix(l.input[l.pos:], rightTrimMarker)
	if trimSpace {
		l.pos += trimMarkerLen
		l.ignore()
	}
	l.pos += Pos(len(l.rightDelim))
	l.emit(itemRightDelim)
	if trimSpace {
		l.pos += leftTrimLength(l.input[l.pos:])
		l.ignore()
	}
	if l.htmlTagMode {
		return nil
	}
	return lexText
}

// lexInsideAction scans the elements inside action delimiters.
func lexInsideAction(l *lexer) stateFn {
	// Either number, quoted string, or identifier.
	// Spaces separate arguments; runs of spaces turn into itemSpace.
	// Pipe symbols separate and are emitted.
	delim, _ := l.atRightDelim()
	if delim {
		if l.parenDepth == 0 {
			return lexRightDelim
		}
		return l.errorf("unclosed left paren")
	}
	var r = l.next()
	switch {
	case r == eof || isEndOfLine(r):
		return l.errorf("unclosed action")
	case isSpace(r):
		l.backup() // Put space back in case we have " -}}".
		return lexSpace
	case r == '=':
		l.emit(itemAssign)
	case r == ':':
		if l.next() != '=' {
			return l.errorf("expected :=")
		}
		l.emit(itemDeclare)
	case r == '|':
		l.emit(itemPipe)
	case r == '"':
		return lexQuote
	case r == '`':
		return lexRawQuote
	case r == '$':
		return lexVariable
	case r == '@':
		return lexImportFieldStart
	case r == '#':
		return lexImportFuncStart
	case r == '\'':
		return lexChar
	case r == '.':
		// special look-ahead for ".field" so we don't break l.backup().
		if l.pos < Pos(len(l.input)) {
			r := l.input[l.pos]
			if r < '0' || '9' < r {
				return lexField
			}
		}
		fallthrough // '.' can start a number.
	case r == '+' || r == '-' || ('0' <= r && r <= '9'):
		l.backup()
		return lexNumber
	case isAlphaNumeric(r):
		l.backup()
		return lexIdentifier
	case r == '(':
		l.emit(itemLeftParen)
		l.parenDepth++
	case r == ')':
		l.emit(itemRightParen)
		l.parenDepth--
		if l.parenDepth < 0 {
			return l.errorf("unexpected right paren %#U", r)
		}
	case r <= unicode.MaxASCII && unicode.IsPrint(r):
		l.emit(itemChar)
	default:
		return l.errorf("unrecognized character in action: %#U", r)
	}

	return lexInsideAction
}

// lexIfDelimiter scans until either it exhaust possible search items
// within provided delimiter string or fails to find corresponding
// token in current input. If delimiter exists then the
// count of found will match the length of delimiter.
//
// It also returns true if this is the case.
func lexIfDelimiter(l *lexer, delimiter string) (bool, int) {
	var numFound int
	var r, rs rune
	var i, width int
	for {
		if len(delimiter[i:]) == 0 {
			break
		}
		rs, width = utf8.DecodeRuneInString(delimiter[i:])
		r = l.peek()
		if r != rs {
			break
		}
		i += width
		numFound++
		l.next()
		continue
	}
	return numFound == len(delimiter), numFound
}

// lexNextCount scans a run of alphanumeric characters.
func lexNextCount(l *lexer, count int) {
	var i int
	for i = 0; i < count; i++ {
		l.next()
	}
}

// lexBackupCount scans a run of alphanumeric characters.
func lexBackupCount(l *lexer, count int) {
	var i int
	for i = 0; i < count; i++ {
		l.backup()
	}
}

// lexAllUntil scans a run of alphanumeric characters.
func lexAllUntil(l *lexer, fn func(r rune) bool) int {
	var r rune
	var numSpaces int
	for {
		r = l.peek()
		if !fn(r) {
			break
		}
		l.next()
		numSpaces++
	}
	return numSpaces
}

// lexTextAndDotUntil scans a run of alphaneumeric characters.
func lexTextAndDotUntil(l *lexer) bool {
	var found = false
	var r rune
	var numSpaces int
	for {
		r = l.peek()
		if !isAlphaNumericAndDot(r) {
			break
		}
		found = true
		l.next()
		numSpaces++
	}
	return found
}

// lexTextUntil scans a run of alphaneumeric characters.
func lexTextUntil(l *lexer) bool {
	var found = false
	var r rune
	var numSpaces int
	for {
		r = l.peek()
		if !isAlphaNumeric(r) {
			break
		}
		found = true
		l.next()
		numSpaces++
	}
	return found
}

// lexSpaceUntil scans a run of alphaneumeric characters.
func lexSpaceUntil(l *lexer) bool {
	var foundSpace = false
	var r rune
	for {
		r = l.peek()
		if !isSpace(r) {
			break
		}
		foundSpace = true
		l.next()
	}
	return foundSpace
}

// lexTextWith scans a run of alphanumeric characters.
func lexTextWith(fn stateFn) stateFn {
	return func(l *lexer) stateFn {
		var r rune
		for {
			r = l.peek()
			if !isAlphaNumeric(r) {
				break
			}
			l.next()
		}
		return fn
	}
}

// lexTextEnclosedInStarterAndEndsWith scans a run of alphanumeric characters.
//
// Returns true if we found both the starting and ending quote.
func lexTextEnclosedInStarterAndEndsWith(l *lexer, starter rune, ender rune, ignoreStarter bool) bool {
	var firstDelim, _ = utf8.DecodeRuneInString(l.leftDelim)
	var found = false

	if !ignoreStarter {
		if l.peek() != starter {
			return false
		}

		l.next()
	}

	var r rune
	for {
		r = l.peek()

		// is it an unexpected value or possible params marker ?
		if isGreaterThan(r) {
			break
		}

		// is this some usage of template substitution ?
		if r == firstDelim {
			var foundDelim, foundCount = lexIfDelimiter(l, l.leftDelim)

			// if it is then let this be handled specially else ignore
			if foundDelim {
				lexBackupCount(l, foundCount)
				break
			}

			l.next()
			continue
		}
		// if we see a quote then we have met the
		if r == ender {
			found = true
			l.next()
			break
		}
		l.next()
	}
	return found
}

// lexTextEnclosedInStarterAndEndsWithStarter scans a run of alphanumeric characters.
//
// Returns true if we found both the starting and ending quote.
func lexTextEnclosedInStarterAndEndsWithStarter(l *lexer, starter rune, ignoreStarter bool) bool {
	var firstDelim, _ = utf8.DecodeRuneInString(l.leftDelim)
	var found = false

	if !ignoreStarter {
		if l.peek() != starter {
			return false
		}

		l.next()
	}

	var r rune
	for {
		r = l.peek()
		// is it an unexpected value or possible params marker ?
		if isOtherTypeOfSpaceOrGreaterThan(r) {
			break
		}
		// is this some usage of template substitution ?
		if r == firstDelim {
			var foundDelim, foundCount = lexIfDelimiter(l, l.leftDelim)

			// if it is then let this be handled specially else ignore
			if foundDelim {
				lexBackupCount(l, foundCount)
				break
			}

			l.next()
			continue
		}
		// if we see a quote then we have met the
		if r == starter {
			found = true
			l.next()
			break
		}
		l.next()
	}
	return found
}

func isGreaterThan(r rune) bool {
	return r == rightAngleBracketDelim
}

func isOtherTypeOfSpaceOrGreaterThan(r rune) bool {
	return r == '\t' || r == '\n' || r == '\f' || r == '\r' || r == rightAngleBracketDelim
}

// lexAttrName scans a run of alphaneumeric characters.
func lexAttrName(l *lexer) bool {
	var found = false
	var r rune
	for {
		r = l.peek()
		if isSpaceEqualOrGreaterThan(r) {
			break
		}
		found = true
		l.next()
	}
	return found
}

func isSpaceEqualOrGreaterThan(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\f' || r == '\r' || r == '=' || r == rightAngleBracketDelim
}

// lexTagName scans a run of alphaneumeric characters.
func lexTagName(l *lexer) bool {
	var found = false
	var r rune
	for {
		r = l.peek()
		if isAlphaNumeric(r) {
			found = true
			l.next()
			continue
		}

		// Allow "x-y" or "x:y" but not "x-", "-y", or "x--y".
		if r == ':' || r == '-' {
			l.next()
			if isAlphaNumeric(l.peek()) {
				continue
			}
			l.backup()
		}
		break
	}
	return found
}

// lexSpaceWith scans a run of space characters.
// We have not consumed the first space, which is known to be present.
// Take care if there is a trim-marked right delimiter, which starts with a space.
func lexSpaceWith(fn stateFn) stateFn {
	return func(l *lexer) stateFn {
		var r rune
		var numSpaces int
		for {
			r = l.peek()
			if !isSpace(r) {
				break
			}
			l.next()
			numSpaces++
		}
		// Be careful about a trim-marked closing delimiter, which has a minus
		// after a space. We know there is a space, so check for the '-' that might follow.
		if strings.HasPrefix(l.input[l.pos-1:], l.trimRightDelim) {
			l.backup() // Before the space.
			if numSpaces == 1 {
				return lexRightDelim // On the delim, so go right to that.
			}
		}
		l.emit(itemSpace)
		return fn
	}
}

// lexSpace scans a run of space characters.
// We have not consumed the first space, which is known to be present.
// Take care if there is a trim-marked right delimiter, which starts with a space.
func lexSpace(l *lexer) stateFn {
	var r rune
	var numSpaces int
	for {
		r = l.peek()
		if !isSpace(r) {
			break
		}
		l.next()
		numSpaces++
	}
	// Be careful about a trim-marked closing delimiter, which has a minus
	// after a space. We know there is a space, so check for the '-' that might follow.
	if strings.HasPrefix(l.input[l.pos-1:], l.trimRightDelim) {
		l.backup() // Before the space.
		if numSpaces == 1 {
			return lexRightDelim // On the delim, so go right to that.
		}
	}
	l.emit(itemSpace)
	return lexInsideAction
}

// lexIdentifier scans an alphanumeric.
func lexIdentifier(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// absorb.
		default:
			l.backup()
			word := l.input[l.start:l.pos]
			if !l.atTerminator() {
				return l.errorf("bad character %#U", r)
			}
			switch {
			case key[word] > itemKeyword:
				switch word {
				case "import":
					l.ignore()
					return lexImportDeclaration(l)
				case "modelType":
					l.ignore()
					return lexModelType(l)
				case "model":
					l.ignore()
					return lexModel(l)
				case "mountList":
					l.ignore()
					return lexMountList(l)
				case "mount":
					l.ignore()
					return lexMount(l)
				case "rootType":
					l.ignore()
					return lexRootType(l)
				case "method":
					l.ignore()
					return lexMethod(l)
				default:
					l.emit(key[word])
				}
			case word[0] == '.':
				l.emit(itemField)
			case word == "true", word == "false":
				l.emit(itemBool)
			default:
				l.emit(itemIdentifier)
			}
			break Loop
		}
	}
	return lexInsideAction
}

// lexField scans a field: .Alphanumeric.
// The . has been scanned.
func lexField(l *lexer) stateFn {
	return lexFieldOrVariable(l, itemField)
}

// lexVariable scans a Variable: $Alphanumeric.
// The $ has been scanned.
func lexVariable(l *lexer) stateFn {
	if l.atTerminator() { // Nothing interesting follows -> "$".
		l.emit(itemVariable)
		return lexInsideAction
	}
	return lexFieldOrVariable(l, itemVariable)
}

// lexVariable scans a field or variable: [.$]Alphanumeric.
// The . or $ has been scanned.
func lexFieldOrVariable(l *lexer, typ itemType) stateFn {
	if l.atTerminator() { // Nothing interesting follows -> "." or "$".
		if typ == itemVariable {
			l.emit(itemVariable)
		} else {
			l.emit(itemDot)
		}

		return lexInsideAction
	}
	var r rune
	for {
		r = l.next()
		if !isAlphaNumeric(r) {
			l.backup()
			break
		}
	}

	if !l.atTerminator() {
		return l.errorf("bad character %#U", r)
	}
	l.emit(typ)
	return lexInsideAction
}

// atTerminator reports whether the input is at valid termination character to
// appear after an identifier. Breaks .X.Y into two pieces. Also catches cases
// like "$x+2" not being acceptable without a space, in case we decide one
// day to implement arithmetic.
func (l *lexer) atTerminator() bool {
	r := l.peek()
	if isSpace(r) || isEndOfLine(r) {
		return true
	}
	switch r {
	case eof, '.', ',', '|', ':', ')', '(':
		return true
	}
	// Does r start the delimiter? This can be ambiguous (with delim=="//", $x/2 will
	// succeed but should fail) but only in extremely rare cases caused by willfully
	// bad choice of delimiter.
	if rd, _ := utf8.DecodeRuneInString(l.rightDelim); rd == r {
		return true
	}
	return false
}

// lexChar scans a character constant. The initial quote is already
// scanned. Syntax checking is done by the parser.
func lexChar(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated character constant")
		case '\'':
			break Loop
		}
	}
	l.emit(itemCharConstant)
	return lexInsideAction
}

// lexNumber scans a number: decimal, octal, hex, float, or imaginary. This
// isn't a perfect number scanner - for instance it accepts "." and "0x0.2"
// and "089" - but when it's wrong the input is invalid and the parser (via
// strconv) will notice.
func lexNumber(l *lexer) stateFn {
	if !l.scanNumber() {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}
	if sign := l.peek(); sign == '+' || sign == '-' {
		// Complex: 1+2i. No spaces, must end in 'i'.
		if !l.scanNumber() || l.input[l.pos-1] != 'i' {
			return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
		}
		l.emit(itemComplex)
	} else {
		l.emit(itemNumber)
	}
	return lexInsideAction
}

func (l *lexer) scanNumber() bool {
	// Optional leading sign.
	l.accept("+-")
	// Is it hex?
	digits := "0123456789_"
	if l.accept("0") {
		// Note: Leading 0 does not mean octal in floats.
		if l.accept("xX") {
			digits = "0123456789abcdefABCDEF_"
		} else if l.accept("oO") {
			digits = "01234567_"
		} else if l.accept("bB") {
			digits = "01_"
		}
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if len(digits) == 10+1 && l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789_")
	}
	if len(digits) == 16+6+1 && l.accept("pP") {
		l.accept("+-")
		l.acceptRun("0123456789_")
	}
	// Is it imaginary?
	l.accept("i")
	// Next thing mustn't be alphanumeric.
	if isAlphaNumeric(l.peek()) {
		l.next()
		return false
	}
	return true
}

// lexQuote scans a quoted string.
func lexQuote(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			break Loop
		}
	}
	l.emit(itemString)
	return lexInsideAction
}

// lexRawQuote scans a raw quoted string.
func lexRawQuote(l *lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case eof:
			return l.errorf("unterminated raw quoted string")
		case '`':
			break Loop
		}
	}
	l.emit(itemRawString)
	return lexInsideAction
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumericAndDot(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.'
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func getIndexOfLatestAlphaOrSymbol(s string, c rune) int {
	for index, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == c {
			return index
		}
	}
	return -1
}

func getIndexOfLatestAlpha(s string) int {
	for index, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return index
		}
	}
	return -1
}

func getIndexOfLastAlpha(s string) int {
	for index, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			continue
		}
		return index
	}
	return -1
}

func hasAnyAplhanumeric(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return true
		}
	}
	return false
}

// asciiAlpha reports whether c is an ASCII letter.
func asciiAlpha(c byte) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z'
}

// asciiAlphaNum reports whether c is an ASCII letter or digit.
func asciiAlphaNum(c byte) bool {
	return asciiAlpha(c) || '0' <= c && c <= '9'
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
