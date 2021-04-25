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
	AddExerciseToWorkoutLog(wlog wl.WorkoutLog, exercise wl.Exercise) (ExerciseRes, error)
	AddSetToExercise(exercise wl.Exercise, set wl.Set) (SetRes, error)
	DeleteExercise(id string) error
	DeleteSet(id string) error
	UpdateWorkoutLog(req EditWorkoutLogReq) (WorkoutLogRes, error)
	UpdateExercise(req EditExerciseReq) (ExerciseRes, error)
	UpdateSet(req EditSetReq) (SetRes, error)
	FindWorkoutLogByID(id string) (WorkoutLogRes, bool, error)
	FindAllWorkoutLogsForAthlete(ath athlete.Athlete) ([]WorkoutLogRes, error)
	FindAllExercisesForWorkoutLog(wlog wl.WorkoutLog) ([]ExerciseRes, error)
	FindAllSetsForExercise(exercise wl.Exercise) ([]SetRes, error)
}
