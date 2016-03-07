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

func handleLine(line string) float64 {
	tokens := strings.Split(line, ",")
	if len(tokens) != 3 {
		log.Fatal("invalid line: ", line)
	}
	if tokens[1] != "Singapore" {
		return math.NaN()
	}
	ts, err := time.Parse(time.RFC3339, tokens[0])
	if err != nil {
		log.Fatal("unable to parse time: ", err)
	}
	if ts.Month() != time.March {
		return math.NaN()
	}
	t, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		log.Fatal("unable to parse temperature: ", err)
	}
	return t
}

// START_MAIN OMIT
func main() {
	s := bufio.NewScanner(os.Stdin)
	sum := 0.0
	cnt := 0
	for s.Scan() {
		t := handleLine(s.Text()) // HL
		if !math.IsNaN(t) {
			sum += t
			cnt++
		}
	}
	if s.Err() != nil {
		log.Fatal("scan: ", s.Err())
	}
	fmt.Println(sum / float64(cnt))
}

// END_MAIN OMIT
