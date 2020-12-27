package styled

import (
	json2 "encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalette(t *testing.T) {
	var palette = CreatePalette(HuePalette{
		Steps:              10,
		HueStart:           100,
		HueEnd:             359,
		HueCurve:           EaseInCubicTCurve,
		SaturationRate:     30,
		SaturationStart:    10,
		SaturationEnd:      80,
		SaturationCurve:    EaseInCubicTCurve,
		LuminousityStart:   5,
		LuminousityEnd:     80,
		LuminosityCurve:    EaseInCubicTCurve,
		ContrastColorLight: nil,
		ContrastColorDark:  nil,
	})
	require.NotNil(t, palette.DarkMode)
	require.NotNil(t, palette.LightMode)
	require.Len(t, palette.LightMode, 10)
	require.Len(t, palette.DarkMode, 10)
}

func ExamplePalette() {
	var palette = CreatePalette(HuePalette{
		Steps:              10,
		HueStart:           100,
		HueEnd:             359,
		HueCurve:           EaseInCubicTCurve,
		SaturationRate:     30,
		SaturationStart:    10,
		SaturationEnd:      80,
		SaturationCurve:    EaseInCubicTCurve,
		LuminousityStart:   5,
		LuminousityEnd:     80,
		LuminosityCurve:    EaseInCubicTCurve,
		ContrastColorLight: nil,
		ContrastColorDark:  nil,
	})

	var json, err = json2.Marshal(palette)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(json)
}
