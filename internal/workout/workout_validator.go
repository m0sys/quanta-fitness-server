package workout

import (
	"errors"
	"github.com/go-playground/validator/v10"

	"github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type WorkoutValidator interface {
	ValidateCreateWorkout(title string) error
	ValidateUpdateWorkout(workout entity.Workout) error
	ValidateDeleteWorkout(wid int64) error
	ValidateGetWorkoutExercises(wid int64) error
}

type workoutValidator struct{}

var (
	validate = validator.New()
	vws      workoutstore.WorkoutStore
)

func NewWorkoutValidator(
	workoutstore workoutstore.WorkoutStore) WorkoutValidator {
	vws = workoutstore
	return &workoutValidator{}
}

func (*workoutValidator) ValidateCreateWorkout(title string) error {

	err := checkTitleLength(title)
	if err != nil {
		return err
	}

	return nil

}

func checkTitleLength(title string) error {
	if validate.Var(title, "required,max=75") != nil {
		return errors.New("Title must be less than 76 characters!")
	}
	return nil
}

func (*workoutValidator) ValidateUpdateWorkout(workout entity.Workout) error {
	return nil

}

func (*workoutValidator) ValidateDeleteWorkout(wid int64) error {
	return nil

}

func (*workoutValidator) ValidateGetWorkoutExercises(wid int64) error {
	return nil

}
