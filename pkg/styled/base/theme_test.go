package base

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/influx6/groundlayer/pkg/domu"
)

var _ ThemeResolver = (*hoverResolver)(nil)

type focusResolver struct{}

func (h focusResolver) Can(directive Directive, variant string) bool {
	return variant == "focus"
}

func (h focusResolver) Resolve(theme *Theme, directive Directive, variant string, target string) (string, Stylus) {
	var sectionName = fmt.Sprintf("%s:focus", target)
	var styling = Stylus{
		"text-size": "10px",
	}
	return sectionName, styling
}

type redColorResolver struct{}

func (h redColorResolver) Can(directive Directive, variant string) bool {
	return variant == "hover"
}

func (h redColorResolver) Resolve(theme *Theme, directive Directive, variant string, target string) (string, Stylus) {
	var styling = Stylus{
		"color": "red",
	}
	return target, styling
}

type hoverRedResolver struct{}

func (h hoverRedResolver) Can(directive Directive, variant string) bool {
	return variant == "hover"
}

func (h hoverRedResolver) Resolve(theme *Theme, directive Directive, variant string, target string) (string, Stylus) {
	var sectionName = fmt.Sprintf("%s:hover", target)
	var styling = Stylus{
		"color": "red",
	}
	return sectionName, styling
}

type hoverResolver struct{}

func (h hoverResolver) Can(directive Directive, variant string) bool {
	return variant == "hover"
}

func (h hoverResolver) Resolve(theme *Theme, directive Directive, variant string, target string) (string, Stylus) {
	var sectionName = fmt.Sprintf("%s:hover", target)
	var styling = Stylus{
		"text-size": "30px",
	}
	return sectionName, styling
}

var defaultScreenScales = StringScale("sm", "md", "lg", "xl", "2xl")

var chaka = NewString(defaultScreenScales).
	Key("sm", `@media (min-width: 30em)`).
	Key("md", `@media (min-width: 48em)`).
	Key("lg", `@media (min-width: 62em)`).
	Key("xl", `@media (min-width: 80em)`).
	Key("2xl", `@media (min-width: 94em)`)

func TestThemeWithNoResolution(t *testing.T) {
	var th = Theme{
		Screens: chaka,
		Resolvers: ThemeResolvers{
			&hoverResolver{},
			&focusResolver{},
			&redColorResolver{},
			&hoverRedResolver{},
		},
	}

	var targetNode = domu.HTMLDiv()
	targetNode.UseID("wallet")

	var rootNode = domu.HTMLDiv(targetNode)
	require.Equal(t, rootNode.ChildCount(), 1)

	var directives = &domu.ThemeDirective{
		Directives: []string{"hover:font-10", "focus:font-30"},
		Node:       targetNode,
	}

	require.NoError(t, th.Resolve(directives, rootNode))

	var content strings.Builder
	require.NoError(t, rootNode.RenderHTML(&content, false))

	require.Equal(t, rootNode.ChildCount(), 2)
	require.Contains(t, content.String(), "<style", content.String())
	require.Contains(t, content.String(), "</style>", content.String())
	require.NotContains(t, content.String(), "@media (min-width: 30em)", content.String())
	require.NotContains(t, content.String(), "_resolution=\"sm\"", content.String())
	require.Contains(t, content.String(), "_resolution=\"*\"", content.String())
}

func TestTheme(t *testing.T) {
	var th = Theme{
		Screens: chaka,
		Resolvers: ThemeResolvers{
			&hoverResolver{},
			&focusResolver{},
			&redColorResolver{},
			&hoverRedResolver{},
		},
	}

	var targetNode = domu.HTMLDiv()
	targetNode.UseID("wallet")

	var rootNode = domu.HTMLDiv(targetNode)
	require.Equal(t, rootNode.ChildCount(), 1)

	var directives = &domu.ThemeDirective{
		Directives: []string{"hover:font-10", "focus:font-30", "sm:focus:focus-30"},
		Node:       targetNode,
	}

	require.NoError(t, th.Resolve(directives, rootNode))

	var content strings.Builder
	require.NoError(t, rootNode.RenderHTML(&content, false))

	require.Equal(t, rootNode.ChildCount(), 3)
	require.Contains(t, content.String(), "<style", content.String())
	require.Contains(t, content.String(), "</style>", content.String())
	require.Contains(t, content.String(), "@media (min-width: 30em)", content.String())
	require.Contains(t, content.String(), "_resolution=\"sm\"", content.String())
	require.Contains(t, content.String(), "_resolution=\"*\"", content.String())
}
