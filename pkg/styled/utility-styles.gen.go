package styled

import "time"

// FontFamily represent the CSS style "font-family" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-family
type FontFamilyStyle string

func (t FontFamilyStyle) Name() string {
	return "font-family"
}

func (t FontFamilyStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontFamilyStyle) Utilities() map[string]string {
	return map[string]string{
		"font-family": "font-family: dependsOnUserAgent",
	}
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

func (t TextAlignLastStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextAlignLastStyle) Utilities() map[string]string {
	return map[string]string{
		"text-align-last":         "text-align-last: auto",
		"text-align-last-auto":    "text-align-last: auto",
		"text-align-last-start":   "text-align-last: start",
		"text-align-last-end":     "text-align-last: end",
		"text-align-last-left":    "text-align-last: left",
		"text-align-last-right":   "text-align-last: right",
		"text-align-last-center":  "text-align-last: center",
		"text-align-last-justify": "text-align-last: justify",
	}
}

// BorderImageOutset represent the CSS style "border-image-outset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-outset
type BorderImageOutsetStyle string

func (t BorderImageOutsetStyle) Name() string {
	return "border-image-outset"
}

func (t BorderImageOutsetStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderImageOutsetStyle) Utilities() map[string]string {
	return map[string]string{
		"border-image-outset": "border-image-outset: 0",
	}
}

// FontSizeAdjust represent the CSS style "font-size-adjust" with value "none"
// See https://developer.mozilla.org/docs/Web/CSS/font-size-adjust
type FontSizeAdjustStyle string

func (t FontSizeAdjustStyle) Name() string {
	return "font-size-adjust"
}

const FontSizeAdjustStyleNone FontSizeAdjustStyle = "none"

func (t FontSizeAdjustStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontSizeAdjustStyle) Utilities() map[string]string {
	return map[string]string{
		"font-size-adjust":      "font-size-adjust: none",
		"font-size-adjust-none": "font-size-adjust: none",
	}
}

// InsetInline represent the CSS style "inset-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-inline
type InsetInlineStyle string

func (t InsetInlineStyle) Name() string {
	return "inset-inline"
}

func (t InsetInlineStyle) BrowserVariants() []string {
	return []string{}
}

func (t InsetInlineStyle) Utilities() map[string]string {
	return map[string]string{
		"inset-inline": "inset-inline: auto",
	}
}

// ScrollMarginBlock represent the CSS style "scroll-margin-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-block
type ScrollMarginBlockStyle float64

func (t ScrollMarginBlockStyle) Name() string {
	return "scroll-margin-block"
}

func (t ScrollMarginBlockStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginBlockStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin-block": "scroll-margin-block: 0",
	}
}

// ScrollPaddingInline represent the CSS style "scroll-padding-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-inline
type ScrollPaddingInlineStyle string

func (t ScrollPaddingInlineStyle) Name() string {
	return "scroll-padding-inline"
}

func (t ScrollPaddingInlineStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingInlineStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding-inline": "scroll-padding-inline: auto",
	}
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

func (t AlignItemsStyle) BrowserVariants() []string {
	return []string{}
}

func (t AlignItemsStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// BorderImageSource represent the CSS style "border-image-source" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-source
type BorderImageSourceStyle string

func (t BorderImageSourceStyle) Name() string {
	return "border-image-source"
}

func (t BorderImageSourceStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderImageSourceStyle) Utilities() map[string]string {
	return map[string]string{
		"border-image-source": "border-image-source: none",
	}
}

// Widows represent the CSS style "widows" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/widows
type WidowsStyle float64

func (t WidowsStyle) Name() string {
	return "widows"
}

func (t WidowsStyle) BrowserVariants() []string {
	return []string{}
}

func (t WidowsStyle) Utilities() map[string]string {
	return map[string]string{
		"widows": "widows: 2",
	}
}

// BorderSpacing represent the CSS style "border-spacing" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-spacing
type BorderSpacingStyle string

func (t BorderSpacingStyle) Name() string {
	return "border-spacing"
}

func (t BorderSpacingStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderSpacingStyle) Utilities() map[string]string {
	return map[string]string{
		"border-spacing": "border-spacing: 0",
	}
}

// MaskClip represent the CSS style "mask-clip" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-clip
type MaskClipStyle string

func (t MaskClipStyle) Name() string {
	return "mask-clip"
}

func (t MaskClipStyle) BrowserVariants() []string {
	return []string{
		"-webkit-mask-clip",
	}
}

func (t MaskClipStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-clip": "mask-clip: border-box",
	}
}

// MaxHeight represent the CSS style "max-height" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-height
type MaxHeightStyle string

func (t MaxHeightStyle) Name() string {
	return "max-height"
}

func (t MaxHeightStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaxHeightStyle) Utilities() map[string]string {
	return map[string]string{
		"max-height": "max-height: none",
	}
}

// TextEmphasisPositionFirst represent the CSS style "text-emphasis-position-first" with value " over | under"
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-position
type TextEmphasisPositionFirstStyle string

func (t TextEmphasisPositionFirstStyle) Name() string {
	return "text-emphasis-position-first"
}

const TextEmphasisPositionFirstStyleOver TextEmphasisPositionFirstStyle = "over"

const TextEmphasisPositionFirstStyleUnder TextEmphasisPositionFirstStyle = "under"

func (t TextEmphasisPositionFirstStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextEmphasisPositionFirstStyle) Utilities() map[string]string {
	return map[string]string{
		"text-emphasis-position-first":       "text-emphasis-position-first: over right",
		"text-emphasis-position-first-over":  "text-emphasis-position-first: over",
		"text-emphasis-position-first-under": "text-emphasis-position-first: under",
	}
}

// BorderBlockEndWidth represent the CSS style "border-block-end-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-end-width
type BorderBlockEndWidthStyle string

func (t BorderBlockEndWidthStyle) Name() string {
	return "border-block-end-width"
}

func (t BorderBlockEndWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBlockEndWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-block-end-width": "border-block-end-width: medium",
	}
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

func (t ImeModeStyle) BrowserVariants() []string {
	return []string{}
}

func (t ImeModeStyle) Utilities() map[string]string {
	return map[string]string{
		"ime-mode":          "ime-mode: auto",
		"ime-mode-auto":     "ime-mode: auto",
		"ime-mode-normal":   "ime-mode: normal",
		"ime-mode-active":   "ime-mode: active",
		"ime-mode-inactive": "ime-mode: inactive",
		"ime-mode-disabled": "ime-mode: disabled",
	}
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

func (t BackgroundPositionXStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundPositionXStyle) Utilities() map[string]string {
	return map[string]string{
		"background-position-x":         "background-position-x: left",
		"background-position-x-center":  "background-position-x: center",
		"background-position-x-left":    "background-position-x: left",
		"background-position-x-right":   "background-position-x: right",
		"background-position-x-x-start": "background-position-x: x-start",
		"background-position-x-x-end":   "background-position-x: x-end",
	}
}

// MaskMode represent the CSS style "mask-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-mode
type MaskModeStyle string

func (t MaskModeStyle) Name() string {
	return "mask-mode"
}

func (t MaskModeStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaskModeStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-mode": "mask-mode: match-source",
	}
}

// InsetInlineEnd represent the CSS style "inset-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-inline-end
type InsetInlineEndStyle string

func (t InsetInlineEndStyle) Name() string {
	return "inset-inline-end"
}

func (t InsetInlineEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t InsetInlineEndStyle) Utilities() map[string]string {
	return map[string]string{
		"inset-inline-end": "inset-inline-end: auto",
	}
}

// MarginInlineEnd represent the CSS style "margin-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-inline-end
type MarginInlineEndStyle string

func (t MarginInlineEndStyle) Name() string {
	return "margin-inline-end"
}

func (t MarginInlineEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginInlineEndStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-inline-end": "margin-inline-end: 0",
	}
}

// ScrollPaddingBlock represent the CSS style "scroll-padding-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-block
type ScrollPaddingBlockStyle string

func (t ScrollPaddingBlockStyle) Name() string {
	return "scroll-padding-block"
}

func (t ScrollPaddingBlockStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingBlockStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding-block": "scroll-padding-block: auto",
	}
}

// TransformScaleY represent the CSS style "transform-scale-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformScaleYStyle string

func (t TransformScaleYStyle) Name() string {
	return "transform-scale-y"
}

func (t TransformScaleYStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformScaleYStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-scale-y": "transform-scale-y: scaleY(0)",
	}
}

// TransformRotateX represent the CSS style "transform-rotate-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformRotateXStyle string

func (t TransformRotateXStyle) Name() string {
	return "transform-rotate-x"
}

func (t TransformRotateXStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformRotateXStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-rotate-x": "transform-rotate-x: rotateX(0)",
	}
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

func (t BorderImageRepeatStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderImageRepeatStyle) Utilities() map[string]string {
	return map[string]string{
		"border-image-repeat":         "border-image-repeat: stretch",
		"border-image-repeat-stretch": "border-image-repeat: stretch",
		"border-image-repeat-repeat":  "border-image-repeat: repeat",
		"border-image-repeat-round":   "border-image-repeat: round",
		"border-image-repeat-space":   "border-image-repeat: space",
	}
}

// BorderRightWidth represent the CSS style "border-right-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-right-width
type BorderRightWidthStyle string

func (t BorderRightWidthStyle) Name() string {
	return "border-right-width"
}

func (t BorderRightWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderRightWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-right-width": "border-right-width: medium",
	}
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

func (t ScrollbarGutterStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollbarGutterStyle) Utilities() map[string]string {
	return map[string]string{
		"scrollbar-gutter":        "scrollbar-gutter: auto",
		"scrollbar-gutter-auto":   "scrollbar-gutter: auto",
		"scrollbar-gutter-stable": "scrollbar-gutter: stable",
		"scrollbar-gutter-always": "scrollbar-gutter: always",
		"scrollbar-gutter-both":   "scrollbar-gutter: both",
		"scrollbar-gutter-force":  "scrollbar-gutter: force",
	}
}

// ScrollMarginBlockEnd represent the CSS style "scroll-margin-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-block-end
type ScrollMarginBlockEndStyle float64

func (t ScrollMarginBlockEndStyle) Name() string {
	return "scroll-margin-block-end"
}

func (t ScrollMarginBlockEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginBlockEndStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin-block-end": "scroll-margin-block-end: 0",
	}
}

// ScrollPaddingBlockStart represent the CSS style "scroll-padding-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-block-start
type ScrollPaddingBlockStartStyle string

func (t ScrollPaddingBlockStartStyle) Name() string {
	return "scroll-padding-block-start"
}

func (t ScrollPaddingBlockStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingBlockStartStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding-block-start": "scroll-padding-block-start: auto",
	}
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

func (t TextDecorationSkipStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextDecorationSkipStyle) Utilities() map[string]string {
	return map[string]string{
		"text-decoration-skip":                 "text-decoration-skip: objects",
		"text-decoration-skip-none":            "text-decoration-skip: none",
		"text-decoration-skip-objects":         "text-decoration-skip: objects",
		"text-decoration-skip-spaces":          "text-decoration-skip: spaces",
		"text-decoration-skip-leading-spaces":  "text-decoration-skip: leading-spaces",
		"text-decoration-skip-trailing-spaces": "text-decoration-skip: trailing-spaces",
		"text-decoration-skip-edges":           "text-decoration-skip: edges",
		"text-decoration-skip-box-decoration":  "text-decoration-skip: box-decoration",
	}
}

// FloodColor represent the CSS style "flood-color" with value ""
// See
type FloodColorStyle Color

func (t FloodColorStyle) Name() string {
	return "flood-color"
}

