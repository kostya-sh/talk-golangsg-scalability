package benchmarks

import (
	"math/rand"
	"sync"
	"testing"
)

// START_RAND OMIT
func BenchmarkRandFloat64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Float64() // HL
		}
	})
}

// END_RAND OMIT

// START_CHANNELS OMIT
func BenchmarkRandFloat64_Channels(b *testing.B) {
	ch := make(chan float64, 1000)
	go func() {
		for {
			ch <- rand.Float64() // HL
		}
	}()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = <-ch // HL
		}
	})
}

// END_CHANNELS OMIT

// START_POOL OMIT

func BenchmarkRandFloat64_Pool(b *testing.B) {
	var pool = sync.Pool{
		New: func() interface{} {
			return rand.New(rand.NewSource(rand.Int63()))
		},
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := pool.Get().(*rand.Rand)
			r.Float64()
			pool.Put(r)
		}
	})
}

// END_POOL OMIT

// START_SOURCE OMIT
func BenchmarkRandFloat64_Source(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(rand.Int63()))
		for pb.Next() {
			r.Float64()
		}
	})
}

// END_SOURCE OMIT
