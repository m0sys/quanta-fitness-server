package managertest

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/stretchr/testify/require"
)

func TestFetchWeightHistory(t *testing.T) {
	t.Run("When Weight History is empty", func(t *testing.T) {
		manager := setup()

		fetched, err := manager.FetchWeightHistory()
		require.NoError(t, err)
		require.Equal(t, 0, len(fetched))

	})

	t.Run("When Weight added", func(t *testing.T) {
		manager := setup()
		n := 5

		weight := random.Weight()

		for i := 0; i < n; i++ {
			err := manager.UpdateWeight(weight)
			require.NoError(t, err)
		}
		fetched, err := manager.FetchWeightHistory()
		require.NoError(t, err)
		require.Equal(t, 5, len(fetched))
		require.Equal(t, weight, fetched[0].Amount)
	})
}
