package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleLine(res chan float64, line string) {
	ss := strings.Split(line, ",")
	if len(ss) != 3 {
		log.Fatal("invalid line", line)
	}
	if ss[1] != "Singapore" {
		res <- 0.0
		return
	}
	ts, err := time.Parse(time.RFC3339, ss[0])
	if err != nil {
		log.Fatal("unable to parse time", err)
	}
	if ts.Month() != time.March {
		res <- 0.0
		return
	}
	temp, err := strconv.ParseFloat(ss[2], 64)
	if err != nil {
		log.Fatal("unable to parse temperature", err)
	}
	res <- temp
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	res := make(chan float64, 100)
	n := 0
	for s.Scan() {
		n++
		go handleLine(res, s.Text())
	}
	sum := 0.0
	k := 0
	for i := 0; i < n; i++ {
		t := <-res
		if t != 0 {
			sum += t
			k++
		}
	}
	fmt.Println(sum / float64(k))
}
