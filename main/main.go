package main

import (
	// "adventures-in-go/compute_pi"
	"adventures-in-go/traffic_model"
	// "github.com/davecgh/go-spew/spew"
)

func main() {
	model := traffic_model.NewNagelSchreckenberg(25, 5, 3, 0.33)
	for i := 1; i <= 1; i++ {
		model.SimulateStep()
	}
}
