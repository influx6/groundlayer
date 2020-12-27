package domu

import (
	"fmt"
	"strings"

	"github.com/influx6/npkg"

	"github.com/influx6/npkg/nerror"
)

var _ npkg.EncodableObject = (*JSONNode)(nil)

type JSONAttr struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type JSONEvent struct {
	Name    string   `json:"name"`
	Targets []string `json:"targets"`
}

func NewJSONEvent(name string, targets ...string) JSONEvent {
	return JSONEvent{Name: name, Targets: targets}
}

func (nj JSONEvent) Text() string {
	return fmt.Sprintf("event-%s=%q", nj.Name, strings.Join(nj.Targets, "|"))
}

func (nj JSONEvent) Mount(root *Node) error {
	root.Events.AddJSONEvent(nj)
	return nil
}

type JSONNodes []JSONNode

type JSONNode struct {
	Type        int         `json:"type"`
	Removed     bool        `json:"removed"`
	Changed     bool        `json:"changed"`
	AttrChanged bool        `json:"attr_changed"`
	Id          string      `json:"id"`
	Tid         string      `json:"tid"`
	AtId        string      `json:"atid"`
	Ref         string      `json:"ref"`
	TypeName    string      `json:"type_name"`
	Name        string      `json:"name"`
	Content     string      `json:"content"`
	Namespace   string      `json:"namespace"`
	Attrs       []JSONAttr  `json:"attrs"`
	Events      []JSONEvent `json:"events"`
	Children    []JSONNode  `json:"children"`
}

func (n *Node) RenderShallowJSONNode() JSONNode {
	var rootNode JSONNode
	rootNode.Content = n.Text()
	rootNode.Id = n.ID()
	rootNode.AtId = n.atid
	rootNode.Changed = n.changed
	rootNode.AttrChanged = n.attrChanged
	rootNode.Removed = n.removed
	rootNode.Tid = n.tid
	rootNode.Name = n.tagName
	rootNode.Type = int(n.nodeType)
	rootNode.Ref = n.RefTree()
	rootNode.TypeName = n.Type().String()
	rootNode.Attrs = make([]JSONAttr, 0, n.Attrs.Len())
	rootNode.Events = make([]JSONEvent, 0, n.Events.Len())

	n.Attrs.Each(func(attr Attr) bool {
		rootNode.Attrs = append(rootNode.Attrs, JSONAttr{
			Key:   attr.Key(),
			Value: attr.Text(),
		})
		return true
	})

	for eventName, eventTargets := range n.Events.nodes {
		rootNode.Events = append(rootNode.Events, JSONEvent{
			Name:    eventName,
			Targets: append([]string{}, eventTargets...),
		})
	}

	return rootNode
}

func (n *Node) RenderJSONNode() JSONNode {
	var rootNode JSONNode
	rootNode.Content = n.Text()
	rootNode.Id = n.ID()
	rootNode.AtId = n.atid
	rootNode.Changed = n.changed
	rootNode.AttrChanged = n.attrChanged
	rootNode.Removed = n.removed
	rootNode.Tid = n.tid
	rootNode.Name = n.tagName
	rootNode.Type = int(n.nodeType)
	rootNode.Ref = n.RefTree()
	rootNode.TypeName = n.Type().String()
	rootNode.Attrs = make([]JSONAttr, 0, n.Attrs.Len())
	rootNode.Events = make([]JSONEvent, 0, n.Events.Len())

	n.Attrs.Each(func(attr Attr) bool {
		rootNode.Attrs = append(rootNode.Attrs, JSONAttr{
			Key:   attr.Key(),
			Value: attr.Text(),
		})
		return true
	})

	for eventName, eventTargets := range n.Events.nodes {
		rootNode.Events = append(rootNode.Events, JSONEvent{
			Name:    eventName,
			Targets: append([]string{}, eventTargets...),
		})
	}

	n.kids.Each(func(node *Node, i int) bool {
		rootNode.Children = append(rootNode.Children, node.RenderJSONNode())
		return true
	})

	return rootNode
}

