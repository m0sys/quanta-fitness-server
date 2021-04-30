package planningtest

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/mhd53/quanta-fitness-server/planner/adapters"
	p "github.com/mhd53/quanta-fitness-server/planner/planning"
	"github.com/stretchr/testify/require"
)

func TestCreateNewWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When title is more than 75 chars", func(t *testing.T) {
		err := service.CreateNewWorkoutPlan(ath, random.String(76))
		require.Error(t, err)
	})

	t.Run("When success", func(t *testing.T) {
		err := service.CreateNewWorkoutPlan(ath, random.String(75))
		require.NoError(t, err)
	})

	t.Run("When WorkoutPlan with given title already exists", func(t *testing.T) {
		gen := random.String(75)
		err := service.CreateNewWorkoutPlan(ath, gen)
		require.NoError(t, err)

		err = service.CreateNewWorkoutPlan(ath, gen)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialTitle.Error(), err.Error())

	})
}

func setup() (p.PlanningService, athlete.Athlete) {
	repo := adapters.NewInMemRepo()
	return p.NewPlanningService(repo), athlete.NewAthlete()
}
