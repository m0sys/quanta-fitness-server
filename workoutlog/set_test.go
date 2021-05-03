package workoutlog

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/stretchr/testify/require"
)

func TestNewSet(t *testing.T) {
	t.Run("When negative rep", func(t *testing.T) {
		set, err := NewSet(-1)
		require.Error(t, err)
		require.Empty(t, set)
	})

	t.Run("When success", func(t *testing.T) {
		gen := random.RepCount()
		set, err := NewSet(gen)
		require.NoError(t, err)
		require.NotEmpty(t, set)
		require.Equal(t, gen, set.ActualRepCount())
	})
}

func TestEditSet(t *testing.T) {
	gen := random.RepCount()
	set, err := NewSet(gen)
	require.NoError(t, err)
	require.NotEmpty(t, set)
	require.Equal(t, gen, set.ActualRepCount())

	t.Run("When negative rep", func(t *testing.T) {
		err = set.EditSet(-1)
		require.Error(t, err)
	})

	t.Run("When success", func(t *testing.T) {
		gen2 := random.RepCount()
		err = set.EditSet(gen2)
		require.NoError(t, err)
		require.Equal(t, gen2, set.ActualRepCount())
	})
}
