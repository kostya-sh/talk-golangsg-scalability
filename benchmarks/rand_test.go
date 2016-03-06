package benchmarks

import (
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkInt63Threadsafe(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Int63()
		}
	})
}

func BenchmarkInt63Unthreadsafe(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(1))
		for pb.Next() {
			r.Int63()
		}
	})
}

func newRand() interface{} {
	return rand.New(rand.NewSource(rand.Int63()))
}

func BenchmarkInt63ThreadsafePooledRand(b *testing.B) {
	var pool = sync.Pool{New: newRand}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r := pool.Get().(*rand.Rand)
			r.Int63()
			pool.Put(r)
		}
	})
}

func BenchmarkInt63Channels(b *testing.B) {
	randCh := make(chan int64, 100*runtime.NumCPU())

	go func() {
		for {
			randCh <- rand.Int63()
		}
	}()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			<-randCh
		}
	})
}
