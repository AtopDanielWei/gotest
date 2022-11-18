package typeconversions

import (
	"math"
)

func Typeconv(a,b int) (x int,y int,z uint) {
	x, y = a, b
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z = uint(f)
	return
}
