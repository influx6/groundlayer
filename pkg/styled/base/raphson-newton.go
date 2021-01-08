package base

import "math"

const (
	// These values are established by empiricism with
	// tests (tradeoff: performance VS precision)
	NEWTON_ITERATIONS          = 4
	NEWTON_MIN_SLOPE           = 0.001
	SUBDIVISION_PRECISION      = 0.0000001
	SUBDIVISION_MAX_ITERATIONS = float64(10)

	kSplineTableSize = 11
	kSampleStepSize  = 1.0 / (kSplineTableSize - 1.0)
)

func aPoint(aA1 float64, aA2 float64) float64 { return 1.0 - 3.0*aA2 + 3.0*aA1 }
func bPoint(aA1 float64, aA2 float64) float64 { return 3.0*aA2 - 6.0*aA1 }
func cPoint(aA1 float64) float64              { return 3.0 * aA1 }

// Returns x(t) given t, x1, and x2, or y(t) given t, y1, and y2.
func calcBezier(aT float64, aA1 float64, aA2 float64) float64 {
	return ((aPoint(aA1, aA2)*aT+bPoint(aA1, aA2))*aT + cPoint(aA1)) * aT
}

// Returns dx/dt given t, x1, and x2, or dy/dt given t, y1, and y2.
func getSlope(aT float64, aA1 float64, aA2 float64) float64 {
	return 3.0*aPoint(aA1, aA2)*aT*aT + 2.0*bPoint(aA1, aA2)*aT + cPoint(aA1)
}

func binarySubdivide(aX, aA, aB, mX1, mX2 float64) float64 {
	var currentX float64 = 0
	var currentT float64 = 0
	var i float64 = 0
	for true {
		i = +1
		var condition = (math.Abs(currentX) > SUBDIVISION_PRECISION) && (i < SUBDIVISION_MAX_ITERATIONS)
		if !condition {
			break
		}

		currentT = aA + (aB-aA)/2.0
		currentX = calcBezier(currentT, mX1, mX2) - aX
		if currentX > 0.0 {
			aB = currentT
		} else {
			aA = currentT
		}
	}
	return currentT
}

func newtonRaphsonIterate(aX, aGuessT, mX1, mX2 float64) float64 {
	for i := 0; i < NEWTON_ITERATIONS; i++ {
		var currentSlope = getSlope(aGuessT, mX1, mX2)
		if currentSlope == 0.0 {
			return aGuessT
		}
		var currentX = calcBezier(aGuessT, mX1, mX2) - aX
		aGuessT -= currentX / currentSlope
	}
	return aGuessT
}

func linearEasing(x float64) float64 {
	return x
}

type CurveFunc func(t float64) float64

func (c CurveFunc) Point(t float64) float64 {
	return c(t)
}

func BezierTPoints(x1, y1, x2, y2 float64) Curve {
	if !(0 <= x1 && x1 <= 1 && 0 <= x2 && x2 <= 1) {
		panic("bezier x values must be in [0, 1] range")
	}

	if x1 == y1 && x2 == y2 {
		return CurveFunc(linearEasing)
	}

	var sampleValues = make([]float64, kSplineTableSize)
	for i := 0; i < kSplineTableSize; i++ {
		sampleValues[i] = calcBezier(float64(i)*float64(kSampleStepSize), x1, x2)
	}

	var getTForX = func(aX float64) float64 {
		var intervalStart = 0.0
		var lastSample = kSplineTableSize - 1

		var currentSample = 1
		for ; currentSample != lastSample && sampleValues[currentSample] <= aX; currentSample++ {
			intervalStart += float64(kSampleStepSize)
		}

		currentSample--

		// Interpolate to provide an initial guess for t
		var dist = (aX - sampleValues[currentSample]) / (sampleValues[currentSample+1] - sampleValues[currentSample])
		var guessForT = intervalStart + dist*float64(kSampleStepSize)

		var initialSlope = getSlope(guessForT, x1, x2)
		if initialSlope >= NEWTON_MIN_SLOPE {
			return newtonRaphsonIterate(aX, guessForT, x1, x2)
		} else if initialSlope == 0.0 {
			return guessForT
		} else {
			return binarySubdivide(aX, intervalStart, intervalStart+float64(kSampleStepSize), x1, x2)
		}
	}

	return CurveFunc(func(x float64) float64 {
		if x == 0 || x == 1 {
			return x
		}
		return calcBezier(getTForX(x), y1, y2)
	})
}
