package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func handleLine(line []byte) float64 {
	p1 := bytes.IndexByte(line, ',')
	p2 := bytes.IndexByte(line[p1+1:], ',') + p1 + 1
	if bytes.Equal(line[p1+1:p2], []byte("Singapore")) {
		return 0
	}
	ts, err := time.Parse(time.RFC3339, string(line[:p1]))
	if err != nil {
		log.Fatal("unable to parse time", err)
	}
	if ts.Month() != time.March {
		return 0
	}
	temp, err := strconv.ParseFloat(string(line[p2+1:]), 64)
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
		t := handleLine(s.Bytes())
		if t != 0 {
			sum += t
			k++
		}
	}
	fmt.Println(sum/float64(k), " ", k)
}
