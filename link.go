package rand

import (
	r "math/rand"
	_ "unsafe"
)

const (
	rngLen   = 607
	rngTap   = 273
	rngMax   = 1 << 63
	rngMask  = rngMax - 1
	int32max = (1 << 31) - 1
)

//go:linkname fastrand64 runtime.fastrand64
func fastrand64() uint64

//go:linkname fastrandn runtime.fastrandn
func fastrandn(n uint32) uint32

type internalRng struct{}

func (i *internalRng) Int63() int64 {
	return int64(fastrand64() & rngMask)
}

func (i *internalRng) Uint64() uint64 {
	return fastrand64()
}

func (i *internalRng) Seed(seed int64) {
	return
}
func newInternalRNG() r.Source {
	return &internalRng{}
}
