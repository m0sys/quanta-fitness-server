package planningtest

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	"github.com/mhd53/quanta-fitness-server/planner/adapters"
	e "github.com/mhd53/quanta-fitness-server/planner/exercise"
	p "github.com/mhd53/quanta-fitness-server/planner/planning"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
	"github.com/stretchr/testify/require"
)

func TestCreateNewWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When success", func(t *testing.T) {
		wplan, title := workoutPlanSuccessSetup(t, ath, service)
		require.NotEmpty(t, wplan)

		wplans, err := service.FetchWorkoutPlans(ath)
		require.NoError(t, err)
		require.NotEmpty(t, wplans)
		require.Equal(t, title, wplans[0].Title())
	})

	t.Run("When WorkoutPlan with given title already exists", func(t *testing.T) {
		_, title := workoutPlanSuccessSetup(t, ath, service)

		wplan, err := service.CreateNewWorkoutPlan(ath, title)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialTitle.Error(), err.Error())
		require.Empty(t, wplan)
	})
}

func TestAddNewExerciseToWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		name := random.String(75)
		wplan := workoutPlanNotFoundSetup(t, ath)

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
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
		require.Empty(t, exercise)
	})

	t.Run("When WorkoutPlan doesn't belong to Athlete", func(t *testing.T) {
		name := random.String(75)
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		ath2 := athlete.NewAthlete()

		metrics, err := e.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		require.NoError(t, err)

		exercise, err := service.AddNewExerciseToWorkoutPlan(
			ath2,
			wplan,
			name,
			metrics,
		)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
		require.Empty(t, exercise)
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _, _ := exerciseSuccessSetup(t, ath, wplan, service)
		require.NotEmpty(t, exercise)

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, exercise.ID(), exercises[0].ID())
	})

	t.Run("When Exercise with same name already in WorkoutPlan", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		_, metrics, name := exerciseSuccessSetup(t, ath, wplan, service)

		exercise, err := service.AddNewExerciseToWorkoutPlan(
			ath,
			wplan,
			name,
			metrics,
		)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialName.Error(), err.Error())
		require.Empty(t, exercise)
	})

}

func TestRemoveExerciseFromWorkoutPlan(t *testing.T) {
	service, ath := setup()

	t.Run("When unauthorized to access WorkoutPlan", func(t *testing.T) {
		wplan := workoutPlanUnauthorizedSetup(t)
		exercise := exerciseNotFoundSetup(t, ath, wplan)

		err := service.RemoveExerciseFromWorkoutPlan(
			ath,
			wplan,
			exercise,
		)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When unauthorized to access Exercise", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseUnauthorizedSetup(t, wplan)

		err := service.RemoveExerciseFromWorkoutPlan(
			ath,
			wplan,
			exercise,
		)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan := workoutPlanNotFoundSetup(t, ath)
		exercise := exerciseNotFoundSetup(t, ath, wplan)

		err := service.RemoveExerciseFromWorkoutPlan(
			ath,
			wplan,
			exercise,
		)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})
	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseNotFoundSetup(t, ath, wplan)

		err := service.RemoveExerciseFromWorkoutPlan(
			ath,
			wplan,
			exercise,
		)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _, _ := exerciseSuccessSetup(t, ath, wplan, service)

		err := service.RemoveExerciseFromWorkoutPlan(
			ath,
			wplan,
			exercise,
		)
		require.NoError(t, err)

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.NoError(t, err)
		require.Empty(t, exercises)
	})
}

