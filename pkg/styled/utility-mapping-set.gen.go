package styled

var UtilityToMappingSet = map[string]map[string]string{

	"animation-fill-mode": AnimationFillModeStyleUtilitiesMapping,

	"font-variant-alternates": FontVariantAlternatesStyleUtilitiesMapping,

	"transform-rotate-y": TransformRotateYStyleUtilitiesMapping,

	"stroke-miterlimit": StrokeMiterlimitStyleUtilitiesMapping,

	"dominant-baseline": DominantBaselineStyleUtilitiesMapping,

	"fill": FillStyleUtilitiesMapping,

	"stroke": StrokeStyleUtilitiesMapping,

	"background-blend-mode": BackgroundBlendModeStyleUtilitiesMapping,

	"min-width": MinWidthStyleUtilitiesMapping,

	"translate": TranslateStyleUtilitiesMapping,

	"ruby-position": RubyPositionStyleUtilitiesMapping,

	"marker-end": MarkerEndStyleUtilitiesMapping,

	"background-repeat": BackgroundRepeatStyleUtilitiesMapping,

	"inset-inline": InsetInlineStyleUtilitiesMapping,

	"paint-order": PaintOrderStyleUtilitiesMapping,

	"stop-opacity": StopOpacityStyleUtilitiesMapping,

	"scroll-padding-inline-end": ScrollPaddingInlineEndStyleUtilitiesMapping,

	"border-top-right-radius": BorderTopRightRadiusStyleUtilitiesMapping,

	"font-stretch": FontStretchStyleUtilitiesMapping,

	"position": PositionStyleUtilitiesMapping,

	"column-rule-style": ColumnRuleStyleStyleUtilitiesMapping,

	"line-clamp": LineClampStyleUtilitiesMapping,

	"font-feature-settings": FontFeatureSettingsStyleUtilitiesMapping,

	"flood-opacity": FloodOpacityStyleUtilitiesMapping,

	"stop-color": StopColorStyleUtilitiesMapping,

	"border-left-color": BorderLeftColorStyleUtilitiesMapping,

	"width": WidthStyleUtilitiesMapping,

	"fill-rule": FillRuleStyleUtilitiesMapping,

	"box-ordinal-group": BoxOrdinalGroupStyleUtilitiesMapping,

	"grid-auto-rows": GridAutoRowsStyleUtilitiesMapping,

	"shape-margin": ShapeMarginStyleUtilitiesMapping,

	"flood-color": FloodColorStyleUtilitiesMapping,

	"transform-translate-3d": TransformTranslate3dStyleUtilitiesMapping,

	"tab-size": TabSizeStyleUtilitiesMapping,

	"margin-top": MarginTopStyleUtilitiesMapping,

	"mask-position": MaskPositionStyleUtilitiesMapping,

	"line-break": LineBreakStyleUtilitiesMapping,

	"overscroll-behavior-block": OverscrollBehaviorBlockStyleUtilitiesMapping,

	"padding-block-start": PaddingBlockStartStyleUtilitiesMapping,

	"text-decoration-line": TextDecorationLineStyleUtilitiesMapping,

	"text-transform": TextTransformStyleUtilitiesMapping,

	"align-tracks": AlignTracksStyleUtilitiesMapping,

	"font-kerning": FontKerningStyleUtilitiesMapping,

	"transform-rotate-x": TransformRotateXStyleUtilitiesMapping,

	"border-image-source": BorderImageSourceStyleUtilitiesMapping,

	"flex-wrap": FlexWrapStyleUtilitiesMapping,

	"font-variant-caps": FontVariantCapsStyleUtilitiesMapping,

	"overflow-y": OverflowYStyleUtilitiesMapping,

	"color-interpolation": ColorInterpolationStyleUtilitiesMapping,

	"break-before": BreakBeforeStyleUtilitiesMapping,

	"empty-cells": EmptyCellsStyleUtilitiesMapping,

	"padding-block-end": PaddingBlockEndStyleUtilitiesMapping,

	"scroll-margin-block": ScrollMarginBlockStyleUtilitiesMapping,

	"scroll-padding-block": ScrollPaddingBlockStyleUtilitiesMapping,

	"scroll-margin-top": ScrollMarginTopStyleUtilitiesMapping,

	"background-position": BackgroundPositionStyleUtilitiesMapping,

	"mask-border-source": MaskBorderSourceStyleUtilitiesMapping,

	"text-decoration-color": TextDecorationColorStyleUtilitiesMapping,

	"marker-mid": MarkerMidStyleUtilitiesMapping,

	"list-style-type": ListStyleTypeStyleUtilitiesMapping,

	"mask-clip": MaskClipStyleUtilitiesMapping,

	"page-break-inside": PageBreakInsideStyleUtilitiesMapping,

	"pointer-events": PointerEventsStyleUtilitiesMapping,

	"text-emphasis-position-first": TextEmphasisPositionFirstStyleUtilitiesMapping,

	"border-bottom-color": BorderBottomColorStyleUtilitiesMapping,

	"border-block-start-color": BorderBlockStartColorStyleUtilitiesMapping,

	"border-left-style": BorderLeftStyleStyleUtilitiesMapping,

	"clear": ClearStyleUtilitiesMapping,

	"overscroll-behavior": OverscrollBehaviorStyleUtilitiesMapping,

	"transform-skew-x": TransformSkewXStyleUtilitiesMapping,

	"display": DisplayStyleUtilitiesMapping,

	"visibility": VisibilityStyleUtilitiesMapping,

	"content": ContentStyleUtilitiesMapping,

	"alignment-baseline": AlignmentBaselineStyleUtilitiesMapping,

	"border-bottom-width": BorderBottomWidthStyleUtilitiesMapping,

	"initial-letter-align": InitialLetterAlignStyleUtilitiesMapping,

	"border-bottom-right-radius": BorderBottomRightRadiusStyleUtilitiesMapping,

	"margin-left": MarginLeftStyleUtilitiesMapping,

	"text-align-last": TextAlignLastStyleUtilitiesMapping,

	"color-adjust": ColorAdjustStyleUtilitiesMapping,

	"margin-bottom": MarginBottomStyleUtilitiesMapping,

	"mask-border-width": MaskBorderWidthStyleUtilitiesMapping,

	"transform-translate-z": TransformTranslateZStyleUtilitiesMapping,

	"grid-column-gap": GridColumnGapStyleUtilitiesMapping,

	"inset-block-start": InsetBlockStartStyleUtilitiesMapping,

	"min-inline-size": MinInlineSizeStyleUtilitiesMapping,

	"word-spacing": WordSpacingStyleUtilitiesMapping,

	"animation-name": AnimationNameStyleUtilitiesMapping,

	"font-variant-east-asian": FontVariantEastAsianStyleUtilitiesMapping,

	"grid-rows": GridRowsStyleUtilitiesMapping,

	"inset-inline-start": InsetInlineStartStyleUtilitiesMapping,

	"all": AllStyleUtilitiesMapping,

	"max-lines": MaxLinesStyleUtilitiesMapping,

	"text-emphasis-position-second": TextEmphasisPositionSecondStyleUtilitiesMapping,

	"order": OrderStyleUtilitiesMapping,

	"outline-style": OutlineStyleStyleUtilitiesMapping,

	"transition-property": TransitionPropertyStyleUtilitiesMapping,

	"column-fill": ColumnFillStyleUtilitiesMapping,

	"mask-composite": MaskCompositeStyleUtilitiesMapping,

	"overflow-wrap": OverflowWrapStyleUtilitiesMapping,

	"aspect-ratio": AspectRatioStyleUtilitiesMapping,

	"line-height": LineHeightStyleUtilitiesMapping,

	"transition-duration": TransitionDurationStyleUtilitiesMapping,

	"border-start-end-radius": BorderStartEndRadiusStyleUtilitiesMapping,

	"border-top-color": BorderTopColorStyleUtilitiesMapping,

	"ime-mode": ImeModeStyleUtilitiesMapping,

	"border-inline-end-width": BorderInlineEndWidthStyleUtilitiesMapping,

	"grid-column-start": GridColumnStartStyleUtilitiesMapping,

	"text-decoration-style": TextDecorationStyleStyleUtilitiesMapping,

	"transform-translate": TransformTranslateStyleUtilitiesMapping,

	"border-start-start-radius": BorderStartStartRadiusStyleUtilitiesMapping,

	"font-synthesis": FontSynthesisStyleUtilitiesMapping,

	"scroll-margin-bottom": ScrollMarginBottomStyleUtilitiesMapping,

	"filter": FilterStyleUtilitiesMapping,

	"flex-basis": FlexBasisStyleUtilitiesMapping,

	"border-width": BorderWidthStyleUtilitiesMapping,

	"scroll-padding": ScrollPaddingStyleUtilitiesMapping,

	"border-block-end-color": BorderBlockEndColorStyleUtilitiesMapping,

	"widows": WidowsStyleUtilitiesMapping,

	"column-gap": ColumnGapStyleUtilitiesMapping,

	"transition-delay": TransitionDelayStyleUtilitiesMapping,

	"grid-row-gap": GridRowGapStyleUtilitiesMapping,

	"list-style-image": ListStyleImageStyleUtilitiesMapping,

	"transform": TransformStyleUtilitiesMapping,

	"transform-rotate-3d": TransformRotate3dStyleUtilitiesMapping,

	"scroll-snap-stop": ScrollSnapStopStyleUtilitiesMapping,

	"font-size-adjust": FontSizeAdjustStyleUtilitiesMapping,

	"font-variant-position": FontVariantPositionStyleUtilitiesMapping,

	"grid-row-start": GridRowStartStyleUtilitiesMapping,

	"animation-iteration-count": AnimationIterationCountStyleUtilitiesMapping,

	"border-end-end-radius": BorderEndEndRadiusStyleUtilitiesMapping,

	"image-rendering": ImageRenderingStyleUtilitiesMapping,

	"inset-block-end": InsetBlockEndStyleUtilitiesMapping,

	"mask-origin": MaskOriginStyleUtilitiesMapping,

	"text-decoration-skip": TextDecorationSkipStyleUtilitiesMapping,

	"stroke-dashoffset": StrokeDashoffsetStyleUtilitiesMapping,

	"quotes": QuotesStyleUtilitiesMapping,

	"align-items": AlignItemsStyleUtilitiesMapping,

	"box-sizing": BoxSizingStyleUtilitiesMapping,

	"clip-path": ClipPathStyleUtilitiesMapping,

	"column-rule-width": ColumnRuleWidthStyleUtilitiesMapping,

	"shape-outside": ShapeOutsideStyleUtilitiesMapping,

	"border-inline-width": BorderInlineWidthStyleUtilitiesMapping,

	"forced-color-adjust": ForcedColorAdjustStyleUtilitiesMapping,

	"place-content": PlaceContentStyleUtilitiesMapping,

	"table-layout": TableLayoutStyleUtilitiesMapping,

	"text-decoration-thickness": TextDecorationThicknessStyleUtilitiesMapping,

	"transform-origin": TransformOriginStyleUtilitiesMapping,

	"zoom": ZoomStyleUtilitiesMapping,

	"background-color": BackgroundColorStyleUtilitiesMapping,

	"margin-inline-start": MarginInlineStartStyleUtilitiesMapping,

	"outline-color": OutlineColorStyleUtilitiesMapping,

	"marker-start": MarkerStartStyleUtilitiesMapping,

	"break-after": BreakAfterStyleUtilitiesMapping,

	"text-rendering": TextRenderingStyleUtilitiesMapping,

	"transform-scale-x": TransformScaleXStyleUtilitiesMapping,

	"animation-direction": AnimationDirectionStyleUtilitiesMapping,

	"border-inline-style": BorderInlineStyleStyleUtilitiesMapping,

	"scroll-padding-right": ScrollPaddingRightStyleUtilitiesMapping,

	"scroll-snap-destination": ScrollSnapDestinationStyleUtilitiesMapping,

	"transform-translate-x": TransformTranslateXStyleUtilitiesMapping,

	"ruby-merge": RubyMergeStyleUtilitiesMapping,

	"transform-box": TransformBoxStyleUtilitiesMapping,

	"transform-scale-z": TransformScaleZStyleUtilitiesMapping,

	"outline-offset": OutlineOffsetStyleUtilitiesMapping,

	"unicode-bidi": UnicodeBidiStyleUtilitiesMapping,

	"inline-size": InlineSizeStyleUtilitiesMapping,

	"padding-bottom": PaddingBottomStyleUtilitiesMapping,

	"scrollbar-color": ScrollbarColorStyleUtilitiesMapping,

	"text-combine-upright": TextCombineUprightStyleUtilitiesMapping,

	"border-block-end-style": BorderBlockEndStyleStyleUtilitiesMapping,

	"background-position-x": BackgroundPositionXStyleUtilitiesMapping,

	"border-inline-start-style": BorderInlineStartStyleStyleUtilitiesMapping,

	"justify-self": JustifySelfStyleUtilitiesMapping,

	"padding-left": PaddingLeftStyleUtilitiesMapping,

	"scroll-margin-block-end": ScrollMarginBlockEndStyleUtilitiesMapping,

	"scroll-snap-type-y": ScrollSnapTypeYStyleUtilitiesMapping,

	"initial-letter": InitialLetterStyleUtilitiesMapping,

	"overflow-clip-box": OverflowClipBoxStyleUtilitiesMapping,

	"transform-rotate-z": TransformRotateZStyleUtilitiesMapping,

	"background-image": BackgroundImageStyleUtilitiesMapping,

	"border-top-left-radius": BorderTopLeftRadiusStyleUtilitiesMapping,

	"math-style": MathStyleStyleUtilitiesMapping,

	"text-underline-position": TextUnderlinePositionStyleUtilitiesMapping,

	"baseline-shift": BaselineShiftStyleUtilitiesMapping,

	"border-block-color": BorderBlockColorStyleUtilitiesMapping,

	"border-inline-start-color": BorderInlineStartColorStyleUtilitiesMapping,

	"bottom": BottomStyleUtilitiesMapping,

	"caret-color": CaretColorStyleUtilitiesMapping,

	"column-rule-color": ColumnRuleColorStyleUtilitiesMapping,

	"inset-block": InsetBlockStyleUtilitiesMapping,

	"clip-rule": ClipRuleStyleUtilitiesMapping,

	"masonry-auto-flow": MasonryAutoFlowStyleUtilitiesMapping,

	"overflow-x": OverflowXStyleUtilitiesMapping,

	"scroll-snap-type-x": ScrollSnapTypeXStyleUtilitiesMapping,

	"border-image-width": BorderImageWidthStyleUtilitiesMapping,

	"touch-action": TouchActionStyleUtilitiesMapping,

	"vector-effect": VectorEffectStyleUtilitiesMapping,

	"box-align": BoxAlignStyleUtilitiesMapping,

	"font-smooth": FontSmoothStyleUtilitiesMapping,

	"scroll-behavior": ScrollBehaviorStyleUtilitiesMapping,

	"box-pack": BoxPackStyleUtilitiesMapping,

	"content-visibility": ContentVisibilityStyleUtilitiesMapping,

	"grid-template-rows": GridTemplateRowsStyleUtilitiesMapping,

	"justify-tracks": JustifyTracksStyleUtilitiesMapping,

	"transform-translate-y": TransformTranslateYStyleUtilitiesMapping,

	"column-width": ColumnWidthStyleUtilitiesMapping,

	"image-orientation": ImageOrientationStyleUtilitiesMapping,

	"padding-block": PaddingBlockStyleUtilitiesMapping,

	"text-indent": TextIndentStyleUtilitiesMapping,

	"justify-content": JustifyContentStyleUtilitiesMapping,

	"page-break-before": PageBreakBeforeStyleUtilitiesMapping,

	"color-rendering": ColorRenderingStyleUtilitiesMapping,

	"text-emphasis-color": TextEmphasisColorStyleUtilitiesMapping,

	"padding-right": PaddingRightStyleUtilitiesMapping,

	"perspective-origin": PerspectiveOriginStyleUtilitiesMapping,

	"stroke-width": StrokeWidthStyleUtilitiesMapping,

	"background-clip": BackgroundClipStyleUtilitiesMapping,

	"stroke-opacity": StrokeOpacityStyleUtilitiesMapping,

	"mask-repeat": MaskRepeatStyleUtilitiesMapping,

	"color": ColorStyleUtilitiesMapping,

	"mask-border-repeat": MaskBorderRepeatStyleUtilitiesMapping,

	"overflow": OverflowStyleUtilitiesMapping,

	"overscroll-behavior-inline": OverscrollBehaviorInlineStyleUtilitiesMapping,

	"text-orientation": TextOrientationStyleUtilitiesMapping,

	"appearance": AppearanceStyleUtilitiesMapping,

	"backface-visibility": BackfaceVisibilityStyleUtilitiesMapping,

	"margin-block-start": MarginBlockStartStyleUtilitiesMapping,

	"scroll-margin-right": ScrollMarginRightStyleUtilitiesMapping,

	"text-anchor": TextAnchorStyleUtilitiesMapping,

	"counter-set": CounterSetStyleUtilitiesMapping,

	"grid-row-end": GridRowEndStyleUtilitiesMapping,

	"overscroll-behavior-x": OverscrollBehaviorXStyleUtilitiesMapping,

	"scroll-margin-inline": ScrollMarginInlineStyleUtilitiesMapping,

	"hyphens": HyphensStyleUtilitiesMapping,

	"grid-template-columns": GridTemplateColumnsStyleUtilitiesMapping,

	"counter-increment": CounterIncrementStyleUtilitiesMapping,

	"height": HeightStyleUtilitiesMapping,

	"margin-trim": MarginTrimStyleUtilitiesMapping,

	"stroke-linecap": StrokeLinecapStyleUtilitiesMapping,

	"padding-top": PaddingTopStyleUtilitiesMapping,

	"scrollbar-width": ScrollbarWidthStyleUtilitiesMapping,

	"scroll-padding-top": ScrollPaddingTopStyleUtilitiesMapping,

	"word-break": WordBreakStyleUtilitiesMapping,

	"word-wrap": WordWrapStyleUtilitiesMapping,

	"writing-mode": WritingModeStyleUtilitiesMapping,

	"marker": MarkerStyleUtilitiesMapping,

	"flex-grow": FlexGrowStyleUtilitiesMapping,

	"mix-blend-mode": MixBlendModeStyleUtilitiesMapping,

	"lighting-color": LightingColorStyleUtilitiesMapping,

	"border-image-repeat": BorderImageRepeatStyleUtilitiesMapping,

	"grid-auto-flow": GridAutoFlowStyleUtilitiesMapping,

	"offset-path": OffsetPathStyleUtilitiesMapping,

	"animation-delay": AnimationDelayStyleUtilitiesMapping,

	"box-decoration-break": BoxDecorationBreakStyleUtilitiesMapping,

	"shape-image-threshold": ShapeImageThresholdStyleUtilitiesMapping,

	"transform-skew-y": TransformSkewYStyleUtilitiesMapping,

	"hanging-punctuation": HangingPunctuationStyleUtilitiesMapping,

	"overflow-anchor": OverflowAnchorStyleUtilitiesMapping,

	"font-optical-sizing": FontOpticalSizingStyleUtilitiesMapping,

	"border-block-style": BorderBlockStyleStyleUtilitiesMapping,

	"image-resolution": ImageResolutionStyleUtilitiesMapping,

	"margin-right": MarginRightStyleUtilitiesMapping,

	"mask-border-outset": MaskBorderOutsetStyleUtilitiesMapping,

	"scrollbar-gutter": ScrollbarGutterStyleUtilitiesMapping,

	"text-decoration-skip-ink": TextDecorationSkipInkStyleUtilitiesMapping,

	"transition-timing-function": TransitionTimingFunctionStyleUtilitiesMapping,

	"font-language-override": FontLanguageOverrideStyleUtilitiesMapping,

	"padding-inline": PaddingInlineStyleUtilitiesMapping,

	"perspective": PerspectiveStyleUtilitiesMapping,

	"scale-right": ScaleRightStyleUtilitiesMapping,

	"border-block-end-width": BorderBlockEndWidthStyleUtilitiesMapping,

	"background-origin": BackgroundOriginStyleUtilitiesMapping,

	"border-block-start-style": BorderBlockStartStyleStyleUtilitiesMapping,

	"clip": ClipStyleUtilitiesMapping,

	"offset-anchor": OffsetAnchorStyleUtilitiesMapping,

	"box-flex-group": BoxFlexGroupStyleUtilitiesMapping,

	"font-variation-settings": FontVariationSettingsStyleUtilitiesMapping,

	"mask-size": MaskSizeStyleUtilitiesMapping,

	"min-block-size": MinBlockSizeStyleUtilitiesMapping,

	"offset-rotate": OffsetRotateStyleUtilitiesMapping,

	"white-space": WhiteSpaceStyleUtilitiesMapping,

	"column-span": ColumnSpanStyleUtilitiesMapping,

	"list-style-position": ListStylePositionStyleUtilitiesMapping,

	"offset-position": OffsetPositionStyleUtilitiesMapping,

	"border-end-start-radius": BorderEndStartRadiusStyleUtilitiesMapping,

	"border-right-style": BorderRightStyleStyleUtilitiesMapping,

	"background-attachment": BackgroundAttachmentStyleUtilitiesMapping,

	"background-size": BackgroundSizeStyleUtilitiesMapping,

	"float": FloatStyleUtilitiesMapping,

	"gap": GapStyleUtilitiesMapping,

	"max-height": MaxHeightStyleUtilitiesMapping,

	"text-justify": TextJustifyStyleUtilitiesMapping,

	"text-size-adjust": TextSizeAdjustStyleUtilitiesMapping,

	"margin-block-end": MarginBlockEndStyleUtilitiesMapping,

	"border-bottom-style": BorderBottomStyleStyleUtilitiesMapping,

	"box-shadow": BoxShadowStyleUtilitiesMapping,

	"transform-rotate": TransformRotateStyleUtilitiesMapping,

	"animation-timing-function": AnimationTimingFunctionStyleUtilitiesMapping,

	"flex-direction": FlexDirectionStyleUtilitiesMapping,

	"margin-block": MarginBlockStyleUtilitiesMapping,

	"margin-inline-end": MarginInlineEndStyleUtilitiesMapping,

	"fill-opacity": FillOpacityStyleUtilitiesMapping,

	"text-emphasis-style": TextEmphasisStyleStyleUtilitiesMapping,

	"column-count": ColumnCountStyleUtilitiesMapping,

	"mask-border-slice": MaskBorderSliceStyleUtilitiesMapping,

	"scroll-padding-block-start": ScrollPaddingBlockStartStyleUtilitiesMapping,

	"transform-style": TransformStyleStyleUtilitiesMapping,

	"grid-auto-columns": GridAutoColumnsStyleUtilitiesMapping,

	"outline-width": OutlineWidthStyleUtilitiesMapping,

	"row-gap": RowGapStyleUtilitiesMapping,

	"scroll-padding-left": ScrollPaddingLeftStyleUtilitiesMapping,

	"scroll-margin-block-start": ScrollMarginBlockStartStyleUtilitiesMapping,

	"text-align": TextAlignStyleUtilitiesMapping,

	"border-right-width": BorderRightWidthStyleUtilitiesMapping,

	"border-block-width": BorderBlockWidthStyleUtilitiesMapping,

	"min-height": MinHeightStyleUtilitiesMapping,

	"animation-duration": AnimationDurationStyleUtilitiesMapping,

	"max-block-size": MaxBlockSizeStyleUtilitiesMapping,

	"max-width": MaxWidthStyleUtilitiesMapping,

	"scroll-margin-left": ScrollMarginLeftStyleUtilitiesMapping,

	"scroll-snap-type": ScrollSnapTypeStyleUtilitiesMapping,

	"direction": DirectionStyleUtilitiesMapping,

	"overflow-inline": OverflowInlineStyleUtilitiesMapping,

	"scroll-snap-points-x": ScrollSnapPointsXStyleUtilitiesMapping,

	"text-underline-offset": TextUnderlineOffsetStyleUtilitiesMapping,

	"border-bottom-left-radius": BorderBottomLeftRadiusStyleUtilitiesMapping,

	"border-spacing": BorderSpacingStyleUtilitiesMapping,

	"font-variant-ligatures": FontVariantLigaturesStyleUtilitiesMapping,

	"padding-inline-start": PaddingInlineStartStyleUtilitiesMapping,

	"border-block-start-width": BorderBlockStartWidthStyleUtilitiesMapping,

	"top": TopStyleUtilitiesMapping,

	"mask-border-mode": MaskBorderModeStyleUtilitiesMapping,

	"scroll-margin": ScrollMarginStyleUtilitiesMapping,

	"scroll-margin-inline-start": ScrollMarginInlineStartStyleUtilitiesMapping,

	"text-shadow": TextShadowStyleUtilitiesMapping,

	"counter-reset": CounterResetStyleUtilitiesMapping,

	"vertical-align": VerticalAlignStyleUtilitiesMapping,

	"grid-template-areas": GridTemplateAreasStyleUtilitiesMapping,

	"page-break-after": PageBreakAfterStyleUtilitiesMapping,

	"scroll-snap-points-y": ScrollSnapPointsYStyleUtilitiesMapping,

	"font-style": FontStyleStyleUtilitiesMapping,

	"overscroll-behavior-y": OverscrollBehaviorYStyleUtilitiesMapping,

	"scroll-padding-block-end": ScrollPaddingBlockEndStyleUtilitiesMapping,

	"shape-rendering": ShapeRenderingStyleUtilitiesMapping,

	"resize": ResizeStyleUtilitiesMapping,

	"z-index": ZIndexStyleUtilitiesMapping,

	"block-overflow": BlockOverflowStyleUtilitiesMapping,

	"border-image-outset": BorderImageOutsetStyleUtilitiesMapping,

	"line-height-step": LineHeightStepStyleUtilitiesMapping,

	"transform-skew": TransformSkewStyleUtilitiesMapping,

	"border-left-width": BorderLeftWidthStyleUtilitiesMapping,

	"padding-inline-end": PaddingInlineEndStyleUtilitiesMapping,

	"left": LeftStyleUtilitiesMapping,

	"will-change": WillChangeStyleUtilitiesMapping,

	"letter-spacing": LetterSpacingStyleUtilitiesMapping,

	"box-flex": BoxFlexStyleUtilitiesMapping,

	"border-image-slice": BorderImageSliceStyleUtilitiesMapping,

	"box-lines": BoxLinesStyleUtilitiesMapping,

	"caption-side": CaptionSideStyleUtilitiesMapping,

	"mask-mode": MaskModeStyleUtilitiesMapping,

	"font-variant": FontVariantStyleUtilitiesMapping,

	"scroll-margin-inline-end": ScrollMarginInlineEndStyleUtilitiesMapping,

	"stroke-linejoin": StrokeLinejoinStyleUtilitiesMapping,

	"rotate": RotateStyleUtilitiesMapping,

	"transform-scale-3d": TransformScale3dStyleUtilitiesMapping,

	"cursor": CursorStyleUtilitiesMapping,

	"box-direction": BoxDirectionStyleUtilitiesMapping,

	"grid-column-end": GridColumnEndStyleUtilitiesMapping,

	"user-select": UserSelectStyleUtilitiesMapping,

	"border-inline-start-width": BorderInlineStartWidthStyleUtilitiesMapping,

	"border-inline-color": BorderInlineColorStyleUtilitiesMapping,

	"border-top-width": BorderTopWidthStyleUtilitiesMapping,

	"flex-shrink": FlexShrinkStyleUtilitiesMapping,

	"isolation": IsolationStyleUtilitiesMapping,

	"mask-type": MaskTypeStyleUtilitiesMapping,

	"backdrop-filter": BackdropFilterStyleUtilitiesMapping,

	"animation-play-state": AnimationPlayStateStyleUtilitiesMapping,

	"border-inline-end-style": BorderInlineEndStyleStyleUtilitiesMapping,

	"ruby-align": RubyAlignStyleUtilitiesMapping,

	"scroll-snap-align": ScrollSnapAlignStyleUtilitiesMapping,

	"glyph-orientation-vertical": GlyphOrientationVerticalStyleUtilitiesMapping,

	"font-family": FontFamilyStyleUtilitiesMapping,

	"inset-inline-end": InsetInlineEndStyleUtilitiesMapping,

	"scroll-padding-inline-start": ScrollPaddingInlineStartStyleUtilitiesMapping,

	"border-inline-end-color": BorderInlineEndColorStyleUtilitiesMapping,

	"font-variant-numeric": FontVariantNumericStyleUtilitiesMapping,

	"inset": InsetStyleUtilitiesMapping,

	"align-content": AlignContentStyleUtilitiesMapping,

	"transform-scale": TransformScaleStyleUtilitiesMapping,

	"scroll-snap-coordinate": ScrollSnapCoordinateStyleUtilitiesMapping,

	"azimuth": AzimuthStyleUtilitiesMapping,

	"align-self": AlignSelfStyleUtilitiesMapping,

	"block-size": BlockSizeStyleUtilitiesMapping,

	"opacity": OpacityStyleUtilitiesMapping,

	"break-inside": BreakInsideStyleUtilitiesMapping,

	"transform-scale-y": TransformScaleYStyleUtilitiesMapping,

	"font-weight": FontWeightStyleUtilitiesMapping,

	"overflow-block": OverflowBlockStyleUtilitiesMapping,

	"text-overflow": TextOverflowStyleUtilitiesMapping,

	"font-size": FontSizeStyleUtilitiesMapping,

	"border-collapse": BorderCollapseStyleUtilitiesMapping,

	"contain": ContainStyleUtilitiesMapping,

	"orphans": OrphansStyleUtilitiesMapping,

	"border-top-style": BorderTopStyleStyleUtilitiesMapping,

	"justify-items": JustifyItemsStyleUtilitiesMapping,

	"margin-inline": MarginInlineStyleUtilitiesMapping,

	"object-fit": ObjectFitStyleUtilitiesMapping,

	"background-position-y": BackgroundPositionYStyleUtilitiesMapping,

	"offset-distance": OffsetDistanceStyleUtilitiesMapping,

	"scroll-padding-bottom": ScrollPaddingBottomStyleUtilitiesMapping,

	"border-right-color": BorderRightColorStyleUtilitiesMapping,

	"max-inline-size": MaxInlineSizeStyleUtilitiesMapping,

	"right": RightStyleUtilitiesMapping,

	"scale-left": ScaleLeftStyleUtilitiesMapping,

	"mask-image": MaskImageStyleUtilitiesMapping,

	"object-position": ObjectPositionStyleUtilitiesMapping,

	"scroll-padding-inline": ScrollPaddingInlineStyleUtilitiesMapping,

	"box-orient": BoxOrientStyleUtilitiesMapping,
}
