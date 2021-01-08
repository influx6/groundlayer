package styled

import (
	"strings"

	"github.com/influx6/npkg/nerror"

	"github.com/influx6/groundlayer/pkg/domu"
)

// Color represents a color type for a css style value.
type Color Palette

// URL represent a style url, image unit type in css (used by background-image, etc).
type URL string

type Directive struct {
	Screen   string   // screen variant
	Variants []string // other variants
	Style    string   // the target style e.g text-gray-500, font-thin, ...etc
	Negated  bool     // has a ! at the end of the style to negate giving styling.
}

// ParseDirective parses incoming directive string, which we follow a simple format:
//
// 1. If directive has one or multiple `text:` parts then this is sorted
//    as variants on the giving directive. e.g sm: for (small screens size), hover: for hover state, ..etc
//    e.g sm:hover:text-red-500
//
// 2. if a directive has a `1` at the end then we enable negation on directive.
//      e.g font-8 font-8!
//              where font-8! => resets the font size to `inherit`
//
func ParseDirective(t string) (Directive, error) {
	t = strings.TrimSpace(t)

	var directive Directive
	if t == "" {
		return directive, nerror.New("empty directive not allowed")
	}

	var parts = strings.Split(t, ":")
	if len(parts) == 1 {
		directive.Style = parts[0]
		return directive, nil
	}

	var style = parts[len(parts)-1]
	var variants = parts[0 : len(parts)-1]

	if strings.HasSuffix(style, "!") {
		directive.Negated = true
		style = strings.TrimSuffix(style, "!")
	}

	directive.Style = style
	directive.Variants = variants
	return directive, nil
}

// Stylus is a series of css strings which will be contained
// as the
type Stylus []string

// Adds new style into the style list.
func (st *Stylus) Add(style string) {
	*st = append(*st, style)
}

// ThemeResolver will resolve the giving theme directives as <style> nodes using node's special 'tid' or `id`
// to apply direct changes.
type ThemeResolver interface {
	Resolve(theme *Theme, directive string, ts *Stylus)
}

type ThemeResolvers []ThemeResolver

func (t *ThemeResolvers) Add(r ThemeResolver) {
	*t = append(*t, r)
}

func (t ThemeResolvers) Resolve(themeDirective *domu.ThemeDirective, theme *Theme, root *domu.Node) {
	// var stylus Stylus
	// for _, directiveValue := range themeDirective.Directives {
	//
	// 	var directive, directiveErr = ParseDirective(directiveValue)
	// 	if directiveErr != nil {
	// 		log.Printf("Failed to parse directive %q: %s\n", directiveValue, directiveErr)
	// 		continue
	// 	}
	//
	// 	// for _, resolver := range t {
	// 	// 	// resolver.
	// 	// }
	// }
	// for _, resolver := range t {
	//
	// 	resolver.Resolve(theme, directive, root)
	// }
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

func (t *Theme) Resolve(directive *domu.ThemeDirective, root *domu.Node) {
	t.Resolvers.Resolve(directive, t, root)
}
