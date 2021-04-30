package exercise

import (
	"errors"
	"math"

	"github.com/mhd53/quanta-fitness-server/units"
)

type Exercise struct {
	uuid    string
	name    string
	metrics Metrics
}

type Metrics struct {
	targetRep    int
	weight       units.Kilogram
	restDuration units.Second
	numSets      int
}

func NewExercise(name string, metrics Metrics) Exercise {
	err := validateName(name)
	if err != nil {
		return Exercise{}, err
	}

}

// Helper funcs.

func (e *Exercise) validateName() error {
	if len(e.name) > 75 {
		return errors.New("name must be less than 76 characters")
	}

	return nil
}

func validateWeight(weight float64) error {
	if weight < 0 {
		return errors.New("weight must be a positive number")
	}
	return nil
}

func validateRestTime(restTime float64) error {
	if restTime < 0 {
		return errors.New("rest time must be a positive number")
	}
	return nil
}

func validateTargetRep(targetRep int) error {
	if targetRep < 0 {
		return errors.New("target rep must be a positive number")
	}
	return nil
}

func roundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}
