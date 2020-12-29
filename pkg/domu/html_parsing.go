package domu

import (
	"strings"
	"sync"

	"bytes"
	"text/template"

	"golang.org/x/net/html"
)

// ParseTemplateInto parses the provided string has a template which
// is processed with the provided binding and passed into the root.
func ParseTemplateInto(root *Node, markup string, binding interface{}) {
	var bu bytes.Buffer

	tmpl := template.Must(template.New("Parsed").Parse(markup))
	if err := tmpl.Execute(&bu, binding); err != nil {
		return
	}

	ParseToRoot(root, bu.String())
}

// ParseTemplate parses the provided string has a template which
// is processed with the provided binding.
func ParseTemplate(markup string, binding interface{}) *Node {
	var bu bytes.Buffer

	tmpl := template.Must(template.New("Parsed").Parse(markup))
	if err := tmpl.Execute(&bu, binding); err != nil {
		return nil
	}

	return ParseTree(bu.String())
}

// ParseToRoot passes the markup generated from the markup added to the provided
// root.
func ParseToRoot(root *Node, markup string) {
	ParseTree(markup).EachChild(func(node *Node, i int) bool {
		root.AppendChild(node)
		return true
	})
}

// ParseAndFirst expects the markup provided to only have one root element which
// will be returned.
func ParseAndFirst(markup string) *Node {
	return ParseTree(markup)
}

type counter interface {
	Next() int
}

type cn struct {
	ml  sync.Mutex
	val int
}

func (c *cn) Next() int {
	c.ml.Lock()
	c.val++
	c.ml.Unlock()
	return c.val
}

// ParseTree takes a html string and returns a *Node representing giving markup.
//
// It returns a Document node always, hence to gain access to the nodes described
// by markup, use Node.EachChild.
func ParseTree(markup string) *Node {
	tokens := html.NewTokenizer(strings.NewReader(markup))

	rootElem := Document()
	pullNode(tokens, rootElem)
	return rootElem
}

func pullNode(tokens *html.Tokenizer, root *Node) {
	for {
		token := tokens.Next()

		switch token {
		case html.ErrorToken:
			return

		case html.TextToken, html.CommentToken, html.DoctypeToken:
			text := strings.TrimSpace(string(tokens.Text()))

			if text == "" {
				continue
			}

			if token == html.CommentToken {
				text = "<!--" + text + "-->"
			}

			root.AppendChild(Text(text))
			continue

		case html.StartTagToken, html.EndTagToken, html.SelfClosingTagToken:
			tagName, hasAttr := tokens.TagName()

			if token == html.EndTagToken && string(tagName) == root.tagName {
				return
			}

			node := NewNode(ElementNode, string(tagName))
			root.AppendChild(node)

			if hasAttr {
			attrLoop:
				for {
					key, val, more := tokens.TagAttr()

					var keyText = string(key)
					if keyText == "id" {
						node.idAttr = string(val)
						continue
					}

					if keyText != "" {
						if strings.HasPrefix(keyText, "event-") {
							var eventName = strings.TrimPrefix(keyText, "event-")
							node.Events.Add(eventName, strings.Split(string(val), "|")...)
							continue
						}

						node.Attrs.Add(NewStringAttr(string(key), string(val)))
					}

					if !more {
						break attrLoop
					}
				}
			}

			if token == html.SelfClosingTagToken {
				continue
			}

			pullNode(tokens, node)
		}
	}
}
