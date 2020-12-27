package peji

import (
	"io"

	"github.com/influx6/npkg"
	"github.com/influx6/npkg/nerror"
	"github.com/influx6/npkg/njson"

	"github.com/influx6/groundlayer/pkg/domu"
)

func RenderVoidHTML(node *domu.Node, target io.Writer) error {
	if renderErr := node.RenderHTML(target, false); renderErr != nil {
		return renderErr
	}
	return nil
}

func RenderVoidHTMLDiff(node *domu.Node, target io.Writer) error {
	if renderErr := node.RenderHTMLDiff(target, false); renderErr != nil {
		return nerror.WrapOnly(renderErr)
	}
	return nil
}

func RenderVoidJSON(node *domu.Node, target io.Writer) error {
	var nodeJSON = node.RenderJSONNode()
	var jsonWriter = njson.JSONB()
	nodeJSON.EncodeObject(jsonWriter)
	if jsonWriter.Err() != nil {
		return nerror.WrapOnly(jsonWriter.Err())
	}
	if _, renderErr := jsonWriter.WriteTo(target); renderErr != nil {
		return nerror.WrapOnly(renderErr)
	}
	return nil
}

func RenderVoidJSONStream(node *domu.Node, target io.Writer) error {
	var jsonWriter = njson.JSONL(func(event npkg.Encoder) {
		var nodeJSON = node.GetChangeStream()
		for _, node := range nodeJSON {
			func(n domu.JSONNode) {
				event.AddObjectWith(func(encoder npkg.ObjectEncoder) {
					n.EncodeObject(encoder)
				})
			}(node)
		}
	})
	if jsonWriter.Err() != nil {
		return nerror.WrapOnly(jsonWriter.Err())
	}
	if _, renderErr := jsonWriter.WriteTo(target); renderErr != nil {
		return nerror.WrapOnly(renderErr)
	}
	return nil
}
