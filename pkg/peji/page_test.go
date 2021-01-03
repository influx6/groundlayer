package peji_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/influx6/groundlayer/pkg/domu"
	"github.com/influx6/groundlayer/pkg/peji"
	"github.com/influx6/groundlayer/pkg/styled"
)

var sampleLayout = peji.LayoutFunc(func(page *peji.Page, data peji.Data, parent *domu.Node) {
	var provider = page.Live("users")
	var item = provider.Render(data)
	var result = domu.Carrier(domu.HTMLHead(), domu.HTMLBody(item))
	result.Mount(parent)
})

func TestPageLiveDOM(t *testing.T) {
	var sample = peji.WithPage("samples", &styled.Theme{}, sampleLayout, peji.DefaultNotFound{})

	var user = &UserComponent{}
	sample.AddLive("users", user)

	require.True(t, sample.Has("/users"))

	var response = sample.Serve(peji.Data{
		Path:   "/samples/users",
		Params: map[string]string{},
		Data:   "welcome",
	})

	require.NotNil(t, response)

	var content strings.Builder
	var err = response.RenderHTML(&content, true)
	require.NoError(t, err)

	var contentText = content.String()
	require.NotContains(t, contentText, "not found", contentText)
}
