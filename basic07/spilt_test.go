package split

import (
	"testing"
)

func TestSplit(t *testing.T) {
	a := 17
	var tmp_got1,tmp_got2 int
	got1,got2 := Split(a)
	
	if tmp_got1 = a * 4 / 9; tmp_got1 != got1 {
		t.Errorf("error.")
	}else{
		t.Log("success.")
	}
	if tmp_got2 = a- tmp_got1 ; tmp_got2 != got2 {
		t.Errorf("error.")
	}else{
		t.Log("success.")
	}
}
