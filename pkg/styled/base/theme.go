package base

import (
	"fmt"
	"strings"

	"github.com/influx6/npkg/nerror"

	"github.com/influx6/groundlayer/pkg/domu"
)

// Color represents a color type for a css style value.
type Color Palette

// URL represent a style url, image unit type in css (used by background-image, etc).
type URL string

const darkVariant = "dark"

type Directive struct {
	Screens  map[string]bool // screen variant
	Variants []string        // other variants
	Style    string          // the target style e.g text-gray-500, font-thin, ...etc
	Negated  bool            // has a ! at the end of the style to negate giving styling.
	Dark     bool
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
		if theme.Screens.GetScale().Key(target) {
			directive.Screens[target] = true
			continue
		}
		if strings.ToLower(target) == darkVariant {
			directive.Dark = true
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
	return StylusGroup{}
}

func (st StylusGroup) GetOrAdd(name string) *Stylus {
	var tl, ok = st[name]
	if !ok {
		tl = Stylus{}
		st[name] = tl
	}
	return &tl
}

func (st StylusGroup) Add(name string, sl Stylus) {
	var vt, ok = st[name]
	if !ok {
		st[name] = sl
		return
	}
	for key, value := range sl {
		vt[key] = value
	}
	st[name] = vt
}

// ThemeResolver are implementations which exists to handle specific
// forms of styling based on the Directive and variant supplied.
//
// Each can indicate if it can handle such a directive and variant
// and if so, will return the group of css properties (key:value pairs)
// and the full name for the giving style.
//
// Please ensure to take into account the exact css style name as the
// selector of the target node is supplied and it's the responsibility
// of the resolver to return what should be the css property selector
// for it's generated styles.
type ThemeResolver interface {
	Can(directive Directive, variant string) bool
	Resolve(theme *Theme, directive Directive, variant string, targetSelector string) (string, Stylus)
}

type ThemeGenerator interface {
	Generate(themeDirective *domu.ThemeDirective, theme *Theme, root *domu.Node) error
}

var _ ThemeGenerator = (*ThemeResolvers)(nil)

const baseResolution = "*"

type ThemeResolvers []ThemeResolver

func (t *ThemeResolvers) Add(r ThemeResolver) {
	*t = append(*t, r)
}

func (t ThemeResolvers) parseDirectives(themeDirective *domu.ThemeDirective, theme *Theme) (map[string][]Directive, error) {
	var resolutionGroupings = map[string][]Directive{}
	for _, directiveValue := range themeDirective.Directives {
		var directive, directiveErr = ParseDirective(directiveValue, theme)
		if directiveErr != nil {
			return nil, nerror.WrapOnly(directiveErr)
		}

		// if there is no screen then this is a non-scoped styling
		if len(directive.Screens) == 0 {
			resolutionGroupings[baseResolution] = append(resolutionGroupings[baseResolution], directive)
			continue
		}

		for resolution := range directive.Screens {
			resolutionGroupings[resolution] = append(resolutionGroupings[resolution], directive)
		}
	}
	return resolutionGroupings, nil
}

func (t ThemeResolvers) Generate(themeDirective *domu.ThemeDirective, theme *Theme, root *domu.Node) error {
	var screenGroups, groupingErr = t.parseDirectives(themeDirective, theme)
	if groupingErr != nil {
		return nerror.WrapOnly(groupingErr)
	}

	var screenNodes = map[string]*domu.Node{}

	for screen, directives := range screenGroups {
		// if the screen is not part of our scale, then skip.
		if screen == baseResolution || !theme.Screens.GetScale().Key(screen) {
			continue
		}
		screenNodes[screen] = t.writeStylingForGroup(screen, directives, themeDirective, theme)
	}

	// append in order of screen resolution mapping
	var count = theme.Screens.GetScale().Count()
	for i := count - 1; i > -1; i-- {
		var keyName = theme.Screens.KeyAtIndex(i)
		var targetNode, hasNode = screenNodes[keyName]
		if !hasNode {
			continue
		}

		root.PrependChild(targetNode)
	}

	var baseDirectives = screenGroups[baseResolution]
	var baseNode = t.writeStylingForGroup(baseResolution, baseDirectives, themeDirective, theme)
	root.PrependChild(baseNode)
	return nil
}

func (t ThemeResolvers) writeStylingForGroup(screen string, directives []Directive, themeDirective *domu.ThemeDirective, theme *Theme) *domu.Node {
	var group = NewStylusGroup()

	for _, directive := range directives {
		t.parseDirective(directive, themeDirective, theme, group)
	}

	return t.writeGroupStyling(screen, theme, group)
}

func (t ThemeResolvers) parseDirective(directive Directive, themeDirective *domu.ThemeDirective, theme *Theme, group StylusGroup) {
	// if there is no variants then run against the root
	if len(directive.Variants) == 0 {
		for _, resolver := range t {
			if !resolver.Can(directive, "") {
				continue
			}
			var styleName, styling = resolver.Resolve(theme, directive, "", themeDirective.Node.Selector())
			group.Add(styleName, styling)
		}

		return
	}

	// get the styling for each variant
	for _, variant := range directive.Variants {
		for _, resolver := range t {
			if !resolver.Can(directive, variant) {
				continue
			}
			var styleName, styling = resolver.Resolve(theme, directive, variant, themeDirective.Node.Selector())
			group.Add(styleName, styling)
		}
	}
}

func (t ThemeResolvers) writeGroupStyling(resolution string, theme *Theme, group StylusGroup) *domu.Node {
	var contents strings.Builder
	contents.Reset()

	if resolution != baseResolution {
		var resolutionFormat = theme.Screens.ValueAtKey(resolution)
		contents.WriteString(resolutionFormat)
		contents.WriteString(" {\t")
	}

	for sectionName, styling := range group {
		contents.WriteString(sectionName)
		contents.WriteString("\t{\t")
		for key, value := range styling {
			contents.WriteString(fmt.Sprintf("\t\t%s: %s;", key, value))
		}
		contents.WriteString("\t}\t")
	}

	if resolution != baseResolution {
		contents.WriteString("\t}\t")
	}

	return domu.HTMLStyle(
		domu.Text(contents.String()),
		domu.NewStringAttr("type", "text/css"),
		domu.NewStringAttr("_resolution", resolution),
	)
}

/**
Theme provides a clear description of a predefined styling which will
be used to generate unique styles for each target element. It provides
fields to cover the following styling.

	Main Styles
*/
type Theme struct {
	Resolvers ThemeGenerator

	Screens *String
	Colors  map[string]*ColorDirective
	// spacing: ThemeSection
	// durations: ThemeSection<string | string[]>
	// animation: ThemeSection<string | string[]>
	// backgroundColor: ThemeSection<ThemeColor>
	// backgroundImage: ThemeSection<string | string[]>
	// backgroundOpacity: ThemeSection
	// borderColor: ThemeSection<ThemeColor>
	// borderOpacity: ThemeSection
	// borderRadius: ThemeSection
	// borderWidth: ThemeSection
	// boxShadow: ThemeSection<string | string[]>
	// container: ThemeContainer | ThemeSectionResolver<ThemeContainer>
	// fill: ThemeSection<ThemeColor>
	// flex: ThemeSection
	// fontFamily: ThemeSection<string | string[]>
	// fontSize: ThemeSection<ThemeFontSize>
	// fontWeight: ThemeSection
	// gap: ThemeSection
	// gradientColorStops: ThemeSection<ThemeColor>
	// height: ThemeSection
	// inset: ThemeSection
	// letterSpacing: ThemeSection
	// lineHeight: ThemeSection
	// margin: ThemeSection
	// maxHeight: ThemeSection
	// maxWidth: ThemeSection
	// minHeight: ThemeSection
	// minWidth: ThemeSection
	// opacity: ThemeSection
	// padding: ThemeSection
	// placeholderColor: ThemeSection<ThemeColor>
	// placeholderOpacity: ThemeSection
	// rotate: ThemeSection
	// scale: ThemeSection
	// skew: ThemeSection
	// space: ThemeSection
	// stroke: ThemeSection<ThemeColor>
	// strokeWidth: ThemeSection
	// textColor: ThemeSection<ThemeColor>
	// textOpacity: ThemeSection
	// width: ThemeSection
	// zIndex: ThemeSection
	//
	// order: ThemeSection
	// outline: ThemeSection<ThemeOutline>
	// keyframes: ThemeSection<CSSAtKeyframes>
	// transitionDelay: ThemeSection<string | string[]>
	// transitionDuration: ThemeSection<string | string[]>
	// transitionProperty: ThemeSection<string | string[]>
	// transitionTimingFunction: ThemeSection<string | string[]>
	// translate: ThemeSection
	//
	// divideColor: ThemeSection<ThemeColor>
	// divideOpacity: ThemeSection
	// divideWidth: ThemeSection
	//
	// ringColor: ThemeSection<ThemeColor>
	// ringOffsetColor: ThemeSection<ThemeColor>
	// ringOffsetWidth: ThemeSection
	// ringOpacity: ThemeSection
	// ringWidth: ThemeSection
}

func (t *Theme) Resolve(directive *domu.ThemeDirective, root *domu.Node) error {
	var err = t.Resolvers.Generate(directive, t, root)
	if err != nil {
		return nerror.WrapOnly(err)
	}
	return nil
}

func (t *Theme) MustResolve(directive *domu.ThemeDirective, root *domu.Node) {
	if err := t.Resolvers.Generate(directive, t, root); err != nil {
		panic(nerror.WrapOnly(err).Error())
	}
}
