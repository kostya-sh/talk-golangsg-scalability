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

func handleLine(line string) float64 {
	ss := strings.Split(line, ",")
	if len(ss) != 3 {
		log.Fatal("invalid line", line)
	}
	if ss[1] != "Singapore" {
		return 0
	}
	ts, err := time.Parse(time.RFC3339, ss[0])
	if err != nil {
		log.Fatal("unable to parse time", err)
	}
	if ts.Month() != time.March {
		return 0
	}
	temp, err := strconv.ParseFloat(ss[2], 64)
	if err != nil {
		log.Fatal("unable to parse temperature", err)
	}
	return temp
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	sum := 0.0
	k := 0
	for s.Scan() {
		t := handleLine(s.Text())
		if t != 0 {
			sum += t
			k++
		}
	}
	fmt.Println(sum / float64(k))
}
