package base

import (
	"github.com/influx6/npkg/nerror"
	gocolorful "github.com/lucasb-eyer/go-colorful"
	"gonum.org/v1/plot/vg"
)

var (
	origin = []float64{0, 1}
	white  = gocolorful.Color{R: 1, G: 1, B: 1}
	black  = gocolorful.Color{R: 0, G: 0, B: 0}
)

type Curve interface {
	Point(t float64) float64
}

type XYCurve interface {
	Point(t float64) vg.Point
}

// HuePalette implements the color algorithm defined in the lyft color algorithm
// which provides a more concise way of generating colors for specific hues, saturation
// and luminousity levels that match how we perceive colors.
//
// See - [Re-Approaching Color by Lyft Design](https://design.lyft.com/re-approaching-color-9e604ba22c88)
// and [Life Color Algorithm](https://github.com/lyft/coloralgorithm).
//
type HuePalette struct {
	Steps              int              // number of colors to be generated
	HueStart           int              // Hue value (0-360)
	HueEnd             int              // Hue value (0-360)
	HueCurve           Curve            // Hue bezier curve points
	SaturationRate     float64          // rate of saturation increase (0-200)
	SaturationStart    int              // Saturation value (0-100)
	SaturationEnd      int              // Saturation value (0-100)
	SaturationCurve    Curve            // Bezier curve points for saturation
	LuminousityStart   int              // Luminous value (0-100)
	LuminousityEnd     int              // Luminous value (0-100)
	LuminosityCurve    Curve            // Luminosity bezier curve points
	ContrastColorLight gocolorful.Color // Provides light color contrast to be used for calculating contrast
	ContrastColorDark  gocolorful.Color // Provides dark color contrast to be used for calculating contrast
}

func (h *HuePalette) ensure() error {
	if h.Steps <= 0 {
		return nerror.New("HuePalette.Steps is required and should be above 0")
	}
	if h.LuminosityCurve == nil {
		return nerror.New("HuePalette.LuminosityCurve is required")
	}
	if h.HueCurve == nil {
		return nerror.New("HuePalette.HueCurve is required")
	}
	if h.SaturationCurve == nil {
		return nerror.New("HuePalette.SaturationCurve is required")
	}
	if h.SaturationRate <= 0 {
		return nerror.New("HuePalette.SaturationRate is required and should be above 0")
	}
	if h.SaturationStart <= 0 {
		return nerror.New("HuePalette.SaturationStart is required and should be above 0")
	}
	if h.SaturationEnd > 100 {
		return nerror.New("HuePalette.SaturationEnd is required and should be below or equal to 100")
	}
	if h.LuminousityStart <= 0 {
		return nerror.New("HuePalette.LuminousityStart is required and should be above 0")
	}
	if h.LuminousityEnd > 100 {
		return nerror.New("HuePalette.LuminousityEnd is required and should be below or equal to 100")
	}
	if h.HueStart <= 0 {
		return nerror.New("HuePalette.HueStart is required and should be above 0")
	}
	if h.HueEnd >= 360 {
		return nerror.New("HuePalette.HueEnd is required and should be below 360")
	}
	return nil
}

func Hex(hex string) (HuePalette, error) {
	var config HuePalette

	var hexColor, hexErr = gocolorful.Hex(hex)
	if hexErr != nil {
		return config, nerror.WrapOnly(hexErr)
	}

	var hexH, hexS, hexL = hexColor.Hsl()

	config.Steps = 10
	config.HueStart = int(hexH)
	config.HueEnd = 360
	config.HueCurve = EaseInCubicTCurve
	config.SaturationStart = int(hexS * 100)
	config.SaturationRate = 20
	config.SaturationEnd = 100
	config.SaturationCurve = EaseInCubicTCurve
	config.LuminousityStart = int(hexL)
	config.LuminousityEnd = 1
	config.LuminosityCurve = EaseInCubicTCurve
	return config, nil
}

// Palette defines the generated color set which is produced from
// provided HuePalette.
type Palette struct {
	Step          int
	Steps         int
	Hex           string
	RGB           gocolorful.Color
	HCL           gocolorful.Color
	HSV           gocolorful.Color
	HSL           gocolorful.Color
	ContrastDark  float64
	ContrastLight float64
	HueRange      []int
	Hue           float64
	Saturation    float64
	Luminosity    float64
}

func MustUserPalettes(hexColors ...string) []Palette {
	var pls, err = UserPalettes(hexColors...)
	if err != nil {
		panic(err.Error())
	}
	return pls
}

func UserPalettes(hexColors ...string) ([]Palette, error) {
	var palettes = make([]Palette, 0, len(hexColors))
	for index, hex := range hexColors {
		var ul, err = UserPalette(hex, index, len(hexColors))
		if err != nil {
			return palettes, nerror.WrapOnly(err)
		}
		palettes = append(palettes, ul)
	}
	return palettes, nil
}

