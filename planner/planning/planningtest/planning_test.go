package planningtest

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	p "github.com/mhd53/quanta-fitness-server/planner/planning"
	"github.com/mhd53/quanta-fitness-server/planner/planning/adapters"
	"github.com/stretchr/testify/require"
)

func TestCreateNewWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When success", func(t *testing.T) {
		wplan, title := workoutPlanSuccessSetup(t, ath, service)
		require.NotEmpty(t, wplan)

		wplans, err := service.FetchWorkoutPlans(ath.AthleteID())
		require.NoError(t, err)
		require.NotEmpty(t, wplans)
		require.Equal(t, title, wplans[0].Title)
	})

	t.Run("When WorkoutPlan with given title already exists", func(t *testing.T) {
		_, title := workoutPlanSuccessSetup(t, ath, service)
		req := p.CreateNewWorkoutPlanReq{
			AthleteID: ath.AthleteID(),
			Title:     title,
		}

		wplan, err := service.CreateNewWorkoutPlan(req)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialTitle.Error(), err.Error())
		require.Empty(t, wplan)
	})
}

func TestAddNewExerciseToWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		name := random.String(75)

		req := p.AddNewExerciseToWorkoutPlanReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: "1234",
			Name:          name,
			TargetRep:     random.RepCount(),
			NumSets:       random.NumSets(),
			Weight:        random.Weight(),
			RestDur:       random.RestTime(),
		}
		exercise, err := service.AddNewExerciseToWorkoutPlan(req)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
		require.Empty(t, exercise)
	})

	t.Run("When WorkoutPlan doesn't belong to Athlete", func(t *testing.T) {
		name := random.String(75)
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		ath2 := athlete.NewAthlete()

		req := p.AddNewExerciseToWorkoutPlanReq{
			AthleteID:     ath2.AthleteID(),
			WorkoutPlanID: wplan.ID,
			Name:          name,
			TargetRep:     random.RepCount(),
			NumSets:       random.NumSets(),
			Weight:        random.Weight(),
			RestDur:       random.RestTime(),
		}

		exercise, err := service.AddNewExerciseToWorkoutPlan(req)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
		require.Empty(t, exercise)
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		require.NotEmpty(t, exercise)

		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}
		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, exercise.ID, exercises[0].ID)
	})

	t.Run("When Exercise with same name already in WorkoutPlan", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		_, name := exerciseSuccessSetup(t, ath, wplan, service)

		req := p.AddNewExerciseToWorkoutPlanReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			Name:          name,
			TargetRep:     random.RepCount(),
			NumSets:       random.NumSets(),
			Weight:        random.Weight(),
			RestDur:       random.RestTime(),
		}

		exercise, err := service.AddNewExerciseToWorkoutPlan(req)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialName.Error(), err.Error())
		require.Empty(t, exercise)
	})

}

func TestRemoveExerciseFromWorkoutPlan(t *testing.T) {
	service, ath := setup()

	t.Run("When unauthorized to access WorkoutPlan", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		ath2 := athlete.NewAthlete()

		req := p.RemoveExerciseFromWorkoutPlanReq{
			AthleteID:     ath2.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    "1234",
		}

		err := service.RemoveExerciseFromWorkoutPlan(req)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	/*
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
	*/

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)

		req := p.RemoveExerciseFromWorkoutPlanReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: "1234",
			ExerciseID:    exercise.ID,
		}

		err := service.RemoveExerciseFromWorkoutPlan(req)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})
	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)

		req := p.RemoveExerciseFromWorkoutPlanReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    "1234",
		}

		err := service.RemoveExerciseFromWorkoutPlan(req)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)

		req := p.RemoveExerciseFromWorkoutPlanReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
		}

		err := service.RemoveExerciseFromWorkoutPlan(req)
		require.NoError(t, err)

		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}
		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.NoError(t, err)
		require.Empty(t, exercises)
	})
}

