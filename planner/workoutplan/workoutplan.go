package workoutplan

import (
	"errors"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
)

type WorkoutPlan struct {
	uuid  string
	title string
}

// NewWorkoutPlan create a new Workout.
func NewWorkoutPlan(title string) (WorkoutPlan, error) {
	err := validateTitle(title)
	if err != nil {
		return WorkoutPlan{}, err

	}

	return WorkoutPlan{
		uuid:  uuid.GenerateUUID(),
		title: title,
	}, nil
}

// FIXME: find alternative solution.
func RestoreWorkoutPlan(id, title string) (WorkoutPlan, error) {
	err := validateTitle(title)
	if err != nil {
		return WorkoutPlan{}, err

	}

	return WorkoutPlan{
		uuid:  id,
		title: title,
	}, nil
}

// EditTitle edit title of WorkoutPlan.
func (w *WorkoutPlan) EditTitle(title string) error {
	err := validateTitle(title)
	if err != nil {
		return err
	}

	w.title = title
	return nil
}

func (w *WorkoutPlan) ID() string {
	return w.uuid
}

func (w *WorkoutPlan) Title() string {
	return w.title
}

// Helper funcs.

func validateTitle(title string) error {
	if len(title) > 75 {
		return errors.New("Title must be less than 76 characters")
	}

	return nil
}
