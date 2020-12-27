package peji_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/influx6/groundlayer/pkg/domu"
	"github.com/influx6/groundlayer/pkg/peji"
)

var sampleLayout = peji.LayoutFunc(func(page *peji.Page, data peji.Data, parent *domu.Node) {
	page.Static("users").Render(data).Mount(parent)
})

func TestPageStaticDOM(t *testing.T) {
	var sample = peji.WithPage("samples", sampleLayout, peji.DefaultNotFound{})

	var user = &UserComponent{}
	sample.AddStatic("users_list", user)

	require.True(t, sample.HasStatic("users_list"))
	require.False(t, sample.HasLive("users"))

	var response = sample.Serve(peji.Data{
		Path:   "/samples/users_list",
		Params: map[string]string{},
		Data:   "welcome",
	})

	require.NotNil(t, response)
	require.NotContains(t, response.Text(), "not found")
}

func TestPageLiveDOM(t *testing.T) {
	var sample = peji.WithPage("samples", sampleLayout, peji.DefaultNotFound{})

	var user = &UserComponent{}
	sample.AddStatic("users_list", user)
	sample.AddLive("/users_live_list", "users_list")

	require.True(t, sample.HasStatic("users_list"))
	require.False(t, sample.HasLive("users"))

	var response = sample.Serve(peji.Data{
		Path:   "/samples/users_live_list",
		Params: map[string]string{},
		Data:   "welcome",
	})

	require.NotNil(t, response)
	require.NotContains(t, response.Text(), "not found")
}
