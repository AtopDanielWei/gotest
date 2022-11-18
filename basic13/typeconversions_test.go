package typeconversions

import (
	"testing"
	"math"
)

func TestTypeconv(t *testing.T) {
	x,y,z := Typeconv(3,4)
	if x != 3 {
		t.Errorf("error.")
	}else{
		t.Log("success.")
	}
	if y != 4 {
		t.Errorf("error.")
	}else{
		t.Log("success.")
	}
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var got uint = uint(f)
	if got != z {
		t.Errorf("error.")
	}else{
		t.Log("success.")
	}
}

