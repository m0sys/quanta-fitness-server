package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/mhd53/quanta-fitness-server/graph/generated"

	"github.com/mhd53/quanta-fitness-server/api/auth"
)

type Resolver struct {
	AuthServer auth.ServerAuth
}

func NewResolver() generated.Config {
	return generated.Config{Resolvers: &Resolver{
		AuthServer: auth.NewServerAuth(),
	}}
}