func TestEditWorkoutPlanTitle(t *testing.T) {
	service, ath := setup()
	t.Run("When unauthorized", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		title2 := random.String(75)

		ath2 := athlete.NewAthlete()
		req := p.EditWorkoutPlanTitleReq{
			AthleteID:     ath2.AthleteID(),
			WorkoutPlanID: wplan.ID,
			Title:         title2,
		}
		err := service.EditWorkoutPlanTitle(req)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		title2 := random.String(75)

		req := p.EditWorkoutPlanTitleReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: "1234",
			Title:         title2,
		}

		err := service.EditWorkoutPlanTitle(req)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When same title", func(t *testing.T) {
		wplan, title := workoutPlanSuccessSetup(t, ath, service)

		req := p.EditWorkoutPlanTitleReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			Title:         wplan.Title,
		}

		err := service.EditWorkoutPlanTitle(req)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialTitle.Error(), err.Error())

		wplans, err := service.FetchWorkoutPlans(ath.AthleteID())
		require.NoError(t, err)
		require.NotEmpty(t, wplans)
		for _, val := range wplans {
			if val.ID == wplan.ID {
				require.Equal(t, title, val.Title)
			}
		}
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		title2 := random.String(75)

		req := p.EditWorkoutPlanTitleReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			Title:         title2,
		}

		err := service.EditWorkoutPlanTitle(req)
		require.NoError(t, err)

		wplans, err := service.FetchWorkoutPlans(ath.AthleteID())
		require.NoError(t, err)
		require.NotEmpty(t, wplans)
		for _, val := range wplans {
			if val.ID == wplan.ID {
				require.Equal(t, title2, val.Title)
			}
		}
	})
}

func TestFetchWorkoutPlans(t *testing.T) {
	service, ath := setup()
	t.Run("When no WorkoutPlan for Athlete", func(t *testing.T) {
		wplans, err := service.FetchWorkoutPlans(ath.AthleteID())
		require.NoError(t, err)
		require.Empty(t, wplans)
	})

	t.Run("After creating WorkoutPlans for Athlete", func(t *testing.T) {
		n := 5
		for i := 0; i < n; i++ {
			workoutPlanSuccessSetup(t, ath, service)
		}
		wplans, err := service.FetchWorkoutPlans(ath.AthleteID())
		require.NoError(t, err)
		require.NotEmpty(t, wplans)
		require.Equal(t, n, len(wplans))
	})
}

func TestFetchWorkoutPlanExercises(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)

		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath2.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}
		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
		require.Empty(t, exercises)
	})
	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: "1234",
		}

		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
		require.Empty(t, exercises)
	})

	t.Run("When no Exercises for WorkoutPlan", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}

		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.NoError(t, err)
		require.Empty(t, exercises)
	})

	t.Run("After Exercises have been added to WorkoutPlan", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}

		n := 5
		for i := 0; i < n; i++ {
			exerciseSuccessSetup(t, ath, wplan, service)
		}

		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, n, len(exercises))
	})

}

func TestEditExerciseName(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized WorkoutPlan", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		name := random.String(75)

		req := p.EditExerciseNameReq{
			AthleteID:     ath2.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			Name:          name,
		}

		err := service.EditExerciseName(req)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	/*
		t.Run("When Unauthorized Exercise", func(t *testing.T) {
			wplan, _ := workoutPlanSuccessSetup(t, ath, service)
			exercise := exerciseUnauthorizedSetup(t, wplan)
			name := random.String(75)

			err := service.EditExerciseName(ath, wplan, exercise, name)
			require.Error(t, err)
			require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
		})
	*/

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		name := random.String(75)

		req := p.EditExerciseNameReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: "1234",
			ExerciseID:    exercise.ID,
			Name:          name,
		}

		err := service.EditExerciseName(req)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		name := random.String(75)

		req := p.EditExerciseNameReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    "1234",
			Name:          name,
		}

		err := service.EditExerciseName(req)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When same Exercise name", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, name := exerciseSuccessSetup(t, ath, wplan, service)

		req := p.EditExerciseNameReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			Name:          name,
		}

		err := service.EditExerciseName(req)
		require.Error(t, err)
		require.Equal(t, p.ErrIdentialName.Error(), err.Error())

		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}
		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, name, exercises[0].Name)
		require.Equal(t, name, exercise.Name)
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		name := random.String(75)

		req := p.EditExerciseNameReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			Name:          name,
		}

		err := service.EditExerciseName(req)
		require.NoError(t, err)

		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}
		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, name, exercises[0].Name)
	})
}

