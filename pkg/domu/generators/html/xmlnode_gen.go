// Credit to Richard Musiol (https://github.com/neelance/dom)
// His code was crafted to fit my use

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var workers sync.WaitGroup
var elemNameMap = map[string]string{
	"g":                "Group",
	"font-face-face":   "Fontface",
	"font-face-format": "FontFaceFormat",
	"font-face-name":   "FontfaceName",
	"font-face-src":    "FontFaceSrc",
	"font-face-uri":    "FontfaceURI",
	"missing-glyph":    "MissingGlyph",
	"a":                "Anchor",
	"article":          "Article",
	"aside":            "Aside",
	"area":             "Area",
	"abbr":             "Abbreviation",
	"b":                "Bold",
	"base":             "Base",
	"bdi":              "BidirectionalIsolation",
	"bdo":              "BidirectionalOverride",
	"blockquote":       "BlockQuote",
	"br":               "Break",
	"cite":             "Citation",
	"col":              "Column",
	"colgroup":         "ColumnGroup",
	"datalist":         "DataList",
	"dialog":           "Dialog",
	"details":          "Details",
	"dd":               "Description",
	"del":              "DeletedText",
	"dfn":              "Definition",
	"Def":              "Definition",
	"dl":               "DescriptionList",
	"dt":               "DefinitionTerm",
	"G":                "Group",
	"em":               "Emphasis",
	"embed":            "Embed",
	"footer":           "Footer",
	"figure":           "Figure",
	"figcaption":       "FigureCaption",
	"fieldset":         "FieldSet",
	"h1":               "Header1",
	"h2":               "Header2",
	"h3":               "Header3",
	"h4":               "Header4",
	"h5":               "Header5",
	"h6":               "Header6",
	"hgroup":           "HeadingsGroup",
	"header":           "Header",
	"hr":               "HorizontalRule",
	"i":                "Italic",
	"iframe":           "InlineFrame",
	"img":              "Image",
	"image":            "Image",
	"ins":              "InsertedText",
	"kbd":              "KeyboardInput",
	"keygen":           "KeyGen",
	"li":               "ListItem",
	"meta":             "Meta",
	"menuitem":         "MenuItem",
	"nav":              "Navigation",
	"noframes":         "NoFrames",
	"noscript":         "NoScript",
	"ol":               "OrderedList",
	"option":           "Option",
	"optgroup":         "OptionsGroup",
	"p":                "Paragraph",
	"param":            "Parameter",
	"pre":              "Preformatted",
	"q":                "Quote",
	"rp":               "RubyParenthesis",
	"Ref":              "Reference",
	"rt":               "RubyText",
	"s":                "Strikethrough",
	"samp":             "Sample",
	"source":           "Source",
	"section":          "Section",
	"sub":              "Subscript",
	"sup":              "Superscript",
	"tbody":            "TableBody",
	"textarea":         "TextArea",
	"td":               "TableData",
	"tfoot":            "TableFoot",
	"th":               "TableHeader",
	"thead":            "TableHead",
	"tr":               "TableRow",
	"u":                "Underline",
	"ul":               "UnorderedList",
	"var":              "Variable",
	"track":            "Track",
	"wbr":              "WordBreakOpportunity",
}

// list of self closing tags
var autoclosers = map[string]bool{
	"use":     true,
	"area":    true,
	"base":    true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"input":   true,
	"keygen":  true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
	"br":      true,
}

var code = regexp.MustCompile("</?code>")
var unwanted = regexp.MustCompile("[^\\w\\d-]+")

func pullDoc(url string, fx func(doc *goquery.Document)) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	fx(doc)
	return nil
}

func main() {
	htmlFile, err := os.Create("xml_nodes.gen.go")
	if err != nil {
		panic(err)
	}

	defer htmlFile.Close()

	fmt.Fprint(htmlFile, fileHeader)

	var htmlErr = make(chan error, 2)

	var html, svg bytes.Buffer

	workers.Add(2)
	go buildHTML(&html, htmlErr)
	go buildSVG(&svg, htmlErr)

	workers.Wait()

	select {
	case err := <-htmlErr:
		log.Fatalf(err.Error())
		return
	default:
		if _, err := html.WriteTo(htmlFile); err != nil {
			log.Fatalf(err.Error())
			return
		}
		if _, err := svg.WriteTo(htmlFile); err != nil {
			log.Fatalf(err.Error())
			return
		}
		log.Println("XML Nodes Finished!")
	}
}

