package main

import (
	"fmt"

	"github.com/russmack/bitarray-go"
)

func main() {
	fmt.Println("Bit array.")

	// Currently only supporting arrays of up to 18446744073709551614 bits
	// so this argument to NewBitArray is being ignored for now.
	numBits := uint64(8 * 4)
	b := bitarraygo.NewBitArray(numBits)

	i := 2
	n := uint64(i)

	b.Set(n, true)
	fmt.Printf("Set %d: %b\n", n, b.AsNumber())

	got := b.Get(n)
	fmt.Printf("Get %d: %t\n", n, got)

	n = 1
	fmt.Printf("Pre-flip: %b\n", b.AsNumber())
	b.Flip(n)
	fmt.Printf("Post-flip: %b\n", b.AsNumber())

}