func TestEditExerciseTargetRep(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized WorkoutPlan", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		targetRep := random.RepCount()

		req := p.EditExerciseTargetRepReq{
			AthleteID:     ath2.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			TargetRep:     targetRep,
		}

		err := service.EditExerciseTargetRep(req)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	/*
		t.Run("When Unauthorized Exercise", func(t *testing.T) {
			wplan, _ := workoutPlanSuccessSetup(t, ath, service)
			exercise := exerciseUnauthorizedSetup(t, wplan)
			name := random.String(75)

			err := service.EditExerciseName(ath, wplan, exercise, name)
			require.Error(t, err)
			require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
		})
	*/

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		targetRep := random.RepCount()

		req := p.EditExerciseTargetRepReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: "1234",
			ExerciseID:    exercise.ID,
			TargetRep:     targetRep,
		}

		err := service.EditExerciseTargetRep(req)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		targetRep := random.RepCount()

		req := p.EditExerciseTargetRepReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    "1234",
			TargetRep:     targetRep,
		}

		err := service.EditExerciseTargetRep(req)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		targetRep := random.RepCount()

		req := p.EditExerciseTargetRepReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			TargetRep:     targetRep,
		}

		err := service.EditExerciseTargetRep(req)
		require.NoError(t, err)

		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}
		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, targetRep, exercises[0].TargetRep)
	})
}

func TestEditExerciseNumSets(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized WorkoutPlan", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		numSets := random.NumSets()

		req := p.EditExerciseNumSetsReq{
			AthleteID:     ath2.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			NumSets:       numSets,
		}

		err := service.EditExerciseNumSets(req)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	/*
		t.Run("When Unauthorized Exercise", func(t *testing.T) {
			wplan, _ := workoutPlanSuccessSetup(t, ath, service)
			exercise := exerciseUnauthorizedSetup(t, wplan)
			name := random.String(75)

			err := service.EditExerciseName(ath, wplan, exercise, name)
			require.Error(t, err)
			require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
		})
	*/

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		numSets := random.NumSets()

		req := p.EditExerciseNumSetsReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: "1234",
			ExerciseID:    exercise.ID,
			NumSets:       numSets,
		}

		err := service.EditExerciseNumSets(req)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		numSets := random.NumSets()

		req := p.EditExerciseNumSetsReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    "1234",
			NumSets:       numSets,
		}

		err := service.EditExerciseNumSets(req)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		numSets := random.NumSets()

		req := p.EditExerciseNumSetsReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			NumSets:       numSets,
		}

		err := service.EditExerciseNumSets(req)
		require.NoError(t, err)

		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}
		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, numSets, exercises[0].NumSets)
	})
}

func TestEditExerciseWeight(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized WorkoutPlan", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		weight := random.Weight()

		req := p.EditExerciseWeightReq{
			AthleteID:     ath2.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			Weight:        weight,
		}

		err := service.EditExerciseWeight(req)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	/*
		t.Run("When Unauthorized Exercise", func(t *testing.T) {
			wplan, _ := workoutPlanSuccessSetup(t, ath, service)
			exercise := exerciseUnauthorizedSetup(t, wplan)
			name := random.String(75)

			err := service.EditExerciseName(ath, wplan, exercise, name)
			require.Error(t, err)
			require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
		})
	*/

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		weight := random.Weight()

		req := p.EditExerciseWeightReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: "1234",
			ExerciseID:    exercise.ID,
			Weight:        weight,
		}

		err := service.EditExerciseWeight(req)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		weight := random.Weight()

		req := p.EditExerciseWeightReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    "1234",
			Weight:        weight,
		}

		err := service.EditExerciseWeight(req)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		weight := random.Weight()

		req := p.EditExerciseWeightReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			Weight:        weight,
		}

		err := service.EditExerciseWeight(req)
		require.NoError(t, err)

		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}
		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, weight, float64(exercises[0].Weight))
	})
}