func buildHTML(w io.Writer, errs chan error) {
	defer workers.Done()

	var doneHtml = map[string]bool{}
	if err := pullDoc("https://developer.mozilla.org/en-US/docs/Web/HTML/Element", func(doc *goquery.Document) {
		doc.Find("main#content table.standard-table tr td a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if !strings.HasPrefix(link, "/en-US/docs/Web/HTML/Element/") {
				return
			}

			var parent = s.Parent()
			if parent.Find(".icon-trash, .icon-thumbs-down-alt, .icon-warning-sign").Length() > 0 {
				return
			}

			var desc, _ = s.Attr("title")
			text := s.Text()

			var name = text
			if strings.Contains(text, "<") && strings.Contains(text, ">") {
				name = text[1 : len(text)-1]
			}

			var funName = name
			if bName, ok := elemNameMap[strings.ToLower(name)]; ok {
				funName = bName
			}

			funName = restruct(funName)

			if unwanted.MatchString(name) {
				return
			}

			if doneHtml[name] {
				return
			}

			if doneHtml[funName] {
				return
			}

			doneHtml[name] = true
			doneHtml[funName] = true

			writeElem(w, name, funName, desc, link, "Html", "HTML", "XHTML/HTML")
		})
	}); err != nil {
		errs <- err
	}
}

func buildSVG(w io.Writer, errs chan error) {
	defer workers.Done()

	doneSvg := make(map[string]bool)
	if err := pullDoc("https://developer.mozilla.org/en-US/docs/Web/SVG/Element", func(doc *goquery.Document) {
		doc.Find(".index ul li a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")

			if !strings.HasPrefix(link, "/en-US/docs/Web/SVG/Element/") {
				return
			}

			if s.Parent().Find(".icon-trash, .icon-thumbs-down-alt, .icon-warning-sign").Length() > 0 {
				return
			}

			desc, _ := s.Attr("title")

			text := code.ReplaceAllString(s.Text(), "")

			name := text[1 : len(text)-1]
			var funName = name
			if bName, ok := elemNameMap[strings.ToLower(name)]; ok {
				funName = bName
			}

			funName = restruct(funName)

			if doneSvg[name] || doneSvg[funName] || unwanted.MatchString(name) {
				return
			}

			writeElem(w, name, funName, desc, link, "Svg", "SVG", "XML/SVG")
			doneSvg[name] = true
			doneSvg[funName] = true
		})
	}); err != nil {
		errs <- err
	}
}

var badSymbs = regexp.MustCompile("-(.+)")

func writeElem(w io.Writer, name, funName, desc, link string, exception string, prefix string, ty string) {
	for badSymbs.MatchString(funName) {
		if simbs := badSymbs.FindStringSubmatch(funName); len(simbs) > 0 {
			item := capitalize(simbs[1])
			funName = badSymbs.ReplaceAllString(funName, item)
		}
	}

	funName = capitalize(funName)
	if funName != exception {
		funName = prefix + funName
	}

	fmt.Fprintf(w, nodeFormat, funName, name, ty, desc, link, funName, name)
}

// capitalize capitalizes the first character in a string
func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func restruct(s string) string {
	if strings.Contains(s, "-") {
		mo := strings.Split(s, "-")
		for index, mi := range mo {
			if index == 0 {
				continue
			}

			mo[index] = capitalize(mi)
		}

		return strings.Join(mo, "")
	}

	return s
}

const fileHeader = `// Code auto-generated to provide HTML and SVG DOM Nodes.
// Documentation source: "HTML element reference" by Mozilla Contributors.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under CC-BY-SA 2.5.

package domu
`
const nodeFormat = `
// %s provides Node representation for the element %q in %s DOM 
// %s
// https://developer.mozilla.org%s
func %s(renders ...Mounter) *Node {
	return Element("%s", renders...)
}

`
