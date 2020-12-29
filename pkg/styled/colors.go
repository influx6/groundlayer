package styled

// Colors represents measure/scale of possible color variants.
var DefaultColorScale = NumberScale(50, 100, 200, 300, 400, 500, 600, 700, 800, 900)

type ColorDirective struct {
	Scale *StringPalette
}

func NewColorDirective() *ColorDirective {
	return &ColorDirective{}
}
