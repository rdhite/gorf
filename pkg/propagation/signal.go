package propagation

import "math"

var c float64 = 299_792_458

type Signal struct {
	Pattern   Pattern
	Watts     float64
	Frequency float64
	Direction Vec3
	Location  Vec3
}

// Provides the decibel difference of `p1` compared to `p2` - i.e.
// if p1 > p2, then the result is positive.
func RelativeDecibels(p1, p2 float64) float64 {
	return 10 * math.Log10(p1/p2)
}

func PowerAtPosition(sig Signal, location Vec3) (dBm float64) {
	mat := rotateToX(sig.Direction)
	diff := location.Sub(sig.Location)
	dist := Magnitude(diff)
	dBi := sig.Pattern.CalcGainVec(matmul(mat, diff))

	fspl := math.Pow(4*math.Pi*dist/freqToWavelength(sig.Frequency), 2)

	foo := sig.Watts * 1 /* turn dBi into linear */ / fspl

	return foo /*after making it logarithmic again*/
}

// Returns the rotation matrix that transforms `vec` to {1, 0, 0}
func rotateToX(vec Vec3) (mat [3]Vec3) {
	// TODO: implement
	return
}

func matmul(mat [3]Vec3, vec Vec3) Vec3 {
	// TODO: implement
	return Vec3{}
}

func freqToWavelength(freq float64) float64 {
	return c / freq
}
