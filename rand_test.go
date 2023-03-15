package rand

import (
	r "math/rand"
	"sync"
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

func BenchmarkReadSmall(b *testing.B) {
	// 1 KB
	buf := make([]byte, 1024)
	b.ResetTimer()
	Read(buf)
}

func BenchmarkGoReadSmall(b *testing.B) {
	// 32 KB
	buf := make([]byte, 1024)
	b.ResetTimer()
	r.Read(buf)
}

func BenchmarkReadMedium(b *testing.B) {
	buf := make([]byte, 32*1024)
	b.ResetTimer()
	Read(buf)
}

func BenchmarkGoReadMedium(b *testing.B) {
	buf := make([]byte, 32*1024)
	b.ResetTimer()
	r.Read(buf)
}

func BenchmarkReadLarge(b *testing.B) {
	// 128 KB
	buf := make([]byte, 128*1024)
	b.ResetTimer()
	Read(buf)
}

func BenchmarkGoReadLarge(b *testing.B) {
	buf := make([]byte, 128*1024)
	b.ResetTimer()
	r.Read(buf)
}

func BenchmarkParallel(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			_ = Int()
		}()
	}
	wg.Wait()
}

func BenchmarkGoParallel(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			_ = r.Int()
		}()
	}
	wg.Wait()
}

func BenchmarkParallelRead(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			buf := make([]byte, 64)
			Read(buf)
		}()
	}
	wg.Wait()
}

func BenchmarkGoParallelRead(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			buf := make([]byte, 64)
			r.Read(buf)
		}()
	}
	wg.Wait()
}
