package workouttests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	// ws "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/workout"
)

func TestValidateCreateWorkoutWhenTitleLengthMaxed(t *testing.T) {
	// mockWS := ws.NewMockWorkoutStore()
	testValidator := workout.NewWorkoutValidator()

	err := testValidator.ValidateCreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus..")

	assert.NotNil(t, err)
	assert.Equal(t, "Title must be less than 76 characters!", err.Error())
}

func TestValidateCreateWorkoutWhenTitleIsValid(t *testing.T) {
	// mockWS := ws.NewMockWorkoutStore()
	testValidator := workout.NewWorkoutValidator()

	err := testValidator.ValidateCreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus.")

	assert.Nil(t, err)
}

func TestValidateUpdateWorkoutWhenTitleLengthMaxed(t *testing.T) {
	// mockWS := ws.NewMockWorkoutStore()
	testValidator := workout.NewWorkoutValidator()

	mockWorkout := CreateInvalidMockBaseWorkout()
	err := testValidator.ValidateUpdateWorkout(mockWorkout)

	assert.NotNil(t, err)
	assert.Equal(t, "Title must be less than 76 characters!", err.Error())
}

func TestValidateUpdateWorkoutWhenTitleIsValid(t *testing.T) {
	// mockWS := ws.NewMockWorkoutStore()
	testValidator := workout.NewWorkoutValidator()

	mockWorkout := CreateValidMockBaseWorkout()
	err := testValidator.ValidateUpdateWorkout(mockWorkout)

	assert.Nil(t, err)
}
