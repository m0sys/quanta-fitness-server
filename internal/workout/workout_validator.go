package workout

import (
	"errors"
	"github.com/go-playground/validator/v10"

	"github.com/m0sys/quanta-fitness-server/internal/entity"
)

type WorkoutValidator interface {
	ValidateCreateWorkout(title string) error
	ValidateUpdateWorkout(workout entity.BaseWorkout) error
}

type workoutValidator struct{}

func NewWorkoutValidator() WorkoutValidator {
	return &workoutValidator{}
}

func (*workoutValidator) ValidateCreateWorkout(title string) error {
	return validateTitle(title)
}

func validateTitle(title string) error {
	return checkTitleLength(title)
}

func checkTitleLength(title string) error {
	validate := validator.New()
	if validate.Var(title, "required,max=75") != nil {
		return errors.New("Title must be less than 76 characters!")
	}
	return nil
}

func (*workoutValidator) ValidateUpdateWorkout(workout entity.BaseWorkout) error {
	return validateTitle(workout.Title)

}
