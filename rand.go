package rand

import (
	"math"
	r "math/rand"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/MeteorsLiu/wyhash"
)

var (
	NV_MAGICCONST = 4 * math.Exp(-0.5) / math.Sqrt(2.0)
)

const (
	UNLOCK int32 = iota
	LOCKED
)

type lockfreeRNG struct {
	Rng     *r.Rand
	grabbed int32
}

func newlockfreeRng() *lockfreeRNG {
	return &lockfreeRNG{
		Rng: r.New(r.NewSource(int64(rng.Uint64()))),
	}
}

func (l *lockfreeRNG) Grab() bool {
	if l.grabbed == LOCKED {
		return false
	}
	return atomic.CompareAndSwapInt32(&l.grabbed, UNLOCK, LOCKED)
}

func (l *lockfreeRNG) Release() {
	atomic.StoreInt32(&l.grabbed, UNLOCK)
}

var (
	// pesudo RNG seed generator
	rng       wyhash.SRNG
	globalRng = newlockfreeRng()
	// lock free rand pool
	defaultRandPool = sync.Pool{
		New: func() any {
			return r.New(r.NewSource(int64(rng.Uint64())))
		},
	}

	safeRng = r.New(newInternalRNG())
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

func ExpFloat64() float64 {
	return safeRng.ExpFloat64()
}
func Float32() float32 {
	return safeRng.Float32()
}
func Float64() float64 {
	return safeRng.Float64()
}
func Int() int {
	return safeRng.Int()
}
func Int31() int32 {
	return safeRng.Int31()
}
func Int31n(n int32) int32 {
	return safeRng.Int31n(n)
}
func Int63() int64 {
	return safeRng.Int63()
}
func Int63n(n int64) int64 {
	return safeRng.Int63n(n)
}
func Intn(n int) int {
	return safeRng.Int()
}
func NormFloat64() float64 {
	return safeRng.NormFloat64()
}
func Perm(n int) []int {
	return safeRng.Perm(n)
}
func Read(p []byte) (n int, err error) {
	var pos int8
	var val uint64
	for n = 0; n < len(p); n++ {
		if pos == 0 {
			val = fastrand64()
			pos = 7
		}
		p[n] = byte(val)
		val >>= 8
		pos--
	}
	return
}
func Seed(seed int64) {
	return
}
func Shuffle(n int, swap func(i, j int)) {
	safeRng.Shuffle(n, swap)
}
func Uint32() uint32 {
	return safeRng.Uint32()
}
func Uint64() uint64 {
	return safeRng.Uint64()
}

func Do(f func(*r.Rand)) {
	if globalRng.Grab() {
		f(globalRng.Rng)
		globalRng.Release()
		return
	}
	rd := defaultRandPool.Get().(*r.Rand)
	f(rd)
	defaultRandPool.Put(rd)
}
func Intrange(from, to int, _step ...int) int {
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

func RandBytes(n int) []byte {
	if n <= 0 {
		return nil
	}
	b := make([]byte, n)
	Read(b)
	return b
}

func ChoiceString(a string) byte {
	if len(a) == 0 {
		return byte(0)
	}
	return a[Intn(len(a))]
}

func SampleString(s string, n int) string {
	if len(s) == 0 || n >= len(s) {
		return s
	}
	if n == 0 {
		return ""
	}
	var sb strings.Builder
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(s[Intn(len(s))])
	}
	return sb.String()
}
func Choice[T any](a []T) (r T) {
	if len(a) == 0 {
		return
	}
	r = a[Intn(len(a))]
	return
}

func Sample[T any](p []T, n int) []T {
	if len(p) == 0 || n >= len(p) {
		return p
	}

	if n == 0 {
		return p[:0]
	}
	new := make([]T, n)
	for i := 0; i < n; i++ {
		new[i] = p[Intn(len(p))]
	}
	return new
}

func Uniform32(a, b float32) float32 {
	return a + (b-a)*Float32()
}
func Uniform64(a, b float64) float64 {
	return a + (b-a)*Float64()
}

func Triangular(low, high float64, mode ...float64) float64 {
	c := 0.5
	if len(mode) > 0 {
		if high-low == 0 {
			return 0
		}
		c = (mode[0] - low) / (high - low)
	}
	u := Float64()
	if u > c {
		u = 1.0 - u
		c = 1.0 - c
		low, high = high, low
	}
	return low + (high-low)*math.Sqrt(u*c)
}
func Normalvariate(mu, sigma float64) float64 {
	var z float64
	for {
		u1 := Float64()
		u2 := 1.0 - Float64()
		z = NV_MAGICCONST * (u1 - 0.5) / u2
		zz := z * z / 4.0
		if zz <= math.Log(u2) {
			break
		}
	}
	return mu + z*sigma
}

func Gauss(mu, sigma float64) float64 {
	x2pi := Float64() * 2 * math.Pi
	g2rad := math.Sqrt(-2.0 * math.Log(1.0-Float64()))
	z := math.Cos(x2pi) * g2rad
	return mu + z*sigma
}

func Lognormvariate(mu, sigma float64) float64 {
	return math.Exp(Normalvariate(mu, sigma))
}

func Expovariate(lambd float64) float64 {
	if lambd == 0 {
		lambd = 1
	}
	return math.Log(1-Float64()) / lambd
}

func Vonmisesvariate(mu, kappa float64) float64 {
	if kappa <= 1e-6 {
		return 2 * math.Pi * Float64()
	}
	s := 0.5 / kappa
	r := s + math.Sqrt(1.0+s*s)
	var z float64
	for {
		u1 := Float64()
		z = math.Cos(math.Pi * u1)
		d := z / (r + z)
		u2 := Float64()
		if u2 < 1.0-d*d || u2 <= (1.0-d)*math.Exp(d) {
			break
		}
	}
	q := 1.0 / r
	f := (q + z) / (1.0 + q*z)
	u3 := Float64()
	var theta float64
	if u3 > 0.5 {
		theta = math.Mod(mu+math.Acos(f), 2*math.Pi)
	} else {
		theta = math.Mod(mu-math.Acos(f), 2*math.Pi)
	}

	return theta
}

func Gammavariate(alpha, beta float64) float64 {
	if alpha <= 0.0 || beta <= 0.0 {
		return 0
	}
	if alpha > 1 {
		ainv := math.Sqrt(2.0*alpha - 1.0)
		bbb := alpha - math.Log(4)
		ccc := alpha + ainv
		SG_MAGICCONST := 1.0 + math.Log(4.5)
		for {
			u1 := Float64()
			if 1e-7 < u1 && u1 < 0.9999999 {
				continue
			}
			u2 := 1.0 - Float64()
			v := math.Log(u1/(1.0-u1)) / ainv
			x := alpha * math.Exp(v)
			z := u1 * u1 * u2
			r := bbb + ccc*v - x
			if r+SG_MAGICCONST-4.5*z >= 0.0 || r >= math.Log(z) {
				return x * beta
			}
		}
	} else if alpha == 1 {
		return -math.Log(1.0-Float64()) * beta
	}
	var x float64
	for {
		u := Float64()
		b := (math.E + alpha) / math.E
		p := b * u
		if p <= 1.0 {
			x = math.Pow(p, 1.0/alpha)
		} else {
			x = -math.Log((b - p) / alpha)
		}
		u1 := Float64()
		if p > 1.0 {
			if u1 <= math.Pow(x, alpha-1.0) {
				break
			}
		} else if u1 <= math.Exp(-x) {
			break
		}
	}
	return x * beta
}
func Betavariate(alpha, beta float64) float64 {
	y := Gammavariate(alpha, 1.0)
	if y > 0 {
		return y / (y + Gammavariate(beta, 1.0))
	}
	return 0.0
}

// TODO
// func Binomialvariate(n, p float64) float64

func Paretovariate(alpha float64) float64 {
	u := 1.0 - Float64()
	return math.Pow(u, -1.0/alpha)
}

func Weibullvariate(alpha, beta float64) float64 {
	u := 1.0 - Float64()
	return alpha * math.Pow(-math.Log(u), 1.0/beta)
}
