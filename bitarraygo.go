// Package bitarraygo is a bit array with methods to manipulate it.
package bitarraygo

import (
	"fmt"
)

// BitArray is the structure representing the... bit array.
type BitArray struct {
	words []uint64
}

// NewBitArray returns a BitArray of the specified size.
// Currently the size is fixed at 18446744073709551614. WIP.
func NewBitArray(n uint64) *BitArray {
	// Todo: implement arrays bigger than 18446744073709551614
	// Until then, pin 1 to n.
	n = 1
	return &BitArray{words: make([]uint64, n)}
}

// AsNumber returns the value of the bit array as a number.
func (a *BitArray) AsNumber() uint64 {
	return a.words[0]
}

func (a *BitArray) AsString() string {
	return fmt.Sprintf("%b", a.words[0])
}

// FromNumber sets the bits of the bit array according to a specified number.
func (a *BitArray) FromNumber(n uint64) {
	a.words[0] = n
}

// FromBinary set the bits of the BitArray from a string of 1s and 0s.
func (a *BitArray) FromBinary(s string) {
	var b uint64 = 0
	for i, j := range s {
		if i != 0 {
			b = b << 1
		}

		if j == 49 {
			b |= 1
		}
	}
	a.words[0] = b
}

// Set turns on or off, as specified, the bit at the specified index.
func (a *BitArray) Set(i uint64, b bool) *BitArray {
	if b == true {
		a.words[0] |= 1 << i
	} else {
		a.words[0] &^= 1 << i
	}
	return a
}

// Flip toggles the bit at the specified index.
func (a *BitArray) Flip(i uint64) {
	a.words[0] ^= 1 << i
}

// Get returns the state of the bit at the specified index.
func (a *BitArray) Get(i uint64) bool {
	b := a.words[0] & (1 << i)

	if b != 0 {
		return true
	}

	return false
}
