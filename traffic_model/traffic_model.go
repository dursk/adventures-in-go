/*
A naive Monte Carlo simulation
using the Nagel-Schreckenberg traffic model.

See http://statweb.stanford.edu/~owen/mc/Ch-intro.pdf
for more information.
*/
package traffic_model

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/davecgh/go-spew/spew"
)

type Vehicle struct {
	velocity int
}

type Zone struct {
	vehicle *Vehicle
}

type NagelSchreckenberg struct {
	zones []Zone
	maxSpeed int
	pDecrease float64
}

func (model *NagelSchreckenberg) SimulateStep() {
	var vehicle *Vehicle
	spew.Dump(model)
	for i, zone := range model.zones {
		vehicle = zone.vehicle
		if vehicle != nil {
			// Step 1: Increase velocity by a unit
			if vehicle.velocity < model.maxSpeed {
				vehicle.velocity++
			}
			// Step 2: If velocity > distance to next car,
			// reduce velocity by a unit to avoid collision.
			y := i + 1
			for model.zones[y].vehicle == nil {
				if y == len(model.zones) - 1 {
					y = 0
				} else {
					y++;
				}
			}
			if vehicle.velocity >= (y - i) {
				vehicle.velocity--
			}
			// Step 3: if velocity > 0 decrease with probability pDecrease
			randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
			p := randGen.Float64()
			if p <= model.pDecrease && vehicle.velocity > 0 {
				vehicle.velocity--
			}
		}
	}
	fmt.Println()
	spew.Dump(model)
	fmt.Println()
	// Step 4: All vehicles move ahead by their velocity in parallel
	for i, zone := range model.zones {
		vehicle = zone.vehicle
		if vehicle == nil {
			continue
		}
		moveTo := i + vehicle.velocity
		if moveTo > len(model.zones) - 1 {
			moveTo = moveTo - len(model.zones)
		}
		if moveTo > 0 {
			model.zones[moveTo].vehicle = vehicle
			model.zones[i].vehicle = nil
		}
	}
	spew.Dump(model)
	fmt.Println()
}

func NewNagelSchreckenberg(numZones, numVehicles, maxSpeed int,
	                       pDecrease float64) *NagelSchreckenberg {
	model := NagelSchreckenberg{
		maxSpeed: maxSpeed,
		zones: make([]Zone, numZones),
    	pDecrease: pDecrease,
	}
	// numZones must be a multiple of numVehicles pl0x
	spacing := (numZones / numVehicles)
	for i := 0; i < numZones; i++ {
		if i % spacing == 0 {
			model.zones[i].vehicle = &Vehicle{velocity: 0}
		} else {
			model.zones[i].vehicle = nil
		}
	}
	return &model
}
