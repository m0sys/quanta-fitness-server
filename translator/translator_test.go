package translator

import (
	"testing"

	"github.com/m0sys/quanta-fitness-server/internal/random"
	"github.com/m0sys/quanta-fitness-server/manager/athlete"
	p "github.com/m0sys/quanta-fitness-server/planner/planning"
	pa "github.com/m0sys/quanta-fitness-server/planner/planning/adapters"
	el "github.com/m0sys/quanta-fitness-server/tracker/exerciselog"
	ts "github.com/m0sys/quanta-fitness-server/tracker/training"
	ta "github.com/m0sys/quanta-fitness-server/tracker/training/adapters"
	wl "github.com/m0sys/quanta-fitness-server/tracker/workoutlog"
	"github.com/stretchr/testify/require"
)

func TestConvertWorkout(t *testing.T) {
	wt, pService, tService, ath := setup()

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
		require.Equal(t, wplan.Title, wlog.Title())

		wlogs, err := tService.FetchWorkoutLogs(ath.AthleteID())
		require.NoError(t, err)
		require.NotEmpty(t, wlogs)
	})
}

func TestFetchWorkoutLogs(t *testing.T) {
	wt, pService, tService, ath := setup()

	t.Run("When no WorkoutLog for Athlete", func(t *testing.T) {
		wlogs, err := tService.FetchWorkoutLogs(ath.AthleteID())
		require.NoError(t, err)
		require.Empty(t, wlogs)
	})

	t.Run("When  WorkoutLogs for Athlete exist", func(t *testing.T) {
		n := 5
		for i := 0; i < n; i++ {
			wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
			require.NotEmpty(t, wlog)
		}

		wlogs, err := tService.FetchWorkoutLogs(ath.AthleteID())
		require.NoError(t, err)
		require.NotEmpty(t, wlogs)
		require.Equal(t, n, len(wlogs))
	})

}

func TestFetchWorkoutLogExerciseLogs(t *testing.T) {
	wt, pService, tService, ath := setup()

	t.Run("When WorkoutLog not found", func(t *testing.T) {
		req := ts.FetchWorkoutLogExerciseLogsReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: "1234",
		}

		elogs, err := tService.FetchWorkoutLogExerciseLogs(req)
		require.Error(t, err)
		require.Empty(t, elogs)
		require.Equal(t, ts.ErrWorkoutLogNotFound.Error(), err.Error())
	})

	t.Run("When Unauthorized", func(t *testing.T) {
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)

		ath2 := athlete.NewAthlete()
		req := ts.FetchWorkoutLogExerciseLogsReq{
			AthleteID:    ath2.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}

		elogs, err := tService.FetchWorkoutLogExerciseLogs(req)
		require.Error(t, err)
		require.Equal(t, ts.ErrUnauthorizedAccess.Error(), err.Error())
		require.Empty(t, elogs)
	})

	t.Run("When success", func(t *testing.T) {
		wlog, n := wlogSuccesSetup(t, ath, pService, wt)

		req := ts.FetchWorkoutLogExerciseLogsReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}
		elogs, err := tService.FetchWorkoutLogExerciseLogs(req)
		require.NoError(t, err)
		require.NotEmpty(t, elogs)
		require.Equal(t, n, len(elogs))
	})
}

func TestAddSetLogToExerciseLog(t *testing.T) {
	wt, pService, tService, ath := setup()
	t.Run("When WorkoutLog not found", func(t *testing.T) {
		req := ts.AddSetLogToExerciseLogReq{
			AthleteID:      ath.AthleteID(),
			WorkoutLogID:   "1234",
			ExerciseLogID:  "1234",
			ActualRepCount: random.RepCount(),
			Duration:       random.RestTime(),
		}

		err := tService.AddSetLogToExerciseLog(req)
		require.Error(t, err)
		require.Equal(t, ts.ErrWorkoutLogNotFound.Error(), err.Error())
	})

	t.Run("When ExerciseLog not found", func(t *testing.T) {
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		req := ts.AddSetLogToExerciseLogReq{
			AthleteID:      ath.AthleteID(),
			WorkoutLogID:   wlog.ID(),
			ExerciseLogID:  "1234",
			ActualRepCount: random.RepCount(),
			Duration:       random.RestTime(),
		}

		err := tService.AddSetLogToExerciseLog(req)
		require.Error(t, err)
		require.Equal(t, ts.ErrExerciseLogNotFound.Error(), err.Error())
	})

	t.Run("When Unauthorized WorkoutLog", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		req := ts.FetchWorkoutLogExerciseLogsReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}
		elogs, err := tService.FetchWorkoutLogExerciseLogs(req)
		require.NoError(t, err)
		elog := elogs[0]

		req2 := ts.AddSetLogToExerciseLogReq{
			AthleteID:      ath2.AthleteID(),
			WorkoutLogID:   wlog.ID(),
			ExerciseLogID:  elog.ID,
			ActualRepCount: random.RepCount(),
			Duration:       random.RestTime(),
		}

		err = tService.AddSetLogToExerciseLog(req2)
		require.Error(t, err)
		require.Equal(t, ts.ErrUnauthorizedAccess.Error(), err.Error())

	})

	t.Run("When success", func(t *testing.T) {
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		req := ts.FetchWorkoutLogExerciseLogsReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}
		repCount := random.RepCount()
		elogs, err := tService.FetchWorkoutLogExerciseLogs(req)
		require.NoError(t, err)
		elog := elogs[0]

		req2 := ts.AddSetLogToExerciseLogReq{
			AthleteID:      ath.AthleteID(),
			WorkoutLogID:   wlog.ID(),
			ExerciseLogID:  elog.ID,
			ActualRepCount: repCount,
			Duration:       random.RestTime(),
		}
		err = tService.AddSetLogToExerciseLog(req2)
		require.NoError(t, err)

		req3 := ts.FetchSetLogsForExerciseLogReq{
			AthleteID:     ath.AthleteID(),
			WorkoutLogID:  wlog.ID(),
			ExerciseLogID: elog.ID,
		}
		slogs, err := tService.FetchSetLogsForExerciseLog(req3)
		require.NoError(t, err)
		require.NotEmpty(t, slogs)
		require.Equal(t, repCount, slogs[0].ActualRepCount)
	})

	t.Run("When attempting to exceed num sets for ExerciseLog", func(t *testing.T) {
		n := 4
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		req := ts.FetchWorkoutLogExerciseLogsReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}

		elogs, err := tService.FetchWorkoutLogExerciseLogs(req)
		require.NoError(t, err)
		elog := elogs[0]
		req2 := ts.AddSetLogToExerciseLogReq{
			AthleteID:      ath.AthleteID(),
			WorkoutLogID:   wlog.ID(),
			ExerciseLogID:  elog.ID,
			ActualRepCount: random.RepCount(),
			Duration:       random.RestTime(),
		}

		for i := 0; i < n; i++ {
			err = tService.AddSetLogToExerciseLog(req2)
			require.NoError(t, err)
		}

		err = tService.AddSetLogToExerciseLog(req2)
		require.Error(t, err)
		require.Equal(t, ts.ErrCannotExceedNumSets.Error(), err.Error())
	})

}

