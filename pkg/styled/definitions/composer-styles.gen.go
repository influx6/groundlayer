package definitions

// These are styles which define how a giving parent style will be compose from it's
// utilities which break down a style into individual bits.
var ParentCompositionStyles = map[string][]string{

	"transition": {
		"transition-delay",
		"transition-duration",
		"transition-property",
		"transition-timing-function",
	},

	"mask-border": {
		"mask-border-mode",
		"mask-border-outset",
		"mask-border-repeat",
		"mask-border-slice",
		"mask-border-source",
		"mask-border-width",
	},

	"margin": {
		"margin-bottom",
		"margin-left",
		"margin-right",
		"margin-top",
	},

	"grid-template": {
		"grid-template-columns",
		"grid-template-rows",
		"grid-template-areas",
	},

	"grid": {
		"grid-template-rows",
		"grid-template-columns",
		"grid-template-areas",
		"grid-auto-rows",
		"grid-auto-columns",
		"grid-auto-flow",
		"grid-column-gap",
		"grid-row-gap",
		"column-gap",
		"row-gap",
	},

	"grid-area": {
		"grid-row-start",
		"grid-column-start",
		"grid-row-end",
		"grid-column-end",
	},

	"text-emphasis": {
		"text-emphasis-style",
		"text-emphasis-color",
	},

	"border-radius": {
		"border-top-left-radius",
		"border-top-right-radius",
		"border-bottom-right-radius",
		"border-bottom-left-radius",
	},

	"columns": {
		"column-width",
		"column-count",
	},

	"column-rule": {
		"column-rule-width",
		"column-rule-style",
		"column-rule-color",
	},

	"border-width": {
		"border-top-width",
		"border-right-width",
		"border-bottom-width",
		"border-left-width",
	},

	"-ms-content-zoom-snap": {
		"-ms-content-zoom-snap-type",
		"-ms-content-zoom-snap-points",
	},

	"outline": {
		"outline-color",
		"outline-style",
		"outline-width",
	},

	"mask": {
		"mask-image",
		"mask-mode",
		"mask-repeat",
		"mask-position",
		"mask-clip",
		"mask-origin",
		"mask-size",
		"mask-composite",
	},

	"-webkit-mask": {
		"mask-image",
		"mask-mode",
		"mask-repeat",
		"mask-position",
		"mask-clip",
		"mask-origin",
		"mask-size",
		"mask-composite",
	},

	"list-style": {
		"list-style-type",
		"list-style-position",
		"list-style-image",
	},

	"text-emphasis-position": {
		"text-emphasis-position-first text-emphasis-position-second",
	},

	"border": {
		"border-width",
		"border-style",
		"border-color",
	},

	"border-left": {
		"border-left-width",
		"border-left-style",
		"border-left-color",
	},

	"transform": {
		"transform-translate-x",
		"transform-translate-y",
		"transform-scale",
		"transform-scale-x",
		"transform-scale-y",
		"transform-rotate",
		"transform-skew",
		"transform-skew-x",
		"transform-skew-y",
		"transform-matrix3d",
		"transform-translate-z",
		"transform-scale3d",
		"transform-scale-z",
		"transform-rotate-x",
		"transform-rotate-y",
		"transform-rotate-z",
		"transform-perspective",
	},

	"padding": {
		"padding-bottom",
		"padding-left",
		"padding-right",
		"padding-top",
	},

	"border-block-end": {
		"border-top-width",
		"border-top-style",
		"border-top-color",
	},

	"border-inline": {
		"border-top-width",
		"border-top-style",
		"border-top-color",
	},

	"place-self": {
		"align-self",
		"justify-self",
	},

	"border-bottom": {
		"border-bottom-width",
		"border-bottom-style",
		"border-bottom-color",
	},

	"border-inline-start": {
		"border-width",
		"border-style",
		"color",
	},

	"flex-flow": {
		"flex-direction",
		"flex-wrap",
	},

	"background": {
		"background-image",
		"background-position",
		"background-size",
		"background-repeat",
		"background-origin",
		"background-clip",
		"background-attachment",
		"background-color",
	},

	"flex": {
		"flex-grow",
		"flex-shrink",
		"flex-basis",
	},

	"-ms-content-zoom-limit": {},

	"scale": {
		"scale-left",
		"scale-right",
	},

	"place-items": {
		"align-items",
		"justify-items",
	},

	"grid-gap": {
		"grid-row-gap",
		"grid-column-gap",
	},

	"border-right": {
		"border-right-width",
		"border-right-style",
		"border-right-color",
	},

	"border-style": {
		"border-top-style",
		"border-right-style",
		"border-bottom-style",
		"border-left-style",
	},

	"border-color": {
		"border-top-color",
		"border-right-color",
		"border-bottom-color",
		"border-left-color",
	},

	"grid-row": {
		"grid-row-start",
		"grid-row-end",
	},

	"grid-column": {
		"grid-column-start",
		"grid-column-end",
	},

	"font": {
		"font-style",
		"font-variant",
		"font-weight",
		"font-stretch",
		"font-size",
		"line-height",
		"font-family",
	},

	"border-image": {
		"border-image-source",
		"border-image-slice",
		"border-image-width",
		"border-image-outset",
		"border-image-repeat",
	},

	"text-decoration": {
		"text-decoration-color",
		"text-decoration-style",
		"text-decoration-line",
	},

	"offset": {
		"offset-position",
		"offset-path",
		"offset-distance",
		"offset-anchor",
		"offset-rotate",
	},

	"animation": {
		"animation-name",
		"animation-duration",
		"animation-timing-function",
		"animation-delay",
		"animation-iteration-count",
		"animation-direction",
		"animation-fill-mode",
		"animation-play-state",
	},

	"border-block": {
		"border-top-width",
		"border-top-style",
		"border-top-color",
	},

	"border-block-start": {
		"border-width",
		"border-style",
		"color",
	},

	"border-inline-end": {
		"border-width",
		"border-style",
		"color",
	},

	"border-top": {
		"border-top-width",
		"border-top-style",
		"border-top-color",
	},
}
