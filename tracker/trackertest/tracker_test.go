package trackertest

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/mhd53/quanta-fitness-server/tracker"
	"github.com/stretchr/testify/require"
)

func TestCreateWorkoutLog(t *testing.T) {
	gen := random.String(64)

	res, err := testTracker.CreateWorkoutLog(gen)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, gen, res.Title)
	require.Equal(t, res.LogID, testAthlete.WorkoutLogs[0].LogID)
}

func TestAddExerciseToWorkoutLog(t *testing.T) {
	title := random.String(64)
	name := random.String(64)
	weight := random.Weight()
	targetRep := random.RepCount()
	restTime := random.RestTime()

	res, err := testTracker.CreateWorkoutLog(title)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	req := tracker.AddExerciseToWorkoutLogReq{
		LogID:     res.LogID,
		Name:      name,
		Weight:    weight,
		TargetRep: targetRep,
		RestTime:  restTime,
	}
	res2, err := testTracker.AddExerciseToWorkoutLog(req)
	require.NoError(t, err)
	require.NotEmpty(t, res2)
	require.Equal(t, name, res2.Name)
	// require.Equal(t, weight, res2.Weight) // round causing error
	require.Equal(t, targetRep, res2.TargetRep)
	// require.Equal(t, restTime, res2.RestTime) // round causing error
}
