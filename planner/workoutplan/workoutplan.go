package workoutplan

import (
	"errors"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
)

type WorkoutPlan struct {
	uuid  string
	aid   string
	title string
}

// NewWorkoutPlan create a new Workout.
func NewWorkoutPlan(aid, title string) (WorkoutPlan, error) {
	err := validateTitle(title)
	if err != nil {
		return WorkoutPlan{}, err

	}

	return WorkoutPlan{
		uuid:  uuid.GenerateUUID(),
		aid:   aid,
		title: title,
	}, nil
}

// FIXME: find alternative solution for id checking...
func RestoreWorkoutPlan(id, aid, title string) (WorkoutPlan, error) {
	err := validateTitle(title)
	if err != nil {
		return WorkoutPlan{}, err

	}

	return WorkoutPlan{
		uuid:  id,
		aid:   aid,
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

func (w *WorkoutPlan) AthleteID() string {
	return w.aid
}

// Helper funcs.

func validateTitle(title string) error {
	if len(title) > 75 {
		return errors.New("Title must be less than 76 characters")
	}

	return nil
}