func (n *Node) FromJSONNode(jsonNode JSONNode) error {
	n.nodeType = NodeType(jsonNode.Type)
	n.tagName = jsonNode.Name
	n.tid = jsonNode.Tid
	n.idAttr = jsonNode.Id
	n.removed = jsonNode.Removed
	n.atid = jsonNode.AtId
	n.content = TextContent(jsonNode.Content)

	for _, event := range jsonNode.Events {
		n.Events.AddJSONEvent(event)
	}

	for _, attr := range jsonNode.Attrs {
		n.Attrs.Add(NewStringAttr(attr.Key, attr.Value))
	}

	for _, child := range jsonNode.Children {
		var newChild = new(Node)
		if err := newChild.FromJSONNode(child); err != nil {
			return nerror.WrapOnly(err)
		}
		n.AppendChild(newChild)
	}

	return nil
}

// ReconcileStream runs the reconciliation process for this node against
// a provided old version. It produces a JSON list containing
// all changed nodes from removed nodes and nodes which have seen an update from
// a previous version in the old node.
//
// It will will not produce nodes in proper order, but in partial order where a changed
// node is rendered top-down till it's children but never it's parent. The node should contain
// adequate information (mainly through it's ref retrieved by Node.RefTree()) of it's ancestry
// and location.
func (n *Node) ReconcileStream(old *Node) []JSONNode {
	var changes []JSONNode
	if !n.Reconcile(old) {
		return changes
	}
	return n.GetChangeStream()
}

func (n *Node) GetChangeStream() []JSONNode {
	var changes []JSONNode

	// did we ourselves change ?
	if n.attrChanged {
		changes = append(changes, n.RenderJSONNode())
		return changes
	}

	n.EachChild(func(nm *Node, index int) bool {
		if nm.removed {
			changes = append(changes, n.RenderJSONNode())
			return false
		}

		// if not changed occurred, then move to next
		if !nm.changed {
			return true
		}

		// if we have empty kids and we still have changed
		// then we must be the target change then add parent.
		if nm.kids.Empty() {
			changes = append(changes, n.RenderJSONNode())
			return false
		}

		// if we are a text or comment node and then add parent.
		if nm.Type() == TextNode || nm.Type() == CommentNode {
			changes = append(changes, n.RenderJSONNode())
			return false
		}

		changes = append(changes, nm.GetChangeStream()...)
		return true
	})
	return changes
}

// EncodeObject implements the npkg.EncodableObject interface.
func (n *JSONNode) EncodeObject(encoder npkg.ObjectEncoder) {
	encoder.Int("type", n.Type)

	encoder.String("ref", n.Ref)

	encoder.String("typeName", n.TypeName)
	encoder.String("atid", n.AtId)
	encoder.Bool("changed", n.Changed)
	encoder.Bool("attr_changed", n.AttrChanged)
	encoder.String("name", n.Name)

	if n.Removed {
		encoder.Bool("removed", n.Removed)
	}

	if n.Content != "" {
		encoder.String("content", n.Content)
	}

	encoder.String("id", n.Id)
	encoder.String("tid", n.Tid)

	encoder.ListFor("attrs", func(encoder npkg.ListEncoder) {
		for _, attrItem := range n.Attrs {
			func(attr JSONAttr) {
				encoder.AddObjectWith(func(obj npkg.ObjectEncoder) {
					obj.String(attr.Key, attr.Value)
				})
			}(attrItem)
		}
	})

	encoder.ListFor("events", func(encoder npkg.ListEncoder) {
		for _, eventItem := range n.Events {
			func(event JSONEvent) {
				encoder.AddObjectWith(func(obj npkg.ObjectEncoder) {
					obj.String("name", event.Name)
					obj.ListFor("targets", func(targets npkg.ListEncoder) {
						for _, targetItem := range event.Targets {
							targets.AddString(targetItem)
						}
					})
				})
			}(eventItem)
		}
	})

	encoder.ListFor("children", func(kids npkg.ListEncoder) {
		for _, child := range n.Children {
			kids.AddObject(&child)
		}
	})
}
