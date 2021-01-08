package base

import (
	"fmt"
	"log"
	"strings"

	"github.com/influx6/npkg/nerror"

	"github.com/influx6/groundlayer/pkg/domu"
)

// Color represents a color type for a css style value.
type Color Palette

// URL represent a style url, image unit type in css (used by background-image, etc).
type URL string

type Directive struct {
	Screens  map[string]bool // screen variant
	Variants []string        // other variants
	Style    string          // the target style e.g text-gray-500, font-thin, ...etc
	Negated  bool            // has a ! at the end of the style to negate giving styling.
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
func ParseDirective(t string, theme *Theme) (Directive, error) {
	t = strings.TrimSpace(t)

	var directive Directive
	if t == "" {
		return directive, nerror.New("empty directive not allowed")
	}

	directive.Screens = map[string]bool{}

	var parts = strings.Split(t, ":")
	if len(parts) == 1 {
		directive.Style = parts[0]
		return directive, nil
	}

	var style = parts[len(parts)-1]

	if strings.HasSuffix(style, "!") {
		directive.Negated = true
		style = strings.TrimSuffix(style, "!")
	}

	var filteredVariants = make([]string, 0, len(parts))
	var variantList = parts[0 : len(parts)-1]
	for i := 0; i < len(variantList); i++ {
		var target = strings.TrimSpace(variantList[i])
		if theme.Screens.Values.Scale().IsKeyType(target) {
			directive.Screens[target] = true
			continue
		}
		filteredVariants = append(filteredVariants, target)
	}

	directive.Style = style
	directive.Variants = filteredVariants
	return directive, nil
}

// Stylus is a series of css strings which will be contained
// as the
type Stylus map[string]string

// Adds new style into the style list.
func (st Stylus) Add(key string, value string) {
	st[key] = value
}

type StylusGroup map[string]Stylus

func NewStylusGroup() StylusGroup {
	return StylusGroup{
		"*": Stylus{},
	}
}

func (st StylusGroup) GetOrAdd(variant string) *Stylus {
	var tl, ok = st[variant]
	if !ok {
		tl = Stylus{}
		st[variant] = tl
	}
	return &tl
}

func (st StylusGroup) Add(variant string, sl Stylus) {
	st[variant] = sl
}

type VariantToCSS interface {
	CSSFormat(variant string) string
}

// ThemeResolver will resolve the giving theme directives where the
// value (css property:key;) is stored in the provided Stylus.
//
// We provide the directive (style) and variant (hover, focus, ..etc)
// for you to generate accordingly the css key:value pair
type ThemeResolver interface {
	Resolve(theme *Theme, directive Directive, variant string, ts *Stylus)
}

type ThemeGenerator interface {
	Generate(themeDirective *domu.ThemeDirective, theme *Theme, root *domu.Node)
}

type ThemeResolvers []ThemeResolver

func (t *ThemeResolvers) Add(r ThemeResolver) {
	*t = append(*t, r)
}

func (t ThemeResolvers) Resolve(themeDirective *domu.ThemeDirective, theme *Theme, root *domu.Node) {
	for _, directiveValue := range themeDirective.Directives {
		var group = NewStylusGroup()

		var directive, directiveErr = ParseDirective(directiveValue, theme)
		if directiveErr != nil {
			log.Printf(`{"message":"failed to parse theme directive", "_level": 2, "directive": %q, "error": %+q}\n`, directiveValue, directiveErr)
			continue
		}

		// if there is no variants then run against the root
		if len(directive.Variants) == 0 {
			var rootStyle = group.GetOrAdd("*")
			for _, resolver := range t {
				resolver.Resolve(theme, directive, "", rootStyle)
			}

			continue
		}

		// get the styling for each variant
		for _, variant := range directive.Variants {
			var variantCSSName = theme.VariantNaming.CSSFormat(variant)
			var rs = group.GetOrAdd(variantCSSName)
			for _, resolver := range t {
				resolver.Resolve(theme, directive, variant, rs)
			}
		}

		t.writeGroupStyling(themeDirective, theme, root, group)
	}

}

func (t ThemeResolvers) writeGroupStyling(themeDirective *domu.ThemeDirective, theme *Theme, root *domu.Node, group StylusGroup) {
	var nodeId = themeDirective.Node.ID()
	var nodeTag = themeDirective.Node.Name()

	var resolutionCount = theme.Screens.Values.Scale().Count()

	var contents strings.Builder
	for i := 0; i < resolutionCount; i++ {
		contents.Reset()

		var resolution = theme.Screens.ValuesAtIndex(i)

		contents.WriteString(resolution)
		contents.WriteString(" {\n")

		for variantName, styling := range group {
			if variantName == "*" {
				// write root style.
				contents.WriteString(fmt.Sprintf("%s#%s", nodeTag, nodeId))
				contents.WriteString("\t{\n")
				for key, value := range styling {
					contents.WriteString(fmt.Sprintf("\t\t%s: %s;\n", key, value))
				}
				contents.WriteString("\t\n}\n")
				continue
			}

			contents.WriteString(fmt.Sprintf("%s#%s:%s", nodeTag, nodeId, variantName))
			contents.WriteString("\t{\n")
			for key, value := range styling {
				contents.WriteString(fmt.Sprintf("\t\t%s: %s;\n", key, value))
			}
			contents.WriteString("\t\n}\n")
		}

		contents.WriteString("\n}\n")

		root.AppendChild(domu.HTMLStyle(domu.Text(contents.String())))
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
	Screens       *String
	Colors        *ColorDirective
	Resolvers     ThemeGenerator
	VariantNaming VariantToCSS
}

func (t *Theme) Resolve(directive *domu.ThemeDirective, root *domu.Node) {
	t.Resolvers.Generate(directive, t, root)
}
