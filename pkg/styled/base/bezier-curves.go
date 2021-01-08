package base

import (
	"gonum.org/v1/plot/tools/bezier"
	"gonum.org/v1/plot/vg"
)

func isEven(v int) bool {
	return (v & 1) == 0
}

type CurveXYPoints []vg.Point

func BezierXYPointsList(points ...float64) CurveXYPoints {
	if !isEven(len(points)) {
		panic("Only works with even length points")
	}

	var pointCount = len(points)

	var pointList = make([]vg.Point, 0, pointCount/2)
	for i := 0; i < pointCount; i += 2 {
		pointList = append(pointList, vg.Point{
			X: vg.Length(points[i]),
			Y: vg.Length(points[i+1]),
		})
	}
	return pointList
}

func BezierXYPoints(points ...float64) XYCurve {
	return bezier.New(BezierXYPointsList(points...)...)
}

//------------------------------------------------------------/
// XYCurve for Bezier producing single float value on curve
//------------------------------------------------------------/

var Linear = BezierXYPoints(0.5, 0.5, 0.5, 0.5)

var EaseInCubic = BezierXYPoints(0.55, 0.055, 0.675, 0.19)
var EaseOutCubic = BezierXYPoints(0.215, 0.61, 0.355, 1)
var EaseInOutCubic = BezierXYPoints(0.645, 0.045, 0.355, 1)
var EaseInSine = BezierXYPoints(0.47, 0, 0.745, 0.715)
var EaseOutSine = BezierXYPoints(0.39, 0.575, 0.565, 1)
var EaseInOutSine = BezierXYPoints(0.445, 0.05, 0.55, 0.95)

var EaseInQuad = BezierXYPoints(0.55, 0.085, 0.68, 0.53)
var EaseOutQuad = BezierXYPoints(0.25, 0.46, 0.45, 0.94)
var EaseInOutQuad = BezierXYPoints(0.455, 0.03, 0.515, 0.955)

var EaseInQuart = BezierXYPoints(0.895, 0.03, 0.685, 0.22)
var EaseOutQuart = BezierXYPoints(0.165, 0.84, 0.44, 1)
var EaseInOutQuart = BezierXYPoints(0.77, 0, 0.175, 1)

var EaseInCirc = BezierXYPoints(0.6, 0.04, 0.98, 0.335)
var EaseOutCirc = BezierXYPoints(0.075, 0.82, 0.165, 1)
var EaseInOutCirc = BezierXYPoints(0.785, 0.135, 0.15, 0.86)

var EaseInQuint = BezierXYPoints(0.755, 0.05, 0.855, 0.06)
var EaseOutQuint = BezierXYPoints(0.23, 1, 0.32, 1)
var EaseInOutQuint = BezierXYPoints(0.86, 0, 0.07, 1)

var EaseInExpo = BezierXYPoints(0.95, 0.05, 0.795, 0.035)
var EaseOutExpo = BezierXYPoints(0.19, 1, 0.22, 1)
var EaseInOutExpo = BezierXYPoints(1, 0, 0, 1)

var EaseInBack = BezierXYPoints(0.6, -0.28, 0.735, 0.045)
var EaseOutBack = BezierXYPoints(0.175, 0.885, 0.32, 1.275)
var EaseInOutBack = BezierXYPoints(0.68, -0.55, 0.265, 1.55)

//------------------------------------------------------------/
// Curve for Bezier producing single float value on curve
//------------------------------------------------------------/

var LinearTCurve = BezierTPoints(0.5, 0.5, 0.5, 0.5)

var EaseInCubicTCurve = BezierTPoints(0.55, 0.055, 0.675, 0.19)
var EaseOutCubicTCurve = BezierTPoints(0.215, 0.61, 0.355, 1)
var EaseInOutCubicTCurve = BezierTPoints(0.645, 0.045, 0.355, 1)
var EaseInSineTCurve = BezierTPoints(0.47, 0, 0.745, 0.715)
var EaseOutSineTCurve = BezierTPoints(0.39, 0.575, 0.565, 1)
var EaseInOutSineTCurve = BezierTPoints(0.445, 0.05, 0.55, 0.95)

var EaseInQuadTCurve = BezierTPoints(0.55, 0.085, 0.68, 0.53)
var EaseOutQuadTCurve = BezierTPoints(0.25, 0.46, 0.45, 0.94)
var EaseInOutQuadTCurve = BezierTPoints(0.455, 0.03, 0.515, 0.955)

var EaseInQuartTCurve = BezierTPoints(0.895, 0.03, 0.685, 0.22)
var EaseOutQuartTCurve = BezierTPoints(0.165, 0.84, 0.44, 1)
var EaseInOutQuartTCurve = BezierTPoints(0.77, 0, 0.175, 1)

var EaseInCircTCurve = BezierTPoints(0.6, 0.04, 0.98, 0.335)
var EaseOutCircTCurve = BezierTPoints(0.075, 0.82, 0.165, 1)
var EaseInOutCircTCurve = BezierTPoints(0.785, 0.135, 0.15, 0.86)

var EaseInQuintTCurve = BezierTPoints(0.755, 0.05, 0.855, 0.06)
var EaseOutQuintTCurve = BezierTPoints(0.23, 1, 0.32, 1)
var EaseInOutQuintTCurve = BezierTPoints(0.86, 0, 0.07, 1)

var EaseInExpoTCurve = BezierTPoints(0.95, 0.05, 0.795, 0.035)
var EaseOutExpoTCurve = BezierTPoints(0.19, 1, 0.22, 1)
var EaseInOutExpoTCurve = BezierTPoints(1, 0, 0, 1)

var EaseInBackTCurve = BezierTPoints(0.6, -0.28, 0.735, 0.045)
var EaseOutBackTCurve = BezierTPoints(0.175, 0.885, 0.32, 1.275)
var EaseInOutBackTCurve = BezierTPoints(0.68, -0.55, 0.265, 1.55)