func MustUserPalette(hex string, step int, steps int) Palette {
	var upl, err = UserPalette(hex, step, steps)
	if err != nil {
		panic(err)
	}
	return upl
}

func UserPalette(hex string, step int, steps int) (Palette, error) {
	var config Palette

	var hexColor, hexErr = gocolorful.Hex(hex)
	if hexErr != nil {
		return config, nerror.WrapOnly(hexErr)
	}

	var hexHV, hexSV, hexV = hexColor.Hsv()
	config.HSV = gocolorful.Hsl(hexHV, hexSV, hexV)

	var hexH, hexS, hexL = hexColor.Hsl()
	config.HSL = gocolorful.Hsl(hexH, hexS, hexL)

	var hexHL, hexCL, hexLL = hexColor.Hcl()
	config.HCL = gocolorful.Hcl(hexHL, hexCL, hexLL)

	config.RGB = gocolorful.Color{
		R: hexColor.R,
		G: hexColor.G,
		B: hexColor.B,
	}

	config.ContrastLight = getColorContrastFor(config.HCL, white)
	config.ContrastDark = getColorContrastFor(config.HCL, black)

	config.Steps = steps
	config.Step = step
	config.Hex = hex
	config.HueRange = []int{int(hexH), 360}
	config.Hue = hexH
	config.Saturation = hexS
	config.Luminosity = hexL
	return config, nil
}

type Palettes struct {
	LightMode []Palette
	DarkMode  []Palette // will be a list of LightMode palette set reversed.
}

func MustCreatePalettes(config HuePalette) (lightPalettes []Palette, darkPalettes []Palette, err error) {
	lightPalettes, darkPalettes, err = CreatePalettes(config)
	if err != nil {
		panic(err.Error())
	}
	return
}

func CreatePalettes(config HuePalette) (lightPalettes []Palette, darkPalettes []Palette, err error) {
	if ensureErr := config.ensure(); ensureErr != nil {
		err = nerror.WrapOnly(ensureErr)
		return
	}

	var lightContrast = config.ContrastColorLight
	if isDefaultValue(lightContrast) {
		lightContrast = white
	}

	var darkContrast = config.ContrastColorDark
	if isDefaultValue(darkContrast) {
		darkContrast = black
	}

	var huePoints = getStepCurvePoints(config.HueCurve, config.Steps)
	var luminosityPoints = getStepCurvePoints(config.LuminosityCurve, config.Steps)
	var saturationPoints = getStepCurvePoints(config.SaturationCurve, config.Steps)

	var hueCount = len(huePoints)
	var hueValues = make([]float64, hueCount)
	for index := 0; index < hueCount; index++ {
		var step = huePoints[index]
		hueValues[hueCount-index-1] = distributeValue(
			step,
			origin,
			[]float64{
				float64(config.HueStart),
				float64(config.HueEnd),
			},
		)
	}

	var hueDarkValues = make([]float64, hueCount)
	for index := hueCount - 1; index > -1; index-- {
		var step = huePoints[index]
		hueDarkValues[hueCount-index-1] = distributeValue(
			step,
			origin,
			[]float64{
				float64(config.HueEnd),
				float64(config.HueStart),
			},
		)
	}

	var saturatedCount = len(saturationPoints)
	var saturatedValues = make([]float64, saturatedCount)
	for index := 0; index < saturatedCount; index++ {
		var step = saturationPoints[index]
		saturatedValues[saturatedCount-index-1] = distributeValue(
			step,
			origin,
			[]float64{
				float64(config.SaturationStart) * 0.01,
				float64(config.SaturationEnd) * 0.01,
			},
		) * config.SaturationRate
	}

	var saturatedDarkValues = make([]float64, saturatedCount)
	for index := saturatedCount - 1; index > -1; index-- {
		var step = saturationPoints[index]
		saturatedDarkValues[saturatedCount-index-1] = distributeValue(
			step,
			origin,
			[]float64{
				float64(config.SaturationEnd) * 0.01,
				float64(config.SaturationStart) * 0.01,
			},
		) * config.SaturationRate
	}

	var luminousCount = len(luminosityPoints)
	var luminousValues = make([]float64, luminousCount)
	for index := 0; index < luminousCount; index++ {
		var step = luminosityPoints[index]
		luminousValues[luminousCount-index-1] = distributeValue(
			step,
			origin,
			[]float64{
				float64(config.LuminousityEnd) * 0.01,
				float64(config.LuminousityStart) * 0.01,
			},
		)
	}

	var luminousDarkValues = make([]float64, luminousCount)
	for index := luminousCount - 1; index > -1; index-- {
		var step = luminosityPoints[index]
		luminousDarkValues[luminousCount-index-1] = distributeValue(
			step,
			origin,
			[]float64{
				float64(config.LuminousityEnd) * 0.01,
				float64(config.LuminousityStart) * 0.01,
			},
		)
	}

	lightPalettes = make([]Palette, 0, config.Steps)
	darkPalettes = make([]Palette, 0, config.Steps)

	// use the luminosity as base for palette color list.
	for index := 0; index < luminousCount; index++ {
		var lightPalette Palette
		lightPalette.Step = index + 1
		lightPalette.Steps = config.Steps
		lightPalette.Hue = hueValues[index]
		lightPalette.Luminosity = luminousValues[index]
		lightPalette.Saturation = saturatedValues[index]
		lightPalette.HueRange = []int{config.HueStart, config.HueEnd}

		if lightPalette.Saturation > 1 {
			lightPalette.Saturation = 1
		}

		lightPalette.HCL = gocolorful.Hcl(lightPalette.Hue, lightPalette.Saturation, lightPalette.Luminosity)
		lightPalette.Hex = lightPalette.HCL.Hex()

		var h, s, v = lightPalette.HCL.Hsv()
		lightPalette.HSV = gocolorful.Hsv(h, s, v)

		var h1, s1, l = lightPalette.HCL.Hsl()
		lightPalette.HSL = gocolorful.Hsl(h1, s1, l)

		var r, g, b = lightPalette.HCL.LinearRgb()
		lightPalette.RGB = gocolorful.Color{
			R: r,
			G: g,
			B: b,
		}

		lightPalette.ContrastLight = getColorContrastFor(lightPalette.HCL, lightContrast)
		lightPalette.ContrastDark = getColorContrastFor(lightPalette.HCL, darkContrast)

		lightPalettes = append(lightPalettes, lightPalette)

		var darkPalette Palette
		darkPalette.Step = index + 1
		darkPalette.Steps = config.Steps
		darkPalette.Hue = hueDarkValues[index]
		darkPalette.Luminosity = luminousDarkValues[index]
		darkPalette.Saturation = saturatedDarkValues[index]
		darkPalette.HueRange = []int{config.HueEnd, config.HueStart}

		if darkPalette.Saturation > 1 {
			darkPalette.Saturation = 1
		}

		darkPalette.HCL = gocolorful.Hcl(darkPalette.Hue, darkPalette.Saturation, darkPalette.Luminosity)
		darkPalette.Hex = darkPalette.HCL.Hex()

		var dh, ds, dv = darkPalette.HCL.Hsv()
		darkPalette.HSV = gocolorful.Hsv(dh, ds, dv)

		var dh1, ds1, dl = darkPalette.HCL.Hsl()
		darkPalette.HSL = gocolorful.Hsl(dh1, ds1, dl)

		var dr, dg, db = darkPalette.HCL.LinearRgb()
		darkPalette.RGB = gocolorful.Color{
			R: dr,
			G: dg,
			B: db,
		}

		darkPalette.ContrastLight = getColorContrastFor(darkPalette.HCL, lightContrast)
		darkPalette.ContrastDark = getColorContrastFor(darkPalette.HCL, darkContrast)

		darkPalettes = append(darkPalettes, darkPalette)
	}

	return
}

