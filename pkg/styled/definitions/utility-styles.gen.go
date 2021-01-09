package definitions

import "time"

// BorderInlineStartColor represent the CSS style "border-inline-start-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-start-color
type BorderInlineStartColorStyle Color

func (t BorderInlineStartColorStyle) Name() string {
	return "border-inline-start-color"
}

var BorderInlineStartColorStyleBrowserVariantsList = []string{}

func (t BorderInlineStartColorStyle) BrowserVariants() []string {
	return BorderInlineStartColorStyleBrowserVariantsList
}

var BorderInlineStartColorStyleUtilitiesMapping = map[string]string{
	"border-inline-start-color": "border-inline-start-color: currentcolor",
}

func (t BorderInlineStartColorStyle) Utilities() map[string]string {
	return BorderInlineStartColorStyleUtilitiesMapping
}

func (t BorderInlineStartColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderInlineStartColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderLeftColor represent the CSS style "border-left-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-left-color
type BorderLeftColorStyle Color

func (t BorderLeftColorStyle) Name() string {
	return "border-left-color"
}

var BorderLeftColorStyleBrowserVariantsList = []string{}

func (t BorderLeftColorStyle) BrowserVariants() []string {
	return BorderLeftColorStyleBrowserVariantsList
}

var BorderLeftColorStyleUtilitiesMapping = map[string]string{
	"border-left-color": "border-left-color: currentcolor",
}

func (t BorderLeftColorStyle) Utilities() map[string]string {
	return BorderLeftColorStyleUtilitiesMapping
}

func (t BorderLeftColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderLeftColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// RubyAlign represent the CSS style "ruby-align" with value "start | center | space-between | space-around"
// See https://developer.mozilla.org/docs/Web/CSS/ruby-align
type RubyAlignStyle string

func (t RubyAlignStyle) Name() string {
	return "ruby-align"
}

const RubyAlignStyleStart RubyAlignStyle = "start"

const RubyAlignStyleCenter RubyAlignStyle = "center"

const RubyAlignStyleSpaceBetween RubyAlignStyle = "space-between"

const RubyAlignStyleSpaceAround RubyAlignStyle = "space-around"

var RubyAlignStyleBrowserVariantsList = []string{}

func (t RubyAlignStyle) BrowserVariants() []string {
	return RubyAlignStyleBrowserVariantsList
}

var RubyAlignStyleUtilitiesMapping = map[string]string{
	"ruby-align":               "ruby-align: space-around",
	"ruby-align-start":         "ruby-align: start",
	"ruby-align-center":        "ruby-align: center",
	"ruby-align-space-between": "ruby-align: space-between",
	"ruby-align-space-around":  "ruby-align: space-around",
}

func (t RubyAlignStyle) Utilities() map[string]string {
	return RubyAlignStyleUtilitiesMapping
}

func (t RubyAlignStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = RubyAlignStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformTranslateZ represent the CSS style "transform-translate-z" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformTranslateZStyle string

func (t TransformTranslateZStyle) Name() string {
	return "transform-translate-z"
}

var TransformTranslateZStyleBrowserVariantsList = []string{}

func (t TransformTranslateZStyle) BrowserVariants() []string {
	return TransformTranslateZStyleBrowserVariantsList
}

var TransformTranslateZStyleUtilitiesMapping = map[string]string{
	"transform-translate-z": "transform-translate-z: translateZ(0)",
}

func (t TransformTranslateZStyle) Utilities() map[string]string {
	return TransformTranslateZStyleUtilitiesMapping
}

func (t TransformTranslateZStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformTranslateZStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformTranslateX represent the CSS style "transform-translate-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformTranslateXStyle string

func (t TransformTranslateXStyle) Name() string {
	return "transform-translate-x"
}

var TransformTranslateXStyleBrowserVariantsList = []string{}

func (t TransformTranslateXStyle) BrowserVariants() []string {
	return TransformTranslateXStyleBrowserVariantsList
}

var TransformTranslateXStyleUtilitiesMapping = map[string]string{
	"transform-translate-x": "transform-translate-x: translate(0)",
}

func (t TransformTranslateXStyle) Utilities() map[string]string {
	return TransformTranslateXStyleUtilitiesMapping
}

func (t TransformTranslateXStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformTranslateXStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderImageSource represent the CSS style "border-image-source" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-source
type BorderImageSourceStyle string

func (t BorderImageSourceStyle) Name() string {
	return "border-image-source"
}

var BorderImageSourceStyleBrowserVariantsList = []string{}

func (t BorderImageSourceStyle) BrowserVariants() []string {
	return BorderImageSourceStyleBrowserVariantsList
}

var BorderImageSourceStyleUtilitiesMapping = map[string]string{
	"border-image-source": "border-image-source: none",
}

func (t BorderImageSourceStyle) Utilities() map[string]string {
	return BorderImageSourceStyleUtilitiesMapping
}

func (t BorderImageSourceStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderImageSourceStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginInlineStart represent the CSS style "margin-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-inline-start
type MarginInlineStartStyle string

func (t MarginInlineStartStyle) Name() string {
	return "margin-inline-start"
}

var MarginInlineStartStyleBrowserVariantsList = []string{}

func (t MarginInlineStartStyle) BrowserVariants() []string {
	return MarginInlineStartStyleBrowserVariantsList
}

var MarginInlineStartStyleUtilitiesMapping = map[string]string{
	"margin-inline-start": "margin-inline-start: 0",
}

func (t MarginInlineStartStyle) Utilities() map[string]string {
	return MarginInlineStartStyleUtilitiesMapping
}

func (t MarginInlineStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginInlineStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBlockColor represent the CSS style "border-block-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-color
type BorderBlockColorStyle Color

func (t BorderBlockColorStyle) Name() string {
	return "border-block-color"
}

var BorderBlockColorStyleBrowserVariantsList = []string{}

func (t BorderBlockColorStyle) BrowserVariants() []string {
	return BorderBlockColorStyleBrowserVariantsList
}

var BorderBlockColorStyleUtilitiesMapping = map[string]string{
	"border-block-color": "border-block-color: currentcolor",
}

func (t BorderBlockColorStyle) Utilities() map[string]string {
	return BorderBlockColorStyleUtilitiesMapping
}

func (t BorderBlockColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBlockColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ZIndex represent the CSS style "z-index" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/z-index
type ZIndexStyle string

func (t ZIndexStyle) Name() string {
	return "z-index"
}

var ZIndexStyleBrowserVariantsList = []string{}

func (t ZIndexStyle) BrowserVariants() []string {
	return ZIndexStyleBrowserVariantsList
}

var ZIndexStyleUtilitiesMapping = map[string]string{
	"z-index": "z-index: auto",
}

func (t ZIndexStyle) Utilities() map[string]string {
	return ZIndexStyleUtilitiesMapping
}

func (t ZIndexStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ZIndexStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderEndEndRadius represent the CSS style "border-end-end-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-end-end-radius
type BorderEndEndRadiusStyle string

func (t BorderEndEndRadiusStyle) Name() string {
	return "border-end-end-radius"
}

var BorderEndEndRadiusStyleBrowserVariantsList = []string{}

func (t BorderEndEndRadiusStyle) BrowserVariants() []string {
	return BorderEndEndRadiusStyleBrowserVariantsList
}

var BorderEndEndRadiusStyleUtilitiesMapping = map[string]string{
	"border-end-end-radius": "border-end-end-radius: 0",
}

func (t BorderEndEndRadiusStyle) Utilities() map[string]string {
	return BorderEndEndRadiusStyleUtilitiesMapping
}

func (t BorderEndEndRadiusStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderEndEndRadiusStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskRepeat represent the CSS style "mask-repeat" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-repeat
type MaskRepeatStyle string

func (t MaskRepeatStyle) Name() string {
	return "mask-repeat"
}

var MaskRepeatStyleBrowserVariantsList = []string{
	"-webkit-mask-repeat",
}

func (t MaskRepeatStyle) BrowserVariants() []string {
	return MaskRepeatStyleBrowserVariantsList
}

var MaskRepeatStyleUtilitiesMapping = map[string]string{
	"mask-repeat": "mask-repeat: no-repeat",
}

func (t MaskRepeatStyle) Utilities() map[string]string {
	return MaskRepeatStyleUtilitiesMapping
}

func (t MaskRepeatStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskRepeatStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMarginInlineStart represent the CSS style "scroll-margin-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-inline-start
type ScrollMarginInlineStartStyle float64

func (t ScrollMarginInlineStartStyle) Name() string {
	return "scroll-margin-inline-start"
}

var ScrollMarginInlineStartStyleBrowserVariantsList = []string{}

func (t ScrollMarginInlineStartStyle) BrowserVariants() []string {
	return ScrollMarginInlineStartStyleBrowserVariantsList
}

var ScrollMarginInlineStartStyleUtilitiesMapping = map[string]string{
	"scroll-margin-inline-start": "scroll-margin-inline-start: 0",
}

func (t ScrollMarginInlineStartStyle) Utilities() map[string]string {
	return ScrollMarginInlineStartStyleUtilitiesMapping
}

func (t ScrollMarginInlineStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginInlineStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskBorderSource represent the CSS style "mask-border-source" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-source
type MaskBorderSourceStyle string

func (t MaskBorderSourceStyle) Name() string {
	return "mask-border-source"
}

var MaskBorderSourceStyleBrowserVariantsList = []string{}

func (t MaskBorderSourceStyle) BrowserVariants() []string {
	return MaskBorderSourceStyleBrowserVariantsList
}

var MaskBorderSourceStyleUtilitiesMapping = map[string]string{
	"mask-border-source": "mask-border-source: none",
}

func (t MaskBorderSourceStyle) Utilities() map[string]string {
	return MaskBorderSourceStyleUtilitiesMapping
}

func (t MaskBorderSourceStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskBorderSourceStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TableLayout represent the CSS style "table-layout" with value "auto | fixed"
// See https://developer.mozilla.org/docs/Web/CSS/table-layout
type TableLayoutStyle string

func (t TableLayoutStyle) Name() string {
	return "table-layout"
}

const TableLayoutStyleAuto TableLayoutStyle = "auto"

const TableLayoutStyleFixed TableLayoutStyle = "fixed"

var TableLayoutStyleBrowserVariantsList = []string{}

func (t TableLayoutStyle) BrowserVariants() []string {
	return TableLayoutStyleBrowserVariantsList
}

var TableLayoutStyleUtilitiesMapping = map[string]string{
	"table-layout":       "table-layout: auto",
	"table-layout-auto":  "table-layout: auto",
	"table-layout-fixed": "table-layout: fixed",
}

func (t TableLayoutStyle) Utilities() map[string]string {
	return TableLayoutStyleUtilitiesMapping
}

func (t TableLayoutStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TableLayoutStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridTemplateAreas represent the CSS style "grid-template-areas" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-template-areas
type GridTemplateAreasStyle string

func (t GridTemplateAreasStyle) Name() string {
	return "grid-template-areas"
}

var GridTemplateAreasStyleBrowserVariantsList = []string{}

func (t GridTemplateAreasStyle) BrowserVariants() []string {
	return GridTemplateAreasStyleBrowserVariantsList
}

var GridTemplateAreasStyleUtilitiesMapping = map[string]string{
	"grid-template-areas": "grid-template-areas: none",
}

func (t GridTemplateAreasStyle) Utilities() map[string]string {
	return GridTemplateAreasStyleUtilitiesMapping
}

func (t GridTemplateAreasStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridTemplateAreasStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollSnapTypeY represent the CSS style "scroll-snap-type-y" with value "none | mandatory | proximity"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-type-y
type ScrollSnapTypeYStyle string

func (t ScrollSnapTypeYStyle) Name() string {
	return "scroll-snap-type-y"
}

const ScrollSnapTypeYStyleNone ScrollSnapTypeYStyle = "none"

const ScrollSnapTypeYStyleMandatory ScrollSnapTypeYStyle = "mandatory"

const ScrollSnapTypeYStyleProximity ScrollSnapTypeYStyle = "proximity"

var ScrollSnapTypeYStyleBrowserVariantsList = []string{}

func (t ScrollSnapTypeYStyle) BrowserVariants() []string {
	return ScrollSnapTypeYStyleBrowserVariantsList
}

var ScrollSnapTypeYStyleUtilitiesMapping = map[string]string{
	"scroll-snap-type-y":           "scroll-snap-type-y: none",
	"scroll-snap-type-y-none":      "scroll-snap-type-y: none",
	"scroll-snap-type-y-mandatory": "scroll-snap-type-y: mandatory",
	"scroll-snap-type-y-proximity": "scroll-snap-type-y: proximity",
}

func (t ScrollSnapTypeYStyle) Utilities() map[string]string {
	return ScrollSnapTypeYStyleUtilitiesMapping
}

func (t ScrollSnapTypeYStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollSnapTypeYStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformRotate3d represent the CSS style "transform-rotate-3d" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformRotate3dStyle string

func (t TransformRotate3dStyle) Name() string {
	return "transform-rotate-3d"
}

var TransformRotate3dStyleBrowserVariantsList = []string{}

func (t TransformRotate3dStyle) BrowserVariants() []string {
	return TransformRotate3dStyleBrowserVariantsList
}

var TransformRotate3dStyleUtilitiesMapping = map[string]string{
	"transform-rotate-3d": "transform-rotate-3d: rotate3d(0, 0, 0)",
}

func (t TransformRotate3dStyle) Utilities() map[string]string {
	return TransformRotate3dStyleUtilitiesMapping
}

func (t TransformRotate3dStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformRotate3dStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarkerStart represent the CSS style "marker-start" with value ""
// See
type MarkerStartStyle string

func (t MarkerStartStyle) Name() string {
	return "marker-start"
}

var MarkerStartStyleBrowserVariantsList = []string{}

func (t MarkerStartStyle) BrowserVariants() []string {
	return MarkerStartStyleBrowserVariantsList
}

var MarkerStartStyleUtilitiesMapping = map[string]string{
	"marker-start": "marker-start: none",
}

func (t MarkerStartStyle) Utilities() map[string]string {
	return MarkerStartStyleUtilitiesMapping
}

func (t MarkerStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarkerStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Color represent the CSS style "color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/color
type ColorStyle Color

func (t ColorStyle) Name() string {
	return "color"
}

var ColorStyleBrowserVariantsList = []string{}

func (t ColorStyle) BrowserVariants() []string {
	return ColorStyleBrowserVariantsList
}

var ColorStyleUtilitiesMapping = map[string]string{
	"color": "color: variesFromBrowserToBrowser",
}

func (t ColorStyle) Utilities() map[string]string {
	return ColorStyleUtilitiesMapping
}

func (t ColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MinWidth represent the CSS style "min-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-width
type MinWidthStyle string

func (t MinWidthStyle) Name() string {
	return "min-width"
}

var MinWidthStyleBrowserVariantsList = []string{}

func (t MinWidthStyle) BrowserVariants() []string {
	return MinWidthStyleBrowserVariantsList
}

var MinWidthStyleUtilitiesMapping = map[string]string{
	"min-width": "min-width: auto",
}

func (t MinWidthStyle) Utilities() map[string]string {
	return MinWidthStyleUtilitiesMapping
}

func (t MinWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MinWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OutlineWidth represent the CSS style "outline-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-width
type OutlineWidthStyle string

func (t OutlineWidthStyle) Name() string {
	return "outline-width"
}

var OutlineWidthStyleBrowserVariantsList = []string{}

func (t OutlineWidthStyle) BrowserVariants() []string {
	return OutlineWidthStyleBrowserVariantsList
}

var OutlineWidthStyleUtilitiesMapping = map[string]string{
	"outline-width": "outline-width: medium",
}

func (t OutlineWidthStyle) Utilities() map[string]string {
	return OutlineWidthStyleUtilitiesMapping
}

func (t OutlineWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OutlineWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AnimationTimingFunction represent the CSS style "animation-timing-function" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-timing-function
type AnimationTimingFunctionStyle string

func (t AnimationTimingFunctionStyle) Name() string {
	return "animation-timing-function"
}

var AnimationTimingFunctionStyleBrowserVariantsList = []string{}

func (t AnimationTimingFunctionStyle) BrowserVariants() []string {
	return AnimationTimingFunctionStyleBrowserVariantsList
}

var AnimationTimingFunctionStyleUtilitiesMapping = map[string]string{
	"animation-timing-function": "animation-timing-function: ease",
}

func (t AnimationTimingFunctionStyle) Utilities() map[string]string {
	return AnimationTimingFunctionStyleUtilitiesMapping
}

func (t AnimationTimingFunctionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AnimationTimingFunctionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontVariationSettings represent the CSS style "font-variation-settings" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variation-settings
type FontVariationSettingsStyle string

func (t FontVariationSettingsStyle) Name() string {
	return "font-variation-settings"
}

var FontVariationSettingsStyleBrowserVariantsList = []string{}

func (t FontVariationSettingsStyle) BrowserVariants() []string {
	return FontVariationSettingsStyleBrowserVariantsList
}

var FontVariationSettingsStyleUtilitiesMapping = map[string]string{
	"font-variation-settings": "font-variation-settings: normal",
}

func (t FontVariationSettingsStyle) Utilities() map[string]string {
	return FontVariationSettingsStyleUtilitiesMapping
}

func (t FontVariationSettingsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontVariationSettingsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextDecorationSkipInk represent the CSS style "text-decoration-skip-ink" with value "auto | all | none"
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-skip-ink
type TextDecorationSkipInkStyle string

func (t TextDecorationSkipInkStyle) Name() string {
	return "text-decoration-skip-ink"
}

const TextDecorationSkipInkStyleAuto TextDecorationSkipInkStyle = "auto"

const TextDecorationSkipInkStyleAll TextDecorationSkipInkStyle = "all"

const TextDecorationSkipInkStyleNone TextDecorationSkipInkStyle = "none"

var TextDecorationSkipInkStyleBrowserVariantsList = []string{}

func (t TextDecorationSkipInkStyle) BrowserVariants() []string {
	return TextDecorationSkipInkStyleBrowserVariantsList
}

var TextDecorationSkipInkStyleUtilitiesMapping = map[string]string{
	"text-decoration-skip-ink":      "text-decoration-skip-ink: auto",
	"text-decoration-skip-ink-auto": "text-decoration-skip-ink: auto",
	"text-decoration-skip-ink-all":  "text-decoration-skip-ink: all",
	"text-decoration-skip-ink-none": "text-decoration-skip-ink: none",
}

func (t TextDecorationSkipInkStyle) Utilities() map[string]string {
	return TextDecorationSkipInkStyleUtilitiesMapping
}

func (t TextDecorationSkipInkStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextDecorationSkipInkStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextEmphasisColor represent the CSS style "text-emphasis-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-color
type TextEmphasisColorStyle Color

func (t TextEmphasisColorStyle) Name() string {
	return "text-emphasis-color"
}

var TextEmphasisColorStyleBrowserVariantsList = []string{}

func (t TextEmphasisColorStyle) BrowserVariants() []string {
	return TextEmphasisColorStyleBrowserVariantsList
}

var TextEmphasisColorStyleUtilitiesMapping = map[string]string{
	"text-emphasis-color": "text-emphasis-color: currentcolor",
}

func (t TextEmphasisColorStyle) Utilities() map[string]string {
	return TextEmphasisColorStyleUtilitiesMapping
}

func (t TextEmphasisColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextEmphasisColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AnimationFillMode represent the CSS style "animation-fill-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-fill-mode
type AnimationFillModeStyle string

func (t AnimationFillModeStyle) Name() string {
	return "animation-fill-mode"
}

var AnimationFillModeStyleBrowserVariantsList = []string{}

func (t AnimationFillModeStyle) BrowserVariants() []string {
	return AnimationFillModeStyleBrowserVariantsList
}

var AnimationFillModeStyleUtilitiesMapping = map[string]string{
	"animation-fill-mode": "animation-fill-mode: none",
}

func (t AnimationFillModeStyle) Utilities() map[string]string {
	return AnimationFillModeStyleUtilitiesMapping
}

func (t AnimationFillModeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AnimationFillModeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Contain represent the CSS style "contain" with value "none | strict | content | size  layout | style | paint "
// See https://developer.mozilla.org/docs/Web/CSS/contain
type ContainStyle string

func (t ContainStyle) Name() string {
	return "contain"
}

const ContainStyleNone ContainStyle = "none"

const ContainStyleStrict ContainStyle = "strict"

const ContainStyleContent ContainStyle = "content"

const ContainStyleSizeLayout ContainStyle = "size--layout"

const ContainStyleStyle ContainStyle = "style"

const ContainStylePaint ContainStyle = "paint"

var ContainStyleBrowserVariantsList = []string{}

func (t ContainStyle) BrowserVariants() []string {
	return ContainStyleBrowserVariantsList
}

var ContainStyleUtilitiesMapping = map[string]string{
	"contain":              "contain: none",
	"contain-none":         "contain: none",
	"contain-strict":       "contain: strict",
	"contain-content":      "contain: content",
	"contain-size--layout": "contain: size--layout",
	"contain-style":        "contain: style",
	"contain-paint":        "contain: paint",
}

func (t ContainStyle) Utilities() map[string]string {
	return ContainStyleUtilitiesMapping
}

func (t ContainStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ContainStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FlexBasis represent the CSS style "flex-basis" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/flex-basis
type FlexBasisStyle string

func (t FlexBasisStyle) Name() string {
	return "flex-basis"
}

var FlexBasisStyleBrowserVariantsList = []string{}

func (t FlexBasisStyle) BrowserVariants() []string {
	return FlexBasisStyleBrowserVariantsList
}

var FlexBasisStyleUtilitiesMapping = map[string]string{
	"flex-basis": "flex-basis: auto",
}

func (t FlexBasisStyle) Utilities() map[string]string {
	return FlexBasisStyleUtilitiesMapping
}

func (t FlexBasisStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FlexBasisStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontVariantAlternates represent the CSS style "font-variant-alternates" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-alternates
type FontVariantAlternatesStyle string

func (t FontVariantAlternatesStyle) Name() string {
	return "font-variant-alternates"
}

var FontVariantAlternatesStyleBrowserVariantsList = []string{}

func (t FontVariantAlternatesStyle) BrowserVariants() []string {
	return FontVariantAlternatesStyleBrowserVariantsList
}

var FontVariantAlternatesStyleUtilitiesMapping = map[string]string{
	"font-variant-alternates": "font-variant-alternates: normal",
}

func (t FontVariantAlternatesStyle) Utilities() map[string]string {
	return FontVariantAlternatesStyleUtilitiesMapping
}

func (t FontVariantAlternatesStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontVariantAlternatesStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridRowStart represent the CSS style "grid-row-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-row-start
type GridRowStartStyle string

func (t GridRowStartStyle) Name() string {
	return "grid-row-start"
}

var GridRowStartStyleBrowserVariantsList = []string{}

func (t GridRowStartStyle) BrowserVariants() []string {
	return GridRowStartStyleBrowserVariantsList
}

var GridRowStartStyleUtilitiesMapping = map[string]string{
	"grid-row-start": "grid-row-start: auto",
}

func (t GridRowStartStyle) Utilities() map[string]string {
	return GridRowStartStyleUtilitiesMapping
}

func (t GridRowStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridRowStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollbarColor represent the CSS style "scrollbar-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scrollbar-color
type ScrollbarColorStyle Color

func (t ScrollbarColorStyle) Name() string {
	return "scrollbar-color"
}

var ScrollbarColorStyleBrowserVariantsList = []string{}

func (t ScrollbarColorStyle) BrowserVariants() []string {
	return ScrollbarColorStyleBrowserVariantsList
}

var ScrollbarColorStyleUtilitiesMapping = map[string]string{
	"scrollbar-color": "scrollbar-color: auto",
}

func (t ScrollbarColorStyle) Utilities() map[string]string {
	return ScrollbarColorStyleUtilitiesMapping
}

func (t ScrollbarColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollbarColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// LineClamp represent the CSS style "line-clamp" with value ""
// See
type LineClampStyle string

func (t LineClampStyle) Name() string {
	return "line-clamp"
}

var LineClampStyleBrowserVariantsList = []string{
	"-webkit-line-clamp",
}

func (t LineClampStyle) BrowserVariants() []string {
	return LineClampStyleBrowserVariantsList
}

var LineClampStyleUtilitiesMapping = map[string]string{
	"line-clamp": "line-clamp: none",
}

func (t LineClampStyle) Utilities() map[string]string {
	return LineClampStyleUtilitiesMapping
}

func (t LineClampStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = LineClampStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderStartEndRadius represent the CSS style "border-start-end-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-start-end-radius
type BorderStartEndRadiusStyle string

func (t BorderStartEndRadiusStyle) Name() string {
	return "border-start-end-radius"
}

var BorderStartEndRadiusStyleBrowserVariantsList = []string{}

func (t BorderStartEndRadiusStyle) BrowserVariants() []string {
	return BorderStartEndRadiusStyleBrowserVariantsList
}

var BorderStartEndRadiusStyleUtilitiesMapping = map[string]string{
	"border-start-end-radius": "border-start-end-radius: 0",
}

func (t BorderStartEndRadiusStyle) Utilities() map[string]string {
	return BorderStartEndRadiusStyleUtilitiesMapping
}

func (t BorderStartEndRadiusStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderStartEndRadiusStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontFeatureSettings represent the CSS style "font-feature-settings" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-feature-settings
type FontFeatureSettingsStyle string

func (t FontFeatureSettingsStyle) Name() string {
	return "font-feature-settings"
}

var FontFeatureSettingsStyleBrowserVariantsList = []string{}

func (t FontFeatureSettingsStyle) BrowserVariants() []string {
	return FontFeatureSettingsStyleBrowserVariantsList
}

var FontFeatureSettingsStyleUtilitiesMapping = map[string]string{
	"font-feature-settings": "font-feature-settings: normal",
}

func (t FontFeatureSettingsStyle) Utilities() map[string]string {
	return FontFeatureSettingsStyleUtilitiesMapping
}

func (t FontFeatureSettingsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontFeatureSettingsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskBorderWidth represent the CSS style "mask-border-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-width
type MaskBorderWidthStyle string

func (t MaskBorderWidthStyle) Name() string {
	return "mask-border-width"
}

var MaskBorderWidthStyleBrowserVariantsList = []string{}

func (t MaskBorderWidthStyle) BrowserVariants() []string {
	return MaskBorderWidthStyleBrowserVariantsList
}

var MaskBorderWidthStyleUtilitiesMapping = map[string]string{
	"mask-border-width": "mask-border-width: auto",
}

func (t MaskBorderWidthStyle) Utilities() map[string]string {
	return MaskBorderWidthStyleUtilitiesMapping
}

func (t MaskBorderWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskBorderWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverscrollBehaviorBlock represent the CSS style "overscroll-behavior-block" with value "contain | none | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overscroll-behavior-block
type OverscrollBehaviorBlockStyle string

func (t OverscrollBehaviorBlockStyle) Name() string {
	return "overscroll-behavior-block"
}

const OverscrollBehaviorBlockStyleContain OverscrollBehaviorBlockStyle = "contain"

const OverscrollBehaviorBlockStyleNone OverscrollBehaviorBlockStyle = "none"

const OverscrollBehaviorBlockStyleAuto OverscrollBehaviorBlockStyle = "auto"

var OverscrollBehaviorBlockStyleBrowserVariantsList = []string{}

func (t OverscrollBehaviorBlockStyle) BrowserVariants() []string {
	return OverscrollBehaviorBlockStyleBrowserVariantsList
}

var OverscrollBehaviorBlockStyleUtilitiesMapping = map[string]string{
	"overscroll-behavior-block":         "overscroll-behavior-block: auto",
	"overscroll-behavior-block-contain": "overscroll-behavior-block: contain",
	"overscroll-behavior-block-none":    "overscroll-behavior-block: none",
	"overscroll-behavior-block-auto":    "overscroll-behavior-block: auto",
}

func (t OverscrollBehaviorBlockStyle) Utilities() map[string]string {
	return OverscrollBehaviorBlockStyleUtilitiesMapping
}

func (t OverscrollBehaviorBlockStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverscrollBehaviorBlockStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontVariantPosition represent the CSS style "font-variant-position" with value "normal | sub | super"
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-position
type FontVariantPositionStyle string

func (t FontVariantPositionStyle) Name() string {
	return "font-variant-position"
}

const FontVariantPositionStyleNormal FontVariantPositionStyle = "normal"

const FontVariantPositionStyleSub FontVariantPositionStyle = "sub"

const FontVariantPositionStyleSuper FontVariantPositionStyle = "super"

var FontVariantPositionStyleBrowserVariantsList = []string{}

func (t FontVariantPositionStyle) BrowserVariants() []string {
	return FontVariantPositionStyleBrowserVariantsList
}

var FontVariantPositionStyleUtilitiesMapping = map[string]string{
	"font-variant-position":        "font-variant-position: normal",
	"font-variant-position-normal": "font-variant-position: normal",
	"font-variant-position-sub":    "font-variant-position: sub",
	"font-variant-position-super":  "font-variant-position: super",
}

func (t FontVariantPositionStyle) Utilities() map[string]string {
	return FontVariantPositionStyleUtilitiesMapping
}

func (t FontVariantPositionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontVariantPositionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaddingBlockStart represent the CSS style "padding-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-block-start
type PaddingBlockStartStyle string

func (t PaddingBlockStartStyle) Name() string {
	return "padding-block-start"
}

var PaddingBlockStartStyleBrowserVariantsList = []string{}

func (t PaddingBlockStartStyle) BrowserVariants() []string {
	return PaddingBlockStartStyleBrowserVariantsList
}

var PaddingBlockStartStyleUtilitiesMapping = map[string]string{
	"padding-block-start": "padding-block-start: 0",
}

func (t PaddingBlockStartStyle) Utilities() map[string]string {
	return PaddingBlockStartStyleUtilitiesMapping
}

func (t PaddingBlockStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaddingBlockStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AnimationName represent the CSS style "animation-name" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-name
type AnimationNameStyle string

func (t AnimationNameStyle) Name() string {
	return "animation-name"
}

var AnimationNameStyleBrowserVariantsList = []string{}

func (t AnimationNameStyle) BrowserVariants() []string {
	return AnimationNameStyleBrowserVariantsList
}

var AnimationNameStyleUtilitiesMapping = map[string]string{
	"animation-name": "animation-name: none",
}

func (t AnimationNameStyle) Utilities() map[string]string {
	return AnimationNameStyleUtilitiesMapping
}

func (t AnimationNameStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AnimationNameStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundSize represent the CSS style "background-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-size
type BackgroundSizeStyle string

func (t BackgroundSizeStyle) Name() string {
	return "background-size"
}

var BackgroundSizeStyleBrowserVariantsList = []string{}

func (t BackgroundSizeStyle) BrowserVariants() []string {
	return BackgroundSizeStyleBrowserVariantsList
}

var BackgroundSizeStyleUtilitiesMapping = map[string]string{
	"background-size": "background-size: auto auto",
}

func (t BackgroundSizeStyle) Utilities() map[string]string {
	return BackgroundSizeStyleUtilitiesMapping
}

func (t BackgroundSizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundSizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxAlign represent the CSS style "box-align" with value "start | center | end | baseline | stretch"
// See https://developer.mozilla.org/docs/Web/CSS/box-align
type BoxAlignStyle string

func (t BoxAlignStyle) Name() string {
	return "box-align"
}

const BoxAlignStyleStart BoxAlignStyle = "start"

const BoxAlignStyleCenter BoxAlignStyle = "center"

const BoxAlignStyleEnd BoxAlignStyle = "end"

const BoxAlignStyleBaseline BoxAlignStyle = "baseline"

const BoxAlignStyleStretch BoxAlignStyle = "stretch"

var BoxAlignStyleBrowserVariantsList = []string{}

func (t BoxAlignStyle) BrowserVariants() []string {
	return BoxAlignStyleBrowserVariantsList
}

var BoxAlignStyleUtilitiesMapping = map[string]string{
	"box-align":          "box-align: stretch",
	"box-align-start":    "box-align: start",
	"box-align-center":   "box-align: center",
	"box-align-end":      "box-align: end",
	"box-align-baseline": "box-align: baseline",
	"box-align-stretch":  "box-align: stretch",
}

func (t BoxAlignStyle) Utilities() map[string]string {
	return BoxAlignStyleUtilitiesMapping
}

func (t BoxAlignStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxAlignStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxFlex represent the CSS style "box-flex" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/box-flex
type BoxFlexStyle float64

func (t BoxFlexStyle) Name() string {
	return "box-flex"
}

var BoxFlexStyleBrowserVariantsList = []string{}

func (t BoxFlexStyle) BrowserVariants() []string {
	return BoxFlexStyleBrowserVariantsList
}

var BoxFlexStyleUtilitiesMapping = map[string]string{
	"box-flex": "box-flex: 0",
}

func (t BoxFlexStyle) Utilities() map[string]string {
	return BoxFlexStyleUtilitiesMapping
}

func (t BoxFlexStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxFlexStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ImageOrientation represent the CSS style "image-orientation" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/image-orientation
type ImageOrientationStyle string

func (t ImageOrientationStyle) Name() string {
	return "image-orientation"
}

var ImageOrientationStyleBrowserVariantsList = []string{}

func (t ImageOrientationStyle) BrowserVariants() []string {
	return ImageOrientationStyleBrowserVariantsList
}

var ImageOrientationStyleUtilitiesMapping = map[string]string{
	"image-orientation": "image-orientation: from-image",
}

func (t ImageOrientationStyle) Utilities() map[string]string {
	return ImageOrientationStyleUtilitiesMapping
}

func (t ImageOrientationStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ImageOrientationStyleUtilitiesMapping[tu]
	return value, hasValue
}

// StrokeOpacity represent the CSS style "stroke-opacity" with value ""
// See
type StrokeOpacityStyle float64

func (t StrokeOpacityStyle) Name() string {
	return "stroke-opacity"
}

var StrokeOpacityStyleBrowserVariantsList = []string{}

func (t StrokeOpacityStyle) BrowserVariants() []string {
	return StrokeOpacityStyleBrowserVariantsList
}

var StrokeOpacityStyleUtilitiesMapping = map[string]string{
	"stroke-opacity": "stroke-opacity: 1",
}

func (t StrokeOpacityStyle) Utilities() map[string]string {
	return StrokeOpacityStyleUtilitiesMapping
}

func (t StrokeOpacityStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = StrokeOpacityStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxFlexGroup represent the CSS style "box-flex-group" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/box-flex-group
type BoxFlexGroupStyle float64

func (t BoxFlexGroupStyle) Name() string {
	return "box-flex-group"
}

var BoxFlexGroupStyleBrowserVariantsList = []string{}

func (t BoxFlexGroupStyle) BrowserVariants() []string {
	return BoxFlexGroupStyleBrowserVariantsList
}

var BoxFlexGroupStyleUtilitiesMapping = map[string]string{
	"box-flex-group": "box-flex-group: 1",
}

func (t BoxFlexGroupStyle) Utilities() map[string]string {
	return BoxFlexGroupStyleUtilitiesMapping
}

func (t BoxFlexGroupStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxFlexGroupStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskClip represent the CSS style "mask-clip" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-clip
type MaskClipStyle string

func (t MaskClipStyle) Name() string {
	return "mask-clip"
}

var MaskClipStyleBrowserVariantsList = []string{
	"-webkit-mask-clip",
}

func (t MaskClipStyle) BrowserVariants() []string {
	return MaskClipStyleBrowserVariantsList
}

var MaskClipStyleUtilitiesMapping = map[string]string{
	"mask-clip": "mask-clip: border-box",
}

func (t MaskClipStyle) Utilities() map[string]string {
	return MaskClipStyleUtilitiesMapping
}

func (t MaskClipStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskClipStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPaddingBlockEnd represent the CSS style "scroll-padding-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-block-end
type ScrollPaddingBlockEndStyle string

func (t ScrollPaddingBlockEndStyle) Name() string {
	return "scroll-padding-block-end"
}

var ScrollPaddingBlockEndStyleBrowserVariantsList = []string{}

func (t ScrollPaddingBlockEndStyle) BrowserVariants() []string {
	return ScrollPaddingBlockEndStyleBrowserVariantsList
}

var ScrollPaddingBlockEndStyleUtilitiesMapping = map[string]string{
	"scroll-padding-block-end": "scroll-padding-block-end: auto",
}

func (t ScrollPaddingBlockEndStyle) Utilities() map[string]string {
	return ScrollPaddingBlockEndStyleUtilitiesMapping
}

func (t ScrollPaddingBlockEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingBlockEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextShadow represent the CSS style "text-shadow" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-shadow
type TextShadowStyle string

func (t TextShadowStyle) Name() string {
	return "text-shadow"
}

var TextShadowStyleBrowserVariantsList = []string{}

func (t TextShadowStyle) BrowserVariants() []string {
	return TextShadowStyleBrowserVariantsList
}

var TextShadowStyleUtilitiesMapping = map[string]string{
	"text-shadow": "text-shadow: none",
}

func (t TextShadowStyle) Utilities() map[string]string {
	return TextShadowStyleUtilitiesMapping
}

func (t TextShadowStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextShadowStyleUtilitiesMapping[tu]
	return value, hasValue
}

// CaretColor represent the CSS style "caret-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/caret-color
type CaretColorStyle Color

func (t CaretColorStyle) Name() string {
	return "caret-color"
}

var CaretColorStyleBrowserVariantsList = []string{}

func (t CaretColorStyle) BrowserVariants() []string {
	return CaretColorStyleBrowserVariantsList
}

var CaretColorStyleUtilitiesMapping = map[string]string{
	"caret-color": "caret-color: auto",
}

func (t CaretColorStyle) Utilities() map[string]string {
	return CaretColorStyleUtilitiesMapping
}

func (t CaretColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = CaretColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskImage represent the CSS style "mask-image" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-image
type MaskImageStyle string

func (t MaskImageStyle) Name() string {
	return "mask-image"
}

var MaskImageStyleBrowserVariantsList = []string{
	"-webkit-mask-image",
}

func (t MaskImageStyle) BrowserVariants() []string {
	return MaskImageStyleBrowserVariantsList
}

var MaskImageStyleUtilitiesMapping = map[string]string{
	"mask-image": "mask-image: none",
}

func (t MaskImageStyle) Utilities() map[string]string {
	return MaskImageStyleUtilitiesMapping
}

func (t MaskImageStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskImageStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridRows represent the CSS style "grid-rows" with value "none"
// See https://developer.mozilla.org/docs/Web/CSS/-ms-grid-rows
type GridRowsStyle string

func (t GridRowsStyle) Name() string {
	return "grid-rows"
}

const GridRowsStyleNone GridRowsStyle = "none"

var GridRowsStyleBrowserVariantsList = []string{
	"-ms-grid-rows",
}

func (t GridRowsStyle) BrowserVariants() []string {
	return GridRowsStyleBrowserVariantsList
}

var GridRowsStyleUtilitiesMapping = map[string]string{
	"grid-rows":      "grid-rows: none",
	"grid-rows-none": "grid-rows: none",
}

func (t GridRowsStyle) Utilities() map[string]string {
	return GridRowsStyleUtilitiesMapping
}

func (t GridRowsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridRowsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AnimationPlayState represent the CSS style "animation-play-state" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-play-state
type AnimationPlayStateStyle string

func (t AnimationPlayStateStyle) Name() string {
	return "animation-play-state"
}

var AnimationPlayStateStyleBrowserVariantsList = []string{}

func (t AnimationPlayStateStyle) BrowserVariants() []string {
	return AnimationPlayStateStyleBrowserVariantsList
}

var AnimationPlayStateStyleUtilitiesMapping = map[string]string{
	"animation-play-state": "animation-play-state: running",
}

func (t AnimationPlayStateStyle) Utilities() map[string]string {
	return AnimationPlayStateStyleUtilitiesMapping
}

func (t AnimationPlayStateStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AnimationPlayStateStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridColumnStart represent the CSS style "grid-column-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-column-start
type GridColumnStartStyle string

func (t GridColumnStartStyle) Name() string {
	return "grid-column-start"
}

var GridColumnStartStyleBrowserVariantsList = []string{}

func (t GridColumnStartStyle) BrowserVariants() []string {
	return GridColumnStartStyleBrowserVariantsList
}

var GridColumnStartStyleUtilitiesMapping = map[string]string{
	"grid-column-start": "grid-column-start: auto",
}

func (t GridColumnStartStyle) Utilities() map[string]string {
	return GridColumnStartStyleUtilitiesMapping
}

func (t GridColumnStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridColumnStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskPosition represent the CSS style "mask-position" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-position
type MaskPositionStyle string

func (t MaskPositionStyle) Name() string {
	return "mask-position"
}

var MaskPositionStyleBrowserVariantsList = []string{
	"-webkit-mask-position",
}

func (t MaskPositionStyle) BrowserVariants() []string {
	return MaskPositionStyleBrowserVariantsList
}

var MaskPositionStyleUtilitiesMapping = map[string]string{
	"mask-position": "mask-position: center",
}

func (t MaskPositionStyle) Utilities() map[string]string {
	return MaskPositionStyleUtilitiesMapping
}

func (t MaskPositionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskPositionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Orphans represent the CSS style "orphans" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/orphans
type OrphansStyle float64

func (t OrphansStyle) Name() string {
	return "orphans"
}

var OrphansStyleBrowserVariantsList = []string{}

func (t OrphansStyle) BrowserVariants() []string {
	return OrphansStyleBrowserVariantsList
}

var OrphansStyleUtilitiesMapping = map[string]string{
	"orphans": "orphans: 2",
}

func (t OrphansStyle) Utilities() map[string]string {
	return OrphansStyleUtilitiesMapping
}

func (t OrphansStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OrphansStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollSnapAlign represent the CSS style "scroll-snap-align" with value "none | start | end | center | start end | start center | end start | end center | center start | center end"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-align
type ScrollSnapAlignStyle string

func (t ScrollSnapAlignStyle) Name() string {
	return "scroll-snap-align"
}

const ScrollSnapAlignStyleNone ScrollSnapAlignStyle = "none"

const ScrollSnapAlignStyleStart ScrollSnapAlignStyle = "start"

const ScrollSnapAlignStyleEnd ScrollSnapAlignStyle = "end"

const ScrollSnapAlignStyleCenter ScrollSnapAlignStyle = "center"

const ScrollSnapAlignStyleStartEnd ScrollSnapAlignStyle = "start-end"

const ScrollSnapAlignStyleStartCenter ScrollSnapAlignStyle = "start-center"

const ScrollSnapAlignStyleEndStart ScrollSnapAlignStyle = "end-start"

const ScrollSnapAlignStyleEndCenter ScrollSnapAlignStyle = "end-center"

const ScrollSnapAlignStyleCenterStart ScrollSnapAlignStyle = "center-start"

const ScrollSnapAlignStyleCenterEnd ScrollSnapAlignStyle = "center-end"

var ScrollSnapAlignStyleBrowserVariantsList = []string{}

func (t ScrollSnapAlignStyle) BrowserVariants() []string {
	return ScrollSnapAlignStyleBrowserVariantsList
}

var ScrollSnapAlignStyleUtilitiesMapping = map[string]string{
	"scroll-snap-align":              "scroll-snap-align: none",
	"scroll-snap-align-none":         "scroll-snap-align: none",
	"scroll-snap-align-start":        "scroll-snap-align: start",
	"scroll-snap-align-end":          "scroll-snap-align: end",
	"scroll-snap-align-center":       "scroll-snap-align: center",
	"scroll-snap-align-start-end":    "scroll-snap-align: start-end",
	"scroll-snap-align-start-center": "scroll-snap-align: start-center",
	"scroll-snap-align-end-start":    "scroll-snap-align: end-start",
	"scroll-snap-align-end-center":   "scroll-snap-align: end-center",
	"scroll-snap-align-center-start": "scroll-snap-align: center-start",
	"scroll-snap-align-center-end":   "scroll-snap-align: center-end",
}

func (t ScrollSnapAlignStyle) Utilities() map[string]string {
	return ScrollSnapAlignStyleUtilitiesMapping
}

func (t ScrollSnapAlignStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollSnapAlignStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPaddingRight represent the CSS style "scroll-padding-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-right
type ScrollPaddingRightStyle string

func (t ScrollPaddingRightStyle) Name() string {
	return "scroll-padding-right"
}

var ScrollPaddingRightStyleBrowserVariantsList = []string{}

func (t ScrollPaddingRightStyle) BrowserVariants() []string {
	return ScrollPaddingRightStyleBrowserVariantsList
}

var ScrollPaddingRightStyleUtilitiesMapping = map[string]string{
	"scroll-padding-right": "scroll-padding-right: auto",
}

func (t ScrollPaddingRightStyle) Utilities() map[string]string {
	return ScrollPaddingRightStyleUtilitiesMapping
}

func (t ScrollPaddingRightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingRightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollSnapCoordinate represent the CSS style "scroll-snap-coordinate" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-coordinate
type ScrollSnapCoordinateStyle string

func (t ScrollSnapCoordinateStyle) Name() string {
	return "scroll-snap-coordinate"
}

var ScrollSnapCoordinateStyleBrowserVariantsList = []string{}

func (t ScrollSnapCoordinateStyle) BrowserVariants() []string {
	return ScrollSnapCoordinateStyleBrowserVariantsList
}

var ScrollSnapCoordinateStyleUtilitiesMapping = map[string]string{
	"scroll-snap-coordinate": "scroll-snap-coordinate: none",
}

func (t ScrollSnapCoordinateStyle) Utilities() map[string]string {
	return ScrollSnapCoordinateStyleUtilitiesMapping
}

func (t ScrollSnapCoordinateStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollSnapCoordinateStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollSnapStop represent the CSS style "scroll-snap-stop" with value "normal | always"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-stop
type ScrollSnapStopStyle string

func (t ScrollSnapStopStyle) Name() string {
	return "scroll-snap-stop"
}

const ScrollSnapStopStyleNormal ScrollSnapStopStyle = "normal"

const ScrollSnapStopStyleAlways ScrollSnapStopStyle = "always"

var ScrollSnapStopStyleBrowserVariantsList = []string{}

func (t ScrollSnapStopStyle) BrowserVariants() []string {
	return ScrollSnapStopStyleBrowserVariantsList
}

var ScrollSnapStopStyleUtilitiesMapping = map[string]string{
	"scroll-snap-stop":        "scroll-snap-stop: normal",
	"scroll-snap-stop-normal": "scroll-snap-stop: normal",
	"scroll-snap-stop-always": "scroll-snap-stop: always",
}

func (t ScrollSnapStopStyle) Utilities() map[string]string {
	return ScrollSnapStopStyleUtilitiesMapping
}

func (t ScrollSnapStopStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollSnapStopStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ListStyleType represent the CSS style "list-style-type" with value "disc | circle | square | decimal | georgian | trad-chinese-informal | kannada"
// See https://developer.mozilla.org/docs/Web/CSS/list-style-type
type ListStyleTypeStyle string

func (t ListStyleTypeStyle) Name() string {
	return "list-style-type"
}

const ListStyleTypeStyleDisc ListStyleTypeStyle = "disc"

const ListStyleTypeStyleCircle ListStyleTypeStyle = "circle"

const ListStyleTypeStyleSquare ListStyleTypeStyle = "square"

const ListStyleTypeStyleDecimal ListStyleTypeStyle = "decimal"

const ListStyleTypeStyleGeorgian ListStyleTypeStyle = "georgian"

const ListStyleTypeStyleTradChineseInformal ListStyleTypeStyle = "trad-chinese-informal"

const ListStyleTypeStyleKannada ListStyleTypeStyle = "kannada"

var ListStyleTypeStyleBrowserVariantsList = []string{}

func (t ListStyleTypeStyle) BrowserVariants() []string {
	return ListStyleTypeStyleBrowserVariantsList
}

var ListStyleTypeStyleUtilitiesMapping = map[string]string{
	"list-style-type":                       "list-style-type: disc",
	"list-style-type-disc":                  "list-style-type: disc",
	"list-style-type-circle":                "list-style-type: circle",
	"list-style-type-square":                "list-style-type: square",
	"list-style-type-decimal":               "list-style-type: decimal",
	"list-style-type-georgian":              "list-style-type: georgian",
	"list-style-type-trad-chinese-informal": "list-style-type: trad-chinese-informal",
	"list-style-type-kannada":               "list-style-type: kannada",
}

func (t ListStyleTypeStyle) Utilities() map[string]string {
	return ListStyleTypeStyleUtilitiesMapping
}

func (t ListStyleTypeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ListStyleTypeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaddingInlineEnd represent the CSS style "padding-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-inline-end
type PaddingInlineEndStyle string

func (t PaddingInlineEndStyle) Name() string {
	return "padding-inline-end"
}

var PaddingInlineEndStyleBrowserVariantsList = []string{}

func (t PaddingInlineEndStyle) BrowserVariants() []string {
	return PaddingInlineEndStyleBrowserVariantsList
}

var PaddingInlineEndStyleUtilitiesMapping = map[string]string{
	"padding-inline-end": "padding-inline-end: 0",
}

func (t PaddingInlineEndStyle) Utilities() map[string]string {
	return PaddingInlineEndStyleUtilitiesMapping
}

func (t PaddingInlineEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaddingInlineEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// UnicodeBidi represent the CSS style "unicode-bidi" with value "normal | embed | isolate | bidi-override | isolate-override | plaintext"
// See https://developer.mozilla.org/docs/Web/CSS/unicode-bidi
type UnicodeBidiStyle string

func (t UnicodeBidiStyle) Name() string {
	return "unicode-bidi"
}

const UnicodeBidiStyleNormal UnicodeBidiStyle = "normal"

const UnicodeBidiStyleEmbed UnicodeBidiStyle = "embed"

const UnicodeBidiStyleIsolate UnicodeBidiStyle = "isolate"

const UnicodeBidiStyleBidiOverride UnicodeBidiStyle = "bidi-override"

const UnicodeBidiStyleIsolateOverride UnicodeBidiStyle = "isolate-override"

const UnicodeBidiStylePlaintext UnicodeBidiStyle = "plaintext"

var UnicodeBidiStyleBrowserVariantsList = []string{}

func (t UnicodeBidiStyle) BrowserVariants() []string {
	return UnicodeBidiStyleBrowserVariantsList
}

var UnicodeBidiStyleUtilitiesMapping = map[string]string{
	"unicode-bidi":                  "unicode-bidi: normal",
	"unicode-bidi-normal":           "unicode-bidi: normal",
	"unicode-bidi-embed":            "unicode-bidi: embed",
	"unicode-bidi-isolate":          "unicode-bidi: isolate",
	"unicode-bidi-bidi-override":    "unicode-bidi: bidi-override",
	"unicode-bidi-isolate-override": "unicode-bidi: isolate-override",
	"unicode-bidi-plaintext":        "unicode-bidi: plaintext",
}

func (t UnicodeBidiStyle) Utilities() map[string]string {
	return UnicodeBidiStyleUtilitiesMapping
}

func (t UnicodeBidiStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = UnicodeBidiStyleUtilitiesMapping[tu]
	return value, hasValue
}

// HangingPunctuation represent the CSS style "hanging-punctuation" with value "none | first  force-end | allow-end | last"
// See https://developer.mozilla.org/docs/Web/CSS/hanging-punctuation
type HangingPunctuationStyle string

func (t HangingPunctuationStyle) Name() string {
	return "hanging-punctuation"
}

const HangingPunctuationStyleNone HangingPunctuationStyle = "none"

const HangingPunctuationStyleFirstForceEnd HangingPunctuationStyle = "first--force-end"

const HangingPunctuationStyleAllowEnd HangingPunctuationStyle = "allow-end"

const HangingPunctuationStyleLast HangingPunctuationStyle = "last"

var HangingPunctuationStyleBrowserVariantsList = []string{}

func (t HangingPunctuationStyle) BrowserVariants() []string {
	return HangingPunctuationStyleBrowserVariantsList
}

var HangingPunctuationStyleUtilitiesMapping = map[string]string{
	"hanging-punctuation":                  "hanging-punctuation: none",
	"hanging-punctuation-none":             "hanging-punctuation: none",
	"hanging-punctuation-first--force-end": "hanging-punctuation: first--force-end",
	"hanging-punctuation-allow-end":        "hanging-punctuation: allow-end",
	"hanging-punctuation-last":             "hanging-punctuation: last",
}

func (t HangingPunctuationStyle) Utilities() map[string]string {
	return HangingPunctuationStyleUtilitiesMapping
}

func (t HangingPunctuationStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = HangingPunctuationStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMarginBlockStart represent the CSS style "scroll-margin-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-block-start
type ScrollMarginBlockStartStyle float64

func (t ScrollMarginBlockStartStyle) Name() string {
	return "scroll-margin-block-start"
}

var ScrollMarginBlockStartStyleBrowserVariantsList = []string{}

func (t ScrollMarginBlockStartStyle) BrowserVariants() []string {
	return ScrollMarginBlockStartStyleBrowserVariantsList
}

var ScrollMarginBlockStartStyleUtilitiesMapping = map[string]string{
	"scroll-margin-block-start": "scroll-margin-block-start: 0",
}

func (t ScrollMarginBlockStartStyle) Utilities() map[string]string {
	return ScrollMarginBlockStartStyleUtilitiesMapping
}

func (t ScrollMarginBlockStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginBlockStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// RowGap represent the CSS style "row-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/row-gap
type RowGapStyle float64

func (t RowGapStyle) Name() string {
	return "row-gap"
}

var RowGapStyleBrowserVariantsList = []string{}

func (t RowGapStyle) BrowserVariants() []string {
	return RowGapStyleBrowserVariantsList
}

var RowGapStyleUtilitiesMapping = map[string]string{
	"row-gap": "row-gap: normal",
}

func (t RowGapStyle) Utilities() map[string]string {
	return RowGapStyleUtilitiesMapping
}

func (t RowGapStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = RowGapStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BreakInside represent the CSS style "break-inside" with value "auto | avoid | avoid-page | avoid-column | avoid-region"
// See https://developer.mozilla.org/docs/Web/CSS/break-inside
type BreakInsideStyle string

func (t BreakInsideStyle) Name() string {
	return "break-inside"
}

const BreakInsideStyleAuto BreakInsideStyle = "auto"

const BreakInsideStyleAvoid BreakInsideStyle = "avoid"

const BreakInsideStyleAvoidPage BreakInsideStyle = "avoid-page"

const BreakInsideStyleAvoidColumn BreakInsideStyle = "avoid-column"

const BreakInsideStyleAvoidRegion BreakInsideStyle = "avoid-region"

var BreakInsideStyleBrowserVariantsList = []string{}

func (t BreakInsideStyle) BrowserVariants() []string {
	return BreakInsideStyleBrowserVariantsList
}

var BreakInsideStyleUtilitiesMapping = map[string]string{
	"break-inside":              "break-inside: auto",
	"break-inside-auto":         "break-inside: auto",
	"break-inside-avoid":        "break-inside: avoid",
	"break-inside-avoid-page":   "break-inside: avoid-page",
	"break-inside-avoid-column": "break-inside: avoid-column",
	"break-inside-avoid-region": "break-inside: avoid-region",
}

func (t BreakInsideStyle) Utilities() map[string]string {
	return BreakInsideStyleUtilitiesMapping
}

func (t BreakInsideStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BreakInsideStyleUtilitiesMapping[tu]
	return value, hasValue
}

// StrokeLinejoin represent the CSS style "stroke-linejoin" with value "miter | round | bevel"
// See
type StrokeLinejoinStyle string

func (t StrokeLinejoinStyle) Name() string {
	return "stroke-linejoin"
}

const StrokeLinejoinStyleMiter StrokeLinejoinStyle = "miter"

const StrokeLinejoinStyleRound StrokeLinejoinStyle = "round"

const StrokeLinejoinStyleBevel StrokeLinejoinStyle = "bevel"

var StrokeLinejoinStyleBrowserVariantsList = []string{}

func (t StrokeLinejoinStyle) BrowserVariants() []string {
	return StrokeLinejoinStyleBrowserVariantsList
}

var StrokeLinejoinStyleUtilitiesMapping = map[string]string{
	"stroke-linejoin":       "stroke-linejoin: miter",
	"stroke-linejoin-miter": "stroke-linejoin: miter",
	"stroke-linejoin-round": "stroke-linejoin: round",
	"stroke-linejoin-bevel": "stroke-linejoin: bevel",
}

func (t StrokeLinejoinStyle) Utilities() map[string]string {
	return StrokeLinejoinStyleUtilitiesMapping
}

func (t StrokeLinejoinStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = StrokeLinejoinStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundPosition represent the CSS style "background-position" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-position
type BackgroundPositionStyle string

func (t BackgroundPositionStyle) Name() string {
	return "background-position"
}

var BackgroundPositionStyleBrowserVariantsList = []string{}

func (t BackgroundPositionStyle) BrowserVariants() []string {
	return BackgroundPositionStyleBrowserVariantsList
}

var BackgroundPositionStyleUtilitiesMapping = map[string]string{
	"background-position": "background-position: 0% 0%",
}

func (t BackgroundPositionStyle) Utilities() map[string]string {
	return BackgroundPositionStyleUtilitiesMapping
}

func (t BackgroundPositionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundPositionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBottomColor represent the CSS style "border-bottom-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-color
type BorderBottomColorStyle Color

func (t BorderBottomColorStyle) Name() string {
	return "border-bottom-color"
}

var BorderBottomColorStyleBrowserVariantsList = []string{}

func (t BorderBottomColorStyle) BrowserVariants() []string {
	return BorderBottomColorStyleBrowserVariantsList
}

var BorderBottomColorStyleUtilitiesMapping = map[string]string{
	"border-bottom-color": "border-bottom-color: currentcolor",
}

func (t BorderBottomColorStyle) Utilities() map[string]string {
	return BorderBottomColorStyleUtilitiesMapping
}

func (t BorderBottomColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBottomColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPaddingInline represent the CSS style "scroll-padding-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-inline
type ScrollPaddingInlineStyle string

func (t ScrollPaddingInlineStyle) Name() string {
	return "scroll-padding-inline"
}

var ScrollPaddingInlineStyleBrowserVariantsList = []string{}

func (t ScrollPaddingInlineStyle) BrowserVariants() []string {
	return ScrollPaddingInlineStyleBrowserVariantsList
}

var ScrollPaddingInlineStyleUtilitiesMapping = map[string]string{
	"scroll-padding-inline": "scroll-padding-inline: auto",
}

func (t ScrollPaddingInlineStyle) Utilities() map[string]string {
	return ScrollPaddingInlineStyleUtilitiesMapping
}

func (t ScrollPaddingInlineStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingInlineStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Isolation represent the CSS style "isolation" with value "auto | isolate"
// See https://developer.mozilla.org/docs/Web/CSS/isolation
type IsolationStyle string

func (t IsolationStyle) Name() string {
	return "isolation"
}

const IsolationStyleAuto IsolationStyle = "auto"

const IsolationStyleIsolate IsolationStyle = "isolate"

var IsolationStyleBrowserVariantsList = []string{}

func (t IsolationStyle) BrowserVariants() []string {
	return IsolationStyleBrowserVariantsList
}

var IsolationStyleUtilitiesMapping = map[string]string{
	"isolation":         "isolation: auto",
	"isolation-auto":    "isolation: auto",
	"isolation-isolate": "isolation: isolate",
}

func (t IsolationStyle) Utilities() map[string]string {
	return IsolationStyleUtilitiesMapping
}

func (t IsolationStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = IsolationStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontVariantCaps represent the CSS style "font-variant-caps" with value "normal | small-caps | all-small-caps | petite-caps | all-petite-caps | unicase | titling-caps"
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-caps
type FontVariantCapsStyle string

func (t FontVariantCapsStyle) Name() string {
	return "font-variant-caps"
}

const FontVariantCapsStyleNormal FontVariantCapsStyle = "normal"

const FontVariantCapsStyleSmallCaps FontVariantCapsStyle = "small-caps"

const FontVariantCapsStyleAllSmallCaps FontVariantCapsStyle = "all-small-caps"

const FontVariantCapsStylePetiteCaps FontVariantCapsStyle = "petite-caps"

const FontVariantCapsStyleAllPetiteCaps FontVariantCapsStyle = "all-petite-caps"

const FontVariantCapsStyleUnicase FontVariantCapsStyle = "unicase"

const FontVariantCapsStyleTitlingCaps FontVariantCapsStyle = "titling-caps"

var FontVariantCapsStyleBrowserVariantsList = []string{}

func (t FontVariantCapsStyle) BrowserVariants() []string {
	return FontVariantCapsStyleBrowserVariantsList
}

var FontVariantCapsStyleUtilitiesMapping = map[string]string{
	"font-variant-caps":                 "font-variant-caps: normal",
	"font-variant-caps-normal":          "font-variant-caps: normal",
	"font-variant-caps-small-caps":      "font-variant-caps: small-caps",
	"font-variant-caps-all-small-caps":  "font-variant-caps: all-small-caps",
	"font-variant-caps-petite-caps":     "font-variant-caps: petite-caps",
	"font-variant-caps-all-petite-caps": "font-variant-caps: all-petite-caps",
	"font-variant-caps-unicase":         "font-variant-caps: unicase",
	"font-variant-caps-titling-caps":    "font-variant-caps: titling-caps",
}

func (t FontVariantCapsStyle) Utilities() map[string]string {
	return FontVariantCapsStyleUtilitiesMapping
}

func (t FontVariantCapsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontVariantCapsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontStyle represent the CSS style "font-style" with value "normal | italic | oblique | angle"
// See https://developer.mozilla.org/docs/Web/CSS/font-style
type FontStyleStyle string

func (t FontStyleStyle) Name() string {
	return "font-style"
}

const FontStyleStyleNormal FontStyleStyle = "normal"

const FontStyleStyleItalic FontStyleStyle = "italic"

const FontStyleStyleOblique FontStyleStyle = "oblique"

const FontStyleStyleAngle FontStyleStyle = "angle"

var FontStyleStyleBrowserVariantsList = []string{}

func (t FontStyleStyle) BrowserVariants() []string {
	return FontStyleStyleBrowserVariantsList
}

var FontStyleStyleUtilitiesMapping = map[string]string{
	"font-style":         "font-style: normal",
	"font-style-normal":  "font-style: normal",
	"font-style-italic":  "font-style: italic",
	"font-style-oblique": "font-style: oblique",
	"font-style-angle":   "font-style: angle",
}

func (t FontStyleStyle) Utilities() map[string]string {
	return FontStyleStyleUtilitiesMapping
}

func (t FontStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// JustifyContent represent the CSS style "justify-content" with value "auto | normal | stretch | end | start | flex-start | flex-end | center | left | right |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe center | unsafe center"
// See https://developer.mozilla.org/docs/Web/CSS/justify-content
type JustifyContentStyle string

func (t JustifyContentStyle) Name() string {
	return "justify-content"
}

const JustifyContentStyleAuto JustifyContentStyle = "auto"

const JustifyContentStyleNormal JustifyContentStyle = "normal"

const JustifyContentStyleStretch JustifyContentStyle = "stretch"

const JustifyContentStyleEnd JustifyContentStyle = "end"

const JustifyContentStyleStart JustifyContentStyle = "start"

const JustifyContentStyleFlexStart JustifyContentStyle = "flex-start"

const JustifyContentStyleFlexEnd JustifyContentStyle = "flex-end"

const JustifyContentStyleCenter JustifyContentStyle = "center"

const JustifyContentStyleLeft JustifyContentStyle = "left"

const JustifyContentStyleRight JustifyContentStyle = "right"

const JustifyContentStyleBaseline JustifyContentStyle = "baseline"

const JustifyContentStyleFirstBaseline JustifyContentStyle = "first-baseline"

const JustifyContentStyleLastBaseline JustifyContentStyle = "last-baseline"

const JustifyContentStyleSpaceBetween JustifyContentStyle = "space-between"

const JustifyContentStyleSpaceAround JustifyContentStyle = "space-around"

const JustifyContentStyleSpaceEvenly JustifyContentStyle = "space-evenly"

const JustifyContentStyleSafeCenter JustifyContentStyle = "safe-center"

const JustifyContentStyleUnsafeCenter JustifyContentStyle = "unsafe-center"

var JustifyContentStyleBrowserVariantsList = []string{}

func (t JustifyContentStyle) BrowserVariants() []string {
	return JustifyContentStyleBrowserVariantsList
}

var JustifyContentStyleUtilitiesMapping = map[string]string{
	"justify-content":                "justify-content: normal",
	"justify-content-auto":           "justify-content: auto",
	"justify-content-normal":         "justify-content: normal",
	"justify-content-stretch":        "justify-content: stretch",
	"justify-content-end":            "justify-content: end",
	"justify-content-start":          "justify-content: start",
	"justify-content-flex-start":     "justify-content: flex-start",
	"justify-content-flex-end":       "justify-content: flex-end",
	"justify-content-center":         "justify-content: center",
	"justify-content-left":           "justify-content: left",
	"justify-content-right":          "justify-content: right",
	"justify-content-baseline":       "justify-content: baseline",
	"justify-content-first-baseline": "justify-content: first-baseline",
	"justify-content-last-baseline":  "justify-content: last-baseline",
	"justify-content-space-between":  "justify-content: space-between",
	"justify-content-space-around":   "justify-content: space-around",
	"justify-content-space-evenly":   "justify-content: space-evenly",
	"justify-content-safe-center":    "justify-content: safe-center",
	"justify-content-unsafe-center":  "justify-content: unsafe-center",
}

func (t JustifyContentStyle) Utilities() map[string]string {
	return JustifyContentStyleUtilitiesMapping
}

func (t JustifyContentStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = JustifyContentStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginLeft represent the CSS style "margin-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-left
type MarginLeftStyle string

func (t MarginLeftStyle) Name() string {
	return "margin-left"
}

var MarginLeftStyleBrowserVariantsList = []string{}

func (t MarginLeftStyle) BrowserVariants() []string {
	return MarginLeftStyleBrowserVariantsList
}

var MarginLeftStyleUtilitiesMapping = map[string]string{
	"margin-left": "margin-left: 0",
}

func (t MarginLeftStyle) Utilities() map[string]string {
	return MarginLeftStyleUtilitiesMapping
}

func (t MarginLeftStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginLeftStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPaddingBottom represent the CSS style "scroll-padding-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-bottom
type ScrollPaddingBottomStyle string

func (t ScrollPaddingBottomStyle) Name() string {
	return "scroll-padding-bottom"
}

var ScrollPaddingBottomStyleBrowserVariantsList = []string{}

func (t ScrollPaddingBottomStyle) BrowserVariants() []string {
	return ScrollPaddingBottomStyleBrowserVariantsList
}

var ScrollPaddingBottomStyleUtilitiesMapping = map[string]string{
	"scroll-padding-bottom": "scroll-padding-bottom: auto",
}

func (t ScrollPaddingBottomStyle) Utilities() map[string]string {
	return ScrollPaddingBottomStyleUtilitiesMapping
}

func (t ScrollPaddingBottomStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingBottomStyleUtilitiesMapping[tu]
	return value, hasValue
}

// StrokeLinecap represent the CSS style "stroke-linecap" with value "butt | round | square"
// See
type StrokeLinecapStyle string

func (t StrokeLinecapStyle) Name() string {
	return "stroke-linecap"
}

const StrokeLinecapStyleButt StrokeLinecapStyle = "butt"

const StrokeLinecapStyleRound StrokeLinecapStyle = "round"

const StrokeLinecapStyleSquare StrokeLinecapStyle = "square"

var StrokeLinecapStyleBrowserVariantsList = []string{}

func (t StrokeLinecapStyle) BrowserVariants() []string {
	return StrokeLinecapStyleBrowserVariantsList
}

var StrokeLinecapStyleUtilitiesMapping = map[string]string{
	"stroke-linecap":        "stroke-linecap: butt",
	"stroke-linecap-butt":   "stroke-linecap: butt",
	"stroke-linecap-round":  "stroke-linecap: round",
	"stroke-linecap-square": "stroke-linecap: square",
}

func (t StrokeLinecapStyle) Utilities() map[string]string {
	return StrokeLinecapStyleUtilitiesMapping
}

func (t StrokeLinecapStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = StrokeLinecapStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ListStylePosition represent the CSS style "list-style-position" with value "inside | outside"
// See https://developer.mozilla.org/docs/Web/CSS/list-style-position
type ListStylePositionStyle string

func (t ListStylePositionStyle) Name() string {
	return "list-style-position"
}

const ListStylePositionStyleInside ListStylePositionStyle = "inside"

const ListStylePositionStyleOutside ListStylePositionStyle = "outside"

var ListStylePositionStyleBrowserVariantsList = []string{}

func (t ListStylePositionStyle) BrowserVariants() []string {
	return ListStylePositionStyleBrowserVariantsList
}

var ListStylePositionStyleUtilitiesMapping = map[string]string{
	"list-style-position":         "list-style-position: outside",
	"list-style-position-inside":  "list-style-position: inside",
	"list-style-position-outside": "list-style-position: outside",
}

func (t ListStylePositionStyle) Utilities() map[string]string {
	return ListStylePositionStyleUtilitiesMapping
}

func (t ListStylePositionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ListStylePositionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaxBlockSize represent the CSS style "max-block-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-block-size
type MaxBlockSizeStyle string

func (t MaxBlockSizeStyle) Name() string {
	return "max-block-size"
}

var MaxBlockSizeStyleBrowserVariantsList = []string{}

func (t MaxBlockSizeStyle) BrowserVariants() []string {
	return MaxBlockSizeStyleBrowserVariantsList
}

var MaxBlockSizeStyleUtilitiesMapping = map[string]string{
	"max-block-size": "max-block-size: 0",
}

func (t MaxBlockSizeStyle) Utilities() map[string]string {
	return MaxBlockSizeStyleUtilitiesMapping
}

func (t MaxBlockSizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaxBlockSizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBlockStyle represent the CSS style "border-block-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-style
type BorderBlockStyleStyle string

func (t BorderBlockStyleStyle) Name() string {
	return "border-block-style"
}

var BorderBlockStyleStyleBrowserVariantsList = []string{}

func (t BorderBlockStyleStyle) BrowserVariants() []string {
	return BorderBlockStyleStyleBrowserVariantsList
}

var BorderBlockStyleStyleUtilitiesMapping = map[string]string{
	"border-block-style": "border-block-style: none",
}

func (t BorderBlockStyleStyle) Utilities() map[string]string {
	return BorderBlockStyleStyleUtilitiesMapping
}

func (t BorderBlockStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBlockStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBlockStartWidth represent the CSS style "border-block-start-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-start-width
type BorderBlockStartWidthStyle string

func (t BorderBlockStartWidthStyle) Name() string {
	return "border-block-start-width"
}

var BorderBlockStartWidthStyleBrowserVariantsList = []string{}

func (t BorderBlockStartWidthStyle) BrowserVariants() []string {
	return BorderBlockStartWidthStyleBrowserVariantsList
}

var BorderBlockStartWidthStyleUtilitiesMapping = map[string]string{
	"border-block-start-width": "border-block-start-width: medium",
}

func (t BorderBlockStartWidthStyle) Utilities() map[string]string {
	return BorderBlockStartWidthStyleUtilitiesMapping
}

func (t BorderBlockStartWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBlockStartWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxLines represent the CSS style "box-lines" with value "single | multiple"
// See https://developer.mozilla.org/docs/Web/CSS/box-lines
type BoxLinesStyle string

func (t BoxLinesStyle) Name() string {
	return "box-lines"
}

const BoxLinesStyleSingle BoxLinesStyle = "single"

const BoxLinesStyleMultiple BoxLinesStyle = "multiple"

var BoxLinesStyleBrowserVariantsList = []string{}

func (t BoxLinesStyle) BrowserVariants() []string {
	return BoxLinesStyleBrowserVariantsList
}

var BoxLinesStyleUtilitiesMapping = map[string]string{
	"box-lines":          "box-lines: single",
	"box-lines-single":   "box-lines: single",
	"box-lines-multiple": "box-lines: multiple",
}

func (t BoxLinesStyle) Utilities() map[string]string {
	return BoxLinesStyleUtilitiesMapping
}

func (t BoxLinesStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxLinesStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxOrient represent the CSS style "box-orient" with value "horizontal | vertical | inline-axis | block-axis | inherit"
// See https://developer.mozilla.org/docs/Web/CSS/box-orient
type BoxOrientStyle string

func (t BoxOrientStyle) Name() string {
	return "box-orient"
}

const BoxOrientStyleHorizontal BoxOrientStyle = "horizontal"

const BoxOrientStyleVertical BoxOrientStyle = "vertical"

const BoxOrientStyleInlineAxis BoxOrientStyle = "inline-axis"

const BoxOrientStyleBlockAxis BoxOrientStyle = "block-axis"

const BoxOrientStyleInherit BoxOrientStyle = "inherit"

var BoxOrientStyleBrowserVariantsList = []string{}

func (t BoxOrientStyle) BrowserVariants() []string {
	return BoxOrientStyleBrowserVariantsList
}

var BoxOrientStyleUtilitiesMapping = map[string]string{
	"box-orient":             "box-orient: inlineAxisHorizontalInXUL",
	"box-orient-horizontal":  "box-orient: horizontal",
	"box-orient-vertical":    "box-orient: vertical",
	"box-orient-inline-axis": "box-orient: inline-axis",
	"box-orient-block-axis":  "box-orient: block-axis",
	"box-orient-inherit":     "box-orient: inherit",
}

func (t BoxOrientStyle) Utilities() map[string]string {
	return BoxOrientStyleUtilitiesMapping
}

func (t BoxOrientStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxOrientStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ImageResolution represent the CSS style "image-resolution" with value ""
// See
type ImageResolutionStyle string

func (t ImageResolutionStyle) Name() string {
	return "image-resolution"
}

var ImageResolutionStyleBrowserVariantsList = []string{}

func (t ImageResolutionStyle) BrowserVariants() []string {
	return ImageResolutionStyleBrowserVariantsList
}

var ImageResolutionStyleUtilitiesMapping = map[string]string{
	"image-resolution": "image-resolution: 1dppx",
}

func (t ImageResolutionStyle) Utilities() map[string]string {
	return ImageResolutionStyleUtilitiesMapping
}

func (t ImageResolutionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ImageResolutionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPaddingInlineStart represent the CSS style "scroll-padding-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-inline-start
type ScrollPaddingInlineStartStyle string

func (t ScrollPaddingInlineStartStyle) Name() string {
	return "scroll-padding-inline-start"
}

var ScrollPaddingInlineStartStyleBrowserVariantsList = []string{}

func (t ScrollPaddingInlineStartStyle) BrowserVariants() []string {
	return ScrollPaddingInlineStartStyleBrowserVariantsList
}

var ScrollPaddingInlineStartStyleUtilitiesMapping = map[string]string{
	"scroll-padding-inline-start": "scroll-padding-inline-start: auto",
}

func (t ScrollPaddingInlineStartStyle) Utilities() map[string]string {
	return ScrollPaddingInlineStartStyleUtilitiesMapping
}

func (t ScrollPaddingInlineStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingInlineStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AnimationIterationCount represent the CSS style "animation-iteration-count" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-iteration-count
type AnimationIterationCountStyle float64

func (t AnimationIterationCountStyle) Name() string {
	return "animation-iteration-count"
}

var AnimationIterationCountStyleBrowserVariantsList = []string{}

func (t AnimationIterationCountStyle) BrowserVariants() []string {
	return AnimationIterationCountStyleBrowserVariantsList
}

var AnimationIterationCountStyleUtilitiesMapping = map[string]string{
	"animation-iteration-count": "animation-iteration-count: 1",
}

func (t AnimationIterationCountStyle) Utilities() map[string]string {
	return AnimationIterationCountStyleUtilitiesMapping
}

func (t AnimationIterationCountStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AnimationIterationCountStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderTopLeftRadius represent the CSS style "border-top-left-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-left-radius
type BorderTopLeftRadiusStyle string

func (t BorderTopLeftRadiusStyle) Name() string {
	return "border-top-left-radius"
}

var BorderTopLeftRadiusStyleBrowserVariantsList = []string{}

func (t BorderTopLeftRadiusStyle) BrowserVariants() []string {
	return BorderTopLeftRadiusStyleBrowserVariantsList
}

var BorderTopLeftRadiusStyleUtilitiesMapping = map[string]string{
	"border-top-left-radius": "border-top-left-radius: 0",
}

func (t BorderTopLeftRadiusStyle) Utilities() map[string]string {
	return BorderTopLeftRadiusStyleUtilitiesMapping
}

func (t BorderTopLeftRadiusStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderTopLeftRadiusStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverflowAnchor represent the CSS style "overflow-anchor" with value "auto | none"
// See
type OverflowAnchorStyle string

func (t OverflowAnchorStyle) Name() string {
	return "overflow-anchor"
}

const OverflowAnchorStyleAuto OverflowAnchorStyle = "auto"

const OverflowAnchorStyleNone OverflowAnchorStyle = "none"

var OverflowAnchorStyleBrowserVariantsList = []string{}

func (t OverflowAnchorStyle) BrowserVariants() []string {
	return OverflowAnchorStyleBrowserVariantsList
}

var OverflowAnchorStyleUtilitiesMapping = map[string]string{
	"overflow-anchor":      "overflow-anchor: auto",
	"overflow-anchor-auto": "overflow-anchor: auto",
	"overflow-anchor-none": "overflow-anchor: none",
}

func (t OverflowAnchorStyle) Utilities() map[string]string {
	return OverflowAnchorStyleUtilitiesMapping
}

func (t OverflowAnchorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverflowAnchorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverflowInline represent the CSS style "overflow-inline" with value "visible | hidden | clip | scroll | auto"
// See
type OverflowInlineStyle string

func (t OverflowInlineStyle) Name() string {
	return "overflow-inline"
}

const OverflowInlineStyleVisible OverflowInlineStyle = "visible"

const OverflowInlineStyleHidden OverflowInlineStyle = "hidden"

const OverflowInlineStyleClip OverflowInlineStyle = "clip"

const OverflowInlineStyleScroll OverflowInlineStyle = "scroll"

const OverflowInlineStyleAuto OverflowInlineStyle = "auto"

var OverflowInlineStyleBrowserVariantsList = []string{}

func (t OverflowInlineStyle) BrowserVariants() []string {
	return OverflowInlineStyleBrowserVariantsList
}

var OverflowInlineStyleUtilitiesMapping = map[string]string{
	"overflow-inline":         "overflow-inline: auto",
	"overflow-inline-visible": "overflow-inline: visible",
	"overflow-inline-hidden":  "overflow-inline: hidden",
	"overflow-inline-clip":    "overflow-inline: clip",
	"overflow-inline-scroll":  "overflow-inline: scroll",
	"overflow-inline-auto":    "overflow-inline: auto",
}

func (t OverflowInlineStyle) Utilities() map[string]string {
	return OverflowInlineStyleUtilitiesMapping
}

func (t OverflowInlineStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverflowInlineStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridRowGap represent the CSS style "grid-row-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/row-gap
type GridRowGapStyle string

func (t GridRowGapStyle) Name() string {
	return "grid-row-gap"
}

var GridRowGapStyleBrowserVariantsList = []string{}

func (t GridRowGapStyle) BrowserVariants() []string {
	return GridRowGapStyleBrowserVariantsList
}

var GridRowGapStyleUtilitiesMapping = map[string]string{
	"grid-row-gap": "grid-row-gap: 0",
}

func (t GridRowGapStyle) Utilities() map[string]string {
	return GridRowGapStyleUtilitiesMapping
}

func (t GridRowGapStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridRowGapStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ContentVisibility represent the CSS style "content-visibility" with value "visible | auto | hidden"
// See https://developer.mozilla.org/docs/Web/CSS/content-visibility
type ContentVisibilityStyle string

func (t ContentVisibilityStyle) Name() string {
	return "content-visibility"
}

const ContentVisibilityStyleVisible ContentVisibilityStyle = "visible"

const ContentVisibilityStyleAuto ContentVisibilityStyle = "auto"

const ContentVisibilityStyleHidden ContentVisibilityStyle = "hidden"

var ContentVisibilityStyleBrowserVariantsList = []string{}

func (t ContentVisibilityStyle) BrowserVariants() []string {
	return ContentVisibilityStyleBrowserVariantsList
}

var ContentVisibilityStyleUtilitiesMapping = map[string]string{
	"content-visibility":         "content-visibility: visible",
	"content-visibility-visible": "content-visibility: visible",
	"content-visibility-auto":    "content-visibility: auto",
	"content-visibility-hidden":  "content-visibility: hidden",
}

func (t ContentVisibilityStyle) Utilities() map[string]string {
	return ContentVisibilityStyleUtilitiesMapping
}

func (t ContentVisibilityStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ContentVisibilityStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridAutoRows represent the CSS style "grid-auto-rows" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-auto-rows
type GridAutoRowsStyle string

func (t GridAutoRowsStyle) Name() string {
	return "grid-auto-rows"
}

var GridAutoRowsStyleBrowserVariantsList = []string{}

func (t GridAutoRowsStyle) BrowserVariants() []string {
	return GridAutoRowsStyleBrowserVariantsList
}

var GridAutoRowsStyleUtilitiesMapping = map[string]string{
	"grid-auto-rows": "grid-auto-rows: auto",
}

func (t GridAutoRowsStyle) Utilities() map[string]string {
	return GridAutoRowsStyleUtilitiesMapping
}

func (t GridAutoRowsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridAutoRowsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridTemplateColumns represent the CSS style "grid-template-columns" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-template-columns
type GridTemplateColumnsStyle string

func (t GridTemplateColumnsStyle) Name() string {
	return "grid-template-columns"
}

var GridTemplateColumnsStyleBrowserVariantsList = []string{}

func (t GridTemplateColumnsStyle) BrowserVariants() []string {
	return GridTemplateColumnsStyleBrowserVariantsList
}

var GridTemplateColumnsStyleUtilitiesMapping = map[string]string{
	"grid-template-columns": "grid-template-columns: none",
}

func (t GridTemplateColumnsStyle) Utilities() map[string]string {
	return GridTemplateColumnsStyleUtilitiesMapping
}

func (t GridTemplateColumnsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridTemplateColumnsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PageBreakInside represent the CSS style "page-break-inside" with value "auto | avoid"
// See https://developer.mozilla.org/docs/Web/CSS/page-break-inside
type PageBreakInsideStyle string

func (t PageBreakInsideStyle) Name() string {
	return "page-break-inside"
}

const PageBreakInsideStyleAuto PageBreakInsideStyle = "auto"

const PageBreakInsideStyleAvoid PageBreakInsideStyle = "avoid"

var PageBreakInsideStyleBrowserVariantsList = []string{}

func (t PageBreakInsideStyle) BrowserVariants() []string {
	return PageBreakInsideStyleBrowserVariantsList
}

var PageBreakInsideStyleUtilitiesMapping = map[string]string{
	"page-break-inside":       "page-break-inside: auto",
	"page-break-inside-auto":  "page-break-inside: auto",
	"page-break-inside-avoid": "page-break-inside: avoid",
}

func (t PageBreakInsideStyle) Utilities() map[string]string {
	return PageBreakInsideStyleUtilitiesMapping
}

func (t PageBreakInsideStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PageBreakInsideStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScaleRight represent the CSS style "scale-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scale
type ScaleRightStyle string

func (t ScaleRightStyle) Name() string {
	return "scale-right"
}

var ScaleRightStyleBrowserVariantsList = []string{}

func (t ScaleRightStyle) BrowserVariants() []string {
	return ScaleRightStyleBrowserVariantsList
}

var ScaleRightStyleUtilitiesMapping = map[string]string{
	"scale-right": "scale-right: none",
}

func (t ScaleRightStyle) Utilities() map[string]string {
	return ScaleRightStyleUtilitiesMapping
}

func (t ScaleRightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScaleRightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextDecorationColor represent the CSS style "text-decoration-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-color
type TextDecorationColorStyle Color

func (t TextDecorationColorStyle) Name() string {
	return "text-decoration-color"
}

var TextDecorationColorStyleBrowserVariantsList = []string{}

func (t TextDecorationColorStyle) BrowserVariants() []string {
	return TextDecorationColorStyleBrowserVariantsList
}

var TextDecorationColorStyleUtilitiesMapping = map[string]string{
	"text-decoration-color": "text-decoration-color: currentcolor",
}

func (t TextDecorationColorStyle) Utilities() map[string]string {
	return TextDecorationColorStyleUtilitiesMapping
}

func (t TextDecorationColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextDecorationColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBlockEndWidth represent the CSS style "border-block-end-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-end-width
type BorderBlockEndWidthStyle string

func (t BorderBlockEndWidthStyle) Name() string {
	return "border-block-end-width"
}

var BorderBlockEndWidthStyleBrowserVariantsList = []string{}

func (t BorderBlockEndWidthStyle) BrowserVariants() []string {
	return BorderBlockEndWidthStyleBrowserVariantsList
}

var BorderBlockEndWidthStyleUtilitiesMapping = map[string]string{
	"border-block-end-width": "border-block-end-width: medium",
}

func (t BorderBlockEndWidthStyle) Utilities() map[string]string {
	return BorderBlockEndWidthStyleUtilitiesMapping
}

func (t BorderBlockEndWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBlockEndWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderTopWidth represent the CSS style "border-top-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-width
type BorderTopWidthStyle string

func (t BorderTopWidthStyle) Name() string {
	return "border-top-width"
}

var BorderTopWidthStyleBrowserVariantsList = []string{}

func (t BorderTopWidthStyle) BrowserVariants() []string {
	return BorderTopWidthStyleBrowserVariantsList
}

var BorderTopWidthStyleUtilitiesMapping = map[string]string{
	"border-top-width": "border-top-width: medium",
}

func (t BorderTopWidthStyle) Utilities() map[string]string {
	return BorderTopWidthStyleUtilitiesMapping
}

func (t BorderTopWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderTopWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaxWidth represent the CSS style "max-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-width
type MaxWidthStyle string

func (t MaxWidthStyle) Name() string {
	return "max-width"
}

var MaxWidthStyleBrowserVariantsList = []string{}

func (t MaxWidthStyle) BrowserVariants() []string {
	return MaxWidthStyleBrowserVariantsList
}

var MaxWidthStyleUtilitiesMapping = map[string]string{
	"max-width": "max-width: none",
}

func (t MaxWidthStyle) Utilities() map[string]string {
	return MaxWidthStyleUtilitiesMapping
}

func (t MaxWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaxWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaxHeight represent the CSS style "max-height" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-height
type MaxHeightStyle string

func (t MaxHeightStyle) Name() string {
	return "max-height"
}

var MaxHeightStyleBrowserVariantsList = []string{}

func (t MaxHeightStyle) BrowserVariants() []string {
	return MaxHeightStyleBrowserVariantsList
}

var MaxHeightStyleUtilitiesMapping = map[string]string{
	"max-height": "max-height: none",
}

func (t MaxHeightStyle) Utilities() map[string]string {
	return MaxHeightStyleUtilitiesMapping
}

func (t MaxHeightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaxHeightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PointerEvents represent the CSS style "pointer-events" with value "auto | none | visiblePainted | visibleFill | visibleStroke | visible | painted | fill | stroke | all | inherit"
// See https://developer.mozilla.org/docs/Web/CSS/pointer-events
type PointerEventsStyle string

func (t PointerEventsStyle) Name() string {
	return "pointer-events"
}

const PointerEventsStyleAuto PointerEventsStyle = "auto"

const PointerEventsStyleNone PointerEventsStyle = "none"

const PointerEventsStyleVisiblePainted PointerEventsStyle = "visiblePainted"

const PointerEventsStyleVisibleFill PointerEventsStyle = "visibleFill"

const PointerEventsStyleVisibleStroke PointerEventsStyle = "visibleStroke"

const PointerEventsStyleVisible PointerEventsStyle = "visible"

const PointerEventsStylePainted PointerEventsStyle = "painted"

const PointerEventsStyleFill PointerEventsStyle = "fill"

const PointerEventsStyleStroke PointerEventsStyle = "stroke"

const PointerEventsStyleAll PointerEventsStyle = "all"

const PointerEventsStyleInherit PointerEventsStyle = "inherit"

var PointerEventsStyleBrowserVariantsList = []string{}

func (t PointerEventsStyle) BrowserVariants() []string {
	return PointerEventsStyleBrowserVariantsList
}

var PointerEventsStyleUtilitiesMapping = map[string]string{
	"pointer-events":                "pointer-events: auto",
	"pointer-events-auto":           "pointer-events: auto",
	"pointer-events-none":           "pointer-events: none",
	"pointer-events-visiblePainted": "pointer-events: visiblePainted",
	"pointer-events-visibleFill":    "pointer-events: visibleFill",
	"pointer-events-visibleStroke":  "pointer-events: visibleStroke",
	"pointer-events-visible":        "pointer-events: visible",
	"pointer-events-painted":        "pointer-events: painted",
	"pointer-events-fill":           "pointer-events: fill",
	"pointer-events-stroke":         "pointer-events: stroke",
	"pointer-events-all":            "pointer-events: all",
	"pointer-events-inherit":        "pointer-events: inherit",
}

func (t PointerEventsStyle) Utilities() map[string]string {
	return PointerEventsStyleUtilitiesMapping
}

func (t PointerEventsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PointerEventsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransitionDelay represent the CSS style "transition-delay" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-delay
type TransitionDelayStyle time.Duration

func (t TransitionDelayStyle) Name() string {
	return "transition-delay"
}

var TransitionDelayStyleBrowserVariantsList = []string{}

func (t TransitionDelayStyle) BrowserVariants() []string {
	return TransitionDelayStyleBrowserVariantsList
}

var TransitionDelayStyleUtilitiesMapping = map[string]string{
	"transition-delay": "transition-delay: 0s",
}

func (t TransitionDelayStyle) Utilities() map[string]string {
	return TransitionDelayStyleUtilitiesMapping
}

func (t TransitionDelayStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransitionDelayStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Cursor represent the CSS style "cursor" with value "auto | default | none | context-menu | help | pointer | progress | wait | cell | crosshair | text | vertical-text | alias | copy | move | no-drop | not-allowed | e-resize | n-resize | ne-resize | nw-resize | s-resize | se-resize | sw-resize | w-resize | ew-resize | ns-resize | nesw-resize | nwse-resize | col-resize | row-resize | all-scroll | zoom-in | zoom-out | grab | grabbing"
// See https://developer.mozilla.org/docs/Web/CSS/cursor
type CursorStyle string

func (t CursorStyle) Name() string {
	return "cursor"
}

const CursorStyleAuto CursorStyle = "auto"

const CursorStyleDefault CursorStyle = "default"

const CursorStyleNone CursorStyle = "none"

const CursorStyleContextMenu CursorStyle = "context-menu"

const CursorStyleHelp CursorStyle = "help"

const CursorStylePointer CursorStyle = "pointer"

const CursorStyleProgress CursorStyle = "progress"

const CursorStyleWait CursorStyle = "wait"

const CursorStyleCell CursorStyle = "cell"

const CursorStyleCrosshair CursorStyle = "crosshair"

const CursorStyleText CursorStyle = "text"

const CursorStyleVerticalText CursorStyle = "vertical-text"

const CursorStyleAlias CursorStyle = "alias"

const CursorStyleCopy CursorStyle = "copy"

const CursorStyleMove CursorStyle = "move"

const CursorStyleNoDrop CursorStyle = "no-drop"

const CursorStyleNotAllowed CursorStyle = "not-allowed"

const CursorStyleEResize CursorStyle = "e-resize"

const CursorStyleNResize CursorStyle = "n-resize"

const CursorStyleNeResize CursorStyle = "ne-resize"

const CursorStyleNwResize CursorStyle = "nw-resize"

const CursorStyleSResize CursorStyle = "s-resize"

const CursorStyleSeResize CursorStyle = "se-resize"

const CursorStyleSwResize CursorStyle = "sw-resize"

const CursorStyleWResize CursorStyle = "w-resize"

const CursorStyleEwResize CursorStyle = "ew-resize"

const CursorStyleNsResize CursorStyle = "ns-resize"

const CursorStyleNeswResize CursorStyle = "nesw-resize"

const CursorStyleNwseResize CursorStyle = "nwse-resize"

const CursorStyleColResize CursorStyle = "col-resize"

const CursorStyleRowResize CursorStyle = "row-resize"

const CursorStyleAllScroll CursorStyle = "all-scroll"

const CursorStyleZoomIn CursorStyle = "zoom-in"

const CursorStyleZoomOut CursorStyle = "zoom-out"

const CursorStyleGrab CursorStyle = "grab"

const CursorStyleGrabbing CursorStyle = "grabbing"

var CursorStyleBrowserVariantsList = []string{}

func (t CursorStyle) BrowserVariants() []string {
	return CursorStyleBrowserVariantsList
}

var CursorStyleUtilitiesMapping = map[string]string{
	"cursor":               "cursor: auto",
	"cursor-auto":          "cursor: auto",
	"cursor-default":       "cursor: default",
	"cursor-none":          "cursor: none",
	"cursor-context-menu":  "cursor: context-menu",
	"cursor-help":          "cursor: help",
	"cursor-pointer":       "cursor: pointer",
	"cursor-progress":      "cursor: progress",
	"cursor-wait":          "cursor: wait",
	"cursor-cell":          "cursor: cell",
	"cursor-crosshair":     "cursor: crosshair",
	"cursor-text":          "cursor: text",
	"cursor-vertical-text": "cursor: vertical-text",
	"cursor-alias":         "cursor: alias",
	"cursor-copy":          "cursor: copy",
	"cursor-move":          "cursor: move",
	"cursor-no-drop":       "cursor: no-drop",
	"cursor-not-allowed":   "cursor: not-allowed",
	"cursor-e-resize":      "cursor: e-resize",
	"cursor-n-resize":      "cursor: n-resize",
	"cursor-ne-resize":     "cursor: ne-resize",
	"cursor-nw-resize":     "cursor: nw-resize",
	"cursor-s-resize":      "cursor: s-resize",
	"cursor-se-resize":     "cursor: se-resize",
	"cursor-sw-resize":     "cursor: sw-resize",
	"cursor-w-resize":      "cursor: w-resize",
	"cursor-ew-resize":     "cursor: ew-resize",
	"cursor-ns-resize":     "cursor: ns-resize",
	"cursor-nesw-resize":   "cursor: nesw-resize",
	"cursor-nwse-resize":   "cursor: nwse-resize",
	"cursor-col-resize":    "cursor: col-resize",
	"cursor-row-resize":    "cursor: row-resize",
	"cursor-all-scroll":    "cursor: all-scroll",
	"cursor-zoom-in":       "cursor: zoom-in",
	"cursor-zoom-out":      "cursor: zoom-out",
	"cursor-grab":          "cursor: grab",
	"cursor-grabbing":      "cursor: grabbing",
}

func (t CursorStyle) Utilities() map[string]string {
	return CursorStyleUtilitiesMapping
}

func (t CursorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = CursorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformRotate represent the CSS style "transform-rotate" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformRotateStyle string

func (t TransformRotateStyle) Name() string {
	return "transform-rotate"
}

var TransformRotateStyleBrowserVariantsList = []string{}

func (t TransformRotateStyle) BrowserVariants() []string {
	return TransformRotateStyleBrowserVariantsList
}

var TransformRotateStyleUtilitiesMapping = map[string]string{
	"transform-rotate": "transform-rotate: rotate(0)",
}

func (t TransformRotateStyle) Utilities() map[string]string {
	return TransformRotateStyleUtilitiesMapping
}

func (t TransformRotateStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformRotateStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ShapeRendering represent the CSS style "shape-rendering" with value "auto | optimizeSpeed | crispEdges | geometricPrecision"
// See
type ShapeRenderingStyle string

func (t ShapeRenderingStyle) Name() string {
	return "shape-rendering"
}

const ShapeRenderingStyleAuto ShapeRenderingStyle = "auto"

const ShapeRenderingStyleOptimizeSpeed ShapeRenderingStyle = "optimizeSpeed"

const ShapeRenderingStyleCrispEdges ShapeRenderingStyle = "crispEdges"

const ShapeRenderingStyleGeometricPrecision ShapeRenderingStyle = "geometricPrecision"

var ShapeRenderingStyleBrowserVariantsList = []string{}

func (t ShapeRenderingStyle) BrowserVariants() []string {
	return ShapeRenderingStyleBrowserVariantsList
}

var ShapeRenderingStyleUtilitiesMapping = map[string]string{
	"shape-rendering":                    "shape-rendering: auto",
	"shape-rendering-auto":               "shape-rendering: auto",
	"shape-rendering-optimizeSpeed":      "shape-rendering: optimizeSpeed",
	"shape-rendering-crispEdges":         "shape-rendering: crispEdges",
	"shape-rendering-geometricPrecision": "shape-rendering: geometricPrecision",
}

func (t ShapeRenderingStyle) Utilities() map[string]string {
	return ShapeRenderingStyleUtilitiesMapping
}

func (t ShapeRenderingStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ShapeRenderingStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AlignSelf represent the CSS style "align-self" with value "auto | normal | stretch | end | start | flex-start | flex-end | center |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe | unsafe"
// See https://developer.mozilla.org/docs/Web/CSS/align-self
type AlignSelfStyle string

func (t AlignSelfStyle) Name() string {
	return "align-self"
}

const AlignSelfStyleAuto AlignSelfStyle = "auto"

const AlignSelfStyleNormal AlignSelfStyle = "normal"

const AlignSelfStyleStretch AlignSelfStyle = "stretch"

const AlignSelfStyleEnd AlignSelfStyle = "end"

const AlignSelfStyleStart AlignSelfStyle = "start"

const AlignSelfStyleFlexStart AlignSelfStyle = "flex-start"

const AlignSelfStyleFlexEnd AlignSelfStyle = "flex-end"

const AlignSelfStyleCenter AlignSelfStyle = "center"

const AlignSelfStyleBaseline AlignSelfStyle = "baseline"

const AlignSelfStyleFirstBaseline AlignSelfStyle = "first-baseline"

const AlignSelfStyleLastBaseline AlignSelfStyle = "last-baseline"

const AlignSelfStyleSpaceBetween AlignSelfStyle = "space-between"

const AlignSelfStyleSpaceAround AlignSelfStyle = "space-around"

const AlignSelfStyleSpaceEvenly AlignSelfStyle = "space-evenly"

const AlignSelfStyleSafe AlignSelfStyle = "safe"

const AlignSelfStyleUnsafe AlignSelfStyle = "unsafe"

var AlignSelfStyleBrowserVariantsList = []string{}

func (t AlignSelfStyle) BrowserVariants() []string {
	return AlignSelfStyleBrowserVariantsList
}

var AlignSelfStyleUtilitiesMapping = map[string]string{
	"align-self":                "align-self: auto",
	"align-self-auto":           "align-self: auto",
	"align-self-normal":         "align-self: normal",
	"align-self-stretch":        "align-self: stretch",
	"align-self-end":            "align-self: end",
	"align-self-start":          "align-self: start",
	"align-self-flex-start":     "align-self: flex-start",
	"align-self-flex-end":       "align-self: flex-end",
	"align-self-center":         "align-self: center",
	"align-self-baseline":       "align-self: baseline",
	"align-self-first-baseline": "align-self: first-baseline",
	"align-self-last-baseline":  "align-self: last-baseline",
	"align-self-space-between":  "align-self: space-between",
	"align-self-space-around":   "align-self: space-around",
	"align-self-space-evenly":   "align-self: space-evenly",
	"align-self-safe":           "align-self: safe",
	"align-self-unsafe":         "align-self: unsafe",
}

func (t AlignSelfStyle) Utilities() map[string]string {
	return AlignSelfStyleUtilitiesMapping
}

func (t AlignSelfStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AlignSelfStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderImageRepeat represent the CSS style "border-image-repeat" with value "stretch | repeat | round | space"
// See https://developer.mozilla.org/docs/Web/CSS/border-image-repeat
type BorderImageRepeatStyle string

func (t BorderImageRepeatStyle) Name() string {
	return "border-image-repeat"
}

const BorderImageRepeatStyleStretch BorderImageRepeatStyle = "stretch"

const BorderImageRepeatStyleRepeat BorderImageRepeatStyle = "repeat"

const BorderImageRepeatStyleRound BorderImageRepeatStyle = "round"

const BorderImageRepeatStyleSpace BorderImageRepeatStyle = "space"

var BorderImageRepeatStyleBrowserVariantsList = []string{}

func (t BorderImageRepeatStyle) BrowserVariants() []string {
	return BorderImageRepeatStyleBrowserVariantsList
}

var BorderImageRepeatStyleUtilitiesMapping = map[string]string{
	"border-image-repeat":         "border-image-repeat: stretch",
	"border-image-repeat-stretch": "border-image-repeat: stretch",
	"border-image-repeat-repeat":  "border-image-repeat: repeat",
	"border-image-repeat-round":   "border-image-repeat: round",
	"border-image-repeat-space":   "border-image-repeat: space",
}

func (t BorderImageRepeatStyle) Utilities() map[string]string {
	return BorderImageRepeatStyleUtilitiesMapping
}

func (t BorderImageRepeatStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderImageRepeatStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Gap represent the CSS style "gap" with value "none"
// See https://developer.mozilla.org/docs/Web/CSS/gap
type GapStyle string

func (t GapStyle) Name() string {
	return "gap"
}

const GapStyleNone GapStyle = "none"

var GapStyleBrowserVariantsList = []string{}

func (t GapStyle) BrowserVariants() []string {
	return GapStyleBrowserVariantsList
}

var GapStyleUtilitiesMapping = map[string]string{
	"gap":      "gap: none",
	"gap-none": "gap: none",
}

func (t GapStyle) Utilities() map[string]string {
	return GapStyleUtilitiesMapping
}

func (t GapStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GapStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBlockWidth represent the CSS style "border-block-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-width
type BorderBlockWidthStyle string

func (t BorderBlockWidthStyle) Name() string {
	return "border-block-width"
}

var BorderBlockWidthStyleBrowserVariantsList = []string{}

func (t BorderBlockWidthStyle) BrowserVariants() []string {
	return BorderBlockWidthStyleBrowserVariantsList
}

var BorderBlockWidthStyleUtilitiesMapping = map[string]string{
	"border-block-width": "border-block-width: medium",
}

func (t BorderBlockWidthStyle) Utilities() map[string]string {
	return BorderBlockWidthStyleUtilitiesMapping
}

func (t BorderBlockWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBlockWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColorAdjust represent the CSS style "color-adjust" with value "economy | exact"
// See https://developer.mozilla.org/docs/Web/CSS/color-adjust
type ColorAdjustStyle string

func (t ColorAdjustStyle) Name() string {
	return "color-adjust"
}

const ColorAdjustStyleEconomy ColorAdjustStyle = "economy"

const ColorAdjustStyleExact ColorAdjustStyle = "exact"

var ColorAdjustStyleBrowserVariantsList = []string{}

func (t ColorAdjustStyle) BrowserVariants() []string {
	return ColorAdjustStyleBrowserVariantsList
}

var ColorAdjustStyleUtilitiesMapping = map[string]string{
	"color-adjust":         "color-adjust: economy",
	"color-adjust-economy": "color-adjust: economy",
	"color-adjust-exact":   "color-adjust: exact",
}

func (t ColorAdjustStyle) Utilities() map[string]string {
	return ColorAdjustStyleUtilitiesMapping
}

func (t ColorAdjustStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColorAdjustStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridRowEnd represent the CSS style "grid-row-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-row-end
type GridRowEndStyle string

func (t GridRowEndStyle) Name() string {
	return "grid-row-end"
}

var GridRowEndStyleBrowserVariantsList = []string{}

func (t GridRowEndStyle) BrowserVariants() []string {
	return GridRowEndStyleBrowserVariantsList
}

var GridRowEndStyleUtilitiesMapping = map[string]string{
	"grid-row-end": "grid-row-end: auto",
}

func (t GridRowEndStyle) Utilities() map[string]string {
	return GridRowEndStyleUtilitiesMapping
}

func (t GridRowEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridRowEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// JustifyItems represent the CSS style "justify-items" with value "auto | normal | stretch | end | start | flex-start | flex-end | center | left | right |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe center | unsafe center"
// See https://developer.mozilla.org/docs/Web/CSS/justify-items
type JustifyItemsStyle string

func (t JustifyItemsStyle) Name() string {
	return "justify-items"
}

const JustifyItemsStyleAuto JustifyItemsStyle = "auto"

const JustifyItemsStyleNormal JustifyItemsStyle = "normal"

const JustifyItemsStyleStretch JustifyItemsStyle = "stretch"

const JustifyItemsStyleEnd JustifyItemsStyle = "end"

const JustifyItemsStyleStart JustifyItemsStyle = "start"

const JustifyItemsStyleFlexStart JustifyItemsStyle = "flex-start"

const JustifyItemsStyleFlexEnd JustifyItemsStyle = "flex-end"

const JustifyItemsStyleCenter JustifyItemsStyle = "center"

const JustifyItemsStyleLeft JustifyItemsStyle = "left"

const JustifyItemsStyleRight JustifyItemsStyle = "right"

const JustifyItemsStyleBaseline JustifyItemsStyle = "baseline"

const JustifyItemsStyleFirstBaseline JustifyItemsStyle = "first-baseline"

const JustifyItemsStyleLastBaseline JustifyItemsStyle = "last-baseline"

const JustifyItemsStyleSpaceBetween JustifyItemsStyle = "space-between"

const JustifyItemsStyleSpaceAround JustifyItemsStyle = "space-around"

const JustifyItemsStyleSpaceEvenly JustifyItemsStyle = "space-evenly"

const JustifyItemsStyleSafeCenter JustifyItemsStyle = "safe-center"

const JustifyItemsStyleUnsafeCenter JustifyItemsStyle = "unsafe-center"

var JustifyItemsStyleBrowserVariantsList = []string{}

func (t JustifyItemsStyle) BrowserVariants() []string {
	return JustifyItemsStyleBrowserVariantsList
}

var JustifyItemsStyleUtilitiesMapping = map[string]string{
	"justify-items":                "justify-items: legacy",
	"justify-items-auto":           "justify-items: auto",
	"justify-items-normal":         "justify-items: normal",
	"justify-items-stretch":        "justify-items: stretch",
	"justify-items-end":            "justify-items: end",
	"justify-items-start":          "justify-items: start",
	"justify-items-flex-start":     "justify-items: flex-start",
	"justify-items-flex-end":       "justify-items: flex-end",
	"justify-items-center":         "justify-items: center",
	"justify-items-left":           "justify-items: left",
	"justify-items-right":          "justify-items: right",
	"justify-items-baseline":       "justify-items: baseline",
	"justify-items-first-baseline": "justify-items: first-baseline",
	"justify-items-last-baseline":  "justify-items: last-baseline",
	"justify-items-space-between":  "justify-items: space-between",
	"justify-items-space-around":   "justify-items: space-around",
	"justify-items-space-evenly":   "justify-items: space-evenly",
	"justify-items-safe-center":    "justify-items: safe-center",
	"justify-items-unsafe-center":  "justify-items: unsafe-center",
}

func (t JustifyItemsStyle) Utilities() map[string]string {
	return JustifyItemsStyleUtilitiesMapping
}

func (t JustifyItemsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = JustifyItemsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Fill represent the CSS style "fill" with value ""
// See
type FillStyle Color

func (t FillStyle) Name() string {
	return "fill"
}

var FillStyleBrowserVariantsList = []string{}

func (t FillStyle) BrowserVariants() []string {
	return FillStyleBrowserVariantsList
}

var FillStyleUtilitiesMapping = map[string]string{
	"fill": "fill: black",
}

func (t FillStyle) Utilities() map[string]string {
	return FillStyleUtilitiesMapping
}

func (t FillStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FillStyleUtilitiesMapping[tu]
	return value, hasValue
}

// LightingColor represent the CSS style "lighting-color" with value ""
// See
type LightingColorStyle Color

func (t LightingColorStyle) Name() string {
	return "lighting-color"
}

var LightingColorStyleBrowserVariantsList = []string{}

func (t LightingColorStyle) BrowserVariants() []string {
	return LightingColorStyleBrowserVariantsList
}

var LightingColorStyleUtilitiesMapping = map[string]string{
	"lighting-color": "lighting-color: white",
}

func (t LightingColorStyle) Utilities() map[string]string {
	return LightingColorStyleUtilitiesMapping
}

func (t LightingColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = LightingColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundImage represent the CSS style "background-image" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-image
type BackgroundImageStyle URL

func (t BackgroundImageStyle) Name() string {
	return "background-image"
}

var BackgroundImageStyleBrowserVariantsList = []string{}

func (t BackgroundImageStyle) BrowserVariants() []string {
	return BackgroundImageStyleBrowserVariantsList
}

var BackgroundImageStyleUtilitiesMapping = map[string]string{
	"background-image": "background-image: none",
}

func (t BackgroundImageStyle) Utilities() map[string]string {
	return BackgroundImageStyleUtilitiesMapping
}

func (t BackgroundImageStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundImageStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverscrollBehaviorX represent the CSS style "overscroll-behavior-x" with value "contain | none | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overscroll-behavior-x
type OverscrollBehaviorXStyle string

func (t OverscrollBehaviorXStyle) Name() string {
	return "overscroll-behavior-x"
}

const OverscrollBehaviorXStyleContain OverscrollBehaviorXStyle = "contain"

const OverscrollBehaviorXStyleNone OverscrollBehaviorXStyle = "none"

const OverscrollBehaviorXStyleAuto OverscrollBehaviorXStyle = "auto"

var OverscrollBehaviorXStyleBrowserVariantsList = []string{}

func (t OverscrollBehaviorXStyle) BrowserVariants() []string {
	return OverscrollBehaviorXStyleBrowserVariantsList
}

var OverscrollBehaviorXStyleUtilitiesMapping = map[string]string{
	"overscroll-behavior-x":         "overscroll-behavior-x: auto",
	"overscroll-behavior-x-contain": "overscroll-behavior-x: contain",
	"overscroll-behavior-x-none":    "overscroll-behavior-x: none",
	"overscroll-behavior-x-auto":    "overscroll-behavior-x: auto",
}

func (t OverscrollBehaviorXStyle) Utilities() map[string]string {
	return OverscrollBehaviorXStyleUtilitiesMapping
}

func (t OverscrollBehaviorXStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverscrollBehaviorXStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextSizeAdjust represent the CSS style "text-size-adjust" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-size-adjust
type TextSizeAdjustStyle string

func (t TextSizeAdjustStyle) Name() string {
	return "text-size-adjust"
}

var TextSizeAdjustStyleBrowserVariantsList = []string{}

func (t TextSizeAdjustStyle) BrowserVariants() []string {
	return TextSizeAdjustStyleBrowserVariantsList
}

var TextSizeAdjustStyleUtilitiesMapping = map[string]string{
	"text-size-adjust": "text-size-adjust: autoForSmartphoneBrowsersSupportingInflation",
}

func (t TextSizeAdjustStyle) Utilities() map[string]string {
	return TextSizeAdjustStyleUtilitiesMapping
}

func (t TextSizeAdjustStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextSizeAdjustStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Float represent the CSS style "float" with value "left | right | none | inline-start | inline-end"
// See https://developer.mozilla.org/docs/Web/CSS/float
type FloatStyle string

func (t FloatStyle) Name() string {
	return "float"
}

const FloatStyleLeft FloatStyle = "left"

const FloatStyleRight FloatStyle = "right"

const FloatStyleNone FloatStyle = "none"

const FloatStyleInlineStart FloatStyle = "inline-start"

const FloatStyleInlineEnd FloatStyle = "inline-end"

var FloatStyleBrowserVariantsList = []string{}

func (t FloatStyle) BrowserVariants() []string {
	return FloatStyleBrowserVariantsList
}

var FloatStyleUtilitiesMapping = map[string]string{
	"float":              "float: none",
	"float-left":         "float: left",
	"float-right":        "float: right",
	"float-none":         "float: none",
	"float-inline-start": "float: inline-start",
	"float-inline-end":   "float: inline-end",
}

func (t FloatStyle) Utilities() map[string]string {
	return FloatStyleUtilitiesMapping
}

func (t FloatStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FloatStyleUtilitiesMapping[tu]
	return value, hasValue
}

// StopColor represent the CSS style "stop-color" with value ""
// See
type StopColorStyle Color

func (t StopColorStyle) Name() string {
	return "stop-color"
}

var StopColorStyleBrowserVariantsList = []string{}

func (t StopColorStyle) BrowserVariants() []string {
	return StopColorStyleBrowserVariantsList
}

var StopColorStyleUtilitiesMapping = map[string]string{
	"stop-color": "stop-color: black",
}

func (t StopColorStyle) Utilities() map[string]string {
	return StopColorStyleUtilitiesMapping
}

func (t StopColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = StopColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColumnRuleColor represent the CSS style "column-rule-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-rule-color
type ColumnRuleColorStyle Color

func (t ColumnRuleColorStyle) Name() string {
	return "column-rule-color"
}

var ColumnRuleColorStyleBrowserVariantsList = []string{}

func (t ColumnRuleColorStyle) BrowserVariants() []string {
	return ColumnRuleColorStyleBrowserVariantsList
}

var ColumnRuleColorStyleUtilitiesMapping = map[string]string{
	"column-rule-color": "column-rule-color: currentcolor",
}

func (t ColumnRuleColorStyle) Utilities() map[string]string {
	return ColumnRuleColorStyleUtilitiesMapping
}

func (t ColumnRuleColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColumnRuleColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Overflow represent the CSS style "overflow" with value "visible | hidden | clip | scroll | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overflow
type OverflowStyle string

func (t OverflowStyle) Name() string {
	return "overflow"
}

const OverflowStyleVisible OverflowStyle = "visible"

const OverflowStyleHidden OverflowStyle = "hidden"

const OverflowStyleClip OverflowStyle = "clip"

const OverflowStyleScroll OverflowStyle = "scroll"

const OverflowStyleAuto OverflowStyle = "auto"

var OverflowStyleBrowserVariantsList = []string{}

func (t OverflowStyle) BrowserVariants() []string {
	return OverflowStyleBrowserVariantsList
}

var OverflowStyleUtilitiesMapping = map[string]string{
	"overflow":         "overflow: visible",
	"overflow-visible": "overflow: visible",
	"overflow-hidden":  "overflow: hidden",
	"overflow-clip":    "overflow: clip",
	"overflow-scroll":  "overflow: scroll",
	"overflow-auto":    "overflow: auto",
}

func (t OverflowStyle) Utilities() map[string]string {
	return OverflowStyleUtilitiesMapping
}

func (t OverflowStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverflowStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaddingBlockEnd represent the CSS style "padding-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-block-end
type PaddingBlockEndStyle string

func (t PaddingBlockEndStyle) Name() string {
	return "padding-block-end"
}

var PaddingBlockEndStyleBrowserVariantsList = []string{}

func (t PaddingBlockEndStyle) BrowserVariants() []string {
	return PaddingBlockEndStyleBrowserVariantsList
}

var PaddingBlockEndStyleUtilitiesMapping = map[string]string{
	"padding-block-end": "padding-block-end: 0",
}

func (t PaddingBlockEndStyle) Utilities() map[string]string {
	return PaddingBlockEndStyleUtilitiesMapping
}

func (t PaddingBlockEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaddingBlockEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderInlineEndColor represent the CSS style "border-inline-end-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-end-color
type BorderInlineEndColorStyle Color

func (t BorderInlineEndColorStyle) Name() string {
	return "border-inline-end-color"
}

var BorderInlineEndColorStyleBrowserVariantsList = []string{}

func (t BorderInlineEndColorStyle) BrowserVariants() []string {
	return BorderInlineEndColorStyleBrowserVariantsList
}

var BorderInlineEndColorStyleUtilitiesMapping = map[string]string{
	"border-inline-end-color": "border-inline-end-color: currentcolor",
}

func (t BorderInlineEndColorStyle) Utilities() map[string]string {
	return BorderInlineEndColorStyleUtilitiesMapping
}

func (t BorderInlineEndColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderInlineEndColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// LineHeight represent the CSS style "line-height" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/line-height
type LineHeightStyle string

func (t LineHeightStyle) Name() string {
	return "line-height"
}

var LineHeightStyleBrowserVariantsList = []string{}

func (t LineHeightStyle) BrowserVariants() []string {
	return LineHeightStyleBrowserVariantsList
}

var LineHeightStyleUtilitiesMapping = map[string]string{
	"line-height": "line-height: normal",
}

func (t LineHeightStyle) Utilities() map[string]string {
	return LineHeightStyleUtilitiesMapping
}

func (t LineHeightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = LineHeightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MinInlineSize represent the CSS style "min-inline-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-inline-size
type MinInlineSizeStyle string

func (t MinInlineSizeStyle) Name() string {
	return "min-inline-size"
}

var MinInlineSizeStyleBrowserVariantsList = []string{}

func (t MinInlineSizeStyle) BrowserVariants() []string {
	return MinInlineSizeStyleBrowserVariantsList
}

var MinInlineSizeStyleUtilitiesMapping = map[string]string{
	"min-inline-size": "min-inline-size: 0",
}

func (t MinInlineSizeStyle) Utilities() map[string]string {
	return MinInlineSizeStyleUtilitiesMapping
}

func (t MinInlineSizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MinInlineSizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPaddingInlineEnd represent the CSS style "scroll-padding-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-inline-end
type ScrollPaddingInlineEndStyle string

func (t ScrollPaddingInlineEndStyle) Name() string {
	return "scroll-padding-inline-end"
}

var ScrollPaddingInlineEndStyleBrowserVariantsList = []string{}

func (t ScrollPaddingInlineEndStyle) BrowserVariants() []string {
	return ScrollPaddingInlineEndStyleBrowserVariantsList
}

var ScrollPaddingInlineEndStyleUtilitiesMapping = map[string]string{
	"scroll-padding-inline-end": "scroll-padding-inline-end: auto",
}

func (t ScrollPaddingInlineEndStyle) Utilities() map[string]string {
	return ScrollPaddingInlineEndStyleUtilitiesMapping
}

func (t ScrollPaddingInlineEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingInlineEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextAlign represent the CSS style "text-align" with value "start | end | left | right | center | justify | match-parent"
// See https://developer.mozilla.org/docs/Web/CSS/text-align
type TextAlignStyle string

func (t TextAlignStyle) Name() string {
	return "text-align"
}

const TextAlignStyleStart TextAlignStyle = "start"

const TextAlignStyleEnd TextAlignStyle = "end"

const TextAlignStyleLeft TextAlignStyle = "left"

const TextAlignStyleRight TextAlignStyle = "right"

const TextAlignStyleCenter TextAlignStyle = "center"

const TextAlignStyleJustify TextAlignStyle = "justify"

const TextAlignStyleMatchParent TextAlignStyle = "match-parent"

var TextAlignStyleBrowserVariantsList = []string{}

func (t TextAlignStyle) BrowserVariants() []string {
	return TextAlignStyleBrowserVariantsList
}

var TextAlignStyleUtilitiesMapping = map[string]string{
	"text-align":              "text-align: startOrNamelessValueIfLTRRightIfRTL",
	"text-align-start":        "text-align: start",
	"text-align-end":          "text-align: end",
	"text-align-left":         "text-align: left",
	"text-align-right":        "text-align: right",
	"text-align-center":       "text-align: center",
	"text-align-justify":      "text-align: justify",
	"text-align-match-parent": "text-align: match-parent",
}

func (t TextAlignStyle) Utilities() map[string]string {
	return TextAlignStyleUtilitiesMapping
}

func (t TextAlignStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextAlignStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AnimationDirection represent the CSS style "animation-direction" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-direction
type AnimationDirectionStyle string

func (t AnimationDirectionStyle) Name() string {
	return "animation-direction"
}

var AnimationDirectionStyleBrowserVariantsList = []string{}

func (t AnimationDirectionStyle) BrowserVariants() []string {
	return AnimationDirectionStyleBrowserVariantsList
}

var AnimationDirectionStyleUtilitiesMapping = map[string]string{
	"animation-direction": "animation-direction: normal",
}

func (t AnimationDirectionStyle) Utilities() map[string]string {
	return AnimationDirectionStyleUtilitiesMapping
}

func (t AnimationDirectionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AnimationDirectionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBottomLeftRadius represent the CSS style "border-bottom-left-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-left-radius
type BorderBottomLeftRadiusStyle string

func (t BorderBottomLeftRadiusStyle) Name() string {
	return "border-bottom-left-radius"
}

var BorderBottomLeftRadiusStyleBrowserVariantsList = []string{}

func (t BorderBottomLeftRadiusStyle) BrowserVariants() []string {
	return BorderBottomLeftRadiusStyleBrowserVariantsList
}

var BorderBottomLeftRadiusStyleUtilitiesMapping = map[string]string{
	"border-bottom-left-radius": "border-bottom-left-radius: 0",
}

func (t BorderBottomLeftRadiusStyle) Utilities() map[string]string {
	return BorderBottomLeftRadiusStyleUtilitiesMapping
}

func (t BorderBottomLeftRadiusStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBottomLeftRadiusStyleUtilitiesMapping[tu]
	return value, hasValue
}

// CounterReset represent the CSS style "counter-reset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/counter-reset
type CounterResetStyle string

func (t CounterResetStyle) Name() string {
	return "counter-reset"
}

var CounterResetStyleBrowserVariantsList = []string{}

func (t CounterResetStyle) BrowserVariants() []string {
	return CounterResetStyleBrowserVariantsList
}

var CounterResetStyleUtilitiesMapping = map[string]string{
	"counter-reset": "counter-reset: none",
}

func (t CounterResetStyle) Utilities() map[string]string {
	return CounterResetStyleUtilitiesMapping
}

func (t CounterResetStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = CounterResetStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextEmphasisStyle represent the CSS style "text-emphasis-style" with value "none | filled | open | dot | circle | double-circle | triangle | sesame | string"
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-style
type TextEmphasisStyleStyle string

func (t TextEmphasisStyleStyle) Name() string {
	return "text-emphasis-style"
}

const TextEmphasisStyleStyleNone TextEmphasisStyleStyle = "none"

const TextEmphasisStyleStyleFilled TextEmphasisStyleStyle = "filled"

const TextEmphasisStyleStyleOpen TextEmphasisStyleStyle = "open"

const TextEmphasisStyleStyleDot TextEmphasisStyleStyle = "dot"

const TextEmphasisStyleStyleCircle TextEmphasisStyleStyle = "circle"

const TextEmphasisStyleStyleDoubleCircle TextEmphasisStyleStyle = "double-circle"

const TextEmphasisStyleStyleTriangle TextEmphasisStyleStyle = "triangle"

const TextEmphasisStyleStyleSesame TextEmphasisStyleStyle = "sesame"

const TextEmphasisStyleStyleString TextEmphasisStyleStyle = "string"

var TextEmphasisStyleStyleBrowserVariantsList = []string{}

func (t TextEmphasisStyleStyle) BrowserVariants() []string {
	return TextEmphasisStyleStyleBrowserVariantsList
}

var TextEmphasisStyleStyleUtilitiesMapping = map[string]string{
	"text-emphasis-style":               "text-emphasis-style: none",
	"text-emphasis-style-none":          "text-emphasis-style: none",
	"text-emphasis-style-filled":        "text-emphasis-style: filled",
	"text-emphasis-style-open":          "text-emphasis-style: open",
	"text-emphasis-style-dot":           "text-emphasis-style: dot",
	"text-emphasis-style-circle":        "text-emphasis-style: circle",
	"text-emphasis-style-double-circle": "text-emphasis-style: double-circle",
	"text-emphasis-style-triangle":      "text-emphasis-style: triangle",
	"text-emphasis-style-sesame":        "text-emphasis-style: sesame",
	"text-emphasis-style-string":        "text-emphasis-style: string",
}

func (t TextEmphasisStyleStyle) Utilities() map[string]string {
	return TextEmphasisStyleStyleUtilitiesMapping
}

func (t TextEmphasisStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextEmphasisStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontLanguageOverride represent the CSS style "font-language-override" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-language-override
type FontLanguageOverrideStyle string

func (t FontLanguageOverrideStyle) Name() string {
	return "font-language-override"
}

var FontLanguageOverrideStyleBrowserVariantsList = []string{}

func (t FontLanguageOverrideStyle) BrowserVariants() []string {
	return FontLanguageOverrideStyleBrowserVariantsList
}

var FontLanguageOverrideStyleUtilitiesMapping = map[string]string{
	"font-language-override": "font-language-override: normal",
}

func (t FontLanguageOverrideStyle) Utilities() map[string]string {
	return FontLanguageOverrideStyleUtilitiesMapping
}

func (t FontLanguageOverrideStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontLanguageOverrideStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformOrigin represent the CSS style "transform-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformOriginStyle string

func (t TransformOriginStyle) Name() string {
	return "transform-origin"
}

var TransformOriginStyleBrowserVariantsList = []string{}

func (t TransformOriginStyle) BrowserVariants() []string {
	return TransformOriginStyleBrowserVariantsList
}

var TransformOriginStyleUtilitiesMapping = map[string]string{
	"transform-origin": "transform-origin: 50% 50% 0",
}

func (t TransformOriginStyle) Utilities() map[string]string {
	return TransformOriginStyleUtilitiesMapping
}

func (t TransformOriginStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformOriginStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColumnSpan represent the CSS style "column-span" with value "none | all"
// See https://developer.mozilla.org/docs/Web/CSS/column-span
type ColumnSpanStyle string

func (t ColumnSpanStyle) Name() string {
	return "column-span"
}

const ColumnSpanStyleNone ColumnSpanStyle = "none"

const ColumnSpanStyleAll ColumnSpanStyle = "all"

var ColumnSpanStyleBrowserVariantsList = []string{}

func (t ColumnSpanStyle) BrowserVariants() []string {
	return ColumnSpanStyleBrowserVariantsList
}

var ColumnSpanStyleUtilitiesMapping = map[string]string{
	"column-span":      "column-span: none",
	"column-span-none": "column-span: none",
	"column-span-all":  "column-span: all",
}

func (t ColumnSpanStyle) Utilities() map[string]string {
	return ColumnSpanStyleUtilitiesMapping
}

func (t ColumnSpanStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColumnSpanStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Opacity represent the CSS style "opacity" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/opacity
type OpacityStyle float64

func (t OpacityStyle) Name() string {
	return "opacity"
}

var OpacityStyleBrowserVariantsList = []string{}

func (t OpacityStyle) BrowserVariants() []string {
	return OpacityStyleBrowserVariantsList
}

var OpacityStyleUtilitiesMapping = map[string]string{
	"opacity": "opacity: 1.0",
}

func (t OpacityStyle) Utilities() map[string]string {
	return OpacityStyleUtilitiesMapping
}

func (t OpacityStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OpacityStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverscrollBehaviorY represent the CSS style "overscroll-behavior-y" with value "contain | none | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overscroll-behavior-y
type OverscrollBehaviorYStyle string

func (t OverscrollBehaviorYStyle) Name() string {
	return "overscroll-behavior-y"
}

const OverscrollBehaviorYStyleContain OverscrollBehaviorYStyle = "contain"

const OverscrollBehaviorYStyleNone OverscrollBehaviorYStyle = "none"

const OverscrollBehaviorYStyleAuto OverscrollBehaviorYStyle = "auto"

var OverscrollBehaviorYStyleBrowserVariantsList = []string{}

func (t OverscrollBehaviorYStyle) BrowserVariants() []string {
	return OverscrollBehaviorYStyleBrowserVariantsList
}

var OverscrollBehaviorYStyleUtilitiesMapping = map[string]string{
	"overscroll-behavior-y":         "overscroll-behavior-y: auto",
	"overscroll-behavior-y-contain": "overscroll-behavior-y: contain",
	"overscroll-behavior-y-none":    "overscroll-behavior-y: none",
	"overscroll-behavior-y-auto":    "overscroll-behavior-y: auto",
}

func (t OverscrollBehaviorYStyle) Utilities() map[string]string {
	return OverscrollBehaviorYStyleUtilitiesMapping
}

func (t OverscrollBehaviorYStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverscrollBehaviorYStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OffsetDistance represent the CSS style "offset-distance" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/offset-distance
type OffsetDistanceStyle string

func (t OffsetDistanceStyle) Name() string {
	return "offset-distance"
}

var OffsetDistanceStyleBrowserVariantsList = []string{}

func (t OffsetDistanceStyle) BrowserVariants() []string {
	return OffsetDistanceStyleBrowserVariantsList
}

var OffsetDistanceStyleUtilitiesMapping = map[string]string{
	"offset-distance": "offset-distance: 0",
}

func (t OffsetDistanceStyle) Utilities() map[string]string {
	return OffsetDistanceStyleUtilitiesMapping
}

func (t OffsetDistanceStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OffsetDistanceStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaddingRight represent the CSS style "padding-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-right
type PaddingRightStyle string

func (t PaddingRightStyle) Name() string {
	return "padding-right"
}

var PaddingRightStyleBrowserVariantsList = []string{}

func (t PaddingRightStyle) BrowserVariants() []string {
	return PaddingRightStyleBrowserVariantsList
}

var PaddingRightStyleUtilitiesMapping = map[string]string{
	"padding-right": "padding-right: 0",
}

func (t PaddingRightStyle) Utilities() map[string]string {
	return PaddingRightStyleUtilitiesMapping
}

func (t PaddingRightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaddingRightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Quotes represent the CSS style "quotes" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/quotes
type QuotesStyle string

func (t QuotesStyle) Name() string {
	return "quotes"
}

var QuotesStyleBrowserVariantsList = []string{}

func (t QuotesStyle) BrowserVariants() []string {
	return QuotesStyleBrowserVariantsList
}

var QuotesStyleUtilitiesMapping = map[string]string{
	"quotes": "quotes: dependsOnUserAgent",
}

func (t QuotesStyle) Utilities() map[string]string {
	return QuotesStyleUtilitiesMapping
}

func (t QuotesStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = QuotesStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollSnapPointsY represent the CSS style "scroll-snap-points-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-points-y
type ScrollSnapPointsYStyle string

func (t ScrollSnapPointsYStyle) Name() string {
	return "scroll-snap-points-y"
}

var ScrollSnapPointsYStyleBrowserVariantsList = []string{
	"-ms-scroll-snap-points-y",
}

func (t ScrollSnapPointsYStyle) BrowserVariants() []string {
	return ScrollSnapPointsYStyleBrowserVariantsList
}

var ScrollSnapPointsYStyleUtilitiesMapping = map[string]string{
	"scroll-snap-points-y": "scroll-snap-points-y: none",
}

func (t ScrollSnapPointsYStyle) Utilities() map[string]string {
	return ScrollSnapPointsYStyleUtilitiesMapping
}

func (t ScrollSnapPointsYStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollSnapPointsYStyleUtilitiesMapping[tu]
	return value, hasValue
}

// LineHeightStep represent the CSS style "line-height-step" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/line-height-step
type LineHeightStepStyle float64

func (t LineHeightStepStyle) Name() string {
	return "line-height-step"
}

var LineHeightStepStyleBrowserVariantsList = []string{}

func (t LineHeightStepStyle) BrowserVariants() []string {
	return LineHeightStepStyleBrowserVariantsList
}

var LineHeightStepStyleUtilitiesMapping = map[string]string{
	"line-height-step": "line-height-step: 0",
}

func (t LineHeightStepStyle) Utilities() map[string]string {
	return LineHeightStepStyleUtilitiesMapping
}

func (t LineHeightStepStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = LineHeightStepStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ShapeMargin represent the CSS style "shape-margin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/shape-margin
type ShapeMarginStyle string

func (t ShapeMarginStyle) Name() string {
	return "shape-margin"
}

var ShapeMarginStyleBrowserVariantsList = []string{}

func (t ShapeMarginStyle) BrowserVariants() []string {
	return ShapeMarginStyleBrowserVariantsList
}

var ShapeMarginStyleUtilitiesMapping = map[string]string{
	"shape-margin": "shape-margin: 0",
}

func (t ShapeMarginStyle) Utilities() map[string]string {
	return ShapeMarginStyleUtilitiesMapping
}

func (t ShapeMarginStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ShapeMarginStyleUtilitiesMapping[tu]
	return value, hasValue
}

// WritingMode represent the CSS style "writing-mode" with value "horizontal-tb | vertical-rl | vertical-lr | sideways-rl | sideways-lr"
// See https://developer.mozilla.org/docs/Web/CSS/writing-mode
type WritingModeStyle string

func (t WritingModeStyle) Name() string {
	return "writing-mode"
}

const WritingModeStyleHorizontalTb WritingModeStyle = "horizontal-tb"

const WritingModeStyleVerticalRl WritingModeStyle = "vertical-rl"

const WritingModeStyleVerticalLr WritingModeStyle = "vertical-lr"

const WritingModeStyleSidewaysRl WritingModeStyle = "sideways-rl"

const WritingModeStyleSidewaysLr WritingModeStyle = "sideways-lr"

var WritingModeStyleBrowserVariantsList = []string{}

func (t WritingModeStyle) BrowserVariants() []string {
	return WritingModeStyleBrowserVariantsList
}

var WritingModeStyleUtilitiesMapping = map[string]string{
	"writing-mode":               "writing-mode: horizontal-tb",
	"writing-mode-horizontal-tb": "writing-mode: horizontal-tb",
	"writing-mode-vertical-rl":   "writing-mode: vertical-rl",
	"writing-mode-vertical-lr":   "writing-mode: vertical-lr",
	"writing-mode-sideways-rl":   "writing-mode: sideways-rl",
	"writing-mode-sideways-lr":   "writing-mode: sideways-lr",
}

func (t WritingModeStyle) Utilities() map[string]string {
	return WritingModeStyleUtilitiesMapping
}

func (t WritingModeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = WritingModeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridAutoColumns represent the CSS style "grid-auto-columns" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-auto-columns
type GridAutoColumnsStyle string

func (t GridAutoColumnsStyle) Name() string {
	return "grid-auto-columns"
}

var GridAutoColumnsStyleBrowserVariantsList = []string{}

func (t GridAutoColumnsStyle) BrowserVariants() []string {
	return GridAutoColumnsStyleBrowserVariantsList
}

var GridAutoColumnsStyleUtilitiesMapping = map[string]string{
	"grid-auto-columns": "grid-auto-columns: auto",
}

func (t GridAutoColumnsStyle) Utilities() map[string]string {
	return GridAutoColumnsStyleUtilitiesMapping
}

func (t GridAutoColumnsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridAutoColumnsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridColumnGap represent the CSS style "grid-column-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-gap
type GridColumnGapStyle string

func (t GridColumnGapStyle) Name() string {
	return "grid-column-gap"
}

var GridColumnGapStyleBrowserVariantsList = []string{}

func (t GridColumnGapStyle) BrowserVariants() []string {
	return GridColumnGapStyleBrowserVariantsList
}

var GridColumnGapStyleUtilitiesMapping = map[string]string{
	"grid-column-gap": "grid-column-gap: 0",
}

func (t GridColumnGapStyle) Utilities() map[string]string {
	return GridColumnGapStyleUtilitiesMapping
}

func (t GridColumnGapStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridColumnGapStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderImageOutset represent the CSS style "border-image-outset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-outset
type BorderImageOutsetStyle string

func (t BorderImageOutsetStyle) Name() string {
	return "border-image-outset"
}

var BorderImageOutsetStyleBrowserVariantsList = []string{}

func (t BorderImageOutsetStyle) BrowserVariants() []string {
	return BorderImageOutsetStyleBrowserVariantsList
}

var BorderImageOutsetStyleUtilitiesMapping = map[string]string{
	"border-image-outset": "border-image-outset: 0",
}

func (t BorderImageOutsetStyle) Utilities() map[string]string {
	return BorderImageOutsetStyleUtilitiesMapping
}

func (t BorderImageOutsetStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderImageOutsetStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxOrdinalGroup represent the CSS style "box-ordinal-group" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/box-ordinal-group
type BoxOrdinalGroupStyle float64

func (t BoxOrdinalGroupStyle) Name() string {
	return "box-ordinal-group"
}

var BoxOrdinalGroupStyleBrowserVariantsList = []string{}

func (t BoxOrdinalGroupStyle) BrowserVariants() []string {
	return BoxOrdinalGroupStyleBrowserVariantsList
}

var BoxOrdinalGroupStyleUtilitiesMapping = map[string]string{
	"box-ordinal-group": "box-ordinal-group: 1",
}

func (t BoxOrdinalGroupStyle) Utilities() map[string]string {
	return BoxOrdinalGroupStyleUtilitiesMapping
}

func (t BoxOrdinalGroupStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxOrdinalGroupStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskComposite represent the CSS style "mask-composite" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-composite
type MaskCompositeStyle string

func (t MaskCompositeStyle) Name() string {
	return "mask-composite"
}

var MaskCompositeStyleBrowserVariantsList = []string{
	"-webkit-mask-composite",
}

func (t MaskCompositeStyle) BrowserVariants() []string {
	return MaskCompositeStyleBrowserVariantsList
}

var MaskCompositeStyleUtilitiesMapping = map[string]string{
	"mask-composite": "mask-composite: add",
}

func (t MaskCompositeStyle) Utilities() map[string]string {
	return MaskCompositeStyleUtilitiesMapping
}

func (t MaskCompositeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskCompositeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Inset represent the CSS style "inset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset
type InsetStyle string

func (t InsetStyle) Name() string {
	return "inset"
}

var InsetStyleBrowserVariantsList = []string{}

func (t InsetStyle) BrowserVariants() []string {
	return InsetStyleBrowserVariantsList
}

var InsetStyleUtilitiesMapping = map[string]string{
	"inset": "inset: auto",
}

func (t InsetStyle) Utilities() map[string]string {
	return InsetStyleUtilitiesMapping
}

func (t InsetStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = InsetStyleUtilitiesMapping[tu]
	return value, hasValue
}

// RubyMerge represent the CSS style "ruby-merge" with value "separate | collapse | auto"
// See
type RubyMergeStyle string

func (t RubyMergeStyle) Name() string {
	return "ruby-merge"
}

const RubyMergeStyleSeparate RubyMergeStyle = "separate"

const RubyMergeStyleCollapse RubyMergeStyle = "collapse"

const RubyMergeStyleAuto RubyMergeStyle = "auto"

var RubyMergeStyleBrowserVariantsList = []string{}

func (t RubyMergeStyle) BrowserVariants() []string {
	return RubyMergeStyleBrowserVariantsList
}

var RubyMergeStyleUtilitiesMapping = map[string]string{
	"ruby-merge":          "ruby-merge: separate",
	"ruby-merge-separate": "ruby-merge: separate",
	"ruby-merge-collapse": "ruby-merge: collapse",
	"ruby-merge-auto":     "ruby-merge: auto",
}

func (t RubyMergeStyle) Utilities() map[string]string {
	return RubyMergeStyleUtilitiesMapping
}

func (t RubyMergeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = RubyMergeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColorRendering represent the CSS style "color-rendering" with value "auto | optimizeSpeed | optimizeQuality"
// See
type ColorRenderingStyle string

func (t ColorRenderingStyle) Name() string {
	return "color-rendering"
}

const ColorRenderingStyleAuto ColorRenderingStyle = "auto"

const ColorRenderingStyleOptimizeSpeed ColorRenderingStyle = "optimizeSpeed"

const ColorRenderingStyleOptimizeQuality ColorRenderingStyle = "optimizeQuality"

var ColorRenderingStyleBrowserVariantsList = []string{}

func (t ColorRenderingStyle) BrowserVariants() []string {
	return ColorRenderingStyleBrowserVariantsList
}

var ColorRenderingStyleUtilitiesMapping = map[string]string{
	"color-rendering":                 "color-rendering: auto",
	"color-rendering-auto":            "color-rendering: auto",
	"color-rendering-optimizeSpeed":   "color-rendering: optimizeSpeed",
	"color-rendering-optimizeQuality": "color-rendering: optimizeQuality",
}

func (t ColorRenderingStyle) Utilities() map[string]string {
	return ColorRenderingStyleUtilitiesMapping
}

func (t ColorRenderingStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColorRenderingStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxPack represent the CSS style "box-pack" with value "start | center | end | justify"
// See https://developer.mozilla.org/docs/Web/CSS/box-pack
type BoxPackStyle string

func (t BoxPackStyle) Name() string {
	return "box-pack"
}

const BoxPackStyleStart BoxPackStyle = "start"

const BoxPackStyleCenter BoxPackStyle = "center"

const BoxPackStyleEnd BoxPackStyle = "end"

const BoxPackStyleJustify BoxPackStyle = "justify"

var BoxPackStyleBrowserVariantsList = []string{}

func (t BoxPackStyle) BrowserVariants() []string {
	return BoxPackStyleBrowserVariantsList
}

var BoxPackStyleUtilitiesMapping = map[string]string{
	"box-pack":         "box-pack: start",
	"box-pack-start":   "box-pack: start",
	"box-pack-center":  "box-pack: center",
	"box-pack-end":     "box-pack: end",
	"box-pack-justify": "box-pack: justify",
}

func (t BoxPackStyle) Utilities() map[string]string {
	return BoxPackStyleUtilitiesMapping
}

func (t BoxPackStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxPackStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridAutoFlow represent the CSS style "grid-auto-flow" with value "row | column | dense"
// See https://developer.mozilla.org/docs/Web/CSS/grid-auto-flow
type GridAutoFlowStyle string

func (t GridAutoFlowStyle) Name() string {
	return "grid-auto-flow"
}

const GridAutoFlowStyleRow GridAutoFlowStyle = "row"

const GridAutoFlowStyleColumn GridAutoFlowStyle = "column"

const GridAutoFlowStyleDense GridAutoFlowStyle = "dense"

var GridAutoFlowStyleBrowserVariantsList = []string{}

func (t GridAutoFlowStyle) BrowserVariants() []string {
	return GridAutoFlowStyleBrowserVariantsList
}

var GridAutoFlowStyleUtilitiesMapping = map[string]string{
	"grid-auto-flow":        "grid-auto-flow: row",
	"grid-auto-flow-row":    "grid-auto-flow: row",
	"grid-auto-flow-column": "grid-auto-flow: column",
	"grid-auto-flow-dense":  "grid-auto-flow: dense",
}

func (t GridAutoFlowStyle) Utilities() map[string]string {
	return GridAutoFlowStyleUtilitiesMapping
}

func (t GridAutoFlowStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridAutoFlowStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBlockStartStyle represent the CSS style "border-block-start-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-start-style
type BorderBlockStartStyleStyle string

func (t BorderBlockStartStyleStyle) Name() string {
	return "border-block-start-style"
}

var BorderBlockStartStyleStyleBrowserVariantsList = []string{}

func (t BorderBlockStartStyleStyle) BrowserVariants() []string {
	return BorderBlockStartStyleStyleBrowserVariantsList
}

var BorderBlockStartStyleStyleUtilitiesMapping = map[string]string{
	"border-block-start-style": "border-block-start-style: none",
}

func (t BorderBlockStartStyleStyle) Utilities() map[string]string {
	return BorderBlockStartStyleStyleUtilitiesMapping
}

func (t BorderBlockStartStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBlockStartStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColumnFill represent the CSS style "column-fill" with value "auto | balance | balance-all"
// See https://developer.mozilla.org/docs/Web/CSS/column-fill
type ColumnFillStyle string

func (t ColumnFillStyle) Name() string {
	return "column-fill"
}

const ColumnFillStyleAuto ColumnFillStyle = "auto"

const ColumnFillStyleBalance ColumnFillStyle = "balance"

const ColumnFillStyleBalanceAll ColumnFillStyle = "balance-all"

var ColumnFillStyleBrowserVariantsList = []string{}

func (t ColumnFillStyle) BrowserVariants() []string {
	return ColumnFillStyleBrowserVariantsList
}

var ColumnFillStyleUtilitiesMapping = map[string]string{
	"column-fill":             "column-fill: balance",
	"column-fill-auto":        "column-fill: auto",
	"column-fill-balance":     "column-fill: balance",
	"column-fill-balance-all": "column-fill: balance-all",
}

func (t ColumnFillStyle) Utilities() map[string]string {
	return ColumnFillStyleUtilitiesMapping
}

func (t ColumnFillStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColumnFillStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontSynthesis represent the CSS style "font-synthesis" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-synthesis
type FontSynthesisStyle string

func (t FontSynthesisStyle) Name() string {
	return "font-synthesis"
}

var FontSynthesisStyleBrowserVariantsList = []string{}

func (t FontSynthesisStyle) BrowserVariants() []string {
	return FontSynthesisStyleBrowserVariantsList
}

var FontSynthesisStyleUtilitiesMapping = map[string]string{
	"font-synthesis": "font-synthesis: weight style",
}

func (t FontSynthesisStyle) Utilities() map[string]string {
	return FontSynthesisStyleUtilitiesMapping
}

func (t FontSynthesisStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontSynthesisStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColorInterpolation represent the CSS style "color-interpolation" with value "auto | sRGB | linearRGB"
// See
type ColorInterpolationStyle string

func (t ColorInterpolationStyle) Name() string {
	return "color-interpolation"
}

const ColorInterpolationStyleAuto ColorInterpolationStyle = "auto"

const ColorInterpolationStyleSRGB ColorInterpolationStyle = "sRGB"

const ColorInterpolationStyleLinearRGB ColorInterpolationStyle = "linearRGB"

var ColorInterpolationStyleBrowserVariantsList = []string{}

func (t ColorInterpolationStyle) BrowserVariants() []string {
	return ColorInterpolationStyleBrowserVariantsList
}

var ColorInterpolationStyleUtilitiesMapping = map[string]string{
	"color-interpolation":           "color-interpolation: sRGB",
	"color-interpolation-auto":      "color-interpolation: auto",
	"color-interpolation-sRGB":      "color-interpolation: sRGB",
	"color-interpolation-linearRGB": "color-interpolation: linearRGB",
}

func (t ColorInterpolationStyle) Utilities() map[string]string {
	return ColorInterpolationStyleUtilitiesMapping
}

func (t ColorInterpolationStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColorInterpolationStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Transform represent the CSS style "transform" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform
type TransformStyle string

func (t TransformStyle) Name() string {
	return "transform"
}

var TransformStyleBrowserVariantsList = []string{}

func (t TransformStyle) BrowserVariants() []string {
	return TransformStyleBrowserVariantsList
}

var TransformStyleUtilitiesMapping = map[string]string{
	"transform": "transform: none",
}

func (t TransformStyle) Utilities() map[string]string {
	return TransformStyleUtilitiesMapping
}

func (t TransformStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaintOrder represent the CSS style "paint-order" with value "normal | fill | stroke | markers"
// See https://developer.mozilla.org/docs/Web/CSS/paint-order
type PaintOrderStyle string

func (t PaintOrderStyle) Name() string {
	return "paint-order"
}

const PaintOrderStyleNormal PaintOrderStyle = "normal"

const PaintOrderStyleFill PaintOrderStyle = "fill"

const PaintOrderStyleStroke PaintOrderStyle = "stroke"

const PaintOrderStyleMarkers PaintOrderStyle = "markers"

var PaintOrderStyleBrowserVariantsList = []string{}

func (t PaintOrderStyle) BrowserVariants() []string {
	return PaintOrderStyleBrowserVariantsList
}

var PaintOrderStyleUtilitiesMapping = map[string]string{
	"paint-order":         "paint-order: normal",
	"paint-order-normal":  "paint-order: normal",
	"paint-order-fill":    "paint-order: fill",
	"paint-order-stroke":  "paint-order: stroke",
	"paint-order-markers": "paint-order: markers",
}

func (t PaintOrderStyle) Utilities() map[string]string {
	return PaintOrderStyleUtilitiesMapping
}

func (t PaintOrderStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaintOrderStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundClip represent the CSS style "background-clip" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-clip
type BackgroundClipStyle string

func (t BackgroundClipStyle) Name() string {
	return "background-clip"
}

var BackgroundClipStyleBrowserVariantsList = []string{}

func (t BackgroundClipStyle) BrowserVariants() []string {
	return BackgroundClipStyleBrowserVariantsList
}

var BackgroundClipStyleUtilitiesMapping = map[string]string{
	"background-clip": "background-clip: border-box",
}

func (t BackgroundClipStyle) Utilities() map[string]string {
	return BackgroundClipStyleUtilitiesMapping
}

func (t BackgroundClipStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundClipStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformSkewX represent the CSS style "transform-skew-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformSkewXStyle string

func (t TransformSkewXStyle) Name() string {
	return "transform-skew-x"
}

var TransformSkewXStyleBrowserVariantsList = []string{}

func (t TransformSkewXStyle) BrowserVariants() []string {
	return TransformSkewXStyleBrowserVariantsList
}

var TransformSkewXStyleUtilitiesMapping = map[string]string{
	"transform-skew-x": "transform-skew-x: skewX(0)",
}

func (t TransformSkewXStyle) Utilities() map[string]string {
	return TransformSkewXStyleUtilitiesMapping
}

func (t TransformSkewXStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformSkewXStyleUtilitiesMapping[tu]
	return value, hasValue
}

// InsetBlock represent the CSS style "inset-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-block
type InsetBlockStyle string

func (t InsetBlockStyle) Name() string {
	return "inset-block"
}

var InsetBlockStyleBrowserVariantsList = []string{}

func (t InsetBlockStyle) BrowserVariants() []string {
	return InsetBlockStyleBrowserVariantsList
}

var InsetBlockStyleUtilitiesMapping = map[string]string{
	"inset-block": "inset-block: auto",
}

func (t InsetBlockStyle) Utilities() map[string]string {
	return InsetBlockStyleUtilitiesMapping
}

func (t InsetBlockStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = InsetBlockStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontVariantEastAsian represent the CSS style "font-variant-east-asian" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-east-asian
type FontVariantEastAsianStyle string

func (t FontVariantEastAsianStyle) Name() string {
	return "font-variant-east-asian"
}

var FontVariantEastAsianStyleBrowserVariantsList = []string{}

func (t FontVariantEastAsianStyle) BrowserVariants() []string {
	return FontVariantEastAsianStyleBrowserVariantsList
}

var FontVariantEastAsianStyleUtilitiesMapping = map[string]string{
	"font-variant-east-asian": "font-variant-east-asian: normal",
}

func (t FontVariantEastAsianStyle) Utilities() map[string]string {
	return FontVariantEastAsianStyleUtilitiesMapping
}

func (t FontVariantEastAsianStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontVariantEastAsianStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderInlineStartStyle represent the CSS style "border-inline-start-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-start-style
type BorderInlineStartStyleStyle string

func (t BorderInlineStartStyleStyle) Name() string {
	return "border-inline-start-style"
}

var BorderInlineStartStyleStyleBrowserVariantsList = []string{}

func (t BorderInlineStartStyleStyle) BrowserVariants() []string {
	return BorderInlineStartStyleStyleBrowserVariantsList
}

var BorderInlineStartStyleStyleUtilitiesMapping = map[string]string{
	"border-inline-start-style": "border-inline-start-style: none",
}

func (t BorderInlineStartStyleStyle) Utilities() map[string]string {
	return BorderInlineStartStyleStyleUtilitiesMapping
}

func (t BorderInlineStartStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderInlineStartStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxDirection represent the CSS style "box-direction" with value "normal | reverse | inherit"
// See https://developer.mozilla.org/docs/Web/CSS/box-direction
type BoxDirectionStyle string

func (t BoxDirectionStyle) Name() string {
	return "box-direction"
}

const BoxDirectionStyleNormal BoxDirectionStyle = "normal"

const BoxDirectionStyleReverse BoxDirectionStyle = "reverse"

const BoxDirectionStyleInherit BoxDirectionStyle = "inherit"

var BoxDirectionStyleBrowserVariantsList = []string{}

func (t BoxDirectionStyle) BrowserVariants() []string {
	return BoxDirectionStyleBrowserVariantsList
}

var BoxDirectionStyleUtilitiesMapping = map[string]string{
	"box-direction":         "box-direction: normal",
	"box-direction-normal":  "box-direction: normal",
	"box-direction-reverse": "box-direction: reverse",
	"box-direction-inherit": "box-direction: inherit",
}

func (t BoxDirectionStyle) Utilities() map[string]string {
	return BoxDirectionStyleUtilitiesMapping
}

func (t BoxDirectionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxDirectionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// LineBreak represent the CSS style "line-break" with value "auto | loose | normal | strict | anywhere"
// See https://developer.mozilla.org/docs/Web/CSS/line-break
type LineBreakStyle string

func (t LineBreakStyle) Name() string {
	return "line-break"
}

const LineBreakStyleAuto LineBreakStyle = "auto"

const LineBreakStyleLoose LineBreakStyle = "loose"

const LineBreakStyleNormal LineBreakStyle = "normal"

const LineBreakStyleStrict LineBreakStyle = "strict"

const LineBreakStyleAnywhere LineBreakStyle = "anywhere"

var LineBreakStyleBrowserVariantsList = []string{}

func (t LineBreakStyle) BrowserVariants() []string {
	return LineBreakStyleBrowserVariantsList
}

var LineBreakStyleUtilitiesMapping = map[string]string{
	"line-break":          "line-break: auto",
	"line-break-auto":     "line-break: auto",
	"line-break-loose":    "line-break: loose",
	"line-break-normal":   "line-break: normal",
	"line-break-strict":   "line-break: strict",
	"line-break-anywhere": "line-break: anywhere",
}

func (t LineBreakStyle) Utilities() map[string]string {
	return LineBreakStyleUtilitiesMapping
}

func (t LineBreakStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = LineBreakStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OutlineOffset represent the CSS style "outline-offset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-offset
type OutlineOffsetStyle float64

func (t OutlineOffsetStyle) Name() string {
	return "outline-offset"
}

var OutlineOffsetStyleBrowserVariantsList = []string{}

func (t OutlineOffsetStyle) BrowserVariants() []string {
	return OutlineOffsetStyleBrowserVariantsList
}

var OutlineOffsetStyleUtilitiesMapping = map[string]string{
	"outline-offset": "outline-offset: 0",
}

func (t OutlineOffsetStyle) Utilities() map[string]string {
	return OutlineOffsetStyleUtilitiesMapping
}

func (t OutlineOffsetStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OutlineOffsetStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextOverflow represent the CSS style "text-overflow" with value "clip | ellipsis | clip ellipsis | ellipsis clip"
// See https://developer.mozilla.org/docs/Web/CSS/text-overflow
type TextOverflowStyle string

func (t TextOverflowStyle) Name() string {
	return "text-overflow"
}

const TextOverflowStyleClip TextOverflowStyle = "clip"

const TextOverflowStyleEllipsis TextOverflowStyle = "ellipsis"

const TextOverflowStyleClipEllipsis TextOverflowStyle = "clip-ellipsis"

const TextOverflowStyleEllipsisClip TextOverflowStyle = "ellipsis-clip"

var TextOverflowStyleBrowserVariantsList = []string{}

func (t TextOverflowStyle) BrowserVariants() []string {
	return TextOverflowStyleBrowserVariantsList
}

var TextOverflowStyleUtilitiesMapping = map[string]string{
	"text-overflow":               "text-overflow: clip",
	"text-overflow-clip":          "text-overflow: clip",
	"text-overflow-ellipsis":      "text-overflow: ellipsis",
	"text-overflow-clip-ellipsis": "text-overflow: clip-ellipsis",
	"text-overflow-ellipsis-clip": "text-overflow: ellipsis-clip",
}

func (t TextOverflowStyle) Utilities() map[string]string {
	return TextOverflowStyleUtilitiesMapping
}

func (t TextOverflowStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextOverflowStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBlockStartColor represent the CSS style "border-block-start-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-start-color
type BorderBlockStartColorStyle Color

func (t BorderBlockStartColorStyle) Name() string {
	return "border-block-start-color"
}

var BorderBlockStartColorStyleBrowserVariantsList = []string{}

func (t BorderBlockStartColorStyle) BrowserVariants() []string {
	return BorderBlockStartColorStyleBrowserVariantsList
}

var BorderBlockStartColorStyleUtilitiesMapping = map[string]string{
	"border-block-start-color": "border-block-start-color: currentcolor",
}

func (t BorderBlockStartColorStyle) Utilities() map[string]string {
	return BorderBlockStartColorStyleUtilitiesMapping
}

func (t BorderBlockStartColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBlockStartColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextIndent represent the CSS style "text-indent" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-indent
type TextIndentStyle string

func (t TextIndentStyle) Name() string {
	return "text-indent"
}

var TextIndentStyleBrowserVariantsList = []string{}

func (t TextIndentStyle) BrowserVariants() []string {
	return TextIndentStyleBrowserVariantsList
}

var TextIndentStyleUtilitiesMapping = map[string]string{
	"text-indent": "text-indent: 0",
}

func (t TextIndentStyle) Utilities() map[string]string {
	return TextIndentStyleUtilitiesMapping
}

func (t TextIndentStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextIndentStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FlexDirection represent the CSS style "flex-direction" with value "row | row-reverse | column | column-reverse"
// See https://developer.mozilla.org/docs/Web/CSS/flex-direction
type FlexDirectionStyle string

func (t FlexDirectionStyle) Name() string {
	return "flex-direction"
}

const FlexDirectionStyleRow FlexDirectionStyle = "row"

const FlexDirectionStyleRowReverse FlexDirectionStyle = "row-reverse"

const FlexDirectionStyleColumn FlexDirectionStyle = "column"

const FlexDirectionStyleColumnReverse FlexDirectionStyle = "column-reverse"

var FlexDirectionStyleBrowserVariantsList = []string{}

func (t FlexDirectionStyle) BrowserVariants() []string {
	return FlexDirectionStyleBrowserVariantsList
}

var FlexDirectionStyleUtilitiesMapping = map[string]string{
	"flex-direction":                "flex-direction: row",
	"flex-direction-row":            "flex-direction: row",
	"flex-direction-row-reverse":    "flex-direction: row-reverse",
	"flex-direction-column":         "flex-direction: column",
	"flex-direction-column-reverse": "flex-direction: column-reverse",
}

func (t FlexDirectionStyle) Utilities() map[string]string {
	return FlexDirectionStyleUtilitiesMapping
}

func (t FlexDirectionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FlexDirectionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// InlineSize represent the CSS style "inline-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inline-size
type InlineSizeStyle string

func (t InlineSizeStyle) Name() string {
	return "inline-size"
}

var InlineSizeStyleBrowserVariantsList = []string{}

func (t InlineSizeStyle) BrowserVariants() []string {
	return InlineSizeStyleBrowserVariantsList
}

var InlineSizeStyleUtilitiesMapping = map[string]string{
	"inline-size": "inline-size: auto",
}

func (t InlineSizeStyle) Utilities() map[string]string {
	return InlineSizeStyleUtilitiesMapping
}

func (t InlineSizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = InlineSizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// LetterSpacing represent the CSS style "letter-spacing" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/letter-spacing
type LetterSpacingStyle string

func (t LetterSpacingStyle) Name() string {
	return "letter-spacing"
}

var LetterSpacingStyleBrowserVariantsList = []string{}

func (t LetterSpacingStyle) BrowserVariants() []string {
	return LetterSpacingStyleBrowserVariantsList
}

var LetterSpacingStyleUtilitiesMapping = map[string]string{
	"letter-spacing": "letter-spacing: normal",
}

func (t LetterSpacingStyle) Utilities() map[string]string {
	return LetterSpacingStyleUtilitiesMapping
}

func (t LetterSpacingStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = LetterSpacingStyleUtilitiesMapping[tu]
	return value, hasValue
}

// CounterIncrement represent the CSS style "counter-increment" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/counter-increment
type CounterIncrementStyle string

func (t CounterIncrementStyle) Name() string {
	return "counter-increment"
}

var CounterIncrementStyleBrowserVariantsList = []string{}

func (t CounterIncrementStyle) BrowserVariants() []string {
	return CounterIncrementStyleBrowserVariantsList
}

var CounterIncrementStyleUtilitiesMapping = map[string]string{
	"counter-increment": "counter-increment: none",
}

func (t CounterIncrementStyle) Utilities() map[string]string {
	return CounterIncrementStyleUtilitiesMapping
}

func (t CounterIncrementStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = CounterIncrementStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MinHeight represent the CSS style "min-height" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-height
type MinHeightStyle string

func (t MinHeightStyle) Name() string {
	return "min-height"
}

var MinHeightStyleBrowserVariantsList = []string{}

func (t MinHeightStyle) BrowserVariants() []string {
	return MinHeightStyleBrowserVariantsList
}

var MinHeightStyleUtilitiesMapping = map[string]string{
	"min-height": "min-height: auto",
}

func (t MinHeightStyle) Utilities() map[string]string {
	return MinHeightStyleUtilitiesMapping
}

func (t MinHeightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MinHeightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollSnapType represent the CSS style "scroll-snap-type" with value "none | block | inline | both | mandatory | proximity"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-type
type ScrollSnapTypeStyle string

func (t ScrollSnapTypeStyle) Name() string {
	return "scroll-snap-type"
}

const ScrollSnapTypeStyleNone ScrollSnapTypeStyle = "none"

const ScrollSnapTypeStyleBlock ScrollSnapTypeStyle = "block"

const ScrollSnapTypeStyleInline ScrollSnapTypeStyle = "inline"

const ScrollSnapTypeStyleBoth ScrollSnapTypeStyle = "both"

const ScrollSnapTypeStyleMandatory ScrollSnapTypeStyle = "mandatory"

const ScrollSnapTypeStyleProximity ScrollSnapTypeStyle = "proximity"

var ScrollSnapTypeStyleBrowserVariantsList = []string{
	"-ms-scroll-snap-type",
}

func (t ScrollSnapTypeStyle) BrowserVariants() []string {
	return ScrollSnapTypeStyleBrowserVariantsList
}

var ScrollSnapTypeStyleUtilitiesMapping = map[string]string{
	"scroll-snap-type":           "scroll-snap-type: none",
	"scroll-snap-type-none":      "scroll-snap-type: none",
	"scroll-snap-type-block":     "scroll-snap-type: block",
	"scroll-snap-type-inline":    "scroll-snap-type: inline",
	"scroll-snap-type-both":      "scroll-snap-type: both",
	"scroll-snap-type-mandatory": "scroll-snap-type: mandatory",
	"scroll-snap-type-proximity": "scroll-snap-type: proximity",
}

func (t ScrollSnapTypeStyle) Utilities() map[string]string {
	return ScrollSnapTypeStyleUtilitiesMapping
}

func (t ScrollSnapTypeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollSnapTypeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BlockOverflow represent the CSS style "block-overflow" with value "clip | ellipsis | string"
// See
type BlockOverflowStyle string

func (t BlockOverflowStyle) Name() string {
	return "block-overflow"
}

const BlockOverflowStyleClip BlockOverflowStyle = "clip"

const BlockOverflowStyleEllipsis BlockOverflowStyle = "ellipsis"

const BlockOverflowStyleString BlockOverflowStyle = "string"

var BlockOverflowStyleBrowserVariantsList = []string{}

func (t BlockOverflowStyle) BrowserVariants() []string {
	return BlockOverflowStyleBrowserVariantsList
}

var BlockOverflowStyleUtilitiesMapping = map[string]string{
	"block-overflow":          "block-overflow: clip",
	"block-overflow-clip":     "block-overflow: clip",
	"block-overflow-ellipsis": "block-overflow: ellipsis",
	"block-overflow-string":   "block-overflow: string",
}

func (t BlockOverflowStyle) Utilities() map[string]string {
	return BlockOverflowStyleUtilitiesMapping
}

func (t BlockOverflowStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BlockOverflowStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColumnGap represent the CSS style "column-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-gap
type ColumnGapStyle string

func (t ColumnGapStyle) Name() string {
	return "column-gap"
}

var ColumnGapStyleBrowserVariantsList = []string{}

func (t ColumnGapStyle) BrowserVariants() []string {
	return ColumnGapStyleBrowserVariantsList
}

var ColumnGapStyleUtilitiesMapping = map[string]string{
	"column-gap": "column-gap: normal",
}

func (t ColumnGapStyle) Utilities() map[string]string {
	return ColumnGapStyleUtilitiesMapping
}

func (t ColumnGapStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColumnGapStyleUtilitiesMapping[tu]
	return value, hasValue
}

// InsetBlockEnd represent the CSS style "inset-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-block-end
type InsetBlockEndStyle string

func (t InsetBlockEndStyle) Name() string {
	return "inset-block-end"
}

var InsetBlockEndStyleBrowserVariantsList = []string{}

func (t InsetBlockEndStyle) BrowserVariants() []string {
	return InsetBlockEndStyleBrowserVariantsList
}

var InsetBlockEndStyleUtilitiesMapping = map[string]string{
	"inset-block-end": "inset-block-end: auto",
}

func (t InsetBlockEndStyle) Utilities() map[string]string {
	return InsetBlockEndStyleUtilitiesMapping
}

func (t InsetBlockEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = InsetBlockEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverflowY represent the CSS style "overflow-y" with value "visible | hidden | clip | scroll | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overflow-y
type OverflowYStyle string

func (t OverflowYStyle) Name() string {
	return "overflow-y"
}

const OverflowYStyleVisible OverflowYStyle = "visible"

const OverflowYStyleHidden OverflowYStyle = "hidden"

const OverflowYStyleClip OverflowYStyle = "clip"

const OverflowYStyleScroll OverflowYStyle = "scroll"

const OverflowYStyleAuto OverflowYStyle = "auto"

var OverflowYStyleBrowserVariantsList = []string{}

func (t OverflowYStyle) BrowserVariants() []string {
	return OverflowYStyleBrowserVariantsList
}

var OverflowYStyleUtilitiesMapping = map[string]string{
	"overflow-y":         "overflow-y: visible",
	"overflow-y-visible": "overflow-y: visible",
	"overflow-y-hidden":  "overflow-y: hidden",
	"overflow-y-clip":    "overflow-y: clip",
	"overflow-y-scroll":  "overflow-y: scroll",
	"overflow-y-auto":    "overflow-y: auto",
}

func (t OverflowYStyle) Utilities() map[string]string {
	return OverflowYStyleUtilitiesMapping
}

func (t OverflowYStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverflowYStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextAnchor represent the CSS style "text-anchor" with value "start | middle | end"
// See
type TextAnchorStyle string

func (t TextAnchorStyle) Name() string {
	return "text-anchor"
}

const TextAnchorStyleStart TextAnchorStyle = "start"

const TextAnchorStyleMiddle TextAnchorStyle = "middle"

const TextAnchorStyleEnd TextAnchorStyle = "end"

var TextAnchorStyleBrowserVariantsList = []string{}

func (t TextAnchorStyle) BrowserVariants() []string {
	return TextAnchorStyleBrowserVariantsList
}

var TextAnchorStyleUtilitiesMapping = map[string]string{
	"text-anchor":        "text-anchor: start",
	"text-anchor-start":  "text-anchor: start",
	"text-anchor-middle": "text-anchor: middle",
	"text-anchor-end":    "text-anchor: end",
}

func (t TextAnchorStyle) Utilities() map[string]string {
	return TextAnchorStyleUtilitiesMapping
}

func (t TextAnchorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextAnchorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// InsetBlockStart represent the CSS style "inset-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-block-start
type InsetBlockStartStyle string

func (t InsetBlockStartStyle) Name() string {
	return "inset-block-start"
}

var InsetBlockStartStyleBrowserVariantsList = []string{}

func (t InsetBlockStartStyle) BrowserVariants() []string {
	return InsetBlockStartStyleBrowserVariantsList
}

var InsetBlockStartStyleUtilitiesMapping = map[string]string{
	"inset-block-start": "inset-block-start: auto",
}

func (t InsetBlockStartStyle) Utilities() map[string]string {
	return InsetBlockStartStyleUtilitiesMapping
}

func (t InsetBlockStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = InsetBlockStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginBottom represent the CSS style "margin-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-bottom
type MarginBottomStyle string

func (t MarginBottomStyle) Name() string {
	return "margin-bottom"
}

var MarginBottomStyleBrowserVariantsList = []string{}

func (t MarginBottomStyle) BrowserVariants() []string {
	return MarginBottomStyleBrowserVariantsList
}

var MarginBottomStyleUtilitiesMapping = map[string]string{
	"margin-bottom": "margin-bottom: 0",
}

func (t MarginBottomStyle) Utilities() map[string]string {
	return MarginBottomStyleUtilitiesMapping
}

func (t MarginBottomStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginBottomStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextJustify represent the CSS style "text-justify" with value "auto | inter-character | inter-word | none"
// See https://developer.mozilla.org/docs/Web/CSS/text-justify
type TextJustifyStyle string

func (t TextJustifyStyle) Name() string {
	return "text-justify"
}

const TextJustifyStyleAuto TextJustifyStyle = "auto"

const TextJustifyStyleInterCharacter TextJustifyStyle = "inter-character"

const TextJustifyStyleInterWord TextJustifyStyle = "inter-word"

const TextJustifyStyleNone TextJustifyStyle = "none"

var TextJustifyStyleBrowserVariantsList = []string{}

func (t TextJustifyStyle) BrowserVariants() []string {
	return TextJustifyStyleBrowserVariantsList
}

var TextJustifyStyleUtilitiesMapping = map[string]string{
	"text-justify":                 "text-justify: auto",
	"text-justify-auto":            "text-justify: auto",
	"text-justify-inter-character": "text-justify: inter-character",
	"text-justify-inter-word":      "text-justify: inter-word",
	"text-justify-none":            "text-justify: none",
}

func (t TextJustifyStyle) Utilities() map[string]string {
	return TextJustifyStyleUtilitiesMapping
}

func (t TextJustifyStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextJustifyStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Top represent the CSS style "top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/top
type TopStyle string

func (t TopStyle) Name() string {
	return "top"
}

var TopStyleBrowserVariantsList = []string{}

func (t TopStyle) BrowserVariants() []string {
	return TopStyleBrowserVariantsList
}

var TopStyleUtilitiesMapping = map[string]string{
	"top": "top: auto",
}

func (t TopStyle) Utilities() map[string]string {
	return TopStyleUtilitiesMapping
}

func (t TopStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TopStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontVariant represent the CSS style "font-variant" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant
type FontVariantStyle string

func (t FontVariantStyle) Name() string {
	return "font-variant"
}

var FontVariantStyleBrowserVariantsList = []string{}

func (t FontVariantStyle) BrowserVariants() []string {
	return FontVariantStyleBrowserVariantsList
}

var FontVariantStyleUtilitiesMapping = map[string]string{
	"font-variant": "font-variant: normal",
}

func (t FontVariantStyle) Utilities() map[string]string {
	return FontVariantStyleUtilitiesMapping
}

func (t FontVariantStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontVariantStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderLeftWidth represent the CSS style "border-left-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-left-width
type BorderLeftWidthStyle string

func (t BorderLeftWidthStyle) Name() string {
	return "border-left-width"
}

var BorderLeftWidthStyleBrowserVariantsList = []string{}

func (t BorderLeftWidthStyle) BrowserVariants() []string {
	return BorderLeftWidthStyleBrowserVariantsList
}

var BorderLeftWidthStyleUtilitiesMapping = map[string]string{
	"border-left-width": "border-left-width: medium",
}

func (t BorderLeftWidthStyle) Utilities() map[string]string {
	return BorderLeftWidthStyleUtilitiesMapping
}

func (t BorderLeftWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderLeftWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColumnWidth represent the CSS style "column-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-width
type ColumnWidthStyle string

func (t ColumnWidthStyle) Name() string {
	return "column-width"
}

var ColumnWidthStyleBrowserVariantsList = []string{}

func (t ColumnWidthStyle) BrowserVariants() []string {
	return ColumnWidthStyleBrowserVariantsList
}

var ColumnWidthStyleUtilitiesMapping = map[string]string{
	"column-width": "column-width: auto",
}

func (t ColumnWidthStyle) Utilities() map[string]string {
	return ColumnWidthStyleUtilitiesMapping
}

func (t ColumnWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColumnWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackdropFilter represent the CSS style "backdrop-filter" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/backdrop-filter
type BackdropFilterStyle string

func (t BackdropFilterStyle) Name() string {
	return "backdrop-filter"
}

var BackdropFilterStyleBrowserVariantsList = []string{}

func (t BackdropFilterStyle) BrowserVariants() []string {
	return BackdropFilterStyleBrowserVariantsList
}

var BackdropFilterStyleUtilitiesMapping = map[string]string{
	"backdrop-filter": "backdrop-filter: none",
}

func (t BackdropFilterStyle) Utilities() map[string]string {
	return BackdropFilterStyleUtilitiesMapping
}

func (t BackdropFilterStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackdropFilterStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderInlineColor represent the CSS style "border-inline-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-color
type BorderInlineColorStyle Color

func (t BorderInlineColorStyle) Name() string {
	return "border-inline-color"
}

var BorderInlineColorStyleBrowserVariantsList = []string{}

func (t BorderInlineColorStyle) BrowserVariants() []string {
	return BorderInlineColorStyleBrowserVariantsList
}

var BorderInlineColorStyleUtilitiesMapping = map[string]string{
	"border-inline-color": "border-inline-color: currentcolor",
}

func (t BorderInlineColorStyle) Utilities() map[string]string {
	return BorderInlineColorStyleUtilitiesMapping
}

func (t BorderInlineColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderInlineColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// WordWrap represent the CSS style "word-wrap" with value "normal | break-word"
// See https://developer.mozilla.org/docs/Web/CSS/overflow-wrap
type WordWrapStyle string

func (t WordWrapStyle) Name() string {
	return "word-wrap"
}

const WordWrapStyleNormal WordWrapStyle = "normal"

const WordWrapStyleBreakWord WordWrapStyle = "break-word"

var WordWrapStyleBrowserVariantsList = []string{}

func (t WordWrapStyle) BrowserVariants() []string {
	return WordWrapStyleBrowserVariantsList
}

var WordWrapStyleUtilitiesMapping = map[string]string{
	"word-wrap":            "word-wrap: normal",
	"word-wrap-normal":     "word-wrap: normal",
	"word-wrap-break-word": "word-wrap: break-word",
}

func (t WordWrapStyle) Utilities() map[string]string {
	return WordWrapStyleUtilitiesMapping
}

func (t WordWrapStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = WordWrapStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Content represent the CSS style "content" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/content
type ContentStyle string

func (t ContentStyle) Name() string {
	return "content"
}

var ContentStyleBrowserVariantsList = []string{}

func (t ContentStyle) BrowserVariants() []string {
	return ContentStyleBrowserVariantsList
}

var ContentStyleUtilitiesMapping = map[string]string{
	"content": "content: normal",
}

func (t ContentStyle) Utilities() map[string]string {
	return ContentStyleUtilitiesMapping
}

func (t ContentStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ContentStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Height represent the CSS style "height" with value "auto |  min-content | max-content"
// See https://developer.mozilla.org/docs/Web/CSS/height
type HeightStyle string

func (t HeightStyle) Name() string {
	return "height"
}

const HeightStyleAuto HeightStyle = "auto"

const HeightStyleMinContent HeightStyle = "min-content"

const HeightStyleMaxContent HeightStyle = "max-content"

var HeightStyleBrowserVariantsList = []string{}

func (t HeightStyle) BrowserVariants() []string {
	return HeightStyleBrowserVariantsList
}

var HeightStyleUtilitiesMapping = map[string]string{
	"height":             "height: auto",
	"height-auto":        "height: auto",
	"height-min-content": "height: min-content",
	"height-max-content": "height: max-content",
}

func (t HeightStyle) Utilities() map[string]string {
	return HeightStyleUtilitiesMapping
}

func (t HeightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = HeightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// RubyPosition represent the CSS style "ruby-position" with value "over | under | inter-character"
// See https://developer.mozilla.org/docs/Web/CSS/ruby-position
type RubyPositionStyle string

func (t RubyPositionStyle) Name() string {
	return "ruby-position"
}

const RubyPositionStyleOver RubyPositionStyle = "over"

const RubyPositionStyleUnder RubyPositionStyle = "under"

const RubyPositionStyleInterCharacter RubyPositionStyle = "inter-character"

var RubyPositionStyleBrowserVariantsList = []string{}

func (t RubyPositionStyle) BrowserVariants() []string {
	return RubyPositionStyleBrowserVariantsList
}

var RubyPositionStyleUtilitiesMapping = map[string]string{
	"ruby-position":                 "ruby-position: over",
	"ruby-position-over":            "ruby-position: over",
	"ruby-position-under":           "ruby-position: under",
	"ruby-position-inter-character": "ruby-position: inter-character",
}

func (t RubyPositionStyle) Utilities() map[string]string {
	return RubyPositionStyleUtilitiesMapping
}

func (t RubyPositionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = RubyPositionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GlyphOrientationVertical represent the CSS style "glyph-orientation-vertical" with value ""
// See
type GlyphOrientationVerticalStyle string

func (t GlyphOrientationVerticalStyle) Name() string {
	return "glyph-orientation-vertical"
}

var GlyphOrientationVerticalStyleBrowserVariantsList = []string{}

func (t GlyphOrientationVerticalStyle) BrowserVariants() []string {
	return GlyphOrientationVerticalStyleBrowserVariantsList
}

var GlyphOrientationVerticalStyleUtilitiesMapping = map[string]string{
	"glyph-orientation-vertical": "glyph-orientation-vertical: auto",
}

func (t GlyphOrientationVerticalStyle) Utilities() map[string]string {
	return GlyphOrientationVerticalStyleUtilitiesMapping
}

func (t GlyphOrientationVerticalStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GlyphOrientationVerticalStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundAttachment represent the CSS style "background-attachment" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-attachment
type BackgroundAttachmentStyle string

func (t BackgroundAttachmentStyle) Name() string {
	return "background-attachment"
}

var BackgroundAttachmentStyleBrowserVariantsList = []string{}

func (t BackgroundAttachmentStyle) BrowserVariants() []string {
	return BackgroundAttachmentStyleBrowserVariantsList
}

var BackgroundAttachmentStyleUtilitiesMapping = map[string]string{
	"background-attachment": "background-attachment: scroll",
}

func (t BackgroundAttachmentStyle) Utilities() map[string]string {
	return BackgroundAttachmentStyleUtilitiesMapping
}

func (t BackgroundAttachmentStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundAttachmentStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollbarWidth represent the CSS style "scrollbar-width" with value "auto | thin | none"
// See https://developer.mozilla.org/docs/Web/CSS/scrollbar-width
type ScrollbarWidthStyle string

func (t ScrollbarWidthStyle) Name() string {
	return "scrollbar-width"
}

const ScrollbarWidthStyleAuto ScrollbarWidthStyle = "auto"

const ScrollbarWidthStyleThin ScrollbarWidthStyle = "thin"

const ScrollbarWidthStyleNone ScrollbarWidthStyle = "none"

var ScrollbarWidthStyleBrowserVariantsList = []string{}

func (t ScrollbarWidthStyle) BrowserVariants() []string {
	return ScrollbarWidthStyleBrowserVariantsList
}

var ScrollbarWidthStyleUtilitiesMapping = map[string]string{
	"scrollbar-width":      "scrollbar-width: auto",
	"scrollbar-width-auto": "scrollbar-width: auto",
	"scrollbar-width-thin": "scrollbar-width: thin",
	"scrollbar-width-none": "scrollbar-width: none",
}

func (t ScrollbarWidthStyle) Utilities() map[string]string {
	return ScrollbarWidthStyleUtilitiesMapping
}

func (t ScrollbarWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollbarWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPadding represent the CSS style "scroll-padding" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding
type ScrollPaddingStyle string

func (t ScrollPaddingStyle) Name() string {
	return "scroll-padding"
}

var ScrollPaddingStyleBrowserVariantsList = []string{}

func (t ScrollPaddingStyle) BrowserVariants() []string {
	return ScrollPaddingStyleBrowserVariantsList
}

var ScrollPaddingStyleUtilitiesMapping = map[string]string{
	"scroll-padding": "scroll-padding: auto",
}

func (t ScrollPaddingStyle) Utilities() map[string]string {
	return ScrollPaddingStyleUtilitiesMapping
}

func (t ScrollPaddingStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformSkew represent the CSS style "transform-skew" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformSkewStyle string

func (t TransformSkewStyle) Name() string {
	return "transform-skew"
}

var TransformSkewStyleBrowserVariantsList = []string{}

func (t TransformSkewStyle) BrowserVariants() []string {
	return TransformSkewStyleBrowserVariantsList
}

var TransformSkewStyleUtilitiesMapping = map[string]string{
	"transform-skew": "transform-skew: skew(0,0)",
}

func (t TransformSkewStyle) Utilities() map[string]string {
	return TransformSkewStyleUtilitiesMapping
}

func (t TransformSkewStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformSkewStyleUtilitiesMapping[tu]
	return value, hasValue
}

// All represent the CSS style "all" with value "initial | inherit | unset | revert"
// See https://developer.mozilla.org/docs/Web/CSS/all
type AllStyle string

func (t AllStyle) Name() string {
	return "all"
}

const AllStyleInitial AllStyle = "initial"

const AllStyleInherit AllStyle = "inherit"

const AllStyleUnset AllStyle = "unset"

const AllStyleRevert AllStyle = "revert"

var AllStyleBrowserVariantsList = []string{}

func (t AllStyle) BrowserVariants() []string {
	return AllStyleBrowserVariantsList
}

var AllStyleUtilitiesMapping = map[string]string{
	"all":         "all: noPracticalInitialValue",
	"all-initial": "all: initial",
	"all-inherit": "all: inherit",
	"all-unset":   "all: unset",
	"all-revert":  "all: revert",
}

func (t AllStyle) Utilities() map[string]string {
	return AllStyleUtilitiesMapping
}

func (t AllStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AllStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderInlineEndWidth represent the CSS style "border-inline-end-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-end-width
type BorderInlineEndWidthStyle string

func (t BorderInlineEndWidthStyle) Name() string {
	return "border-inline-end-width"
}

var BorderInlineEndWidthStyleBrowserVariantsList = []string{}

func (t BorderInlineEndWidthStyle) BrowserVariants() []string {
	return BorderInlineEndWidthStyleBrowserVariantsList
}

var BorderInlineEndWidthStyleUtilitiesMapping = map[string]string{
	"border-inline-end-width": "border-inline-end-width: medium",
}

func (t BorderInlineEndWidthStyle) Utilities() map[string]string {
	return BorderInlineEndWidthStyleUtilitiesMapping
}

func (t BorderInlineEndWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderInlineEndWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FlexGrow represent the CSS style "flex-grow" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/flex-grow
type FlexGrowStyle float64

func (t FlexGrowStyle) Name() string {
	return "flex-grow"
}

var FlexGrowStyleBrowserVariantsList = []string{}

func (t FlexGrowStyle) BrowserVariants() []string {
	return FlexGrowStyleBrowserVariantsList
}

var FlexGrowStyleUtilitiesMapping = map[string]string{
	"flex-grow": "flex-grow: 0",
}

func (t FlexGrowStyle) Utilities() map[string]string {
	return FlexGrowStyleUtilitiesMapping
}

func (t FlexGrowStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FlexGrowStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ShapeImageThreshold represent the CSS style "shape-image-threshold" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/shape-image-threshold
type ShapeImageThresholdStyle float64

func (t ShapeImageThresholdStyle) Name() string {
	return "shape-image-threshold"
}

var ShapeImageThresholdStyleBrowserVariantsList = []string{}

func (t ShapeImageThresholdStyle) BrowserVariants() []string {
	return ShapeImageThresholdStyleBrowserVariantsList
}

var ShapeImageThresholdStyleUtilitiesMapping = map[string]string{
	"shape-image-threshold": "shape-image-threshold: 0.0",
}

func (t ShapeImageThresholdStyle) Utilities() map[string]string {
	return ShapeImageThresholdStyleUtilitiesMapping
}

func (t ShapeImageThresholdStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ShapeImageThresholdStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformBox represent the CSS style "transform-box" with value "content-box | border-box | fill-box | stroke-box | view-box"
// See https://developer.mozilla.org/docs/Web/CSS/transform-box
type TransformBoxStyle string

func (t TransformBoxStyle) Name() string {
	return "transform-box"
}

const TransformBoxStyleContentBox TransformBoxStyle = "content-box"

const TransformBoxStyleBorderBox TransformBoxStyle = "border-box"

const TransformBoxStyleFillBox TransformBoxStyle = "fill-box"

const TransformBoxStyleStrokeBox TransformBoxStyle = "stroke-box"

const TransformBoxStyleViewBox TransformBoxStyle = "view-box"

var TransformBoxStyleBrowserVariantsList = []string{}

func (t TransformBoxStyle) BrowserVariants() []string {
	return TransformBoxStyleBrowserVariantsList
}

var TransformBoxStyleUtilitiesMapping = map[string]string{
	"transform-box":             "transform-box: view-box",
	"transform-box-content-box": "transform-box: content-box",
	"transform-box-border-box":  "transform-box: border-box",
	"transform-box-fill-box":    "transform-box: fill-box",
	"transform-box-stroke-box":  "transform-box: stroke-box",
	"transform-box-view-box":    "transform-box: view-box",
}

func (t TransformBoxStyle) Utilities() map[string]string {
	return TransformBoxStyleUtilitiesMapping
}

func (t TransformBoxStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformBoxStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBlockEndStyle represent the CSS style "border-block-end-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-end-style
type BorderBlockEndStyleStyle string

func (t BorderBlockEndStyleStyle) Name() string {
	return "border-block-end-style"
}

var BorderBlockEndStyleStyleBrowserVariantsList = []string{}

func (t BorderBlockEndStyleStyle) BrowserVariants() []string {
	return BorderBlockEndStyleStyleBrowserVariantsList
}

var BorderBlockEndStyleStyleUtilitiesMapping = map[string]string{
	"border-block-end-style": "border-block-end-style: none",
}

func (t BorderBlockEndStyleStyle) Utilities() map[string]string {
	return BorderBlockEndStyleStyleUtilitiesMapping
}

func (t BorderBlockEndStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBlockEndStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Clip represent the CSS style "clip" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/clip
type ClipStyle string

func (t ClipStyle) Name() string {
	return "clip"
}

var ClipStyleBrowserVariantsList = []string{}

func (t ClipStyle) BrowserVariants() []string {
	return ClipStyleBrowserVariantsList
}

var ClipStyleUtilitiesMapping = map[string]string{
	"clip": "clip: auto",
}

func (t ClipStyle) Utilities() map[string]string {
	return ClipStyleUtilitiesMapping
}

func (t ClipStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ClipStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextCombineUpright represent the CSS style "text-combine-upright" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-combine-upright
type TextCombineUprightStyle string

func (t TextCombineUprightStyle) Name() string {
	return "text-combine-upright"
}

var TextCombineUprightStyleBrowserVariantsList = []string{}

func (t TextCombineUprightStyle) BrowserVariants() []string {
	return TextCombineUprightStyleBrowserVariantsList
}

var TextCombineUprightStyleUtilitiesMapping = map[string]string{
	"text-combine-upright": "text-combine-upright: none",
}

func (t TextCombineUprightStyle) Utilities() map[string]string {
	return TextCombineUprightStyleUtilitiesMapping
}

func (t TextCombineUprightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextCombineUprightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextRendering represent the CSS style "text-rendering" with value "auto | optimizeSpeed | optimizeLegibility | geometricPrecision"
// See https://developer.mozilla.org/docs/Web/CSS/text-rendering
type TextRenderingStyle string

func (t TextRenderingStyle) Name() string {
	return "text-rendering"
}

const TextRenderingStyleAuto TextRenderingStyle = "auto"

const TextRenderingStyleOptimizeSpeed TextRenderingStyle = "optimizeSpeed"

const TextRenderingStyleOptimizeLegibility TextRenderingStyle = "optimizeLegibility"

const TextRenderingStyleGeometricPrecision TextRenderingStyle = "geometricPrecision"

var TextRenderingStyleBrowserVariantsList = []string{}

func (t TextRenderingStyle) BrowserVariants() []string {
	return TextRenderingStyleBrowserVariantsList
}

var TextRenderingStyleUtilitiesMapping = map[string]string{
	"text-rendering":                    "text-rendering: auto",
	"text-rendering-auto":               "text-rendering: auto",
	"text-rendering-optimizeSpeed":      "text-rendering: optimizeSpeed",
	"text-rendering-optimizeLegibility": "text-rendering: optimizeLegibility",
	"text-rendering-geometricPrecision": "text-rendering: geometricPrecision",
}

func (t TextRenderingStyle) Utilities() map[string]string {
	return TextRenderingStyleUtilitiesMapping
}

func (t TextRenderingStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextRenderingStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontFamily represent the CSS style "font-family" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-family
type FontFamilyStyle string

func (t FontFamilyStyle) Name() string {
	return "font-family"
}

var FontFamilyStyleBrowserVariantsList = []string{}

func (t FontFamilyStyle) BrowserVariants() []string {
	return FontFamilyStyleBrowserVariantsList
}

var FontFamilyStyleUtilitiesMapping = map[string]string{
	"font-family": "font-family: dependsOnUserAgent",
}

func (t FontFamilyStyle) Utilities() map[string]string {
	return FontFamilyStyleUtilitiesMapping
}

func (t FontFamilyStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontFamilyStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ForcedColorAdjust represent the CSS style "forced-color-adjust" with value "auto | none"
// See https://developer.mozilla.org/docs/Web/CSS/forced-color-adjust
type ForcedColorAdjustStyle string

func (t ForcedColorAdjustStyle) Name() string {
	return "forced-color-adjust"
}

const ForcedColorAdjustStyleAuto ForcedColorAdjustStyle = "auto"

const ForcedColorAdjustStyleNone ForcedColorAdjustStyle = "none"

var ForcedColorAdjustStyleBrowserVariantsList = []string{}

func (t ForcedColorAdjustStyle) BrowserVariants() []string {
	return ForcedColorAdjustStyleBrowserVariantsList
}

var ForcedColorAdjustStyleUtilitiesMapping = map[string]string{
	"forced-color-adjust":      "forced-color-adjust: auto",
	"forced-color-adjust-auto": "forced-color-adjust: auto",
	"forced-color-adjust-none": "forced-color-adjust: none",
}

func (t ForcedColorAdjustStyle) Utilities() map[string]string {
	return ForcedColorAdjustStyleUtilitiesMapping
}

func (t ForcedColorAdjustStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ForcedColorAdjustStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverflowBlock represent the CSS style "overflow-block" with value "visible | hidden | clip | scroll | auto"
// See
type OverflowBlockStyle string

func (t OverflowBlockStyle) Name() string {
	return "overflow-block"
}

const OverflowBlockStyleVisible OverflowBlockStyle = "visible"

const OverflowBlockStyleHidden OverflowBlockStyle = "hidden"

const OverflowBlockStyleClip OverflowBlockStyle = "clip"

const OverflowBlockStyleScroll OverflowBlockStyle = "scroll"

const OverflowBlockStyleAuto OverflowBlockStyle = "auto"

var OverflowBlockStyleBrowserVariantsList = []string{}

func (t OverflowBlockStyle) BrowserVariants() []string {
	return OverflowBlockStyleBrowserVariantsList
}

var OverflowBlockStyleUtilitiesMapping = map[string]string{
	"overflow-block":         "overflow-block: auto",
	"overflow-block-visible": "overflow-block: visible",
	"overflow-block-hidden":  "overflow-block: hidden",
	"overflow-block-clip":    "overflow-block: clip",
	"overflow-block-scroll":  "overflow-block: scroll",
	"overflow-block-auto":    "overflow-block: auto",
}

func (t OverflowBlockStyle) Utilities() map[string]string {
	return OverflowBlockStyleUtilitiesMapping
}

func (t OverflowBlockStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverflowBlockStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMarginLeft represent the CSS style "scroll-margin-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-left
type ScrollMarginLeftStyle float64

func (t ScrollMarginLeftStyle) Name() string {
	return "scroll-margin-left"
}

var ScrollMarginLeftStyleBrowserVariantsList = []string{}

func (t ScrollMarginLeftStyle) BrowserVariants() []string {
	return ScrollMarginLeftStyleBrowserVariantsList
}

var ScrollMarginLeftStyleUtilitiesMapping = map[string]string{
	"scroll-margin-left": "scroll-margin-left: 0",
}

func (t ScrollMarginLeftStyle) Utilities() map[string]string {
	return ScrollMarginLeftStyleUtilitiesMapping
}

func (t ScrollMarginLeftStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginLeftStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AlignItems represent the CSS style "align-items" with value "auto | normal | stretch | end | start | flex-start | flex-end | center |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe | unsafe"
// See https://developer.mozilla.org/docs/Web/CSS/align-items
type AlignItemsStyle string

func (t AlignItemsStyle) Name() string {
	return "align-items"
}

const AlignItemsStyleAuto AlignItemsStyle = "auto"

const AlignItemsStyleNormal AlignItemsStyle = "normal"

const AlignItemsStyleStretch AlignItemsStyle = "stretch"

const AlignItemsStyleEnd AlignItemsStyle = "end"

const AlignItemsStyleStart AlignItemsStyle = "start"

const AlignItemsStyleFlexStart AlignItemsStyle = "flex-start"

const AlignItemsStyleFlexEnd AlignItemsStyle = "flex-end"

const AlignItemsStyleCenter AlignItemsStyle = "center"

const AlignItemsStyleBaseline AlignItemsStyle = "baseline"

const AlignItemsStyleFirstBaseline AlignItemsStyle = "first-baseline"

const AlignItemsStyleLastBaseline AlignItemsStyle = "last-baseline"

const AlignItemsStyleSpaceBetween AlignItemsStyle = "space-between"

const AlignItemsStyleSpaceAround AlignItemsStyle = "space-around"

const AlignItemsStyleSpaceEvenly AlignItemsStyle = "space-evenly"

const AlignItemsStyleSafe AlignItemsStyle = "safe"

const AlignItemsStyleUnsafe AlignItemsStyle = "unsafe"

var AlignItemsStyleBrowserVariantsList = []string{}

func (t AlignItemsStyle) BrowserVariants() []string {
	return AlignItemsStyleBrowserVariantsList
}

var AlignItemsStyleUtilitiesMapping = map[string]string{
	"align-items":                "align-items: normal",
	"align-items-auto":           "align-items: auto",
	"align-items-normal":         "align-items: normal",
	"align-items-stretch":        "align-items: stretch",
	"align-items-end":            "align-items: end",
	"align-items-start":          "align-items: start",
	"align-items-flex-start":     "align-items: flex-start",
	"align-items-flex-end":       "align-items: flex-end",
	"align-items-center":         "align-items: center",
	"align-items-baseline":       "align-items: baseline",
	"align-items-first-baseline": "align-items: first-baseline",
	"align-items-last-baseline":  "align-items: last-baseline",
	"align-items-space-between":  "align-items: space-between",
	"align-items-space-around":   "align-items: space-around",
	"align-items-space-evenly":   "align-items: space-evenly",
	"align-items-safe":           "align-items: safe",
	"align-items-unsafe":         "align-items: unsafe",
}

func (t AlignItemsStyle) Utilities() map[string]string {
	return AlignItemsStyleUtilitiesMapping
}

func (t AlignItemsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AlignItemsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextUnderlineOffset represent the CSS style "text-underline-offset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-underline-offset
type TextUnderlineOffsetStyle string

func (t TextUnderlineOffsetStyle) Name() string {
	return "text-underline-offset"
}

var TextUnderlineOffsetStyleBrowserVariantsList = []string{}

func (t TextUnderlineOffsetStyle) BrowserVariants() []string {
	return TextUnderlineOffsetStyleBrowserVariantsList
}

var TextUnderlineOffsetStyleUtilitiesMapping = map[string]string{
	"text-underline-offset": "text-underline-offset: auto",
}

func (t TextUnderlineOffsetStyle) Utilities() map[string]string {
	return TextUnderlineOffsetStyleUtilitiesMapping
}

func (t TextUnderlineOffsetStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextUnderlineOffsetStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollSnapDestination represent the CSS style "scroll-snap-destination" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-destination
type ScrollSnapDestinationStyle string

func (t ScrollSnapDestinationStyle) Name() string {
	return "scroll-snap-destination"
}

var ScrollSnapDestinationStyleBrowserVariantsList = []string{}

func (t ScrollSnapDestinationStyle) BrowserVariants() []string {
	return ScrollSnapDestinationStyleBrowserVariantsList
}

var ScrollSnapDestinationStyleUtilitiesMapping = map[string]string{
	"scroll-snap-destination": "scroll-snap-destination: 0px 0px",
}

func (t ScrollSnapDestinationStyle) Utilities() map[string]string {
	return ScrollSnapDestinationStyleUtilitiesMapping
}

func (t ScrollSnapDestinationStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollSnapDestinationStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBottomRightRadius represent the CSS style "border-bottom-right-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-right-radius
type BorderBottomRightRadiusStyle string

func (t BorderBottomRightRadiusStyle) Name() string {
	return "border-bottom-right-radius"
}

var BorderBottomRightRadiusStyleBrowserVariantsList = []string{}

func (t BorderBottomRightRadiusStyle) BrowserVariants() []string {
	return BorderBottomRightRadiusStyleBrowserVariantsList
}

var BorderBottomRightRadiusStyleUtilitiesMapping = map[string]string{
	"border-bottom-right-radius": "border-bottom-right-radius: 0",
}

func (t BorderBottomRightRadiusStyle) Utilities() map[string]string {
	return BorderBottomRightRadiusStyleUtilitiesMapping
}

func (t BorderBottomRightRadiusStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBottomRightRadiusStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextDecorationSkip represent the CSS style "text-decoration-skip" with value "none | objects | spaces | leading-spaces | trailing-spaces | edges | box-decoration"
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-skip
type TextDecorationSkipStyle string

func (t TextDecorationSkipStyle) Name() string {
	return "text-decoration-skip"
}

const TextDecorationSkipStyleNone TextDecorationSkipStyle = "none"

const TextDecorationSkipStyleObjects TextDecorationSkipStyle = "objects"

const TextDecorationSkipStyleSpaces TextDecorationSkipStyle = "spaces"

const TextDecorationSkipStyleLeadingSpaces TextDecorationSkipStyle = "leading-spaces"

const TextDecorationSkipStyleTrailingSpaces TextDecorationSkipStyle = "trailing-spaces"

const TextDecorationSkipStyleEdges TextDecorationSkipStyle = "edges"

const TextDecorationSkipStyleBoxDecoration TextDecorationSkipStyle = "box-decoration"

var TextDecorationSkipStyleBrowserVariantsList = []string{}

func (t TextDecorationSkipStyle) BrowserVariants() []string {
	return TextDecorationSkipStyleBrowserVariantsList
}

var TextDecorationSkipStyleUtilitiesMapping = map[string]string{
	"text-decoration-skip":                 "text-decoration-skip: objects",
	"text-decoration-skip-none":            "text-decoration-skip: none",
	"text-decoration-skip-objects":         "text-decoration-skip: objects",
	"text-decoration-skip-spaces":          "text-decoration-skip: spaces",
	"text-decoration-skip-leading-spaces":  "text-decoration-skip: leading-spaces",
	"text-decoration-skip-trailing-spaces": "text-decoration-skip: trailing-spaces",
	"text-decoration-skip-edges":           "text-decoration-skip: edges",
	"text-decoration-skip-box-decoration":  "text-decoration-skip: box-decoration",
}

func (t TextDecorationSkipStyle) Utilities() map[string]string {
	return TextDecorationSkipStyleUtilitiesMapping
}

func (t TextDecorationSkipStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextDecorationSkipStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderEndStartRadius represent the CSS style "border-end-start-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-end-start-radius
type BorderEndStartRadiusStyle string

func (t BorderEndStartRadiusStyle) Name() string {
	return "border-end-start-radius"
}

var BorderEndStartRadiusStyleBrowserVariantsList = []string{}

func (t BorderEndStartRadiusStyle) BrowserVariants() []string {
	return BorderEndStartRadiusStyleBrowserVariantsList
}

var BorderEndStartRadiusStyleUtilitiesMapping = map[string]string{
	"border-end-start-radius": "border-end-start-radius: 0",
}

func (t BorderEndStartRadiusStyle) Utilities() map[string]string {
	return BorderEndStartRadiusStyleUtilitiesMapping
}

func (t BorderEndStartRadiusStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderEndStartRadiusStyleUtilitiesMapping[tu]
	return value, hasValue
}

// InitialLetter represent the CSS style "initial-letter" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/initial-letter
type InitialLetterStyle string

func (t InitialLetterStyle) Name() string {
	return "initial-letter"
}

var InitialLetterStyleBrowserVariantsList = []string{}

func (t InitialLetterStyle) BrowserVariants() []string {
	return InitialLetterStyleBrowserVariantsList
}

var InitialLetterStyleUtilitiesMapping = map[string]string{
	"initial-letter": "initial-letter: normal",
}

func (t InitialLetterStyle) Utilities() map[string]string {
	return InitialLetterStyleUtilitiesMapping
}

func (t InitialLetterStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = InitialLetterStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverscrollBehavior represent the CSS style "overscroll-behavior" with value "contain | none | auto | contain none | contain auto | none contain | none auto | auto contain | auto none"
// See https://developer.mozilla.org/docs/Web/CSS/overscroll-behavior
type OverscrollBehaviorStyle string

func (t OverscrollBehaviorStyle) Name() string {
	return "overscroll-behavior"
}

const OverscrollBehaviorStyleContain OverscrollBehaviorStyle = "contain"

const OverscrollBehaviorStyleNone OverscrollBehaviorStyle = "none"

const OverscrollBehaviorStyleAuto OverscrollBehaviorStyle = "auto"

const OverscrollBehaviorStyleContainNone OverscrollBehaviorStyle = "contain-none"

const OverscrollBehaviorStyleContainAuto OverscrollBehaviorStyle = "contain-auto"

const OverscrollBehaviorStyleNoneContain OverscrollBehaviorStyle = "none-contain"

const OverscrollBehaviorStyleNoneAuto OverscrollBehaviorStyle = "none-auto"

const OverscrollBehaviorStyleAutoContain OverscrollBehaviorStyle = "auto-contain"

const OverscrollBehaviorStyleAutoNone OverscrollBehaviorStyle = "auto-none"

var OverscrollBehaviorStyleBrowserVariantsList = []string{}

func (t OverscrollBehaviorStyle) BrowserVariants() []string {
	return OverscrollBehaviorStyleBrowserVariantsList
}

var OverscrollBehaviorStyleUtilitiesMapping = map[string]string{
	"overscroll-behavior":              "overscroll-behavior: auto",
	"overscroll-behavior-contain":      "overscroll-behavior: contain",
	"overscroll-behavior-none":         "overscroll-behavior: none",
	"overscroll-behavior-auto":         "overscroll-behavior: auto",
	"overscroll-behavior-contain-none": "overscroll-behavior: contain-none",
	"overscroll-behavior-contain-auto": "overscroll-behavior: contain-auto",
	"overscroll-behavior-none-contain": "overscroll-behavior: none-contain",
	"overscroll-behavior-none-auto":    "overscroll-behavior: none-auto",
	"overscroll-behavior-auto-contain": "overscroll-behavior: auto-contain",
	"overscroll-behavior-auto-none":    "overscroll-behavior: auto-none",
}

func (t OverscrollBehaviorStyle) Utilities() map[string]string {
	return OverscrollBehaviorStyleUtilitiesMapping
}

func (t OverscrollBehaviorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverscrollBehaviorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformTranslate3d represent the CSS style "transform-translate-3d" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformTranslate3dStyle string

func (t TransformTranslate3dStyle) Name() string {
	return "transform-translate-3d"
}

var TransformTranslate3dStyleBrowserVariantsList = []string{}

func (t TransformTranslate3dStyle) BrowserVariants() []string {
	return TransformTranslate3dStyleBrowserVariantsList
}

var TransformTranslate3dStyleUtilitiesMapping = map[string]string{
	"transform-translate-3d": "transform-translate-3d: translate3d(0, 0, 0)",
}

func (t TransformTranslate3dStyle) Utilities() map[string]string {
	return TransformTranslate3dStyleUtilitiesMapping
}

func (t TransformTranslate3dStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformTranslate3dStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderTopStyle represent the CSS style "border-top-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-style
type BorderTopStyleStyle string

func (t BorderTopStyleStyle) Name() string {
	return "border-top-style"
}

var BorderTopStyleStyleBrowserVariantsList = []string{}

func (t BorderTopStyleStyle) BrowserVariants() []string {
	return BorderTopStyleStyleBrowserVariantsList
}

var BorderTopStyleStyleUtilitiesMapping = map[string]string{
	"border-top-style": "border-top-style: none",
}

func (t BorderTopStyleStyle) Utilities() map[string]string {
	return BorderTopStyleStyleUtilitiesMapping
}

func (t BorderTopStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderTopStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// UserSelect represent the CSS style "user-select" with value "auto | text | none | contain | all"
// See https://developer.mozilla.org/docs/Web/CSS/user-select
type UserSelectStyle string

func (t UserSelectStyle) Name() string {
	return "user-select"
}

const UserSelectStyleAuto UserSelectStyle = "auto"

const UserSelectStyleText UserSelectStyle = "text"

const UserSelectStyleNone UserSelectStyle = "none"

const UserSelectStyleContain UserSelectStyle = "contain"

const UserSelectStyleAll UserSelectStyle = "all"

var UserSelectStyleBrowserVariantsList = []string{
	"-ms-user-select",
}

func (t UserSelectStyle) BrowserVariants() []string {
	return UserSelectStyleBrowserVariantsList
}

var UserSelectStyleUtilitiesMapping = map[string]string{
	"user-select":         "user-select: auto",
	"user-select-auto":    "user-select: auto",
	"user-select-text":    "user-select: text",
	"user-select-none":    "user-select: none",
	"user-select-contain": "user-select: contain",
	"user-select-all":     "user-select: all",
}

func (t UserSelectStyle) Utilities() map[string]string {
	return UserSelectStyleUtilitiesMapping
}

func (t UserSelectStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = UserSelectStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Direction represent the CSS style "direction" with value "ltr | rtl"
// See https://developer.mozilla.org/docs/Web/CSS/direction
type DirectionStyle string

func (t DirectionStyle) Name() string {
	return "direction"
}

const DirectionStyleLtr DirectionStyle = "ltr"

const DirectionStyleRtl DirectionStyle = "rtl"

var DirectionStyleBrowserVariantsList = []string{}

func (t DirectionStyle) BrowserVariants() []string {
	return DirectionStyleBrowserVariantsList
}

var DirectionStyleUtilitiesMapping = map[string]string{
	"direction":     "direction: ltr",
	"direction-ltr": "direction: ltr",
	"direction-rtl": "direction: rtl",
}

func (t DirectionStyle) Utilities() map[string]string {
	return DirectionStyleUtilitiesMapping
}

func (t DirectionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = DirectionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// InsetInlineEnd represent the CSS style "inset-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-inline-end
type InsetInlineEndStyle string

func (t InsetInlineEndStyle) Name() string {
	return "inset-inline-end"
}

var InsetInlineEndStyleBrowserVariantsList = []string{}

func (t InsetInlineEndStyle) BrowserVariants() []string {
	return InsetInlineEndStyleBrowserVariantsList
}

var InsetInlineEndStyleUtilitiesMapping = map[string]string{
	"inset-inline-end": "inset-inline-end: auto",
}

func (t InsetInlineEndStyle) Utilities() map[string]string {
	return InsetInlineEndStyleUtilitiesMapping
}

func (t InsetInlineEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = InsetInlineEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// InsetInline represent the CSS style "inset-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-inline
type InsetInlineStyle string

func (t InsetInlineStyle) Name() string {
	return "inset-inline"
}

var InsetInlineStyleBrowserVariantsList = []string{}

func (t InsetInlineStyle) BrowserVariants() []string {
	return InsetInlineStyleBrowserVariantsList
}

var InsetInlineStyleUtilitiesMapping = map[string]string{
	"inset-inline": "inset-inline: auto",
}

func (t InsetInlineStyle) Utilities() map[string]string {
	return InsetInlineStyleUtilitiesMapping
}

func (t InsetInlineStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = InsetInlineStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundBlendMode represent the CSS style "background-blend-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-blend-mode
type BackgroundBlendModeStyle string

func (t BackgroundBlendModeStyle) Name() string {
	return "background-blend-mode"
}

var BackgroundBlendModeStyleBrowserVariantsList = []string{}

func (t BackgroundBlendModeStyle) BrowserVariants() []string {
	return BackgroundBlendModeStyleBrowserVariantsList
}

var BackgroundBlendModeStyleUtilitiesMapping = map[string]string{
	"background-blend-mode": "background-blend-mode: normal",
}

func (t BackgroundBlendModeStyle) Utilities() map[string]string {
	return BackgroundBlendModeStyleUtilitiesMapping
}

func (t BackgroundBlendModeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundBlendModeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransitionProperty represent the CSS style "transition-property" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-property
type TransitionPropertyStyle string

func (t TransitionPropertyStyle) Name() string {
	return "transition-property"
}

var TransitionPropertyStyleBrowserVariantsList = []string{}

func (t TransitionPropertyStyle) BrowserVariants() []string {
	return TransitionPropertyStyleBrowserVariantsList
}

var TransitionPropertyStyleUtilitiesMapping = map[string]string{
	"transition-property": "transition-property: all",
}

func (t TransitionPropertyStyle) Utilities() map[string]string {
	return TransitionPropertyStyleUtilitiesMapping
}

func (t TransitionPropertyStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransitionPropertyStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverscrollBehaviorInline represent the CSS style "overscroll-behavior-inline" with value "contain | none | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overscroll-behavior-inline
type OverscrollBehaviorInlineStyle string

func (t OverscrollBehaviorInlineStyle) Name() string {
	return "overscroll-behavior-inline"
}

const OverscrollBehaviorInlineStyleContain OverscrollBehaviorInlineStyle = "contain"

const OverscrollBehaviorInlineStyleNone OverscrollBehaviorInlineStyle = "none"

const OverscrollBehaviorInlineStyleAuto OverscrollBehaviorInlineStyle = "auto"

var OverscrollBehaviorInlineStyleBrowserVariantsList = []string{}

func (t OverscrollBehaviorInlineStyle) BrowserVariants() []string {
	return OverscrollBehaviorInlineStyleBrowserVariantsList
}

var OverscrollBehaviorInlineStyleUtilitiesMapping = map[string]string{
	"overscroll-behavior-inline":         "overscroll-behavior-inline: auto",
	"overscroll-behavior-inline-contain": "overscroll-behavior-inline: contain",
	"overscroll-behavior-inline-none":    "overscroll-behavior-inline: none",
	"overscroll-behavior-inline-auto":    "overscroll-behavior-inline: auto",
}

func (t OverscrollBehaviorInlineStyle) Utilities() map[string]string {
	return OverscrollBehaviorInlineStyleUtilitiesMapping
}

func (t OverscrollBehaviorInlineStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverscrollBehaviorInlineStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMarginBlock represent the CSS style "scroll-margin-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-block
type ScrollMarginBlockStyle float64

func (t ScrollMarginBlockStyle) Name() string {
	return "scroll-margin-block"
}

var ScrollMarginBlockStyleBrowserVariantsList = []string{}

func (t ScrollMarginBlockStyle) BrowserVariants() []string {
	return ScrollMarginBlockStyleBrowserVariantsList
}

var ScrollMarginBlockStyleUtilitiesMapping = map[string]string{
	"scroll-margin-block": "scroll-margin-block: 0",
}

func (t ScrollMarginBlockStyle) Utilities() map[string]string {
	return ScrollMarginBlockStyleUtilitiesMapping
}

func (t ScrollMarginBlockStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginBlockStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformTranslate represent the CSS style "transform-translate" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformTranslateStyle string

func (t TransformTranslateStyle) Name() string {
	return "transform-translate"
}

var TransformTranslateStyleBrowserVariantsList = []string{}

func (t TransformTranslateStyle) BrowserVariants() []string {
	return TransformTranslateStyleBrowserVariantsList
}

var TransformTranslateStyleUtilitiesMapping = map[string]string{
	"transform-translate": "transform-translate: translate(0)",
}

func (t TransformTranslateStyle) Utilities() map[string]string {
	return TransformTranslateStyleUtilitiesMapping
}

func (t TransformTranslateStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformTranslateStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Marker represent the CSS style "marker" with value ""
// See
type MarkerStyle string

func (t MarkerStyle) Name() string {
	return "marker"
}

var MarkerStyleBrowserVariantsList = []string{}

func (t MarkerStyle) BrowserVariants() []string {
	return MarkerStyleBrowserVariantsList
}

var MarkerStyleUtilitiesMapping = map[string]string{
	"marker": "marker: none",
}

func (t MarkerStyle) Utilities() map[string]string {
	return MarkerStyleUtilitiesMapping
}

func (t MarkerStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarkerStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderImageWidth represent the CSS style "border-image-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-width
type BorderImageWidthStyle string

func (t BorderImageWidthStyle) Name() string {
	return "border-image-width"
}

var BorderImageWidthStyleBrowserVariantsList = []string{}

func (t BorderImageWidthStyle) BrowserVariants() []string {
	return BorderImageWidthStyleBrowserVariantsList
}

var BorderImageWidthStyleUtilitiesMapping = map[string]string{
	"border-image-width": "border-image-width: 1",
}

func (t BorderImageWidthStyle) Utilities() map[string]string {
	return BorderImageWidthStyleUtilitiesMapping
}

func (t BorderImageWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderImageWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OutlineStyle represent the CSS style "outline-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-style
type OutlineStyleStyle string

func (t OutlineStyleStyle) Name() string {
	return "outline-style"
}

var OutlineStyleStyleBrowserVariantsList = []string{}

func (t OutlineStyleStyle) BrowserVariants() []string {
	return OutlineStyleStyleBrowserVariantsList
}

var OutlineStyleStyleUtilitiesMapping = map[string]string{
	"outline-style": "outline-style: none",
}

func (t OutlineStyleStyle) Utilities() map[string]string {
	return OutlineStyleStyleUtilitiesMapping
}

func (t OutlineStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OutlineStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverflowClipBox represent the CSS style "overflow-clip-box" with value "padding-box | content-box"
// See https://developer.mozilla.org/docs/Mozilla/CSS/overflow-clip-box
type OverflowClipBoxStyle string

func (t OverflowClipBoxStyle) Name() string {
	return "overflow-clip-box"
}

const OverflowClipBoxStylePaddingBox OverflowClipBoxStyle = "padding-box"

const OverflowClipBoxStyleContentBox OverflowClipBoxStyle = "content-box"

var OverflowClipBoxStyleBrowserVariantsList = []string{}

func (t OverflowClipBoxStyle) BrowserVariants() []string {
	return OverflowClipBoxStyleBrowserVariantsList
}

var OverflowClipBoxStyleUtilitiesMapping = map[string]string{
	"overflow-clip-box":             "overflow-clip-box: padding-box",
	"overflow-clip-box-padding-box": "overflow-clip-box: padding-box",
	"overflow-clip-box-content-box": "overflow-clip-box: content-box",
}

func (t OverflowClipBoxStyle) Utilities() map[string]string {
	return OverflowClipBoxStyleUtilitiesMapping
}

func (t OverflowClipBoxStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverflowClipBoxStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPaddingTop represent the CSS style "scroll-padding-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-top
type ScrollPaddingTopStyle string

func (t ScrollPaddingTopStyle) Name() string {
	return "scroll-padding-top"
}

var ScrollPaddingTopStyleBrowserVariantsList = []string{}

func (t ScrollPaddingTopStyle) BrowserVariants() []string {
	return ScrollPaddingTopStyleBrowserVariantsList
}

var ScrollPaddingTopStyleUtilitiesMapping = map[string]string{
	"scroll-padding-top": "scroll-padding-top: auto",
}

func (t ScrollPaddingTopStyle) Utilities() map[string]string {
	return ScrollPaddingTopStyleUtilitiesMapping
}

func (t ScrollPaddingTopStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingTopStyleUtilitiesMapping[tu]
	return value, hasValue
}

// StrokeWidth represent the CSS style "stroke-width" with value ""
// See
type StrokeWidthStyle float64

func (t StrokeWidthStyle) Name() string {
	return "stroke-width"
}

var StrokeWidthStyleBrowserVariantsList = []string{}

func (t StrokeWidthStyle) BrowserVariants() []string {
	return StrokeWidthStyleBrowserVariantsList
}

var StrokeWidthStyleUtilitiesMapping = map[string]string{
	"stroke-width": "stroke-width: 1",
}

func (t StrokeWidthStyle) Utilities() map[string]string {
	return StrokeWidthStyleUtilitiesMapping
}

func (t StrokeWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = StrokeWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ListStyleImage represent the CSS style "list-style-image" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/list-style-image
type ListStyleImageStyle URL

func (t ListStyleImageStyle) Name() string {
	return "list-style-image"
}

var ListStyleImageStyleBrowserVariantsList = []string{}

func (t ListStyleImageStyle) BrowserVariants() []string {
	return ListStyleImageStyleBrowserVariantsList
}

var ListStyleImageStyleUtilitiesMapping = map[string]string{
	"list-style-image": "list-style-image: none",
}

func (t ListStyleImageStyle) Utilities() map[string]string {
	return ListStyleImageStyleUtilitiesMapping
}

func (t ListStyleImageStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ListStyleImageStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPaddingBlockStart represent the CSS style "scroll-padding-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-block-start
type ScrollPaddingBlockStartStyle string

func (t ScrollPaddingBlockStartStyle) Name() string {
	return "scroll-padding-block-start"
}

var ScrollPaddingBlockStartStyleBrowserVariantsList = []string{}

func (t ScrollPaddingBlockStartStyle) BrowserVariants() []string {
	return ScrollPaddingBlockStartStyleBrowserVariantsList
}

var ScrollPaddingBlockStartStyleUtilitiesMapping = map[string]string{
	"scroll-padding-block-start": "scroll-padding-block-start: auto",
}

func (t ScrollPaddingBlockStartStyle) Utilities() map[string]string {
	return ScrollPaddingBlockStartStyleUtilitiesMapping
}

func (t ScrollPaddingBlockStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingBlockStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderImageSlice represent the CSS style "border-image-slice" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-slice
type BorderImageSliceStyle string

func (t BorderImageSliceStyle) Name() string {
	return "border-image-slice"
}

var BorderImageSliceStyleBrowserVariantsList = []string{}

func (t BorderImageSliceStyle) BrowserVariants() []string {
	return BorderImageSliceStyleBrowserVariantsList
}

var BorderImageSliceStyleUtilitiesMapping = map[string]string{
	"border-image-slice": "border-image-slice: 100%",
}

func (t BorderImageSliceStyle) Utilities() map[string]string {
	return BorderImageSliceStyleUtilitiesMapping
}

func (t BorderImageSliceStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderImageSliceStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderTopColor represent the CSS style "border-top-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-color
type BorderTopColorStyle Color

func (t BorderTopColorStyle) Name() string {
	return "border-top-color"
}

var BorderTopColorStyleBrowserVariantsList = []string{}

func (t BorderTopColorStyle) BrowserVariants() []string {
	return BorderTopColorStyleBrowserVariantsList
}

var BorderTopColorStyleUtilitiesMapping = map[string]string{
	"border-top-color": "border-top-color: currentcolor",
}

func (t BorderTopColorStyle) Utilities() map[string]string {
	return BorderTopColorStyleUtilitiesMapping
}

func (t BorderTopColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderTopColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformScale3d represent the CSS style "transform-scale-3d" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformScale3dStyle string

func (t TransformScale3dStyle) Name() string {
	return "transform-scale-3d"
}

var TransformScale3dStyleBrowserVariantsList = []string{}

func (t TransformScale3dStyle) BrowserVariants() []string {
	return TransformScale3dStyleBrowserVariantsList
}

var TransformScale3dStyleUtilitiesMapping = map[string]string{
	"transform-scale-3d": "transform-scale-3d: scale3d(1, 1, 1)",
}

func (t TransformScale3dStyle) Utilities() map[string]string {
	return TransformScale3dStyleUtilitiesMapping
}

func (t TransformScale3dStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformScale3dStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AspectRatio represent the CSS style "aspect-ratio" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/aspect-ratio
type AspectRatioStyle string

func (t AspectRatioStyle) Name() string {
	return "aspect-ratio"
}

var AspectRatioStyleBrowserVariantsList = []string{}

func (t AspectRatioStyle) BrowserVariants() []string {
	return AspectRatioStyleBrowserVariantsList
}

var AspectRatioStyleUtilitiesMapping = map[string]string{
	"aspect-ratio": "aspect-ratio: auto",
}

func (t AspectRatioStyle) Utilities() map[string]string {
	return AspectRatioStyleUtilitiesMapping
}

func (t AspectRatioStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AspectRatioStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Position represent the CSS style "position" with value "static | relative | absolute | sticky | fixed"
// See https://developer.mozilla.org/docs/Web/CSS/position
type PositionStyle string

func (t PositionStyle) Name() string {
	return "position"
}

const PositionStyleStatic PositionStyle = "static"

const PositionStyleRelative PositionStyle = "relative"

const PositionStyleAbsolute PositionStyle = "absolute"

const PositionStyleSticky PositionStyle = "sticky"

const PositionStyleFixed PositionStyle = "fixed"

var PositionStyleBrowserVariantsList = []string{}

func (t PositionStyle) BrowserVariants() []string {
	return PositionStyleBrowserVariantsList
}

var PositionStyleUtilitiesMapping = map[string]string{
	"position":          "position: static",
	"position-static":   "position: static",
	"position-relative": "position: relative",
	"position-absolute": "position: absolute",
	"position-sticky":   "position: sticky",
	"position-fixed":    "position: fixed",
}

func (t PositionStyle) Utilities() map[string]string {
	return PositionStyleUtilitiesMapping
}

func (t PositionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PositionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskBorderOutset represent the CSS style "mask-border-outset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-outset
type MaskBorderOutsetStyle string

func (t MaskBorderOutsetStyle) Name() string {
	return "mask-border-outset"
}

var MaskBorderOutsetStyleBrowserVariantsList = []string{}

func (t MaskBorderOutsetStyle) BrowserVariants() []string {
	return MaskBorderOutsetStyleBrowserVariantsList
}

var MaskBorderOutsetStyleUtilitiesMapping = map[string]string{
	"mask-border-outset": "mask-border-outset: 0",
}

func (t MaskBorderOutsetStyle) Utilities() map[string]string {
	return MaskBorderOutsetStyleUtilitiesMapping
}

func (t MaskBorderOutsetStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskBorderOutsetStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ClipRule represent the CSS style "clip-rule" with value "nonzero | evenodd"
// See
type ClipRuleStyle string

func (t ClipRuleStyle) Name() string {
	return "clip-rule"
}

const ClipRuleStyleNonzero ClipRuleStyle = "nonzero"

const ClipRuleStyleEvenodd ClipRuleStyle = "evenodd"

var ClipRuleStyleBrowserVariantsList = []string{}

func (t ClipRuleStyle) BrowserVariants() []string {
	return ClipRuleStyleBrowserVariantsList
}

var ClipRuleStyleUtilitiesMapping = map[string]string{
	"clip-rule":         "clip-rule: nonzero",
	"clip-rule-nonzero": "clip-rule: nonzero",
	"clip-rule-evenodd": "clip-rule: evenodd",
}

func (t ClipRuleStyle) Utilities() map[string]string {
	return ClipRuleStyleUtilitiesMapping
}

func (t ClipRuleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ClipRuleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontStretch represent the CSS style "font-stretch" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-stretch
type FontStretchStyle string

func (t FontStretchStyle) Name() string {
	return "font-stretch"
}

var FontStretchStyleBrowserVariantsList = []string{}

func (t FontStretchStyle) BrowserVariants() []string {
	return FontStretchStyleBrowserVariantsList
}

var FontStretchStyleUtilitiesMapping = map[string]string{
	"font-stretch": "font-stretch: normal",
}

func (t FontStretchStyle) Utilities() map[string]string {
	return FontStretchStyleUtilitiesMapping
}

func (t FontStretchStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontStretchStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginTop represent the CSS style "margin-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-top
type MarginTopStyle string

func (t MarginTopStyle) Name() string {
	return "margin-top"
}

var MarginTopStyleBrowserVariantsList = []string{}

func (t MarginTopStyle) BrowserVariants() []string {
	return MarginTopStyleBrowserVariantsList
}

var MarginTopStyleUtilitiesMapping = map[string]string{
	"margin-top": "margin-top: 0",
}

func (t MarginTopStyle) Utilities() map[string]string {
	return MarginTopStyleUtilitiesMapping
}

func (t MarginTopStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginTopStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMarginBottom represent the CSS style "scroll-margin-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-bottom
type ScrollMarginBottomStyle float64

func (t ScrollMarginBottomStyle) Name() string {
	return "scroll-margin-bottom"
}

var ScrollMarginBottomStyleBrowserVariantsList = []string{}

func (t ScrollMarginBottomStyle) BrowserVariants() []string {
	return ScrollMarginBottomStyleBrowserVariantsList
}

var ScrollMarginBottomStyleUtilitiesMapping = map[string]string{
	"scroll-margin-bottom": "scroll-margin-bottom: 0",
}

func (t ScrollMarginBottomStyle) Utilities() map[string]string {
	return ScrollMarginBottomStyleUtilitiesMapping
}

func (t ScrollMarginBottomStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginBottomStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AnimationDelay represent the CSS style "animation-delay" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-delay
type AnimationDelayStyle time.Duration

func (t AnimationDelayStyle) Name() string {
	return "animation-delay"
}

var AnimationDelayStyleBrowserVariantsList = []string{}

func (t AnimationDelayStyle) BrowserVariants() []string {
	return AnimationDelayStyleBrowserVariantsList
}

var AnimationDelayStyleUtilitiesMapping = map[string]string{
	"animation-delay": "animation-delay: 0s",
}

func (t AnimationDelayStyle) Utilities() map[string]string {
	return AnimationDelayStyleUtilitiesMapping
}

func (t AnimationDelayStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AnimationDelayStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BlockSize represent the CSS style "block-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/block-size
type BlockSizeStyle string

func (t BlockSizeStyle) Name() string {
	return "block-size"
}

var BlockSizeStyleBrowserVariantsList = []string{}

func (t BlockSizeStyle) BrowserVariants() []string {
	return BlockSizeStyleBrowserVariantsList
}

var BlockSizeStyleUtilitiesMapping = map[string]string{
	"block-size": "block-size: auto",
}

func (t BlockSizeStyle) Utilities() map[string]string {
	return BlockSizeStyleUtilitiesMapping
}

func (t BlockSizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BlockSizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontOpticalSizing represent the CSS style "font-optical-sizing" with value "auto | none"
// See https://developer.mozilla.org/docs/Web/CSS/font-optical-sizing
type FontOpticalSizingStyle string

func (t FontOpticalSizingStyle) Name() string {
	return "font-optical-sizing"
}

const FontOpticalSizingStyleAuto FontOpticalSizingStyle = "auto"

const FontOpticalSizingStyleNone FontOpticalSizingStyle = "none"

var FontOpticalSizingStyleBrowserVariantsList = []string{}

func (t FontOpticalSizingStyle) BrowserVariants() []string {
	return FontOpticalSizingStyleBrowserVariantsList
}

var FontOpticalSizingStyleUtilitiesMapping = map[string]string{
	"font-optical-sizing":      "font-optical-sizing: auto",
	"font-optical-sizing-auto": "font-optical-sizing: auto",
	"font-optical-sizing-none": "font-optical-sizing: none",
}

func (t FontOpticalSizingStyle) Utilities() map[string]string {
	return FontOpticalSizingStyleUtilitiesMapping
}

func (t FontOpticalSizingStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontOpticalSizingStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ShapeOutside represent the CSS style "shape-outside" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/shape-outside
type ShapeOutsideStyle string

func (t ShapeOutsideStyle) Name() string {
	return "shape-outside"
}

var ShapeOutsideStyleBrowserVariantsList = []string{}

func (t ShapeOutsideStyle) BrowserVariants() []string {
	return ShapeOutsideStyleBrowserVariantsList
}

var ShapeOutsideStyleUtilitiesMapping = map[string]string{
	"shape-outside": "shape-outside: none",
}

func (t ShapeOutsideStyle) Utilities() map[string]string {
	return ShapeOutsideStyleUtilitiesMapping
}

func (t ShapeOutsideStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ShapeOutsideStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FloodOpacity represent the CSS style "flood-opacity" with value ""
// See
type FloodOpacityStyle float64

func (t FloodOpacityStyle) Name() string {
	return "flood-opacity"
}

var FloodOpacityStyleBrowserVariantsList = []string{}

func (t FloodOpacityStyle) BrowserVariants() []string {
	return FloodOpacityStyleBrowserVariantsList
}

var FloodOpacityStyleUtilitiesMapping = map[string]string{
	"flood-opacity": "flood-opacity: 1",
}

func (t FloodOpacityStyle) Utilities() map[string]string {
	return FloodOpacityStyleUtilitiesMapping
}

func (t FloodOpacityStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FloodOpacityStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaddingInline represent the CSS style "padding-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-inline
type PaddingInlineStyle string

func (t PaddingInlineStyle) Name() string {
	return "padding-inline"
}

var PaddingInlineStyleBrowserVariantsList = []string{}

func (t PaddingInlineStyle) BrowserVariants() []string {
	return PaddingInlineStyleBrowserVariantsList
}

var PaddingInlineStyleUtilitiesMapping = map[string]string{
	"padding-inline": "padding-inline: 0",
}

func (t PaddingInlineStyle) Utilities() map[string]string {
	return PaddingInlineStyleUtilitiesMapping
}

func (t PaddingInlineStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaddingInlineStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBottomWidth represent the CSS style "border-bottom-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-width
type BorderBottomWidthStyle string

func (t BorderBottomWidthStyle) Name() string {
	return "border-bottom-width"
}

var BorderBottomWidthStyleBrowserVariantsList = []string{}

func (t BorderBottomWidthStyle) BrowserVariants() []string {
	return BorderBottomWidthStyleBrowserVariantsList
}

var BorderBottomWidthStyleUtilitiesMapping = map[string]string{
	"border-bottom-width": "border-bottom-width: medium",
}

func (t BorderBottomWidthStyle) Utilities() map[string]string {
	return BorderBottomWidthStyleUtilitiesMapping
}

func (t BorderBottomWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBottomWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridTemplateRows represent the CSS style "grid-template-rows" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-template-rows
type GridTemplateRowsStyle string

func (t GridTemplateRowsStyle) Name() string {
	return "grid-template-rows"
}

var GridTemplateRowsStyleBrowserVariantsList = []string{}

func (t GridTemplateRowsStyle) BrowserVariants() []string {
	return GridTemplateRowsStyleBrowserVariantsList
}

var GridTemplateRowsStyleUtilitiesMapping = map[string]string{
	"grid-template-rows": "grid-template-rows: none",
}

func (t GridTemplateRowsStyle) Utilities() map[string]string {
	return GridTemplateRowsStyleUtilitiesMapping
}

func (t GridTemplateRowsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridTemplateRowsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OffsetAnchor represent the CSS style "offset-anchor" with value ""
// See
type OffsetAnchorStyle string

func (t OffsetAnchorStyle) Name() string {
	return "offset-anchor"
}

var OffsetAnchorStyleBrowserVariantsList = []string{}

func (t OffsetAnchorStyle) BrowserVariants() []string {
	return OffsetAnchorStyleBrowserVariantsList
}

var OffsetAnchorStyleUtilitiesMapping = map[string]string{
	"offset-anchor": "offset-anchor: auto",
}

func (t OffsetAnchorStyle) Utilities() map[string]string {
	return OffsetAnchorStyleUtilitiesMapping
}

func (t OffsetAnchorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OffsetAnchorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// VectorEffect represent the CSS style "vector-effect" with value "non-scaling-stroke | none"
// See
type VectorEffectStyle string

func (t VectorEffectStyle) Name() string {
	return "vector-effect"
}

const VectorEffectStyleNonScalingStroke VectorEffectStyle = "non-scaling-stroke"

const VectorEffectStyleNone VectorEffectStyle = "none"

var VectorEffectStyleBrowserVariantsList = []string{}

func (t VectorEffectStyle) BrowserVariants() []string {
	return VectorEffectStyleBrowserVariantsList
}

var VectorEffectStyleUtilitiesMapping = map[string]string{
	"vector-effect":                    "vector-effect: none",
	"vector-effect-non-scaling-stroke": "vector-effect: non-scaling-stroke",
	"vector-effect-none":               "vector-effect: none",
}

func (t VectorEffectStyle) Utilities() map[string]string {
	return VectorEffectStyleUtilitiesMapping
}

func (t VectorEffectStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = VectorEffectStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarkerEnd represent the CSS style "marker-end" with value ""
// See
type MarkerEndStyle string

func (t MarkerEndStyle) Name() string {
	return "marker-end"
}

var MarkerEndStyleBrowserVariantsList = []string{}

func (t MarkerEndStyle) BrowserVariants() []string {
	return MarkerEndStyleBrowserVariantsList
}

var MarkerEndStyleUtilitiesMapping = map[string]string{
	"marker-end": "marker-end: none",
}

func (t MarkerEndStyle) Utilities() map[string]string {
	return MarkerEndStyleUtilitiesMapping
}

func (t MarkerEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarkerEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarkerMid represent the CSS style "marker-mid" with value ""
// See
type MarkerMidStyle string

func (t MarkerMidStyle) Name() string {
	return "marker-mid"
}

var MarkerMidStyleBrowserVariantsList = []string{}

func (t MarkerMidStyle) BrowserVariants() []string {
	return MarkerMidStyleBrowserVariantsList
}

var MarkerMidStyleUtilitiesMapping = map[string]string{
	"marker-mid": "marker-mid: none",
}

func (t MarkerMidStyle) Utilities() map[string]string {
	return MarkerMidStyleUtilitiesMapping
}

func (t MarkerMidStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarkerMidStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBlockEndColor represent the CSS style "border-block-end-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-end-color
type BorderBlockEndColorStyle Color

func (t BorderBlockEndColorStyle) Name() string {
	return "border-block-end-color"
}

var BorderBlockEndColorStyleBrowserVariantsList = []string{}

func (t BorderBlockEndColorStyle) BrowserVariants() []string {
	return BorderBlockEndColorStyleBrowserVariantsList
}

var BorderBlockEndColorStyleUtilitiesMapping = map[string]string{
	"border-block-end-color": "border-block-end-color: currentcolor",
}

func (t BorderBlockEndColorStyle) Utilities() map[string]string {
	return BorderBlockEndColorStyleUtilitiesMapping
}

func (t BorderBlockEndColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBlockEndColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ObjectPosition represent the CSS style "object-position" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/object-position
type ObjectPositionStyle string

func (t ObjectPositionStyle) Name() string {
	return "object-position"
}

var ObjectPositionStyleBrowserVariantsList = []string{}

func (t ObjectPositionStyle) BrowserVariants() []string {
	return ObjectPositionStyleBrowserVariantsList
}

var ObjectPositionStyleUtilitiesMapping = map[string]string{
	"object-position": "object-position: 50% 50%",
}

func (t ObjectPositionStyle) Utilities() map[string]string {
	return ObjectPositionStyleUtilitiesMapping
}

func (t ObjectPositionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ObjectPositionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OffsetPath represent the CSS style "offset-path" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/offset-path
type OffsetPathStyle string

func (t OffsetPathStyle) Name() string {
	return "offset-path"
}

var OffsetPathStyleBrowserVariantsList = []string{}

func (t OffsetPathStyle) BrowserVariants() []string {
	return OffsetPathStyleBrowserVariantsList
}

var OffsetPathStyleUtilitiesMapping = map[string]string{
	"offset-path": "offset-path: none",
}

func (t OffsetPathStyle) Utilities() map[string]string {
	return OffsetPathStyleUtilitiesMapping
}

func (t OffsetPathStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OffsetPathStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMarginInlineEnd represent the CSS style "scroll-margin-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-inline-end
type ScrollMarginInlineEndStyle float64

func (t ScrollMarginInlineEndStyle) Name() string {
	return "scroll-margin-inline-end"
}

var ScrollMarginInlineEndStyleBrowserVariantsList = []string{}

func (t ScrollMarginInlineEndStyle) BrowserVariants() []string {
	return ScrollMarginInlineEndStyleBrowserVariantsList
}

var ScrollMarginInlineEndStyleUtilitiesMapping = map[string]string{
	"scroll-margin-inline-end": "scroll-margin-inline-end: 0",
}

func (t ScrollMarginInlineEndStyle) Utilities() map[string]string {
	return ScrollMarginInlineEndStyleUtilitiesMapping
}

func (t ScrollMarginInlineEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginInlineEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderCollapse represent the CSS style "border-collapse" with value "collapse | separate"
// See https://developer.mozilla.org/docs/Web/CSS/border-collapse
type BorderCollapseStyle string

func (t BorderCollapseStyle) Name() string {
	return "border-collapse"
}

const BorderCollapseStyleCollapse BorderCollapseStyle = "collapse"

const BorderCollapseStyleSeparate BorderCollapseStyle = "separate"

var BorderCollapseStyleBrowserVariantsList = []string{}

func (t BorderCollapseStyle) BrowserVariants() []string {
	return BorderCollapseStyleBrowserVariantsList
}

var BorderCollapseStyleUtilitiesMapping = map[string]string{
	"border-collapse":          "border-collapse: separate",
	"border-collapse-collapse": "border-collapse: collapse",
	"border-collapse-separate": "border-collapse: separate",
}

func (t BorderCollapseStyle) Utilities() map[string]string {
	return BorderCollapseStyleUtilitiesMapping
}

func (t BorderCollapseStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderCollapseStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderTopRightRadius represent the CSS style "border-top-right-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-right-radius
type BorderTopRightRadiusStyle string

func (t BorderTopRightRadiusStyle) Name() string {
	return "border-top-right-radius"
}

var BorderTopRightRadiusStyleBrowserVariantsList = []string{}

func (t BorderTopRightRadiusStyle) BrowserVariants() []string {
	return BorderTopRightRadiusStyleBrowserVariantsList
}

var BorderTopRightRadiusStyleUtilitiesMapping = map[string]string{
	"border-top-right-radius": "border-top-right-radius: 0",
}

func (t BorderTopRightRadiusStyle) Utilities() map[string]string {
	return BorderTopRightRadiusStyleUtilitiesMapping
}

func (t BorderTopRightRadiusStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderTopRightRadiusStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxDecorationBreak represent the CSS style "box-decoration-break" with value "slice | clone"
// See https://developer.mozilla.org/docs/Web/CSS/box-decoration-break
type BoxDecorationBreakStyle string

func (t BoxDecorationBreakStyle) Name() string {
	return "box-decoration-break"
}

const BoxDecorationBreakStyleSlice BoxDecorationBreakStyle = "slice"

const BoxDecorationBreakStyleClone BoxDecorationBreakStyle = "clone"

var BoxDecorationBreakStyleBrowserVariantsList = []string{}

func (t BoxDecorationBreakStyle) BrowserVariants() []string {
	return BoxDecorationBreakStyleBrowserVariantsList
}

var BoxDecorationBreakStyleUtilitiesMapping = map[string]string{
	"box-decoration-break":       "box-decoration-break: slice",
	"box-decoration-break-slice": "box-decoration-break: slice",
	"box-decoration-break-clone": "box-decoration-break: clone",
}

func (t BoxDecorationBreakStyle) Utilities() map[string]string {
	return BoxDecorationBreakStyleUtilitiesMapping
}

func (t BoxDecorationBreakStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxDecorationBreakStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskOrigin represent the CSS style "mask-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-origin
type MaskOriginStyle string

func (t MaskOriginStyle) Name() string {
	return "mask-origin"
}

var MaskOriginStyleBrowserVariantsList = []string{
	"-webkit-mask-origin",
}

func (t MaskOriginStyle) BrowserVariants() []string {
	return MaskOriginStyleBrowserVariantsList
}

var MaskOriginStyleUtilitiesMapping = map[string]string{
	"mask-origin": "mask-origin: border-box",
}

func (t MaskOriginStyle) Utilities() map[string]string {
	return MaskOriginStyleUtilitiesMapping
}

func (t MaskOriginStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskOriginStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PageBreakAfter represent the CSS style "page-break-after" with value "auto | always | avoid | left | right | recto | verso"
// See https://developer.mozilla.org/docs/Web/CSS/page-break-after
type PageBreakAfterStyle string

func (t PageBreakAfterStyle) Name() string {
	return "page-break-after"
}

const PageBreakAfterStyleAuto PageBreakAfterStyle = "auto"

const PageBreakAfterStyleAlways PageBreakAfterStyle = "always"

const PageBreakAfterStyleAvoid PageBreakAfterStyle = "avoid"

const PageBreakAfterStyleLeft PageBreakAfterStyle = "left"

const PageBreakAfterStyleRight PageBreakAfterStyle = "right"

const PageBreakAfterStyleRecto PageBreakAfterStyle = "recto"

const PageBreakAfterStyleVerso PageBreakAfterStyle = "verso"

var PageBreakAfterStyleBrowserVariantsList = []string{}

func (t PageBreakAfterStyle) BrowserVariants() []string {
	return PageBreakAfterStyleBrowserVariantsList
}

var PageBreakAfterStyleUtilitiesMapping = map[string]string{
	"page-break-after":        "page-break-after: auto",
	"page-break-after-auto":   "page-break-after: auto",
	"page-break-after-always": "page-break-after: always",
	"page-break-after-avoid":  "page-break-after: avoid",
	"page-break-after-left":   "page-break-after: left",
	"page-break-after-right":  "page-break-after: right",
	"page-break-after-recto":  "page-break-after: recto",
	"page-break-after-verso":  "page-break-after: verso",
}

func (t PageBreakAfterStyle) Utilities() map[string]string {
	return PageBreakAfterStyleUtilitiesMapping
}

func (t PageBreakAfterStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PageBreakAfterStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextTransform represent the CSS style "text-transform" with value "none | capitalize | uppercase | lowercase | full-width | full-size-kana"
// See https://developer.mozilla.org/docs/Web/CSS/text-transform
type TextTransformStyle string

func (t TextTransformStyle) Name() string {
	return "text-transform"
}

const TextTransformStyleNone TextTransformStyle = "none"

const TextTransformStyleCapitalize TextTransformStyle = "capitalize"

const TextTransformStyleUppercase TextTransformStyle = "uppercase"

const TextTransformStyleLowercase TextTransformStyle = "lowercase"

const TextTransformStyleFullWidth TextTransformStyle = "full-width"

const TextTransformStyleFullSizeKana TextTransformStyle = "full-size-kana"

var TextTransformStyleBrowserVariantsList = []string{}

func (t TextTransformStyle) BrowserVariants() []string {
	return TextTransformStyleBrowserVariantsList
}

var TextTransformStyleUtilitiesMapping = map[string]string{
	"text-transform":                "text-transform: none",
	"text-transform-none":           "text-transform: none",
	"text-transform-capitalize":     "text-transform: capitalize",
	"text-transform-uppercase":      "text-transform: uppercase",
	"text-transform-lowercase":      "text-transform: lowercase",
	"text-transform-full-width":     "text-transform: full-width",
	"text-transform-full-size-kana": "text-transform: full-size-kana",
}

func (t TextTransformStyle) Utilities() map[string]string {
	return TextTransformStyleUtilitiesMapping
}

func (t TextTransformStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextTransformStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformRotateZ represent the CSS style "transform-rotate-z" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformRotateZStyle string

func (t TransformRotateZStyle) Name() string {
	return "transform-rotate-z"
}

var TransformRotateZStyleBrowserVariantsList = []string{}

func (t TransformRotateZStyle) BrowserVariants() []string {
	return TransformRotateZStyleBrowserVariantsList
}

var TransformRotateZStyleUtilitiesMapping = map[string]string{
	"transform-rotate-z": "transform-rotate-z: rotateZ(0)",
}

func (t TransformRotateZStyle) Utilities() map[string]string {
	return TransformRotateZStyleUtilitiesMapping
}

func (t TransformRotateZStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformRotateZStyleUtilitiesMapping[tu]
	return value, hasValue
}

// JustifySelf represent the CSS style "justify-self" with value "auto | normal | stretch | end | start | flex-start | flex-end | center | left | right |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe center | unsafe center"
// See https://developer.mozilla.org/docs/Web/CSS/justify-self
type JustifySelfStyle string

func (t JustifySelfStyle) Name() string {
	return "justify-self"
}

const JustifySelfStyleAuto JustifySelfStyle = "auto"

const JustifySelfStyleNormal JustifySelfStyle = "normal"

const JustifySelfStyleStretch JustifySelfStyle = "stretch"

const JustifySelfStyleEnd JustifySelfStyle = "end"

const JustifySelfStyleStart JustifySelfStyle = "start"

const JustifySelfStyleFlexStart JustifySelfStyle = "flex-start"

const JustifySelfStyleFlexEnd JustifySelfStyle = "flex-end"

const JustifySelfStyleCenter JustifySelfStyle = "center"

const JustifySelfStyleLeft JustifySelfStyle = "left"

const JustifySelfStyleRight JustifySelfStyle = "right"

const JustifySelfStyleBaseline JustifySelfStyle = "baseline"

const JustifySelfStyleFirstBaseline JustifySelfStyle = "first-baseline"

const JustifySelfStyleLastBaseline JustifySelfStyle = "last-baseline"

const JustifySelfStyleSpaceBetween JustifySelfStyle = "space-between"

const JustifySelfStyleSpaceAround JustifySelfStyle = "space-around"

const JustifySelfStyleSpaceEvenly JustifySelfStyle = "space-evenly"

const JustifySelfStyleSafeCenter JustifySelfStyle = "safe-center"

const JustifySelfStyleUnsafeCenter JustifySelfStyle = "unsafe-center"

var JustifySelfStyleBrowserVariantsList = []string{}

func (t JustifySelfStyle) BrowserVariants() []string {
	return JustifySelfStyleBrowserVariantsList
}

var JustifySelfStyleUtilitiesMapping = map[string]string{
	"justify-self":                "justify-self: auto",
	"justify-self-auto":           "justify-self: auto",
	"justify-self-normal":         "justify-self: normal",
	"justify-self-stretch":        "justify-self: stretch",
	"justify-self-end":            "justify-self: end",
	"justify-self-start":          "justify-self: start",
	"justify-self-flex-start":     "justify-self: flex-start",
	"justify-self-flex-end":       "justify-self: flex-end",
	"justify-self-center":         "justify-self: center",
	"justify-self-left":           "justify-self: left",
	"justify-self-right":          "justify-self: right",
	"justify-self-baseline":       "justify-self: baseline",
	"justify-self-first-baseline": "justify-self: first-baseline",
	"justify-self-last-baseline":  "justify-self: last-baseline",
	"justify-self-space-between":  "justify-self: space-between",
	"justify-self-space-around":   "justify-self: space-around",
	"justify-self-space-evenly":   "justify-self: space-evenly",
	"justify-self-safe-center":    "justify-self: safe-center",
	"justify-self-unsafe-center":  "justify-self: unsafe-center",
}

func (t JustifySelfStyle) Utilities() map[string]string {
	return JustifySelfStyleUtilitiesMapping
}

func (t JustifySelfStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = JustifySelfStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Resize represent the CSS style "resize" with value "none | both | horizontal | vertical | block | inline"
// See https://developer.mozilla.org/docs/Web/CSS/resize
type ResizeStyle string

func (t ResizeStyle) Name() string {
	return "resize"
}

const ResizeStyleNone ResizeStyle = "none"

const ResizeStyleBoth ResizeStyle = "both"

const ResizeStyleHorizontal ResizeStyle = "horizontal"

const ResizeStyleVertical ResizeStyle = "vertical"

const ResizeStyleBlock ResizeStyle = "block"

const ResizeStyleInline ResizeStyle = "inline"

var ResizeStyleBrowserVariantsList = []string{}

func (t ResizeStyle) BrowserVariants() []string {
	return ResizeStyleBrowserVariantsList
}

var ResizeStyleUtilitiesMapping = map[string]string{
	"resize":            "resize: none",
	"resize-none":       "resize: none",
	"resize-both":       "resize: both",
	"resize-horizontal": "resize: horizontal",
	"resize-vertical":   "resize: vertical",
	"resize-block":      "resize: block",
	"resize-inline":     "resize: inline",
}

func (t ResizeStyle) Utilities() map[string]string {
	return ResizeStyleUtilitiesMapping
}

func (t ResizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ResizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderInlineStartWidth represent the CSS style "border-inline-start-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-start-width
type BorderInlineStartWidthStyle string

func (t BorderInlineStartWidthStyle) Name() string {
	return "border-inline-start-width"
}

var BorderInlineStartWidthStyleBrowserVariantsList = []string{}

func (t BorderInlineStartWidthStyle) BrowserVariants() []string {
	return BorderInlineStartWidthStyleBrowserVariantsList
}

var BorderInlineStartWidthStyleUtilitiesMapping = map[string]string{
	"border-inline-start-width": "border-inline-start-width: medium",
}

func (t BorderInlineStartWidthStyle) Utilities() map[string]string {
	return BorderInlineStartWidthStyleUtilitiesMapping
}

func (t BorderInlineStartWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderInlineStartWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontSize represent the CSS style "font-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-size
type FontSizeStyle string

func (t FontSizeStyle) Name() string {
	return "font-size"
}

var FontSizeStyleBrowserVariantsList = []string{}

func (t FontSizeStyle) BrowserVariants() []string {
	return FontSizeStyleBrowserVariantsList
}

var FontSizeStyleUtilitiesMapping = map[string]string{
	"font-size": "font-size: medium",
}

func (t FontSizeStyle) Utilities() map[string]string {
	return FontSizeStyleUtilitiesMapping
}

func (t FontSizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontSizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskBorderRepeat represent the CSS style "mask-border-repeat" with value "stretch | repeat | round | space | stretch repeat | stretch round | stretch space | repeat stretch | repeat round | repeat space | round stretch | round repeat | round space | space stretch | space repeat | space round"
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-repeat
type MaskBorderRepeatStyle string

func (t MaskBorderRepeatStyle) Name() string {
	return "mask-border-repeat"
}

const MaskBorderRepeatStyleStretch MaskBorderRepeatStyle = "stretch"

const MaskBorderRepeatStyleRepeat MaskBorderRepeatStyle = "repeat"

const MaskBorderRepeatStyleRound MaskBorderRepeatStyle = "round"

const MaskBorderRepeatStyleSpace MaskBorderRepeatStyle = "space"

const MaskBorderRepeatStyleStretchRepeat MaskBorderRepeatStyle = "stretch-repeat"

const MaskBorderRepeatStyleStretchRound MaskBorderRepeatStyle = "stretch-round"

const MaskBorderRepeatStyleStretchSpace MaskBorderRepeatStyle = "stretch-space"

const MaskBorderRepeatStyleRepeatStretch MaskBorderRepeatStyle = "repeat-stretch"

const MaskBorderRepeatStyleRepeatRound MaskBorderRepeatStyle = "repeat-round"

const MaskBorderRepeatStyleRepeatSpace MaskBorderRepeatStyle = "repeat-space"

const MaskBorderRepeatStyleRoundStretch MaskBorderRepeatStyle = "round-stretch"

const MaskBorderRepeatStyleRoundRepeat MaskBorderRepeatStyle = "round-repeat"

const MaskBorderRepeatStyleRoundSpace MaskBorderRepeatStyle = "round-space"

const MaskBorderRepeatStyleSpaceStretch MaskBorderRepeatStyle = "space-stretch"

const MaskBorderRepeatStyleSpaceRepeat MaskBorderRepeatStyle = "space-repeat"

const MaskBorderRepeatStyleSpaceRound MaskBorderRepeatStyle = "space-round"

var MaskBorderRepeatStyleBrowserVariantsList = []string{}

func (t MaskBorderRepeatStyle) BrowserVariants() []string {
	return MaskBorderRepeatStyleBrowserVariantsList
}

var MaskBorderRepeatStyleUtilitiesMapping = map[string]string{
	"mask-border-repeat":                "mask-border-repeat: stretch",
	"mask-border-repeat-stretch":        "mask-border-repeat: stretch",
	"mask-border-repeat-repeat":         "mask-border-repeat: repeat",
	"mask-border-repeat-round":          "mask-border-repeat: round",
	"mask-border-repeat-space":          "mask-border-repeat: space",
	"mask-border-repeat-stretch-repeat": "mask-border-repeat: stretch-repeat",
	"mask-border-repeat-stretch-round":  "mask-border-repeat: stretch-round",
	"mask-border-repeat-stretch-space":  "mask-border-repeat: stretch-space",
	"mask-border-repeat-repeat-stretch": "mask-border-repeat: repeat-stretch",
	"mask-border-repeat-repeat-round":   "mask-border-repeat: repeat-round",
	"mask-border-repeat-repeat-space":   "mask-border-repeat: repeat-space",
	"mask-border-repeat-round-stretch":  "mask-border-repeat: round-stretch",
	"mask-border-repeat-round-repeat":   "mask-border-repeat: round-repeat",
	"mask-border-repeat-round-space":    "mask-border-repeat: round-space",
	"mask-border-repeat-space-stretch":  "mask-border-repeat: space-stretch",
	"mask-border-repeat-space-repeat":   "mask-border-repeat: space-repeat",
	"mask-border-repeat-space-round":    "mask-border-repeat: space-round",
}

func (t MaskBorderRepeatStyle) Utilities() map[string]string {
	return MaskBorderRepeatStyleUtilitiesMapping
}

func (t MaskBorderRepeatStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskBorderRepeatStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextDecorationStyle represent the CSS style "text-decoration-style" with value "solid | double | dotted | dashed | wavy"
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-style
type TextDecorationStyleStyle string

func (t TextDecorationStyleStyle) Name() string {
	return "text-decoration-style"
}

const TextDecorationStyleStyleSolid TextDecorationStyleStyle = "solid"

const TextDecorationStyleStyleDouble TextDecorationStyleStyle = "double"

const TextDecorationStyleStyleDotted TextDecorationStyleStyle = "dotted"

const TextDecorationStyleStyleDashed TextDecorationStyleStyle = "dashed"

const TextDecorationStyleStyleWavy TextDecorationStyleStyle = "wavy"

var TextDecorationStyleStyleBrowserVariantsList = []string{}

func (t TextDecorationStyleStyle) BrowserVariants() []string {
	return TextDecorationStyleStyleBrowserVariantsList
}

var TextDecorationStyleStyleUtilitiesMapping = map[string]string{
	"text-decoration-style":        "text-decoration-style: solid",
	"text-decoration-style-solid":  "text-decoration-style: solid",
	"text-decoration-style-double": "text-decoration-style: double",
	"text-decoration-style-dotted": "text-decoration-style: dotted",
	"text-decoration-style-dashed": "text-decoration-style: dashed",
	"text-decoration-style-wavy":   "text-decoration-style: wavy",
}

func (t TextDecorationStyleStyle) Utilities() map[string]string {
	return TextDecorationStyleStyleUtilitiesMapping
}

func (t TextDecorationStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextDecorationStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformScale represent the CSS style "transform-scale" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformScaleStyle string

func (t TransformScaleStyle) Name() string {
	return "transform-scale"
}

var TransformScaleStyleBrowserVariantsList = []string{}

func (t TransformScaleStyle) BrowserVariants() []string {
	return TransformScaleStyleBrowserVariantsList
}

var TransformScaleStyleUtilitiesMapping = map[string]string{
	"transform-scale": "transform-scale: scale(1,1)",
}

func (t TransformScaleStyle) Utilities() map[string]string {
	return TransformScaleStyleUtilitiesMapping
}

func (t TransformScaleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformScaleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Clear represent the CSS style "clear" with value "none | left | right | both | inline-start | inline-end"
// See https://developer.mozilla.org/docs/Web/CSS/clear
type ClearStyle string

func (t ClearStyle) Name() string {
	return "clear"
}

const ClearStyleNone ClearStyle = "none"

const ClearStyleLeft ClearStyle = "left"

const ClearStyleRight ClearStyle = "right"

const ClearStyleBoth ClearStyle = "both"

const ClearStyleInlineStart ClearStyle = "inline-start"

const ClearStyleInlineEnd ClearStyle = "inline-end"

var ClearStyleBrowserVariantsList = []string{}

func (t ClearStyle) BrowserVariants() []string {
	return ClearStyleBrowserVariantsList
}

var ClearStyleUtilitiesMapping = map[string]string{
	"clear":              "clear: none",
	"clear-none":         "clear: none",
	"clear-left":         "clear: left",
	"clear-right":        "clear: right",
	"clear-both":         "clear: both",
	"clear-inline-start": "clear: inline-start",
	"clear-inline-end":   "clear: inline-end",
}

func (t ClearStyle) Utilities() map[string]string {
	return ClearStyleUtilitiesMapping
}

func (t ClearStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ClearStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontSmooth represent the CSS style "font-smooth" with value "auto | never | always"
// See https://developer.mozilla.org/docs/Web/CSS/font-smooth
type FontSmoothStyle string

func (t FontSmoothStyle) Name() string {
	return "font-smooth"
}

const FontSmoothStyleAuto FontSmoothStyle = "auto"

const FontSmoothStyleNever FontSmoothStyle = "never"

const FontSmoothStyleAlways FontSmoothStyle = "always"

var FontSmoothStyleBrowserVariantsList = []string{}

func (t FontSmoothStyle) BrowserVariants() []string {
	return FontSmoothStyleBrowserVariantsList
}

var FontSmoothStyleUtilitiesMapping = map[string]string{
	"font-smooth":        "font-smooth: auto",
	"font-smooth-auto":   "font-smooth: auto",
	"font-smooth-never":  "font-smooth: never",
	"font-smooth-always": "font-smooth: always",
}

func (t FontSmoothStyle) Utilities() map[string]string {
	return FontSmoothStyleUtilitiesMapping
}

func (t FontSmoothStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontSmoothStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaxLines represent the CSS style "max-lines" with value ""
// See
type MaxLinesStyle string

func (t MaxLinesStyle) Name() string {
	return "max-lines"
}

var MaxLinesStyleBrowserVariantsList = []string{}

func (t MaxLinesStyle) BrowserVariants() []string {
	return MaxLinesStyleBrowserVariantsList
}

var MaxLinesStyleUtilitiesMapping = map[string]string{
	"max-lines": "max-lines: none",
}

func (t MaxLinesStyle) Utilities() map[string]string {
	return MaxLinesStyleUtilitiesMapping
}

func (t MaxLinesStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaxLinesStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPaddingLeft represent the CSS style "scroll-padding-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-left
type ScrollPaddingLeftStyle string

func (t ScrollPaddingLeftStyle) Name() string {
	return "scroll-padding-left"
}

var ScrollPaddingLeftStyleBrowserVariantsList = []string{}

func (t ScrollPaddingLeftStyle) BrowserVariants() []string {
	return ScrollPaddingLeftStyleBrowserVariantsList
}

var ScrollPaddingLeftStyleUtilitiesMapping = map[string]string{
	"scroll-padding-left": "scroll-padding-left: auto",
}

func (t ScrollPaddingLeftStyle) Utilities() map[string]string {
	return ScrollPaddingLeftStyleUtilitiesMapping
}

func (t ScrollPaddingLeftStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingLeftStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FloodColor represent the CSS style "flood-color" with value ""
// See
type FloodColorStyle Color

func (t FloodColorStyle) Name() string {
	return "flood-color"
}

var FloodColorStyleBrowserVariantsList = []string{}

func (t FloodColorStyle) BrowserVariants() []string {
	return FloodColorStyleBrowserVariantsList
}

var FloodColorStyleUtilitiesMapping = map[string]string{
	"flood-color": "flood-color: black",
}

func (t FloodColorStyle) Utilities() map[string]string {
	return FloodColorStyleUtilitiesMapping
}

func (t FloodColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FloodColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColumnRuleWidth represent the CSS style "column-rule-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-rule-width
type ColumnRuleWidthStyle string

func (t ColumnRuleWidthStyle) Name() string {
	return "column-rule-width"
}

var ColumnRuleWidthStyleBrowserVariantsList = []string{}

func (t ColumnRuleWidthStyle) BrowserVariants() []string {
	return ColumnRuleWidthStyleBrowserVariantsList
}

var ColumnRuleWidthStyleUtilitiesMapping = map[string]string{
	"column-rule-width": "column-rule-width: medium",
}

func (t ColumnRuleWidthStyle) Utilities() map[string]string {
	return ColumnRuleWidthStyleUtilitiesMapping
}

func (t ColumnRuleWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColumnRuleWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformScaleX represent the CSS style "transform-scale-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformScaleXStyle string

func (t TransformScaleXStyle) Name() string {
	return "transform-scale-x"
}

var TransformScaleXStyleBrowserVariantsList = []string{}

func (t TransformScaleXStyle) BrowserVariants() []string {
	return TransformScaleXStyleBrowserVariantsList
}

var TransformScaleXStyleUtilitiesMapping = map[string]string{
	"transform-scale-x": "transform-scale-x: scaleX(0)",
}

func (t TransformScaleXStyle) Utilities() map[string]string {
	return TransformScaleXStyleUtilitiesMapping
}

func (t TransformScaleXStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformScaleXStyleUtilitiesMapping[tu]
	return value, hasValue
}

// WordSpacing represent the CSS style "word-spacing" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/word-spacing
type WordSpacingStyle string

func (t WordSpacingStyle) Name() string {
	return "word-spacing"
}

var WordSpacingStyleBrowserVariantsList = []string{}

func (t WordSpacingStyle) BrowserVariants() []string {
	return WordSpacingStyleBrowserVariantsList
}

var WordSpacingStyleUtilitiesMapping = map[string]string{
	"word-spacing": "word-spacing: normal",
}

func (t WordSpacingStyle) Utilities() map[string]string {
	return WordSpacingStyleUtilitiesMapping
}

func (t WordSpacingStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = WordSpacingStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontWeight represent the CSS style "font-weight" with value "100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | bolder | lighter"
// See https://developer.mozilla.org/docs/Web/CSS/font-weight
type FontWeightStyle string

func (t FontWeightStyle) Name() string {
	return "font-weight"
}

const FontWeightStyle100 FontWeightStyle = "100"

const FontWeightStyle200 FontWeightStyle = "200"

const FontWeightStyle300 FontWeightStyle = "300"

const FontWeightStyle400 FontWeightStyle = "400"

const FontWeightStyle500 FontWeightStyle = "500"

const FontWeightStyle600 FontWeightStyle = "600"

const FontWeightStyle700 FontWeightStyle = "700"

const FontWeightStyle800 FontWeightStyle = "800"

const FontWeightStyle900 FontWeightStyle = "900"

const FontWeightStyleBolder FontWeightStyle = "bolder"

const FontWeightStyleLighter FontWeightStyle = "lighter"

var FontWeightStyleBrowserVariantsList = []string{}

func (t FontWeightStyle) BrowserVariants() []string {
	return FontWeightStyleBrowserVariantsList
}

var FontWeightStyleUtilitiesMapping = map[string]string{
	"font-weight":         "font-weight: normal",
	"font-weight-100":     "font-weight: 100",
	"font-weight-200":     "font-weight: 200",
	"font-weight-300":     "font-weight: 300",
	"font-weight-400":     "font-weight: 400",
	"font-weight-500":     "font-weight: 500",
	"font-weight-600":     "font-weight: 600",
	"font-weight-700":     "font-weight: 700",
	"font-weight-800":     "font-weight: 800",
	"font-weight-900":     "font-weight: 900",
	"font-weight-bolder":  "font-weight: bolder",
	"font-weight-lighter": "font-weight: lighter",
}

func (t FontWeightStyle) Utilities() map[string]string {
	return FontWeightStyleUtilitiesMapping
}

func (t FontWeightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontWeightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ImageRendering represent the CSS style "image-rendering" with value "auto | crisp-edges | pixelated"
// See https://developer.mozilla.org/docs/Web/CSS/image-rendering
type ImageRenderingStyle string

func (t ImageRenderingStyle) Name() string {
	return "image-rendering"
}

const ImageRenderingStyleAuto ImageRenderingStyle = "auto"

const ImageRenderingStyleCrispEdges ImageRenderingStyle = "crisp-edges"

const ImageRenderingStylePixelated ImageRenderingStyle = "pixelated"

var ImageRenderingStyleBrowserVariantsList = []string{}

func (t ImageRenderingStyle) BrowserVariants() []string {
	return ImageRenderingStyleBrowserVariantsList
}

var ImageRenderingStyleUtilitiesMapping = map[string]string{
	"image-rendering":             "image-rendering: auto",
	"image-rendering-auto":        "image-rendering: auto",
	"image-rendering-crisp-edges": "image-rendering: crisp-edges",
	"image-rendering-pixelated":   "image-rendering: pixelated",
}

func (t ImageRenderingStyle) Utilities() map[string]string {
	return ImageRenderingStyleUtilitiesMapping
}

func (t ImageRenderingStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ImageRenderingStyleUtilitiesMapping[tu]
	return value, hasValue
}

// CounterSet represent the CSS style "counter-set" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/counter-set
type CounterSetStyle string

func (t CounterSetStyle) Name() string {
	return "counter-set"
}

var CounterSetStyleBrowserVariantsList = []string{}

func (t CounterSetStyle) BrowserVariants() []string {
	return CounterSetStyleBrowserVariantsList
}

var CounterSetStyleUtilitiesMapping = map[string]string{
	"counter-set": "counter-set: none",
}

func (t CounterSetStyle) Utilities() map[string]string {
	return CounterSetStyleUtilitiesMapping
}

func (t CounterSetStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = CounterSetStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginInlineEnd represent the CSS style "margin-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-inline-end
type MarginInlineEndStyle string

func (t MarginInlineEndStyle) Name() string {
	return "margin-inline-end"
}

var MarginInlineEndStyleBrowserVariantsList = []string{}

func (t MarginInlineEndStyle) BrowserVariants() []string {
	return MarginInlineEndStyleBrowserVariantsList
}

var MarginInlineEndStyleUtilitiesMapping = map[string]string{
	"margin-inline-end": "margin-inline-end: 0",
}

func (t MarginInlineEndStyle) Utilities() map[string]string {
	return MarginInlineEndStyleUtilitiesMapping
}

func (t MarginInlineEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginInlineEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Order represent the CSS style "order" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/order
type OrderStyle float64

func (t OrderStyle) Name() string {
	return "order"
}

var OrderStyleBrowserVariantsList = []string{}

func (t OrderStyle) BrowserVariants() []string {
	return OrderStyleBrowserVariantsList
}

var OrderStyleUtilitiesMapping = map[string]string{
	"order": "order: 0",
}

func (t OrderStyle) Utilities() map[string]string {
	return OrderStyleUtilitiesMapping
}

func (t OrderStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OrderStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderStartStartRadius represent the CSS style "border-start-start-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-start-start-radius
type BorderStartStartRadiusStyle string

func (t BorderStartStartRadiusStyle) Name() string {
	return "border-start-start-radius"
}

var BorderStartStartRadiusStyleBrowserVariantsList = []string{}

func (t BorderStartStartRadiusStyle) BrowserVariants() []string {
	return BorderStartStartRadiusStyleBrowserVariantsList
}

var BorderStartStartRadiusStyleUtilitiesMapping = map[string]string{
	"border-start-start-radius": "border-start-start-radius: 0",
}

func (t BorderStartStartRadiusStyle) Utilities() map[string]string {
	return BorderStartStartRadiusStyleUtilitiesMapping
}

func (t BorderStartStartRadiusStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderStartStartRadiusStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginBlockEnd represent the CSS style "margin-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-block-end
type MarginBlockEndStyle string

func (t MarginBlockEndStyle) Name() string {
	return "margin-block-end"
}

var MarginBlockEndStyleBrowserVariantsList = []string{}

func (t MarginBlockEndStyle) BrowserVariants() []string {
	return MarginBlockEndStyleBrowserVariantsList
}

var MarginBlockEndStyleUtilitiesMapping = map[string]string{
	"margin-block-end": "margin-block-end: 0",
}

func (t MarginBlockEndStyle) Utilities() map[string]string {
	return MarginBlockEndStyleUtilitiesMapping
}

func (t MarginBlockEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginBlockEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ClipPath represent the CSS style "clip-path" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/clip-path
type ClipPathStyle string

func (t ClipPathStyle) Name() string {
	return "clip-path"
}

var ClipPathStyleBrowserVariantsList = []string{}

func (t ClipPathStyle) BrowserVariants() []string {
	return ClipPathStyleBrowserVariantsList
}

var ClipPathStyleUtilitiesMapping = map[string]string{
	"clip-path": "clip-path: none",
}

func (t ClipPathStyle) Utilities() map[string]string {
	return ClipPathStyleUtilitiesMapping
}

func (t ClipPathStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ClipPathStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaddingTop represent the CSS style "padding-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-top
type PaddingTopStyle string

func (t PaddingTopStyle) Name() string {
	return "padding-top"
}

var PaddingTopStyleBrowserVariantsList = []string{}

func (t PaddingTopStyle) BrowserVariants() []string {
	return PaddingTopStyleBrowserVariantsList
}

var PaddingTopStyleUtilitiesMapping = map[string]string{
	"padding-top": "padding-top: 0",
}

func (t PaddingTopStyle) Utilities() map[string]string {
	return PaddingTopStyleUtilitiesMapping
}

func (t PaddingTopStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaddingTopStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextOrientation represent the CSS style "text-orientation" with value "mixed | upright | sideways"
// See https://developer.mozilla.org/docs/Web/CSS/text-orientation
type TextOrientationStyle string

func (t TextOrientationStyle) Name() string {
	return "text-orientation"
}

const TextOrientationStyleMixed TextOrientationStyle = "mixed"

const TextOrientationStyleUpright TextOrientationStyle = "upright"

const TextOrientationStyleSideways TextOrientationStyle = "sideways"

var TextOrientationStyleBrowserVariantsList = []string{}

func (t TextOrientationStyle) BrowserVariants() []string {
	return TextOrientationStyleBrowserVariantsList
}

var TextOrientationStyleUtilitiesMapping = map[string]string{
	"text-orientation":          "text-orientation: mixed",
	"text-orientation-mixed":    "text-orientation: mixed",
	"text-orientation-upright":  "text-orientation: upright",
	"text-orientation-sideways": "text-orientation: sideways",
}

func (t TextOrientationStyle) Utilities() map[string]string {
	return TextOrientationStyleUtilitiesMapping
}

func (t TextOrientationStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextOrientationStyleUtilitiesMapping[tu]
	return value, hasValue
}

// StrokeDashoffset represent the CSS style "stroke-dashoffset" with value ""
// See
type StrokeDashoffsetStyle float64

func (t StrokeDashoffsetStyle) Name() string {
	return "stroke-dashoffset"
}

var StrokeDashoffsetStyleBrowserVariantsList = []string{}

func (t StrokeDashoffsetStyle) BrowserVariants() []string {
	return StrokeDashoffsetStyleBrowserVariantsList
}

var StrokeDashoffsetStyleUtilitiesMapping = map[string]string{
	"stroke-dashoffset": "stroke-dashoffset: 0",
}

func (t StrokeDashoffsetStyle) Utilities() map[string]string {
	return StrokeDashoffsetStyleUtilitiesMapping
}

func (t StrokeDashoffsetStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = StrokeDashoffsetStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderInlineWidth represent the CSS style "border-inline-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-width
type BorderInlineWidthStyle string

func (t BorderInlineWidthStyle) Name() string {
	return "border-inline-width"
}

var BorderInlineWidthStyleBrowserVariantsList = []string{}

func (t BorderInlineWidthStyle) BrowserVariants() []string {
	return BorderInlineWidthStyleBrowserVariantsList
}

var BorderInlineWidthStyleUtilitiesMapping = map[string]string{
	"border-inline-width": "border-inline-width: medium",
}

func (t BorderInlineWidthStyle) Utilities() map[string]string {
	return BorderInlineWidthStyleUtilitiesMapping
}

func (t BorderInlineWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderInlineWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Rotate represent the CSS style "rotate" with value "angle"
// See https://developer.mozilla.org/docs/Web/CSS/rotate
type RotateStyle string

func (t RotateStyle) Name() string {
	return "rotate"
}

const RotateStyleAngle RotateStyle = "angle"

var RotateStyleBrowserVariantsList = []string{}

func (t RotateStyle) BrowserVariants() []string {
	return RotateStyleBrowserVariantsList
}

var RotateStyleUtilitiesMapping = map[string]string{
	"rotate":       "rotate: none",
	"rotate-angle": "rotate: angle",
}

func (t RotateStyle) Utilities() map[string]string {
	return RotateStyleUtilitiesMapping
}

func (t RotateStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = RotateStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformRotateX represent the CSS style "transform-rotate-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformRotateXStyle string

func (t TransformRotateXStyle) Name() string {
	return "transform-rotate-x"
}

var TransformRotateXStyleBrowserVariantsList = []string{}

func (t TransformRotateXStyle) BrowserVariants() []string {
	return TransformRotateXStyleBrowserVariantsList
}

var TransformRotateXStyleUtilitiesMapping = map[string]string{
	"transform-rotate-x": "transform-rotate-x: rotateX(0)",
}

func (t TransformRotateXStyle) Utilities() map[string]string {
	return TransformRotateXStyleUtilitiesMapping
}

func (t TransformRotateXStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformRotateXStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Azimuth represent the CSS style "azimuth" with value "left-side | far-left | left | center-left | center | center-right | right | far-right | right-side | behind | leftwards | rightwards"
// See https://developer.mozilla.org/docs/Web/CSS/azimuth
type AzimuthStyle string

func (t AzimuthStyle) Name() string {
	return "azimuth"
}

const AzimuthStyleLeftSide AzimuthStyle = "left-side"

const AzimuthStyleFarLeft AzimuthStyle = "far-left"

const AzimuthStyleLeft AzimuthStyle = "left"

const AzimuthStyleCenterLeft AzimuthStyle = "center-left"

const AzimuthStyleCenter AzimuthStyle = "center"

const AzimuthStyleCenterRight AzimuthStyle = "center-right"

const AzimuthStyleRight AzimuthStyle = "right"

const AzimuthStyleFarRight AzimuthStyle = "far-right"

const AzimuthStyleRightSide AzimuthStyle = "right-side"

const AzimuthStyleBehind AzimuthStyle = "behind"

const AzimuthStyleLeftwards AzimuthStyle = "leftwards"

const AzimuthStyleRightwards AzimuthStyle = "rightwards"

var AzimuthStyleBrowserVariantsList = []string{}

func (t AzimuthStyle) BrowserVariants() []string {
	return AzimuthStyleBrowserVariantsList
}

var AzimuthStyleUtilitiesMapping = map[string]string{
	"azimuth":              "azimuth: center",
	"azimuth-left-side":    "azimuth: left-side",
	"azimuth-far-left":     "azimuth: far-left",
	"azimuth-left":         "azimuth: left",
	"azimuth-center-left":  "azimuth: center-left",
	"azimuth-center":       "azimuth: center",
	"azimuth-center-right": "azimuth: center-right",
	"azimuth-right":        "azimuth: right",
	"azimuth-far-right":    "azimuth: far-right",
	"azimuth-right-side":   "azimuth: right-side",
	"azimuth-behind":       "azimuth: behind",
	"azimuth-leftwards":    "azimuth: leftwards",
	"azimuth-rightwards":   "azimuth: rightwards",
}

func (t AzimuthStyle) Utilities() map[string]string {
	return AzimuthStyleUtilitiesMapping
}

func (t AzimuthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AzimuthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Bottom represent the CSS style "bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/bottom
type BottomStyle string

func (t BottomStyle) Name() string {
	return "bottom"
}

var BottomStyleBrowserVariantsList = []string{}

func (t BottomStyle) BrowserVariants() []string {
	return BottomStyleBrowserVariantsList
}

var BottomStyleUtilitiesMapping = map[string]string{
	"bottom": "bottom: auto",
}

func (t BottomStyle) Utilities() map[string]string {
	return BottomStyleUtilitiesMapping
}

func (t BottomStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BottomStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BreakBefore represent the CSS style "break-before" with value "auto | avoid | always | all | avoid-page | page | left | right | recto | verso | avoid-column | column | avoid-region | region"
// See https://developer.mozilla.org/docs/Web/CSS/break-before
type BreakBeforeStyle string

func (t BreakBeforeStyle) Name() string {
	return "break-before"
}

const BreakBeforeStyleAuto BreakBeforeStyle = "auto"

const BreakBeforeStyleAvoid BreakBeforeStyle = "avoid"

const BreakBeforeStyleAlways BreakBeforeStyle = "always"

const BreakBeforeStyleAll BreakBeforeStyle = "all"

const BreakBeforeStyleAvoidPage BreakBeforeStyle = "avoid-page"

const BreakBeforeStylePage BreakBeforeStyle = "page"

const BreakBeforeStyleLeft BreakBeforeStyle = "left"

const BreakBeforeStyleRight BreakBeforeStyle = "right"

const BreakBeforeStyleRecto BreakBeforeStyle = "recto"

const BreakBeforeStyleVerso BreakBeforeStyle = "verso"

const BreakBeforeStyleAvoidColumn BreakBeforeStyle = "avoid-column"

const BreakBeforeStyleColumn BreakBeforeStyle = "column"

const BreakBeforeStyleAvoidRegion BreakBeforeStyle = "avoid-region"

const BreakBeforeStyleRegion BreakBeforeStyle = "region"

var BreakBeforeStyleBrowserVariantsList = []string{}

func (t BreakBeforeStyle) BrowserVariants() []string {
	return BreakBeforeStyleBrowserVariantsList
}

var BreakBeforeStyleUtilitiesMapping = map[string]string{
	"break-before":              "break-before: auto",
	"break-before-auto":         "break-before: auto",
	"break-before-avoid":        "break-before: avoid",
	"break-before-always":       "break-before: always",
	"break-before-all":          "break-before: all",
	"break-before-avoid-page":   "break-before: avoid-page",
	"break-before-page":         "break-before: page",
	"break-before-left":         "break-before: left",
	"break-before-right":        "break-before: right",
	"break-before-recto":        "break-before: recto",
	"break-before-verso":        "break-before: verso",
	"break-before-avoid-column": "break-before: avoid-column",
	"break-before-column":       "break-before: column",
	"break-before-avoid-region": "break-before: avoid-region",
	"break-before-region":       "break-before: region",
}

func (t BreakBeforeStyle) Utilities() map[string]string {
	return BreakBeforeStyleUtilitiesMapping
}

func (t BreakBeforeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BreakBeforeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransitionDuration represent the CSS style "transition-duration" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-duration
type TransitionDurationStyle time.Duration

func (t TransitionDurationStyle) Name() string {
	return "transition-duration"
}

var TransitionDurationStyleBrowserVariantsList = []string{}

func (t TransitionDurationStyle) BrowserVariants() []string {
	return TransitionDurationStyleBrowserVariantsList
}

var TransitionDurationStyleUtilitiesMapping = map[string]string{
	"transition-duration": "transition-duration: 0s",
}

func (t TransitionDurationStyle) Utilities() map[string]string {
	return TransitionDurationStyleUtilitiesMapping
}

func (t TransitionDurationStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransitionDurationStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Right represent the CSS style "right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/right
type RightStyle string

func (t RightStyle) Name() string {
	return "right"
}

var RightStyleBrowserVariantsList = []string{}

func (t RightStyle) BrowserVariants() []string {
	return RightStyleBrowserVariantsList
}

var RightStyleUtilitiesMapping = map[string]string{
	"right": "right: auto",
}

func (t RightStyle) Utilities() map[string]string {
	return RightStyleUtilitiesMapping
}

func (t RightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = RightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextDecorationLine represent the CSS style "text-decoration-line" with value "none | underline | overline | line-through | blink | spelling-error | grammar-error"
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-line
type TextDecorationLineStyle string

func (t TextDecorationLineStyle) Name() string {
	return "text-decoration-line"
}

const TextDecorationLineStyleNone TextDecorationLineStyle = "none"

const TextDecorationLineStyleUnderline TextDecorationLineStyle = "underline"

const TextDecorationLineStyleOverline TextDecorationLineStyle = "overline"

const TextDecorationLineStyleLineThrough TextDecorationLineStyle = "line-through"

const TextDecorationLineStyleBlink TextDecorationLineStyle = "blink"

const TextDecorationLineStyleSpellingError TextDecorationLineStyle = "spelling-error"

const TextDecorationLineStyleGrammarError TextDecorationLineStyle = "grammar-error"

var TextDecorationLineStyleBrowserVariantsList = []string{}

func (t TextDecorationLineStyle) BrowserVariants() []string {
	return TextDecorationLineStyleBrowserVariantsList
}

var TextDecorationLineStyleUtilitiesMapping = map[string]string{
	"text-decoration-line":                "text-decoration-line: none",
	"text-decoration-line-none":           "text-decoration-line: none",
	"text-decoration-line-underline":      "text-decoration-line: underline",
	"text-decoration-line-overline":       "text-decoration-line: overline",
	"text-decoration-line-line-through":   "text-decoration-line: line-through",
	"text-decoration-line-blink":          "text-decoration-line: blink",
	"text-decoration-line-spelling-error": "text-decoration-line: spelling-error",
	"text-decoration-line-grammar-error":  "text-decoration-line: grammar-error",
}

func (t TextDecorationLineStyle) Utilities() map[string]string {
	return TextDecorationLineStyleUtilitiesMapping
}

func (t TextDecorationLineStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextDecorationLineStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderRightColor represent the CSS style "border-right-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-right-color
type BorderRightColorStyle Color

func (t BorderRightColorStyle) Name() string {
	return "border-right-color"
}

var BorderRightColorStyleBrowserVariantsList = []string{}

func (t BorderRightColorStyle) BrowserVariants() []string {
	return BorderRightColorStyleBrowserVariantsList
}

var BorderRightColorStyleUtilitiesMapping = map[string]string{
	"border-right-color": "border-right-color: currentcolor",
}

func (t BorderRightColorStyle) Utilities() map[string]string {
	return BorderRightColorStyleUtilitiesMapping
}

func (t BorderRightColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderRightColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TouchAction represent the CSS style "touch-action" with value "auto | none | pan-x | pan-left | pan-right | pan-y | pan-up | pan-down | pinch-zoom | manipulation"
// See https://developer.mozilla.org/docs/Web/CSS/touch-action
type TouchActionStyle string

func (t TouchActionStyle) Name() string {
	return "touch-action"
}

const TouchActionStyleAuto TouchActionStyle = "auto"

const TouchActionStyleNone TouchActionStyle = "none"

const TouchActionStylePanX TouchActionStyle = "pan-x"

const TouchActionStylePanLeft TouchActionStyle = "pan-left"

const TouchActionStylePanRight TouchActionStyle = "pan-right"

const TouchActionStylePanY TouchActionStyle = "pan-y"

const TouchActionStylePanUp TouchActionStyle = "pan-up"

const TouchActionStylePanDown TouchActionStyle = "pan-down"

const TouchActionStylePinchZoom TouchActionStyle = "pinch-zoom"

const TouchActionStyleManipulation TouchActionStyle = "manipulation"

var TouchActionStyleBrowserVariantsList = []string{}

func (t TouchActionStyle) BrowserVariants() []string {
	return TouchActionStyleBrowserVariantsList
}

var TouchActionStyleUtilitiesMapping = map[string]string{
	"touch-action":              "touch-action: auto",
	"touch-action-auto":         "touch-action: auto",
	"touch-action-none":         "touch-action: none",
	"touch-action-pan-x":        "touch-action: pan-x",
	"touch-action-pan-left":     "touch-action: pan-left",
	"touch-action-pan-right":    "touch-action: pan-right",
	"touch-action-pan-y":        "touch-action: pan-y",
	"touch-action-pan-up":       "touch-action: pan-up",
	"touch-action-pan-down":     "touch-action: pan-down",
	"touch-action-pinch-zoom":   "touch-action: pinch-zoom",
	"touch-action-manipulation": "touch-action: manipulation",
}

func (t TouchActionStyle) Utilities() map[string]string {
	return TouchActionStyleUtilitiesMapping
}

func (t TouchActionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TouchActionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// WillChange represent the CSS style "will-change" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/will-change
type WillChangeStyle string

func (t WillChangeStyle) Name() string {
	return "will-change"
}

var WillChangeStyleBrowserVariantsList = []string{}

func (t WillChangeStyle) BrowserVariants() []string {
	return WillChangeStyleBrowserVariantsList
}

var WillChangeStyleUtilitiesMapping = map[string]string{
	"will-change": "will-change: auto",
}

func (t WillChangeStyle) Utilities() map[string]string {
	return WillChangeStyleUtilitiesMapping
}

func (t WillChangeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = WillChangeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaddingInlineStart represent the CSS style "padding-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-inline-start
type PaddingInlineStartStyle string

func (t PaddingInlineStartStyle) Name() string {
	return "padding-inline-start"
}

var PaddingInlineStartStyleBrowserVariantsList = []string{}

func (t PaddingInlineStartStyle) BrowserVariants() []string {
	return PaddingInlineStartStyleBrowserVariantsList
}

var PaddingInlineStartStyleUtilitiesMapping = map[string]string{
	"padding-inline-start": "padding-inline-start: 0",
}

func (t PaddingInlineStartStyle) Utilities() map[string]string {
	return PaddingInlineStartStyleUtilitiesMapping
}

func (t PaddingInlineStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaddingInlineStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMarginInline represent the CSS style "scroll-margin-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-inline
type ScrollMarginInlineStyle float64

func (t ScrollMarginInlineStyle) Name() string {
	return "scroll-margin-inline"
}

var ScrollMarginInlineStyleBrowserVariantsList = []string{}

func (t ScrollMarginInlineStyle) BrowserVariants() []string {
	return ScrollMarginInlineStyleBrowserVariantsList
}

var ScrollMarginInlineStyleUtilitiesMapping = map[string]string{
	"scroll-margin-inline": "scroll-margin-inline: 0",
}

func (t ScrollMarginInlineStyle) Utilities() map[string]string {
	return ScrollMarginInlineStyleUtilitiesMapping
}

func (t ScrollMarginInlineStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginInlineStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformScaleY represent the CSS style "transform-scale-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformScaleYStyle string

func (t TransformScaleYStyle) Name() string {
	return "transform-scale-y"
}

var TransformScaleYStyleBrowserVariantsList = []string{}

func (t TransformScaleYStyle) BrowserVariants() []string {
	return TransformScaleYStyleBrowserVariantsList
}

var TransformScaleYStyleUtilitiesMapping = map[string]string{
	"transform-scale-y": "transform-scale-y: scaleY(0)",
}

func (t TransformScaleYStyle) Utilities() map[string]string {
	return TransformScaleYStyleUtilitiesMapping
}

func (t TransformScaleYStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformScaleYStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformTranslateY represent the CSS style "transform-translate-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformTranslateYStyle string

func (t TransformTranslateYStyle) Name() string {
	return "transform-translate-y"
}

var TransformTranslateYStyleBrowserVariantsList = []string{}

func (t TransformTranslateYStyle) BrowserVariants() []string {
	return TransformTranslateYStyleBrowserVariantsList
}

var TransformTranslateYStyleUtilitiesMapping = map[string]string{
	"transform-translate-y": "transform-translate-y: translateY(0)",
}

func (t TransformTranslateYStyle) Utilities() map[string]string {
	return TransformTranslateYStyleUtilitiesMapping
}

func (t TransformTranslateYStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformTranslateYStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Left represent the CSS style "left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/left
type LeftStyle string

func (t LeftStyle) Name() string {
	return "left"
}

var LeftStyleBrowserVariantsList = []string{}

func (t LeftStyle) BrowserVariants() []string {
	return LeftStyleBrowserVariantsList
}

var LeftStyleUtilitiesMapping = map[string]string{
	"left": "left: auto",
}

func (t LeftStyle) Utilities() map[string]string {
	return LeftStyleUtilitiesMapping
}

func (t LeftStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = LeftStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskMode represent the CSS style "mask-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-mode
type MaskModeStyle string

func (t MaskModeStyle) Name() string {
	return "mask-mode"
}

var MaskModeStyleBrowserVariantsList = []string{}

func (t MaskModeStyle) BrowserVariants() []string {
	return MaskModeStyleBrowserVariantsList
}

var MaskModeStyleUtilitiesMapping = map[string]string{
	"mask-mode": "mask-mode: match-source",
}

func (t MaskModeStyle) Utilities() map[string]string {
	return MaskModeStyleUtilitiesMapping
}

func (t MaskModeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskModeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaddingBlock represent the CSS style "padding-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-block
type PaddingBlockStyle string

func (t PaddingBlockStyle) Name() string {
	return "padding-block"
}

var PaddingBlockStyleBrowserVariantsList = []string{}

func (t PaddingBlockStyle) BrowserVariants() []string {
	return PaddingBlockStyleBrowserVariantsList
}

var PaddingBlockStyleUtilitiesMapping = map[string]string{
	"padding-block": "padding-block: 0",
}

func (t PaddingBlockStyle) Utilities() map[string]string {
	return PaddingBlockStyleUtilitiesMapping
}

func (t PaddingBlockStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaddingBlockStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FillOpacity represent the CSS style "fill-opacity" with value ""
// See
type FillOpacityStyle float64

func (t FillOpacityStyle) Name() string {
	return "fill-opacity"
}

var FillOpacityStyleBrowserVariantsList = []string{}

func (t FillOpacityStyle) BrowserVariants() []string {
	return FillOpacityStyleBrowserVariantsList
}

var FillOpacityStyleUtilitiesMapping = map[string]string{
	"fill-opacity": "fill-opacity: 1",
}

func (t FillOpacityStyle) Utilities() map[string]string {
	return FillOpacityStyleUtilitiesMapping
}

func (t FillOpacityStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FillOpacityStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Stroke represent the CSS style "stroke" with value ""
// See
type StrokeStyle Color

func (t StrokeStyle) Name() string {
	return "stroke"
}

var StrokeStyleBrowserVariantsList = []string{}

func (t StrokeStyle) BrowserVariants() []string {
	return StrokeStyleBrowserVariantsList
}

var StrokeStyleUtilitiesMapping = map[string]string{
	"stroke": "stroke: none",
}

func (t StrokeStyle) Utilities() map[string]string {
	return StrokeStyleUtilitiesMapping
}

func (t StrokeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = StrokeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BreakAfter represent the CSS style "break-after" with value "auto | avoid | always | all | avoid-page | page | left | right | recto | verso | avoid-column | column | avoid-region | region"
// See https://developer.mozilla.org/docs/Web/CSS/break-after
type BreakAfterStyle string

func (t BreakAfterStyle) Name() string {
	return "break-after"
}

const BreakAfterStyleAuto BreakAfterStyle = "auto"

const BreakAfterStyleAvoid BreakAfterStyle = "avoid"

const BreakAfterStyleAlways BreakAfterStyle = "always"

const BreakAfterStyleAll BreakAfterStyle = "all"

const BreakAfterStyleAvoidPage BreakAfterStyle = "avoid-page"

const BreakAfterStylePage BreakAfterStyle = "page"

const BreakAfterStyleLeft BreakAfterStyle = "left"

const BreakAfterStyleRight BreakAfterStyle = "right"

const BreakAfterStyleRecto BreakAfterStyle = "recto"

const BreakAfterStyleVerso BreakAfterStyle = "verso"

const BreakAfterStyleAvoidColumn BreakAfterStyle = "avoid-column"

const BreakAfterStyleColumn BreakAfterStyle = "column"

const BreakAfterStyleAvoidRegion BreakAfterStyle = "avoid-region"

const BreakAfterStyleRegion BreakAfterStyle = "region"

var BreakAfterStyleBrowserVariantsList = []string{}

func (t BreakAfterStyle) BrowserVariants() []string {
	return BreakAfterStyleBrowserVariantsList
}

var BreakAfterStyleUtilitiesMapping = map[string]string{
	"break-after":              "break-after: auto",
	"break-after-auto":         "break-after: auto",
	"break-after-avoid":        "break-after: avoid",
	"break-after-always":       "break-after: always",
	"break-after-all":          "break-after: all",
	"break-after-avoid-page":   "break-after: avoid-page",
	"break-after-page":         "break-after: page",
	"break-after-left":         "break-after: left",
	"break-after-right":        "break-after: right",
	"break-after-recto":        "break-after: recto",
	"break-after-verso":        "break-after: verso",
	"break-after-avoid-column": "break-after: avoid-column",
	"break-after-column":       "break-after: column",
	"break-after-avoid-region": "break-after: avoid-region",
	"break-after-region":       "break-after: region",
}

func (t BreakAfterStyle) Utilities() map[string]string {
	return BreakAfterStyleUtilitiesMapping
}

func (t BreakAfterStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BreakAfterStyleUtilitiesMapping[tu]
	return value, hasValue
}

// InitialLetterAlign represent the CSS style "initial-letter-align" with value "auto | alphabetic | hanging | ideographic"
// See https://developer.mozilla.org/docs/Web/CSS/initial-letter-align
type InitialLetterAlignStyle string

func (t InitialLetterAlignStyle) Name() string {
	return "initial-letter-align"
}

const InitialLetterAlignStyleAuto InitialLetterAlignStyle = "auto"

const InitialLetterAlignStyleAlphabetic InitialLetterAlignStyle = "alphabetic"

const InitialLetterAlignStyleHanging InitialLetterAlignStyle = "hanging"

const InitialLetterAlignStyleIdeographic InitialLetterAlignStyle = "ideographic"

var InitialLetterAlignStyleBrowserVariantsList = []string{}

func (t InitialLetterAlignStyle) BrowserVariants() []string {
	return InitialLetterAlignStyleBrowserVariantsList
}

var InitialLetterAlignStyleUtilitiesMapping = map[string]string{
	"initial-letter-align":             "initial-letter-align: auto",
	"initial-letter-align-auto":        "initial-letter-align: auto",
	"initial-letter-align-alphabetic":  "initial-letter-align: alphabetic",
	"initial-letter-align-hanging":     "initial-letter-align: hanging",
	"initial-letter-align-ideographic": "initial-letter-align: ideographic",
}

func (t InitialLetterAlignStyle) Utilities() map[string]string {
	return InitialLetterAlignStyleUtilitiesMapping
}

func (t InitialLetterAlignStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = InitialLetterAlignStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextAlignLast represent the CSS style "text-align-last" with value "auto | start | end | left | right | center | justify"
// See https://developer.mozilla.org/docs/Web/CSS/text-align-last
type TextAlignLastStyle string

func (t TextAlignLastStyle) Name() string {
	return "text-align-last"
}

const TextAlignLastStyleAuto TextAlignLastStyle = "auto"

const TextAlignLastStyleStart TextAlignLastStyle = "start"

const TextAlignLastStyleEnd TextAlignLastStyle = "end"

const TextAlignLastStyleLeft TextAlignLastStyle = "left"

const TextAlignLastStyleRight TextAlignLastStyle = "right"

const TextAlignLastStyleCenter TextAlignLastStyle = "center"

const TextAlignLastStyleJustify TextAlignLastStyle = "justify"

var TextAlignLastStyleBrowserVariantsList = []string{}

func (t TextAlignLastStyle) BrowserVariants() []string {
	return TextAlignLastStyleBrowserVariantsList
}

var TextAlignLastStyleUtilitiesMapping = map[string]string{
	"text-align-last":         "text-align-last: auto",
	"text-align-last-auto":    "text-align-last: auto",
	"text-align-last-start":   "text-align-last: start",
	"text-align-last-end":     "text-align-last: end",
	"text-align-last-left":    "text-align-last: left",
	"text-align-last-right":   "text-align-last: right",
	"text-align-last-center":  "text-align-last: center",
	"text-align-last-justify": "text-align-last: justify",
}

func (t TextAlignLastStyle) Utilities() map[string]string {
	return TextAlignLastStyleUtilitiesMapping
}

func (t TextAlignLastStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextAlignLastStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontSizeAdjust represent the CSS style "font-size-adjust" with value "none"
// See https://developer.mozilla.org/docs/Web/CSS/font-size-adjust
type FontSizeAdjustStyle string

func (t FontSizeAdjustStyle) Name() string {
	return "font-size-adjust"
}

const FontSizeAdjustStyleNone FontSizeAdjustStyle = "none"

var FontSizeAdjustStyleBrowserVariantsList = []string{}

func (t FontSizeAdjustStyle) BrowserVariants() []string {
	return FontSizeAdjustStyleBrowserVariantsList
}

var FontSizeAdjustStyleUtilitiesMapping = map[string]string{
	"font-size-adjust":      "font-size-adjust: none",
	"font-size-adjust-none": "font-size-adjust: none",
}

func (t FontSizeAdjustStyle) Utilities() map[string]string {
	return FontSizeAdjustStyleUtilitiesMapping
}

func (t FontSizeAdjustStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontSizeAdjustStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollBehavior represent the CSS style "scroll-behavior" with value "auto | smooth"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-behavior
type ScrollBehaviorStyle string

func (t ScrollBehaviorStyle) Name() string {
	return "scroll-behavior"
}

const ScrollBehaviorStyleAuto ScrollBehaviorStyle = "auto"

const ScrollBehaviorStyleSmooth ScrollBehaviorStyle = "smooth"

var ScrollBehaviorStyleBrowserVariantsList = []string{}

func (t ScrollBehaviorStyle) BrowserVariants() []string {
	return ScrollBehaviorStyleBrowserVariantsList
}

var ScrollBehaviorStyleUtilitiesMapping = map[string]string{
	"scroll-behavior":        "scroll-behavior: auto",
	"scroll-behavior-auto":   "scroll-behavior: auto",
	"scroll-behavior-smooth": "scroll-behavior: smooth",
}

func (t ScrollBehaviorStyle) Utilities() map[string]string {
	return ScrollBehaviorStyleUtilitiesMapping
}

func (t ScrollBehaviorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollBehaviorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// StopOpacity represent the CSS style "stop-opacity" with value ""
// See
type StopOpacityStyle float64

func (t StopOpacityStyle) Name() string {
	return "stop-opacity"
}

var StopOpacityStyleBrowserVariantsList = []string{}

func (t StopOpacityStyle) BrowserVariants() []string {
	return StopOpacityStyleBrowserVariantsList
}

var StopOpacityStyleUtilitiesMapping = map[string]string{
	"stop-opacity": "stop-opacity: 1",
}

func (t StopOpacityStyle) Utilities() map[string]string {
	return StopOpacityStyleUtilitiesMapping
}

func (t StopOpacityStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = StopOpacityStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackfaceVisibility represent the CSS style "backface-visibility" with value "visible | hidden"
// See https://developer.mozilla.org/docs/Web/CSS/backface-visibility
type BackfaceVisibilityStyle string

func (t BackfaceVisibilityStyle) Name() string {
	return "backface-visibility"
}

const BackfaceVisibilityStyleVisible BackfaceVisibilityStyle = "visible"

const BackfaceVisibilityStyleHidden BackfaceVisibilityStyle = "hidden"

var BackfaceVisibilityStyleBrowserVariantsList = []string{}

func (t BackfaceVisibilityStyle) BrowserVariants() []string {
	return BackfaceVisibilityStyleBrowserVariantsList
}

var BackfaceVisibilityStyleUtilitiesMapping = map[string]string{
	"backface-visibility":         "backface-visibility: visible",
	"backface-visibility-visible": "backface-visibility: visible",
	"backface-visibility-hidden":  "backface-visibility: hidden",
}

func (t BackfaceVisibilityStyle) Utilities() map[string]string {
	return BackfaceVisibilityStyleUtilitiesMapping
}

func (t BackfaceVisibilityStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackfaceVisibilityStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderLeftStyle represent the CSS style "border-left-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-left-style
type BorderLeftStyleStyle string

func (t BorderLeftStyleStyle) Name() string {
	return "border-left-style"
}

var BorderLeftStyleStyleBrowserVariantsList = []string{}

func (t BorderLeftStyleStyle) BrowserVariants() []string {
	return BorderLeftStyleStyleBrowserVariantsList
}

var BorderLeftStyleStyleUtilitiesMapping = map[string]string{
	"border-left-style": "border-left-style: none",
}

func (t BorderLeftStyleStyle) Utilities() map[string]string {
	return BorderLeftStyleStyleUtilitiesMapping
}

func (t BorderLeftStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderLeftStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// JustifyTracks represent the CSS style "justify-tracks" with value "auto | normal | stretch | end | start | flex-start | flex-end | center | left | right |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe center | unsafe center"
// See https://developer.mozilla.org/docs/Web/CSS/justify-tracks
type JustifyTracksStyle string

func (t JustifyTracksStyle) Name() string {
	return "justify-tracks"
}

const JustifyTracksStyleAuto JustifyTracksStyle = "auto"

const JustifyTracksStyleNormal JustifyTracksStyle = "normal"

const JustifyTracksStyleStretch JustifyTracksStyle = "stretch"

const JustifyTracksStyleEnd JustifyTracksStyle = "end"

const JustifyTracksStyleStart JustifyTracksStyle = "start"

const JustifyTracksStyleFlexStart JustifyTracksStyle = "flex-start"

const JustifyTracksStyleFlexEnd JustifyTracksStyle = "flex-end"

const JustifyTracksStyleCenter JustifyTracksStyle = "center"

const JustifyTracksStyleLeft JustifyTracksStyle = "left"

const JustifyTracksStyleRight JustifyTracksStyle = "right"

const JustifyTracksStyleBaseline JustifyTracksStyle = "baseline"

const JustifyTracksStyleFirstBaseline JustifyTracksStyle = "first-baseline"

const JustifyTracksStyleLastBaseline JustifyTracksStyle = "last-baseline"

const JustifyTracksStyleSpaceBetween JustifyTracksStyle = "space-between"

const JustifyTracksStyleSpaceAround JustifyTracksStyle = "space-around"

const JustifyTracksStyleSpaceEvenly JustifyTracksStyle = "space-evenly"

const JustifyTracksStyleSafeCenter JustifyTracksStyle = "safe-center"

const JustifyTracksStyleUnsafeCenter JustifyTracksStyle = "unsafe-center"

var JustifyTracksStyleBrowserVariantsList = []string{}

func (t JustifyTracksStyle) BrowserVariants() []string {
	return JustifyTracksStyleBrowserVariantsList
}

var JustifyTracksStyleUtilitiesMapping = map[string]string{
	"justify-tracks":                "justify-tracks: normal",
	"justify-tracks-auto":           "justify-tracks: auto",
	"justify-tracks-normal":         "justify-tracks: normal",
	"justify-tracks-stretch":        "justify-tracks: stretch",
	"justify-tracks-end":            "justify-tracks: end",
	"justify-tracks-start":          "justify-tracks: start",
	"justify-tracks-flex-start":     "justify-tracks: flex-start",
	"justify-tracks-flex-end":       "justify-tracks: flex-end",
	"justify-tracks-center":         "justify-tracks: center",
	"justify-tracks-left":           "justify-tracks: left",
	"justify-tracks-right":          "justify-tracks: right",
	"justify-tracks-baseline":       "justify-tracks: baseline",
	"justify-tracks-first-baseline": "justify-tracks: first-baseline",
	"justify-tracks-last-baseline":  "justify-tracks: last-baseline",
	"justify-tracks-space-between":  "justify-tracks: space-between",
	"justify-tracks-space-around":   "justify-tracks: space-around",
	"justify-tracks-space-evenly":   "justify-tracks: space-evenly",
	"justify-tracks-safe-center":    "justify-tracks: safe-center",
	"justify-tracks-unsafe-center":  "justify-tracks: unsafe-center",
}

func (t JustifyTracksStyle) Utilities() map[string]string {
	return JustifyTracksStyleUtilitiesMapping
}

func (t JustifyTracksStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = JustifyTracksStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskBorderMode represent the CSS style "mask-border-mode" with value "luminance | alpha"
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-mode
type MaskBorderModeStyle string

func (t MaskBorderModeStyle) Name() string {
	return "mask-border-mode"
}

const MaskBorderModeStyleLuminance MaskBorderModeStyle = "luminance"

const MaskBorderModeStyleAlpha MaskBorderModeStyle = "alpha"

var MaskBorderModeStyleBrowserVariantsList = []string{}

func (t MaskBorderModeStyle) BrowserVariants() []string {
	return MaskBorderModeStyleBrowserVariantsList
}

var MaskBorderModeStyleUtilitiesMapping = map[string]string{
	"mask-border-mode":           "mask-border-mode: alpha",
	"mask-border-mode-luminance": "mask-border-mode: luminance",
	"mask-border-mode-alpha":     "mask-border-mode: alpha",
}

func (t MaskBorderModeStyle) Utilities() map[string]string {
	return MaskBorderModeStyleUtilitiesMapping
}

func (t MaskBorderModeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskBorderModeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MixBlendMode represent the CSS style "mix-blend-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mix-blend-mode
type MixBlendModeStyle string

func (t MixBlendModeStyle) Name() string {
	return "mix-blend-mode"
}

var MixBlendModeStyleBrowserVariantsList = []string{}

func (t MixBlendModeStyle) BrowserVariants() []string {
	return MixBlendModeStyleBrowserVariantsList
}

var MixBlendModeStyleUtilitiesMapping = map[string]string{
	"mix-blend-mode": "mix-blend-mode: normal",
}

func (t MixBlendModeStyle) Utilities() map[string]string {
	return MixBlendModeStyleUtilitiesMapping
}

func (t MixBlendModeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MixBlendModeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollSnapTypeX represent the CSS style "scroll-snap-type-x" with value "none | mandatory | proximity"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-type-x
type ScrollSnapTypeXStyle string

func (t ScrollSnapTypeXStyle) Name() string {
	return "scroll-snap-type-x"
}

const ScrollSnapTypeXStyleNone ScrollSnapTypeXStyle = "none"

const ScrollSnapTypeXStyleMandatory ScrollSnapTypeXStyle = "mandatory"

const ScrollSnapTypeXStyleProximity ScrollSnapTypeXStyle = "proximity"

var ScrollSnapTypeXStyleBrowserVariantsList = []string{}

func (t ScrollSnapTypeXStyle) BrowserVariants() []string {
	return ScrollSnapTypeXStyleBrowserVariantsList
}

var ScrollSnapTypeXStyleUtilitiesMapping = map[string]string{
	"scroll-snap-type-x":           "scroll-snap-type-x: none",
	"scroll-snap-type-x-none":      "scroll-snap-type-x: none",
	"scroll-snap-type-x-mandatory": "scroll-snap-type-x: mandatory",
	"scroll-snap-type-x-proximity": "scroll-snap-type-x: proximity",
}

func (t ScrollSnapTypeXStyle) Utilities() map[string]string {
	return ScrollSnapTypeXStyleUtilitiesMapping
}

func (t ScrollSnapTypeXStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollSnapTypeXStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundRepeat represent the CSS style "background-repeat" with value "none | repeat | no-repeat"
// See https://developer.mozilla.org/docs/Web/CSS/background-repeat
type BackgroundRepeatStyle string

func (t BackgroundRepeatStyle) Name() string {
	return "background-repeat"
}

const BackgroundRepeatStyleNone BackgroundRepeatStyle = "none"

const BackgroundRepeatStyleRepeat BackgroundRepeatStyle = "repeat"

const BackgroundRepeatStyleNoRepeat BackgroundRepeatStyle = "no-repeat"

var BackgroundRepeatStyleBrowserVariantsList = []string{}

func (t BackgroundRepeatStyle) BrowserVariants() []string {
	return BackgroundRepeatStyleBrowserVariantsList
}

var BackgroundRepeatStyleUtilitiesMapping = map[string]string{
	"background-repeat":           "background-repeat: repeat",
	"background-repeat-none":      "background-repeat: none",
	"background-repeat-repeat":    "background-repeat: repeat",
	"background-repeat-no-repeat": "background-repeat: no-repeat",
}

func (t BackgroundRepeatStyle) Utilities() map[string]string {
	return BackgroundRepeatStyleUtilitiesMapping
}

func (t BackgroundRepeatStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundRepeatStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PerspectiveOrigin represent the CSS style "perspective-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/perspective-origin
type PerspectiveOriginStyle string

func (t PerspectiveOriginStyle) Name() string {
	return "perspective-origin"
}

var PerspectiveOriginStyleBrowserVariantsList = []string{}

func (t PerspectiveOriginStyle) BrowserVariants() []string {
	return PerspectiveOriginStyleBrowserVariantsList
}

var PerspectiveOriginStyleUtilitiesMapping = map[string]string{
	"perspective-origin": "perspective-origin: 50% 50%",
}

func (t PerspectiveOriginStyle) Utilities() map[string]string {
	return PerspectiveOriginStyleUtilitiesMapping
}

func (t PerspectiveOriginStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PerspectiveOriginStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollbarGutter represent the CSS style "scrollbar-gutter" with value "auto |  stable | always | both | force"
// See https://developer.mozilla.org/docs/Web/CSS/scrollbar-gutter
type ScrollbarGutterStyle string

func (t ScrollbarGutterStyle) Name() string {
	return "scrollbar-gutter"
}

const ScrollbarGutterStyleAuto ScrollbarGutterStyle = "auto"

const ScrollbarGutterStyleStable ScrollbarGutterStyle = "stable"

const ScrollbarGutterStyleAlways ScrollbarGutterStyle = "always"

const ScrollbarGutterStyleBoth ScrollbarGutterStyle = "both"

const ScrollbarGutterStyleForce ScrollbarGutterStyle = "force"

var ScrollbarGutterStyleBrowserVariantsList = []string{}

func (t ScrollbarGutterStyle) BrowserVariants() []string {
	return ScrollbarGutterStyleBrowserVariantsList
}

var ScrollbarGutterStyleUtilitiesMapping = map[string]string{
	"scrollbar-gutter":        "scrollbar-gutter: auto",
	"scrollbar-gutter-auto":   "scrollbar-gutter: auto",
	"scrollbar-gutter-stable": "scrollbar-gutter: stable",
	"scrollbar-gutter-always": "scrollbar-gutter: always",
	"scrollbar-gutter-both":   "scrollbar-gutter: both",
	"scrollbar-gutter-force":  "scrollbar-gutter: force",
}

func (t ScrollbarGutterStyle) Utilities() map[string]string {
	return ScrollbarGutterStyleUtilitiesMapping
}

func (t ScrollbarGutterStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollbarGutterStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FillRule represent the CSS style "fill-rule" with value "nonzero | evenodd"
// See
type FillRuleStyle string

func (t FillRuleStyle) Name() string {
	return "fill-rule"
}

const FillRuleStyleNonzero FillRuleStyle = "nonzero"

const FillRuleStyleEvenodd FillRuleStyle = "evenodd"

var FillRuleStyleBrowserVariantsList = []string{}

func (t FillRuleStyle) BrowserVariants() []string {
	return FillRuleStyleBrowserVariantsList
}

var FillRuleStyleUtilitiesMapping = map[string]string{
	"fill-rule":         "fill-rule: nonzero",
	"fill-rule-nonzero": "fill-rule: nonzero",
	"fill-rule-evenodd": "fill-rule: evenodd",
}

func (t FillRuleStyle) Utilities() map[string]string {
	return FillRuleStyleUtilitiesMapping
}

func (t FillRuleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FillRuleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Width represent the CSS style "width" with value "auto |  min-content | max-content"
// See https://developer.mozilla.org/docs/Web/CSS/width
type WidthStyle string

func (t WidthStyle) Name() string {
	return "width"
}

const WidthStyleAuto WidthStyle = "auto"

const WidthStyleMinContent WidthStyle = "min-content"

const WidthStyleMaxContent WidthStyle = "max-content"

var WidthStyleBrowserVariantsList = []string{}

func (t WidthStyle) BrowserVariants() []string {
	return WidthStyleBrowserVariantsList
}

var WidthStyleUtilitiesMapping = map[string]string{
	"width":             "width: auto",
	"width-auto":        "width: auto",
	"width-min-content": "width: min-content",
	"width-max-content": "width: max-content",
}

func (t WidthStyle) Utilities() map[string]string {
	return WidthStyleUtilitiesMapping
}

func (t WidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = WidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundPositionY represent the CSS style "background-position-y" with value "center | top | bottom | y-start | y-end"
// See https://developer.mozilla.org/docs/Web/CSS/background-position-y
type BackgroundPositionYStyle string

func (t BackgroundPositionYStyle) Name() string {
	return "background-position-y"
}

const BackgroundPositionYStyleCenter BackgroundPositionYStyle = "center"

const BackgroundPositionYStyleTop BackgroundPositionYStyle = "top"

const BackgroundPositionYStyleBottom BackgroundPositionYStyle = "bottom"

const BackgroundPositionYStyleYStart BackgroundPositionYStyle = "y-start"

const BackgroundPositionYStyleYEnd BackgroundPositionYStyle = "y-end"

var BackgroundPositionYStyleBrowserVariantsList = []string{}

func (t BackgroundPositionYStyle) BrowserVariants() []string {
	return BackgroundPositionYStyleBrowserVariantsList
}

var BackgroundPositionYStyleUtilitiesMapping = map[string]string{
	"background-position-y":         "background-position-y: top",
	"background-position-y-center":  "background-position-y: center",
	"background-position-y-top":     "background-position-y: top",
	"background-position-y-bottom":  "background-position-y: bottom",
	"background-position-y-y-start": "background-position-y: y-start",
	"background-position-y-y-end":   "background-position-y: y-end",
}

func (t BackgroundPositionYStyle) Utilities() map[string]string {
	return BackgroundPositionYStyleUtilitiesMapping
}

func (t BackgroundPositionYStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundPositionYStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderInlineEndStyle represent the CSS style "border-inline-end-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-end-style
type BorderInlineEndStyleStyle string

func (t BorderInlineEndStyleStyle) Name() string {
	return "border-inline-end-style"
}

var BorderInlineEndStyleStyleBrowserVariantsList = []string{}

func (t BorderInlineEndStyleStyle) BrowserVariants() []string {
	return BorderInlineEndStyleStyleBrowserVariantsList
}

var BorderInlineEndStyleStyleUtilitiesMapping = map[string]string{
	"border-inline-end-style": "border-inline-end-style: none",
}

func (t BorderInlineEndStyleStyle) Utilities() map[string]string {
	return BorderInlineEndStyleStyleUtilitiesMapping
}

func (t BorderInlineEndStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderInlineEndStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginBlockStart represent the CSS style "margin-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-block-start
type MarginBlockStartStyle string

func (t MarginBlockStartStyle) Name() string {
	return "margin-block-start"
}

var MarginBlockStartStyleBrowserVariantsList = []string{}

func (t MarginBlockStartStyle) BrowserVariants() []string {
	return MarginBlockStartStyleBrowserVariantsList
}

var MarginBlockStartStyleUtilitiesMapping = map[string]string{
	"margin-block-start": "margin-block-start: 0",
}

func (t MarginBlockStartStyle) Utilities() map[string]string {
	return MarginBlockStartStyleUtilitiesMapping
}

func (t MarginBlockStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginBlockStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TabSize represent the CSS style "tab-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/tab-size
type TabSizeStyle float64

func (t TabSizeStyle) Name() string {
	return "tab-size"
}

var TabSizeStyleBrowserVariantsList = []string{}

func (t TabSizeStyle) BrowserVariants() []string {
	return TabSizeStyleBrowserVariantsList
}

var TabSizeStyleUtilitiesMapping = map[string]string{
	"tab-size": "tab-size: 8",
}

func (t TabSizeStyle) Utilities() map[string]string {
	return TabSizeStyleUtilitiesMapping
}

func (t TabSizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TabSizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderSpacing represent the CSS style "border-spacing" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-spacing
type BorderSpacingStyle string

func (t BorderSpacingStyle) Name() string {
	return "border-spacing"
}

var BorderSpacingStyleBrowserVariantsList = []string{}

func (t BorderSpacingStyle) BrowserVariants() []string {
	return BorderSpacingStyleBrowserVariantsList
}

var BorderSpacingStyleUtilitiesMapping = map[string]string{
	"border-spacing": "border-spacing: 0",
}

func (t BorderSpacingStyle) Utilities() map[string]string {
	return BorderSpacingStyleUtilitiesMapping
}

func (t BorderSpacingStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderSpacingStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskType represent the CSS style "mask-type" with value "luminance | alpha"
// See https://developer.mozilla.org/docs/Web/CSS/mask-type
type MaskTypeStyle string

func (t MaskTypeStyle) Name() string {
	return "mask-type"
}

const MaskTypeStyleLuminance MaskTypeStyle = "luminance"

const MaskTypeStyleAlpha MaskTypeStyle = "alpha"

var MaskTypeStyleBrowserVariantsList = []string{}

func (t MaskTypeStyle) BrowserVariants() []string {
	return MaskTypeStyleBrowserVariantsList
}

var MaskTypeStyleUtilitiesMapping = map[string]string{
	"mask-type":           "mask-type: luminance",
	"mask-type-luminance": "mask-type: luminance",
	"mask-type-alpha":     "mask-type: alpha",
}

func (t MaskTypeStyle) Utilities() map[string]string {
	return MaskTypeStyleUtilitiesMapping
}

func (t MaskTypeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskTypeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaddingLeft represent the CSS style "padding-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-left
type PaddingLeftStyle string

func (t PaddingLeftStyle) Name() string {
	return "padding-left"
}

var PaddingLeftStyleBrowserVariantsList = []string{}

func (t PaddingLeftStyle) BrowserVariants() []string {
	return PaddingLeftStyleBrowserVariantsList
}

var PaddingLeftStyleUtilitiesMapping = map[string]string{
	"padding-left": "padding-left: 0",
}

func (t PaddingLeftStyle) Utilities() map[string]string {
	return PaddingLeftStyleUtilitiesMapping
}

func (t PaddingLeftStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaddingLeftStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextEmphasisPositionSecond represent the CSS style "text-emphasis-position-second" with value "right | left"
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-position
type TextEmphasisPositionSecondStyle string

func (t TextEmphasisPositionSecondStyle) Name() string {
	return "text-emphasis-position-second"
}

const TextEmphasisPositionSecondStyleRight TextEmphasisPositionSecondStyle = "right"

const TextEmphasisPositionSecondStyleLeft TextEmphasisPositionSecondStyle = "left"

var TextEmphasisPositionSecondStyleBrowserVariantsList = []string{}

func (t TextEmphasisPositionSecondStyle) BrowserVariants() []string {
	return TextEmphasisPositionSecondStyleBrowserVariantsList
}

var TextEmphasisPositionSecondStyleUtilitiesMapping = map[string]string{
	"text-emphasis-position-second":       "text-emphasis-position-second: over right",
	"text-emphasis-position-second-right": "text-emphasis-position-second: right",
	"text-emphasis-position-second-left":  "text-emphasis-position-second: left",
}

func (t TextEmphasisPositionSecondStyle) Utilities() map[string]string {
	return TextEmphasisPositionSecondStyleUtilitiesMapping
}

func (t TextEmphasisPositionSecondStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextEmphasisPositionSecondStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AnimationDuration represent the CSS style "animation-duration" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-duration
type AnimationDurationStyle time.Duration

func (t AnimationDurationStyle) Name() string {
	return "animation-duration"
}

var AnimationDurationStyleBrowserVariantsList = []string{}

func (t AnimationDurationStyle) BrowserVariants() []string {
	return AnimationDurationStyleBrowserVariantsList
}

var AnimationDurationStyleUtilitiesMapping = map[string]string{
	"animation-duration": "animation-duration: 0s",
}

func (t AnimationDurationStyle) Utilities() map[string]string {
	return AnimationDurationStyleUtilitiesMapping
}

func (t AnimationDurationStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AnimationDurationStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundOrigin represent the CSS style "background-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-origin
type BackgroundOriginStyle string

func (t BackgroundOriginStyle) Name() string {
	return "background-origin"
}

var BackgroundOriginStyleBrowserVariantsList = []string{}

func (t BackgroundOriginStyle) BrowserVariants() []string {
	return BackgroundOriginStyleBrowserVariantsList
}

var BackgroundOriginStyleUtilitiesMapping = map[string]string{
	"background-origin": "background-origin: padding-box",
}

func (t BackgroundOriginStyle) Utilities() map[string]string {
	return BackgroundOriginStyleUtilitiesMapping
}

func (t BackgroundOriginStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundOriginStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Zoom represent the CSS style "zoom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/zoom
type ZoomStyle string

func (t ZoomStyle) Name() string {
	return "zoom"
}

var ZoomStyleBrowserVariantsList = []string{}

func (t ZoomStyle) BrowserVariants() []string {
	return ZoomStyleBrowserVariantsList
}

var ZoomStyleUtilitiesMapping = map[string]string{
	"zoom": "zoom: normal",
}

func (t ZoomStyle) Utilities() map[string]string {
	return ZoomStyleUtilitiesMapping
}

func (t ZoomStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ZoomStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Display represent the CSS style "display" with value "contents | none | inline-block | inline-table | inline-flex| inline-grid | | block | inline | run-in | flow | flow-root | table | flex | grid | ruby | list-item | table-row-group | table-header-group | table-footer-group | table-row | table-cell | table-column-group | table-column | table-caption | ruby-base | ruby-text | ruby-base-container | ruby-text-container"
// See https://developer.mozilla.org/docs/Web/CSS/display
type DisplayStyle string

func (t DisplayStyle) Name() string {
	return "display"
}

const DisplayStyleContents DisplayStyle = "contents"

const DisplayStyleNone DisplayStyle = "none"

const DisplayStyleInlineBlock DisplayStyle = "inline-block"

const DisplayStyleInlineTable DisplayStyle = "inline-table"

const DisplayStyleInlineFlex DisplayStyle = "inline-flex"

const DisplayStyleInlineGrid DisplayStyle = "inline-grid"

const DisplayStyleBlock DisplayStyle = "block"

const DisplayStyleInline DisplayStyle = "inline"

const DisplayStyleRunIn DisplayStyle = "run-in"

const DisplayStyleFlow DisplayStyle = "flow"

const DisplayStyleFlowRoot DisplayStyle = "flow-root"

const DisplayStyleTable DisplayStyle = "table"

const DisplayStyleFlex DisplayStyle = "flex"

const DisplayStyleGrid DisplayStyle = "grid"

const DisplayStyleRuby DisplayStyle = "ruby"

const DisplayStyleListItem DisplayStyle = "list-item"

const DisplayStyleTableRowGroup DisplayStyle = "table-row-group"

const DisplayStyleTableHeaderGroup DisplayStyle = "table-header-group"

const DisplayStyleTableFooterGroup DisplayStyle = "table-footer-group"

const DisplayStyleTableRow DisplayStyle = "table-row"

const DisplayStyleTableCell DisplayStyle = "table-cell"

const DisplayStyleTableColumnGroup DisplayStyle = "table-column-group"

const DisplayStyleTableColumn DisplayStyle = "table-column"

const DisplayStyleTableCaption DisplayStyle = "table-caption"

const DisplayStyleRubyBase DisplayStyle = "ruby-base"

const DisplayStyleRubyText DisplayStyle = "ruby-text"

const DisplayStyleRubyBaseContainer DisplayStyle = "ruby-base-container"

const DisplayStyleRubyTextContainer DisplayStyle = "ruby-text-container"

var DisplayStyleBrowserVariantsList = []string{}

func (t DisplayStyle) BrowserVariants() []string {
	return DisplayStyleBrowserVariantsList
}

var DisplayStyleUtilitiesMapping = map[string]string{
	"display":                     "display: inline",
	"display-contents":            "display: contents",
	"display-none":                "display: none",
	"display-inline-block":        "display: inline-block",
	"display-inline-table":        "display: inline-table",
	"display-inline-flex":         "display: inline-flex",
	"display-inline-grid":         "display: inline-grid",
	"display-block":               "display: block",
	"display-inline":              "display: inline",
	"display-run-in":              "display: run-in",
	"display-flow":                "display: flow",
	"display-flow-root":           "display: flow-root",
	"display-table":               "display: table",
	"display-flex":                "display: flex",
	"display-grid":                "display: grid",
	"display-ruby":                "display: ruby",
	"display-list-item":           "display: list-item",
	"display-table-row-group":     "display: table-row-group",
	"display-table-header-group":  "display: table-header-group",
	"display-table-footer-group":  "display: table-footer-group",
	"display-table-row":           "display: table-row",
	"display-table-cell":          "display: table-cell",
	"display-table-column-group":  "display: table-column-group",
	"display-table-column":        "display: table-column",
	"display-table-caption":       "display: table-caption",
	"display-ruby-base":           "display: ruby-base",
	"display-ruby-text":           "display: ruby-text",
	"display-ruby-base-container": "display: ruby-base-container",
	"display-ruby-text-container": "display: ruby-text-container",
}

func (t DisplayStyle) Utilities() map[string]string {
	return DisplayStyleUtilitiesMapping
}

func (t DisplayStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = DisplayStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MathStyle represent the CSS style "math-style" with value "normal | compact"
// See https://developer.mozilla.org/docs/Web/CSS/math-style
type MathStyleStyle string

func (t MathStyleStyle) Name() string {
	return "math-style"
}

const MathStyleStyleNormal MathStyleStyle = "normal"

const MathStyleStyleCompact MathStyleStyle = "compact"

var MathStyleStyleBrowserVariantsList = []string{}

func (t MathStyleStyle) BrowserVariants() []string {
	return MathStyleStyleBrowserVariantsList
}

var MathStyleStyleUtilitiesMapping = map[string]string{
	"math-style":         "math-style: normal",
	"math-style-normal":  "math-style: normal",
	"math-style-compact": "math-style: compact",
}

func (t MathStyleStyle) Utilities() map[string]string {
	return MathStyleStyleUtilitiesMapping
}

func (t MathStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MathStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMarginTop represent the CSS style "scroll-margin-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-top
type ScrollMarginTopStyle float64

func (t ScrollMarginTopStyle) Name() string {
	return "scroll-margin-top"
}

var ScrollMarginTopStyleBrowserVariantsList = []string{}

func (t ScrollMarginTopStyle) BrowserVariants() []string {
	return ScrollMarginTopStyleBrowserVariantsList
}

var ScrollMarginTopStyleUtilitiesMapping = map[string]string{
	"scroll-margin-top": "scroll-margin-top: 0",
}

func (t ScrollMarginTopStyle) Utilities() map[string]string {
	return ScrollMarginTopStyleUtilitiesMapping
}

func (t ScrollMarginTopStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginTopStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollPaddingBlock represent the CSS style "scroll-padding-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-block
type ScrollPaddingBlockStyle string

func (t ScrollPaddingBlockStyle) Name() string {
	return "scroll-padding-block"
}

var ScrollPaddingBlockStyleBrowserVariantsList = []string{}

func (t ScrollPaddingBlockStyle) BrowserVariants() []string {
	return ScrollPaddingBlockStyleBrowserVariantsList
}

var ScrollPaddingBlockStyleUtilitiesMapping = map[string]string{
	"scroll-padding-block": "scroll-padding-block: auto",
}

func (t ScrollPaddingBlockStyle) Utilities() map[string]string {
	return ScrollPaddingBlockStyleUtilitiesMapping
}

func (t ScrollPaddingBlockStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollPaddingBlockStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundColor represent the CSS style "background-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-color
type BackgroundColorStyle Color

func (t BackgroundColorStyle) Name() string {
	return "background-color"
}

var BackgroundColorStyleBrowserVariantsList = []string{}

func (t BackgroundColorStyle) BrowserVariants() []string {
	return BackgroundColorStyleBrowserVariantsList
}

var BackgroundColorStyleUtilitiesMapping = map[string]string{
	"background-color": "background-color: transparent",
}

func (t BackgroundColorStyle) Utilities() map[string]string {
	return BackgroundColorStyleUtilitiesMapping
}

func (t BackgroundColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderBottomStyle represent the CSS style "border-bottom-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-style
type BorderBottomStyleStyle string

func (t BorderBottomStyleStyle) Name() string {
	return "border-bottom-style"
}

var BorderBottomStyleStyleBrowserVariantsList = []string{}

func (t BorderBottomStyleStyle) BrowserVariants() []string {
	return BorderBottomStyleStyleBrowserVariantsList
}

var BorderBottomStyleStyleUtilitiesMapping = map[string]string{
	"border-bottom-style": "border-bottom-style: none",
}

func (t BorderBottomStyleStyle) Utilities() map[string]string {
	return BorderBottomStyleStyleUtilitiesMapping
}

func (t BorderBottomStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderBottomStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColumnRuleStyle represent the CSS style "column-rule-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-rule-style
type ColumnRuleStyleStyle string

func (t ColumnRuleStyleStyle) Name() string {
	return "column-rule-style"
}

var ColumnRuleStyleStyleBrowserVariantsList = []string{}

func (t ColumnRuleStyleStyle) BrowserVariants() []string {
	return ColumnRuleStyleStyleBrowserVariantsList
}

var ColumnRuleStyleStyleUtilitiesMapping = map[string]string{
	"column-rule-style": "column-rule-style: none",
}

func (t ColumnRuleStyleStyle) Utilities() map[string]string {
	return ColumnRuleStyleStyleUtilitiesMapping
}

func (t ColumnRuleStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColumnRuleStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Translate represent the CSS style "translate" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/translate
type TranslateStyle string

func (t TranslateStyle) Name() string {
	return "translate"
}

var TranslateStyleBrowserVariantsList = []string{}

func (t TranslateStyle) BrowserVariants() []string {
	return TranslateStyleBrowserVariantsList
}

var TranslateStyleUtilitiesMapping = map[string]string{
	"translate": "translate: none",
}

func (t TranslateStyle) Utilities() map[string]string {
	return TranslateStyleUtilitiesMapping
}

func (t TranslateStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TranslateStyleUtilitiesMapping[tu]
	return value, hasValue
}

// DominantBaseline represent the CSS style "dominant-baseline" with value "auto | use-script | no-change | reset-size | ideographic | alphabetic | hanging | mathematical | central | middle | text-after-edge | text-before-edge"
// See
type DominantBaselineStyle string

func (t DominantBaselineStyle) Name() string {
	return "dominant-baseline"
}

const DominantBaselineStyleAuto DominantBaselineStyle = "auto"

const DominantBaselineStyleUseScript DominantBaselineStyle = "use-script"

const DominantBaselineStyleNoChange DominantBaselineStyle = "no-change"

const DominantBaselineStyleResetSize DominantBaselineStyle = "reset-size"

const DominantBaselineStyleIdeographic DominantBaselineStyle = "ideographic"

const DominantBaselineStyleAlphabetic DominantBaselineStyle = "alphabetic"

const DominantBaselineStyleHanging DominantBaselineStyle = "hanging"

const DominantBaselineStyleMathematical DominantBaselineStyle = "mathematical"

const DominantBaselineStyleCentral DominantBaselineStyle = "central"

const DominantBaselineStyleMiddle DominantBaselineStyle = "middle"

const DominantBaselineStyleTextAfterEdge DominantBaselineStyle = "text-after-edge"

const DominantBaselineStyleTextBeforeEdge DominantBaselineStyle = "text-before-edge"

var DominantBaselineStyleBrowserVariantsList = []string{}

func (t DominantBaselineStyle) BrowserVariants() []string {
	return DominantBaselineStyleBrowserVariantsList
}

var DominantBaselineStyleUtilitiesMapping = map[string]string{
	"dominant-baseline":                  "dominant-baseline: auto",
	"dominant-baseline-auto":             "dominant-baseline: auto",
	"dominant-baseline-use-script":       "dominant-baseline: use-script",
	"dominant-baseline-no-change":        "dominant-baseline: no-change",
	"dominant-baseline-reset-size":       "dominant-baseline: reset-size",
	"dominant-baseline-ideographic":      "dominant-baseline: ideographic",
	"dominant-baseline-alphabetic":       "dominant-baseline: alphabetic",
	"dominant-baseline-hanging":          "dominant-baseline: hanging",
	"dominant-baseline-mathematical":     "dominant-baseline: mathematical",
	"dominant-baseline-central":          "dominant-baseline: central",
	"dominant-baseline-middle":           "dominant-baseline: middle",
	"dominant-baseline-text-after-edge":  "dominant-baseline: text-after-edge",
	"dominant-baseline-text-before-edge": "dominant-baseline: text-before-edge",
}

func (t DominantBaselineStyle) Utilities() map[string]string {
	return DominantBaselineStyleUtilitiesMapping
}

func (t DominantBaselineStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = DominantBaselineStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderInlineStyle represent the CSS style "border-inline-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-style
type BorderInlineStyleStyle string

func (t BorderInlineStyleStyle) Name() string {
	return "border-inline-style"
}

var BorderInlineStyleStyleBrowserVariantsList = []string{}

func (t BorderInlineStyleStyle) BrowserVariants() []string {
	return BorderInlineStyleStyleBrowserVariantsList
}

var BorderInlineStyleStyleUtilitiesMapping = map[string]string{
	"border-inline-style": "border-inline-style: none",
}

func (t BorderInlineStyleStyle) Utilities() map[string]string {
	return BorderInlineStyleStyleUtilitiesMapping
}

func (t BorderInlineStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderInlineStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// VerticalAlign represent the CSS style "vertical-align" with value "baseline | sub | super | text-top | text-bottom | middle | top | bottom"
// See https://developer.mozilla.org/docs/Web/CSS/vertical-align
type VerticalAlignStyle string

func (t VerticalAlignStyle) Name() string {
	return "vertical-align"
}

const VerticalAlignStyleBaseline VerticalAlignStyle = "baseline"

const VerticalAlignStyleSub VerticalAlignStyle = "sub"

const VerticalAlignStyleSuper VerticalAlignStyle = "super"

const VerticalAlignStyleTextTop VerticalAlignStyle = "text-top"

const VerticalAlignStyleTextBottom VerticalAlignStyle = "text-bottom"

const VerticalAlignStyleMiddle VerticalAlignStyle = "middle"

const VerticalAlignStyleTop VerticalAlignStyle = "top"

const VerticalAlignStyleBottom VerticalAlignStyle = "bottom"

var VerticalAlignStyleBrowserVariantsList = []string{}

func (t VerticalAlignStyle) BrowserVariants() []string {
	return VerticalAlignStyleBrowserVariantsList
}

var VerticalAlignStyleUtilitiesMapping = map[string]string{
	"vertical-align":             "vertical-align: baseline",
	"vertical-align-baseline":    "vertical-align: baseline",
	"vertical-align-sub":         "vertical-align: sub",
	"vertical-align-super":       "vertical-align: super",
	"vertical-align-text-top":    "vertical-align: text-top",
	"vertical-align-text-bottom": "vertical-align: text-bottom",
	"vertical-align-middle":      "vertical-align: middle",
	"vertical-align-top":         "vertical-align: top",
	"vertical-align-bottom":      "vertical-align: bottom",
}

func (t VerticalAlignStyle) Utilities() map[string]string {
	return VerticalAlignStyleUtilitiesMapping
}

func (t VerticalAlignStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = VerticalAlignStyleUtilitiesMapping[tu]
	return value, hasValue
}

// EmptyCells represent the CSS style "empty-cells" with value "show | hide"
// See https://developer.mozilla.org/docs/Web/CSS/empty-cells
type EmptyCellsStyle string

func (t EmptyCellsStyle) Name() string {
	return "empty-cells"
}

const EmptyCellsStyleShow EmptyCellsStyle = "show"

const EmptyCellsStyleHide EmptyCellsStyle = "hide"

var EmptyCellsStyleBrowserVariantsList = []string{}

func (t EmptyCellsStyle) BrowserVariants() []string {
	return EmptyCellsStyleBrowserVariantsList
}

var EmptyCellsStyleUtilitiesMapping = map[string]string{
	"empty-cells":      "empty-cells: show",
	"empty-cells-show": "empty-cells: show",
	"empty-cells-hide": "empty-cells: hide",
}

func (t EmptyCellsStyle) Utilities() map[string]string {
	return EmptyCellsStyleUtilitiesMapping
}

func (t EmptyCellsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = EmptyCellsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// GridColumnEnd represent the CSS style "grid-column-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-column-end
type GridColumnEndStyle string

func (t GridColumnEndStyle) Name() string {
	return "grid-column-end"
}

var GridColumnEndStyleBrowserVariantsList = []string{}

func (t GridColumnEndStyle) BrowserVariants() []string {
	return GridColumnEndStyleBrowserVariantsList
}

var GridColumnEndStyleUtilitiesMapping = map[string]string{
	"grid-column-end": "grid-column-end: auto",
}

func (t GridColumnEndStyle) Utilities() map[string]string {
	return GridColumnEndStyleUtilitiesMapping
}

func (t GridColumnEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = GridColumnEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderWidth represent the CSS style "border-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-width
type BorderWidthStyle string

func (t BorderWidthStyle) Name() string {
	return "border-width"
}

var BorderWidthStyleBrowserVariantsList = []string{}

func (t BorderWidthStyle) BrowserVariants() []string {
	return BorderWidthStyleBrowserVariantsList
}

var BorderWidthStyleUtilitiesMapping = map[string]string{
	"border-width": "border-width: auto",
}

func (t BorderWidthStyle) Utilities() map[string]string {
	return BorderWidthStyleUtilitiesMapping
}

func (t BorderWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginRight represent the CSS style "margin-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-right
type MarginRightStyle string

func (t MarginRightStyle) Name() string {
	return "margin-right"
}

var MarginRightStyleBrowserVariantsList = []string{}

func (t MarginRightStyle) BrowserVariants() []string {
	return MarginRightStyleBrowserVariantsList
}

var MarginRightStyleUtilitiesMapping = map[string]string{
	"margin-right": "margin-right: 0",
}

func (t MarginRightStyle) Utilities() map[string]string {
	return MarginRightStyleUtilitiesMapping
}

func (t MarginRightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginRightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MinBlockSize represent the CSS style "min-block-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-block-size
type MinBlockSizeStyle string

func (t MinBlockSizeStyle) Name() string {
	return "min-block-size"
}

var MinBlockSizeStyleBrowserVariantsList = []string{}

func (t MinBlockSizeStyle) BrowserVariants() []string {
	return MinBlockSizeStyleBrowserVariantsList
}

var MinBlockSizeStyleUtilitiesMapping = map[string]string{
	"min-block-size": "min-block-size: 0",
}

func (t MinBlockSizeStyle) Utilities() map[string]string {
	return MinBlockSizeStyleUtilitiesMapping
}

func (t MinBlockSizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MinBlockSizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScaleLeft represent the CSS style "scale-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scale
type ScaleLeftStyle string

func (t ScaleLeftStyle) Name() string {
	return "scale-left"
}

var ScaleLeftStyleBrowserVariantsList = []string{}

func (t ScaleLeftStyle) BrowserVariants() []string {
	return ScaleLeftStyleBrowserVariantsList
}

var ScaleLeftStyleUtilitiesMapping = map[string]string{
	"scale-left": "scale-left: none",
}

func (t ScaleLeftStyle) Utilities() map[string]string {
	return ScaleLeftStyleUtilitiesMapping
}

func (t ScaleLeftStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScaleLeftStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontVariantLigatures represent the CSS style "font-variant-ligatures" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-ligatures
type FontVariantLigaturesStyle string

func (t FontVariantLigaturesStyle) Name() string {
	return "font-variant-ligatures"
}

var FontVariantLigaturesStyleBrowserVariantsList = []string{}

func (t FontVariantLigaturesStyle) BrowserVariants() []string {
	return FontVariantLigaturesStyleBrowserVariantsList
}

var FontVariantLigaturesStyleUtilitiesMapping = map[string]string{
	"font-variant-ligatures": "font-variant-ligatures: normal",
}

func (t FontVariantLigaturesStyle) Utilities() map[string]string {
	return FontVariantLigaturesStyleUtilitiesMapping
}

func (t FontVariantLigaturesStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontVariantLigaturesStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextDecorationThickness represent the CSS style "text-decoration-thickness" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-thickness
type TextDecorationThicknessStyle string

func (t TextDecorationThicknessStyle) Name() string {
	return "text-decoration-thickness"
}

var TextDecorationThicknessStyleBrowserVariantsList = []string{}

func (t TextDecorationThicknessStyle) BrowserVariants() []string {
	return TextDecorationThicknessStyleBrowserVariantsList
}

var TextDecorationThicknessStyleUtilitiesMapping = map[string]string{
	"text-decoration-thickness": "text-decoration-thickness: auto",
}

func (t TextDecorationThicknessStyle) Utilities() map[string]string {
	return TextDecorationThicknessStyleUtilitiesMapping
}

func (t TextDecorationThicknessStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextDecorationThicknessStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaxInlineSize represent the CSS style "max-inline-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-inline-size
type MaxInlineSizeStyle string

func (t MaxInlineSizeStyle) Name() string {
	return "max-inline-size"
}

var MaxInlineSizeStyleBrowserVariantsList = []string{}

func (t MaxInlineSizeStyle) BrowserVariants() []string {
	return MaxInlineSizeStyleBrowserVariantsList
}

var MaxInlineSizeStyleUtilitiesMapping = map[string]string{
	"max-inline-size": "max-inline-size: 0",
}

func (t MaxInlineSizeStyle) Utilities() map[string]string {
	return MaxInlineSizeStyleUtilitiesMapping
}

func (t MaxInlineSizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaxInlineSizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverflowX represent the CSS style "overflow-x" with value "visible | hidden | clip | scroll | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overflow-x
type OverflowXStyle string

func (t OverflowXStyle) Name() string {
	return "overflow-x"
}

const OverflowXStyleVisible OverflowXStyle = "visible"

const OverflowXStyleHidden OverflowXStyle = "hidden"

const OverflowXStyleClip OverflowXStyle = "clip"

const OverflowXStyleScroll OverflowXStyle = "scroll"

const OverflowXStyleAuto OverflowXStyle = "auto"

var OverflowXStyleBrowserVariantsList = []string{}

func (t OverflowXStyle) BrowserVariants() []string {
	return OverflowXStyleBrowserVariantsList
}

var OverflowXStyleUtilitiesMapping = map[string]string{
	"overflow-x":         "overflow-x: visible",
	"overflow-x-visible": "overflow-x: visible",
	"overflow-x-hidden":  "overflow-x: hidden",
	"overflow-x-clip":    "overflow-x: clip",
	"overflow-x-scroll":  "overflow-x: scroll",
	"overflow-x-auto":    "overflow-x: auto",
}

func (t OverflowXStyle) Utilities() map[string]string {
	return OverflowXStyleUtilitiesMapping
}

func (t OverflowXStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverflowXStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PaddingBottom represent the CSS style "padding-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-bottom
type PaddingBottomStyle string

func (t PaddingBottomStyle) Name() string {
	return "padding-bottom"
}

var PaddingBottomStyleBrowserVariantsList = []string{}

func (t PaddingBottomStyle) BrowserVariants() []string {
	return PaddingBottomStyleBrowserVariantsList
}

var PaddingBottomStyleUtilitiesMapping = map[string]string{
	"padding-bottom": "padding-bottom: 0",
}

func (t PaddingBottomStyle) Utilities() map[string]string {
	return PaddingBottomStyleUtilitiesMapping
}

func (t PaddingBottomStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PaddingBottomStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextUnderlinePosition represent the CSS style "text-underline-position" with value "auto | from-font | under | left | right"
// See https://developer.mozilla.org/docs/Web/CSS/text-underline-position
type TextUnderlinePositionStyle string

func (t TextUnderlinePositionStyle) Name() string {
	return "text-underline-position"
}

const TextUnderlinePositionStyleAuto TextUnderlinePositionStyle = "auto"

const TextUnderlinePositionStyleFromFont TextUnderlinePositionStyle = "from-font"

const TextUnderlinePositionStyleUnder TextUnderlinePositionStyle = "under"

const TextUnderlinePositionStyleLeft TextUnderlinePositionStyle = "left"

const TextUnderlinePositionStyleRight TextUnderlinePositionStyle = "right"

var TextUnderlinePositionStyleBrowserVariantsList = []string{}

func (t TextUnderlinePositionStyle) BrowserVariants() []string {
	return TextUnderlinePositionStyleBrowserVariantsList
}

var TextUnderlinePositionStyleUtilitiesMapping = map[string]string{
	"text-underline-position":           "text-underline-position: auto",
	"text-underline-position-auto":      "text-underline-position: auto",
	"text-underline-position-from-font": "text-underline-position: from-font",
	"text-underline-position-under":     "text-underline-position: under",
	"text-underline-position-left":      "text-underline-position: left",
	"text-underline-position-right":     "text-underline-position: right",
}

func (t TextUnderlinePositionStyle) Utilities() map[string]string {
	return TextUnderlinePositionStyleUtilitiesMapping
}

func (t TextUnderlinePositionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextUnderlinePositionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderRightStyle represent the CSS style "border-right-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-right-style
type BorderRightStyleStyle string

func (t BorderRightStyleStyle) Name() string {
	return "border-right-style"
}

var BorderRightStyleStyleBrowserVariantsList = []string{}

func (t BorderRightStyleStyle) BrowserVariants() []string {
	return BorderRightStyleStyleBrowserVariantsList
}

var BorderRightStyleStyleUtilitiesMapping = map[string]string{
	"border-right-style": "border-right-style: none",
}

func (t BorderRightStyleStyle) Utilities() map[string]string {
	return BorderRightStyleStyleUtilitiesMapping
}

func (t BorderRightStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderRightStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginInline represent the CSS style "margin-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-inline
type MarginInlineStyle string

func (t MarginInlineStyle) Name() string {
	return "margin-inline"
}

var MarginInlineStyleBrowserVariantsList = []string{}

func (t MarginInlineStyle) BrowserVariants() []string {
	return MarginInlineStyleBrowserVariantsList
}

var MarginInlineStyleUtilitiesMapping = map[string]string{
	"margin-inline": "margin-inline: 0",
}

func (t MarginInlineStyle) Utilities() map[string]string {
	return MarginInlineStyleUtilitiesMapping
}

func (t MarginInlineStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginInlineStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformRotateY represent the CSS style "transform-rotate-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformRotateYStyle string

func (t TransformRotateYStyle) Name() string {
	return "transform-rotate-y"
}

var TransformRotateYStyleBrowserVariantsList = []string{}

func (t TransformRotateYStyle) BrowserVariants() []string {
	return TransformRotateYStyleBrowserVariantsList
}

var TransformRotateYStyleUtilitiesMapping = map[string]string{
	"transform-rotate-y": "transform-rotate-y: rotateY(0)",
}

func (t TransformRotateYStyle) Utilities() map[string]string {
	return TransformRotateYStyleUtilitiesMapping
}

func (t TransformRotateYStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformRotateYStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AlignContent represent the CSS style "align-content" with value "auto | normal | stretch | end | start | flex-start | flex-end | center |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe | unsafe"
// See https://developer.mozilla.org/docs/Web/CSS/align-content
type AlignContentStyle string

func (t AlignContentStyle) Name() string {
	return "align-content"
}

const AlignContentStyleAuto AlignContentStyle = "auto"

const AlignContentStyleNormal AlignContentStyle = "normal"

const AlignContentStyleStretch AlignContentStyle = "stretch"

const AlignContentStyleEnd AlignContentStyle = "end"

const AlignContentStyleStart AlignContentStyle = "start"

const AlignContentStyleFlexStart AlignContentStyle = "flex-start"

const AlignContentStyleFlexEnd AlignContentStyle = "flex-end"

const AlignContentStyleCenter AlignContentStyle = "center"

const AlignContentStyleBaseline AlignContentStyle = "baseline"

const AlignContentStyleFirstBaseline AlignContentStyle = "first-baseline"

const AlignContentStyleLastBaseline AlignContentStyle = "last-baseline"

const AlignContentStyleSpaceBetween AlignContentStyle = "space-between"

const AlignContentStyleSpaceAround AlignContentStyle = "space-around"

const AlignContentStyleSpaceEvenly AlignContentStyle = "space-evenly"

const AlignContentStyleSafe AlignContentStyle = "safe"

const AlignContentStyleUnsafe AlignContentStyle = "unsafe"

var AlignContentStyleBrowserVariantsList = []string{}

func (t AlignContentStyle) BrowserVariants() []string {
	return AlignContentStyleBrowserVariantsList
}

var AlignContentStyleUtilitiesMapping = map[string]string{
	"align-content":                "align-content: normal",
	"align-content-auto":           "align-content: auto",
	"align-content-normal":         "align-content: normal",
	"align-content-stretch":        "align-content: stretch",
	"align-content-end":            "align-content: end",
	"align-content-start":          "align-content: start",
	"align-content-flex-start":     "align-content: flex-start",
	"align-content-flex-end":       "align-content: flex-end",
	"align-content-center":         "align-content: center",
	"align-content-baseline":       "align-content: baseline",
	"align-content-first-baseline": "align-content: first-baseline",
	"align-content-last-baseline":  "align-content: last-baseline",
	"align-content-space-between":  "align-content: space-between",
	"align-content-space-around":   "align-content: space-around",
	"align-content-space-evenly":   "align-content: space-evenly",
	"align-content-safe":           "align-content: safe",
	"align-content-unsafe":         "align-content: unsafe",
}

func (t AlignContentStyle) Utilities() map[string]string {
	return AlignContentStyleUtilitiesMapping
}

func (t AlignContentStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AlignContentStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginBlock represent the CSS style "margin-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-block
type MarginBlockStyle string

func (t MarginBlockStyle) Name() string {
	return "margin-block"
}

var MarginBlockStyleBrowserVariantsList = []string{}

func (t MarginBlockStyle) BrowserVariants() []string {
	return MarginBlockStyleBrowserVariantsList
}

var MarginBlockStyleUtilitiesMapping = map[string]string{
	"margin-block": "margin-block: 0",
}

func (t MarginBlockStyle) Utilities() map[string]string {
	return MarginBlockStyleUtilitiesMapping
}

func (t MarginBlockStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginBlockStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OffsetRotate represent the CSS style "offset-rotate" with value "auto | reverse | angle"
// See https://developer.mozilla.org/docs/Web/CSS/offset-rotate
type OffsetRotateStyle string

func (t OffsetRotateStyle) Name() string {
	return "offset-rotate"
}

const OffsetRotateStyleAuto OffsetRotateStyle = "auto"

const OffsetRotateStyleReverse OffsetRotateStyle = "reverse"

const OffsetRotateStyleAngle OffsetRotateStyle = "angle"

var OffsetRotateStyleBrowserVariantsList = []string{}

func (t OffsetRotateStyle) BrowserVariants() []string {
	return OffsetRotateStyleBrowserVariantsList
}

var OffsetRotateStyleUtilitiesMapping = map[string]string{
	"offset-rotate":         "offset-rotate: auto",
	"offset-rotate-auto":    "offset-rotate: auto",
	"offset-rotate-reverse": "offset-rotate: reverse",
	"offset-rotate-angle":   "offset-rotate: angle",
}

func (t OffsetRotateStyle) Utilities() map[string]string {
	return OffsetRotateStyleUtilitiesMapping
}

func (t OffsetRotateStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OffsetRotateStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformStyle represent the CSS style "transform-style" with value "flat | preserve-3d"
// See https://developer.mozilla.org/docs/Web/CSS/transform-style
type TransformStyleStyle string

func (t TransformStyleStyle) Name() string {
	return "transform-style"
}

const TransformStyleStyleFlat TransformStyleStyle = "flat"

const TransformStyleStylePreserve3d TransformStyleStyle = "preserve-3d"

var TransformStyleStyleBrowserVariantsList = []string{}

func (t TransformStyleStyle) BrowserVariants() []string {
	return TransformStyleStyleBrowserVariantsList
}

var TransformStyleStyleUtilitiesMapping = map[string]string{
	"transform-style":             "transform-style: flat",
	"transform-style-flat":        "transform-style: flat",
	"transform-style-preserve-3d": "transform-style: preserve-3d",
}

func (t TransformStyleStyle) Utilities() map[string]string {
	return TransformStyleStyleUtilitiesMapping
}

func (t TransformStyleStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformStyleStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OffsetPosition represent the CSS style "offset-position" with value ""
// See
type OffsetPositionStyle string

func (t OffsetPositionStyle) Name() string {
	return "offset-position"
}

var OffsetPositionStyleBrowserVariantsList = []string{}

func (t OffsetPositionStyle) BrowserVariants() []string {
	return OffsetPositionStyleBrowserVariantsList
}

var OffsetPositionStyleUtilitiesMapping = map[string]string{
	"offset-position": "offset-position: auto",
}

func (t OffsetPositionStyle) Utilities() map[string]string {
	return OffsetPositionStyleUtilitiesMapping
}

func (t OffsetPositionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OffsetPositionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskBorderSlice represent the CSS style "mask-border-slice" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-slice
type MaskBorderSliceStyle string

func (t MaskBorderSliceStyle) Name() string {
	return "mask-border-slice"
}

var MaskBorderSliceStyleBrowserVariantsList = []string{}

func (t MaskBorderSliceStyle) BrowserVariants() []string {
	return MaskBorderSliceStyleBrowserVariantsList
}

var MaskBorderSliceStyleUtilitiesMapping = map[string]string{
	"mask-border-slice": "mask-border-slice: 0",
}

func (t MaskBorderSliceStyle) Utilities() map[string]string {
	return MaskBorderSliceStyleUtilitiesMapping
}

func (t MaskBorderSliceStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskBorderSliceStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OutlineColor represent the CSS style "outline-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-color
type OutlineColorStyle Color

func (t OutlineColorStyle) Name() string {
	return "outline-color"
}

var OutlineColorStyleBrowserVariantsList = []string{}

func (t OutlineColorStyle) BrowserVariants() []string {
	return OutlineColorStyleBrowserVariantsList
}

var OutlineColorStyleUtilitiesMapping = map[string]string{
	"outline-color": "outline-color: invertOrCurrentColor",
}

func (t OutlineColorStyle) Utilities() map[string]string {
	return OutlineColorStyleUtilitiesMapping
}

func (t OutlineColorStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OutlineColorStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Filter represent the CSS style "filter" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/filter
type FilterStyle string

func (t FilterStyle) Name() string {
	return "filter"
}

var FilterStyleBrowserVariantsList = []string{
	"-ms-filter",
}

func (t FilterStyle) BrowserVariants() []string {
	return FilterStyleBrowserVariantsList
}

var FilterStyleUtilitiesMapping = map[string]string{
	"filter": "filter: none",
}

func (t FilterStyle) Utilities() map[string]string {
	return FilterStyleUtilitiesMapping
}

func (t FilterStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FilterStyleUtilitiesMapping[tu]
	return value, hasValue
}

// OverflowWrap represent the CSS style "overflow-wrap" with value "normal | break-word | anywhere"
// See https://developer.mozilla.org/docs/Web/CSS/overflow-wrap
type OverflowWrapStyle string

func (t OverflowWrapStyle) Name() string {
	return "overflow-wrap"
}

const OverflowWrapStyleNormal OverflowWrapStyle = "normal"

const OverflowWrapStyleBreakWord OverflowWrapStyle = "break-word"

const OverflowWrapStyleAnywhere OverflowWrapStyle = "anywhere"

var OverflowWrapStyleBrowserVariantsList = []string{}

func (t OverflowWrapStyle) BrowserVariants() []string {
	return OverflowWrapStyleBrowserVariantsList
}

var OverflowWrapStyleUtilitiesMapping = map[string]string{
	"overflow-wrap":            "overflow-wrap: normal",
	"overflow-wrap-normal":     "overflow-wrap: normal",
	"overflow-wrap-break-word": "overflow-wrap: break-word",
	"overflow-wrap-anywhere":   "overflow-wrap: anywhere",
}

func (t OverflowWrapStyle) Utilities() map[string]string {
	return OverflowWrapStyleUtilitiesMapping
}

func (t OverflowWrapStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = OverflowWrapStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMarginRight represent the CSS style "scroll-margin-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-right
type ScrollMarginRightStyle float64

func (t ScrollMarginRightStyle) Name() string {
	return "scroll-margin-right"
}

var ScrollMarginRightStyleBrowserVariantsList = []string{}

func (t ScrollMarginRightStyle) BrowserVariants() []string {
	return ScrollMarginRightStyleBrowserVariantsList
}

var ScrollMarginRightStyleUtilitiesMapping = map[string]string{
	"scroll-margin-right": "scroll-margin-right: 0",
}

func (t ScrollMarginRightStyle) Utilities() map[string]string {
	return ScrollMarginRightStyleUtilitiesMapping
}

func (t ScrollMarginRightStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginRightStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformSkewY represent the CSS style "transform-skew-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformSkewYStyle string

func (t TransformSkewYStyle) Name() string {
	return "transform-skew-y"
}

var TransformSkewYStyleBrowserVariantsList = []string{}

func (t TransformSkewYStyle) BrowserVariants() []string {
	return TransformSkewYStyleBrowserVariantsList
}

var TransformSkewYStyleUtilitiesMapping = map[string]string{
	"transform-skew-y": "transform-skew-y: skewY(0)",
}

func (t TransformSkewYStyle) Utilities() map[string]string {
	return TransformSkewYStyleUtilitiesMapping
}

func (t TransformSkewYStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformSkewYStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxShadow represent the CSS style "box-shadow" with value "string"
// See https://developer.mozilla.org/docs/Web/CSS/box-shadow
type BoxShadowStyle string

func (t BoxShadowStyle) Name() string {
	return "box-shadow"
}

const BoxShadowStyleString BoxShadowStyle = "string"

var BoxShadowStyleBrowserVariantsList = []string{}

func (t BoxShadowStyle) BrowserVariants() []string {
	return BoxShadowStyleBrowserVariantsList
}

var BoxShadowStyleUtilitiesMapping = map[string]string{
	"box-shadow":        "box-shadow: none",
	"box-shadow-string": "box-shadow: string",
}

func (t BoxShadowStyle) Utilities() map[string]string {
	return BoxShadowStyleUtilitiesMapping
}

func (t BoxShadowStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxShadowStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BoxSizing represent the CSS style "box-sizing" with value "content-box | border-box"
// See https://developer.mozilla.org/docs/Web/CSS/box-sizing
type BoxSizingStyle string

func (t BoxSizingStyle) Name() string {
	return "box-sizing"
}

const BoxSizingStyleContentBox BoxSizingStyle = "content-box"

const BoxSizingStyleBorderBox BoxSizingStyle = "border-box"

var BoxSizingStyleBrowserVariantsList = []string{}

func (t BoxSizingStyle) BrowserVariants() []string {
	return BoxSizingStyleBrowserVariantsList
}

var BoxSizingStyleUtilitiesMapping = map[string]string{
	"box-sizing":             "box-sizing: content-box",
	"box-sizing-content-box": "box-sizing: content-box",
	"box-sizing-border-box":  "box-sizing: border-box",
}

func (t BoxSizingStyle) Utilities() map[string]string {
	return BoxSizingStyleUtilitiesMapping
}

func (t BoxSizingStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BoxSizingStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransformScaleZ represent the CSS style "transform-scale-z" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformScaleZStyle string

func (t TransformScaleZStyle) Name() string {
	return "transform-scale-z"
}

var TransformScaleZStyleBrowserVariantsList = []string{}

func (t TransformScaleZStyle) BrowserVariants() []string {
	return TransformScaleZStyleBrowserVariantsList
}

var TransformScaleZStyleUtilitiesMapping = map[string]string{
	"transform-scale-z": "transform-scale-z: scaleZ(0)",
}

func (t TransformScaleZStyle) Utilities() map[string]string {
	return TransformScaleZStyleUtilitiesMapping
}

func (t TransformScaleZStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransformScaleZStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BorderRightWidth represent the CSS style "border-right-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-right-width
type BorderRightWidthStyle string

func (t BorderRightWidthStyle) Name() string {
	return "border-right-width"
}

var BorderRightWidthStyleBrowserVariantsList = []string{}

func (t BorderRightWidthStyle) BrowserVariants() []string {
	return BorderRightWidthStyleBrowserVariantsList
}

var BorderRightWidthStyleUtilitiesMapping = map[string]string{
	"border-right-width": "border-right-width: medium",
}

func (t BorderRightWidthStyle) Utilities() map[string]string {
	return BorderRightWidthStyleUtilitiesMapping
}

func (t BorderRightWidthStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BorderRightWidthStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FlexWrap represent the CSS style "flex-wrap" with value "nowrap | wrap | wrap-reverse"
// See https://developer.mozilla.org/docs/Web/CSS/flex-wrap
type FlexWrapStyle string

func (t FlexWrapStyle) Name() string {
	return "flex-wrap"
}

const FlexWrapStyleNowrap FlexWrapStyle = "nowrap"

const FlexWrapStyleWrap FlexWrapStyle = "wrap"

const FlexWrapStyleWrapReverse FlexWrapStyle = "wrap-reverse"

var FlexWrapStyleBrowserVariantsList = []string{}

func (t FlexWrapStyle) BrowserVariants() []string {
	return FlexWrapStyleBrowserVariantsList
}

var FlexWrapStyleUtilitiesMapping = map[string]string{
	"flex-wrap":              "flex-wrap: nowrap",
	"flex-wrap-nowrap":       "flex-wrap: nowrap",
	"flex-wrap-wrap":         "flex-wrap: wrap",
	"flex-wrap-wrap-reverse": "flex-wrap: wrap-reverse",
}

func (t FlexWrapStyle) Utilities() map[string]string {
	return FlexWrapStyleUtilitiesMapping
}

func (t FlexWrapStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FlexWrapStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMargin represent the CSS style "scroll-margin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin
type ScrollMarginStyle float64

func (t ScrollMarginStyle) Name() string {
	return "scroll-margin"
}

var ScrollMarginStyleBrowserVariantsList = []string{}

func (t ScrollMarginStyle) BrowserVariants() []string {
	return ScrollMarginStyleBrowserVariantsList
}

var ScrollMarginStyleUtilitiesMapping = map[string]string{
	"scroll-margin": "scroll-margin: 0",
}

func (t ScrollMarginStyle) Utilities() map[string]string {
	return ScrollMarginStyleUtilitiesMapping
}

func (t ScrollMarginStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollMarginBlockEnd represent the CSS style "scroll-margin-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-block-end
type ScrollMarginBlockEndStyle float64

func (t ScrollMarginBlockEndStyle) Name() string {
	return "scroll-margin-block-end"
}

var ScrollMarginBlockEndStyleBrowserVariantsList = []string{}

func (t ScrollMarginBlockEndStyle) BrowserVariants() []string {
	return ScrollMarginBlockEndStyleBrowserVariantsList
}

var ScrollMarginBlockEndStyleUtilitiesMapping = map[string]string{
	"scroll-margin-block-end": "scroll-margin-block-end: 0",
}

func (t ScrollMarginBlockEndStyle) Utilities() map[string]string {
	return ScrollMarginBlockEndStyleUtilitiesMapping
}

func (t ScrollMarginBlockEndStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollMarginBlockEndStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AlignmentBaseline represent the CSS style "alignment-baseline" with value "auto | baseline | before-edge | text-before-edge | middle | central | after-edge | text-after-edge | ideographic | alphabetic | hanging | mathematical"
// See
type AlignmentBaselineStyle string

func (t AlignmentBaselineStyle) Name() string {
	return "alignment-baseline"
}

const AlignmentBaselineStyleAuto AlignmentBaselineStyle = "auto"

const AlignmentBaselineStyleBaseline AlignmentBaselineStyle = "baseline"

const AlignmentBaselineStyleBeforeEdge AlignmentBaselineStyle = "before-edge"

const AlignmentBaselineStyleTextBeforeEdge AlignmentBaselineStyle = "text-before-edge"

const AlignmentBaselineStyleMiddle AlignmentBaselineStyle = "middle"

const AlignmentBaselineStyleCentral AlignmentBaselineStyle = "central"

const AlignmentBaselineStyleAfterEdge AlignmentBaselineStyle = "after-edge"

const AlignmentBaselineStyleTextAfterEdge AlignmentBaselineStyle = "text-after-edge"

const AlignmentBaselineStyleIdeographic AlignmentBaselineStyle = "ideographic"

const AlignmentBaselineStyleAlphabetic AlignmentBaselineStyle = "alphabetic"

const AlignmentBaselineStyleHanging AlignmentBaselineStyle = "hanging"

const AlignmentBaselineStyleMathematical AlignmentBaselineStyle = "mathematical"

var AlignmentBaselineStyleBrowserVariantsList = []string{}

func (t AlignmentBaselineStyle) BrowserVariants() []string {
	return AlignmentBaselineStyleBrowserVariantsList
}

var AlignmentBaselineStyleUtilitiesMapping = map[string]string{
	"alignment-baseline":                  "alignment-baseline: auto",
	"alignment-baseline-auto":             "alignment-baseline: auto",
	"alignment-baseline-baseline":         "alignment-baseline: baseline",
	"alignment-baseline-before-edge":      "alignment-baseline: before-edge",
	"alignment-baseline-text-before-edge": "alignment-baseline: text-before-edge",
	"alignment-baseline-middle":           "alignment-baseline: middle",
	"alignment-baseline-central":          "alignment-baseline: central",
	"alignment-baseline-after-edge":       "alignment-baseline: after-edge",
	"alignment-baseline-text-after-edge":  "alignment-baseline: text-after-edge",
	"alignment-baseline-ideographic":      "alignment-baseline: ideographic",
	"alignment-baseline-alphabetic":       "alignment-baseline: alphabetic",
	"alignment-baseline-hanging":          "alignment-baseline: hanging",
	"alignment-baseline-mathematical":     "alignment-baseline: mathematical",
}

func (t AlignmentBaselineStyle) Utilities() map[string]string {
	return AlignmentBaselineStyleUtilitiesMapping
}

func (t AlignmentBaselineStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AlignmentBaselineStyleUtilitiesMapping[tu]
	return value, hasValue
}

// StrokeMiterlimit represent the CSS style "stroke-miterlimit" with value ""
// See
type StrokeMiterlimitStyle float64

func (t StrokeMiterlimitStyle) Name() string {
	return "stroke-miterlimit"
}

var StrokeMiterlimitStyleBrowserVariantsList = []string{}

func (t StrokeMiterlimitStyle) BrowserVariants() []string {
	return StrokeMiterlimitStyleBrowserVariantsList
}

var StrokeMiterlimitStyleUtilitiesMapping = map[string]string{
	"stroke-miterlimit": "stroke-miterlimit: 4",
}

func (t StrokeMiterlimitStyle) Utilities() map[string]string {
	return StrokeMiterlimitStyleUtilitiesMapping
}

func (t StrokeMiterlimitStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = StrokeMiterlimitStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Perspective represent the CSS style "perspective" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/perspective
type PerspectiveStyle string

func (t PerspectiveStyle) Name() string {
	return "perspective"
}

var PerspectiveStyleBrowserVariantsList = []string{}

func (t PerspectiveStyle) BrowserVariants() []string {
	return PerspectiveStyleBrowserVariantsList
}

var PerspectiveStyleUtilitiesMapping = map[string]string{
	"perspective": "perspective: none",
}

func (t PerspectiveStyle) Utilities() map[string]string {
	return PerspectiveStyleUtilitiesMapping
}

func (t PerspectiveStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PerspectiveStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MasonryAutoFlow represent the CSS style "masonry-auto-flow" with value "pack | next | pack definite-first | next definite-first | pack ordered | next ordered"
// See https://developer.mozilla.org/docs/Web/CSS/masonry-auto-flow
type MasonryAutoFlowStyle string

func (t MasonryAutoFlowStyle) Name() string {
	return "masonry-auto-flow"
}

const MasonryAutoFlowStylePack MasonryAutoFlowStyle = "pack"

const MasonryAutoFlowStyleNext MasonryAutoFlowStyle = "next"

const MasonryAutoFlowStylePackDefiniteFirst MasonryAutoFlowStyle = "pack-definite-first"

const MasonryAutoFlowStyleNextDefiniteFirst MasonryAutoFlowStyle = "next-definite-first"

const MasonryAutoFlowStylePackOrdered MasonryAutoFlowStyle = "pack-ordered"

const MasonryAutoFlowStyleNextOrdered MasonryAutoFlowStyle = "next-ordered"

var MasonryAutoFlowStyleBrowserVariantsList = []string{}

func (t MasonryAutoFlowStyle) BrowserVariants() []string {
	return MasonryAutoFlowStyleBrowserVariantsList
}

var MasonryAutoFlowStyleUtilitiesMapping = map[string]string{
	"masonry-auto-flow":                     "masonry-auto-flow: pack",
	"masonry-auto-flow-pack":                "masonry-auto-flow: pack",
	"masonry-auto-flow-next":                "masonry-auto-flow: next",
	"masonry-auto-flow-pack-definite-first": "masonry-auto-flow: pack-definite-first",
	"masonry-auto-flow-next-definite-first": "masonry-auto-flow: next-definite-first",
	"masonry-auto-flow-pack-ordered":        "masonry-auto-flow: pack-ordered",
	"masonry-auto-flow-next-ordered":        "masonry-auto-flow: next-ordered",
}

func (t MasonryAutoFlowStyle) Utilities() map[string]string {
	return MasonryAutoFlowStyleUtilitiesMapping
}

func (t MasonryAutoFlowStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MasonryAutoFlowStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ScrollSnapPointsX represent the CSS style "scroll-snap-points-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-points-x
type ScrollSnapPointsXStyle string

func (t ScrollSnapPointsXStyle) Name() string {
	return "scroll-snap-points-x"
}

var ScrollSnapPointsXStyleBrowserVariantsList = []string{
	"-ms-scroll-snap-points-x",
}

func (t ScrollSnapPointsXStyle) BrowserVariants() []string {
	return ScrollSnapPointsXStyleBrowserVariantsList
}

var ScrollSnapPointsXStyleUtilitiesMapping = map[string]string{
	"scroll-snap-points-x": "scroll-snap-points-x: none",
}

func (t ScrollSnapPointsXStyle) Utilities() map[string]string {
	return ScrollSnapPointsXStyleUtilitiesMapping
}

func (t ScrollSnapPointsXStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ScrollSnapPointsXStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Widows represent the CSS style "widows" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/widows
type WidowsStyle float64

func (t WidowsStyle) Name() string {
	return "widows"
}

var WidowsStyleBrowserVariantsList = []string{}

func (t WidowsStyle) BrowserVariants() []string {
	return WidowsStyleBrowserVariantsList
}

var WidowsStyleUtilitiesMapping = map[string]string{
	"widows": "widows: 2",
}

func (t WidowsStyle) Utilities() map[string]string {
	return WidowsStyleUtilitiesMapping
}

func (t WidowsStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = WidowsStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BaselineShift represent the CSS style "baseline-shift" with value "baseline | sub | super"
// See
type BaselineShiftStyle string

func (t BaselineShiftStyle) Name() string {
	return "baseline-shift"
}

const BaselineShiftStyleBaseline BaselineShiftStyle = "baseline"

const BaselineShiftStyleSub BaselineShiftStyle = "sub"

const BaselineShiftStyleSuper BaselineShiftStyle = "super"

var BaselineShiftStyleBrowserVariantsList = []string{}

func (t BaselineShiftStyle) BrowserVariants() []string {
	return BaselineShiftStyleBrowserVariantsList
}

var BaselineShiftStyleUtilitiesMapping = map[string]string{
	"baseline-shift":          "baseline-shift: baseline",
	"baseline-shift-baseline": "baseline-shift: baseline",
	"baseline-shift-sub":      "baseline-shift: sub",
	"baseline-shift-super":    "baseline-shift: super",
}

func (t BaselineShiftStyle) Utilities() map[string]string {
	return BaselineShiftStyleUtilitiesMapping
}

func (t BaselineShiftStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BaselineShiftStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MarginTrim represent the CSS style "margin-trim" with value "none | in-flow | all"
// See https://developer.mozilla.org/docs/Web/CSS/margin-trim
type MarginTrimStyle string

func (t MarginTrimStyle) Name() string {
	return "margin-trim"
}

const MarginTrimStyleNone MarginTrimStyle = "none"

const MarginTrimStyleInFlow MarginTrimStyle = "in-flow"

const MarginTrimStyleAll MarginTrimStyle = "all"

var MarginTrimStyleBrowserVariantsList = []string{}

func (t MarginTrimStyle) BrowserVariants() []string {
	return MarginTrimStyleBrowserVariantsList
}

var MarginTrimStyleUtilitiesMapping = map[string]string{
	"margin-trim":         "margin-trim: none",
	"margin-trim-none":    "margin-trim: none",
	"margin-trim-in-flow": "margin-trim: in-flow",
	"margin-trim-all":     "margin-trim: all",
}

func (t MarginTrimStyle) Utilities() map[string]string {
	return MarginTrimStyleUtilitiesMapping
}

func (t MarginTrimStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MarginTrimStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TransitionTimingFunction represent the CSS style "transition-timing-function" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-timing-function
type TransitionTimingFunctionStyle string

func (t TransitionTimingFunctionStyle) Name() string {
	return "transition-timing-function"
}

var TransitionTimingFunctionStyleBrowserVariantsList = []string{}

func (t TransitionTimingFunctionStyle) BrowserVariants() []string {
	return TransitionTimingFunctionStyleBrowserVariantsList
}

var TransitionTimingFunctionStyleUtilitiesMapping = map[string]string{
	"transition-timing-function": "transition-timing-function: ease",
}

func (t TransitionTimingFunctionStyle) Utilities() map[string]string {
	return TransitionTimingFunctionStyleUtilitiesMapping
}

func (t TransitionTimingFunctionStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TransitionTimingFunctionStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FlexShrink represent the CSS style "flex-shrink" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/flex-shrink
type FlexShrinkStyle float64

func (t FlexShrinkStyle) Name() string {
	return "flex-shrink"
}

var FlexShrinkStyleBrowserVariantsList = []string{}

func (t FlexShrinkStyle) BrowserVariants() []string {
	return FlexShrinkStyleBrowserVariantsList
}

var FlexShrinkStyleUtilitiesMapping = map[string]string{
	"flex-shrink": "flex-shrink: 1",
}

func (t FlexShrinkStyle) Utilities() map[string]string {
	return FlexShrinkStyleUtilitiesMapping
}

func (t FlexShrinkStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FlexShrinkStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Hyphens represent the CSS style "hyphens" with value "none | manual | auto"
// See https://developer.mozilla.org/docs/Web/CSS/hyphens
type HyphensStyle string

func (t HyphensStyle) Name() string {
	return "hyphens"
}

const HyphensStyleNone HyphensStyle = "none"

const HyphensStyleManual HyphensStyle = "manual"

const HyphensStyleAuto HyphensStyle = "auto"

var HyphensStyleBrowserVariantsList = []string{}

func (t HyphensStyle) BrowserVariants() []string {
	return HyphensStyleBrowserVariantsList
}

var HyphensStyleUtilitiesMapping = map[string]string{
	"hyphens":        "hyphens: manual",
	"hyphens-none":   "hyphens: none",
	"hyphens-manual": "hyphens: manual",
	"hyphens-auto":   "hyphens: auto",
}

func (t HyphensStyle) Utilities() map[string]string {
	return HyphensStyleUtilitiesMapping
}

func (t HyphensStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = HyphensStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PlaceContent represent the CSS style "place-content" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/place-content
type PlaceContentStyle string

func (t PlaceContentStyle) Name() string {
	return "place-content"
}

var PlaceContentStyleBrowserVariantsList = []string{}

func (t PlaceContentStyle) BrowserVariants() []string {
	return PlaceContentStyleBrowserVariantsList
}

var PlaceContentStyleUtilitiesMapping = map[string]string{
	"place-content": "place-content: normal",
}

func (t PlaceContentStyle) Utilities() map[string]string {
	return PlaceContentStyleUtilitiesMapping
}

func (t PlaceContentStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PlaceContentStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Visibility represent the CSS style "visibility" with value "visible | hidden | collapse"
// See https://developer.mozilla.org/docs/Web/CSS/visibility
type VisibilityStyle string

func (t VisibilityStyle) Name() string {
	return "visibility"
}

const VisibilityStyleVisible VisibilityStyle = "visible"

const VisibilityStyleHidden VisibilityStyle = "hidden"

const VisibilityStyleCollapse VisibilityStyle = "collapse"

var VisibilityStyleBrowserVariantsList = []string{}

func (t VisibilityStyle) BrowserVariants() []string {
	return VisibilityStyleBrowserVariantsList
}

var VisibilityStyleUtilitiesMapping = map[string]string{
	"visibility":          "visibility: visible",
	"visibility-visible":  "visibility: visible",
	"visibility-hidden":   "visibility: hidden",
	"visibility-collapse": "visibility: collapse",
}

func (t VisibilityStyle) Utilities() map[string]string {
	return VisibilityStyleUtilitiesMapping
}

func (t VisibilityStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = VisibilityStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ObjectFit represent the CSS style "object-fit" with value "fill | contain | cover | none | scale-down"
// See https://developer.mozilla.org/docs/Web/CSS/object-fit
type ObjectFitStyle string

func (t ObjectFitStyle) Name() string {
	return "object-fit"
}

const ObjectFitStyleFill ObjectFitStyle = "fill"

const ObjectFitStyleContain ObjectFitStyle = "contain"

const ObjectFitStyleCover ObjectFitStyle = "cover"

const ObjectFitStyleNone ObjectFitStyle = "none"

const ObjectFitStyleScaleDown ObjectFitStyle = "scale-down"

var ObjectFitStyleBrowserVariantsList = []string{}

func (t ObjectFitStyle) BrowserVariants() []string {
	return ObjectFitStyleBrowserVariantsList
}

var ObjectFitStyleUtilitiesMapping = map[string]string{
	"object-fit":            "object-fit: fill",
	"object-fit-fill":       "object-fit: fill",
	"object-fit-contain":    "object-fit: contain",
	"object-fit-cover":      "object-fit: cover",
	"object-fit-none":       "object-fit: none",
	"object-fit-scale-down": "object-fit: scale-down",
}

func (t ObjectFitStyle) Utilities() map[string]string {
	return ObjectFitStyleUtilitiesMapping
}

func (t ObjectFitStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ObjectFitStyleUtilitiesMapping[tu]
	return value, hasValue
}

// WordBreak represent the CSS style "word-break" with value "normal | break-all | keep-all | break-word"
// See https://developer.mozilla.org/docs/Web/CSS/word-break
type WordBreakStyle string

func (t WordBreakStyle) Name() string {
	return "word-break"
}

const WordBreakStyleNormal WordBreakStyle = "normal"

const WordBreakStyleBreakAll WordBreakStyle = "break-all"

const WordBreakStyleKeepAll WordBreakStyle = "keep-all"

const WordBreakStyleBreakWord WordBreakStyle = "break-word"

var WordBreakStyleBrowserVariantsList = []string{}

func (t WordBreakStyle) BrowserVariants() []string {
	return WordBreakStyleBrowserVariantsList
}

var WordBreakStyleUtilitiesMapping = map[string]string{
	"word-break":            "word-break: normal",
	"word-break-normal":     "word-break: normal",
	"word-break-break-all":  "word-break: break-all",
	"word-break-keep-all":   "word-break: keep-all",
	"word-break-break-word": "word-break: break-word",
}

func (t WordBreakStyle) Utilities() map[string]string {
	return WordBreakStyleUtilitiesMapping
}

func (t WordBreakStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = WordBreakStyleUtilitiesMapping[tu]
	return value, hasValue
}

// Appearance represent the CSS style "appearance" with value "none | auto | textfield | menulist-button | compat-auto"
// See https://developer.mozilla.org/docs/Web/CSS/appearance
type AppearanceStyle string

func (t AppearanceStyle) Name() string {
	return "appearance"
}

const AppearanceStyleNone AppearanceStyle = "none"

const AppearanceStyleAuto AppearanceStyle = "auto"

const AppearanceStyleTextfield AppearanceStyle = "textfield"

const AppearanceStyleMenulistButton AppearanceStyle = "menulist-button"

const AppearanceStyleCompatAuto AppearanceStyle = "compat-auto"

var AppearanceStyleBrowserVariantsList = []string{
	"-webkit-appearance",
	"-moz-appearance",
}

func (t AppearanceStyle) BrowserVariants() []string {
	return AppearanceStyleBrowserVariantsList
}

var AppearanceStyleUtilitiesMapping = map[string]string{
	"appearance":                 "appearance: auto",
	"appearance-none":            "appearance: none",
	"appearance-auto":            "appearance: auto",
	"appearance-textfield":       "appearance: textfield",
	"appearance-menulist-button": "appearance: menulist-button",
	"appearance-compat-auto":     "appearance: compat-auto",
}

func (t AppearanceStyle) Utilities() map[string]string {
	return AppearanceStyleUtilitiesMapping
}

func (t AppearanceStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AppearanceStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontVariantNumeric represent the CSS style "font-variant-numeric" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-numeric
type FontVariantNumericStyle string

func (t FontVariantNumericStyle) Name() string {
	return "font-variant-numeric"
}

var FontVariantNumericStyleBrowserVariantsList = []string{}

func (t FontVariantNumericStyle) BrowserVariants() []string {
	return FontVariantNumericStyleBrowserVariantsList
}

var FontVariantNumericStyleUtilitiesMapping = map[string]string{
	"font-variant-numeric": "font-variant-numeric: normal",
}

func (t FontVariantNumericStyle) Utilities() map[string]string {
	return FontVariantNumericStyleUtilitiesMapping
}

func (t FontVariantNumericStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontVariantNumericStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ImeMode represent the CSS style "ime-mode" with value "auto | normal | active | inactive | disabled"
// See https://developer.mozilla.org/docs/Web/CSS/ime-mode
type ImeModeStyle string

func (t ImeModeStyle) Name() string {
	return "ime-mode"
}

const ImeModeStyleAuto ImeModeStyle = "auto"

const ImeModeStyleNormal ImeModeStyle = "normal"

const ImeModeStyleActive ImeModeStyle = "active"

const ImeModeStyleInactive ImeModeStyle = "inactive"

const ImeModeStyleDisabled ImeModeStyle = "disabled"

var ImeModeStyleBrowserVariantsList = []string{}

func (t ImeModeStyle) BrowserVariants() []string {
	return ImeModeStyleBrowserVariantsList
}

var ImeModeStyleUtilitiesMapping = map[string]string{
	"ime-mode":          "ime-mode: auto",
	"ime-mode-auto":     "ime-mode: auto",
	"ime-mode-normal":   "ime-mode: normal",
	"ime-mode-active":   "ime-mode: active",
	"ime-mode-inactive": "ime-mode: inactive",
	"ime-mode-disabled": "ime-mode: disabled",
}

func (t ImeModeStyle) Utilities() map[string]string {
	return ImeModeStyleUtilitiesMapping
}

func (t ImeModeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ImeModeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// WhiteSpace represent the CSS style "white-space" with value "normal | pre | nowrap | pre-wrap | pre-line | break-spaces"
// See https://developer.mozilla.org/docs/Web/CSS/white-space
type WhiteSpaceStyle string

func (t WhiteSpaceStyle) Name() string {
	return "white-space"
}

const WhiteSpaceStyleNormal WhiteSpaceStyle = "normal"

const WhiteSpaceStylePre WhiteSpaceStyle = "pre"

const WhiteSpaceStyleNowrap WhiteSpaceStyle = "nowrap"

const WhiteSpaceStylePreWrap WhiteSpaceStyle = "pre-wrap"

const WhiteSpaceStylePreLine WhiteSpaceStyle = "pre-line"

const WhiteSpaceStyleBreakSpaces WhiteSpaceStyle = "break-spaces"

var WhiteSpaceStyleBrowserVariantsList = []string{}

func (t WhiteSpaceStyle) BrowserVariants() []string {
	return WhiteSpaceStyleBrowserVariantsList
}

var WhiteSpaceStyleUtilitiesMapping = map[string]string{
	"white-space":              "white-space: normal",
	"white-space-normal":       "white-space: normal",
	"white-space-pre":          "white-space: pre",
	"white-space-nowrap":       "white-space: nowrap",
	"white-space-pre-wrap":     "white-space: pre-wrap",
	"white-space-pre-line":     "white-space: pre-line",
	"white-space-break-spaces": "white-space: break-spaces",
}

func (t WhiteSpaceStyle) Utilities() map[string]string {
	return WhiteSpaceStyleUtilitiesMapping
}

func (t WhiteSpaceStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = WhiteSpaceStyleUtilitiesMapping[tu]
	return value, hasValue
}

// CaptionSide represent the CSS style "caption-side" with value "top | bottom | block-start | block-end | inline-start | inline-end"
// See https://developer.mozilla.org/docs/Web/CSS/caption-side
type CaptionSideStyle string

func (t CaptionSideStyle) Name() string {
	return "caption-side"
}

const CaptionSideStyleTop CaptionSideStyle = "top"

const CaptionSideStyleBottom CaptionSideStyle = "bottom"

const CaptionSideStyleBlockStart CaptionSideStyle = "block-start"

const CaptionSideStyleBlockEnd CaptionSideStyle = "block-end"

const CaptionSideStyleInlineStart CaptionSideStyle = "inline-start"

const CaptionSideStyleInlineEnd CaptionSideStyle = "inline-end"

var CaptionSideStyleBrowserVariantsList = []string{}

func (t CaptionSideStyle) BrowserVariants() []string {
	return CaptionSideStyleBrowserVariantsList
}

var CaptionSideStyleUtilitiesMapping = map[string]string{
	"caption-side":              "caption-side: top",
	"caption-side-top":          "caption-side: top",
	"caption-side-bottom":       "caption-side: bottom",
	"caption-side-block-start":  "caption-side: block-start",
	"caption-side-block-end":    "caption-side: block-end",
	"caption-side-inline-start": "caption-side: inline-start",
	"caption-side-inline-end":   "caption-side: inline-end",
}

func (t CaptionSideStyle) Utilities() map[string]string {
	return CaptionSideStyleUtilitiesMapping
}

func (t CaptionSideStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = CaptionSideStyleUtilitiesMapping[tu]
	return value, hasValue
}

// ColumnCount represent the CSS style "column-count" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-count
type ColumnCountStyle string

func (t ColumnCountStyle) Name() string {
	return "column-count"
}

var ColumnCountStyleBrowserVariantsList = []string{}

func (t ColumnCountStyle) BrowserVariants() []string {
	return ColumnCountStyleBrowserVariantsList
}

var ColumnCountStyleUtilitiesMapping = map[string]string{
	"column-count": "column-count: auto",
}

func (t ColumnCountStyle) Utilities() map[string]string {
	return ColumnCountStyleUtilitiesMapping
}

func (t ColumnCountStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = ColumnCountStyleUtilitiesMapping[tu]
	return value, hasValue
}

// AlignTracks represent the CSS style "align-tracks" with value "normal | end | start | stretch | flex-start | flex-end | center |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe | unsafe"
// See https://developer.mozilla.org/docs/Web/CSS/align-tracks
type AlignTracksStyle string

func (t AlignTracksStyle) Name() string {
	return "align-tracks"
}

const AlignTracksStyleNormal AlignTracksStyle = "normal"

const AlignTracksStyleEnd AlignTracksStyle = "end"

const AlignTracksStyleStart AlignTracksStyle = "start"

const AlignTracksStyleStretch AlignTracksStyle = "stretch"

const AlignTracksStyleFlexStart AlignTracksStyle = "flex-start"

const AlignTracksStyleFlexEnd AlignTracksStyle = "flex-end"

const AlignTracksStyleCenter AlignTracksStyle = "center"

const AlignTracksStyleBaseline AlignTracksStyle = "baseline"

const AlignTracksStyleFirstBaseline AlignTracksStyle = "first-baseline"

const AlignTracksStyleLastBaseline AlignTracksStyle = "last-baseline"

const AlignTracksStyleSpaceBetween AlignTracksStyle = "space-between"

const AlignTracksStyleSpaceAround AlignTracksStyle = "space-around"

const AlignTracksStyleSpaceEvenly AlignTracksStyle = "space-evenly"

const AlignTracksStyleSafe AlignTracksStyle = "safe"

const AlignTracksStyleUnsafe AlignTracksStyle = "unsafe"

var AlignTracksStyleBrowserVariantsList = []string{}

func (t AlignTracksStyle) BrowserVariants() []string {
	return AlignTracksStyleBrowserVariantsList
}

var AlignTracksStyleUtilitiesMapping = map[string]string{
	"align-tracks":                "align-tracks: normal",
	"align-tracks-normal":         "align-tracks: normal",
	"align-tracks-end":            "align-tracks: end",
	"align-tracks-start":          "align-tracks: start",
	"align-tracks-stretch":        "align-tracks: stretch",
	"align-tracks-flex-start":     "align-tracks: flex-start",
	"align-tracks-flex-end":       "align-tracks: flex-end",
	"align-tracks-center":         "align-tracks: center",
	"align-tracks-baseline":       "align-tracks: baseline",
	"align-tracks-first-baseline": "align-tracks: first-baseline",
	"align-tracks-last-baseline":  "align-tracks: last-baseline",
	"align-tracks-space-between":  "align-tracks: space-between",
	"align-tracks-space-around":   "align-tracks: space-around",
	"align-tracks-space-evenly":   "align-tracks: space-evenly",
	"align-tracks-safe":           "align-tracks: safe",
	"align-tracks-unsafe":         "align-tracks: unsafe",
}

func (t AlignTracksStyle) Utilities() map[string]string {
	return AlignTracksStyleUtilitiesMapping
}

func (t AlignTracksStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = AlignTracksStyleUtilitiesMapping[tu]
	return value, hasValue
}

// FontKerning represent the CSS style "font-kerning" with value "auto | normal | none"
// See https://developer.mozilla.org/docs/Web/CSS/font-kerning
type FontKerningStyle string

func (t FontKerningStyle) Name() string {
	return "font-kerning"
}

const FontKerningStyleAuto FontKerningStyle = "auto"

const FontKerningStyleNormal FontKerningStyle = "normal"

const FontKerningStyleNone FontKerningStyle = "none"

var FontKerningStyleBrowserVariantsList = []string{}

func (t FontKerningStyle) BrowserVariants() []string {
	return FontKerningStyleBrowserVariantsList
}

var FontKerningStyleUtilitiesMapping = map[string]string{
	"font-kerning":        "font-kerning: auto",
	"font-kerning-auto":   "font-kerning: auto",
	"font-kerning-normal": "font-kerning: normal",
	"font-kerning-none":   "font-kerning: none",
}

func (t FontKerningStyle) Utilities() map[string]string {
	return FontKerningStyleUtilitiesMapping
}

func (t FontKerningStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = FontKerningStyleUtilitiesMapping[tu]
	return value, hasValue
}

// InsetInlineStart represent the CSS style "inset-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-inline-start
type InsetInlineStartStyle string

func (t InsetInlineStartStyle) Name() string {
	return "inset-inline-start"
}

var InsetInlineStartStyleBrowserVariantsList = []string{}

func (t InsetInlineStartStyle) BrowserVariants() []string {
	return InsetInlineStartStyleBrowserVariantsList
}

var InsetInlineStartStyleUtilitiesMapping = map[string]string{
	"inset-inline-start": "inset-inline-start: auto",
}

func (t InsetInlineStartStyle) Utilities() map[string]string {
	return InsetInlineStartStyleUtilitiesMapping
}

func (t InsetInlineStartStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = InsetInlineStartStyleUtilitiesMapping[tu]
	return value, hasValue
}

// MaskSize represent the CSS style "mask-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-size
type MaskSizeStyle string

func (t MaskSizeStyle) Name() string {
	return "mask-size"
}

var MaskSizeStyleBrowserVariantsList = []string{
	"-webkit-mask-size",
}

func (t MaskSizeStyle) BrowserVariants() []string {
	return MaskSizeStyleBrowserVariantsList
}

var MaskSizeStyleUtilitiesMapping = map[string]string{
	"mask-size": "mask-size: auto",
}

func (t MaskSizeStyle) Utilities() map[string]string {
	return MaskSizeStyleUtilitiesMapping
}

func (t MaskSizeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = MaskSizeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// PageBreakBefore represent the CSS style "page-break-before" with value "auto | always | avoid | left | right | recto | verso"
// See https://developer.mozilla.org/docs/Web/CSS/page-break-before
type PageBreakBeforeStyle string

func (t PageBreakBeforeStyle) Name() string {
	return "page-break-before"
}

const PageBreakBeforeStyleAuto PageBreakBeforeStyle = "auto"

const PageBreakBeforeStyleAlways PageBreakBeforeStyle = "always"

const PageBreakBeforeStyleAvoid PageBreakBeforeStyle = "avoid"

const PageBreakBeforeStyleLeft PageBreakBeforeStyle = "left"

const PageBreakBeforeStyleRight PageBreakBeforeStyle = "right"

const PageBreakBeforeStyleRecto PageBreakBeforeStyle = "recto"

const PageBreakBeforeStyleVerso PageBreakBeforeStyle = "verso"

var PageBreakBeforeStyleBrowserVariantsList = []string{}

func (t PageBreakBeforeStyle) BrowserVariants() []string {
	return PageBreakBeforeStyleBrowserVariantsList
}

var PageBreakBeforeStyleUtilitiesMapping = map[string]string{
	"page-break-before":        "page-break-before: auto",
	"page-break-before-auto":   "page-break-before: auto",
	"page-break-before-always": "page-break-before: always",
	"page-break-before-avoid":  "page-break-before: avoid",
	"page-break-before-left":   "page-break-before: left",
	"page-break-before-right":  "page-break-before: right",
	"page-break-before-recto":  "page-break-before: recto",
	"page-break-before-verso":  "page-break-before: verso",
}

func (t PageBreakBeforeStyle) Utilities() map[string]string {
	return PageBreakBeforeStyleUtilitiesMapping
}

func (t PageBreakBeforeStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = PageBreakBeforeStyleUtilitiesMapping[tu]
	return value, hasValue
}

// TextEmphasisPositionFirst represent the CSS style "text-emphasis-position-first" with value " over | under"
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-position
type TextEmphasisPositionFirstStyle string

func (t TextEmphasisPositionFirstStyle) Name() string {
	return "text-emphasis-position-first"
}

const TextEmphasisPositionFirstStyleOver TextEmphasisPositionFirstStyle = "over"

const TextEmphasisPositionFirstStyleUnder TextEmphasisPositionFirstStyle = "under"

var TextEmphasisPositionFirstStyleBrowserVariantsList = []string{}

func (t TextEmphasisPositionFirstStyle) BrowserVariants() []string {
	return TextEmphasisPositionFirstStyleBrowserVariantsList
}

var TextEmphasisPositionFirstStyleUtilitiesMapping = map[string]string{
	"text-emphasis-position-first":       "text-emphasis-position-first: over right",
	"text-emphasis-position-first-over":  "text-emphasis-position-first: over",
	"text-emphasis-position-first-under": "text-emphasis-position-first: under",
}

func (t TextEmphasisPositionFirstStyle) Utilities() map[string]string {
	return TextEmphasisPositionFirstStyleUtilitiesMapping
}

func (t TextEmphasisPositionFirstStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = TextEmphasisPositionFirstStyleUtilitiesMapping[tu]
	return value, hasValue
}

// BackgroundPositionX represent the CSS style "background-position-x" with value "center | left | right | x-start | x-end"
// See https://developer.mozilla.org/docs/Web/CSS/background-position-x
type BackgroundPositionXStyle string

func (t BackgroundPositionXStyle) Name() string {
	return "background-position-x"
}

const BackgroundPositionXStyleCenter BackgroundPositionXStyle = "center"

const BackgroundPositionXStyleLeft BackgroundPositionXStyle = "left"

const BackgroundPositionXStyleRight BackgroundPositionXStyle = "right"

const BackgroundPositionXStyleXStart BackgroundPositionXStyle = "x-start"

const BackgroundPositionXStyleXEnd BackgroundPositionXStyle = "x-end"

var BackgroundPositionXStyleBrowserVariantsList = []string{}

func (t BackgroundPositionXStyle) BrowserVariants() []string {
	return BackgroundPositionXStyleBrowserVariantsList
}

var BackgroundPositionXStyleUtilitiesMapping = map[string]string{
	"background-position-x":         "background-position-x: left",
	"background-position-x-center":  "background-position-x: center",
	"background-position-x-left":    "background-position-x: left",
	"background-position-x-right":   "background-position-x: right",
	"background-position-x-x-start": "background-position-x: x-start",
	"background-position-x-x-end":   "background-position-x: x-end",
}

func (t BackgroundPositionXStyle) Utilities() map[string]string {
	return BackgroundPositionXStyleUtilitiesMapping
}

func (t BackgroundPositionXStyle) UtilityFor(tu string) (string, bool) {
	var value, hasValue = BackgroundPositionXStyleUtilitiesMapping[tu]
	return value, hasValue
}
