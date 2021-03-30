package resolver

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/api/auth"
	"github.com/mhd53/quanta-fitness-server/graph"
	"github.com/mhd53/quanta-fitness-server/graph/generated"
)

func TestAuthFlow(t *testing.T) {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.NewResolver()))
	authSrv := auth.Middleware()(srv)
	c := client.New(authSrv)

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

	t.Run("Login user with username", func(t *testing.T) {
		var resp struct {
			Login struct{ Token string }
		}

		c.MustPost(`mutation {
		login:login(input: {username: "hero", password: "nero"}) {
		token
		}}`, &resp)
		assert.NotEmpty(t, resp.Login.Token)
	})

	t.Run("Login user with email", func(t *testing.T) {
		var resp struct {
			Login struct{ Token string }
		}

		c.MustPost(`mutation {
		login:login(input: {email: "hero@gmail.com", password: "nero"}) {
		token
		}}`, &resp)
		assert.NotEmpty(t, resp.Login.Token)
	})
}

func TestWorkoutQueries(t *testing.T) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.NewResolver()))
	authSrv := auth.Middleware()(srv)
	c := client.New(authSrv)
	var token string

	// Create Workout Mutation Tests.
	t.Run("Create Workout: When not unauthenticated", func(t *testing.T) {
		var resp struct{}

		err := c.Post(`mutation {
				createWorkout(input: {title: "Chest Day 2"}) {
				    id
					title
				    createdAt
					updatedAt
								  
			}
			
		}`, &resp)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["createWorkout"]}]`)
	})

	t.Run("Create Workout: When authenticated", func(t *testing.T) {
		var resp struct {
			Register struct{ Token string }
		}

		c.MustPost(`mutation {
register:register(input: {username: "hero", email: "hero@gmail.com", password: "nero", confirm: "nero"}) {
				    token
			}
			
		}`, &resp)
		assert.NotEmpty(t, resp.Register.Token)

		token = resp.Register.Token

		var resp2 struct {
			Workout struct {
				Title string
			}
		}

		c.MustPost(`mutation {
workout:createWorkout(input: {title: "Chest Day 2"}) {
					 title
			}
			
		}`, &resp2, client.AddHeader("Authorization", token))
		assert.NotEmpty(t, resp2.Workout)
		assert.Equal(t, "Chest Day 2", resp2.Workout.Title)

	})

	// Update Workout Mutation Tests.
	t.Run("Update Workout: When not unauthenticated", func(t *testing.T) {
		var resp struct{}

		err := c.Post(`mutation {
updateWorkout(input: {id: "0", title: "Pull ups"})}`, &resp)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["updateWorkout"]}]`)
	})

	t.Run("Update Workout: When authenticated", func(t *testing.T) {
		var resp2 struct {
			Workout bool
		}

		c.MustPost(`mutation {
workout:updateWorkout(input: {id: "0", title: "Pull ups"})}`, &resp2, client.AddHeader("Authorization", token))
		assert.True(t, resp2.Workout)

	})

	// Get Workout Query Tests.
	t.Run("Get Workout: When not unauthenticated", func(t *testing.T) {
		var resp struct{}

		err := c.Post(`{ workout(id: 0) {
		id
		title
		}}`, &resp)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["workout"]}]`)
	})

	t.Run("Get Workout: When authenticated", func(t *testing.T) {
		var resp struct {
			Workout struct {
				Id    string
				Title string
			}
		}

		c.MustPost(`{ workout(id: 0) {
		id
		title
		}}`, &resp, client.AddHeader("Authorization", token))
		assert.NotEmpty(t, resp.Workout)
		assert.Equal(t, "0", resp.Workout.Id)
		assert.Equal(t, "Pull ups", resp.Workout.Title)

	})

	// Delete Workout Query Tests.
	t.Run("Delete Workout: When not unauthenticated", func(t *testing.T) {
		var resp struct{}

		err := c.Post(`mutation { deleteWorkout(id: 0) }`, &resp)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["deleteWorkout"]}]`)
	})

	t.Run("Delete Workout: When authenticated", func(t *testing.T) {
		var resp struct {
			Workout bool
		}

		c.MustPost(`mutation { workout:deleteWorkout(id: 0) }`, &resp, client.AddHeader("Authorization", token))
		assert.True(t, resp.Workout)

	})

	// Get Workouts Query Tests.
	t.Run("Get Workouts: When not unauthenticated", func(t *testing.T) {
		var resp struct{}

		err := c.Post(`{ workouts(username: "hero") {
		id
		title
		}}`, &resp)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["workouts"]}]`)
	})

	t.Run("Get Workouts: When authenticated", func(t *testing.T) {
		var resp2 struct {
			Workout struct {
				Title string
			}
		}

		c.MustPost(`mutation {
workout:createWorkout(input: {title: "Chest Day 2"}) {
					 title
			}
			
		}`, &resp2, client.AddHeader("Authorization", token))

		c.MustPost(`mutation {
workout:createWorkout(input: {title: "Chest Day 2"}) {
					 title
			}
			
		}`, &resp2, client.AddHeader("Authorization", token))

		c.MustPost(`mutation {
workout:createWorkout(input: {title: "Chest Day 2"}) {
					 title
			}
			
		}`, &resp2, client.AddHeader("Authorization", token))

		var resp struct {
			Workouts []struct {
				Id    string
				Title string
			}
		}

		c.MustPost(`{ workouts(username: "hero") {
		id
		title
		}}`, &resp, client.AddHeader("Authorization", token))
		assert.NotEmpty(t, resp.Workouts)
		assert.Equal(t, 3, len(resp.Workouts))

	})
}