func (t FloodColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t FloodColorStyle) Utilities() map[string]string {
	return map[string]string{
		"flood-color": "flood-color: black",
	}
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

func (t PointerEventsStyle) BrowserVariants() []string {
	return []string{}
}

func (t PointerEventsStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// InsetBlock represent the CSS style "inset-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-block
type InsetBlockStyle string

func (t InsetBlockStyle) Name() string {
	return "inset-block"
}

func (t InsetBlockStyle) BrowserVariants() []string {
	return []string{}
}

func (t InsetBlockStyle) Utilities() map[string]string {
	return map[string]string{
		"inset-block": "inset-block: auto",
	}
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

func (t OverflowInlineStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverflowInlineStyle) Utilities() map[string]string {
	return map[string]string{
		"overflow-inline":         "overflow-inline: auto",
		"overflow-inline-visible": "overflow-inline: visible",
		"overflow-inline-hidden":  "overflow-inline: hidden",
		"overflow-inline-clip":    "overflow-inline: clip",
		"overflow-inline-scroll":  "overflow-inline: scroll",
		"overflow-inline-auto":    "overflow-inline: auto",
	}
}

// TransformSkewX represent the CSS style "transform-skew-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformSkewXStyle string

func (t TransformSkewXStyle) Name() string {
	return "transform-skew-x"
}

func (t TransformSkewXStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformSkewXStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-skew-x": "transform-skew-x: skewX(0)",
	}
}

// BackfaceVisibility represent the CSS style "backface-visibility" with value "visible | hidden"
// See https://developer.mozilla.org/docs/Web/CSS/backface-visibility
type BackfaceVisibilityStyle string

func (t BackfaceVisibilityStyle) Name() string {
	return "backface-visibility"
}

const BackfaceVisibilityStyleVisible BackfaceVisibilityStyle = "visible"

const BackfaceVisibilityStyleHidden BackfaceVisibilityStyle = "hidden"

func (t BackfaceVisibilityStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackfaceVisibilityStyle) Utilities() map[string]string {
	return map[string]string{
		"backface-visibility":         "backface-visibility: visible",
		"backface-visibility-visible": "backface-visibility: visible",
		"backface-visibility-hidden":  "backface-visibility: hidden",
	}
}

// BackdropFilter represent the CSS style "backdrop-filter" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/backdrop-filter
type BackdropFilterStyle string

func (t BackdropFilterStyle) Name() string {
	return "backdrop-filter"
}

func (t BackdropFilterStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackdropFilterStyle) Utilities() map[string]string {
	return map[string]string{
		"backdrop-filter": "backdrop-filter: none",
	}
}

// BorderLeftColor represent the CSS style "border-left-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-left-color
type BorderLeftColorStyle Color

func (t BorderLeftColorStyle) Name() string {
	return "border-left-color"
}

func (t BorderLeftColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderLeftColorStyle) Utilities() map[string]string {
	return map[string]string{
		"border-left-color": "border-left-color: currentcolor",
	}
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

func (t ScrollSnapTypeXStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollSnapTypeXStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-snap-type-x":           "scroll-snap-type-x: none",
		"scroll-snap-type-x-none":      "scroll-snap-type-x: none",
		"scroll-snap-type-x-mandatory": "scroll-snap-type-x: mandatory",
		"scroll-snap-type-x-proximity": "scroll-snap-type-x: proximity",
	}
}

// ShapeOutside represent the CSS style "shape-outside" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/shape-outside
type ShapeOutsideStyle string

func (t ShapeOutsideStyle) Name() string {
	return "shape-outside"
}

func (t ShapeOutsideStyle) BrowserVariants() []string {
	return []string{}
}

func (t ShapeOutsideStyle) Utilities() map[string]string {
	return map[string]string{
		"shape-outside": "shape-outside: none",
	}
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

func (t StrokeLinejoinStyle) BrowserVariants() []string {
	return []string{}
}

func (t StrokeLinejoinStyle) Utilities() map[string]string {
	return map[string]string{
		"stroke-linejoin":       "stroke-linejoin: miter",
		"stroke-linejoin-miter": "stroke-linejoin: miter",
		"stroke-linejoin-round": "stroke-linejoin: round",
		"stroke-linejoin-bevel": "stroke-linejoin: bevel",
	}
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

func (t BoxPackStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxPackStyle) Utilities() map[string]string {
	return map[string]string{
		"box-pack":         "box-pack: start",
		"box-pack-start":   "box-pack: start",
		"box-pack-center":  "box-pack: center",
		"box-pack-end":     "box-pack: end",
		"box-pack-justify": "box-pack: justify",
	}
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

func (t VerticalAlignStyle) BrowserVariants() []string {
	return []string{}
}

func (t VerticalAlignStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// ClipRule represent the CSS style "clip-rule" with value "nonzero | evenodd"
// See
type ClipRuleStyle string

func (t ClipRuleStyle) Name() string {
	return "clip-rule"
}

const ClipRuleStyleNonzero ClipRuleStyle = "nonzero"

const ClipRuleStyleEvenodd ClipRuleStyle = "evenodd"

func (t ClipRuleStyle) BrowserVariants() []string {
	return []string{}
}

func (t ClipRuleStyle) Utilities() map[string]string {
	return map[string]string{
		"clip-rule":         "clip-rule: nonzero",
		"clip-rule-nonzero": "clip-rule: nonzero",
		"clip-rule-evenodd": "clip-rule: evenodd",
	}
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

func (t DominantBaselineStyle) BrowserVariants() []string {
	return []string{}
}

func (t DominantBaselineStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// PageBreakInside represent the CSS style "page-break-inside" with value "auto | avoid"
// See https://developer.mozilla.org/docs/Web/CSS/page-break-inside
type PageBreakInsideStyle string

func (t PageBreakInsideStyle) Name() string {
	return "page-break-inside"
}

const PageBreakInsideStyleAuto PageBreakInsideStyle = "auto"

const PageBreakInsideStyleAvoid PageBreakInsideStyle = "avoid"

func (t PageBreakInsideStyle) BrowserVariants() []string {
	return []string{}
}

func (t PageBreakInsideStyle) Utilities() map[string]string {
	return map[string]string{
		"page-break-inside":       "page-break-inside: auto",
		"page-break-inside-auto":  "page-break-inside: auto",
		"page-break-inside-avoid": "page-break-inside: avoid",
	}
}

// Content represent the CSS style "content" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/content
type ContentStyle string

func (t ContentStyle) Name() string {
	return "content"
}

func (t ContentStyle) BrowserVariants() []string {
	return []string{}
}

func (t ContentStyle) Utilities() map[string]string {
	return map[string]string{
		"content": "content: normal",
	}
}

// LineHeightStep represent the CSS style "line-height-step" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/line-height-step
type LineHeightStepStyle float64

func (t LineHeightStepStyle) Name() string {
	return "line-height-step"
}

func (t LineHeightStepStyle) BrowserVariants() []string {
	return []string{}
}

func (t LineHeightStepStyle) Utilities() map[string]string {
	return map[string]string{
		"line-height-step": "line-height-step: 0",
	}
}

// ScrollMarginBlockStart represent the CSS style "scroll-margin-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-block-start
type ScrollMarginBlockStartStyle float64

func (t ScrollMarginBlockStartStyle) Name() string {
	return "scroll-margin-block-start"
}

func (t ScrollMarginBlockStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginBlockStartStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin-block-start": "scroll-margin-block-start: 0",
	}
}

// ScrollPaddingBlockEnd represent the CSS style "scroll-padding-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-block-end
type ScrollPaddingBlockEndStyle string

func (t ScrollPaddingBlockEndStyle) Name() string {
	return "scroll-padding-block-end"
}

func (t ScrollPaddingBlockEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingBlockEndStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding-block-end": "scroll-padding-block-end: auto",
	}
}

// BorderTopStyle represent the CSS style "border-top-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-style
type BorderTopStyleStyle string

func (t BorderTopStyleStyle) Name() string {
	return "border-top-style"
}

func (t BorderTopStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderTopStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"border-top-style": "border-top-style: none",
	}
}

// ScrollSnapPointsX represent the CSS style "scroll-snap-points-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-points-x
type ScrollSnapPointsXStyle string

func (t ScrollSnapPointsXStyle) Name() string {
	return "scroll-snap-points-x"
}

func (t ScrollSnapPointsXStyle) BrowserVariants() []string {
	return []string{
		"-ms-scroll-snap-points-x",
	}
}

func (t ScrollSnapPointsXStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-snap-points-x": "scroll-snap-points-x: none",
	}
}

// MinBlockSize represent the CSS style "min-block-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-block-size
type MinBlockSizeStyle string

func (t MinBlockSizeStyle) Name() string {
	return "min-block-size"
}

func (t MinBlockSizeStyle) BrowserVariants() []string {
	return []string{}
}

func (t MinBlockSizeStyle) Utilities() map[string]string {
	return map[string]string{
		"min-block-size": "min-block-size: 0",
	}
}

// ScrollMarginInline represent the CSS style "scroll-margin-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-inline
type ScrollMarginInlineStyle float64

func (t ScrollMarginInlineStyle) Name() string {
	return "scroll-margin-inline"
}

func (t ScrollMarginInlineStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginInlineStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin-inline": "scroll-margin-inline: 0",
	}
}

// TransformScale represent the CSS style "transform-scale" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformScaleStyle string

func (t TransformScaleStyle) Name() string {
	return "transform-scale"
}

func (t TransformScaleStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformScaleStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-scale": "transform-scale: scale(1,1)",
	}
}

// ForcedColorAdjust represent the CSS style "forced-color-adjust" with value "auto | none"
// See https://developer.mozilla.org/docs/Web/CSS/forced-color-adjust
type ForcedColorAdjustStyle string

func (t ForcedColorAdjustStyle) Name() string {
	return "forced-color-adjust"
}

const ForcedColorAdjustStyleAuto ForcedColorAdjustStyle = "auto"

const ForcedColorAdjustStyleNone ForcedColorAdjustStyle = "none"

func (t ForcedColorAdjustStyle) BrowserVariants() []string {
	return []string{}
}

func (t ForcedColorAdjustStyle) Utilities() map[string]string {
	return map[string]string{
		"forced-color-adjust":      "forced-color-adjust: auto",
		"forced-color-adjust-auto": "forced-color-adjust: auto",
		"forced-color-adjust-none": "forced-color-adjust: none",
	}
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

func (t TextJustifyStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextJustifyStyle) Utilities() map[string]string {
	return map[string]string{
		"text-justify":                 "text-justify: auto",
		"text-justify-auto":            "text-justify: auto",
		"text-justify-inter-character": "text-justify: inter-character",
		"text-justify-inter-word":      "text-justify: inter-word",
		"text-justify-none":            "text-justify: none",
	}
}

// FontLanguageOverride represent the CSS style "font-language-override" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-language-override
type FontLanguageOverrideStyle string

func (t FontLanguageOverrideStyle) Name() string {
	return "font-language-override"
}

func (t FontLanguageOverrideStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontLanguageOverrideStyle) Utilities() map[string]string {
	return map[string]string{
		"font-language-override": "font-language-override: normal",
	}
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

func (t ColorRenderingStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColorRenderingStyle) Utilities() map[string]string {
	return map[string]string{
		"color-rendering":                 "color-rendering: auto",
		"color-rendering-auto":            "color-rendering: auto",
		"color-rendering-optimizeSpeed":   "color-rendering: optimizeSpeed",
		"color-rendering-optimizeQuality": "color-rendering: optimizeQuality",
	}
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

func (t RubyMergeStyle) BrowserVariants() []string {
	return []string{}
}

func (t RubyMergeStyle) Utilities() map[string]string {
	return map[string]string{
		"ruby-merge":          "ruby-merge: separate",
		"ruby-merge-separate": "ruby-merge: separate",
		"ruby-merge-collapse": "ruby-merge: collapse",
		"ruby-merge-auto":     "ruby-merge: auto",
	}
}

// ScrollBehavior represent the CSS style "scroll-behavior" with value "auto | smooth"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-behavior
type ScrollBehaviorStyle string

func (t ScrollBehaviorStyle) Name() string {
	return "scroll-behavior"
}

const ScrollBehaviorStyleAuto ScrollBehaviorStyle = "auto"

const ScrollBehaviorStyleSmooth ScrollBehaviorStyle = "smooth"

func (t ScrollBehaviorStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollBehaviorStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-behavior":        "scroll-behavior: auto",
		"scroll-behavior-auto":   "scroll-behavior: auto",
		"scroll-behavior-smooth": "scroll-behavior: smooth",
	}
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

func (t TextTransformStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextTransformStyle) Utilities() map[string]string {
	return map[string]string{
		"text-transform":                "text-transform: none",
		"text-transform-none":           "text-transform: none",
		"text-transform-capitalize":     "text-transform: capitalize",
		"text-transform-uppercase":      "text-transform: uppercase",
		"text-transform-lowercase":      "text-transform: lowercase",
		"text-transform-full-width":     "text-transform: full-width",
		"text-transform-full-size-kana": "text-transform: full-size-kana",
	}
}

// EmptyCells represent the CSS style "empty-cells" with value "show | hide"
// See https://developer.mozilla.org/docs/Web/CSS/empty-cells
type EmptyCellsStyle string

func (t EmptyCellsStyle) Name() string {
	return "empty-cells"
}

const EmptyCellsStyleShow EmptyCellsStyle = "show"

const EmptyCellsStyleHide EmptyCellsStyle = "hide"

func (t EmptyCellsStyle) BrowserVariants() []string {
	return []string{}
}

func (t EmptyCellsStyle) Utilities() map[string]string {
	return map[string]string{
		"empty-cells":      "empty-cells: show",
		"empty-cells-show": "empty-cells: show",
		"empty-cells-hide": "empty-cells: hide",
	}
}

// FlexBasis represent the CSS style "flex-basis" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/flex-basis
type FlexBasisStyle string

func (t FlexBasisStyle) Name() string {
	return "flex-basis"
}

func (t FlexBasisStyle) BrowserVariants() []string {
	return []string{}
}

func (t FlexBasisStyle) Utilities() map[string]string {
	return map[string]string{
		"flex-basis": "flex-basis: auto",
	}
}

// MarginBottom represent the CSS style "margin-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-bottom
type MarginBottomStyle string

func (t MarginBottomStyle) Name() string {
	return "margin-bottom"
}

func (t MarginBottomStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginBottomStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-bottom": "margin-bottom: 0",
	}
}

// ScrollbarColor represent the CSS style "scrollbar-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scrollbar-color
type ScrollbarColorStyle Color

func (t ScrollbarColorStyle) Name() string {
	return "scrollbar-color"
}

func (t ScrollbarColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollbarColorStyle) Utilities() map[string]string {
	return map[string]string{
		"scrollbar-color": "scrollbar-color: auto",
	}
}

// Marker represent the CSS style "marker" with value ""
// See
type MarkerStyle string

func (t MarkerStyle) Name() string {
	return "marker"
}

func (t MarkerStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarkerStyle) Utilities() map[string]string {
	return map[string]string{
		"marker": "marker: none",
	}
}

// ScrollSnapPointsY represent the CSS style "scroll-snap-points-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-points-y
type ScrollSnapPointsYStyle string

func (t ScrollSnapPointsYStyle) Name() string {
	return "scroll-snap-points-y"
}

func (t ScrollSnapPointsYStyle) BrowserVariants() []string {
	return []string{
		"-ms-scroll-snap-points-y",
	}
}

func (t ScrollSnapPointsYStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-snap-points-y": "scroll-snap-points-y: none",
	}
}

// TransformSkewY represent the CSS style "transform-skew-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformSkewYStyle string

func (t TransformSkewYStyle) Name() string {
	return "transform-skew-y"
}

func (t TransformSkewYStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformSkewYStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-skew-y": "transform-skew-y: skewY(0)",
	}
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

func (t HangingPunctuationStyle) BrowserVariants() []string {
	return []string{}
}

func (t HangingPunctuationStyle) Utilities() map[string]string {
	return map[string]string{
		"hanging-punctuation":                  "hanging-punctuation: none",
		"hanging-punctuation-none":             "hanging-punctuation: none",
		"hanging-punctuation-first--force-end": "hanging-punctuation: first--force-end",
		"hanging-punctuation-allow-end":        "hanging-punctuation: allow-end",
		"hanging-punctuation-last":             "hanging-punctuation: last",
	}
}

// ScrollPaddingBottom represent the CSS style "scroll-padding-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-bottom
type ScrollPaddingBottomStyle string

func (t ScrollPaddingBottomStyle) Name() string {
	return "scroll-padding-bottom"
}

func (t ScrollPaddingBottomStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingBottomStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding-bottom": "scroll-padding-bottom: auto",
	}
}

// TextEmphasisColor represent the CSS style "text-emphasis-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-color
type TextEmphasisColorStyle Color

func (t TextEmphasisColorStyle) Name() string {
	return "text-emphasis-color"
}

func (t TextEmphasisColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextEmphasisColorStyle) Utilities() map[string]string {
	return map[string]string{
		"text-emphasis-color": "text-emphasis-color: currentcolor",
	}
}

// BackgroundClip represent the CSS style "background-clip" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-clip
type BackgroundClipStyle string

func (t BackgroundClipStyle) Name() string {
	return "background-clip"
}

func (t BackgroundClipStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundClipStyle) Utilities() map[string]string {
	return map[string]string{
		"background-clip": "background-clip: border-box",
	}
}

// OutlineColor represent the CSS style "outline-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-color
type OutlineColorStyle Color

func (t OutlineColorStyle) Name() string {
	return "outline-color"
}

func (t OutlineColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t OutlineColorStyle) Utilities() map[string]string {
	return map[string]string{
		"outline-color": "outline-color: invertOrCurrentColor",
	}
}

// InsetInlineStart represent the CSS style "inset-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-inline-start
type InsetInlineStartStyle string

func (t InsetInlineStartStyle) Name() string {
	return "inset-inline-start"
}

func (t InsetInlineStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t InsetInlineStartStyle) Utilities() map[string]string {
	return map[string]string{
		"inset-inline-start": "inset-inline-start: auto",
	}
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

func (t MarginTrimStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginTrimStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-trim":         "margin-trim: none",
		"margin-trim-none":    "margin-trim: none",
		"margin-trim-in-flow": "margin-trim: in-flow",
		"margin-trim-all":     "margin-trim: all",
	}
}

// OutlineStyle represent the CSS style "outline-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-style
type OutlineStyleStyle string

func (t OutlineStyleStyle) Name() string {
	return "outline-style"
}

func (t OutlineStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t OutlineStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"outline-style": "outline-style: none",
	}
}

// ScrollSnapDestination represent the CSS style "scroll-snap-destination" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-destination
type ScrollSnapDestinationStyle string

func (t ScrollSnapDestinationStyle) Name() string {
	return "scroll-snap-destination"
}

func (t ScrollSnapDestinationStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollSnapDestinationStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-snap-destination": "scroll-snap-destination: 0px 0px",
	}
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

func (t AlignmentBaselineStyle) BrowserVariants() []string {
	return []string{}
}

func (t AlignmentBaselineStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// BorderLeftStyle represent the CSS style "border-left-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-left-style
type BorderLeftStyleStyle string

func (t BorderLeftStyleStyle) Name() string {
	return "border-left-style"
}

func (t BorderLeftStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderLeftStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"border-left-style": "border-left-style: none",
	}
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

func (t TextUnderlinePositionStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextUnderlinePositionStyle) Utilities() map[string]string {
	return map[string]string{
		"text-underline-position":           "text-underline-position: auto",
		"text-underline-position-auto":      "text-underline-position: auto",
		"text-underline-position-from-font": "text-underline-position: from-font",
		"text-underline-position-under":     "text-underline-position: under",
		"text-underline-position-left":      "text-underline-position: left",
		"text-underline-position-right":     "text-underline-position: right",
	}
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

func (t VisibilityStyle) BrowserVariants() []string {
	return []string{}
}

func (t VisibilityStyle) Utilities() map[string]string {
	return map[string]string{
		"visibility":          "visibility: visible",
		"visibility-visible":  "visibility: visible",
		"visibility-hidden":   "visibility: hidden",
		"visibility-collapse": "visibility: collapse",
	}
}

// MaskSize represent the CSS style "mask-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-size
type MaskSizeStyle string

func (t MaskSizeStyle) Name() string {
	return "mask-size"
}

func (t MaskSizeStyle) BrowserVariants() []string {
	return []string{
		"-webkit-mask-size",
	}
}

func (t MaskSizeStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-size": "mask-size: auto",
	}
}

// PaddingLeft represent the CSS style "padding-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-left
type PaddingLeftStyle string

func (t PaddingLeftStyle) Name() string {
	return "padding-left"
}

func (t PaddingLeftStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaddingLeftStyle) Utilities() map[string]string {
	return map[string]string{
		"padding-left": "padding-left: 0",
	}
}

// MarkerEnd represent the CSS style "marker-end" with value ""
// See
type MarkerEndStyle string

func (t MarkerEndStyle) Name() string {
	return "marker-end"
}

func (t MarkerEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarkerEndStyle) Utilities() map[string]string {
	return map[string]string{
		"marker-end": "marker-end: none",
	}
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

func (t ImageRenderingStyle) BrowserVariants() []string {
	return []string{}
}

func (t ImageRenderingStyle) Utilities() map[string]string {
	return map[string]string{
		"image-rendering":             "image-rendering: auto",
		"image-rendering-auto":        "image-rendering: auto",
		"image-rendering-crisp-edges": "image-rendering: crisp-edges",
		"image-rendering-pixelated":   "image-rendering: pixelated",
	}
}

// BorderBlockEndColor represent the CSS style "border-block-end-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-end-color
type BorderBlockEndColorStyle Color

func (t BorderBlockEndColorStyle) Name() string {
	return "border-block-end-color"
}

func (t BorderBlockEndColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBlockEndColorStyle) Utilities() map[string]string {
	return map[string]string{
		"border-block-end-color": "border-block-end-color: currentcolor",
	}
}

// Clip represent the CSS style "clip" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/clip
type ClipStyle string

func (t ClipStyle) Name() string {
	return "clip"
}

func (t ClipStyle) BrowserVariants() []string {
	return []string{}
}

func (t ClipStyle) Utilities() map[string]string {
	return map[string]string{
		"clip": "clip: auto",
	}
}

// TextShadow represent the CSS style "text-shadow" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-shadow
type TextShadowStyle string

func (t TextShadowStyle) Name() string {
	return "text-shadow"
}

func (t TextShadowStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextShadowStyle) Utilities() map[string]string {
	return map[string]string{
		"text-shadow": "text-shadow: none",
	}
}

// BackgroundOrigin represent the CSS style "background-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-origin
type BackgroundOriginStyle string

func (t BackgroundOriginStyle) Name() string {
	return "background-origin"
}

func (t BackgroundOriginStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundOriginStyle) Utilities() map[string]string {
	return map[string]string{
		"background-origin": "background-origin: padding-box",
	}
}

// BorderStartStartRadius represent the CSS style "border-start-start-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-start-start-radius
type BorderStartStartRadiusStyle string

func (t BorderStartStartRadiusStyle) Name() string {
	return "border-start-start-radius"
}

func (t BorderStartStartRadiusStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderStartStartRadiusStyle) Utilities() map[string]string {
	return map[string]string{
		"border-start-start-radius": "border-start-start-radius: 0",
	}
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

func (t ListStyleTypeStyle) BrowserVariants() []string {
	return []string{}
}

func (t ListStyleTypeStyle) Utilities() map[string]string {
	return map[string]string{
		"list-style-type":                       "list-style-type: disc",
		"list-style-type-disc":                  "list-style-type: disc",
		"list-style-type-circle":                "list-style-type: circle",
		"list-style-type-square":                "list-style-type: square",
		"list-style-type-decimal":               "list-style-type: decimal",
		"list-style-type-georgian":              "list-style-type: georgian",
		"list-style-type-trad-chinese-informal": "list-style-type: trad-chinese-informal",
		"list-style-type-kannada":               "list-style-type: kannada",
	}
}

// RowGap represent the CSS style "row-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/row-gap
type RowGapStyle float64

func (t RowGapStyle) Name() string {
	return "row-gap"
}

func (t RowGapStyle) BrowserVariants() []string {
	return []string{}
}

func (t RowGapStyle) Utilities() map[string]string {
	return map[string]string{
		"row-gap": "row-gap: normal",
	}
}

// FillRule represent the CSS style "fill-rule" with value "nonzero | evenodd"
// See
type FillRuleStyle string

func (t FillRuleStyle) Name() string {
	return "fill-rule"
}

const FillRuleStyleNonzero FillRuleStyle = "nonzero"

const FillRuleStyleEvenodd FillRuleStyle = "evenodd"

func (t FillRuleStyle) BrowserVariants() []string {
	return []string{}
}

func (t FillRuleStyle) Utilities() map[string]string {
	return map[string]string{
		"fill-rule":         "fill-rule: nonzero",
		"fill-rule-nonzero": "fill-rule: nonzero",
		"fill-rule-evenodd": "fill-rule: evenodd",
	}
}

// GridRows represent the CSS style "grid-rows" with value "none"
// See https://developer.mozilla.org/docs/Web/CSS/-ms-grid-rows
type GridRowsStyle string

func (t GridRowsStyle) Name() string {
	return "grid-rows"
}

const GridRowsStyleNone GridRowsStyle = "none"

func (t GridRowsStyle) BrowserVariants() []string {
	return []string{
		"-ms-grid-rows",
	}
}

func (t GridRowsStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-rows":      "grid-rows: none",
		"grid-rows-none": "grid-rows: none",
	}
}

// Transform represent the CSS style "transform" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform
type TransformStyle string

func (t TransformStyle) Name() string {
	return "transform"
}

func (t TransformStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformStyle) Utilities() map[string]string {
	return map[string]string{
		"transform": "transform: none",
	}
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

func (t HyphensStyle) BrowserVariants() []string {
	return []string{}
}

func (t HyphensStyle) Utilities() map[string]string {
	return map[string]string{
		"hyphens":        "hyphens: manual",
		"hyphens-none":   "hyphens: none",
		"hyphens-manual": "hyphens: manual",
		"hyphens-auto":   "hyphens: auto",
	}
}

// LightingColor represent the CSS style "lighting-color" with value ""
// See
type LightingColorStyle Color

func (t LightingColorStyle) Name() string {
	return "lighting-color"
}

func (t LightingColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t LightingColorStyle) Utilities() map[string]string {
	return map[string]string{
		"lighting-color": "lighting-color: white",
	}
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

func (t AlignSelfStyle) BrowserVariants() []string {
	return []string{}
}

func (t AlignSelfStyle) Utilities() map[string]string {
	return map[string]string{
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

func (t OffsetRotateStyle) BrowserVariants() []string {
	return []string{}
}

func (t OffsetRotateStyle) Utilities() map[string]string {
	return map[string]string{
		"offset-rotate":         "offset-rotate: auto",
		"offset-rotate-auto":    "offset-rotate: auto",
		"offset-rotate-reverse": "offset-rotate: reverse",
		"offset-rotate-angle":   "offset-rotate: angle",
	}
}

// Opacity represent the CSS style "opacity" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/opacity
type OpacityStyle float64

func (t OpacityStyle) Name() string {
	return "opacity"
}

func (t OpacityStyle) BrowserVariants() []string {
	return []string{}
}

func (t OpacityStyle) Utilities() map[string]string {
	return map[string]string{
		"opacity": "opacity: 1.0",
	}
}

// ScaleLeft represent the CSS style "scale-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scale
type ScaleLeftStyle string

func (t ScaleLeftStyle) Name() string {
	return "scale-left"
}

func (t ScaleLeftStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScaleLeftStyle) Utilities() map[string]string {
	return map[string]string{
		"scale-left": "scale-left: none",
	}
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

func (t WhiteSpaceStyle) BrowserVariants() []string {
	return []string{}
}

func (t WhiteSpaceStyle) Utilities() map[string]string {
	return map[string]string{
		"white-space":              "white-space: normal",
		"white-space-normal":       "white-space: normal",
		"white-space-pre":          "white-space: pre",
		"white-space-nowrap":       "white-space: nowrap",
		"white-space-pre-wrap":     "white-space: pre-wrap",
		"white-space-pre-line":     "white-space: pre-line",
		"white-space-break-spaces": "white-space: break-spaces",
	}
}

// MaxLines represent the CSS style "max-lines" with value ""
// See
type MaxLinesStyle string

func (t MaxLinesStyle) Name() string {
	return "max-lines"
}

func (t MaxLinesStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaxLinesStyle) Utilities() map[string]string {
	return map[string]string{
		"max-lines": "max-lines: none",
	}
}

// OffsetAnchor represent the CSS style "offset-anchor" with value ""
// See
type OffsetAnchorStyle string

func (t OffsetAnchorStyle) Name() string {
	return "offset-anchor"
}

func (t OffsetAnchorStyle) BrowserVariants() []string {
	return []string{}
}

func (t OffsetAnchorStyle) Utilities() map[string]string {
	return map[string]string{
		"offset-anchor": "offset-anchor: auto",
	}
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

func (t DisplayStyle) BrowserVariants() []string {
	return []string{}
}

func (t DisplayStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// GridTemplateColumns represent the CSS style "grid-template-columns" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-template-columns
type GridTemplateColumnsStyle string

func (t GridTemplateColumnsStyle) Name() string {
	return "grid-template-columns"
}

func (t GridTemplateColumnsStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridTemplateColumnsStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-template-columns": "grid-template-columns: none",
	}
}

// PaddingBlockStart represent the CSS style "padding-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-block-start
type PaddingBlockStartStyle string

func (t PaddingBlockStartStyle) Name() string {
	return "padding-block-start"
}

func (t PaddingBlockStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaddingBlockStartStyle) Utilities() map[string]string {
	return map[string]string{
		"padding-block-start": "padding-block-start: 0",
	}
}

// BoxSizing represent the CSS style "box-sizing" with value "content-box | border-box"
// See https://developer.mozilla.org/docs/Web/CSS/box-sizing
type BoxSizingStyle string

func (t BoxSizingStyle) Name() string {
	return "box-sizing"
}

const BoxSizingStyleContentBox BoxSizingStyle = "content-box"

const BoxSizingStyleBorderBox BoxSizingStyle = "border-box"

func (t BoxSizingStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxSizingStyle) Utilities() map[string]string {
	return map[string]string{
		"box-sizing":             "box-sizing: content-box",
		"box-sizing-content-box": "box-sizing: content-box",
		"box-sizing-border-box":  "box-sizing: border-box",
	}
}

// ClipPath represent the CSS style "clip-path" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/clip-path
type ClipPathStyle string

func (t ClipPathStyle) Name() string {
	return "clip-path"
}

func (t ClipPathStyle) BrowserVariants() []string {
	return []string{}
}

func (t ClipPathStyle) Utilities() map[string]string {
	return map[string]string{
		"clip-path": "clip-path: none",
	}
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

func (t FontVariantCapsStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontVariantCapsStyle) Utilities() map[string]string {
	return map[string]string{
		"font-variant-caps":                 "font-variant-caps: normal",
		"font-variant-caps-normal":          "font-variant-caps: normal",
		"font-variant-caps-small-caps":      "font-variant-caps: small-caps",
		"font-variant-caps-all-small-caps":  "font-variant-caps: all-small-caps",
		"font-variant-caps-petite-caps":     "font-variant-caps: petite-caps",
		"font-variant-caps-all-petite-caps": "font-variant-caps: all-petite-caps",
		"font-variant-caps-unicase":         "font-variant-caps: unicase",
		"font-variant-caps-titling-caps":    "font-variant-caps: titling-caps",
	}
}

// InitialLetter represent the CSS style "initial-letter" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/initial-letter
type InitialLetterStyle string

func (t InitialLetterStyle) Name() string {
	return "initial-letter"
}

func (t InitialLetterStyle) BrowserVariants() []string {
	return []string{}
}

func (t InitialLetterStyle) Utilities() map[string]string {
	return map[string]string{
		"initial-letter": "initial-letter: normal",
	}
}

// LetterSpacing represent the CSS style "letter-spacing" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/letter-spacing
type LetterSpacingStyle string

func (t LetterSpacingStyle) Name() string {
	return "letter-spacing"
}

func (t LetterSpacingStyle) BrowserVariants() []string {
	return []string{}
}

func (t LetterSpacingStyle) Utilities() map[string]string {
	return map[string]string{
		"letter-spacing": "letter-spacing: normal",
	}
}

// Stroke represent the CSS style "stroke" with value ""
// See
type StrokeStyle Color

func (t StrokeStyle) Name() string {
	return "stroke"
}

func (t StrokeStyle) BrowserVariants() []string {
	return []string{}
}

func (t StrokeStyle) Utilities() map[string]string {
	return map[string]string{
		"stroke": "stroke: none",
	}
}

// BorderLeftWidth represent the CSS style "border-left-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-left-width
type BorderLeftWidthStyle string

func (t BorderLeftWidthStyle) Name() string {
	return "border-left-width"
}

func (t BorderLeftWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderLeftWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-left-width": "border-left-width: medium",
	}
}

// OffsetPath represent the CSS style "offset-path" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/offset-path
type OffsetPathStyle string

func (t OffsetPathStyle) Name() string {
	return "offset-path"
}

func (t OffsetPathStyle) BrowserVariants() []string {
	return []string{}
}

func (t OffsetPathStyle) Utilities() map[string]string {
	return map[string]string{
		"offset-path": "offset-path: none",
	}
}

// TextEmphasisPositionSecond represent the CSS style "text-emphasis-position-second" with value "right | left"
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-position
type TextEmphasisPositionSecondStyle string

func (t TextEmphasisPositionSecondStyle) Name() string {
	return "text-emphasis-position-second"
}

const TextEmphasisPositionSecondStyleRight TextEmphasisPositionSecondStyle = "right"

const TextEmphasisPositionSecondStyleLeft TextEmphasisPositionSecondStyle = "left"

func (t TextEmphasisPositionSecondStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextEmphasisPositionSecondStyle) Utilities() map[string]string {
	return map[string]string{
		"text-emphasis-position-second":       "text-emphasis-position-second: over right",
		"text-emphasis-position-second-right": "text-emphasis-position-second: right",
		"text-emphasis-position-second-left":  "text-emphasis-position-second: left",
	}
}

// MaxBlockSize represent the CSS style "max-block-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-block-size
type MaxBlockSizeStyle string

func (t MaxBlockSizeStyle) Name() string {
	return "max-block-size"
}

func (t MaxBlockSizeStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaxBlockSizeStyle) Utilities() map[string]string {
	return map[string]string{
		"max-block-size": "max-block-size: 0",
	}
}

// OverflowClipBox represent the CSS style "overflow-clip-box" with value "padding-box | content-box"
// See https://developer.mozilla.org/docs/Mozilla/CSS/overflow-clip-box
type OverflowClipBoxStyle string

func (t OverflowClipBoxStyle) Name() string {
	return "overflow-clip-box"
}

const OverflowClipBoxStylePaddingBox OverflowClipBoxStyle = "padding-box"

const OverflowClipBoxStyleContentBox OverflowClipBoxStyle = "content-box"

func (t OverflowClipBoxStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverflowClipBoxStyle) Utilities() map[string]string {
	return map[string]string{
		"overflow-clip-box":             "overflow-clip-box: padding-box",
		"overflow-clip-box-padding-box": "overflow-clip-box: padding-box",
		"overflow-clip-box-content-box": "overflow-clip-box: content-box",
	}
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

func (t CaptionSideStyle) BrowserVariants() []string {
	return []string{}
}

func (t CaptionSideStyle) Utilities() map[string]string {
	return map[string]string{
		"caption-side":              "caption-side: top",
		"caption-side-top":          "caption-side: top",
		"caption-side-bottom":       "caption-side: bottom",
		"caption-side-block-start":  "caption-side: block-start",
		"caption-side-block-end":    "caption-side: block-end",
		"caption-side-inline-start": "caption-side: inline-start",
		"caption-side-inline-end":   "caption-side: inline-end",
	}
}

// ListStyleImage represent the CSS style "list-style-image" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/list-style-image
type ListStyleImageStyle URL

func (t ListStyleImageStyle) Name() string {
	return "list-style-image"
}

func (t ListStyleImageStyle) BrowserVariants() []string {
	return []string{}
}

func (t ListStyleImageStyle) Utilities() map[string]string {
	return map[string]string{
		"list-style-image": "list-style-image: none",
	}
}

// TransformTranslateX represent the CSS style "transform-translate-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformTranslateXStyle string

func (t TransformTranslateXStyle) Name() string {
	return "transform-translate-x"
}

func (t TransformTranslateXStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformTranslateXStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-translate-x": "transform-translate-x: translate(0)",
	}
}

// StrokeMiterlimit represent the CSS style "stroke-miterlimit" with value ""
// See
type StrokeMiterlimitStyle float64

func (t StrokeMiterlimitStyle) Name() string {
	return "stroke-miterlimit"
}

func (t StrokeMiterlimitStyle) BrowserVariants() []string {
	return []string{}
}

func (t StrokeMiterlimitStyle) Utilities() map[string]string {
	return map[string]string{
		"stroke-miterlimit": "stroke-miterlimit: 4",
	}
}

// PaddingBlockEnd represent the CSS style "padding-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-block-end
type PaddingBlockEndStyle string

func (t PaddingBlockEndStyle) Name() string {
	return "padding-block-end"
}

func (t PaddingBlockEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaddingBlockEndStyle) Utilities() map[string]string {
	return map[string]string{
		"padding-block-end": "padding-block-end: 0",
	}
}

// TabSize represent the CSS style "tab-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/tab-size
type TabSizeStyle float64

func (t TabSizeStyle) Name() string {
	return "tab-size"
}

func (t TabSizeStyle) BrowserVariants() []string {
	return []string{}
}

func (t TabSizeStyle) Utilities() map[string]string {
	return map[string]string{
		"tab-size": "tab-size: 8",
	}
}

// OutlineWidth represent the CSS style "outline-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-width
type OutlineWidthStyle string

func (t OutlineWidthStyle) Name() string {
	return "outline-width"
}

func (t OutlineWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t OutlineWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"outline-width": "outline-width: medium",
	}
}

// OverflowAnchor represent the CSS style "overflow-anchor" with value "auto | none"
// See
type OverflowAnchorStyle string

func (t OverflowAnchorStyle) Name() string {
	return "overflow-anchor"
}

const OverflowAnchorStyleAuto OverflowAnchorStyle = "auto"

const OverflowAnchorStyleNone OverflowAnchorStyle = "none"

func (t OverflowAnchorStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverflowAnchorStyle) Utilities() map[string]string {
	return map[string]string{
		"overflow-anchor":      "overflow-anchor: auto",
		"overflow-anchor-auto": "overflow-anchor: auto",
		"overflow-anchor-none": "overflow-anchor: none",
	}
}

// ScrollPadding represent the CSS style "scroll-padding" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding
type ScrollPaddingStyle string

func (t ScrollPaddingStyle) Name() string {
	return "scroll-padding"
}

func (t ScrollPaddingStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding": "scroll-padding: auto",
	}
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

func (t FontKerningStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontKerningStyle) Utilities() map[string]string {
	return map[string]string{
		"font-kerning":        "font-kerning: auto",
		"font-kerning-auto":   "font-kerning: auto",
		"font-kerning-normal": "font-kerning: normal",
		"font-kerning-none":   "font-kerning: none",
	}
}

// BackgroundPosition represent the CSS style "background-position" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-position
type BackgroundPositionStyle string

func (t BackgroundPositionStyle) Name() string {
	return "background-position"
}

func (t BackgroundPositionStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundPositionStyle) Utilities() map[string]string {
	return map[string]string{
		"background-position": "background-position: 0% 0%",
	}
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

func (t FloatStyle) BrowserVariants() []string {
	return []string{}
}

func (t FloatStyle) Utilities() map[string]string {
	return map[string]string{
		"float":              "float: none",
		"float-left":         "float: left",
		"float-right":        "float: right",
		"float-none":         "float: none",
		"float-inline-start": "float: inline-start",
		"float-inline-end":   "float: inline-end",
	}
}

// InlineSize represent the CSS style "inline-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inline-size
type InlineSizeStyle string

func (t InlineSizeStyle) Name() string {
	return "inline-size"
}

func (t InlineSizeStyle) BrowserVariants() []string {
	return []string{}
}

func (t InlineSizeStyle) Utilities() map[string]string {
	return map[string]string{
		"inline-size": "inline-size: auto",
	}
}

// MaxInlineSize represent the CSS style "max-inline-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-inline-size
type MaxInlineSizeStyle string

func (t MaxInlineSizeStyle) Name() string {
	return "max-inline-size"
}

func (t MaxInlineSizeStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaxInlineSizeStyle) Utilities() map[string]string {
	return map[string]string{
		"max-inline-size": "max-inline-size: 0",
	}
}

// TransitionProperty represent the CSS style "transition-property" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-property
type TransitionPropertyStyle string

func (t TransitionPropertyStyle) Name() string {
	return "transition-property"
}

func (t TransitionPropertyStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransitionPropertyStyle) Utilities() map[string]string {
	return map[string]string{
		"transition-property": "transition-property: all",
	}
}

// MaskBorderSource represent the CSS style "mask-border-source" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-source
type MaskBorderSourceStyle string

func (t MaskBorderSourceStyle) Name() string {
	return "mask-border-source"
}

func (t MaskBorderSourceStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaskBorderSourceStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-border-source": "mask-border-source: none",
	}
}

// TransitionTimingFunction represent the CSS style "transition-timing-function" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-timing-function
type TransitionTimingFunctionStyle string

func (t TransitionTimingFunctionStyle) Name() string {
	return "transition-timing-function"
}

func (t TransitionTimingFunctionStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransitionTimingFunctionStyle) Utilities() map[string]string {
	return map[string]string{
		"transition-timing-function": "transition-timing-function: ease",
	}
}

// ColumnSpan represent the CSS style "column-span" with value "none | all"
// See https://developer.mozilla.org/docs/Web/CSS/column-span
type ColumnSpanStyle string

func (t ColumnSpanStyle) Name() string {
	return "column-span"
}

const ColumnSpanStyleNone ColumnSpanStyle = "none"

const ColumnSpanStyleAll ColumnSpanStyle = "all"

func (t ColumnSpanStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColumnSpanStyle) Utilities() map[string]string {
	return map[string]string{
		"column-span":      "column-span: none",
		"column-span-none": "column-span: none",
		"column-span-all":  "column-span: all",
	}
}

// Direction represent the CSS style "direction" with value "ltr | rtl"
// See https://developer.mozilla.org/docs/Web/CSS/direction
type DirectionStyle string

func (t DirectionStyle) Name() string {
	return "direction"
}

const DirectionStyleLtr DirectionStyle = "ltr"

const DirectionStyleRtl DirectionStyle = "rtl"

func (t DirectionStyle) BrowserVariants() []string {
	return []string{}
}

func (t DirectionStyle) Utilities() map[string]string {
	return map[string]string{
		"direction":     "direction: ltr",
		"direction-ltr": "direction: ltr",
		"direction-rtl": "direction: rtl",
	}
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

func (t OverflowStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverflowStyle) Utilities() map[string]string {
	return map[string]string{
		"overflow":         "overflow: visible",
		"overflow-visible": "overflow: visible",
		"overflow-hidden":  "overflow: hidden",
		"overflow-clip":    "overflow: clip",
		"overflow-scroll":  "overflow: scroll",
		"overflow-auto":    "overflow: auto",
	}
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

func (t StrokeLinecapStyle) BrowserVariants() []string {
	return []string{}
}

func (t StrokeLinecapStyle) Utilities() map[string]string {
	return map[string]string{
		"stroke-linecap":        "stroke-linecap: butt",
		"stroke-linecap-butt":   "stroke-linecap: butt",
		"stroke-linecap-round":  "stroke-linecap: round",
		"stroke-linecap-square": "stroke-linecap: square",
	}
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

func (t OverscrollBehaviorInlineStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverscrollBehaviorInlineStyle) Utilities() map[string]string {
	return map[string]string{
		"overscroll-behavior-inline":         "overscroll-behavior-inline: auto",
		"overscroll-behavior-inline-contain": "overscroll-behavior-inline: contain",
		"overscroll-behavior-inline-none":    "overscroll-behavior-inline: none",
		"overscroll-behavior-inline-auto":    "overscroll-behavior-inline: auto",
	}
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

func (t OverscrollBehaviorStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverscrollBehaviorStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// ScrollPaddingInlineEnd represent the CSS style "scroll-padding-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-inline-end
type ScrollPaddingInlineEndStyle string

func (t ScrollPaddingInlineEndStyle) Name() string {
	return "scroll-padding-inline-end"
}

func (t ScrollPaddingInlineEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingInlineEndStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding-inline-end": "scroll-padding-inline-end: auto",
	}
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

func (t TouchActionStyle) BrowserVariants() []string {
	return []string{}
}

func (t TouchActionStyle) Utilities() map[string]string {
	return map[string]string{
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

func (t AppearanceStyle) BrowserVariants() []string {
	return []string{
		"-webkit-appearance",
		"-moz-appearance",
	}
}

func (t AppearanceStyle) Utilities() map[string]string {
	return map[string]string{
		"appearance":                 "appearance: auto",
		"appearance-none":            "appearance: none",
		"appearance-auto":            "appearance: auto",
		"appearance-textfield":       "appearance: textfield",
		"appearance-menulist-button": "appearance: menulist-button",
		"appearance-compat-auto":     "appearance: compat-auto",
	}
}

// MinWidth represent the CSS style "min-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-width
type MinWidthStyle string

func (t MinWidthStyle) Name() string {
	return "min-width"
}

func (t MinWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t MinWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"min-width": "min-width: auto",
	}
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

func (t HeightStyle) BrowserVariants() []string {
	return []string{}
}

func (t HeightStyle) Utilities() map[string]string {
	return map[string]string{
		"height":             "height: auto",
		"height-auto":        "height: auto",
		"height-min-content": "height: min-content",
		"height-max-content": "height: max-content",
	}
}

// TransformScaleZ represent the CSS style "transform-scale-z" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformScaleZStyle string

func (t TransformScaleZStyle) Name() string {
	return "transform-scale-z"
}

func (t TransformScaleZStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformScaleZStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-scale-z": "transform-scale-z: scaleZ(0)",
	}
}

// BorderBlockStyle represent the CSS style "border-block-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-style
type BorderBlockStyleStyle string

func (t BorderBlockStyleStyle) Name() string {
	return "border-block-style"
}

func (t BorderBlockStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBlockStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"border-block-style": "border-block-style: none",
	}
}

// BorderBottomLeftRadius represent the CSS style "border-bottom-left-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-left-radius
type BorderBottomLeftRadiusStyle string

func (t BorderBottomLeftRadiusStyle) Name() string {
	return "border-bottom-left-radius"
}

func (t BorderBottomLeftRadiusStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBottomLeftRadiusStyle) Utilities() map[string]string {
	return map[string]string{
		"border-bottom-left-radius": "border-bottom-left-radius: 0",
	}
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

func (t TextOverflowStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextOverflowStyle) Utilities() map[string]string {
	return map[string]string{
		"text-overflow":               "text-overflow: clip",
		"text-overflow-clip":          "text-overflow: clip",
		"text-overflow-ellipsis":      "text-overflow: ellipsis",
		"text-overflow-clip-ellipsis": "text-overflow: clip-ellipsis",
		"text-overflow-ellipsis-clip": "text-overflow: ellipsis-clip",
	}
}

// ZIndex represent the CSS style "z-index" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/z-index
type ZIndexStyle string

func (t ZIndexStyle) Name() string {
	return "z-index"
}

func (t ZIndexStyle) BrowserVariants() []string {
	return []string{}
}

func (t ZIndexStyle) Utilities() map[string]string {
	return map[string]string{
		"z-index": "z-index: auto",
	}
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

func (t AzimuthStyle) BrowserVariants() []string {
	return []string{}
}

func (t AzimuthStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// MaskPosition represent the CSS style "mask-position" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-position
type MaskPositionStyle string

func (t MaskPositionStyle) Name() string {
	return "mask-position"
}

func (t MaskPositionStyle) BrowserVariants() []string {
	return []string{
		"-webkit-mask-position",
	}
}

func (t MaskPositionStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-position": "mask-position: center",
	}
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

func (t TextAlignStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextAlignStyle) Utilities() map[string]string {
	return map[string]string{
		"text-align":              "text-align: startOrNamelessValueIfLTRRightIfRTL",
		"text-align-start":        "text-align: start",
		"text-align-end":          "text-align: end",
		"text-align-left":         "text-align: left",
		"text-align-right":        "text-align: right",
		"text-align-center":       "text-align: center",
		"text-align-justify":      "text-align: justify",
		"text-align-match-parent": "text-align: match-parent",
	}
}

// BorderStartEndRadius represent the CSS style "border-start-end-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-start-end-radius
type BorderStartEndRadiusStyle string

func (t BorderStartEndRadiusStyle) Name() string {
	return "border-start-end-radius"
}

func (t BorderStartEndRadiusStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderStartEndRadiusStyle) Utilities() map[string]string {
	return map[string]string{
		"border-start-end-radius": "border-start-end-radius: 0",
	}
}

// BorderTopLeftRadius represent the CSS style "border-top-left-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-left-radius
type BorderTopLeftRadiusStyle string

func (t BorderTopLeftRadiusStyle) Name() string {
	return "border-top-left-radius"
}

func (t BorderTopLeftRadiusStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderTopLeftRadiusStyle) Utilities() map[string]string {
	return map[string]string{
		"border-top-left-radius": "border-top-left-radius: 0",
	}
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

func (t LineBreakStyle) BrowserVariants() []string {
	return []string{}
}

func (t LineBreakStyle) Utilities() map[string]string {
	return map[string]string{
		"line-break":          "line-break: auto",
		"line-break-auto":     "line-break: auto",
		"line-break-loose":    "line-break: loose",
		"line-break-normal":   "line-break: normal",
		"line-break-strict":   "line-break: strict",
		"line-break-anywhere": "line-break: anywhere",
	}
}

// MarginLeft represent the CSS style "margin-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-left
type MarginLeftStyle string

func (t MarginLeftStyle) Name() string {
	return "margin-left"
}

func (t MarginLeftStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginLeftStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-left": "margin-left: 0",
	}
}

// ObjectPosition represent the CSS style "object-position" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/object-position
type ObjectPositionStyle string

func (t ObjectPositionStyle) Name() string {
	return "object-position"
}

func (t ObjectPositionStyle) BrowserVariants() []string {
	return []string{}
}

func (t ObjectPositionStyle) Utilities() map[string]string {
	return map[string]string{
		"object-position": "object-position: 50% 50%",
	}
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

func (t CursorStyle) BrowserVariants() []string {
	return []string{}
}

func (t CursorStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// TransformTranslate3d represent the CSS style "transform-translate-3d" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformTranslate3dStyle string

func (t TransformTranslate3dStyle) Name() string {
	return "transform-translate-3d"
}

func (t TransformTranslate3dStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformTranslate3dStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-translate-3d": "transform-translate-3d: translate3d(0, 0, 0)",
	}
}

// OffsetPosition represent the CSS style "offset-position" with value ""
// See
type OffsetPositionStyle string

func (t OffsetPositionStyle) Name() string {
	return "offset-position"
}

func (t OffsetPositionStyle) BrowserVariants() []string {
	return []string{}
}

func (t OffsetPositionStyle) Utilities() map[string]string {
	return map[string]string{
		"offset-position": "offset-position: auto",
	}
}

// BoxFlex represent the CSS style "box-flex" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/box-flex
type BoxFlexStyle float64

func (t BoxFlexStyle) Name() string {
	return "box-flex"
}

func (t BoxFlexStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxFlexStyle) Utilities() map[string]string {
	return map[string]string{
		"box-flex": "box-flex: 0",
	}
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

func (t MaskBorderRepeatStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaskBorderRepeatStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// StrokeOpacity represent the CSS style "stroke-opacity" with value ""
// See
type StrokeOpacityStyle float64

func (t StrokeOpacityStyle) Name() string {
	return "stroke-opacity"
}

func (t StrokeOpacityStyle) BrowserVariants() []string {
	return []string{}
}

func (t StrokeOpacityStyle) Utilities() map[string]string {
	return map[string]string{
		"stroke-opacity": "stroke-opacity: 1",
	}
}

// BorderTopRightRadius represent the CSS style "border-top-right-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-right-radius
type BorderTopRightRadiusStyle string

func (t BorderTopRightRadiusStyle) Name() string {
	return "border-top-right-radius"
}

func (t BorderTopRightRadiusStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderTopRightRadiusStyle) Utilities() map[string]string {
	return map[string]string{
		"border-top-right-radius": "border-top-right-radius: 0",
	}
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

func (t FontSmoothStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontSmoothStyle) Utilities() map[string]string {
	return map[string]string{
		"font-smooth":        "font-smooth: auto",
		"font-smooth-auto":   "font-smooth: auto",
		"font-smooth-never":  "font-smooth: never",
		"font-smooth-always": "font-smooth: always",
	}
}

// Gap represent the CSS style "gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/gap
type GapStyle string

func (t GapStyle) Name() string {
	return "gap"
}

func (t GapStyle) BrowserVariants() []string {
	return []string{}
}

func (t GapStyle) Utilities() map[string]string {
	return map[string]string{
		"gap": "gap: ",
	}
}

// InsetBlockEnd represent the CSS style "inset-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-block-end
type InsetBlockEndStyle string

func (t InsetBlockEndStyle) Name() string {
	return "inset-block-end"
}

func (t InsetBlockEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t InsetBlockEndStyle) Utilities() map[string]string {
	return map[string]string{
		"inset-block-end": "inset-block-end: auto",
	}
}

// StrokeDashoffset represent the CSS style "stroke-dashoffset" with value ""
// See
type StrokeDashoffsetStyle float64

func (t StrokeDashoffsetStyle) Name() string {
	return "stroke-dashoffset"
}

func (t StrokeDashoffsetStyle) BrowserVariants() []string {
	return []string{}
}

func (t StrokeDashoffsetStyle) Utilities() map[string]string {
	return map[string]string{
		"stroke-dashoffset": "stroke-dashoffset: 0",
	}
}

// AnimationTimingFunction represent the CSS style "animation-timing-function" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-timing-function
type AnimationTimingFunctionStyle string

func (t AnimationTimingFunctionStyle) Name() string {
	return "animation-timing-function"
}

func (t AnimationTimingFunctionStyle) BrowserVariants() []string {
	return []string{}
}

func (t AnimationTimingFunctionStyle) Utilities() map[string]string {
	return map[string]string{
		"animation-timing-function": "animation-timing-function: ease",
	}
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

func (t TextDecorationLineStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextDecorationLineStyle) Utilities() map[string]string {
	return map[string]string{
		"text-decoration-line":                "text-decoration-line: none",
		"text-decoration-line-none":           "text-decoration-line: none",
		"text-decoration-line-underline":      "text-decoration-line: underline",
		"text-decoration-line-overline":       "text-decoration-line: overline",
		"text-decoration-line-line-through":   "text-decoration-line: line-through",
		"text-decoration-line-blink":          "text-decoration-line: blink",
		"text-decoration-line-spelling-error": "text-decoration-line: spelling-error",
		"text-decoration-line-grammar-error":  "text-decoration-line: grammar-error",
	}
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

func (t TextDecorationSkipInkStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextDecorationSkipInkStyle) Utilities() map[string]string {
	return map[string]string{
		"text-decoration-skip-ink":      "text-decoration-skip-ink: auto",
		"text-decoration-skip-ink-auto": "text-decoration-skip-ink: auto",
		"text-decoration-skip-ink-all":  "text-decoration-skip-ink: all",
		"text-decoration-skip-ink-none": "text-decoration-skip-ink: none",
	}
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

func (t JustifySelfStyle) BrowserVariants() []string {
	return []string{}
}

func (t JustifySelfStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// ScrollPaddingTop represent the CSS style "scroll-padding-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-top
type ScrollPaddingTopStyle string

func (t ScrollPaddingTopStyle) Name() string {
	return "scroll-padding-top"
}

func (t ScrollPaddingTopStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingTopStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding-top": "scroll-padding-top: auto",
	}
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

func (t RubyAlignStyle) BrowserVariants() []string {
	return []string{}
}

func (t RubyAlignStyle) Utilities() map[string]string {
	return map[string]string{
		"ruby-align":               "ruby-align: space-around",
		"ruby-align-start":         "ruby-align: start",
		"ruby-align-center":        "ruby-align: center",
		"ruby-align-space-between": "ruby-align: space-between",
		"ruby-align-space-around":  "ruby-align: space-around",
	}
}

// ListStylePosition represent the CSS style "list-style-position" with value "inside | outside"
// See https://developer.mozilla.org/docs/Web/CSS/list-style-position
type ListStylePositionStyle string

func (t ListStylePositionStyle) Name() string {
	return "list-style-position"
}

const ListStylePositionStyleInside ListStylePositionStyle = "inside"

const ListStylePositionStyleOutside ListStylePositionStyle = "outside"

func (t ListStylePositionStyle) BrowserVariants() []string {
	return []string{}
}

func (t ListStylePositionStyle) Utilities() map[string]string {
	return map[string]string{
		"list-style-position":         "list-style-position: outside",
		"list-style-position-inside":  "list-style-position: inside",
		"list-style-position-outside": "list-style-position: outside",
	}
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

func (t ScrollbarWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollbarWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"scrollbar-width":      "scrollbar-width: auto",
		"scrollbar-width-auto": "scrollbar-width: auto",
		"scrollbar-width-thin": "scrollbar-width: thin",
		"scrollbar-width-none": "scrollbar-width: none",
	}
}

// FloodOpacity represent the CSS style "flood-opacity" with value ""
// See
type FloodOpacityStyle float64

func (t FloodOpacityStyle) Name() string {
	return "flood-opacity"
}

func (t FloodOpacityStyle) BrowserVariants() []string {
	return []string{}
}

func (t FloodOpacityStyle) Utilities() map[string]string {
	return map[string]string{
		"flood-opacity": "flood-opacity: 1",
	}
}

// FontVariant represent the CSS style "font-variant" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant
type FontVariantStyle string

func (t FontVariantStyle) Name() string {
	return "font-variant"
}

func (t FontVariantStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontVariantStyle) Utilities() map[string]string {
	return map[string]string{
		"font-variant": "font-variant: normal",
	}
}

// BoxShadow represent the CSS style "box-shadow" with value "string"
// See https://developer.mozilla.org/docs/Web/CSS/box-shadow
type BoxShadowStyle string

func (t BoxShadowStyle) Name() string {
	return "box-shadow"
}

const BoxShadowStyleString BoxShadowStyle = "string"

func (t BoxShadowStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxShadowStyle) Utilities() map[string]string {
	return map[string]string{
		"box-shadow":        "box-shadow: none",
		"box-shadow-string": "box-shadow: string",
	}
}

// GridRowEnd represent the CSS style "grid-row-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-row-end
type GridRowEndStyle string

func (t GridRowEndStyle) Name() string {
	return "grid-row-end"
}

func (t GridRowEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridRowEndStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-row-end": "grid-row-end: auto",
	}
}

// Translate represent the CSS style "translate" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/translate
type TranslateStyle string

func (t TranslateStyle) Name() string {
	return "translate"
}

func (t TranslateStyle) BrowserVariants() []string {
	return []string{}
}

func (t TranslateStyle) Utilities() map[string]string {
	return map[string]string{
		"translate": "translate: none",
	}
}

// ColumnRuleWidth represent the CSS style "column-rule-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-rule-width
type ColumnRuleWidthStyle string

func (t ColumnRuleWidthStyle) Name() string {
	return "column-rule-width"
}

func (t ColumnRuleWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColumnRuleWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"column-rule-width": "column-rule-width: medium",
	}
}

// PaddingInline represent the CSS style "padding-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-inline
type PaddingInlineStyle string

func (t PaddingInlineStyle) Name() string {
	return "padding-inline"
}

func (t PaddingInlineStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaddingInlineStyle) Utilities() map[string]string {
	return map[string]string{
		"padding-inline": "padding-inline: 0",
	}
}

// PaddingInlineStart represent the CSS style "padding-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-inline-start
type PaddingInlineStartStyle string

func (t PaddingInlineStartStyle) Name() string {
	return "padding-inline-start"
}

func (t PaddingInlineStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaddingInlineStartStyle) Utilities() map[string]string {
	return map[string]string{
		"padding-inline-start": "padding-inline-start: 0",
	}
}

// TableLayout represent the CSS style "table-layout" with value "auto | fixed"
// See https://developer.mozilla.org/docs/Web/CSS/table-layout
type TableLayoutStyle string

func (t TableLayoutStyle) Name() string {
	return "table-layout"
}

const TableLayoutStyleAuto TableLayoutStyle = "auto"

const TableLayoutStyleFixed TableLayoutStyle = "fixed"

func (t TableLayoutStyle) BrowserVariants() []string {
	return []string{}
}

func (t TableLayoutStyle) Utilities() map[string]string {
	return map[string]string{
		"table-layout":       "table-layout: auto",
		"table-layout-auto":  "table-layout: auto",
		"table-layout-fixed": "table-layout: fixed",
	}
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

func (t AllStyle) BrowserVariants() []string {
	return []string{}
}

func (t AllStyle) Utilities() map[string]string {
	return map[string]string{
		"all":         "all: noPracticalInitialValue",
		"all-initial": "all: initial",
		"all-inherit": "all: inherit",
		"all-unset":   "all: unset",
		"all-revert":  "all: revert",
	}
}

// ColorAdjust represent the CSS style "color-adjust" with value "economy | exact"
// See https://developer.mozilla.org/docs/Web/CSS/color-adjust
type ColorAdjustStyle string

func (t ColorAdjustStyle) Name() string {
	return "color-adjust"
}

const ColorAdjustStyleEconomy ColorAdjustStyle = "economy"

const ColorAdjustStyleExact ColorAdjustStyle = "exact"

func (t ColorAdjustStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColorAdjustStyle) Utilities() map[string]string {
	return map[string]string{
		"color-adjust":         "color-adjust: economy",
		"color-adjust-economy": "color-adjust: economy",
		"color-adjust-exact":   "color-adjust: exact",
	}
}

// MarginInlineStart represent the CSS style "margin-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-inline-start
type MarginInlineStartStyle string

func (t MarginInlineStartStyle) Name() string {
	return "margin-inline-start"
}

func (t MarginInlineStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginInlineStartStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-inline-start": "margin-inline-start: 0",
	}
}

// MarginTop represent the CSS style "margin-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-top
type MarginTopStyle string

func (t MarginTopStyle) Name() string {
	return "margin-top"
}

func (t MarginTopStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginTopStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-top": "margin-top: 0",
	}
}

// Orphans represent the CSS style "orphans" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/orphans
type OrphansStyle float64

func (t OrphansStyle) Name() string {
	return "orphans"
}

func (t OrphansStyle) BrowserVariants() []string {
	return []string{}
}

func (t OrphansStyle) Utilities() map[string]string {
	return map[string]string{
		"orphans": "orphans: 2",
	}
}

// ScrollMarginInlineStart represent the CSS style "scroll-margin-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-inline-start
type ScrollMarginInlineStartStyle float64

func (t ScrollMarginInlineStartStyle) Name() string {
	return "scroll-margin-inline-start"
}

func (t ScrollMarginInlineStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginInlineStartStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin-inline-start": "scroll-margin-inline-start: 0",
	}
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

func (t PageBreakAfterStyle) BrowserVariants() []string {
	return []string{}
}

func (t PageBreakAfterStyle) Utilities() map[string]string {
	return map[string]string{
		"page-break-after":        "page-break-after: auto",
		"page-break-after-auto":   "page-break-after: auto",
		"page-break-after-always": "page-break-after: always",
		"page-break-after-avoid":  "page-break-after: avoid",
		"page-break-after-left":   "page-break-after: left",
		"page-break-after-right":  "page-break-after: right",
		"page-break-after-recto":  "page-break-after: recto",
		"page-break-after-verso":  "page-break-after: verso",
	}
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

func (t ResizeStyle) BrowserVariants() []string {
	return []string{}
}

func (t ResizeStyle) Utilities() map[string]string {
	return map[string]string{
		"resize":            "resize: none",
		"resize-none":       "resize: none",
		"resize-both":       "resize: both",
		"resize-horizontal": "resize: horizontal",
		"resize-vertical":   "resize: vertical",
		"resize-block":      "resize: block",
		"resize-inline":     "resize: inline",
	}
}

// BorderCollapse represent the CSS style "border-collapse" with value "collapse | separate"
// See https://developer.mozilla.org/docs/Web/CSS/border-collapse
type BorderCollapseStyle string

func (t BorderCollapseStyle) Name() string {
	return "border-collapse"
}

const BorderCollapseStyleCollapse BorderCollapseStyle = "collapse"

const BorderCollapseStyleSeparate BorderCollapseStyle = "separate"

func (t BorderCollapseStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderCollapseStyle) Utilities() map[string]string {
	return map[string]string{
		"border-collapse":          "border-collapse: separate",
		"border-collapse-collapse": "border-collapse: collapse",
		"border-collapse-separate": "border-collapse: separate",
	}
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

func (t BreakAfterStyle) BrowserVariants() []string {
	return []string{}
}

func (t BreakAfterStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// ScrollSnapCoordinate represent the CSS style "scroll-snap-coordinate" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-coordinate
type ScrollSnapCoordinateStyle string

func (t ScrollSnapCoordinateStyle) Name() string {
	return "scroll-snap-coordinate"
}

func (t ScrollSnapCoordinateStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollSnapCoordinateStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-snap-coordinate": "scroll-snap-coordinate: none",
	}
}

// BorderBlockWidth represent the CSS style "border-block-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-width
type BorderBlockWidthStyle string

func (t BorderBlockWidthStyle) Name() string {
	return "border-block-width"
}

func (t BorderBlockWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBlockWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-block-width": "border-block-width: medium",
	}
}

// FontSynthesis represent the CSS style "font-synthesis" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-synthesis
type FontSynthesisStyle string

func (t FontSynthesisStyle) Name() string {
	return "font-synthesis"
}

func (t FontSynthesisStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontSynthesisStyle) Utilities() map[string]string {
	return map[string]string{
		"font-synthesis": "font-synthesis: weight style",
	}
}

// BorderInlineStartColor represent the CSS style "border-inline-start-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-start-color
type BorderInlineStartColorStyle Color

func (t BorderInlineStartColorStyle) Name() string {
	return "border-inline-start-color"
}

func (t BorderInlineStartColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderInlineStartColorStyle) Utilities() map[string]string {
	return map[string]string{
		"border-inline-start-color": "border-inline-start-color: currentcolor",
	}
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

func (t ContentVisibilityStyle) BrowserVariants() []string {
	return []string{}
}

func (t ContentVisibilityStyle) Utilities() map[string]string {
	return map[string]string{
		"content-visibility":         "content-visibility: visible",
		"content-visibility-visible": "content-visibility: visible",
		"content-visibility-auto":    "content-visibility: auto",
		"content-visibility-hidden":  "content-visibility: hidden",
	}
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

func (t JustifyItemsStyle) BrowserVariants() []string {
	return []string{}
}

func (t JustifyItemsStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// VectorEffect represent the CSS style "vector-effect" with value "non-scaling-stroke | none"
// See
type VectorEffectStyle string

func (t VectorEffectStyle) Name() string {
	return "vector-effect"
}

const VectorEffectStyleNonScalingStroke VectorEffectStyle = "non-scaling-stroke"

const VectorEffectStyleNone VectorEffectStyle = "none"

func (t VectorEffectStyle) BrowserVariants() []string {
	return []string{}
}

func (t VectorEffectStyle) Utilities() map[string]string {
	return map[string]string{
		"vector-effect":                    "vector-effect: none",
		"vector-effect-non-scaling-stroke": "vector-effect: non-scaling-stroke",
		"vector-effect-none":               "vector-effect: none",
	}
}

// Quotes represent the CSS style "quotes" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/quotes
type QuotesStyle string

func (t QuotesStyle) Name() string {
	return "quotes"
}

func (t QuotesStyle) BrowserVariants() []string {
	return []string{}
}

func (t QuotesStyle) Utilities() map[string]string {
	return map[string]string{
		"quotes": "quotes: dependsOnUserAgent",
	}
}

// ScrollPaddingRight represent the CSS style "scroll-padding-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-right
type ScrollPaddingRightStyle string

func (t ScrollPaddingRightStyle) Name() string {
	return "scroll-padding-right"
}

func (t ScrollPaddingRightStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingRightStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding-right": "scroll-padding-right: auto",
	}
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

func (t ColumnFillStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColumnFillStyle) Utilities() map[string]string {
	return map[string]string{
		"column-fill":             "column-fill: balance",
		"column-fill-auto":        "column-fill: auto",
		"column-fill-balance":     "column-fill: balance",
		"column-fill-balance-all": "column-fill: balance-all",
	}
}

// GridAutoRows represent the CSS style "grid-auto-rows" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-auto-rows
type GridAutoRowsStyle string

func (t GridAutoRowsStyle) Name() string {
	return "grid-auto-rows"
}

func (t GridAutoRowsStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridAutoRowsStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-auto-rows": "grid-auto-rows: auto",
	}
}

// MathStyle represent the CSS style "math-style" with value "normal | compact"
// See https://developer.mozilla.org/docs/Web/CSS/math-style
type MathStyleStyle string

func (t MathStyleStyle) Name() string {
	return "math-style"
}

const MathStyleStyleNormal MathStyleStyle = "normal"

const MathStyleStyleCompact MathStyleStyle = "compact"

func (t MathStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t MathStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"math-style":         "math-style: normal",
		"math-style-normal":  "math-style: normal",
		"math-style-compact": "math-style: compact",
	}
}

// PaddingTop represent the CSS style "padding-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-top
type PaddingTopStyle string

func (t PaddingTopStyle) Name() string {
	return "padding-top"
}

func (t PaddingTopStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaddingTopStyle) Utilities() map[string]string {
	return map[string]string{
		"padding-top": "padding-top: 0",
	}
}

// BorderInlineStartWidth represent the CSS style "border-inline-start-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-start-width
type BorderInlineStartWidthStyle string

func (t BorderInlineStartWidthStyle) Name() string {
	return "border-inline-start-width"
}

func (t BorderInlineStartWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderInlineStartWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-inline-start-width": "border-inline-start-width: medium",
	}
}

// BoxDecorationBreak represent the CSS style "box-decoration-break" with value "slice | clone"
// See https://developer.mozilla.org/docs/Web/CSS/box-decoration-break
type BoxDecorationBreakStyle string

func (t BoxDecorationBreakStyle) Name() string {
	return "box-decoration-break"
}

const BoxDecorationBreakStyleSlice BoxDecorationBreakStyle = "slice"

const BoxDecorationBreakStyleClone BoxDecorationBreakStyle = "clone"

func (t BoxDecorationBreakStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxDecorationBreakStyle) Utilities() map[string]string {
	return map[string]string{
		"box-decoration-break":       "box-decoration-break: slice",
		"box-decoration-break-slice": "box-decoration-break: slice",
		"box-decoration-break-clone": "box-decoration-break: clone",
	}
}

// Inset represent the CSS style "inset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset
type InsetStyle string

func (t InsetStyle) Name() string {
	return "inset"
}

func (t InsetStyle) BrowserVariants() []string {
	return []string{}
}

func (t InsetStyle) Utilities() map[string]string {
	return map[string]string{
		"inset": "inset: auto",
	}
}

// AnimationFillMode represent the CSS style "animation-fill-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-fill-mode
type AnimationFillModeStyle string

func (t AnimationFillModeStyle) Name() string {
	return "animation-fill-mode"
}

func (t AnimationFillModeStyle) BrowserVariants() []string {
	return []string{}
}

func (t AnimationFillModeStyle) Utilities() map[string]string {
	return map[string]string{
		"animation-fill-mode": "animation-fill-mode: none",
	}
}

// Bottom represent the CSS style "bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/bottom
type BottomStyle string

func (t BottomStyle) Name() string {
	return "bottom"
}

func (t BottomStyle) BrowserVariants() []string {
	return []string{}
}

func (t BottomStyle) Utilities() map[string]string {
	return map[string]string{
		"bottom": "bottom: auto",
	}
}

// GridRowGap represent the CSS style "grid-row-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/row-gap
type GridRowGapStyle string

func (t GridRowGapStyle) Name() string {
	return "grid-row-gap"
}

func (t GridRowGapStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridRowGapStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-row-gap": "grid-row-gap: 0",
	}
}

// ScrollMarginInlineEnd represent the CSS style "scroll-margin-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-inline-end
type ScrollMarginInlineEndStyle float64

func (t ScrollMarginInlineEndStyle) Name() string {
	return "scroll-margin-inline-end"
}

func (t ScrollMarginInlineEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginInlineEndStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin-inline-end": "scroll-margin-inline-end: 0",
	}
}

// BorderEndStartRadius represent the CSS style "border-end-start-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-end-start-radius
type BorderEndStartRadiusStyle string

func (t BorderEndStartRadiusStyle) Name() string {
	return "border-end-start-radius"
}

func (t BorderEndStartRadiusStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderEndStartRadiusStyle) Utilities() map[string]string {
	return map[string]string{
		"border-end-start-radius": "border-end-start-radius: 0",
	}
}

// BorderInlineEndColor represent the CSS style "border-inline-end-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-end-color
type BorderInlineEndColorStyle Color

func (t BorderInlineEndColorStyle) Name() string {
	return "border-inline-end-color"
}

func (t BorderInlineEndColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderInlineEndColorStyle) Utilities() map[string]string {
	return map[string]string{
		"border-inline-end-color": "border-inline-end-color: currentcolor",
	}
}

// TransformRotate represent the CSS style "transform-rotate" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformRotateStyle string

func (t TransformRotateStyle) Name() string {
	return "transform-rotate"
}

func (t TransformRotateStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformRotateStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-rotate": "transform-rotate: rotate(0)",
	}
}

// FontVariantNumeric represent the CSS style "font-variant-numeric" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-numeric
type FontVariantNumericStyle string

func (t FontVariantNumericStyle) Name() string {
	return "font-variant-numeric"
}

func (t FontVariantNumericStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontVariantNumericStyle) Utilities() map[string]string {
	return map[string]string{
		"font-variant-numeric": "font-variant-numeric: normal",
	}
}

// MaskBorderOutset represent the CSS style "mask-border-outset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-outset
type MaskBorderOutsetStyle string

func (t MaskBorderOutsetStyle) Name() string {
	return "mask-border-outset"
}

func (t MaskBorderOutsetStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaskBorderOutsetStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-border-outset": "mask-border-outset: 0",
	}
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

func (t PageBreakBeforeStyle) BrowserVariants() []string {
	return []string{}
}

func (t PageBreakBeforeStyle) Utilities() map[string]string {
	return map[string]string{
		"page-break-before":        "page-break-before: auto",
		"page-break-before-auto":   "page-break-before: auto",
		"page-break-before-always": "page-break-before: always",
		"page-break-before-avoid":  "page-break-before: avoid",
		"page-break-before-left":   "page-break-before: left",
		"page-break-before-right":  "page-break-before: right",
		"page-break-before-recto":  "page-break-before: recto",
		"page-break-before-verso":  "page-break-before: verso",
	}
}

// ShapeMargin represent the CSS style "shape-margin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/shape-margin
type ShapeMarginStyle string

func (t ShapeMarginStyle) Name() string {
	return "shape-margin"
}

func (t ShapeMarginStyle) BrowserVariants() []string {
	return []string{}
}

func (t ShapeMarginStyle) Utilities() map[string]string {
	return map[string]string{
		"shape-margin": "shape-margin: 0",
	}
}

// AnimationDirection represent the CSS style "animation-direction" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-direction
type AnimationDirectionStyle string

func (t AnimationDirectionStyle) Name() string {
	return "animation-direction"
}

func (t AnimationDirectionStyle) BrowserVariants() []string {
	return []string{}
}

func (t AnimationDirectionStyle) Utilities() map[string]string {
	return map[string]string{
		"animation-direction": "animation-direction: normal",
	}
}

// ScrollPaddingInlineStart represent the CSS style "scroll-padding-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-inline-start
type ScrollPaddingInlineStartStyle string

func (t ScrollPaddingInlineStartStyle) Name() string {
	return "scroll-padding-inline-start"
}

func (t ScrollPaddingInlineStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingInlineStartStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding-inline-start": "scroll-padding-inline-start: auto",
	}
}

// AnimationDelay represent the CSS style "animation-delay" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-delay
type AnimationDelayStyle time.Duration

func (t AnimationDelayStyle) Name() string {
	return "animation-delay"
}

func (t AnimationDelayStyle) BrowserVariants() []string {
	return []string{}
}

func (t AnimationDelayStyle) Utilities() map[string]string {
	return map[string]string{
		"animation-delay": "animation-delay: 0s",
	}
}

// BorderBottomColor represent the CSS style "border-bottom-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-color
type BorderBottomColorStyle Color

func (t BorderBottomColorStyle) Name() string {
	return "border-bottom-color"
}

func (t BorderBottomColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBottomColorStyle) Utilities() map[string]string {
	return map[string]string{
		"border-bottom-color": "border-bottom-color: currentcolor",
	}
}

// ScrollPaddingLeft represent the CSS style "scroll-padding-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-left
type ScrollPaddingLeftStyle string

func (t ScrollPaddingLeftStyle) Name() string {
	return "scroll-padding-left"
}

func (t ScrollPaddingLeftStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollPaddingLeftStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-padding-left": "scroll-padding-left: auto",
	}
}

// GridAutoColumns represent the CSS style "grid-auto-columns" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-auto-columns
type GridAutoColumnsStyle string

func (t GridAutoColumnsStyle) Name() string {
	return "grid-auto-columns"
}

func (t GridAutoColumnsStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridAutoColumnsStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-auto-columns": "grid-auto-columns: auto",
	}
}

// BorderInlineColor represent the CSS style "border-inline-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-color
type BorderInlineColorStyle Color

func (t BorderInlineColorStyle) Name() string {
	return "border-inline-color"
}

func (t BorderInlineColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderInlineColorStyle) Utilities() map[string]string {
	return map[string]string{
		"border-inline-color": "border-inline-color: currentcolor",
	}
}

// BorderInlineEndStyle represent the CSS style "border-inline-end-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-end-style
type BorderInlineEndStyleStyle string

func (t BorderInlineEndStyleStyle) Name() string {
	return "border-inline-end-style"
}

func (t BorderInlineEndStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderInlineEndStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"border-inline-end-style": "border-inline-end-style: none",
	}
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

func (t TextDecorationStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextDecorationStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"text-decoration-style":        "text-decoration-style: solid",
		"text-decoration-style-solid":  "text-decoration-style: solid",
		"text-decoration-style-double": "text-decoration-style: double",
		"text-decoration-style-dotted": "text-decoration-style: dotted",
		"text-decoration-style-dashed": "text-decoration-style: dashed",
		"text-decoration-style-wavy":   "text-decoration-style: wavy",
	}
}

// BorderEndEndRadius represent the CSS style "border-end-end-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-end-end-radius
type BorderEndEndRadiusStyle string

func (t BorderEndEndRadiusStyle) Name() string {
	return "border-end-end-radius"
}

func (t BorderEndEndRadiusStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderEndEndRadiusStyle) Utilities() map[string]string {
	return map[string]string{
		"border-end-end-radius": "border-end-end-radius: 0",
	}
}

// BorderBottomRightRadius represent the CSS style "border-bottom-right-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-right-radius
type BorderBottomRightRadiusStyle string

func (t BorderBottomRightRadiusStyle) Name() string {
	return "border-bottom-right-radius"
}

func (t BorderBottomRightRadiusStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBottomRightRadiusStyle) Utilities() map[string]string {
	return map[string]string{
		"border-bottom-right-radius": "border-bottom-right-radius: 0",
	}
}

// TextSizeAdjust represent the CSS style "text-size-adjust" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-size-adjust
type TextSizeAdjustStyle string

func (t TextSizeAdjustStyle) Name() string {
	return "text-size-adjust"
}

func (t TextSizeAdjustStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextSizeAdjustStyle) Utilities() map[string]string {
	return map[string]string{
		"text-size-adjust": "text-size-adjust: autoForSmartphoneBrowsersSupportingInflation",
	}
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

func (t ColorInterpolationStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColorInterpolationStyle) Utilities() map[string]string {
	return map[string]string{
		"color-interpolation":           "color-interpolation: sRGB",
		"color-interpolation-auto":      "color-interpolation: auto",
		"color-interpolation-sRGB":      "color-interpolation: sRGB",
		"color-interpolation-linearRGB": "color-interpolation: linearRGB",
	}
}

// AnimationDuration represent the CSS style "animation-duration" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-duration
type AnimationDurationStyle time.Duration

func (t AnimationDurationStyle) Name() string {
	return "animation-duration"
}

func (t AnimationDurationStyle) BrowserVariants() []string {
	return []string{}
}

func (t AnimationDurationStyle) Utilities() map[string]string {
	return map[string]string{
		"animation-duration": "animation-duration: 0s",
	}
}

// FontVariantLigatures represent the CSS style "font-variant-ligatures" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-ligatures
type FontVariantLigaturesStyle string

func (t FontVariantLigaturesStyle) Name() string {
	return "font-variant-ligatures"
}

func (t FontVariantLigaturesStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontVariantLigaturesStyle) Utilities() map[string]string {
	return map[string]string{
		"font-variant-ligatures": "font-variant-ligatures: normal",
	}
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

func (t ScrollSnapTypeStyle) BrowserVariants() []string {
	return []string{
		"-ms-scroll-snap-type",
	}
}

func (t ScrollSnapTypeStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-snap-type":           "scroll-snap-type: none",
		"scroll-snap-type-none":      "scroll-snap-type: none",
		"scroll-snap-type-block":     "scroll-snap-type: block",
		"scroll-snap-type-inline":    "scroll-snap-type: inline",
		"scroll-snap-type-both":      "scroll-snap-type: both",
		"scroll-snap-type-mandatory": "scroll-snap-type: mandatory",
		"scroll-snap-type-proximity": "scroll-snap-type: proximity",
	}
}

// AnimationPlayState represent the CSS style "animation-play-state" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-play-state
type AnimationPlayStateStyle string

func (t AnimationPlayStateStyle) Name() string {
	return "animation-play-state"
}

func (t AnimationPlayStateStyle) BrowserVariants() []string {
	return []string{}
}

func (t AnimationPlayStateStyle) Utilities() map[string]string {
	return map[string]string{
		"animation-play-state": "animation-play-state: running",
	}
}

// TransformRotate3d represent the CSS style "transform-rotate-3d" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformRotate3dStyle string

func (t TransformRotate3dStyle) Name() string {
	return "transform-rotate-3d"
}

func (t TransformRotate3dStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformRotate3dStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-rotate-3d": "transform-rotate-3d: rotate3d(0, 0, 0)",
	}
}

// TransformRotateY represent the CSS style "transform-rotate-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformRotateYStyle string

func (t TransformRotateYStyle) Name() string {
	return "transform-rotate-y"
}

func (t TransformRotateYStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformRotateYStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-rotate-y": "transform-rotate-y: rotateY(0)",
	}
}

// MaskOrigin represent the CSS style "mask-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-origin
type MaskOriginStyle string

func (t MaskOriginStyle) Name() string {
	return "mask-origin"
}

func (t MaskOriginStyle) BrowserVariants() []string {
	return []string{
		"-webkit-mask-origin",
	}
}

func (t MaskOriginStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-origin": "mask-origin: border-box",
	}
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

func (t BackgroundRepeatStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundRepeatStyle) Utilities() map[string]string {
	return map[string]string{
		"background-repeat":           "background-repeat: repeat",
		"background-repeat-none":      "background-repeat: none",
		"background-repeat-repeat":    "background-repeat: repeat",
		"background-repeat-no-repeat": "background-repeat: no-repeat",
	}
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

func (t FontWeightStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontWeightStyle) Utilities() map[string]string {
	return map[string]string{
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

func (t JustifyTracksStyle) BrowserVariants() []string {
	return []string{}
}

func (t JustifyTracksStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// GlyphOrientationVertical represent the CSS style "glyph-orientation-vertical" with value ""
// See
type GlyphOrientationVerticalStyle string

func (t GlyphOrientationVerticalStyle) Name() string {
	return "glyph-orientation-vertical"
}

func (t GlyphOrientationVerticalStyle) BrowserVariants() []string {
	return []string{}
}

func (t GlyphOrientationVerticalStyle) Utilities() map[string]string {
	return map[string]string{
		"glyph-orientation-vertical": "glyph-orientation-vertical: auto",
	}
}

// BackgroundSize represent the CSS style "background-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-size
type BackgroundSizeStyle string

func (t BackgroundSizeStyle) Name() string {
	return "background-size"
}

func (t BackgroundSizeStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundSizeStyle) Utilities() map[string]string {
	return map[string]string{
		"background-size": "background-size: auto auto",
	}
}

// StopOpacity represent the CSS style "stop-opacity" with value ""
// See
type StopOpacityStyle float64

func (t StopOpacityStyle) Name() string {
	return "stop-opacity"
}

func (t StopOpacityStyle) BrowserVariants() []string {
	return []string{}
}

func (t StopOpacityStyle) Utilities() map[string]string {
	return map[string]string{
		"stop-opacity": "stop-opacity: 1",
	}
}

// BackgroundBlendMode represent the CSS style "background-blend-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-blend-mode
type BackgroundBlendModeStyle string

func (t BackgroundBlendModeStyle) Name() string {
	return "background-blend-mode"
}

func (t BackgroundBlendModeStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundBlendModeStyle) Utilities() map[string]string {
	return map[string]string{
		"background-blend-mode": "background-blend-mode: normal",
	}
}

// TransformScale3d represent the CSS style "transform-scale-3d" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformScale3dStyle string

func (t TransformScale3dStyle) Name() string {
	return "transform-scale-3d"
}

func (t TransformScale3dStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformScale3dStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-scale-3d": "transform-scale-3d: scale3d(1, 1, 1)",
	}
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

func (t WritingModeStyle) BrowserVariants() []string {
	return []string{}
}

func (t WritingModeStyle) Utilities() map[string]string {
	return map[string]string{
		"writing-mode":               "writing-mode: horizontal-tb",
		"writing-mode-horizontal-tb": "writing-mode: horizontal-tb",
		"writing-mode-vertical-rl":   "writing-mode: vertical-rl",
		"writing-mode-vertical-lr":   "writing-mode: vertical-lr",
		"writing-mode-sideways-rl":   "writing-mode: sideways-rl",
		"writing-mode-sideways-lr":   "writing-mode: sideways-lr",
	}
}

// MarginBlockEnd represent the CSS style "margin-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-block-end
type MarginBlockEndStyle string

func (t MarginBlockEndStyle) Name() string {
	return "margin-block-end"
}

func (t MarginBlockEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginBlockEndStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-block-end": "margin-block-end: 0",
	}
}

// Order represent the CSS style "order" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/order
type OrderStyle float64

func (t OrderStyle) Name() string {
	return "order"
}

func (t OrderStyle) BrowserVariants() []string {
	return []string{}
}

func (t OrderStyle) Utilities() map[string]string {
	return map[string]string{
		"order": "order: 0",
	}
}

// PaddingBlock represent the CSS style "padding-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-block
type PaddingBlockStyle string

func (t PaddingBlockStyle) Name() string {
	return "padding-block"
}

func (t PaddingBlockStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaddingBlockStyle) Utilities() map[string]string {
	return map[string]string{
		"padding-block": "padding-block: 0",
	}
}

// AspectRatio represent the CSS style "aspect-ratio" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/aspect-ratio
type AspectRatioStyle string

func (t AspectRatioStyle) Name() string {
	return "aspect-ratio"
}

func (t AspectRatioStyle) BrowserVariants() []string {
	return []string{}
}

func (t AspectRatioStyle) Utilities() map[string]string {
	return map[string]string{
		"aspect-ratio": "aspect-ratio: auto",
	}
}

// InsetBlockStart represent the CSS style "inset-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-block-start
type InsetBlockStartStyle string

func (t InsetBlockStartStyle) Name() string {
	return "inset-block-start"
}

func (t InsetBlockStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t InsetBlockStartStyle) Utilities() map[string]string {
	return map[string]string{
		"inset-block-start": "inset-block-start: auto",
	}
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

func (t PaintOrderStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaintOrderStyle) Utilities() map[string]string {
	return map[string]string{
		"paint-order":         "paint-order: normal",
		"paint-order-normal":  "paint-order: normal",
		"paint-order-fill":    "paint-order: fill",
		"paint-order-stroke":  "paint-order: stroke",
		"paint-order-markers": "paint-order: markers",
	}
}

// ScrollMarginTop represent the CSS style "scroll-margin-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-top
type ScrollMarginTopStyle float64

func (t ScrollMarginTopStyle) Name() string {
	return "scroll-margin-top"
}

func (t ScrollMarginTopStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginTopStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin-top": "scroll-margin-top: 0",
	}
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

func (t ScrollSnapTypeYStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollSnapTypeYStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-snap-type-y":           "scroll-snap-type-y: none",
		"scroll-snap-type-y-none":      "scroll-snap-type-y: none",
		"scroll-snap-type-y-mandatory": "scroll-snap-type-y: mandatory",
		"scroll-snap-type-y-proximity": "scroll-snap-type-y: proximity",
	}
}

// BackgroundColor represent the CSS style "background-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-color
type BackgroundColorStyle Color

func (t BackgroundColorStyle) Name() string {
	return "background-color"
}

func (t BackgroundColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundColorStyle) Utilities() map[string]string {
	return map[string]string{
		"background-color": "background-color: transparent",
	}
}

// AnimationIterationCount represent the CSS style "animation-iteration-count" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-iteration-count
type AnimationIterationCountStyle float64

func (t AnimationIterationCountStyle) Name() string {
	return "animation-iteration-count"
}

func (t AnimationIterationCountStyle) BrowserVariants() []string {
	return []string{}
}

func (t AnimationIterationCountStyle) Utilities() map[string]string {
	return map[string]string{
		"animation-iteration-count": "animation-iteration-count: 1",
	}
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

func (t FlexDirectionStyle) BrowserVariants() []string {
	return []string{}
}

func (t FlexDirectionStyle) Utilities() map[string]string {
	return map[string]string{
		"flex-direction":                "flex-direction: row",
		"flex-direction-row":            "flex-direction: row",
		"flex-direction-row-reverse":    "flex-direction: row-reverse",
		"flex-direction-column":         "flex-direction: column",
		"flex-direction-column-reverse": "flex-direction: column-reverse",
	}
}

// GridColumnStart represent the CSS style "grid-column-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-column-start
type GridColumnStartStyle string

func (t GridColumnStartStyle) Name() string {
	return "grid-column-start"
}

func (t GridColumnStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridColumnStartStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-column-start": "grid-column-start: auto",
	}
}

// MaxWidth represent the CSS style "max-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-width
type MaxWidthStyle string

func (t MaxWidthStyle) Name() string {
	return "max-width"
}

func (t MaxWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaxWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"max-width": "max-width: none",
	}
}

// WillChange represent the CSS style "will-change" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/will-change
type WillChangeStyle string

func (t WillChangeStyle) Name() string {
	return "will-change"
}

func (t WillChangeStyle) BrowserVariants() []string {
	return []string{}
}

func (t WillChangeStyle) Utilities() map[string]string {
	return map[string]string{
		"will-change": "will-change: auto",
	}
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

func (t ContainStyle) BrowserVariants() []string {
	return []string{}
}

func (t ContainStyle) Utilities() map[string]string {
	return map[string]string{
		"contain":              "contain: none",
		"contain-none":         "contain: none",
		"contain-strict":       "contain: strict",
		"contain-content":      "contain: content",
		"contain-size--layout": "contain: size--layout",
		"contain-style":        "contain: style",
		"contain-paint":        "contain: paint",
	}
}

// CounterIncrement represent the CSS style "counter-increment" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/counter-increment
type CounterIncrementStyle string

func (t CounterIncrementStyle) Name() string {
	return "counter-increment"
}

func (t CounterIncrementStyle) BrowserVariants() []string {
	return []string{}
}

func (t CounterIncrementStyle) Utilities() map[string]string {
	return map[string]string{
		"counter-increment": "counter-increment: none",
	}
}

// ColumnGap represent the CSS style "column-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-gap
type ColumnGapStyle string

func (t ColumnGapStyle) Name() string {
	return "column-gap"
}

func (t ColumnGapStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColumnGapStyle) Utilities() map[string]string {
	return map[string]string{
		"column-gap": "column-gap: normal",
	}
}

// MaskType represent the CSS style "mask-type" with value "luminance | alpha"
// See https://developer.mozilla.org/docs/Web/CSS/mask-type
type MaskTypeStyle string

func (t MaskTypeStyle) Name() string {
	return "mask-type"
}

const MaskTypeStyleLuminance MaskTypeStyle = "luminance"

const MaskTypeStyleAlpha MaskTypeStyle = "alpha"

func (t MaskTypeStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaskTypeStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-type":           "mask-type: luminance",
		"mask-type-luminance": "mask-type: luminance",
		"mask-type-alpha":     "mask-type: alpha",
	}
}

// StrokeWidth represent the CSS style "stroke-width" with value ""
// See
type StrokeWidthStyle float64

func (t StrokeWidthStyle) Name() string {
	return "stroke-width"
}

func (t StrokeWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t StrokeWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"stroke-width": "stroke-width: 1",
	}
}

// Top represent the CSS style "top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/top
type TopStyle string

func (t TopStyle) Name() string {
	return "top"
}

func (t TopStyle) BrowserVariants() []string {
	return []string{}
}

func (t TopStyle) Utilities() map[string]string {
	return map[string]string{
		"top": "top: auto",
	}
}

// GridRowStart represent the CSS style "grid-row-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-row-start
type GridRowStartStyle string

func (t GridRowStartStyle) Name() string {
	return "grid-row-start"
}

func (t GridRowStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridRowStartStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-row-start": "grid-row-start: auto",
	}
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

func (t ObjectFitStyle) BrowserVariants() []string {
	return []string{}
}

func (t ObjectFitStyle) Utilities() map[string]string {
	return map[string]string{
		"object-fit":            "object-fit: fill",
		"object-fit-fill":       "object-fit: fill",
		"object-fit-contain":    "object-fit: contain",
		"object-fit-cover":      "object-fit: cover",
		"object-fit-none":       "object-fit: none",
		"object-fit-scale-down": "object-fit: scale-down",
	}
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

func (t JustifyContentStyle) BrowserVariants() []string {
	return []string{}
}

func (t JustifyContentStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// TextUnderlineOffset represent the CSS style "text-underline-offset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-underline-offset
type TextUnderlineOffsetStyle string

func (t TextUnderlineOffsetStyle) Name() string {
	return "text-underline-offset"
}

func (t TextUnderlineOffsetStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextUnderlineOffsetStyle) Utilities() map[string]string {
	return map[string]string{
		"text-underline-offset": "text-underline-offset: auto",
	}
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

func (t FlexWrapStyle) BrowserVariants() []string {
	return []string{}
}

func (t FlexWrapStyle) Utilities() map[string]string {
	return map[string]string{
		"flex-wrap":              "flex-wrap: nowrap",
		"flex-wrap-nowrap":       "flex-wrap: nowrap",
		"flex-wrap-wrap":         "flex-wrap: wrap",
		"flex-wrap-wrap-reverse": "flex-wrap: wrap-reverse",
	}
}

// ScrollMarginBottom represent the CSS style "scroll-margin-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-bottom
type ScrollMarginBottomStyle float64

func (t ScrollMarginBottomStyle) Name() string {
	return "scroll-margin-bottom"
}

func (t ScrollMarginBottomStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginBottomStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin-bottom": "scroll-margin-bottom: 0",
	}
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

func (t InitialLetterAlignStyle) BrowserVariants() []string {
	return []string{}
}

func (t InitialLetterAlignStyle) Utilities() map[string]string {
	return map[string]string{
		"initial-letter-align":             "initial-letter-align: auto",
		"initial-letter-align-auto":        "initial-letter-align: auto",
		"initial-letter-align-alphabetic":  "initial-letter-align: alphabetic",
		"initial-letter-align-hanging":     "initial-letter-align: hanging",
		"initial-letter-align-ideographic": "initial-letter-align: ideographic",
	}
}

// BorderRightColor represent the CSS style "border-right-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-right-color
type BorderRightColorStyle Color

func (t BorderRightColorStyle) Name() string {
	return "border-right-color"
}

func (t BorderRightColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderRightColorStyle) Utilities() map[string]string {
	return map[string]string{
		"border-right-color": "border-right-color: currentcolor",
	}
}

// MaskBorderMode represent the CSS style "mask-border-mode" with value "luminance | alpha"
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-mode
type MaskBorderModeStyle string

func (t MaskBorderModeStyle) Name() string {
	return "mask-border-mode"
}

const MaskBorderModeStyleLuminance MaskBorderModeStyle = "luminance"

const MaskBorderModeStyleAlpha MaskBorderModeStyle = "alpha"

func (t MaskBorderModeStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaskBorderModeStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-border-mode":           "mask-border-mode: alpha",
		"mask-border-mode-luminance": "mask-border-mode: luminance",
		"mask-border-mode-alpha":     "mask-border-mode: alpha",
	}
}

// MaskBorderWidth represent the CSS style "mask-border-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-width
type MaskBorderWidthStyle string

func (t MaskBorderWidthStyle) Name() string {
	return "mask-border-width"
}

func (t MaskBorderWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaskBorderWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-border-width": "mask-border-width: auto",
	}
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

func (t RubyPositionStyle) BrowserVariants() []string {
	return []string{}
}

func (t RubyPositionStyle) Utilities() map[string]string {
	return map[string]string{
		"ruby-position":                 "ruby-position: over",
		"ruby-position-over":            "ruby-position: over",
		"ruby-position-under":           "ruby-position: under",
		"ruby-position-inter-character": "ruby-position: inter-character",
	}
}

// AnimationName represent the CSS style "animation-name" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-name
type AnimationNameStyle string

func (t AnimationNameStyle) Name() string {
	return "animation-name"
}

func (t AnimationNameStyle) BrowserVariants() []string {
	return []string{}
}

func (t AnimationNameStyle) Utilities() map[string]string {
	return map[string]string{
		"animation-name": "animation-name: none",
	}
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

func (t UnicodeBidiStyle) BrowserVariants() []string {
	return []string{}
}

func (t UnicodeBidiStyle) Utilities() map[string]string {
	return map[string]string{
		"unicode-bidi":                  "unicode-bidi: normal",
		"unicode-bidi-normal":           "unicode-bidi: normal",
		"unicode-bidi-embed":            "unicode-bidi: embed",
		"unicode-bidi-isolate":          "unicode-bidi: isolate",
		"unicode-bidi-bidi-override":    "unicode-bidi: bidi-override",
		"unicode-bidi-isolate-override": "unicode-bidi: isolate-override",
		"unicode-bidi-plaintext":        "unicode-bidi: plaintext",
	}
}

// TextCombineUpright represent the CSS style "text-combine-upright" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-combine-upright
type TextCombineUprightStyle string

func (t TextCombineUprightStyle) Name() string {
	return "text-combine-upright"
}

func (t TextCombineUprightStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextCombineUprightStyle) Utilities() map[string]string {
	return map[string]string{
		"text-combine-upright": "text-combine-upright: none",
	}
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

func (t BreakInsideStyle) BrowserVariants() []string {
	return []string{}
}

func (t BreakInsideStyle) Utilities() map[string]string {
	return map[string]string{
		"break-inside":              "break-inside: auto",
		"break-inside-auto":         "break-inside: auto",
		"break-inside-avoid":        "break-inside: avoid",
		"break-inside-avoid-page":   "break-inside: avoid-page",
		"break-inside-avoid-column": "break-inside: avoid-column",
		"break-inside-avoid-region": "break-inside: avoid-region",
	}
}

// TransformStyle represent the CSS style "transform-style" with value "flat | preserve-3d"
// See https://developer.mozilla.org/docs/Web/CSS/transform-style
type TransformStyleStyle string

func (t TransformStyleStyle) Name() string {
	return "transform-style"
}

const TransformStyleStyleFlat TransformStyleStyle = "flat"

const TransformStyleStylePreserve3d TransformStyleStyle = "preserve-3d"

func (t TransformStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-style":             "transform-style: flat",
		"transform-style-flat":        "transform-style: flat",
		"transform-style-preserve-3d": "transform-style: preserve-3d",
	}
}

// BorderBlockStartColor represent the CSS style "border-block-start-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-start-color
type BorderBlockStartColorStyle Color

func (t BorderBlockStartColorStyle) Name() string {
	return "border-block-start-color"
}

func (t BorderBlockStartColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBlockStartColorStyle) Utilities() map[string]string {
	return map[string]string{
		"border-block-start-color": "border-block-start-color: currentcolor",
	}
}

// BorderWidth represent the CSS style "border-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-width
type BorderWidthStyle string

func (t BorderWidthStyle) Name() string {
	return "border-width"
}

func (t BorderWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-width": "border-width: auto",
	}
}

// FontVariationSettings represent the CSS style "font-variation-settings" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variation-settings
type FontVariationSettingsStyle string

func (t FontVariationSettingsStyle) Name() string {
	return "font-variation-settings"
}

func (t FontVariationSettingsStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontVariationSettingsStyle) Utilities() map[string]string {
	return map[string]string{
		"font-variation-settings": "font-variation-settings: normal",
	}
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

func (t TextRenderingStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextRenderingStyle) Utilities() map[string]string {
	return map[string]string{
		"text-rendering":                    "text-rendering: auto",
		"text-rendering-auto":               "text-rendering: auto",
		"text-rendering-optimizeSpeed":      "text-rendering: optimizeSpeed",
		"text-rendering-optimizeLegibility": "text-rendering: optimizeLegibility",
		"text-rendering-geometricPrecision": "text-rendering: geometricPrecision",
	}
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

func (t WordBreakStyle) BrowserVariants() []string {
	return []string{}
}

func (t WordBreakStyle) Utilities() map[string]string {
	return map[string]string{
		"word-break":            "word-break: normal",
		"word-break-normal":     "word-break: normal",
		"word-break-break-all":  "word-break: break-all",
		"word-break-keep-all":   "word-break: keep-all",
		"word-break-break-word": "word-break: break-word",
	}
}

// Rotate represent the CSS style "rotate" with value "angle"
// See https://developer.mozilla.org/docs/Web/CSS/rotate
type RotateStyle string

func (t RotateStyle) Name() string {
	return "rotate"
}

const RotateStyleAngle RotateStyle = "angle"

func (t RotateStyle) BrowserVariants() []string {
	return []string{}
}

func (t RotateStyle) Utilities() map[string]string {
	return map[string]string{
		"rotate":       "rotate: none",
		"rotate-angle": "rotate: angle",
	}
}

// BlockSize represent the CSS style "block-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/block-size
type BlockSizeStyle string

func (t BlockSizeStyle) Name() string {
	return "block-size"
}

func (t BlockSizeStyle) BrowserVariants() []string {
	return []string{}
}

func (t BlockSizeStyle) Utilities() map[string]string {
	return map[string]string{
		"block-size": "block-size: auto",
	}
}

// ScrollMargin represent the CSS style "scroll-margin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin
type ScrollMarginStyle float64

func (t ScrollMarginStyle) Name() string {
	return "scroll-margin"
}

func (t ScrollMarginStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin": "scroll-margin: 0",
	}
}

// TextIndent represent the CSS style "text-indent" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-indent
type TextIndentStyle string

func (t TextIndentStyle) Name() string {
	return "text-indent"
}

func (t TextIndentStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextIndentStyle) Utilities() map[string]string {
	return map[string]string{
		"text-indent": "text-indent: 0",
	}
}

// MarkerMid represent the CSS style "marker-mid" with value ""
// See
type MarkerMidStyle string

func (t MarkerMidStyle) Name() string {
	return "marker-mid"
}

func (t MarkerMidStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarkerMidStyle) Utilities() map[string]string {
	return map[string]string{
		"marker-mid": "marker-mid: none",
	}
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

func (t UserSelectStyle) BrowserVariants() []string {
	return []string{
		"-ms-user-select",
	}
}

func (t UserSelectStyle) Utilities() map[string]string {
	return map[string]string{
		"user-select":         "user-select: auto",
		"user-select-auto":    "user-select: auto",
		"user-select-text":    "user-select: text",
		"user-select-none":    "user-select: none",
		"user-select-contain": "user-select: contain",
		"user-select-all":     "user-select: all",
	}
}

// FontOpticalSizing represent the CSS style "font-optical-sizing" with value "auto | none"
// See https://developer.mozilla.org/docs/Web/CSS/font-optical-sizing
type FontOpticalSizingStyle string

func (t FontOpticalSizingStyle) Name() string {
	return "font-optical-sizing"
}

const FontOpticalSizingStyleAuto FontOpticalSizingStyle = "auto"

const FontOpticalSizingStyleNone FontOpticalSizingStyle = "none"

func (t FontOpticalSizingStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontOpticalSizingStyle) Utilities() map[string]string {
	return map[string]string{
		"font-optical-sizing":      "font-optical-sizing: auto",
		"font-optical-sizing-auto": "font-optical-sizing: auto",
		"font-optical-sizing-none": "font-optical-sizing: none",
	}
}

// MarginRight represent the CSS style "margin-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-right
type MarginRightStyle string

func (t MarginRightStyle) Name() string {
	return "margin-right"
}

func (t MarginRightStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginRightStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-right": "margin-right: 0",
	}
}

// PaddingInlineEnd represent the CSS style "padding-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-inline-end
type PaddingInlineEndStyle string

func (t PaddingInlineEndStyle) Name() string {
	return "padding-inline-end"
}

func (t PaddingInlineEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaddingInlineEndStyle) Utilities() map[string]string {
	return map[string]string{
		"padding-inline-end": "padding-inline-end: 0",
	}
}

// Fill represent the CSS style "fill" with value ""
// See
type FillStyle Color

func (t FillStyle) Name() string {
	return "fill"
}

func (t FillStyle) BrowserVariants() []string {
	return []string{}
}

func (t FillStyle) Utilities() map[string]string {
	return map[string]string{
		"fill": "fill: black",
	}
}

// ColumnRuleColor represent the CSS style "column-rule-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-rule-color
type ColumnRuleColorStyle Color

func (t ColumnRuleColorStyle) Name() string {
	return "column-rule-color"
}

func (t ColumnRuleColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColumnRuleColorStyle) Utilities() map[string]string {
	return map[string]string{
		"column-rule-color": "column-rule-color: currentcolor",
	}
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

func (t BackgroundPositionYStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundPositionYStyle) Utilities() map[string]string {
	return map[string]string{
		"background-position-y":         "background-position-y: top",
		"background-position-y-center":  "background-position-y: center",
		"background-position-y-top":     "background-position-y: top",
		"background-position-y-bottom":  "background-position-y: bottom",
		"background-position-y-y-start": "background-position-y: y-start",
		"background-position-y-y-end":   "background-position-y: y-end",
	}
}

// BorderBlockStartWidth represent the CSS style "border-block-start-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-start-width
type BorderBlockStartWidthStyle string

func (t BorderBlockStartWidthStyle) Name() string {
	return "border-block-start-width"
}

func (t BorderBlockStartWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBlockStartWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-block-start-width": "border-block-start-width: medium",
	}
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

func (t BoxDirectionStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxDirectionStyle) Utilities() map[string]string {
	return map[string]string{
		"box-direction":         "box-direction: normal",
		"box-direction-normal":  "box-direction: normal",
		"box-direction-reverse": "box-direction: reverse",
		"box-direction-inherit": "box-direction: inherit",
	}
}

// CaretColor represent the CSS style "caret-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/caret-color
type CaretColorStyle Color

func (t CaretColorStyle) Name() string {
	return "caret-color"
}

func (t CaretColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t CaretColorStyle) Utilities() map[string]string {
	return map[string]string{
		"caret-color": "caret-color: auto",
	}
}

// FlexGrow represent the CSS style "flex-grow" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/flex-grow
type FlexGrowStyle float64

func (t FlexGrowStyle) Name() string {
	return "flex-grow"
}

func (t FlexGrowStyle) BrowserVariants() []string {
	return []string{}
}

func (t FlexGrowStyle) Utilities() map[string]string {
	return map[string]string{
		"flex-grow": "flex-grow: 0",
	}
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

func (t OverflowWrapStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverflowWrapStyle) Utilities() map[string]string {
	return map[string]string{
		"overflow-wrap":            "overflow-wrap: normal",
		"overflow-wrap-normal":     "overflow-wrap: normal",
		"overflow-wrap-break-word": "overflow-wrap: break-word",
		"overflow-wrap-anywhere":   "overflow-wrap: anywhere",
	}
}

// ScrollSnapStop represent the CSS style "scroll-snap-stop" with value "normal | always"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-stop
type ScrollSnapStopStyle string

func (t ScrollSnapStopStyle) Name() string {
	return "scroll-snap-stop"
}

const ScrollSnapStopStyleNormal ScrollSnapStopStyle = "normal"

const ScrollSnapStopStyleAlways ScrollSnapStopStyle = "always"

func (t ScrollSnapStopStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollSnapStopStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-snap-stop":        "scroll-snap-stop: normal",
		"scroll-snap-stop-normal": "scroll-snap-stop: normal",
		"scroll-snap-stop-always": "scroll-snap-stop: always",
	}
}

// WordWrap represent the CSS style "word-wrap" with value "normal | break-word"
// See https://developer.mozilla.org/docs/Web/CSS/overflow-wrap
type WordWrapStyle string

func (t WordWrapStyle) Name() string {
	return "word-wrap"
}

const WordWrapStyleNormal WordWrapStyle = "normal"

const WordWrapStyleBreakWord WordWrapStyle = "break-word"

func (t WordWrapStyle) BrowserVariants() []string {
	return []string{}
}

func (t WordWrapStyle) Utilities() map[string]string {
	return map[string]string{
		"word-wrap":            "word-wrap: normal",
		"word-wrap-normal":     "word-wrap: normal",
		"word-wrap-break-word": "word-wrap: break-word",
	}
}

// BoxOrdinalGroup represent the CSS style "box-ordinal-group" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/box-ordinal-group
type BoxOrdinalGroupStyle float64

func (t BoxOrdinalGroupStyle) Name() string {
	return "box-ordinal-group"
}

func (t BoxOrdinalGroupStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxOrdinalGroupStyle) Utilities() map[string]string {
	return map[string]string{
		"box-ordinal-group": "box-ordinal-group: 1",
	}
}

// CounterSet represent the CSS style "counter-set" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/counter-set
type CounterSetStyle string

func (t CounterSetStyle) Name() string {
	return "counter-set"
}

func (t CounterSetStyle) BrowserVariants() []string {
	return []string{}
}

func (t CounterSetStyle) Utilities() map[string]string {
	return map[string]string{
		"counter-set": "counter-set: none",
	}
}

// MaskBorderSlice represent the CSS style "mask-border-slice" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-slice
type MaskBorderSliceStyle string

func (t MaskBorderSliceStyle) Name() string {
	return "mask-border-slice"
}

func (t MaskBorderSliceStyle) BrowserVariants() []string {
	return []string{}
}

func (t MaskBorderSliceStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-border-slice": "mask-border-slice: 0",
	}
}

// TransformTranslateY represent the CSS style "transform-translate-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformTranslateYStyle string

func (t TransformTranslateYStyle) Name() string {
	return "transform-translate-y"
}

func (t TransformTranslateYStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformTranslateYStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-translate-y": "transform-translate-y: translateY(0)",
	}
}

// MixBlendMode represent the CSS style "mix-blend-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mix-blend-mode
type MixBlendModeStyle string

func (t MixBlendModeStyle) Name() string {
	return "mix-blend-mode"
}

func (t MixBlendModeStyle) BrowserVariants() []string {
	return []string{}
}

func (t MixBlendModeStyle) Utilities() map[string]string {
	return map[string]string{
		"mix-blend-mode": "mix-blend-mode: normal",
	}
}

// Zoom represent the CSS style "zoom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/zoom
type ZoomStyle string

func (t ZoomStyle) Name() string {
	return "zoom"
}

func (t ZoomStyle) BrowserVariants() []string {
	return []string{}
}

func (t ZoomStyle) Utilities() map[string]string {
	return map[string]string{
		"zoom": "zoom: normal",
	}
}

// ColumnRuleStyle represent the CSS style "column-rule-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-rule-style
type ColumnRuleStyleStyle string

func (t ColumnRuleStyleStyle) Name() string {
	return "column-rule-style"
}

func (t ColumnRuleStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColumnRuleStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"column-rule-style": "column-rule-style: none",
	}
}

// MarginBlock represent the CSS style "margin-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-block
type MarginBlockStyle string

func (t MarginBlockStyle) Name() string {
	return "margin-block"
}

func (t MarginBlockStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginBlockStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-block": "margin-block: 0",
	}
}

// BorderRightStyle represent the CSS style "border-right-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-right-style
type BorderRightStyleStyle string

func (t BorderRightStyleStyle) Name() string {
	return "border-right-style"
}

func (t BorderRightStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderRightStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"border-right-style": "border-right-style: none",
	}
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

func (t OverscrollBehaviorYStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverscrollBehaviorYStyle) Utilities() map[string]string {
	return map[string]string{
		"overscroll-behavior-y":         "overscroll-behavior-y: auto",
		"overscroll-behavior-y-contain": "overscroll-behavior-y: contain",
		"overscroll-behavior-y-none":    "overscroll-behavior-y: none",
		"overscroll-behavior-y-auto":    "overscroll-behavior-y: auto",
	}
}

// BorderBlockStartStyle represent the CSS style "border-block-start-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-start-style
type BorderBlockStartStyleStyle string

func (t BorderBlockStartStyleStyle) Name() string {
	return "border-block-start-style"
}

func (t BorderBlockStartStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBlockStartStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"border-block-start-style": "border-block-start-style: none",
	}
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

func (t GridAutoFlowStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridAutoFlowStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-auto-flow":        "grid-auto-flow: row",
		"grid-auto-flow-row":    "grid-auto-flow: row",
		"grid-auto-flow-column": "grid-auto-flow: column",
		"grid-auto-flow-dense":  "grid-auto-flow: dense",
	}
}

// LineClamp represent the CSS style "line-clamp" with value ""
// See
type LineClampStyle string

func (t LineClampStyle) Name() string {
	return "line-clamp"
}

func (t LineClampStyle) BrowserVariants() []string {
	return []string{
		"-webkit-line-clamp",
	}
}

func (t LineClampStyle) Utilities() map[string]string {
	return map[string]string{
		"line-clamp": "line-clamp: none",
	}
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

func (t OverflowBlockStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverflowBlockStyle) Utilities() map[string]string {
	return map[string]string{
		"overflow-block":         "overflow-block: auto",
		"overflow-block-visible": "overflow-block: visible",
		"overflow-block-hidden":  "overflow-block: hidden",
		"overflow-block-clip":    "overflow-block: clip",
		"overflow-block-scroll":  "overflow-block: scroll",
		"overflow-block-auto":    "overflow-block: auto",
	}
}

// TransformTranslateZ represent the CSS style "transform-translate-z" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformTranslateZStyle string

func (t TransformTranslateZStyle) Name() string {
	return "transform-translate-z"
}

func (t TransformTranslateZStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformTranslateZStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-translate-z": "transform-translate-z: translateZ(0)",
	}
}

// BorderBottomStyle represent the CSS style "border-bottom-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-style
type BorderBottomStyleStyle string

func (t BorderBottomStyleStyle) Name() string {
	return "border-bottom-style"
}

func (t BorderBottomStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBottomStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"border-bottom-style": "border-bottom-style: none",
	}
}

// Filter represent the CSS style "filter" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/filter
type FilterStyle string

func (t FilterStyle) Name() string {
	return "filter"
}

func (t FilterStyle) BrowserVariants() []string {
	return []string{
		"-ms-filter",
	}
}

func (t FilterStyle) Utilities() map[string]string {
	return map[string]string{
		"filter": "filter: none",
	}
}

// PerspectiveOrigin represent the CSS style "perspective-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/perspective-origin
type PerspectiveOriginStyle string

func (t PerspectiveOriginStyle) Name() string {
	return "perspective-origin"
}

func (t PerspectiveOriginStyle) BrowserVariants() []string {
	return []string{}
}

func (t PerspectiveOriginStyle) Utilities() map[string]string {
	return map[string]string{
		"perspective-origin": "perspective-origin: 50% 50%",
	}
}

// TextDecorationColor represent the CSS style "text-decoration-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-color
type TextDecorationColorStyle Color

func (t TextDecorationColorStyle) Name() string {
	return "text-decoration-color"
}

func (t TextDecorationColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextDecorationColorStyle) Utilities() map[string]string {
	return map[string]string{
		"text-decoration-color": "text-decoration-color: currentcolor",
	}
}

// TextDecorationThickness represent the CSS style "text-decoration-thickness" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-thickness
type TextDecorationThicknessStyle string

func (t TextDecorationThicknessStyle) Name() string {
	return "text-decoration-thickness"
}

func (t TextDecorationThicknessStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextDecorationThicknessStyle) Utilities() map[string]string {
	return map[string]string{
		"text-decoration-thickness": "text-decoration-thickness: auto",
	}
}

// FillOpacity represent the CSS style "fill-opacity" with value ""
// See
type FillOpacityStyle float64

func (t FillOpacityStyle) Name() string {
	return "fill-opacity"
}

func (t FillOpacityStyle) BrowserVariants() []string {
	return []string{}
}

func (t FillOpacityStyle) Utilities() map[string]string {
	return map[string]string{
		"fill-opacity": "fill-opacity: 1",
	}
}

// ScrollMarginRight represent the CSS style "scroll-margin-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-right
type ScrollMarginRightStyle float64

func (t ScrollMarginRightStyle) Name() string {
	return "scroll-margin-right"
}

func (t ScrollMarginRightStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginRightStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin-right": "scroll-margin-right: 0",
	}
}

// Isolation represent the CSS style "isolation" with value "auto | isolate"
// See https://developer.mozilla.org/docs/Web/CSS/isolation
type IsolationStyle string

func (t IsolationStyle) Name() string {
	return "isolation"
}

const IsolationStyleAuto IsolationStyle = "auto"

const IsolationStyleIsolate IsolationStyle = "isolate"

func (t IsolationStyle) BrowserVariants() []string {
	return []string{}
}

func (t IsolationStyle) Utilities() map[string]string {
	return map[string]string{
		"isolation":         "isolation: auto",
		"isolation-auto":    "isolation: auto",
		"isolation-isolate": "isolation: isolate",
	}
}

// BorderTopWidth represent the CSS style "border-top-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-width
type BorderTopWidthStyle string

func (t BorderTopWidthStyle) Name() string {
	return "border-top-width"
}

func (t BorderTopWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderTopWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-top-width": "border-top-width: medium",
	}
}

// ImageOrientation represent the CSS style "image-orientation" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/image-orientation
type ImageOrientationStyle string

func (t ImageOrientationStyle) Name() string {
	return "image-orientation"
}

func (t ImageOrientationStyle) BrowserVariants() []string {
	return []string{}
}

func (t ImageOrientationStyle) Utilities() map[string]string {
	return map[string]string{
		"image-orientation": "image-orientation: from-image",
	}
}

// Color represent the CSS style "color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/color
type ColorStyle Color

func (t ColorStyle) Name() string {
	return "color"
}

func (t ColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColorStyle) Utilities() map[string]string {
	return map[string]string{
		"color": "color: variesFromBrowserToBrowser",
	}
}

// MinInlineSize represent the CSS style "min-inline-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-inline-size
type MinInlineSizeStyle string

func (t MinInlineSizeStyle) Name() string {
	return "min-inline-size"
}

func (t MinInlineSizeStyle) BrowserVariants() []string {
	return []string{}
}

func (t MinInlineSizeStyle) Utilities() map[string]string {
	return map[string]string{
		"min-inline-size": "min-inline-size: 0",
	}
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

func (t TransformBoxStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformBoxStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-box":             "transform-box: view-box",
		"transform-box-content-box": "transform-box: content-box",
		"transform-box-border-box":  "transform-box: border-box",
		"transform-box-fill-box":    "transform-box: fill-box",
		"transform-box-stroke-box":  "transform-box: stroke-box",
		"transform-box-view-box":    "transform-box: view-box",
	}
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

func (t AlignTracksStyle) BrowserVariants() []string {
	return []string{}
}

func (t AlignTracksStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// BackgroundAttachment represent the CSS style "background-attachment" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-attachment
type BackgroundAttachmentStyle string

func (t BackgroundAttachmentStyle) Name() string {
	return "background-attachment"
}

func (t BackgroundAttachmentStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundAttachmentStyle) Utilities() map[string]string {
	return map[string]string{
		"background-attachment": "background-attachment: scroll",
	}
}

// FlexShrink represent the CSS style "flex-shrink" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/flex-shrink
type FlexShrinkStyle float64

func (t FlexShrinkStyle) Name() string {
	return "flex-shrink"
}

func (t FlexShrinkStyle) BrowserVariants() []string {
	return []string{}
}

func (t FlexShrinkStyle) Utilities() map[string]string {
	return map[string]string{
		"flex-shrink": "flex-shrink: 1",
	}
}

// GridColumnEnd represent the CSS style "grid-column-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-column-end
type GridColumnEndStyle string

func (t GridColumnEndStyle) Name() string {
	return "grid-column-end"
}

func (t GridColumnEndStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridColumnEndStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-column-end": "grid-column-end: auto",
	}
}

// MaskImage represent the CSS style "mask-image" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-image
type MaskImageStyle string

func (t MaskImageStyle) Name() string {
	return "mask-image"
}

func (t MaskImageStyle) BrowserVariants() []string {
	return []string{
		"-webkit-mask-image",
	}
}

func (t MaskImageStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-image": "mask-image: none",
	}
}

// BorderBottomWidth represent the CSS style "border-bottom-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-width
type BorderBottomWidthStyle string

func (t BorderBottomWidthStyle) Name() string {
	return "border-bottom-width"
}

func (t BorderBottomWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBottomWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-bottom-width": "border-bottom-width: medium",
	}
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

func (t OverscrollBehaviorXStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverscrollBehaviorXStyle) Utilities() map[string]string {
	return map[string]string{
		"overscroll-behavior-x":         "overscroll-behavior-x: auto",
		"overscroll-behavior-x-contain": "overscroll-behavior-x: contain",
		"overscroll-behavior-x-none":    "overscroll-behavior-x: none",
		"overscroll-behavior-x-auto":    "overscroll-behavior-x: auto",
	}
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

func (t PositionStyle) BrowserVariants() []string {
	return []string{}
}

func (t PositionStyle) Utilities() map[string]string {
	return map[string]string{
		"position":          "position: static",
		"position-static":   "position: static",
		"position-relative": "position: relative",
		"position-absolute": "position: absolute",
		"position-sticky":   "position: sticky",
		"position-fixed":    "position: fixed",
	}
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

func (t TextAnchorStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextAnchorStyle) Utilities() map[string]string {
	return map[string]string{
		"text-anchor":        "text-anchor: start",
		"text-anchor-start":  "text-anchor: start",
		"text-anchor-middle": "text-anchor: middle",
		"text-anchor-end":    "text-anchor: end",
	}
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

func (t BoxAlignStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxAlignStyle) Utilities() map[string]string {
	return map[string]string{
		"box-align":          "box-align: stretch",
		"box-align-start":    "box-align: start",
		"box-align-center":   "box-align: center",
		"box-align-end":      "box-align: end",
		"box-align-baseline": "box-align: baseline",
		"box-align-stretch":  "box-align: stretch",
	}
}

// MinHeight represent the CSS style "min-height" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-height
type MinHeightStyle string

func (t MinHeightStyle) Name() string {
	return "min-height"
}

func (t MinHeightStyle) BrowserVariants() []string {
	return []string{}
}

func (t MinHeightStyle) Utilities() map[string]string {
	return map[string]string{
		"min-height": "min-height: auto",
	}
}

// FontVariantEastAsian represent the CSS style "font-variant-east-asian" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-east-asian
type FontVariantEastAsianStyle string

func (t FontVariantEastAsianStyle) Name() string {
	return "font-variant-east-asian"
}

func (t FontVariantEastAsianStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontVariantEastAsianStyle) Utilities() map[string]string {
	return map[string]string{
		"font-variant-east-asian": "font-variant-east-asian: normal",
	}
}

// MarginInline represent the CSS style "margin-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-inline
type MarginInlineStyle string

func (t MarginInlineStyle) Name() string {
	return "margin-inline"
}

func (t MarginInlineStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginInlineStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-inline": "margin-inline: 0",
	}
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

func (t OverflowYStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverflowYStyle) Utilities() map[string]string {
	return map[string]string{
		"overflow-y":         "overflow-y: visible",
		"overflow-y-visible": "overflow-y: visible",
		"overflow-y-hidden":  "overflow-y: hidden",
		"overflow-y-clip":    "overflow-y: clip",
		"overflow-y-scroll":  "overflow-y: scroll",
		"overflow-y-auto":    "overflow-y: auto",
	}
}

// TransformRotateZ represent the CSS style "transform-rotate-z" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformRotateZStyle string

func (t TransformRotateZStyle) Name() string {
	return "transform-rotate-z"
}

func (t TransformRotateZStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformRotateZStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-rotate-z": "transform-rotate-z: rotateZ(0)",
	}
}

// GridTemplateRows represent the CSS style "grid-template-rows" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-template-rows
type GridTemplateRowsStyle string

func (t GridTemplateRowsStyle) Name() string {
	return "grid-template-rows"
}

func (t GridTemplateRowsStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridTemplateRowsStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-template-rows": "grid-template-rows: none",
	}
}

// MarginBlockStart represent the CSS style "margin-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-block-start
type MarginBlockStartStyle string

func (t MarginBlockStartStyle) Name() string {
	return "margin-block-start"
}

func (t MarginBlockStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarginBlockStartStyle) Utilities() map[string]string {
	return map[string]string{
		"margin-block-start": "margin-block-start: 0",
	}
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

func (t MasonryAutoFlowStyle) BrowserVariants() []string {
	return []string{}
}

func (t MasonryAutoFlowStyle) Utilities() map[string]string {
	return map[string]string{
		"masonry-auto-flow":                     "masonry-auto-flow: pack",
		"masonry-auto-flow-pack":                "masonry-auto-flow: pack",
		"masonry-auto-flow-next":                "masonry-auto-flow: next",
		"masonry-auto-flow-pack-definite-first": "masonry-auto-flow: pack-definite-first",
		"masonry-auto-flow-next-definite-first": "masonry-auto-flow: next-definite-first",
		"masonry-auto-flow-pack-ordered":        "masonry-auto-flow: pack-ordered",
		"masonry-auto-flow-next-ordered":        "masonry-auto-flow: next-ordered",
	}
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

func (t BlockOverflowStyle) BrowserVariants() []string {
	return []string{}
}

func (t BlockOverflowStyle) Utilities() map[string]string {
	return map[string]string{
		"block-overflow":          "block-overflow: clip",
		"block-overflow-clip":     "block-overflow: clip",
		"block-overflow-ellipsis": "block-overflow: ellipsis",
		"block-overflow-string":   "block-overflow: string",
	}
}

// PaddingRight represent the CSS style "padding-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-right
type PaddingRightStyle string

func (t PaddingRightStyle) Name() string {
	return "padding-right"
}

func (t PaddingRightStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaddingRightStyle) Utilities() map[string]string {
	return map[string]string{
		"padding-right": "padding-right: 0",
	}
}

// GridColumnGap represent the CSS style "grid-column-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-gap
type GridColumnGapStyle string

func (t GridColumnGapStyle) Name() string {
	return "grid-column-gap"
}

func (t GridColumnGapStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridColumnGapStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-column-gap": "grid-column-gap: 0",
	}
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

func (t TextEmphasisStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextEmphasisStyleStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// TransformSkew represent the CSS style "transform-skew" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformSkewStyle string

func (t TransformSkewStyle) Name() string {
	return "transform-skew"
}

func (t TransformSkewStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformSkewStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-skew": "transform-skew: skew(0,0)",
	}
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

func (t ShapeRenderingStyle) BrowserVariants() []string {
	return []string{}
}

func (t ShapeRenderingStyle) Utilities() map[string]string {
	return map[string]string{
		"shape-rendering":                    "shape-rendering: auto",
		"shape-rendering-auto":               "shape-rendering: auto",
		"shape-rendering-optimizeSpeed":      "shape-rendering: optimizeSpeed",
		"shape-rendering-crispEdges":         "shape-rendering: crispEdges",
		"shape-rendering-geometricPrecision": "shape-rendering: geometricPrecision",
	}
}

// FontVariantAlternates represent the CSS style "font-variant-alternates" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-alternates
type FontVariantAlternatesStyle string

func (t FontVariantAlternatesStyle) Name() string {
	return "font-variant-alternates"
}

func (t FontVariantAlternatesStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontVariantAlternatesStyle) Utilities() map[string]string {
	return map[string]string{
		"font-variant-alternates": "font-variant-alternates: normal",
	}
}

// MaskRepeat represent the CSS style "mask-repeat" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-repeat
type MaskRepeatStyle string

func (t MaskRepeatStyle) Name() string {
	return "mask-repeat"
}

func (t MaskRepeatStyle) BrowserVariants() []string {
	return []string{
		"-webkit-mask-repeat",
	}
}

func (t MaskRepeatStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-repeat": "mask-repeat: no-repeat",
	}
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

func (t FontStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"font-style":         "font-style: normal",
		"font-style-normal":  "font-style: normal",
		"font-style-italic":  "font-style: italic",
		"font-style-oblique": "font-style: oblique",
		"font-style-angle":   "font-style: angle",
	}
}

// ColumnCount represent the CSS style "column-count" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-count
type ColumnCountStyle string

func (t ColumnCountStyle) Name() string {
	return "column-count"
}

func (t ColumnCountStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColumnCountStyle) Utilities() map[string]string {
	return map[string]string{
		"column-count": "column-count: auto",
	}
}

// ColumnWidth represent the CSS style "column-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-width
type ColumnWidthStyle string

func (t ColumnWidthStyle) Name() string {
	return "column-width"
}

func (t ColumnWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t ColumnWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"column-width": "column-width: auto",
	}
}

// BorderBlockColor represent the CSS style "border-block-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-color
type BorderBlockColorStyle Color

func (t BorderBlockColorStyle) Name() string {
	return "border-block-color"
}

func (t BorderBlockColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBlockColorStyle) Utilities() map[string]string {
	return map[string]string{
		"border-block-color": "border-block-color: currentcolor",
	}
}

// BorderInlineWidth represent the CSS style "border-inline-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-width
type BorderInlineWidthStyle string

func (t BorderInlineWidthStyle) Name() string {
	return "border-inline-width"
}

func (t BorderInlineWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderInlineWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-inline-width": "border-inline-width: medium",
	}
}

// FontFeatureSettings represent the CSS style "font-feature-settings" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-feature-settings
type FontFeatureSettingsStyle string

func (t FontFeatureSettingsStyle) Name() string {
	return "font-feature-settings"
}

func (t FontFeatureSettingsStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontFeatureSettingsStyle) Utilities() map[string]string {
	return map[string]string{
		"font-feature-settings": "font-feature-settings: normal",
	}
}

// Perspective represent the CSS style "perspective" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/perspective
type PerspectiveStyle string

func (t PerspectiveStyle) Name() string {
	return "perspective"
}

func (t PerspectiveStyle) BrowserVariants() []string {
	return []string{}
}

func (t PerspectiveStyle) Utilities() map[string]string {
	return map[string]string{
		"perspective": "perspective: none",
	}
}

// StopColor represent the CSS style "stop-color" with value ""
// See
type StopColorStyle Color

func (t StopColorStyle) Name() string {
	return "stop-color"
}

func (t StopColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t StopColorStyle) Utilities() map[string]string {
	return map[string]string{
		"stop-color": "stop-color: black",
	}
}

// MaskComposite represent the CSS style "mask-composite" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-composite
type MaskCompositeStyle string

func (t MaskCompositeStyle) Name() string {
	return "mask-composite"
}

func (t MaskCompositeStyle) BrowserVariants() []string {
	return []string{
		"-webkit-mask-composite",
	}
}

func (t MaskCompositeStyle) Utilities() map[string]string {
	return map[string]string{
		"mask-composite": "mask-composite: add",
	}
}

// BoxLines represent the CSS style "box-lines" with value "single | multiple"
// See https://developer.mozilla.org/docs/Web/CSS/box-lines
type BoxLinesStyle string

func (t BoxLinesStyle) Name() string {
	return "box-lines"
}

const BoxLinesStyleSingle BoxLinesStyle = "single"

const BoxLinesStyleMultiple BoxLinesStyle = "multiple"

func (t BoxLinesStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxLinesStyle) Utilities() map[string]string {
	return map[string]string{
		"box-lines":          "box-lines: single",
		"box-lines-single":   "box-lines: single",
		"box-lines-multiple": "box-lines: multiple",
	}
}

// Right represent the CSS style "right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/right
type RightStyle string

func (t RightStyle) Name() string {
	return "right"
}

func (t RightStyle) BrowserVariants() []string {
	return []string{}
}

func (t RightStyle) Utilities() map[string]string {
	return map[string]string{
		"right": "right: auto",
	}
}

// BorderBlockEndStyle represent the CSS style "border-block-end-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-end-style
type BorderBlockEndStyleStyle string

func (t BorderBlockEndStyleStyle) Name() string {
	return "border-block-end-style"
}

func (t BorderBlockEndStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderBlockEndStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"border-block-end-style": "border-block-end-style: none",
	}
}

// BorderInlineEndWidth represent the CSS style "border-inline-end-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-end-width
type BorderInlineEndWidthStyle string

func (t BorderInlineEndWidthStyle) Name() string {
	return "border-inline-end-width"
}

func (t BorderInlineEndWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderInlineEndWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-inline-end-width": "border-inline-end-width: medium",
	}
}

// Left represent the CSS style "left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/left
type LeftStyle string

func (t LeftStyle) Name() string {
	return "left"
}

func (t LeftStyle) BrowserVariants() []string {
	return []string{}
}

func (t LeftStyle) Utilities() map[string]string {
	return map[string]string{
		"left": "left: auto",
	}
}

// ScrollMarginLeft represent the CSS style "scroll-margin-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-left
type ScrollMarginLeftStyle float64

func (t ScrollMarginLeftStyle) Name() string {
	return "scroll-margin-left"
}

func (t ScrollMarginLeftStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollMarginLeftStyle) Utilities() map[string]string {
	return map[string]string{
		"scroll-margin-left": "scroll-margin-left: 0",
	}
}

// BoxFlexGroup represent the CSS style "box-flex-group" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/box-flex-group
type BoxFlexGroupStyle float64

func (t BoxFlexGroupStyle) Name() string {
	return "box-flex-group"
}

func (t BoxFlexGroupStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxFlexGroupStyle) Utilities() map[string]string {
	return map[string]string{
		"box-flex-group": "box-flex-group: 1",
	}
}

// TransformTranslate represent the CSS style "transform-translate" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformTranslateStyle string

func (t TransformTranslateStyle) Name() string {
	return "transform-translate"
}

func (t TransformTranslateStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformTranslateStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-translate": "transform-translate: translate(0)",
	}
}

// BorderImageSlice represent the CSS style "border-image-slice" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-slice
type BorderImageSliceStyle string

func (t BorderImageSliceStyle) Name() string {
	return "border-image-slice"
}

func (t BorderImageSliceStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderImageSliceStyle) Utilities() map[string]string {
	return map[string]string{
		"border-image-slice": "border-image-slice: 100%",
	}
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

func (t BreakBeforeStyle) BrowserVariants() []string {
	return []string{}
}

func (t BreakBeforeStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// WordSpacing represent the CSS style "word-spacing" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/word-spacing
type WordSpacingStyle string

func (t WordSpacingStyle) Name() string {
	return "word-spacing"
}

func (t WordSpacingStyle) BrowserVariants() []string {
	return []string{}
}

func (t WordSpacingStyle) Utilities() map[string]string {
	return map[string]string{
		"word-spacing": "word-spacing: normal",
	}
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

func (t OverflowXStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverflowXStyle) Utilities() map[string]string {
	return map[string]string{
		"overflow-x":         "overflow-x: visible",
		"overflow-x-visible": "overflow-x: visible",
		"overflow-x-hidden":  "overflow-x: hidden",
		"overflow-x-clip":    "overflow-x: clip",
		"overflow-x-scroll":  "overflow-x: scroll",
		"overflow-x-auto":    "overflow-x: auto",
	}
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

func (t OverscrollBehaviorBlockStyle) BrowserVariants() []string {
	return []string{}
}

func (t OverscrollBehaviorBlockStyle) Utilities() map[string]string {
	return map[string]string{
		"overscroll-behavior-block":         "overscroll-behavior-block: auto",
		"overscroll-behavior-block-contain": "overscroll-behavior-block: contain",
		"overscroll-behavior-block-none":    "overscroll-behavior-block: none",
		"overscroll-behavior-block-auto":    "overscroll-behavior-block: auto",
	}
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

func (t AlignContentStyle) BrowserVariants() []string {
	return []string{}
}

func (t AlignContentStyle) Utilities() map[string]string {
	return map[string]string{
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
}

// BorderTopColor represent the CSS style "border-top-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-color
type BorderTopColorStyle Color

func (t BorderTopColorStyle) Name() string {
	return "border-top-color"
}

func (t BorderTopColorStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderTopColorStyle) Utilities() map[string]string {
	return map[string]string{
		"border-top-color": "border-top-color: currentcolor",
	}
}

// OutlineOffset represent the CSS style "outline-offset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-offset
type OutlineOffsetStyle float64

func (t OutlineOffsetStyle) Name() string {
	return "outline-offset"
}

func (t OutlineOffsetStyle) BrowserVariants() []string {
	return []string{}
}

func (t OutlineOffsetStyle) Utilities() map[string]string {
	return map[string]string{
		"outline-offset": "outline-offset: 0",
	}
}

// MarkerStart represent the CSS style "marker-start" with value ""
// See
type MarkerStartStyle string

func (t MarkerStartStyle) Name() string {
	return "marker-start"
}

func (t MarkerStartStyle) BrowserVariants() []string {
	return []string{}
}

func (t MarkerStartStyle) Utilities() map[string]string {
	return map[string]string{
		"marker-start": "marker-start: none",
	}
}

// FontSize represent the CSS style "font-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-size
type FontSizeStyle string

func (t FontSizeStyle) Name() string {
	return "font-size"
}

func (t FontSizeStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontSizeStyle) Utilities() map[string]string {
	return map[string]string{
		"font-size": "font-size: medium",
	}
}

// FontStretch represent the CSS style "font-stretch" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-stretch
type FontStretchStyle string

func (t FontStretchStyle) Name() string {
	return "font-stretch"
}

func (t FontStretchStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontStretchStyle) Utilities() map[string]string {
	return map[string]string{
		"font-stretch": "font-stretch: normal",
	}
}

// OffsetDistance represent the CSS style "offset-distance" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/offset-distance
type OffsetDistanceStyle string

func (t OffsetDistanceStyle) Name() string {
	return "offset-distance"
}

func (t OffsetDistanceStyle) BrowserVariants() []string {
	return []string{}
}

func (t OffsetDistanceStyle) Utilities() map[string]string {
	return map[string]string{
		"offset-distance": "offset-distance: 0",
	}
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

func (t BaselineShiftStyle) BrowserVariants() []string {
	return []string{}
}

func (t BaselineShiftStyle) Utilities() map[string]string {
	return map[string]string{
		"baseline-shift":          "baseline-shift: baseline",
		"baseline-shift-baseline": "baseline-shift: baseline",
		"baseline-shift-sub":      "baseline-shift: sub",
		"baseline-shift-super":    "baseline-shift: super",
	}
}

// BorderInlineStartStyle represent the CSS style "border-inline-start-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-start-style
type BorderInlineStartStyleStyle string

func (t BorderInlineStartStyleStyle) Name() string {
	return "border-inline-start-style"
}

func (t BorderInlineStartStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderInlineStartStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"border-inline-start-style": "border-inline-start-style: none",
	}
}

// TransformOrigin represent the CSS style "transform-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformOriginStyle string

func (t TransformOriginStyle) Name() string {
	return "transform-origin"
}

func (t TransformOriginStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformOriginStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-origin": "transform-origin: 50% 50% 0",
	}
}

// TransitionDelay represent the CSS style "transition-delay" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-delay
type TransitionDelayStyle time.Duration

func (t TransitionDelayStyle) Name() string {
	return "transition-delay"
}

func (t TransitionDelayStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransitionDelayStyle) Utilities() map[string]string {
	return map[string]string{
		"transition-delay": "transition-delay: 0s",
	}
}

// ShapeImageThreshold represent the CSS style "shape-image-threshold" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/shape-image-threshold
type ShapeImageThresholdStyle float64

func (t ShapeImageThresholdStyle) Name() string {
	return "shape-image-threshold"
}

func (t ShapeImageThresholdStyle) BrowserVariants() []string {
	return []string{}
}

func (t ShapeImageThresholdStyle) Utilities() map[string]string {
	return map[string]string{
		"shape-image-threshold": "shape-image-threshold: 0.0",
	}
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

func (t ClearStyle) BrowserVariants() []string {
	return []string{}
}

func (t ClearStyle) Utilities() map[string]string {
	return map[string]string{
		"clear":              "clear: none",
		"clear-none":         "clear: none",
		"clear-left":         "clear: left",
		"clear-right":        "clear: right",
		"clear-both":         "clear: both",
		"clear-inline-start": "clear: inline-start",
		"clear-inline-end":   "clear: inline-end",
	}
}

// CounterReset represent the CSS style "counter-reset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/counter-reset
type CounterResetStyle string

func (t CounterResetStyle) Name() string {
	return "counter-reset"
}

func (t CounterResetStyle) BrowserVariants() []string {
	return []string{}
}

func (t CounterResetStyle) Utilities() map[string]string {
	return map[string]string{
		"counter-reset": "counter-reset: none",
	}
}

// ScaleRight represent the CSS style "scale-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scale
type ScaleRightStyle string

func (t ScaleRightStyle) Name() string {
	return "scale-right"
}

func (t ScaleRightStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScaleRightStyle) Utilities() map[string]string {
	return map[string]string{
		"scale-right": "scale-right: none",
	}
}

// BackgroundImage represent the CSS style "background-image" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-image
type BackgroundImageStyle URL

func (t BackgroundImageStyle) Name() string {
	return "background-image"
}

func (t BackgroundImageStyle) BrowserVariants() []string {
	return []string{}
}

func (t BackgroundImageStyle) Utilities() map[string]string {
	return map[string]string{
		"background-image": "background-image: none",
	}
}

// TransitionDuration represent the CSS style "transition-duration" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-duration
type TransitionDurationStyle time.Duration

func (t TransitionDurationStyle) Name() string {
	return "transition-duration"
}

func (t TransitionDurationStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransitionDurationStyle) Utilities() map[string]string {
	return map[string]string{
		"transition-duration": "transition-duration: 0s",
	}
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

func (t FontVariantPositionStyle) BrowserVariants() []string {
	return []string{}
}

func (t FontVariantPositionStyle) Utilities() map[string]string {
	return map[string]string{
		"font-variant-position":        "font-variant-position: normal",
		"font-variant-position-normal": "font-variant-position: normal",
		"font-variant-position-sub":    "font-variant-position: sub",
		"font-variant-position-super":  "font-variant-position: super",
	}
}

// GridTemplateAreas represent the CSS style "grid-template-areas" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-template-areas
type GridTemplateAreasStyle string

func (t GridTemplateAreasStyle) Name() string {
	return "grid-template-areas"
}

func (t GridTemplateAreasStyle) BrowserVariants() []string {
	return []string{}
}

func (t GridTemplateAreasStyle) Utilities() map[string]string {
	return map[string]string{
		"grid-template-areas": "grid-template-areas: none",
	}
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

func (t ScrollSnapAlignStyle) BrowserVariants() []string {
	return []string{}
}

func (t ScrollSnapAlignStyle) Utilities() map[string]string {
	return map[string]string{
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

func (t WidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t WidthStyle) Utilities() map[string]string {
	return map[string]string{
		"width":             "width: auto",
		"width-auto":        "width: auto",
		"width-min-content": "width: min-content",
		"width-max-content": "width: max-content",
	}
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

func (t TextOrientationStyle) BrowserVariants() []string {
	return []string{}
}

func (t TextOrientationStyle) Utilities() map[string]string {
	return map[string]string{
		"text-orientation":          "text-orientation: mixed",
		"text-orientation-mixed":    "text-orientation: mixed",
		"text-orientation-upright":  "text-orientation: upright",
		"text-orientation-sideways": "text-orientation: sideways",
	}
}

// BorderInlineStyle represent the CSS style "border-inline-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-style
type BorderInlineStyleStyle string

func (t BorderInlineStyleStyle) Name() string {
	return "border-inline-style"
}

func (t BorderInlineStyleStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderInlineStyleStyle) Utilities() map[string]string {
	return map[string]string{
		"border-inline-style": "border-inline-style: none",
	}
}

// ImageResolution represent the CSS style "image-resolution" with value ""
// See
type ImageResolutionStyle string

func (t ImageResolutionStyle) Name() string {
	return "image-resolution"
}

func (t ImageResolutionStyle) BrowserVariants() []string {
	return []string{}
}

func (t ImageResolutionStyle) Utilities() map[string]string {
	return map[string]string{
		"image-resolution": "image-resolution: 1dppx",
	}
}

// BorderImageWidth represent the CSS style "border-image-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-width
type BorderImageWidthStyle string

func (t BorderImageWidthStyle) Name() string {
	return "border-image-width"
}

func (t BorderImageWidthStyle) BrowserVariants() []string {
	return []string{}
}

func (t BorderImageWidthStyle) Utilities() map[string]string {
	return map[string]string{
		"border-image-width": "border-image-width: 1",
	}
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

func (t BoxOrientStyle) BrowserVariants() []string {
	return []string{}
}

func (t BoxOrientStyle) Utilities() map[string]string {
	return map[string]string{
		"box-orient":             "box-orient: inlineAxisHorizontalInXUL",
		"box-orient-horizontal":  "box-orient: horizontal",
		"box-orient-vertical":    "box-orient: vertical",
		"box-orient-inline-axis": "box-orient: inline-axis",
		"box-orient-block-axis":  "box-orient: block-axis",
		"box-orient-inherit":     "box-orient: inherit",
	}
}

// LineHeight represent the CSS style "line-height" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/line-height
type LineHeightStyle string

func (t LineHeightStyle) Name() string {
	return "line-height"
}

func (t LineHeightStyle) BrowserVariants() []string {
	return []string{}
}

func (t LineHeightStyle) Utilities() map[string]string {
	return map[string]string{
		"line-height": "line-height: normal",
	}
}

// PlaceContent represent the CSS style "place-content" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/place-content
type PlaceContentStyle string

func (t PlaceContentStyle) Name() string {
	return "place-content"
}

func (t PlaceContentStyle) BrowserVariants() []string {
	return []string{}
}

func (t PlaceContentStyle) Utilities() map[string]string {
	return map[string]string{
		"place-content": "place-content: normal",
	}
}

// TransformScaleX represent the CSS style "transform-scale-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformScaleXStyle string

func (t TransformScaleXStyle) Name() string {
	return "transform-scale-x"
}

func (t TransformScaleXStyle) BrowserVariants() []string {
	return []string{}
}

func (t TransformScaleXStyle) Utilities() map[string]string {
	return map[string]string{
		"transform-scale-x": "transform-scale-x: scaleX(0)",
	}
}

// PaddingBottom represent the CSS style "padding-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-bottom
type PaddingBottomStyle string

func (t PaddingBottomStyle) Name() string {
	return "padding-bottom"
}

func (t PaddingBottomStyle) BrowserVariants() []string {
	return []string{}
}

func (t PaddingBottomStyle) Utilities() map[string]string {
	return map[string]string{
		"padding-bottom": "padding-bottom: 0",
	}
}
