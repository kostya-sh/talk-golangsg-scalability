package benchmarks

import (
	"math/rand"
	"sync"
	"testing"
)

// START_GLOBAL OMIT
func BenchmarkRandFloat64_Global(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Float64() // HL
		}
	})
}

// END_GLOBAL OMIT

// START_CHANNEL OMIT
func BenchmarkRandFloat64_Channel(b *testing.B) {
	ch := make(chan float64, 1000) // HL
	go func() {
		r := rand.New(rand.NewSource(rand.Int63())) // HL
		for {
			ch <- r.Float64() // HL
		}
	}()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = <-ch // HL
		}
	})
}

// END_CHANNEL OMIT

// START_POOL OMIT

func BenchmarkRandFloat64_Pool(b *testing.B) {
	var pool = sync.Pool{ // HL
		New: func() interface{} { // HL
			return rand.New(rand.NewSource(rand.Int63())) // HL
		}, // HL
	} // HL

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := pool.Get().(*rand.Rand) // HL
			r.Float64()
			pool.Put(r) // HL
		}
	})
}

// END_POOL OMIT

// START_SOURCE OMIT
func BenchmarkRandFloat64_Source(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(rand.Int63())) // HL
		for pb.Next() {
			r.Float64()
		}
	})
}

// END_SOURCE OMIT
