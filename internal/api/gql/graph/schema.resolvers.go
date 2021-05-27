package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mhd53/quanta-fitness-server/internal/api/gql/graph/generated"
	"github.com/mhd53/quanta-fitness-server/internal/api/gql/graph/model"
	p "github.com/mhd53/quanta-fitness-server/planner/planning"
)

func (r *mutationResolver) CreateWorkoutPlan(ctx context.Context, input model.NewWorkoutPlan) (*model.WorkoutPlan, error) {
	req := p.CreateNewWorkoutPlanReq{
		AthleteID: "1234",
		Title:     input.Title,
	}
	res, err := r.planning.CreateNewWorkoutPlan(req)
	if err != nil {
		return &model.WorkoutPlan{}, err
	}

	return &model.WorkoutPlan{
		ID:    res.ID,
		Title: res.Title,
	}, nil
}

func (r *mutationResolver) EditWorkoutPlanTitle(ctx context.Context, input model.EditWorkoutPlanTitle) (bool, error) {
	req := p.EditWorkoutPlanTitleReq{
		AthleteID:     "1234",
		WorkoutPlanID: input.ID,
		Title:         input.Title,
	}

	err := r.planning.EditWorkoutPlanTitle(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) AddExercise(ctx context.Context, input model.NewExercise) (*model.Exercise, error) {
	req := p.AddNewExerciseToWorkoutPlanReq{
		AthleteID:     "1234",
		WorkoutPlanID: input.Wpid,
		Name:          input.Name,
		TargetRep:     input.TargetRep,
		NumSets:       input.NumSets,
		Weight:        input.Weight,
		RestDur:       input.RestDur,
	}

	res, err := r.planning.AddNewExerciseToWorkoutPlan(req)
	if err != nil {
		return &model.Exercise{}, err
	}

	return &model.Exercise{
		ID:        res.ID,
		Wpid:      res.WorkoutPlanID,
		Name:      res.Name,
		TargetRep: res.TargetRep,
		NumSets:   res.NumSets,
		Weight:    float64(res.Weight),
		RestDur:   float64(res.RestDur),
	}, nil
}

func (r *mutationResolver) RemoveExercise(ctx context.Context, id string, wpid string) (bool, error) {
	req := p.RemoveExerciseFromWorkoutPlanReq{
		AthleteID:     "1234",
		WorkoutPlanID: wpid,
		ExerciseID:    id,
	}

	err := r.planning.RemoveExerciseFromWorkoutPlan(req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteWorkoutPlan(ctx context.Context, id string) (bool, error) {
	req := p.RemoveWorkoutPlanReq{
		AthleteID:     "1234",
		WorkoutPlanID: id,
	}

	err := r.planning.RemoveWorkoutPlan(req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) Hello(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) WorkoutPlans(ctx context.Context) ([]*model.WorkoutPlan, error) {
	var wplans []*model.WorkoutPlan

	res, err := r.planning.FetchWorkoutPlans("1234")
	if err != nil {
		return wplans, err
	}

	for _, wplan := range res {
		modelPlan := &model.WorkoutPlan{
			ID:    wplan.ID,
			Title: wplan.Title,
		}

		wplans = append(wplans, modelPlan)
	}
	return wplans, nil
}

func (r *queryResolver) Exercises(ctx context.Context, wpid string) ([]*model.Exercise, error) {
	var exercises []*model.Exercise

	req := p.FetchWorkoutPlanExercisesReq{
		AthleteID:     "1234",
		WorkoutPlanID: wpid,
	}

	res, err := r.planning.FetchWorkoutPlanExercises(req)
	if err != nil {
		return exercises, err
	}

	for _, e := range res {
		modelExercise := &model.Exercise{
			ID:        e.ID,
			Wpid:      e.WorkoutPlanID,
			Name:      e.Name,
			TargetRep: e.TargetRep,
			NumSets:   e.NumSets,
			Weight:    float64(e.Weight),
			RestDur:   float64(e.RestDur),
		}

		exercises = append(exercises, modelExercise)
	}

	return exercises, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
