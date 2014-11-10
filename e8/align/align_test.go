package align

import (
	"testing"
)

func TestAlign(t *testing.T) {
	o := func(b bool) {
		if !b {
			t.Fail()
		}
	}

	o(A32(3) == 0)
	o(A16(3) == 2)
	o(A32(1024) == 1024)
	o(A32(1025) == 1024)
	o(A32(1026) == 1024)
	o(A32(1027) == 1024)
}
