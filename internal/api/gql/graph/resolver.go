package graph

import (
	"github.com/m0sys/quanta-fitness-server/internal/api/gql/graph/generated"
	p "github.com/m0sys/quanta-fitness-server/planner/planning"
	pa "github.com/m0sys/quanta-fitness-server/planner/planning/adapters"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	planning p.PlanningService
}

func NewResolver() generated.Config {
	repo := pa.NewInMemRepo()
	return generated.Config{Resolvers: &Resolver{
		planning: p.NewPlanningService(repo),
	}}
}
