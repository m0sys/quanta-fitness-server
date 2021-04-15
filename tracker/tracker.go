// Facade for all persistant tracking of data for the FitnessTracker
package tracker

type FitnessGateway interface {
	WorkoutGateway
	ExerciseGateway
	SetGateway
	AthleteGateway
}

type WorkoutGateway interface {
	SaveWorkout()
	GetWorkout()
	DeleteWorkout()
	UpdateWorkout()
}

type ExerciseGateway interface {
	SaveExercise()
	GetExercise()
	DeleteExercise()
	UpdateExercise()
}

type SetGateway interface {
	SaveSet()
	GetSet()
	DeleteSet()
	UpdateSet()
}

type AthleteGateway interface {
	SaveAthlete()
	GetAthlete()
	DeleteAthlete()
	UpdatedAthlete()
}
