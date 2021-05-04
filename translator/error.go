package translator

import "fmt"

const (
	errSlur = "translator"
)

var (
	errInternal                  = fmt.Errorf("%s: internal error", errSlur)
	ErrWorkoutPlanHasNoExercises = fmt.Errorf("%s: WorkoutPlan has no Exercises", errSlur)
)
