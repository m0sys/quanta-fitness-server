package translator

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	pa "github.com/mhd53/quanta-fitness-server/planner/adapters"
	e "github.com/mhd53/quanta-fitness-server/planner/exercise"
	p "github.com/mhd53/quanta-fitness-server/planner/planning"
	"github.com/mhd53/quanta-fitness-server/planner/workoutplan"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
	ta "github.com/mhd53/quanta-fitness-server/tracker/adapters"
	el "github.com/mhd53/quanta-fitness-server/tracker/exerciselog"
	sl "github.com/mhd53/quanta-fitness-server/tracker/setlog"
	ts "github.com/mhd53/quanta-fitness-server/tracker/training"
	wl "github.com/mhd53/quanta-fitness-server/tracker/workoutlog"
	"github.com/stretchr/testify/require"
)

func TestConvertWorkout(t *testing.T) {
	wt, pService, _, ath := setup()

	t.Run("When no Exercise in WorkoutPlan", func(t *testing.T) {
		wplan := wplanSetup(t, ath, pService)
		wlog, err := wt.ConvertWorkoutPlan(wplan)
		require.Error(t, err)
		require.Empty(t, wlog)
		require.Equal(t, ErrWorkoutPlanHasNoExercises.Error(), err.Error())
	})

	t.Run("When WorkoutPlan has at least one Exercise", func(t *testing.T) {
		wplan := wplanSetup(t, ath, pService)
		exerciseSetup(t, ath, wplan, pService)
		wlog, err := wt.ConvertWorkoutPlan(wplan)
		require.NoError(t, err)
		require.NotEmpty(t, wlog)
		require.Equal(t, wplan.Title(), wlog.Title())

		// TODO: Check that WorkoutLog is stored.
	})
}

func TestFetchWorkoutLogs(t *testing.T) {
	wt, pService, tService, ath := setup()

	t.Run("When no WorkoutLog for Athlete", func(t *testing.T) {
		wlogs, err := tService.FetchWorkoutLogs(ath)
		require.NoError(t, err)
		require.Empty(t, wlogs)
	})

	t.Run("When  WorkoutLogs for Athlete exist", func(t *testing.T) {
		wplan := wplanSetup(t, ath, pService)
		exerciseSetup(t, ath, wplan, pService)
		n := 5
		for i := 0; i < n; i++ {
			wlog, err := wt.ConvertWorkoutPlan(wplan)
			require.NoError(t, err)
			require.NotEmpty(t, wlog)
		}

		wlogs, err := tService.FetchWorkoutLogs(ath)
		require.NoError(t, err)
		require.NotEmpty(t, wlogs)
		require.Equal(t, n, len(wlogs))
	})

}

func TestFetchWorkoutLogExerciseLogs(t *testing.T) {
	wt, pService, tService, ath := setup()

	t.Run("When WorkoutLog not found", func(t *testing.T) {
		wlog := wlogNotFoundSetup(ath)
		elogs, err := tService.FetchWorkoutLogExerciseLogs(ath, wlog)
		require.Error(t, err)
		require.Empty(t, elogs)
		require.Equal(t, ts.ErrWorkoutLogNotFound.Error(), err.Error())
	})

	t.Run("When Unauthorized", func(t *testing.T) {
		wplan := wplanSetup(t, ath, pService)
		n := 5
		for i := 0; i < n; i++ {
			exerciseSetup(t, ath, wplan, pService)
		}
		wlog, err := wt.ConvertWorkoutPlan(wplan)
		require.NoError(t, err)
		require.NotEmpty(t, wlog)
		require.Equal(t, wplan.Title(), wlog.Title())

		ath2 := athlete.NewAthlete()
		elogs, err := tService.FetchWorkoutLogExerciseLogs(ath2, wlog)
		require.Error(t, err)
		require.Equal(t, ts.ErrUnauthorizedAccess.Error(), err.Error())
		require.Empty(t, elogs)

	})

	t.Run("When success", func(t *testing.T) {
		wlog, n := wlogSuccesSetup(t, ath, pService, wt)

		elogs, err := tService.FetchWorkoutLogExerciseLogs(ath, wlog)
		require.NoError(t, err)
		require.NotEmpty(t, elogs)
		require.Equal(t, n, len(elogs))

	})

}