func TestEditWorkoutPlanTitle(t *testing.T) {
	service, ath := setup()
	t.Run("When unauthorized", func(t *testing.T) {
		wplan := workoutPlanUnauthorizedSetup(t)
		title2 := random.String(75)

		err := service.EditWorkoutPlanTitle(ath, wplan, title2)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan := workoutPlanNotFoundSetup(t, ath)
		title2 := random.String(75)

		err := service.EditWorkoutPlanTitle(ath, wplan, title2)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})
	t.Run("When same title", func(t *testing.T) {
		ath := athlete.NewAthlete()
		wplan, title := workoutPlanSuccessSetup(t, ath, service)

		err := service.EditWorkoutPlanTitle(ath, wplan, title)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialTitle.Error(), err.Error())

		wplans, err := service.FetchWorkoutPlans(ath)
		require.NoError(t, err)
		require.NotEmpty(t, wplans)
		require.Equal(t, title, wplans[0].Title())
		require.Equal(t, title, wplan.Title())
	})

	t.Run("When success", func(t *testing.T) {
		ath := athlete.NewAthlete()
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		title := random.String(75)

		err := service.EditWorkoutPlanTitle(ath, wplan, title)
		require.NoError(t, err)

		wplans, err := service.FetchWorkoutPlans(ath)
		require.NoError(t, err)
		require.NotEmpty(t, wplans)
		require.Equal(t, title, wplans[0].Title())
	})
}

func TestFetchWorkoutPlans(t *testing.T) {
	service, ath := setup()
	t.Run("When no WorkoutPlan for Athlete", func(t *testing.T) {
		wplans, err := service.FetchWorkoutPlans(ath)
		require.NoError(t, err)
		require.Empty(t, wplans)
	})

	t.Run("After creating WorkoutPlans for Athlete", func(t *testing.T) {
		n := 5
		for i := 0; i < n; i++ {
			workoutPlanSuccessSetup(t, ath, service)
		}
		wplans, err := service.FetchWorkoutPlans(ath)
		require.NoError(t, err)
		require.NotEmpty(t, wplans)
		require.Equal(t, n, len(wplans))
	})
}

func TestFetchWorkoutPlanExercises(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized", func(t *testing.T) {
		wplan := workoutPlanUnauthorizedSetup(t)

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
		require.Empty(t, exercises)
	})
	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan := workoutPlanNotFoundSetup(t, ath)

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
		require.Empty(t, exercises)
	})

	t.Run("When no Exercises for WorkoutPlan", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.NoError(t, err)
		require.Empty(t, exercises)
	})

	t.Run("After Exercises have been added to WorkoutPlan", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		n := 5
		for i := 0; i < n; i++ {
			exerciseSuccessSetup(t, ath, wplan, service)
		}

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, n, len(exercises))
	})

}

func TestEditExerciseName(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized WorkoutPlan", func(t *testing.T) {
		wplan := workoutPlanUnauthorizedSetup(t)
		exercise := exerciseUnauthorizedSetup(t, wplan)
		name := random.String(75)

		err := service.EditExerciseName(ath, wplan, exercise, name)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When Unauthorized Exercise", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseUnauthorizedSetup(t, wplan)
		name := random.String(75)

		err := service.EditExerciseName(ath, wplan, exercise, name)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan := workoutPlanNotFoundSetup(t, ath)
		exercise := exerciseNotFoundSetup(t, ath, wplan)
		name := random.String(75)

		err := service.EditExerciseName(ath, wplan, exercise, name)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseNotFoundSetup(t, ath, wplan)
		name := random.String(75)

		err := service.EditExerciseName(ath, wplan, exercise, name)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When same Exercise name", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _, name := exerciseSuccessSetup(t, ath, wplan, service)

		err := service.EditExerciseName(ath, wplan, exercise, name)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialName.Error(), err.Error())

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, name, exercises[0].Name())
		require.Equal(t, name, exercise.Name())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _, _ := exerciseSuccessSetup(t, ath, wplan, service)
		name := random.String(75)

		err := service.EditExerciseName(ath, wplan, exercise, name)
		require.NoError(t, err)

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, name, exercises[0].Name())
	})
}

