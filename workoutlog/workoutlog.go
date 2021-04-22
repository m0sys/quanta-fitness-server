package workoutlog

import (
	"errors"
	"sort"
	"time"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
)

// WorkoutLog entity for representing what a workout log is.
type WorkoutLog struct {
	LogID     string
	Title     string
	Date      time.Time
	Exercises []Exercise
}

// NewWorkoutLog create a new WorkoutLog.
func NewWorkoutLog(title string) (WorkoutLog, error) {
	err := validateTitle(title)
	if err != nil {
		return &WorkoutLog{}, err
	}

	return &WorkoutLog{
		LogID: uuid.GenerateUUID(),
		Title: title,
		Date:  time.Now(),
	}, nil
}

// AddExercise add Exercise to WorkoutLog.
func (l *WorkoutLog) AddExercise(exercise Exercise) error {
	for _, e := range l.Exercises {
		if e.ExerciseID == exercise.ExerciseID {
			return errors.New("Exercise already logged")
		}
	}

	l.Exercises = append(l.Exercises, Exercise)
	return nil
}

// RemoveExercise remove Exercise from WorkoutLog.
func (l *WorkoutLog) RemoveExercise(exercise Exercise) error {
	deleted := false

	for i, e := range l.Exercises {
		if e.ExerciseID == exercise.ExerciseID {
			l.Exercises = remove(l.Exercises, i)
			deleted = true
		}
	}

	if !deleted {
		return errors.New("Exercise not found")
	}
	return nil
}

// EditWorkoutLog edit title of WorkoutLog.
func (l *WorkoutLog) EditWorkoutLog(title string, date time.Time) error {
	err := validateTitle(title)
	if err != nil {
		return err
	}

	l.Title = title
	l.Date = date
	return nil
}

// OrderExercises order Exercises in WorkoutLog by their order.
func (l *WorkoutLog) OrderExercises() {
	sort.Slice(l.Exercises, func(i, j int) bool {
		return l.Exercises[i].order < l.Exercises[j].order
	})
}

// Helper funcs.

func remove(slice []interface{}, idx int) []interface{} {
	return append(slice[:s], slice[s+1:]...)
}

func validateTitle(title string) error {
	if len(title) > 75 {
		return errors.New("Title must be less than 76 characters")
	}

	return nil
}
