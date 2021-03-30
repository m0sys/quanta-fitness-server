package state

import (
	ustore "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	wstore "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
)

type ServerState struct {
	UserStore    ustore.UserStore
	WorkoutStore wstore.WorkoutStore
}

func NewServerState() ServerState {
	return ServerState{
		UserStore:    ustore.NewMockUserStore(),
		WorkoutStore: wstore.NewMockWorkoutStore(),
	}
}
