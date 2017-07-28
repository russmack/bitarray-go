package bitarraygo

import (
	"testing"

	"fmt"
)

func Test_Set_on(t *testing.T) {
	tests := []struct {
		input  uint64
		expect uint64
	}{
		{0, 1},
		{1, 2},
		{2, 4},
		{3, 8},
	}

	for _, j := range tests {
		a := NewBitArray(1)
		a.Set(j.input, true)
		actual := a.AsNumber()
		if actual != j.expect {
			t.Errorf("For %d, expected %d, got %d.", j.input, j.expect, actual)
		}
	}
}

func Test_Set_off(t *testing.T) {
	tests := []struct {
		input  uint64
		expect uint64
	}{
		{0, 170},
		{1, 168},
		{2, 170},
		{3, 162},
	}

	for _, j := range tests {
		a := NewBitArray(1)
		a.FromNumber(170)
		a.Set(j.input, false)
		actual := a.AsNumber()
		if actual != j.expect {
			t.Errorf("For %d, expected %d, got %d.", j.input, j.expect, actual)
		}
	}
}

func Test_Get(t *testing.T) {
	tests := []struct {
		input  uint64
		expect bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{3, false},
		{4, true},
		{5, true},
		{6, true},
	}

	a := NewBitArray(1)
	a.Set(0, true).
		Set(2, true).
		Set(4, true).
		Set(5, true).
		Set(6, true)

	for _, j := range tests {
		actual := a.Get(j.input)
		if actual != j.expect {
			t.Errorf("For %d, expected %t, got %t.", j.input, j.expect, actual)
		}
	}
}

func Test_Flip(t *testing.T) {
	type input struct {
		number uint64
		index  uint64
	}
	tests := []struct {
		input  input
		expect uint64
	}{
		{input{0, 0}, 1},
		{input{1, 0}, 0},
		{input{2, 0}, 3},
		{input{5, 1}, 7},
		{input{13, 2}, 9},
		{input{255, 7}, 127},
	}

	for _, j := range tests {
		a := NewBitArray(1)
		a.FromNumber(j.input.number)
		a.Flip(j.input.index)
		actual := a.AsNumber()
		if actual != j.expect {
			t.Errorf("For %d, %d, expected %d, got %d.",
				j.input.number, j.input.index, j.expect, actual)
		}
	}
}

func Test_FromNumber(t *testing.T) {
	tests := []struct {
		input  uint64
		expect uint64
	}{
		{0, 0},
		{1, 1},
		{37, 37},
		{255, 255},
	}

	for _, j := range tests {
		a := NewBitArray(1)
		a.FromNumber(j.input)
		actual := a.AsNumber()
		if actual != j.expect {
			t.Errorf("For %d, expected %d, got %d.", j.input, j.expect, actual)
		}
	}
}
