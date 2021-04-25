package tracker

import "errors"

var (
	errNilWorkoutLog    = errors.New("no WorkoutLog is assigned to Tracker")
	errLogIDMismatch    = errors.New("WorkoutLog does not match requested LogID")
	errInternal         = errors.New("Internal Error")
	errExerciseNotFound = errors.New("Exercise not found")
)
