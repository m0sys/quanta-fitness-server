package workout

import (
	"errors"
	"log"

	"github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type WorkoutService interface {
	CreateWorkout(title string, uname string) (entity.Workout, error)
	UpdateWorkout(workout entity.Workout, uname string) error
	DeleteWorkout(wid int64, uname string) error
	GetWorkoutExercises(wid int64, uname string) ([]entity.Exercise, error)
}

type service struct{}

var (
	ws   workoutstore.WorkoutStore
	auth WorkoutAuth
	val  WorkoutValidator
)

func NewWorkoutService(
	workoutstore workoutstore.WorkoutStore,
	authorizer WorkoutAuth,
	validator WorkoutValidator) WorkoutService {
	ws = workoutstore
	auth = authorizer
	val = validator
	return &service{}
}

func (*service) CreateWorkout(title string, uname string) (entity.Workout, error) {
	authorized, err := auth.AuthorizeCreateWorkout(uname)

	if err != nil {
		log.Panic(err)
		return entity.Workout{}, errors.New("Internal error!")
	}

	if !authorized {
		return entity.Workout{}, errors.New("Access Denied!")

	}

	err2 := val.ValidateCreateWorkout(title)

	if err2 != nil {
		return entity.Workout{}, err2
	}
	workout := entity.BaseWorkout{
		Title:    title,
		Username: uname,
	}
	workoutDS, err3 := ws.Save(workout)

	if err3 != nil {
		log.Fatal(err3)
		return entity.Workout{}, errors.New("Internal error!")
	}

	return workoutDS, nil
}

func (*service) UpdateWorkout(workout entity.Workout, uname string) error {
	return nil

}

func (*service) DeleteWorkout(wid int64, uname string) error {
	return nil

}

func (*service) GetWorkoutExercises(wid int64, uname string) ([]entity.Exercise, error) {
	var exercises []entity.Exercise
	return exercises, nil
}
