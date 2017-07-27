package bitarraygo

import (
	"testing"

	"fmt"
)

func Test_Set(t *testing.T) {
	tests := []struct {
		in     uint64
		expect string
	}{
		{
			3,
			"[00001000]",
		},
	}

	for _, j := range tests {
		a := NewBitArray(1)
		a.Set(j.in)
		actual := fmt.Sprintf("%08b", a)
		if actual != j.expect {
			t.Errorf("For %d, expected %s, got %s.", j.in, j.expect, actual)
		}
	}
}
