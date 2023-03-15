package rand

import (
	r "math/rand"
	"sync"
	"time"

	"github.com/zeebo/wyhash"
)

var (
	// pesudo RNG seed generator
	rng wyhash.SRNG

	// lock free rand pool
	defaultRandPool = sync.Pool{
		New: func() any {
			return r.New(r.NewSource(getTimeBasedSeed()))
		},
	}
)

func step(s ...int) int {
	defaultStep := 1
	if len(s) > 0 {
		if s[0] > 0 {
			defaultStep = s[0]
		}
	}
	return defaultStep
}

func getTimeBasedSeed() int64 {
	return time.Now().UnixNano() ^ int64(rng.Uint64())
}

func ExpFloat64() float64 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.ExpFloat64()
}
func Float32() float32 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Float32()
}
func Float64() float64 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Float64()
}
func Int() int {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Int()
}
func Int31() int32 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Int31()
}
func Int31n(n int32) int32 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Int31n(n)
}
func Int63() int64 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Int63()
}
func Int63n(n int64) int64 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Int63n(n)
}
func Intn(n int) int {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Intn(n)
}
func NormFloat64() float64 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.NormFloat64()
}
func Perm(n int) []int {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Perm(n)
}
func Read(p []byte) (n int, err error) {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Read(p)
}
func Seed(seed int64) {
	return
}
func Shuffle(n int, swap func(i, j int)) {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	rd.Shuffle(n, swap)
}
func Uint32() uint32 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Uint32()
}
func Uint64() uint64 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	return rd.Uint64()
}

func Intrange(from, to int, _step ...int) int {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	stp := step(_step...)
	width := to - from
	switch {
	case stp == 1:
		return from + Intn(width)
	case stp > 0:
		n := (width + stp - 1) / stp
		return from + stp*Intn(n)
	case stp < 0:
		n := (width + stp + 1) / stp
		return from + stp*Intn(n)
	default:
		panic("error step")
	}
}

func Int63range(from, to int64, _step ...int) int64 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	stp := int64(step(_step...))
	width := to - from
	switch {
	case stp == 1:
		return from + Int63n(width)
	case stp > 0:
		n := (width + stp - 1) / stp
		return from + stp*Int63n(n)
	case stp < 0:
		n := (width + stp + 1) / stp
		return from + stp*Int63n(n)
	default:
		panic("error step")
	}
}

func Int31range(from, to int32, _step ...int) int32 {
	rd := defaultRandPool.Get().(*r.Rand)
	defer defaultRandPool.Put(rd)
	stp := int32(step(_step...))
	width := to - from
	switch {
	case stp == 1:
		return from + Int31n(width)
	case stp > 0:
		n := (width + stp - 1) / stp
		return from + stp*Int31n(n)
	case stp < 0:
		n := (width + stp + 1) / stp
		return from + stp*Int31n(n)
	default:
		panic("error step")
	}
}
