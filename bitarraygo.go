package bitarraygo

type BitArray []uint64

func NewBitArray(n uint) BitArray {
	return make(BitArray, n)
}

func (a BitArray) Set(i uint64) {
	a[0] |= 1 << i
}

func (a BitArray) Flip(i uint64) {
	a[0] ^= 1 << i
}

func (a BitArray) Get(i uint64) bool {
	b := a[0] & (1 << i)

	if b != 0 {
		return true
	}

	return false
}
