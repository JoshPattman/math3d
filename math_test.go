package math3d

import (
	"fmt"
	"testing"
)

func TestCross(t *testing.T) {
	v := V(1, 2, 3)
	q := QAxisAngle(V(0.5, 0, 2), Degrees(45))
	fmt.Println(q.Apply(v))
}
