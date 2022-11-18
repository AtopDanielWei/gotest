package foriswhile

import (
	"testing"
)

func TestForiswhile(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	i := Foriswhile()
    switch {
		case i == sum :
			t.Log("success.")
		default:
			t.Errorf("error.")
    }

}
