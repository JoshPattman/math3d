# math3d
This is a very simple 3d math packge written in pure go. It is not designed to be ultra-fast, but instead is supposed to be easy to use and read. It has no dependancies.
## Example
Vectors and Quaternions are not passed as pointers, but instead as values. This is deliberate to prevent accidental modification of vectors you don't mean to. Below is some example code that has no real purpose:
```go
v := V(0, 1, 0)
q := QAxisAngle(V(1, 0, 0), Degrees(45))

fmt.Println(q.Apply(v))

v2 := v.Add(V(0, 0, 1))
q2 := QFromTo(v2, v)

fmt.Println(q2.Apply(v2))
fmt.Println(v.Cross(v2))
```