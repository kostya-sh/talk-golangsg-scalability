package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func monteCalroGlobalRand(n int) int {
	k := 0
	for i := 0; i < n; i++ {
		x := rand.Float64() - 1
		y := rand.Float64() - 1
		if x*x+y*y <= 1 {
			k++
		}
	}
	return k
}

func monteCalro(n int) int {
	r := rand.New(rand.NewSource(rand.Int63()))

	k := 0
	for i := 0; i < n; i++ {
		x := r.Float64() - 1
		y := r.Float64() - 1
		if x*x+y*y <= 1 {
			k++
		}
	}
	return k
}

func main() {
	const n = 1e8

	cores := runtime.NumCPU()
	res := make(chan int)
	for i := 0; i < cores; i++ {
		go func() {
			res <- monteCalro(n / cores)
		}()
	}

	k := 0
	for i := 0; i < cores; i++ {
		k += <-res
	}

	fmt.Println(4 * float64(k) / n)
}
