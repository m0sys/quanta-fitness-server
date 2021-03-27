package exercise

import (
	// "errors"
	// "log"

	es "github.com/mhd53/quanta-fitness-server/internal/datastore/exercisestore"
	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	// "github.com/mhd53/quanta-fitness-server/internal/util"
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
	return true, nil
}
