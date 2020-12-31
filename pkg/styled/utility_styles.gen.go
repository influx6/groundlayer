package styled

import "time"

// HangingPunctuation represent the CSS style "hanging-punctuation" with value "none | first  force-end | allow-end | last"
// See https://developer.mozilla.org/docs/Web/CSS/hanging-punctuation
type HangingPunctuationStyle string

func (t HangingPunctuationStyle) Name() string {
	return "hanging-punctuation"
}

const HangingPunctuationStyleNone = "none"

const HangingPunctuationStyleFirstForceEnd = "first  force-end"

const HangingPunctuationStyleAllowEnd = "allow-end"

const HangingPunctuationStyleLast = "last"

// InlineSize represent the CSS style "inline-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inline-size
type InlineSizeStyle string

func (t InlineSizeStyle) Name() string {
	return "inline-size"
}

// JustifyContent represent the CSS style "justify-content" with value "auto | normal | stretch | end | start | flex-start | flex-end | center | left | right |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe center | unsafe center"
// See https://developer.mozilla.org/docs/Web/CSS/justify-content
type JustifyContentStyle string

func (t JustifyContentStyle) Name() string {
	return "justify-content"
}

const JustifyContentStyleAuto = "auto"

const JustifyContentStyleNormal = "normal"

const JustifyContentStyleStretch = "stretch"

const JustifyContentStyleEnd = "end"

const JustifyContentStyleStart = "start"

const JustifyContentStyleFlexStart = "flex-start"

const JustifyContentStyleFlexEnd = "flex-end"

const JustifyContentStyleCenter = "center"

const JustifyContentStyleLeft = "left"

const JustifyContentStyleRight = "right"

const JustifyContentStyleBaseline = "baseline"

const JustifyContentStyleFirstBaseline = "first baseline"

const JustifyContentStyleLastBaseline = "last baseline"

const JustifyContentStyleSpaceBetween = "space-between"

const JustifyContentStyleSpaceAround = "space-around"

const JustifyContentStyleSpaceEvenly = "space-evenly"

const JustifyContentStyleSafeCenter = "safe center"

const JustifyContentStyleUnsafeCenter = "unsafe center"

// MaskBorderSlice represent the CSS style "mask-border-slice" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-slice
type MaskBorderSliceStyle string

func (t MaskBorderSliceStyle) Name() string {
	return "mask-border-slice"
}

// TextAlign represent the CSS style "text-align" with value "start | end | left | right | center | justify | match-parent"
// See https://developer.mozilla.org/docs/Web/CSS/text-align
type TextAlignStyle string

func (t TextAlignStyle) Name() string {
	return "text-align"
}

const TextAlignStyleStart = "start"

const TextAlignStyleEnd = "end"

const TextAlignStyleLeft = "left"

const TextAlignStyleRight = "right"

const TextAlignStyleCenter = "center"

const TextAlignStyleJustify = "justify"

const TextAlignStyleMatchParent = "match-parent"

// ZIndex represent the CSS style "z-index" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/z-index
type ZIndexStyle string

func (t ZIndexStyle) Name() string {
	return "z-index"
}

// BorderTopColor represent the CSS style "border-top-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-color
type BorderTopColorStyle Color

func (t BorderTopColorStyle) Name() string {
	return "border-top-color"
}

// BoxLines represent the CSS style "box-lines" with value "single | multiple"
// See https://developer.mozilla.org/docs/Web/CSS/box-lines
type BoxLinesStyle string

func (t BoxLinesStyle) Name() string {
	return "box-lines"
}

const BoxLinesStyleSingle = "single"

const BoxLinesStyleMultiple = "multiple"

// MarginInline represent the CSS style "margin-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-inline
type MarginInlineStyle string

func (t MarginInlineStyle) Name() string {
	return "margin-inline"
}

// OffsetDistance represent the CSS style "offset-distance" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/offset-distance
type OffsetDistanceStyle string

func (t OffsetDistanceStyle) Name() string {
	return "offset-distance"
}

// PaddingBlockEnd represent the CSS style "padding-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-block-end
type PaddingBlockEndStyle string

func (t PaddingBlockEndStyle) Name() string {
	return "padding-block-end"
}

// EmptyCells represent the CSS style "empty-cells" with value "show | hide"
// See https://developer.mozilla.org/docs/Web/CSS/empty-cells
type EmptyCellsStyle string

func (t EmptyCellsStyle) Name() string {
	return "empty-cells"
}

const EmptyCellsStyleShow = "show"

const EmptyCellsStyleHide = "hide"

// FontVariantCaps represent the CSS style "font-variant-caps" with value "normal | small-caps | all-small-caps | petite-caps | all-petite-caps | unicase | titling-caps"
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-caps
type FontVariantCapsStyle string

func (t FontVariantCapsStyle) Name() string {
	return "font-variant-caps"
}

const FontVariantCapsStyleNormal = "normal"

const FontVariantCapsStyleSmallCaps = "small-caps"

const FontVariantCapsStyleAllSmallCaps = "all-small-caps"

const FontVariantCapsStylePetiteCaps = "petite-caps"

const FontVariantCapsStyleAllPetiteCaps = "all-petite-caps"

const FontVariantCapsStyleUnicase = "unicase"

const FontVariantCapsStyleTitlingCaps = "titling-caps"

// LetterSpacing represent the CSS style "letter-spacing" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/letter-spacing
type LetterSpacingStyle string

func (t LetterSpacingStyle) Name() string {
	return "letter-spacing"
}

// OverscrollBehaviorY represent the CSS style "overscroll-behavior-y" with value "contain | none | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overscroll-behavior-y
type OverscrollBehaviorYStyle string

func (t OverscrollBehaviorYStyle) Name() string {
	return "overscroll-behavior-y"
}

const OverscrollBehaviorYStyleContain = "contain"

const OverscrollBehaviorYStyleNone = "none"

const OverscrollBehaviorYStyleAuto = "auto"

// TransitionTimingFunction represent the CSS style "transition-timing-function" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-timing-function
type TransitionTimingFunctionStyle string

func (t TransitionTimingFunctionStyle) Name() string {
	return "transition-timing-function"
}

// BorderInlineStartColor represent the CSS style "border-inline-start-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-start-color
type BorderInlineStartColorStyle Color

func (t BorderInlineStartColorStyle) Name() string {
	return "border-inline-start-color"
}

// FlexGrow represent the CSS style "flex-grow" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/flex-grow
type FlexGrowStyle float64

func (t FlexGrowStyle) Name() string {
	return "flex-grow"
}

// OverflowClipBox represent the CSS style "overflow-clip-box" with value "padding-box | content-box"
// See https://developer.mozilla.org/docs/Mozilla/CSS/overflow-clip-box
type OverflowClipBoxStyle string

func (t OverflowClipBoxStyle) Name() string {
	return "overflow-clip-box"
}

const OverflowClipBoxStylePaddingBox = "padding-box"

const OverflowClipBoxStyleContentBox = "content-box"

// ScrollPaddingBottom represent the CSS style "scroll-padding-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-bottom
type ScrollPaddingBottomStyle string

func (t ScrollPaddingBottomStyle) Name() string {
	return "scroll-padding-bottom"
}

// TextIndent represent the CSS style "text-indent" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-indent
type TextIndentStyle string

func (t TextIndentStyle) Name() string {
	return "text-indent"
}

// ListStylePosition represent the CSS style "list-style-position" with value "inside | outside"
// See https://developer.mozilla.org/docs/Web/CSS/list-style-position
type ListStylePositionStyle string

func (t ListStylePositionStyle) Name() string {
	return "list-style-position"
}

const ListStylePositionStyleInside = "inside"

const ListStylePositionStyleOutside = "outside"

// MarginBlockEnd represent the CSS style "margin-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-block-end
type MarginBlockEndStyle string

func (t MarginBlockEndStyle) Name() string {
	return "margin-block-end"
}

// TextDecorationStyle represent the CSS style "text-decoration-style" with value "solid | double | dotted | dashed | wavy"
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-style
type TextDecorationStyleStyle string

func (t TextDecorationStyleStyle) Name() string {
	return "text-decoration-style"
}

const TextDecorationStyleStyleSolid = "solid"

const TextDecorationStyleStyleDouble = "double"

const TextDecorationStyleStyleDotted = "dotted"

const TextDecorationStyleStyleDashed = "dashed"

const TextDecorationStyleStyleWavy = "wavy"

// BackgroundPositionY represent the CSS style "background-position-y" with value "center | top | bottom | y-start | y-end"
// See https://developer.mozilla.org/docs/Web/CSS/background-position-y
type BackgroundPositionYStyle string

func (t BackgroundPositionYStyle) Name() string {
	return "background-position-y"
}

const BackgroundPositionYStyleCenter = "center"

const BackgroundPositionYStyleTop = "top"

const BackgroundPositionYStyleBottom = "bottom"

const BackgroundPositionYStyleYStart = "y-start"

const BackgroundPositionYStyleYEnd = "y-end"

// BorderBottomLeftRadius represent the CSS style "border-bottom-left-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-left-radius
type BorderBottomLeftRadiusStyle string

func (t BorderBottomLeftRadiusStyle) Name() string {
	return "border-bottom-left-radius"
}

// BorderInlineWidth represent the CSS style "border-inline-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-width
type BorderInlineWidthStyle string

func (t BorderInlineWidthStyle) Name() string {
	return "border-inline-width"
}

// ColumnGap represent the CSS style "column-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-gap
type ColumnGapStyle string

func (t ColumnGapStyle) Name() string {
	return "column-gap"
}

// GridRowStart represent the CSS style "grid-row-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-row-start
type GridRowStartStyle string

func (t GridRowStartStyle) Name() string {
	return "grid-row-start"
}

// Top represent the CSS style "top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/top
type TopStyle string

func (t TopStyle) Name() string {
	return "top"
}

// TouchAction represent the CSS style "touch-action" with value "auto | none | pan-x | pan-left | pan-right | pan-y | pan-up | pan-down | pinch-zoom | manipulation"
// See https://developer.mozilla.org/docs/Web/CSS/touch-action
type TouchActionStyle string

func (t TouchActionStyle) Name() string {
	return "touch-action"
}

const TouchActionStyleAuto = "auto"

const TouchActionStyleNone = "none"

const TouchActionStylePanX = "pan-x"

const TouchActionStylePanLeft = "pan-left"

const TouchActionStylePanRight = "pan-right"

const TouchActionStylePanY = "pan-y"

const TouchActionStylePanUp = "pan-up"

const TouchActionStylePanDown = "pan-down"

const TouchActionStylePinchZoom = "pinch-zoom"

const TouchActionStyleManipulation = "manipulation"

// WhiteSpace represent the CSS style "white-space" with value "normal | pre | nowrap | pre-wrap | pre-line | break-spaces"
// See https://developer.mozilla.org/docs/Web/CSS/white-space
type WhiteSpaceStyle string

func (t WhiteSpaceStyle) Name() string {
	return "white-space"
}

const WhiteSpaceStyleNormal = "normal"

const WhiteSpaceStylePre = "pre"

const WhiteSpaceStyleNowrap = "nowrap"

const WhiteSpaceStylePreWrap = "pre-wrap"

const WhiteSpaceStylePreLine = "pre-line"

const WhiteSpaceStyleBreakSpaces = "break-spaces"

// AnimationDirection represent the CSS style "animation-direction" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-direction
type AnimationDirectionStyle string

func (t AnimationDirectionStyle) Name() string {
	return "animation-direction"
}

// FontVariationSettings represent the CSS style "font-variation-settings" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variation-settings
type FontVariationSettingsStyle string

func (t FontVariationSettingsStyle) Name() string {
	return "font-variation-settings"
}

// FontSynthesis represent the CSS style "font-synthesis" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-synthesis
type FontSynthesisStyle string

func (t FontSynthesisStyle) Name() string {
	return "font-synthesis"
}

// ScrollMargin represent the CSS style "scroll-margin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin
type ScrollMarginStyle float64

func (t ScrollMarginStyle) Name() string {
	return "scroll-margin"
}

// TextOrientation represent the CSS style "text-orientation" with value "mixed | upright | sideways"
// See https://developer.mozilla.org/docs/Web/CSS/text-orientation
type TextOrientationStyle string

func (t TextOrientationStyle) Name() string {
	return "text-orientation"
}

const TextOrientationStyleMixed = "mixed"

const TextOrientationStyleUpright = "upright"

const TextOrientationStyleSideways = "sideways"

// OffsetPath represent the CSS style "offset-path" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/offset-path
type OffsetPathStyle string

func (t OffsetPathStyle) Name() string {
	return "offset-path"
}

// ScrollBehavior represent the CSS style "scroll-behavior" with value "auto | smooth"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-behavior
type ScrollBehaviorStyle string

func (t ScrollBehaviorStyle) Name() string {
	return "scroll-behavior"
}

const ScrollBehaviorStyleAuto = "auto"

const ScrollBehaviorStyleSmooth = "smooth"

// ScrollPaddingTop represent the CSS style "scroll-padding-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-top
type ScrollPaddingTopStyle string

func (t ScrollPaddingTopStyle) Name() string {
	return "scroll-padding-top"
}

// BackgroundRepeat represent the CSS style "background-repeat" with value "none | repeat | no-repeat"
// See https://developer.mozilla.org/docs/Web/CSS/background-repeat
type BackgroundRepeatStyle string

func (t BackgroundRepeatStyle) Name() string {
	return "background-repeat"
}

const BackgroundRepeatStyleNone = "none"

const BackgroundRepeatStyleRepeat = "repeat"

const BackgroundRepeatStyleNoRepeat = "no-repeat"

// BorderBlockEndWidth represent the CSS style "border-block-end-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-end-width
type BorderBlockEndWidthStyle string

func (t BorderBlockEndWidthStyle) Name() string {
	return "border-block-end-width"
}

// BorderImageWidth represent the CSS style "border-image-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-width
type BorderImageWidthStyle string

func (t BorderImageWidthStyle) Name() string {
	return "border-image-width"
}

// FontLanguageOverride represent the CSS style "font-language-override" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-language-override
type FontLanguageOverrideStyle string

func (t FontLanguageOverrideStyle) Name() string {
	return "font-language-override"
}

// FontOpticalSizing represent the CSS style "font-optical-sizing" with value "auto | none"
// See https://developer.mozilla.org/docs/Web/CSS/font-optical-sizing
type FontOpticalSizingStyle string

func (t FontOpticalSizingStyle) Name() string {
	return "font-optical-sizing"
}

const FontOpticalSizingStyleAuto = "auto"

const FontOpticalSizingStyleNone = "none"

// ScrollSnapAlign represent the CSS style "scroll-snap-align" with value "none | start | end | center | start end | start center | end start | end center | center start | center end"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-align
type ScrollSnapAlignStyle string

func (t ScrollSnapAlignStyle) Name() string {
	return "scroll-snap-align"
}

const ScrollSnapAlignStyleNone = "none"

const ScrollSnapAlignStyleStart = "start"

const ScrollSnapAlignStyleEnd = "end"

const ScrollSnapAlignStyleCenter = "center"

const ScrollSnapAlignStyleStartEnd = "start end"

const ScrollSnapAlignStyleStartCenter = "start center"

const ScrollSnapAlignStyleEndStart = "end start"

const ScrollSnapAlignStyleEndCenter = "end center"

const ScrollSnapAlignStyleCenterStart = "center start"

const ScrollSnapAlignStyleCenterEnd = "center end"

// Stroke represent the CSS style "stroke" with value ""
// See
type StrokeStyle Color

func (t StrokeStyle) Name() string {
	return "stroke"
}

// ColumnCount represent the CSS style "column-count" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-count
type ColumnCountStyle string

func (t ColumnCountStyle) Name() string {
	return "column-count"
}

// Gap represent the CSS style "gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/gap
type GapStyle string

func (t GapStyle) Name() string {
	return "gap"
}

// MaskRepeat represent the CSS style "mask-repeat" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-repeat
type MaskRepeatStyle string

