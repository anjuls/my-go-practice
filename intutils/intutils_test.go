package intutils

import "testing"

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2,-2) = %d; want -2", ans)
	}
}
