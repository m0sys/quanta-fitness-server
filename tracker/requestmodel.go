package tracker

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
