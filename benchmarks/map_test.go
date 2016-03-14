package benchmarks

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// START_RO OMIT
const n = 100

func BenchmarkMap_Readonly(b *testing.B) {
	m := map[string]string{} // HL
	for i := 0; i < n; i++ { // HL
		m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i) // HL
	} // HL

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if m[fmt.Sprintf("key%d", i)] == "impossible value" { // HL
				fmt.Println("should not be printed")
			}
			i = (i + 1) % n
		}
	})
}

// END_RO OMIT

// START_NOSYNC OMIT
func BenchmarkMap_NoSync(b *testing.B) {
	// initialization (omitted)
	m := map[string]string{} // OMIT
	for i := 0; i < n; i++ { // OMIT
		m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i) // OMIT
	} // OMIT

	go func() {
		for i := 0; ; i = (i + 1) % n {
			m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("newvalue%d", i) // HL
			time.Sleep(100 * time.Millisecond)
		}
	}()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if m[fmt.Sprintf("key%d", i)] == "impossible value" { // HL
				fmt.Println("should not be printed")
			}
			i = (i + 1) % n
		}
	})
}

// END_NOSYNC OMIT

// START_MUTEX OMIT
func BenchmarkMap_Mutex(b *testing.B) {
	// initialization (omitted)
	m := map[string]string{} // OMIT
	for i := 0; i < n; i++ { // OMIT
		m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i) // OMIT
	} // OMIT
	mu := sync.Mutex{} // HL
	go func() {
		for i := 0; ; i = (i + 1) % n {
			mu.Lock() // HL
			m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("newvalue%d", i)
			mu.Unlock() // HL
			time.Sleep(100 * time.Millisecond)
		}
	}()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			mu.Lock() // HL
			if m[fmt.Sprintf("key%d", i)] == "impossible value" {
				fmt.Println("should not be printed")
			}
			mu.Unlock() // HL
			i = (i + 1) % n
		}
	})
}

// END_MUTEX OMIT

// START_RWMUTEX OMIT
func BenchmarkMap_RWMutex(b *testing.B) {
	// initialization (omitted)
	m := map[string]string{} // OMIT
	for i := 0; i < n; i++ { // OMIT
		m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i) // OMIT
	} // OMIT
	mu := sync.RWMutex{} // HL
	go func() {
		for i := 0; ; i = (i + 1) % n {
			mu.Lock()
			m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("newvalue%d", i)
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			mu.RLock() // HL
			if m[fmt.Sprintf("key%d", i)] == "impossible value" {
				fmt.Println("should not be printed")
			}
			mu.RUnlock() // HL
			i = (i + 1) % n
		}
	})
}

// END_RWMUTEX OMIT

func clone(m map[string]string) map[string]string {
	cm := make(map[string]string)
	for k, v := range m {
		cm[k] = v
	}
	return cm
}

// START_ATOMIC OMIT
func BenchmarkMap_Atomic(b *testing.B) {
	// initialization (omitted)
	m := map[string]string{} // OMIT
	for i := 0; i < n; i++ { // OMIT
		m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i) // OMIT
	} // OMIT
	mv := atomic.Value{} // HL
	mv.Store(m)          // HL
	go func() {
		for i := 0; ; i = (i + 1) % n {
			mm := clone(mv.Load().(map[string]string)) // HL
			mm[fmt.Sprintf("key%d", i)] = fmt.Sprintf("newvalue%d", i)
			mv.Store(mm) // HL
			time.Sleep(100 * time.Millisecond)
		}
	}()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			mm := mv.Load().(map[string]string) // HL
			if mm[fmt.Sprintf("key%d", i)] == "impossible value" {
				fmt.Println("should not be printed")
			}
			i = (i + 1) % n
		}
	})
}

// END_ATOMIC OMIT
