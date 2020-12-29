package styled

type ThemeDirective []string

func (td *ThemeDirective) Add(t string) {
	*td = append(*td, t)
}
