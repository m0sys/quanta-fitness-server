package workout

import (
	"fmt"

	us "github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
	ws "github.com/m0sys/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/m0sys/quanta-fitness-server/internal/util"
)

type WorkoutAuth interface {
	AuthorizeCreateWorkout(uname string) (bool, error)
	AuthorizeAccessWorkout(uname string, wid int64) (bool, error)
}

type authorizer struct {
	ws ws.WorkoutStore
	us us.UserStore
}

func NewWorkoutAuthorizer(workoutstore ws.WorkoutStore,
	userstore us.UserStore) WorkoutAuth {
	return &authorizer{
		ws: workoutstore,
		us: userstore,
	}
}

func (auth *authorizer) AuthorizeCreateWorkout(uname string) (bool, error) {
	return util.CheckUserExists(auth.us, uname)
}

func (auth *authorizer) AuthorizeAccessWorkout(uname string, wid int64) (bool, error) {
	ok, err := util.CheckUserExists(auth.us, uname)

	if err != nil || !ok {
		return ok, err
	}

	return checkUserOwnsWorkout(auth.ws, uname, wid)
}

func checkUserOwnsWorkout(ws ws.WorkoutStore, uname string, wid int64) (bool, error) {
	workoutDS, found, err := ws.FindWorkoutById(wid)
	if err != nil {
		return false, fmt.Errorf("%s: couldn't access db: %w", "workout_auth", err)
	}

	if !found {
		return false, nil
	}

	if workoutDS.Username != uname {
		return false, nil
	}

	return true, nil
}
