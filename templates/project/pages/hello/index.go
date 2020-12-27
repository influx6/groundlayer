package hello

import (
	"fmt"

	"github.com/influx6/sabuhp"

	"github.com/influx6/groundlayer/pkg/domu"
	"github.com/influx6/groundlayer/pkg/peji"
)

var HelloWorldNotFound = peji.HandlerFunc(func(data peji.Data) *domu.Node {
	return domu.Text(fmt.Sprintf("Hello World does not have this %q!", data.Path))
})

type HelloWorldLayout struct{}

func (h HelloWorldLayout) Render(p *peji.Page, data peji.Data, parent *domu.Node) {
	parent.AppendChild(domu.Text(fmt.Sprintf("Hello World %q!", data.Path)))
}

func CreateHelloWorldPage(pageRoute string, tr sabuhp.Transport) *peji.Page {
	var newPage = peji.WithPage(pageRoute, &HelloWorldLayout{}, HelloWorldNotFound)
	return newPage
}
