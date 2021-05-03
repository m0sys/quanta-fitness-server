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
	RemoveWorkoutLog(wlog wl.WorkoutLog) error
	FindAllWorkoutLogsForAthlete(ath athlete.Athlete) ([]wl.WorkoutLog, error)
	FindAllExerciseLogsForWorkoutLog(wlog wl.WorkoutLog) ([]elg.ExerciseLog, error)
	FindAllSetLogsForExerciseLog(elog elg.ExerciseLog) ([]sl.SetLog, error)
}
