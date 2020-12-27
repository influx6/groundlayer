package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/influx6/groundlayer/pkg/domu"
)

func main() {
	var base = domu.Element("section", "767h")

	for i := 0; i < 10; i++ {
		var digit = fmt.Sprintf("%d", i)
		if err := base.AppendChild(domu.Comment("Commentary")); err != nil {
			log.Fatalf("bad things occured: %s\n", err)
		}
		if err := base.AppendChild(
			domu.Element(
				"div",
				digit,
				domu.ClickEvent(),
				domu.MouseOverEvent(),
				domu.NewStringAttr("count-target", digit),
				domu.Text(digit),
			),
		); err != nil {
			log.Fatalf("bad things occured: %+s\n", err)
		}
	}

	// Render html into giving builder.
	var content strings.Builder
	if err := base.RenderHTMLTo(&content, true, true); err != nil {
		log.Fatalf("failed to render: %s\n", err)
	}

	fmt.Println(content.String())
}
