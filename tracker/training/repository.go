package training

import (
	elg "github.com/m0sys/quanta-fitness-server/tracker/exerciselog"
	sl "github.com/m0sys/quanta-fitness-server/tracker/setlog"
	wl "github.com/m0sys/quanta-fitness-server/tracker/workoutlog"
)

type Repository interface {
	StoreWorkoutLog(wlog wl.WorkoutLog) error
	StoreExerciseLog(elog elg.ExerciseLog) error
	StoreSetLog(setlog sl.SetLog) error
	FindAllWorkoutLogsForAthlete(id string) ([]wl.WorkoutLog, error)
	FindAllExerciseLogsForWorkoutLog(wlog wl.WorkoutLog) ([]elg.ExerciseLog, error)
	FindAllSetLogsForExerciseLog(elog elg.ExerciseLog) ([]sl.SetLog, error)
	FindWorkoutLogByID(id string) (wl.WorkoutLog, bool, error)
	FindExerciseLogByID(id string) (elg.ExerciseLog, bool, error)
	UpdateWorkoutLog(wlog wl.WorkoutLog) error
	RemoveWorkoutLog(wlog wl.WorkoutLog) error
	RemoveExerciseLog(elog elg.ExerciseLog) error
	RemoveSetLog(slog sl.SetLog) error
}
