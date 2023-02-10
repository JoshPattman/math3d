package math3d

import (
	"fmt"
	"math"
)

// VecZero is the vector (0,0,0)
var VecZero = V(0, 0, 0)

// Vec represents a vector with 3 dimensions X Y Z
type Vec struct {
	X float64
	Y float64
	Z float64
}

// V creates a new Vec3 with x y z components
func V(x, y, z float64) Vec {
	return Vec{
		X: x,
		Y: y,
		Z: z,
	}
}

// Unit returns the normalized form of this vector. If the length is 0, it will return VecZero
func (v Vec) Unit() Vec {
	l := v.Len()
	if l == 0 {
		return VecZero
	}
	return v.Mul(1.0 / l)
}

// Len returns the length of this vector
func (v Vec) Len() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// Len2 returns the squared length of this vector
func (v Vec) Len2() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

// Mul returns the vector scaled by a scalar number
func (v Vec) Mul(f float64) Vec {
	vn := v
	vn.X *= f
	vn.Y *= f
	vn.Z *= f
	return vn
}

// MulVec returns the elementwise multiplication of this vector and v2
func (v Vec) MulVec(v2 Vec) Vec {
	vn := v
	vn.X *= v2.X
	vn.Y *= v2.Y
	vn.Z *= v2.Z
	return vn
}

// Add returns the elementwise addition of this vector and v2
func (v Vec) Add(v2 Vec) Vec {
	vn := v
	vn.X += v2.X
	vn.Y += v2.Y
	vn.Z += v2.Z
	return vn
}

// Sub returns the elementwise subtraction of this vector and v2
func (v Vec) Sub(v2 Vec) Vec {
	vn := v
	vn.X -= v2.X
	vn.Y -= v2.Y
	vn.Z -= v2.Z
	return vn
}

// Inv returns the inverse of this vector: v.Mul(-1)
func (v Vec) Inv() Vec {
	return v.Mul(-1)
}

// Dot returns the dot product of this vector and v2
func (v Vec) Dot(v2 Vec) float64 {
	return (v.X * v2.X) + (v.Y * v2.Y) + (v.Z * v2.Z)
}

// Cross returns the cross product of this vector and v2
func (v Vec) Cross(v2 Vec) Vec {
	return Vec{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

// Rotated returns the vector rotated by q
func (v Vec) Rotated(q Quat) Vec {
	return q.Apply(v)
}

// AngleTo returns the angle from this vector to v2
func (v Vec) AngleTo(v2 Vec) Angle {
	return Radians(math.Acos(v.Dot(v2) / (v.Len() * v2.Len())))
}

// DistanceTo returns distance between this vector and v2
func (v Vec) DistanceTo(v2 Vec) float64 {
	return v.Sub(v2).Len()
}

// Project returns this vector projected onto the plane with normal
func (v Vec) Project(normal Vec) Vec {
	d := v.Dot(normal) / normal.Len()
	p := normal.Unit().Mul(d)
	return v.Sub(p)
}

// Reflect returns the vector reflected on the plane with normal
func (v Vec) Reflect(normal Vec) Vec {
	normal = normal.Unit()
	return v.Sub(normal.Mul(2 * v.Dot(normal)))
}

// String returns this vector as a readable string
func (v Vec) String() string {
	return fmt.Sprintf("(x%.2f,y%.2f,z%.2f)", v.X, v.Y, v.Z)
}
