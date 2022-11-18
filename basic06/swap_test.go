package swap

import (
	"testing"
)

func TestSwap(t *testing.T) {
	a := "daniel"
	b := "wei"
	got1,got2 := Swap(a,b)
	if got1 == a {
		t.Errorf("a = got1 ; want a = got2")
	}
	if got2 == b {
		t.Errorf("b = got2 ; want b = got1")
	}
}