func TestMoveToNextExerciseLog(t *testing.T) {
	wt, pService, tService, ath := setup()
	t.Run("When WorkoutLog not found", func(t *testing.T) {
		wlog := wlogNotFoundSetup(ath)
		_, elog, err := tService.MoveToNextExerciseLog(ath, wlog)
		require.Error(t, err)
		require.Empty(t, elog)
		require.Equal(t, ts.ErrWorkoutLogNotFound.Error(), err.Error())
	})

	t.Run("When unauthorized access", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		_, elog, err := tService.MoveToNextExerciseLog(ath2, wlog)
		require.Error(t, err)
		require.Empty(t, elog)
		require.Equal(t, ts.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		_, elog, err := tService.MoveToNextExerciseLog(ath, wlog)
		require.NoError(t, err)
		require.NotEmpty(t, elog)
	})

	t.Run("When WorkoutLog already completed", func(t *testing.T) {
		var elog el.ExerciseLog
		var wlog wl.WorkoutLog
		var err error
		n := 5
		wlog, _ = wlogSuccesSetup(t, ath, pService, wt)

		for i := 0; i < n; i++ {
			currPos := wlog.CurrentPos()
			wlog, elog, err = tService.MoveToNextExerciseLog(ath, wlog)
			require.NoError(t, err)
			require.NotEmpty(t, elog)
			require.Equal(t, currPos+1, wlog.CurrentPos())
		}

		currPos := wlog.CurrentPos()
		wlog, elog, err = tService.MoveToNextExerciseLog(ath, wlog)
		require.Error(t, err)
		require.Empty(t, elog)
		require.Equal(t, ts.ErrWorkoutLogAlreadyCompleted.Error(), err.Error())
		require.Equal(t, currPos, wlog.CurrentPos())
	})
}

func TestAddSetLogToExerciseLog(t *testing.T) {
	wt, pService, tService, ath := setup()
	t.Run("When WorkoutLog not found", func(t *testing.T) {
		wlog := wlogNotFoundSetup(ath)
		elog := elogNotFoundSetup(wlog)
		metrics := sl.NewMetrics(random.RepCount(), random.RestTime())
		err := tService.AddSetLogToExerciseLog(ath, wlog, elog, metrics)
		require.Error(t, err)
		require.Equal(t, ts.ErrWorkoutLogNotFound.Error(), err.Error())
	})

	t.Run("When Unauthorized WorkoutLog", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wlog := wlogNotFoundSetup(ath2)
		elog := elogNotFoundSetup(wlog)
		metrics := sl.NewMetrics(random.RepCount(), random.RestTime())
		err := tService.AddSetLogToExerciseLog(ath, wlog, elog, metrics)
		require.Error(t, err)
		require.Equal(t, ts.ErrUnauthorizedAccess.Error(), err.Error())

	})

	t.Run("When ExerciseLog not found", func(t *testing.T) {
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		elog := elogNotFoundSetup(wlog)
		metrics := sl.NewMetrics(random.RepCount(), random.RestTime())
		err := tService.AddSetLogToExerciseLog(ath, wlog, elog, metrics)
		require.Error(t, err)
		require.Equal(t, ts.ErrExerciseLogNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)

		elogs, err := tService.FetchWorkoutLogExerciseLogs(ath, wlog)
		require.NoError(t, err)
		elog := elogs[0]
		repCount := random.RepCount()
		metrics := sl.NewMetrics(repCount, random.RestTime())

		err = tService.AddSetLogToExerciseLog(ath, wlog, elog, metrics)
		require.NoError(t, err)

		slogs, err := tService.FetchSetLogsForExerciseLog(ath, wlog, elog)
		require.NoError(t, err)
		require.NotEmpty(t, slogs)
		metrics = slogs[0].Metrics()
		require.Equal(t, repCount, metrics.ActualRepCount())
	})

	t.Run("When attempting to exceed num sets for ExerciseLog", func(t *testing.T) {
		n := 4
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)

		elogs, err := tService.FetchWorkoutLogExerciseLogs(ath, wlog)
		require.NoError(t, err)
		elog := elogs[0]
		repCount := random.RepCount()
		metrics := sl.NewMetrics(repCount, random.RestTime())

		for i := 0; i < n; i++ {
			err = tService.AddSetLogToExerciseLog(ath, wlog, elog, metrics)
			require.NoError(t, err)
		}

		err = tService.AddSetLogToExerciseLog(ath, wlog, elog, metrics)
		require.Error(t, err)
		require.Equal(t, ts.ErrCannotExceedNumSets.Error(), err.Error())
	})

}

