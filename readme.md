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
    n := rand.Intn(5)
    for i :=0; i < n; i++ {
        randseed := rand.Int()
    }
})

```


# Benchmark

```
goos: windows
goarch: amd64
pkg: github.com/MeteorsLiu/rand
cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
BenchmarkInt-8                          64142653                18.63 ns/op            0 B/op          0 allocs/op
BenchmarkGoInt-8                        85382512                16.28 ns/op            0 B/op          0 allocs/op
BenchmarkReadSmall-8                     1422010               845.7 ns/op             0 B/op          0 allocs/op
BenchmarkGoReadSmall-8                   1387641               852.5 ns/op             0 B/op          0 allocs/op
BenchmarkReadMedium-8                      45312             26719 ns/op               0 B/op          0 allocs/op
BenchmarkGoReadMedium-8                    44556             26519 ns/op               0 B/op          0 allocs/op
BenchmarkReadLarge-8                       10000            108856 ns/op               0 B/op          0 allocs/op
BenchmarkGoReadLarge-8                      9871            106127 ns/op               0 B/op          0 allocs/op
BenchmarkParallel-8                      4756483               247.7 ns/op            16 B/op          1 allocs/op
BenchmarkGoParallel-8                    4388599               275.8 ns/op            16 B/op          1 allocs/op
BenchmarkParallelRead-8                  4715338               254.1 ns/op            16 B/op          1 allocs/op
BenchmarkGoParallelRead-8                2627706               427.8 ns/op            16 B/op          1 allocs/op
BenchmarkWyhashParallelRead-8            3529256               341.9 ns/op            24 B/op          1 allocs/op
BenchmarkWyhashPoolParallelRead-8        4086340               291.5 ns/op            48 B/op          1 allocs/op
BenchmarkGoMultipleDo-8                  1000000              1002 ns/op              19 B/op          1 allocs/op
BenchmarkMultipleDo-8                    4532272               266.9 ns/op            16 B/op          1 allocs/op
```
The prefix of BenchmarkGo stands for **math/rand** function, without which, it stands for this repo.

Optimized for high concurrent cases, I use sync.Pool to avoid the global mutex lock.
