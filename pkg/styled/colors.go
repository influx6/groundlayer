package styled

// Colors represents measure/scale of possible color variants.
var DefaultColorScale = NumberScale(50, 100, 200, 300, 400, 500, 600, 700, 800, 900)

func ColorFromDefaultScale(values ...Palette) *NumberPalette {
	if len(values) < DefaultColorScale.Count() {
		panic("color count can't be less than scale")
	}

	var set = NewNumberPalette(DefaultColorScale)
	for i := 0; i < DefaultColorScale.Count(); i++ {
		set.Index(i, values[i])
	}
	return set
}

type ColorDirective struct {
	Light *NumberPalette
	Dark  *NumberPalette
}

func NewColorDirective(light *NumberPalette, dark *NumberPalette) *ColorDirective {
	return &ColorDirective{
		Light: light,
		Dark:  dark,
	}
}
