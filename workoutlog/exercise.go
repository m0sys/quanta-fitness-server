package workoutlog

import (
	"errors"
	"math"
	"time"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
)

// NOTE: All units will be in SI units.

// Exercise entity for representing what an exercise is.
type Exercise struct {
	ExerciseID string
	Name       string
	Weight     float64 // in kg
	TargetRep  int
	RestTime   time.Duration // in sec
	Sets       []Set
	order      int
}

// NewExercise create a new Exercise.
func NewExercise(name string, weight, restTime float64, targetRep int, order int) (Exercise, error) {
	if err := validateExerciseFields(name, weight, restTime, targetRep); err != nil {
		return &Exercise{}, err
	}

	return &Exercise{
		ExerciseID: uuid.GenerateUUID(),
		Weight:     roundToTwoDecimalPlaces(weight),
		RestTime:   roundToTwoDecimalPlaces(restTime),
		TargetRep:  targetRep,
		order:      order,
	}, nil
}

// AddSet add Set to Exercise.
func (e *Exercise) AddSet(set Set) error {
	for _, s := range e.Sets {
		if s.SetID == set.SetID {
			return errors.New("Set is already logged")
		}
	}

	e.Sets = append(e.Sets, set)
	return nil
}

// RemoveSet remove Set from Exercise.
func (e *Exercise) RemoveSet(set Set) error {
	deleted := false

	for i, s := range e.Sets {
		if s.SetID == set.SetID {
			e.Sets = remove(e.Sets, i)
			deleted = true
		}
	}

	if !deleted {
		return errors.New("Set not found")
	}
	return nil
}

// EditExercise edit fields of Exercise.
func (e *Exercise) EditExercise(name string, weight, restTime float64, targetRep int) error {
	if err := validateExerciseFields(name, weight, restTime, targetRep); err != nil {
		return err
	}

	e.Name = name
	e.Weight = weight
	e.RestTime = restTime
	e.TargetRep = targetRep
	return nil
}

// Helper funcs.

func validateExerciseFields(name string, weight, restTime float64, targetRep int) error {
	if err := validateName(name); err != nil {
		return err
	}

	if err := validateWeight(weight); err != nil {
		return err
	}

	if err := validateRestTime(restTime); err != nil {
		return err
	}

	if err := validateTargetRep(targetRep); err != nil {
		return err
	}

	return nil
}

func validateName(title string) error {
	if len(title) > 75 {
		return errors.New("Title must be less than 76 characters")
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
