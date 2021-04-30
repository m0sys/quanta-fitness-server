package planner

import (
	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/units"
)

type Exercise struct {
	uuid    string
	athlete athlete.Athlete
	name    string
	metrics Metrics
}

type Metrics struct {
	targetRep    int
	weight       units.Kilogram
	restDuration units.Second
	numSets      int
}
