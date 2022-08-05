package resolver

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"

	"github.com/m0sys/quanta-fitness-server/api/auth"
	"github.com/m0sys/quanta-fitness-server/graph"
	"github.com/m0sys/quanta-fitness-server/graph/generated"
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

func TestExerciseQueries(t *testing.T) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.NewResolver()))
	authSrv := auth.Middleware()(srv)
	c := client.New(authSrv)
	var token string

	var resp struct {
		Register struct{ Token string }
	}

	// First create a new user to get a valid token.
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

	// Second create a workout.
	c.MustPost(`mutation {
	workout:createWorkout(input: {title: "Chest Day 2"}) {
					 title
			}
	}`, &resp2, client.AddHeader("Authorization", token))
	assert.NotEmpty(t, resp2.Workout)
	assert.Equal(t, "Chest Day 2", resp2.Workout.Title)

	// Add Exercise To Workout Query Tests.
	t.Run("Add Exercise To Workout: When not unauthenticated", func(t *testing.T) {
		var resp3 struct{}

		err := c.Post(`mutation {
			exercise: addExerciseToWorkout(input: {wid: 0, name: "Leg Day"}) {
				    id
					name
				    wid
			}
			
		}`, &resp3)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["exercise"]}]`)
	})

	t.Run("Add Exercise To Workout: When authenticated", func(t *testing.T) {
		var resp3 struct {
			Exercise struct {
				Id   string
				Name string
				Wid  string
			}
		}

		c.MustPost(`mutation {
			exercise: addExerciseToWorkout(input: {wid: 0, name: "Leg Day"}) {
				    id
					name
				    wid
			}
			
		}`, &resp3, client.AddHeader("Authorization", token))

		assert.NotNil(t, resp3.Exercise)
		assert.Equal(t, "0", resp3.Exercise.Id)
		assert.Equal(t, "0", resp3.Exercise.Wid)
		assert.Equal(t, "Leg Day", resp3.Exercise.Name)

	})

	// Update Exercise Query Tests.
	t.Run("Update Exercise: When not unauthenticated", func(t *testing.T) {
		var resp3 struct{}

		err := c.Post(`mutation {
			updated: updateExercise(input: {id: 0, name: "Flat Bench Press", weight: 192.0, targetRep: 5, restTime: 120.0, numSets: 3}) 
		}`, &resp3)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["updated"]}]`)
	})

	t.Run("Update Exercise: When authenticated", func(t *testing.T) {
		var resp3 struct {
			Updated bool
		}

		c.MustPost(`mutation {
updated: updateExercise(input: {id: 0, name: "Flat Bench Press", weight: 192.0, targetRep: 5, restTime: 120.0, numSets: 3}) 
		}`, &resp3, client.AddHeader("Authorization", token))

		assert.True(t, resp3.Updated)
	})

	// Get Exercise Query Tests.
	t.Run("Get Exercise: When not unauthenticated", func(t *testing.T) {
		var resp3 struct{}

		err := c.Post(`{exercise(id: 0) {
			  id
			  wid
			  name
			  weight
		}}`, &resp3)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["exercise"]}]`)
	})

	t.Run("Get Exercise: When authenticated", func(t *testing.T) {
		var resp3 struct {
			Exercise struct {
				Id     string
				Wid    string
				Name   string
				Weight float64
			}
		}

		c.MustPost(`{exercise(id: 0) {
			  id
			  wid
			  name
			  weight
		}}`, &resp3, client.AddHeader("Authorization", token))

		assert.NotEmpty(t, resp3.Exercise)
		assert.Equal(t, "0", resp3.Exercise.Id)
		assert.Equal(t, "0", resp3.Exercise.Wid)
		assert.Equal(t, "Flat Bench Press", resp3.Exercise.Name)
		assert.Equal(t, 192.0, resp3.Exercise.Weight)
	})

	// Get Exercises For Workout Query Tests.
	t.Run("Get Exercises For Workout: When not unauthenticated", func(t *testing.T) {
		var resp3 struct{}

		err := c.Post(`{exercises(wid: 0) {
			  id
			  wid
			  name
				  
		}}`, &resp3)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["exercises"]}]`)
	})

	t.Run("Get Exercises For Workout: When authenticated", func(t *testing.T) {
		var resp3 struct {
			Exercises []struct {
				Id   string
				Wid  string
				Name string
			}
		}

		c.MustPost(`{exercises(wid: 0) {
			  id
			  wid
			  name
				  
		}}`, &resp3, client.AddHeader("Authorization", token))

		assert.NotEmpty(t, resp3.Exercises)
		assert.Equal(t, 1, len(resp3.Exercises))
		assert.Equal(t, "0", resp3.Exercises[0].Id)
		assert.Equal(t, "0", resp3.Exercises[0].Wid)
		assert.Equal(t, "Flat Bench Press", resp3.Exercises[0].Name)
	})

	// Delete Exercise Query Tests.
	t.Run("Delete Exercise: When not unauthenticated", func(t *testing.T) {
		var resp3 struct{}

		err := c.Post(`mutation {
				deleted:deleteExercise(id: 0)
		}`, &resp3)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["deleted"]}]`)
	})

	t.Run("Delete Exercise: When authenticated", func(t *testing.T) {
		var resp3 struct {
			Deleted bool
		}

		c.MustPost(`mutation {
				deleted:deleteExercise(id: 0)
		}`, &resp3, client.AddHeader("Authorization", token))

		assert.True(t, resp3.Deleted)
	})

}

func TestEsetQueries(t *testing.T) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.NewResolver()))
	authSrv := auth.Middleware()(srv)
	c := client.New(authSrv)
	var token string

	var resp struct {
		Register struct{ Token string }
	}

	// First create a new user to get a valid token.
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

	// Second create a workout.
	c.MustPost(`mutation {
	workout:createWorkout(input: {title: "Chest Day 2"}) {
					 title
			}
	}`, &resp2, client.AddHeader("Authorization", token))
	assert.NotEmpty(t, resp2.Workout)
	assert.Equal(t, "Chest Day 2", resp2.Workout.Title)

	// Third create an exercise.
	var resp3 struct {
		Exercise struct {
			Id   string
			Name string
			Wid  string
		}
	}

	c.MustPost(`mutation {
			exercise: addExerciseToWorkout(input: {wid: 0, name: "Leg Day"}) {
				    id
					name
				    wid
			}
	}`, &resp3, client.AddHeader("Authorization", token))
	assert.NotNil(t, resp3.Exercise)
	assert.Equal(t, "0", resp3.Exercise.Id)
	assert.Equal(t, "0", resp3.Exercise.Wid)
	assert.Equal(t, "Leg Day", resp3.Exercise.Name)

	// Add Eset To Exercise Query Tests.
	t.Run("Add Eset To Exercise: When not unauthenticated", func(t *testing.T) {
		var resp4 struct{}

		err := c.Post(`mutation {
			added:addEsetToExercise(input: {eid: 0, actualRepCount: 7, duration: 120.0, restTimeDuration: 120.0}) {
		    id
			eid
		    actualRepCount
			}
		}`, &resp4)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["added"]}]`)
	})

	t.Run("Add Eset To Exercise: When authenticated", func(t *testing.T) {
		var resp4 struct {
			Added struct {
				Id             string
				Eid            string
				ActualRepCount int
			}
		}

		c.MustPost(`mutation {
			added:addEsetToExercise(input: {eid: 0, actualRepCount: 7, duration: 120.0, restTimeDuration: 120.0}) {
		    id
			eid
		    actualRepCount
			}
		}`, &resp4, client.AddHeader("Authorization", token))

		assert.NotNil(t, resp4.Added)
		assert.Equal(t, "0", resp4.Added.Id)
		assert.Equal(t, "0", resp4.Added.Eid)
		assert.Equal(t, 7, resp4.Added.ActualRepCount)

	})

	// Update Eset Query Tests.
	t.Run("Update Eset: When not unauthenticated", func(t *testing.T) {
		var resp4 struct{}

		err := c.Post(`mutation {
			updated: updateEset(input: {id: 0, actualRepCount: 10, duration: 121.0, restTimeDuration: 125.0})
		}`, &resp4)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["updated"]}]`)
	})

	t.Run("Update Eset: When authenticated", func(t *testing.T) {
		var resp4 struct {
			Updated bool
		}

		c.MustPost(`mutation {
			updated: updateEset(input: {id: 0, actualRepCount: 10, duration: 121.0, restTimeDuration: 125.0})
		}`, &resp4, client.AddHeader("Authorization", token))

		assert.True(t, resp4.Updated)
	})

	// Get Eset Query Tests.
	t.Run("Get Eset: When not unauthenticated", func(t *testing.T) {
		var resp4 struct{}

		err := c.Post(`{eset(id: 0) {
			  id
			  eid
			  actualRepCount
		}}`, &resp4)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["eset"]}]`)
	})

	t.Run("Get Eset: When authenticated", func(t *testing.T) {
		var resp4 struct {
			Eset struct {
				Id             string
				Eid            string
				ActualRepCount int
			}
		}

		c.MustPost(`{eset(id: 0) {
			  id
			  eid
			  actualRepCount
		}}`, &resp4, client.AddHeader("Authorization", token))

		assert.NotNil(t, resp4.Eset)
		assert.Equal(t, "0", resp4.Eset.Id)
		assert.Equal(t, "0", resp4.Eset.Eid)
		assert.Equal(t, 10, resp4.Eset.ActualRepCount)
	})

	// Get Eset For Exercise Query Tests.
	t.Run("Get Eset For Exercise: When not unauthenticated", func(t *testing.T) {
		var resp4 struct{}

		err := c.Post(`{esets(eid: 0) {
			  id
			  eid
			  actualRepCount
		}}`, &resp4)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["esets"]}]`)
	})

	t.Run("Get Eset For Exercise: When authenticated", func(t *testing.T) {
		var resp4 struct {
			Esets []struct {
				Id             string
				Eid            string
				ActualRepCount int
			}
		}

		c.MustPost(`{esets(eid: 0) {
			  id
			  eid
			  actualRepCount
		}}`, &resp4, client.AddHeader("Authorization", token))

		assert.NotEmpty(t, resp4.Esets)
		assert.Equal(t, 1, len(resp4.Esets))
		assert.Equal(t, "0", resp4.Esets[0].Id)
		assert.Equal(t, "0", resp4.Esets[0].Eid)
		assert.Equal(t, 10, resp4.Esets[0].ActualRepCount)
	})

	// Delete Eset Query Tests.
	t.Run("Delete Eset: When not unauthenticated", func(t *testing.T) {
		var resp4 struct{}

		err := c.Post(`mutation {
				deleted:deleteEset(id: 0)
		}`, &resp4)
		assert.NotEmpty(t, err)
		assert.EqualError(t, err, `[{"message":"Access Denied!","path":["deleted"]}]`)
	})

	t.Run("Delete Eset: When authenticated", func(t *testing.T) {
		var resp4 struct {
			Deleted bool
		}

		c.MustPost(`mutation {
				deleted:deleteEset(id: 0)
		}`, &resp4, client.AddHeader("Authorization", token))

		assert.True(t, resp4.Deleted)
	})

}
