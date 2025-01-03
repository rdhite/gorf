package propagation

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

var c float64 = 299_792_458

type Signal struct {
	Pattern   Pattern
	Watts     float64
	Frequency float64
	Direction mgl64.Vec3
	Location  mgl64.Vec3
}

// Provides the decibel difference of `p1` compared to `p2` - i.e.
// if p1 > p2, then the result is positive.
func RelativeDecibels(p1, p2 float64) float64 {
	return 10 * math.Log10(p1/p2)
}

// Calculates power density (mW/m^2)
func PowerAtPosition(sig Signal, location mgl64.Vec3) float64 {
	quat := rotateToX(sig.Direction)
	diff := location.Sub(sig.Location)
	dBi := sig.Pattern.CalcGainVec(quat.Rotate(diff))

	distDb := Linear2Decibel(diff.Len())
	fourPiDb := Linear2Decibel(4 * math.Pi)
	wavelengthDb := Linear2Decibel(freqToWavelength(sig.Frequency))
	transmitDb := Linear2Decibel(sig.Watts * 1000) // convert to milliwatts for dBm

	// Using equation for P_r from https://en.wikipedia.org/wiki/Free-space_path_loss
	// Assuming D_r is just 0 gain
	dBm := transmitDb + dBi /* + 0 + */ + 2*(wavelengthDb-fourPiDb-distDb)

	return Decibel2Linear(dBm)
}

// Returns the rotation matrix that transforms `vec` to {1, 0, 0}
func rotateToX(vec mgl64.Vec3) mgl64.Quat {
	return mgl64.QuatBetweenVectors(vec, mgl64.Vec3{1, 0, 0})
}

func freqToWavelength(freq float64) float64 {
	return c / freq
}
