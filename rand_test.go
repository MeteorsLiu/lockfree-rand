package rand

import (
	r "math/rand"
	"testing"
)

func TestSeed(t *testing.T) {
	t.Log(getTimeBasedSeed())
	t.Log(getTimeBasedSeed())
	t.Log(getTimeBasedSeed())
}

func TestAll(t *testing.T) {
	t.Log(ExpFloat64())
	t.Log(Float32())
	t.Log(Float64())
	t.Log(Int())
	t.Log(Int31())
	t.Log(Int31n(123))
	t.Log(Int63())
	t.Log(Int63n(123))
	t.Log(Intn(123))
	t.Log(NormFloat64())
	t.Log(Perm(5))
	t.Log(Uint32())
	t.Log(Uint64())
	t.Log(Int31range(10, 20))
	t.Log(Intrange(100, 200))
	t.Log(Int63range(10522, 20453))
}

func BenchmarkInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Int()
	}
}

func BenchmarkGoInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.Int()
	}
}
