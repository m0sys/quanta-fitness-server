package adapters

import (
	"testing"

	"github.com/m0sys/quanta-fitness-server/internal/random"
	"github.com/m0sys/quanta-fitness-server/manager/athlete"
	el "github.com/m0sys/quanta-fitness-server/tracker/exerciselog"
	sl "github.com/m0sys/quanta-fitness-server/tracker/setlog"
	wl "github.com/m0sys/quanta-fitness-server/tracker/workoutlog"
	"github.com/stretchr/testify/require"
)

const (
	// FIXME: replace this with a dynamic aid when manger infra is implemented.
	aid = "59be62ed-4c63-4d25-9327-cd29664a1b71"
)

func TestStoreWorkoutLog_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		// TODO: Check that its stored.
	})
}

func TestStoreExerciseLog_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		metrics := el.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		elog := el.NewExerciseLog(wlog.ID(), random.String(10), metrics, 0)
		err = psqlRepo.StoreExerciseLog(elog)
		require.NoError(t, err)

		// TODO: Check that its stored.
	})
}

func TestStoreSetLog_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		metrics := el.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		elog := el.NewExerciseLog(wlog.ID(), random.String(10), metrics, 0)
		err = psqlRepo.StoreExerciseLog(elog)
		require.NoError(t, err)

		metrics2 := sl.NewMetrics(random.RepCount(), random.RestTime())
		slog := sl.NewSetLog(elog.ID(), metrics2)
		err = psqlRepo.StoreSetLog(slog)
		require.NoError(t, err)

		// TODO: Check that its stored.
	})
}

func TestFindAllWorkoutLogsForAthlete_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		wlogs, err := psqlRepo.FindAllWorkoutLogsForAthlete(ath)
		require.NoError(t, err)
		require.NotEmpty(t, wlogs)
	})
}

func TestFindAllExerciseLogsForWorkoutLog_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		n := 5
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		metrics := el.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)

		for i := 0; i < n; i++ {
			elog := el.NewExerciseLog(wlog.ID(), random.String(10), metrics, 0)
			err = psqlRepo.StoreExerciseLog(elog)
			require.NoError(t, err)
		}

		elogs, err := psqlRepo.FindAllExerciseLogsForWorkoutLog(wlog)
		require.NoError(t, err)
		require.NotEmpty(t, elogs)
		require.Equal(t, n, len(elogs))
	})
}

func TestFindAllSetLogsForExerciseLog_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		n := 5
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		metrics := el.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)

		elog := el.NewExerciseLog(wlog.ID(), random.String(10), metrics, 0)
		err = psqlRepo.StoreExerciseLog(elog)
		require.NoError(t, err)

		metrics2 := sl.NewMetrics(random.RepCount(), random.RestTime())
		for i := 0; i < n; i++ {
			slog := sl.NewSetLog(elog.ID(), metrics2)
			err = psqlRepo.StoreSetLog(slog)
			require.NoError(t, err)
		}

		slogs, err := psqlRepo.FindAllSetLogsForExerciseLog(elog)
		require.NoError(t, err)
		require.NotEmpty(t, slogs)
		require.Equal(t, n, len(slogs))
	})
}

func TestFindWorkouLogByID_psql(t *testing.T) {
	psqlRepo, ath := setup()

	t.Run("When not found", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		found, err := psqlRepo.FindWorkoutLogByID(wlog)
		require.NoError(t, err)
		require.False(t, found)
	})

	t.Run("When success", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		found, err := psqlRepo.FindWorkoutLogByID(wlog)
		require.NoError(t, err)
		require.True(t, found)
	})
}

func TestFindExerciseLogByID_psql(t *testing.T) {
	psqlRepo, ath := setup()

	t.Run("When not found", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		metrics := el.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		elog := el.NewExerciseLog(wlog.ID(), random.String(10), metrics, 0)
		found, err := psqlRepo.FindExerciseLogByID(elog)
		require.NoError(t, err)
		require.False(t, found)
	})

	t.Run("When success", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		metrics := el.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		elog := el.NewExerciseLog(wlog.ID(), random.String(10), metrics, 0)
		err = psqlRepo.StoreExerciseLog(elog)
		require.NoError(t, err)

		found, err := psqlRepo.FindExerciseLogByID(elog)
		require.NoError(t, err)
		require.True(t, found)
	})
}

func TestUpdateWorkoutLog_psql(t *testing.T) {
	psqlRepo, ath := setup()

	t.Run("When success", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		wlog.NextPos()
		wlog.Complete()

		err = psqlRepo.UpdateWorkoutLog(wlog)
		require.NoError(t, err)

		wlogs, err := psqlRepo.FindAllWorkoutLogsForAthlete(ath)
		length := len(wlogs)
		require.NoError(t, err)
		require.NotEmpty(t, wlogs)
		require.Equal(t, 1, wlogs[length-1].CurrentPos())
		require.True(t, wlogs[length-1].Completed())
	})
}

func TestRemoveWorkoutLog_psql(t *testing.T) {
	psqlRepo, ath := setup()

	t.Run("When success", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		err = psqlRepo.RemoveWorkoutLog(wlog)
		require.NoError(t, err)

		found, err := psqlRepo.FindWorkoutLogByID(wlog)
		require.NoError(t, err)
		require.False(t, found)
	})
}

func TestRemoveExerciseLog_psql(t *testing.T) {
	psqlRepo, ath := setup()

	t.Run("When success", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		metrics := el.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		elog := el.NewExerciseLog(wlog.ID(), random.String(10), metrics, 0)
		err = psqlRepo.StoreExerciseLog(elog)
		require.NoError(t, err)

		err = psqlRepo.RemoveExerciseLog(elog)
		require.NoError(t, err)

		found, err := psqlRepo.FindExerciseLogByID(elog)
		require.NoError(t, err)
		require.False(t, found)
	})
}

func TestRemoveSetLog_psql(t *testing.T) {
	psqlRepo, ath := setup()

	t.Run("When success", func(t *testing.T) {
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(10))
		require.NotEmpty(t, wlog)

		err := psqlRepo.StoreWorkoutLog(wlog)
		require.NoError(t, err)

		metrics := el.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		elog := el.NewExerciseLog(wlog.ID(), random.String(10), metrics, 0)
		err = psqlRepo.StoreExerciseLog(elog)
		require.NoError(t, err)

		metrics2 := sl.NewMetrics(random.RepCount(), random.RestTime())
		slog := sl.NewSetLog(elog.ID(), metrics2)
		err = psqlRepo.StoreSetLog(slog)
		require.NoError(t, err)

		err = psqlRepo.RemoveSetLog(slog)
		require.NoError(t, err)

		slogs, err := psqlRepo.FindAllSetLogsForExerciseLog(elog)
		require.NoError(t, err)
		require.Empty(t, slogs)
	})
}

// Helper funcs.
func setup() (PsqlTrainingRepository, athlete.Athlete) {
	ath := athlete.RestoreAthlete(aid)
	return NewPSQLRepo(testStore), ath
}
