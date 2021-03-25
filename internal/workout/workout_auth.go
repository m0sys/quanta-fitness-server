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
	AuthorizeUpdateWorkout(uname string, wid int64) (bool, error)
	AuthorizeDeleteWorkout(uname string, wid int64) (bool, error)
	AuthorizeGetWorkoutExercises(uname string, wid int64) (bool, error)
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

func (*authorizer) AuthorizeUpdateWorkout(uname string, wid int64) (bool, error) {
	return true, nil
}

func (*authorizer) AuthorizeDeleteWorkout(uname string, wid int64) (bool, error) {
	return true, nil
}

func (*authorizer) AuthorizeGetWorkoutExercises(uname string, wid int64) (bool, error) {
	return true, nil
}
