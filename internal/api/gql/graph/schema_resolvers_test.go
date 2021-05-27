package graph

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/mhd53/quanta-fitness-server/internal/api/gql/graph"
	"github.com/mhd53/quanta-fitness-server/internal/api/gql/graph/generated"
	"github.com/stretchr/testify/assert"
)

func TestCreateWorkoutPlan(t *testing.T) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.NewResolver()))
	// authSrv := auth.Middleware()(srv)
	// c := client.New(authSrv)
	c := client.New(srv)

	t.Run("Create WorkoutPlan", func(t *testing.T) {
		var res struct {
			WorkoutPlan struct {
				ID    string
				Title string
			}
		}

		c.MustPost(`mutation {
workoutPlan: createWorkoutPlan(input: {title: "Chest Gains"}) {
				    id
					title
						  
			}
			
		}`, &res)
		assert.NotEmpty(t, res.WorkoutPlan)
		assert.Equal(t, "Chest Gains", res.WorkoutPlan.Title)
	})
}
