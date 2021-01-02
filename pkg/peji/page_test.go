package peji_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/influx6/groundlayer/pkg/domu"
	"github.com/influx6/groundlayer/pkg/peji"
	"github.com/influx6/groundlayer/pkg/styled"
)

var sampleLayout = peji.LayoutFunc(func(page *peji.Page, data peji.Data, parent *domu.Node) {
	page.Live("users").Render(data).Mount(parent)
})

func TestPageLiveDOM(t *testing.T) {
	var sample = peji.WithPage("samples", &styled.Theme{}, sampleLayout, peji.DefaultNotFound{})

	var user = &UserComponent{}
	sample.AddLive("/users_list", user)

	require.True(t, sample.Has("users_list"))

	var response = sample.Serve(peji.Data{
		Path:   "/samples/users_list",
		Params: map[string]string{},
		Data:   "welcome",
	})

	require.NotNil(t, response)
	require.NotContains(t, response.Text(), "not found")
}
