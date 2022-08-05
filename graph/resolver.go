package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/m0sys/quanta-fitness-server/graph/generated"

	"github.com/m0sys/quanta-fitness-server/api/auth"
	"github.com/m0sys/quanta-fitness-server/api/eset"
	"github.com/m0sys/quanta-fitness-server/api/exercise"
	"github.com/m0sys/quanta-fitness-server/api/state"
	"github.com/m0sys/quanta-fitness-server/api/workout"
)

type Resolver struct {
	AuthServer     auth.ServerAuth
	WorkoutServer  workout.WorkoutServer
	ExerciseServer exercise.ExerciseServer
	EsetServer     eset.EsetServer
}

func NewResolver() generated.Config {
	s := state.NewServerState()
	return generated.Config{Resolvers: &Resolver{
		AuthServer:     auth.NewServerAuth(s.UserStore),
		WorkoutServer:  workout.NewWorkoutServer(s.UserStore, s.WorkoutStore),
		ExerciseServer: exercise.NewExerciseServer(s.UserStore, s.WorkoutStore, s.ExerciseStore),
		EsetServer:     eset.NewEsetServer(s.UserStore, s.EsetStore, s.ExerciseStore, s.WorkoutStore),
	}}
}
