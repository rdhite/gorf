package propagation

type IsotropicPattern struct{}

func (pattern IsotropicPattern) CalcGainAE(azimuth, elevation float64) float64 {
	return 0.0
}

func (pattern IsotropicPattern) CalcGainVec(direction [3]float64) float64 {
	return 0.0
}
