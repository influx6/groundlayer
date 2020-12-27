package peji_test

import (
	"testing"

	"github.com/influx6/groundlayer/pkg/peji"

	"github.com/stretchr/testify/require"

	"github.com/influx6/groundlayer/pkg/domu"
)

type UserComponent struct{}

func (u *UserComponent) Render(data peji.Data) *domu.Node {
	return domu.HTMLDiv(domu.Text("user"))
}

func TestKomponentRegistry(t *testing.T) {
	var registry = peji.NewDOMRegistry()

	var user = &UserComponent{}
	registry.Add(user, "users")
	require.True(t, registry.Has("users"))
}

func TestKomponentRegistry_WithList(t *testing.T) {
	var registry = peji.NewDOMRegistry()

	var users = []peji.DOM{&UserComponent{}}
	registry.AddList(users, "users")
	require.True(t, registry.HasList("users"))
}
