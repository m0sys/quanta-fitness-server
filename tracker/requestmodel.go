package tracker

import "time"

// AddExerciseToWorkoutLogReq request model for the AddExerciseToWorkoutLog use case.
type AddExerciseToWorkoutLogReq struct {
	LogID     string
	Name      string
	Weight    float64
	TargetRep int
	RestTime  float64
}

// AddSetToExerciseReq request model for the AddSetToExercise use case.
type AddSetToExerciseReq struct {
	LogID          string
	ExerciseID     string
	ActualRepCount int
}

// EditWorkoutLogReq request model for the EditWorkoutLog use case.
type EditWorkoutLogReq struct {
	LogID string
	Title string
	Date  time.Time
}

// EditExerciseReq request model for the EditExercise use case.
type EditExerciseReq struct {
	ExerciseID string
	Name       string
	Weight     float64
	TargetRep  int
	RestTime   float64
}

// EditSetReq request model for the EditSet use case.
type EditSetReq struct {
	SetID          string
	ExerciseID     string
	ActualRepCount int
}
