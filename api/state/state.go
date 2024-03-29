package state

import (
	esstore "github.com/m0sys/quanta-fitness-server/internal/datastore/esetstore"
	estore "github.com/m0sys/quanta-fitness-server/internal/datastore/exercisestore"
	ustore "github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
	wstore "github.com/m0sys/quanta-fitness-server/internal/datastore/workoutstore"
)

type ServerState struct {
	UserStore     ustore.UserStore
	WorkoutStore  wstore.WorkoutStore
	ExerciseStore estore.ExerciseStore
	EsetStore     esstore.EsetStore
}

func NewServerState() ServerState {
	return ServerState{
		UserStore:     ustore.NewMockUserStore(),
		WorkoutStore:  wstore.NewMockWorkoutStore(),
		ExerciseStore: estore.NewMockExerciseStore(),
		EsetStore:     esstore.NewMockEsetStore(),
	}
}
