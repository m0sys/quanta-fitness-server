package record

import "errors"

var (
	errAccessDenied    = errors.New("Access denied!")
	errWorkoutNotFound = errors.New("Workout not found!")
)
