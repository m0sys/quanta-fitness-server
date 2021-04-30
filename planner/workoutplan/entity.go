package workout

import (
	"errors"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
)

type WorkoutPlan struct {
	uuid  string
	title string
}

// NewWorkoutPlan create a new Workout.
func NewWorkoutPlan(title string) (Workout, error) {
	err := validateTitle(title)
	if err != nil {
		return WorkoutPlan{}, err

	}

	return WorkoutPlan{
		uuid:  uuid.GenerateUUID(),
		title: title,
	}, nil
}

// EditWorkoutPlanTitle edit title of Workout.
func (w *WorkoutPlan) EditWorkoutPlanTitle(title string) error {
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
