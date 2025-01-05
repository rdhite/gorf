package propagation

import "github.com/go-gl/mathgl/mgl64"

type AngleGain struct {
	Angle float64
	Gain  float64
}

// Represents the radiation pattern of a signal/antenna
//
// TODO: Swap the concept of main lobe and azimuth to be CW from "north" (y-axis)
// instead of CCW from x-axis, since that seems to be the common convention.
type Pattern interface {
	// Calculate (or retrieve) the gain (in dBi) of the radiation pattern at the
	// given `azimuth` and `elevation`.
	//
	// TODO: Figure out if azimuth/elevation should be 0-2pi or -pi-0pi or some other range.
	CalcGainAE(azimuth, elevation float64) float64

	// Calculate (or retrieve) the gain (in dBi) of the radiation pattern in the
	// given `direction`. NOTE: main lobe is assumed to face positive x-axis
	CalcGainVec(direction mgl64.Vec3) float64
}
