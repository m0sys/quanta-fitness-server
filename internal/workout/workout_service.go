package workout

import (
	"errors"
	"log"

	"github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type WorkoutService interface {
	CreateWorkout(title string, uname string) (entity.Workout, error)
	UpdateWorkout(wid int64, workout entity.BaseWorkout, uname string) error
	DeleteWorkout(wid int64, uname string) error
	GetWorkout(wid int64, uname string) (entity.Workout, error)
	GetWorkoutsForUser(uname string) ([]entity.Workout, error)
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

func (*service) UpdateWorkout(wid int64, workout entity.BaseWorkout, uname string) error {
	authorized, err := auth.AuthorizeAccessWorkout(uname, wid)

	if err != nil {
		return err
	}

	if !authorized {
		return errors.New("Access Denied!")
	}

	err2 := val.ValidateUpdateWorkout(workout)

	if err2 != nil {
		return err2
	}

	err3 := ws.Update(wid, workout)

	if err3 != nil {
		log.Panic(err3)
		return errors.New("Internal Error!")
	}

	return nil

}

func (*service) DeleteWorkout(wid int64, uname string) error {
	return nil

}

func (*service) GetWorkout(wid int64, uname string) (entity.Workout, error) {
	authorized, err := auth.AuthorizeAccessWorkout(uname, wid)
	if err != nil {
		return entity.Workout{}, err
	}

	if !authorized {
		return entity.Workout{}, errors.New("Access Denied!")
	}

	got, _, err2 := ws.FindWorkoutById(wid)

	if err2 != nil {
		log.Panic(err2)
		return entity.Workout{}, errors.New("Internal Error!")
	}

	return got, nil
}

func (*service) GetWorkoutsForUser(uname string) ([]entity.Workout, error) {
	var workouts []entity.Workout
	return workouts, nil
}

func (*service) GetWorkoutExercises(wid int64, uname string) ([]entity.Exercise, error) {
	var exercises []entity.Exercise
	return exercises, nil
}
