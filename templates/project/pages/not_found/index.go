package not_found

import (
	"fmt"

	"github.com/influx6/groundlayer/pkg/domu"
	"github.com/influx6/groundlayer/pkg/peji"
)

var NotFound = peji.HandlerFunc(func(data peji.Data) *domu.Node {
	return domu.Text(fmt.Sprintf("%s not found!", data.Path))
})
