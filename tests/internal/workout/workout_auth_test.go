package workouttests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	ws "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	"github.com/mhd53/quanta-fitness-server/internal/workout"
	"github.com/mhd53/quanta-fitness-server/tests/internal/auth"
)

func skipTest(t *testing.T) {
	t.Skip("Focusing on Authentication.")
}

func TestAuthorizeCreateWorkoutWhenUnauthorized(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()
	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeCreateWorkout("hello")

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeCreateWorkoutWhenAuthorized(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	user := authtests.CreateValidAuthBaseUser()
	mockUS.Save(user)
	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeCreateWorkout("robin")

	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAuthorizeAccessWorkoutWhenUnauthorized(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	created, _ := mockWS.Save(CreateWorkout())
	assert.NotEmpty(t, created)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeAccessWorkout("hello", 0)

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
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	created, _ := mockWS.Save(
		entity.BaseWorkout{
			Username: "bobin",
			Title:    MOCK_TITLE,
		},
	)
	assert.NotEmpty(t, created)

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeAccessWorkout("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeAccessWorkoutWhenWorkoutNotFound(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	created, _ := mockWS.Save(CreateWorkout())
	assert.NotEmpty(t, created)

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeAccessWorkout("robin", 1)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeAccessWorkoutSuccess(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	created, _ := mockWS.Save(CreateWorkout())
	assert.NotEmpty(t, created)

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeAccessWorkout("robin", 0)

	assert.Nil(t, err)
	assert.True(t, ok)
}
