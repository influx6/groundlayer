package domu

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/influx6/npkg"
)

// EventHashList implements the a set list for Nodes using
// their Node.RefID() value as unique keys.
type EventHashList struct {
	nodes map[string][]string
}

// NewEventHashList returns a new instance EventHashList.
func NewEventHashList() *EventHashList {
	return &EventHashList{
		nodes: map[string][]string{},
	}
}

// Len returns the underline length of events in map.
func (na *EventHashList) Len() int {
	return len(na.nodes)
}

// MatchEvents returns true if both are equal in keys and values.
func (na *EventHashList) MatchEvents(other *EventHashList) bool {
	for name, items := range na.nodes {
		if otherItems, hasItem := other.nodes[name]; hasItem {
			for _, itemVal := range items {
				var found = false
				for _, otherVal := range otherItems {
					if itemVal == otherVal {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			}
			continue
		}
		return false
	}
	return true
}

// EncodeList encodes underline events into provided list encoder.
func (na *EventHashList) EncodeList(enc npkg.ListEncoder) {
	for name, events := range na.nodes {
		if len(events) == 0 {
			continue
		}
		name = strcase.ToKebab(name)
		enc.AddObjectWith(func(encoder npkg.ObjectEncoder) {
			encoder.String("name", name)
			encoder.ListFor("targets", func(encoder npkg.ListEncoder) {
				for _, item := range events {
					encoder.AddString(item)
				}
			})
		})
	}
}

// Reset resets the internal hashmap used for storing
// nodes. There by removing all registered nodes.
func (na *EventHashList) Reset() {
	na.nodes = map[string][]string{}
}

// Count returns the total content count of map
func (na *EventHashList) Count() int {
	if na.nodes == nil {
		return 0
	}
	return len(na.nodes)
}

// EncodeEvents encodes all giving event within provided event hash list.
func (na *EventHashList) EncodeEvents(encoder *strings.Builder) error {
	if na.nodes != nil {
		var count int
		for name, events := range na.nodes {
			if len(events) == 0 {
				continue
			}
			if count > 0 {
				if _, err := encoder.WriteString(spacer); err != nil {
					return err
				}
			}

			name = strcase.ToKebab(name)
			var eventString = fmt.Sprintf("event-%s=%q", name, strings.Join(events, "|"))
			if _, err := encoder.WriteString(eventString); err != nil {
				return err
			}
			count++
		}
	}
	return nil
}

// AddJSONEvent adds giving node into giving list if it has
// giving attribute value.
func (na *EventHashList) AddJSONEvent(ev JSONEvent) {
	na.Add(ev.Name, ev.Targets...)
}

// Add adds giving node into giving list if it has
// giving attribute value.
func (na *EventHashList) Add(eventName string, callNames ...string) {
	if na.nodes == nil {
		na.nodes = map[string][]string{}
	}

	eventName = strcase.ToKebab(eventName)
	var name = strings.ToLower(eventName)
	na.nodes[name] = append(na.nodes[name], callNames...)
}

// RemoveAll removes giving node in list if it has
// giving attribute value.
func (na *EventHashList) RemoveAll(event string) {
	if na.nodes == nil {
		na.nodes = map[string][]string{}
	}
	delete(na.nodes, event)
}

// Remove removes giving node in list if it has
// giving handler.
func (na *EventHashList) Remove(eventName string, callNames ...string) {
	if na.nodes == nil {
		na.nodes = map[string][]string{}
	}

	var name = strings.ToLower(eventName)
	var set = na.nodes[name]
	for index, desc := range set {
		for _, callName := range callNames {
			if desc == callName {
				set = append(set[:index], set[index+1:]...)
				na.nodes[name] = set
			}
		}
	}

	if len(set) == 0 {
		delete(na.nodes, name)
	}
}
