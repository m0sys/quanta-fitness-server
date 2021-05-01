package planning

import "fmt"

const (
	errSlur = "planning"
)

var (
	errInternal                 = fmt.Errorf("%s: internal error", errSlur)
	ErrIdentialTitle            = fmt.Errorf("%s: WorkoutPlan with identical title already exists", errSlur)
	ErrWorkoutPlanAlreadyExists = fmt.Errorf("%s: WorkoutPlan already exists", errSlur)
	ErrUnauthorizedAccess       = fmt.Errorf("%s: unauthorized access", errSlur)
	ErrWorkoutPlanNotFound      = fmt.Errorf("%s: WorkoutPlan not found", errSlur)
	ErrIdentialName             = fmt.Errorf("%s: Exercise with identical name already in WorkoutPlan", errSlur)
	ErrExerciseAlreadyExists    = fmt.Errorf("%s: Exercise already Exists", errSlur)
)
