package workout

import (
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/id"
)

type WorkoutBuilder interface {
	AddNewExercise(e Exercise) error
	RemoveExercise(eid int64) error
}

func NewWorkout(title string, date time.Time) (Workout, error) {
	err := checkTitleLength(title)
	if err != nil {
		return Workout{}, err
	}

	return Workout{
		ID:    id.NewID(),
		Title: title,
		Date:  date,
		State: "incomplete",
	}, nil
}
