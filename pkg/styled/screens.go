package styled

type Unit int

const (
	Pixels Unit = iota
	Rems
	Ems
	Vw // viewport's width
	Vh // viewport's height
)

// Screens represents measures/scale of supported screen sizes.
var DefaultScreenScales = StringScale("sm", "md", "lg", "xl", "2xl")
var DefaultScreenValues = NewStringNumber(DefaultScreenScales).
	Key("sm", 30).
	Key("md", 48).
	Key("lg", 62).
	Key("xl", 80)

var TailwindScreenValues = NewStringNumber(DefaultScreenScales).
	Key("sm", 640).
	Key("md", 768).
	Key("lg", 1024).
	Key("xl", 1280).
	Key("2xl", 1280)

func DefaultScreens() *ScreenDirective {
	return NewScreenDirective(Ems, DefaultScreenValues)
}

type ScreenDirective struct {
	Unit   Unit
	Values *StringNumber
}

func NewScreenDirective(unit Unit, values *StringNumber) *ScreenDirective {
	return &ScreenDirective{
		Unit:   unit,
		Values: values,
	}
}
