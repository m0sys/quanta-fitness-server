// Facade for all persistant tracking of data for the FitnessTracker
package tracker

import (
	a "github.com/mhd53/quanta-fitness-server/athlete"
	w "github.com/mhd53/quanta-fitness-server/workout"
)

type TrackerGateway interface {
	WorkoutGateway
	ExerciseGateway
	SetGateway
	AthleteGateway
}

type WorkoutGateway interface {
	SaveWorkout(w.Workout) error
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
	FindAtheleteByUname(uname string) (a.Athlete, bool)
	SaveAthlete()
	GetAthlete()
	DeleteAthlete()
	UpdatedAthlete()
}
