package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/mhd53/quanta-fitness-server/graph/generated"

	"github.com/mhd53/quanta-fitness-server/api/auth"
	"github.com/mhd53/quanta-fitness-server/api/state"
	"github.com/mhd53/quanta-fitness-server/api/workout"
)

type Resolver struct {
	AuthServer    auth.ServerAuth
	WorkoutServer workout.WorkoutServer
}

func NewResolver() generated.Config {
	serverState := state.NewServerState()
	return generated.Config{Resolvers: &Resolver{
		AuthServer:    auth.NewServerAuth(serverState.UserStore),
		WorkoutServer: workout.NewWorkoutServer(serverState.UserStore, serverState.WorkoutStore),
	}}
}
