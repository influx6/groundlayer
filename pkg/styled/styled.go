package styled

import (
	"github.com/influx6/groundlayer/pkg/domu"
)

// Color represents a color type for a css style value.
type Color Palette

// URL represent a style url, image unit type in css (used by background-image, etc).
type URL string

// ThemeResolver will resolve the giving theme directives as <style> nodes using node's special 'tid' or `id`
// to apply direct changes.
type ThemeResolver interface {
	Resolve(theme *Theme, directive domu.ThemeDirective, root *domu.Node)
}

type ThemeResolvers []ThemeResolver

func (t *ThemeResolvers) Add(r ThemeResolver) {
	*t = append(*t, r)
}

func (t ThemeResolvers) Resolve(theme *Theme, directive domu.ThemeDirective, root *domu.Node) {
	for _, resolver := range t {
		resolver.Resolve(theme, directive, root)
	}
}

/**
Theme provides a clear description of a predefined styling which will
be used to generate unique styles for each target element. It provides
fields to cover the following styling.

	Main Styles
		 colors: ThemeSection<ThemeColor>
		 spacing: ThemeSection
		 durations: ThemeSection<string | string[]>
		 screens: ThemeSection
		 animation: ThemeSection<string | string[]>
		 backgroundColor: ThemeSection<ThemeColor>
		 backgroundImage: ThemeSection<string | string[]>
		 backgroundOpacity: ThemeSection
		 borderColor: ThemeSection<ThemeColor>
		 borderOpacity: ThemeSection
		 borderRadius: ThemeSection
		 borderWidth: ThemeSection
		 boxShadow: ThemeSection<string | string[]>
		 container: ThemeContainer | ThemeSectionResolver<ThemeContainer>
		 fill: ThemeSection<ThemeColor>
		 flex: ThemeSection
		 fontFamily: ThemeSection<string | string[]>
		 fontSize: ThemeSection<ThemeFontSize>
		 fontWeight: ThemeSection
		 gap: ThemeSection
		 gradientColorStops: ThemeSection<ThemeColor>
		 height: ThemeSection
		 inset: ThemeSection
		 letterSpacing: ThemeSection
		 lineHeight: ThemeSection
		 margin: ThemeSection
		 maxHeight: ThemeSection
		 maxWidth: ThemeSection
		 minHeight: ThemeSection
		 minWidth: ThemeSection
		 opacity: ThemeSection
		 padding: ThemeSection
		 placeholderColor: ThemeSection<ThemeColor>
		 placeholderOpacity: ThemeSection
		 rotate: ThemeSection
		 scale: ThemeSection
		 skew: ThemeSection
		 space: ThemeSection
		 stroke: ThemeSection<ThemeColor>
		 strokeWidth: ThemeSection
		 textColor: ThemeSection<ThemeColor>
		 textOpacity: ThemeSection
		 width: ThemeSection
		 zIndex: ThemeSection

	Custom Styles
		order: ThemeSection
		outline: ThemeSection<ThemeOutline>
		keyframes: ThemeSection<CSSAtKeyframes>
		transitionDelay: ThemeSection<string | string[]>
		transitionDuration: ThemeSection<string | string[]>
		transitionProperty: ThemeSection<string | string[]>
		transitionTimingFunction: ThemeSection<string | string[]>
		translate: ThemeSection

		divideColor: ThemeSection<ThemeColor>
		divideOpacity: ThemeSection
		divideWidth: ThemeSection

		ringColor: ThemeSection<ThemeColor>
		ringOffsetColor: ThemeSection<ThemeColor>
		ringOffsetWidth: ThemeSection
		ringOpacity: ThemeSection
		ringWidth: ThemeSection
*/
type Theme struct {
	Screens   *ScreenDirective
	Colors    *ColorDirective
	Resolvers ThemeResolvers
}

func (t *Theme) Resolve(directive domu.ThemeDirective, root *domu.Node) {
	t.Resolvers.Resolve(t, directive, root)
}
