package planning

import "github.com/m0sys/quanta-fitness-server/units"

type WorkoutPlanRes struct {
	ID        string
	Title     string
	AthleteID string
}

type ExerciseRes struct {
	ID            string
	WorkoutPlanID string
	AthleteID     string
	Name          string
	TargetRep     int
	NumSets       int
	Weight        units.Kilogram
	RestDur       units.Second
}
