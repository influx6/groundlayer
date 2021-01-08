package styled

import (
	"github.com/influx6/groundlayer/pkg/styled/base"
)

// Screens represents measures/scale of supported screen sizes.
var DefaultScreenScales = base.StringScale("sm", "md", "lg", "xl", "2xl")

var ChakraUIScreens = base.NewString(DefaultScreenScales).
	Key("sm", `@media (min-width: 30em)`).
	Key("md", `@media (min-width: 48em)`).
	Key("lg", `@media (min-width: 62em)`).
	Key("xl", `@media (min-width: 80em)`).
	Key("2xl", `@media (min-width: 94em)`)

var TailwindUIScreens = base.NewString(DefaultScreenScales).
	Key("sm", `@media (min-width: 640px)`).
	Key("md", `@media (min-width: 768px)`).
	Key("lg", `@media (min-width: 1024px)`).
	Key("xl", `@media (min-width: 1280px)`).
	Key("2xl", `@media (min-width: 1536px)`)
