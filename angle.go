package math3d

import "math"

type Angle float64

// Degrees creates an angle from an amount of degrees
func Degrees(degs float64) Angle {
	return Angle(degs / 180.0 * math.Pi)
}

// Radians creates an angle from an amount of radians
func Radians(rads float64) Angle {
	return Angle(rads)
}

// Degrees converts this angle to an amount of degrees
func (a Angle) Degrees() float64 {
	return float64(a) * 180.0 / math.Pi
}

// Radians converts this angle to an amount of radians
func (a Angle) Radians() float64 {
	return float64(a)
}