func (t MaskRepeatStyle) Name() string {
	return "mask-repeat"
}

// ScrollMarginRight represent the CSS style "scroll-margin-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-right
type ScrollMarginRightStyle float64

func (t ScrollMarginRightStyle) Name() string {
	return "scroll-margin-right"
}

// TextEmphasisPositionSecond represent the CSS style "text-emphasis-position-second" with value "right | left"
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-position
type TextEmphasisPositionSecondStyle string

func (t TextEmphasisPositionSecondStyle) Name() string {
	return "text-emphasis-position-second"
}

const TextEmphasisPositionSecondStyleRight = "right"

const TextEmphasisPositionSecondStyleLeft = "left"

// Widows represent the CSS style "widows" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/widows
type WidowsStyle float64

func (t WidowsStyle) Name() string {
	return "widows"
}

// ColorInterpolation represent the CSS style "color-interpolation" with value "auto | sRGB | linearRGB"
// See
type ColorInterpolationStyle string

func (t ColorInterpolationStyle) Name() string {
	return "color-interpolation"
}

const ColorInterpolationStyleAuto = "auto"

const ColorInterpolationStyleSRGB = "sRGB"

const ColorInterpolationStyleLinearRGB = "linearRGB"

// BackgroundAttachment represent the CSS style "background-attachment" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-attachment
type BackgroundAttachmentStyle string

func (t BackgroundAttachmentStyle) Name() string {
	return "background-attachment"
}

// BorderBottomWidth represent the CSS style "border-bottom-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-width
type BorderBottomWidthStyle string

func (t BorderBottomWidthStyle) Name() string {
	return "border-bottom-width"
}

// FontFamily represent the CSS style "font-family" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-family
type FontFamilyStyle string

func (t FontFamilyStyle) Name() string {
	return "font-family"
}

// PerspectiveOrigin represent the CSS style "perspective-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/perspective-origin
type PerspectiveOriginStyle string

func (t PerspectiveOriginStyle) Name() string {
	return "perspective-origin"
}

// ScaleRight represent the CSS style "scale-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scale
type ScaleRightStyle string

func (t ScaleRightStyle) Name() string {
	return "scale-right"
}

// ScrollPaddingRight represent the CSS style "scroll-padding-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-right
type ScrollPaddingRightStyle string

func (t ScrollPaddingRightStyle) Name() string {
	return "scroll-padding-right"
}

// ScrollSnapTypeX represent the CSS style "scroll-snap-type-x" with value "none | mandatory | proximity"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-type-x
type ScrollSnapTypeXStyle string

func (t ScrollSnapTypeXStyle) Name() string {
	return "scroll-snap-type-x"
}

const ScrollSnapTypeXStyleNone = "none"

const ScrollSnapTypeXStyleMandatory = "mandatory"

const ScrollSnapTypeXStyleProximity = "proximity"

// BackdropFilter represent the CSS style "backdrop-filter" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/backdrop-filter
type BackdropFilterStyle string

func (t BackdropFilterStyle) Name() string {
	return "backdrop-filter"
}

// CounterReset represent the CSS style "counter-reset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/counter-reset
type CounterResetStyle string

func (t CounterResetStyle) Name() string {
	return "counter-reset"
}

// ImageOrientation represent the CSS style "image-orientation" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/image-orientation
type ImageOrientationStyle string

func (t ImageOrientationStyle) Name() string {
	return "image-orientation"
}

// ListStyleImage represent the CSS style "list-style-image" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/list-style-image
type ListStyleImageStyle URL

func (t ListStyleImageStyle) Name() string {
	return "list-style-image"
}

// PageBreakBefore represent the CSS style "page-break-before" with value "auto | always | avoid | left | right | recto | verso"
// See https://developer.mozilla.org/docs/Web/CSS/page-break-before
type PageBreakBeforeStyle string

func (t PageBreakBeforeStyle) Name() string {
	return "page-break-before"
}

const PageBreakBeforeStyleAuto = "auto"

const PageBreakBeforeStyleAlways = "always"

const PageBreakBeforeStyleAvoid = "avoid"

const PageBreakBeforeStyleLeft = "left"

const PageBreakBeforeStyleRight = "right"

const PageBreakBeforeStyleRecto = "recto"

const PageBreakBeforeStyleVerso = "verso"

// Height represent the CSS style "height" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/height
type HeightStyle string

func (t HeightStyle) Name() string {
	return "height"
}

// Order represent the CSS style "order" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/order
type OrderStyle float64

func (t OrderStyle) Name() string {
	return "order"
}

// PaddingRight represent the CSS style "padding-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-right
type PaddingRightStyle string

func (t PaddingRightStyle) Name() string {
	return "padding-right"
}

// Rotate represent the CSS style "rotate" with value "angle"
// See https://developer.mozilla.org/docs/Web/CSS/rotate
type RotateStyle string

func (t RotateStyle) Name() string {
	return "rotate"
}

const RotateStyleAngle = "angle"

// OverflowBlock represent the CSS style "overflow-block" with value "visible | hidden | clip | scroll | auto"
// See
type OverflowBlockStyle string

func (t OverflowBlockStyle) Name() string {
	return "overflow-block"
}

const OverflowBlockStyleVisible = "visible"

const OverflowBlockStyleHidden = "hidden"

const OverflowBlockStyleClip = "clip"

const OverflowBlockStyleScroll = "scroll"

const OverflowBlockStyleAuto = "auto"

// ScrollMarginInlineEnd represent the CSS style "scroll-margin-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-inline-end
type ScrollMarginInlineEndStyle float64

func (t ScrollMarginInlineEndStyle) Name() string {
	return "scroll-margin-inline-end"
}

// TextDecorationColor represent the CSS style "text-decoration-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-color
type TextDecorationColorStyle Color

func (t TextDecorationColorStyle) Name() string {
	return "text-decoration-color"
}

// AlignmentBaseline represent the CSS style "alignment-baseline" with value "auto | baseline | before-edge | text-before-edge | middle | central | after-edge | text-after-edge | ideographic | alphabetic | hanging | mathematical"
// See
type AlignmentBaselineStyle string

func (t AlignmentBaselineStyle) Name() string {
	return "alignment-baseline"
}

const AlignmentBaselineStyleAuto = "auto"

const AlignmentBaselineStyleBaseline = "baseline"

const AlignmentBaselineStyleBeforeEdge = "before-edge"

const AlignmentBaselineStyleTextBeforeEdge = "text-before-edge"

const AlignmentBaselineStyleMiddle = "middle"

const AlignmentBaselineStyleCentral = "central"

const AlignmentBaselineStyleAfterEdge = "after-edge"

const AlignmentBaselineStyleTextAfterEdge = "text-after-edge"

const AlignmentBaselineStyleIdeographic = "ideographic"

const AlignmentBaselineStyleAlphabetic = "alphabetic"

const AlignmentBaselineStyleHanging = "hanging"

const AlignmentBaselineStyleMathematical = "mathematical"

// GridColumnGap represent the CSS style "grid-column-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-gap
type GridColumnGapStyle string

func (t GridColumnGapStyle) Name() string {
	return "grid-column-gap"
}

// JustifyTracks represent the CSS style "justify-tracks" with value "auto | normal | stretch | end | start | flex-start | flex-end | center | left | right |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe center | unsafe center"
// See https://developer.mozilla.org/docs/Web/CSS/justify-tracks
type JustifyTracksStyle string

func (t JustifyTracksStyle) Name() string {
	return "justify-tracks"
}

const JustifyTracksStyleAuto = "auto"

const JustifyTracksStyleNormal = "normal"

const JustifyTracksStyleStretch = "stretch"

const JustifyTracksStyleEnd = "end"

const JustifyTracksStyleStart = "start"

const JustifyTracksStyleFlexStart = "flex-start"

const JustifyTracksStyleFlexEnd = "flex-end"

const JustifyTracksStyleCenter = "center"

const JustifyTracksStyleLeft = "left"

const JustifyTracksStyleRight = "right"

const JustifyTracksStyleBaseline = "baseline"

const JustifyTracksStyleFirstBaseline = "first baseline"

const JustifyTracksStyleLastBaseline = "last baseline"

const JustifyTracksStyleSpaceBetween = "space-between"

const JustifyTracksStyleSpaceAround = "space-around"

const JustifyTracksStyleSpaceEvenly = "space-evenly"

const JustifyTracksStyleSafeCenter = "safe center"

const JustifyTracksStyleUnsafeCenter = "unsafe center"

// MaskType represent the CSS style "mask-type" with value "luminance | alpha"
// See https://developer.mozilla.org/docs/Web/CSS/mask-type
type MaskTypeStyle string

func (t MaskTypeStyle) Name() string {
	return "mask-type"
}

const MaskTypeStyleLuminance = "luminance"

const MaskTypeStyleAlpha = "alpha"

// OverflowAnchor represent the CSS style "overflow-anchor" with value "auto | none"
// See
type OverflowAnchorStyle string

func (t OverflowAnchorStyle) Name() string {
	return "overflow-anchor"
}

const OverflowAnchorStyleAuto = "auto"

const OverflowAnchorStyleNone = "none"

// TransformBox represent the CSS style "transform-box" with value "content-box | border-box | fill-box | stroke-box | view-box"
// See https://developer.mozilla.org/docs/Web/CSS/transform-box
type TransformBoxStyle string

func (t TransformBoxStyle) Name() string {
	return "transform-box"
}

const TransformBoxStyleContentBox = "content-box"

const TransformBoxStyleBorderBox = "border-box"

const TransformBoxStyleFillBox = "fill-box"

const TransformBoxStyleStrokeBox = "stroke-box"

const TransformBoxStyleViewBox = "view-box"

// TransitionDuration represent the CSS style "transition-duration" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-duration
type TransitionDurationStyle time.Duration

func (t TransitionDurationStyle) Name() string {
	return "transition-duration"
}

// OffsetRotate represent the CSS style "offset-rotate" with value "auto | reverse | angle"
// See https://developer.mozilla.org/docs/Web/CSS/offset-rotate
type OffsetRotateStyle string

func (t OffsetRotateStyle) Name() string {
	return "offset-rotate"
}

const OffsetRotateStyleAuto = "auto"

const OffsetRotateStyleReverse = "reverse"

const OffsetRotateStyleAngle = "angle"

// OverscrollBehavior represent the CSS style "overscroll-behavior" with value "contain | none | auto | contain none | contain auto | none contain | none auto | auto contain | auto none"
// See https://developer.mozilla.org/docs/Web/CSS/overscroll-behavior
type OverscrollBehaviorStyle string

func (t OverscrollBehaviorStyle) Name() string {
	return "overscroll-behavior"
}

const OverscrollBehaviorStyleContain = "contain"

const OverscrollBehaviorStyleNone = "none"

const OverscrollBehaviorStyleAuto = "auto"

const OverscrollBehaviorStyleContainNone = "contain none"

const OverscrollBehaviorStyleContainAuto = "contain auto"

const OverscrollBehaviorStyleNoneContain = "none contain"

const OverscrollBehaviorStyleNoneAuto = "none auto"

const OverscrollBehaviorStyleAutoContain = "auto contain"

const OverscrollBehaviorStyleAutoNone = "auto none"

// ScrollSnapCoordinate represent the CSS style "scroll-snap-coordinate" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-coordinate
type ScrollSnapCoordinateStyle string

func (t ScrollSnapCoordinateStyle) Name() string {
	return "scroll-snap-coordinate"
}

// BorderBlockWidth represent the CSS style "border-block-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-width
type BorderBlockWidthStyle string

func (t BorderBlockWidthStyle) Name() string {
	return "border-block-width"
}

// BorderInlineEndColor represent the CSS style "border-inline-end-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-end-color
type BorderInlineEndColorStyle Color

func (t BorderInlineEndColorStyle) Name() string {
	return "border-inline-end-color"
}

// BorderStartStartRadius represent the CSS style "border-start-start-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-start-start-radius
type BorderStartStartRadiusStyle string

func (t BorderStartStartRadiusStyle) Name() string {
	return "border-start-start-radius"
}

// ColorAdjust represent the CSS style "color-adjust" with value "economy | exact"
// See https://developer.mozilla.org/docs/Web/CSS/color-adjust
type ColorAdjustStyle string

func (t ColorAdjustStyle) Name() string {
	return "color-adjust"
}

const ColorAdjustStyleEconomy = "economy"

const ColorAdjustStyleExact = "exact"

// FontStyle represent the CSS style "font-style" with value "normal | italic | oblique | angle"
// See https://developer.mozilla.org/docs/Web/CSS/font-style
type FontStyleStyle string

func (t FontStyleStyle) Name() string {
	return "font-style"
}

const FontStyleStyleNormal = "normal"

const FontStyleStyleItalic = "italic"

const FontStyleStyleOblique = "oblique"

const FontStyleStyleAngle = "angle"

// TextCombineUpright represent the CSS style "text-combine-upright" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-combine-upright
type TextCombineUprightStyle string

func (t TextCombineUprightStyle) Name() string {
	return "text-combine-upright"
}

// MarkerEnd represent the CSS style "marker-end" with value ""
// See
type MarkerEndStyle string

func (t MarkerEndStyle) Name() string {
	return "marker-end"
}

// BackgroundSize represent the CSS style "background-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-size
type BackgroundSizeStyle string

func (t BackgroundSizeStyle) Name() string {
	return "background-size"
}

// PaddingInlineEnd represent the CSS style "padding-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-inline-end
type PaddingInlineEndStyle string

func (t PaddingInlineEndStyle) Name() string {
	return "padding-inline-end"
}

// TextTransform represent the CSS style "text-transform" with value "none | capitalize | uppercase | lowercase | full-width | full-size-kana"
// See https://developer.mozilla.org/docs/Web/CSS/text-transform
type TextTransformStyle string

func (t TextTransformStyle) Name() string {
	return "text-transform"
}

const TextTransformStyleNone = "none"

const TextTransformStyleCapitalize = "capitalize"

const TextTransformStyleUppercase = "uppercase"

const TextTransformStyleLowercase = "lowercase"

const TextTransformStyleFullWidth = "full-width"

const TextTransformStyleFullSizeKana = "full-size-kana"

// Zoom represent the CSS style "zoom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/zoom
type ZoomStyle string

func (t ZoomStyle) Name() string {
	return "zoom"
}

// ShapeRendering represent the CSS style "shape-rendering" with value "auto | optimizeSpeed | crispEdges | geometricPrecision"
// See
type ShapeRenderingStyle string

func (t ShapeRenderingStyle) Name() string {
	return "shape-rendering"
}

const ShapeRenderingStyleAuto = "auto"

const ShapeRenderingStyleOptimizeSpeed = "optimizeSpeed"

const ShapeRenderingStyleCrispEdges = "crispEdges"

const ShapeRenderingStyleGeometricPrecision = "geometricPrecision"

// BoxFlex represent the CSS style "box-flex" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/box-flex
type BoxFlexStyle float64

func (t BoxFlexStyle) Name() string {
	return "box-flex"
}

// CaptionSide represent the CSS style "caption-side" with value "top | bottom | block-start | block-end | inline-start | inline-end"
// See https://developer.mozilla.org/docs/Web/CSS/caption-side
type CaptionSideStyle string

func (t CaptionSideStyle) Name() string {
	return "caption-side"
}

const CaptionSideStyleTop = "top"

const CaptionSideStyleBottom = "bottom"

const CaptionSideStyleBlockStart = "block-start"

const CaptionSideStyleBlockEnd = "block-end"

const CaptionSideStyleInlineStart = "inline-start"

const CaptionSideStyleInlineEnd = "inline-end"

// BorderInlineColor represent the CSS style "border-inline-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-color
type BorderInlineColorStyle Color

func (t BorderInlineColorStyle) Name() string {
	return "border-inline-color"
}

// Direction represent the CSS style "direction" with value "ltr | rtl"
// See https://developer.mozilla.org/docs/Web/CSS/direction
type DirectionStyle string

func (t DirectionStyle) Name() string {
	return "direction"
}

