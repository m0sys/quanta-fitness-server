package planningtest

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/mhd53/quanta-fitness-server/planner/adapters"
	e "github.com/mhd53/quanta-fitness-server/planner/exercise"
	p "github.com/mhd53/quanta-fitness-server/planner/planning"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
	"github.com/stretchr/testify/require"
)

func TestCreateNewWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When success", func(t *testing.T) {
		title := random.String(75)

		wplan, err := service.CreateNewWorkoutPlan(ath, title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)
		require.Equal(t, title, wplan.Title())

		// TODO: Check that WorkoutPlan is stored.
	})

	t.Run("When WorkoutPlan with given title already exists", func(t *testing.T) {
		title := random.String(75)

		wplan, err := service.CreateNewWorkoutPlan(ath, title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)
		require.Equal(t, title, wplan.Title())

		wplan, err = service.CreateNewWorkoutPlan(ath, title)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialTitle.Error(), err.Error())
		require.Empty(t, wplan)
	})
}

func TestAddNewExerciseToWorkoutPlan(t *testing.T) {
	t.Skip()
	service, ath := setup()
	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		title := random.String(75)
		name := random.String(75)
		wplan, err := wp.NewWorkoutPlan("1", title)
		require.NoError(t, err)

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
		title := random.String(75)
		name := random.String(75)

		wplan, err := service.CreateNewWorkoutPlan(ath, title)
		require.NoError(t, err)

		ath2 := athlete.NewAthlete()

		metrics, err := e.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		require.NoError(t, err)

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
		title := random.String(75)
		name := random.String(75)

		wplan, err := service.CreateNewWorkoutPlan(ath, title)
		require.NoError(t, err)

		metrics, err := e.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		require.NoError(t, err)

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

		// TODO: Check that exercise is stored in WorkoutPlan
	})

	t.Run("When Exercise with same name already in WorkoutPlan", func(t *testing.T) {
		title := random.String(75)
		name := random.String(75)

		wplan, err := service.CreateNewWorkoutPlan(ath, title)
		require.NoError(t, err)

		metrics, err := e.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		require.NoError(t, err)

		require.NoError(t, err)

		exercise, err := service.AddNewExerciseToWorkoutPlan(
			ath,
			wplan,
			name,
			metrics,
		)

		require.NoError(t, err)
		require.NotEmpty(t, exercise)

		exercise, err = service.AddNewExerciseToWorkoutPlan(
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

/*
func TestRemoveExerciseFromWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		title := random.String(75)
		wplan, err := wp.NewWorkoutPlan(title)
		require.NoError(t, err)

		metrics, err := e.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		require.NoError(t, err)

		exercise, err := e.NewExercise(random.String(75), metrics)
		require.NoError(t, err)

		err = service.AddNewExerciseToWorkoutPlan(
			ath,
			wplan,
			exercise,
		)

		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})
}
*/

func setup() (p.PlanningService, athlete.Athlete) {
	repo := adapters.NewInMemRepo()
	return p.NewPlanningService(repo), athlete.NewAthlete()
}