func TestRemoveWorkoutLog(t *testing.T) {
	wt, pService, tService, ath := setup()
	t.Run("When WorkoutLog not found", func(t *testing.T) {
		wlog := wlogNotFoundSetup(ath)
		err := tService.RemoveWorkoutLog(ath, wlog)
		require.Error(t, err)
		require.Equal(t, ts.ErrWorkoutLogNotFound.Error(), err.Error())
	})

	t.Run("When unauthorized access to WorkoutLog", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wlog, _ := wlogSuccesSetup(t, ath2, pService, wt)
		err := tService.RemoveWorkoutLog(ath, wlog)
		require.Error(t, err)
		require.Equal(t, ts.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		err := tService.RemoveWorkoutLog(ath, wlog)
		require.NoError(t, err)

		// TODO: Check not found
		wlogs, err := tService.FetchWorkoutLogs(ath)
		require.NoError(t, err)
		require.Empty(t, wlogs)
	})
}

func wplanSetup(t *testing.T, ath athlete.Athlete, service p.PlanningService) workoutplan.WorkoutPlan {
	wplan, err := service.CreateNewWorkoutPlan(ath, random.String(75))
	require.NoError(t, err)
	require.NotEmpty(t, wplan)
	return wplan
}

func wlogNotFoundSetup(ath athlete.Athlete) wl.WorkoutLog {
	return wl.NewWorkoutLog(ath.AthleteID(), random.String(75))
}

func wlogSuccesSetup(t *testing.T, ath athlete.Athlete, pService p.PlanningService, wt WorkoutTranslator) (wl.WorkoutLog, int) {
	wplan := wplanSetup(t, ath, pService)
	n := 5
	for i := 0; i < n; i++ {
		exerciseSetup(t, ath, wplan, pService)
	}
	wlog, err := wt.ConvertWorkoutPlan(wplan)
	require.NoError(t, err)
	require.NotEmpty(t, wlog)
	require.Equal(t, wplan.Title(), wlog.Title())

	return wlog, n
}

func exerciseSetup(t *testing.T, ath athlete.Athlete, wplan wp.WorkoutPlan, service p.PlanningService) e.Exercise {
	name := random.String(75)

	metrics, err := e.NewMetrics(
		random.RepCount(),
		4,
		random.Weight(),
		random.RestTime(),
	)
	require.NoError(t, err)

	exercise, err := service.AddNewExerciseToWorkoutPlan(
		ath,
		wplan,
		name,
		metrics,
	)

	require.NoError(t, err)
	require.NotEmpty(t, exercise)
	require.Equal(t, name, exercise.Name())

	return exercise
}

func elogNotFoundSetup(wlog wl.WorkoutLog) el.ExerciseLog {
	name := random.String(75)

	metrics := el.NewMetrics(
		random.RepCount(),
		random.NumSets(),
		random.Weight(),
		random.RestTime(),
	)

	elog := el.NewExerciseLog(wlog.ID(), name, metrics, 0)
	return elog
}

func setup() (WorkoutTranslator, p.PlanningService, ts.TrainingService, athlete.Athlete) {
	planRepo := pa.NewInMemRepo()
	logRepo := ta.NewInMemRepo()
	pService := p.NewPlanningService(planRepo)
	tService := ts.NewTrainingService(logRepo)
	ath := athlete.NewAthlete()

	return NewWorkoutTranslator(planRepo, logRepo), pService, tService, ath
}