const DirectionStyleLtr = "ltr"

const DirectionStyleRtl = "rtl"

// MinInlineSize represent the CSS style "min-inline-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-inline-size
type MinInlineSizeStyle string

func (t MinInlineSizeStyle) Name() string {
	return "min-inline-size"
}

// ShapeMargin represent the CSS style "shape-margin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/shape-margin
type ShapeMarginStyle string

func (t ShapeMarginStyle) Name() string {
	return "shape-margin"
}

// TabSize represent the CSS style "tab-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/tab-size
type TabSizeStyle float64

func (t TabSizeStyle) Name() string {
	return "tab-size"
}

// ScrollPaddingBlockStart represent the CSS style "scroll-padding-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-block-start
type ScrollPaddingBlockStartStyle string

func (t ScrollPaddingBlockStartStyle) Name() string {
	return "scroll-padding-block-start"
}

// WordWrap represent the CSS style "word-wrap" with value "normal | break-word"
// See https://developer.mozilla.org/docs/Web/CSS/overflow-wrap
type WordWrapStyle string

func (t WordWrapStyle) Name() string {
	return "word-wrap"
}

const WordWrapStyleNormal = "normal"

const WordWrapStyleBreakWord = "break-word"

// StrokeLinecap represent the CSS style "stroke-linecap" with value "butt | round | square"
// See
type StrokeLinecapStyle string

func (t StrokeLinecapStyle) Name() string {
	return "stroke-linecap"
}

const StrokeLinecapStyleButt = "butt"

const StrokeLinecapStyleRound = "round"

const StrokeLinecapStyleSquare = "square"

// MarginBottom represent the CSS style "margin-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-bottom
type MarginBottomStyle string

func (t MarginBottomStyle) Name() string {
	return "margin-bottom"
}

// MixBlendMode represent the CSS style "mix-blend-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mix-blend-mode
type MixBlendModeStyle string

func (t MixBlendModeStyle) Name() string {
	return "mix-blend-mode"
}

// OutlineStyle represent the CSS style "outline-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-style
type OutlineStyleStyle string

func (t OutlineStyleStyle) Name() string {
	return "outline-style"
}

// OverscrollBehaviorBlock represent the CSS style "overscroll-behavior-block" with value "contain | none | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overscroll-behavior-block
type OverscrollBehaviorBlockStyle string

func (t OverscrollBehaviorBlockStyle) Name() string {
	return "overscroll-behavior-block"
}

const OverscrollBehaviorBlockStyleContain = "contain"

const OverscrollBehaviorBlockStyleNone = "none"

const OverscrollBehaviorBlockStyleAuto = "auto"

// BorderBottomRightRadius represent the CSS style "border-bottom-right-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-right-radius
type BorderBottomRightRadiusStyle string

func (t BorderBottomRightRadiusStyle) Name() string {
	return "border-bottom-right-radius"
}

// MaskBorderRepeat represent the CSS style "mask-border-repeat" with value "stretch | repeat | round | space | stretch repeat | stretch round | stretch space | repeat stretch | repeat round | repeat space | round stretch | round repeat | round space | space stretch | space repeat | space round"
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-repeat
type MaskBorderRepeatStyle string

func (t MaskBorderRepeatStyle) Name() string {
	return "mask-border-repeat"
}

const MaskBorderRepeatStyleStretch = "stretch"

const MaskBorderRepeatStyleRepeat = "repeat"

const MaskBorderRepeatStyleRound = "round"

const MaskBorderRepeatStyleSpace = "space"

const MaskBorderRepeatStyleStretchRepeat = "stretch repeat"

const MaskBorderRepeatStyleStretchRound = "stretch round"

const MaskBorderRepeatStyleStretchSpace = "stretch space"

const MaskBorderRepeatStyleRepeatStretch = "repeat stretch"

const MaskBorderRepeatStyleRepeatRound = "repeat round"

const MaskBorderRepeatStyleRepeatSpace = "repeat space"

const MaskBorderRepeatStyleRoundStretch = "round stretch"

const MaskBorderRepeatStyleRoundRepeat = "round repeat"

const MaskBorderRepeatStyleRoundSpace = "round space"

const MaskBorderRepeatStyleSpaceStretch = "space stretch"

const MaskBorderRepeatStyleSpaceRepeat = "space repeat"

const MaskBorderRepeatStyleSpaceRound = "space round"

// RubyMerge represent the CSS style "ruby-merge" with value "separate | collapse | auto"
// See
type RubyMergeStyle string

func (t RubyMergeStyle) Name() string {
	return "ruby-merge"
}

const RubyMergeStyleSeparate = "separate"

const RubyMergeStyleCollapse = "collapse"

const RubyMergeStyleAuto = "auto"

// ColorRendering represent the CSS style "color-rendering" with value "auto | optimizeSpeed | optimizeQuality"
// See
type ColorRenderingStyle string

func (t ColorRenderingStyle) Name() string {
	return "color-rendering"
}

const ColorRenderingStyleAuto = "auto"

const ColorRenderingStyleOptimizeSpeed = "optimizeSpeed"

const ColorRenderingStyleOptimizeQuality = "optimizeQuality"

// DominantBaseline represent the CSS style "dominant-baseline" with value "auto | use-script | no-change | reset-size | ideographic | alphabetic | hanging | mathematical | central | middle | text-after-edge | text-before-edge"
// See
type DominantBaselineStyle string

func (t DominantBaselineStyle) Name() string {
	return "dominant-baseline"
}

const DominantBaselineStyleAuto = "auto"

const DominantBaselineStyleUseScript = "use-script"

const DominantBaselineStyleNoChange = "no-change"

const DominantBaselineStyleResetSize = "reset-size"

const DominantBaselineStyleIdeographic = "ideographic"

const DominantBaselineStyleAlphabetic = "alphabetic"

const DominantBaselineStyleHanging = "hanging"

const DominantBaselineStyleMathematical = "mathematical"

const DominantBaselineStyleCentral = "central"

const DominantBaselineStyleMiddle = "middle"

const DominantBaselineStyleTextAfterEdge = "text-after-edge"

const DominantBaselineStyleTextBeforeEdge = "text-before-edge"

// Fill represent the CSS style "fill" with value ""
// See
type FillStyle Color

func (t FillStyle) Name() string {
	return "fill"
}

// AlignContent represent the CSS style "align-content" with value "auto | normal | stretch | end | start | flex-start | flex-end | center |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe | unsafe"
// See https://developer.mozilla.org/docs/Web/CSS/align-content
type AlignContentStyle string

func (t AlignContentStyle) Name() string {
	return "align-content"
}

const AlignContentStyleAuto = "auto"

const AlignContentStyleNormal = "normal"

const AlignContentStyleStretch = "stretch"

const AlignContentStyleEnd = "end"

const AlignContentStyleStart = "start"

const AlignContentStyleFlexStart = "flex-start"

const AlignContentStyleFlexEnd = "flex-end"

const AlignContentStyleCenter = "center"

const AlignContentStyleBaseline = "baseline"

const AlignContentStyleFirstBaseline = "first baseline"

const AlignContentStyleLastBaseline = "last baseline"

const AlignContentStyleSpaceBetween = "space-between"

const AlignContentStyleSpaceAround = "space-around"

const AlignContentStyleSpaceEvenly = "space-evenly"

const AlignContentStyleSafe = "safe"

const AlignContentStyleUnsafe = "unsafe"

// BackgroundOrigin represent the CSS style "background-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-origin
type BackgroundOriginStyle string

func (t BackgroundOriginStyle) Name() string {
	return "background-origin"
}

// ScrollbarGutter represent the CSS style "scrollbar-gutter" with value "auto |  stable | always | both | force"
// See https://developer.mozilla.org/docs/Web/CSS/scrollbar-gutter
type ScrollbarGutterStyle string

func (t ScrollbarGutterStyle) Name() string {
	return "scrollbar-gutter"
}

const ScrollbarGutterStyleAuto = "auto"

const ScrollbarGutterStyleStable = "stable"

const ScrollbarGutterStyleAlways = "always"

const ScrollbarGutterStyleBoth = "both"

const ScrollbarGutterStyleForce = "force"

// ShapeImageThreshold represent the CSS style "shape-image-threshold" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/shape-image-threshold
type ShapeImageThresholdStyle float64

func (t ShapeImageThresholdStyle) Name() string {
	return "shape-image-threshold"
}

// Translate represent the CSS style "translate" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/translate
type TranslateStyle string

func (t TranslateStyle) Name() string {
	return "translate"
}

// UnicodeBidi represent the CSS style "unicode-bidi" with value "normal | embed | isolate | bidi-override | isolate-override | plaintext"
// See https://developer.mozilla.org/docs/Web/CSS/unicode-bidi
type UnicodeBidiStyle string

func (t UnicodeBidiStyle) Name() string {
	return "unicode-bidi"
}

const UnicodeBidiStyleNormal = "normal"

const UnicodeBidiStyleEmbed = "embed"

const UnicodeBidiStyleIsolate = "isolate"

const UnicodeBidiStyleBidiOverride = "bidi-override"

const UnicodeBidiStyleIsolateOverride = "isolate-override"

const UnicodeBidiStylePlaintext = "plaintext"

// TransformOrigin represent the CSS style "transform-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform-origin
type TransformOriginStyle string

func (t TransformOriginStyle) Name() string {
	return "transform-origin"
}

// BorderRightStyle represent the CSS style "border-right-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-right-style
type BorderRightStyleStyle string

func (t BorderRightStyleStyle) Name() string {
	return "border-right-style"
}

// BorderTopWidth represent the CSS style "border-top-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-width
type BorderTopWidthStyle string

func (t BorderTopWidthStyle) Name() string {
	return "border-top-width"
}

// FontVariantNumeric represent the CSS style "font-variant-numeric" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-numeric
type FontVariantNumericStyle string

func (t FontVariantNumericStyle) Name() string {
	return "font-variant-numeric"
}

// MaskComposite represent the CSS style "mask-composite" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-composite
type MaskCompositeStyle string

func (t MaskCompositeStyle) Name() string {
	return "mask-composite"
}

// ScrollPaddingBlock represent the CSS style "scroll-padding-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-block
type ScrollPaddingBlockStyle string

func (t ScrollPaddingBlockStyle) Name() string {
	return "scroll-padding-block"
}

// BoxPack represent the CSS style "box-pack" with value "start | center | end | justify"
// See https://developer.mozilla.org/docs/Web/CSS/box-pack
type BoxPackStyle string

func (t BoxPackStyle) Name() string {
	return "box-pack"
}

const BoxPackStyleStart = "start"

const BoxPackStyleCenter = "center"

const BoxPackStyleEnd = "end"

const BoxPackStyleJustify = "justify"

// MarginLeft represent the CSS style "margin-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-left
type MarginLeftStyle string

func (t MarginLeftStyle) Name() string {
	return "margin-left"
}

// PaddingLeft represent the CSS style "padding-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-left
type PaddingLeftStyle string

func (t PaddingLeftStyle) Name() string {
	return "padding-left"
}

// PageBreakInside represent the CSS style "page-break-inside" with value "auto | avoid"
// See https://developer.mozilla.org/docs/Web/CSS/page-break-inside
type PageBreakInsideStyle string

func (t PageBreakInsideStyle) Name() string {
	return "page-break-inside"
}

const PageBreakInsideStyleAuto = "auto"

const PageBreakInsideStyleAvoid = "avoid"

// Resize represent the CSS style "resize" with value "none | both | horizontal | vertical | block | inline"
// See https://developer.mozilla.org/docs/Web/CSS/resize
type ResizeStyle string

func (t ResizeStyle) Name() string {
	return "resize"
}

const ResizeStyleNone = "none"

const ResizeStyleBoth = "both"

const ResizeStyleHorizontal = "horizontal"

const ResizeStyleVertical = "vertical"

const ResizeStyleBlock = "block"

const ResizeStyleInline = "inline"

// LineHeight represent the CSS style "line-height" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/line-height
type LineHeightStyle string

func (t LineHeightStyle) Name() string {
	return "line-height"
}

// OverscrollBehaviorX represent the CSS style "overscroll-behavior-x" with value "contain | none | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overscroll-behavior-x
type OverscrollBehaviorXStyle string

func (t OverscrollBehaviorXStyle) Name() string {
	return "overscroll-behavior-x"
}

const OverscrollBehaviorXStyleContain = "contain"

const OverscrollBehaviorXStyleNone = "none"

const OverscrollBehaviorXStyleAuto = "auto"

// BoxOrient represent the CSS style "box-orient" with value "horizontal | vertical | inline-axis | block-axis | inherit"
// See https://developer.mozilla.org/docs/Web/CSS/box-orient
type BoxOrientStyle string

func (t BoxOrientStyle) Name() string {
	return "box-orient"
}

const BoxOrientStyleHorizontal = "horizontal"

const BoxOrientStyleVertical = "vertical"

const BoxOrientStyleInlineAxis = "inline-axis"

const BoxOrientStyleBlockAxis = "block-axis"

const BoxOrientStyleInherit = "inherit"

// Filter represent the CSS style "filter" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/filter
type FilterStyle string

func (t FilterStyle) Name() string {
	return "filter"
}

// MarginTop represent the CSS style "margin-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-top
type MarginTopStyle string

func (t MarginTopStyle) Name() string {
	return "margin-top"
}

// UserSelect represent the CSS style "user-select" with value "auto | text | none | contain | all"
// See https://developer.mozilla.org/docs/Web/CSS/user-select
type UserSelectStyle string

func (t UserSelectStyle) Name() string {
	return "user-select"
}

const UserSelectStyleAuto = "auto"

const UserSelectStyleText = "text"

const UserSelectStyleNone = "none"

const UserSelectStyleContain = "contain"

const UserSelectStyleAll = "all"

// Clip represent the CSS style "clip" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/clip
type ClipStyle string

func (t ClipStyle) Name() string {
	return "clip"
}

// FontVariantEastAsian represent the CSS style "font-variant-east-asian" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-east-asian
type FontVariantEastAsianStyle string

func (t FontVariantEastAsianStyle) Name() string {
	return "font-variant-east-asian"
}

// ImageResolution represent the CSS style "image-resolution" with value ""
// See
type ImageResolutionStyle string

func (t ImageResolutionStyle) Name() string {
	return "image-resolution"
}

// ListStyleType represent the CSS style "list-style-type" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/list-style-type
type ListStyleTypeStyle URL

func (t ListStyleTypeStyle) Name() string {
	return "list-style-type"
}

// MaskOrigin represent the CSS style "mask-origin" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-origin
type MaskOriginStyle string

func (t MaskOriginStyle) Name() string {
	return "mask-origin"
}

// ColumnRuleWidth represent the CSS style "column-rule-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-rule-width
type ColumnRuleWidthStyle string

func (t ColumnRuleWidthStyle) Name() string {
	return "column-rule-width"
}

// FontSize represent the CSS style "font-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-size
type FontSizeStyle string

func (t FontSizeStyle) Name() string {
	return "font-size"
}

// FontVariant represent the CSS style "font-variant" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant
type FontVariantStyle string

func (t FontVariantStyle) Name() string {
	return "font-variant"
}

// MaskPosition represent the CSS style "mask-position" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-position
type MaskPositionStyle string

func (t MaskPositionStyle) Name() string {
	return "mask-position"
}

// VerticalAlign represent the CSS style "vertical-align" with value "baseline | sub | super | text-top | text-bottom | middle | top | bottom"
// See https://developer.mozilla.org/docs/Web/CSS/vertical-align
type VerticalAlignStyle string

func (t VerticalAlignStyle) Name() string {
	return "vertical-align"
}

const VerticalAlignStyleBaseline = "baseline"

const VerticalAlignStyleSub = "sub"

const VerticalAlignStyleSuper = "super"

const VerticalAlignStyleTextTop = "text-top"

const VerticalAlignStyleTextBottom = "text-bottom"

const VerticalAlignStyleMiddle = "middle"

const VerticalAlignStyleTop = "top"

