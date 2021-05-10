package planning

import "github.com/mhd53/quanta-fitness-server/units"

type WorkoutPlanRes struct {
	ID    string
	Title string
}

type ExerciseRes struct {
	ID            string
	WorkoutPlanID string
	Name          string
	TargetRep     int
	NumSets       int
	Weight        units.Kilogram
	RestDur       units.Second
}