func TestEditExerciseTargetRep(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized WorkoutPlan", func(t *testing.T) {
		wplan := workoutPlanUnauthorizedSetup(t)
		exercise := exerciseUnauthorizedSetup(t, wplan)
		targetRep := random.RepCount()

		err := service.EditExerciseTargetRep(ath, wplan, exercise, targetRep)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When Unauthorized Exercise", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseUnauthorizedSetup(t, wplan)
		targetRep := random.RepCount()

		err := service.EditExerciseTargetRep(ath, wplan, exercise, targetRep)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan := workoutPlanNotFoundSetup(t, ath)
		exercise := exerciseNotFoundSetup(t, ath, wplan)
		targetRep := random.RepCount()

		err := service.EditExerciseTargetRep(ath, wplan, exercise, targetRep)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseNotFoundSetup(t, ath, wplan)
		targetRep := random.RepCount()

		err := service.EditExerciseTargetRep(ath, wplan, exercise, targetRep)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _, _ := exerciseSuccessSetup(t, ath, wplan, service)
		targetRep := random.RepCount()

		err := service.EditExerciseTargetRep(ath, wplan, exercise, targetRep)
		require.NoError(t, err)

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)

		metrics := exercises[0].Metrics()
		require.Equal(t, targetRep, metrics.TargetRep())
	})
}

func TestEditExerciseNumSets(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized WorkoutPlan", func(t *testing.T) {
		wplan := workoutPlanUnauthorizedSetup(t)
		exercise := exerciseUnauthorizedSetup(t, wplan)
		numSets := random.NumSets()

		err := service.EditExerciseNumSets(ath, wplan, exercise, numSets)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When Unauthorized Exercise", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseUnauthorizedSetup(t, wplan)
		numSets := random.NumSets()

		err := service.EditExerciseNumSets(ath, wplan, exercise, numSets)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan := workoutPlanNotFoundSetup(t, ath)
		exercise := exerciseNotFoundSetup(t, ath, wplan)
		numSets := random.NumSets()

		err := service.EditExerciseNumSets(ath, wplan, exercise, numSets)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseNotFoundSetup(t, ath, wplan)
		numSets := random.NumSets()

		err := service.EditExerciseNumSets(ath, wplan, exercise, numSets)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _, _ := exerciseSuccessSetup(t, ath, wplan, service)
		numSets := random.NumSets()

		err := service.EditExerciseNumSets(ath, wplan, exercise, numSets)
		require.NoError(t, err)

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)

		metrics := exercises[0].Metrics()
		require.Equal(t, numSets, metrics.NumSets())
	})
}

func TestEditExerciseWeight(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized WorkoutPlan", func(t *testing.T) {
		wplan := workoutPlanUnauthorizedSetup(t)
		exercise := exerciseUnauthorizedSetup(t, wplan)
		weight := random.Weight()

		err := service.EditExerciseWeight(ath, wplan, exercise, weight)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When Unauthorized Exercise", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseUnauthorizedSetup(t, wplan)
		weight := random.Weight()

		err := service.EditExerciseWeight(ath, wplan, exercise, weight)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan := workoutPlanNotFoundSetup(t, ath)
		exercise := exerciseNotFoundSetup(t, ath, wplan)
		weight := random.Weight()

		err := service.EditExerciseWeight(ath, wplan, exercise, weight)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseNotFoundSetup(t, ath, wplan)
		weight := random.Weight()

		err := service.EditExerciseWeight(ath, wplan, exercise, weight)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _, _ := exerciseSuccessSetup(t, ath, wplan, service)
		weight := random.Weight()

		err := service.EditExerciseWeight(ath, wplan, exercise, weight)
		require.NoError(t, err)

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)

		metrics := exercises[0].Metrics()
		require.Equal(t, weight, float64(metrics.Weight()))
	})
}

