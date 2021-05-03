package workoutlog

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/stretchr/testify/require"
)

func TestNewExercise(t *testing.T) {
	t.Run("When name is more than 75 chars", func(t *testing.T) {
		exercise, err := NewExercise(random.String(76), random.Weight(), random.RestTime(), random.RepCount(), 0)

		require.Error(t, err)
		require.Equal(t, "name must be less than 76 characters", err.Error())
		require.Empty(t, exercise)
	})

	t.Run("When negative weight", func(t *testing.T) {
		exercise, err := NewExercise(random.String(75), -1, random.RestTime(), random.RepCount(), 0)

		require.Error(t, err)
		require.Equal(t, "weight must be a positive number", err.Error())
		require.Empty(t, exercise)
	})

	t.Run("When negative rest time", func(t *testing.T) {
		exercise, err := NewExercise(random.String(75), random.Weight(), -1.0, random.RepCount(), 0)

		require.Error(t, err)
		require.Equal(t, "rest time must be a positive number", err.Error())
		require.Empty(t, exercise)
	})

	t.Run("When negative target rep", func(t *testing.T) {
		exercise, err := NewExercise(random.String(75), random.Weight(), random.RestTime(), -1, 0)

		require.Error(t, err)
		require.Equal(t, "target rep must be a positive number", err.Error())
		require.Empty(t, exercise)
	})

	t.Run("When success", func(t *testing.T) {
		exercise, err := NewExercise(random.String(75), random.Weight(), random.RestTime(), random.RepCount(), 0)

		require.NoError(t, err)
		require.NotEmpty(t, exercise)
	})
}

func TestAddSet(t *testing.T) {
	t.Run("When Set is new and then already logged", func(t *testing.T) {
		set, err := NewSet(random.RepCount())
		require.NoError(t, err)
		require.NotEmpty(t, set)

		exercise, err := NewExercise(random.String(75), random.Weight(), random.RestTime(), random.RepCount(), 0)
		require.NoError(t, err)
		require.NotEmpty(t, exercise)

		err = exercise.AddSet(set)
		require.NoError(t, err)
		require.Equal(t, 1, len(exercise.sets))
		require.Equal(t, set.SetID(), exercise.sets[0].SetID())

		err = exercise.AddSet(set)
		require.Error(t, err)
	})
}

func TestRemove(t *testing.T) {
	t.Run("When is not found and then found", func(t *testing.T) {
		set, err := NewSet(random.RepCount())
		require.NoError(t, err)
		require.NotEmpty(t, set)

		exercise, err := NewExercise(random.String(75), random.Weight(), random.RestTime(), random.RepCount(), 0)
		require.NoError(t, err)
		require.NotEmpty(t, exercise)

		err = exercise.RemoveSet(set)
		require.Error(t, err)
		require.Equal(t, 0, len(exercise.sets))

		err = exercise.AddSet(set)
		require.NoError(t, err)
		require.Equal(t, 1, len(exercise.sets))
		require.Equal(t, set.SetID(), exercise.sets[0].SetID())

		err = exercise.RemoveSet(set)
		require.NoError(t, err)
		require.Equal(t, 0, len(exercise.sets))

	})
}

func TestEditExercise(t *testing.T) {
	exercise, err := NewExercise(random.String(75), random.Weight(), random.RestTime(), random.RepCount(), 0)
	require.NoError(t, err)
	require.NotEmpty(t, exercise)

	t.Run("When negative weight", func(t *testing.T) {
		err = exercise.EditExercise(random.String(75), -1, random.RestTime(), random.RepCount())
		require.Error(t, err)
	})

	t.Run("When success", func(t *testing.T) {
		gen := random.Weight()
		err = exercise.EditExercise(random.String(75), gen, random.RestTime(), random.RepCount())
		require.NoError(t, err)
		require.Equal(t, gen, exercise.weight)
	})

}
