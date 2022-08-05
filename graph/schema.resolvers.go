package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/m0sys/quanta-fitness-server/api/auth"
	"github.com/m0sys/quanta-fitness-server/graph/generated"
	"github.com/m0sys/quanta-fitness-server/graph/model"
	"github.com/m0sys/quanta-fitness-server/internal/entity"
)

func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.Auth, error) {
	token, err := r.AuthServer.RegisterNewUser(input.Username, input.Email, input.Password, input.Confirm)
	return &model.Auth{
		Token: token,
	}, err
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.Auth, error) {
	if input.Username != nil {
		token, err := r.AuthServer.LoginWithUname(*input.Username, input.Password)
		return &model.Auth{
			Token: token,
		}, err
	}

	if input.Email != nil {
		token, err := r.AuthServer.LoginWithEmail(*input.Email, input.Password)
		return &model.Auth{
			Token: token,
		}, err

	}
	return &model.Auth{
		Token: "",
	}, errors.New("Error: Must provide username or email to login!")
}

func (r *mutationResolver) CreateWorkout(ctx context.Context, input model.NewWorkout) (*model.Workout, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return &model.Workout{}, errors.New("Access Denied!")
	}

	created, err := r.WorkoutServer.CreateWorkout(input.Title, uname)

	if err != nil {
		return &model.Workout{}, err
	}

	return &model.Workout{
		ID:        strconv.FormatInt(created.ID, 10),
		Title:     created.Title,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
	}, nil
}

func (r *mutationResolver) UpdateWorkout(ctx context.Context, input model.WorkoutUpdate) (bool, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return false, errors.New("Access Denied!")
	}

	updates := entity.BaseWorkout{
		Title:    input.Title,
		Username: uname,
	}

	success, err := r.WorkoutServer.UpdateWorkout(input.ID, updates, uname)
	if err != nil {
		return false, err

	}

	return success, nil
}

func (r *mutationResolver) DeleteWorkout(ctx context.Context, id string) (bool, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return false, errors.New("Access Denied!")
	}

	success, err := r.WorkoutServer.DeleteWorkout(id, uname)
	if err != nil {
		return false, err

	}

	return success, nil
}

func (r *mutationResolver) AddExerciseToWorkout(ctx context.Context, input model.NewExercise) (*model.Exercise, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return &model.Exercise{}, errors.New("Access Denied!")
	}

	created, err := r.ExerciseServer.AddExerciseToWorkout(input.Name, uname, input.Wid)

	if err != nil {
		return &model.Exercise{}, err
	}
	return &model.Exercise{
		ID:        strconv.FormatInt(created.ID, 10),
		Wid:       strconv.FormatInt(created.WID, 10),
		Name:      created.Name,
		Weight:    float64(created.Metrics.Weight),
		TargetRep: created.Metrics.TargetRep,
		RestTime:  float64(created.Metrics.RestTime),
		NumSets:   created.Metrics.NumSets,
	}, nil
}

func (r *mutationResolver) UpdateExercise(ctx context.Context, input model.ExerciseUpdate) (bool, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return false, errors.New("Access Denied!")
	}

	updates := entity.ExerciseUpdate{
		Name: input.Name,
		Metrics: entity.Metrics{
			Weight:    float32(input.Weight),
			TargetRep: input.TargetRep,
			RestTime:  float32(input.RestTime),
			NumSets:   input.NumSets,
		},
	}

	success, err := r.ExerciseServer.UpdateExercise(input.ID, uname, updates)
	if err != nil {
		return false, err
	}

	return success, nil
}

func (r *mutationResolver) DeleteExercise(ctx context.Context, id string) (bool, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return false, errors.New("Access Denied!")
	}

	success, err := r.ExerciseServer.DeleteExercise(id, uname)
	if err != nil {
		return false, err

	}

	return success, nil
}

func (r *mutationResolver) AddEsetToExercise(ctx context.Context, input model.NewEset) (*model.Eset, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return &model.Eset{}, errors.New("Access Denied!")
	}

	created, err := r.EsetServer.AddEsetToExercise(uname, input.Eid, input.ActualRepCount, float32(input.Duration), float32(input.RestTimeDuration))

	if err != nil {
		return &model.Eset{}, err
	}
	return &model.Eset{
		ID:               strconv.FormatInt(created.ID, 10),
		Eid:              strconv.FormatInt(created.EID, 10),
		ActualRepCount:   created.SMetric.ActualRepCount,
		Duration:         float64(created.SMetric.Duration),
		RestTimeDuration: float64(created.SMetric.RestTimeDuration),
	}, nil
}

func (r *mutationResolver) UpdateEset(ctx context.Context, input model.EsetUpdate) (bool, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return false, errors.New("Access Denied!")
	}

	updates := entity.EsetUpdate{
		SMetric: entity.SMetric{
			ActualRepCount:   input.ActualRepCount,
			RestTimeDuration: float32(input.RestTimeDuration),
			Duration:         float32(input.Duration),
		},
	}

	success, err := r.EsetServer.UpdateEset(input.ID, uname, updates)
	if err != nil {
		return false, err
	}

	return success, nil
}

