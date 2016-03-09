package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// START_MC OMIT
func monteCalro(n int) int {
	k := 0
	for i := 0; i < n; i++ {
		x := rand.Float64() // HL
		y := rand.Float64() // HL
		if x*x+y*y <= 1 {
			k++
		}
	}
	return k
}

// END_MC OMIT

// START_MAIN OMIT
func main() {
	const n = 1e8

	cores := runtime.NumCPU() // HL

	res := make(chan int) // HL
	for i := 0; i < cores; i++ {
		go func() {
			res <- monteCalro(n / cores) // HL
		}()
	}

	k := 0
	for i := 0; i < cores; i++ {
		k += <-res // HL
	}
	fmt.Println(4 * float64(k) / n)
}

// END_MAIN OMIT
