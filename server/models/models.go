package models

import (
	"math"
)

// Vector represents a point in 2D space
type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Position x y position
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Velocity dx dy velocity
type Velocity struct {
	Dx float64 `json:"dx"`
	Dy float64 `json:"dy"`
}

func (v *Velocity) magnitude() float64 {
	return math.Hypot(v.Dx, v.Dy)
}

func (v *Velocity) normalize() {
	mag := v.magnitude()
	if mag > 0 {
		v.Dx /= mag
		v.Dy /= mag
	}
}
