package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func handler(in chan string, out chan float64) {
	for line := range in {
		ss := strings.Split(line, ",")
		if len(ss) != 3 {
			log.Fatal("invalid line", line)
		}
		if ss[1] != "Singapore" {
			continue
		}
		ts, err := time.Parse(time.RFC3339, ss[0])
		if err != nil {
			log.Fatal("unable to parse time", err)
		}
		if ts.Month() != time.March {
			continue
		}
		temp, err := strconv.ParseFloat(ss[2], 64)
		if err != nil {
			log.Fatal("unable to parse temperature", err)
		}
		out <- temp
	}
	wg.Done()
}

func main() {
	res := make(chan float64, 100)
	in := make(chan string, 100)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go handler(in, res)
	}

	r2 := make(chan float64)
	go func() {
		sum := 0.0
		k := 0
		for t := range res {
			sum += t
			k++
		}
		r2 <- sum / float64(k)
	}()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		in <- s.Text()
	}
	close(in)
	wg.Wait()
	close(res)

	fmt.Println(<-r2)
}