const VerticalAlignStyleBottom = "bottom"

// TextDecorationLine represent the CSS style "text-decoration-line" with value "none | underline | overline | line-through | blink | spelling-error | grammar-error"
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-line
type TextDecorationLineStyle string

func (t TextDecorationLineStyle) Name() string {
	return "text-decoration-line"
}

const TextDecorationLineStyleNone = "none"

const TextDecorationLineStyleUnderline = "underline"

const TextDecorationLineStyleOverline = "overline"

const TextDecorationLineStyleLineThrough = "line-through"

const TextDecorationLineStyleBlink = "blink"

const TextDecorationLineStyleSpellingError = "spelling-error"

const TextDecorationLineStyleGrammarError = "grammar-error"

// AlignItems represent the CSS style "align-items" with value "auto | normal | stretch | end | start | flex-start | flex-end | center |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe | unsafe"
// See https://developer.mozilla.org/docs/Web/CSS/align-items
type AlignItemsStyle string

func (t AlignItemsStyle) Name() string {
	return "align-items"
}

const AlignItemsStyleAuto = "auto"

const AlignItemsStyleNormal = "normal"

const AlignItemsStyleStretch = "stretch"

const AlignItemsStyleEnd = "end"

const AlignItemsStyleStart = "start"

const AlignItemsStyleFlexStart = "flex-start"

const AlignItemsStyleFlexEnd = "flex-end"

const AlignItemsStyleCenter = "center"

const AlignItemsStyleBaseline = "baseline"

const AlignItemsStyleFirstBaseline = "first baseline"

const AlignItemsStyleLastBaseline = "last baseline"

const AlignItemsStyleSpaceBetween = "space-between"

const AlignItemsStyleSpaceAround = "space-around"

const AlignItemsStyleSpaceEvenly = "space-evenly"

const AlignItemsStyleSafe = "safe"

const AlignItemsStyleUnsafe = "unsafe"

// BorderBlockEndStyle represent the CSS style "border-block-end-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-end-style
type BorderBlockEndStyleStyle string

func (t BorderBlockEndStyleStyle) Name() string {
	return "border-block-end-style"
}

// LineHeightStep represent the CSS style "line-height-step" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/line-height-step
type LineHeightStepStyle float64

func (t LineHeightStepStyle) Name() string {
	return "line-height-step"
}

// ObjectPosition represent the CSS style "object-position" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/object-position
type ObjectPositionStyle string

func (t ObjectPositionStyle) Name() string {
	return "object-position"
}

// ScrollSnapType represent the CSS style "scroll-snap-type" with value "none |  x | y | block | inline | both | mandatory | proximity"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-type
type ScrollSnapTypeStyle string

func (t ScrollSnapTypeStyle) Name() string {
	return "scroll-snap-type"
}

const ScrollSnapTypeStyleNone = "none"

const ScrollSnapTypeStyleX = "x"

const ScrollSnapTypeStyleY = "y"

const ScrollSnapTypeStyleBlock = "block"

const ScrollSnapTypeStyleInline = "inline"

const ScrollSnapTypeStyleBoth = "both"

const ScrollSnapTypeStyleMandatory = "mandatory"

const ScrollSnapTypeStyleProximity = "proximity"

// GridTemplateColumns represent the CSS style "grid-template-columns" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-template-columns
type GridTemplateColumnsStyle string

func (t GridTemplateColumnsStyle) Name() string {
	return "grid-template-columns"
}

// Orphans represent the CSS style "orphans" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/orphans
type OrphansStyle float64

func (t OrphansStyle) Name() string {
	return "orphans"
}

// ScrollSnapPointsY represent the CSS style "scroll-snap-points-y" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-points-y
type ScrollSnapPointsYStyle string

func (t ScrollSnapPointsYStyle) Name() string {
	return "scroll-snap-points-y"
}

// VectorEffect represent the CSS style "vector-effect" with value "non-scaling-stroke | none"
// See
type VectorEffectStyle string

func (t VectorEffectStyle) Name() string {
	return "vector-effect"
}

const VectorEffectStyleNonScalingStroke = "non-scaling-stroke"

const VectorEffectStyleNone = "none"

// BackfaceVisibility represent the CSS style "backface-visibility" with value "visible | hidden"
// See https://developer.mozilla.org/docs/Web/CSS/backface-visibility
type BackfaceVisibilityStyle string

func (t BackfaceVisibilityStyle) Name() string {
	return "backface-visibility"
}

const BackfaceVisibilityStyleVisible = "visible"

const BackfaceVisibilityStyleHidden = "hidden"

// BorderImageSlice represent the CSS style "border-image-slice" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-slice
type BorderImageSliceStyle string

func (t BorderImageSliceStyle) Name() string {
	return "border-image-slice"
}

// BorderInlineStartWidth represent the CSS style "border-inline-start-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-start-width
type BorderInlineStartWidthStyle string

func (t BorderInlineStartWidthStyle) Name() string {
	return "border-inline-start-width"
}

// FlexBasis represent the CSS style "flex-basis" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/flex-basis
type FlexBasisStyle string

func (t FlexBasisStyle) Name() string {
	return "flex-basis"
}

// JustifyItems represent the CSS style "justify-items" with value "auto | normal | stretch | end | start | flex-start | flex-end | center | left | right |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe center | unsafe center"
// See https://developer.mozilla.org/docs/Web/CSS/justify-items
type JustifyItemsStyle string

func (t JustifyItemsStyle) Name() string {
	return "justify-items"
}

const JustifyItemsStyleAuto = "auto"

const JustifyItemsStyleNormal = "normal"

const JustifyItemsStyleStretch = "stretch"

const JustifyItemsStyleEnd = "end"

const JustifyItemsStyleStart = "start"

const JustifyItemsStyleFlexStart = "flex-start"

const JustifyItemsStyleFlexEnd = "flex-end"

const JustifyItemsStyleCenter = "center"

const JustifyItemsStyleLeft = "left"

const JustifyItemsStyleRight = "right"

const JustifyItemsStyleBaseline = "baseline"

const JustifyItemsStyleFirstBaseline = "first baseline"

const JustifyItemsStyleLastBaseline = "last baseline"

const JustifyItemsStyleSpaceBetween = "space-between"

const JustifyItemsStyleSpaceAround = "space-around"

const JustifyItemsStyleSpaceEvenly = "space-evenly"

const JustifyItemsStyleSafeCenter = "safe center"

const JustifyItemsStyleUnsafeCenter = "unsafe center"

// MaxInlineSize represent the CSS style "max-inline-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-inline-size
type MaxInlineSizeStyle string

func (t MaxInlineSizeStyle) Name() string {
	return "max-inline-size"
}

// AnimationFillMode represent the CSS style "animation-fill-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-fill-mode
type AnimationFillModeStyle string

func (t AnimationFillModeStyle) Name() string {
	return "animation-fill-mode"
}

// AnimationTimingFunction represent the CSS style "animation-timing-function" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-timing-function
type AnimationTimingFunctionStyle string

func (t AnimationTimingFunctionStyle) Name() string {
	return "animation-timing-function"
}

// BorderStartEndRadius represent the CSS style "border-start-end-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-start-end-radius
type BorderStartEndRadiusStyle string

func (t BorderStartEndRadiusStyle) Name() string {
	return "border-start-end-radius"
}

// BorderTopRightRadius represent the CSS style "border-top-right-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-right-radius
type BorderTopRightRadiusStyle string

func (t BorderTopRightRadiusStyle) Name() string {
	return "border-top-right-radius"
}

// GridRowGap represent the CSS style "grid-row-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/row-gap
type GridRowGapStyle string

func (t GridRowGapStyle) Name() string {
	return "grid-row-gap"
}

// ColumnRuleColor represent the CSS style "column-rule-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-rule-color
type ColumnRuleColorStyle Color

func (t ColumnRuleColorStyle) Name() string {
	return "column-rule-color"
}

// InsetBlock represent the CSS style "inset-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-block
type InsetBlockStyle string

func (t InsetBlockStyle) Name() string {
	return "inset-block"
}

// InsetInlineStart represent the CSS style "inset-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-inline-start
type InsetInlineStartStyle string

func (t InsetInlineStartStyle) Name() string {
	return "inset-inline-start"
}

// All represent the CSS style "all" with value "initial | inherit | unset | revert"
// See https://developer.mozilla.org/docs/Web/CSS/all
type AllStyle string

func (t AllStyle) Name() string {
	return "all"
}

const AllStyleInitial = "initial"

const AllStyleInherit = "inherit"

const AllStyleUnset = "unset"

const AllStyleRevert = "revert"

// BorderEndStartRadius represent the CSS style "border-end-start-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-end-start-radius
type BorderEndStartRadiusStyle string

func (t BorderEndStartRadiusStyle) Name() string {
	return "border-end-start-radius"
}

// BorderLeftStyle represent the CSS style "border-left-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-left-style
type BorderLeftStyleStyle string

func (t BorderLeftStyleStyle) Name() string {
	return "border-left-style"
}

// BoxDecorationBreak represent the CSS style "box-decoration-break" with value "slice | clone"
// See https://developer.mozilla.org/docs/Web/CSS/box-decoration-break
type BoxDecorationBreakStyle string

func (t BoxDecorationBreakStyle) Name() string {
	return "box-decoration-break"
}

const BoxDecorationBreakStyleSlice = "slice"

const BoxDecorationBreakStyleClone = "clone"

// BreakInside represent the CSS style "break-inside" with value "auto | avoid | avoid-page | avoid-column | avoid-region"
// See https://developer.mozilla.org/docs/Web/CSS/break-inside
type BreakInsideStyle string

func (t BreakInsideStyle) Name() string {
	return "break-inside"
}

const BreakInsideStyleAuto = "auto"

const BreakInsideStyleAvoid = "avoid"

const BreakInsideStyleAvoidPage = "avoid-page"

const BreakInsideStyleAvoidColumn = "avoid-column"

const BreakInsideStyleAvoidRegion = "avoid-region"

// MaskBorderSource represent the CSS style "mask-border-source" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-source
type MaskBorderSourceStyle string

func (t MaskBorderSourceStyle) Name() string {
	return "mask-border-source"
}

// MaskMode represent the CSS style "mask-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-mode
type MaskModeStyle string

func (t MaskModeStyle) Name() string {
	return "mask-mode"
}

// MasonryAutoFlow represent the CSS style "masonry-auto-flow" with value "pack | next | pack definite-first | next definite-first | pack ordered | next ordered"
// See https://developer.mozilla.org/docs/Web/CSS/masonry-auto-flow
type MasonryAutoFlowStyle string

func (t MasonryAutoFlowStyle) Name() string {
	return "masonry-auto-flow"
}

const MasonryAutoFlowStylePack = "pack"

const MasonryAutoFlowStyleNext = "next"

const MasonryAutoFlowStylePackDefiniteFirst = "pack definite-first"

const MasonryAutoFlowStyleNextDefiniteFirst = "next definite-first"

const MasonryAutoFlowStylePackOrdered = "pack ordered"

const MasonryAutoFlowStyleNextOrdered = "next ordered"

// OverflowInline represent the CSS style "overflow-inline" with value "visible | hidden | clip | scroll | auto"
// See
type OverflowInlineStyle string

func (t OverflowInlineStyle) Name() string {
	return "overflow-inline"
}

const OverflowInlineStyleVisible = "visible"

const OverflowInlineStyleHidden = "hidden"

const OverflowInlineStyleClip = "clip"

const OverflowInlineStyleScroll = "scroll"

const OverflowInlineStyleAuto = "auto"

// PaddingBlockStart represent the CSS style "padding-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-block-start
type PaddingBlockStartStyle string

func (t PaddingBlockStartStyle) Name() string {
	return "padding-block-start"
}

// ScrollbarColor represent the CSS style "scrollbar-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scrollbar-color
type ScrollbarColorStyle Color

func (t ScrollbarColorStyle) Name() string {
	return "scrollbar-color"
}

// ScrollPaddingInline represent the CSS style "scroll-padding-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-inline
type ScrollPaddingInlineStyle string

func (t ScrollPaddingInlineStyle) Name() string {
	return "scroll-padding-inline"
}

// TextDecorationThickness represent the CSS style "text-decoration-thickness" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-thickness
type TextDecorationThicknessStyle string

func (t TextDecorationThicknessStyle) Name() string {
	return "text-decoration-thickness"
}

// ClipRule represent the CSS style "clip-rule" with value "nonzero | evenodd"
// See
type ClipRuleStyle string

func (t ClipRuleStyle) Name() string {
	return "clip-rule"
}

const ClipRuleStyleNonzero = "nonzero"

const ClipRuleStyleEvenodd = "evenodd"

// TextAnchor represent the CSS style "text-anchor" with value "start | middle | end"
// See
type TextAnchorStyle string

func (t TextAnchorStyle) Name() string {
	return "text-anchor"
}

const TextAnchorStyleStart = "start"

const TextAnchorStyleMiddle = "middle"

const TextAnchorStyleEnd = "end"

// BlockOverflow represent the CSS style "block-overflow" with value "clip | ellipsis | string"
// See
type BlockOverflowStyle string

func (t BlockOverflowStyle) Name() string {
	return "block-overflow"
}

const BlockOverflowStyleClip = "clip"

const BlockOverflowStyleEllipsis = "ellipsis"

const BlockOverflowStyleString = "string"

// Hyphens represent the CSS style "hyphens" with value "none | manual | auto"
// See https://developer.mozilla.org/docs/Web/CSS/hyphens
type HyphensStyle string

func (t HyphensStyle) Name() string {
	return "hyphens"
}

const HyphensStyleNone = "none"

const HyphensStyleManual = "manual"

const HyphensStyleAuto = "auto"

// MaxWidth represent the CSS style "max-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-width
type MaxWidthStyle string

func (t MaxWidthStyle) Name() string {
	return "max-width"
}

// Opacity represent the CSS style "opacity" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/opacity
type OpacityStyle float64

func (t OpacityStyle) Name() string {
	return "opacity"
}

// OverflowWrap represent the CSS style "overflow-wrap" with value "normal | break-word | anywhere"
// See https://developer.mozilla.org/docs/Web/CSS/overflow-wrap
type OverflowWrapStyle string

func (t OverflowWrapStyle) Name() string {
	return "overflow-wrap"
}

const OverflowWrapStyleNormal = "normal"

const OverflowWrapStyleBreakWord = "break-word"

const OverflowWrapStyleAnywhere = "anywhere"

// ScrollMarginBottom represent the CSS style "scroll-margin-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-bottom
type ScrollMarginBottomStyle float64

func (t ScrollMarginBottomStyle) Name() string {
	return "scroll-margin-bottom"
}

// FloodColor represent the CSS style "flood-color" with value ""
// See
type FloodColorStyle Color

func (t FloodColorStyle) Name() string {
	return "flood-color"
}

// BoxSizing represent the CSS style "box-sizing" with value "content-box | border-box"
// See https://developer.mozilla.org/docs/Web/CSS/box-sizing
type BoxSizingStyle string

func (t BoxSizingStyle) Name() string {
	return "box-sizing"
}

const BoxSizingStyleContentBox = "content-box"

const BoxSizingStyleBorderBox = "border-box"

// GridAutoColumns represent the CSS style "grid-auto-columns" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-auto-columns
type GridAutoColumnsStyle string

func (t GridAutoColumnsStyle) Name() string {
	return "grid-auto-columns"
}

// LineClamp represent the CSS style "line-clamp" with value ""
// See
type LineClampStyle string

func (t LineClampStyle) Name() string {
	return "line-clamp"
}

// MinBlockSize represent the CSS style "min-block-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-block-size
type MinBlockSizeStyle string

func (t MinBlockSizeStyle) Name() string {
	return "min-block-size"
}

// FlexWrap represent the CSS style "flex-wrap" with value "nowrap | wrap | wrap-reverse"
// See https://developer.mozilla.org/docs/Web/CSS/flex-wrap
type FlexWrapStyle string

func (t FlexWrapStyle) Name() string {
	return "flex-wrap"
}

