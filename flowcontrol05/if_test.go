package iftest

import (
	"testing"
	"fmt"
	"math"
)

func TestSqrt(t *testing.T) {
	var i,j string
	var a,b,c string
	
	i = Sqrt(2)
	a = fmt.Sprint(math.Sqrt(float64(2)))
	
	j = Sqrt(-4)
	c = fmt.Sprint( math.Sqrt( float64(-1)*float64(-4) ) )
	b = c + "i"

	if i == a {
		t.Log("success.")
	}else{
		t.Errorf("error.")
	}
	if j == b {
		t.Log("success.")
	}else{
		t.Errorf("error.")
	}
}
