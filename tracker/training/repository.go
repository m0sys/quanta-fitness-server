package training

import (
	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	elg "github.com/mhd53/quanta-fitness-server/tracker/exerciselog"
	sl "github.com/mhd53/quanta-fitness-server/tracker/setlog"
	wl "github.com/mhd53/quanta-fitness-server/tracker/workoutlog"
)

type Repository interface {
	StoreWorkoutLog(wlog wl.WorkoutLog) error
	StoreExerciseLog(elog elg.ExerciseLog) error
	StoreSetLog(setlog sl.SetLog) error
	FindAllWorkoutLogsForAthlete(ath athlete.Athlete) ([]wl.WorkoutLog, error)
	FindAllExerciseLogsForWorkoutLog(wlog wl.WorkoutLog) ([]elg.ExerciseLog, error)
	FindAllSetLogsForExerciseLog(elog elg.ExerciseLog) ([]sl.SetLog, error)
	FindWorkoutLogByID(wlog wl.WorkoutLog) (bool, error)
	FindExerciseLogByID(elog elg.ExerciseLog) (bool, error)
	UpdateWorkoutLog(wlog wl.WorkoutLog) error
	RemoveWorkoutLog(wlog wl.WorkoutLog) error
	RemoveExerciseLog(elog elg.ExerciseLog) error
	RemoveSetLog(slog sl.SetLog) error
}
