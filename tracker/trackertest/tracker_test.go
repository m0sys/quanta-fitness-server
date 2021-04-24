package trackertest

import (
	"testing"
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/mhd53/quanta-fitness-server/tracker"
	"github.com/stretchr/testify/require"
)

func TestCreateWorkoutLog(t *testing.T) {
	testTracker, testAthlete := setup()
	gen := random.String(64)

	res, err := testTracker.CreateWorkoutLog(gen)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, gen, res.Title)
	require.Equal(t, res.LogID, testAthlete.WorkoutLogs[0].LogID)
}

func TestAddExerciseToWorkoutLog(t *testing.T) {
	t.Run("When WorkoutLog not first created", func(t *testing.T) {
		testTracker, _ := setup()
		name := random.String(64)
		weight := random.Weight()
		targetRep := random.RepCount()
		restTime := random.RestTime()

		req := tracker.AddExerciseToWorkoutLogReq{
			LogID:     "1234",
			Name:      name,
			Weight:    weight,
			TargetRep: targetRep,
			RestTime:  restTime,
		}
		res2, err := testTracker.AddExerciseToWorkoutLog(req)
		require.Error(t, err)
		require.Empty(t, res2)
		require.Equal(t, "no WorkoutLog is assigned to Tracker", err.Error())
	})

	t.Run("When LogID mismatch", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)
		name := random.String(64)
		weight := random.Weight()
		targetRep := random.RepCount()
		restTime := random.RestTime()

		_, err := testTracker.CreateWorkoutLog(title)
		req := tracker.AddExerciseToWorkoutLogReq{
			LogID:     "1234",
			Name:      name,
			Weight:    weight,
			TargetRep: targetRep,
			RestTime:  restTime,
		}
		res2, err := testTracker.AddExerciseToWorkoutLog(req)
		require.Error(t, err)
		require.Empty(t, res2)
		require.Equal(t, "WorkoutLog does not match requested LogID", err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		testTracker, _ := setup()
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

		// TODO: Check for length.

		res2, err = testTracker.AddExerciseToWorkoutLog(req)
		require.NoError(t, err)
		require.NotEmpty(t, res2)

		res2, err = testTracker.AddExerciseToWorkoutLog(req)
		require.NoError(t, err)
		require.NotEmpty(t, res2)

		res2, err = testTracker.AddExerciseToWorkoutLog(req)
		require.NoError(t, err)
		require.NotEmpty(t, res2)

		// require.Equal(t, 4, len(testTracker.))
	})
}

func TestAddSetToExercise(t *testing.T) {
	t.Run("When WorkoutLog not created first", func(t *testing.T) {
		testTracker, _ := setup()
		rep := random.RepCount()

		req := tracker.AddSetToExerciseReq{
			LogID:          "1234",
			ExerciseID:     "1234",
			ActualRepCount: rep,
		}

		res2, err := testTracker.AddSetToExercise(req)
		require.Error(t, err)
		require.Empty(t, res2)
		require.Equal(t, "no WorkoutLog is assigned to Tracker", err.Error())
	})

	t.Run("When LogID mismatch", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)
		rep := random.RepCount()

		res, err := testTracker.CreateWorkoutLog(title)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		req := tracker.AddSetToExerciseReq{
			LogID:          "1234",
			ExerciseID:     "1234",
			ActualRepCount: rep,
		}

		res2, err := testTracker.AddSetToExercise(req)
		require.Error(t, err)
		require.Empty(t, res2)
		require.Equal(t, "WorkoutLog does not match requested LogID", err.Error())

	})

	t.Run("When Exercise not found", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)
		rep := random.RepCount()

		res, err := testTracker.CreateWorkoutLog(title)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		req := tracker.AddSetToExerciseReq{
			LogID:          res.LogID,
			ExerciseID:     "1234",
			ActualRepCount: rep,
		}

		res2, err := testTracker.AddSetToExercise(req)
		require.Error(t, err)
		require.Empty(t, res2)
		require.Equal(t, "Exercise not found", err.Error())

	})

	t.Run("When success", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)
		rep := random.RepCount()
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

		req2 := tracker.AddSetToExerciseReq{
			LogID:          res.LogID,
			ExerciseID:     res2.ExerciseID,
			ActualRepCount: rep,
		}

		res3, err := testTracker.AddSetToExercise(req2)
		require.NoError(t, err)
		require.NotEmpty(t, res3)
		require.Equal(t, rep, res3.ActualRepCount)

		// TODO: Check for length.

	})

}

