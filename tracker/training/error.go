package training

import "fmt"

const (
	errSlur = "training"
)

var (
	errInternal                   = fmt.Errorf("%s: internal error", errSlur)
	ErrWorkoutLogNotFound         = fmt.Errorf("%s: WorkoutLog not found", errSlur)
	ErrUnauthorizedAccess         = fmt.Errorf("%s: unauthorized access", errSlur)
	ErrExerciseLogNotFound        = fmt.Errorf("%s: ExerciseLog not found", errSlur)
	ErrCannotExceedNumSets        = fmt.Errorf("%s: cannot exceed num sets for ExerciseLog", errSlur)
	ErrWorkoutLogAlreadyCompleted = fmt.Errorf("%s: WorkoutLog already completed", errSlur)
)
