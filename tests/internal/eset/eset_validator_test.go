package esettest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/m0sys/quanta-fitness-server/internal/entity"
	es "github.com/m0sys/quanta-fitness-server/internal/eset"
)

func TestValidateAddEsetToExerciseWhenNegativeActualRC(t *testing.T) {
	testValidator := es.NewEsetValidator()

	err := testValidator.ValidateAddEsetToExercise(-10, 120.0, 123.3)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Actual rep count must be positive or zero!", err.Error())
}

func TestValidateAddEsetToExerciseWhenNegativeDur(t *testing.T) {
	testValidator := es.NewEsetValidator()

	err := testValidator.ValidateAddEsetToExercise(10, -120.0, 123.3)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Duration must be positive or zero!", err.Error())
}

func TestValidateAddEsetToExerciseWhenNegativeRestDur(t *testing.T) {
	testValidator := es.NewEsetValidator()

	err := testValidator.ValidateAddEsetToExercise(10, 120.0, -123.3)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Rest duration must be positive or zero!", err.Error())
}

func TestValidateAddEsetToExerciseSuccess(t *testing.T) {
	testValidator := es.NewEsetValidator()

	err := testValidator.ValidateAddEsetToExercise(10, 120.0, 123.3)

	assert.Nil(t, err)
}

func TestValidateUpdateEsetWhenNegativeActualRC(t *testing.T) {
	testValidator := es.NewEsetValidator()

	updates := entity.EsetUpdate{
		SMetric: entity.SMetric{

			ActualRepCount:   -5,
			Duration:         120.0,
			RestTimeDuration: 123.3,
		},
	}

	err := testValidator.ValidateUpdateEset(updates)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Actual rep count must be positive or zero!", err.Error())
}

func TestValidateUpdateEsetWhenNegativeDuration(t *testing.T) {
	testValidator := es.NewEsetValidator()

	updates := entity.EsetUpdate{
		SMetric: entity.SMetric{

			ActualRepCount:   5,
			Duration:         -120.0,
			RestTimeDuration: 123.3,
		},
	}

	err := testValidator.ValidateUpdateEset(updates)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Duration must be positive or zero!", err.Error())
}

func TestValidateUpdateEsetWhenNegativeRestTimeDuration(t *testing.T) {
	testValidator := es.NewEsetValidator()

	updates := entity.EsetUpdate{
		SMetric: entity.SMetric{

			ActualRepCount:   5,
			Duration:         120.0,
			RestTimeDuration: -123.3,
		},
	}

	err := testValidator.ValidateUpdateEset(updates)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Rest duration must be positive or zero!", err.Error())
}

func TestValidateUpdateEsetSuccess(t *testing.T) {
	testValidator := es.NewEsetValidator()

	updates := entity.EsetUpdate{
		SMetric: entity.SMetric{

			ActualRepCount:   5,
			Duration:         120.0,
			RestTimeDuration: 123.3,
		},
	}

	err := testValidator.ValidateUpdateEset(updates)

	assert.Nil(t, err)
}
