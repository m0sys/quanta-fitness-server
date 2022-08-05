package eset

import (
	"fmt"

	esstore "github.com/m0sys/quanta-fitness-server/internal/datastore/esetstore"
	estore "github.com/m0sys/quanta-fitness-server/internal/datastore/exercisestore"
	ustore "github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
	wstore "github.com/m0sys/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/m0sys/quanta-fitness-server/internal/entity"
	serv "github.com/m0sys/quanta-fitness-server/internal/eset"
	eServ "github.com/m0sys/quanta-fitness-server/internal/exercise"
	wServ "github.com/m0sys/quanta-fitness-server/internal/workout"
	"github.com/m0sys/quanta-fitness-server/pkg/format"
)

var (
	service serv.EsetService
)

type server struct{}

type EsetServer interface {
	AddEsetToExercise(uname, eid string, actualRC int, dur, restDur float32) (entity.Eset, error)
	UpdateEset(id, uname string, updates entity.EsetUpdate) (bool, error)
	DeleteEset(id, uname string) (bool, error)
	GetEset(id, uname string) (entity.Eset, error)
	GetEsetsForExercise(eid, uname string) ([]entity.Eset, error)
}

func NewEsetServer(us ustore.UserStore, ess esstore.EsetStore, es estore.ExerciseStore, ws wstore.WorkoutStore) EsetServer {
	wAuthorizer := wServ.NewWorkoutAuthorizer(ws, us)
	eAuthorizer := eServ.NewExerciseAuthorizer(es, us, wAuthorizer)
	authorizer := serv.NewEsetAuthorizer(ess, us, eAuthorizer)
	validator := serv.NewEsetValidator()
	service = serv.NewEsetService(ess, authorizer, validator)
	return &server{}
}

func (*server) AddEsetToExercise(uname, eid string, actualRC int, dur, restDur float32) (entity.Eset, error) {
	intID, err := format.ConvertToBase64(eid)
	if err != nil {
		return entity.Eset{}, formatErr(err)

	}

	return service.AddEsetToExercise(uname, intID, actualRC, dur, restDur)
}

func (*server) UpdateEset(id, uname string, updates entity.EsetUpdate) (bool, error) {
	intID, err := format.ConvertToBase64(id)
	if err != nil {
		return false, formatErr(err)
	}

	err2 := service.UpdateEset(intID, uname, updates)
	if err2 != nil {
		return false, err2
	}

	return true, nil

}

func (*server) DeleteEset(id, uname string) (bool, error) {
	intID, err := format.ConvertToBase64(id)
	if err != nil {
		return false, formatErr(err)
	}

	err2 := service.DeleteEset(intID, uname)
	if err2 != nil {
		return false, err2
	}

	return true, nil

}

func (*server) GetEset(id, uname string) (entity.Eset, error) {
	intID, err := format.ConvertToBase64(id)
	if err != nil {
		return entity.Eset{}, formatErr(err)
	}

	got, err2 := service.GetEset(intID, uname)
	if err2 != nil {
		return entity.Eset{}, err2
	}
	return got, nil

}

func (*server) GetEsetsForExercise(eid, uname string) ([]entity.Eset, error) {
	var esets []entity.Eset

	intID, err := format.ConvertToBase64(eid)
	if err != nil {
		return esets, formatErr(err)
	}

	esets, err2 := service.GetEsetsForExercise(intID, uname)
	if err2 != nil {
		return esets, err2
	}

	return esets, nil
}

func formatErr(err error) error {
	return fmt.Errorf("%s: couldn't format id: %w", "API Eset", err)
}
