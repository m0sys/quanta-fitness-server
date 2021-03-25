package workouttests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
	"github.com/mhd53/quanta-fitness-server/internal/workout"
	"github.com/mhd53/quanta-fitness-server/tests/internal/auth"
)

func skipTest(t *testing.T) {
	t.Skip("Focusing on Authentication.")
}

func TestAuthorizeCreateWorkoutWhenUnauthorized(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := new(MockStore)
	mockUS.On("FindUserByUsername").Return(entity.User{}, false, nil)
	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeCreateWorkout("hello")

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeCreateWorkoutWhenAuthorized(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := new(MockStore)

	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)
	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)

	ok, err := testAuthorizer.AuthorizeCreateWorkout("robin")

	assert.Nil(t, err)
	assert.True(t, ok)
}
