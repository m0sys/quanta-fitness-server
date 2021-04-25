package tracker

import "time"

// WorkoutLogRes response model for the use cases that return a WorkoutLog.
type WorkoutLogRes struct {
	LogID     string
	Title     string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ExerciseRes response model for the use cases that return an Exercise.
type ExerciseRes struct {
	ExerciseID string
	Name       string
	Weight     float64
	TargetRep  int
	RestTime   float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// SetRes response model for the use cases that return a Set.
type SetRes struct {
	SetID          string
	ActualRepCount int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
