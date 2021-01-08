package styled

// These are styles which define how a giving parent style will be compose from it's
// utilities which break down a style into individual bits.
var ParentCompositionStyles = map[string][]string{

	"place-items": {
		"align-items",
		"justify-items",
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

	"font": {
		"font-style",
		"font-variant",
		"font-weight",
		"font-stretch",
		"font-size",
		"line-height",
		"font-family",
	},

	"border-inline-start": {
		"border-width",
		"border-style",
		"color",
	},

	"border-top": {
		"border-top-width",
		"border-top-style",
		"border-top-color",
	},

	"column-rule": {
		"column-rule-width",
		"column-rule-style",
		"column-rule-color",
	},

	"flex-flow": {
		"flex-direction",
		"flex-wrap",
	},

	"place-self": {
		"align-self",
		"justify-self",
	},

	"margin": {
		"margin-bottom",
		"margin-left",
		"margin-right",
		"margin-top",
	},

	"text-emphasis": {
		"text-emphasis-style",
		"text-emphasis-color",
	},

	"text-emphasis-position": {
		"text-emphasis-position-first text-emphasis-position-second",
	},

	"border-radius": {
		"border-top-left-radius",
		"border-top-right-radius",
		"border-bottom-right-radius",
		"border-bottom-left-radius",
	},

	"border-right": {
		"border-right-width",
		"border-right-style",
		"border-right-color",
	},

	"padding": {
		"padding-bottom",
		"padding-left",
		"padding-right",
		"padding-top",
	},

	"grid-template": {
		"grid-template-columns",
		"grid-template-rows",
		"grid-template-areas",
	},

	"border-width": {
		"border-top-width",
		"border-right-width",
		"border-bottom-width",
		"border-left-width",
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

	"mask-border": {
		"mask-border-mode",
		"mask-border-outset",
		"mask-border-repeat",
		"mask-border-slice",
		"mask-border-source",
		"mask-border-width",
	},

	"border-inline-end": {
		"border-width",
		"border-style",
		"color",
	},

	"border-block-end": {
		"border-top-width",
		"border-top-style",
		"border-top-color",
	},

	"border-bottom": {
		"border-bottom-width",
		"border-bottom-style",
		"border-bottom-color",
	},

	"text-decoration": {
		"text-decoration-color",
		"text-decoration-style",
		"text-decoration-line",
	},

	"scale": {
		"scale-left",
		"scale-right",
	},

	"outline": {
		"outline-color",
		"outline-style",
		"outline-width",
	},

	"grid-gap": {
		"grid-row-gap",
		"grid-column-gap",
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

	"border-color": {
		"border-top-color",
		"border-right-color",
		"border-bottom-color",
		"border-left-color",
	},

	"border-inline": {
		"border-top-width",
		"border-top-style",
		"border-top-color",
	},

	"border-style": {
		"border-top-style",
		"border-right-style",
		"border-bottom-style",
		"border-left-style",
	},

	"-ms-content-zoom-snap": {
		"-ms-content-zoom-snap-type",
		"-ms-content-zoom-snap-points",
	},

	"-ms-content-zoom-limit": {},

	"transition": {
		"transition-delay",
		"transition-duration",
		"transition-property",
		"transition-timing-function",
	},

	"grid-area": {
		"grid-row-start",
		"grid-column-start",
		"grid-row-end",
		"grid-column-end",
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

	"border-left": {
		"border-left-width",
		"border-left-style",
		"border-left-color",
	},

	"list-style": {
		"list-style-type",
		"list-style-position",
		"list-style-image",
	},

	"grid-column": {
		"grid-column-start",
		"grid-column-end",
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

	"border-image": {
		"border-image-source",
		"border-image-slice",
		"border-image-width",
		"border-image-outset",
		"border-image-repeat",
	},

	"flex": {
		"flex-grow",
		"flex-shrink",
		"flex-basis",
	},

	"offset": {
		"offset-position",
		"offset-path",
		"offset-distance",
		"offset-anchor",
		"offset-rotate",
	},

	"grid-row": {
		"grid-row-start",
		"grid-row-end",
	},

	"border": {
		"border-width",
		"border-style",
		"border-color",
	},

	"columns": {
		"column-width",
		"column-count",
	},
}
