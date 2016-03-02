package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var cities = []string{
	"Jakarta",
	"Bangkok",
	"Ho Chi Minh City",
	"Hanoi",
	"Singapore",
	"Yangon",
	"Surabaya",
	"Quezon City",
	"Bandung",
	"Bekasi",
	"Medan",
	"Tangerang",
	"Hai Phong",
	"Depok",
	"Manila",
	"Semarang",
	"Palembang",
	"Caloocan",
	"Kuala Lumpur",
	"Davao City",
	"South Tangerang",
	"Makassar",
	"Phnom Penh",
	"Batam",
	"Bogor",
	"Johor Bahru",
	"Mandalay",
	"Padang",
	"Cebu City",
	"Denpasar",
	"Malang",
	"Samarinda",
	"Zamboanga City",
	"George Town",
	"Ipoh",
}

const (
	n = 1e6
)

func main() {
	w := csv.NewWriter(os.Stdout)
	now := time.Now()
	for i := 0; i < n; i++ {
		ts := now.Add(time.Duration(-rand.Int63n(60*60*24*1000) * int64(time.Second))).Format(time.RFC3339)
		city := cities[rand.Intn(len(cities))]
		temp := fmt.Sprintf("%.3f", rand.Float64()*15+20)
		_ = w.Write([]string{ts, city, temp})
	}
	w.Flush()
}
