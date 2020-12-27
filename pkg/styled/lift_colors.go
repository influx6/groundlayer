package styled

import (
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
	Steps              int               // number of colors to be generated
	HueStart           int               // Hue value (0-359)
	HueEnd             int               // Hue value (0-359)
	HueCurve           Curve             // Hue bezier curve points
	SaturationRate     float64           // rate of saturation increase (0-200)
	SaturationStart    int               // Saturation value (0-100)
	SaturationEnd      int               // Saturation value (0-100)
	SaturationCurve    Curve             // Bezier curve points for saturation
	LuminousityStart   int               // Luminous value (0-100)
	LuminousityEnd     int               // Luminous value (0-100)
	LuminosityCurve    Curve             // Luminosity bezier curve points
	ContrastColorLight *gocolorful.Color // Provides light color contrast to be used for calculating contrast
	ContrastColorDark  *gocolorful.Color // Provides dark color contrast to be used for calculating contrast
}

func (h *HuePalette) ensure() {
	if h.Steps <= 0 {
		panic("HuePalette.Steps is required and should be above 0")
	}
	if h.LuminosityCurve == nil {
		panic("HuePalette.LuminosityCurve is required")
	}
	if h.HueCurve == nil {
		panic("HuePalette.HueCurve is required")
	}
	if h.SaturationCurve == nil {
		panic("HuePalette.SaturationCurve is required")
	}
	if h.SaturationRate <= 0 {
		panic("HuePalette.SaturationRate is required and should be above 0")
	}
	if h.SaturationStart <= 0 {
		panic("HuePalette.SaturationStart is required and should be above 0")
	}
	if h.SaturationEnd > 100 {
		panic("HuePalette.SaturationEnd is required and should be below or equal to 100")
	}
	if h.LuminousityStart <= 0 {
		panic("HuePalette.LuminousityStart is required and should be above 0")
	}
	if h.LuminousityEnd > 100 {
		panic("HuePalette.LuminousityEnd is required and should be below or equal to 100")
	}
	if h.HueStart <= 0 {
		panic("HuePalette.HueStart is required and should be above 0")
	}
	if h.HueEnd >= 360 {
		panic("HuePalette.HueEnd is required and should be below 360")
	}
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

type Palettes struct {
	LightMode []Palette
	DarkMode  []Palette // will be a list of LightMode palette set reversed.
}

func CreatePalette(config HuePalette) Palettes {
	config.ensure()

	var lightContrast = config.ContrastColorLight
	if lightContrast == nil {
		lightContrast = &white
	}

	var darkContrast = config.ContrastColorDark
	if darkContrast == nil {
		darkContrast = &black
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

	var lightPalettes = make([]Palette, 0, config.Steps)
	var darkPalettes = make([]Palette, 0, config.Steps)

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

		lightPalette.ContrastLight = getColorContrastFor(&lightPalette.HCL, lightContrast)
		lightPalette.ContrastDark = getColorContrastFor(&lightPalette.HCL, darkContrast)

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

		darkPalette.ContrastLight = getColorContrastFor(&darkPalette.HCL, lightContrast)
		darkPalette.ContrastDark = getColorContrastFor(&darkPalette.HCL, darkContrast)

		darkPalettes = append(darkPalettes, darkPalette)
	}

	return Palettes{
		LightMode: lightPalettes,
		DarkMode:  darkPalettes,
	}
}

// getColorContrastFor returns the contrast of giving color against white.
// WCAG contrast ratio
// see http://www.w3.org/TR/2008/REC-WCAG20-20081211/#contrast-ratiodef
func getColorContrastFor(c *gocolorful.Color, contender *gocolorful.Color) float64 {
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
	if len(fromRange) != 2 {
		panic("from range requires only 2 items")
	}
	if len(toRange) != 2 {
		panic("to range requires only 2 items")
	}

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
