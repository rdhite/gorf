package propagation

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

type SincPattern struct {
	// Scaling factor added to sinc(x) to go from max-relative (dB) to
	// isotropic-relative (dBi).
	Factor float64
}

func (pattern SincPattern) CalcGainAE(azimuth, elevation float64) float64 {
	// sinc is axially symmetrical, so we just have to get the angle between the
	// antenna's forward and the target azimuth/elevation compound direction.
	// The math for that happens to simplify down nicely.
	theta := calc_compound(azimuth, elevation)
	if theta == 0 {
		theta = math.SmallestNonzeroFloat64
	}

	// Honestly not sure if this is correct or how it's supposed to work.
	return pattern.Factor + 10*math.Log10(math.Abs(sinc(theta)))
}

func (pattern SincPattern) CalcGainVec(direction mgl64.Vec3) float64 {
	theta := angleBetween(mgl64.Vec3{1, 0, 0}, direction)
	if theta == 0 {
		theta = math.SmallestNonzeroFloat64
	}
	return pattern.Factor + 10*math.Log10(math.Abs(sinc(theta)))
}

func sinc(theta float64) float64 {
	return math.Sin(math.Pi*theta) / (math.Pi * theta)
}

// Calculate the angle between "forward" and the resulting look
// angle that `azimuth` and `elevation` produce.
func calc_compound(azimuth, elevation float64) float64 {
	return math.Acos(math.Cos(azimuth) * math.Cos(elevation))
}

func angleBetween(v1, v2 mgl64.Vec3) float64 {
	return math.Acos(v1.Dot(v2) / (v1.Len() * v2.Len()))
}
