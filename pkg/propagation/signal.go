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

// Calculates power density (mW/m^2)
func PowerAtPosition(sig Signal, location mgl64.Vec3) float64 {
	rotateToXQuat := mgl64.QuatBetweenVectors(sig.Direction, mgl64.Vec3{1, 0, 0})
	diff := location.Sub(sig.Location)
	dBi := sig.Pattern.CalcGainVec(rotateToXQuat.Rotate(diff))

	distDb := Linear2Decibel(diff.Len())
	fourPiDb := Linear2Decibel(4 * math.Pi)
	wavelengthDb := Linear2Decibel(freqToWavelength(sig.Frequency))
	transmitDb := Linear2Decibel(sig.Watts * 1000) // convert to milliwatts for dBm

	// Using equation for P_r from https://en.wikipedia.org/wiki/Free-space_path_loss
	// Assuming D_r is just 0 gain
	dBm := transmitDb + dBi /* + 0 + */ + 2*(wavelengthDb-fourPiDb-distDb)

	return Decibel2Linear(dBm)
}

func freqToWavelength(freq float64) float64 {
	return c / freq
}
