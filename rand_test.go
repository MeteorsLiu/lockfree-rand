package rand

import (
	r "math/rand"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/MeteorsLiu/wyhash"
)

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
	t.Log(Uniform32(30.5, 55.5))
	t.Log(Uniform64(30.5, 55.5))

	Do(func(rd *r.Rand) {
		t.Log(rd.Int())
		t.Log(rd.Intn(50))
		t.Log(rd.Uint64())
	})

	buf := make([]byte, 64)
	Read(buf)
	t.Log(buf)
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
	for i := 0; i < b.N; i++ {
		Read(buf)
	}
}

func BenchmarkGoReadSmall(b *testing.B) {
	// 32 KB
	buf := make([]byte, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Read(buf)
	}
}

func BenchmarkReadMedium(b *testing.B) {
	buf := make([]byte, 32*1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Read(buf)
	}
}

func BenchmarkGoReadMedium(b *testing.B) {
	buf := make([]byte, 32*1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Read(buf)
	}
}

func BenchmarkReadLarge(b *testing.B) {
	// 128 KB
	buf := make([]byte, 128*1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Read(buf)
	}
}

func BenchmarkGoReadLarge(b *testing.B) {
	buf := make([]byte, 128*1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Read(buf)
	}
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

func BenchmarkWyhashParallelRead(b *testing.B) {
	var rng wyhash.SRNG
	var wg sync.WaitGroup
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			buf := make([]byte, 64)
			rng.Read(buf)
		}()
	}
	wg.Wait()
}

func BenchmarkWyhashPoolParallelRead(b *testing.B) {
	var lock int32
	var rng wyhash.RNG
	var srng wyhash.SRNG
	var wg sync.WaitGroup
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			buf := make([]byte, 64)
			if lock == 0 {
				if atomic.CompareAndSwapInt32(&lock, 0, 1) {
					rng.Read(buf)
					atomic.StoreInt32(&lock, 0)
					return
				}
			}
			srng.Read(buf)
		}()
	}
	wg.Wait()
}

func BenchmarkGoMultipleDo(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			buf := make([]byte, 64)
			r.Read(buf)
			_ = r.Int()
			_ = r.Intn(50)
			_ = r.Uint64()
		}()
	}
	wg.Wait()
}

func BenchmarkMultipleDo(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			buf := make([]byte, 64)
			Do(func(rd *r.Rand) {
				rd.Read(buf)
				_ = rd.Int()
				_ = rd.Intn(50)
				_ = rd.Uint64()
			})
		}()
	}
	wg.Wait()
}