func TestEditExerciseRestDur(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized WorkoutPlan", func(t *testing.T) {
		wplan := workoutPlanUnauthorizedSetup(t)
		exercise := exerciseUnauthorizedSetup(t, wplan)
		restDur := random.RestTime()

		err := service.EditExerciseRestDur(ath, wplan, exercise, restDur)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When Unauthorized Exercise", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseUnauthorizedSetup(t, wplan)
		restDur := random.RestTime()

		err := service.EditExerciseRestDur(ath, wplan, exercise, restDur)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan := workoutPlanNotFoundSetup(t, ath)
		exercise := exerciseNotFoundSetup(t, ath, wplan)
		restDur := random.RestTime()

		err := service.EditExerciseRestDur(ath, wplan, exercise, restDur)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise := exerciseNotFoundSetup(t, ath, wplan)
		restDur := random.RestTime()

		err := service.EditExerciseRestDur(ath, wplan, exercise, restDur)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _, _ := exerciseSuccessSetup(t, ath, wplan, service)
		restDur := random.RestTime()

		err := service.EditExerciseRestDur(ath, wplan, exercise, restDur)
		require.NoError(t, err)

		exercises, err := service.FetchWorkoutPlanExercises(ath, wplan)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)

		metrics := exercises[0].Metrics()
		require.Equal(t, restDur, float64(metrics.RestDur()))
	})
}

func TestRemoveWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized", func(t *testing.T) {
		wplan := workoutPlanUnauthorizedSetup(t)

		err := service.RemoveWorkoutPlan(ath, wplan)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan := workoutPlanNotFoundSetup(t, ath)

		err := service.RemoveWorkoutPlan(ath, wplan)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)

		err := service.RemoveWorkoutPlan(ath, wplan)
		require.NoError(t, err)

		wplans, err := service.FetchWorkoutPlans(ath)
		require.NoError(t, err)
		require.Empty(t, wplans)
	})
}

func workoutPlanNotFoundSetup(t *testing.T, ath athlete.Athlete) wp.WorkoutPlan {
	title := random.String(75)
	wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), title)
	require.NoError(t, err)

	return wplan
}

func workoutPlanSuccessSetup(t *testing.T, ath athlete.Athlete, service p.PlanningService) (wp.WorkoutPlan, string) {
	title := random.String(75)

	wplan, err := service.CreateNewWorkoutPlan(ath, title)
	require.NoError(t, err)
	return wplan, title
}

func workoutPlanUnauthorizedSetup(t *testing.T) wp.WorkoutPlan {
	title := random.String(75)
	ath := athlete.NewAthlete()
	wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), title)
	require.NoError(t, err)

	return wplan
}

func exerciseSuccessSetup(t *testing.T, ath athlete.Athlete, wplan wp.WorkoutPlan, service p.PlanningService) (e.Exercise, e.Metrics, string) {
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

	return exercise, metrics, name
}

func exerciseNotFoundSetup(t *testing.T, ath athlete.Athlete, wplan wp.WorkoutPlan) e.Exercise {
	name := random.String(75)

	metrics, err := e.NewMetrics(
		random.RepCount(),
		random.NumSets(),
		random.Weight(),
		random.RestTime(),
	)
	require.NoError(t, err)

	exercise, err := e.NewExercise(wplan.ID(), ath.AthleteID(), name, metrics, 0)
	require.NoError(t, err)
	require.NotEmpty(t, exercise)
	require.Equal(t, name, exercise.Name())
	return exercise
}

func exerciseUnauthorizedSetup(t *testing.T, wplan wp.WorkoutPlan) e.Exercise {
	ath := athlete.NewAthlete()
	name := random.String(75)

	metrics, err := e.NewMetrics(
		random.RepCount(),
		random.NumSets(),
		random.Weight(),
		random.RestTime(),
	)
	require.NoError(t, err)

	exercise, err := e.NewExercise(wplan.ID(), ath.AthleteID(), name, metrics, 0)
	require.NoError(t, err)
	require.NotEmpty(t, exercise)
	require.Equal(t, name, exercise.Name())

	return exercise
}

func setup() (p.PlanningService, athlete.Athlete) {
	repo := adapters.NewInMemRepo()
	return p.NewPlanningService(repo), athlete.NewAthlete()
}
