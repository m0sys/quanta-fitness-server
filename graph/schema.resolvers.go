package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/mhd53/quanta-fitness-server/api/auth"
	"github.com/mhd53/quanta-fitness-server/graph/generated"
	"github.com/mhd53/quanta-fitness-server/graph/model"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	"github.com/mhd53/quanta-fitness-server/pkg/format"
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
		log.Panic("GraphQL Error: ", err.Error())
		return &model.Workout{}, errors.New("Internal Error!")
	}

	return &model.Workout{
		ID:        strconv.FormatInt(created.ID, 10),
		Title:     created.Title,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
	}, nil
}

func (r *mutationResolver) UpdateWorkout(ctx context.Context, input model.WorkoutUpdate) (*model.Workout, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return &model.Workout{}, errors.New("Access Denied!")
	}

	id, err := format.ConvertToBase64(input.ID)
	if err != nil {
		log.Panic("GraphQL Error: ", err.Error())
		return &model.Workout{}, errors.New("Internal Error!")

	}

	updates := entity.BaseWorkout{
		Title:    input.Title,
		Username: uname,
	}

	err2 := r.WorkoutServer.UpdateWorkout(id, updates, uname)
	if err2 != nil {
		log.Panic("GraphQL Error: ", err2.Error())
		return &model.Workout{}, errors.New("Internal Error!")

	}

	got, err3 := r.WorkoutServer.GetWorkout(id, uname)
	if err3 != nil {
		log.Panic("GraphQL Error: ", err3.Error())
		return &model.Workout{}, errors.New("Internal Error!")

	}

	return &model.Workout{
		ID:        strconv.FormatInt(got.ID, 10),
		Title:     got.Title,
		CreatedAt: got.CreatedAt,
		UpdatedAt: got.UpdatedAt,
	}, nil
}

func (r *mutationResolver) DeleteWorkout(ctx context.Context, id string) (*model.Workout, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return &model.Workout{}, errors.New("Access Denied!")
	}

	intID, err := format.ConvertToBase64(id)
	if err != nil {
		log.Panic("GraphQL Error: ", err.Error())
		return &model.Workout{}, errors.New("Internal Error!")

	}

	got, err2 := r.WorkoutServer.GetWorkout(intID, uname)
	if err2 != nil {
		log.Panic("GraphQL Error: ", err2.Error())
		return &model.Workout{}, errors.New("Internal Error!")

	}

	err3 := r.WorkoutServer.DeleteWorkout(intID, uname)
	if err3 != nil {
		log.Panic("GraphQL Error: ", err3.Error())
		return &model.Workout{}, errors.New("Internal Error!")

	}

	return &model.Workout{
		ID:        strconv.FormatInt(got.ID, 10),
		Title:     got.Title,
		CreatedAt: got.CreatedAt,
		UpdatedAt: got.UpdatedAt,
	}, nil

}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Workouts(ctx context.Context) ([]*model.Workout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Workout(ctx context.Context, id string) (*model.Workout, error) {
	uname := auth.ForContext(ctx)

	if uname == "" {
		return &model.Workout{}, errors.New("Access Denied!")
	}

	intID, err := format.ConvertToBase64(id)
	if err != nil {
		log.Panic("GraphQL Error: ", err.Error())
		return &model.Workout{}, errors.New("Internal Error!")

	}

	got, err2 := r.WorkoutServer.GetWorkout(intID, uname)
	if err2 != nil {
		log.Panic("GraphQL Error: ", err2.Error())
		return &model.Workout{}, errors.New("Internal Error!")

	}

	return &model.Workout{
		ID:        strconv.FormatInt(got.ID, 10),
		Title:     got.Title,
		CreatedAt: got.CreatedAt,
		UpdatedAt: got.UpdatedAt,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
