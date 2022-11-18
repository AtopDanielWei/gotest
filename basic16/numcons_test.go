package numcons

import (
	"testing"
)

func TestNeedInt(t *testing.T) {
	var i,j int

	i = NeedInt(Small)
	j = (Small*10)+ 1
	if i == j {
		t.Log("success.")
	}else{
		t.Errorf("error.")
	}

}
func TestNeedFloat(t *testing.T) {
	var i,j float64
	i = NeedFloat(Big)
	j = Big * 0.1

	if i == j {
		t.Log("success.")
	}else{
		t.Errorf("error.")
	}

	i = NeedFloat(Small)
	j = Small * 0.1

	if i == j {
		t.Log("success.")
	}else{
		t.Errorf("error.")
	}
}

