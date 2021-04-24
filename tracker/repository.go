package tracker

import (
	"github.com/mhd53/quanta-fitness-server/athlete"
	wl "github.com/mhd53/quanta-fitness-server/workoutlog"
)

// Repository repo for persiting all WorkoutLog related data.
type Repository interface {
	workoutLogRepo
}

type workoutLogRepo interface {
	StoreWorkoutLog(wlog wl.WorkoutLog, ath athlete.Athlete) (WorkoutLogRes, error)
	FindWorkoutLogByID(id string) (WorkoutLogRes, bool, error)
	FindAllExercisesForWorkoutLog(wlog wl.WorkoutLog) ([]ExerciseRes, error)
	AddExerciseToWorkoutLog(wlog wl.WorkoutLog, exercise wl.Exercise) (ExerciseRes, error)
	AddSetToExercise(exercise wl.Exercise, set wl.Set) (SetRes, error)
	DeleteExercise(id string) error
	DeleteSet(id string) error
	// FindAllWorkoutLogsForAthlete(ath athlete.Athlete) ([]WorkoutLogRes, error)
}
