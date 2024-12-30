package propagation

type AngleGain struct {
	Angle float64
	Gain  float64
}

type Pattern interface {
	// Calculate (or retrieve) the gain (in dBi) of the radiation pattern at the
	// given `azimuth` and `elevation`.
	CalcGainAE(azimuth, elevation float64) float64

	// Calculate (or retrieve) the gain (in dBi) of the radiation pattern in the
	// given `direction`. NOTE: main lobe is assumed to face positive x-axis
	CalcGainVec(direction Vec3) float64
}
