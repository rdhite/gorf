package propagation

import "math"

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Add(vec Vec3) Vec3 {
	return Vec3{v.X + vec.X, v.Y + vec.Y, v.Z + vec.Z}
}

func (v Vec3) Sub(vec Vec3) Vec3 {
	return Vec3{v.X - vec.X, v.Y - vec.Y, v.Z - vec.Z}
}

func Normalize(v Vec3) Vec3 {
	mag := Magnitude(v)
	return Vec3{X: v.X / mag, Y: v.Y / mag, Z: v.Z / mag}
}

func Magnitude(v Vec3) float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func Dot(v1, v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func AngleBetween(v1, v2 Vec3) float64 {
	return math.Acos(Dot(v1, v2) / (Magnitude(v1) * Magnitude(v2)))
}

// Calculate the angle between "forward" and the resulting look
// angle that `azimuth` and `elevation` produce.
func calc_compound(azimuth, elevation float64) float64 {
	return math.Acos(math.Cos(azimuth) * math.Cos(elevation))
}

func interpolate(x1, y1, x2, y2, xVal float64) float64 {
	slope := (y2 - y1) / (x2 - x1)
	return y1 + slope*(xVal-x1)
}
