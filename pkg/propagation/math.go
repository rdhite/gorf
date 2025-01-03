package propagation

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

func AngleBetween(v1, v2 mgl64.Vec3) float64 {
	// TODO: this technically does what it the function indicates, but is
	// typically not what is desired (such as for bearing calculations where 30 != -30 degrees)
	return math.Acos(v1.Dot(v2) / (v1.Len() * v2.Len()))
}

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
