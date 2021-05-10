package planning

import "fmt"

const (
	errSlug = "planning"
)

var (
	errInternal                   = fmt.Errorf("%s: internal error", errSlug)
	ErrIdentialTitle              = fmt.Errorf("%s: WorkoutPlan with identical title already exists", errSlug)
	ErrWorkoutPlanAlreadyExists   = fmt.Errorf("%s: WorkoutPlan already exists", errSlug)
	ErrUnauthorizedAccess         = fmt.Errorf("%s: unauthorized access", errSlug)
	ErrWorkoutPlanNotFound        = fmt.Errorf("%s: WorkoutPlan not found", errSlug)
	ErrIdentialName               = fmt.Errorf("%s: Exercise with identical name already in WorkoutPlan", errSlug)
	ErrExerciseAlreadyExists      = fmt.Errorf("%s: Exercise already Exists", errSlug)
	ErrExerciseNotFound           = fmt.Errorf("%s: Exercise not found", errSlug)
	ErrOrderOutOfRange            = fmt.Errorf("%s: Exercise order is out of range", errSlug)
	ErrAthleteIDCannotBeEmpty     = fieldCannotBeEmpty("AthleteID")
	ErrTitleCannotBeEmpty         = fieldCannotBeEmpty("Title")
	ErrWorkoutPlanIDCannotBeEmpty = fieldCannotBeEmpty("WorkoutPlanID")
	ErrNameCannotBeEmpty          = fieldCannotBeEmpty("Name")
	ErrTargetRepCannotBeEmpty     = fieldCannotBeEmpty("TargetRep")
	ErrNumSetsCannotBeEmpty       = fieldCannotBeEmpty("NumSets")
	ErrWeightCannotBeEmpty        = fieldCannotBeEmpty("Weight")
	ErrRestDurCannotBeEmpty       = fieldCannotBeEmpty("RestDur")
	ErrExerciseIDCannotBeEmpty    = fieldCannotBeEmpty("ExerciseID")
)

func fieldCannotBeEmpty(fieldName string) error {
	return fmt.Errorf("%s: %s cannot be empty", errSlug, fieldName)
}
