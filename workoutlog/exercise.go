// Exercise contains the Exercise entity.

package workoutlog

import (
	"errors"
	"math"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
)

// NOTE: All units will be in SI units.

// Exercise entity for representing what an exercise is.
type Exercise struct {
	uuid      string
	name      string
	weight    float64 // in kg
	targetRep int
	restTime  float64 // in sec
	sets      []Set
	order     int
}

// NewExercise create a new Exercise.
func NewExercise(name string, weight, restTime float64, targetRep int, order int) (Exercise, error) {
	if err := validateExerciseFields(name, weight, restTime, targetRep); err != nil {
		return Exercise{}, err
	}

	weight = roundToTwoDecimalPlaces(weight)
	restTime = roundToTwoDecimalPlaces(restTime)

	return Exercise{
		uuid:      uuid.GenerateUUID(),
		name:      name,
		weight:    roundToTwoDecimalPlaces(weight),
		restTime:  roundToTwoDecimalPlaces(restTime),
		targetRep: targetRep,
		order:     order,
	}, nil
}

func (e *Exercise) ExerciseID() string {
	return e.uuid
}

func (e *Exercise) Name() string {
	return e.name
}

func (e *Exercise) Weight() float64 {
	return e.weight
}

func (e *Exercise) TargetRep() int {
	return e.targetRep
}

func (e *Exercise) RestTime() float64 {
	return e.restTime
}

func (e *Exercise) InsertSet(s Set, i int) {
	e.sets[i] = s
}

func (e *Exercise) Sets() []Set {
	tmp := make([]Set, len(e.sets))
	copy(tmp, e.sets)
	return tmp
}

// AddSet add Set to Exercise.
func (e *Exercise) AddSet(set Set) error {
	for _, s := range e.sets {
		if s.SetID() == set.SetID() {
			return errors.New("Set is already logged")
		}
	}

	e.sets = append(e.sets, set)
	return nil
}

// RemoveSet remove Set from Exercise.
func (e *Exercise) RemoveSet(set Set) error {
	deleted := false

	for i, s := range e.sets {
		if s.SetID() == set.SetID() {
			e.sets = removeSet(e.sets, i)
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

	e.name = name
	e.weight = weight
	e.restTime = restTime
	e.targetRep = targetRep
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

func removeSet(slice []Set, idx int) []Set {
	return append(slice[:idx], slice[idx+1:]...)
}
