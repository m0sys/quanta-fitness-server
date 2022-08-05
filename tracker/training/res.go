package training

import (
	"time"

	"github.com/m0sys/quanta-fitness-server/units"
)

type ExerciseLogRes struct {
	ID           string
	WorkoutLogID string
	Name         string
	TargetRep    int
	NumSets      int
	Weight       units.Kilogram
	RestDur      units.Second
	Completed    bool
	Pos          int
}

type WorkoutLogRes struct {
	ID         string
	Title      string
	Date       time.Time
	CurrentPos int
	Completed  bool
}

type SetLogRes struct {
	ID             string
	ExerciseLogID  string
	ActualRepCount int
	Duration       units.Second
}
