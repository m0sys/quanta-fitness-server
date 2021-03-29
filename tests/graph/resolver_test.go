package resolver

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/graph"
	"github.com/mhd53/quanta-fitness-server/graph/generated"
)

func TestAuthFlow(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(graph.NewResolver())))

	t.Run("Register user", func(t *testing.T) {
		var resp struct {
			Register struct{ Token string }
		}

		c.MustPost(`mutation {
register:register(input: {username: "hero", email: "hero@gmail.com", password: "nero", confirm: "nero"}) {
				    token
			}
			
		}`, &resp)
		assert.NotEmpty(t, resp.Register.Token)
	})
}