const FlexWrapStyleNowrap = "nowrap"

const FlexWrapStyleWrap = "wrap"

const FlexWrapStyleWrapReverse = "wrap-reverse"

// GridColumnStart represent the CSS style "grid-column-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-column-start
type GridColumnStartStyle string

func (t GridColumnStartStyle) Name() string {
	return "grid-column-start"
}

// InsetInlineEnd represent the CSS style "inset-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-inline-end
type InsetInlineEndStyle string

func (t InsetInlineEndStyle) Name() string {
	return "inset-inline-end"
}

// OutlineOffset represent the CSS style "outline-offset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-offset
type OutlineOffsetStyle float64

func (t OutlineOffsetStyle) Name() string {
	return "outline-offset"
}

// MarginTrim represent the CSS style "margin-trim" with value "none | in-flow | all"
// See https://developer.mozilla.org/docs/Web/CSS/margin-trim
type MarginTrimStyle string

func (t MarginTrimStyle) Name() string {
	return "margin-trim"
}

const MarginTrimStyleNone = "none"

const MarginTrimStyleInFlow = "in-flow"

const MarginTrimStyleAll = "all"

// ScrollPaddingInlineStart represent the CSS style "scroll-padding-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-inline-start
type ScrollPaddingInlineStartStyle string

func (t ScrollPaddingInlineStartStyle) Name() string {
	return "scroll-padding-inline-start"
}

// GlyphOrientationVertical represent the CSS style "glyph-orientation-vertical" with value ""
// See
type GlyphOrientationVerticalStyle string

func (t GlyphOrientationVerticalStyle) Name() string {
	return "glyph-orientation-vertical"
}

// StopOpacity represent the CSS style "stop-opacity" with value ""
// See
type StopOpacityStyle float64

func (t StopOpacityStyle) Name() string {
	return "stop-opacity"
}

// BackgroundClip represent the CSS style "background-clip" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-clip
type BackgroundClipStyle string

func (t BackgroundClipStyle) Name() string {
	return "background-clip"
}

// BorderInlineStartStyle represent the CSS style "border-inline-start-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-start-style
type BorderInlineStartStyleStyle string

func (t BorderInlineStartStyleStyle) Name() string {
	return "border-inline-start-style"
}

// CounterIncrement represent the CSS style "counter-increment" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/counter-increment
type CounterIncrementStyle string

func (t CounterIncrementStyle) Name() string {
	return "counter-increment"
}

// TextJustify represent the CSS style "text-justify" with value "auto | inter-character | inter-word | none"
// See https://developer.mozilla.org/docs/Web/CSS/text-justify
type TextJustifyStyle string

func (t TextJustifyStyle) Name() string {
	return "text-justify"
}

const TextJustifyStyleAuto = "auto"

const TextJustifyStyleInterCharacter = "inter-character"

const TextJustifyStyleInterWord = "inter-word"

const TextJustifyStyleNone = "none"

// TextUnderlinePosition represent the CSS style "text-underline-position" with value "auto | from-font | under | left | right"
// See https://developer.mozilla.org/docs/Web/CSS/text-underline-position
type TextUnderlinePositionStyle string

func (t TextUnderlinePositionStyle) Name() string {
	return "text-underline-position"
}

const TextUnderlinePositionStyleAuto = "auto"

const TextUnderlinePositionStyleFromFont = "from-font"

const TextUnderlinePositionStyleUnder = "under"

const TextUnderlinePositionStyleLeft = "left"

const TextUnderlinePositionStyleRight = "right"

// BorderRightWidth represent the CSS style "border-right-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-right-width
type BorderRightWidthStyle string

func (t BorderRightWidthStyle) Name() string {
	return "border-right-width"
}

// Bottom represent the CSS style "bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/bottom
type BottomStyle string

func (t BottomStyle) Name() string {
	return "bottom"
}

// ForcedColorAdjust represent the CSS style "forced-color-adjust" with value "auto | none"
// See https://developer.mozilla.org/docs/Web/CSS/forced-color-adjust
type ForcedColorAdjustStyle string

func (t ForcedColorAdjustStyle) Name() string {
	return "forced-color-adjust"
}

const ForcedColorAdjustStyleAuto = "auto"

const ForcedColorAdjustStyleNone = "none"

// OutlineColor represent the CSS style "outline-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-color
type OutlineColorStyle Color

func (t OutlineColorStyle) Name() string {
	return "outline-color"
}

// ScrollMarginBlockEnd represent the CSS style "scroll-margin-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-block-end
type ScrollMarginBlockEndStyle float64

func (t ScrollMarginBlockEndStyle) Name() string {
	return "scroll-margin-block-end"
}

// ScrollPaddingInlineEnd represent the CSS style "scroll-padding-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-inline-end
type ScrollPaddingInlineEndStyle string

func (t ScrollPaddingInlineEndStyle) Name() string {
	return "scroll-padding-inline-end"
}

// TextDecorationSkipInk represent the CSS style "text-decoration-skip-ink" with value "auto | all | none"
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-skip-ink
type TextDecorationSkipInkStyle string

func (t TextDecorationSkipInkStyle) Name() string {
	return "text-decoration-skip-ink"
}

const TextDecorationSkipInkStyleAuto = "auto"

const TextDecorationSkipInkStyleAll = "all"

const TextDecorationSkipInkStyleNone = "none"

// StrokeLinejoin represent the CSS style "stroke-linejoin" with value "miter | round | bevel"
// See
type StrokeLinejoinStyle string

func (t StrokeLinejoinStyle) Name() string {
	return "stroke-linejoin"
}

const StrokeLinejoinStyleMiter = "miter"

const StrokeLinejoinStyleRound = "round"

const StrokeLinejoinStyleBevel = "bevel"

// AlignTracks represent the CSS style "align-tracks" with value "normal | end | start | stretch | flex-start | flex-end | center |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe | unsafe"
// See https://developer.mozilla.org/docs/Web/CSS/align-tracks
type AlignTracksStyle string

func (t AlignTracksStyle) Name() string {
	return "align-tracks"
}

const AlignTracksStyleNormal = "normal"

const AlignTracksStyleEnd = "end"

const AlignTracksStyleStart = "start"

const AlignTracksStyleStretch = "stretch"

const AlignTracksStyleFlexStart = "flex-start"

const AlignTracksStyleFlexEnd = "flex-end"

const AlignTracksStyleCenter = "center"

const AlignTracksStyleBaseline = "baseline"

const AlignTracksStyleFirstBaseline = "first baseline"

const AlignTracksStyleLastBaseline = "last baseline"

const AlignTracksStyleSpaceBetween = "space-between"

const AlignTracksStyleSpaceAround = "space-around"

const AlignTracksStyleSpaceEvenly = "space-evenly"

const AlignTracksStyleSafe = "safe"

const AlignTracksStyleUnsafe = "unsafe"

// ColumnSpan represent the CSS style "column-span" with value "none | all"
// See https://developer.mozilla.org/docs/Web/CSS/column-span
type ColumnSpanStyle string

func (t ColumnSpanStyle) Name() string {
	return "column-span"
}

const ColumnSpanStyleNone = "none"

const ColumnSpanStyleAll = "all"

// InsetBlockEnd represent the CSS style "inset-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-block-end
type InsetBlockEndStyle string

func (t InsetBlockEndStyle) Name() string {
	return "inset-block-end"
}

// Right represent the CSS style "right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/right
type RightStyle string

func (t RightStyle) Name() string {
	return "right"
}

// TextEmphasisPositionFirst represent the CSS style "text-emphasis-position-first" with value " over | under"
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-position
type TextEmphasisPositionFirstStyle string

func (t TextEmphasisPositionFirstStyle) Name() string {
	return "text-emphasis-position-first"
}

const TextEmphasisPositionFirstStyleOver = "over"

const TextEmphasisPositionFirstStyleUnder = "under"

// BorderBlockStyle represent the CSS style "border-block-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-style
type BorderBlockStyleStyle string

func (t BorderBlockStyleStyle) Name() string {
	return "border-block-style"
}

// BoxAlign represent the CSS style "box-align" with value "start | center | end | baseline | stretch"
// See https://developer.mozilla.org/docs/Web/CSS/box-align
type BoxAlignStyle string

func (t BoxAlignStyle) Name() string {
	return "box-align"
}

const BoxAlignStyleStart = "start"

const BoxAlignStyleCenter = "center"

const BoxAlignStyleEnd = "end"

const BoxAlignStyleBaseline = "baseline"

const BoxAlignStyleStretch = "stretch"

// FontWeight represent the CSS style "font-weight" with value "100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900 | bolder | lighter"
// See https://developer.mozilla.org/docs/Web/CSS/font-weight
type FontWeightStyle string

func (t FontWeightStyle) Name() string {
	return "font-weight"
}

const FontWeightStyle100 = "100"

const FontWeightStyle200 = "200"

const FontWeightStyle300 = "300"

const FontWeightStyle400 = "400"

const FontWeightStyle500 = "500"

const FontWeightStyle600 = "600"

const FontWeightStyle700 = "700"

const FontWeightStyle800 = "800"

const FontWeightStyle900 = "900"

const FontWeightStyleBolder = "bolder"

const FontWeightStyleLighter = "lighter"

// ScrollSnapPointsX represent the CSS style "scroll-snap-points-x" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-points-x
type ScrollSnapPointsXStyle string

func (t ScrollSnapPointsXStyle) Name() string {
	return "scroll-snap-points-x"
}

// TextAlignLast represent the CSS style "text-align-last" with value "auto | start | end | left | right | center | justify"
// See https://developer.mozilla.org/docs/Web/CSS/text-align-last
type TextAlignLastStyle string

func (t TextAlignLastStyle) Name() string {
	return "text-align-last"
}

const TextAlignLastStyleAuto = "auto"

const TextAlignLastStyleStart = "start"

const TextAlignLastStyleEnd = "end"

const TextAlignLastStyleLeft = "left"

const TextAlignLastStyleRight = "right"

const TextAlignLastStyleCenter = "center"

const TextAlignLastStyleJustify = "justify"

// Display represent the CSS style "display" with value "contents | none | inline-block | inline-table | inline-flex| inline-grid | | block | inline | run-in | flow | flow-root | table | flex | grid | ruby | list-item | table-row-group | table-header-group | table-footer-group | table-row | table-cell | table-column-group | table-column | table-caption | ruby-base | ruby-text | ruby-base-container | ruby-text-container"
// See https://developer.mozilla.org/docs/Web/CSS/display
type DisplayStyle string

func (t DisplayStyle) Name() string {
	return "display"
}

const DisplayStyleContents = "contents"

const DisplayStyleNone = "none"

const DisplayStyleInlineBlock = "inline-block"

const DisplayStyleInlineTable = "inline-table"

const DisplayStyleInlineFlex = "inline-flex"

const DisplayStyleInlineGrid = "inline-grid"

const DisplayStyleBlock = "block"

const DisplayStyleInline = "inline"

const DisplayStyleRunIn = "run-in"

const DisplayStyleFlow = "flow"

const DisplayStyleFlowRoot = "flow-root"

const DisplayStyleTable = "table"

const DisplayStyleFlex = "flex"

const DisplayStyleGrid = "grid"

const DisplayStyleRuby = "ruby"

const DisplayStyleListItem = "list-item"

const DisplayStyleTableRowGroup = "table-row-group"

const DisplayStyleTableHeaderGroup = "table-header-group"

const DisplayStyleTableFooterGroup = "table-footer-group"

const DisplayStyleTableRow = "table-row"

const DisplayStyleTableCell = "table-cell"

const DisplayStyleTableColumnGroup = "table-column-group"

const DisplayStyleTableColumn = "table-column"

const DisplayStyleTableCaption = "table-caption"

const DisplayStyleRubyBase = "ruby-base"

const DisplayStyleRubyText = "ruby-text"

const DisplayStyleRubyBaseContainer = "ruby-base-container"

const DisplayStyleRubyTextContainer = "ruby-text-container"

// Inset represent the CSS style "inset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset
type InsetStyle string

func (t InsetStyle) Name() string {
	return "inset"
}

// RubyPosition represent the CSS style "ruby-position" with value "over | under | inter-character"
// See https://developer.mozilla.org/docs/Web/CSS/ruby-position
type RubyPositionStyle string

func (t RubyPositionStyle) Name() string {
	return "ruby-position"
}

const RubyPositionStyleOver = "over"

const RubyPositionStyleUnder = "under"

const RubyPositionStyleInterCharacter = "inter-character"

// TableLayout represent the CSS style "table-layout" with value "auto | fixed"
// See https://developer.mozilla.org/docs/Web/CSS/table-layout
type TableLayoutStyle string

func (t TableLayoutStyle) Name() string {
	return "table-layout"
}

const TableLayoutStyleAuto = "auto"

const TableLayoutStyleFixed = "fixed"

// WordSpacing represent the CSS style "word-spacing" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/word-spacing
type WordSpacingStyle string

func (t WordSpacingStyle) Name() string {
	return "word-spacing"
}

// MaskBorderWidth represent the CSS style "mask-border-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-width
type MaskBorderWidthStyle string

func (t MaskBorderWidthStyle) Name() string {
	return "mask-border-width"
}

// PageBreakAfter represent the CSS style "page-break-after" with value "auto | always | avoid | left | right | recto | verso"
// See https://developer.mozilla.org/docs/Web/CSS/page-break-after
type PageBreakAfterStyle string

func (t PageBreakAfterStyle) Name() string {
	return "page-break-after"
}

const PageBreakAfterStyleAuto = "auto"

const PageBreakAfterStyleAlways = "always"

const PageBreakAfterStyleAvoid = "avoid"

const PageBreakAfterStyleLeft = "left"

const PageBreakAfterStyleRight = "right"

const PageBreakAfterStyleRecto = "recto"

const PageBreakAfterStyleVerso = "verso"

// RubyAlign represent the CSS style "ruby-align" with value "start | center | space-between | space-around"
// See https://developer.mozilla.org/docs/Web/CSS/ruby-align
type RubyAlignStyle string

func (t RubyAlignStyle) Name() string {
	return "ruby-align"
}

const RubyAlignStyleStart = "start"

const RubyAlignStyleCenter = "center"

const RubyAlignStyleSpaceBetween = "space-between"

const RubyAlignStyleSpaceAround = "space-around"

// AnimationName represent the CSS style "animation-name" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-name
type AnimationNameStyle string

func (t AnimationNameStyle) Name() string {
	return "animation-name"
}

// Azimuth represent the CSS style "azimuth" with value "left-side | far-left | left | center-left | center | center-right | right | far-right | right-side | behind | leftwards | rightwards"
// See https://developer.mozilla.org/docs/Web/CSS/azimuth
type AzimuthStyle string

func (t AzimuthStyle) Name() string {
	return "azimuth"
}

const AzimuthStyleLeftSide = "left-side"

const AzimuthStyleFarLeft = "far-left"

const AzimuthStyleLeft = "left"

const AzimuthStyleCenterLeft = "center-left"

const AzimuthStyleCenter = "center"

const AzimuthStyleCenterRight = "center-right"

const AzimuthStyleRight = "right"

const AzimuthStyleFarRight = "far-right"

const AzimuthStyleRightSide = "right-side"

const AzimuthStyleBehind = "behind"

const AzimuthStyleLeftwards = "leftwards"

const AzimuthStyleRightwards = "rightwards"

// BlockSize represent the CSS style "block-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/block-size
type BlockSizeStyle string

func (t BlockSizeStyle) Name() string {
	return "block-size"
}

// BorderSpacing represent the CSS style "border-spacing" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-spacing
type BorderSpacingStyle string

func (t BorderSpacingStyle) Name() string {
	return "border-spacing"
}

// CaretColor represent the CSS style "caret-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/caret-color
type CaretColorStyle Color

func (t CaretColorStyle) Name() string {
	return "caret-color"
}

// ScrollSnapStop represent the CSS style "scroll-snap-stop" with value "normal | always"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-stop
type ScrollSnapStopStyle string

