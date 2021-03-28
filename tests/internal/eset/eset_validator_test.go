package esettest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	es "github.com/mhd53/quanta-fitness-server/internal/eset"
)

func TestValidateAddEsetToExerciseWhenNegativeActualRC(t *testing.T) {
	testValidator := es.NewEsetValidator()

	err := testValidator.ValidateAddEsetToExercise(-10, 120.0, 123.3)

	assert.NotNil(t, err)
}

func TestValidateAddEsetToExerciseWhenNegativeDur(t *testing.T) {
	testValidator := es.NewEsetValidator()

	err := testValidator.ValidateAddEsetToExercise(10, -120.0, 123.3)

	assert.NotNil(t, err)
}

func TestValidateAddEsetToExerciseWhenNegativeRestDur(t *testing.T) {
	testValidator := es.NewEsetValidator()

	err := testValidator.ValidateAddEsetToExercise(10, 120.0, -123.3)

	assert.NotNil(t, err)
}

func TestValidateAddEsetToExerciseSuccess(t *testing.T) {
	testValidator := es.NewEsetValidator()

	err := testValidator.ValidateAddEsetToExercise(10, 120.0, 123.3)

	assert.Nil(t, err)
}
