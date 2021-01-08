package base

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPalette(t *testing.T) {
	var lightList, darkList, perr = CreatePalettes(HuePalette{
		Steps:            10,
		HueStart:         100,
		HueEnd:           359,
		HueCurve:         EaseInCubicTCurve,
		SaturationRate:   30,
		SaturationStart:  10,
		SaturationEnd:    80,
		SaturationCurve:  EaseInCubicTCurve,
		LuminousityStart: 5,
		LuminousityEnd:   80,
		LuminosityCurve:  EaseInCubicTCurve,
	})
	require.NoError(t, perr)
	require.Len(t, lightList, 10)
	require.Len(t, darkList, 10)
}

func ExamplePalette() {
	var light, dark, perr = CreatePalettes(HuePalette{
		Steps:            10,
		HueStart:         100,
		HueEnd:           359,
		HueCurve:         EaseInCubicTCurve,
		SaturationRate:   30,
		SaturationStart:  10,
		SaturationEnd:    80,
		SaturationCurve:  EaseInCubicTCurve,
		LuminousityStart: 5,
		LuminousityEnd:   80,
		LuminosityCurve:  EaseInCubicTCurve,
	})

	if perr != nil {
		log.Fatalf("Error: %s\n", perr)
	}

	var lightJSON, err = json.Marshal(light)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(lightJSON)

	var darkJSON, err2 = json.Marshal(dark)
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	fmt.Println(darkJSON)
}
