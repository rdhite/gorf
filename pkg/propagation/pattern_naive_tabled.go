package propagation

import (
	"cmp"
	"math"
	"slices"

	"github.com/go-gl/mathgl/mgl64"
)

// Built from one azimuth and one elevation gain pattern. It is assumed that the
// elevation pattern holds for all azimuths, not just in the direction of the
// main lobe. Elevation gain is normalized and then added to azimuth.
type NaiveTabledPattern struct {
	azimuths, elevations []AngleGain
}

// Create a naive gain pattern based on an azimuth plot and elevation plot. Plot's are
// represented by a list of angle/gain pairs - e.g. (0, 52) implies 52 dBi gain at 0 degrees.
func CreateNaiveTabledPattern(azimuths, elevations []AngleGain) (pattern NaiveTabledPattern) {
	pattern.azimuths = make([]AngleGain, len(azimuths))
	copy(azimuths, pattern.azimuths)
	slices.SortFunc(pattern.azimuths, func(a, b AngleGain) int {
		return int(a.Angle - b.Angle)
	})

	pattern.elevations = make([]AngleGain, len(elevations))
	copy(elevations, pattern.elevations)

	// elevation get's normalized so these values can just be added to
	// azimuth gains to get the compound gain.
	maxGain := -math.MaxFloat64
	for _, v := range elevations {
		maxGain = math.Max(maxGain, v.Gain)
	}
	for i, v := range pattern.elevations {
		pattern.elevations[i].Gain = v.Gain - maxGain
	}

	slices.SortFunc(pattern.elevations, func(a, b AngleGain) int {
		return cmp.Compare(a.Angle, b.Angle)
	})

	return
}

func (pattern NaiveTabledPattern) CalcGainAE(azimuth, elevation float64) float64 {
	var azGain, elGain float64

	azGain = interpolateGain(pattern.azimuths, azimuth)
	elGain = interpolateGain(pattern.elevations, elevation)

	return azGain + elGain
}

func (pattern NaiveTabledPattern) CalcGainVec(direction mgl64.Vec3) float64 {
	// azimuth will be the angle from x-axis to the x,y subcomponent of `direction`
	azimuth := math.Atan2(direction[0], direction[1])

	// elevation is angle from x,y plane to `direction`
	// arctan the "base" of the direction triangle (the x,y subcomponent) and the "height" (z subcomponent)
	elevation := math.Atan2(mgl64.Vec2{direction[0], direction[1]}.Len(), direction[2])

	return pattern.CalcGainAE(azimuth, elevation)
}

// Finds the gain for `angle` by interpolating within `gains`.
//
// Assumes that gains angles are sorted, non repeating, and in the range [0, 2*pi)
func interpolateGain(gains []AngleGain, angle float64) float64 {
	if len(gains) == 1 {
		return gains[0].Gain
	}

	// TODO: de-loop via modulo operator
	for angle < 0 {
		angle += 2 * math.Pi
	}

	// TODO: de-loop via modulo operator
	for angle >= 2*math.Pi {
		angle -= 2 * math.Pi
	}

	for i := 0; i < len(gains)-1; i++ {
		currV, nextV := gains[i], gains[i+1]
		if currV.Angle <= angle && angle < nextV.Angle {
			return interpolate(currV.Angle, currV.Gain, nextV.Angle, nextV.Gain, angle)
		}
	}

	// need to check the "last link" that makes the `gains` chain a loop
	currV, nextV := gains[len(gains)-1], gains[0]
	return interpolate(currV.Angle, currV.Gain, nextV.Angle+2*math.Pi, nextV.Gain, angle)
}