func (t ScrollSnapStopStyle) Name() string {
	return "scroll-snap-stop"
}

const ScrollSnapStopStyleNormal = "normal"

const ScrollSnapStopStyleAlways = "always"

// ScrollSnapDestination represent the CSS style "scroll-snap-destination" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-destination
type ScrollSnapDestinationStyle string

func (t ScrollSnapDestinationStyle) Name() string {
	return "scroll-snap-destination"
}

// TransitionProperty represent the CSS style "transition-property" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-property
type TransitionPropertyStyle string

func (t TransitionPropertyStyle) Name() string {
	return "transition-property"
}

// ClipPath represent the CSS style "clip-path" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/clip-path
type ClipPathStyle string

func (t ClipPathStyle) Name() string {
	return "clip-path"
}

// MarginInlineEnd represent the CSS style "margin-inline-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-inline-end
type MarginInlineEndStyle string

func (t MarginInlineEndStyle) Name() string {
	return "margin-inline-end"
}

// ObjectFit represent the CSS style "object-fit" with value "fill | contain | cover | none | scale-down"
// See https://developer.mozilla.org/docs/Web/CSS/object-fit
type ObjectFitStyle string

func (t ObjectFitStyle) Name() string {
	return "object-fit"
}

const ObjectFitStyleFill = "fill"

const ObjectFitStyleContain = "contain"

const ObjectFitStyleCover = "cover"

const ObjectFitStyleNone = "none"

const ObjectFitStyleScaleDown = "scale-down"

// PointerEvents represent the CSS style "pointer-events" with value "auto | none | visiblePainted | visibleFill | visibleStroke | visible | painted | fill | stroke | all | inherit"
// See https://developer.mozilla.org/docs/Web/CSS/pointer-events
type PointerEventsStyle string

func (t PointerEventsStyle) Name() string {
	return "pointer-events"
}

const PointerEventsStyleAuto = "auto"

const PointerEventsStyleNone = "none"

const PointerEventsStyleVisiblePainted = "visiblePainted"

const PointerEventsStyleVisibleFill = "visibleFill"

const PointerEventsStyleVisibleStroke = "visibleStroke"

const PointerEventsStyleVisible = "visible"

const PointerEventsStylePainted = "painted"

const PointerEventsStyleFill = "fill"

const PointerEventsStyleStroke = "stroke"

const PointerEventsStyleAll = "all"

const PointerEventsStyleInherit = "inherit"

// ScrollMarginLeft represent the CSS style "scroll-margin-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-left
type ScrollMarginLeftStyle float64

func (t ScrollMarginLeftStyle) Name() string {
	return "scroll-margin-left"
}

// TextEmphasisColor represent the CSS style "text-emphasis-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-color
type TextEmphasisColorStyle Color

func (t TextEmphasisColorStyle) Name() string {
	return "text-emphasis-color"
}

// TextRendering represent the CSS style "text-rendering" with value "auto | optimizeSpeed | optimizeLegibility | geometricPrecision"
// See https://developer.mozilla.org/docs/Web/CSS/text-rendering
type TextRenderingStyle string

func (t TextRenderingStyle) Name() string {
	return "text-rendering"
}

const TextRenderingStyleAuto = "auto"

const TextRenderingStyleOptimizeSpeed = "optimizeSpeed"

const TextRenderingStyleOptimizeLegibility = "optimizeLegibility"

const TextRenderingStyleGeometricPrecision = "geometricPrecision"

// TextSizeAdjust represent the CSS style "text-size-adjust" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-size-adjust
type TextSizeAdjustStyle string

func (t TextSizeAdjustStyle) Name() string {
	return "text-size-adjust"
}

// Appearance represent the CSS style "appearance" with value "none | auto | textfield | menulist-button | compat-auto"
// See https://developer.mozilla.org/docs/Web/CSS/appearance
type AppearanceStyle string

func (t AppearanceStyle) Name() string {
	return "appearance"
}

const AppearanceStyleNone = "none"

const AppearanceStyleAuto = "auto"

const AppearanceStyleTextfield = "textfield"

const AppearanceStyleMenulistButton = "menulist-button"

const AppearanceStyleCompatAuto = "compat-auto"

// BorderBlockStartColor represent the CSS style "border-block-start-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-start-color
type BorderBlockStartColorStyle Color

func (t BorderBlockStartColorStyle) Name() string {
	return "border-block-start-color"
}

// BreakAfter represent the CSS style "break-after" with value "auto | avoid | always | all | avoid-page | page | left | right | recto | verso | avoid-column | column | avoid-region | region"
// See https://developer.mozilla.org/docs/Web/CSS/break-after
type BreakAfterStyle string

func (t BreakAfterStyle) Name() string {
	return "break-after"
}

const BreakAfterStyleAuto = "auto"

const BreakAfterStyleAvoid = "avoid"

const BreakAfterStyleAlways = "always"

const BreakAfterStyleAll = "all"

const BreakAfterStyleAvoidPage = "avoid-page"

const BreakAfterStylePage = "page"

const BreakAfterStyleLeft = "left"

const BreakAfterStyleRight = "right"

const BreakAfterStyleRecto = "recto"

const BreakAfterStyleVerso = "verso"

const BreakAfterStyleAvoidColumn = "avoid-column"

const BreakAfterStyleColumn = "column"

const BreakAfterStyleAvoidRegion = "avoid-region"

const BreakAfterStyleRegion = "region"

// MarginBlock represent the CSS style "margin-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-block
type MarginBlockStyle string

func (t MarginBlockStyle) Name() string {
	return "margin-block"
}

// PaintOrder represent the CSS style "paint-order" with value "normal | fill | stroke | markers"
// See https://developer.mozilla.org/docs/Web/CSS/paint-order
type PaintOrderStyle string

func (t PaintOrderStyle) Name() string {
	return "paint-order"
}

const PaintOrderStyleNormal = "normal"

const PaintOrderStyleFill = "fill"

const PaintOrderStyleStroke = "stroke"

const PaintOrderStyleMarkers = "markers"

// Transform represent the CSS style "transform" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transform
type TransformStyle string

func (t TransformStyle) Name() string {
	return "transform"
}

// BorderInlineEndStyle represent the CSS style "border-inline-end-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-end-style
type BorderInlineEndStyleStyle string

func (t BorderInlineEndStyleStyle) Name() string {
	return "border-inline-end-style"
}

// BoxShadow represent the CSS style "box-shadow" with value "string"
// See https://developer.mozilla.org/docs/Web/CSS/box-shadow
type BoxShadowStyle string

func (t BoxShadowStyle) Name() string {
	return "box-shadow"
}

const BoxShadowStyleString = "string"

// MarginRight represent the CSS style "margin-right" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-right
type MarginRightStyle string

func (t MarginRightStyle) Name() string {
	return "margin-right"
}

// MaskBorderMode represent the CSS style "mask-border-mode" with value "luminance | alpha"
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-mode
type MaskBorderModeStyle string

func (t MaskBorderModeStyle) Name() string {
	return "mask-border-mode"
}

const MaskBorderModeStyleLuminance = "luminance"

const MaskBorderModeStyleAlpha = "alpha"

// Position represent the CSS style "position" with value "static | relative | absolute | sticky | fixed"
// See https://developer.mozilla.org/docs/Web/CSS/position
type PositionStyle string

func (t PositionStyle) Name() string {
	return "position"
}

const PositionStyleStatic = "static"

const PositionStyleRelative = "relative"

const PositionStyleAbsolute = "absolute"

const PositionStyleSticky = "sticky"

const PositionStyleFixed = "fixed"

// BackgroundImage represent the CSS style "background-image" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-image
type BackgroundImageStyle URL

func (t BackgroundImageStyle) Name() string {
	return "background-image"
}

// BorderBottomStyle represent the CSS style "border-bottom-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-style
type BorderBottomStyleStyle string

func (t BorderBottomStyleStyle) Name() string {
	return "border-bottom-style"
}

// BoxFlexGroup represent the CSS style "box-flex-group" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/box-flex-group
type BoxFlexGroupStyle float64

func (t BoxFlexGroupStyle) Name() string {
	return "box-flex-group"
}

// BreakBefore represent the CSS style "break-before" with value "auto | avoid | always | all | avoid-page | page | left | right | recto | verso | avoid-column | column | avoid-region | region"
// See https://developer.mozilla.org/docs/Web/CSS/break-before
type BreakBeforeStyle string

func (t BreakBeforeStyle) Name() string {
	return "break-before"
}

const BreakBeforeStyleAuto = "auto"

const BreakBeforeStyleAvoid = "avoid"

const BreakBeforeStyleAlways = "always"

const BreakBeforeStyleAll = "all"

const BreakBeforeStyleAvoidPage = "avoid-page"

const BreakBeforeStylePage = "page"

const BreakBeforeStyleLeft = "left"

const BreakBeforeStyleRight = "right"

const BreakBeforeStyleRecto = "recto"

const BreakBeforeStyleVerso = "verso"

const BreakBeforeStyleAvoidColumn = "avoid-column"

const BreakBeforeStyleColumn = "column"

const BreakBeforeStyleAvoidRegion = "avoid-region"

const BreakBeforeStyleRegion = "region"

// Color represent the CSS style "color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/color
type ColorStyle Color

func (t ColorStyle) Name() string {
	return "color"
}

// TransitionDelay represent the CSS style "transition-delay" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/transition-delay
type TransitionDelayStyle time.Duration

func (t TransitionDelayStyle) Name() string {
	return "transition-delay"
}

// WordBreak represent the CSS style "word-break" with value "normal | break-all | keep-all | break-word"
// See https://developer.mozilla.org/docs/Web/CSS/word-break
type WordBreakStyle string

func (t WordBreakStyle) Name() string {
	return "word-break"
}

const WordBreakStyleNormal = "normal"

const WordBreakStyleBreakAll = "break-all"

const WordBreakStyleKeepAll = "keep-all"

const WordBreakStyleBreakWord = "break-word"

// InitialLetterAlign represent the CSS style "initial-letter-align" with value "auto | alphabetic | hanging | ideographic"
// See https://developer.mozilla.org/docs/Web/CSS/initial-letter-align
type InitialLetterAlignStyle string

func (t InitialLetterAlignStyle) Name() string {
	return "initial-letter-align"
}

const InitialLetterAlignStyleAuto = "auto"

const InitialLetterAlignStyleAlphabetic = "alphabetic"

const InitialLetterAlignStyleHanging = "hanging"

const InitialLetterAlignStyleIdeographic = "ideographic"

// MarginInlineStart represent the CSS style "margin-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-inline-start
type MarginInlineStartStyle string

func (t MarginInlineStartStyle) Name() string {
	return "margin-inline-start"
}

// MaskBorderOutset represent the CSS style "mask-border-outset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-border-outset
type MaskBorderOutsetStyle string

func (t MaskBorderOutsetStyle) Name() string {
	return "mask-border-outset"
}

// Marker represent the CSS style "marker" with value ""
// See
type MarkerStyle string

func (t MarkerStyle) Name() string {
	return "marker"
}

// BorderInlineEndWidth represent the CSS style "border-inline-end-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-end-width
type BorderInlineEndWidthStyle string

func (t BorderInlineEndWidthStyle) Name() string {
	return "border-inline-end-width"
}

// ColumnFill represent the CSS style "column-fill" with value "auto | balance | balance-all"
// See https://developer.mozilla.org/docs/Web/CSS/column-fill
type ColumnFillStyle string

func (t ColumnFillStyle) Name() string {
	return "column-fill"
}

const ColumnFillStyleAuto = "auto"

const ColumnFillStyleBalance = "balance"

const ColumnFillStyleBalanceAll = "balance-all"

// FontSizeAdjust represent the CSS style "font-size-adjust" with value "none"
// See https://developer.mozilla.org/docs/Web/CSS/font-size-adjust
type FontSizeAdjustStyle string

func (t FontSizeAdjustStyle) Name() string {
	return "font-size-adjust"
}

const FontSizeAdjustStyleNone = "none"

// GridTemplateAreas represent the CSS style "grid-template-areas" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-template-areas
type GridTemplateAreasStyle string

func (t GridTemplateAreasStyle) Name() string {
	return "grid-template-areas"
}

// ScrollMarginTop represent the CSS style "scroll-margin-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-top
type ScrollMarginTopStyle float64

func (t ScrollMarginTopStyle) Name() string {
	return "scroll-margin-top"
}

// TransformStyle represent the CSS style "transform-style" with value "flat | preserve-3d"
// See https://developer.mozilla.org/docs/Web/CSS/transform-style
type TransformStyleStyle string

func (t TransformStyleStyle) Name() string {
	return "transform-style"
}

const TransformStyleStyleFlat = "flat"

const TransformStyleStylePreserve3d = "preserve-3d"

// LightingColor represent the CSS style "lighting-color" with value ""
// See
type LightingColorStyle Color

func (t LightingColorStyle) Name() string {
	return "lighting-color"
}

// AnimationDuration represent the CSS style "animation-duration" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-duration
type AnimationDurationStyle time.Duration

func (t AnimationDurationStyle) Name() string {
	return "animation-duration"
}

// OffsetAnchor represent the CSS style "offset-anchor" with value ""
// See
type OffsetAnchorStyle string

func (t OffsetAnchorStyle) Name() string {
	return "offset-anchor"
}

// Overflow represent the CSS style "overflow" with value "visible | hidden | clip | scroll | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overflow
type OverflowStyle string

func (t OverflowStyle) Name() string {
	return "overflow"
}

const OverflowStyleVisible = "visible"

const OverflowStyleHidden = "hidden"

const OverflowStyleClip = "clip"

const OverflowStyleScroll = "scroll"

const OverflowStyleAuto = "auto"

// ScrollMarginBlockStart represent the CSS style "scroll-margin-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-block-start
type ScrollMarginBlockStartStyle float64

func (t ScrollMarginBlockStartStyle) Name() string {
	return "scroll-margin-block-start"
}

// BackgroundBlendMode represent the CSS style "background-blend-mode" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-blend-mode
type BackgroundBlendModeStyle string

func (t BackgroundBlendModeStyle) Name() string {
	return "background-blend-mode"
}

// FontStretch represent the CSS style "font-stretch" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-stretch
type FontStretchStyle string

func (t FontStretchStyle) Name() string {
	return "font-stretch"
}

// OffsetPosition represent the CSS style "offset-position" with value ""
// See
type OffsetPositionStyle string

func (t OffsetPositionStyle) Name() string {
	return "offset-position"
}

// FontVariantLigatures represent the CSS style "font-variant-ligatures" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-ligatures
type FontVariantLigaturesStyle string

func (t FontVariantLigaturesStyle) Name() string {
	return "font-variant-ligatures"
}

// InsetInline represent the CSS style "inset-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-inline
type InsetInlineStyle string

func (t InsetInlineStyle) Name() string {
	return "inset-inline"
}

// ScrollSnapTypeY represent the CSS style "scroll-snap-type-y" with value "none | mandatory | proximity"
// See https://developer.mozilla.org/docs/Web/CSS/scroll-snap-type-y
type ScrollSnapTypeYStyle string

func (t ScrollSnapTypeYStyle) Name() string {
	return "scroll-snap-type-y"
}

const ScrollSnapTypeYStyleNone = "none"

const ScrollSnapTypeYStyleMandatory = "mandatory"

const ScrollSnapTypeYStyleProximity = "proximity"

// BaselineShift represent the CSS style "baseline-shift" with value "baseline | sub | super"
// See
type BaselineShiftStyle string

func (t BaselineShiftStyle) Name() string {
	return "baseline-shift"
}

const BaselineShiftStyleBaseline = "baseline"

const BaselineShiftStyleSub = "sub"

const BaselineShiftStyleSuper = "super"

// BorderImageSource represent the CSS style "border-image-source" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-source
type BorderImageSourceStyle string

func (t BorderImageSourceStyle) Name() string {
	return "border-image-source"
}

