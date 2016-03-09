package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// START_HL OMIT
func handleLine(line string) float64 {
	tokens := strings.Split(line, ",") // HL
	if len(tokens) != 3 {
		log.Fatal("invalid line: ", line)
	}
	if tokens[1] != "Singapore" { // HL
		return math.NaN() // HL
	} // HL
	ts, err := time.Parse(time.RFC3339, tokens[0])
	if err != nil {
		log.Fatal("unable to parse time: ", err)
	}
	if ts.Month() != time.March { // HL
		return math.NaN() // HL
	} // HL
	t, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		log.Fatal("unable to parse temperature: ", err)
	}
	return t // HL
}

// END_HL OMIT

// START_MAIN OMIT
func main() {
	s := bufio.NewScanner(os.Stdin)
	out := make(chan float64, 100) // HL
	n := 0
	for s.Scan() {
		n++
		line := s.Text()
		go func() {
			out <- handleLine(line) // HL
		}()
	}
	if s.Err() != nil {
		log.Fatal("scan: ", s.Err())
	}
	sum := 0.0
	cnt := 0
	for i := 0; i < n; i++ {
		if t := <-out; !math.IsNaN(t) { // HL
			sum += t
			cnt++
		}
	}
	fmt.Println(sum / float64(cnt))
}

// END_MAIN OMIT
