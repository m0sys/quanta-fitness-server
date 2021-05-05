package athlete

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/stretchr/testify/require"
)

func TestSetHeight(t *testing.T) {
	athlete := NewAthlete()
	t.Run("When negative height", func(t *testing.T) {
		err := athlete.SetHeight(-1.0)
		require.Error(t, err)
		require.Equal(t, 0.0, athlete.Height())
	})

	t.Run("When success", func(t *testing.T) {
		gen := random.Height()
		err := athlete.SetHeight(gen)
		require.NoError(t, err)
		require.Equal(t, gen, athlete.Height())
	})
}

func TestUpdateHeight(t *testing.T) {
	athlete := NewAthlete()
	require.Equal(t, 0, len(athlete.weightHistory))

	t.Run("When weight is negative", func(t *testing.T) {
		res, err := athlete.UpdateWeight(-1)
		require.Error(t, err)
		require.Empty(t, res)
		require.Equal(t, 0, len(athlete.weightHistory))
	})

	t.Run("When success", func(t *testing.T) {
		gen := random.Weight()
		res, err := athlete.UpdateWeight(gen)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		require.Equal(t, 1, len(athlete.weightHistory))
		require.Equal(t, gen, athlete.weightHistory[0].amount)
		require.Equal(t, gen, res.amount)

	})
}