// Clear represent the CSS style "clear" with value "none | left | right | both | inline-start | inline-end"
// See https://developer.mozilla.org/docs/Web/CSS/clear
type ClearStyle string

func (t ClearStyle) Name() string {
	return "clear"
}

const ClearStyleNone = "none"

const ClearStyleLeft = "left"

const ClearStyleRight = "right"

const ClearStyleBoth = "both"

const ClearStyleInlineStart = "inline-start"

const ClearStyleInlineEnd = "inline-end"

// ColumnWidth represent the CSS style "column-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-width
type ColumnWidthStyle string

func (t ColumnWidthStyle) Name() string {
	return "column-width"
}

// Cursor represent the CSS style "cursor" with value "auto | default | none | context-menu | help | pointer | progress | wait | cell | crosshair | text | vertical-text | alias | copy | move | no-drop | not-allowed | e-resize | n-resize | ne-resize | nw-resize | s-resize | se-resize | sw-resize | w-resize | ew-resize | ns-resize | nesw-resize | nwse-resize | col-resize | row-resize | all-scroll | zoom-in | zoom-out | grab | grabbing"
// See https://developer.mozilla.org/docs/Web/CSS/cursor
type CursorStyle string

func (t CursorStyle) Name() string {
	return "cursor"
}

const CursorStyleAuto = "auto"

const CursorStyleDefault = "default"

const CursorStyleNone = "none"

const CursorStyleContextMenu = "context-menu"

const CursorStyleHelp = "help"

const CursorStylePointer = "pointer"

const CursorStyleProgress = "progress"

const CursorStyleWait = "wait"

const CursorStyleCell = "cell"

const CursorStyleCrosshair = "crosshair"

const CursorStyleText = "text"

const CursorStyleVerticalText = "vertical-text"

const CursorStyleAlias = "alias"

const CursorStyleCopy = "copy"

const CursorStyleMove = "move"

const CursorStyleNoDrop = "no-drop"

const CursorStyleNotAllowed = "not-allowed"

const CursorStyleEResize = "e-resize"

const CursorStyleNResize = "n-resize"

const CursorStyleNeResize = "ne-resize"

const CursorStyleNwResize = "nw-resize"

const CursorStyleSResize = "s-resize"

const CursorStyleSeResize = "se-resize"

const CursorStyleSwResize = "sw-resize"

const CursorStyleWResize = "w-resize"

const CursorStyleEwResize = "ew-resize"

const CursorStyleNsResize = "ns-resize"

const CursorStyleNeswResize = "nesw-resize"

const CursorStyleNwseResize = "nwse-resize"

const CursorStyleColResize = "col-resize"

const CursorStyleRowResize = "row-resize"

const CursorStyleAllScroll = "all-scroll"

const CursorStyleZoomIn = "zoom-in"

const CursorStyleZoomOut = "zoom-out"

const CursorStyleGrab = "grab"

const CursorStyleGrabbing = "grabbing"

// TextEmphasisStyle represent the CSS style "text-emphasis-style" with value "none | filled | open | dot | circle | double-circle | triangle | sesame | string"
// See https://developer.mozilla.org/docs/Web/CSS/text-emphasis-style
type TextEmphasisStyleStyle string

func (t TextEmphasisStyleStyle) Name() string {
	return "text-emphasis-style"
}

const TextEmphasisStyleStyleNone = "none"

const TextEmphasisStyleStyleFilled = "filled"

const TextEmphasisStyleStyleOpen = "open"

const TextEmphasisStyleStyleDot = "dot"

const TextEmphasisStyleStyleCircle = "circle"

const TextEmphasisStyleStyleDoubleCircle = "double-circle"

const TextEmphasisStyleStyleTriangle = "triangle"

const TextEmphasisStyleStyleSesame = "sesame"

const TextEmphasisStyleStyleString = "string"

// BorderTopLeftRadius represent the CSS style "border-top-left-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-left-radius
type BorderTopLeftRadiusStyle string

func (t BorderTopLeftRadiusStyle) Name() string {
	return "border-top-left-radius"
}

// ScrollMarginInlineStart represent the CSS style "scroll-margin-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-inline-start
type ScrollMarginInlineStartStyle float64

func (t ScrollMarginInlineStartStyle) Name() string {
	return "scroll-margin-inline-start"
}

// ScrollPaddingBlockEnd represent the CSS style "scroll-padding-block-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-block-end
type ScrollPaddingBlockEndStyle string

func (t ScrollPaddingBlockEndStyle) Name() string {
	return "scroll-padding-block-end"
}

// ShapeOutside represent the CSS style "shape-outside" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/shape-outside
type ShapeOutsideStyle string

func (t ShapeOutsideStyle) Name() string {
	return "shape-outside"
}

// BorderBottomColor represent the CSS style "border-bottom-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-bottom-color
type BorderBottomColorStyle Color

func (t BorderBottomColorStyle) Name() string {
	return "border-bottom-color"
}

// Quotes represent the CSS style "quotes" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/quotes
type QuotesStyle string

func (t QuotesStyle) Name() string {
	return "quotes"
}

// FillRule represent the CSS style "fill-rule" with value "nonzero | evenodd"
// See
type FillRuleStyle string

func (t FillRuleStyle) Name() string {
	return "fill-rule"
}

const FillRuleStyleNonzero = "nonzero"

const FillRuleStyleEvenodd = "evenodd"

// StopColor represent the CSS style "stop-color" with value ""
// See
type StopColorStyle Color

func (t StopColorStyle) Name() string {
	return "stop-color"
}

// JustifySelf represent the CSS style "justify-self" with value "auto | normal | stretch | end | start | flex-start | flex-end | center | left | right |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe center | unsafe center"
// See https://developer.mozilla.org/docs/Web/CSS/justify-self
type JustifySelfStyle string

func (t JustifySelfStyle) Name() string {
	return "justify-self"
}

const JustifySelfStyleAuto = "auto"

const JustifySelfStyleNormal = "normal"

const JustifySelfStyleStretch = "stretch"

const JustifySelfStyleEnd = "end"

const JustifySelfStyleStart = "start"

const JustifySelfStyleFlexStart = "flex-start"

const JustifySelfStyleFlexEnd = "flex-end"

const JustifySelfStyleCenter = "center"

const JustifySelfStyleLeft = "left"

const JustifySelfStyleRight = "right"

const JustifySelfStyleBaseline = "baseline"

const JustifySelfStyleFirstBaseline = "first baseline"

const JustifySelfStyleLastBaseline = "last baseline"

const JustifySelfStyleSpaceBetween = "space-between"

const JustifySelfStyleSpaceAround = "space-around"

const JustifySelfStyleSpaceEvenly = "space-evenly"

const JustifySelfStyleSafeCenter = "safe center"

const JustifySelfStyleUnsafeCenter = "unsafe center"

// ScrollMarginBlock represent the CSS style "scroll-margin-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-block
type ScrollMarginBlockStyle float64

func (t ScrollMarginBlockStyle) Name() string {
	return "scroll-margin-block"
}

// TextDecorationSkip represent the CSS style "text-decoration-skip" with value "none | objects | spaces | leading-spaces | trailing-spaces | edges | box-decoration"
// See https://developer.mozilla.org/docs/Web/CSS/text-decoration-skip
type TextDecorationSkipStyle string

func (t TextDecorationSkipStyle) Name() string {
	return "text-decoration-skip"
}

const TextDecorationSkipStyleNone = "none"

const TextDecorationSkipStyleObjects = "objects"

const TextDecorationSkipStyleSpaces = "spaces"

const TextDecorationSkipStyleLeadingSpaces = "leading-spaces"

const TextDecorationSkipStyleTrailingSpaces = "trailing-spaces"

const TextDecorationSkipStyleEdges = "edges"

const TextDecorationSkipStyleBoxDecoration = "box-decoration"

// FloodOpacity represent the CSS style "flood-opacity" with value ""
// See
type FloodOpacityStyle float64

func (t FloodOpacityStyle) Name() string {
	return "flood-opacity"
}

// AnimationDelay represent the CSS style "animation-delay" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-delay
type AnimationDelayStyle time.Duration

func (t AnimationDelayStyle) Name() string {
	return "animation-delay"
}

// BackgroundPosition represent the CSS style "background-position" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-position
type BackgroundPositionStyle string

func (t BackgroundPositionStyle) Name() string {
	return "background-position"
}

// BorderInlineStyle represent the CSS style "border-inline-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-inline-style
type BorderInlineStyleStyle string

func (t BorderInlineStyleStyle) Name() string {
	return "border-inline-style"
}

// MaskImage represent the CSS style "mask-image" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-image
type MaskImageStyle string

func (t MaskImageStyle) Name() string {
	return "mask-image"
}

// PaddingTop represent the CSS style "padding-top" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-top
type PaddingTopStyle string

func (t PaddingTopStyle) Name() string {
	return "padding-top"
}

// StrokeMiterlimit represent the CSS style "stroke-miterlimit" with value ""
// See
type StrokeMiterlimitStyle float64

func (t StrokeMiterlimitStyle) Name() string {
	return "stroke-miterlimit"
}

// BorderBlockColor represent the CSS style "border-block-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-color
type BorderBlockColorStyle Color

func (t BorderBlockColorStyle) Name() string {
	return "border-block-color"
}

// Float represent the CSS style "float" with value "left | right | none | inline-start | inline-end"
// See https://developer.mozilla.org/docs/Web/CSS/float
type FloatStyle string

func (t FloatStyle) Name() string {
	return "float"
}

const FloatStyleLeft = "left"

const FloatStyleRight = "right"

const FloatStyleNone = "none"

const FloatStyleInlineStart = "inline-start"

const FloatStyleInlineEnd = "inline-end"

// GridAutoFlow represent the CSS style "grid-auto-flow" with value "row | column | dense"
// See https://developer.mozilla.org/docs/Web/CSS/grid-auto-flow
type GridAutoFlowStyle string

func (t GridAutoFlowStyle) Name() string {
	return "grid-auto-flow"
}

const GridAutoFlowStyleRow = "row"

const GridAutoFlowStyleColumn = "column"

const GridAutoFlowStyleDense = "dense"

// PaddingBottom represent the CSS style "padding-bottom" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-bottom
type PaddingBottomStyle string

func (t PaddingBottomStyle) Name() string {
	return "padding-bottom"
}

// OverflowY represent the CSS style "overflow-y" with value "visible | hidden | clip | scroll | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overflow-y
type OverflowYStyle string

func (t OverflowYStyle) Name() string {
	return "overflow-y"
}

const OverflowYStyleVisible = "visible"

const OverflowYStyleHidden = "hidden"

const OverflowYStyleClip = "clip"

const OverflowYStyleScroll = "scroll"

const OverflowYStyleAuto = "auto"

// PaddingInline represent the CSS style "padding-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-inline
type PaddingInlineStyle string

func (t PaddingInlineStyle) Name() string {
	return "padding-inline"
}

// ScrollMarginInline represent the CSS style "scroll-margin-inline" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-margin-inline
type ScrollMarginInlineStyle float64

func (t ScrollMarginInlineStyle) Name() string {
	return "scroll-margin-inline"
}

// StrokeDashoffset represent the CSS style "stroke-dashoffset" with value ""
// See
type StrokeDashoffsetStyle float64

func (t StrokeDashoffsetStyle) Name() string {
	return "stroke-dashoffset"
}

// BorderLeftWidth represent the CSS style "border-left-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-left-width
type BorderLeftWidthStyle string

func (t BorderLeftWidthStyle) Name() string {
	return "border-left-width"
}

// ContentVisibility represent the CSS style "content-visibility" with value "visible | auto | hidden"
// See https://developer.mozilla.org/docs/Web/CSS/content-visibility
type ContentVisibilityStyle string

func (t ContentVisibilityStyle) Name() string {
	return "content-visibility"
}

const ContentVisibilityStyleVisible = "visible"

const ContentVisibilityStyleAuto = "auto"

const ContentVisibilityStyleHidden = "hidden"

// FlexShrink represent the CSS style "flex-shrink" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/flex-shrink
type FlexShrinkStyle float64

func (t FlexShrinkStyle) Name() string {
	return "flex-shrink"
}

// FontVariantPosition represent the CSS style "font-variant-position" with value "normal | sub | super"
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-position
type FontVariantPositionStyle string

func (t FontVariantPositionStyle) Name() string {
	return "font-variant-position"
}

const FontVariantPositionStyleNormal = "normal"

const FontVariantPositionStyleSub = "sub"

const FontVariantPositionStyleSuper = "super"

// TextUnderlineOffset represent the CSS style "text-underline-offset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-underline-offset
type TextUnderlineOffsetStyle string

func (t TextUnderlineOffsetStyle) Name() string {
	return "text-underline-offset"
}

// TextShadow represent the CSS style "text-shadow" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/text-shadow
type TextShadowStyle string

func (t TextShadowStyle) Name() string {
	return "text-shadow"
}

// WillChange represent the CSS style "will-change" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/will-change
type WillChangeStyle string

func (t WillChangeStyle) Name() string {
	return "will-change"
}

// BorderEndEndRadius represent the CSS style "border-end-end-radius" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-end-end-radius
type BorderEndEndRadiusStyle string

func (t BorderEndEndRadiusStyle) Name() string {
	return "border-end-end-radius"
}

// Isolation represent the CSS style "isolation" with value "auto | isolate"
// See https://developer.mozilla.org/docs/Web/CSS/isolation
type IsolationStyle string

func (t IsolationStyle) Name() string {
	return "isolation"
}

const IsolationStyleAuto = "auto"

const IsolationStyleIsolate = "isolate"

// MaskSize represent the CSS style "mask-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-size
type MaskSizeStyle string

func (t MaskSizeStyle) Name() string {
	return "mask-size"
}

// MathStyle represent the CSS style "math-style" with value "normal | compact"
// See https://developer.mozilla.org/docs/Web/CSS/math-style
type MathStyleStyle string

func (t MathStyleStyle) Name() string {
	return "math-style"
}

const MathStyleStyleNormal = "normal"

const MathStyleStyleCompact = "compact"

// TextOverflow represent the CSS style "text-overflow" with value "clip | ellipsis | clip ellipsis | ellipsis clip"
// See https://developer.mozilla.org/docs/Web/CSS/text-overflow
type TextOverflowStyle string

func (t TextOverflowStyle) Name() string {
	return "text-overflow"
}

const TextOverflowStyleClip = "clip"

const TextOverflowStyleEllipsis = "ellipsis"

const TextOverflowStyleClipEllipsis = "clip ellipsis"

const TextOverflowStyleEllipsisClip = "ellipsis clip"

// BorderBlockEndColor represent the CSS style "border-block-end-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-end-color
type BorderBlockEndColorStyle Color

func (t BorderBlockEndColorStyle) Name() string {
	return "border-block-end-color"
}

// BorderImageRepeat represent the CSS style "border-image-repeat" with value "stretch | repeat | round | space"
// See https://developer.mozilla.org/docs/Web/CSS/border-image-repeat
type BorderImageRepeatStyle string

func (t BorderImageRepeatStyle) Name() string {
	return "border-image-repeat"
}

const BorderImageRepeatStyleStretch = "stretch"

const BorderImageRepeatStyleRepeat = "repeat"

const BorderImageRepeatStyleRound = "round"

const BorderImageRepeatStyleSpace = "space"

// Content represent the CSS style "content" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/content
type ContentStyle string

func (t ContentStyle) Name() string {
	return "content"
}

// FontKerning represent the CSS style "font-kerning" with value "auto | normal | none"
// See https://developer.mozilla.org/docs/Web/CSS/font-kerning
type FontKerningStyle string

func (t FontKerningStyle) Name() string {
	return "font-kerning"
}

const FontKerningStyleAuto = "auto"

const FontKerningStyleNormal = "normal"

const FontKerningStyleNone = "none"

// ScrollPaddingLeft represent the CSS style "scroll-padding-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding-left
type ScrollPaddingLeftStyle string

func (t ScrollPaddingLeftStyle) Name() string {
	return "scroll-padding-left"
}

