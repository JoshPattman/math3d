# 'github.com/JoshPattman/math3d' - Simple 3D math library for golang
This is a very simple 3d math packge written in pure go with no dependancies. It prioritises simplicity and readability over speed (that's not to say the pacakge is slow, but its just not as fast as `gonum` for example).
## Pacakge Methodology
Vectors and Quaternions are passed by value, not pointer. This means that functions such as `Add` and `Cross` return a new value, which allows methods to be chained on a single line, such as `z := a.Add(b).Mul(0.5).Sub(c)`
## Example
Below is some example code:
```golang
// Create a new vector with y component of 1
v := V(0, 1, 0)
// Create a new rotation of 45 degrees around the x axis
q := QAxisAngle(V(1, 0, 0), Degrees(45))
// Apply the rotation to the vector
fmt.Println(q.Apply(v))

// Create a new vector that is the addition of our first vector and a vector with z component 1
v2 := v.Add(V(0, 0, 1))
// Create a rotation from our second vector to our first
q2 := QFromTo(v2, v)
// Apply that rotation to our second vector
fmt.Println(q2.Apply(v2))

// Cross vector 1 with vector 2
fmt.Println(v.Cross(v2))
```