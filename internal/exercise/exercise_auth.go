package exercise

import (
	"fmt"

	es "github.com/mhd53/quanta-fitness-server/internal/datastore/exercisestore"
	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	"github.com/mhd53/quanta-fitness-server/internal/util"
	w "github.com/mhd53/quanta-fitness-server/internal/workout"
)

type ExerciseAuth interface {
	AuthorizeReadAccess(uname string) (bool, error)
	AuthorizeWorkoutAccess(uname string, wid int64) (bool, error)
	AuthorizeExerciseAccess(uname string, eid int64) (bool, error)
}

type authorizer struct{}

var (
	aes   es.ExerciseStore
	aus   us.UserStore
	wauth w.WorkoutAuth
)

func NewExerciseAuthorizer(estore es.ExerciseStore,
	ustore us.UserStore,
	wauthorizer w.WorkoutAuth) ExerciseAuth {
	aes = estore
	aus = ustore
	wauth = wauthorizer
	return &authorizer{}
}

func (*authorizer) AuthorizeReadAccess(uname string) (bool, error) {
	return true, nil

}

func (*authorizer) AuthorizeWorkoutAccess(uname string, wid int64) (bool, error) {
	return wauth.AuthorizeAccessWorkout(uname, wid)
}

func (*authorizer) AuthorizeExerciseAccess(uname string, eid int64) (bool, error) {
	ok, err := util.CheckUserExists(aus, uname)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}

	ok2, err2 := checkUserOwnsExercise(uname, eid)

	if err2 != nil {
		return false, err2
	}

	if !ok2 {
		return false, nil
	}

	return true, nil
}

func checkUserOwnsExercise(uname string, eid int64) (bool, error) {
	exerciseDS, found, err := aes.FindExerciseById(eid)
	if err != nil {
		return false, fmt.Errorf("%s: couldn't access db: %w", "exercise_auth", err)

	}

	if !found {
		return false, nil
	}

	if exerciseDS.Username != uname {
		return false, nil
	}
	return true, nil
}
