package styled

var (
	// Screens represents measures/scale of supported screen sizes.
	Screens = StringScale("sm", "md", "lg", "xl")

	// Colors represents measure/scale of possible color variants.
	Colors = NumberScale(100, 200, 300, 400, 500, 600, 700, 800, 900)
)
