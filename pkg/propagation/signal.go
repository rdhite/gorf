package propagation

import "math"

type Signal struct {
}

// Provides the decibel difference of `p1` compared to `p2` - i.e.
// if p1 > p2, then the result is positive.
func RelativeDecibels(p1, p2 float64) float64 {
	return 10 * math.Log10(p1/p2)
}