func TestEditExerciseRestDur(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized WorkoutPlan", func(t *testing.T) {
		ath2 := athlete.NewAthlete()
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		restDur := random.RestTime()

		req := p.EditExerciseRestDurReq{
			AthleteID:     ath2.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			RestDur:       restDur,
		}

		err := service.EditExerciseRestDur(req)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	/*
		t.Run("When Unauthorized Exercise", func(t *testing.T) {
			wplan, _ := workoutPlanSuccessSetup(t, ath, service)
			exercise := exerciseUnauthorizedSetup(t, wplan)
			name := random.String(75)

			err := service.EditExerciseName(ath, wplan, exercise, name)
			require.Error(t, err)
			require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
		})
	*/

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		restDur := random.RestTime()

		req := p.EditExerciseRestDurReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: "1234",
			ExerciseID:    exercise.ID,
			RestDur:       restDur,
		}

		err := service.EditExerciseRestDur(req)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When Exercise not found", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		restDur := random.RestTime()

		req := p.EditExerciseRestDurReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    "1234",
			RestDur:       restDur,
		}

		err := service.EditExerciseRestDur(req)
		require.Error(t, err)
		require.Equal(t, p.ErrExerciseNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		exercise, _ := exerciseSuccessSetup(t, ath, wplan, service)
		restDur := random.RestTime()

		req := p.EditExerciseRestDurReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
			ExerciseID:    exercise.ID,
			RestDur:       restDur,
		}

		err := service.EditExerciseRestDur(req)
		require.NoError(t, err)

		req2 := p.FetchWorkoutPlanExercisesReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}
		exercises, err := service.FetchWorkoutPlanExercises(req2)
		require.NoError(t, err)
		require.NotEmpty(t, exercises)
		require.Equal(t, restDur, float64(exercises[0].RestDur))
	})
}

func TestRemoveWorkoutPlan(t *testing.T) {
	service, ath := setup()
	t.Run("When Unauthorized", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		ath2 := athlete.NewAthlete()

		req := p.RemoveWorkoutPlanReq{
			AthleteID:     ath2.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}

		err := service.RemoveWorkoutPlan(req)
		require.Error(t, err)
		require.Equal(t, p.ErrUnauthorizedAccess.Error(), err.Error())
	})

	t.Run("When WorkoutPlan not found", func(t *testing.T) {
		req := p.RemoveWorkoutPlanReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: "1234",
		}

		err := service.RemoveWorkoutPlan(req)
		require.Error(t, err)
		require.Equal(t, p.ErrWorkoutPlanNotFound.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		wplan, _ := workoutPlanSuccessSetup(t, ath, service)
		req := p.RemoveWorkoutPlanReq{
			AthleteID:     ath.AthleteID(),
			WorkoutPlanID: wplan.ID,
		}

		err := service.RemoveWorkoutPlan(req)
		require.NoError(t, err)

		wplans, err := service.FetchWorkoutPlans(ath.AthleteID())
		require.NoError(t, err)
		for _, val := range wplans {
			require.NotEqual(t, val.ID, wplan.ID)
		}
	})
}

func workoutPlanSuccessSetup(t *testing.T, ath athlete.Athlete, service p.PlanningService) (p.WorkoutPlanRes, string) {
	title := random.String(75)

	req := p.CreateNewWorkoutPlanReq{
		AthleteID: ath.AthleteID(),
		Title:     title,
	}
	res, err := service.CreateNewWorkoutPlan(req)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	return res, title
}

func exerciseSuccessSetup(t *testing.T, ath athlete.Athlete, wplan p.WorkoutPlanRes, service p.PlanningService) (p.ExerciseRes, string) {
	name := random.String(75)

	req := p.AddNewExerciseToWorkoutPlanReq{
		AthleteID:     ath.AthleteID(),
		WorkoutPlanID: wplan.ID,
		Name:          name,
		TargetRep:     random.RepCount(),
		NumSets:       random.NumSets(),
		Weight:        random.Weight(),
		RestDur:       random.RestTime(),
	}

	res, err := service.AddNewExerciseToWorkoutPlan(req)

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, name, res.Name)

	return res, name
}

func setup() (p.PlanningService, athlete.Athlete) {
	repo := adapters.NewInMemRepo()
	return p.NewPlanningService(repo), athlete.NewAthlete()
}
