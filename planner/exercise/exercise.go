package exercise

/*
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
	TargetRep    int
	Weight       units.Kilogram
	RestDuration units.Second
	NumSets      int
}

func NewExercise(name string, metrics Metrics) Exercise {
	err := validateExerciseFields(
		name,
		metrics.Weight,
		metrics.RestDuration,
		metrics.TargetRep,
		metrics.NumSets,
	)

	if err != nil {
		return Exercise{}, err
	}
}

// Helper funcs.

func validateExerciseFields(name string, weight, restDur float64, targetRep int, numSets int) error {
	if err := validateName(name); err != nil {
		return err
	}

	if err := validateWeight(weight); err != nil {
		return err
	}

	if err := validateRestDur(restDur); err != nil {
		return err
	}

	if err := validateTargetRep(targetRep); err != nil {
		return err
	}

	if err := validateNumSets(numSets); err != nil {
		return err
	}

	return nil
}

func validateName(name string) error {
	if len(name) > 75 {
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

func validateRestDur(restDur float64) error {
	if restTime < 0 {
		return errors.New("rest duration must be a positive number")
	}
	return nil
}

func validateTargetRep(targetRep int) error {
	if targetRep < 0 {
		return errors.New("target rep must be a positive number")
	}
	return nil
}

func validateNumSets(numSets int) error {
	if numSets < 0 {
		return errors.New("num sets must be a positive number")
	}
	return nil
}

func roundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}
*/
