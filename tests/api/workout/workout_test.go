package workouttest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/m0sys/quanta-fitness-server/api/auth"
	wapi "github.com/m0sys/quanta-fitness-server/api/workout"
	ustore "github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
	wstore "github.com/m0sys/quanta-fitness-server/internal/datastore/workoutstore"
)

func TestCreateWorkoutWhenInvalidToken(t *testing.T) {
	skipTest(t)
	mockUS := ustore.NewMockUserStore()
	mockWS := wstore.NewMockWorkoutStore()
	server := wapi.NewWorkoutServer(mockUS, mockWS)

	created, err := server.CreateWorkout("Chest Day", "")
	assert.NotNil(t, err)
	assert.Empty(t, created)
}

func TestCreateWorkoutWhenValidToken(t *testing.T) {
	skipTest(t)
	mockUS := ustore.NewMockUserStore()
	mockWS := wstore.NewMockWorkoutStore()
	authServer := auth.NewServerAuth(mockUS)
	server := wapi.NewWorkoutServer(mockUS, mockWS)

	token, _ := authServer.RegisterNewUser("robin", "robin@gmail.com", "password", "password")

	created, err := server.CreateWorkout("Chest Day", token)
	assert.Nil(t, err)
	assert.NotEmpty(t, created)
	assert.Equal(t, "robin", created.Username)
	assert.Equal(t, "Chest Day", created.Title)
}

// Util funcs.

func skipTest(t *testing.T) {
	t.Skip("Implement AuthorizedReadAccess first!.")
}
