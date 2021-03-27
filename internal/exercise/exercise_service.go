package exercise

import (
	"errors"
	"log"

	store "github.com/mhd53/quanta-fitness-server/internal/datastore/exercisestore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type ExerciseService interface {
	AddExerciseToWorkout(name, uname string, wid int64) (entity.Exercise, error)
	UpdateExercise(eid int64, uname string, updates entity.ExerciseUpdate) error
	DeleteExercise(eid int64, uname string) error
	GetExercise(eid int64, uname string) (entity.Exercise, error)
	GetExercisesForUser(uname string) ([]entity.Exercise, error)
}

type service struct{}

var (
	ses  store.ExerciseStore
	auth ExerciseAuth
	val  ExerciseValidator
)

func NewExerciseService(
	estore store.ExerciseStore,
	authorizer ExerciseAuth,
	validator ExerciseValidator) ExerciseService {
	ses = estore
	auth = authorizer
	val = validator
	return &service{}
}

func (*service) AddExerciseToWorkout(name, uname string, wid int64) (entity.Exercise, error) {
	ok, err := auth.AuthorizeWorkoutAccess(uname, wid)

	if err != nil {
		log.Fatal(err)
		return entity.Exercise{}, errors.New("Internal Error!")
	}
	if !ok {
		return entity.Exercise{}, errors.New("Access Denied!")
	}

	err2 := val.ValidateCreateExercise(name)

	if err2 != nil {
		return entity.Exercise{}, err2
	}

	created, err3 := ses.Save(entity.BaseExercise{
		Name:     name,
		WID:      wid,
		Username: uname,
	})

	if err3 != nil {
		return entity.Exercise{}, errors.New("Internal Error!")
	}

	return created, nil
}

func (*service) UpdateExercise(eid int64, uname string, updates entity.ExerciseUpdate) error {
	ok, err := auth.AuthorizeExerciseAccess(uname, eid)

	if err != nil {
		log.Fatal(err)
		return errors.New("Internal Error!")
	}

	if !ok {
		return errors.New("Access Denied!")
	}

	err2 := val.ValidateUpdateExercise(updates)

	if err2 != nil {
		return err2
	}

	err3 := ses.Update(eid, updates)

	if err3 != nil {
		return err3
	}

	return nil
}

func (*service) DeleteExercise(eid int64, uname string) error {
	return nil
}

func (*service) GetExercise(eid int64, uname string) (entity.Exercise, error) {
	return entity.Exercise{}, nil
}

func (*service) GetExercisesForUser(uname string) ([]entity.Exercise, error) {
	var exercises []entity.Exercise
	return exercises, nil
}
