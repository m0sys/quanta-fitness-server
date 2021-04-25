package athlete

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	wl "github.com/mhd53/quanta-fitness-server/workoutlog"
	"github.com/stretchr/testify/require"
)

func TestSetHeight(t *testing.T) {
	athlete := NewAthlete()
	t.Run("When negative height", func(t *testing.T) {
		err := athlete.SetHeight(-1.0)
		require.Error(t, err)
		require.Equal(t, 0.0, athlete.Height)
	})

	t.Run("When success", func(t *testing.T) {
		gen := random.Height()
		err := athlete.SetHeight(gen)
		require.NoError(t, err)
		require.Equal(t, gen, athlete.Height)
	})
}

func TestUpdateHeight(t *testing.T) {
	athlete := NewAthlete()
	require.Equal(t, 0, len(athlete.WeightHistory))

	t.Run("When weight is negative", func(t *testing.T) {
		res, err := athlete.UpdateWeight(-1)
		require.Error(t, err)
		require.Empty(t, res)
		require.Equal(t, 0, len(athlete.WeightHistory))
	})

	t.Run("When success", func(t *testing.T) {
		gen := random.Weight()
		res, err := athlete.UpdateWeight(gen)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		require.Equal(t, 1, len(athlete.WeightHistory))
		require.Equal(t, gen, athlete.WeightHistory[0].Amount)
		require.Equal(t, gen, res.Amount)

	})
}

func TestAddWorkoutLog(t *testing.T) {
	athlete := NewAthlete()
	wlog, err := wl.NewWorkoutLog(random.String(64))
	require.NoError(t, err)

	t.Run("When success", func(t *testing.T) {
		err = athlete.AddWorkoutLog(wlog)
		require.NoError(t, err)
		require.Equal(t, 1, len(athlete.WorkoutLogs))
		require.Equal(t, wlog.LogID, athlete.WorkoutLogs[0].LogID)
	})

	t.Run("When already logged", func(t *testing.T) {
		err = athlete.AddWorkoutLog(wlog)
		require.Error(t, err)
		require.Equal(t, 1, len(athlete.WorkoutLogs))
		require.Equal(t, wlog.LogID, athlete.WorkoutLogs[0].LogID)
	})
}

func TestRemoveWorkoutLog(t *testing.T) {
	athlete := NewAthlete()
	wlog, err := wl.NewWorkoutLog(random.String(64))
	require.NoError(t, err)

	t.Run("When success", func(t *testing.T) {
		err = athlete.AddWorkoutLog(wlog)
		require.NoError(t, err)
		require.Equal(t, 1, len(athlete.WorkoutLogs))
		require.Equal(t, wlog.LogID, athlete.WorkoutLogs[0].LogID)

		err := athlete.RemoveWorkoutLog(wlog)
		require.NoError(t, err)
		require.Equal(t, 0, len(athlete.WorkoutLogs))
	})

	t.Run("When not found", func(t *testing.T) {
		err := athlete.RemoveWorkoutLog(wlog)
		require.Error(t, err)
		require.Equal(t, 0, len(athlete.WorkoutLogs))
	})
}
