package workouttests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/workout"
	// "github.com/mhd53/quanta-fitness-server/internal/entity"
)

func TestValidateCreateWorkoutWhenTitleLengthMaxed(t *testing.T) {
	mockWS := new(MockStore)
	testValidator := workout.NewWorkoutValidator(mockWS)

	err := testValidator.ValidateCreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus..")

	assert.NotNil(t, err)
	assert.Equal(t, "Title must be less than 76 characters!", err.Error())
}

func TestValidateCreateWorkoutWhenTitleIsValid(t *testing.T) {
	mockWS := new(MockStore)
	testValidator := workout.NewWorkoutValidator(mockWS)

	err := testValidator.ValidateCreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus.")

	assert.Nil(t, err)
}
