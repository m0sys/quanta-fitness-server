package workoutplan

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/stretchr/testify/require"
)

func TestNewWorkoutPlan(t *testing.T) {
	t.Run("When title is more than 75 chars", func(t *testing.T) {
		wplan, err := NewWorkoutPlan("1", random.String(76))
		require.Error(t, err)
		require.Empty(t, wplan)
	})

	t.Run("When title is 75 chars less", func(t *testing.T) {
		gen := random.String(75)
		wplan, err := NewWorkoutPlan("1", gen)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)
		require.Equal(t, gen, wplan.Title())
	})
}

func TestEditTitle(t *testing.T) {
	gen := random.String(75)
	wplan, err := NewWorkoutPlan("1", gen)
	require.NoError(t, err)
	require.NotEmpty(t, wplan)
	require.Equal(t, gen, wplan.Title())
	t.Run("When title is more than 75 chars", func(t *testing.T) {
		err = wplan.EditTitle(random.String(76))
		require.Error(t, err)
	})

	t.Run("When title is 75 chars less", func(t *testing.T) {
		gen2 := random.String(75)
		err = wplan.EditTitle(gen2)
		require.NoError(t, err)

		require.NotEqual(t, gen, wplan.Title())
		require.Equal(t, gen2, wplan.Title())
	})
}
