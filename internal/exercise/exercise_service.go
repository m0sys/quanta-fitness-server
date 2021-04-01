package exercise

import (
	"errors"
	"fmt"
	"log"

	store "github.com/mhd53/quanta-fitness-server/internal/datastore/exercisestore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type ExerciseService interface {
	AddExerciseToWorkout(name, uname, wid string) (entity.Exercise, error)
	UpdateExercise(eid, uname string, updates entity.ExerciseUpdate) error
	DeleteExercise(eid, uname string) error
	GetExercise(eid, uname string) (entity.Exercise, error)
	GetExercisesForWorkout(wid, uname string) ([]entity.Exercise, error)
	GetExercisesForUser(uname string) ([]entity.Exercise, error)
}

type service struct{}

var (
	ses         store.ExerciseStore
	auth        ExerciseAuth
	val         ExerciseValidator
	deniedErr   = errors.New("Access Denied!")
	internalErr = errors.New("Internal Error!")
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

func (*service) AddExerciseToWorkout(name, uname, wid string) (entity.Exercise, error) {
	ok, err := auth.AuthorizeWorkoutAccess(uname, wid)

	if err != nil {
		logErr(err)
		return entity.Exercise{}, internalErr
	}
	if !ok {
		return entity.Exercise{}, deniedErr
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
		return entity.Exercise{}, formatErr(err3)
	}

	return created, nil
}

func (*service) UpdateExercise(eid, uname string, updates entity.ExerciseUpdate) error {
	ok, err := auth.AuthorizeExerciseAccess(uname, eid)

	if err != nil {
		logErr(err)
		return internalErr
	}

	if !ok {
		return deniedErr
	}

	err2 := val.ValidateUpdateExercise(updates)

	if err2 != nil {
		return err2
	}

	err3 := ses.Update(eid, updates)

	if err3 != nil {
		return formatErr(err3)
	}

	return nil
}

func (*service) DeleteExercise(eid, uname string) error {
	authorized, err := auth.AuthorizeExerciseAccess(uname, eid)

	if err != nil {
		logErr(err)
		return internalErr
	}

	if !authorized {
		return deniedErr
	}

	err2 := ses.Delete(eid)

	if err2 != nil {
		return formatErr(err2)
	}

	return nil
}

func (*service) GetExercise(eid, uname string) (entity.Exercise, error) {
	authorized, err := auth.AuthorizeExerciseAccess(uname, eid)

	if err != nil {
		logErr(err)
		return entity.Exercise{}, internalErr
	}

	if !authorized {
		return entity.Exercise{}, deniedErr
	}

	got, _, err2 := ses.FindExerciseById(eid)

	if err2 != nil {
		return entity.Exercise{}, formatErr(err2)
	}

	return got, nil

}

func (*service) GetExercisesForWorkout(wid, uname string) ([]entity.Exercise, error) {
	var exercises []entity.Exercise

	authorized, err := auth.AuthorizeWorkoutAccess(uname, wid)

	if err != nil {
		logErr(err)
		return exercises, internalErr
	}

	if !authorized {
		return exercises, deniedErr
	}

	exercises, err2 := ses.FindAllExercisesByWID(wid)

	if err2 != nil {
		return exercises, formatErr(err2)
	}

	return exercises, nil
}

// FIXME: have to implement `AuthorizeReadAccess` first!
func (*service) GetExercisesForUser(uname string) ([]entity.Exercise, error) {
	var exercises []entity.Exercise

	authorized, err := auth.AuthorizeReadAccess(uname)

	if err != nil {
		logErr(err)
		return exercises, internalErr
	}

	if !authorized {
		return exercises, deniedErr
	}

	return exercises, nil
}

func logErr(err error) {
	log.Printf("%s: %s", "exercise_service", err.Error())
}

func formatErr(err error) error {
	return fmt.Errorf("%s: couldn't access db: %w", "exercise_service", err)
}
