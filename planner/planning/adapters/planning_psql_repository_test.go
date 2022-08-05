package adapters

import (
	"testing"

	"github.com/m0sys/quanta-fitness-server/internal/random"
	"github.com/m0sys/quanta-fitness-server/manager/athlete"
	e "github.com/m0sys/quanta-fitness-server/planner/exercise"
	wp "github.com/m0sys/quanta-fitness-server/planner/workoutplan"
	"github.com/stretchr/testify/require"
)

const (
	// FIXME: replace this with a dynamic aid when manger infra is implemented.
	aid = "59be62ed-4c63-4d25-9327-cd29664a1b71"
)

func TestStoreWorkoutPlan_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), random.String(10))
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)
	})
}

func TestFindWorkoutPlanByTitleAndAthleteID_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When title not found", func(t *testing.T) {
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), random.String(10))
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		found, err := psqlRepo.FindWorkoutPlanByTitleAndAthleteID(wplan, ath)
		require.NoError(t, err)
		require.False(t, found)
	})

	t.Run("When Athlete not found", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wplan, err := wp.NewWorkoutPlan(ath2.AthleteID(), random.String(10))
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		found, err := psqlRepo.FindWorkoutPlanByTitleAndAthleteID(wplan, ath)
		require.NoError(t, err)
		require.False(t, found)
	})

	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		found, err := psqlRepo.FindWorkoutPlanByTitleAndAthleteID(wplan, ath)
		require.NoError(t, err)
		require.True(t, found)
	})
}

func TestFindWorkoutPlanByID_psql(t *testing.T) {
	psqlRepo, ath := setup()

	t.Run("When ID not found", func(t *testing.T) {
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), random.String(10))
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		found, err := psqlRepo.FindWorkoutPlanByID(wplan)
		require.NoError(t, err)
		require.False(t, found)
	})

	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		found, err := psqlRepo.FindWorkoutPlanByID(wplan)
		require.NoError(t, err)
		require.True(t, found)
	})
}

func TestFindWorkoutPlanByIDAndAthleteID_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When ID not found", func(t *testing.T) {
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), random.String(10))
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		found, err := psqlRepo.FindWorkoutPlanByIDAndAthleteID(wplan, ath)
		require.NoError(t, err)
		require.False(t, found)
	})

	t.Run("When Athlete not found", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wplan, err := wp.NewWorkoutPlan(ath2.AthleteID(), random.String(10))
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		found, err := psqlRepo.FindWorkoutPlanByIDAndAthleteID(wplan, ath)
		require.NoError(t, err)
		require.False(t, found)
	})

	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		found, err := psqlRepo.FindWorkoutPlanByIDAndAthleteID(wplan, ath)
		require.NoError(t, err)
		require.True(t, found)
	})
}

func TestStoreExercise_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), random.String(10))
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		// Exercise setup.
		exercise := exerciseRawSetup(t, ath, wplan)

		err = psqlRepo.StoreExercise(wplan, exercise, ath)
		require.NoError(t, err)
	})
}

func TestFindExerciseByID_psql(t *testing.T) {
	psqlRepo, ath := setup()

	t.Run("When ID not found", func(t *testing.T) {
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), random.String(10))
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		exercise := exerciseRawSetup(t, ath, wplan)

		found, err := psqlRepo.FindExerciseByID(exercise)
		require.NoError(t, err)
		require.False(t, found)
	})

	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		exercise := exerciseRawSetup(t, ath, wplan)
		err = psqlRepo.StoreExercise(wplan, exercise, ath)
		require.NoError(t, err)

		found, err := psqlRepo.FindExerciseByID(exercise)
		require.NoError(t, err)
		require.True(t, found)
	})
}

func TestFindExerciseByNameAndWorkoutPlanID_psql(t *testing.T) {
	psqlRepo, ath := setup()

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), random.String(10))
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		exercise := exerciseRawSetup(t, ath, wplan)

		found, err := psqlRepo.FindExerciseByNameAndWorkoutPlanID(wplan, exercise)
		require.NoError(t, err)
		require.False(t, found)
	})

	t.Run("When Name not found", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(aid, title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		exercise := exerciseRawSetup(t, ath, wplan)

		found, err := psqlRepo.FindExerciseByNameAndWorkoutPlanID(wplan, exercise)
		require.NoError(t, err)
		require.False(t, found)
	})

	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(aid, title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		exercise := exerciseRawSetup(t, ath, wplan)
		err = psqlRepo.StoreExercise(wplan, exercise, ath)
		require.NoError(t, err)

		found, err := psqlRepo.FindExerciseByNameAndWorkoutPlanID(wplan, exercise)
		require.NoError(t, err)
		require.True(t, found)
	})
}

func TestRemoveExercise_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(aid, title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		exercise := exerciseRawSetup(t, ath, wplan)
		err = psqlRepo.StoreExercise(wplan, exercise, ath)
		require.NoError(t, err)

		err = psqlRepo.RemoveExercise(exercise)
		require.NoError(t, err)

		found, err := psqlRepo.FindExerciseByID(exercise)
		require.NoError(t, err)
		require.False(t, found)
	})
}

func TestRemoveWorkoutPlan_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(aid, title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		err = psqlRepo.RemoveWorkoutPlan(wplan)
		require.NoError(t, err)

		found, err := psqlRepo.FindWorkoutPlanByID(wplan)
		require.NoError(t, err)
		require.False(t, found)
	})
}

func TestUpdateWorkoutPlan_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(aid, title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		newTitle := random.String(10)
		err = wplan.EditTitle(newTitle)
		err = psqlRepo.UpdateWorkoutPlan(wplan)
		require.NoError(t, err)

		// TODO: Check that its been updated.
	})
}

func TestFindAllWorkoutPlansForAthlete_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(aid, title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		ath2 := athlete.RestoreAthlete(aid)
		wplans, err := psqlRepo.FindAllWorkoutPlansForAthlete(ath2)
		require.NoError(t, err)
		require.NotEmpty(t, wplans)
	})
}

func TestFindAllExercisesForWorkoutPlan_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(aid, title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		exercise := exerciseRawSetup(t, ath, wplan)
		err = psqlRepo.StoreExercise(wplan, exercise, ath)
		require.NoError(t, err)

		exercises, err := psqlRepo.FindAllExercisesForWorkoutPlan(wplan)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
	})
}

func TestUpdateExercise_psql(t *testing.T) {
	psqlRepo, ath := setup()
	t.Run("When success", func(t *testing.T) {
		title := random.String(10)
		wplan, err := wp.NewWorkoutPlan(aid, title)
		require.NoError(t, err)
		require.NotEmpty(t, wplan)

		err = psqlRepo.StoreWorkoutPlan(wplan, ath)
		require.NoError(t, err)

		exercise := exerciseRawSetup(t, ath, wplan)
		err = psqlRepo.StoreExercise(wplan, exercise, ath)
		require.NoError(t, err)

		name := random.String(10)
		err = exercise.EditName(name)
		require.NoError(t, err)

		err = psqlRepo.UpdateExercise(exercise)
		require.NoError(t, err)

		exercises, err := psqlRepo.FindAllExercisesForWorkoutPlan(wplan)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, name, exercises[0].Name())
	})
}

// Helper funcs.
func setup() (PsqlPlanningRepository, athlete.Athlete) {
	ath := athlete.RestoreAthlete(aid)
	return NewPSQLRepo(testStore), ath
}

func exerciseRawSetup(t *testing.T, ath athlete.Athlete, wplan wp.WorkoutPlan) e.Exercise {
	name := random.String(10)
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

	return exercise
}
