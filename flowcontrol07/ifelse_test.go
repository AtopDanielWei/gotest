package ifelse

import (
	"testing"
	"math"
)

func TestPow(t *testing.T) {
	var a,b,x,y float64
	
	a = Pow(3,2,10)
	b = Pow(3,3,20)

	if x = math.Pow(3, 2); x >= 10 {
		x = 10
	}
	if y = math.Pow(3, 3); y >= 20 {
		y = 20
	}

	if a == x {
		t.Log("success.")
	}else{
		t.Errorf("error.")
	}
	if b == y {
		t.Log("success.")
	}else{
		t.Errorf("error.")
	}
}
