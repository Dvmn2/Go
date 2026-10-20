package mascot_test

import (
	"testing"

	"github.com/Dvmn2/Go.git/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.Mascot() != "Mascot" {
		t.Fatal("Wrong Mascot")
	}
}
