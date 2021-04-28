package workoutlog

import (
	"testing"
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/stretchr/testify/require"
)

func TestNewWorkoutLog(t *testing.T) {
	t.Run("When title is more than 75 chars", func(t *testing.T) {
		wlog, err := NewWorkoutLog(random.String(76))
		require.Error(t, err)
		require.Empty(t, wlog)
	})

	t.Run("When title is 75 chars less", func(t *testing.T) {
		gen := random.String(75)
		wlog, err := NewWorkoutLog(gen)
		require.NoError(t, err)
		require.NotEmpty(t, wlog)
		require.Equal(t, gen, wlog.Title())
	})
}

func TestAddExercise(t *testing.T) {
	t.Run("When exercise is new", func(t *testing.T) {

		wlog, err := NewWorkoutLog(random.String(75))
		require.NoError(t, err)
		require.NotEmpty(t, wlog)

		exercise, err := NewExercise(random.String(75), random.Weight(), random.RestTime(), random.RepCount(), 0)
		require.NoError(t, err)
		require.NotEmpty(t, exercise)

		err = wlog.AddExercise(exercise)
		require.NoError(t, err)
		require.Equal(t, 1, len(wlog.exercises))

	})

	t.Run("When exercise is already logged", func(t *testing.T) {

		wlog, err := NewWorkoutLog(random.String(75))
		require.NoError(t, err)
		require.NotEmpty(t, wlog)

		exercise, err := NewExercise(random.String(75), random.Weight(), random.RestTime(), random.RepCount(), 0)
		require.NoError(t, err)
		require.NotEmpty(t, exercise)

		err = wlog.AddExercise(exercise)
		require.NoError(t, err)
		require.Equal(t, 1, len(wlog.exercises))

		err = wlog.AddExercise(exercise)
		require.Error(t, err)
		require.Equal(t, 1, len(wlog.exercises))

	})
}

func TestRemoveExercise(t *testing.T) {
	t.Run("When exercise not found", func(t *testing.T) {
		wlog, err := NewWorkoutLog(random.String(75))
		require.NoError(t, err)
		require.NotEmpty(t, wlog)

		exercise, err := NewExercise(random.String(75), random.Weight(), random.RestTime(), random.RepCount(), 0)
		require.NoError(t, err)
		require.NotEmpty(t, exercise)

		err = wlog.RemoveExercise(exercise)
		require.Error(t, err)
		require.Equal(t, 0, len(wlog.exercises))

	})

	t.Run("When exercise found", func(t *testing.T) {
		wlog, err := NewWorkoutLog(random.String(75))
		require.NoError(t, err)
		require.NotEmpty(t, wlog)

		exercise, err := NewExercise(random.String(75), random.Weight(), random.RestTime(), random.RepCount(), 0)
		require.NoError(t, err)
		require.NotEmpty(t, exercise)

		err = wlog.AddExercise(exercise)
		require.NoError(t, err)
		require.Equal(t, 1, len(wlog.exercises))

		err = wlog.RemoveExercise(exercise)
		require.NoError(t, err)
		require.Equal(t, 0, len(wlog.exercises))
	})
}

func TestEditWorkoutLog(t *testing.T) {
	gen := random.String(75)
	wlog, err := NewWorkoutLog(gen)
	require.NoError(t, err)
	require.NotEmpty(t, wlog)
	require.Equal(t, gen, wlog.Title())
	t.Run("When title is more than 75 chars", func(t *testing.T) {
		err = wlog.EditWorkoutLog(random.String(76), time.Now())
		require.Error(t, err)
	})

	t.Run("When title is 75 chars less", func(t *testing.T) {
		gen2 := random.String(75)
		err = wlog.EditWorkoutLog(gen2, time.Now())
		require.NoError(t, err)

		require.NotEqual(t, gen, wlog.Title())
		require.Equal(t, gen2, wlog.Title())
	})
}
