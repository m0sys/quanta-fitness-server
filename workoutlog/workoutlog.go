// WorkoutLog contains the WorkoutLog entity.

package workoutlog

import (
	"errors"
	"sort"
	"time"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
)

// WorkoutLog entity for representing what a workout log is.
type WorkoutLog struct {
	uuid      string
	title     string
	date      time.Time
	exercises []Exercise
}

// NewWorkoutLog create a new WorkoutLog.
func NewWorkoutLog(title string) (WorkoutLog, error) {
	err := validateTitle(title)
	if err != nil {
		return WorkoutLog{}, err
	}

	return WorkoutLog{
		uuid:  uuid.GenerateUUID(),
		title: title,
		date:  time.Now(),
	}, nil
}

func (l *WorkoutLog) LogID() string {
	return l.uuid
}

func (l *WorkoutLog) Title() string {
	return l.title
}

func (l *WorkoutLog) Date() time.Time {
	return l.date
}

func (l *WorkoutLog) NumExercises() int {
	return len(l.exercises)
}

func (l *WorkoutLog) InsertExercise(e Exercise, i int) {
	l.exercises[i] = e
}

func (l *WorkoutLog) Exercises() []Exercise {
	tmp := make([]Exercise, len(l.exercises))
	copy(tmp, l.exercises)
	return tmp
}

// AddExercise add Exercise to WorkoutLog.
func (l *WorkoutLog) AddExercise(exercise Exercise) error {
	for _, e := range l.exercises {
		if e.ExerciseID() == exercise.ExerciseID() {
			return errors.New("Exercise already logged")
		}
	}

	l.exercises = append(l.exercises, exercise)
	return nil
}

// RemoveExercise remove Exercise from WorkoutLog.
func (l *WorkoutLog) RemoveExercise(exercise Exercise) error {
	deleted := false

	for i, e := range l.exercises {
		if e.ExerciseID() == exercise.ExerciseID() {
			l.exercises = removeExercise(l.exercises, i)
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

	l.title = title
	l.date = date
	return nil
}

// OrderExercises order Exercises in WorkoutLog by their order.
func (l *WorkoutLog) OrderExercises() {
	sort.Slice(l.exercises, func(i, j int) bool {
		return l.exercises[i].order < l.exercises[j].order
	})
}

// Helper funcs.

func removeExercise(slice []Exercise, idx int) []Exercise {
	return append(slice[:idx], slice[idx+1:]...)
}

func validateTitle(title string) error {
	if len(title) > 75 {
		return errors.New("Title must be less than 76 characters")
	}

	return nil
}
