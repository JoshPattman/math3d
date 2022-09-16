package math3d

import (
	"fmt"
	"math"
)

// VecZero is the vector (0,0,0)
var VecZero = V(0, 0, 0)

// Vec3 represents a vector with 3 dimensions X Y Z
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// V creates a new Vec3 with x y z components
func V(x, y, z float64) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

// Unit returns the normalized form of this vector. If the length is 0, it will return VecZero
func (v Vec3) Unit() Vec3 {
	l := v.Len()
	if l == 0 {
		return VecZero
	}
	return v.Mul(1.0 / l)
}

// Len returns the length of this vector
func (v Vec3) Len() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// Len2 returns the squared length of this vector
func (v Vec3) Len2() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

// Mul returns the vector scaled by a scalar number
func (v Vec3) Mul(f float64) Vec3 {
	vn := v
	vn.X *= f
	vn.Y *= f
	vn.Z *= f
	return vn
}

// MulVec returns the elementwise multiplication of this vector and v2
func (v Vec3) MulVec(v2 Vec3) Vec3 {
	vn := v
	vn.X *= v2.X
	vn.Y *= v2.Y
	vn.Z *= v2.Z
	return vn
}

// Add returns the elementwise addition of this vector and v2
func (v Vec3) Add(v2 Vec3) Vec3 {
	vn := v
	vn.X += v2.X
	vn.Y += v2.Y
	vn.Z += v2.Z
	return vn
}

// Sub returns the elementwise subtraction of this vector and v2
func (v Vec3) Sub(v2 Vec3) Vec3 {
	vn := v
	vn.X -= v2.X
	vn.Y -= v2.Y
	vn.Z -= v2.Z
	return vn
}

// Inv returns the inverse of this vector: v.Mul(-1)
func (v Vec3) Inv() Vec3 {
	return v.Mul(-1)
}

// Dot returns the dot product of this vector and v2
func (v Vec3) Dot(v2 Vec3) float64 {
	return (v.X * v2.X) + (v.Y * v2.Y) + (v.Z * v2.Z)
}

// Cross returns the cross product of this vector and v2
func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

// Rotated returns the vector rotated by q
func (v Vec3) Rotated(q Quat) Vec3 {
	return q.Apply(v)
}

// AngleTo returns the angle from this vector to v2
func (v Vec3) AngleTo(v2 Vec3) Angle {
	return Radians(math.Acos(v.Dot(v2) / (v.Len() * v2.Len())))
}

// DistanceTo returns distance between this vector and v2
func (v Vec3) DistanceTo(v2 Vec3) float64 {
	return v.Sub(v2).Len()
}

// Project returns this vector projected onto the plane with normal
func (v Vec3) Project(normal Vec3) Vec3 {
	d := v.Dot(normal) / normal.Len()
	p := normal.Unit().Mul(d)
	return v.Sub(p)
}

// Reflect returns the vector reflected on the plane with normal
func (v Vec3) Reflect(normal Vec3) Vec3 {
	normal = normal.Unit()
	return v.Sub(normal.Mul(2 * v.Dot(normal)))
}

// String returns this vector as a readable string
func (v Vec3) String() string {
	return fmt.Sprintf("(x%.2f,y%.2f,z%.2f)", v.X, v.Y, v.Z)
}
