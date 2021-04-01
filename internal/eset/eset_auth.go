package eset

import (
	"fmt"

	ess "github.com/mhd53/quanta-fitness-server/internal/datastore/esetstore"
	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	e "github.com/mhd53/quanta-fitness-server/internal/exercise"
	"github.com/mhd53/quanta-fitness-server/internal/util"
)

type EsetAuth interface {
	AuthorizeExerciseAccess(uname string, eid string) (bool, error)
	AuthorizeEsetAccess(uname string, esid string) (bool, error)
}

type authorizer struct{}

var (
	aess  ess.EsetStore
	aus   us.UserStore
	eauth e.ExerciseAuth
)

func NewEsetAuthorizer(store ess.EsetStore,
	ustore us.UserStore,
	eauthorizer e.ExerciseAuth) EsetAuth {
	aess = store
	aus = ustore
	eauth = eauthorizer
	return &authorizer{}
}

func (*authorizer) AuthorizeExerciseAccess(uname string, eid string) (bool, error) {
	return eauth.AuthorizeExerciseAccess(uname, eid)
}

func (*authorizer) AuthorizeEsetAccess(uname string, esid string) (bool, error) {
	ok, err := util.CheckUserExists(aus, uname)

	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}

	ok2, err2 := checkUserOwnsEset(uname, esid)

	if err2 != nil {
		return false, err2
	}

	if !ok2 {
		return false, nil
	}

	return true, nil
}

func checkUserOwnsEset(uname string, esid string) (bool, error) {
	esetDS, found, err := aess.FindEsetById(esid)
	if err != nil {
		return false, fmt.Errorf("%s: couldn't access db: %w", "eset_auth", err)
	}

	if !found {
		return false, nil
	}

	if esetDS.Username != uname {
		return false, nil
	}
	return true, nil
}
