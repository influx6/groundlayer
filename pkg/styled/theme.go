package styled

import (
	"strings"

	"github.com/influx6/groundlayer/pkg/domu"
)

type ThemeDirective []string

func Themes(t ...string) ThemeDirective {
	return ThemeDirective(t)
}

func (td *ThemeDirective) Add(t string) {
	*td = append(*td, t)
}

func (td *ThemeDirective) Mount(p *domu.Node) {
	p.Themes = *td
}

var rootResolvers ThemeResolvers

func AddResolver(resolver ...ThemeResolver) {
	for _, resolver := range resolver {
		rootResolvers = append(rootResolvers, resolver)
	}
}

func GetResolver() ThemeResolvers {
	return rootResolvers[0:len(rootResolvers)]
}

type ThemeResolver interface {
	Resolve(directive ThemeDirective, builder *strings.Builder)
}

type ThemeResolvers []ThemeResolver

func (t ThemeResolvers) Resolve(directive ThemeDirective, builder *strings.Builder) {
	for _, resolver := range t {
		resolver.Resolve(directive, builder)
	}
}
