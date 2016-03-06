package benchmarks

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkRWRead(b *testing.B) {
	//var rwmu sync.RWMutex
	var mu sync.Mutex
	x := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}

	go func() {
		var n = 0
		for {
			mu.Lock()
			x[n]++
			n = (n + 1) % len(x)
			mu.Unlock()

			time.Sleep(100 * time.Millisecond)
		}
	}()

	b.RunParallel(func(pb *testing.PB) {
		var n int
		for pb.Next() {
			mu.Lock()
			if x[n] != 0 {
				n = (n + 1) % len(x)
			}
			mu.Unlock()
		}
	})
}

func BenchmarkAtomic(b *testing.B) {
	x := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
	var m atomic.Value
	m.Store(x)

	go func() {
		var n = 0
		for {
			xx := m.Load().(map[int]int)
			xxx := make(map[int]int)
			for k, v := range xx {
				xxx[k] = v
			}
			xxx[n]++
			n = (n + 1) % len(xxx)
			m.Store(xxx)

			time.Sleep(100 * time.Millisecond)
		}
	}()

	b.RunParallel(func(pb *testing.PB) {
		var n int
		for pb.Next() {
			xx := m.Load().(map[int]int)
			if xx[n] != 0 {
				n = (n + 1) % len(xx)
			}
		}
	})
}
