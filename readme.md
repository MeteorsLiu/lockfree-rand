# An optimized math/rand

Benchmark:

```
goos: linux
goarch: amd64
pkg: github.com/MeteorsLiu/lockfree-rand
cpu: Common KVM processor
BenchmarkInt-8              	67152585	        18.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoInt-8            	75535684	        16.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkReadSmall-8        	1000000000	         0.0000029 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoReadSmall-8      	1000000000	         0.0000026 ns/op	       0 B/op	       0 allocs/op
BenchmarkReadMedium-8       	1000000000	         0.0000399 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoReadMedium-8     	1000000000	         0.0000399 ns/op	       0 B/op	       0 allocs/op
BenchmarkReadLarge-8        	1000000000	         0.0001642 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoReadLarge-8      	1000000000	         0.0001864 ns/op	       0 B/op	       0 allocs/op
BenchmarkParallel-8         	 4170608	       303.4 ns/op	      16 B/op	       1 allocs/op
BenchmarkGoParallel-8       	 3891436	       321.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkParallelRead-8     	 3827606	       316.3 ns/op	      16 B/op	       1 allocs/op
BenchmarkGoParallelRead-8   	 2025992	       587.8 ns/op	      20 B/op	       1 allocs/op
```
The prefix of BenchmarkGo stands for **math/rand** function, without which, it stands for this repo.

Optimized for high concurrent cases, I use sync.Pool to avoid the global mutex lock.
