package training

import "fmt"

const (
	errSlur = "training"
)

var (
	errInternal           = fmt.Errorf("%s: internal error", errSlur)
	ErrWorkoutLogNotFound = fmt.Errorf("%s: WorkoutLog not found", errSlur)
	ErrUnauthorizedAccess = fmt.Errorf("%s: unauthorized access", errSlur)
)
