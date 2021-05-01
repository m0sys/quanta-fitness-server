package planningtest

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/mhd53/quanta-fitness-server/planner/adapters"
	"github.com/mhd53/quanta-fitness-server/planner/exercise"
	p "github.com/mhd53/quanta-fitness-server/planner/planning"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
	"github.com/stretchr/testify/require"
)

func TestCreateNewWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When success", func(t *testing.T) {
		title := random.String(75)
		wplan, err := wp.NewWorkoutPlan(title)
		require.NoError(t, err)

		err = service.CreateNewWorkoutPlan(ath, wplan)
		require.NoError(t, err)

		// TODO: Check that WorkoutPlan is stored.
	})

	t.Run("When WorkoutPlan with given title already exists", func(t *testing.T) {
		title := random.String(75)
		wplan, err := wp.NewWorkoutPlan(title)
		require.NoError(t, err)

		err = service.CreateNewWorkoutPlan(ath, wplan)
		require.NoError(t, err)

		wplan2, err := wp.NewWorkoutPlan(title)
		require.NoError(t, err)
		err = service.CreateNewWorkoutPlan(ath, wplan2)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialTitle.Error(), err.Error())

	})

	t.Run("When WorkoutPlan already exists", func(t *testing.T) {
		title := random.String(75)
		wplan, err := wp.NewWorkoutPlan(title)
		require.NoError(t, err)

		err = service.CreateNewWorkoutPlan(ath, wplan)
		require.NoError(t, err)

		err = service.CreateNewWorkoutPlan(ath, wplan)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanAlreadyExists.Error(), err.Error())
	})
}

func TestAddNewExerciseToWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		title := random.String(75)
		wplan, err := wp.NewWorkoutPlan(title)
		require.NoError(t, err)

		metrics, err := exercise.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		require.NoError(t, err)

		exercise, err := exercise.NewExercise(random.String(75), metrics)
		require.NoError(t, err)

		err = service.AddNewExerciseToWorkoutPlan(
			ath,
			wplan,
			exercise,
		)

		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When WorkoutPlan doesn't belong to Athlete", func(t *testing.T) {
		title := random.String(75)
		wplan, err := wp.NewWorkoutPlan(title)
		require.NoError(t, err)

		err = service.CreateNewWorkoutPlan(ath, wplan)
		require.NoError(t, err)

		ath2 := athlete.NewAthlete()

		metrics, err := exercise.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		require.NoError(t, err)

		exercise, err := exercise.NewExercise(random.String(75), metrics)
		require.NoError(t, err)

		err = service.AddNewExerciseToWorkoutPlan(
			ath2,
			wplan,
			exercise,
		)

		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		title := random.String(75)
		wplan, err := wp.NewWorkoutPlan(title)
		require.NoError(t, err)

		err = service.CreateNewWorkoutPlan(ath, wplan)
		require.NoError(t, err)

		metrics, err := exercise.NewMetrics(
			random.RepCount(),
			random.NumSets(),
			random.Weight(),
			random.RestTime(),
		)
		require.NoError(t, err)

		exercise, err := exercise.NewExercise(random.String(75), metrics)
		require.NoError(t, err)

		err = service.AddNewExerciseToWorkoutPlan(
			ath,
			wplan,
			exercise,
		)

		require.NoError(t, err)

		// TODO: Check that exercise is stored in WorkoutPlan
	})

}

func setup() (p.PlanningService, athlete.Athlete) {
	repo := adapters.NewInMemRepo()
	return p.NewPlanningService(repo), athlete.NewAthlete()
}
