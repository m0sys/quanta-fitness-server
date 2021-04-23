package tracker

// AddExerciseToWorkoutLogReq request model for the AddExerciseToWorkoutLog use case.
type AddExerciseToWorkoutLogReq struct {
	LogID     string
	Name      string
	Weight    float64
	TargetRep int
	RestTime  float64
}