func isDefaultValue(c gocolorful.Color) bool {
	return c.R == 0 && c.G == 0 && c.B == 0
}

// getColorContrastFor returns the contrast of giving color against white.
// WCAG contrast ratio
// see http://www.w3.org/TR/2008/REC-WCAG20-20081211/#contrast-ratiodef
func getColorContrastFor(c gocolorful.Color, contender gocolorful.Color) float64 {
	var _, _, l1 = c.Hcl()
	var _, _, l2 = contender.Hcl()
	if l1 > l2 {
		return (l1 + 0.05) / (l2 + 0.05)
	}
	return (l2 + 0.05) / (l1 + 0.05)
}

func getStepCurvePoints(curve Curve, steps int) []float64 {
	var values = make([]float64, steps)
	for index := 1; index < steps+1; index++ {
		var curvePoint = float64(index) / float64(steps-1)
		values[steps-index] = curve.Point(curvePoint)
	}
	return values
}

func distributeValue(value float64, fromRange []float64, toRange []float64) float64 {
	var fromLow, fromHigh = fromRange[0], fromRange[1]
	var toLow, toHigh = toRange[0], toRange[1]

	var valueFromLow = value - fromLow
	var diffBetweenFroms = fromHigh - fromLow
	var diffBetweenTos = toHigh - toLow
	var result = toLow + ((valueFromLow / diffBetweenFroms) * diffBetweenTos)

	if toLow < toHigh {
		if result < toLow {
			return toLow
		}
		if result > toHigh {
			return toHigh
		}

		return result
	}

	if result > toLow {
		return toLow
	}

	if result < toHigh {
		return toHigh
	}

	return result
}

func reversePaletteList(set []Palette) []Palette {
	var setCount = len(set)
	var reversed = make([]Palette, setCount)
	for index := setCount - 1; index > -1; index-- {
		reversed[index] = set[index]
	}
	return reversed
}

func reverseFloatList(set []float64) []float64 {
	var setCount = len(set)
	var reversed = make([]float64, setCount)
	for index := setCount - 1; index > -1; index-- {
		reversed[index] = set[index]
	}
	return reversed
}