func (r *mutationResolver) DeleteEset(ctx context.Context, id string) (bool, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return false, errors.New("Access Denied!")
	}

	success, err := r.EsetServer.DeleteEset(id, uname)
	if err != nil {
		return false, err

	}

	return success, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.PublicUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Workouts(ctx context.Context, username string) ([]*model.Workout, error) {
	var modelWorkouts []*model.Workout

	uname := auth.ForContext(ctx)

	if uname == "" {
		return modelWorkouts, errors.New("Access Denied!")
	}

	got, err := r.WorkoutServer.GetWorkouts(uname)
	if err != nil {
		return modelWorkouts, err
	}

	for _, v := range got {
		mWorkout := &model.Workout{
			ID:        strconv.FormatInt(v.ID, 10),
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt}

		modelWorkouts = append(modelWorkouts, mWorkout)
	}

	return modelWorkouts, nil
}

func (r *queryResolver) Workout(ctx context.Context, id string) (*model.Workout, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return &model.Workout{}, errors.New("Access Denied!")
	}

	got, err := r.WorkoutServer.GetWorkout(id, uname)
	if err != nil {
		return &model.Workout{}, err

	}

	return &model.Workout{
		ID:        strconv.FormatInt(got.ID, 10),
		Title:     got.Title,
		CreatedAt: got.CreatedAt,
		UpdatedAt: got.UpdatedAt,
	}, nil
}

func (r *queryResolver) Exercise(ctx context.Context, id string) (*model.Exercise, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return &model.Exercise{}, errors.New("Access Denied!")
	}

	got, err := r.ExerciseServer.GetExercise(id, uname)
	if err != nil {
		return &model.Exercise{}, err

	}

	return &model.Exercise{
		ID:        strconv.FormatInt(got.ID, 10),
		Wid:       strconv.FormatInt(got.WID, 10),
		Name:      got.Name,
		Weight:    float64(got.Metrics.Weight),
		TargetRep: got.Metrics.TargetRep,
		RestTime:  float64(got.Metrics.RestTime),
		NumSets:   got.Metrics.NumSets,
		CreatedAt: got.CreatedAt,
		UpdatedAt: got.UpdatedAt,
	}, nil
}

func (r *queryResolver) Exercises(ctx context.Context, wid string) ([]*model.Exercise, error) {
	var modelExercises []*model.Exercise

	uname := auth.ForContext(ctx)

	if uname == "" {
		return modelExercises, errors.New("Access Denied!")
	}

	got, err := r.ExerciseServer.GetExercisesForWorkout(wid, uname)
	if err != nil {
		return modelExercises, err
	}

	for _, v := range got {
		mExercise := &model.Exercise{
			ID:        strconv.FormatInt(v.ID, 10),
			Wid:       strconv.FormatInt(v.WID, 10),
			Name:      v.Name,
			Weight:    float64(v.Metrics.Weight),
			TargetRep: v.Metrics.TargetRep,
			RestTime:  float64(v.Metrics.RestTime),
			NumSets:   v.Metrics.NumSets,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt}

		modelExercises = append(modelExercises, mExercise)
	}

	return modelExercises, nil
}

func (r *queryResolver) Eset(ctx context.Context, id string) (*model.Eset, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return &model.Eset{}, errors.New("Access Denied!")
	}

	got, err := r.EsetServer.GetEset(id, uname)
	if err != nil {
		return &model.Eset{}, err

	}

	return &model.Eset{
		ID:               strconv.FormatInt(got.ID, 10),
		Eid:              strconv.FormatInt(got.EID, 10),
		ActualRepCount:   got.SMetric.ActualRepCount,
		RestTimeDuration: float64(got.BaseEset.SMetric.RestTimeDuration),
		Duration:         float64(got.BaseEset.SMetric.Duration),
		CreatedAt:        got.CreatedAt,
		UpdatedAt:        got.UpdatedAt,
	}, nil
}

func (r *queryResolver) Esets(ctx context.Context, eid string) ([]*model.Eset, error) {
	var modelEsets []*model.Eset

	uname := auth.ForContext(ctx)

	if uname == "" {
		return modelEsets, errors.New("Access Denied!")
	}

	got, err := r.EsetServer.GetEsetsForExercise(eid, uname)
	if err != nil {
		return modelEsets, err
	}

	for _, v := range got {
		mEset := &model.Eset{
			ID:               strconv.FormatInt(v.ID, 10),
			Eid:              strconv.FormatInt(v.EID, 10),
			ActualRepCount:   v.BaseEset.SMetric.ActualRepCount,
			RestTimeDuration: float64(v.BaseEset.SMetric.RestTimeDuration),
			Duration:         float64(v.BaseEset.SMetric.Duration),
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		}

		modelEsets = append(modelEsets, mEset)
	}

	return modelEsets, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
