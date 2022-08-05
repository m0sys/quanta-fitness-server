package exercisetest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/m0sys/quanta-fitness-server/internal/entity"
	e "github.com/m0sys/quanta-fitness-server/internal/exercise"
)

func TestValidateCreateExerciseWhenNameTooLong(t *testing.T) {
	testValidator := e.NewExerciseValidator()

	err := testValidator.ValidateCreateExercise(MOCK_INVALID_NAME)

	assert.NotNil(t, err)
	assert.Equal(t, "Name must be less than 38 characters!", err.Error())

}

func TestValidateCreateExerciseSuccess(t *testing.T) {
	testValidator := e.NewExerciseValidator()

	err := testValidator.ValidateCreateExercise(MOCK_VALID_NAME)

	assert.Nil(t, err)
}

func TestValidateUpdateExerciseWhenNameTooLong(t *testing.T) {
	testValidator := e.NewExerciseValidator()

	updates := createUpdates(MOCK_INVALID_NAME, 0, 0, 0, 0)
	err := testValidator.ValidateUpdateExercise(updates)

	assert.NotNil(t, err)
	assert.Equal(t, "Name must be less than 38 characters!", err.Error())
}

func createUpdates(name string, weight, restTime float32, targetRep, numSets int) entity.ExerciseUpdate {
	return entity.ExerciseUpdate{
		Name: name,
		Metrics: entity.Metrics{
			Weight:    weight,
			RestTime:  restTime,
			TargetRep: targetRep,
			NumSets:   numSets,
		},
	}
}

func TestValidateUpdateExerciseWhenNegativeWeight(t *testing.T) {
	testValidator := e.NewExerciseValidator()

	updates := createUpdates(MOCK_VALID_NAME, -10, 0, 0, 0)
	err := testValidator.ValidateUpdateExercise(updates)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Weight must be positive or zero!", err.Error())
}

func TestValidateUpdateExerciseWhenNegativeRestTime(t *testing.T) {
	testValidator := e.NewExerciseValidator()

	updates := createUpdates(MOCK_VALID_NAME, 0, -10, 0, 0)
	err := testValidator.ValidateUpdateExercise(updates)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Rest time must be positive or zero!", err.Error())
}

func TestValidateUpdateExerciseWhenNegativeTargetRep(t *testing.T) {
	testValidator := e.NewExerciseValidator()

	updates := createUpdates(MOCK_VALID_NAME, 0, 0, -10, 0)
	err := testValidator.ValidateUpdateExercise(updates)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Target rep must be positive or zero!", err.Error())
}

func TestValidateUpdateExerciseWhenNegativeNumSets(t *testing.T) {
	testValidator := e.NewExerciseValidator()

	updates := createUpdates(MOCK_VALID_NAME, 0, 0, 0, -10)
	err := testValidator.ValidateUpdateExercise(updates)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Number of sets must be positive or zero!", err.Error())
}
