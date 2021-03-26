package workout

import (
	"errors"
	"log"

	"github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	"github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	// "github.com/mhd53/quanta-fitness-server/internal/entity"
)

type WorkoutAuth interface {
	AuthorizeCreateWorkout(uname string) (bool, error)
	AuthorizeAccessWorkout(uname string, wid int64) (bool, error)
}

type authorizer struct{}

var (
	aws workoutstore.WorkoutStore
	aus userstore.UserStore
)

func NewWorkoutAuthorizer(workoutstore workoutstore.WorkoutStore,
	userstore userstore.UserStore) WorkoutAuth {
	aws = workoutstore
	aus = userstore
	return &authorizer{}
}

func (*authorizer) AuthorizeCreateWorkout(uname string) (bool, error) {
	return checkUserExists(uname)
}

func checkUserExists(uname string) (bool, error) {
	_, found, err := aus.FindUserByUsername(uname)
	if err != nil {
		log.Fatal(err)
		return false, errors.New("Internal error! Please try again later.")
	}

	if !found {
		return false, nil
	}

	return true, nil

}

func (*authorizer) AuthorizeAccessWorkout(uname string, wid int64) (bool, error) {
	ok, err := checkUserExists(uname)

	if err != nil {
		return ok, err
	}

	if !ok {
		return false, errors.New("Access Denied!")
	}

	return checkUserOwnsWorkout(uname, wid)

}

func checkUserOwnsWorkout(uname string, wid int64) (bool, error) {
	workoutDS, found, err := ws.FindWorkoutById(wid)
	if err != nil {
		log.Fatal(err)
		return false, errors.New("Internal error! Please try again later.")

	}

	if !found {
		return false, nil
	}

	if workoutDS.Username != uname {
		return false, nil
	}

	return true, nil
}