func TestMoveToNextExerciseLog(t *testing.T) {
	wt, pService, tService, ath := setup()
	t.Run("When WorkoutLog not found", func(t *testing.T) {
		req := ts.MoveToNextExerciseLogReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: "1234",
		}

		err := tService.MoveToNextExerciseLog(req)
		require.Error(t, err)
		require.Equal(t, ts.ErrWorkoutLogNotFound.Error(), err.Error())
	})

	t.Run("When unauthorized access", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		req := ts.MoveToNextExerciseLogReq{
			AthleteID:    ath2.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}
		err := tService.MoveToNextExerciseLog(req)
		require.Error(t, err)
		require.Equal(t, ts.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		req := ts.MoveToNextExerciseLogReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}
		err := tService.MoveToNextExerciseLog(req)
		require.NoError(t, err)

		req2 := ts.FetchCurrentExerciseLogReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}
		res, err := tService.FetchCurrentExerciseLog(req2)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		req3 := ts.FetchWorkoutLogExerciseLogsReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}

		results, err := tService.FetchWorkoutLogExerciseLogs(req3)
		require.NoError(t, err)
		require.NotEmpty(t, results)
		require.Equal(t, results[1].ID, res.ID)
	})

	t.Run("When WorkoutLog already completed", func(t *testing.T) {
		n := 5
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		req := ts.MoveToNextExerciseLogReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}

		for i := 0; i < n; i++ {
			err := tService.MoveToNextExerciseLog(req)
			require.NoError(t, err)
		}

		err := tService.MoveToNextExerciseLog(req)
		require.Error(t, err)
		require.Equal(t, ts.ErrWorkoutLogAlreadyCompleted.Error(), err.Error())
	})
}

func TestRemoveWorkoutLog(t *testing.T) {
	wt, pService, tService, ath := setup()
	t.Run("When WorkoutLog not found", func(t *testing.T) {
		req := ts.RemoveWorkoutLogReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: "1234",
		}
		err := tService.RemoveWorkoutLog(req)
		require.Error(t, err)
		require.Equal(t, ts.ErrWorkoutLogNotFound.Error(), err.Error())
	})

	t.Run("When unauthorized access to WorkoutLog", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		req := ts.RemoveWorkoutLogReq{
			AthleteID:    ath2.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}
		err := tService.RemoveWorkoutLog(req)
		require.Error(t, err)
		require.Equal(t, ts.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wlog, _ := wlogSuccesSetup(t, ath, pService, wt)
		req := ts.RemoveWorkoutLogReq{
			AthleteID:    ath.AthleteID(),
			WorkoutLogID: wlog.ID(),
		}
		err := tService.RemoveWorkoutLog(req)
		require.NoError(t, err)

		wlogs, err := tService.FetchWorkoutLogs(ath.AthleteID())
		require.NoError(t, err)
		for _, wlog2 := range wlogs {
			require.NotEqual(t, wlog.ID(), wlog2.ID)
		}
	})
}

func wplanSetup(t *testing.T, ath athlete.Athlete, service p.PlanningService) p.WorkoutPlanRes {
	req := p.CreateNewWorkoutPlanReq{
		AthleteID: ath.AthleteID(),
		Title:     random.String(75),
	}
	res, err := service.CreateNewWorkoutPlan(req)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	return res
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
	require.Equal(t, wplan.Title, wlog.Title())

	return wlog, n
}

func exerciseSetup(t *testing.T, ath athlete.Athlete, wplan p.WorkoutPlanRes, service p.PlanningService) p.ExerciseRes {
	name := random.String(75)

	req := p.AddNewExerciseToWorkoutPlanReq{
		AthleteID:     ath.AthleteID(),
		WorkoutPlanID: wplan.ID,
		Name:          name,
		TargetRep:     random.RepCount(),
		NumSets:       4,
		Weight:        random.Weight(),
		RestDur:       random.RestTime(),
	}

	res, err := service.AddNewExerciseToWorkoutPlan(req)

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, name, res.Name)

	return res
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