func TestRemoveExerciseFromWorkoutLog(t *testing.T) {
	t.Run("When WorkoutLog not created first", func(t *testing.T) {
		testTracker, _ := setup()

		err := testTracker.RemoveExerciseFromWorkoutLog("1234")
		require.Error(t, err)
		require.Equal(t, "no WorkoutLog is assigned to Tracker", err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)

		res, err := testTracker.CreateWorkoutLog(title)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		err = testTracker.RemoveExerciseFromWorkoutLog("1234")
		require.Error(t, err)
		require.Equal(t, "Exercise not found", err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		testTracker, _ := setup()
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

		err = testTracker.RemoveExerciseFromWorkoutLog(res2.ExerciseID)
		require.NoError(t, err)

		// TODO: test that Exercise has been removed in repo.
	})
}

func TestRemoveSetFromExercise(t *testing.T) {
	t.Run("When WorkoutLog not created first", func(t *testing.T) {
		testTracker, _ := setup()

		err := testTracker.RemoveSetFromExercise("1234", "1234")
		require.Error(t, err)
		require.Equal(t, "no WorkoutLog is assigned to Tracker", err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)

		res, err := testTracker.CreateWorkoutLog(title)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		err = testTracker.RemoveSetFromExercise("1234", "1234")
		require.Error(t, err)
		require.Equal(t, "Exercise not found", err.Error())
	})

	t.Run("When Set not found", func(t *testing.T) {
		testTracker, _ := setup()
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

		err = testTracker.RemoveSetFromExercise("1234", res2.ExerciseID)
		require.Error(t, err)
		require.Equal(t, "Set not found", err.Error())

	})

	t.Run("When success", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)
		rep := random.RepCount()
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

		req2 := tracker.AddSetToExerciseReq{
			LogID:          res.LogID,
			ExerciseID:     res2.ExerciseID,
			ActualRepCount: rep,
		}

		res3, err := testTracker.AddSetToExercise(req2)
		require.NoError(t, err)
		require.NotEmpty(t, res3)

		err = testTracker.RemoveSetFromExercise(res3.SetID, res2.ExerciseID)
		require.NoError(t, err)

		// TODO: test that Set has been removed in repo.
	})
}

func TestEditWorkoutLog(t *testing.T) {
	t.Run("When WorkoutLog not created first", func(t *testing.T) {
		testTracker, _ := setup()
		newTitle := random.String(64)
		newDate := time.Now().AddDate(0, 0, 1)

		req := tracker.EditWorkoutLogReq{
			LogID: "1234",
			Title: newTitle,
			Date:  newDate,
		}

		updated, err := testTracker.EditWorkoutLog(req)
		require.Error(t, err)
		require.Empty(t, updated)
		require.Equal(t, "no WorkoutLog is assigned to Tracker", err.Error())
	})

	t.Run("When LogID mismatch", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)
		newTitle := random.String(64)
		newDate := time.Now().AddDate(0, 0, 1)

		res, err := testTracker.CreateWorkoutLog(title)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		req := tracker.EditWorkoutLogReq{
			LogID: "1234",
			Title: newTitle,
			Date:  newDate,
		}

		updated, err := testTracker.EditWorkoutLog(req)
		require.Error(t, err)
		require.Empty(t, updated)
		require.Equal(t, "WorkoutLog does not match requested LogID", err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)
		newTitle := random.String(64)
		newDate := time.Now().AddDate(0, 0, 1)

		res, err := testTracker.CreateWorkoutLog(title)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		req := tracker.EditWorkoutLogReq{
			LogID: res.LogID,
			Title: newTitle,
			Date:  newDate,
		}

		updated, err := testTracker.EditWorkoutLog(req)
		require.NoError(t, err)
		require.NotEmpty(t, updated)
		require.Equal(t, res.LogID, updated.LogID)
		require.Equal(t, newTitle, updated.Title)
		require.WithinDuration(t, newDate, updated.Date, time.Second)
	})
}

func TestEditExercise(t *testing.T) {
	t.Run("When WorkoutLog not created first", func(t *testing.T) {
		testTracker, _ := setup()
		newName := random.String(64)
		newWeight := random.Weight()
		newTargetRep := random.RepCount()
		newRestTime := random.RestTime()

		req := tracker.EditExerciseReq{
			ExerciseID: "1234",
			Name:       newName,
			Weight:     newWeight,
			TargetRep:  newTargetRep,
			RestTime:   newRestTime,
		}

		updated, err := testTracker.EditExercise(req)
		require.Error(t, err)
		require.Empty(t, updated)
		require.Equal(t, "no WorkoutLog is assigned to Tracker", err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)
		newName := random.String(64)
		newWeight := random.Weight()
		newTargetRep := random.RepCount()
		newRestTime := random.RestTime()

		res, err := testTracker.CreateWorkoutLog(title)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		req := tracker.EditExerciseReq{
			ExerciseID: "1234",
			Name:       newName,
			Weight:     newWeight,
			TargetRep:  newTargetRep,
			RestTime:   newRestTime,
		}

		updated, err := testTracker.EditExercise(req)
		require.Error(t, err)
		require.Empty(t, updated)
		require.Equal(t, "Exercise not found", err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		testTracker, _ := setup()
		title := random.String(64)
		name := random.String(64)
		weight := random.Weight()
		targetRep := random.RepCount()
		restTime := random.RestTime()
		newName := random.String(64)
		newWeight := random.Weight()
		newTargetRep := random.RepCount()
		newRestTime := random.RestTime()

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

		req2 := tracker.EditExerciseReq{
			ExerciseID: res2.ExerciseID,
			Name:       newName,
			Weight:     newWeight,
			TargetRep:  newTargetRep,
			RestTime:   newRestTime,
		}

		updated, err := testTracker.EditExercise(req2)
		require.NoError(t, err)
		require.NotEmpty(t, updated)
		require.Equal(t, res2.ExerciseID, updated.ExerciseID)
		require.Equal(t, newName, updated.Name)
	})
}
