# An optimized math/rand

Concurrency safe math/rand, optimized for high concurrent pseudo number generating.

# Why?
Go's random algorithm is great, its pseudo randomness behaves well.

However, the standard library uses a global mutex lock to ensure its safety of concurrency.

That's not a good things.

In a high concorrent pseudo number generating case, like a online game, the global mutex will increase the latency of requests when the parallel requests is incrasing.

## How to solve?

It's easy to avoid the lock.

1. Pool
2. Atomic operation

In this porject, I choose to use the pool. The other one([Wyhash](https://github.com/MeteorsLiu/wyhash)) uses Atomic Operation.

However, if only one goroutine is using, there's no necessity to use the pool, because pulling out the pool and pushing into the pool cost, it's not free.


So I use a CAS lock to check whether only one goroutine uses it.

If yes, use a global lockless one.No, grab it from the pool.


# What's new?

I not only optimize the performance of concurrency, but also add some new things.

Like Do(), Intrange(), Uniform(), Sample(), ReadBytes(), Triangular()...

They are quite similar with the Python Standard libray.

Exclude Do().

What does Do use for?

In some cases, we might call Int(), Intn() multiple times.

Like
```
n := rand.Intn(5)
for i :=0; i < n; i++ {
    randseed := rand.Int()
}
```

If using a global mutex, locking and unlocking cost too much, that's not good.

However, if we can get a lockless one from the pool, and generate it multiple times.

There's no cost of locks.

That's the Do() do.

```
Do(func (r *rand.Rand) {
    // No more lock here
    n := r.Intn(5)
    for i :=0; i < n; i++ {
        randseed := r.Int()
    }
})

```


# Benchmark

```
goos: windows
goarch: amd64
pkg: github.com/MeteorsLiu/lockfree-rand
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkInt-8                          247086846                4.794 ns/op           0 B/op          0 allocs/op
BenchmarkGoInt-8                        73754803                16.29 ns/op            0 B/op          0 allocs/op
BenchmarkReadSmall-8                     1595437               759.7 ns/op             0 B/op          0 allocs/op
BenchmarkGoReadSmall-8                   1356376               887.3 ns/op             0 B/op          0 allocs/op
BenchmarkReadMedium-8                      49314             24041 ns/op               0 B/op          0 allocs/op
BenchmarkGoReadMedium-8                    43911             27435 ns/op               0 B/op          0 allocs/op
BenchmarkReadLarge-8                       12649             97187 ns/op               0 B/op          0 allocs/op
BenchmarkGoReadLarge-8                     10000            110534 ns/op               0 B/op          0 allocs/op
BenchmarkParallel-8                      4934991               253.0 ns/op            16 B/op          1 allocs/op
BenchmarkGoParallel-8                    4218584               282.6 ns/op            16 B/op          1 allocs/op
BenchmarkParallelRead-8                  4548276               258.9 ns/op            16 B/op          1 allocs/op
BenchmarkGoParallelRead-8                2720745               443.8 ns/op            17 B/op          1 allocs/op
BenchmarkWyhashParallelRead-8            3443152               341.4 ns/op            24 B/op          1 allocs/op
BenchmarkWyhashPoolParallelRead-8        3987759               300.0 ns/op            48 B/op          1 allocs/op
BenchmarkGoMultipleDo-8                  1210926               995.7 ns/op            18 B/op          1 allocs/op
BenchmarkMultipleDo-8                    4430743               276.3 ns/op            16 B/op          1 allocs/op
```
The prefix of BenchmarkGo stands for **math/rand** function, without which, it stands for this repo.

Optimized for high concurrent cases, I use sync.Pool to avoid the global mutex lock.
