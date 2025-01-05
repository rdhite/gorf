package propagation

import (
	"math"
)

func Decibel2Linear(dB float64) float64 {
	return math.Pow(10, dB/10)
}

func Linear2Decibel(lin float64) float64 {
	return 10 * math.Log10(lin)
}

func interpolate(x1, y1, x2, y2, xVal float64) float64 {
	slope := (y2 - y1) / (x2 - x1)
	return y1 + slope*(xVal-x1)
}
