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
		wlog := wl.NewWorkoutLog(ath.AthleteID(), random.String(75))
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

		elogs, err := tService.FetchWorkoutLogExerciseLogs(ath, wlog)
		require.NoError(t, err)
		require.NotEmpty(t, elogs)
		require.Equal(t, n, len(elogs))

	})

	t.Run("When success", func(t *testing.T) {

	})

}

func wplanSetup(t *testing.T, ath athlete.Athlete, service p.PlanningService) workoutplan.WorkoutPlan {
	wplan, err := service.CreateNewWorkoutPlan(ath, random.String(75))
	require.NoError(t, err)
	require.NotEmpty(t, wplan)
	return wplan
}

func exerciseSetup(t *testing.T, ath athlete.Athlete, wplan wp.WorkoutPlan, service p.PlanningService) e.Exercise {
	name := random.String(75)

	metrics, err := e.NewMetrics(
		random.RepCount(),
		random.NumSets(),
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

func setup() (WorkoutTranslator, p.PlanningService, ts.TrainingService, athlete.Athlete) {
	planRepo := pa.NewInMemRepo()
	logRepo := ta.NewInMemRepo()
	pService := p.NewPlanningService(planRepo)
	tService := ts.NewTrainingService(logRepo)
	ath := athlete.NewAthlete()

	return NewWorkoutTranslator(planRepo, logRepo), pService, tService, ath
}
