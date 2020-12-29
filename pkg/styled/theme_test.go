package styled

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThemeDirective(t *testing.T) {
	var themes = ThemeDirective{}
	themes.Add("color-red-100")
	themes.Add("color-red-10")

	require.Len(t, themes, 2)
	require.Contains(t, themes, "color-red-100", "colors-red-10")
}
