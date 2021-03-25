package workouttests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
	"github.com/mhd53/quanta-fitness-server/internal/workout"
	"github.com/mhd53/quanta-fitness-server/tests/internal/auth"
)

func TestCreateWorkoutWhenUnauthorized(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := new(MockStore)
	mockUS.On("FindUserByUsername").Return(entity.User{}, false, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	workout, err := testService.CreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus.", "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, workout)
}

func TestCreateWorkoutWhenTitleIsInvalid(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := new(MockStore)

	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	workout, err := testService.CreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus..", "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Title must be less than 76 characters!", err.Error())
	assert.Empty(t, workout)
}

func TestCreateWorkoutSuccess(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := new(MockStore)

	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)

	mockWorkout := CreateValidMockWorkout(id)
	mockWS.On("Save").Return(mockWorkout, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	workoutDS, err := testService.CreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus.", "robin")

	assert.Nil(t, err)
	assert.NotEmpty(t, workoutDS)
	assert.Equal(t, workoutDS.Username, mockWorkout.Username)
	assert.Equal(t, workoutDS.Title, mockWorkout.Title)
}
