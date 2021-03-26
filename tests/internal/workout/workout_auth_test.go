package workouttests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	ws "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	"github.com/mhd53/quanta-fitness-server/internal/workout"
	"github.com/mhd53/quanta-fitness-server/tests/internal/auth"
)

func skipTest(t *testing.T) {
	t.Skip("Focusing on Authentication.")
}

func TestAuthorizeCreateWorkoutWhenUnauthorized(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()
	mockUS.On("FindUserByUsername").Return(entity.User{}, false, nil)
	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeCreateWorkout("hello")

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeCreateWorkoutWhenAuthorized(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()

	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)
	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeCreateWorkout("robin")

	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAuthorizeAccessWorkoutWhenUnauthorized(t *testing.T) {
	skipTest(t)
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()
	mockUS.On("FindUserByUsername").Return(entity.User{}, false, nil)

	created, _ := mockWS.Save(CreateWorkout())
	assert.NotEmpty(t, created)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeAccessWorkout("hello", 1)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func CreateWorkout() entity.BaseWorkout {
	return entity.BaseWorkout{
		Username: "robin",
		Title:    MOCK_TITLE,
	}
}

func TestAuthorizeAccessWorkoutWhenUserNotOwnWorkout(t *testing.T) {
	skipTest(t)
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()
	created, _ := mockWS.Save(CreateWorkout())
	assert.NotEmpty(t, created)

	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeAccessWorkout("bobin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeAccessWorkoutWhenWorkoutNotFound(t *testing.T) {
	skipTest(t)
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()

	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeAccessWorkout("bobin", 1)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeAccessWorkoutSuccess(t *testing.T) {
	skipTest(t)
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()
	created, _ := mockWS.Save(CreateWorkout())
	assert.NotEmpty(t, created)

	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeAccessWorkout("robin", 0)

	assert.Nil(t, err)
	assert.True(t, ok)
}
