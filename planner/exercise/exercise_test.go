package exercise

import (
	"testing"

	"github.com/m0sys/quanta-fitness-server/internal/random"
	"github.com/m0sys/quanta-fitness-server/units"
	"github.com/stretchr/testify/require"
)

func TestNewExercise(t *testing.T) {
	t.Run("When name is more than 75 chars", func(t *testing.T) {
		metrics, err := NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		require.NoError(t, err)
		exercise, err := NewExercise("1", "1", random.String(76), metrics, 0)

		require.Error(t, err)
		require.Equal(t, ErrInvalidName.Error(), err.Error())
		require.Empty(t, exercise)
	})

	t.Run("When success", func(t *testing.T) {
		metrics, err := NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		require.NoError(t, err)

		exercise, err := NewExercise("1", "1", random.String(75), metrics, 0)
		require.NoError(t, err)
		require.NotEmpty(t, exercise)
	})
}

func TestNewMetrics(t *testing.T) {
	t.Run("When negative target rep", func(t *testing.T) {
		metrics, err := NewMetrics(
			-1,
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)

		require.Error(t, err)
		require.Equal(t, ErrInvalidTargetRep.Error(), err.Error())
		require.Empty(t, metrics)
	})

	t.Run("When negative num sets", func(t *testing.T) {
		metrics, err := NewMetrics(
			random.RepCount(),
			-1,
			random.Weight(),
			random.RestTime(),
		)

		require.Error(t, err)
		require.Equal(t, ErrInvalidNumSets.Error(), err.Error())
		require.Empty(t, metrics)
	})

	t.Run("When negative weight", func(t *testing.T) {
		metrics, err := NewMetrics(
			random.RepCount(),
			random.NumSets(),
			-1,
			random.RestTime(),
		)

		require.Error(t, err)
		require.Equal(t, ErrInvalidWeight.Error(), err.Error())
		require.Empty(t, metrics)
	})

	t.Run("When negative rest duration", func(t *testing.T) {
		metrics, err := NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			-1,
		)

		require.Error(t, err)
		require.Equal(t, ErrInvalidRestDur.Error(), err.Error())
		require.Empty(t, metrics)
	})

	t.Run("When success", func(t *testing.T) {
		metrics, err := NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)

		require.NoError(t, err)
		require.NotEmpty(t, metrics)
	})
}

func TestEditExercise(t *testing.T) {
	metrics, err := NewMetrics(
		random.RepCount(),
		random.NumSets(),
		random.Weight(),
		random.RestTime(),
	)
	require.NoError(t, err)

	exercise, err := NewExercise("1", "1", random.String(75), metrics, 0)

	t.Run("When long name", func(t *testing.T) {
		err = exercise.EditName(random.String(76))
		require.Error(t, err)
		require.Equal(t, ErrInvalidName.Error(), err.Error())
	})

	t.Run("When name good", func(t *testing.T) {
		gen := random.String(75)
		err = exercise.EditName(gen)
		require.NoError(t, err)
		require.Equal(t, gen, exercise.Name())
	})

	t.Run("When negative target rep", func(t *testing.T) {
		err = exercise.EditTargetRep(-1)
		require.Error(t, err)
		require.Equal(t, ErrInvalidTargetRep.Error(), err.Error())
	})

	t.Run("When target rep good", func(t *testing.T) {
		gen := random.RepCount()
		err = exercise.EditTargetRep(gen)
		require.NoError(t, err)

		metrics := exercise.Metrics()
		require.Equal(t, gen, metrics.TargetRep())
	})

	t.Run("When negative num sets", func(t *testing.T) {
		err = exercise.EditNumSets(-1)
		require.Error(t, err)
		require.Equal(t, ErrInvalidNumSets.Error(), err.Error())
	})

	t.Run("When num sets good", func(t *testing.T) {
		gen := random.NumSets()
		err = exercise.EditNumSets(gen)
		require.NoError(t, err)

		metrics := exercise.Metrics()
		require.Equal(t, gen, metrics.NumSets())
	})

	t.Run("When negative weight", func(t *testing.T) {
		err = exercise.EditWeight(-1)
		require.Error(t, err)
		require.Equal(t, ErrInvalidWeight.Error(), err.Error())
	})

	t.Run("When weight good", func(t *testing.T) {
		gen := random.Weight()
		err = exercise.EditWeight(gen)
		require.NoError(t, err)

		metrics := exercise.Metrics()
		// Very fragile test...!
		require.Equal(t, units.Kilogram(roundToTwoDecimalPlaces(gen)), metrics.Weight())
	})

	t.Run("When negative rest duration", func(t *testing.T) {
		err = exercise.EditRestDur(-1)
		require.Error(t, err)
		require.Equal(t, ErrInvalidRestDur.Error(), err.Error())
	})

	t.Run("When rest duration good", func(t *testing.T) {
		gen := random.RestTime()
		err = exercise.EditRestDur(gen)
		require.NoError(t, err)

		metrics := exercise.Metrics()
		// Very fragile test...!
		require.Equal(t, units.Second(roundToTwoDecimalPlaces(gen)), metrics.RestDur())
	})
}
