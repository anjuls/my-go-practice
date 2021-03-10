package mascot_test

import (
	"testing"

	"example.com/mascot/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Errorf("Wrong Mascot")
	}
}
