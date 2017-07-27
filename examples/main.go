package main

import (
	"fmt"

	"github.com/russmack/bitarray-go"
)

func main() {
	fmt.Println("Bit array.")

	numBits := 8 * 4
	_ = numBits
	b := bitarraygo.NewBitArray(1)

	i := 2
	n := uint64(i)
	b.Set(n)
	fmt.Printf("Set %d: %b\n", n, b)

	got := b.Get(0)
	fmt.Println("got:", got)

	got = b.Get(1)
	fmt.Println("got:", got)

	got = b.Get(2)
	fmt.Println("got:", got)

	got = b.Get(3)
	fmt.Println("got:", got)

	got = b.Get(4)
	fmt.Println("got:", got)

	fmt.Printf("pre : %b\n", b)
	b.Flip(1)
	fmt.Printf("post: %b\n", b)

}