// Visibility represent the CSS style "visibility" with value "visible | hidden | collapse"
// See https://developer.mozilla.org/docs/Web/CSS/visibility
type VisibilityStyle string

func (t VisibilityStyle) Name() string {
	return "visibility"
}

const VisibilityStyleVisible = "visible"

const VisibilityStyleHidden = "hidden"

const VisibilityStyleCollapse = "collapse"

// BorderLeftColor represent the CSS style "border-left-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-left-color
type BorderLeftColorStyle Color

func (t BorderLeftColorStyle) Name() string {
	return "border-left-color"
}

// ColumnRuleStyle represent the CSS style "column-rule-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/column-rule-style
type ColumnRuleStyleStyle string

func (t ColumnRuleStyleStyle) Name() string {
	return "column-rule-style"
}

// CounterSet represent the CSS style "counter-set" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/counter-set
type CounterSetStyle string

func (t CounterSetStyle) Name() string {
	return "counter-set"
}

// GridColumnEnd represent the CSS style "grid-column-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-column-end
type GridColumnEndStyle string

func (t GridColumnEndStyle) Name() string {
	return "grid-column-end"
}

// MinHeight represent the CSS style "min-height" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-height
type MinHeightStyle string

func (t MinHeightStyle) Name() string {
	return "min-height"
}

// BackgroundPositionX represent the CSS style "background-position-x" with value "center | left | right | x-start | x-end"
// See https://developer.mozilla.org/docs/Web/CSS/background-position-x
type BackgroundPositionXStyle string

func (t BackgroundPositionXStyle) Name() string {
	return "background-position-x"
}

const BackgroundPositionXStyleCenter = "center"

const BackgroundPositionXStyleLeft = "left"

const BackgroundPositionXStyleRight = "right"

const BackgroundPositionXStyleXStart = "x-start"

const BackgroundPositionXStyleXEnd = "x-end"

// LineBreak represent the CSS style "line-break" with value "auto | loose | normal | strict | anywhere"
// See https://developer.mozilla.org/docs/Web/CSS/line-break
type LineBreakStyle string

func (t LineBreakStyle) Name() string {
	return "line-break"
}

const LineBreakStyleAuto = "auto"

const LineBreakStyleLoose = "loose"

const LineBreakStyleNormal = "normal"

const LineBreakStyleStrict = "strict"

const LineBreakStyleAnywhere = "anywhere"

// MaxHeight represent the CSS style "max-height" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-height
type MaxHeightStyle string

func (t MaxHeightStyle) Name() string {
	return "max-height"
}

// MarkerMid represent the CSS style "marker-mid" with value ""
// See
type MarkerMidStyle string

func (t MarkerMidStyle) Name() string {
	return "marker-mid"
}

// AlignSelf represent the CSS style "align-self" with value "auto | normal | stretch | end | start | flex-start | flex-end | center |  baseline | first baseline | last baseline | space-between | space-around | space-evenly | safe | unsafe"
// See https://developer.mozilla.org/docs/Web/CSS/align-self
type AlignSelfStyle string

func (t AlignSelfStyle) Name() string {
	return "align-self"
}

const AlignSelfStyleAuto = "auto"

const AlignSelfStyleNormal = "normal"

const AlignSelfStyleStretch = "stretch"

const AlignSelfStyleEnd = "end"

const AlignSelfStyleStart = "start"

const AlignSelfStyleFlexStart = "flex-start"

const AlignSelfStyleFlexEnd = "flex-end"

const AlignSelfStyleCenter = "center"

const AlignSelfStyleBaseline = "baseline"

const AlignSelfStyleFirstBaseline = "first baseline"

const AlignSelfStyleLastBaseline = "last baseline"

const AlignSelfStyleSpaceBetween = "space-between"

const AlignSelfStyleSpaceAround = "space-around"

const AlignSelfStyleSpaceEvenly = "space-evenly"

const AlignSelfStyleSafe = "safe"

const AlignSelfStyleUnsafe = "unsafe"

// AnimationIterationCount represent the CSS style "animation-iteration-count" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-iteration-count
type AnimationIterationCountStyle float64

func (t AnimationIterationCountStyle) Name() string {
	return "animation-iteration-count"
}

// FlexDirection represent the CSS style "flex-direction" with value "row | row-reverse | column | column-reverse"
// See https://developer.mozilla.org/docs/Web/CSS/flex-direction
type FlexDirectionStyle string

func (t FlexDirectionStyle) Name() string {
	return "flex-direction"
}

const FlexDirectionStyleRow = "row"

const FlexDirectionStyleRowReverse = "row-reverse"

const FlexDirectionStyleColumn = "column"

const FlexDirectionStyleColumnReverse = "column-reverse"

// ImeMode represent the CSS style "ime-mode" with value "auto | normal | active | inactive | disabled"
// See https://developer.mozilla.org/docs/Web/CSS/ime-mode
type ImeModeStyle string

func (t ImeModeStyle) Name() string {
	return "ime-mode"
}

const ImeModeStyleAuto = "auto"

const ImeModeStyleNormal = "normal"

const ImeModeStyleActive = "active"

const ImeModeStyleInactive = "inactive"

const ImeModeStyleDisabled = "disabled"

// OverscrollBehaviorInline represent the CSS style "overscroll-behavior-inline" with value "contain | none | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overscroll-behavior-inline
type OverscrollBehaviorInlineStyle string

func (t OverscrollBehaviorInlineStyle) Name() string {
	return "overscroll-behavior-inline"
}

const OverscrollBehaviorInlineStyleContain = "contain"

const OverscrollBehaviorInlineStyleNone = "none"

const OverscrollBehaviorInlineStyleAuto = "auto"

// PaddingBlock represent the CSS style "padding-block" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-block
type PaddingBlockStyle string

func (t PaddingBlockStyle) Name() string {
	return "padding-block"
}

// Width represent the CSS style "width" with value "auto |  min-content | max-content"
// See https://developer.mozilla.org/docs/Web/CSS/width
type WidthStyle string

func (t WidthStyle) Name() string {
	return "width"
}

const WidthStyleAuto = "auto"

const WidthStyleMinContent = "min-content"

const WidthStyleMaxContent = "max-content"

// ScaleLeft represent the CSS style "scale-left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scale
type ScaleLeftStyle string

func (t ScaleLeftStyle) Name() string {
	return "scale-left"
}

// StrokeOpacity represent the CSS style "stroke-opacity" with value ""
// See
type StrokeOpacityStyle float64

func (t StrokeOpacityStyle) Name() string {
	return "stroke-opacity"
}

// BorderImageOutset represent the CSS style "border-image-outset" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-image-outset
type BorderImageOutsetStyle string

func (t BorderImageOutsetStyle) Name() string {
	return "border-image-outset"
}

// InsetBlockStart represent the CSS style "inset-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/inset-block-start
type InsetBlockStartStyle string

func (t InsetBlockStartStyle) Name() string {
	return "inset-block-start"
}

// MaxBlockSize represent the CSS style "max-block-size" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/max-block-size
type MaxBlockSizeStyle string

func (t MaxBlockSizeStyle) Name() string {
	return "max-block-size"
}

// OutlineWidth represent the CSS style "outline-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/outline-width
type OutlineWidthStyle string

func (t OutlineWidthStyle) Name() string {
	return "outline-width"
}

// GridRowEnd represent the CSS style "grid-row-end" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-row-end
type GridRowEndStyle string

func (t GridRowEndStyle) Name() string {
	return "grid-row-end"
}

// InitialLetter represent the CSS style "initial-letter" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/initial-letter
type InitialLetterStyle string

func (t InitialLetterStyle) Name() string {
	return "initial-letter"
}

// MarginBlockStart represent the CSS style "margin-block-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/margin-block-start
type MarginBlockStartStyle string

func (t MarginBlockStartStyle) Name() string {
	return "margin-block-start"
}

// FillOpacity represent the CSS style "fill-opacity" with value ""
// See
type FillOpacityStyle float64

func (t FillOpacityStyle) Name() string {
	return "fill-opacity"
}

// StrokeWidth represent the CSS style "stroke-width" with value ""
// See
type StrokeWidthStyle float64

func (t StrokeWidthStyle) Name() string {
	return "stroke-width"
}

// BorderCollapse represent the CSS style "border-collapse" with value "collapse | separate"
// See https://developer.mozilla.org/docs/Web/CSS/border-collapse
type BorderCollapseStyle string

func (t BorderCollapseStyle) Name() string {
	return "border-collapse"
}

const BorderCollapseStyleCollapse = "collapse"

const BorderCollapseStyleSeparate = "separate"

// BorderRightColor represent the CSS style "border-right-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-right-color
type BorderRightColorStyle Color

func (t BorderRightColorStyle) Name() string {
	return "border-right-color"
}

// BorderTopStyle represent the CSS style "border-top-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-top-style
type BorderTopStyleStyle string

func (t BorderTopStyleStyle) Name() string {
	return "border-top-style"
}

// ScrollPadding represent the CSS style "scroll-padding" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/scroll-padding
type ScrollPaddingStyle string

func (t ScrollPaddingStyle) Name() string {
	return "scroll-padding"
}

// BorderBlockStartWidth represent the CSS style "border-block-start-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-start-width
type BorderBlockStartWidthStyle string

func (t BorderBlockStartWidthStyle) Name() string {
	return "border-block-start-width"
}

// Contain represent the CSS style "contain" with value "none | strict | content | size  layout | style | paint "
// See https://developer.mozilla.org/docs/Web/CSS/contain
type ContainStyle string

func (t ContainStyle) Name() string {
	return "contain"
}

const ContainStyleNone = "none"

const ContainStyleStrict = "strict"

const ContainStyleContent = "content"

const ContainStyleSizeLayout = "size  layout"

const ContainStyleStyle = "style"

const ContainStylePaint = "paint"

// GridAutoRows represent the CSS style "grid-auto-rows" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-auto-rows
type GridAutoRowsStyle string

func (t GridAutoRowsStyle) Name() string {
	return "grid-auto-rows"
}

// Left represent the CSS style "left" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/left
type LeftStyle string

func (t LeftStyle) Name() string {
	return "left"
}

// PaddingInlineStart represent the CSS style "padding-inline-start" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/padding-inline-start
type PaddingInlineStartStyle string

func (t PaddingInlineStartStyle) Name() string {
	return "padding-inline-start"
}

// BackgroundColor represent the CSS style "background-color" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/background-color
type BackgroundColorStyle Color

func (t BackgroundColorStyle) Name() string {
	return "background-color"
}

// MaskClip represent the CSS style "mask-clip" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/mask-clip
type MaskClipStyle string

func (t MaskClipStyle) Name() string {
	return "mask-clip"
}

// RowGap represent the CSS style "row-gap" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/row-gap
type RowGapStyle float64

func (t RowGapStyle) Name() string {
	return "row-gap"
}

// ScrollbarWidth represent the CSS style "scrollbar-width" with value "auto | thin | none"
// See https://developer.mozilla.org/docs/Web/CSS/scrollbar-width
type ScrollbarWidthStyle string

func (t ScrollbarWidthStyle) Name() string {
	return "scrollbar-width"
}

const ScrollbarWidthStyleAuto = "auto"

const ScrollbarWidthStyleThin = "thin"

const ScrollbarWidthStyleNone = "none"

// WritingMode represent the CSS style "writing-mode" with value "horizontal-tb | vertical-rl | vertical-lr | sideways-rl | sideways-lr"
// See https://developer.mozilla.org/docs/Web/CSS/writing-mode
type WritingModeStyle string

func (t WritingModeStyle) Name() string {
	return "writing-mode"
}

const WritingModeStyleHorizontalTb = "horizontal-tb"

const WritingModeStyleVerticalRl = "vertical-rl"

const WritingModeStyleVerticalLr = "vertical-lr"

const WritingModeStyleSidewaysRl = "sideways-rl"

const WritingModeStyleSidewaysLr = "sideways-lr"

// BoxDirection represent the CSS style "box-direction" with value "normal | reverse | inherit"
// See https://developer.mozilla.org/docs/Web/CSS/box-direction
type BoxDirectionStyle string

func (t BoxDirectionStyle) Name() string {
	return "box-direction"
}

const BoxDirectionStyleNormal = "normal"

const BoxDirectionStyleReverse = "reverse"

const BoxDirectionStyleInherit = "inherit"

// FontVariantAlternates represent the CSS style "font-variant-alternates" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-variant-alternates
type FontVariantAlternatesStyle string

func (t FontVariantAlternatesStyle) Name() string {
	return "font-variant-alternates"
}

// ImageRendering represent the CSS style "image-rendering" with value "auto | crisp-edges | pixelated"
// See https://developer.mozilla.org/docs/Web/CSS/image-rendering
type ImageRenderingStyle string

func (t ImageRenderingStyle) Name() string {
	return "image-rendering"
}

const ImageRenderingStyleAuto = "auto"

const ImageRenderingStyleCrispEdges = "crisp-edges"

const ImageRenderingStylePixelated = "pixelated"

// Perspective represent the CSS style "perspective" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/perspective
type PerspectiveStyle string

func (t PerspectiveStyle) Name() string {
	return "perspective"
}

// PlaceContent represent the CSS style "place-content" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/place-content
type PlaceContentStyle string

func (t PlaceContentStyle) Name() string {
	return "place-content"
}

// OverflowX represent the CSS style "overflow-x" with value "visible | hidden | clip | scroll | auto"
// See https://developer.mozilla.org/docs/Web/CSS/overflow-x
type OverflowXStyle string

func (t OverflowXStyle) Name() string {
	return "overflow-x"
}

const OverflowXStyleVisible = "visible"

const OverflowXStyleHidden = "hidden"

const OverflowXStyleClip = "clip"

const OverflowXStyleScroll = "scroll"

const OverflowXStyleAuto = "auto"

// AspectRatio represent the CSS style "aspect-ratio" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/aspect-ratio
type AspectRatioStyle string

func (t AspectRatioStyle) Name() string {
	return "aspect-ratio"
}

// BorderBlockStartStyle represent the CSS style "border-block-start-style" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/border-block-start-style
type BorderBlockStartStyleStyle string

func (t BorderBlockStartStyleStyle) Name() string {
	return "border-block-start-style"
}

// FontSmooth represent the CSS style "font-smooth" with value "auto | never | always"
// See https://developer.mozilla.org/docs/Web/CSS/font-smooth
type FontSmoothStyle string

func (t FontSmoothStyle) Name() string {
	return "font-smooth"
}

const FontSmoothStyleAuto = "auto"

const FontSmoothStyleNever = "never"

const FontSmoothStyleAlways = "always"

// GridTemplateRows represent the CSS style "grid-template-rows" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/grid-template-rows
type GridTemplateRowsStyle string

func (t GridTemplateRowsStyle) Name() string {
	return "grid-template-rows"
}

// MinWidth represent the CSS style "min-width" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/min-width
type MinWidthStyle string

func (t MinWidthStyle) Name() string {
	return "min-width"
}

// MarkerStart represent the CSS style "marker-start" with value ""
// See
type MarkerStartStyle string

func (t MarkerStartStyle) Name() string {
	return "marker-start"
}

// AnimationPlayState represent the CSS style "animation-play-state" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/animation-play-state
type AnimationPlayStateStyle string

func (t AnimationPlayStateStyle) Name() string {
	return "animation-play-state"
}

// BoxOrdinalGroup represent the CSS style "box-ordinal-group" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/box-ordinal-group
type BoxOrdinalGroupStyle float64

func (t BoxOrdinalGroupStyle) Name() string {
	return "box-ordinal-group"
}

// FontFeatureSettings represent the CSS style "font-feature-settings" with value ""
// See https://developer.mozilla.org/docs/Web/CSS/font-feature-settings
type FontFeatureSettingsStyle string

func (t FontFeatureSettingsStyle) Name() string {
	return "font-feature-settings"
}

// MaxLines represent the CSS style "max-lines" with value ""
// See
type MaxLinesStyle string

func (t MaxLinesStyle) Name() string {
	return "max-lines"
}
