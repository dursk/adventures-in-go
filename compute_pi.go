package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var totalRuns int

func init() {
	flag.IntVar(&totalRuns, "n", 30000, "Number of times to aggregate")
}

func main() {
	flag.Parse()
	inside := 0
	for i := 0; i < totalRuns; i++ {
		randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
		x, y := randGen.Intn(101), randGen.Intn(101)
		x2 := math.Pow(float64(x), 2)
		y2 := math.Pow(float64(y), 2)
		if x2 + y2 <= 10000 {
			inside++
		}
	}
	pi := float64(inside) / float64(totalRuns) * 4
	realPi := 3.14159265359
	fmt.Printf("Pi: %f\n", pi)
	errorMargin := math.Abs(pi - realPi) / realPi * 100
	fmt.Printf("Error margin: %f\n", errorMargin)
}
