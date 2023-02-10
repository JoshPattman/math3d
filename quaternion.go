package math3d

// Based on github.com/westphae/quaternion

import (
	"fmt"
	"math"
)

// QIdentity is the identity quaternion (zero rotation)
var QIdentity = Q(1, 0, 0, 0)

// Q creates a new quation wxyz
func Q(w, x, y, z float64) Quat {
	return Quat{W: w, X: x, Y: y, Z: z}
}

// QAxisAngle returns a quaternion with rotation a around axis v
func QAxisAngle(v Vec, a Angle) Quat {
	v = v.Unit()
	// Here we calculate the sin( theta / 2) once for optimization
	factor := math.Sin(a.Radians() / 2.0)

	// Calculate the x, y and z of the quaternion
	x := v.X * factor
	y := v.Y * factor
	z := v.Z * factor

	// Calculate the w value by cos( theta / 2 )
	w := math.Cos(a.Radians() / 2.0)

	return Q(w, x, y, z).Unit()
}

// QFromTo returns quaternion that rotates from v1 to v2
func QFromTo(v1, v2 Vec) Quat {
	axis := v1.Cross(v2)
	angle := v1.AngleTo(v2)
	return QAxisAngle(axis, angle)
}

// Quat represents a quaternion W X Y Z
type Quat struct {
	W float64 // Scalar component
	X float64 // i component
	Y float64 // j component
	Z float64 // k component
}

// Conj returns the conjugate of a Quat (W,X,Y,Z) -> (W,-X,-Y,-Z)
func (qin Quat) Conj() Quat {
	qin.X = -qin.X
	qin.Y = -qin.Y
	qin.Z = -qin.Z
	return qin
}

// Norm returns the L1-Norm of a Quat (W,X,Y,Z) -> Sqrt(W*W+X*X+Y*Y+Z*Z)
func (qin Quat) Norm() float64 {
	return math.Sqrt(qin.Norm2())
}

// Norm2 returns the L2-Norm of a Quat (W,X,Y,Z) -> W*W+X*X+Y*Y+Z*Z
func (qin Quat) Norm2() float64 {
	return qin.W*qin.W + qin.X*qin.X + qin.Y*qin.Y + qin.Z*qin.Z
}

// Neg returns the negative
func (qin Quat) Neg() Quat {
	qin.W = -qin.W
	qin.X = -qin.X
	qin.Y = -qin.Y
	qin.Z = -qin.Z
	return qin
}

// Prod returns the product of q and q2 (q*q2)
func (q Quat) Prod(q2 Quat) Quat {
	return Quat{
		q.W*q2.W - q.X*q2.X - q.Y*q2.Y - q.Z*q2.Z,
		q.W*q2.X + q.X*q2.W + q.Y*q2.Z - q.Z*q2.Y,
		q.W*q2.Y + q.Y*q2.W + q.Z*q2.X - q.X*q2.Z,
		q.W*q2.Z + q.Z*q2.W + q.X*q2.Y - q.Y*q2.X,
	}
}

// Unit returns the Quat rescaled to unit-L1-norm
func (qin Quat) Unit() Quat {
	k := qin.Norm()
	return Quat{qin.W / k, qin.X / k, qin.Y / k, qin.Z / k}
}

// Inv returns the Quat conjugate rescaled so that Q Q* = 1. This is the reversed rotation
func (qin Quat) Inv() Quat {
	k2 := qin.Norm2()
	q := qin.Conj()
	return Quat{q.W / k2, q.X / k2, q.Y / k2, q.Z / k2}
}

// Apply returns the vector rotated by the quaternion.
func (qin Quat) Apply(vec Vec) Vec {
	conj := qin.Conj()
	aug := Quat{0, vec.X, vec.Y, vec.Z}
	rot := qin.Prod(aug).Prod(conj)
	return Vec{rot.X, rot.Y, rot.Z}
}

// RotateByGlobal rotates quaternion a by b along the local axis relative to a. It is sugar for a*b
func (a Quat) RotateByLocal(b Quat) Quat {
	return a.Prod(b)
}

// RotateByGlobal rotates quaternion a by b along the global axis. It is sugar for b*a
func (a Quat) RotateByGlobal(b Quat) Quat {
	return b.Prod(a)
}

// String converts this quaternion to a string
func (q Quat) String() string {
	return fmt.Sprintf("(w%.2f,x%.2f,y%.2f,z%.2f)", q.W, q.X, q.Y, q.Z)
}
