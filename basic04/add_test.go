package add

import "testing"

func TestAdd(t *testing.T) {
	n := Add(1, 2)
	if n == 3 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
